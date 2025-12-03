/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package vcr

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
)

func Test_Redactor_givenInput_returnsExpectedResult(t *testing.T) {
	t.Parallel()

	//TODO: Add more test cases, including ones to cover existing redactions
	cases := map[string]struct {
		input    string
		expected string
	}{
		"Empty string": {
			input:    "",
			expected: "",
		},
		//
		// Test cases for hideAppConfigurationKeySecrets
		//
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

	ids := creds.AzureIDs{}
	redactor := NewRedactor(ids)

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := redactor.HideRecordingData(c.input)

			g.Expect(actual).To(Equal(c.expected))
		})
	}
}
