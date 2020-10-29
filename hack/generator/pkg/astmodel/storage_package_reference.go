/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

const (
	StoragePackageSuffix = "storage"
)

type StoragePackageReference struct {
	LocalPackageReference
}

var _ PackageReference = StoragePackageReference{}

// MakeStoragePackageReference creates a new storage package reference from a local package reference
func MakeStoragePackageReference(local LocalPackageReference) StoragePackageReference {
	return StoragePackageReference{
		LocalPackageReference{
			group:   local.group,
			version: local.version + StoragePackageSuffix,
		},
	}
}

// IsStoragePackageReference returns true if the reference is to a storage package
func IsStoragePackageReference(reference PackageReference) bool {
	_, ok := reference.(StoragePackageReference)
	return ok
}
