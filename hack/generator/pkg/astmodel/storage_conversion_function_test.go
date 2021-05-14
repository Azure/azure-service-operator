/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bytes"
	"testing"

	"github.com/sebdah/goldie/v2"

	. "github.com/onsi/gomega"
)

type StorageConversionPropertyTestCase struct {
	name          string
	currentObject TypeDefinition
	otherObject   TypeDefinition
	types         Types
}

func CreateStorageConversionFunctionTestCases() []*StorageConversionPropertyTestCase {

	vCurrent := makeTestLocalPackageReference("Verification", "vCurrent")
	vNext := makeTestLocalPackageReference("Verification", "vNext")

	alpha := EnumValue{Identifier: "Alpha", Value: "alpha"}
	beta := EnumValue{Identifier: "Beta", Value: "beta"}

	enumType := NewEnumType(StringType, alpha, beta)
	currentEnum := MakeTypeDefinition(MakeTypeName(vCurrent, "Bucket"), enumType)
	nextEnum := MakeTypeDefinition(MakeTypeName(vNext, "Bucket"), enumType)

	// These are aliases of primitive types
	currentAge := MakeTypeDefinition(MakeTypeName(vCurrent, "Age"), IntType)
	nextAge := MakeTypeDefinition(MakeTypeName(vNext, "Age"), IntType)

	requiredStringProperty := NewPropertyDefinition("name", "name", StringType)
	optionalStringProperty := NewPropertyDefinition("name", "name", NewOptionalType(StringType))
	requiredIntProperty := NewPropertyDefinition("age", "age", IntType)
	optionalIntProperty := NewPropertyDefinition("age", "age", NewOptionalType(IntType))

	arrayOfRequiredIntProperty := NewPropertyDefinition("scores", "scores", NewArrayType(IntType))
	arrayOfOptionalIntProperty := NewPropertyDefinition("scores", "scores", NewArrayType(NewOptionalType(IntType)))

	mapOfRequiredIntsProperty := NewPropertyDefinition("ratings", "ratings", NewMapType(StringType, IntType))
	mapOfOptionalIntsProperty := NewPropertyDefinition("ratings", "ratings", NewMapType(StringType, NewOptionalType(IntType)))

	requiredCurrentEnumProperty := NewPropertyDefinition("release", "release", currentEnum.name)
	requiredHubEnumProperty := NewPropertyDefinition("release", "release", nextEnum.name)
	optionalCurrentEnumProperty := NewPropertyDefinition("release", "release", NewOptionalType(currentEnum.name))
	optionalHubEnumProperty := NewPropertyDefinition("release", "release", NewOptionalType(nextEnum.name))

	roleType := NewObjectType().WithProperty(requiredStringProperty).WithProperty(arrayOfRequiredIntProperty)
	currentRole := MakeTypeDefinition(MakeTypeName(vCurrent, "Release"), roleType)
	hubRole := MakeTypeDefinition(MakeTypeName(vNext, "Release"), roleType)

	requiredCurrentRoleProperty := NewPropertyDefinition("role", "role", currentRole.Name())
	requiredHubRoleProperty := NewPropertyDefinition("role", "role", hubRole.Name())
	optionalCurrentRoleProperty := NewPropertyDefinition("role", "role", NewOptionalType(currentRole.Name()))
	optionalNextRoleProperty := NewPropertyDefinition("role", "role", NewOptionalType(hubRole.Name()))

	requiredSkuStringProperty := NewPropertyDefinition("sku", "sku", StringType)
	requiredNextSkuEnumProperty := NewPropertyDefinition("sku", "sku", nextEnum.name)
	optionalSkuStringProperty := NewPropertyDefinition("sku", "sku", NewOptionalType(StringType))
	optionalNextSkuEnumProperty := NewPropertyDefinition("sku", "sku", NewOptionalType(nextEnum.name))

	requiredPrimitiveAgeProperty := NewPropertyDefinition("age", "age", IntType)
	optionalPrimitiveAgeProperty := NewPropertyDefinition("age", "age", NewOptionalType(IntType))
	requiredCurrentAgeProperty := NewPropertyDefinition("age", "age", currentAge.name)
	requiredNextAgeProperty := NewPropertyDefinition("age", "age", nextAge.name)
	optionalCurrentAgeProperty := NewPropertyDefinition("age", "age", NewOptionalType(currentAge.name))
	optionalNextAgeProperty := NewPropertyDefinition("age", "age", NewOptionalType(nextAge.name))

	nastyProperty := NewPropertyDefinition(
		"nasty",
		"nasty",
		NewMapType(
			StringType,
			NewArrayType(
				NewMapType(StringType, BoolType))))

	createTest := func(
		name string,
		currentProperty *PropertyDefinition,
		hubProperty *PropertyDefinition,
		otherDefinitions ...TypeDefinition) *StorageConversionPropertyTestCase {

		currentType := NewObjectType().WithProperty(currentProperty)
		currentDefinition := MakeTypeDefinition(
			MakeTypeName(vCurrent, "Person"),
			currentType)

		hubType := NewObjectType().WithProperty(hubProperty)
		hubDefinition := MakeTypeDefinition(
			MakeTypeName(vNext, "Person"),
			hubType)

		types := make(Types)
		types.Add(currentDefinition)
		types.Add(hubDefinition)
		types.AddAll(otherDefinitions)

		return &StorageConversionPropertyTestCase{
			name:          name,
			currentObject: currentDefinition,
			otherObject:   hubDefinition,
			types:         types,
		}
	}

	return []*StorageConversionPropertyTestCase{
		createTest("SetStringFromString", requiredStringProperty, requiredStringProperty),
		createTest("SetStringFromOptionalString", requiredStringProperty, optionalStringProperty),
		createTest("SetOptionalStringFromString", optionalStringProperty, requiredStringProperty),
		createTest("SetOptionalStringFromOptionalString", optionalStringProperty, optionalStringProperty),

		createTest("SetIntFromInt", requiredIntProperty, requiredIntProperty),
		createTest("SetIntFromOptionalInt", requiredIntProperty, optionalIntProperty),

		createTest("SetArrayOfRequiredFromArrayOfRequired", arrayOfRequiredIntProperty, arrayOfRequiredIntProperty),
		createTest("SetArrayOfRequiredFromArrayOfOptional", arrayOfRequiredIntProperty, arrayOfOptionalIntProperty),
		createTest("SetArrayOfOptionalFromArrayOfRequired", arrayOfOptionalIntProperty, arrayOfRequiredIntProperty),

		createTest("SetMapOfRequiredFromMapOfRequired", mapOfRequiredIntsProperty, mapOfRequiredIntsProperty),
		createTest("SetMapOfRequiredFromMapOfOptional", mapOfRequiredIntsProperty, mapOfOptionalIntsProperty),
		createTest("SetMapOfOptionalFromMapOfRequired", mapOfOptionalIntsProperty, mapOfRequiredIntsProperty),

		createTest("NastyTest", nastyProperty, nastyProperty),

		createTest("ConvertBetweenRequiredEnumAndRequiredEnum", requiredCurrentEnumProperty, requiredHubEnumProperty, currentEnum, nextEnum),
		createTest("ConvertBetweenRequiredEnumAndOptionalEnum", requiredCurrentEnumProperty, optionalHubEnumProperty, currentEnum, nextEnum),
		createTest("ConvertBetweenOptionalEnumAndOptionalEnum", optionalCurrentEnumProperty, optionalHubEnumProperty, currentEnum, nextEnum),

		createTest("ConvertBetweenRequiredObjectAndRequiredObject", requiredCurrentRoleProperty, requiredHubRoleProperty, currentRole, hubRole),
		createTest("ConvertBetweenRequiredObjectAndOptionalObject", requiredCurrentRoleProperty, optionalNextRoleProperty, currentRole, hubRole),
		createTest("ConvertBetweenOptionalObjectAndOptionalObject", optionalCurrentRoleProperty, optionalNextRoleProperty, currentRole, hubRole),

		createTest("ConvertBetweenEnumAndBaseType", requiredSkuStringProperty, requiredNextSkuEnumProperty, nextEnum),
		createTest("ConvertBetweenEnumAndOptionalBaseType", optionalSkuStringProperty, requiredNextSkuEnumProperty, nextEnum),
		createTest("ConvertBetweenOptionalEnumAndBaseType", requiredSkuStringProperty, optionalNextSkuEnumProperty, nextEnum),
		createTest("ConvertBetweenOptionalEnumAndOptionalBaseType", optionalSkuStringProperty, optionalNextSkuEnumProperty, nextEnum),

		createTest("ConvertBetweenAliasAndAliasType", requiredCurrentAgeProperty, requiredNextAgeProperty, currentAge, nextAge),
		createTest("ConvertBetweenAliasAndOptionalAliasType", requiredCurrentAgeProperty, optionalNextAgeProperty, currentAge, nextAge),
		createTest("ConvertBetweenOptionalAliasAndOptionalAliasType", optionalCurrentAgeProperty, optionalNextAgeProperty, currentAge, nextAge),
		createTest("ConvertBetweenAliasAndBaseType", requiredCurrentAgeProperty, requiredPrimitiveAgeProperty, currentAge, nextAge),
		createTest("ConvertBetweenOptionalAliasAndBaseType", optionalCurrentAgeProperty, requiredPrimitiveAgeProperty, currentAge),
		createTest("ConvertBetweenOptionalAliasAndOptionalBaseType", optionalCurrentAgeProperty, optionalPrimitiveAgeProperty, currentAge),
	}
}

