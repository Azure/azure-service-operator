/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"os"

	"github.com/rotisserie/eris"
	"gopkg.in/yaml.v2"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

// CassetteFileExists returns true if a cassette file exists AND it contains a go-vcr V1 recording
func CassetteFileExists(cassetteName string) (bool, error) {
	exists, err := vcr.CassetteFileExists(cassetteName)
	if err != nil {
		return false, eris.Wrapf(err, "checking whether v1 cassette exists")
	}
	if !exists {
		return false, nil
	}

	filename := vcr.CassetteFileName(cassetteName)
	var content struct {
		Version int `json:"version"`
	}

	file, err := os.Open(filename)
	if err != nil {
		return false, eris.Wrapf(err, "opening cassette file %q", filename)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&content)
	if err != nil {
		return false, eris.Wrapf(err, "parsing cassette file %q", filename)
	}

	return content.Version == 1, nil
}
