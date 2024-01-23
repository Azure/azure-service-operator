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
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// recorderDetailsV3 is an implementation of testRecorder using go-vcr v3.
type recorderDetailsV3 struct {
	cassetteName string
	creds        azcore.TokenCredential
	ids          AzureIDs
	recorder     *recorder.Recorder
	cfg          config.Values
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
func (r recorderDetailsV3) Ids() AzureIDs {
	return r.ids
}

// Stop recording
func (r recorderDetailsV3) Stop() error {
	return r.recorder.Stop()
}

// IsReplaying returns true if we're replaying a recorded test, false if we're recording a new test
func (r recorderDetailsV3) IsReplaying() bool {
	return r.recorder.Mode() == recorder.ModeReplaying
}

// CreateRoundTripper creates a client RoundTripper that can be used to record or replay HTTP requests.
// t is a reference to the test currently executing.
// TODO: Remove the reference to t to reduce coupling
func (r recorderDetailsV3) CreateRoundTripper(t *testing.T) http.RoundTripper {
	return addCountHeader(translateErrorsV1(r.recorder, r.cassetteName, t))
}

func createRecorderV3(cassetteName string, cfg config.Values, recordReplay bool) (testRecorder, error) {
	var err error
	var r *recorder.Recorder
	if recordReplay {
		r, err = recorder.New(cassetteName)
	} else {
		r, err = recorder.NewAsMode(cassetteName, recorder.ModeDisabled, nil)
	}

	if err != nil {
		return recorderDetailsV3{}, errors.Wrapf(err, "creating recorder")
	}

	var creds azcore.TokenCredential
	var azureIDs AzureIDs
	if r.Mode() == recorder.ModeRecording ||
		r.Mode() == recorder.ModeDisabled {
		// if we are recording, we need auth
		creds, azureIDs, err = getCreds()
		if err != nil {
			return nil, err
		}
	} else {
		// if we are replaying, we won't need auth
		// and we use a dummy subscription ID/tenant ID
		creds = MockTokenCredential{}
		azureIDs.tenantID = uuid.Nil.String()
		azureIDs.subscriptionID = uuid.Nil.String()
		azureIDs.billingInvoiceID = DummyBillingId
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

	r.AddSaveFilter(func(i *cassette.Interaction) error {
		// rewrite all request/response fields to hide the real subscription ID
		// this is *not* a security measure but intended to make the tests updateable from
		// any subscription, so a contributor can update the tests against their own sub.
		hide := func(s string, id string, replacement string) string {
			return strings.ReplaceAll(s, id, replacement)
		}

		// Note that this changes the cassette in-place so there's no return needed
		hideCassetteString := func(cas *cassette.Interaction, id string, replacement string) {
			i.Request.Body = strings.ReplaceAll(cas.Request.Body, id, replacement)
			i.Response.Body = strings.ReplaceAll(cas.Response.Body, id, replacement)
			i.Request.URL = strings.ReplaceAll(cas.Request.URL, id, replacement)
		}

		// Hide the subscription ID
		hideCassetteString(i, azureIDs.subscriptionID, uuid.Nil.String())
		// Hide the tenant ID
		hideCassetteString(i, azureIDs.tenantID, uuid.Nil.String())
		// Hide the billing ID
		if azureIDs.billingInvoiceID != "" {
			hideCassetteString(i, azureIDs.billingInvoiceID, DummyBillingId)
		}

		// Hiding other sensitive fields
		i.Request.Body = hideRecordingData(i.Request.Body)
		i.Response.Body = hideRecordingData(i.Response.Body)
		i.Request.URL = hideURLData(i.Request.URL)

		for _, values := range i.Request.Headers {
			for i := range values {
				values[i] = hide(values[i], azureIDs.subscriptionID, uuid.Nil.String())
				values[i] = hide(values[i], azureIDs.tenantID, uuid.Nil.String())
				if azureIDs.billingInvoiceID != "" {
					values[i] = hide(values[i], azureIDs.billingInvoiceID, DummyBillingId)
				}
			}
		}

		for key, values := range i.Response.Headers {
			for i := range values {
				values[i] = hide(values[i], azureIDs.subscriptionID, uuid.Nil.String())
				values[i] = hide(values[i], azureIDs.tenantID, uuid.Nil.String())
				if azureIDs.billingInvoiceID != "" {
					values[i] = hide(values[i], azureIDs.billingInvoiceID, DummyBillingId)
				}
			}
			// Hide the base request URL in the AzureOperation and Location headers
			if key == genericarmclient.AsyncOperationHeader || key == genericarmclient.LocationHeader {
				for i := range values {
					values[i] = hideBaseRequestURL(values[i])
				}
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

	return recorderDetailsV3{
		cassetteName: cassetteName,
		creds:        creds,
		ids:          azureIDs,
		recorder:     r,
		cfg:          cfg,
	}, nil
}
