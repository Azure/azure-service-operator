/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting_test

import (
	"bytes"
	"testing"

	"github.com/onsi/gomega"
	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
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
	rpt := reporting.NewTypeCatalogReport(defs, reporting.InlineTypes)

	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())
	golden.Assert(t, t.Name(), content.Bytes())
}

func Test_TypeCatalogReport_GivenDirectlyRecursiveType_WhenInlined_ShowsExpectedDetails(t *testing.T) {
	t.Parallel()

	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	personName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")

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
	rpt := reporting.NewTypeCatalogReport(defs, reporting.InlineTypes)

	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())
	golden.Assert(t, t.Name(), content.Bytes())
}

func Test_TypeCatalogReport_GivenMapsAndArrays_ShowsExpectedDetails(t *testing.T) {
	t.Parallel()

	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	personName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")

	name := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Name"),
		astmodel.StringType)

	aliasesProperty := astmodel.NewPropertyDefinition(
		"Aliases",
		"aliases",
		astmodel.NewArrayType(astmodel.StringType)).
		WithDescription("Array of aliases")

	friendsProperty := astmodel.NewPropertyDefinition(
		"Friends",
		"friends",
		astmodel.NewArrayType(name.Name())).
		WithDescription("Array of friends")

	cohortProperty := astmodel.NewPropertyDefinition(
		"Cohort",
		"cohort",
		astmodel.NewMapType(astmodel.StringType, name.Name())).
		WithDescription("Map of nickname to actual name")

	person := astmodel.MakeTypeDefinition(
		personName,
		astmodel.NewObjectType().
			WithProperties(
				aliasesProperty,
				friendsProperty,
				cohortProperty))

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(person)
	defs.Add(name)

	var content bytes.Buffer
	rpt := reporting.NewTypeCatalogReport(defs)

	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())
	golden.Assert(t, t.Name(), content.Bytes())
}

func Test_TypeCatalogReport_GivenValidatedAndOptionalTypes_ShowsExpectedDetails(t *testing.T) {
	t.Parallel()

	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	personName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")

	maxLength := int64(100)
	minLength := int64(1)
	name := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Name"),
		astmodel.NewValidatedType(
			astmodel.StringType,
			astmodel.StringValidations{
				MaxLength: &maxLength,
				MinLength: &minLength,
			}))

	knownAsProperty := astmodel.NewPropertyDefinition(
		"KnownAs",
		"knownAs",
		astmodel.StringType).
		WithDescription("Required string property")

	familyNameProperty := astmodel.NewPropertyDefinition(
		"FamilyName",
		"familyName",
		astmodel.OptionalStringType).
		WithDescription("Optional string property")

	fullNameProperty := astmodel.NewPropertyDefinition(
		"FullName",
		"fullName",
		name.Type()).
		WithDescription("Validated required string")

	legalNameProperty := astmodel.NewPropertyDefinition(
		"LegalName",
		"legalName",
		astmodel.NewValidatedType(
			astmodel.OptionalStringType,
			astmodel.StringValidations{
				MaxLength: &maxLength,
				MinLength: &minLength,
			})).
		WithDescription("Validated optional string")

	nicknameProperty := astmodel.NewPropertyDefinition(
		"Nickname",
		"nickname",
		astmodel.NewOptionalType(
			astmodel.NewValidatedType(
				astmodel.StringType,
				astmodel.StringValidations{
					MaxLength: &maxLength,
					MinLength: &minLength,
				}))).
		WithDescription("Optional validated string")

	akaProperty := astmodel.NewPropertyDefinition(
		"AKA",
		"aka",
		name.Name()).
		WithDescription("Mandatory via type name")

	neeProperty := astmodel.NewPropertyDefinition(
		"Nee",
		"nee",
		astmodel.NewOptionalType(name.Name())).
		WithDescription("Optional via type name")

	person := astmodel.MakeTypeDefinition(
		personName,
		astmodel.NewObjectType().
			WithProperties(
				knownAsProperty,
				familyNameProperty,
				fullNameProperty,
				legalNameProperty,
				nicknameProperty,
				akaProperty,
				neeProperty))

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(person)
	defs.Add(name)

	var content bytes.Buffer
	rpt := reporting.NewTypeCatalogReport(defs, reporting.InlineTypes)

	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())
	golden.Assert(t, t.Name(), content.Bytes())
}

func Test_TypeCatalogReport_GivenInterface_ShowsExpectedDetails(t *testing.T) {
	t.Parallel()

	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	personName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")

	idFactory := astmodel.NewIdentifierFactory()
	f1 := test.NewFakeFunction("Hello", idFactory)
	f1.TypeReturned = astmodel.StringType

	f2 := test.NewFakeFunction("Goodbye", idFactory)
	f2.TypeReturned = astmodel.BoolType

	person := astmodel.MakeTypeDefinition(
		personName,
		astmodel.NewInterfaceType(f1, f2))

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(person)

	var content bytes.Buffer
	rpt := reporting.NewTypeCatalogReport(defs, reporting.InlineTypes, reporting.IncludeFunctions)

	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())
	golden.Assert(t, t.Name(), content.Bytes())
}

func Test_TypeCatalogReport_GivenObject_ShowsExpectedDetails(t *testing.T) {
	t.Parallel()

	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	personName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")

	personObj := astmodel.NewObjectType().
		WithProperties(
			test.FullNameProperty,
			test.FamilyNameProperty,
			test.KnownAsProperty)

	functionNames := []string{"Hello", "Goodbye"}
	for _, name := range functionNames {
		fn := test.NewFakeFunction(name, idFactory)
		personObj = personObj.WithFunction(fn)
	}

	person := astmodel.MakeTypeDefinition(personName, personObj)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(person)

	var content bytes.Buffer
	rpt := reporting.NewTypeCatalogReport(defs, reporting.InlineTypes, reporting.IncludeFunctions)

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
