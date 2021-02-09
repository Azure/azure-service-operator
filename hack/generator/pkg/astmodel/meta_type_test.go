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
	enumType := NewEnumType(StringType, []EnumValue{})
	nameType := MakeTypeName(makeTestLocalPackageReference("g", "v"), "foo")

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesArePrimitives", StringType, StringType},
		{"ObjectAreNotPrimitives", objectType, nil},
		{"ArraysAreNotPrimitives", arrayType, nil},
		{"MapsAreNotPrimitives", mapType, nil},
		{"EnumsAreNotPrimitives", enumType, nil},
		{"NamesAreNotPrimitives", nameType, nil},
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

			actual, ok := AsPrimitiveType(c.subject)

			if c.expected == nil {
				g.Expect(ok).To(BeFalse())
			} else {
				g.Expect(actual).To(Equal(c.expected))
				g.Expect(ok).To(BeTrue())
			}

		})
	}
}

func TestAsObjectType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(StringType)
	enumType := NewEnumType(StringType, []EnumValue{})
	nameType := MakeTypeName(makeTestLocalPackageReference("g", "v"), "foo")

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotObjects", StringType, nil},
		{"ObjectsAreObjects", objectType, objectType},
		{"ArraysAreNotObjects", arrayType, nil},
		{"MapsAreNotObjects", mapType, nil},
		{"EnumsAreNotObjects", enumType, nil},
		{"NamesAreNotObjects", nameType, nil},
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

			actual, ok := AsObjectType(c.subject)

			if c.expected == nil {
				g.Expect(ok).To(BeFalse())
			} else {
				g.Expect(actual).To(Equal(c.expected))
				g.Expect(ok).To(BeTrue())
			}

		})
	}
}

func TestAsArrayType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(objectType)
	enumType := NewEnumType(StringType, []EnumValue{})
	nameType := MakeTypeName(makeTestLocalPackageReference("g", "v"), "foo")

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotArrays", StringType, nil},
		{"ObjectsAreNotArrays", objectType, nil},
		{"ArraysAreArrays", arrayType, arrayType},
		{"MapsAreNotArrays", mapType, nil},
		{"EnumsAreNotArrays", enumType, nil},
		{"NamesAreNotArrays", nameType, nil},
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

			actual, ok := AsArrayType(c.subject)

			if c.expected == nil {
				g.Expect(ok).To(BeFalse())
			} else {
				g.Expect(actual).To(Equal(c.expected))
				g.Expect(ok).To(BeTrue())
			}

		})
	}
}

func TestAsMapType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(objectType)
	enumType := NewEnumType(StringType, []EnumValue{})
	nameType := MakeTypeName(makeTestLocalPackageReference("g", "v"), "foo")

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotMaps", StringType, nil},
		{"ObjectsAreNotMaps", objectType, nil},
		{"ArraysAreNotMaps", arrayType, nil},
		{"MapsAreMaps", mapType, mapType},
		{"EnumsAreNotMaps", enumType, nil},
		{"NamesAreNotMaps", nameType, nil},
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

			actual, ok := AsMapType(c.subject)

			if c.expected == nil {
				g.Expect(ok).To(BeFalse())
			} else {
				g.Expect(actual).To(Equal(c.expected))
				g.Expect(ok).To(BeTrue())
			}

		})
	}
}

func TestAsOptionalType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(objectType)
	enumType := NewEnumType(StringType, []EnumValue{})
	nameType := MakeTypeName(makeTestLocalPackageReference("g", "v"), "foo")

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotOptional", StringType, nil},
		{"ObjectsAreNotOptional", objectType, nil},
		{"ArraysAreNotOptional", arrayType, nil},
		{"MapsAreNotOptional", mapType, nil},
		{"EnumsAreNotOptional", enumType, nil},
		{"NamesAreNotOptional", nameType, nil},
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

			actual, ok := AsOptionalType(c.subject)

			if c.expected == nil {
				g.Expect(ok).To(BeFalse())
			} else {
				g.Expect(actual).To(Equal(c.expected))
				g.Expect(ok).To(BeTrue())
			}

		})
	}
}

func TestAsEnumType(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(objectType)
	enumType := NewEnumType(StringType, []EnumValue{})
	nameType := MakeTypeName(makeTestLocalPackageReference("g", "v"), "foo")

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotEnums", StringType, nil},
		{"ObjectsAreNotEnums", objectType, nil},
		{"ArraysAreNotEnums", arrayType, nil},
		{"MapsAreNotEnums", mapType, nil},
		{"EnumsAreEnums", enumType, enumType},
		{"NamesAreNotEnums", nameType, nil},
		{"OptionalAreNotEnums", optionalType, nil},
		{"FlaggedContainingEnums", OneOfFlag.ApplyTo(enumType), enumType},
		{"FlaggedNotContainingEnums", OneOfFlag.ApplyTo(StringType), nil},
		{"ValidatedContainingEnums", NewValidatedType(enumType, nil), enumType},
		{"ValidatedNotContainingEnums", NewValidatedType(StringType, nil), nil},
		{"ErroredContainingEnums", NewErroredType(enumType, nil, nil), enumType},
		{"ErroredNotContainingEnums", NewErroredType(StringType, nil, nil), nil},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual, ok := AsEnumType(c.subject)

			if c.expected == nil {
				g.Expect(ok).To(BeFalse())
			} else {
				g.Expect(actual).To(Equal(c.expected))
				g.Expect(ok).To(BeTrue())
			}

		})
	}
}

func TestAsTypeName(t *testing.T) {

	objectType := NewObjectType()
	arrayType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optionalType := NewOptionalType(objectType)
	enumType := NewEnumType(StringType, []EnumValue{})
	nameType := MakeTypeName(makeTestLocalPackageReference("g", "v"), "foo")

	cases := []struct {
		name     string
		subject  Type
		expected Type
	}{
		{"PrimitivesAreNotNames", StringType, nil},
		{"ObjectsAreNotNames", objectType, nil},
		{"ArraysAreNotNames", arrayType, nil},
		{"MapsAreNotNames", mapType, nil},
		{"EnumsAreNotNames", enumType, nil},
		{"NamesAreNames", nameType, nameType},
		{"OptionalAreNotNames", optionalType, nil},
		{"FlaggedContainingNames", OneOfFlag.ApplyTo(nameType), nameType},
		{"FlaggedNotContainingNames", OneOfFlag.ApplyTo(StringType), nil},
		{"ValidatedContainingNames", NewValidatedType(nameType, nil), nameType},
		{"ValidatedNotContainingNames", NewValidatedType(StringType, nil), nil},
		{"ErroredContainingNames", NewErroredType(nameType, nil, nil), nameType},
		{"ErroredNotContainingNames", NewErroredType(StringType, nil, nil), nil},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual, ok := AsTypeName(c.subject)

			if c.expected == nil {
				g.Expect(ok).To(BeFalse())
			} else {
				g.Expect(actual).To(Equal(c.expected))
				g.Expect(ok).To(BeTrue())
			}

		})
	}
}
