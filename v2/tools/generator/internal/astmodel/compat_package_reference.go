/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

func MakeCompatPackageReference(pkg InternalPackageReference) SubPackageReference {
	return MakeSubPackageReference("compat", pkg)
}
