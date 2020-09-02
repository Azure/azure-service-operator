/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	. "github.com/onsi/gomega"

	"testing"
)

/*
 * IsArmType() tests
 */

func Test_IsArmType_GivenType_ReturnsExpectedResult(t *testing.T) {
	distanceProperty := NewPropertyDefinition("Meters", "meters", IntType)
	objectType := EmptyObjectType.WithProperty(distanceProperty)
	armType := MakeArmType(ObjectType{})

	cases := []struct {
		name     string
		thisType Type
		expected bool
	}{
		{"String is not ARM type", StringType, false},
		{"Map is not ARM type", NewMapType(StringType, BoolType), false},
		{"Object is not ARM type", objectType, false},
		{"ARM is ARM type", armType, true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			isArm := IsArmType(c.thisType)

			g.Expect(isArm).To(Equal(c.expected))
		})
	}
}

/*
 * IsArmDefinition() tests
 */

func Test_IsArmDefinition_GivenType_ReturnsExpectedResult(t *testing.T) {
	distanceProperty := NewPropertyDefinition("Meters", "meters", IntType)
	objectType := EmptyObjectType.WithProperty(distanceProperty)
	stringDefinition := NewTestObject("Range", distanceProperty)
	armType := MakeArmType(*objectType)
	armDef := MakeTypeDefinition(stringDefinition.name, armType)

	cases := []struct {
		name       string
		definition TypeDefinition
		expected   bool
	}{
		{"String Definition is not ARM definition", stringDefinition, false},
		{"ARM Definition is an ARM definition", armDef, true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			isArm := IsArmDefinition(c.definition)

			g.Expect(isArm).To(Equal(c.expected))
		})
	}
}
