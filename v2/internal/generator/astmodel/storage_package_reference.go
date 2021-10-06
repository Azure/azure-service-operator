/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"
)

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
			localPathPrefix: local.localPathPrefix,
			group:           local.group,
			version:         local.version + StoragePackageSuffix,
		},
	}
}

// String returns the string representation of the package reference
func (s StoragePackageReference) String() string {
	return fmt.Sprintf("storage:%s/%s", s.group, s.version)
}

// IsPreview returns true if this package reference is a preview
func (s StoragePackageReference) IsPreview() bool {
	return containsPreviewVersionLabel(strings.ToLower(s.version))
}

// IsStoragePackageReference returns true if the reference is to a storage package
func IsStoragePackageReference(reference PackageReference) bool {
	_, ok := reference.(StoragePackageReference)
	return ok
}
