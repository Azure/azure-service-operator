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

	// TODO: Add more test cases, including ones to cover existing redactions
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
		"Redis Cache Access Keys": {
			input:    `"accessKeys":{"primaryKey":"SECRETVALUEFROMAZURE=","secondaryKey":"SECRETVALUEFROMAZURE="}`,
			expected: `"accessKeys":{"primaryKey":"{KEY}","secondaryKey":"{KEY}"}`,
		},
		"SignalR Connection String": {
			input:    `body: '{"primaryConnectionString":"Endpoint=https://asotest-signalr-mposrl.service.signalr.net;AccessKey=SECRETVALUEFROMAZURE;Version=1.0;","secondaryConnectionString":"Endpoint=https://asotest-signalr-mposrl.service.signalr.net;AccessKey=SECRETVALUEFROMAZURE;Version=1.0;"}'`,
			expected: `body: '{"primaryConnectionString":"Endpoint=https://asotest-signalr-mposrl.service.signalr.net;AccessKey={KEY};Version=1.0;","secondaryConnectionString":"Endpoint=https://asotest-signalr-mposrl.service.signalr.net;AccessKey={KEY};Version=1.0;"}'`,
		},
		"Communication Service Connection String": {
			input:    `{"primaryKey":"{KEY}","secondaryKey":"{KEY}","primaryConnectionString":"endpoint=https://asotest-commssvc.unitedstates.communication.azure.com/;accesskey=FAKESECRETVALUEFROMAZURE1234567890abcdef","secondaryConnectionString":"endpoint=https://asotest-commssvc.unitedstates.communication.azure.com/;accesskey=FAKESECRETVALUEFROMAZURE0987654321fedcba"}`,
			expected: `{"primaryKey":"{KEY}","secondaryKey":"{KEY}","primaryConnectionString":"endpoint=https://asotest-commssvc.unitedstates.communication.azure.com/;accesskey={KEY}","secondaryConnectionString":"endpoint=https://asotest-commssvc.unitedstates.communication.azure.com/;accesskey={KEY}"}`,
		},
		"Mixed casing AccessKey preserves original case": {
			input:    `primary;AccessKey=SECRET1;Version=1.0; secondary;accesskey=SECRET2`,
			expected: `primary;AccessKey={KEY};Version=1.0; secondary;accesskey={KEY}`,
		},
		//
		// Test cases for the identity of whoever recorded the test
		//
		"Recording identity": {
			input:    `{"createdBy":"someone@example.org","createdByType":"User","lastModifiedBy":"someone@example.org"}`,
			expected: `{"createdBy":"redacted@example.com","createdByType":"User","lastModifiedBy":"redacted@example.com"}`,
		},
		"Recording identity of a service principal is left alone": {
			input:    `{"createdBy":"00000000-1111-2222-3333-444444444444","createdByType":"Application"}`,
			expected: `{"createdBy":"00000000-1111-2222-3333-444444444444","createdByType":"Application"}`,
		},
		"Caller object ID": {
			input:    `operationId=1a2b3c4d,objectId=80e349d8-21c2-4322-907d-c1dbac5915dd`,
			expected: `operationId=1a2b3c4d,objectId=00000000-0000-0000-0000-000000000000`,
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
