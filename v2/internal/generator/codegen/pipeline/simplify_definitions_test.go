/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
)

func Test_SimplifyDefinitionsPipelineStage_GivenTypes_FlattensToExpectedTypes(t *testing.T) {
	fullName := astmodel.NewPropertyDefinition("FullName", "full-name", astmodel.StringType)
	familyName := astmodel.NewPropertyDefinition("FamilyName", "family-name", astmodel.StringType)
	knownAs := astmodel.NewPropertyDefinition("KnownAs", "known-as", astmodel.StringType)

	obj := astmodel.NewObjectType().
		WithProperties(fullName, familyName, knownAs)

	flagged := astmodel.ARMFlag.ApplyTo(obj)

	cases := []struct {
		name     string
		original astmodel.Type
		expected astmodel.Type
	}{
		{"Object Type remains unchanged", obj, obj},
		{"Flagged type gets simplified", flagged, obj},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			visitor := createSimplifyingVisitor()
			result, err := visitor.Visit(c.original, nil)
			g.Expect(err).To(BeNil())
			g.Expect(astmodel.TypeEquals(result, c.expected)).To(BeTrue())
		})
	}
}
