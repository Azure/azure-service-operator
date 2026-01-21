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

	cases := map[string]struct {
		group           string
		version         string
		expectedVersion string
	}{
		"group": {
			group:           "group",
			version:         "1",
			expectedVersion: "storage",
		},
		"microsoft.network": {
			group:           "microsoft.network",
			version:         "2018-05-01",
			expectedVersion: "storage",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
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

	cases := map[string]struct {
		storageRef    InternalPackageReference
		otherRef      InternalPackageReference
		expectedEqual bool
	}{
		"Equal to self": {
			storageRef:    storageRef,
			otherRef:      storageRef,
			expectedEqual: true,
		},
		"Equal to other": {
			storageRef:    storageRef,
			otherRef:      otherRef,
			expectedEqual: true,
		},
		"Equal to other (reversed)": {
			storageRef:    otherRef,
			otherRef:      storageRef,
			expectedEqual: true,
		},
		"Not equal to local": {
			storageRef:    storageRef,
			otherRef:      localRef,
			expectedEqual: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.storageRef.Equals(c.otherRef)

			g.Expect(areEqual).To(Equal(c.expectedEqual))
		})
	}
}

func TestStoragePackageReferenceIsPreview(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		version   string
		isPreview bool
	}{
		"GA storage Release is not preview": {
			version:   "v20200901",
			isPreview: false,
		},
		"Preview storage release is preview": {
			version:   "v20200901preview",
			isPreview: true,
		},
		"Preview storage re-release is preview": {
			version:   "v20200901preview2",
			isPreview: true,
		},
		"Alpha storage release is preview": {
			version:   "v20200901alpha",
			isPreview: true,
		},
		"Beta storage release is preview": {
			version:   "v20200901betas",
			isPreview: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
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
