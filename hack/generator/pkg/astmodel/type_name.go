/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
	"github.com/gobuffalo/flect"
)

// TypeName is a name associated with another Type (it also is usable as a Type)
type TypeName struct {
	PackageReference PackageReference // Note: This has to be a value and not a ptr because this type is used as the key in a map
	name             string
}

func SortTypeName(left, right TypeName) bool {
	leftRef := left.PackageReference
	rightRef := right.PackageReference
	return leftRef.PackagePath() < rightRef.PackagePath() ||
		(leftRef.PackagePath() == rightRef.PackagePath() && left.name < right.name)
}

// MakeTypeName is a factory method for creating a TypeName
func MakeTypeName(pr PackageReference, name string) TypeName {
	return TypeName{pr, name}
}

// Name returns the package-local name of the type
func (typeName TypeName) Name() string {
	return typeName.name
}

// A TypeName can be used as a Type,
// it is simply a reference to the name.
var _ Type = TypeName{}

func (typeName TypeName) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	return AsSimpleDeclarations(codeGenerationContext, declContext, typeName)
}

// AsType implements Type for TypeName
func (typeName TypeName) AsType(codeGenerationContext *CodeGenerationContext) dst.Expr {
	// If our package is being referenced, we need to ensure we include a selector for that reference
	packageName, err := codeGenerationContext.GetImportedPackageName(typeName.PackageReference)
	if err == nil {
		return &dst.SelectorExpr{
			X:   dst.NewIdent(packageName),
			Sel: dst.NewIdent(typeName.Name()),
		}
	}

	// Safety assertion that the type we're generating is in the same package that the context is for
	if !codeGenerationContext.currentPackage.Equals(typeName.PackageReference) {
		panic(fmt.Sprintf(
			"no reference for %v included in package %v",
			typeName.name,
			codeGenerationContext.currentPackage))
	}

	return dst.NewIdent(typeName.name)
}

// AsZero renders an expression for the "zero" value of the type.
// The exact thing we need to generate depends on the actual type we reference
func (typeName TypeName) AsZero(types Types, ctx *CodeGenerationContext) dst.Expr {
	actualType, err := types.FullyResolve(typeName)
	if err != nil {
		// This should never happen
		panic(err)
	}

	if _, isObject := AsObjectType(actualType); isObject {
		// We reference an object type, so our zero value is an empty struct
		// But, we need to qualify it if it is from another package
		if typeName.PackageReference.Equals(ctx.CurrentPackage()) {
			// Current package, no qualification needed
			return &dst.BasicLit{
				Value: fmt.Sprintf("%s{}", typeName.Name()),
			}
		}

		packageName := ctx.MustGetImportedPackageName(typeName.PackageReference)

		return &dst.SelectorExpr{
			X:   dst.NewIdent(packageName),
			Sel: dst.NewIdent(fmt.Sprintf("%s{}", typeName.Name())),
		}
	}

	// Otherwise we need the underlying type (e.g. enums, primitive type, etc)
	return actualType.AsZero(types, ctx)
}

// References returns a set containing this type name.
func (typeName TypeName) References() TypeNameSet {
	return NewTypeNameSet(typeName)
}

// RequiredPackageReferences returns all the imports required for this definition
func (typeName TypeName) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet(typeName.PackageReference)
}

// Equals returns true if the passed type is the same TypeName, false otherwise
func (typeName TypeName) Equals(t Type) bool {
	if d, ok := t.(TypeName); ok {
		return typeName.name == d.name && typeName.PackageReference.Equals(d.PackageReference)
	}

	return false
}

// String returns the string representation of the type name, and implements fmt.Stringer.
func (typeName TypeName) String() string {
	return fmt.Sprintf("%s/%s", typeName.PackageReference, typeName.name)
}

var typeNamePluralToSingularOverrides = map[string]string{
	"Services":  "Service",
	"services":  "service",
	"Redis":     "Redis",
	"redis":     "redis",
	"Addresses": "Address",
	"addresses": "address",
}

var typeNameSingularToPluralOverrides map[string]string

// Singular returns a TypeName with the name singularized.
func (typeName TypeName) Singular() TypeName {
	// work around bug in flect: https://github.com/Azure/k8s-infra/issues/319
	name := typeName.name
	for plural, single := range typeNamePluralToSingularOverrides {
		if strings.HasSuffix(name, plural) {
			n := name[0:len(name)-len(plural)] + single
			return MakeTypeName(typeName.PackageReference, n)
		}
	}

	return MakeTypeName(typeName.PackageReference, flect.Singularize(typeName.name))
}

// Plural returns a TypeName with the name pluralized.
func (typeName TypeName) Plural() TypeName {
	name := typeName.name

	if typeNamePluralToSingularOverrides == nil {
		typeNameSingularToPluralOverrides = make(map[string]string, len(typeNamePluralToSingularOverrides))
		for plural, single := range typeNamePluralToSingularOverrides {
			typeNameSingularToPluralOverrides[single] = plural
		}
	}

	for single, plural := range typeNameSingularToPluralOverrides {
		if strings.HasSuffix(name, single) {
			n := name[0:len(name)-len(single)] + plural
			return MakeTypeName(typeName.PackageReference, n)
		}
	}

	return MakeTypeName(typeName.PackageReference, flect.Pluralize(typeName.name))
}

// WriteDebugDescription adds a description of the current type to the passed builder
// builder receives the full description, including nested types
// types is a dictionary for resolving named types
func (typeName TypeName) WriteDebugDescription(builder *strings.Builder, types Types) {
	builder.WriteString(typeName.PackageReference.String())
	builder.WriteString("/")
	builder.WriteString(typeName.name)

	if _, ok := typeName.PackageReference.AsLocalPackage(); ok {
		builder.WriteString(":")
		if definition, ok := types[typeName]; ok {
			definition.Type().WriteDebugDescription(builder, types)
		} else {
			builder.WriteString("NOTDEFINED")
		}
	}
}
