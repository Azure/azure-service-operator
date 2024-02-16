/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

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

// recorderDetailsV3 is an implementation of testRecorder using go-vcr v3.
type recorderDetailsV3 struct {
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

func newTestRecorderV3(
	cassetteName string,
	cfg config.Values,
	log logr.Logger,
) (vcr.Interface, error) {
	opts := &recorder.Options{
		CassetteName: cassetteName,
	}

	cassetteExists, err := cassetteFileExists(cassetteName)
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
		return recorderDetailsV3{}, errors.Wrapf(err, "creating recorder")
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
		credentials = MockTokenCredential{}
		azureIDs.TenantID = nilGuid
		azureIDs.SubscriptionID = nilGuid
		azureIDs.BillingInvoiceID = DummyBillingId

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

	r.AddHook(redactRecording(azureIDs), recorder.BeforeSaveHook)

	return recorderDetailsV3{
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
			hideCassetteString(i, azureIDs.BillingInvoiceID, DummyBillingId)
		}

		// Hiding other sensitive fields
		i.Request.Body = hideRecordingData(i.Request.Body)
		i.Response.Body = hideRecordingData(i.Response.Body)
		i.Request.URL = hideURLData(i.Request.URL)

		// Hide sensitive request headers
		for _, values := range i.Request.Headers {
			for i := range values {
				values[i] = hide(values[i], azureIDs.SubscriptionID, nilGuid)
				values[i] = hide(values[i], azureIDs.TenantID, nilGuid)
				if azureIDs.BillingInvoiceID != "" {
					values[i] = hide(values[i], azureIDs.BillingInvoiceID, DummyBillingId)
				}
			}
		}

		// Hide sensitive response headers
		for key, values := range i.Response.Headers {
			for i := range values {
				values[i] = hide(values[i], azureIDs.SubscriptionID, nilGuid)
				values[i] = hide(values[i], azureIDs.TenantID, nilGuid)
				if azureIDs.BillingInvoiceID != "" {
					values[i] = hide(values[i], azureIDs.BillingInvoiceID, DummyBillingId)
				}
			}

			// Hide the base request URL in the AzureOperation and Location headers
			if key == genericarmclient.AsyncOperationHeader || key == genericarmclient.LocationHeader {
				for i := range values {
					values[i] = hideBaseRequestURL(values[i])
				}
			}
		}

		// Remove request headers
		for _, header := range requestHeadersToRemove {
			delete(i.Request.Headers, header)
		}

		// Remove response headers
		for _, header := range responseHeadersToRemove {
			delete(i.Response.Headers, header)
		}

		return nil
	}
}

// Cfg returns the available configuration for the test
func (r recorderDetailsV3) Cfg() config.Values {
	return r.cfg
}

// Creds returns Azure credentials when running for real
func (r recorderDetailsV3) Creds() azcore.TokenCredential {
	return r.creds
}

// Ids returns the available Azure resource IDs for the test
func (r recorderDetailsV3) Ids() creds.AzureIDs {
	return r.ids
}

// Stop recording
func (r recorderDetailsV3) Stop() error {
	return r.recorder.Stop()
}

// IsReplaying returns true if we're replaying a recorded test, false if we're recording a new test
func (r recorderDetailsV3) IsReplaying() bool {
	return r.recorder.Mode() == recorder.ModeReplayOnly
}

// CreateClient creates an HTTP client configured to record or replay HTTP requests.
// t is a reference to the test currently executing.
// TODO: Remove the reference to t to reduce coupling
func (r recorderDetailsV3) CreateClient(t *testing.T) *http.Client {
	withReplay := newReplayRoundTripper(r.recorder, r.log)
	withErrorTranslation := translateErrors(withReplay, r.cassetteName, t)
	withCountHeader := addCountHeader(withErrorTranslation)

	return &http.Client{
		Transport: withCountHeader,
	}
}
