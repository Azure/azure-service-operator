/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting_test

import (
	"bytes"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"

	"github.com/onsi/gomega"

	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
	"github.com/sebdah/goldie/v2"
)

func Test_TypeCatalogReport_GivenTypes_ShowsExpectedDetails(t *testing.T) {
	t.Parallel()

	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	defs := createDefinitionSet()

	var content bytes.Buffer
	rpt := reporting.NewTypeCatalogReport(defs)
	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())

	golden.Assert(t, t.Name(), content.Bytes())
}

func Test_TypeCatalogReport_GivenTypes_WhenInlined_ShowsExpectedDetails(t *testing.T) {
	t.Parallel()

	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	defs := createDefinitionSet()

	var content bytes.Buffer
	rpt := reporting.NewTypeCatalogReport(defs)
	rpt.InlineTypes()

	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())
	golden.Assert(t, t.Name(), content.Bytes())
}

func Test_TypeCatalogReport_GivenDirectlyRecursiveType_WhenInlined_ShowsExpectedDetails(t *testing.T) {
	t.Parallel()

	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	personName := astmodel.MakeTypeName(test.Pkg2020, "Person")

	parentProperty := astmodel.NewPropertyDefinition(
		"Parent",
		"parentName",
		astmodel.NewOptionalType(personName)).
		WithDescription("Optional reference to our parent")

	personObj := astmodel.NewObjectType().
		WithProperties(
			test.FullNameProperty,
			test.FamilyNameProperty,
			test.KnownAsProperty,
			parentProperty)

	person := astmodel.MakeTypeDefinition(personName, personObj)

	relationship := test.CreateObjectDefinition(
		test.Pkg2020,
		"Relationship",
		astmodel.NewPropertyDefinition(
			"FirstParty",
			"firstparty",
			astmodel.NewOptionalType(personName)),
		astmodel.NewPropertyDefinition(
			"SecondParty",
			"secondparty",
			astmodel.NewOptionalType(personName)))

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(person)
	defs.Add(relationship)

	var content bytes.Buffer
	rpt := reporting.NewTypeCatalogReport(defs)
	rpt.InlineTypes()

	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())
	golden.Assert(t, t.Name(), content.Bytes())
}

func createDefinitionSet() astmodel.TypeDefinitionSet {
	testSpec := test.CreateSpec(
		test.Pkg2020,
		"TestResource",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)

	testStatus := test.CreateStatus(
		test.Pkg2020,
		"TestResource")

	testResource := test.CreateResource(
		test.Pkg2020,
		"TestResource",
		testSpec,
		testStatus)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(testResource)
	defs.Add(testSpec)
	defs.Add(testStatus)
	return defs
}
