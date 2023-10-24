/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

var t1 = MakeInternalTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "t1")
var t2 = MakeInternalTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "t2")

func TestEmptyTypeAssociation_AreEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := make(TypeAssociation)
	y := make(TypeAssociation)

	g.Expect(x).To(Equal(y))
	g.Expect(y).To(Equal(x))
}

func TestIdenticalTypeAssociation_AreEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := make(TypeAssociation)
	x[t1] = t2
	x[t2] = t2
	y := make(TypeAssociation)
	y[t1] = t2
	y[t2] = t2

	g.Expect(x.Equals(y)).To(BeTrue())
	g.Expect(y.Equals(x)).To(BeTrue())
}

func TestDifferentTypeAssociation_AreNotEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := make(TypeAssociation)
	x[t1] = t2
	y := make(TypeAssociation)
	y[t1] = t2
	y[t2] = t2

	g.Expect(x.Equals(y)).To(BeFalse())
	g.Expect(y.Equals(x)).To(BeFalse())
}
