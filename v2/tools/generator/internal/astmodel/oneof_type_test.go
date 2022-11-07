/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestOneOfEqualityDoesNotCareAboutOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := NewOneOfType("x", StringType, BoolType)
	y := NewOneOfType("y", BoolType, StringType)

	g.Expect(TypeEquals(x, y)).To(BeTrue())
	g.Expect(TypeEquals(y, x)).To(BeTrue())
}

func TestOneOfMustHaveAllTypesToBeEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := NewOneOfType("x", StringType, BoolType, FloatType)
	y := NewOneOfType("y", BoolType, StringType)

	g.Expect(TypeEquals(x, y)).To(BeFalse())
	g.Expect(TypeEquals(y, x)).To(BeFalse())
}

func TestOneOfsWithDifferentTypesAreNotEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := NewOneOfType("x", StringType, FloatType)
	y := NewOneOfType("y", BoolType, StringType)

	g.Expect(TypeEquals(x, y)).To(BeFalse())
	g.Expect(TypeEquals(y, x)).To(BeFalse())
}

var expectedOneOfPanic = "OneOfType should have been replaced by generation time by 'convertAllOfAndOneOf' phase"

func TestOneOfAsTypePanics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		x.AsType(&CodeGenerationContext{})
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}

func TestOneOfAsDeclarationsPanics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		x.AsDeclarations(&CodeGenerationContext{}, DeclarationContext{})
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}

func TestOneOfRequiredImportsPanics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		x.RequiredPackageReferences()
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}

func TestOneOfType_WithAdditionalPropertiesFromObject_GivenProperties_ReturnsOneOfContainingProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	oneOf := NewOneOfType("oneOf")
	objA := NewObjectType().WithProperties(
		NewPropertyDefinition("FullName", "fullName", StringType),
		NewPropertyDefinition("KnownAs", "knownAs", StringType))
	objB := NewObjectType().WithProperties(
		NewPropertyDefinition("FullName", "fullName", StringType),
		NewPropertyDefinition("KnownAs", "knownAs", StringType))

	result := oneOf.WithAdditionalPropertyObject(objA).WithAdditionalPropertyObject(objB)
	g.Expect(result).NotTo(BeNil())

	g.Expect(result.PropertyObjects()).To(HaveLen(2))
}

func TestOneOfType_WithoutAnyPropertyObjects_GivenProperties_ReturnsOneOfWithNone(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	oneOf := NewOneOfType("oneOf")
	obj := NewObjectType().WithProperties(
		NewPropertyDefinition("FullName", "fullName", StringType),
		NewPropertyDefinition("KnownAs", "knownAs", StringType))

	oneOf = oneOf.WithAdditionalPropertyObject(obj)

	result := oneOf.WithoutAnyPropertyObjects()
	g.Expect(result).NotTo(BeNil())
	g.Expect(result.PropertyObjects()).To(HaveLen(0))
}
