/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func makeTestLocalPackageReference(group string, version string) LocalPackageReference {
	// We use a fixed path and version prefixes to ensure consistency across testing
	// For convenience, we tolerate the prefix already being present
	version = strings.TrimPrefix(version, "v")
	return MakeLocalPackageReference("github.com/Azure/azure-service-operator/v2/api", group, "v", version)
}

func TestMakeLocalPackageReference_GivenGroupAndPackage_ReturnsInstanceWithProperties(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name       string
		group      string
		apiVersion string
		pkg        string
	}{
		{"Networking", "microsoft.networking", "2020-09-01", "v20200901"},
		{"Batch (new)", "microsoft.batch", "2020-09-01", "v20200901"},
		{"Batch (old)", "microsoft.batch", "2015-01-01", "v20150101"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := makeTestLocalPackageReference(c.group, c.apiVersion)
			g.Expect(ref.Group()).To(Equal(c.group))
			g.Expect(ref.PackageName()).To(Equal(c.pkg))
			g.Expect(ref.ApiVersion()).To(Equal(c.apiVersion))
		})
	}
}

func TestLocalPackageReferences_ReturnExpectedProperties(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		group       string
		version     string
		pkg         string
		packagePath string
		folderPath  string
	}{
		{
			"Networking",
			"microsoft.networking",
			"2020-09-01",
			"v20200901",
			"github.com/Azure/azure-service-operator/v2/api/microsoft.networking/v20200901",
			"microsoft.networking/v20200901",
		},
		{
			"Batch (new)",
			"microsoft.batch",
			"2020-09-01",
			"v20200901",
			"github.com/Azure/azure-service-operator/v2/api/microsoft.batch/v20200901",
			"microsoft.batch/v20200901",
		},
		{
			"Batch (old)",
			"microsoft.batch",
			"2015-01-01",
			"v20150101",
			"github.com/Azure/azure-service-operator/v2/api/microsoft.batch/v20150101",
			"microsoft.batch/v20150101",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := makeTestLocalPackageReference(c.group, c.version)
			g.Expect(ref.PackageName()).To(Equal(c.pkg))
			g.Expect(ref.PackagePath()).To(Equal(c.packagePath))
			g.Expect(ref.String()).To(Equal(c.packagePath))
			g.Expect(ref.Group()).To(Equal(c.group))
			g.Expect(ref.FolderPath()).To(Equal(c.folderPath))
		})
	}
}

func TestLocalPackageReferences_Equals_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	batchRef := makeTestLocalPackageReference("microsoft.batch", "v20200901")
	olderRef := makeTestLocalPackageReference("microsoft.batch", "v20150101")
	networkingRef := makeTestLocalPackageReference("microsoft.networking", "v20200901")
	fmtRef := MakeExternalPackageReference("fmt")

	cases := []struct {
		name     string
		this     LocalPackageReference
		other    PackageReference
		areEqual bool
	}{
		{"Equal self", batchRef, batchRef, true},
		{"Equal self", olderRef, olderRef, true},
		{"Not equal other library name", batchRef, networkingRef, false},
		{"Not equal other library name", networkingRef, batchRef, false},
		{"Not equal other library version", batchRef, olderRef, false},
		{"Not equal other library version", olderRef, batchRef, false},
		{"Not equal other kind", batchRef, fmtRef, false},
		{"Not equal other kind", networkingRef, fmtRef, false},
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

func TestLocalPackageReferenceIsPreview(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		version   string
		isPreview bool
	}{
		{"GA Release is not preview", "v20200901", false},
		{"Preview release is preview", "v20200901preview", true},
		{"Preview rerelease is preview", "v20200901preview2", true},
		{"Alpha release is preview", "v20200901alpha", true},
		{"Beta release is preview", "v20200901beta", true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			// Using GeneratorVersion here to make sure IsPreview isn't fooled
			ref := MakeLocalPackageReference("prefix", "storage", GeneratorVersion, c.version)

			g.Expect(ref.IsPreview()).To(Equal(c.isPreview))
		})
	}
}

func Test_LocalPackageReference_ImportAlias_ReturnsExpectedAlias(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name             string
		group            string
		generatorVersion string
		apiVersion       string
		style            PackageImportStyle
		expected         string
	}{
		// Current generator version
		{"GeneratorVersionOnly", "storage", GeneratorVersion, "20200901", VersionOnly, "v20200901"},
		{"GeneratorGroupOnly", "storage", GeneratorVersion, "20200901", GroupOnly, "storage"},
		{"GeneratorGroupAndVersion", "storage", GeneratorVersion, "20200901", GroupAndVersion, "storage_v20200901"},
		{"GeneratorPreviewVersionOnly", "storage", GeneratorVersion, "20200901preview", VersionOnly, "v20200901p"},
		{"GeneratorPreviewGroupOnly", "storage", GeneratorVersion, "20200901preview", GroupOnly, "storage"},
		{"GeneratorPreviewGroupAndVersion", "storage", GeneratorVersion, "20200901preview", GroupAndVersion, "storage_v20200901p"},
		// Hard coded to v1api
		{"v1apiVersionOnly", "storage", "v1api", "20200901", VersionOnly, "v20200901"},
		{"v1apiGroupOnly", "storage", "v1api", "20200901", GroupOnly, "storage"},
		{"v1apiGroupAndVersion", "storage", "v1api", "20200901", GroupAndVersion, "storage_v20200901"},
		// Hard coded to v1beta
		{"v1betaVersionOnly", "storage", "v1beta", "20200901", VersionOnly, "v1beta20200901"},
		{"v1betaGroupOnly", "storage", "v1beta", "20200901", GroupOnly, "storage"},
		{"v1betaGroupAndVersion", "storage", "v1beta", "20200901", GroupAndVersion, "storage_v1beta20200901"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := MakeLocalPackageReference("v", c.group, c.generatorVersion, c.apiVersion)
			g.Expect(ref.ImportAlias(c.style)).To(Equal(c.expected))
		})
	}
}
