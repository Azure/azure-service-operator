/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestAsPrimitiveType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(arrayType)

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesArePrimitives", StringType, StringType},
		{"ObjectAreNotPrimitives", objectType, nil},
		{"ArraysAreNotPrimitives", arrayType, nil},
		{"MapsAreNotPrimitives", mapType, nil},
		{"OptionalAreNotPrimitives", optionalType, nil},
		{"OptionalContainingPrimitive", NewOptionalType(StringType), StringType},
		{"OptionalNotContainingPrimitive", NewOptionalType(objectType), nil},
		{"FlaggedContainingPrimitive", OneOfFlag.ApplyTo(StringType), StringType},
		{"FlaggedNotContainingPrimitive", OneOfFlag.ApplyTo(objectType), nil},
		{"ValidatedContainingPrimitive", NewValidatedType(StringType, nil), StringType},
		{"ValidatedNotContainingPrimitive", NewValidatedType(objectType, nil), nil},
		{"ErroredContainingPrimitive", NewErroredType(StringType, nil, nil), StringType},
		{"ErroredNotContainingPrimitive", NewErroredType(objectType, nil, nil), nil},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := AsPrimitiveType(c.subject)

			if c.expected == nil {
				g.Expect(actual).To(BeNil())
			} else {
				g.Expect(actual).To(Equal(c.expected))
			}

		})
	}
}

func TestAsObjectType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(StringType)

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotObjects", StringType, nil},
		{"ObjectsAreObjects", objectType, objectType},
		{"ArraysAreNotObjects", arrayType, nil},
		{"MapsAreNotObjects", mapType, nil},
		{"OptionalAreNotObjects", optionalType, nil},
		{"OptionalContainingObject", NewOptionalType(objectType), objectType},
		{"OptionalNotContainingObject", NewOptionalType(StringType), nil},
		{"FlaggedContainingObject", OneOfFlag.ApplyTo(objectType), objectType},
		{"FlaggedNotContainingObject", OneOfFlag.ApplyTo(StringType), nil},
		{"ValidatedContainingObject", NewValidatedType(objectType, nil), objectType},
		{"ValidatedNotContainingObject", NewValidatedType(StringType, nil), nil},
		{"ErroredContainingObject", NewErroredType(objectType, nil, nil), objectType},
		{"ErroredNotContainingObject", NewErroredType(StringType, nil, nil), nil},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := AsObjectType(c.subject)

			if c.expected == nil {
				g.Expect(actual).To(BeNil())
			} else {
				g.Expect(actual).To(Equal(c.expected))
			}

		})
	}
}

func TestAsArrayType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(objectType)

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotArrays", StringType, nil},
		{"ObjectsAreNotArrays", objectType, nil},
		{"ArraysAreArrays", arrayType, arrayType},
		{"MapsAreNotArrays", mapType, nil},
		{"OptionalAreNotArrays", optionalType, nil},
		{"OptionalContainingArray", NewOptionalType(arrayType), arrayType},
		{"OptionalNotContainingArray", NewOptionalType(StringType), nil},
		{"FlaggedContainingArray", OneOfFlag.ApplyTo(arrayType), arrayType},
		{"FlaggedNotContainingArray", OneOfFlag.ApplyTo(StringType), nil},
		{"ValidatedContainingArray", NewValidatedType(arrayType, nil), arrayType},
		{"ValidatedNotContainingArray", NewValidatedType(StringType, nil), nil},
		{"ErroredContainingArray", NewErroredType(arrayType, nil, nil), arrayType},
		{"ErroredNotContainingArray", NewErroredType(StringType, nil, nil), nil},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := AsArrayType(c.subject)

			if c.expected == nil {
				g.Expect(actual).To(BeNil())
			} else {
				g.Expect(actual).To(Equal(c.expected))
			}

		})
	}
}

func TestAsMapType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(objectType)

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotMaps", StringType, nil},
		{"ObjectsAreNotMaps", objectType, nil},
		{"ArraysAreNotMaps", arrayType, nil},
		{"MapsAreMaps", mapType, mapType},
		{"OptionalAreNotMaps", optionalType, nil},
		{"OptionalContainingMaps", NewOptionalType(mapType), mapType},
		{"OptionalNotContainingMaps", NewOptionalType(StringType), nil},
		{"FlaggedContainingMaps", OneOfFlag.ApplyTo(mapType), mapType},
		{"FlaggedNotContainingMaps", OneOfFlag.ApplyTo(StringType), nil},
		{"ValidatedContainingMaps", NewValidatedType(mapType, nil), mapType},
		{"ValidatedNotContainingMaps", NewValidatedType(StringType, nil), nil},
		{"ErroredContainingMaps", NewErroredType(mapType, nil, nil), mapType},
		{"ErroredNotContainingMaps", NewErroredType(StringType, nil, nil), nil},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := AsMapType(c.subject)

			if c.expected == nil {
				g.Expect(actual).To(BeNil())
			} else {
				g.Expect(actual).To(Equal(c.expected))
			}

		})
	}
}

func TestAsOptionalType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(objectType)

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotOptional", StringType, nil},
		{"ObjectsAreNotOptional", objectType, nil},
		{"ArraysAreNotOptional", arrayType, nil},
		{"MapsAreNotOptional", mapType, nil},
		{"OptionalAreOptional", optionalType, optionalType},
		{"FlaggedContainingOptional", OneOfFlag.ApplyTo(optionalType), optionalType},
		{"FlaggedNotContainingOptional", OneOfFlag.ApplyTo(StringType), nil},
		{"ValidatedContainingOptional", NewValidatedType(optionalType, nil), optionalType},
		{"ValidatedNotContainingOptional", NewValidatedType(StringType, nil), nil},
		{"ErroredContainingOptional", NewErroredType(optionalType, nil, nil), optionalType},
		{"ErroredNotContainingOptional", NewErroredType(StringType, nil, nil), nil},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := AsOptionalType(c.subject)

			if c.expected == nil {
				g.Expect(actual).To(BeNil())
			} else {
				g.Expect(actual).To(Equal(c.expected))
			}

		})
	}
}
