// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package app

import (
	"flag"
	"testing"

	. "github.com/onsi/gomega"
)

func TestCRDLabelsFlag(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flags := InitFlags(flagSet)

	g.Expect(flagSet.Parse([]string{"--crd-labels=cluster.x-k8s.io/provider=infrastructure-azure,example.com/owner=aso"})).To(Succeed())
	g.Expect(flags.CRDLabels).To(Equal("cluster.x-k8s.io/provider=infrastructure-azure,example.com/owner=aso"))

	parsed, err := parseCRDLabels(flags.CRDLabels)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(parsed).To(HaveLen(2))
	g.Expect(parsed).To(HaveKeyWithValue("cluster.x-k8s.io/provider", "infrastructure-azure"))
	g.Expect(parsed).To(HaveKeyWithValue("example.com/owner", "aso"))
}

// TestParseCRDLabels covers the CRD specific behaviour layered on top of labels.ParseMap.
// General label parsing and validation is covered by TestParseMap in the labels package.
func TestParseCRDLabels(t *testing.T) {
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
		"reserved app label": {
			value:         "app.kubernetes.io/name=mine",
			expectedError: `label "app.kubernetes.io/name" is reserved`,
		},
		"reserved version label": {
			value:         "app.kubernetes.io/version=v1.0.0",
			expectedError: `label "app.kubernetes.io/version" is reserved`,
		},
		"reserved old version label": {
			value:         "serviceoperator.azure.com/version=v1.0.0",
			expectedError: `label "serviceoperator.azure.com/version" is reserved`,
		},
		"reserved prefix": {
			value:         "serviceoperator.azure.com/anything=value",
			expectedError: `label "serviceoperator.azure.com/anything" is reserved`,
		},
		"reserved label alongside a valid one": {
			value:         "example.com/owner=aso,app.kubernetes.io/name=mine",
			expectedError: `label "app.kubernetes.io/name" is reserved`,
		},
		"reserved label with surrounding whitespace": {
			value:         " app.kubernetes.io/version = v1.0.0 ",
			expectedError: `label "app.kubernetes.io/version" is reserved`,
		},
		"unreserved app.kubernetes.io label": {
			value:    "app.kubernetes.io/part-of=platform",
			expected: map[string]string{"app.kubernetes.io/part-of": "platform"},
		},
		"valid labels are passed through": {
			value:    "environment=production;example.com/owner=aso",
			expected: map[string]string{"environment": "production", "example.com/owner": "aso"},
		},
		"parse errors are propagated": {
			value:         "not a label=value",
			expectedError: `invalid label key "not a label"`,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)

			actual, err := parseCRDLabels(c.value)

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
