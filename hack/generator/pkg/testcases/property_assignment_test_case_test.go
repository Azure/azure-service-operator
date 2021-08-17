/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestPropertyAssignmentTestCase_AsFunc(t *testing.T) {
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

	types := make(astmodel.Types)

	conversionContext := conversions.NewPropertyConversionContext(types, idFactory)
	convertFrom, err := functions.NewPropertyAssignmentFunction(currentSpec, otherSpec, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	convertTo, err := functions.NewPropertyAssignmentFunction(currentSpec, otherSpec, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	currentSpec, err = functionInjector.Inject(currentSpec, convertTo, convertFrom)
	g.Expect(err).To(Succeed())

	container, ok := astmodel.AsFunctionContainer(currentSpec.Type())
	g.Expect(ok).To(BeTrue())

	testcase := NewPropertyAssignmentTestCase(currentSpec.Name(), container, idFactory)

	currentSpec, err = testCaseInjector.Inject(currentSpec, testcase)
	g.Expect(err).To(Succeed())

	fileDef := test.CreateTestFileDefinition(currentSpec)
	test.AssertFileGeneratesExpectedCode(t, fileDef, "person_test")
}
