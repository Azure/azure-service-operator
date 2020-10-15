/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"

	"github.com/pkg/errors"
)

// Type represents something that is a Go type
type Type interface {
	// RequiredPackageReferences returns a set of packages imports required by this type
	RequiredPackageReferences() []PackageReference

	// References returns the names of all types that this type
	// references. For example, an Array of Persons references a
	// Person.
	References() TypeNameSet

	// AsType renders as a Go abstract syntax tree for a type
	// (yes this says ast.Expr but that is what the Go 'ast' package uses for types)
	AsType(codeGenerationContext *CodeGenerationContext) ast.Expr

	// AsDeclarations renders as a Go abstract syntax tree for a declaration
	AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description []string) []ast.Decl

	// Equals returns true if the passed type is the same as this one, false otherwise
	Equals(t Type) bool

	// Make sure all Types have a printable version for debugging/user info.
	// This doesn't need to be a full representation of the type.
	fmt.Stringer
}

// TypeEquals decides if the types are the same and handles the `nil` case
func TypeEquals(left, right Type) bool {
	if left == nil {
		return right == nil
	}

	return left.Equals(right)
}

// TypeAsObjectType extracts the underlying object type (if any)
func TypeAsObjectType(t Type) (*ObjectType, error) {
	rt, err := cachedAsObjectTypeVisitor.Visit(t, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to determine if %v contains an object", t)
	}

	ot, ok := rt.(*ObjectType)
	if !ok {
		return nil, fmt.Errorf("type %v is not an object", rt)
	}

	return ot, nil
}

var cachedAsObjectTypeVisitor = createAsObjectTypeVisitor()

func createAsObjectTypeVisitor() *TypeVisitor {
	result := MakeTypeVisitor()

	/*
	 * Found the object we want
	 */

	result.VisitObjectType = func(_ *TypeVisitor, ot *ObjectType, _ interface{}) (Type, error) {
		return ot, nil
	}

	/*
	 * Unwrap if needed
	 */

	result.VisitArmType = func(_ *TypeVisitor, at *ArmType, _ interface{}) (Type, error) {
		return &at.objectType, nil
	}
	result.VisitOptionalType = func(tv *TypeVisitor, ot *OptionalType, ctx interface{}) (Type, error) {
		return tv.Visit(ot.element, ctx)
	}

	/*
	 * Not the type we are looking for
	 */

	result.VisitMapType = func(_ *TypeVisitor, _ *MapType, _ interface{}) (Type, error) {
		return nil, fmt.Errorf("a map is not an object type")
	}
	result.VisitArrayType = func(_ *TypeVisitor, _ *ArrayType, _ interface{}) (Type, error) {
		return nil, fmt.Errorf("an array is not an object type")
	}
	result.VisitEnumType = func(_ *TypeVisitor, _ *EnumType, _ interface{}) (Type, error) {
		return nil, fmt.Errorf("an enum is not an object type")
	}
	result.VisitPrimitive = func(_ *TypeVisitor, pt *PrimitiveType, _ interface{}) (Type, error) {
		return nil, fmt.Errorf("a %v is not an object type", pt.name)
	}
	result.VisitTypeName = func(_ *TypeVisitor, _ TypeName, _ interface{}) (Type, error) {
		return nil, fmt.Errorf("a name is not an object type")
	}

	result.VisitResourceType = func(_ *TypeVisitor, _ *ResourceType, _ interface{}) (Type, error) {
		return nil, fmt.Errorf("a resource type contains multiple object types (ambiguous result)")
	}

	return &result
}
