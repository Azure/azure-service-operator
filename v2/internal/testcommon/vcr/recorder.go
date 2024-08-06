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

// Interface is a lightweight interface that allows us to swap out implementations of Go-VCR
// as required.
type Interface interface {
	// Cfg returns the available configuration for the test
	Cfg() config.Values

	// Creds returns Azure credentials when running for real
	Creds() azcore.TokenCredential

	// IDs returns the available Azure resource IDs for the test
	IDs() creds.AzureIDs

	// AddLiteralRedaction adds literal redaction value to redactor
	AddLiteralRedaction(redactionValue, replacementValue string)

	// AddRegexpRedaction adds regular expression redaction value to redactor
	AddRegexpRedaction(pattern, replacementValue string)

	// Stop recording
	Stop() error

	// IsReplaying returns true if we're replaying a recorded test, false if we're recording a new test
	IsReplaying() bool

	// CreateClient creates an HTTP client configured to record or replay HTTP requests.
	// t is a reference to the test currently executing.
	CreateClient(t *testing.T) *http.Client
}
