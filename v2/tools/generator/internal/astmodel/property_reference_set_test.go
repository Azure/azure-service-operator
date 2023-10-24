/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPropertyReferenceSet_NewPropertyReferenceSet_ReturnsEmptySet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPropertyReferenceSet()
	g.Expect(set.IsEmpty()).To(BeTrue())
}

func TestPropertyReferenceSet_AfterAdd_ContainsReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	pkg := makeTestLocalPackageReference("Demo", "v1")
	typ := MakeInternalTypeName(pkg, "Person")
	ref := MakePropertyReference(typ, "FullName")

	set := NewPropertyReferenceSet()
	set.Add(ref)

	g.Expect(set.Contains(ref)).To(BeTrue())
}

func TestPropertyReferenceSet_Except_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	pkg := makeTestLocalPackageReference("Demo", "v1")
	typ := MakeInternalTypeName(pkg, "Person")
	alpha := MakePropertyReference(typ, "Alpha")
	beta := MakePropertyReference(typ, "Beta")
	gamma := MakePropertyReference(typ, "Gamma")
	delta := MakePropertyReference(typ, "Delta")

	left := NewPropertyReferenceSet()
	left.Add(alpha, beta, gamma)

	right := NewPropertyReferenceSet()
	right.Add(beta, gamma, delta)

	actual := left.Except(right)

	g.Expect(actual.Contains(alpha)).To(BeTrue())
	g.Expect(actual.Contains(beta)).To(BeFalse())
	g.Expect(actual.Contains(gamma)).To(BeFalse())
	g.Expect(actual.Contains(delta)).To(BeFalse())
}
