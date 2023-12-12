/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"

	. "github.com/onsi/gomega"
)

type StorageConversionPropertyTestCase struct {
	name        string
	current     astmodel.TypeDefinition
	other       astmodel.TypeDefinition
	definitions astmodel.TypeDefinitionSet
}

// StorageConversionPropertyTestCaseFactory creates test cases for property conversion testing.
// We have a lot of test cases to cover, so we use a factory to create them.
type StorageConversionPropertyTestCaseFactory struct {
	cases          map[string]*StorageConversionPropertyTestCase
	currentPackage astmodel.LocalPackageReference
	nextPackage    astmodel.LocalPackageReference
}

// newStorageConversionPropertyTestCaseFactory creates a new test case factory for property conversion testing
func newStorageConversionPropertyTestCaseFactory() *StorageConversionPropertyTestCaseFactory {
	return &StorageConversionPropertyTestCaseFactory{
		cases:          make(map[string]*StorageConversionPropertyTestCase),
		currentPackage: test.MakeLocalPackageReference("verification", "vCurrent"),
		nextPackage:    test.MakeLocalPackageReference("verification", "vNext"),
	}
}

func (factory *StorageConversionPropertyTestCaseFactory) CreatePropertyAssignmentFunctionTestCases() map[string]*StorageConversionPropertyTestCase {
	// Test cases are created in batches to make it easier to see what's going on
	// In case of failures, troubleshoot earlier batches first, as it's very likely
	// that fixing those problems will fix later tests as well.
	factory.createPrimitiveTypeTestCases()
	factory.createCollectionTestCases()
	factory.createEnumTestCases()
	factory.createJSONTestCases()
	factory.createPropertyBagTestCases()
	factory.createObjectAssignmentTestCases()
	factory.createAssignmentViaFunctionTestCases()
	factory.createAssignmentViaHelperMethodsTestCases()
	factory.createPathologicalTestCases()

	return factory.cases
}

// createPrimitiveTypeTestCases creates test cases for conversion of primitive types (aka string, int, etc),
// optional variations, and aliases of those types.
func (factory *StorageConversionPropertyTestCaseFactory) createPrimitiveTypeTestCases() {
	requiredStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType)
	optionalStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.OptionalStringType)
	requiredIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.IntType)
	optionalIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.OptionalIntType)

	factory.createPropertyAssignmentTest("SetStringFromString", requiredStringProperty, requiredStringProperty)
	factory.createPropertyAssignmentTest("SetStringFromOptionalString", requiredStringProperty, optionalStringProperty)
	factory.createPropertyAssignmentTest("SetOptionalStringFromString", optionalStringProperty, requiredStringProperty)
	factory.createPropertyAssignmentTest("SetOptionalStringFromOptionalString", optionalStringProperty, optionalStringProperty)

	factory.createPropertyAssignmentTest("SetIntFromInt", requiredIntProperty, requiredIntProperty)
	factory.createPropertyAssignmentTest("SetIntFromOptionalInt", requiredIntProperty, optionalIntProperty)
	factory.createPropertyAssignmentTest("SetOptionalIntFromOptionalInt", optionalIntProperty, optionalIntProperty)

	// Aliases of primitive types
	currentAge := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Age"), astmodel.IntType)
	nextAge := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "Age"), astmodel.IntType)

	requiredPrimitiveAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.IntType)
	optionalPrimitiveAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.OptionalIntType)
	requiredCurrentAgeProperty := astmodel.NewPropertyDefinition("Age", "age", currentAge.Name())
	requiredNextAgeProperty := astmodel.NewPropertyDefinition("Age", "age", nextAge.Name())
	optionalCurrentAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.NewOptionalType(currentAge.Name()))
	optionalNextAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.NewOptionalType(nextAge.Name()))

	factory.createPropertyAssignmentTest("ConvertBetweenAliasAndAliasType", requiredCurrentAgeProperty, requiredNextAgeProperty, currentAge, nextAge)
	factory.createPropertyAssignmentTest("ConvertBetweenAliasAndOptionalAliasType", requiredCurrentAgeProperty, optionalNextAgeProperty, currentAge, nextAge)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalAliasAndOptionalAliasType", optionalCurrentAgeProperty, optionalNextAgeProperty, currentAge, nextAge)
	factory.createPropertyAssignmentTest("ConvertBetweenAliasAndBaseType", requiredCurrentAgeProperty, requiredPrimitiveAgeProperty, currentAge, nextAge)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalAliasAndBaseType", optionalCurrentAgeProperty, requiredPrimitiveAgeProperty, currentAge)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalAliasAndOptionalBaseType", optionalCurrentAgeProperty, optionalPrimitiveAgeProperty, currentAge)
}

