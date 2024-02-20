/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v3

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

// recorderDetails is an implementation of testRecorder using go-vcr v3.
type recorderDetails struct {
	cassetteName string
	creds        azcore.TokenCredential
	ids          creds.AzureIDs
	recorder     *recorder.Recorder
	cfg          config.Values
	log          logr.Logger
}

var (
	nilGuid = uuid.Nil.String()
)

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
		azureIDs.TenantID = nilGuid
		azureIDs.SubscriptionID = nilGuid
		azureIDs.BillingInvoiceID = creds.DummyBillingId

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

		// verify custom request count header (see counting_roundtripper.go)
		if r.Header.Get(vcr.COUNT_HEADER) != i.Headers.Get(vcr.COUNT_HEADER) {
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
		return b.String() == "" || vcr.HideRecordingData(b.String()) == i.Body
	})

	r.AddHook(redactRecording(azureIDs), recorder.BeforeSaveHook)

	return &recorderDetails{
		cassetteName: cassetteName,
		creds:        credentials,
		ids:          azureIDs,
		recorder:     r,
		cfg:          cfg,
		log:          log,
	}, nil
}

// redactRecording is a BeforeSaveHook that rewrites data in the cassette
// This incldues hiding the SubscriptionID, TenantID, and BillingInvoiceID, but is not
// a a security measure but intended to make the tests updateable from
// any subscription, so a contributor can update the tests against their own sub.
func redactRecording(
	azureIDs creds.AzureIDs,
) recorder.HookFunc {
	hide := func(s string, id string, replacement string) string {
		return strings.ReplaceAll(s, id, replacement)
	}

	return func(i *cassette.Interaction) error {

		// Note that this changes the cassette in-place so there's no return needed
		hideCassetteString := func(cas *cassette.Interaction, id string, replacement string) {
			i.Request.Body = hide(cas.Request.Body, id, replacement)
			i.Response.Body = hide(cas.Response.Body, id, replacement)
			i.Request.URL = hide(cas.Request.URL, id, replacement)
		}

		// Hide the subscription ID
		hideCassetteString(i, azureIDs.SubscriptionID, nilGuid)

		// Hide the tenant ID
		hideCassetteString(i, azureIDs.TenantID, nilGuid)

		// Hide the billing ID
		if azureIDs.BillingInvoiceID != "" {
			hideCassetteString(i, azureIDs.BillingInvoiceID, creds.DummyBillingId)
		}

		// Hiding other sensitive fields
		i.Request.Body = vcr.HideRecordingData(i.Request.Body)
		i.Response.Body = vcr.HideRecordingData(i.Response.Body)
		i.Request.URL = vcr.HideURLData(i.Request.URL)

		// Hide sensitive request headers
		for _, values := range i.Request.Headers {
			for i := range values {
				values[i] = hide(values[i], azureIDs.SubscriptionID, nilGuid)
				values[i] = hide(values[i], azureIDs.TenantID, nilGuid)
				if azureIDs.BillingInvoiceID != "" {
					values[i] = hide(values[i], azureIDs.BillingInvoiceID, creds.DummyBillingId)
				}
			}
		}

		// Hide sensitive response headers
		for key, values := range i.Response.Headers {
			for i := range values {
				values[i] = hide(values[i], azureIDs.SubscriptionID, nilGuid)
				values[i] = hide(values[i], azureIDs.TenantID, nilGuid)
				if azureIDs.BillingInvoiceID != "" {
					values[i] = hide(values[i], azureIDs.BillingInvoiceID, creds.DummyBillingId)
				}
			}

			// Hide the base request URL in the AzureOperation and Location headers
			if key == genericarmclient.AsyncOperationHeader || key == genericarmclient.LocationHeader {
				for i := range values {
					values[i] = vcr.HideBaseRequestURL(values[i])
				}
			}
		}

		vcr.RedactRequestHeaders(i.Request.Headers)
		vcr.RedactResponseHeaders(i.Response.Headers)

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
	withErrorTranslation := translateErrors(withReplay, r.cassetteName, t)
	withCountHeader := vcr.AddCountHeader(withErrorTranslation)

	return &http.Client{
		Transport: withCountHeader,
	}
}