func TestStorageConversionFunction_AsFunc(t *testing.T) {
	for _, c := range CreateStorageConversionFunctionTestCases() {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			RunTestStorageConversionFunction_AsFunc(c, t)
		})
	}
}

func RunTestStorageConversionFunction_AsFunc(c *StorageConversionPropertyTestCase, t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := NewIdentifierFactory()

	conversionContext := NewStorageConversionContext(c.types)

	currentType, ok := AsObjectType(c.currentObject.Type())
	g.Expect(ok).To(BeTrue())

	convertFrom, errs := NewStorageConversionFromFunction(c.currentObject, c.otherObject, idFactory, conversionContext)
	g.Expect(errs).To(BeNil())

	convertTo, errs := NewStorageConversionToFunction(c.currentObject, c.otherObject, idFactory, conversionContext)
	g.Expect(errs).To(BeNil())

	receiverDefinition := c.currentObject.WithType(currentType.WithFunction(convertFrom).WithFunction(convertTo))

	defs := []TypeDefinition{receiverDefinition}
	packages := make(map[PackageReference]*PackageDefinition)

	currentPackage := receiverDefinition.Name().PackageReference.(LocalPackageReference)

	packageDefinition := NewPackageDefinition(currentPackage.Group(), currentPackage.PackageName(), "1")
	packageDefinition.AddDefinition(receiverDefinition)

	packages[currentPackage] = packageDefinition

	// put all definitions in one file, regardless.
	// the package reference isn't really used here.
	fileDef := NewFileDefinition(currentPackage, defs, packages)

	assertFileGeneratesExpectedCode(t, fileDef, c.name)
}

func assertFileGeneratesExpectedCode(t *testing.T, fileDef *FileDefinition, testName string) {
	g := goldie.New(t)

	buf := &bytes.Buffer{}
	fileWriter := NewGoSourceFileWriter(fileDef)
	err := fileWriter.SaveToWriter(buf)
	if err != nil {
		t.Fatalf("could not generate file: %v", err)
	}

	g.Assert(t, testName, buf.Bytes())
}
