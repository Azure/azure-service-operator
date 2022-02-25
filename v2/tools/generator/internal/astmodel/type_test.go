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

	here := MakeLocalPackageReference("local", "test", "1")

	age := MakeTypeName(here, "Age")
	ageDefinition := MakeTypeDefinition(age, IntType)

	personId := MakeTypeName(here, "PersonId")
	personIdDefinition := MakeTypeDefinition(personId, StringType)

	suit := MakeTypeName(here, "Suit")
	diamonds := EnumValue{"diamonds", "Diamonds"}
	hearts := EnumValue{"hearts", "Hearts"}
	clubs := EnumValue{"clubs", "Clubs"}
	spades := EnumValue{"spades", "Spades"}
	suitEnum := NewEnumType(StringType, diamonds, hearts, clubs, spades)
	suitDefinition := MakeTypeDefinition(suit, suitEnum)

	armAge := ARMFlag.ApplyTo(age)
	armSuit := ARMFlag.ApplyTo(suit)

	erroredAge := NewErroredType(age, []string{"boom"}, []string{"oh oh"})

	defs := make(TypeDefinitionSet)
	defs.Add(ageDefinition)
	defs.Add(personIdDefinition)
	defs.Add(suitDefinition)

	cases := []struct {
		name     string
		subject  Type
		expected string
	}{
		{"Integer", IntType, "int"},
		{"String", StringType, "string"},
		{"OptionalInteger", NewOptionalType(IntType), "Optional[int]"},
		{"OptionalString", NewOptionalType(StringType), "Optional[string]"},
		{"ArrayOfString", NewArrayType(StringType), "Array[string]"},
		{"ArrayOfOptionalString", NewArrayType(NewOptionalType(StringType)), "Array[Optional[string]]"},
		{"MapOfStringToInt", NewMapType(StringType, IntType), "Map[string]int"},
		{"MapOfStringToOptionalString", NewMapType(StringType, NewOptionalType(StringType)), "Map[string]Optional[string]"},
		{"AliasedType", age, "local/test/v1alpha1api1/Age:int"},
		{"MapOfStringToAge", NewMapType(StringType, age), "Map[string]local/test/v1alpha1api1/Age:int"},
		{"MapOfPersonIDToAge", NewMapType(personId, age), "Map[local/test/v1alpha1api1/PersonId:string]local/test/v1alpha1api1/Age:int"},
		{"SuitEnum", suitEnum, "Enum[string:clubs|diamonds|hearts|spades]"}, // alphabetical
		{"SuitName", suit, "local/test/v1alpha1api1/Suit:Enum[string:clubs|diamonds|hearts|spades]"},
		{"MapOfSuitToAge", NewMapType(suit, age), "Map[local/test/v1alpha1api1/Suit:Enum[string:clubs|diamonds|hearts|spades]]local/test/v1alpha1api1/Age:int"},
		{"FlaggedAge", armAge, "local/test/v1alpha1api1/Age:int[#arm]"},
		{"FlaggedSuit", armSuit, "local/test/v1alpha1api1/Suit:Enum[string:clubs|diamonds|hearts|spades][#arm]"},
		{"ErroredAge", erroredAge, "Error[local/test/v1alpha1api1/Age:int|boom|oh oh]"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var builder strings.Builder
			c.subject.WriteDebugDescription(&builder, defs)
			g.Expect(builder.String()).To(Equal(c.expected))
		})
	}
}

// TestWriteDebugDescriptionNils ensures that WriteDebugDescription() doesn't panic even if called on a nil reference,
//(this is a diagnostic method that should pretty much always do something useful)
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

	var n TypeName
	g.Expect(func() { n.WriteDebugDescription(&buffer, nil) }).ShouldNot(Panic())

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