// createCollectionTestCases creates test cases for conversion of collection types (aka array, map), and aliases of those types.
// We don't handle/test optional array/map aliases because they are always made non-optional by the FixOptionalCollectionAliases stage
func (factory *StorageConversionPropertyTestCaseFactory) createCollectionTestCases() {
	arrayOfRequiredIntProperty := astmodel.NewPropertyDefinition("Scores", "scores", astmodel.NewArrayType(astmodel.IntType))
	arrayOfOptionalIntProperty := astmodel.NewPropertyDefinition("Scores", "scores", astmodel.NewArrayType(astmodel.OptionalIntType))

	mapOfRequiredIntsProperty := astmodel.NewPropertyDefinition("Ratings", "ratings", astmodel.NewMapType(astmodel.StringType, astmodel.IntType))
	mapOfOptionalIntsProperty := astmodel.NewPropertyDefinition("Ratings", "ratings", astmodel.NewMapType(astmodel.StringType, astmodel.OptionalIntType))

	factory.createPropertyAssignmentTest("SetArrayOfRequiredFromArrayOfRequired", arrayOfRequiredIntProperty, arrayOfRequiredIntProperty)
	factory.createPropertyAssignmentTest("SetArrayOfRequiredFromArrayOfOptional", arrayOfRequiredIntProperty, arrayOfOptionalIntProperty)
	factory.createPropertyAssignmentTest("SetArrayOfOptionalFromArrayOfRequired", arrayOfOptionalIntProperty, arrayOfRequiredIntProperty)

	factory.createPropertyAssignmentTest("SetMapOfRequiredFromMapOfRequired", mapOfRequiredIntsProperty, mapOfRequiredIntsProperty)
	factory.createPropertyAssignmentTest("SetMapOfRequiredFromMapOfOptional", mapOfRequiredIntsProperty, mapOfOptionalIntsProperty)
	factory.createPropertyAssignmentTest("SetMapOfOptionalFromMapOfRequired", mapOfOptionalIntsProperty, mapOfRequiredIntsProperty)

	requiredStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType)
	roleType := astmodel.NewObjectType().WithProperty(requiredStringProperty).WithProperty(arrayOfRequiredIntProperty)
	currentRole := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Release"), roleType)
	hubRole := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "Release"), roleType)

	// Array of optional object
	optionalIndexesProperty := astmodel.NewPropertyDefinition("Indexes", "indexes", astmodel.NewArrayType(astmodel.NewOptionalType(currentRole.Name())))

	// Map of optional object
	optionalKeysProperty := astmodel.NewPropertyDefinition("Keys", "keys", astmodel.NewMapType(astmodel.StringType, astmodel.NewOptionalType(currentRole.Name())))

	// Array of array of int
	arrayofArraysProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewArrayType(astmodel.NewArrayType(astmodel.IntType)))

	// Map of map of int
	mapOfMapsProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewMapType(astmodel.StringType, astmodel.NewMapType(astmodel.StringType, astmodel.IntType)))

	factory.createPropertyAssignmentTest("ConvertSliceNamedIndexesOfOptionalObject", optionalIndexesProperty, optionalIndexesProperty, currentRole)
	factory.createPropertyAssignmentTest("ConvertMapNamedKeysOfOptionalObject", optionalKeysProperty, optionalKeysProperty, currentRole)
	factory.createPropertyAssignmentTest("SetSliceOfSlices", arrayofArraysProperty, arrayofArraysProperty)
	factory.createPropertyAssignmentTest("SetMapOfMaps", mapOfMapsProperty, mapOfMapsProperty)

	// Aliases of collection types
	currentPhoneNumbers := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "PhoneNumbers"), astmodel.NewValidatedType(astmodel.NewArrayType(astmodel.StringType), astmodel.ArrayValidations{}))
	nextPhoneNumbers := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "PhoneNumbers"), astmodel.NewValidatedType(astmodel.NewArrayType(astmodel.StringType), astmodel.ArrayValidations{}))
	currentAddresses := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Addresses"), astmodel.NewValidatedType(astmodel.NewMapType(astmodel.StringType, astmodel.StringType), nil))
	nextAddresses := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "Addresses"), astmodel.NewValidatedType(astmodel.NewMapType(astmodel.StringType, astmodel.StringType), nil))

	requiredCurrentPhoneNumbersProperty := astmodel.NewPropertyDefinition("PhoneNumbers", "phoneNumbers", currentPhoneNumbers.Name())
	requiredNextPhoneNumbersProperty := astmodel.NewPropertyDefinition("PhoneNumbers", "phoneNumbers", nextPhoneNumbers.Name())
	requiredCurrentPhoneNumbersBaseTypeProperty := astmodel.NewPropertyDefinition("PhoneNumbers", "phoneNumbers", astmodel.NewArrayType(astmodel.StringType))
	requiredCurrentAddressesProperty := astmodel.NewPropertyDefinition("Addresses", "addresses", currentAddresses.Name())
	requiredNextAddressesProperty := astmodel.NewPropertyDefinition("Addresses", "addresses", nextAddresses.Name())
	requiredCurrentAddressesBaseTypeProperty := astmodel.NewPropertyDefinition("Addresses", "addresses", astmodel.NewMapType(astmodel.StringType, astmodel.StringType))

	factory.createPropertyAssignmentTest("ConvertBetweenArrayAliasAndArrayAlias", requiredCurrentPhoneNumbersProperty, requiredNextPhoneNumbersProperty, currentPhoneNumbers, nextPhoneNumbers)
	factory.createPropertyAssignmentTest("ConvertBetweenArrayAliasAndBaseType", requiredCurrentPhoneNumbersBaseTypeProperty, requiredNextPhoneNumbersProperty, nextPhoneNumbers)
	factory.createPropertyAssignmentTest("ConvertBetweenMapAliasAndMapAlias", requiredCurrentAddressesProperty, requiredNextAddressesProperty, currentAddresses, nextAddresses)
	factory.createPropertyAssignmentTest("ConvertBetweenMapAliasAndBaseType", requiredCurrentAddressesBaseTypeProperty, requiredNextAddressesProperty, nextAddresses)

	// Some pretty contrived cases to catch shadowing bugs in conversion code for arrays and maps.

	// Array of object
	indexesProperty := astmodel.NewPropertyDefinition("Indexes", "indexes", astmodel.NewArrayType(currentRole.Name()))
	hubIndexesProperty := astmodel.NewPropertyDefinition("Indexes", "indexes", astmodel.NewArrayType(hubRole.Name()))

	// Map of object
	keysProperty := astmodel.NewPropertyDefinition("Keys", "keys", astmodel.NewMapType(astmodel.StringType, currentRole.Name()))
	hubKeysProperty := astmodel.NewPropertyDefinition("Keys", "keys", astmodel.NewMapType(astmodel.StringType, hubRole.Name()))

	factory.createPropertyAssignmentTest("ConvertSliceNamedIndexes", indexesProperty, hubIndexesProperty, currentRole, hubRole)
	factory.createPropertyAssignmentTest("ConvertMapNamedKeys", keysProperty, hubKeysProperty, currentRole, hubRole)
}

