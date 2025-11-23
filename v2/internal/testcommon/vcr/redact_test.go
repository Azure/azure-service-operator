/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package vcr

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_hideAppConfigurationKeySecrets(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input    string
		expected string
	}{
		"Empty string": {
			input:    "",
			expected: "",
		},
		"No secrets": {
			input:    `{"someKey":"someValue"}`,
			expected: `{"someKey":"someValue"}`,
		},
		"Connection string": {
			input:    `"Endpoint=https://asotest-confstore-fsrajl.azconfig.io;Id=GmDj;Secret=SECRETVALUEFROMAZURE"`,
			expected: `"Endpoint=https://asotest-confstore-fsrajl.azconfig.io;Id=GmDj;Secret={KEY}"`,
		},
		"Configuration Key": {
			input:    `{"id":"GmDj","name":"Primary","value":"SECRETVALUEFROMAZURE","connectionString":"Endpoint=https://asotest-confstore-fsrajl.azconfig.io;Id=GmDj;Secret=SECRETVALUEFROMAZURE","lastModified":"2025-11-19T01:22:31+00:00","readOnly":false}`,
			expected: `{"id":"GmDj","name":"Primary","value":"{KEY}","connectionString":"Endpoint=https://asotest-confstore-fsrajl.azconfig.io;Id=GmDj;Secret={KEY}","lastModified":"2025-11-19T01:22:31+00:00","readOnly":false}`,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := hideAppConfigurationKeySecrets(c.input)

			g.Expect(actual).To(Equal(c.expected))
		})
	}
}
