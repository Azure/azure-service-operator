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

// TestGolden_NewSpecChainedConversionFunction_Conversion_GeneratesExpectedCode tests the code when the ConvertToSpec() and
// ConvertFromSpec() functions are converting to/from spec types that aren't the hub type
func TestGolden_NewSpecChainedConversionFunction_Conversion_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our upstream type
	personSpec2020 := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	// Create our downstream type
	personSpec2021 := test.CreateSpec(
		test.Pkg2021,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty,
		test.CityProperty)

	// Create Property Assignment functions
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(personSpec2020)
	defs.AddAll(personSpec2021)

	conversionContext := conversions.NewPropertyConversionContext(defs, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentFunction(
		personSpec2020, personSpec2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFunction(
		personSpec2020, personSpec2021, conversionContext, conversions.ConvertFrom)
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
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(
		t, "Person", personSpec2020updated, test.DiffWith(personSpec2020))
}

// TestGolden_NewStatusChainedConversionFunction_Conversion_GeneratesExpectedCode tests the code when the ConvertToStatus()
// and ConvertFromStatus() functions are converting to/from status types that aren't the hub type
func TestGolden_NewStatusChainedConversionFunction_Conversion_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our upstream type
	personStatus2020 := test.CreateStatus(test.Pkg2020, "Person")

	// Create our downstream type
	personStatus2021 := test.CreateStatus(test.Pkg2021, "Person")

	// Create Property Assignment functions
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(personStatus2020)
	defs.AddAll(personStatus2021)

	conversionContext := conversions.NewPropertyConversionContext(defs, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentFunction(
		personStatus2020, personStatus2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFunction(
		personStatus2020, personStatus2021, conversionContext, conversions.ConvertFrom)
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
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(
		t, "Person", personStatus2020updated, test.DiffWith(personStatus2020))
}
