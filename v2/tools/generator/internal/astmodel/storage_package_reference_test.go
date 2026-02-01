/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestMakeStoragePackageReference(t *testing.T) {
	t.Parallel()

	cases := []struct {
		group           string
		version         string
		expectedVersion string
	}{
		{"group", "1", "storage"},
		{"microsoft.network", "2018-05-01", "storage"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.group, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			localRef := makeTestLocalPackageReference(c.group, c.version)
			storageRef := MakeStoragePackageReference(localRef)

			g.Expect(storageRef).To(BeAssignableToTypeOf(SubPackageReference{}))
			g.Expect(storageRef.PackageName()).To(Equal(c.expectedVersion))
		})
	}
}

func TestStoragePackageReferenceEquals(t *testing.T) {
	t.Parallel()

	localRef := makeTestLocalPackageReference("group", "v1")
	storageRef := MakeStoragePackageReference(localRef)
	otherRef := MakeStoragePackageReference(localRef)

	cases := []struct {
		name          string
		storageRef    InternalPackageReference
		otherRef      InternalPackageReference
		expectedEqual bool
	}{
		{"Equal to self", storageRef, storageRef, true},
		{"Equal to other", storageRef, otherRef, true},
		{"Equal to other (reversed)", otherRef, storageRef, true},
		{"Not equal to local", storageRef, localRef, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.storageRef.Equals(c.otherRef)

			g.Expect(areEqual).To(Equal(c.expectedEqual))
		})
	}
}

func TestStoragePackageReferenceIsPreview(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		version   string
		isPreview bool
	}{
		{"GA storage Release is not preview", "v20200901", false},
		{"Preview storage release is preview", "v20200901preview", true},
		{"Preview storage re-release is preview", "v20200901preview2", true},
		{"Alpha storage release is preview", "v20200901alpha", true},
		{"Beta storage release is preview", "v20200901betas", true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			// Using GeneratorVersion here to make sure IsPreview isn't fooled
			local := MakeVersionedLocalPackageReference("prefix", "microsoft.storage", c.version)
			ref := MakeStoragePackageReference(local)

			g.Expect(ref.IsPreview()).To(Equal(c.isPreview))
		})
	}
}

func Test_StoragePackageReference_ImportAlias_ReturnsExpectedAlias(t *testing.T) {
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
			expected:   "v20200901s",
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
			expected:   "storage_v20200901s",
		},
		"GeneratorPreviewVersionOnly": {
			group:      "storage",
			apiVersion: "20200901preview",
			style:      VersionOnly,
			expected:   "v20200901ps",
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
			expected:   "storage_v20200901ps",
		},
		// Hard coded to v1api
		"v1apiVersionOnly": {
			group:      "storage",
			apiVersion: "20200901",
			style:      VersionOnly,
			expected:   "v20200901s",
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
			expected:   "storage_v20200901s",
		},
		"v1apiGroupAndFullVersion": {
			group:      "storage",
			apiVersion: "20200901",
			style:      GroupAndFullVersion,
			expected:   "storage_v1api20200901s",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			lpr := MakeVersionedLocalPackageReference("v", c.group, c.apiVersion)
			ref := MakeStoragePackageReference(lpr)
			g.Expect(ref.ImportAlias(c.style)).To(Equal(c.expected))
		})
	}
}
