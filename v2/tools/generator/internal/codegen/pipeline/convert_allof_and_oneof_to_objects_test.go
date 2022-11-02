/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"

	. "github.com/onsi/gomega"
)

func makeSynth(definitions ...astmodel.TypeDefinition) synthesizer {
	defs := make(astmodel.TypeDefinitionSet)
	for _, d := range definitions {
		defs.Add(d)
	}

	return newSynthesizer(defs, astmodel.NewIdentifierFactory())
}

var (
	mapStringInterface = astmodel.NewStringMapType(astmodel.AnyType)
	mapStringString    = astmodel.NewStringMapType(astmodel.StringType)
	mapInterfaceString = astmodel.NewMapType(astmodel.AnyType, astmodel.StringType)
)

var emptyObject = astmodel.NewObjectType()

func defineEnum(strings ...string) astmodel.Type {
	values := make([]astmodel.EnumValue, 0, len(strings))
	for _, value := range strings {
		values = append(values, astmodel.MakeEnumValue(value, value))
	}

	return astmodel.NewEnumType(
		astmodel.StringType,
		values...)
}

// any type merged with AnyType is just the type
func TestMergeWithAny(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	synth := makeSynth()
	g.Expect(synth.intersectTypes(astmodel.StringType, astmodel.AnyType)).To(Equal(astmodel.StringType))
	g.Expect(synth.intersectTypes(astmodel.AnyType, astmodel.StringType)).To(Equal(astmodel.StringType))
}

// merging maps is merging their keys and values
func TestMergeMaps(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	synth := makeSynth()
	g.Expect(synth.intersectTypes(mapStringInterface, mapStringString)).To(Equal(mapStringString))
	g.Expect(synth.intersectTypes(mapStringString, mapStringInterface)).To(Equal(mapStringString))
	g.Expect(synth.intersectTypes(mapInterfaceString, mapStringInterface)).To(Equal(mapStringString))
	g.Expect(synth.intersectTypes(mapStringInterface, mapInterfaceString)).To(Equal(mapStringString))
}

// merging a map with string keys with an empty object results in the map
func TestMergeMapEmptyObject(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	synth := makeSynth()
	g.Expect(synth.intersectTypes(mapStringString, emptyObject)).To(Equal(mapStringString))
	g.Expect(synth.intersectTypes(emptyObject, mapStringString)).To(Equal(mapStringString))
}

// merging a map with an object puts the map into 'AdditionalProperties'
func TestMergeMapObject(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	newMap := astmodel.NewMapType(astmodel.StringType, astmodel.FloatType)
	oneProp := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("x", "x", astmodel.IntType),
	)

	expected := oneProp.WithProperties(
		astmodel.NewPropertyDefinition(
			astmodel.AdditionalPropertiesPropertyName,
			astmodel.AdditionalPropertiesJsonName,
			newMap),
	)

	synth := makeSynth()
	g.Expect(synth.intersectTypes(newMap, oneProp)).To(Equal(expected))
	g.Expect(synth.intersectTypes(oneProp, newMap)).To(Equal(expected))
}

// merging two objects results in the union of their properties
func TestMergeObjectObject(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	propX := astmodel.NewPropertyDefinition("x", "x", astmodel.IntType)
	obj1 := astmodel.NewObjectType().WithProperties(propX)

	propY := astmodel.NewPropertyDefinition("y", "y", astmodel.FloatType)
	obj2 := astmodel.NewObjectType().WithProperties(propY)

	expected := astmodel.NewObjectType().WithProperties(propX, propY)

	synth := makeSynth()
	g.Expect(synth.intersectTypes(obj1, obj2)).To(Equal(expected))
	g.Expect(synth.intersectTypes(obj2, obj1)).To(Equal(expected))
}

// merging two objects preserves isResource
func TestMergeObjectIsResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	propX := astmodel.NewPropertyDefinition("x", "x", astmodel.IntType)
	obj1 := astmodel.NewObjectType().WithProperties(propX).WithIsResource(true)

	propY := astmodel.NewPropertyDefinition("y", "y", astmodel.FloatType)
	obj2 := astmodel.NewObjectType().WithProperties(propY)

	expected := astmodel.NewObjectType().WithProperties(propX, propY).WithIsResource(true)

	synth := makeSynth()
	g.Expect(synth.intersectTypes(obj1, obj2)).To(Equal(expected))
	g.Expect(synth.intersectTypes(obj2, obj1)).To(Equal(expected))
}

