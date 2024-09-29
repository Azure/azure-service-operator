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
	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/names"
)

// InternalTypeName is a name associated with another Type (it also is usable as a Type)
type InternalTypeName struct {
	packageReference InternalPackageReference // Note: This has to be a value and not a ptr because this type is used as the key in a map
	name             string
}

// MakeInternalTypeName is a factory method for creating a TypeName
func MakeInternalTypeName(ref InternalPackageReference, name string) InternalTypeName {
	return InternalTypeName{
		packageReference: ref,
		name:             name,
	}
}

// Name returns the package-local name of the type
func (tn InternalTypeName) Name() string {
	return tn.name
}

// PackageReference returns the package to which the type belongs
func (tn InternalTypeName) PackageReference() PackageReference {
	return tn.packageReference
}

// InternalPackageReference returns the internal package to which the type belongs
func (tn InternalTypeName) InternalPackageReference() InternalPackageReference {
	return tn.packageReference
}

// WithName returns a new InternalTypeName in the same package but with a different name
func (tn InternalTypeName) WithName(name string) InternalTypeName {
	return MakeInternalTypeName(tn.packageReference, name)
}

// WithPackageReference returns a new InternalTypeName in a different package but with the same name
func (tn InternalTypeName) WithPackageReference(ref InternalPackageReference) InternalTypeName {
	return MakeInternalTypeName(ref, tn.name)
}

// A TypeName can be used as a Type,
// it is simply a reference to the name.
var _ Type = InternalTypeName{}

func (tn InternalTypeName) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) ([]dst.Decl, error) {
	return AsSimpleDeclarations(codeGenerationContext, declContext, tn)
}

// AsType implements Type for TypeName
func (tn InternalTypeName) AsTypeExpr(codeGenerationContext *CodeGenerationContext) (dst.Expr, error) {
	// If our name is in the current package, we don't need to qualify it
	if codeGenerationContext.currentPackage.Equals(tn.packageReference) {
		return dst.NewIdent(tn.name), nil
	}

	// Need to ensure we include a selector for that reference
	packageName, err := codeGenerationContext.GetImportedPackageName(tn.packageReference)
	if err != nil {
		panic(fmt.Sprintf(
			"no reference for %s from %s available in package %s",
			tn.Name(),
			tn.packageReference,
			codeGenerationContext.currentPackage))
	}

	return astbuilder.Selector(dst.NewIdent(packageName), tn.Name()), nil
}

// AsZero renders an expression for the "zero" value of the type.
// The exact thing we need to generate depends on the actual type we reference
func (tn InternalTypeName) AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	actualType, err := definitions.FullyResolve(tn)
	if err != nil {
		// This should never happen
		panic(err)
	}

	if _, isObject := AsObjectType(actualType); isObject {
		// We reference an object type, so our zero value is an empty struct
		// But, we need to qualify it if it is from another package
		if tn.packageReference.Equals(ctx.CurrentPackage()) {
			// Current package, no qualification needed
			return &dst.BasicLit{
				Value: fmt.Sprintf("%s{}", tn.Name()),
			}
		}

		packageName := ctx.MustGetImportedPackageName(tn.packageReference)

		return &dst.SelectorExpr{
			X:   dst.NewIdent(packageName),
			Sel: dst.NewIdent(fmt.Sprintf("%s{}", tn.Name())),
		}
	}

	// Otherwise we need the underlying type (e.g. enums, primitive type, etc.)
	return actualType.AsZero(definitions, ctx)
}

// References returns a set containing this type name.
func (tn InternalTypeName) References() TypeNameSet {
	return NewTypeNameSet(tn)
}

// RequiredPackageReferences returns all the imports required for this definition
func (tn InternalTypeName) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet(tn.packageReference)
}

// Equals returns true if the passed type is the same TypeName, false otherwise
func (tn InternalTypeName) Equals(t Type, override EqualityOverrides) bool {
	if tn == t && override.InternalTypeName == nil {
		return true
	}

	other, ok := t.(InternalTypeName)
	if !ok {
		return false
	}

	if override.InternalTypeName != nil {
		return override.InternalTypeName(tn, other)
	}

	return tn.name == other.Name() &&
		tn.packageReference.Equals(other.PackageReference())
}

// String returns the string representation of the type name, and implements fmt.Stringer.
func (tn InternalTypeName) String() string {
	return fmt.Sprintf("%s/%s", tn.packageReference, tn.name)
}

// Singular returns a TypeName with the name singularized.
func (tn InternalTypeName) Singular() InternalTypeName {
	name := names.Singularize(tn.Name())
	return tn.WithName(name)
}

// Plural returns a TypeName with the name pluralized.
func (tn InternalTypeName) Plural() InternalTypeName {
	name := names.Pluralize(tn.Name())
	return tn.WithName(name)
}

// WriteDebugDescription adds a description of the current type to the passed builder.
// builder receives the full description, including nested types
// definitions is a dictionary for resolving named types
func (tn InternalTypeName) WriteDebugDescription(builder *strings.Builder, currentPackage InternalPackageReference) {
	if tn.packageReference != nil && !tn.packageReference.Equals(currentPackage) {
		// Reference to a different package, so qualify the output.
		builder.WriteString(tn.packageReference.FolderPath())
		builder.WriteString(".")
	}

	builder.WriteString(tn.name)
}

// IsSpec returns true if the type name specifies a spec
// Sometimes we build type names by adding a suffix after _Spec, so we need to use a contains check
func (tn InternalTypeName) IsSpec() bool {
	return strings.Contains(tn.Name(), SpecSuffix)
}

// IsStatus returns true if the type name specifies a status
// Sometimes we build type names by adding a suffix after _STATUS, so we need to use a contains check
func (tn InternalTypeName) IsStatus() bool {
	return strings.Contains(tn.Name(), StatusSuffix)
}

// IsARMType returns true if the TypeName identifies an ARM specific type, false otherwise.
func (tn InternalTypeName) IsARMType() bool {
	if IsARMPackageReference(tn.InternalPackageReference()) {
		return true
	}

	return strings.HasSuffix(tn.Name(), ARMSuffix)
}

func (tn InternalTypeName) IsEmpty() bool {
	return tn.name == "" && tn.packageReference == nil
}
