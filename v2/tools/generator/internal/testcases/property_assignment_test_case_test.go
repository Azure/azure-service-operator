/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGolden_PropertyAssignmentTestCase_AsFunc(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	functionInjector := astmodel.NewFunctionInjector()
	testCaseInjector := astmodel.NewTestCaseInjector()

	currentSpec := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	otherSpec := test.CreateSpec(
		test.Pkg2021,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	defs := make(astmodel.TypeDefinitionSet)
	cfg := config.NewObjectModelConfiguration()
	conversionContext := conversions.NewPropertyConversionContext("AssignProperties", defs, idFactory).
		WithConfiguration(cfg)
	convertFrom, err := functions.NewPropertyAssignmentFunction(currentSpec, otherSpec, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	convertTo, err := functions.NewPropertyAssignmentFunction(currentSpec, otherSpec, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	currentSpec, err = functionInjector.Inject(currentSpec, convertTo, convertFrom)
	g.Expect(err).To(Succeed())

	container, ok := astmodel.AsFunctionContainer(currentSpec.Type())
	g.Expect(ok).To(BeTrue())

	testcase := NewPropertyAssignmentTestCase(currentSpec.Name(), container, idFactory)

	modifiedSpec, err := testCaseInjector.Inject(currentSpec, testcase)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "person_test", modifiedSpec, test.DiffWith(currentSpec), test.IncludeTestFiles())
}
