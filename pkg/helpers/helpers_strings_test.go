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

func TestNewPasswordLengthCheck(t *testing.T) {
	expectedOutput := passwordLength
	testOutput := len(NewPassword())

	if testOutput != expectedOutput {
		t.Errorf("Test output password length %d is not as expected: %d. ", testOutput, expectedOutput)
	}
}

func TestNewPasswordGenerateRule(t *testing.T) {
	// To verify the generated password must contain at least one caracter from
	// upperAlpha, lowerAlpha and number
	expectedOutput1 := 1
	expectedOutput2 := 1
	expectedOutput3 := 1

	testOutput1 := 0
	testOutput2 := 0
	testOutput3 := 0
	testPassword := NewPassword()

	for _, p := range testPassword {

		switch {
		case unicode.IsLower(p):
			testOutput1++
		case unicode.IsUpper(p):
			testOutput2++
		case unicode.IsNumber(p):
			testOutput3++

		}
	}

	if testOutput1 < expectedOutput1 {
		t.Errorf("Test password '%s' doesnot contain any lower alphabet! ", testPassword)
	}

	if testOutput2 < expectedOutput2 {
		t.Errorf("Test password '%s' doesnot contain any upper alphabet! ", testPassword)
	}

	if testOutput3 < expectedOutput3 {
		t.Errorf("Test password '%s' doesnot contain any number! ", testPassword)
	}

}

func TestNewPasswordCheckRandom(t *testing.T) {
	testPassword1 := NewPassword()
	testPassword2 := NewPassword()

	if testPassword1 == testPassword2 {
		t.Errorf("The two random passwords '%s' and '%s' are the same!", testPassword1, testPassword2)
	}
}
