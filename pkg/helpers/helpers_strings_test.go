// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import (
	"testing"
	"unicode"
)

// TestGenerateRandomUsername tests whether the username generated will ever start with a non letter
func TestGenerateRandomUsername(t *testing.T) {
	for i := 0; i < 1000; i++ {
		u := GenerateRandomUsername(10)
		if !unicode.IsLetter(rune(u[0])) {
			t.Errorf("index 0 of username '%s' is not a letter '%s'", u, string(u[0]))
		}
	}
}
