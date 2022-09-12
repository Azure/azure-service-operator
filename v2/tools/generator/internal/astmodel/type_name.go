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

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// TypeName is a name associated with another Type (it also is usable as a Type)
type TypeName struct {
	PackageReference PackageReference // Note: This has to be a value and not a ptr because this type is used as the key in a map
	name             string
}

var EmptyTypeName TypeName = TypeName{}

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

// WithName returns a new TypeName in the same package but with a different name
func (typeName TypeName) WithName(name string) TypeName {
	return MakeTypeName(typeName.PackageReference, name)
}

// WithPackageReference returns a new TypeName in a different package but with the same name
func (typeName TypeName) WithPackageReference(ref PackageReference) TypeName {
	return MakeTypeName(ref, typeName.name)
}

// A TypeName can be used as a Type,
// it is simply a reference to the name.
var _ Type = TypeName{}

func (typeName TypeName) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	return AsSimpleDeclarations(codeGenerationContext, declContext, typeName)
}

// AsType implements Type for TypeName
func (typeName TypeName) AsType(codeGenerationContext *CodeGenerationContext) dst.Expr {
	// If our name is in the current package, we don't need to qualify it
	if codeGenerationContext.currentPackage.Equals(typeName.PackageReference) {
		return dst.NewIdent(typeName.name)
	}

	// Need to ensure we include a selector for that reference
	packageName, err := codeGenerationContext.GetImportedPackageName(typeName.PackageReference)
	if err != nil {
		panic(fmt.Sprintf(
			"no reference for %s from %s available in package %s",
			typeName.Name(),
			typeName.PackageReference,
			codeGenerationContext.currentPackage))
	}

	return astbuilder.Selector(dst.NewIdent(packageName), typeName.Name())
}

// AsZero renders an expression for the "zero" value of the type.
// The exact thing we need to generate depends on the actual type we reference
func (typeName TypeName) AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	if IsExternalPackageReference(typeName.PackageReference) {
		// TypeName is external, zero value is a qualified empty struct
		// (we might not actually use this, if the property is optional, but we still need to generate the right thing)

		packageName := ctx.MustGetImportedPackageName(typeName.PackageReference)
		return &dst.SelectorExpr{
			X:   dst.NewIdent(packageName),
			Sel: dst.NewIdent(fmt.Sprintf("%s{}", typeName.Name())),
		}
	}

	actualType, err := definitions.FullyResolve(typeName)
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
	return actualType.AsZero(definitions, ctx)
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
func (typeName TypeName) Equals(t Type, override EqualityOverrides) bool {
	if typeName == t && override.TypeName == nil {
		return true
	}

	other, ok := t.(TypeName)
	if !ok {
		return false
	}

	if override.TypeName != nil {
		return override.TypeName(typeName, other)
	}

	return typeName.name == other.name && typeName.PackageReference.Equals(other.PackageReference)
}

// String returns the string representation of the type name, and implements fmt.Stringer.
func (typeName TypeName) String() string {
	return fmt.Sprintf("%s/%s", typeName.PackageReference, typeName.name)
}

var typeNamePluralToSingularOverrides = map[string]string{
	"Redis": "Redis",
	"redis": "redis",
}

var typeNameSingularToPluralOverrides map[string]string

// Singular returns a TypeName with the name singularized.
func (typeName TypeName) Singular(idFactory IdentifierFactory) TypeName {
	// work around bug in flect: https://github.com/Azure/azure-service-operator/issues/1454
	name := typeName.name
	for plural, single := range typeNamePluralToSingularOverrides {
		if strings.HasSuffix(name, plural) {
			n := name[0:len(name)-len(plural)] + single
			return MakeTypeName(typeName.PackageReference, n)
		}
	}

	// Flect isn't consistent about what case it returns. If it's just removing an 's', it will maintain
	// case, but if it's performing a more complicated transformation the result will be all lower case.
	singular := idFactory.CreateIdentifier(flect.Singularize(typeName.name), Exported)
	return MakeTypeName(typeName.PackageReference, singular)
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
// definitions is a dictionary for resolving named types
func (typeName TypeName) WriteDebugDescription(builder *strings.Builder, currentPackage PackageReference) {
	if typeName.PackageReference != nil && !typeName.PackageReference.Equals(currentPackage) {
		// Reference to a different package, so qualify the output.
		// External packages are just qualified by name, other packages by full path
		if IsExternalPackageReference(typeName.PackageReference) {
			builder.WriteString(typeName.PackageReference.PackageName())
		} else {
			builder.WriteString(typeName.PackageReference.String())
		}

		builder.WriteString(".")
	}

	builder.WriteString(typeName.name)
}

// IsEmpty is a predicate that returns true if the TypeName is empty, false otherwise
func (typeName TypeName) IsEmpty() bool {
	return typeName == EmptyTypeName
}

const (
	// SpecSuffix is the suffix used for all Spec types
	SpecSuffix = "_Spec"
	// StatusSuffix is the suffix used for all Status types
	StatusSuffix = "_STATUS"
	// ArmSuffix is the suffix used for all ARM types
	ArmSuffix = "_ARM"
)

// IsSpec returns true if the type name specifies a spec
// Sometimes we build type names by adding a suffix after _Spec, so we need to use a contains check
func (typeName TypeName) IsSpec() bool {
	return strings.Contains(typeName.Name(), SpecSuffix)
}

// IsStatus returns true if the type name specifies a status
// Sometimes we build type names by adding a suffix after _STATUS, so we need to use a contains check
func (typeName TypeName) IsStatus() bool {
	return strings.Contains(typeName.Name(), StatusSuffix)
}

// CreateARMTypeName creates an ARM object type name
func CreateARMTypeName(name TypeName) TypeName {
	return MakeTypeName(name.PackageReference, name.Name()+ArmSuffix)
}