// createEnumTestCases creates test cases for conversion of enum types
func (factory *StorageConversionPropertyTestCaseFactory) createEnumTestCases() {
	// Values for the enum
	alpha := astmodel.MakeEnumValue("Alpha", "alpha")
	beta := astmodel.MakeEnumValue("Beta", "beta")

	// Create the enum itself
	enumType := astmodel.NewEnumType(astmodel.StringType, alpha, beta)

	// Two enum based type definitions
	currentEnum := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Bucket"), enumType)
	nextEnum := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "Bucket"), enumType)

	requiredCurrentEnumProperty := astmodel.NewPropertyDefinition("Release", "release", currentEnum.Name())
	requiredHubEnumProperty := astmodel.NewPropertyDefinition("Release", "release", nextEnum.Name())
	optionalCurrentEnumProperty := astmodel.NewPropertyDefinition("Release", "release", astmodel.NewOptionalType(currentEnum.Name()))
	optionalHubEnumProperty := astmodel.NewPropertyDefinition("Release", "release", astmodel.NewOptionalType(nextEnum.Name()))

	factory.createPropertyAssignmentTest("ConvertBetweenRequiredEnumAndRequiredEnum", requiredCurrentEnumProperty, requiredHubEnumProperty, currentEnum, nextEnum)
	factory.createPropertyAssignmentTest("ConvertBetweenRequiredEnumAndOptionalEnum", requiredCurrentEnumProperty, optionalHubEnumProperty, currentEnum, nextEnum)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalEnumAndOptionalEnum", optionalCurrentEnumProperty, optionalHubEnumProperty, currentEnum, nextEnum)

	requiredSkuStringProperty := astmodel.NewPropertyDefinition("Sku", "sku", astmodel.StringType)
	requiredNextSkuEnumProperty := astmodel.NewPropertyDefinition("Sku", "sku", nextEnum.Name())
	optionalSkuStringProperty := astmodel.NewPropertyDefinition("Sku", "sku", astmodel.OptionalStringType)
	optionalNextSkuEnumProperty := astmodel.NewPropertyDefinition("Sku", "sku", astmodel.NewOptionalType(nextEnum.Name()))

	factory.createPropertyAssignmentTest("ConvertBetweenEnumAndBaseType", requiredSkuStringProperty, requiredNextSkuEnumProperty, nextEnum)
	factory.createPropertyAssignmentTest("ConvertBetweenEnumAndOptionalBaseType", optionalSkuStringProperty, requiredNextSkuEnumProperty, nextEnum)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalEnumAndBaseType", requiredSkuStringProperty, optionalNextSkuEnumProperty, nextEnum)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalEnumAndOptionalBaseType", optionalSkuStringProperty, optionalNextSkuEnumProperty, nextEnum)
}

