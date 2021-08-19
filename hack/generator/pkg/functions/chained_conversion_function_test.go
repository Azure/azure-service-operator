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
	. "github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

// Test_NewSpecChainedConversionFunction_Conversion_GeneratesExpectedCode tests the code when the ConvertToSpec() and
// ConvertFromSpec() functions are converting to/from spec types that aren't the hub  type
func Test_NewSpecChainedConversionFunction_Conversion_GeneratesExpectedCode(t *testing.T) {
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our upstream type
	personSpec2020 := CreateSpec(Pkg2020, "Person", FullNameProperty, KnownAsProperty, FamilyNameProperty)

	// Create our downstream type
	personSpec2021 := CreateSpec(Pkg2021, "Person", FullNameProperty, KnownAsProperty, FamilyNameProperty, CityProperty)

	// Create Property Assignment functions
	types := make(astmodel.Types)
	types.AddAll(personSpec2020)
	types.AddAll(personSpec2021)

	conversionContext := conversions.NewPropertyConversionContext(types, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentFunction(personSpec2020, personSpec2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFunction(personSpec2020, personSpec2021, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	// Create Spec Conversion Functions
	convertSpecTo := NewSpecChainedConversionFunction(propertyAssignTo, idFactory)
	convertSpecFrom := NewSpecChainedConversionFunction(propertyAssignFrom, idFactory)

	// Inject these methods into personSpec2020
	// We omit the PropertyAssignment functions to reduce the amount of clutter, those are tested elsewhere
	injector := astmodel.NewFunctionInjector()
	personSpec2020updated, err := injector.Inject(personSpec2020, convertSpecTo, convertSpecFrom)
	g.Expect(err).To(Succeed())

	/*
	 * When verifying the golden file, look for the changes flagged between the reference and the updated definition
	 * and check that the implementations of ConvertSpecTo() and ConvertSpecFrom() are what you expect.
	 */
	AssertTypeDefinitionGeneratesExpectedCode(t, personSpec2020updated, "Person", DiffWithTypeDefinition(personSpec2020))
}

// Test_NewStatusChainedConversionFunction_Conversion_GeneratesExpectedCode tests the code when the ConvertToStatus()
// and ConvertFromStatus() functions are converting to/from status types that aren't the hub type
func Test_NewStatusChainedConversionFunction_Conversion_GeneratesExpectedCode(t *testing.T) {
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our upstream type
	personStatus2020 := CreateStatus(Pkg2020, "Person")

	// Create our downstream type
	personStatus2021 := CreateStatus(Pkg2021, "Person")

	// Create Property Assignment functions
	types := make(astmodel.Types)
	types.AddAll(personStatus2020)
	types.AddAll(personStatus2021)

	conversionContext := conversions.NewPropertyConversionContext(types, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentFunction(personStatus2020, personStatus2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFunction(personStatus2020, personStatus2021, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	// Create Spec Conversion Functions
	convertSpecTo := NewStatusChainedConversionFunction(propertyAssignTo, idFactory)
	convertSpecFrom := NewStatusChainedConversionFunction(propertyAssignFrom, idFactory)

	// Inject these methods into personStatus2020
	// We omit the PropertyAssignment functions to reduce the amount of clutter, those are tested elsewhere
	injector := astmodel.NewFunctionInjector()
	personStatus2020updated, err := injector.Inject(personStatus2020, convertSpecTo, convertSpecFrom)
	g.Expect(err).To(Succeed())

	/*
	 * When verifying the golden file, look for the changes flagged between the reference and the updated definition
	 * and check that the implementations of ConvertSpecTo() and ConvertSpecFrom() are what you expect.
	 */
	AssertTypeDefinitionGeneratesExpectedCode(t, personStatus2020updated, "Person", DiffWithTypeDefinition(personStatus2020))
}
