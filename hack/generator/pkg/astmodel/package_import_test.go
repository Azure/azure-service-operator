/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPackageImport_Equals(t *testing.T) {
	localPkgRef := NewLocalPackageReference("group", "ver")
	localPkgImport := NewPackageImport(*localPkgRef)

	cases := []struct {
		name     string
		lhs      *PackageImport
		rhs      *PackageImport
		expected bool
	}{
		{"package import is equal to itself", localPkgImport, localPkgImport, true},
		{"package import is equal to same import different reference", NewPackageImport(*localPkgRef), NewPackageImport(*localPkgRef), true},
		{"package import is not equal to import with name", localPkgImport, localPkgImport.WithName("ref"), false},
		{"package import differs by name is not equal", localPkgImport.WithName("ref1"), localPkgImport.WithName("ref2"), false},
		{"package imports with same name are equal", localPkgImport.WithName("ref"), localPkgImport.WithName("ref"), true},
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
