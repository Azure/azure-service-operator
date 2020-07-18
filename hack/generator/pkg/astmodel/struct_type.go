/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
	"sort"
)

// StructType represents an (unnamed) struct type
type StructType struct {
	properties map[PropertyName]*PropertyDefinition
	functions  map[string]Function
}

// EmptyStructType is an empty struct
var EmptyStructType = NewStructType()

// Ensure StructType implements the Type interface correctly
var _ Type = (*StructType)(nil)

// NewStructType is a factory method for creating a new StructTypeDefinition
func NewStructType() *StructType {
	return &StructType{
		properties: make(map[PropertyName]*PropertyDefinition),
		functions:  make(map[string]Function),
	}
}

func (structType *StructType) AsDeclarations(codeGenerationContext *CodeGenerationContext, name *TypeName, description *string) []ast.Decl {
	identifier := ast.NewIdent(name.Name())
	declaration := &ast.GenDecl{
		Tok: token.TYPE,
		Doc: &ast.CommentGroup{},
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: identifier,
				Type: structType.AsType(codeGenerationContext),
			},
		},
	}

	if description != nil {
		addDocComment(&declaration.Doc.List, *description, 200)
	}

	result := []ast.Decl{declaration}
	result = append(result, structType.generateMethodDecls(codeGenerationContext, name)...)
	return result
}

func (structType *StructType) generateMethodDecls(codeGenerationContext *CodeGenerationContext, typeName *TypeName) []ast.Decl {
	var result []ast.Decl
	for methodName, function := range structType.functions {
		funcDef := function.AsFunc(codeGenerationContext, typeName, methodName)
		result = append(result, funcDef)
	}

	return result
}

func defineField(fieldName string, fieldType ast.Expr, tag string) *ast.Field {

	result := &ast.Field{
		Type: fieldType,
		Tag:  &ast.BasicLit{Kind: token.STRING, Value: tag},
	}

	if fieldName != "" {
		result.Names = []*ast.Ident{ast.NewIdent(fieldName)}
	}

	return result
}

// Properties returns all our property definitions
// A sorted slice is returned to preserve immutability and provide determinism
func (structType *StructType) Properties() []*PropertyDefinition {
	var result []*PropertyDefinition
	for _, property := range structType.properties {
		result = append(result, property)
	}

	sort.Slice(result, func(left int, right int) bool {
		return result[left].propertyName < result[right].propertyName
	})

	return result
}

// AsType implements Type for StructType
func (structType *StructType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {

	// Copy the slice of properties and sort it
	properties := structType.Properties()
	sort.Slice(properties, func(i int, j int) bool {
		return properties[i].propertyName < properties[j].propertyName
	})

	fields := make([]*ast.Field, len(properties))
	for i, f := range properties {
		fields[i] = f.AsField(codeGenerationContext)
	}

	return &ast.StructType{
		Fields: &ast.FieldList{
			List: fields,
		},
	}
}

// RequiredImports returns a list of packages required by this
func (structType *StructType) RequiredImports() []*PackageReference {
	var result []*PackageReference
	for _, property := range structType.properties {
		result = append(result, property.PropertyType().RequiredImports()...)
	}

	for _, function := range structType.functions {
		result = append(result, function.RequiredImports()...)
	}

	return result
}

// References returns the set of all the types referred to by any property.
func (structType *StructType) References() TypeNameSet {
	var results TypeNameSet
	for _, property := range structType.properties {
		for ref := range property.PropertyType().References() {
			results = results.Add(ref)
		}
	}
	// Not collecting types from functions deliberately.
	return results
}

// Equals returns true if the passed type is a struct type with the same properties, false otherwise
// The order of the properties is not relevant
func (structType *StructType) Equals(t Type) bool {
	if structType == t {
		return true
	}

	if st, ok := t.(*StructType); ok {
		if len(structType.properties) != len(st.properties) {
			// Different number of properties, not equal
			return false
		}

		for n, f := range st.properties {
			ourProperty, ok := structType.properties[n]
			if !ok {
				// Didn't find the property, not equal
				return false
			}

			if !ourProperty.Equals(f) {
				// Different property, even though same name; not-equal
				return false
			}
		}

		if len(structType.functions) != len(st.functions) {
			// Different number of functions, not equal
			return false
		}

		for functionName, function := range st.functions {
			ourFunction, ok := structType.functions[functionName]
			if !ok {
				// Didn't find the func, not equal
				return false
			}

			if !ourFunction.Equals(function) {
				// Different function, even though same name; not-equal
				return false
			}
		}

		// All properties match, equal
		return true
	}

	return false
}

// WithProperty creates a new StructType with another property attached to it
// Properties are unique by name, so this can be used to Add and Replace a property
func (structType *StructType) WithProperty(property *PropertyDefinition) *StructType {
	// Create a copy of structType to preserve immutability
	result := structType.copy()
	result.properties[property.propertyName] = property

	return result
}

// WithProperties creates a new StructType with additional properties included
// Properties are unique by name, so this can be used to both Add and Replace properties.
func (structType *StructType) WithProperties(properties ...*PropertyDefinition) *StructType {
	// Create a copy of structType to preserve immutability
	result := structType.copy()
	for _, f := range properties {
		result.properties[f.propertyName] = f
	}

	return result
}

// WithFunction creates a new StructType with a function (method) attached to it
func (structType *StructType) WithFunction(name string, function Function) *StructType {
	// Create a copy of structType to preserve immutability
	result := structType.copy()
	result.functions[name] = function

	return result
}

func (structType *StructType) copy() *StructType {
	result := NewStructType()

	for key, value := range structType.properties {
		result.properties[key] = value
	}

	for key, value := range structType.functions {
		result.functions[key] = value
	}

	return result
}
