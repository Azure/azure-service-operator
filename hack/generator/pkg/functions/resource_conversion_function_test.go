/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

// Test_ResourceConversionFunction_DirectConversion_GeneratesExpectedCode tests the code when the ConvertTo() and
// ConvertFrom() functions are directly converting to/from the Hub type, without any intermediate step.
func Test_ResourceConversionFunction_DirectConversion_GeneratesExpectedCode(t *testing.T) {
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
	types := make(astmodel.Types)
	types.AddAll(person2020, personSpec2020, personStatus2020)
	types.AddAll(person2021, personSpec2021, personStatus2021)

	conversionContext := conversions.NewPropertyConversionContext(types, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	// Create Resource Conversion Functions
	convertTo := NewResourceConversionFunction(person2021.Name(), propertyAssignTo, idFactory)
	convertFrom := NewResourceConversionFunction(person2021.Name(), propertyAssignFrom, idFactory)

	// Inject these methods into person2020
	injector := astmodel.NewFunctionInjector()
	person2020, err = injector.Inject(person2020, propertyAssignTo, propertyAssignFrom, convertTo, convertFrom)
	g.Expect(err).To(Succeed())

	// Write to a file
	fileDef := test.CreateFileDefinition(person2020)
	test.AssertFileGeneratesExpectedCode(t, fileDef, "ResourceConversionFunction")
}

// Test_ResourceConversionFunction_IndirectConversion_GeneratesExpectedCode tests the code when the ConvertTo() and
// ConvertFrom() functions can't convert directly to/from the hub type and are forced to stage the conversion on an
// intermediate type.
func Test_ResourceConversionFunction_IndirectConversion_GeneratesExpectedCode(t *testing.T) {
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
	types := make(astmodel.Types)
	types.AddAll(person2020, personSpec2020, personStatus2020)
	types.AddAll(person2021, personSpec2021, personStatus2021)
	types.AddAll(person2022, personSpec2022, personStatus2022)

	conversionContext := conversions.NewPropertyConversionContext(types, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	// Create Resource Conversion Functions
	convertTo := NewResourceConversionFunction(person2022.Name(), propertyAssignTo, idFactory)
	convertFrom := NewResourceConversionFunction(person2022.Name(), propertyAssignFrom, idFactory)

	// Inject these methods into person2020
	injector := astmodel.NewFunctionInjector()
	person2020, err = injector.Inject(person2020, propertyAssignTo, propertyAssignFrom, convertTo, convertFrom)
	g.Expect(err).To(Succeed())

	// Write to a file
	fileDef := test.CreateFileDefinition(person2020)
	test.AssertFileGeneratesExpectedCode(t, fileDef, "ResourceConversionFunction")
}
