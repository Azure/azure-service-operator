/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

var (
	crmV1 = test.MakeLocalPackageReference("crm", "v1")
	crmV2 = test.MakeLocalPackageReference("crm", "v2")
	crmV3 = test.MakeLocalPackageReference("crm", "v3")
	crmV4 = test.MakeLocalPackageReference("crm", "v4")

	batchV1 = test.MakeLocalPackageReference("batch", "v1")
	batchV2 = test.MakeLocalPackageReference("batch", "v2")
)

func TestConversionGraph_GivenLink_ReturnsLink(t *testing.T) {
	g := NewGomegaWithT(t)

	graph := NewConversionGraph()
	graph.AddLink(crmV1, crmV2)
	graph.AddLink(batchV1, batchV2)

	crmActual, ok := graph.LookupTransition(crmV1)
	g.Expect(ok).To(BeTrue())
	g.Expect(crmActual).To(Equal(crmV2))
}

func TestConversionGraph_GivenTypeName_ReturnsExpectedHubTypeName(t *testing.T) {
	g := NewGomegaWithT(t)

	personV1 := astmodel.MakeTypeName(crmV1, "Person")
	personV2 := astmodel.MakeTypeName(crmV2, "Person")
	personV3 := astmodel.MakeTypeName(crmV3, "Person")
	personV4 := astmodel.MakeTypeName(crmV4, "Person")

	graph := NewConversionGraph()
	graph.AddLink(crmV1, crmV2)
	graph.AddLink(crmV2, crmV3)
	graph.AddLink(crmV3, crmV4)

	cases := []struct {
		name     string
		typeName astmodel.TypeName
		expected astmodel.TypeName
	}{
		{"Hub type resolves to self", personV4, personV4},
		{"Directly linked resolves", personV3, personV4},
		{"Doubly linked resolves", personV2, personV4},
		{"Indirectly linked resolves", personV1, personV4},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := graph.FindHubTypeName(c.typeName)
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}
