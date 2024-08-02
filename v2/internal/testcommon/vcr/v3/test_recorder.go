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
	"github.com/pkg/errors"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

// recorderDetails is an implementation of testRecorder using go-vcr v3.
type recorderDetails struct {
	cassetteName    string
	creds           azcore.TokenCredential
	ids             creds.AzureIDs
	recorder        *recorder.Recorder
	cfg             config.Values
	log             logr.Logger
	customRedactMap map[string]string
}

func NewTestRecorder(
	cassetteName string,
	cfg config.Values,
	log logr.Logger,
	customRedactMap map[string]string,
) (vcr.Interface, error) {
	opts := &recorder.Options{
		CassetteName: cassetteName,
	}

	cassetteExists, err := vcr.CassetteFileExists(cassetteName)
	if err != nil {
		return nil, errors.Wrapf(err, "checking existence of cassette %s", cassetteName)
	}

	// Work out whether we are recording or replaying
	if cassetteExists {
		opts.Mode = recorder.ModeReplayOnly
	} else {
		opts.Mode = recorder.ModeRecordOnly
	}

	r, err := recorder.NewWithOptions(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "creating recorder")
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
		cfg.ResourceManagerEndpoint = config.DefaultEndpoint
		cfg.ResourceManagerAudience = config.DefaultAudience
		cfg.AzureAuthorityHost = config.DefaultAADAuthorityHost
	}

	// check body as well as URL/Method (copied from go-vcr documentation)
	r.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		if !cassette.DefaultMatcher(r, i) {
			return false
		}

		// verify custom request count header matches, if present
		if header := r.Header.Get(COUNT_HEADER); header != "" {
			interactionHeader := i.Headers.Get(COUNT_HEADER)
			if header != interactionHeader {
				log.Info("Request count header mismatch", COUNT_HEADER, header, "interaction", interactionHeader)
				return false
			}
		}

		// verify custom body hash header matches, if present
		if header := r.Header.Get(HASH_HEADER); header != "" {
			interactionHeader := i.Headers.Get(HASH_HEADER)
			if header != interactionHeader {
				log.Info("Request body hash header mismatch", HASH_HEADER, header, "interaction", interactionHeader)
				return false
			}
		}

		if r.Body == nil {
			// Empty bodies are stored by go-vcr as empty strings
			return i.Body == ""
		}

		return true
	})

	r.AddHook(redactRecording(azureIDs, customRedactMap), recorder.BeforeSaveHook)

	return &recorderDetails{
		cassetteName:    cassetteName,
		creds:           credentials,
		ids:             azureIDs,
		recorder:        r,
		cfg:             cfg,
		log:             log,
		customRedactMap: customRedactMap,
	}, nil
}

// redactRecording is a BeforeSaveHook that rewrites data in the cassette
// This includes hiding the SubscriptionID, TenantID, and BillingInvoiceID. This is
// to make the tests updatable from
// any subscription, so a contributor can update the tests against their own sub.
func redactRecording(
	azureIDs creds.AzureIDs,
	customRedactMap map[string]string,
) recorder.HookFunc {
	return func(i *cassette.Interaction) error {
		i.Request.Body = vcr.HideRecordingDataWithCustomRedaction(azureIDs, i.Request.Body, customRedactMap)
		i.Response.Body = vcr.HideRecordingDataWithCustomRedaction(azureIDs, i.Response.Body, customRedactMap)
		i.Request.URL = vcr.HideURLData(azureIDs, i.Request.URL)

		vcr.RedactRequestHeaders(azureIDs, i.Request.Headers)
		vcr.RedactResponseHeaders(azureIDs, i.Response.Headers)

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
	withReplay := NewReplayRoundTripper(r.recorder, r.log)
	withErrorTranslation := translateErrors(withReplay, r.cassetteName, r.customRedactMap, t)
	withTrackingHeaders := AddTrackingHeaders(r.ids, withErrorTranslation, r.customRedactMap)

	return &http.Client{
		Transport: withTrackingHeaders,
	}
}
