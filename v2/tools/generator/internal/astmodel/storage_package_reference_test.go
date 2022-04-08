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

func TestStoragePackageReference_CreateImportAlias_GivenVersion_ReturnsExpectedResult(t *testing.T) {
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

			g.Expect(ref.CreateImportAlias(c.style)).To(Equal(c.expected))
		})
	}
}
