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
