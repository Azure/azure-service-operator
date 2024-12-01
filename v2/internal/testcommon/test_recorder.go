/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
	v1 "github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr/v1"
	v3 "github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr/v3"
)

// createTestRecorder returns an instance of testRecorder to allow recording and playback of HTTP requests.
func createTestRecorder(
	cassetteName string,
	cfg config.Values,
	recordReplay bool,
	log logr.Logger,
) (vcr.Interface, error) {
	if !recordReplay {
		// We're not using VCR, so just pass through the requests
		return vcr.NewTestPassthroughRecorder(cfg)
	}

	// If a cassette file exists in the old format, use the old player
	v1Exists, err := v1.CassetteFileExists(cassetteName)
	if err != nil {
		return nil, eris.Wrapf(err, "checking existence of cassette %s", cassetteName)
	}

	if v1Exists {
		return v1.NewTestPlayer(cassetteName, cfg)
	}

	return v3.NewTestRecorder(cassetteName, cfg, log)
}
