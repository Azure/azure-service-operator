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
	// ARMPackageName is the name used for ARM subpackages
	ARMPackageName = "arm"
)

// CreateARMTypeName creates an ARM object type name
func CreateARMTypeName(name InternalTypeName) InternalTypeName {
	armPackage := MakeSubPackageReference(ARMPackageName, name.InternalPackageReference())
	return MakeInternalTypeName(armPackage, name.Name())
}
