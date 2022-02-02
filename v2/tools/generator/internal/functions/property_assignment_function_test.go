/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"

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
	vCurrent := test.MakeLocalPackageReference("verification", "vCurrent")
	vNext := test.MakeLocalPackageReference("verification", "vNext")

	// Custom Types
	alpha := astmodel.EnumValue{Identifier: "Alpha", Value: "alpha"}
	beta := astmodel.EnumValue{Identifier: "Beta", Value: "beta"}

	enumType := astmodel.NewEnumType(astmodel.StringType, alpha, beta)
	currentEnum := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vCurrent, "Bucket"), enumType)
	nextEnum := astmodel.MakeTypeDefinition(astmodel.MakeTypeName(vNext, "Bucket"), enumType)

	jsonObjectType := astmodel.NewMapType(astmodel.StringType, astmodel.JSONType)

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

	referenceProperty := astmodel.NewPropertyDefinition("Reference", "reference", astmodel.ResourceReferenceType)
	knownReferenceProperty := astmodel.NewPropertyDefinition("KnownReference", "known-reference", astmodel.KnownResourceReferenceType)
	jsonProperty := astmodel.NewPropertyDefinition("JSONBlob", "jsonBlob", astmodel.JSONType)
	optionalJSONProperty := astmodel.NewPropertyDefinition("JSONBlob", "jsonBlob", astmodel.NewOptionalType(astmodel.JSONType))
	jsonObjectProperty := astmodel.NewPropertyDefinition("JSONObject", "jsonObject", jsonObjectType)
	optionalJSONObjectProperty := astmodel.NewPropertyDefinition("JSONObject", "jsonObject", astmodel.NewOptionalType(jsonObjectType))

	// Some pretty contrived cases to catch shadowing bugs in conversion code for arrays and maps.
	// Array of object
	indexesProperty := astmodel.NewPropertyDefinition("Indexes", "indexes", astmodel.NewArrayType(currentRole.Name()))
	hubIndexesProperty := astmodel.NewPropertyDefinition("Indexes", "indexes", astmodel.NewArrayType(hubRole.Name()))

	// Map of object
	keysProperty := astmodel.NewPropertyDefinition("Keys", "keys", astmodel.NewMapType(astmodel.StringType, currentRole.Name()))
	hubKeysProperty := astmodel.NewPropertyDefinition("Keys", "keys", astmodel.NewMapType(astmodel.StringType, hubRole.Name()))

	// Array of optional object
	optionalIndexesProperty := astmodel.NewPropertyDefinition("Indexes", "indexes", astmodel.NewArrayType(astmodel.NewOptionalType(currentRole.Name())))

	// Map of optional object
	optionalKeysProperty := astmodel.NewPropertyDefinition("Keys", "keys", astmodel.NewMapType(astmodel.StringType, astmodel.NewOptionalType(currentRole.Name())))

	// Array of array of int
	arrayofArraysProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewArrayType(astmodel.NewArrayType(astmodel.IntType)))

	// Map of map of int
	mapOfMapsProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewMapType(astmodel.StringType, astmodel.NewMapType(astmodel.StringType, astmodel.IntType)))

	// Handcrafted impls
	sliceOfStringProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewArrayType(astmodel.StringType))
	mapOfStringToStringProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewMapType(astmodel.StringType, astmodel.StringType))

	bagProperty := astmodel.NewPropertyDefinition("propertyBag", "$propertyBag", astmodel.PropertyBagType)

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
		createPropertyAssignmentTest("SetOptionalIntFromOptionalInt", optionalIntProperty, optionalIntProperty),

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

		createPropertyAssignmentTest("CopyRequiredStringToPropertyBag", requiredStringProperty, bagProperty),
		createPropertyAssignmentTest("CopyOptionalStringToPropertyBag", optionalStringProperty, bagProperty),
		createPropertyAssignmentTest("CopyRequiredIntToPropertyBag", requiredIntProperty, bagProperty),
		createPropertyAssignmentTest("CopyOptionalIntToPropertyBag", optionalIntProperty, bagProperty),

		createPropertyAssignmentTest("CopyRequiredObjectToPropertyBag", requiredCurrentRoleProperty, bagProperty, currentRole),
		createPropertyAssignmentTest("CopyOptionalObjectToPropertyBag", optionalCurrentRoleProperty, bagProperty, currentRole),

		createPropertyAssignmentTest("CopySliceToPropertyBag", arrayOfRequiredIntProperty, bagProperty),
		createPropertyAssignmentTest("CopyMapToPropertyBag", mapOfRequiredIntsProperty, bagProperty),

		createPropertyAssignmentTest("ConvertSliceNamedIndexes", indexesProperty, hubIndexesProperty, currentRole, hubRole),
		createPropertyAssignmentTest("ConvertMapNamedKeys", keysProperty, hubKeysProperty, currentRole, hubRole),

		createPropertyAssignmentTest("ConvertSliceNamedIndexesOfOptionalObject", optionalIndexesProperty, optionalIndexesProperty, currentRole),
		createPropertyAssignmentTest("ConvertMapNamedKeysOfOptionalObject", optionalKeysProperty, optionalKeysProperty, currentRole),
		createPropertyAssignmentTest("SetSliceOfSlices", arrayofArraysProperty, arrayofArraysProperty),
		createPropertyAssignmentTest("SetMapOfMaps", mapOfMapsProperty, mapOfMapsProperty),

		createPropertyAssignmentTest("SetSliceOfString", sliceOfStringProperty, sliceOfStringProperty),
		createPropertyAssignmentTest("SetMapOfStringToString", mapOfStringToStringProperty, mapOfStringToStringProperty),
	}
}

