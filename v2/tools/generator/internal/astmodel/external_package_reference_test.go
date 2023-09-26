/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestExternalPackageReferences_ReturnExpectedProperties(t *testing.T) {
	t.Parallel()

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

			ref := MakeExternalPackageReference(c.path)
			g.Expect(ref.String()).To(Equal(c.path))
		})
	}
}

func TestExternalPackageReferences_Equals_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	fmtRef := MakeExternalPackageReference("fmt")
	astRef := MakeExternalPackageReference("go/ast")
	otherRef := makeTestLocalPackageReference("group", "package")

	cases := []struct {
		name     string
		this     ExternalPackageReference
		other    PackageReference
		areEqual bool
	}{
		{"Equal self", fmtRef, fmtRef, true},
		{"Equal self", astRef, astRef, true},
		{"Not equal other external reference", fmtRef, astRef, false},
		{"Not equal other external reference", astRef, fmtRef, false},
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
