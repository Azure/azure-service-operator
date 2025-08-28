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
	t.Parallel()
	g := NewGomegaWithT(t)
	ft := NewFlaggedType(StringType, ARMFlag)

	g.Expect(ft).ToNot(BeNil())
	g.Expect(ft.HasFlag(ARMFlag)).To(BeTrue())
	g.Expect(ft.HasFlag(StorageFlag)).To(BeFalse())
}

func TestNewFlaggedType_GivenFlaggedType_DoesNotWrapTypeFurther(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, ARMFlag)
	outer := NewFlaggedType(inner, StorageFlag)
	g.Expect(outer).ToNot(BeNil())
	g.Expect(outer.HasFlag(ARMFlag)).To(BeTrue())
	g.Expect(outer.HasFlag(StorageFlag)).To(BeTrue())
}

/*
 * WithFlag() tests
 */

func TestFlaggedType_WithFlag_GivenFlag_ReturnsFlaggedTypeWithExpectedFlags(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, ARMFlag)
	outer := inner.WithFlag(StorageFlag)
	g.Expect(outer).ToNot(BeNil())
	g.Expect(outer.HasFlag(ARMFlag)).To(BeTrue())
	g.Expect(outer.HasFlag(StorageFlag)).To(BeTrue())
}

/*
 * WithoutFlag() tests
 */

func TestFlaggedType_WithoutFlag_GivenOnlyFlag_ReturnsWrappedType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, ARMFlag)
	final := inner.WithoutFlag(ARMFlag)
	g.Expect(final).To(Equal(StringType))
}

func TestFlaggedType_WithoutFlag_GivenExistingFlag_ReturnsFlaggedTypeWithExpectedFlags(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, ARMFlag, StorageFlag)
	final := inner.WithoutFlag(ARMFlag).(*FlaggedType)
	g.Expect(final.HasFlag(ARMFlag)).To(BeFalse())
	g.Expect(final.HasFlag(StorageFlag)).To(BeTrue())
}

func TestFlaggedType_WithoutFlag_GivenUnusedFlag_ReturnsSameInstance(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	inner := NewFlaggedType(StringType, StorageFlag)
	final := inner.WithoutFlag(ARMFlag).(*FlaggedType)
	g.Expect(final.HasFlag(ARMFlag)).To(BeFalse())
	g.Expect(final.HasFlag(StorageFlag)).To(BeTrue())
	g.Expect(final).To(Equal(inner))
}

/*
 * Equals() tests
 */

func TestFlaggedType_Equals_GivenOther_HasExpectedResult(t *testing.T) {
	t.Parallel()

	armString := NewFlaggedType(StringType, ARMFlag)
	armInt := NewFlaggedType(IntType, ARMFlag)
	storageString := NewFlaggedType(StringType, StorageFlag)
	manyFlagsString := NewFlaggedType(StringType, ARMFlag, StorageFlag, OneOfFlag)

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
			equal := TypeEquals(c.left, c.right)
			g.Expect(equal).To(Equal(c.expectedEqual))
		})
	}
}

/*
 * WithoutFlags() tests
 */

func TestFlaggedType_WithoutFlags_GivenNoFlags_ReturnsSameInstance(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewFlaggedType(StringType, ARMFlag, StorageFlag)
	result := original.WithoutFlags()
	g.Expect(result).To(BeIdenticalTo(original))
}

func TestFlaggedType_WithoutFlags_GivenNonMatchingFlags_ReturnsSameInstance(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewFlaggedType(StringType, ARMFlag, StorageFlag)
	result := original.WithoutFlags(OneOfFlag, DoNotPrune)
	g.Expect(result).To(BeIdenticalTo(original))
}

func TestFlaggedType_WithoutFlags_GivenPartiallyMatchingFlags_ReturnsNewFlaggedTypeWithRemainingFlags(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewFlaggedType(StringType, ARMFlag, StorageFlag, OneOfFlag)
	result := original.WithoutFlags(ARMFlag, DoNotPrune).(*FlaggedType)
	
	g.Expect(result).ToNot(BeIdenticalTo(original))
	g.Expect(result.HasFlag(ARMFlag)).To(BeFalse())
	g.Expect(result.HasFlag(StorageFlag)).To(BeTrue())
	g.Expect(result.HasFlag(OneOfFlag)).To(BeTrue())
	g.Expect(result.Element()).To(Equal(StringType))
}

func TestFlaggedType_WithoutFlags_GivenAllMatchingFlags_ReturnsUnderlyingType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewFlaggedType(StringType, ARMFlag, StorageFlag)
	result := original.WithoutFlags(ARMFlag, StorageFlag)
	
	g.Expect(result).To(Equal(StringType))
	g.Expect(result).ToNot(BeAssignableToTypeOf(&FlaggedType{}))
}

func TestFlaggedType_WithoutFlags_GivenAllMatchingFlagsAndMore_ReturnsUnderlyingType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewFlaggedType(StringType, ARMFlag, StorageFlag)
	result := original.WithoutFlags(ARMFlag, StorageFlag, OneOfFlag, DoNotPrune)
	
	g.Expect(result).To(Equal(StringType))
	g.Expect(result).ToNot(BeAssignableToTypeOf(&FlaggedType{}))
}

func TestFlaggedType_WithoutFlags_GivenSingleFlagMatch_ReturnsNewFlaggedTypeWithRemainingFlags(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewFlaggedType(StringType, ARMFlag, StorageFlag, OneOfFlag)
	result := original.WithoutFlags(StorageFlag).(*FlaggedType)
	
	g.Expect(result).ToNot(BeIdenticalTo(original))
	g.Expect(result.HasFlag(ARMFlag)).To(BeTrue())
	g.Expect(result.HasFlag(StorageFlag)).To(BeFalse())
	g.Expect(result.HasFlag(OneOfFlag)).To(BeTrue())
	g.Expect(result.Element()).To(Equal(StringType))
}

func TestFlaggedType_WithoutFlags_GivenSingleFlagType_ReturnsUnderlyingType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewFlaggedType(StringType, ARMFlag)
	result := original.WithoutFlags(ARMFlag)
	
	g.Expect(result).To(Equal(StringType))
	g.Expect(result).ToNot(BeAssignableToTypeOf(&FlaggedType{}))
}

func TestFlaggedType_WithoutFlags_GivenMixedMatchingAndNonMatchingFlags_ReturnsNewFlaggedTypeWithCorrectFlags(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewFlaggedType(StringType, ARMFlag, StorageFlag, OneOfFlag, WellknownFlag)
	result := original.WithoutFlags(StorageFlag, DoNotPrune, WellknownFlag, CompatibilityFlag).(*FlaggedType)
	
	g.Expect(result).ToNot(BeIdenticalTo(original))
	g.Expect(result.HasFlag(ARMFlag)).To(BeTrue())
	g.Expect(result.HasFlag(StorageFlag)).To(BeFalse())
	g.Expect(result.HasFlag(OneOfFlag)).To(BeTrue())
	g.Expect(result.HasFlag(WellknownFlag)).To(BeFalse())
	g.Expect(result.Element()).To(Equal(StringType))
}

/*
 * String() tests
 */

func TestFlaggedType_String_GivenTypeAndTag_ReturnsExpectedString(t *testing.T) {
	t.Parallel()

	flaggedString := NewFlaggedType(StringType, OneOfFlag)
	cases := []struct {
		name       string
		underlying Type
		flag       TypeFlag
		expected   string
	}{
		{"String with tag ARM", StringType, ARMFlag, "string[arm]"},
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
