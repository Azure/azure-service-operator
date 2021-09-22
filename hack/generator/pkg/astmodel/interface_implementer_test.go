/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func makeInterfaceImplForTest() *InterfaceImplementation {
	pr := makeTestLocalPackageReference("group", "package")

	return NewInterfaceImplementation(MakeTypeName(pr, "foo"))
}

func Test_InterfaceImplementerWithInterface_ReturnsNewInstance(t *testing.T) {

	g := NewGomegaWithT(t)

	original := InterfaceImplementer{}
	updated := original.WithInterface(makeInterfaceImplForTest())

	g.Expect(updated).ToNot(Equal(original))
	g.Expect(len(updated.interfaces)).To(Equal(1))
}
