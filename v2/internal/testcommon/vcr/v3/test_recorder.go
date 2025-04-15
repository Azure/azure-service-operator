/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v3

import (
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
	asocloud "github.com/Azure/azure-service-operator/v2/pkg/common/cloud"
)

// recorderDetails is an implementation of testRecorder using go-vcr v3.
type recorderDetails struct {
	cassetteName string
	creds        azcore.TokenCredential
	ids          creds.AzureIDs
	recorder     *recorder.Recorder
	cfg          config.Values
	log          logr.Logger
	redactor     *vcr.Redactor
}

func NewTestRecorder(
	cassetteName string,
	cfg config.Values,
	log logr.Logger,
) (vcr.Interface, error) {
	opts := &recorder.Options{
		CassetteName: cassetteName,
	}

	cassetteExists, err := vcr.CassetteFileExists(cassetteName)
	if err != nil {
		return nil, eris.Wrapf(err, "checking existence of cassette %s", cassetteName)
	}

	// Work out whether we are recording or replaying
	if cassetteExists {
		opts.Mode = recorder.ModeReplayOnly
	} else {
		opts.Mode = recorder.ModeRecordOnly
	}

	r, err := recorder.NewWithOptions(opts)
	if err != nil {
		return nil, eris.Wrapf(err, "creating recorder")
	}

	var credentials azcore.TokenCredential
	var azureIDs creds.AzureIDs
	if r.Mode() == recorder.ModeRecordOnly {
		// if we are recording, we need auth
		credentials, azureIDs, err = creds.GetCreds()
		if err != nil {
			return nil, err
		}
	} else {
		// if we are replaying, we won't need auth
		// and we use a dummy subscription ID/tenant ID
		credentials = creds.MockTokenCredential{}
		azureIDs = creds.DummyAzureIDs()

		// Force these values to be the default
		cfg.ResourceManagerEndpoint = asocloud.DefaultEndpoint
		cfg.ResourceManagerAudience = asocloud.DefaultAudience
		cfg.AzureAuthorityHost = asocloud.DefaultAADAuthorityHost
	}

	// check body as well as URL/Method (copied from go-vcr documentation)
	r.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		if !cassette.DefaultMatcher(r, i) {
			return false
		}

		// verify custom request count header matches, if present
		if header := r.Header.Get(CountHeader); header != "" {
			interactionHeader := i.Headers.Get(CountHeader)
			if header != interactionHeader {
				log.Info("Request count header mismatch", CountHeader, header, "interaction", interactionHeader)
				return false
			}
		}

		// verify custom body hash header matches, if present
		if header := r.Header.Get(HashHeader); header != "" {
			interactionHeader := i.Headers.Get(HashHeader)
			if header != interactionHeader {
				log.Info("Request body hash header mismatch", HashHeader, header, "interaction", interactionHeader)
				return false
			}
		}

		if r.Body == nil {
			// Empty bodies are stored by go-vcr as empty strings
			return i.Body == ""
		}

		return true
	})

	redactor := vcr.NewRedactor(azureIDs)
	r.AddHook(redactRecording(redactor), recorder.BeforeSaveHook)

	return &recorderDetails{
		cassetteName: cassetteName,
		creds:        credentials,
		ids:          azureIDs,
		recorder:     r,
		cfg:          cfg,
		log:          log,
		redactor:     redactor,
	}, nil
}

// redactRecording is a BeforeSaveHook that rewrites data in the cassette
// This includes hiding the SubscriptionID, TenantID, and BillingInvoiceID. This is
// to make the tests updatable from
// any subscription, so a contributor can update the tests against their own sub.
func redactRecording(
	redactor *vcr.Redactor,
) recorder.HookFunc {
	return func(i *cassette.Interaction) error {
		i.Request.Body = redactor.HideRecordingData(i.Request.Body)
		i.Response.Body = redactor.HideRecordingData(i.Response.Body)
		i.Request.URL = redactor.HideURLData(i.Request.URL)

		redactor.RedactRequestHeaders(i.Request.Headers)
		redactor.RedactResponseHeaders(i.Response.Headers)

		return nil
	}
}

// Cfg returns the available configuration for the test
func (r *recorderDetails) Cfg() config.Values {
	return r.cfg
}

// Creds returns Azure credentials when running for real
func (r *recorderDetails) Creds() azcore.TokenCredential {
	return r.creds
}

// IDs returns the available Azure resource IDs for the test
func (r *recorderDetails) IDs() creds.AzureIDs {
	return r.ids
}

// AddLiteralRedaction adds literal redaction value to redactor
func (r *recorderDetails) AddLiteralRedaction(redactionValue string, replacementValue string) {
	r.redactor.AddLiteralRedaction(redactionValue, replacementValue)
}

// AddRegexpRedaction adds regular expression redaction value to redactor
func (r *recorderDetails) AddRegexpRedaction(pattern string, replacementValue string) {
	r.redactor.AddRegexRedaction(pattern, replacementValue)
}

// Stop recording
func (r *recorderDetails) Stop() error {
	return r.recorder.Stop()
}

// IsReplaying returns true if we're replaying a recorded test, false if we're recording a new test
func (r *recorderDetails) IsReplaying() bool {
	return r.recorder.Mode() == recorder.ModeReplayOnly
}

// CreateClient creates an HTTP client configured to record or replay HTTP requests.
// t is a reference to the test currently executing.
func (r *recorderDetails) CreateClient(t *testing.T) *http.Client {
	withReplay := NewReplayRoundTripper(r.recorder, r.log, r.redactor)
	withErrorTranslation := translateErrors(withReplay, r.cassetteName, r.redactor, t)
	withTrackingHeaders := AddTrackingHeaders(withErrorTranslation, r.redactor)

	return &http.Client{
		Transport: withTrackingHeaders,
	}
}
