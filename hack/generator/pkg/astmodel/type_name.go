/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
)

// TypeName is a name associated with another Type (it also is usable as a Type)
type TypeName struct {
	PackageReference PackageReference // Note: This has to be a value and not a ptr because this type is used as the key in a map
	name             string
}

// NewTypeName is a factory method for creating a TypeName
func NewTypeName(pr PackageReference, name string) *TypeName {
	return &TypeName{pr, name}
}

// Name returns the package-local name of the type
func (typeName *TypeName) Name() string {
	return typeName.name
}

// A TypeName can be used as a Type,
// it is simply a reference to the name.
var _ Type = (*TypeName)(nil)

// AsType implements Type for TypeName
func (typeName *TypeName) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	// If our package is being referenced, we need to ensure we include a selector for that reference
	packageName, err := codeGenerationContext.GetImportedPackageName(&typeName.PackageReference)
	if err == nil {
		return &ast.SelectorExpr{
			X:   ast.NewIdent(packageName),
			Sel: ast.NewIdent(typeName.Name()),
		}
	}

	// Sanity assertion that the type we're generating is in the same package that the context is for
	if !codeGenerationContext.currentPackage.Equals(&typeName.PackageReference) {
		panic(fmt.Sprintf(
			"no reference for %v included in package %v",
			typeName.name,
			codeGenerationContext.currentPackage))
	}

	return ast.NewIdent(typeName.name)
}

// References indicates whether this Type includes any direct references to the given Type
func (typeName *TypeName) References(d *TypeName) bool {
	return typeName.Equals(d)
}

// RequiredImports returns all the imports required for this definition
func (typeName *TypeName) RequiredImports() []*PackageReference {
	return []*PackageReference{&typeName.PackageReference}
}

// Equals returns true if the passed type is the same TypeName, false otherwise
func (typeName *TypeName) Equals(t Type) bool {
	if d, ok := t.(*TypeName); ok {
		return typeName.name == d.name && typeName.PackageReference.Equals(&d.PackageReference)
	}

	return false
}

// CreateInternalDefinitions does nothing
func (typeName *TypeName) CreateInternalDefinitions(_ *TypeName, _ IdentifierFactory) (Type, []TypeDefiner) {
	// there is nothing internal to a TypeName, return it unchanged
	return typeName, nil
}

// CreateDefinitions adds another name to this already-named type
func (typeName *TypeName) CreateDefinitions(name *TypeName, _ IdentifierFactory) (TypeDefiner, []TypeDefiner) {
	return NewSimpleTypeDefiner(name, typeName), nil
}

// String returns the string representation of the type name
func (typeName *TypeName) String() string {
	return fmt.Sprintf("%s/%s", typeName.PackageReference, typeName.name)
}

// Ensure we implement Stringer
var _ fmt.Stringer = &TypeName{}
