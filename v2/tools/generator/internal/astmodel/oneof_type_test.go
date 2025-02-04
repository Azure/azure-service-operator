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

	version1 := NewOneOfType("one", StringType, BoolType)
	version2 := NewOneOfType("one", BoolType, StringType)

	g.Expect(TypeEquals(version1, version2)).To(BeTrue())
	g.Expect(TypeEquals(version2, version1)).To(BeTrue())
}

func TestOneOfMustHaveAllTypesToBeEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	version1 := NewOneOfType("one", StringType, BoolType, FloatType)
	version2 := NewOneOfType("one", BoolType, StringType)

	g.Expect(TypeEquals(version1, version2)).To(BeFalse())
	g.Expect(TypeEquals(version2, version1)).To(BeFalse())
}

func TestOneOfsWithDifferentTypesAreNotEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	version1 := NewOneOfType("one", StringType, FloatType)
	version2 := NewOneOfType("one", BoolType, StringType)

	g.Expect(TypeEquals(version1, version2)).To(BeFalse())
	g.Expect(TypeEquals(version2, version1)).To(BeFalse())
}

var expectedOneOfPanic = "OneOfType should have been replaced by generation time by 'convertAllOfAndOneOf' phase"

func TestOneOfAsTypePanics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		_, err := x.AsTypeExpr(&CodeGenerationContext{})
		g.Expect(err).To(HaveOccurred()) // Unreachable, but keeps lint happy
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}

func TestOneOfAsDeclarationsPanics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		//nolint:errcheck // error will never be returned due to panic
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

func TestOneOfType_Equals_GivenChange_RecognisesDifference(t *testing.T) {
	t.Parallel()

	nameMixin := NewObjectType().WithProperties(
		NewPropertyDefinition("Name", "name", StringType))
	locationMixin := NewObjectType().WithProperties(
		NewPropertyDefinition("Location", "location", StringType))

	cases := map[string]struct {
		change         func(*OneOfType) *OneOfType
		expectedEquals bool
	}{
		"adding a property object": {
			change: func(oneOf *OneOfType) *OneOfType {
				return oneOf.WithAdditionalPropertyObject(locationMixin)
			},
			expectedEquals: false,
		},
		"without property objects": {
			change: func(oneOf *OneOfType) *OneOfType {
				return oneOf.WithoutAnyPropertyObjects()
			},
			expectedEquals: false,
		},
		"with different property object": {
			change: func(oneOf *OneOfType) *OneOfType {
				return oneOf.WithoutAnyPropertyObjects().WithAdditionalPropertyObject(locationMixin)
			},
			expectedEquals: false,
		},
		"with option type already present": {
			change: func(oneOf *OneOfType) *OneOfType {
				return oneOf.WithType(StringType)
			},
			expectedEquals: true,
		},
		"with additional option type": {
			change: func(oneOf *OneOfType) *OneOfType {
				return oneOf.WithType(FloatType)
			},
			expectedEquals: false,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			oneOf := NewOneOfType("oneOf").
				WithType(StringType).
				WithType(BoolType).
				WithAdditionalPropertyObject(nameMixin)

			actual := c.change(oneOf)
			g.Expect(oneOf.Equals(actual, EqualityOverrides{})).To(Equal(c.expectedEquals))
		})
	}
}
