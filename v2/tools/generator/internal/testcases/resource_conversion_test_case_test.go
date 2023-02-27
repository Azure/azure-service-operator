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

func TestGolden_ResourceConversionTestCase_AsFunc(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	functionInjector := astmodel.NewFunctionInjector()
	testCaseInjector := astmodel.NewTestCaseInjector()
	interfaceInjector := astmodel.NewInterfaceInjector()

	personSpec2020 := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	personStatus2020 := test.CreateStatus(test.Pkg2020, "Person")

	person2020 := test.CreateResource(test.Pkg2020, "Person", personSpec2020, personStatus2020)

	personSpec2021 := test.CreateSpec(
		test.Pkg2021,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty,
		test.FullAddressProperty)

	personStatus2021 := test.CreateStatus(test.Pkg2021, "Person")

	person2021 := test.CreateResource(test.Pkg2021, "Person", personSpec2021, personStatus2021)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, personSpec2020, personStatus2020)
	defs.AddAll(person2021, personSpec2021, personStatus2021)

	cfg := config.NewObjectModelConfiguration()
	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, defs, idFactory).
		WithConfiguration(cfg)
	assignFrom, err := functions.NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	assignTo, err := functions.NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	person2020, err = functionInjector.Inject(person2020, assignTo, assignFrom)
	g.Expect(err).To(Succeed())

	convertFrom := functions.NewResourceConversionFunction(person2021.Name(), assignFrom, idFactory)
	convertTo := functions.NewResourceConversionFunction(person2021.Name(), assignTo, idFactory)

	implementation := astmodel.NewInterfaceImplementation(astmodel.ConvertibleInterface, convertFrom, convertTo)
	person2020, err = interfaceInjector.Inject(person2020, implementation)
	g.Expect(err).To(Succeed())

	resource := astmodel.MustBeResourceType(person2020.Type())
	testcase, err := NewResourceConversionTestCase(person2020.Name(), resource, idFactory)
	g.Expect(err).To(Succeed())

	person2020modified, err := testCaseInjector.Inject(person2020, testcase)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(
		t,
		"person",
		person2020modified,
		test.DiffWith(person2020),
		test.IncludeTestFiles(),
		test.ExcludeCodeFiles())
}
