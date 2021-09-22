/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCanMakeEmptyTypeSet(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeTypeSet()
	y := MakeTypeSet()

	g.Expect(x).To(Equal(y))
	g.Expect(y).To(Equal(x))
}

func TestSetsWithIdenticalObjectsAreStructurallyEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeTypeSet(BoolType, StringType)
	y := MakeTypeSet(StringType, BoolType)

	// not structurally equal
	g.Expect(x).ToNot(Equal(y))
	g.Expect(y).ToNot(Equal(x))

	// but .equals() equal
	g.Expect(x.Equals(y)).To(BeTrue())
	g.Expect(y.Equals(x)).To(BeTrue())
}

func TestSetsWithDifferentNumbersOfObjectAreNotEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeTypeSet(BoolType)
	y := MakeTypeSet(StringType, BoolType)

	// both not structurally equal
	g.Expect(x).ToNot(Equal(y))
	g.Expect(y).ToNot(Equal(x))

	// and not .equals() equal
	g.Expect(x.Equals(y)).To(BeFalse())
	g.Expect(y.Equals(x)).To(BeFalse())
}

func TestSetsWithDifferentContentsAreNotEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeTypeSet(BoolType)
	y := MakeTypeSet(StringType)

	// both not structurally equal
	g.Expect(x).ToNot(Equal(y))
	g.Expect(y).ToNot(Equal(x))

	// and not .equals() equal
	g.Expect(x.Equals(y)).To(BeFalse())
	g.Expect(y.Equals(x)).To(BeFalse())
}

func TestSetsWithEqualObjectsAreEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeTypeSet(NewOptionalType(StringType))
	y := MakeTypeSet(NewOptionalType(StringType))

	// structurally equal
	g.Expect(x).To(Equal(y))
	g.Expect(y).To(Equal(x))

	// and .equals() equal
	g.Expect(x.Equals(y)).To(BeTrue())
	g.Expect(y.Equals(x)).To(BeTrue())
}

func TestAddSameTypeTwiceDoesNotAddTwice(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeTypeSet()

	g.Expect(x.Add(StringType)).To(BeTrue())
	g.Expect(x.Add(StringType)).To(BeFalse())
	g.Expect(x.Len()).To(Equal(1))
}

func TestAddEquivalentTypeTwiceDoesNotAddTwice(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeTypeSet()

	g.Expect(x.Add(NewOptionalType(StringType))).To(BeTrue())
	g.Expect(x.Add(NewOptionalType(StringType))).To(BeFalse())
	g.Expect(x.Len()).To(Equal(1))
}
