/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

const (
	CompatPackageName = "compat"
)

// MakeCompatPackageReference creates a subpackage reference using the standard name 'compat'.
// We use subpackages of this form to contain compatibility types we create to handle skipping properties.
func MakeCompatPackageReference(pkg InternalPackageReference) SubPackageReference {
	return MakeSubPackageReference(CompatPackageName, pkg)
}

// IsCompatPackageReference returns true if the reference is to a compat package
func IsCompatPackageReference(reference PackageReference) bool {
	if sub, ok := reference.(SubPackageReference); ok {
		if sub.name == CompatPackageName {
			return true
		}

		return IsCompatPackageReference(sub.parent)
	}

	return false
}
