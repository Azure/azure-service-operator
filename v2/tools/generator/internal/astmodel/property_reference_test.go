/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPropertyReference_IsEmpty_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	pkg := makeTestLocalPackageReference("Demo", "v1")
	declaringType := MakeTypeName(pkg, "Person")

	cases := []struct {
		name          string
		declaringType TypeName
		property      PropertyName
		expected      bool
	}{
		{"Property name present means not empty", nil, "property", false},
		{"Declaring type present means not empty", declaringType, "", false},
		{"Fully populated means not empty", declaringType, "property", false},
		{"Totally empty not empty", nil, "", true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			ref := MakePropertyReference(c.declaringType, c.property)
			g.Expect(ref.IsEmpty()).To(Equal(c.expected))
		})
	}
}

func TestPropertyReference_String_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	pkg := makeTestLocalPackageReference("Demo", "1")
	declaringType := MakeTypeName(pkg, "Person")
	property := PropertyName("FullName")

	ref := MakePropertyReference(declaringType, property)
	str := ref.String()

	g.Expect(str).To(Equal("Demo/v1/Person.FullName"))
}
