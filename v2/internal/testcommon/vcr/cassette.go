/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package vcr

import (
	"os"

	"github.com/rotisserie/eris"
)

func EnsureCassetteFileExists(cassetteName string) error {
	exists, err := CassetteFileExists(cassetteName)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	filename := CassetteFileName(cassetteName)

	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0o644)
	if err != nil {
		return eris.Wrapf(err, "creating empty cassette %q", filename)
	}
	if err := f.Close(); err != nil {
		return eris.Wrapf(err, "failed to close empty cassette %q", filename)
	}

	return nil
}

func CassetteFileExists(cassetteName string) (bool, error) {
	filename := CassetteFileName(cassetteName)

	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func CassetteFileName(cassetteName string) string {
	return cassetteName + ".yaml"
}
