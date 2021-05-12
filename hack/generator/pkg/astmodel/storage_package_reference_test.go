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

	cases := []struct {
		group           string
		version         string
		expectedVersion string
	}{
		{"group", "v1", "v1storage"},
		{"microsoft.network", "v20180501", "v20180501storage"},
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

			ref := MakeStoragePackageReference(
				makeTestLocalPackageReference(
					"microsoft.storage",
					CreateLocalPackageNameFromVersion(c.version)))

			g.Expect(ref.IsPreview()).To(Equal(c.isPreview))
		})
	}
}
