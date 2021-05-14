/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package zips

import (
	"os"
)

type (
	Enver interface {
		Getenv(key string) string
	}

	stdEnv struct{}
)

// Getenv will return os.Getenv for a given key
func (*stdEnv) Getenv(key string) string {
	return os.Getenv(key)
}
