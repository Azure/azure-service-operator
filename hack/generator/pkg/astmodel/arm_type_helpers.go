/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// CreateArmTypeName creates an ARM object type name
func CreateArmTypeName(name TypeName) TypeName {
	return MakeTypeName(name.PackageReference, name.Name()+"Arm")
}
