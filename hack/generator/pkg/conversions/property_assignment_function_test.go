/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"bytes"
	"testing"

	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"

	. "github.com/onsi/gomega"
)

type StorageConversionPropertyTestCase struct {
	name          string
	currentObject astmodel.TypeDefinition
	otherObject   astmodel.TypeDefinition
	types         astmodel.Types
}

func CreateStorageConversionFunctionTestCases() []*StorageConversionPropertyTestCase {

	vCurrent := makeTestLocalPackageReference("Verification", "vCurrent")
	vNext := makeTestLocalPackageReference("Verification", "vNext")

	alpha := astmodel.EnumValue{Identifier: "Alpha", Value: "alpha"}
	beta := astmodel.EnumValue{Identifier: "Beta", Value: "beta"}

	enumType := astmodel.NewEnumType(astmodel.StringType, alpha, beta)
	currentEnum := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vCurrent, "Bucket"), enumType)
	nextEnum := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vNext, "Bucket"), enumType)

	// These are aliases of primitive types
	currentAge := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vCurrent, "Age"), astmodel.IntType)
	nextAge := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vNext, "Age"), astmodel.IntType)

	requiredStringProperty := astmodel.NewPropertyDefinition("name", "name", astmodel.StringType)
	optionalStringProperty := astmodel.NewPropertyDefinition("name", "name", astmodel.NewOptionalType(astmodel.StringType))
	requiredIntProperty := astmodel.NewPropertyDefinition("age", "age", astmodel.IntType)
	optionalIntProperty := astmodel.NewPropertyDefinition("age", "age", astmodel.NewOptionalType(astmodel.IntType))

	arrayOfRequiredIntProperty := astmodel.NewPropertyDefinition("scores", "scores", astmodel.NewArrayType(astmodel.IntType))
	arrayOfOptionalIntProperty := astmodel.NewPropertyDefinition("scores", "scores", astmodel.NewArrayType(astmodel.NewOptionalType(astmodel.IntType)))

	mapOfRequiredIntsProperty := astmodel.NewPropertyDefinition("ratings", "ratings", astmodel.NewMapType(astmodel.StringType, astmodel.IntType))
	mapOfOptionalIntsProperty := astmodel.NewPropertyDefinition("ratings", "ratings", astmodel.NewMapType(astmodel.StringType, astmodel.NewOptionalType(astmodel.IntType)))

	requiredCurrentEnumProperty := astmodel.NewPropertyDefinition("release", "release", currentEnum.Name())
	requiredHubEnumProperty := astmodel.NewPropertyDefinition("release", "release", nextEnum.Name())
	optionalCurrentEnumProperty := astmodel.NewPropertyDefinition("release", "release", astmodel.NewOptionalType(currentEnum.Name()))
	optionalHubEnumProperty := astmodel.NewPropertyDefinition("release", "release", astmodel.NewOptionalType(nextEnum.Name()))

	roleType := astmodel.NewObjectType().WithProperty(requiredStringProperty).WithProperty(arrayOfRequiredIntProperty)
	currentRole := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vCurrent, "Release"), roleType)
	hubRole := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vNext, "Release"), roleType)

	requiredCurrentRoleProperty := astmodel.NewPropertyDefinition("role", "role", currentRole.Name())
	requiredHubRoleProperty := astmodel.NewPropertyDefinition("role", "role", hubRole.Name())
	optionalCurrentRoleProperty := astmodel.NewPropertyDefinition("role", "role", astmodel.NewOptionalType(currentRole.Name()))
	optionalNextRoleProperty := astmodel.NewPropertyDefinition("role", "role", astmodel.NewOptionalType(hubRole.Name()))

	requiredSkuStringProperty := astmodel.NewPropertyDefinition("sku", "sku", astmodel.StringType)
	requiredNextSkuEnumProperty := astmodel.NewPropertyDefinition("sku", "sku", nextEnum.Name())
	optionalSkuStringProperty := astmodel.NewPropertyDefinition("sku", "sku", astmodel.NewOptionalType(astmodel.StringType))
	optionalNextSkuEnumProperty := astmodel.NewPropertyDefinition("sku", "sku", astmodel.NewOptionalType(nextEnum.Name()))

	requiredPrimitiveAgeProperty := astmodel.NewPropertyDefinition("age", "age", astmodel.IntType)
	optionalPrimitiveAgeProperty := astmodel.NewPropertyDefinition("age", "age", astmodel.NewOptionalType(astmodel.IntType))
	requiredCurrentAgeProperty := astmodel.NewPropertyDefinition("age", "age", currentAge.Name())
	requiredNextAgeProperty := astmodel.NewPropertyDefinition("age", "age", nextAge.Name())
	optionalCurrentAgeProperty := astmodel.NewPropertyDefinition("age", "age", astmodel.NewOptionalType(currentAge.Name()))
	optionalNextAgeProperty := astmodel.NewPropertyDefinition("age", "age", astmodel.NewOptionalType(nextAge.Name()))

	nastyProperty := astmodel.NewPropertyDefinition(
		"nasty",
		"nasty",
		astmodel.NewMapType(
			astmodel.StringType,
			astmodel.NewArrayType(
				astmodel.NewMapType(astmodel.StringType, astmodel.BoolType))))

	createTest := func(
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

	idFactory := astmodel.NewIdentifierFactory()

	conversionContext := NewStorageConversionContext(c.types, idFactory)

	currentType, ok := astmodel.AsObjectType(c.currentObject.Type())
	g.Expect(ok).To(BeTrue())

	convertFrom, errs := NewPropertyAssignmentFromFunction(c.currentObject, c.otherObject, idFactory, conversionContext)
	g.Expect(errs).To(BeNil())

	convertTo, errs := NewPropertyAssignmentToFunction(c.currentObject, c.otherObject, idFactory, conversionContext)
	g.Expect(errs).To(BeNil())

	receiverDefinition := c.currentObject.WithType(currentType.WithFunction(convertFrom).WithFunction(convertTo))

	defs := []astmodel.TypeDefinition{receiverDefinition}
	packages := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)

	currentPackage := receiverDefinition.Name().PackageReference.(astmodel.LocalPackageReference)

	packageDefinition := astmodel.NewPackageDefinition(currentPackage.Group(), currentPackage.PackageName(), "1")
	packageDefinition.AddDefinition(receiverDefinition)

	packages[currentPackage] = packageDefinition

	// put all definitions in one file, regardless.
	// the package reference isn't really used here.
	fileDef := astmodel.NewFileDefinition(currentPackage, defs, packages)

	assertFileGeneratesExpectedCode(t, fileDef, c.name)
}

func assertFileGeneratesExpectedCode(t *testing.T, fileDef *astmodel.FileDefinition, testName string) {
	g := goldie.New(t)

	buf := &bytes.Buffer{}
	fileWriter := astmodel.NewGoSourceFileWriter(fileDef)
	err := fileWriter.SaveToWriter(buf)
	if err != nil {
		t.Fatalf("could not generate file: %v", err)
	}

	g.Assert(t, testName, buf.Bytes())
}