// createPropertyBagTestCases creates test cases for conversion values written-to and read-from property bags
func (factory *StorageConversionPropertyTestCaseFactory) createPropertyBagTestCases() {
	bagProperty := astmodel.NewPropertyDefinition("propertyBag", "$propertyBag", astmodel.PropertyBagType)

	requiredStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType)
	optionalStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.OptionalStringType)
	requiredIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.IntType)
	optionalIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.OptionalIntType)

	arrayOfRequiredIntProperty := astmodel.NewPropertyDefinition("Scores", "scores", astmodel.NewArrayType(astmodel.IntType))

	roleType := astmodel.NewObjectType().WithProperty(requiredStringProperty).WithProperty(arrayOfRequiredIntProperty)
	currentRole := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Release"), roleType)

	requiredCurrentRoleProperty := astmodel.NewPropertyDefinition("Role", "role", currentRole.Name())
	optionalCurrentRoleProperty := astmodel.NewPropertyDefinition("Role", "role", astmodel.NewOptionalType(currentRole.Name()))

	mapOfRequiredIntsProperty := astmodel.NewPropertyDefinition("Ratings", "ratings", astmodel.NewMapType(astmodel.StringType, astmodel.IntType))

	factory.createPropertyAssignmentTest("CopyRequiredStringToPropertyBag", requiredStringProperty, bagProperty)
	factory.createPropertyAssignmentTest("CopyOptionalStringToPropertyBag", optionalStringProperty, bagProperty)
	factory.createPropertyAssignmentTest("CopyRequiredIntToPropertyBag", requiredIntProperty, bagProperty)
	factory.createPropertyAssignmentTest("CopyOptionalIntToPropertyBag", optionalIntProperty, bagProperty)

	factory.createPropertyAssignmentTest("CopyRequiredObjectToPropertyBag", requiredCurrentRoleProperty, bagProperty, currentRole)
	factory.createPropertyAssignmentTest("CopyOptionalObjectToPropertyBag", optionalCurrentRoleProperty, bagProperty, currentRole)

	factory.createPropertyAssignmentTest("CopySliceToPropertyBag", arrayOfRequiredIntProperty, bagProperty)
	factory.createPropertyAssignmentTest("CopyMapToPropertyBag", mapOfRequiredIntsProperty, bagProperty)
}

