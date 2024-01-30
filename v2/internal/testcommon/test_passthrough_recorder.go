/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-service-operator/v2/internal/config"
)

// testPassthroughRecorder is an implementation of testRecorder that does not record or replay HTTP requests,
// but which instead just passes them through to a real HTTP endpoint.
type testPassthroughRecorder struct {
	cfg   config.Values
	creds azcore.TokenCredential
	ids   AzureIDs
}

var _ testRecorder = &testPassthroughRecorder{}

// newTestPassthroughRecorder returns an instance of testRecorder that does not record or replay HTTP requests,
func newTestPassthroughRecorder(cfg config.Values) (testRecorder, error) {
	creds, azureIDs, err := getCreds()
	if err != nil {
		return nil, err
	}

	return &testPassthroughRecorder{
		cfg:   cfg,
		creds: creds,
		ids:   azureIDs,
	}, nil
}

// Cfg implements testRecorder.
func (r *testPassthroughRecorder) Cfg() config.Values {
	return r.cfg
}

// CreateClient implements testRecorder.
func (*testPassthroughRecorder) CreateClient(t *testing.T) *http.Client {
	return http.DefaultClient
}

// Creds implements testRecorder.
func (r *testPassthroughRecorder) Creds() azcore.TokenCredential {
	return r.creds
}

// Ids implements testRecorder.
func (r *testPassthroughRecorder) Ids() AzureIDs {
	return r.ids
}

// IsReplaying implements testRecorder.
func (*testPassthroughRecorder) IsReplaying() bool {
	return false
}

// Stop implements testRecorder.
func (*testPassthroughRecorder) Stop() error {
	// Nothing to do
	return nil
}
