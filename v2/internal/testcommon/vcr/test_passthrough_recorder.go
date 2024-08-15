/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package vcr

import (
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
)

// testPassthroughRecorder is an implementation of testRecorder that does not record or replay HTTP requests,
// but which instead just passes them through to a real HTTP endpoint.
type testPassthroughRecorder struct {
	cfg      config.Values
	creds    azcore.TokenCredential
	ids      creds.AzureIDs
	redactor *Redactor
}

var _ Interface = &testPassthroughRecorder{}

// NewTestPassthroughRecorder returns an instance of testRecorder that does not record or replay HTTP requests,
func NewTestPassthroughRecorder(cfg config.Values) (Interface, error) {
	creds, azureIDs, err := creds.GetCreds()
	if err != nil {
		return nil, err
	}

	return &testPassthroughRecorder{
		cfg:      cfg,
		creds:    creds,
		ids:      azureIDs,
		redactor: NewRedactor(azureIDs),
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

// IDs implements testRecorder.
func (r *testPassthroughRecorder) IDs() creds.AzureIDs {
	return r.ids
}

// AddLiteralRedaction adds literal redaction value to redactor
func (r *testPassthroughRecorder) AddLiteralRedaction(redactionValue string, replacementValue string) {
	r.redactor.AddLiteralRedaction(redactionValue, replacementValue)
}

// AddRegexpRedaction adds regular expression redaction value to redactor
func (r *testPassthroughRecorder) AddRegexpRedaction(pattern string, replacementValue string) {
	r.redactor.AddRegexRedaction(pattern, replacementValue)
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
