/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPropertySet_NewPropertySet_ReturnsEmptySet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPropertySet()
	g.Expect(set).To(HaveLen(0))
}

func TestPropertySet_NewPropertySetWithProperties_ReturnsSetContainingProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPropertySet(fullName, familyName, knownAs, gender)
	g.Expect(set).To(ContainElements(fullName, familyName, knownAs, gender))
	g.Expect(set).To(HaveLen(4))
}

func TestPropertySet_AsSlice_ReturnsSliceContainingProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	slice := NewPropertySet(fullName, familyName, knownAs, gender).AsSlice()
	g.Expect(slice).To(ContainElements(fullName, familyName, knownAs, gender))
	g.Expect(slice).To(HaveLen(4))
}

func TestPropertySet_Add_ReturnsExpectedSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPropertySet()
	set.Add(fullName)
	set.Add(familyName)
	set.Add(knownAs)
	set.Add(gender)
	g.Expect(set).To(ContainElements(fullName, familyName, knownAs, gender))
	g.Expect(set).To(HaveLen(4))
}

func TestPropertySet_Copy_ReturnsExpectedSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	clone := NewPropertySet(fullName, familyName, knownAs, gender).Copy()
	g.Expect(clone).To(ContainElements(fullName, familyName, knownAs, gender))
	g.Expect(clone).To(HaveLen(4))
}

func TestPropertySet_FindAll_ReturnsExpectedSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPropertySet(fullName, familyName, knownAs, gender)
	set.FindAll(func(prop *PropertyDefinition) bool {
		prim, ok := AsPrimitiveType(prop.PropertyType())
		if !ok {
			return false
		}
		return prim == StringType
	})
	g.Expect(set).To(ContainElements(fullName, familyName, knownAs, gender))
	g.Expect(set).To(HaveLen(4))
}
