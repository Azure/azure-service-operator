/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"strings"

	"github.com/dave/dst"
)

type TypeName interface {
	Name() string
	PackageReference() PackageReference
	AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl
	AsType(codeGenerationContext *CodeGenerationContext) dst.Expr
	AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr
	References() TypeNameSet
	RequiredPackageReferences() *PackageReferenceSet
	Equals(t Type, override EqualityOverrides) bool
	String() string
	WriteDebugDescription(builder *strings.Builder, currentPackage InternalPackageReference)
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