// merging two enums results in the intersection of their values
func TestMergeEnumEnum(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	enum1 := defineEnum("x", "y", "z", "g")
	enum2 := defineEnum("g", "a", "b", "c")

	expected := defineEnum("g")

	synth := makeSynth()
	g.Expect(synth.intersectTypes(enum1, enum2)).To(Equal(expected))
	g.Expect(synth.intersectTypes(enum2, enum1)).To(Equal(expected))
}

// cannot merge enums with different base types
func TestMergeBadEnums(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	enum := defineEnum("a", "b")
	enumInt := astmodel.NewEnumType(astmodel.IntType)

	var err error

	synth := makeSynth()
	_, err = synth.intersectTypes(enum, enumInt)
	g.Expect(err.Error()).To(ContainSubstring("differing base types"))

	_, err = synth.intersectTypes(enumInt, enum)
	g.Expect(err.Error()).To(ContainSubstring("differing base types"))
}

// merging an enum with its base type results in the enum
func TestMergeEnumBaseType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	enum := defineEnum("a", "b")

	synth := makeSynth()
	g.Expect(synth.intersectTypes(enum, astmodel.StringType)).To(Equal(enum))
	g.Expect(synth.intersectTypes(astmodel.StringType, enum)).To(Equal(enum))
}

// cannot merge an enum with another non-base type
func TestMergeEnumWrongBaseType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	enum := defineEnum("a", "b")

	var err error

	synth := makeSynth()
	_, err = synth.intersectTypes(enum, astmodel.IntType)
	g.Expect(err.Error()).To(ContainSubstring("don't know how to merge enum type"))

	_, err = synth.intersectTypes(astmodel.IntType, enum)
	g.Expect(err.Error()).To(ContainSubstring("don't know how to merge enum type"))
}

// merging two optionals merges their contents
func TestMergeOptionalOptional(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	enum1 := astmodel.NewOptionalType(defineEnum("a", "b"))
	enum2 := astmodel.NewOptionalType(defineEnum("b", "c"))

	expected := astmodel.NewOptionalType(defineEnum("b"))

	synth := makeSynth()
	g.Expect(synth.intersectTypes(enum1, enum2)).To(Equal(expected))
	g.Expect(synth.intersectTypes(enum2, enum1)).To(Equal(expected))
}

// merging an optional with something else that it can be merged with results in that result
// TODO: dubious?
func TestMergeOptionalEnum(t *testing.T) {
	// this feels a bit wrong, but it seems to be expected in real life specs
	t.Parallel()
	g := NewGomegaWithT(t)

	enum1 := defineEnum("a", "b")
	enum2 := astmodel.NewOptionalType(defineEnum("b", "c"))

	expected := defineEnum("b")

	synth := makeSynth()
	g.Expect(synth.intersectTypes(enum1, enum2)).To(Equal(expected))
	g.Expect(synth.intersectTypes(enum2, enum1)).To(Equal(expected))
}

// merging objects with common properties merges the types of those properties
func TestMergeObjectObjectCommonProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	obj1 := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("x", "x", defineEnum("a", "b")))

	obj2 := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("x", "x", defineEnum("b", "c")))

	expected := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("x", "x", defineEnum("b")))

	synth := makeSynth()
	g.Expect(synth.intersectTypes(obj1, obj2)).To(Equal(expected))
	g.Expect(synth.intersectTypes(obj2, obj1)).To(Equal(expected))
}

// merging resources merges their spec & status
func TestMergeResourceResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	r1 := astmodel.NewResourceType(defineEnum("a", "b"), defineEnum("x", "y"))
	r2 := astmodel.NewResourceType(defineEnum("b", "c"), defineEnum("y", "z"))

	expected := astmodel.NewResourceType(defineEnum("b"), defineEnum("y"))

	synth := makeSynth()
	g.Expect(synth.intersectTypes(r1, r2)).To(Equal(expected))
	g.Expect(synth.intersectTypes(r2, r1)).To(Equal(expected))
}

func TestMergeResourceMissingStatus(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	r1 := astmodel.NewResourceType(defineEnum("a", "b"), defineEnum("x", "y"))
	r2 := astmodel.NewResourceType(defineEnum("b", "c"), nil)

	expected := astmodel.NewResourceType(defineEnum("b"), defineEnum("x", "y"))

	synth := makeSynth()
	g.Expect(synth.intersectTypes(r1, r2)).To(Equal(expected))
	g.Expect(synth.intersectTypes(r2, r1)).To(Equal(expected))
}

