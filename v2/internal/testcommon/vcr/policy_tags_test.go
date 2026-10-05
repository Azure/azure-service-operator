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

func Test_Redactor_givenPolicyTags_removesThemFromTagsObjects(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input    string
		expected string
	}{
		"Tag following another": {
			input:    `{"tags":{"CreatedAt":"then","Owner":"someone"}}`,
			expected: `{"tags":{"CreatedAt":"then"}}`,
		},
		"Tag preceding another": {
			input:    `{"tags":{"Owner":"someone","CreatedAt":"then"}}`,
			expected: `{"tags":{"CreatedAt":"then"}}`,
		},
		"Tag between two others": {
			input:    `{"tags":{"CreatedAt":"then","Owner":"someone","Extra":"kept"}}`,
			expected: `{"tags":{"CreatedAt":"then","Extra":"kept"}}`,
		},
		"Every tag of the object": {
			input:    `{"tags":{"Owner":"someone","Project":"something"}}`,
			expected: `{"tags":{}}`,
		},
		"Value containing an escaped quote": {
			input:    `{"tags":{"Quoted":"some\"one","CreatedAt":"then"}}`,
			expected: `{"tags":{"CreatedAt":"then"}}`,
		},
		"Value containing a brace": {
			input:    `{"tags":{"Braced":"}","CreatedAt":"then"}}`,
			expected: `{"tags":{"CreatedAt":"then"}}`,
		},
		"Value ARM echoes back unescaped": {
			input:    `{"tags":{"Angled":"a<b","CreatedAt":"then"}}`,
			expected: `{"tags":{"CreatedAt":"then"}}`,
		},
		"Value the SDK writes escaped": {
			input:    `{"tags":{"Angled":"a\u003cb","CreatedAt":"then"}}`,
			expected: `{"tags":{"CreatedAt":"then"}}`,
		},
		"Tag of the same name with a different value": {
			input:    `{"tags":{"Owner":"someone else","CreatedAt":"then"}}`,
			expected: `{"tags":{"Owner":"someone else","CreatedAt":"then"}}`,
		},
		"Field of the same name outside a tags object": {
			input:    `{"properties":{"Owner":"someone"},"tags":{"Owner":"someone"}}`,
			expected: `{"properties":{"Owner":"someone"},"tags":{}}`,
		},
		"Second tags object in the same body": {
			input:    `{"value":[{"tags":{"Owner":"someone","CreatedAt":"then"}},{"tags":{"Owner":"someone"}}]}`,
			expected: `{"value":[{"tags":{"CreatedAt":"then"}},{"tags":{}}]}`,
		},
		"Unterminated tags object is left alone": {
			input:    `{"tags":{"Owner":"someone"`,
			expected: `{"tags":{"Owner":"someone"`,
		},
	}

	azureIDs := creds.AzureIDs{
		ResourceGroupTags: map[string]string{
			"Owner":   "someone",
			"Project": "something",
			"Quoted":  `some"one`,
			"Braced":  "}",
			"Angled":  "a<b",
		},
	}
	redactor := NewRedactor(azureIDs)

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(redactor.HideRecordingData(c.input)).To(Equal(c.expected))
		})
	}
}

func Test_Redactor_givenNoPolicyTags_leavesTagsAlone(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	redactor := NewRedactor(creds.AzureIDs{})
	body := `{"tags":{"Owner":"someone","CreatedAt":"then"}}`

	g.Expect(redactor.HideRecordingData(body)).To(Equal(body))
}
