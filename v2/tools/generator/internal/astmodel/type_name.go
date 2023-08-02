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
	WithName(name string) TypeName
	WithPackageReference(ref PackageReference) TypeName
	AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl
	AsType(codeGenerationContext *CodeGenerationContext) dst.Expr
	AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr
	References() TypeNameSet
	RequiredPackageReferences() *PackageReferenceSet
	Equals(t Type, override EqualityOverrides) bool
	String() string
	Singular() TypeName
	Plural() TypeName
	WriteDebugDescription(builder *strings.Builder, currentPackage PackageReference)
	IsSpec() bool
	IsStatus() bool
	IsARMType() bool
}

func SortTypeName(left, right TypeName) bool {
	leftRef := left.PackageReference()
	rightRef := right.PackageReference()
	return leftRef.ImportPath() < rightRef.ImportPath() ||
		(leftRef.ImportPath() == rightRef.ImportPath() && left.Name() < right.Name())
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
func CreateARMTypeName(name TypeName) InternalTypeName {
	return MakeInternalTypeName(name.PackageReference(), name.Name()+ARMSuffix)
}