func TestMergeValidated(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	maxLen := int64(32)
	validatedString := astmodel.NewValidatedType(
		astmodel.StringType,
		astmodel.StringValidations{
			MaxLength: &maxLen,
		})

	synth := makeSynth()
	g.Expect(synth.intersectTypes(validatedString, astmodel.StringType)).To(Equal(validatedString))
	g.Expect(synth.intersectTypes(astmodel.StringType, validatedString)).To(Equal(validatedString))
}

func TestMergeValidatedOfOptional(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	maxLen := int64(32)
	validatedOptionalString := astmodel.NewValidatedType(
		astmodel.OptionalStringType,
		astmodel.StringValidations{
			MaxLength: &maxLen,
		})

	validatedString := astmodel.NewValidatedType(
		astmodel.StringType,
		astmodel.StringValidations{
			MaxLength: &maxLen,
		})

	synth := makeSynth()
	g.Expect(synth.intersectTypes(validatedOptionalString, astmodel.StringType)).To(Equal(validatedString))
	g.Expect(synth.intersectTypes(astmodel.StringType, validatedOptionalString)).To(Equal(validatedString))
}

func TestMergeResourceWithOtherDependsOnSpecVsStatus(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	r := astmodel.NewResourceType(defineEnum("a", "b"), defineEnum("c", "d"))
	e := defineEnum("b", "c")

	synth := makeSynth()
	synth.specOrStatus = chooseSpec
	{
		expected := astmodel.NewResourceType(defineEnum("b"), defineEnum("c", "d"))
		g.Expect(synth.intersectTypes(r, e)).To(Equal(expected))
		g.Expect(synth.intersectTypes(e, r)).To(Equal(expected))
	}

	synth.specOrStatus = chooseStatus
	{
		expected := astmodel.NewResourceType(defineEnum("a", "b"), defineEnum("c"))
		g.Expect(synth.intersectTypes(r, e)).To(Equal(expected))
		g.Expect(synth.intersectTypes(e, r)).To(Equal(expected))
	}

	synth.specOrStatus = ""
}

// merging a oneOf with a type that is in the oneOf results in that type
func TestMergeOneOf(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	oneOf := astmodel.NewOneOfType("oneOf", astmodel.IntType, astmodel.StringType, astmodel.BoolType)

	synth := makeSynth()
	g.Expect(synth.intersectTypes(oneOf, astmodel.BoolType)).To(Equal(astmodel.BoolType))
	g.Expect(synth.intersectTypes(astmodel.BoolType, oneOf)).To(Equal(astmodel.BoolType))
}

// wooh!
func TestMergeOneOfEnum(t *testing.T) {
	// this is based on a real-world example
	// base type is defined like:
	//
	// 	ResourceBase { Location: oneOf{ string, enum('Global', 'USWest', 'USSouth', etc...) }
	//
	// derived type is defined like:
	//
	// 	Derived: allOf{ ResourceBase, { Location: enum('Global') } }
	//
	// which then reduces to:
	//
	// 	Derived: { Location: allOf{ oneOf{ string, enum('Global', ...) }, enum('Global') } }
	//
	// the Location type should be reduced to:
	//
	// 	Derived: { Location: enum('Global') }
	t.Parallel()
	g := NewGomegaWithT(t)

	oneOf := astmodel.NewOneOfType(
		"oneOf",
		astmodel.StringType,
		defineEnum("a", "b", "c"),
	)

	enum := defineEnum("b")

	expected := defineEnum("b")

	synth := makeSynth()
	g.Expect(synth.intersectTypes(oneOf, enum)).To(Equal(expected))
	g.Expect(synth.intersectTypes(enum, oneOf)).To(Equal(expected))
}

