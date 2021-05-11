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
 * ApplyTo() tests
 */

func TestTypeFlag_Wrap_GivenType_ReturnsWrappedType(t *testing.T) {
	g := NewGomegaWithT(t)
	ft := ARMFlag.ApplyTo(StringType)
	g.Expect(ft.element).To(Equal(StringType))
	g.Expect(ft.HasFlag(ARMFlag)).To(BeTrue())
}

/*
 * IsOn() tests
 */

func TestTypeFlag_IsOn_GivenType_ReturnsExpectedValue(t *testing.T) {

	armString := ARMFlag.ApplyTo(StringType)

	cases := []struct {
		name     string
		subject  Type
		flag     TypeFlag
		expected bool
	}{
		{"String does not have ArmFlag", StringType, ARMFlag, false},
		{"String does not have StorageFlag", StringType, StorageFlag, false},
		{"ArmString does have ArmFlag", armString, ARMFlag, true},
		{"ArmString does not have StorageFlag", armString, StorageFlag, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			g.Expect(c.flag.IsOn(c.subject)).To(Equal(c.expected))
		})
	}
}

/*
 * RemoveFrom() tests
 */

func TestTypeFlag_RemoveFrom_ReturnsExpectedValue(t *testing.T) {

	armString := ARMFlag.ApplyTo(StringType)
	armStorageString := StorageFlag.ApplyTo(armString)

	cases := []struct {
		name     string
		subject  Type
		flag     TypeFlag
		expected Type
	}{
		{"RemoveFrom Type with no flags returns the type unmodified", StringType, ARMFlag, StringType},
		{"RemoveFrom flag type for flag that isn't present returns the type unmodified", armString, StorageFlag, armString},
		{"RemoveFrom flag type with multiple flags returns flag type without specified flag", armStorageString, StorageFlag, armString},
		{"RemoveFrom flag type with only that flag returns element type", armString, ARMFlag, StringType},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			removed, err := c.flag.RemoveFrom(c.subject)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(removed.Equals(c.expected)).To(BeTrue())
		})
	}
}
