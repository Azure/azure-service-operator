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
 * ApplyTo() tests
 */

func TestTypeFlag_Wrap_GivenType_ReturnsWrappedType(t *testing.T) {
	g := NewGomegaWithT(t)
	ft := ArmFlag.ApplyTo(StringType)
	g.Expect(ft.element).To(Equal(StringType))
	g.Expect(ft.HasFlag(ArmFlag)).To(BeTrue())
}

/*
 * IsOn() tests
 */

func TestTypeFlag_IsOn_GivenType_ReturnsExpectedValue(t *testing.T) {

	armString := ArmFlag.ApplyTo(StringType)

	cases := []struct {
		name     string
		subject  Type
		flag     TypeFlag
		expected bool
	}{
		{"String does not have ArmFlag", StringType, ArmFlag, false},
		{"String does not have StorageFlag", StringType, StorageFlag, false},
		{"ArmString does have ArmFlag", armString, ArmFlag, true},
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
