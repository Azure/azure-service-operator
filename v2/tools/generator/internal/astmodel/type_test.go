/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func TestWriteDebugDescription(t *testing.T) {
	t.Parallel()

	here := MakeLocalPackageReference("local", "test", "v", "1")
	there := MakeLocalPackageReference("local", "test", "v", "2")

	age := MakeTypeName(here, "Age")

	personId := MakeTypeName(here, "PersonId")
	otherPersonId := MakeTypeName(there, "PersonId")

	suit := MakeTypeName(here, "Suit")
	diamonds := MakeEnumValue("diamonds", "Diamonds")
	hearts := MakeEnumValue("hearts", "Hearts")
	clubs := MakeEnumValue("clubs", "Clubs")
	spades := MakeEnumValue("spades", "Spades")
	suitEnum := NewEnumType(StringType, diamonds, hearts, clubs, spades)

	armAge := ARMFlag.ApplyTo(age)
	armSuit := ARMFlag.ApplyTo(suit)

	erroredAge := NewErroredType(age, []string{"boom"}, []string{"oh oh"})

	cases := []struct {
		name     string
		subject  Type
		expected string
	}{
		{"Primitive type shows name (int)", IntType, "int"},
		{"Primitive type shows name (string)", StringType, "string"},
		{"Optional primitive type shows * prefix (int)", OptionalIntType, "*int"},
		{"Optional primitive type shows * prefix (string)", OptionalStringType, "*string"},
		{"Array of primitive type shows type of member", NewArrayType(StringType), "Array[string]"},
		{"Array of optional primitive type shows type of member", NewArrayType(OptionalStringType), "Array[*string]"},
		{"Map of primitive types shows types of keys and values", NewMapType(StringType, IntType), "Map[string]int"},
		{"Map with optional value type shows optionality", NewMapType(StringType, OptionalStringType), "Map[string]*string"},
		{"Alias of type shows name of alias", age, "Age"},
		{"Map with alias value shows name of alias", NewMapType(StringType, age), "Map[string]Age"},
		{"Map with alias key shows name of alias", NewMapType(personId, age), "Map[PersonId]Age"},
		{"Enumeration shows base type and values in alphabetical order", suitEnum, "enum:string[clubs|diamonds|hearts|spades]"},
		{"Alias of enumeration shows name of alias", suit, "Suit"},
		{"Map using aliases shows names of aliases", NewMapType(suit, age), "Map[Suit]Age"},
		{"Flagged type shows details of flags", armAge, "Age[Flag:arm]"},
		{"Flagged alias shows details of flags", armSuit, "Suit[Flag:arm]"},
		{"Errored type shows details of errors", erroredAge, "Error[Age|boom|oh oh]"},
		{"Type name from current package has simple form", personId, "PersonId"},
		{"Type name from other packages shows full path", otherPersonId, "local/test/v2.PersonId"},
		{"Type name from external package qualified by package name", SecretReferenceType, "genruntime.SecretReference"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var builder strings.Builder
			c.subject.WriteDebugDescription(&builder, here)
			g.Expect(builder.String()).To(Equal(c.expected))
		})
	}
}

// TestWriteDebugDescriptionNils ensures that WriteDebugDescription() doesn't panic even if called on a nil reference,
// (this is a diagnostic method that should pretty much always do something useful)
func TestWriteDebugDescriptionNils(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var buffer strings.Builder

	// We call WriteDebugDescription on a bunch of zero variables
	// Declarations are based on the receiver used for the implementation of WriteDebugDescription()

	var p *PrimitiveType
	g.Expect(func() { p.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var o *OptionalType
	g.Expect(func() { o.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var a *ArrayType
	g.Expect(func() { a.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var m *MapType
	g.Expect(func() { m.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var e *EnumType
	g.Expect(func() { e.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var f *FlaggedType
	g.Expect(func() { f.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var all *AllOfType
	g.Expect(func() { all.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var one *OneOfType
	g.Expect(func() { one.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var obj *ObjectType
	g.Expect(func() { obj.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var rsrc *ResourceType
	g.Expect(func() { rsrc.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

	var v ValidatedType
	g.Expect(func() { v.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())
}
