/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

/*
 * NewFlaggedType() tests
 */

func TestNewFlaggedType_GivenTypeAndFlag_ReturnsFlaggedTypeWithFlag(t *testing.T) {
	g := NewGomegaWithT(t)
	ft := NewFlaggedType(StringType, ArmFlag)

	g.Expect(ft).ToNot(BeNil())
	g.Expect(ft.HasFlag(ArmFlag)).To(BeTrue())
	g.Expect(ft.HasFlag(StorageFlag)).To(BeFalse())
}

func TestNewFlaggedType_GivenFlaggedType_DoesNotWrapTypeFurther(t *testing.T) {
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, ArmFlag)
	outer := NewFlaggedType(inner, StorageFlag)
	g.Expect(outer).ToNot(BeNil())
	g.Expect(outer.HasFlag(ArmFlag)).To(BeTrue())
	g.Expect(outer.HasFlag(StorageFlag)).To(BeTrue())
}

/*
 * WithFlag() tests
 */

func TestFlaggedType_WithFlag_GivenFlag_ReturnsFlaggedTypeWithExpectedFlags(t *testing.T) {
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, ArmFlag)
	outer := inner.WithFlag(StorageFlag)
	g.Expect(outer).ToNot(BeNil())
	g.Expect(outer.HasFlag(ArmFlag)).To(BeTrue())
	g.Expect(outer.HasFlag(StorageFlag)).To(BeTrue())
}

/*
 * WithoutFlag() tests
 */

func TestFlaggedType_WithoutFlag_GivenOnlyFlag_ReturnsWrappedType(t *testing.T) {
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, ArmFlag)
	final := inner.WithoutFlag(ArmFlag)
	g.Expect(final).To(Equal(StringType))
}

func TestFlaggedType_WithoutFlag_GivenExistingFlag_ReturnsFlaggedTypeWithExpectedFlags(t *testing.T) {
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, ArmFlag, StorageFlag)
	final := inner.WithoutFlag(ArmFlag).(*FlaggedType)
	g.Expect(final.HasFlag(ArmFlag)).To(BeFalse())
	g.Expect(final.HasFlag(StorageFlag)).To(BeTrue())
}

func TestFlaggedType_WithoutFlag_GivenUnusedFlag_ReturnsSameInstance(t *testing.T) {
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, StorageFlag)
	final := inner.WithoutFlag(ArmFlag).(*FlaggedType)
	g.Expect(final.HasFlag(ArmFlag)).To(BeFalse())
	g.Expect(final.HasFlag(StorageFlag)).To(BeTrue())
	g.Expect(final).To(Equal(inner))
}

/*
 * Equals() tests
 */

func TestFlaggedType_Equals_GivenOther_HasExpectedResult(t *testing.T) {
	armString := NewFlaggedType(StringType, ArmFlag)
	armInt := NewFlaggedType(IntType, ArmFlag)
	storageString := NewFlaggedType(StringType, StorageFlag)
	manyFlagsString := NewFlaggedType(StringType, ArmFlag, StorageFlag, OneOfFlag)

	cases := []struct {
		name          string
		left          *FlaggedType
		right         *FlaggedType
		expectedEqual bool
	}{
		{"Equal to self (i)", armString, armString, true},
		{"Equal to self (ii)", storageString, storageString, true},
		{"Equal to self (iii)", armInt, armInt, true},
		{"Equal to self (iv)", manyFlagsString, manyFlagsString, true},
		{"Different if only flag is different (i)", armString, storageString, false},
		{"Different if only flag is different (ii)", storageString, armString, false},
		{"Different if # flags different (i)", manyFlagsString, storageString, false},
		{"Different if # flags different (ii)", manyFlagsString, armString, false},
		{"Different if # flags different (iii)", storageString, manyFlagsString, false},
		{"Different if # flags different (iv)", armString, manyFlagsString, false},
		{"Different if underlying type different (i)", armString, armInt, false},
		{"Different if underlying type different (ii)", armInt, armString, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			equal := c.left.Equals(c.right)
			g.Expect(equal).To(Equal(c.expectedEqual))
		})
	}
}

/*
 * String() tests
 */

func TestFlaggedType_String_GivenTypeAndTag_ReturnsExpectedString(t *testing.T) {
	flaggedString := NewFlaggedType(StringType, OneOfFlag)
	cases := []struct {
		name       string
		underlying Type
		flag       TypeFlag
		expected   string
	}{
		{"String with tag ARM", StringType, ArmFlag, "string[arm]"},
		{"Bool with tag Storage", BoolType, StorageFlag, "bool[storage]"},
		{"String with multiple tags", flaggedString, StorageFlag, "string[oneof|storage]"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ft := NewFlaggedType(c.underlying, c.flag)
			s := ft.String()
			g.Expect(s).To(Equal(c.expected))
		})
	}
}
