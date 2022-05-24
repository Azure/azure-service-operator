/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"

	. "github.com/onsi/gomega"
)

func makeSynth() synthesizer {
	return synthesizer{
		defs:      make(astmodel.TypeDefinitionSet),
		idFactory: astmodel.NewIdentifierFactory(),
	}
}

var (
	mapStringInterface = astmodel.NewStringMapType(astmodel.AnyType)
	mapStringString    = astmodel.NewStringMapType(astmodel.StringType)
	mapInterfaceString = astmodel.NewMapType(astmodel.AnyType, astmodel.StringType)
)

var emptyObject = astmodel.NewObjectType()

func defineEnum(strings ...string) astmodel.Type {
	var values []astmodel.EnumValue
	for _, value := range strings {
		values = append(values, astmodel.EnumValue{
			Identifier: value,
			Value:      value,
		})
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
	// this feels a bit wrong but it seems to be expected in real life specs
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
		astmodel.NewOptionalType(astmodel.StringType),
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

	oneOf := astmodel.BuildOneOfType(astmodel.IntType, astmodel.StringType, astmodel.BoolType)

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

	oneOf := astmodel.BuildOneOfType(
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
	oneOf := astmodel.BuildOneOfType(astmodel.BoolType, r).(*astmodel.OneOfType)

	synth := makeSynth()
	synth.specOrStatus = chooseSpec

	expected := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition(astmodel.PropertyName("Bool0"), "bool0", astmodel.BoolType).
			MakeTypeOptional().WithDescription("Mutually exclusive with all other properties"),
		astmodel.NewPropertyDefinition(astmodel.PropertyName("Resource1"), "resource1", r).
			MakeTypeOptional().WithDescription("Mutually exclusive with all other properties"),
	)

	names, err := synth.getOneOfPropNames(oneOf)
	g.Expect(err).To(BeNil())

	result := synth.oneOfObject(oneOf, names)
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
