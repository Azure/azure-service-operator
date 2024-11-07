/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "github.com/Azure/azure-service-operator/v2/internal/set"

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
	// ARMPackageName is the name used for ARM subpackages
	ARMPackageName = "arm"
)

var armPackageDenyList = set.Make(
	"kusto",
	"resources")

// CreateARMTypeName creates an ARM object type name
func CreateARMTypeName(name InternalTypeName) InternalTypeName {
	pkg := name.InternalPackageReference()
	if armPackageDenyList.Contains(pkg.Group()) {
		return MakeInternalTypeName(pkg, name.Name()+ARMSuffix)
	}

	armPackage := MakeSubPackageReference(ARMPackageName, name.InternalPackageReference())
	return MakeInternalTypeName(armPackage, name.Name())
}
