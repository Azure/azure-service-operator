/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// CreateARMTypeName creates an ARM object type name
func CreateARMTypeName(name TypeName) TypeName {
	return MakeTypeName(name.PackageReference, name.Name()+"ARM")
}
