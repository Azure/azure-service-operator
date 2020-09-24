/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestMakeLibraryPackageReference_GivenPath_ReturnsInstanceWithPath(t *testing.T) {
	cases := []struct {
		name string
		path string
	}{
		{"fmt library", "fmt"},
		{"ast library", "go/ast"},
		{"gomega library", "github.com/onsi/gomega"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			ref := MakeLibraryPackageReference(c.path)
			g.Expect(ref.PackagePath()).To(Equal(c.path))
		})
	}
}

func TestLibraryPackageReferences_ReturnExpectedProperties(t *testing.T) {
	cases := []struct {
		name        string
		path        string
		packageName string
	}{
		{"fmt library", "fmt", "fmt"},
		{"ast library", "go/ast", "ast"},
		{"gomega library", "github.com/onsi/gomega", "gomega"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := MakeLibraryPackageReference(c.path)
			_, err := ref.Group()

			g.Expect(ref.IsLocalPackage()).To(BeFalse())
			g.Expect(ref.Package()).To(Equal(c.packageName))
			g.Expect(ref.PackagePath()).To(Equal(c.path))
			g.Expect(ref.String()).To(Equal(c.path))
			g.Expect(err).NotTo(BeNil())
		})
	}
}

func TestLibraryPackageReferences_Equals_GivesExpectedResults(t *testing.T) {

	fmtRef := MakeLibraryPackageReference("fmt")
	astRef := MakeLibraryPackageReference("go/ast")
	otherRef := MakeLocalPackageReference("group", "package")

	cases := []struct {
		name     string
		this     LibraryPackageReference
		other    PackageReference
		areEqual bool
	}{
		{"Equal self", fmtRef, fmtRef, true},
		{"Equal self", astRef, astRef, true},
		{"Not equal other library", fmtRef, astRef, false},
		{"Not equal other library", astRef, fmtRef, false},
		{"Not equal other kind", fmtRef, otherRef, false},
		{"Not equal other kind", astRef, otherRef, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.this.Equals(c.other)
			g.Expect(areEqual).To(Equal(c.areEqual))
		})
	}
}
