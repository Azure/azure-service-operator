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
	t.Parallel()
	g := NewGomegaWithT(t)

	pr := makeTestLocalPackageReference("group", "ver")
	pi := NewPackageImport(pr)

	g.Expect(pi.packageReference).To(Equal(pr))
	g.Expect(pi.name).To(BeEmpty())
}

/*
 * WithName() Tests
 */

func Test_PackageImportWithName_GivenName_SetsField(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := "foo"
	pr := makeTestLocalPackageReference("group", "ver")
	pi := NewPackageImport(pr).WithName(name)
	g.Expect(pi.name).To(Equal(name))
}

func Test_PackageImportWithName_GivenName_DoesNotModifyOriginal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pr := makeTestLocalPackageReference("group", "ver")
	original := NewPackageImport(pr)
	modified := original.WithName("foo")
	g.Expect(original.name).NotTo(Equal(modified.name))
}

func Test_PackageImportWithName_GivenName_ReturnsDifferentInstance(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pr := makeTestLocalPackageReference("group", "ver")
	original := NewPackageImport(pr)
	modified := original.WithName("foo")
	g.Expect(original.name).NotTo(Equal(modified.name))
}

func Test_PackageImportWithName_GivenExistingName_ReturnsEqualInstance(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := "foo"
	pr := makeTestLocalPackageReference("group", "ver")
	original := NewPackageImport(pr).WithName(name)
	modified := original.WithName(name)
	g.Expect(modified).To(Equal(original))
}

/*
 * Equals() tests
 */

func TestPackageImport_Equals(t *testing.T) {
	t.Parallel()

	var zeroPkgRef PackageImport
	localPkgRef := makeTestLocalPackageReference("group", "ver")
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

/*
 * CreateImportAlias Tests
 */

func TestPackageImport_CreateImportAlias_GivenLocalPackageReference_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		group    string
		version  string
		style    PackageImportStyle
		expected string
	}{
		// VersionOnly
		{"GA release", "batch", "2020-10-10", VersionOnly, "v20201010"},
		{"Preview release", "compute", "2020-10-10preview", VersionOnly, "v20201010p"},
		{"Alpha release", "network", "2020-10-10alpha", VersionOnly, "v20201010a"},
		{"Beta release", "cache", "2020-10-10beta", VersionOnly, "v20201010b"},
		// GroupOnly
		{"Batch", "batch", "2020-10-10", GroupOnly, "batch"},
		{"Compute", "compute", "2020-10-10", GroupOnly, "compute"},
		{"Network", "network", "2020-10-10", GroupOnly, "network"},
		// GroupAndVersion
		{"GA Batch Release", "batch", "2020-10-10", GroupAndVersion, "batch_v20201010"},
		{"Preview Compute Release", "compute", "2020-10-10preview", GroupAndVersion, "compute_v20201010p"},
		{"Alpha Network Release", "network", "2020-10-10alpha", GroupAndVersion, "network_v20201010a"},
		{"Beta Cache Release", "cache", "2020-10-10beta", GroupAndVersion, "cache_v20201010b"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := makeTestLocalPackageReference(c.group, c.version)
			imp := NewPackageImport(ref)

			g.Expect(imp.createImportAliasForLocalPackageReference(ref, c.style)).To(Equal(c.expected))
		})
	}
}

func TestPackageImport_CreateImportAlias_GivenStoragePackageReference_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		group    string
		version  string
		style    PackageImportStyle
		expected string
	}{
		// VersionOnly
		{"GA release", "batch", "2020-10-10", VersionOnly, "v20201010s"},
		{"Preview release", "compute", "2020-10-10preview", VersionOnly, "v20201010ps"},
		{"Alpha release", "network", "2020-10-10alpha", VersionOnly, "v20201010as"},
		{"Beta release", "cache", "2020-10-10beta", VersionOnly, "v20201010bs"},
		// GroupOnly
		{"Batch", "batch", "2020-10-10", GroupOnly, "batch"},
		{"Compute", "compute", "2020-10-10", GroupOnly, "compute"},
		{"Network", "network", "2020-10-10", GroupOnly, "network"},
		// GroupAndVersion
		{"GA Batch Release", "batch", "2020-10-10", GroupAndVersion, "batch_v20201010s"},
		{"Preview Compute Release", "compute", "2020-10-10preview", GroupAndVersion, "compute_v20201010ps"},
		{"Alpha Network Release", "network", "2020-10-10alpha", GroupAndVersion, "network_v20201010as"},
		{"Beta Cache Release", "cache", "2020-10-10beta", GroupAndVersion, "cache_v20201010bs"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := MakeStoragePackageReference(makeTestLocalPackageReference(c.group, c.version))
			imp := NewPackageImport(ref)

			g.Expect(imp.createImportAliasForStoragePackageReference(ref, c.style)).To(Equal(c.expected))
		})
	}
}
