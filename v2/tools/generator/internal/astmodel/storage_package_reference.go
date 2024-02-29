/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
)

const (
	StoragePackageName = "storage"
)

// MakeStoragePackageReference creates a new storage package reference from a local package reference
func MakeStoragePackageReference(ref InternalPackageReference) InternalPackageReference {
	switch r := ref.(type) {
	case LocalPackageReference:
		return MakeSubPackageReference(StoragePackageName, r)
	case SubPackageReference:
		parent := MakeStoragePackageReference(r.parent)
		return MakeSubPackageReference(r.name, parent)
	default:
		panic(fmt.Sprintf("unknown package reference type %T", ref))
	}
}

// IsStoragePackageReference returns true if the reference is to a storage package OR to a subpackage for storage
func IsStoragePackageReference(reference PackageReference) bool {
	if sub, ok := reference.(SubPackageReference); ok {
		if sub.name == StoragePackageName {
			return true
		}

		return IsStoragePackageReference(sub.parent)
	}

	return false
}
