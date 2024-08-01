/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"

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
	hideCustomData map[string]string) (vcr.Interface, error) {
	if !recordReplay {
		// We're not using VCR, so just pass through the requests
		return vcr.NewTestPassthroughRecorder(cfg)
	}

	// If a cassette file exists in the old format, use the old player
	v1Exists, err := v1.CassetteFileExists(cassetteName)
	if err != nil {
		return nil, errors.Wrapf(err, "checking existence of cassette %s", cassetteName)
	}

	if v1Exists {
		return v1.NewTestPlayer(cassetteName, cfg, hideCustomData)
	}

	return v3.NewTestRecorder(cassetteName, cfg, log, hideCustomData)
}
