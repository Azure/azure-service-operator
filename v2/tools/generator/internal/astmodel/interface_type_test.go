/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestNewInterfaceType_ReturnsEmptyType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	st := NewInterfaceType()
	g.Expect(st.functions.Len()).To(Equal(0))
}

/*
 * Equals() tests
 */

func TestInterfaceType_Equals_WhenGivenType_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	funcA := NewFakeFunction("A")
	funcB := NewFakeFunction("B")

	interfaceA := NewInterfaceType().WithFunction(funcA)
	otherInterfaceA := NewInterfaceType().WithFunction(funcA)
	interfaceB := NewInterfaceType().WithFunction(funcB)
	interfaceAB := NewInterfaceType().WithFunction(funcA).WithFunction(funcB)
	interfaceBA := NewInterfaceType().WithFunction(funcB).WithFunction(funcA)
	mapType := NewMapType(StringType, StringType)

	cases := []struct {
		name      string
		thisType  Type
		otherType Type
		expected  bool
	}{
		{"Equal to self", interfaceA, interfaceA, true},
		{"Equal to same", interfaceA, otherInterfaceA, true},
		{"Equal when properties reordered", interfaceAB, interfaceBA, true},
		{"Not-equal when properties missing", interfaceAB, interfaceA, false},

		{"Not-equal when different type", interfaceB, mapType, false},
		{"Not-equal when different function", interfaceA, interfaceB, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := TypeEquals(c.thisType, c.otherType)

			g.Expect(areEqual).To(Equal(c.expected))
		})
	}
}

/*
 * WithFunction() tests
 */

func Test_WithFunction_GivenEmptyInterface_ReturnsPopulatedInterface(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	empty := EmptyInterfaceType
	fn := NewFakeFunction("Activate")
	iface := empty.WithFunction(fn)
	g.Expect(empty).NotTo(Equal(iface)) // Ensure the original wasn't modified
	g.Expect(iface.functions.Len()).To(Equal(1))

	activateFn, ok := iface.functions.Get("Activate")
	g.Expect(ok).To(BeTrue())
	g.Expect(activateFn.Equals(fn, EqualityOverrides{})).To(BeTrue())
}
