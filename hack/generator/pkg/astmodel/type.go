/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
)

// Type represents something that is a Go type
type Type interface {
	// RequiredImports returns a list of packages required by this type
	RequiredImports() []PackageReference

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

// TypeVisitor represents a visitor for a tree of types.
// The `ctx` argument can be used to “smuggle” additional data down the call-chain.
type TypeVisitor struct {
	VisitTypeName         func(this *TypeVisitor, it TypeName, ctx interface{}) Type
	VisitArrayType        func(this *TypeVisitor, it *ArrayType, ctx interface{}) Type
	VisitPrimitive        func(this *TypeVisitor, it *PrimitiveType, ctx interface{}) Type
	VisitObjectType       func(this *TypeVisitor, it *ObjectType, ctx interface{}) Type
	VisitMapType          func(this *TypeVisitor, it *MapType, ctx interface{}) Type
	VisitOptionalType     func(this *TypeVisitor, it *OptionalType, ctx interface{}) Type
	VisitEnumType         func(this *TypeVisitor, it *EnumType, ctx interface{}) Type
	VisitResourceType     func(this *TypeVisitor, it *ResourceType, ctx interface{}) Type
	VisitResourceListType func(this *TypeVisitor, it *ResourceListType, ctx interface{}) Type
}

// Visit invokes the appropriate VisitX on TypeVisitor
func (tv *TypeVisitor) Visit(t Type, ctx interface{}) Type {
	if t == nil {
		return nil
	}

	switch it := t.(type) {
	case TypeName:
		return tv.VisitTypeName(tv, it, ctx)
	case *ArrayType:
		return tv.VisitArrayType(tv, it, ctx)
	case *PrimitiveType:
		return tv.VisitPrimitive(tv, it, ctx)
	case *ObjectType:
		return tv.VisitObjectType(tv, it, ctx)
	case *MapType:
		return tv.VisitMapType(tv, it, ctx)
	case *OptionalType:
		return tv.VisitOptionalType(tv, it, ctx)
	case *EnumType:
		return tv.VisitEnumType(tv, it, ctx)
	case *ResourceType:
		return tv.VisitResourceType(tv, it, ctx)
	case *ResourceListType:
		return tv.VisitResourceListType(tv, it, ctx)
	}

	panic(fmt.Sprintf("unhandled type: (%T) %v", t, t))
}

// VisitDefinition invokes the TypeVisitor on both the name and type of the definition
// NB: this is only valid if VisitTypeName returns a TypeName and not generally a Type
func (tv *TypeVisitor) VisitDefinition(td TypeDefinition, ctx interface{}) TypeDefinition {
	return MakeTypeDefinition(
		tv.VisitTypeName(tv, td.Name(), ctx).(TypeName),
		tv.Visit(td.Type(), ctx))
}

// MakeTypeVisitor returns a default (identity transform) visitor, which
// visits every type in the tree. If you want to actually do something you will
// need to override the properties on the returned TypeVisitor.
func MakeTypeVisitor() TypeVisitor {
	// TODO [performance]: we can do reference comparisons on the results of
	// recursive invocations of Visit to avoid having to rebuild the tree if the
	// leafs do not actually change.
	return TypeVisitor{
		VisitTypeName: func(_ *TypeVisitor, it TypeName, _ interface{}) Type {
			return it
		},
		VisitArrayType: func(this *TypeVisitor, it *ArrayType, ctx interface{}) Type {
			newElement := this.Visit(it.element, ctx)
			return NewArrayType(newElement)
		},
		VisitPrimitive: func(_ *TypeVisitor, it *PrimitiveType, _ interface{}) Type {
			return it
		},
		VisitObjectType: func(this *TypeVisitor, it *ObjectType, ctx interface{}) Type {
			// just map the property types
			var newProps []*PropertyDefinition
			for _, prop := range it.properties {
				newProps = append(newProps, prop.WithType(this.Visit(prop.propertyType, ctx)))
			}
			return it.WithProperties(newProps...)
		},
		VisitMapType: func(this *TypeVisitor, it *MapType, ctx interface{}) Type {
			newKey := this.Visit(it.key, ctx)
			newValue := this.Visit(it.value, ctx)
			return NewMapType(newKey, newValue)
		},
		VisitEnumType: func(_ *TypeVisitor, it *EnumType, _ interface{}) Type {
			// if we visit the enum base type then we will also have to do something
			// about the values. so by default don't do anything with the enum base
			return it
		},
		VisitOptionalType: func(this *TypeVisitor, it *OptionalType, ctx interface{}) Type {
			return NewOptionalType(this.Visit(it.element, ctx))
		},
		VisitResourceType: func(this *TypeVisitor, it *ResourceType, ctx interface{}) Type {
			spec := this.Visit(it.spec, ctx)
			status := this.Visit(it.status, ctx)
			return NewResourceType(spec, status).WithOwner(it.Owner())
		},
		VisitResourceListType: func(this *TypeVisitor, it *ResourceListType, ctx interface{}) Type {
			resource := this.VisitTypeName(this, it.resource, ctx)
			// There's an expectation that this is a typeName
			resourceTypeName := resource.(TypeName)
			return NewResourceListType(resourceTypeName)
		},
	}
}
