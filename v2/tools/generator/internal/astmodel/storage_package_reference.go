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
var _ LocalLikePackageReference = StoragePackageReference{}
var _ DerivedPackageReference = StoragePackageReference{}

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

// ImportPath returns the path to use when importing this package
func (s StoragePackageReference) ImportPath() string {
	return s.inner.ImportPath() + StoragePackageSuffix
}

// FolderPath returns the path to this package on disk
func (s StoragePackageReference) FolderPath() string {
	return s.inner.FolderPath() + StoragePackageSuffix
}

func (s StoragePackageReference) Version() string {
	return s.inner.Version() + StoragePackageSuffix
}

func (s StoragePackageReference) Group() string {
	return s.inner.Group()
}

func (s StoragePackageReference) LocalPathPrefix() string {
	return s.inner.LocalPathPrefix()
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

// TryGroupVersion returns the group and version of this storage reference.
func (s StoragePackageReference) TryGroupVersion() (string, string, bool) {
	g, v, _ := s.inner.TryGroupVersion()
	return g, v + StoragePackageSuffix, true
}

// MustGroupVersion returns the group and version of this storage reference.
func (s StoragePackageReference) GroupVersion() (string, string) {
	g, v := s.inner.GroupVersion()
	return g, v + StoragePackageSuffix
}

// Local returns the local package reference wrapped by this reference
func (s StoragePackageReference) Local() LocalPackageReference {
	return s.inner
}

// ImportAlias returns the import alias to use for this package reference
func (s StoragePackageReference) ImportAlias(style PackageImportStyle) string {
	base := s.inner.ImportAlias(style)
	switch style {
	case VersionOnly:
		return base + "s"
	case GroupOnly:
		return base
	case GroupAndVersion:
		return base + "s"
	default:
		panic(fmt.Sprintf("didn't expect PackageImportStyle %q", style))
	}
}

// Base implements DerivedPackageReference.
func (s StoragePackageReference) Base() PackageReference {
	return s.inner
}