// createJSONTestCases creates test cases for conversion of JSON types
func (factory *StorageConversionPropertyTestCaseFactory) createJSONTestCases() {
	jsonObjectType := astmodel.NewMapType(astmodel.StringType, astmodel.JSONType)

	jsonProperty := astmodel.NewPropertyDefinition("JSONBlob", "jsonBlob", astmodel.JSONType)
	optionalJSONProperty := astmodel.NewPropertyDefinition("JSONBlob", "jsonBlob", astmodel.NewOptionalType(astmodel.JSONType))
	jsonObjectProperty := astmodel.NewPropertyDefinition("JSONObject", "jsonObject", jsonObjectType)
	optionalJSONObjectProperty := astmodel.NewPropertyDefinition("JSONObject", "jsonObject", astmodel.NewOptionalType(jsonObjectType))

	factory.createPropertyAssignmentTest("CopyJSONProperty", jsonProperty, jsonProperty)
	factory.createPropertyAssignmentTest("CopyJSONObjectProperty", jsonObjectProperty, jsonObjectProperty)
	factory.createPropertyAssignmentTest("CopyOptionalJSONProperty", optionalJSONProperty, jsonProperty)
	factory.createPropertyAssignmentTest("CopyJSONObjectProperty", optionalJSONObjectProperty, jsonObjectProperty)
}

// createObjectAssignmentTestCases creates test cases for conversion of object types and optional variations.
// We don't handle/test Object Aliases here because they are always removed by the RemoveTypeAliases pipeline stage.
func (factory *StorageConversionPropertyTestCaseFactory) createObjectAssignmentTestCases() {
	requiredStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType)
	arrayOfRequiredIntProperty := astmodel.NewPropertyDefinition("Scores", "scores", astmodel.NewArrayType(astmodel.IntType))

	roleType := astmodel.NewObjectType().WithProperty(requiredStringProperty).WithProperty(arrayOfRequiredIntProperty)
	currentRole := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Release"), roleType)
	hubRole := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "Release"), roleType)

	requiredCurrentRoleProperty := astmodel.NewPropertyDefinition("Role", "role", currentRole.Name())
	requiredHubRoleProperty := astmodel.NewPropertyDefinition("Role", "role", hubRole.Name())
	optionalCurrentRoleProperty := astmodel.NewPropertyDefinition("Role", "role", astmodel.NewOptionalType(currentRole.Name()))
	optionalNextRoleProperty := astmodel.NewPropertyDefinition("Role", "role", astmodel.NewOptionalType(hubRole.Name()))

	factory.createPropertyAssignmentTest("ConvertBetweenRequiredObjectAndRequiredObject", requiredCurrentRoleProperty, requiredHubRoleProperty, currentRole, hubRole)
	factory.createPropertyAssignmentTest("ConvertBetweenRequiredObjectAndOptionalObject", requiredCurrentRoleProperty, optionalNextRoleProperty, currentRole, hubRole)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalObjectAndOptionalObject", optionalCurrentRoleProperty, optionalNextRoleProperty, currentRole, hubRole)
}

// createAssignmentViaFunctionTestCases creates test cases where reading the value is handled
// by a provided function, along with optional variations.
func (factory *StorageConversionPropertyTestCaseFactory) createAssignmentViaFunctionTestCases() {
	idFactory := astmodel.NewIdentifierFactory()
	ageFunction := test.NewFakeFunction("Age", idFactory)
	ageFunction.TypeReturned = astmodel.IntType

	requiredIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.IntType)
	optionalIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.OptionalIntType)

	factory.createFunctionAssignmentTest("ReadFromFunctionIntoProperty", requiredIntProperty, ageFunction)
	factory.createFunctionAssignmentTest("ReadFromFunctionIntoOptionalProperty", optionalIntProperty, ageFunction)
}

// createPathologicalTestCases creates test cases designed to be nasty to handle, as a check that the recursive
// approach to breaking down the assignment problem works as expected.
func (factory *StorageConversionPropertyTestCaseFactory) createPathologicalTestCases() {
	nastyProperty := astmodel.NewPropertyDefinition(
		"nasty",
		"nasty",
		astmodel.NewMapType(
			astmodel.StringType,
			astmodel.NewArrayType(
				astmodel.NewMapType(astmodel.StringType, astmodel.BoolType))))

	factory.createPropertyAssignmentTest("NastyTest", nastyProperty, nastyProperty)
}

