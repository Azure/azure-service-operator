/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

var (
	aboveValue = EnumValue{
		Identifier: "Above",
		Value:      "Above",
	}

	underValue = EnumValue{
		Identifier: "Under",
		Value:      "Under",
	}

	leftValue = EnumValue{
		Identifier: "Left",
		Value:      "left",
	}

	rightValue = EnumValue{
		Identifier: "Right",
		Value:      "Right",
	}
)

/*
 * NewEnumType tests
 */

func Test_NewEnumType_GivenValues_InitializesFields(t *testing.T) {
	g := NewGomegaWithT(t)

	e := NewEnumType(StringType, []EnumValue{aboveValue, underValue})
	g.Expect(e.baseType).To(Equal(StringType))
	g.Expect(e.options).To(ContainElements(aboveValue, underValue))
}

/*
 * BaseType() tests
 */

func Test_EnumTypeBaseType_AfterConstruction_ReturnsExpectedType(t *testing.T) {
	g := NewGomegaWithT(t)

	e := NewEnumType(StringType, []EnumValue{aboveValue, underValue})
	g.Expect(e.BaseType()).To(Equal(StringType))
}

/*
 * Equals() tests
 */

func Test_EnumTypeEquals_GivenEnums_ReturnsExpectedResult(t *testing.T) {

	enum := NewEnumType(StringType, []EnumValue{aboveValue, underValue})

	other := NewEnumType(StringType, []EnumValue{aboveValue, underValue})
	differentBase := NewEnumType(IntType, []EnumValue{aboveValue, underValue})
	differentValueOrder := NewEnumType(StringType, []EnumValue{underValue, aboveValue})
	supersetOfValues := NewEnumType(StringType, []EnumValue{underValue, aboveValue, leftValue, rightValue})
	subsetOfValues := NewEnumType(StringType, []EnumValue{underValue})
	differentValues := NewEnumType(StringType, []EnumValue{leftValue, rightValue})

	cases := []struct {
		name      string
		thisEnum  *EnumType
		otherEnum *EnumType
		expected  bool
	}{
		// Expect equal to self
		{"Equal to self", enum, enum, true},
		{"Equal to self", other, other, true},
		// Expect equal to same
		{"Equal to same", enum, other, true},
		{"Equal to same", other, enum, true},
		// Expect equal when values are reordered
		{"Equal when values are reordered", enum, differentValueOrder, true},
		{"Equal when values reordered", differentValueOrder, enum, true},
		// Expect not-equal when values missing
		{"Not-equal when values missing", enum, subsetOfValues, false},
		{"Not-equal when values missing", supersetOfValues, enum, false},
		// Expect not-equal when properties added
		{"Not-equal when values added", enum, supersetOfValues, false},
		{"Not-equal when values added", supersetOfValues, enum, false},
		// Expect not-equal for different base type
		{"Not-equal when different base type", enum, differentBase, false},
		{"Not-equal when different base type", differentBase, enum, false},
		// Expect not-equal for different values (but same value count)
		{"Not-equal when different values", enum, differentValues, false},
		{"Not-equal when different values", differentValues, enum, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.thisEnum.Equals(c.otherEnum)

			g.Expect(areEqual).To(Equal(c.expected))
		})
	}
}
