// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package postgresql

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

			err := FindBadChars(c.input)
			if c.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}
		})
	}
}

func TestEscapeStringLiteral(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no special chars",
			input:    "simplepassword",
			expected: "'simplepassword'",
		},
		{
			name:     "single quote is doubled",
			input:    "pass'word",
			expected: "'pass''word'",
		},
		{
			name:     "multiple single quotes",
			input:    "it's a test's password",
			expected: "'it''s a test''s password'",
		},
		{
			name:     "semicolons are preserved",
			input:    "pass;word",
			expected: "'pass;word'",
		},
		{
			name:     "double dashes are preserved",
			input:    "pass--word",
			expected: "'pass--word'",
		},
		{
			name:     "block comment is preserved",
			input:    "pass/*word",
			expected: "'pass/*word'",
		},
		{
			name:     "double quotes are preserved",
			input:    "pass\"word",
			expected: "'pass\"word'",
		},
		{
			name:     "complex special chars",
			input:    "p@ss;w'rd--/*test",
			expected: "'p@ss;w''rd--/*test'",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "''",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := EscapeStringLiteral(c.input)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}

func TestEscapeIdentifier(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no special chars",
			input:    "username",
			expected: "\"username\"",
		},
		{
			name:     "double quote is doubled",
			input:    "user\"name",
			expected: "\"user\"\"name\"",
		},
		{
			name:     "multiple double quotes",
			input:    "user\"na\"me",
			expected: "\"user\"\"na\"\"me\"",
		},
		{
			name:     "other special chars are preserved",
			input:    "user;name",
			expected: "\"user;name\"",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := EscapeIdentifier(c.input)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}
