/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
	asocloud "github.com/Azure/azure-service-operator/v2/pkg/common/cloud"
)

// recorderDetails is an implementation of testRecorder using go-vcr v4.
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
	cassetteExists, err := vcr.CassetteFileExists(cassetteName)
	if err != nil {
		return nil, eris.Wrapf(err, "checking existence of cassette %s", cassetteName)
	}

	// Work out whether we are recording or replaying
	var mode recorder.Mode
	if cassetteExists {
		mode = recorder.ModeReplayOnly
	} else {
		mode = recorder.ModeRecordOnly
	}

	var credentials azcore.TokenCredential
	var azureIDs creds.AzureIDs
	if mode == recorder.ModeRecordOnly {
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

	redactor := vcr.NewRedactor(azureIDs)

	barrier := newBarrierMatcher(matchOnHeadersAndBody(log), log)
	r, err := recorder.New(
		cassetteName,
		recorder.WithMode(mode),
		recorder.WithMatcher(barrier.Match),
		recorder.WithHook(redactRecording(redactor), recorder.BeforeSaveHook),
	)
	if err != nil {
		return nil, eris.Wrapf(err, "creating recorder")
	}

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

// matchOnHeadersAndBody returns a MatcherFunc to identify which interaction matches an incoming request.
// In go-vcr v3, the default matching logic only considers the HTTP method and URL.
// But in go-vcr v4, this extended to headers and body as well, breaking our functionality (since the incoming request)
// has an unredacted body, but the interaction has a redacted one - they will never match.
// So here we do the entire match ourselves.
func matchOnHeadersAndBody(log logr.Logger) recorder.MatcherFunc {
	return func(r *http.Request, i cassette.Request) bool {
		// Require method to match
		if r.Method != i.Method {
			return false
		}

		// Require URL to match (including query parameters)
		if r.URL.String() != i.URL {
			return false
		}

		// Verify body hash headers match, if present.
		// We check both the raw hash and the canonical hash. A match on either is sufficient.
		// The raw hash provides backward compatibility with existing recordings.
		// The canonical hash handles non-deterministic JSON key ordering from Go map serialization.
		if header := r.Header.Get(HashHeader); header != "" {
			rawMatch := header == i.Headers.Get(HashHeader)
			canonicalMatch := r.Header.Get(CanonicalHashHeader) == i.Headers.Get(CanonicalHashHeader)

			if !rawMatch && !canonicalMatch {
				log.Info(
					"Request body hash header mismatch",
					HashHeader, header,
					"interaction", i.Headers.Get(HashHeader),
					CanonicalHashHeader, r.Header.Get(CanonicalHashHeader),
					"interactionCanonical", i.Headers.Get(CanonicalHashHeader),
				)
				return false
			}
		}

		if r.Body == nil {
			// Empty bodies are stored by go-vcr as empty strings
			return i.Body == ""
		}

		return true
	}
}
