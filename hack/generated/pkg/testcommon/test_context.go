/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
	resources "github.com/Azure/azure-service-operator/v2/api/microsoft.resources/v1alpha1api20200601"
)

// Use WestUS2 as some things (such as VM quota) are hard to get in West US.
var DefaultTestRegion = "westus2" // Could make this an env variable if we wanted

type TestContext struct {
	AzureRegion  string
	NameConfig   *ResourceNameConfig
	RecordReplay bool
}

type PerTestContext struct {
	TestContext
	T                   *testing.T
	logger              logr.Logger
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
	logger := NewTestLogger(t)

	cassetteName := "recordings/" + t.Name()
	authorizer, subscriptionID, recorder, err := createRecorder(cassetteName, tc.RecordReplay)
	if err != nil {
		return PerTestContext{}, errors.Wrapf(err, "creating recorder")
	}

	armClient, err := armclient.NewAzureTemplateClient(authorizer, subscriptionID)
	if err != nil {
		return PerTestContext{}, errors.Wrapf(err, "creating ARM client")
	}

	// As of https://github.com/Azure/go-autorest/releases/tag/autorest%2Fv0.11.9, the autorest
	// client reuses HTTP clients among instances. We add handlers to the HTTP client based on
	// the specific test in question, which means that clients cannot be reused.
	// We explicitly create a new http.Client so that the recording from one test doesn't
	// get used for all other parallel tests.
	httpClient := &http.Client{
		Jar:       armClient.RawClient.Jar,
		Transport: armClient.RawClient.Sender.(*http.Client).Transport,
	}

	// replace the ARM client transport (a bit hacky)
	httpClient.Transport = addCountHeader(translateErrors(recorder, cassetteName, t))
	armClient.RawClient.Sender = httpClient

	t.Cleanup(func() {
		if !t.Failed() {
			logger.Info("saving ARM client recorder")
			err := recorder.Stop()
			if err != nil {
				// cleanup function should not error-out
				logger.Error(err, "unable to stop ARM client recorder")
				t.Fail()
			}
		}
	})

	return PerTestContext{
		TestContext:         tc,
		T:                   t,
		logger:              logger,
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
		return b.String() == "" || hideRecordingData(b.String()) == i.Body
	})

	r.AddSaveFilter(func(i *cassette.Interaction) error {
		// rewrite all request/response fields to hide the real subscription ID
		// this is *not* a security measure but intended to make the tests updateable from
		// any subscription, so a contributer can update the tests against their own sub
		hideSubID := func(s string) string {
			return strings.ReplaceAll(s, subscriptionID, uuid.Nil.String())
		}

		i.Request.Body = hideRecordingData(hideSubID(i.Request.Body))
		i.Response.Body = hideRecordingData(hideSubID(i.Response.Body))
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

		for _, header := range requestHeadersToRemove {
			delete(i.Request.Headers, header)
		}

		for _, header := range responseHeadersToRemove {
			delete(i.Response.Headers, header)
		}

		return nil
	})

	return authorizer, subscriptionID, r, nil
}

var requestHeadersToRemove = []string{
	// remove all Authorization headers from stored requests
	"Authorization",

	// Not needed, adds to diff churn:
	"User-Agent",
}

var responseHeadersToRemove = []string{
	// Request IDs
	"X-Ms-Arm-Service-Request-Id",
	"X-Ms-Correlation-Request-Id",
	"X-Ms-Request-Id",
	"X-Ms-Routing-Request-Id",

	// Quota limits
	"X-Ms-Ratelimit-Remaining-Subscription-Deletes",
	"X-Ms-Ratelimit-Remaining-Subscription-Reads",
	"X-Ms-Ratelimit-Remaining-Subscription-Writes",

	// Not needed, adds to diff churn
	"Date",
}

var (
	dateMatcher     = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(.\d+)?Z`)
	sshKeyMatcher   = regexp.MustCompile("ssh-rsa [0-9a-zA-Z+/=]+")
	passwordMatcher = regexp.MustCompile("pass.*?pass")
)

// hideDates replaces all ISO8601 datetimes with a fixed value
// this lets us match requests that may contain time-sensitive information (timestamps, etc)
func hideDates(s string) string {
	return dateMatcher.ReplaceAllLiteralString(s, "2001-02-03T04:05:06Z") // this should be recognizable/parseable as a fake date
}

// hideSSHKeys hides anything that looks like SSH keys
func hideSSHKeys(s string) string {
	return sshKeyMatcher.ReplaceAllLiteralString(s, "ssh-rsa {KEY}")
}

// hidePasswords hides anything that looks like a generated password
func hidePasswords(s string) string {
	return passwordMatcher.ReplaceAllLiteralString(s, "{PASSWORD}")
}

func hideRecordingData(s string) string {
	result := hideDates(s)
	result = hideSSHKeys(result)
	result = hidePasswords(result)

	return result
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

// GenerateSSHKey generates an SSH key.
func (tc PerTestContext) GenerateSSHKey(size int) (*string, error) {
	// Note: If we ever want to make sure that the SSH keys are the same between
	// test runs, we can base it off of a hash of subscription ID. Right now since
	// we just replace the SSH key in the recordings regardless of what the value is
	// there's no need for uniformity between runs though.

	key, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, err
	}

	err = key.Validate()
	if err != nil {
		return nil, err
	}

	sshPublicKey, err := ssh.NewPublicKey(&key.PublicKey)
	if err != nil {
		return nil, err
	}

	resultBytes := ssh.MarshalAuthorizedKey(sshPublicKey)
	result := string(resultBytes)

	return &result, nil
}

func (tc PerTestContext) MakeARMId(resourceGroup string, provider string, params ...string) string {
	return armclient.MakeResourceGroupScopeARMID(tc.AzureSubscription, resourceGroup, provider, params...)
}
