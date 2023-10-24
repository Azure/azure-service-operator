/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// ExternalTypeName is a name associated with a type that we don't own
type ExternalTypeName struct {
	packageReference ExternalPackageReference // Note: This has to be a value and not a ptr because this type is used as the key in a map
	name             string
}

// A TypeName can be used as a Type, it is simply a reference to the name.
var _ Type = ExternalTypeName{}

var _ TypeName = ExternalTypeName{}

// MakeExternalTypeName is a factory method for creating an ExternalTypeName
func MakeExternalTypeName(ref ExternalPackageReference, name string) ExternalTypeName {
	return ExternalTypeName{
		packageReference: ref,
		name:             name,
	}
}

// Name returns the package-local name of the type
func (tn ExternalTypeName) Name() string {
	return tn.name
}

// PackageReference returns the package to which the type belongs
func (tn ExternalTypeName) PackageReference() PackageReference {
	return tn.packageReference
}

func (tn ExternalTypeName) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	return AsSimpleDeclarations(codeGenerationContext, declContext, tn)
}

// AsType implements Type for TypeName
func (tn ExternalTypeName) AsType(codeGenerationContext *CodeGenerationContext) dst.Expr {
	// Need to ensure we include a selector for our external reference
	packageName := codeGenerationContext.MustGetImportedPackageName(tn.packageReference)
	return astbuilder.Selector(dst.NewIdent(packageName), tn.Name())
}

// AsZero renders an expression for the "zero" value of the type.
// The exact thing we need to generate depends on the actual type we reference
func (tn ExternalTypeName) AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	// TypeName is external, zero value is a qualified empty struct
	// (we might not actually use this, if the property is optional, but we still need to generate the right thing)
	packageName := ctx.MustGetImportedPackageName(tn.packageReference)
	return &dst.SelectorExpr{
		X:   dst.NewIdent(packageName),
		Sel: dst.NewIdent(fmt.Sprintf("%s{}", tn.Name())),
	}
}

// References returns a set containing this type name.
func (tn ExternalTypeName) References() TypeNameSet {
	return NewTypeNameSet(tn)
}

// RequiredPackageReferences returns all the imports required for this definition
func (tn ExternalTypeName) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet(tn.packageReference)
}

// Equals returns true if the passed type is the same TypeName, false otherwise
func (tn ExternalTypeName) Equals(t Type, _ EqualityOverrides) bool {
	if tn == t {
		return true
	}

	other, ok := t.(ExternalTypeName)
	if !ok {
		return false
	}

	return tn.name == other.Name() &&
		tn.packageReference.Equals(other.PackageReference())
}

// String returns the string representation of the type name, and implements fmt.Stringer.
func (tn ExternalTypeName) String() string {
	return fmt.Sprintf("%s/%s", tn.packageReference, tn.name)
}

// WriteDebugDescription adds a description of the current type to the passed builder.
// builder receives the full description, including nested types.
// definitions is a dictionary for resolving named types.
func (tn ExternalTypeName) WriteDebugDescription(builder *strings.Builder, currentPackage InternalPackageReference) {
	// External reference will point to a different package, so qualify the output.
	// External packages are just qualified by name.
	builder.WriteString(tn.packageReference.PackageName())
	builder.WriteString(".")
	builder.WriteString(tn.name)
}
