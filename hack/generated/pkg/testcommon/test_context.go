/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"log"
	"net/http"
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/go-autorest/autorest"
	resources "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.resources/v20200601"
	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
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
	AzureMatch          *ArmMatcher
	Namer               ResourceNamer
	TestName            string
}

// If you modify this make sure to modify the cleanup-test-azure-resources target in the Makefile too
const ResourcePrefix = "k8sinfratest"

func NewTestContext(region string, recordReplay bool) TestContext {
	return TestContext{
		AzureRegion:  region,
		RecordReplay: recordReplay,
		NameConfig:   NewResourceNameConfig(ResourcePrefix, "-", 6),
	}
}

func (tc TestContext) ForTest(t *testing.T) (PerTestContext, error) {
	authorizer, subscriptionID, recorder, err := createRecorder(t.Name(), tc.RecordReplay)
	if err != nil {
		return PerTestContext{}, errors.Wrapf(err, "creating recorder")
	}

	armClient, err := armclient.NewAzureTemplateClient(authorizer, subscriptionID)
	if err != nil {
		return PerTestContext{}, errors.Wrapf(err, "creating ARM client")
	}

	// replace the ARM client transport (a bit hacky)
	httpClient := armClient.RawClient.Sender.(*http.Client)
	httpClient.Transport = recorder

	t.Cleanup(func() {
		log.Printf("stopping ARM client recorder")
		err := recorder.Stop()
		if err != nil {
			// cleanup function should not error-out
			log.Printf("unable to stop ARM client recorder: %s", err.Error())
		}
	})

	return PerTestContext{
		TestContext:         tc,
		T:                   t,
		Namer:               tc.NameConfig.NewResourceNamer(t.Name()),
		AzureClient:         armClient,
		AzureMatch:          NewArmMatcher(armClient),
		AzureClientRecorder: recorder,
		TestName:            t.Name(),
	}, nil
}

func createRecorder(testName string, recordReplay bool) (autorest.Authorizer, string, *recorder.Recorder, error) {
	cassetteName := "recordings/" + testName

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

	r.AddSaveFilter(func(i *cassette.Interaction) error {
		// rewrite all request/response fields to hide the real subscription ID
		// this is *not* a security measure but intended to make the tests updateable from
		// any subscription, so a contributer can update the tests against their own sub
		hideSubID := func(s string) string {
			return strings.ReplaceAll(s, subscriptionID, uuid.Nil.String())
		}

		i.Request.Body = hideSubID(i.Request.Body)
		i.Response.Body = hideSubID(i.Response.Body)
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

		return nil
	})

	// request must match URI & METHOD & our custom header
	r.SetMatcher(func(request *http.Request, i cassette.Request) bool {
		return cassette.DefaultMatcher(request, i) &&
			request.Header.Get(COUNT_HEADER) == i.Headers.Get(COUNT_HEADER)
	})

	return authorizer, subscriptionID, r, nil
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
