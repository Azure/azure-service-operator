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

// testRecorder is a lightweight interface that allows us to swap out implementations of Go-VCR
// as required.
type testRecorder interface {
	// Cfg returns the available configuration for the test
	Cfg() config.Values

	// Creds returns Azure credentials when running for real
	Creds() azcore.TokenCredential

	// Ids returns the available Azure resource IDs for the test
	Ids() AzureIDs

	// Stop recording
	Stop() error

	// IsReplaying returns true if we're replaying a recorded test, false if we're recording a new test
	IsReplaying() bool

	// CreateClient creates an HTTP client configured to record or replay HTTP requests.
	// t is a reference to the test currently executing.
	CreateClient(t *testing.T) *http.Client
}