func TestOneOfResourceSpec(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	r := astmodel.NewResourceType(astmodel.StringType, astmodel.IntType)
	oneOf := astmodel.NewOneOfType("oneOf", astmodel.BoolType, r)

	synth := makeSynth()
	synth.specOrStatus = chooseSpec

	expected := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("Bool0", "bool0", astmodel.BoolType).
			MakeTypeOptional().WithDescription("Mutually exclusive with all other properties"),
		astmodel.NewPropertyDefinition("Resource1", "resource1", r).
			MakeTypeOptional().WithDescription("Mutually exclusive with all other properties"),
	)

	result, err := synth.oneOfToObject(oneOf)
	g.Expect(err).To(BeNil())

	result, ok := astmodel.AsObjectType(result)
	g.Expect(ok).To(BeTrue())
	result = astmodel.NewObjectType().WithProperties(result.(*astmodel.ObjectType).Properties().AsSlice()...)
	g.Expect(result).To(Equal(expected))
}

func TestCommonUppercasedSuffix(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(commonUppercasedSuffix("GoodDog", "HappyDog")).To(Equal("Dog"))
	g.Expect(commonUppercasedSuffix("Dog", "HappyDog")).To(Equal("Dog"))
	g.Expect(commonUppercasedSuffix("GoodDog", "Dog")).To(Equal("Dog"))
	g.Expect(commonUppercasedSuffix("abc", "bc")).To(Equal(""))
	g.Expect(commonUppercasedSuffix("bc", "abc")).To(Equal(""))
	g.Expect(commonUppercasedSuffix("ab1", "ac1")).To(Equal(""))
	g.Expect(commonUppercasedSuffix("1", "1")).To(Equal(""))
	g.Expect(commonUppercasedSuffix("", "1")).To(Equal(""))
	g.Expect(commonUppercasedSuffix("1", "")).To(Equal(""))
	g.Expect(commonUppercasedSuffix("PPP", "PP")).To(Equal("PP"))
	g.Expect(commonUppercasedSuffix("PP", "PPP")).To(Equal("PP"))
	g.Expect(commonUppercasedSuffix("", "")).To(Equal(""))
	g.Expect(commonUppercasedSuffix("xX", "X")).To(Equal("X"))
	g.Expect(commonUppercasedSuffix("X", "xX")).To(Equal("X"))
	g.Expect(commonUppercasedSuffix("xX", "xX")).To(Equal("X"))
	g.Expect(commonUppercasedSuffix("X", "X")).To(Equal("X"))
}

func TestSimplifyPropNamesDoesNotCreateEmptyNames(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	names := []propertyNames{
		{golang: "SuffixOnly"},
		{golang: "NotSuffixOnly"},
	}

	newNames := simplifyPropNames(names)

	// simplifyPropNames will do nothing,
	// because trimming the common suffix would result in an empty name
	g.Expect(newNames).To(Equal(names))
}

func TestSynthesizerOneOfObject_GivenOneOfUsingTypeNames_ReturnsExpectedObject(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	commonProperties := test.CreateObjectType(
		test.FamilyNameProperty,
		test.FullNameProperty)

	parent := createTestLeafOneOfDefinition(
		"Parent",
		"Senior",
		"Maxima",
		commonProperties,
		test.PostalAddress2021,
		test.ResidentialAddress2021)

	child := createTestLeafOneOfDefinition(
		"Child",
		"Junior",
		"Minima",
		commonProperties,
		test.KnownAsProperty)

	person := createTestRootOneOfDefinition(
		"Person",
		"Kind",
		// Use names to reference leaves
		parent.Name(),
		child.Name())

	synth := makeSynth(parent, child, person)

	oneOf := person.Type().(*astmodel.OneOfType)

	actual, err := synth.oneOfToObject(oneOf)
	g.Expect(err).To(BeNil())

	// Expect actual to have a property for each OneOf Option
	test.AssertPropertyCount(t, actual, 2)
	test.AssertPropertyExistsWithType(t, actual, "Parent", astmodel.NewOptionalType(parent.Name()))
	test.AssertPropertyExistsWithType(t, actual, "Child", astmodel.NewOptionalType(child.Name()))
}

func TestSynthesizerOneOfObject_GivenOneOfUsingNames_ReturnsExpectedObject(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	commonProperties := test.CreateObjectType(
		test.FamilyNameProperty,
		test.FullNameProperty)

	parent := createTestLeafOneOfDefinition(
		"Parent",
		"Senior",
		"Maxima",
		commonProperties,
		test.PostalAddress2021,
		test.ResidentialAddress2021)

	child := createTestLeafOneOfDefinition(
		"Child",
		"Junior", // no swagger name
		"Minima",
		commonProperties,
		test.KnownAsProperty)

	person := createTestRootOneOfDefinition(
		"Person",
		"Kind",
		// Use bodies to ensure we're not using the type names
		parent.Type(),
		child.Type())

	synth := makeSynth(parent, child, person)

	oneOf := person.Type().(*astmodel.OneOfType)

	actual, err := synth.oneOfToObject(oneOf)
	g.Expect(err).To(BeNil())

	// Expect actual to have a property for each OneOf Option
	test.AssertPropertyCount(t, actual, 2)
	test.AssertPropertyExists(t, actual, "Senior")
	test.AssertPropertyExists(t, actual, "Junior")
}

