/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

type TypeName interface {
	Type
	Name() string
	PackageReference() PackageReference
}

const (
	// SpecSuffix is the suffix used for all Spec types
	SpecSuffix = "_Spec"
	// StatusSuffix is the suffix used for all Status types
	StatusSuffix = "_STATUS"
	// ARMSuffix is the suffix used for all ARM types
	ARMSuffix = "_ARM"
)

// CreateARMTypeName creates an ARM object type name
func CreateARMTypeName(name InternalTypeName) InternalTypeName {
	return MakeInternalTypeName(name.InternalPackageReference(), name.Name()+ARMSuffix)
}
