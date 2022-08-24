/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

const (
	// ArmSuffix is the suffix used for all ARM types
	ArmSuffix = "ARM"
)

// CreateARMTypeName creates an ARM object type name
func CreateARMTypeName(name TypeName) TypeName {
	return MakeTypeName(name.PackageReference, name.Name()+ArmSuffix)
}
