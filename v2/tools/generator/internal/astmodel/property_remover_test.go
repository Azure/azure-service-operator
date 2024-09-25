/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PropertyRemover_GivenDefinition_RemovesExpectedProperty(t *testing.T) {
	t.Parallel()

	simpleObject := NewObjectType().WithProperties(fullName, familyName, legalName, knownAs)
	validatedObject := NewValidatedType(simpleObject, StringValidations{})
	optionalObject := NewOptionalType(validatedObject)

	cases := map[string]struct {
		theType               Type
		propName              PropertyName
		expectedPropertyCount int
	}{
		"Remove existing property from object": {
			theType:               simpleObject,
			propName:              legalName.propertyName,
			expectedPropertyCount: 3,
		},
		"Remove existing property from validated object": {
			theType:               validatedObject,
			propName:              legalName.propertyName,
			expectedPropertyCount: 3,
		},
		"Remove existing property from optional object": {
			theType:               optionalObject,
			propName:              legalName.propertyName,
			expectedPropertyCount: 3,
		},
		"Remove non-existent property from object": {
			theType:               simpleObject,
			propName:              "DoesNotExist",
			expectedPropertyCount: 4,
		},
	}

	ref := makeTestLocalPackageReference("group", "2024-09-01")

	remover := NewPropertyRemover()
	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			// Arrange
			def := MakeTypeDefinition(
				MakeInternalTypeName(ref, "TestDef"),
				c.theType)

			// Act
			newDef, err := remover.Remove(def, c.propName)
			g.Expect(err).To(Succeed())

			// Assert
			obj, ok := AsObjectType(newDef.Type())
			g.Expect(ok).To(BeTrue())
			g.Expect(obj.properties.Len()).To(Equal(c.expectedPropertyCount))
		})
	}
}
