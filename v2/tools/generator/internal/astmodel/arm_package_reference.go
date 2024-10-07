/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// IsARMPackageReference returns true if the reference is to an ARM package OR to a subpackage of
// an ARM package, false otherwise.
func IsARMPackageReference(reference PackageReference) bool {
	if sub, ok := reference.(SubPackageReference); ok {
		if sub.name == ARMPackageName {
			return true
		}

		return IsARMPackageReference(sub.parent)
	}

	return false
}
