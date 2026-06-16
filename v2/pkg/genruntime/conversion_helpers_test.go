/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime_test

import (
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestConvertStringToJSON_RoundTripsThroughConvertJSONToString(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		value string
	}{
		"empty string": {
			value: "",
		},
		"simple string": {
			value: "value",
		},
		"numeric string": {
			value: "123",
		},
		"boolean string": {
			value: "true",
		},
		"json object string": {
			value: `{"key":"value"}`,
		},
	}

	for name, c := range cases {
		name, c := name, c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			j := genruntime.ConvertStringToJSON(c.value)
			actual := genruntime.ConvertJSONToString(j)
			g.Expect(actual).To(Equal(c.value))
		})
	}
}

func TestConvertJSONToString_EmptyJSONReturnsEmptyString(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(genruntime.ConvertJSONToString(v1.JSON{})).To(Equal(""))
}
