/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

/*
 * This file contains tests for failures creating property conversions
 *
 * Tests for successfully created property conversions are found in property_assignment_function_test.go as the best
 * way to verify those creations are created properly is to render them as code into a golden file.
 */

func TestCreateTypeConversion_GivenIncompatibleEndpoints_ReturnsExpectedError(t *testing.T) {
	t.Parallel()

	stringEnum := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "StringEnum"),
		astmodel.NewEnumType(astmodel.StringType))

	intEnum := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "IntEnum"),
		astmodel.NewEnumType(astmodel.IntType))

	addressObject := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "Address"),
		astmodel.NewObjectType())

	locationObject := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2022, "Location"),
		astmodel.NewObjectType())

	cases := []struct {
		name           string
		source         astmodel.TypeDefinition
		destination    astmodel.TypeDefinition
		errorSubstring string
	}{
		{
			"Incompatible Enumerations",
			stringEnum,
			intEnum,
			"no conversion from string to int",
		},
		{
			"Object definitions with different names and no active rename",
			addressObject,
			locationObject,
			"no configuration to rename Address to Location",
		},
	}

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(stringEnum, intEnum)
	defs.AddAll(addressObject, locationObject)

	idFactory := astmodel.NewIdentifierFactory()
	conversionContext := NewPropertyConversionContext(defs, idFactory).
		WithDirection(ConvertTo)

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			readable := NewTypedConversionEndpoint(c.source.Name(), "property")
			writable := NewTypedConversionEndpoint(c.destination.Name(), "property")
			conv, err := CreateTypeConversion(readable, writable, conversionContext)
			g.Expect(conv).To(BeNil())
			g.Expect(err).NotTo(BeNil())
			g.Expect(err.Error()).To(ContainSubstring(c.errorSubstring))
		})
	}
}