// createAssignmentViaHelperMethodsTestCases creates test cases where the assignment is handled by known
// helper methods, mostly from the genruntime package
func (factory *StorageConversionPropertyTestCaseFactory) createAssignmentViaHelperMethodsTestCases() {
	referenceProperty := astmodel.NewPropertyDefinition("Reference", "reference", astmodel.ResourceReferenceType)
	knownReferenceProperty := astmodel.NewPropertyDefinition("KnownReference", "known-reference", astmodel.KnownResourceReferenceType)

	// Handcrafted impls
	sliceOfStringProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewArrayType(astmodel.StringType))
	mapOfStringToStringProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewMapType(astmodel.StringType, astmodel.StringType))

	factory.createPropertyAssignmentTest("CopyReferenceProperty", referenceProperty, referenceProperty)
	factory.createPropertyAssignmentTest("CopyKnownReferenceProperty", knownReferenceProperty, knownReferenceProperty)

	factory.createPropertyAssignmentTest("SetSliceOfString", sliceOfStringProperty, sliceOfStringProperty)
	factory.createPropertyAssignmentTest("SetMapOfStringToString", mapOfStringToStringProperty, mapOfStringToStringProperty)
}

func (factory *StorageConversionPropertyTestCaseFactory) createPropertyAssignmentTest(
	name string,
	currentProperty *astmodel.PropertyDefinition,
	hubProperty *astmodel.PropertyDefinition,
	otherDefinitions ...astmodel.TypeDefinition,
) {
	currentType := astmodel.NewObjectType().WithProperty(currentProperty)
	currentDefinition := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(factory.currentPackage, "Person"),
		currentType)

	hubType := astmodel.NewObjectType().WithProperty(hubProperty)
	hubDefinition := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(factory.nextPackage, "Person"),
		hubType)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(currentDefinition)
	defs.Add(hubDefinition)
	defs.AddAll(otherDefinitions...)

	factory.addCase(&StorageConversionPropertyTestCase{
		name:        name,
		current:     currentDefinition,
		other:       hubDefinition,
		definitions: defs,
	})
}

func (factory *StorageConversionPropertyTestCaseFactory) createFunctionAssignmentTest(
	name string,
	property *astmodel.PropertyDefinition,
	function astmodel.Function,
) {
	currentType := astmodel.NewObjectType().WithFunction(function)
	currentDefinition := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(factory.currentPackage, "Person"),
		currentType)

	hubType := astmodel.NewObjectType().WithProperty(property)
	hubDefinition := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(factory.nextPackage, "Person"),
		hubType)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(currentDefinition)
	defs.Add(hubDefinition)

	factory.addCase(&StorageConversionPropertyTestCase{
		name:        name,
		current:     currentDefinition,
		other:       hubDefinition,
		definitions: defs,
	})
}

func (factory *StorageConversionPropertyTestCaseFactory) addCase(c *StorageConversionPropertyTestCase) {
	factory.cases[c.name] = c
}

func TestGolden_PropertyAssignmentFunction_AsFunc(t *testing.T) {
	t.Parallel()

	factory := newStorageConversionPropertyTestCaseFactory()
	for n, c := range factory.CreatePropertyAssignmentFunctionTestCases() {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			runTestPropertyAssignmentFunction_AsFunc(c, t)
		})
	}
}

