/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestPropertyBag_CorrectlyRoundTripsIntegers(t *testing.T) {
	g := NewWithT(t)
	var original int = 42
	var actual int

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(BeNil())

	found, err := bag.Pull("prop", &actual)
	g.Expect(err).To(BeNil())
	g.Expect(found).To(BeTrue())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_CorrectlyRoundTrips64bitIntegers(t *testing.T) {
	g := NewWithT(t)
	var original int64 = 3735928559
	var actual int64

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(BeNil())

	found, err := bag.Pull("prop", &actual)
	g.Expect(err).To(BeNil())
	g.Expect(found).To(BeTrue())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_CorrectlyRoundTripsStrings(t *testing.T) {
	g := NewWithT(t)
	var original string = "Pack my box with five dozen liquor jugs"
	var actual string

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(BeNil())

	found, err := bag.Pull("prop", &actual)
	g.Expect(err).To(BeNil())
	g.Expect(found).To(BeTrue())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_CorrectlyRoundTripsBooleans(t *testing.T) {
	g := NewWithT(t)
	var original bool = true
	var actual bool

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(BeNil())

	found, err := bag.Pull("prop", &actual)
	g.Expect(err).To(BeNil())
	g.Expect(found).To(BeTrue())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_CorrectlyRoundTripsFloats(t *testing.T) {
	g := NewWithT(t)
	var original float64 = 1 / 10 // Deliberately chose a value that can't be represented exactly
	var actual float64

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(BeNil())

	found, err := bag.Pull("prop", &actual)
	g.Expect(err).To(BeNil())
	g.Expect(found).To(BeTrue())
	g.Expect(actual).To(Equal(original))
}

type Person struct {
	LegalName  string
	FamilyName string
	KnownAs    string
	AlphaKey   string
}

func TestPropertyBag_CorrectlyRoundTripsStructs(t *testing.T) {
	g := NewWithT(t)
	var original Person = Person{
		LegalName:  "Michael Theodore Mouse",
		FamilyName: "Mouse",
		KnownAs:    "Mickey",
		AlphaKey:   "Mouse",
	}
	var actual Person

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(BeNil())

	found, err := bag.Pull("prop", &actual)
	g.Expect(err).To(BeNil())
	g.Expect(found).To(BeTrue())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_WhenPropertyNotPresent(t *testing.T) {
	g := NewWithT(t)
	var original int64 = 51741654765
	var actual int64

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(BeNil())

	found, err := bag.Pull("otherProp", &actual)
	g.Expect(err).To(BeNil())
	g.Expect(found).To(BeFalse())
}

func TestPropertyBag_WhenPropertyWrongType(t *testing.T) {
	g := NewWithT(t)
	var original int64 = 51741654765
	var actual bool

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(BeNil())

	found, err := bag.Pull("prop", &actual)
	g.Expect(err).NotTo(BeNil())
	g.Expect(found).To(BeTrue())
}