func TestSynthesizerOneOfObject_GivenOneOfUsingDiscriminatorValues_ReturnsExpectedObject(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	commonProperties := test.CreateObjectType(
		test.FamilyNameProperty,
		test.FullNameProperty)

	parent := createTestLeafOneOfDefinition(
		"Parent",
		"", // no swagger name
		"Maxima",
		commonProperties,
		test.PostalAddress2021,
		test.ResidentialAddress2021)

	child := createTestLeafOneOfDefinition(
		"Child",
		"", // no swagger name
		"Minima",
		commonProperties,
		test.KnownAsProperty)

	person := createTestRootOneOfDefinition(
		"Person",
		"Kind",
		// Use bodies to ensure we're not using the type names
		parent.Type(),
		child.Type())

	synth := makeSynth(parent, child, person)

	oneOf := person.Type().(*astmodel.OneOfType)

	actual, err := synth.oneOfToObject(oneOf)
	g.Expect(err).ToNot(HaveOccurred())

	// Expect actual to have a property for each OneOf Option
	test.AssertPropertyCount(t, actual, 2)
	test.AssertPropertyExists(t, actual, "Maxima")
	test.AssertPropertyExists(t, actual, "Minima")
}

var (
	olympianProperties = test.CreateObjectType(
		test.FullNameProperty,
		test.KnownAsProperty)

	zeus = createTestLeafOneOfDefinition(
		"Zeus",
		"zeus",
		"zeus",
		olympianProperties,
		astmodel.NewPropertyDefinition("LightningBolts", "lightningBolts", astmodel.IntType))

	demeter = createTestLeafOneOfDefinition(
		"Demeter",
		"demeter",
		"demeter",
		olympianProperties,
		astmodel.NewPropertyDefinition("Crops", "crops", astmodel.IntType))

	poscidon = createTestLeafOneOfDefinition(
		"Poscidon",
		"poscidon",
		"poscidon",
		olympianProperties,
		astmodel.NewPropertyDefinition("Tsunamis", "tsunamis", astmodel.IntType))

	hades = createTestLeafOneOfDefinition(
		"Hades",
		"hades",
		"hades",
		olympianProperties,
		astmodel.NewPropertyDefinition("Souls", "souls", astmodel.IntType))

	olympian = createTestRootOneOfDefinition(
		"Olympian",
		"name",
		zeus.Name(),
		demeter.Name(),
		poscidon.Name(),
		hades.Name())
)

func TestGolden_Synthesizer_WhenSynthesizingRootBeforeLeaves_ReturnsExpectedResults(t *testing.T) {
	t.Parallel()

	// Test definitions with root first
	RunSynthesizerTestInOrder(t, olympian, zeus, demeter, poscidon, hades)
}

func TestGolden_Synthesizer_WhenSynthesizingLeavesBeforeRoot_ReturnsExpectedResults(t *testing.T) {
	t.Parallel()

	// Test definitions with root last
	RunSynthesizerTestInOrder(t, zeus, demeter, poscidon, hades, olympian)
}

func RunSynthesizerTestInOrder(t *testing.T, defs ...astmodel.TypeDefinition) {
	g := NewGomegaWithT(t)

	synth := makeSynth(defs...)
	visitor := createVisitorForSynthesizer(synth)

	// Visit everything in order
	updatedDefs := make(astmodel.TypeDefinitionSet)
	for _, def := range defs {
		def, err := visitor.VisitDefinition(def, chooseSpec)
		g.Expect(err).To(BeNil())

		updatedDefs.Add(def)
	}

	// Combine the results
	finalDefs := updatedDefs.OverlayWith(synth.updatedDefs)

	// Check on the final shape
	for _, def := range finalDefs {
		test.AssertDefinitionHasExpectedShape(
			t,
			def.Name().Name(),
			def)
	}
}

