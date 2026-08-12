/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package labels_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/labels"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label     string
		wantKey   string
		wantValue string
		wantErr   bool
	}{
		{"example.com/label=value", "example.com/label", "value", false},
		{"example.com/label=", "example.com/label", "", false},
		{"=value", "", "", true},
		{"example.com/label", "", "", true},
		{"example.com/test/label", "", "", true},
		{"thisisaverylonglabelname_solonginfactthatitisgoingtocauseanerror", "", "", true},
		{"", "", "", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.label, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual, err := labels.Parse(tt.label)

			if tt.wantErr {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}

			g.Expect(actual.Key).To(Equal(tt.wantKey))
			g.Expect(actual.Value).To(Equal(tt.wantValue))
		})
	}
}

func TestParseMap(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		value         string
		expected      map[string]string
		expectedError string
	}{
		"empty": {
			value:    "",
			expected: map[string]string{},
		},
		"empty label value": {
			value:    "example.com/owner=",
			expected: map[string]string{"example.com/owner": ""},
		},
		"missing equals": {
			value:         "example.com/owner",
			expectedError: `label "example.com/owner" must be in key=value form`,
		},
		"selector operator": {
			value:         "environment!=production",
			expectedError: `invalid label key "environment!"`,
		},
		"invalid key": {
			value:         "not a label=value",
			expectedError: `invalid label key "not a label"`,
		},
		"invalid value": {
			value:         "example.com/owner=not a value",
			expectedError: `invalid value for label "example.com/owner"`,
		},
		"duplicate key": {
			value:         "environment=production,environment=staging",
			expectedError: `label "environment" was specified more than once`,
		},
		"semicolon separated": {
			value:    "environment=production;example.com/owner=aso",
			expected: map[string]string{"environment": "production", "example.com/owner": "aso"},
		},
		"mixed separators": {
			value:    "environment=production,example.com/owner=aso;example.com/team=platform",
			expected: map[string]string{"environment": "production", "example.com/owner": "aso", "example.com/team": "platform"},
		},
		"blank entries are ignored": {
			value:    ",environment=production;; ,",
			expected: map[string]string{"environment": "production"},
		},
		"whitespace around keys and values is trimmed": {
			value:    " environment = production , example.com/owner = aso ",
			expected: map[string]string{"environment": "production", "example.com/owner": "aso"},
		},
		"duplicate key across separators": {
			value:         "environment=production;environment=staging",
			expectedError: `label "environment" was specified more than once`,
		},
		"long prefixed key is accepted": {
			value:    "aaaaaaaaaa.bbbbbbbbbb.cccccccccc.dddddddddd.eeeeeeeeee.ffffffffff.gggg.example.com/owner=aso",
			expected: map[string]string{"aaaaaaaaaa.bbbbbbbbbb.cccccccccc.dddddddddd.eeeeeeeeee.ffffffffff.gggg.example.com/owner": "aso"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)

			actual, err := labels.ParseMap(c.value)

			if c.expectedError != "" {
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedError)))
				return
			}

			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(actual).To(HaveLen(len(c.expected)))
			for key, value := range c.expected {
				g.Expect(actual).To(HaveKeyWithValue(key, value))
			}
		})
	}
}
