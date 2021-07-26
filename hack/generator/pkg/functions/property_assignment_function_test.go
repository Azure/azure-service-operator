/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"

	. "github.com/onsi/gomega"
)

type StorageConversionPropertyTestCase struct {
	name          string
	currentObject astmodel.TypeDefinition
	otherObject   astmodel.TypeDefinition
	types         astmodel.Types
}

func CreatePropertyAssignmentFunctionTestCases() []*StorageConversionPropertyTestCase {
	// Package References
	vCurrent := test.MakeLocalPackageReference("Verification", "vCurrent")
	vNext := test.MakeLocalPackageReference("Verification", "vNext")

	// Custom Types
	alpha := astmodel.EnumValue{Identifier: "Alpha", Value: "alpha"}
	beta := astmodel.EnumValue{Identifier: "Beta", Value: "beta"}

	enumType := astmodel.NewEnumType(astmodel.StringType, alpha, beta)
	currentEnum := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vCurrent, "Bucket"), enumType)
	nextEnum := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vNext, "Bucket"), enumType)

	jsonObjectType := astmodel.NewMapType(astmodel.StringType, astmodel.JSONTypeName)

	// Aliases of primitive types
	currentAge := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vCurrent, "Age"), astmodel.IntType)
	nextAge := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vNext, "Age"), astmodel.IntType)

	requiredStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType)
	optionalStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.NewOptionalType(astmodel.StringType))
	requiredIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.IntType)
	optionalIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.NewOptionalType(astmodel.IntType))

	arrayOfRequiredIntProperty := astmodel.NewPropertyDefinition("Scores", "scores", astmodel.NewArrayType(astmodel.IntType))
	arrayOfOptionalIntProperty := astmodel.NewPropertyDefinition("Scores", "scores", astmodel.NewArrayType(astmodel.NewOptionalType(astmodel.IntType)))

	mapOfRequiredIntsProperty := astmodel.NewPropertyDefinition("Ratings", "ratings", astmodel.NewMapType(astmodel.StringType, astmodel.IntType))
	mapOfOptionalIntsProperty := astmodel.NewPropertyDefinition("Ratings", "ratings", astmodel.NewMapType(astmodel.StringType, astmodel.NewOptionalType(astmodel.IntType)))

	requiredCurrentEnumProperty := astmodel.NewPropertyDefinition("Release", "release", currentEnum.Name())
	requiredHubEnumProperty := astmodel.NewPropertyDefinition("Release", "release", nextEnum.Name())
	optionalCurrentEnumProperty := astmodel.NewPropertyDefinition("Release", "release", astmodel.NewOptionalType(currentEnum.Name()))
	optionalHubEnumProperty := astmodel.NewPropertyDefinition("Release", "release", astmodel.NewOptionalType(nextEnum.Name()))

	roleType := astmodel.NewObjectType().WithProperty(requiredStringProperty).WithProperty(arrayOfRequiredIntProperty)
	currentRole := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vCurrent, "Release"), roleType)
	hubRole := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vNext, "Release"), roleType)

	requiredCurrentRoleProperty := astmodel.NewPropertyDefinition("Role", "role", currentRole.Name())
	requiredHubRoleProperty := astmodel.NewPropertyDefinition("Role", "role", hubRole.Name())
	optionalCurrentRoleProperty := astmodel.NewPropertyDefinition("Role", "role", astmodel.NewOptionalType(currentRole.Name()))
	optionalNextRoleProperty := astmodel.NewPropertyDefinition("Role", "role", astmodel.NewOptionalType(hubRole.Name()))

	requiredSkuStringProperty := astmodel.NewPropertyDefinition("Sku", "sku", astmodel.StringType)
	requiredNextSkuEnumProperty := astmodel.NewPropertyDefinition("Sku", "sku", nextEnum.Name())
	optionalSkuStringProperty := astmodel.NewPropertyDefinition("Sku", "sku", astmodel.NewOptionalType(astmodel.StringType))
	optionalNextSkuEnumProperty := astmodel.NewPropertyDefinition("Sku", "sku", astmodel.NewOptionalType(nextEnum.Name()))

	requiredPrimitiveAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.IntType)
	optionalPrimitiveAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.NewOptionalType(astmodel.IntType))
	requiredCurrentAgeProperty := astmodel.NewPropertyDefinition("Age", "age", currentAge.Name())
	requiredNextAgeProperty := astmodel.NewPropertyDefinition("Age", "age", nextAge.Name())
	optionalCurrentAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.NewOptionalType(currentAge.Name()))
	optionalNextAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.NewOptionalType(nextAge.Name()))

	referenceProperty := astmodel.NewPropertyDefinition("Reference", "reference", astmodel.ResourceReferenceTypeName)
	knownReferenceProperty := astmodel.NewPropertyDefinition("KnownReference", "known-reference", astmodel.KnownResourceReferenceTypeName)
	jsonProperty := astmodel.NewPropertyDefinition("JSONBlob", "jsonBlob", astmodel.JSONTypeName)
	optionalJSONProperty := astmodel.NewPropertyDefinition("JSONBlob", "jsonBlob", astmodel.NewOptionalType(astmodel.JSONTypeName))
	jsonObjectProperty := astmodel.NewPropertyDefinition("JSONObject", "jsonObject", jsonObjectType)
	optionalJSONObjectProperty := astmodel.NewPropertyDefinition("JSONObject", "jsonObject", astmodel.NewOptionalType(jsonObjectType))

	idFactory := astmodel.NewIdentifierFactory()
	ageFunction := test.NewFakeFunction("Age", idFactory)
	ageFunction.TypeReturned = astmodel.IntType

	nastyProperty := astmodel.NewPropertyDefinition(
		"nasty",
		"nasty",
		astmodel.NewMapType(
			astmodel.StringType,
			astmodel.NewArrayType(
				astmodel.NewMapType(astmodel.StringType, astmodel.BoolType))))

	createPropertyAssignmentTest := func(
		name string,
		currentProperty *astmodel.PropertyDefinition,
		hubProperty *astmodel.PropertyDefinition,
		otherDefinitions ...astmodel.TypeDefinition) *StorageConversionPropertyTestCase {

		currentType := astmodel.NewObjectType().WithProperty(currentProperty)
		currentDefinition := astmodel.MakeTypeDefinition(
			astmodel.MakeTypeName(vCurrent, "Person"),
			currentType)

		hubType := astmodel.NewObjectType().WithProperty(hubProperty)
		hubDefinition := astmodel.MakeTypeDefinition(
			astmodel.MakeTypeName(vNext, "Person"),
			hubType)

		types := make(astmodel.Types)
		types.Add(currentDefinition)
		types.Add(hubDefinition)
		types.AddAll(otherDefinitions...)

		return &StorageConversionPropertyTestCase{
			name:          name,
			currentObject: currentDefinition,
			otherObject:   hubDefinition,
			types:         types,
		}
	}

	createFunctionAssignmentTest := func(
		name string,
		property *astmodel.PropertyDefinition,
		function astmodel.Function) *StorageConversionPropertyTestCase {

		currentType := astmodel.NewObjectType().WithFunction(function)
		currentDefinition := astmodel.MakeTypeDefinition(
			astmodel.MakeTypeName(vCurrent, "Person"),
			currentType)

		hubType := astmodel.NewObjectType().WithProperty(property)
		hubDefinition := astmodel.MakeTypeDefinition(
			astmodel.MakeTypeName(vNext, "Person"),
			hubType)

		types := make(astmodel.Types)
		types.Add(currentDefinition)
		types.Add(hubDefinition)

		return &StorageConversionPropertyTestCase{
			name:          name,
			currentObject: currentDefinition,
			otherObject:   hubDefinition,
			types:         types,
		}
	}

	return []*StorageConversionPropertyTestCase{
		createPropertyAssignmentTest("SetStringFromString", requiredStringProperty, requiredStringProperty),
		createPropertyAssignmentTest("SetStringFromOptionalString", requiredStringProperty, optionalStringProperty),
		createPropertyAssignmentTest("SetOptionalStringFromString", optionalStringProperty, requiredStringProperty),
		createPropertyAssignmentTest("SetOptionalStringFromOptionalString", optionalStringProperty, optionalStringProperty),

		createPropertyAssignmentTest("SetIntFromInt", requiredIntProperty, requiredIntProperty),
		createPropertyAssignmentTest("SetIntFromOptionalInt", requiredIntProperty, optionalIntProperty),

		createPropertyAssignmentTest("SetArrayOfRequiredFromArrayOfRequired", arrayOfRequiredIntProperty, arrayOfRequiredIntProperty),
		createPropertyAssignmentTest("SetArrayOfRequiredFromArrayOfOptional", arrayOfRequiredIntProperty, arrayOfOptionalIntProperty),
		createPropertyAssignmentTest("SetArrayOfOptionalFromArrayOfRequired", arrayOfOptionalIntProperty, arrayOfRequiredIntProperty),

		createPropertyAssignmentTest("SetMapOfRequiredFromMapOfRequired", mapOfRequiredIntsProperty, mapOfRequiredIntsProperty),
		createPropertyAssignmentTest("SetMapOfRequiredFromMapOfOptional", mapOfRequiredIntsProperty, mapOfOptionalIntsProperty),
		createPropertyAssignmentTest("SetMapOfOptionalFromMapOfRequired", mapOfOptionalIntsProperty, mapOfRequiredIntsProperty),

		createPropertyAssignmentTest("NastyTest", nastyProperty, nastyProperty),

		createPropertyAssignmentTest("ConvertBetweenRequiredEnumAndRequiredEnum", requiredCurrentEnumProperty, requiredHubEnumProperty, currentEnum, nextEnum),
		createPropertyAssignmentTest("ConvertBetweenRequiredEnumAndOptionalEnum", requiredCurrentEnumProperty, optionalHubEnumProperty, currentEnum, nextEnum),
		createPropertyAssignmentTest("ConvertBetweenOptionalEnumAndOptionalEnum", optionalCurrentEnumProperty, optionalHubEnumProperty, currentEnum, nextEnum),

		createPropertyAssignmentTest("ConvertBetweenRequiredObjectAndRequiredObject", requiredCurrentRoleProperty, requiredHubRoleProperty, currentRole, hubRole),
		createPropertyAssignmentTest("ConvertBetweenRequiredObjectAndOptionalObject", requiredCurrentRoleProperty, optionalNextRoleProperty, currentRole, hubRole),
		createPropertyAssignmentTest("ConvertBetweenOptionalObjectAndOptionalObject", optionalCurrentRoleProperty, optionalNextRoleProperty, currentRole, hubRole),

		createPropertyAssignmentTest("ConvertBetweenEnumAndBaseType", requiredSkuStringProperty, requiredNextSkuEnumProperty, nextEnum),
		createPropertyAssignmentTest("ConvertBetweenEnumAndOptionalBaseType", optionalSkuStringProperty, requiredNextSkuEnumProperty, nextEnum),
		createPropertyAssignmentTest("ConvertBetweenOptionalEnumAndBaseType", requiredSkuStringProperty, optionalNextSkuEnumProperty, nextEnum),
		createPropertyAssignmentTest("ConvertBetweenOptionalEnumAndOptionalBaseType", optionalSkuStringProperty, optionalNextSkuEnumProperty, nextEnum),

		createPropertyAssignmentTest("ConvertBetweenAliasAndAliasType", requiredCurrentAgeProperty, requiredNextAgeProperty, currentAge, nextAge),
		createPropertyAssignmentTest("ConvertBetweenAliasAndOptionalAliasType", requiredCurrentAgeProperty, optionalNextAgeProperty, currentAge, nextAge),
		createPropertyAssignmentTest("ConvertBetweenOptionalAliasAndOptionalAliasType", optionalCurrentAgeProperty, optionalNextAgeProperty, currentAge, nextAge),
		createPropertyAssignmentTest("ConvertBetweenAliasAndBaseType", requiredCurrentAgeProperty, requiredPrimitiveAgeProperty, currentAge, nextAge),
		createPropertyAssignmentTest("ConvertBetweenOptionalAliasAndBaseType", optionalCurrentAgeProperty, requiredPrimitiveAgeProperty, currentAge),
		createPropertyAssignmentTest("ConvertBetweenOptionalAliasAndOptionalBaseType", optionalCurrentAgeProperty, optionalPrimitiveAgeProperty, currentAge),

		createFunctionAssignmentTest("ReadFromFunctionIntoProperty", requiredIntProperty, ageFunction),
		createFunctionAssignmentTest("ReadFromFunctionIntoOptionalProperty", optionalIntProperty, ageFunction),

		createPropertyAssignmentTest("CopyReferenceProperty", referenceProperty, referenceProperty),
		createPropertyAssignmentTest("CopyKnownReferenceProperty", knownReferenceProperty, knownReferenceProperty),

		createPropertyAssignmentTest("CopyJSONProperty", jsonProperty, jsonProperty),
		createPropertyAssignmentTest("CopyJSONObjectProperty", jsonObjectProperty, jsonObjectProperty),
		createPropertyAssignmentTest("CopyOptionalJSONProperty", optionalJSONProperty, jsonProperty),
		createPropertyAssignmentTest("CopyJSONObjectProperty", optionalJSONObjectProperty, jsonObjectProperty),
	}
}

func TestPropertyAssignmentFunction_AsFunc(t *testing.T) {
	t.Parallel()

	for _, c := range CreatePropertyAssignmentFunctionTestCases() {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			runTestPropertyAssignmentFunction_AsFunc(c, t)
		})
	}
}

func runTestPropertyAssignmentFunction_AsFunc(c *StorageConversionPropertyTestCase, t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	currentType, ok := astmodel.AsObjectType(c.currentObject.Type())
	g.Expect(ok).To(BeTrue())

	conversionContext := conversions.NewPropertyConversionContext(c.types, idFactory)
	convertFrom, errs := NewPropertyAssignmentFunction(c.currentObject, c.otherObject, conversionContext, conversions.ConvertFrom)
	g.Expect(errs).To(BeNil())

	convertTo, errs := NewPropertyAssignmentFunction(c.currentObject, c.otherObject, conversionContext, conversions.ConvertTo)
	g.Expect(errs).To(BeNil())

	receiverDefinition := c.currentObject.WithType(currentType.WithFunction(convertFrom).WithFunction(convertTo))

	fileDef := test.CreateFileDefinition(receiverDefinition)
	test.AssertFileGeneratesExpectedCode(t, fileDef, c.name)
}