func createTestRootOneOfDefinition(
	name string,
	discriminatorProperty string,
	leaves ...astmodel.Type,
) astmodel.TypeDefinition {
	oneOf := astmodel.NewOneOfType(name, leaves...).
		WithDiscriminatorProperty(discriminatorProperty)
	return astmodel.MakeTypeDefinition(astmodel.MakeTypeName(test.Pkg2020, name), oneOf)
}

func createTestLeafOneOfDefinition(
	typeName string,
	swaggerName string,
	discriminator string,
	commonProperties *astmodel.ObjectType,
	properties ...*astmodel.PropertyDefinition,
) astmodel.TypeDefinition {
	obj := astmodel.NewObjectType().
		WithProperties(properties...)
	body := astmodel.NewOneOfType(swaggerName, obj, commonProperties).
		WithDiscriminatorValue(discriminator)
	return astmodel.MakeTypeDefinition(astmodel.MakeTypeName(test.Pkg2020, typeName), body)
}

func Test_ConversionWithNestedAllOfs_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Testing a scenario found during manual testing
	webtestsResource := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "WebtestsResource"),
		astmodel.NewObjectType().
			WithProperties(
				astmodel.NewPropertyDefinition(
					"Location", "location", astmodel.OptionalStringType),
				astmodel.NewPropertyDefinition(
					"Tags", "tags", astmodel.NewOptionalType(astmodel.AnyType))))

	webTestProperties := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "WebTestProperties"),
		astmodel.NewObjectType().
			WithProperties(
				astmodel.NewPropertyDefinition(
					"Alias", "alias", astmodel.OptionalStringType)))

	webTest := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "WebTest"),
		astmodel.NewAllOfType(
			webtestsResource.Name(),
			astmodel.NewObjectType().WithProperties(
				astmodel.NewPropertyDefinition(
					"Kind", "kind", astmodel.NewOptionalType(astmodel.NewEnumType(
						astmodel.StringType,
						astmodel.MakeEnumValue("MultiStep", "multistep"),
						astmodel.MakeEnumValue("Ping", "ping")))),
				astmodel.NewPropertyDefinition("Properties", "properties", webTestProperties.Name()))))

	webTestSpec := astmodel.NewAllOfType(
		webTest.Name(),
		astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("AzureName", "azurename", astmodel.StringType),
			astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType)))

	webTest_Status := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "WebTest_Status"),
		astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("Status", "status", astmodel.OptionalStringType)))

	webTestResource := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "WebTestResource"),
		astmodel.NewResourceType(webTestSpec, webTest_Status.Name()))

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(webtestsResource, webTestProperties, webTest, webTest_Status, webTestResource)

	state := NewState(defs)
	stage := ConvertAllOfAndOneOfToObjects(idFactory)

	finalState, err := stage.Run(context.TODO(), state)
	g.Expect(err).To(BeNil())

	// Check on the final shape
	for _, def := range finalState.definitions {
		test.AssertDefinitionHasExpectedShape(t, def.Name().Name(), def)
	}
}

func TestConversionOfSequentialOneOf_ReturnsExpectedResults(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	first := astmodel.NewObjectType().
		WithProperties(
			astmodel.NewPropertyDefinition("alpha", "alpha", astmodel.StringType),
			astmodel.NewPropertyDefinition("beta", "beta", astmodel.StringType))

	second := astmodel.NewObjectType().
		WithProperties(
			astmodel.NewPropertyDefinition("gamma", "gamma", astmodel.StringType),
			astmodel.NewPropertyDefinition("delta", "delta", astmodel.StringType))

	third := astmodel.NewObjectType().
		WithProperties(
			astmodel.NewPropertyDefinition("epsilon", "epsilon", astmodel.StringType),
			astmodel.NewPropertyDefinition("zeta", "zeta", astmodel.StringType))

	firstDef := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "First"),
		first)

	allOfDef := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(test.Pkg2020, "AllOf"),
		astmodel.NewAllOfType(firstDef.Name(), second, third))

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(allOfDef, firstDef)

	state := NewState(defs)
	stage := ConvertAllOfAndOneOfToObjects(idFactory)

	finalState, err := stage.Run(context.TODO(), state)
	g.Expect(err).To(BeNil())

	// Check on the final shape
	for _, def := range finalState.definitions {
		test.AssertDefinitionHasExpectedShape(t, def.Name().Name(), def)
	}
}
