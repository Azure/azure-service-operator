/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
)

var (
	DefaultTestRegion = "westus" // Could make this an env variable if we wanted
)

type TestContext struct {
	AzureRegion  string
	NameConfig   *ResourceNameConfig
	RecordReplay bool
}

type PerTestContext struct {
	TestContext
	T                   *testing.T
	AzureClientRecorder *recorder.Recorder
	AzureClient         armclient.Applier
	AzureSubscription   string
	AzureMatch          *ArmMatcher
	Namer               ResourceNamer
	TestName            string
}

// If you modify this make sure to modify the cleanup-test-azure-resources target in the Makefile too
const ResourcePrefix = "asotest"

func NewTestContext(region string, recordReplay bool) TestContext {
	return TestContext{
		AzureRegion:  region,
		RecordReplay: recordReplay,
		NameConfig:   NewResourceNameConfig(ResourcePrefix, "-", 6),
	}
}

func (tc TestContext) ForTest(t *testing.T) (PerTestContext, error) {
	cassetteName := "recordings/" + t.Name()
	authorizer, subscriptionID, recorder, err := createRecorder(cassetteName, tc.RecordReplay)
	if err != nil {
		return PerTestContext{}, errors.Wrapf(err, "creating recorder")
	}

	armClient, err := armclient.NewAzureTemplateClient(authorizer, subscriptionID)
	if err != nil {
		return PerTestContext{}, errors.Wrapf(err, "creating ARM client")
	}

	// replace the ARM client transport (a bit hacky)
	httpClient := armClient.RawClient.Sender.(*http.Client)
	httpClient.Transport = addCountHeader(translateErrors(recorder, cassetteName))

	t.Cleanup(func() {
		if !t.Failed() {
			log.Printf("saving ARM client recorder")
			err := recorder.Stop()
			if err != nil {
				// cleanup function should not error-out
				log.Printf("unable to stop ARM client recorder: %s", err.Error())
			}
		}
	})

	return PerTestContext{
		TestContext:         tc,
		T:                   t,
		Namer:               tc.NameConfig.NewResourceNamer(t.Name()),
		AzureClient:         armClient,
		AzureSubscription:   subscriptionID,
		AzureMatch:          NewArmMatcher(armClient),
		AzureClientRecorder: recorder,
		TestName:            t.Name(),
	}, nil
}

func createRecorder(cassetteName string, recordReplay bool) (autorest.Authorizer, string, *recorder.Recorder, error) {
	var err error
	var r *recorder.Recorder
	if recordReplay {
		r, err = recorder.New(cassetteName)
	} else {
		r, err = recorder.NewAsMode(cassetteName, recorder.ModeDisabled, nil)
	}

	if err != nil {
		return nil, "", nil, errors.Wrapf(err, "creating recorder")
	}

	var authorizer autorest.Authorizer
	var subscriptionID string
	if r.Mode() == recorder.ModeRecording ||
		r.Mode() == recorder.ModeDisabled {
		// if we are recording, we need auth
		authorizer, subscriptionID, err = getAuthorizer()
		if err != nil {
			return nil, "", nil, err
		}
	} else {
		// if we are replaying, we won't need auth
		// and we use a dummy subscription ID
		subscriptionID = uuid.Nil.String()
		authorizer = nil
	}

	// check body as well as URL/Method (copied from go-vcr documentation)
	r.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		if !cassette.DefaultMatcher(r, i) {
			return false
		}

		// verify custom request count header (see counting_roundtripper.go)
		if r.Header.Get(COUNT_HEADER) != i.Headers.Get(COUNT_HEADER) {
			return false
		}

		if r.Body == nil {
			return i.Body == ""
		}

		var b bytes.Buffer
		if _, err := b.ReadFrom(r.Body); err != nil {
			panic(err)
		}

		r.Body = io.NopCloser(&b)
		return b.String() == "" || hideDates(b.String()) == i.Body
	})

	r.AddSaveFilter(func(i *cassette.Interaction) error {
		// rewrite all request/response fields to hide the real subscription ID
		// this is *not* a security measure but intended to make the tests updateable from
		// any subscription, so a contributer can update the tests against their own sub
		hideSubID := func(s string) string {
			return strings.ReplaceAll(s, subscriptionID, uuid.Nil.String())
		}

		i.Request.Body = hideDates(hideSubID(i.Request.Body))
		i.Response.Body = hideDates(hideSubID(i.Response.Body))
		i.Request.URL = hideSubID(i.Request.URL)

		for _, values := range i.Request.Headers {
			for i := range values {
				values[i] = hideSubID(values[i])
			}
		}

		for _, values := range i.Response.Headers {
			for i := range values {
				values[i] = hideSubID(values[i])
			}
		}

		// remove all Authorization headers from stored requests
		delete(i.Request.Headers, "Authorization")

		// remove all request IDs
		delete(i.Response.Headers, "X-Ms-Correlation-Request-Id")
		delete(i.Response.Headers, "X-Ms-Ratelimit-Remaining-Subscription-Reads")
		delete(i.Response.Headers, "X-Ms-Ratelimit-Remaining-Subscription-Writes")
		delete(i.Response.Headers, "X-Ms-Request-Id")
		delete(i.Response.Headers, "X-Ms-Routing-Request-Id")

		// don't need these headers and they add to diff churn
		delete(i.Request.Headers, "User-Agent")
		delete(i.Response.Headers, "Date")

		return nil
	})

	return authorizer, subscriptionID, r, nil
}

var dateMatcher *regexp.Regexp = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(.\d+)?Z`)

// hideDates replaces all ISO8601 datetimes with a fixed value
// this lets us match requests that may contain time-sensitive information (timestamps, etc)
func hideDates(s string) string {
	return dateMatcher.ReplaceAllLiteralString(s, "2001-02-03T04:05:06Z") // this should be recognizable/parseable as a fake date
}

func (tc PerTestContext) NewTestResourceGroup() *resources.ResourceGroup {
	return &resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: tc.Namer.GenerateName("rg"),
		},
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			Tags:     CreateTestResourceGroupDefaultTags(),
		},
	}
}

func (tc PerTestContext) MakeARMId(resourceGroup string, provider string, params ...string) string {
	if len(params) == 0 {
		panic("At least 2 params must be specified")
	}
	if len(params)%2 != 0 {
		panic("ARM Id params must come in resourceKind/name pairs")
	}

	suffix := strings.Join(params, "/")

	return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/%s/%s", tc.AzureSubscription, resourceGroup, provider, suffix)
}
