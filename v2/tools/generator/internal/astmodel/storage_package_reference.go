/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

const (
	StoragePackageSuffix = "storage"
)

type StoragePackageReference struct {
	inner LocalPackageReference // a reference to the API package this storage package mirrors
}

var _ PackageReference = StoragePackageReference{}
var _ InternalPackageReference = StoragePackageReference{}
var _ DerivedPackageReference = StoragePackageReference{}

// legacyStorageGroups is a set of groups for which we generate old style storage packages (siblings of the API
// packages). We only do this to reduce the number of changes in a single PR. Once we've migrated all the packages
// we can remove this.
var legacyStorageGroups = set.Make(
	"authorization",
	"batch",
	"cache",
	"containerinstance",
	"containerregistry",
	"containerservice",
	"datafactory",
	"dataprotection",
	"dbformysql",
	"dbforpostgresql",
	"devices",
	"eventhub",
	"insights",
	"keyvault",
	"machinelearningservices",
	"managedidentity",
	"operationalinsights",
	"search",
	"servicebus",
	"signalrservice",
	"sql",
	"storage",
	"subscription",
	"synapse",
	"web",
)

// MakeStoragePackageReference creates a new storage package reference from a local package reference
func MakeStoragePackageReference(ref InternalPackageReference) InternalPackageReference {
	switch r := ref.(type) {
	case LocalPackageReference:
		if legacyStorageGroups.Contains(r.group) {
			return StoragePackageReference{
				inner: r,
			}
		}

		return MakeSubPackageReference(StoragePackageSuffix, r)
	case StoragePackageReference:
		return r
	case SubPackageReference:
		parent := MakeStoragePackageReference(r.parent)
		return MakeSubPackageReference(r.name, parent)
	default:
		panic(fmt.Sprintf("unknown package reference type %T", ref))
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

// FolderPath returns the relative path to this package on disk.
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

// IsStoragePackageReference returns true if the reference is to a storage package OR to a subpackage for storage
func IsStoragePackageReference(reference PackageReference) bool {
	if _, ok := reference.(StoragePackageReference); ok {
		return true
	}

	if sub, ok := reference.(SubPackageReference); ok {
		return sub.name == StoragePackageSuffix
	}

	return false
}

// GroupVersion returns the group and version of this storage reference.
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
func (s StoragePackageReference) Base() InternalPackageReference {
	return s.inner
}