func TestGolden_PropertyAssignmentFunction_AsFunc(t *testing.T) {
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

	conversionContext := conversions.NewPropertyConversionContext(c.types, idFactory, nil /* ObjectModelConfiguration*/)
	assignFrom, err := NewPropertyAssignmentFunction(c.currentObject, c.otherObject, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(BeNil())

	assignTo, err := NewPropertyAssignmentFunction(c.currentObject, c.otherObject, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(BeNil())

	receiverDefinition := c.currentObject.WithType(currentType.WithFunction(assignFrom).WithFunction(assignTo))

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, c.name, receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenPropertyBagPresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	person2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	person2021 := test.CreateObjectDefinition(
		test.Pkg2021,
		"Person",
		test.FullNameProperty,
		test.PropertyBagProperty)

	conversionContext := conversions.NewPropertyConversionContext(
		make(astmodel.Types), idFactory, nil /* ObjectModelConfiguration*/)
	assignFrom, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	assignTo, err := NewPropertyAssignmentFunction(person2020, person2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(person2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "PropertyBags", receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenTypeRenamed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	// We have a definition for Location
	location := test.CreateObjectDefinition(
		test.Pkg2020,
		"Location",
		test.FullAddressProperty,
		test.CityProperty)

	// ... which gets renamed to Venue in a later version
	venue := test.CreateObjectDefinition(
		test.Pkg2021,
		"Venue",
		test.FullAddressProperty,
		test.CityProperty)

	// The earlier version of Event has a property "Where" of type "Location" ...
	whereLocationProperty := astmodel.NewPropertyDefinition("Where", "where", location.Name())
	event2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Event",
		whereLocationProperty)

	// ... in the later version of Event, the property "Where" is now of type "Venue"
	whereVenueProperty := astmodel.NewPropertyDefinition("Where", "where", venue.Name())
	event2021 := test.CreateObjectDefinition(
		test.Pkg2021,
		"Event",
		whereVenueProperty)

	modelConfig := createConfigurationForRename(location.Name(), venue.Name())

	types := make(astmodel.Types)
	types.AddAll(location, venue)

	conversionContext := conversions.NewPropertyConversionContext(types, idFactory, modelConfig)

	assignFrom, err := NewPropertyAssignmentFunction(event2020, event2021, conversionContext, conversions.ConvertFrom)
	g.Expect(err).To(Succeed())

	assignTo, err := NewPropertyAssignmentFunction(event2020, event2021, conversionContext, conversions.ConvertTo)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(event2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	// The generated code should be using the type "Location" when referencing the earlier version of the property
	// "Where" and "Venue" for the later version. The types are visible in declarations of temporary variables,
	// and in the name of the Assign*() functions.
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "PropertyTypeRenamed", receiverDefinition)
}

// createConfigurationForRename builds up a new configuration for a particular desired rename
func createConfigurationForRename(originalName astmodel.TypeName, newName astmodel.TypeName) *config.ObjectModelConfiguration {
	group, version, _ := originalName.PackageReference.GroupVersion()

	typeConfig := config.NewTypeConfiguration(originalName.Name()).SetNameInNextVersion(newName.Name())
	versionConfig := config.NewVersionConfiguration(version).Add(typeConfig)
	groupConfig := config.NewGroupConfiguration(group).Add(versionConfig)
	return config.NewObjectModelConfiguration().Add(groupConfig)
}
