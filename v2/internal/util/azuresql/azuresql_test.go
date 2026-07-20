/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package azuresql

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestFindBadChars(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "clean string passes",
			input:       "validPassword123!@#",
			expectError: false,
		},
		{
			name:        "single quote fails",
			input:       "pass'word",
			expectError: true,
		},
		{
			name:        "double quote fails",
			input:       "pass\"word",
			expectError: true,
		},
		{
			name:        "semicolon fails",
			input:       "pass;word",
			expectError: true,
		},
		{
			name:        "double dash fails",
			input:       "pass--word",
			expectError: true,
		},
		{
			name:        "block comment fails",
			input:       "pass/*word",
			expectError: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			err := findBadChars(c.input)
			if c.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}
		})
	}
}

func TestEscapeStringContent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no special chars",
			input:    "simplepassword",
			expected: "simplepassword",
		},
		{
			name:     "single quote is doubled",
			input:    "pass'word",
			expected: "pass''word",
		},
		{
			name:     "semicolons are preserved",
			input:    "pass;word",
			expected: "pass;word",
		},
		{
			name:     "double dashes are preserved",
			input:    "pass--word",
			expected: "pass--word",
		},
		{
			name:     "complex special chars",
			input:    "p@ss;w'rd--/*test",
			expected: "p@ss;w''rd--/*test",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := escapeStringContent(c.input)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}

func TestEscapeIdentifierContent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no special chars",
			input:    "username",
			expected: "username",
		},
		{
			name:     "double quote is doubled",
			input:    "user\"name",
			expected: "user\"\"name",
		},
		{
			name:     "other special chars are preserved",
			input:    "user;name",
			expected: "user;name",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := escapeIdentifierContent(c.input)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}
