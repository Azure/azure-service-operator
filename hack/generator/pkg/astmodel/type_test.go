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

	here := MakeLocalPackageReference("local", "test", "v1")

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

	types := make(Types)
	types.Add(ageDefinition)
	types.Add(personIdDefinition)
	types.Add(suitDefinition)

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
		{"AliasedType", age, "local/test/v1/Age:int"},
		{"MapOfStringToAge", NewMapType(StringType, age), "Map[string]local/test/v1/Age:int"},
		{"MapOfPersonIDToAge", NewMapType(personId, age), "Map[local/test/v1/PersonId:string]local/test/v1/Age:int"},
		{"SuitEnum", suitEnum, "Enum[string:clubs|diamonds|hearts|spades]"}, // alphabetical
		{"SuitName", suit, "local/test/v1/Suit:Enum[string:clubs|diamonds|hearts|spades]"},
		{"MapOfSuitToAge", NewMapType(suit, age), "Map[local/test/v1/Suit:Enum[string:clubs|diamonds|hearts|spades]]local/test/v1/Age:int"},
		{"FlaggedAge", armAge, "local/test/v1/Age:int[#arm]"},
		{"FlaggedSuit", armSuit, "local/test/v1/Suit:Enum[string:clubs|diamonds|hearts|spades][#arm]"},
		{"ErroredAge", erroredAge, "Error[local/test/v1/Age:int|boom|oh oh]"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var builder strings.Builder
			c.subject.WriteDebugDescription(&builder, types)
			g.Expect(builder.String()).To(Equal(c.expected))
		})
	}
}
