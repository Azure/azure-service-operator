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
	result := MakeVersionedLocalPackageReference("github.com/Azure/azure-service-operator/v2/api", group, version).
		WithVersionPrefix("v")
	return result
}

func TestMakeLocalPackageReference_GivenGroupAndPackage_ReturnsInstanceWithProperties(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		group      string
		apiVersion string
		pkg        string
	}{
		"Networking": {
			group:      "network",
			apiVersion: "2020-09-01",
			pkg:        "v20200901",
		},
		"Batch (new)": {
			group:      "batch",
			apiVersion: "2020-09-01",
			pkg:        "v20200901",
		},
		"Batch (old)": {
			group:      "batch",
			apiVersion: "2015-01-01",
			pkg:        "v20150101",
		},
		"Networking Frontdoor": {
			group:      "network.frontdoor",
			apiVersion: "2020-09-01",
			pkg:        "v20200901",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := makeTestLocalPackageReference(c.group, c.apiVersion)
			g.Expect(ref.Group()).To(Equal(c.group))
			g.Expect(ref.PackageName()).To(Equal(c.pkg))
			g.Expect(ref.APIVersion()).To(Equal(c.apiVersion))
		})
	}
}

func TestLocalPackageReferences_ReturnExpectedProperties(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		group       string
		version     string
		pkg         string
		packagePath string
		folderPath  string
	}{
		"Network": {
			group:       "network",
			version:     "2020-09-01",
			pkg:         "v20200901",
			packagePath: "github.com/Azure/azure-service-operator/v2/api/network/v20200901",
			folderPath:  "network/v20200901",
		},
		"Batch (new)": {
			group:       "batch",
			version:     "2020-09-01",
			pkg:         "v20200901",
			packagePath: "github.com/Azure/azure-service-operator/v2/api/batch/v20200901",
			folderPath:  "batch/v20200901",
		},
		"Batch (old)": {
			group:       "batch",
			version:     "2015-01-01",
			pkg:         "v20150101",
			packagePath: "github.com/Azure/azure-service-operator/v2/api/batch/v20150101",
			folderPath:  "batch/v20150101",
		},
		"Network Frontdoor": {
			group:       "network.frontdoor",
			version:     "2020-09-01",
			pkg:         "v20200901",
			packagePath: "github.com/Azure/azure-service-operator/v2/api/network.frontdoor/v20200901",
			folderPath:  "network.frontdoor/v20200901",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
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

	batchRef := makeTestLocalPackageReference("batch", "v20200901")
	olderRef := makeTestLocalPackageReference("batch", "v20150101")
	networkRef := makeTestLocalPackageReference("network", "v20200901")
	fmtRef := MakeExternalPackageReference("fmt")

	cases := map[string]struct {
		this     LocalPackageReference
		other    PackageReference
		areEqual bool
	}{
		"Equal self (batch)": {
			this:     batchRef,
			other:    batchRef,
			areEqual: true,
		},
		"Equal self (older)": {
			this:     olderRef,
			other:    olderRef,
			areEqual: true,
		},
		"Not equal other library name (batch vs network)": {
			this:     batchRef,
			other:    networkRef,
			areEqual: false,
		},
		"Not equal other library name (network vs batch)": {
			this:     networkRef,
			other:    batchRef,
			areEqual: false,
		},
		"Not equal other library version (batch vs older)": {
			this:     batchRef,
			other:    olderRef,
			areEqual: false,
		},
		"Not equal other library version (older vs batch)": {
			this:     olderRef,
			other:    batchRef,
			areEqual: false,
		},
		"Not equal other kind (batch vs fmt)": {
			this:     batchRef,
			other:    fmtRef,
			areEqual: false,
		},
		"Not equal other kind (network vs fmt)": {
			this:     networkRef,
			other:    fmtRef,
			areEqual: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.this.Equals(c.other)
			g.Expect(areEqual).To(Equal(c.areEqual))
		})
	}
}

func TestLocalPackageReferenceIsPreview(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		version   string
		isPreview bool
	}{
		"GA Release is not preview": {
			version:   "v20200901",
			isPreview: false,
		},
		"Preview release is preview": {
			version:   "v20200901preview",
			isPreview: true,
		},
		"Preview rerelease is preview": {
			version:   "v20200901preview2",
			isPreview: true,
		},
		"Alpha release is preview": {
			version:   "v20200901alpha",
			isPreview: true,
		},
		"Beta release is preview": {
			version:   "v20200901beta",
			isPreview: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			// Using GeneratorVersion here to make sure IsPreview isn't fooled
			ref := MakeVersionedLocalPackageReference("prefix", "storage", c.version)

			g.Expect(ref.IsPreview()).To(Equal(c.isPreview))
		})
	}
}

func Test_LocalPackageReference_ImportAlias_ReturnsExpectedAlias(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		group      string
		apiVersion string
		style      PackageImportStyle
		expected   string
	}{
		// Current generator version
		"GeneratorVersionOnly": {
			group:      "storage",
			apiVersion: "20200901",
			style:      VersionOnly,
			expected:   "v20200901",
		},
		"GeneratorGroupOnly": {
			group:      "storage",
			apiVersion: "20200901",
			style:      GroupOnly,
			expected:   "storage",
		},
		"GeneratorGroupAndVersion": {
			group:      "storage",
			apiVersion: "20200901",
			style:      GroupAndVersion,
			expected:   "storage_v20200901",
		},
		"GeneratorPreviewVersionOnly": {
			group:      "storage",
			apiVersion: "20200901preview",
			style:      VersionOnly,
			expected:   "v20200901p",
		},
		"GeneratorPreviewGroupOnly": {
			group:      "storage",
			apiVersion: "20200901preview",
			style:      GroupOnly,
			expected:   "storage",
		},
		"GeneratorPreviewGroupAndVersion": {
			group:      "storage",
			apiVersion: "20200901preview",
			style:      GroupAndVersion,
			expected:   "storage_v20200901p",
		},
		// Current generator version with dot in the group name
		"DotGroupGeneratorVersionOnly": {
			group:      "network.frontdoor",
			apiVersion: "20200901",
			style:      VersionOnly,
			expected:   "v20200901",
		},
		"DotGroupGeneratorGroupOnly": {
			group:      "network.frontdoor",
			apiVersion: "20200901",
			style:      GroupOnly,
			expected:   "networkfrontdoor",
		},
		"DotGroupGeneratorGroupAndVersion": {
			group:      "network.frontdoor",
			apiVersion: "20200901",
			style:      GroupAndVersion,
			expected:   "networkfrontdoor_v20200901",
		},
		// Hard coded to v1api
		"v1apiVersionOnly": {
			group:      "storage",
			apiVersion: "20200901",
			style:      VersionOnly,
			expected:   "v20200901",
		},
		"v1apiGroupOnly": {
			group:      "storage",
			apiVersion: "20200901",
			style:      GroupOnly,
			expected:   "storage",
		},
		"v1apiGroupAndVersion": {
			group:      "storage",
			apiVersion: "20200901",
			style:      GroupAndVersion,
			expected:   "storage_v20200901",
		},
		"v1apiGroupAndFullVersion": {
			group:      "storage",
			apiVersion: "20200901",
			style:      GroupAndFullVersion,
			expected:   "storage_v1api20200901",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := MakeVersionedLocalPackageReference("v", c.group, c.apiVersion)
			g.Expect(ref.ImportAlias(c.style)).To(Equal(c.expected))
		})
	}
}
