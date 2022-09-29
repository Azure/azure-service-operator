/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"encoding/json"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func Test_SerializeOptionalConfigMapReference(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		optionalRef OptionalConfigMapReference
		expected    string
	}{
		{
			name: "Serialized as string",
			optionalRef: OptionalConfigMapReference{
				Value: to.StringPtr("test"),
			},
			expected: `"test"`,
		},
		{
			name: "Serialized as ConfigMapReference",
			optionalRef: OptionalConfigMapReference{
				Ref: &ConfigMapReference{
					Name: "foo",
					Key:  "key",
				},
			},
			expected: `{"name":"foo","key":"key"}`,
		},
		{
			name:        "Serialized when empty", // This isn't really an expected scenario, but we try to do something sensible anyway.
			optionalRef: OptionalConfigMapReference{},
			expected:    `{}`,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			// Serialize as non-ptr
			result, err := json.Marshal(c.optionalRef)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(result)).To(Equal(c.expected))

			// Also serialize as ptr, they should be the same
			result, err = json.Marshal(&c.optionalRef)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(result)).To(Equal(c.expected))
		})
	}
}

func Test_DeserializeOptionalConfigMapReference(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected OptionalConfigMapReference
	}{
		{
			name:  "Deserialized from string",
			input: `"test"`,
			expected: OptionalConfigMapReference{
				Value: to.StringPtr("test"),
			},
		},
		{
			name:  "Deserialized from reference",
			input: `{"name":"foo","key":"key"}`,
			expected: OptionalConfigMapReference{
				Ref: &ConfigMapReference{
					Name: "foo",
					Key:  "key",
				},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var result OptionalConfigMapReference
			err := json.Unmarshal([]byte(c.input), &result)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(result).To(Equal(c.expected))
		})
	}
}

func Test_DeserializeOptionalConfigMapReference_InvalidJSONProducesError(t *testing.T) {
	t.Parallel()
	format.TruncatedDiff = false
	cases := []struct {
		name              string
		input             string
		expectedErrString string
	}{
		{
			name:              "Deserialized from int",
			input:             `10`,
			expectedErrString: "unexpected OptionalConfigMapReference format. Expected string or ConfigMapReference: json: cannot unmarshal number into Go value of type genruntime.ConfigMapReference",
		},
		{
			name:              "Deserialized from object with unknown field",
			input:             `{"name":"foo","key":"key","bar":"bar"}`,
			expectedErrString: `unexpected OptionalConfigMapReference format. Expected string or ConfigMapReference: json: unknown field "bar"`,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var result OptionalConfigMapReference
			err := json.Unmarshal([]byte(c.input), &result)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal(c.expectedErrString))
		})
	}
}
