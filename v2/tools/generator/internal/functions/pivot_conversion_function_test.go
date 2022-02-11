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

// TestGolden_NewSpecPivotConversionFunction_Conversion_GeneratesExpectedCode tests the code when the ConvertToSpec() and
// ConvertFromSpec() functions are converting to/from spec types that ARE the hub type
func TestGolden_NewSpecPivotConversionFunction_Conversion_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our spec type
	original := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)

	// Create Spec Conversion Functions
	convertSpecTo := NewSpecPivotConversionFunction(conversions.ConvertTo, idFactory)
	convertSpecFrom := NewSpecPivotConversionFunction(conversions.ConvertFrom, idFactory)

	// Inject these methods into original
	injector := astmodel.NewFunctionInjector()
	modified, err := injector.Inject(original, convertSpecTo, convertSpecFrom)
	g.Expect(err).To(Succeed())

	// When verifying the golden file, check that the implementations of ConvertSpecTo() and ConvertSpecFrom() are
	// correctly injected on the spec. Verify that the code does what you expect. If you don't know what to expect,
	// check that they pivot to finish the conversion on the second spoke. :-)
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, t.Name(), modified, test.DiffWith(original))
}

// TestGolden_NewStatusPivotConversionFunction_Conversion_GeneratesExpectedCode tests the code when the ConvertToStatus() and
// ConvertFromStatus() functions are converting to/from status types that aren't the hub  type
func TestGolden_NewStatusPivotConversionFunction_Conversion_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our status type
	original := test.CreateStatus(test.Pkg2020, "Person")

	// Create Spec Conversion Functions
	convertSpecTo := NewStatusPivotConversionFunction(conversions.ConvertTo, idFactory)
	convertSpecFrom := NewStatusPivotConversionFunction(conversions.ConvertFrom, idFactory)

	// Inject these methods into original
	injector := astmodel.NewFunctionInjector()
	modified, err := injector.Inject(original, convertSpecTo, convertSpecFrom)
	g.Expect(err).To(Succeed())

	// When verifying the golden file, check that the implementations of ConvertStatusTo() and ConvertStatusFrom() are
	// correctly injected on the status. Verify that the code does what you expect. If you don't know what to expect,
	// check that they pivot to finish the conversion on the second spoke. :-)
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, t.Name(), modified, test.DiffWith(original))
}
