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

// ObjectType represents an (unnamed) object type
type ObjectType struct {
	properties map[PropertyName]*PropertyDefinition
	functions  map[string]Function
	interfaces map[TypeName]*InterfaceImplementation
}

// EmptyObjectType is an empty object
var EmptyObjectType = NewObjectType()

// Ensure ObjectType implements the Type interface correctly
var _ Type = (*ObjectType)(nil)

// NewObjectType is a factory method for creating a new ObjectType
func NewObjectType() *ObjectType {
	return &ObjectType{
		properties: make(map[PropertyName]*PropertyDefinition),
		functions:  make(map[string]Function),
		interfaces: make(map[TypeName]*InterfaceImplementation),
	}
}

func (objectType *ObjectType) AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description []string) []ast.Decl {
	identifier := ast.NewIdent(name.Name())
	declaration := &ast.GenDecl{
		Tok: token.TYPE,
		Doc: &ast.CommentGroup{},
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: identifier,
				Type: objectType.AsType(codeGenerationContext),
			},
		},
	}

	addDocComments(&declaration.Doc.List, description, 200)

	result := []ast.Decl{declaration}
	result = append(result, objectType.generateInterfaceDecls(codeGenerationContext, name)...)
	result = append(result, objectType.generateMethodDecls(codeGenerationContext, name)...)
	return result
}

func (objectType *ObjectType) generateMethodDecls(codeGenerationContext *CodeGenerationContext, typeName TypeName) []ast.Decl {
	var result []ast.Decl
	for methodName, function := range objectType.functions {
		funcDef := function.AsFunc(codeGenerationContext, typeName, methodName)
		result = append(result, funcDef)
	}

	return result
}

func (objectType *ObjectType) generateInterfaceDecls(codeGenerationContext *CodeGenerationContext, typeName TypeName) []ast.Decl {
	var result []ast.Decl

	// First interfaces must be ordered by name for deterministic output
	var ifaceNames []TypeName
	for ifaceName := range objectType.interfaces {
		ifaceNames = append(ifaceNames, ifaceName)
	}

	sort.Slice(ifaceNames, func(i int, j int) bool {
		return ifaceNames[i].name < ifaceNames[j].name
	})

	for _, ifaceName := range ifaceNames {
		iface := objectType.interfaces[ifaceName]

		result = append(result, objectType.generateInterfaceImplAssertion(codeGenerationContext, iface, typeName))

		var funcNames []string
		for funcName := range iface.functions {
			funcNames = append(funcNames, funcName)
		}

		sort.Strings(funcNames)

		for _, methodName := range funcNames {
			function := iface.functions[methodName]
			result = append(result, function.AsFunc(codeGenerationContext, typeName, methodName))
		}
	}

	return result
}

func (objectType *ObjectType) generateInterfaceImplAssertion(codeGenerationContext *CodeGenerationContext, iface *InterfaceImplementation, typeName TypeName) ast.Decl {

	ifacePackageName, err := codeGenerationContext.GetImportedPackageName(iface.name.PackageReference)
	if err != nil {
		panic(err)
	}

	typeAssertion := &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Type: &ast.SelectorExpr{
					X:   ast.NewIdent(ifacePackageName),
					Sel: ast.NewIdent(iface.name.name),
				},
				Names: []*ast.Ident{
					ast.NewIdent("_"),
				},
				Values: []ast.Expr{
					&ast.UnaryExpr{
						Op: token.AND,
						X: &ast.CompositeLit{
							Type: ast.NewIdent(typeName.name),
						},
					},
				},
			},
		},
	}

	return typeAssertion
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
func (objectType *ObjectType) Properties() []*PropertyDefinition {
	var result []*PropertyDefinition
	for _, property := range objectType.properties {
		result = append(result, property)
	}

	sort.Slice(result, func(left int, right int) bool {
		return result[left].propertyName < result[right].propertyName
	})

	return result
}

func (objectType *ObjectType) Property(name PropertyName) (*PropertyDefinition, bool) {
	prop, ok := objectType.properties[name]
	return prop, ok
}

// HasFunctionWithName determines if this object has a function with the given name
func (objectType *ObjectType) HasFunctionWithName(name string) bool {
	_, ok := objectType.functions[name]
	return ok
}

// AsType implements Type for ObjectType
func (objectType *ObjectType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {

	// Copy the slice of properties and sort it
	properties := objectType.Properties()
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
func (objectType *ObjectType) RequiredImports() []PackageReference {
	var result []PackageReference
	for _, property := range objectType.properties {
		result = append(result, property.PropertyType().RequiredImports()...)
	}

	for _, function := range objectType.functions {
		result = append(result, function.RequiredImports()...)
	}

	for _, i := range objectType.interfaces {
		result = append(result, i.RequiredImports()...)
	}

	return result
}

// References returns the set of all the types referred to by any property.
func (objectType *ObjectType) References() TypeNameSet {
	var results TypeNameSet
	for _, property := range objectType.properties {
		for ref := range property.PropertyType().References() {
			results = results.Add(ref)
		}
	}
	// Not collecting types from functions deliberately.
	return results
}

// Equals returns true if the passed type is a object type with the same properties, false otherwise
// The order of the properties is not relevant
func (objectType *ObjectType) Equals(t Type) bool {
	if objectType == t {
		return true
	}

	if st, ok := t.(*ObjectType); ok {
		if len(objectType.properties) != len(st.properties) {
			// Different number of properties, not equal
			return false
		}

		for n, f := range st.properties {
			ourProperty, ok := objectType.properties[n]
			if !ok {
				// Didn't find the property, not equal
				return false
			}

			if !ourProperty.Equals(f) {
				// Different property, even though same name; not-equal
				return false
			}
		}

		if len(objectType.functions) != len(st.functions) {
			// Different number of functions, not equal
			return false
		}

		for functionName, function := range st.functions {
			ourFunction, ok := objectType.functions[functionName]
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

// WithProperty creates a new ObjectType with another property attached to it
// Properties are unique by name, so this can be used to Add and Replace a property
func (objectType *ObjectType) WithProperty(property *PropertyDefinition) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.properties[property.propertyName] = property

	return result
}

// WithProperties creates a new ObjectType with additional properties included
// Properties are unique by name, so this can be used to both Add and Replace properties.
func (objectType *ObjectType) WithProperties(properties ...*PropertyDefinition) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	for _, f := range properties {
		result.properties[f.propertyName] = f
	}

	return result
}

// WithoutProperty creates a new ObjectType without the specified field
func (objectType *ObjectType) WithoutProperty(name PropertyName) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	delete(result.properties, name)

	return result
}

// WithFunction creates a new ObjectType with a function (method) attached to it
func (objectType *ObjectType) WithFunction(name string, function Function) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.functions[name] = function

	return result
}

// WithInterface creates a new ObjectType with a function (method) attached to it
func (objectType *ObjectType) WithInterface(iface *InterfaceImplementation) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := *objectType

	// Copy contents of old map into new map
	result.interfaces = make(map[TypeName]*InterfaceImplementation)
	for key, value := range objectType.interfaces {
		result.interfaces[key] = value
	}

	result.interfaces[iface.Name()] = iface

	return &result
}

func (objectType *ObjectType) copy() *ObjectType {
	result := NewObjectType()

	for key, value := range objectType.properties {
		result.properties[key] = value
	}

	for key, value := range objectType.functions {
		result.functions[key] = value
	}

	return result
}
