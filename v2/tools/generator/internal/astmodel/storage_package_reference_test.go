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
		{"group", "1", "v1storage"},
		{"microsoft.network", "2018-05-01", "v20180501storage"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.group, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			localRef := makeTestLocalPackageReference(c.group, c.version)
			storageRef := MakeStoragePackageReference(localRef)

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
		storageRef    StoragePackageReference
		otherRef      PackageReference
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
			local := MakeLocalPackageReference("prefix", "microsoft.storage", GeneratorVersion, c.version)
			ref := MakeStoragePackageReference(local)

			g.Expect(ref.IsPreview()).To(Equal(c.isPreview))
		})
	}
}

func Test_StoragePackageReference_ImportAlias_ReturnsExpectedAlias(t *testing.T) {
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
		{"GeneratorVersionOnly", "storage", GeneratorVersion, "20200901", VersionOnly, "v20200901s"},
		{"GeneratorGroupOnly", "storage", GeneratorVersion, "20200901", GroupOnly, "storage"},
		{"GeneratorGroupAndVersion", "storage", GeneratorVersion, "20200901", GroupAndVersion, "storage_v20200901s"},
		{"GeneratorPreviewVersionOnly", "storage", GeneratorVersion, "20200901preview", VersionOnly, "v20200901ps"},
		{"GeneratorPreviewGroupOnly", "storage", GeneratorVersion, "20200901preview", GroupOnly, "storage"},
		{"GeneratorPreviewGroupAndVersion", "storage", GeneratorVersion, "20200901preview", GroupAndVersion, "storage_v20200901ps"},
		// Hard coded to v1api
		{"v1apiVersionOnly", "storage", "v1api", "20200901", VersionOnly, "v20200901s"},
		{"v1apiGroupOnly", "storage", "v1api", "20200901", GroupOnly, "storage"},
		{"v1apiGroupAndVersion", "storage", "v1api", "20200901", GroupAndVersion, "storage_v20200901s"},
		// Hard coded to v1beta
		{"v1betaVersionOnly", "storage", "v1beta", "20200901", VersionOnly, "v1beta20200901s"},
		{"v1betaGroupOnly", "storage", "v1beta", "20200901", GroupOnly, "storage"},
		{"v1betaGroupAndVersion", "storage", "v1beta", "20200901", GroupAndVersion, "storage_v1beta20200901s"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			lpr := MakeLocalPackageReference("v", c.group, c.generatorVersion, c.apiVersion)
			ref := MakeStoragePackageReference(lpr)
			g.Expect(ref.ImportAlias(c.style)).To(Equal(c.expected))
		})
	}
}