func runTestPropertyAssignmentFunction_AsFunc(c *StorageConversionPropertyTestCase, t *testing.T) {
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	currentType, ok := astmodel.AsObjectType(c.current.Type())
	g.Expect(ok).To(BeTrue())

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, c.definitions, idFactory)
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(c.current, c.other, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(BeNil())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(c.current, c.other, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(BeNil())

	receiverDefinition := c.current.WithType(currentType.WithFunction(assignFrom).WithFunction(assignTo))

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

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, make(astmodel.TypeDefinitionSet), idFactory)
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2021, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2021, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
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

	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			location.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.NameInNextVersion.Set(venue.Name().Name())
				return nil
			})).
		To(Succeed())

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(location, venue)

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, defs, idFactory).
		WithConfiguration(omc)

	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(event2020, event2021, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(event2020, event2021, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(event2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	// The generated code should be using the type "Location" when referencing the earlier version of the property
	// "Where", and "Venue" for the later version. The types are visible in declarations of temporary variables,
	// and in the name of the Assign*() functions.
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "PropertyTypeRenamed", receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenPropertyTypeHasIntermediateVersions_GeneratesExpectedCode(t *testing.T) {
	// Ensure that property assignment functions correctly make use of an intermediate type when
	// the required type conversion is not directly possible for a property.
	//
	// We create two versions of Person (v2020 and v2022), each with a `Residence` property of type `Location`.
	// A v2020.Location can't be directly converted to (or from) a v2022.Location, but only by using a v2021.Location
	// as an intermediate.
	//
	// +----------------+                          +----------------+
	// | v2020.Person   | -----------------------> | v2022.Person   |
	// +----------------+                          +----------------+
	//       |                                           |
	//       v                                           v
	// +----------------+    +----------------+    +----------------+
	// | v2020.Location | -> | v2021.Location | -> | v2022.Location |
	// +----------------+    +----------------+    +----------------+
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	// Arrange - create three different versions of Location - in packages v2020, v2021, and v2022
	location2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Location",
		test.FullAddressProperty)

	location2021 := test.CreateObjectDefinition(
		test.Pkg2021,
		"Location",
		test.FullAddressProperty)

	location2022 := test.CreateObjectDefinition(
		test.Pkg2022,
		"Location",
		test.FullAddressProperty)

	// Arrange - create two different version of Person - in packages v2020, and v2022.
	// Each has a property "Residence" of type "Location",
	// using the local version of "Location"
	person2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2020.Name()))

	person2022 := test.CreateObjectDefinition(
		test.Pkg2022,
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2022.Name()))

	// Arrange - create the conversion graph between all these object versions
	definitions := make(astmodel.TypeDefinitionSet)
	definitions.AddAll(location2020, location2021, location2022)
	definitions.AddAll(person2020, person2022)

	cfg := config.NewObjectModelConfiguration()
	builder := storage.NewConversionGraphBuilder(cfg, "v")
	builder.Add(person2020.Name(), person2022.Name())
	builder.Add(location2020.Name(), location2021.Name(), location2022.Name())

	graph, err := builder.Build()
	g.Expect(err).To(BeNil())

	// Arrange - create the conversion context, including the conversion graph
	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, definitions, idFactory).
		WithConversionGraph(graph)

	// Act - create both AssignTo and AssignFrom between the two versions of Person
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2022, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2022, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(person2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	// Assert that the generated code correctly converts using the intermediate version of Location
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "SharedObject", receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenPropertyTypeHasMultipleIntermediateVersions_GeneratesExpectedCode(t *testing.T) {
	// Ensure that property assignment functions correctly make use of multiple intermediate types when
	// the required type conversion is not directly possible for a property.
	//
	// We create two versions of Person (v2020 and v2022), each with a `Residence` property of type `Location`.
	// A v2020.Location can't be directly converted to (or from) a v2022.Location, but only by using all the intermediate
	// types of Location in turn.
	//
	// +----------------+                                                                                  +----------------+
	// |  v2020.Person  | -------------------------------------------------------------------------------> |  v2022.Person  |
	// +----------------+                                                                                  +----------------+
	//       |                                                                                                  |
	//       v                                                                                                  v
	// +----------------+    +--------------------+    +--------------------+    +--------------------+    +----------------+
	// | v2020.Location | -> | v20210301.Location | -> | v20210306.Location | -> | v20210312.Location | -> | v2022.Location |
	// +----------------+    +--------------------+    +--------------------+    +--------------------+    +----------------+
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	// Arrange - create five different versions of Location - in packages v2020, v202101, v202106, v202112, and v2022
	location2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Location",
		test.FullAddressProperty)

	location202101 := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20210301"),
		"Location",
		test.FullAddressProperty)

	location202106 := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20210306"),
		"Location",
		test.FullAddressProperty)

	location202112 := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20210312"),
		"Location",
		test.FullAddressProperty)

	location2022 := test.CreateObjectDefinition(
		test.Pkg2022,
		"Location",
		test.FullAddressProperty)

	// Arrange - create two different version of Person - in packages v2020, and v2022.
	person2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2020.Name()))

	person2022 := test.CreateObjectDefinition(
		test.Pkg2022,
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2022.Name()))

	definitions := make(astmodel.TypeDefinitionSet)
	definitions.AddAll(location2020, location202101, location202106, location202112, location2022)
	definitions.AddAll(person2020, person2022)

	// Arrange - create the conversion graph between all these object versions
	cfg := config.NewObjectModelConfiguration()
	builder := storage.NewConversionGraphBuilder(cfg, "v")
	builder.Add(
		person2020.Name(),
		person2022.Name(),
		location2020.Name(),
		location202101.Name(),
		location202106.Name(),
		location202112.Name(),
		location2022.Name())

	graph, err := builder.Build()
	g.Expect(err).To(BeNil())

	// Arrange - create the conversion context, including the conversion graph
	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, definitions, idFactory).
		WithConversionGraph(graph)

	// Act - create both AssignTo and AssignFrom between the two versions of Person
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2022, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2022, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(person2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	// Assert - the generated code should correctly convert using all the intermediate versions of Location
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "SharedObjectMultiple", receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenPropertyTypeVersionsAreNotInline_GeneratesExpectedCode(t *testing.T) {
	// Ensure that property assignment functions correctly make use of multiple intermediate types when
	// the two types are not inline with each other in the version graph.
	//
	// Under normal circumstances, you can always start from the 'earlier' type and follow links in the conversion
	// graph forward until you find the 'later' type. However, if the two types are not inline with each other,
	// the conversion instead needs to be split into two parts, based around a pivot point that's visible to both types.
	//
	// We create two versions of Person, v20200601preview and v20210601preview, each with a `Residence` property of type
	// `Location`. The conversion graph for `Location` looks like this:
	//
	// +--------------------+    +---------------------------+    +--------------------+    +---------------------------+    +--------------------+
	// | v20200101.Location | <- | v20200601preview.Location |    | v20210101.Location | <- | v20210601preview.Location |    | v20220101.Location |
	// +--------------------+    +---------------------------+    +--------------------+    +---------------------------+    +--------------------+
	//                  |   			   				              ^            |                                             ^
	//   				+---------------------------------------------+            +---------------------------------------------+
	//
	// Starting from v20200601preview.Location and following the conversion graph, we encounter v20200101.Location,
	// then v20210101.Location, and finally v20220101.Location. We never encounter v20210601preview.Location, so
	// a direct conversion is impossible.
	//
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	// Arrange - create five different versions of Location - in packages v2020, v2020p, v2021, v2021p, and v2022
	location2020 := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20200101"),
		"Location",
		test.FullAddressProperty)

	location2020p := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20200601preview"),
		"Location",
		test.FullAddressProperty)

	location2021 := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20210101"),
		"Location",
		test.FullAddressProperty)

	location2021p := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20210601preview"),
		"Location",
		test.FullAddressProperty)

	location2022 := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20220101"),
		"Location",
		test.FullAddressProperty)

	// Arrange - create two different version of Person - in packages v2020p, and v2022p.
	person2020p := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20200601preview"),
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2020p.Name()))

	person2021p := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20210601preview"),
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2021p.Name()))

	definitions := make(astmodel.TypeDefinitionSet)
	definitions.AddAll(location2020, location2020p, location2021, location2021p, location2022)
	definitions.AddAll(person2020p, person2021p)

	// Arrange - create the conversion graph between all these object versions
	cfg := config.NewObjectModelConfiguration()
	builder := storage.NewConversionGraphBuilder(cfg, "v")
	builder.Add(
		person2020p.Name(),
		person2021p.Name(),
		location2020.Name(),
		location2020p.Name(),
		location2021.Name(),
		location2021p.Name(),
		location2022.Name())

	graph, err := builder.Build()
	g.Expect(err).To(BeNil())

	s, err := graph.String(test.Group, "Location")
	g.Expect(err).To(BeNil())
	t.Log(s)

	// Arrange - create the conversion context, including the conversion graph
	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, definitions, idFactory).
		WithConversionGraph(graph)

	// Act - create both AssignTo and AssignFrom between the two versions of Person
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(person2020p, person2021p, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(person2020p, person2021p, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(person2020p, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	// Assert - the generated code should correctly convert using all the intermediate versions of Location
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "SharedObjectNotInline", receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenOverrideInterfacePresent(t *testing.T) {
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

	overrideInterfaceName := astmodel.MakeInternalTypeName(test.Pkg2020, "personAssignable")

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, make(astmodel.TypeDefinitionSet), idFactory)
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2021, conversions.ConvertFrom)
	assignFromBuilder.UseAugmentationInterface(overrideInterfaceName)

	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2021, conversions.ConvertTo)
	assignToBuilder.UseAugmentationInterface(overrideInterfaceName)

	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(person2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "OverrideInterface", receiverDefinition)
}
