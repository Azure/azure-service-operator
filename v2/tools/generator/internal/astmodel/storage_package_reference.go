/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
)

const (
	StoragePackageSuffix = "storage"
)

type StoragePackageReference struct {
	inner LocalPackageReference // a reference to the API package this storage package mirrors
}

var _ PackageReference = StoragePackageReference{}

// MakeStoragePackageReference creates a new storage package reference from a local package reference
func MakeStoragePackageReference(local LocalPackageReference) StoragePackageReference {
	return StoragePackageReference{
		inner: local,
	}
}

// PackageName returns the package name of this reference
func (s StoragePackageReference) PackageName() string {
	return s.Version()
}

// PackagePath returns the fully qualified package path
func (s StoragePackageReference) PackagePath() string {
	url := s.inner.localPathPrefix + "/" + s.inner.group + "/" + s.Version()
	return url
}

func (s StoragePackageReference) Version() string {
	return s.inner.Version() + StoragePackageSuffix
}

// Equals returns true if the passed package reference is a storage package reference wrapping an identical local package reference
func (s StoragePackageReference) Equals(ref PackageReference) bool {
	other, ok := ref.(StoragePackageReference)
	if !ok {
		return false
	}

	return s.inner.Equals(other.inner)
}

// String returns the string representation of the package reference
func (s StoragePackageReference) String() string {
	return fmt.Sprintf("storage:%s/%s", s.inner.group, s.Version())
}

// IsPreview returns true if this package reference is a preview
func (s StoragePackageReference) IsPreview() bool {
	return s.inner.IsPreview()
}

// IsStoragePackageReference returns true if the reference is to a storage package
func IsStoragePackageReference(reference PackageReference) bool {
	_, ok := reference.(StoragePackageReference)
	return ok
}

// GroupVersion returns the group and version of this local reference.
func (s StoragePackageReference) GroupVersion() (string, string, bool) {
	g, v, _ := s.inner.GroupVersion()
	return g, v + StoragePackageSuffix, true
}

// Local returns the local package reference wrapped by this reference
func (s StoragePackageReference) Local() *LocalPackageReference {
	return &s.inner
}
