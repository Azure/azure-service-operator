/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// MakeCompatPackageReference creates a subpackage reference using the standard name 'compat'.
// We use subpackages of this form to contain compatibility types we create to handle skipping properties.
func MakeCompatPackageReference(pkg InternalPackageReference) SubPackageReference {
	return MakeSubPackageReference("compat", pkg)
}
