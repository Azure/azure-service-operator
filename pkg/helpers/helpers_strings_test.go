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

func TestMakeResourceIDWithSubResource(t *testing.T) {
	testOutput := MakeResourceID(
		"00000000-0000-0000-0000-000000000000",
		"test",
		"Microsoft.Network",
		"networkInterfaces",
		"test",
		"subnets",
		"test",
	)
	expectedOutput := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test/providers/Microsoft.Network/networkInterfaces/test/subnets/test"

	if testOutput != expectedOutput {
		t.Errorf("Test output string '%s' is not as expected: '%s'.", testOutput, expectedOutput)
	}
}

func TestMakeResourceIDWithNoSubResource(t *testing.T) {
	testOutput := MakeResourceID(
		"00000000-0000-0000-0000-000000000000",
		"test",
		"Microsoft.Network",
		"networkInterfaces",
		"test",
		"",
		"",
	)
	expectedOutput := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test/providers/Microsoft.Network/networkInterfaces/test"

	if testOutput != expectedOutput {
		t.Errorf("Test output string '%s' is not as expected: '%s'.", testOutput, expectedOutput)
	}
}

func TestDecodingFromBase64EncodedString(t *testing.T) {
	testOutput1 := FromBase64EncodedString("dGVzdA==")
	expectedOutput1 := "test"

	if testOutput1 != expectedOutput1 {
		t.Errorf("Test output string '%s' is not as expected: '%s'.", testOutput1, expectedOutput1)
	}

	testOutput2 := FromBase64EncodedString("")
	expectedOutput2 := ""

	if testOutput2 != expectedOutput2 {
		t.Errorf("Test output string '%s' is not as expected: '%s'.", testOutput2, expectedOutput2)
	}
}
