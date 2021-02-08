/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bytes"
	"github.com/sebdah/goldie/v2"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

type StorageConversionPropertyTestCase struct {
	name             string
	receiverProperty *PropertyDefinition
	otherProperty    *PropertyDefinition
}

func CreateStorageConversionFunctionTestCases() []StorageConversionPropertyTestCase {
	requiredStringProperty := NewPropertyDefinition("name", "name", StringType)
	optionalStringProperty := NewPropertyDefinition("name", "name", NewOptionalType(StringType))
	requiredIntProperty := NewPropertyDefinition("age", "age", IntType)
	optionalIntProperty := NewPropertyDefinition("age", "age", NewOptionalType(IntType))

	arrayOfRequiredIntProperty := NewPropertyDefinition("scores", "scores", NewArrayType(IntType))
	arrayOfOptionalIntProperty := NewPropertyDefinition("scores", "scores", NewArrayType(NewOptionalType(IntType)))

	mapOfRequiredIntsProperty := NewPropertyDefinition("ratings", "ratings", NewMapType(StringType, IntType))
	mapOfOptionalIntsProperty := NewPropertyDefinition("ratings", "ratings", NewMapType(StringType, NewOptionalType(IntType)))

	nastyProperty := NewPropertyDefinition(
		"nasty",
		"nasty",
		NewMapType(
			StringType,
			NewArrayType(
				NewMapType(StringType, BoolType))))

	return []StorageConversionPropertyTestCase{
		{"SetStringFromString", requiredStringProperty, requiredStringProperty},
		{"SetStringFromOptionalString", requiredStringProperty, optionalStringProperty},
		{"SetOptionalStringFromString", optionalStringProperty, requiredStringProperty},
		{"SetOptionalStringFromOptionalString", optionalStringProperty, optionalStringProperty},

		{"SetIntFromInt", requiredIntProperty, requiredIntProperty},
		{"SetIntFromOptionalInt", requiredIntProperty, optionalIntProperty},

		{"SetArrayOfRequiredFromArrayOfRequired", arrayOfRequiredIntProperty, arrayOfRequiredIntProperty},
		{"SetArrayOfRequiredFromArrayOfOptional", arrayOfRequiredIntProperty, arrayOfOptionalIntProperty},
		{"SetArrayOfOptionalFromArrayOfRequired", arrayOfOptionalIntProperty, arrayOfRequiredIntProperty},

		{"SetMapOfRequiredFromMapOfRequired", mapOfRequiredIntsProperty, mapOfRequiredIntsProperty},
		{"SetMapOfRequiredFromMapOfOptional", mapOfRequiredIntsProperty, mapOfOptionalIntsProperty},
		{"SetMapOfOptionalFromMapOfRequired", mapOfOptionalIntsProperty, mapOfRequiredIntsProperty},

		{"NastyTest", nastyProperty, nastyProperty},
	}
}

func TestStorageConversionFunction_AsFuncForDirectConversions(t *testing.T) {
	for _, c := range CreateStorageConversionFunctionTestCases() {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			RunTestStorageConversionFunction_AsFunc(c, true, t)
		})
	}
}

func TestStorageConversionFunction_AsFuncForIndirectConversions(t *testing.T) {
	for _, c := range CreateStorageConversionFunctionTestCases() {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			RunTestStorageConversionFunction_AsFunc(c, false, t)
		})
	}
}

func RunTestStorageConversionFunction_AsFunc(c StorageConversionPropertyTestCase, direct bool, t *testing.T) {
	gm := NewGomegaWithT(t)

	idFactory := NewIdentifierFactory()
	vCurrent := makeTestLocalPackageReference("Verification", "vCurrent")
	vNext := makeTestLocalPackageReference("Verification", "vNext")
	vHub := makeTestLocalPackageReference("Verification", "vHub")

	receiverType := NewObjectType().WithProperty(c.receiverProperty)
	receiverDefinition := MakeTypeDefinition(
		MakeTypeName(vCurrent, "Person"),
		receiverType)

	hubTypeDefinition := MakeTypeDefinition(
		MakeTypeName(vHub, "Person"),
		NewObjectType().WithProperty(c.otherProperty))

	var intermediateTypeDefinition *TypeDefinition = nil
	if !direct {
		def := MakeTypeDefinition(
			MakeTypeName(vNext, "Person"),
			NewObjectType().WithProperty(c.otherProperty))
		intermediateTypeDefinition = &def
	}

	convertFrom, errs := NewStorageConversionFromFunction(receiverDefinition, hubTypeDefinition, intermediateTypeDefinition, idFactory)
	gm.Expect(errs).To(BeNil())

	convertTo, errs := NewStorageConversionToFunction(receiverDefinition, hubTypeDefinition, intermediateTypeDefinition, idFactory)
	gm.Expect(errs).To(BeNil())

	receiverDefinition = receiverDefinition.WithType(receiverType.WithFunction(convertFrom).WithFunction(convertTo))

	defs := []TypeDefinition{receiverDefinition}
	packages := make(map[PackageReference]*PackageDefinition)

	g := goldie.New(t)

	packageDefinition := NewPackageDefinition(vCurrent.Group(), vCurrent.PackageName(), "1")
	packageDefinition.AddDefinition(receiverDefinition)

	packages[vCurrent] = packageDefinition

	// put all definitions in one file, regardless.
	// the package reference isn't really used here.
	fileDef := NewFileDefinition(vCurrent, defs, packages)

	buf := &bytes.Buffer{}
	fileWriter := NewGoSourceFileWriter(fileDef)
	err := fileWriter.SaveToWriter(buf)
	if err != nil {
		t.Fatalf("could not generate file: %v", err)
	}

	var testName strings.Builder
	testName.WriteString(c.name)

	if direct {
		testName.WriteString("-Direct")
	} else {
		testName.WriteString("-ViaStaging")
	}

	g.Assert(t, testName.String(), buf.Bytes())
}
