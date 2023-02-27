/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// TestGolden_ResourceConversionFunction_DirectConversion_GeneratesExpectedCode tests the code when the ConvertTo() and
// ConvertFrom() functions are directly converting to/from the Hub type, without any intermediate step.
func TestGolden_ResourceConversionFunction_DirectConversion_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our upstream type
	personSpec2020 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	personStatus2020 := test.CreateStatus(test.Pkg2020, "Person")
	person2020 := test.CreateResource(test.Pkg2020, "Person", personSpec2020, personStatus2020)

	// Create our downstream type
	personSpec2021 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty, test.CityProperty)
	personStatus2021 := test.CreateStatus(test.Pkg2021, "Person")
	person2021 := test.CreateResource(test.Pkg2021, "Person", personSpec2021, personStatus2021)

	// Create Property Assignment functions
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, personSpec2020, personStatus2020)
	defs.AddAll(person2021, personSpec2021, personStatus2021)

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, defs, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	// Create Resource Conversion Functions
	convertTo := NewResourceConversionFunction(person2021.Name(), propertyAssignTo, idFactory)
	convertFrom := NewResourceConversionFunction(person2021.Name(), propertyAssignFrom, idFactory)

	// Inject these methods into person2020
	// We omit the propertyAssignment functions as they are tested elsewhere
	injector := astmodel.NewFunctionInjector()
	modified, err := injector.Inject(person2020, convertTo, convertFrom)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "ResourceConversionFunction", modified, test.DiffWith(person2020))
}

// TestGolden_ResourceConversionFunction_IndirectConversion_GeneratesExpectedCode tests the code when the ConvertTo() and
// ConvertFrom() functions can't convert directly to/from the hub type and are forced to stage the conversion on an
// intermediate type.
func TestGolden_ResourceConversionFunction_IndirectConversion_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our upstream type
	personSpec2020 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	personStatus2020 := test.CreateStatus(test.Pkg2020, "Person")
	person2020 := test.CreateResource(test.Pkg2020, "Person", personSpec2020, personStatus2020)

	// Create our downstream type - directly convertible with the upstream type
	personSpec2021 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty, test.CityProperty)
	personStatus2021 := test.CreateStatus(test.Pkg2021, "Person")
	person2021 := test.CreateResource(test.Pkg2021, "Person", personSpec2021, personStatus2021)

	// Create our hub type - multiple conversion steps away from Pkg2020
	personSpec2022 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty, test.CityProperty)
	personStatus2022 := test.CreateStatus(test.Pkg2022, "Person")
	person2022 := test.CreateResource(test.Pkg2022, "Person", personSpec2021, personStatus2021)

	// Create Property Assignment functions
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, personSpec2020, personStatus2020)
	defs.AddAll(person2021, personSpec2021, personStatus2021)
	defs.AddAll(person2022, personSpec2022, personStatus2022)

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, defs, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	// Create Resource Conversion Functions
	convertTo := NewResourceConversionFunction(person2022.Name(), propertyAssignTo, idFactory)
	convertFrom := NewResourceConversionFunction(person2022.Name(), propertyAssignFrom, idFactory)

	// Inject these methods into person2020
	// We omit the propertyAssignment functions as they are tested elsewhere
	injector := astmodel.NewFunctionInjector()
	modified, err := injector.Inject(person2020, convertTo, convertFrom)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "ResourceConversionFunction", modified, test.DiffWith(person2020))
}
