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
 * NewPackageImport() Tests
 */

func Test_NewPackageImport_GivenValues_InitializesFields(t *testing.T) {
	g := NewGomegaWithT(t)

	pr := MakeLocalPackageReference("group", "ver")
	pi := NewPackageImport(pr)

	g.Expect(pi.PackageReference).To(Equal(pr))
	g.Expect(pi.name).To(BeEmpty())
}

/*
 * WithName() Tests
 */

func Test_PackageImportWithName_GivenName_SetsField(t *testing.T) {
	g := NewGomegaWithT(t)

	name := "foo"
	pr := MakeLocalPackageReference("group", "ver")
	pi := NewPackageImport(pr).WithName(name)
	g.Expect(pi.name).To(Equal(name))
}

func Test_PackageImportWithName_GivenName_DoesNotModifyOriginal(t *testing.T) {
	g := NewGomegaWithT(t)

	pr := MakeLocalPackageReference("group", "ver")
	original := NewPackageImport(pr)
	modified := original.WithName("foo")
	g.Expect(original.name).NotTo(Equal(modified.name))
}

func Test_PackageImportWithName_GivenName_ReturnsDifferentInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	pr := MakeLocalPackageReference("group", "ver")
	original := NewPackageImport(pr)
	modified := original.WithName("foo")
	g.Expect(original.name).NotTo(Equal(modified.name))
}

func Test_PackageImportWithName_GivenExistingName_ReturnsEqualInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	name := "foo"
	pr := MakeLocalPackageReference("group", "ver")
	original := NewPackageImport(pr).WithName(name)
	modified := original.WithName(name)
	g.Expect(modified).To(Equal(original))
}

/*
 * Equals() tests
 */

func TestPackageImport_Equals(t *testing.T) {
	var zeroPkgRef PackageImport
	localPkgRef := MakeLocalPackageReference("group", "ver")
	localPkgImport := NewPackageImport(localPkgRef)

	cases := []struct {
		name     string
		lhs      PackageImport
		rhs      PackageImport
		expected bool
	}{
		{"package import is equal to itself", localPkgImport, localPkgImport, true},
		{"package import is equal to same import different reference", NewPackageImport(localPkgRef), NewPackageImport(localPkgRef), true},
		{"package import is not equal to import with name", localPkgImport, localPkgImport.WithName("ref"), false},
		{"package import differs by name is not equal", localPkgImport.WithName("ref1"), localPkgImport.WithName("ref2"), false},
		{"package imports with same name are equal", localPkgImport.WithName("ref"), localPkgImport.WithName("ref"), true},
		{"two zero imports are equal", zeroPkgRef, zeroPkgRef, true},
		{"zero not equal to other reference", zeroPkgRef, localPkgImport, false},
		{"other reference not equal to zero", localPkgImport, zeroPkgRef, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.lhs.Equals(c.rhs)

			g.Expect(areEqual).To(Equal(c.expected))
		})
	}
}
