/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"testing"

	. "github.com/onsi/gomega"
)

/*
 * NewPropertyBag() Tests
 */

func TestPropertyBag_NewPropertyBag_WithNoParameters_ReturnsEmptyBag(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	bag := NewPropertyBag()
	g.Expect(bag).To(BeEmpty())
}

func TestPropertyBag_NewPropertyBag_WithZeroBag_ReturnsEmptyBag(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var zeroBag PropertyBag
	bag := NewPropertyBag(zeroBag)
	g.Expect(bag).To(BeEmpty())
}

func TestPropertyBag_NewPropertyBag_WithEmptyBag_ReturnsEmptyBag(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	emptyBag := NewPropertyBag()
	bag := NewPropertyBag(emptyBag)
	g.Expect(bag).To(BeEmpty())
}

func TestPropertyBag_NewPropertyBag_WithSingleItemBag_ReturnsBagWithExpectedLength(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewPropertyBag()
	g.Expect(original.Add("Answer", 42)).To(Succeed())

	bag := NewPropertyBag(original)
	g.Expect(bag).To(HaveLen(1))
}

func TestPropertyBag_NewPropertyBag_WithMultipleItemBag_ReturnsBagWithExpectedKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewPropertyBag()
	g.Expect(original.Add("Answer", 42)).To(Succeed())
	g.Expect(original.Add("Halloween", "31OCT")).To(Succeed())
	g.Expect(original.Add("Christmas", "25DEC")).To(Succeed())

	bag := NewPropertyBag(original)
	g.Expect(bag).To(HaveKey("Answer"))
	g.Expect(bag).To(HaveKey("Halloween"))
	g.Expect(bag).To(HaveKey("Christmas"))
	g.Expect(bag).To(HaveLen(3))
}

func TestPropertyBag_NewPropertyBag_WithMultipleSingleItemBags_ReturnsBagWithExpectedKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	first := NewPropertyBag()
	g.Expect(first.Add("Answer", 42)).To(Succeed())

	second := NewPropertyBag()
	g.Expect(second.Add("Halloween", "31OCT")).To(Succeed())
	g.Expect(second.Add("Christmas", "25DEC")).To(Succeed())

	bag := NewPropertyBag(first, second)
	g.Expect(bag).To(HaveKey("Answer"))
	g.Expect(bag).To(HaveKey("Halloween"))
	g.Expect(bag).To(HaveKey("Christmas"))
	g.Expect(bag).To(HaveLen(3))
}

func TestPropertyBag_NewPropertyBag_WithOverlappingBags_ReturnsBagWithExpectedKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	first := NewPropertyBag()
	g.Expect(first.Add("Answer", 42)).To(Succeed())
	g.Expect(first.Add("Gift", "Skull")).To(Succeed())

	second := NewPropertyBag()
	g.Expect(second.Add("Halloween", "31OCT")).To(Succeed())
	g.Expect(second.Add("Christmas", "25DEC")).To(Succeed())
	g.Expect(second.Add("Gift", "Gold")).To(Succeed())

	bag := NewPropertyBag(first, second)
	g.Expect(bag).To(HaveKey("Answer"))
	g.Expect(bag).To(HaveKey("Halloween"))
	g.Expect(bag).To(HaveKey("Christmas"))
	g.Expect(bag).To(HaveKey("Gift"))
	g.Expect(bag).To(HaveLen(4))

	var gift string
	err := bag.Pull("Gift", &gift)
	g.Expect(err).To(Succeed())
	g.Expect(gift).To(Equal("Gold"))
}

/*
 * Roundtrip Tests
 */

func TestPropertyBag_CorrectlyRoundTripsIntegers(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var original int = 42
	var actual int

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(Succeed())

	err = bag.Pull("prop", &actual)
	g.Expect(err).To(Succeed())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_CorrectlyRoundTrips64bitIntegers(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var original int64 = 3735928559
	var actual int64

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(Succeed())

	err = bag.Pull("prop", &actual)
	g.Expect(err).To(Succeed())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_CorrectlyRoundTripsStrings(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var original string = "Pack my box with five dozen liquor jugs"
	var actual string

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(Succeed())

	err = bag.Pull("prop", &actual)
	g.Expect(err).To(Succeed())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_CorrectlyRoundTripsBooleans(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var original bool = true
	var actual bool

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(Succeed())

	err = bag.Pull("prop", &actual)
	g.Expect(err).To(Succeed())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_CorrectlyRoundTripsFloats(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var original float64 = 1.0 / 10 // Deliberately chose a value that can't be represented exactly
	var actual float64

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(Succeed())

	err = bag.Pull("prop", &actual)
	g.Expect(err).To(Succeed())
	g.Expect(actual).To(Equal(original))
}

type Person struct {
	LegalName  string
	FamilyName string
	KnownAs    string
	AlphaKey   string
}

func TestPropertyBag_CorrectlyRoundTripsStructs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var original Person = Person{
		LegalName:  "Michael Theodore Mouse",
		FamilyName: "Mouse",
		KnownAs:    "Mickey",
		AlphaKey:   "Mouse",
	}
	var actual Person

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(Succeed())

	err = bag.Pull("prop", &actual)
	g.Expect(err).To(Succeed())
	g.Expect(actual).To(Equal(original))
}

func TestPropertyBag_WhenPropertyNotPresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var original int64 = 51741654765
	var actual int64

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(Succeed())

	err = bag.Pull("otherProp", &actual)
	g.Expect(err).NotTo(Succeed())
}

func TestPropertyBag_WhenPropertyWrongType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var original int64 = 51741654765
	var actual bool

	bag := make(PropertyBag)
	err := bag.Add("prop", original)
	g.Expect(err).To(Succeed())

	err = bag.Pull("prop", &actual)
	g.Expect(err).NotTo(Succeed())
}
