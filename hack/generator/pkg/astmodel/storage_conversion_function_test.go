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
	nextObject    *TypeDefinition
	hubObject     TypeDefinition
	types         Types
}

func CreateStorageConversionFunctionTestCases() []*StorageConversionPropertyTestCase {

	vCurrent := makeTestLocalPackageReference("Verification", "vCurrent")
	vNext := makeTestLocalPackageReference("Verification", "vNext")
	vHub := makeTestLocalPackageReference("Verification", "vHub")

	alpha := EnumValue{Identifier: "Alpha", Value: "alpha"}
	beta := EnumValue{Identifier: "Beta", Value: "beta"}

	enumType := NewEnumType(StringType, alpha, beta)
	currentEnum := MakeTypeDefinition(MakeTypeName(vCurrent, "Bucket"), enumType)
	nextEnum := MakeTypeDefinition(MakeTypeName(vNext, "Container"), enumType)

	requiredStringProperty := NewPropertyDefinition("name", "name", StringType)
	optionalStringProperty := NewPropertyDefinition("name", "name", NewOptionalType(StringType))
	requiredIntProperty := NewPropertyDefinition("age", "age", IntType)
	optionalIntProperty := NewPropertyDefinition("age", "age", NewOptionalType(IntType))

	arrayOfRequiredIntProperty := NewPropertyDefinition("scores", "scores", NewArrayType(IntType))
	arrayOfOptionalIntProperty := NewPropertyDefinition("scores", "scores", NewArrayType(NewOptionalType(IntType)))

	mapOfRequiredIntsProperty := NewPropertyDefinition("ratings", "ratings", NewMapType(StringType, IntType))
	mapOfOptionalIntsProperty := NewPropertyDefinition("ratings", "ratings", NewMapType(StringType, NewOptionalType(IntType)))

	requiredCurrentEnumProperty := NewPropertyDefinition("release", "release", currentEnum.name)
	requiredNextEnumProperty := NewPropertyDefinition("release", "release", nextEnum.name)
	optionalCurrentEnumProperty := NewPropertyDefinition("release", "release", NewOptionalType(currentEnum.name))
	optionalNextEnumProperty := NewPropertyDefinition("release", "release", NewOptionalType(nextEnum.name))

	nastyProperty := NewPropertyDefinition(
		"nasty",
		"nasty",
		NewMapType(
			StringType,
			NewArrayType(
				NewMapType(StringType, BoolType))))

	testDirect := func(
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
			MakeTypeName(vHub, "Person"),
			hubType)

		types := make(Types)
		types.Add(currentDefinition)
		types.Add(hubDefinition)
		types.AddAll(otherDefinitions)

		return &StorageConversionPropertyTestCase{
			name:          name + "-Direct",
			currentObject: currentDefinition,
			hubObject:     hubDefinition,
			types:         types,
		}
	}

	testIndirect := func(name string,
		currentProperty *PropertyDefinition,
		nextProperty *PropertyDefinition,
		otherDefinitions ...TypeDefinition) *StorageConversionPropertyTestCase {

		result := testDirect(name, currentProperty, nextProperty, otherDefinitions...)

		nextType := NewObjectType().WithProperty(nextProperty)
		nextDefinition := MakeTypeDefinition(
			MakeTypeName(vNext, "Person"),
			nextType)

		result.name = name + "-ViaStaging"
		result.nextObject = &nextDefinition
		result.types.Add(nextDefinition)

		return result
	}

	return []*StorageConversionPropertyTestCase{
		testDirect("SetStringFromString", requiredStringProperty, requiredStringProperty),
		testDirect("SetStringFromOptionalString", requiredStringProperty, optionalStringProperty),
		testDirect("SetOptionalStringFromString", optionalStringProperty, requiredStringProperty),
		testDirect("SetOptionalStringFromOptionalString", optionalStringProperty, optionalStringProperty),

		testIndirect("SetStringFromString", requiredStringProperty, requiredStringProperty),
		testIndirect("SetStringFromOptionalString", requiredStringProperty, optionalStringProperty),
		testIndirect("SetOptionalStringFromString", optionalStringProperty, requiredStringProperty),
		testIndirect("SetOptionalStringFromOptionalString", optionalStringProperty, optionalStringProperty),

		testDirect("SetIntFromInt", requiredIntProperty, requiredIntProperty),
		testDirect("SetIntFromOptionalInt", requiredIntProperty, optionalIntProperty),

		testIndirect("SetIntFromInt", requiredIntProperty, requiredIntProperty),
		testIndirect("SetIntFromOptionalInt", requiredIntProperty, optionalIntProperty),

		testDirect("SetArrayOfRequiredFromArrayOfRequired", arrayOfRequiredIntProperty, arrayOfRequiredIntProperty),
		testDirect("SetArrayOfRequiredFromArrayOfOptional", arrayOfRequiredIntProperty, arrayOfOptionalIntProperty),
		testDirect("SetArrayOfOptionalFromArrayOfRequired", arrayOfOptionalIntProperty, arrayOfRequiredIntProperty),

		testIndirect("SetArrayOfRequiredFromArrayOfRequired", arrayOfRequiredIntProperty, arrayOfRequiredIntProperty),
		testIndirect("SetArrayOfRequiredFromArrayOfOptional", arrayOfRequiredIntProperty, arrayOfOptionalIntProperty),
		testIndirect("SetArrayOfOptionalFromArrayOfRequired", arrayOfOptionalIntProperty, arrayOfRequiredIntProperty),

		testDirect("SetMapOfRequiredFromMapOfRequired", mapOfRequiredIntsProperty, mapOfRequiredIntsProperty),
		testDirect("SetMapOfRequiredFromMapOfOptional", mapOfRequiredIntsProperty, mapOfOptionalIntsProperty),
		testDirect("SetMapOfOptionalFromMapOfRequired", mapOfOptionalIntsProperty, mapOfRequiredIntsProperty),

		testIndirect("SetMapOfRequiredFromMapOfRequired", mapOfRequiredIntsProperty, mapOfRequiredIntsProperty),
		testIndirect("SetMapOfRequiredFromMapOfOptional", mapOfRequiredIntsProperty, mapOfOptionalIntsProperty),
		testIndirect("SetMapOfOptionalFromMapOfRequired", mapOfOptionalIntsProperty, mapOfRequiredIntsProperty),

		testDirect("NastyTest", nastyProperty, nastyProperty),
		testIndirect("NastyTest", nastyProperty, nastyProperty),

		testDirect("SetRequiredEnumFromRequiredEnum", requiredCurrentEnumProperty, requiredNextEnumProperty, currentEnum, nextEnum),
		testDirect("SetRequiredEnumFromOptionalEnum", requiredCurrentEnumProperty, optionalNextEnumProperty, currentEnum, nextEnum),
		testDirect("SetOptionalEnumFromRequiredEnum", optionalCurrentEnumProperty, requiredNextEnumProperty, currentEnum, nextEnum),
		testDirect("SetOptionalEnumFromOptionalEnum", optionalCurrentEnumProperty, optionalNextEnumProperty, currentEnum, nextEnum),
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

	convertFrom, errs := NewStorageConversionFromFunction(c.currentObject, c.hubObject, c.nextObject, idFactory, conversionContext)
	g.Expect(errs).To(BeNil())

	convertTo, errs := NewStorageConversionToFunction(c.currentObject, c.hubObject, c.nextObject, idFactory, conversionContext)
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
