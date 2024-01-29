/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

// playerDetailsV1 is an implementation of testRecorder using go-vcr v1 that can only play back
// test recordings, not record them.
type playerDetailsV1 struct {
	cassetteName string
	creds        azcore.TokenCredential
	ids          AzureIDs
	recorder     *recorder.Recorder
	cfg          config.Values
}

// Verify we implement testRecorder
var _ testRecorder = &playerDetailsV1{}

// newTestPlayerV1 creates a TestRecorder that can be used to replay test recorded with go-vcr v1.
// cassetteName is the name of the cassette file to replay.
// cfg is the configuration to use when replaying the test.
func newTestPlayerV1(
	cassetteName string, 
	cfg config.Values,
) (testRecorder, error) {
	cassetteExists, err := cassetteFileExists(cassetteName)
	if err != nil {
		return nil, errors.Wrapf(err, "checking for cassette file")
	}
	if !cassetteExists {
		return nil, errors.Errorf("cassette %s does not exist", cassetteName)
	}

	r, err := recorder.NewAsMode(cassetteName, recorder.ModeReplaying, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "creating player")
	}

	var creds azcore.TokenCredential
	var azureIDs AzureIDs

	// if We are replaying, we won't need auth
	// and we use a dummy subscription ID/tenant ID
	creds = MockTokenCredential{}
	azureIDs.tenantID = uuid.Nil.String()
	azureIDs.subscriptionID = uuid.Nil.String()
	azureIDs.billingInvoiceID = DummyBillingId

	// Force these values to be the default
	cfg.ResourceManagerEndpoint = config.DefaultEndpoint
	cfg.ResourceManagerAudience = config.DefaultAudience
	cfg.AzureAuthorityHost = config.DefaultAADAuthorityHost

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

	return playerDetailsV1{
		cassetteName: cassetteName,
		creds:        creds,
		ids:          azureIDs,
		recorder:     r,
		cfg:          cfg,
	}, nil
}

// Cfg returns the available configuration for the test
func (r playerDetailsV1) Cfg() config.Values {
	return r.cfg
}

// Creds returns Azure credentials when running for real
func (r playerDetailsV1) Creds() azcore.TokenCredential {
	return r.creds
}

// Ids returns the available Azure resource IDs for the test
func (r playerDetailsV1) Ids() AzureIDs {
	return r.ids
}

// Stop recording
func (r playerDetailsV1) Stop() error {
	return r.recorder.Stop()
}

// IsReplaying returns true if we're replaying a recorded test, false if we're recording a new test
func (r playerDetailsV1) IsReplaying() bool {
	return r.recorder.Mode() == recorder.ModeReplaying
}

// CreateRoundTripper creates a client RoundTripper that can be used to record or replay HTTP requests.
// t is a reference to the test currently executing.
// TODO: Remove the reference to t to reduce coupling
func (r playerDetailsV1) CreateRoundTripper(t *testing.T) http.RoundTripper {
	return addCountHeader(translateErrorsV1(r.recorder, r.cassetteName, t))
}
