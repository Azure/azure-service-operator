/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestConversionGraph_GivenLink_ReturnsLink(t *testing.T) {
	g := NewGomegaWithT(t)

	personV1 := test.MakeLocalPackageReference("person", "v1")
	personV2 := test.MakeLocalPackageReference("person", "v2")

	placeV1 := test.MakeLocalPackageReference("place", "v1")
	placeV2 := test.MakeLocalPackageReference("place", "v2")

	graph := NewConversionGraph()
	graph.AddLink(personV1, personV2)
	graph.AddLink(placeV1, placeV2)

	personActual, ok := graph.LookupTransition(personV1)
	g.Expect(ok).To(BeTrue())
	g.Expect(personActual).To(Equal(personV2))
}
