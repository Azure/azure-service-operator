/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
	"sort"

	"github.com/pkg/errors"
)

// ObjectType represents an (unnamed) object type
type ObjectType struct {
	embedded   map[TypeName]*PropertyDefinition
	properties map[PropertyName]*PropertyDefinition
	functions  map[string]Function
	InterfaceImplementer
}

// EmptyObjectType is an empty object
var EmptyObjectType = NewObjectType()

// Ensure ObjectType implements the Type interface correctly
var _ Type = &ObjectType{}

// NewObjectType is a factory method for creating a new ObjectType
func NewObjectType() *ObjectType {
	return &ObjectType{
		embedded:             make(map[TypeName]*PropertyDefinition),
		properties:           make(map[PropertyName]*PropertyDefinition),
		functions:            make(map[string]Function),
		InterfaceImplementer: MakeInterfaceImplementer(),
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

	addWrappedComments(&declaration.Doc.List, description, 200)

	result := []ast.Decl{declaration}
	result = append(result, objectType.InterfaceImplementer.AsDeclarations(codeGenerationContext, name, nil)...)
	result = append(result, objectType.generateMethodDecls(codeGenerationContext, name)...)
	return result
}

func (objectType *ObjectType) generateMethodDecls(codeGenerationContext *CodeGenerationContext, typeName TypeName) []ast.Decl {
	var result []ast.Decl

	// Functions must be ordered by name for deterministic output
	var functions []Function
	for _, f := range objectType.functions {
		functions = append(functions, f)
	}

	sort.Slice(functions, func(i int, j int) bool {
		return functions[i].Name() < functions[j].Name()
	})

	for _, f := range functions {
		funcDef := f.AsFunc(codeGenerationContext, typeName)
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

// Properties returns all the property definitions
// A sorted slice is returned to preserve immutability and provide determinism
func (objectType *ObjectType) Properties() []*PropertyDefinition {
	var result []*PropertyDefinition
	for _, property := range objectType.properties {
		result = append(result, property)
	}

	// Sorted so that it's always consistent
	sort.Slice(result, func(left int, right int) bool {
		return result[left].propertyName < result[right].propertyName
	})

	return result
}

// Property returns the details of a specific property based on its unique case sensitive name
func (objectType *ObjectType) Property(name PropertyName) (*PropertyDefinition, bool) {
	prop, ok := objectType.properties[name]
	return prop, ok
}

// EmbeddedProperties returns all the embedded properties
// A sorted slice is returned to preserve immutability and provide determinism
func (objectType *ObjectType) EmbeddedProperties() []*PropertyDefinition {
	var result []*PropertyDefinition
	for _, embedded := range objectType.embedded {
		result = append(result, embedded)
	}

	sort.Slice(result, func(left int, right int) bool {
		lTypeName, err := extractEmbeddedTypeName(result[left].PropertyType())
		if err != nil {
			// It shouldn't be possible to get an invalid embedded type here, as we should have
			// failed to add it.
			panic(err)
		}
		rTypeName, err := extractEmbeddedTypeName(result[right].PropertyType())
		if err != nil {
			// It shouldn't be possible to get an invalid embedded type here, as we should have
			// failed to add it.
			panic(err)
		}
		return lTypeName.Name() < rTypeName.Name()
	})

	return result
}

// HasFunctionWithName determines if this object has a function with the given name
func (objectType *ObjectType) HasFunctionWithName(name string) bool {
	_, ok := objectType.functions[name]
	return ok
}

// AsType implements Type for ObjectType
func (objectType *ObjectType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	embedded := objectType.EmbeddedProperties()
	properties := objectType.Properties()
	var fields []*ast.Field

	for _, f := range embedded {
		fields = append(fields, f.AsField(codeGenerationContext))
	}

	for _, f := range properties {
		fields = append(fields, f.AsField(codeGenerationContext))
	}

	return &ast.StructType{
		Fields: &ast.FieldList{
			List: fields,
		},
	}
}

// RequiredPackageReferences returns a list of packages required by this
func (objectType *ObjectType) RequiredPackageReferences() []PackageReference {
	var result []PackageReference
	for _, property := range objectType.embedded {
		result = append(result, property.PropertyType().RequiredPackageReferences()...)
	}

	for _, property := range objectType.properties {
		result = append(result, property.PropertyType().RequiredPackageReferences()...)
	}

	for _, function := range objectType.functions {
		result = append(result, function.RequiredPackageReferences()...)
	}

	result = append(result, objectType.InterfaceImplementer.RequiredPackageReferences()...)

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

	other, ok := t.(*ObjectType)
	if !ok {
		return false
	}

	if len(objectType.embedded) != len(other.embedded) {
		// Different number of embedded properties, not equal
		return false
	}

	for n, f := range other.embedded {
		ourProperty, ok := objectType.embedded[n]
		if !ok {
			// Didn't find the property, not equal
			return false
		}

		if !ourProperty.Equals(f) {
			// Different property, even though same name; not-equal
			return false
		}
	}

	if len(objectType.properties) != len(other.properties) {
		// Different number of properties, not equal
		return false
	}

	for n, f := range other.properties {
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

	if len(objectType.functions) != len(other.functions) {
		// Different number of functions, not equal
		return false
	}

	for functionName, function := range other.functions {
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

	return objectType.InterfaceImplementer.Equals(other.InterfaceImplementer)
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

// WithEmbeddedProperties creates a new ObjectType with the additional embedded properties included
func (objectType *ObjectType) WithEmbeddedProperties(properties ...*PropertyDefinition) (*ObjectType, error) {

	// Create a copy of objectType to preserve immutability
	result := objectType.copy()

	for _, p := range properties {
		err := objectType.checkEmbeddedProperty(p)
		if err != nil {
			return objectType, err
		}

		typeName, err := extractEmbeddedTypeName(p.PropertyType())
		if err != nil {
			return objectType, err
		}

		result.embedded[typeName] = p
	}

	return result, nil
}

// WithoutProperties creates a new ObjectType from this one but
// without any properties.
func (objectType *ObjectType) WithoutProperties() *ObjectType {
	result := objectType.copy()
	result.properties = make(map[PropertyName]*PropertyDefinition)
	return result
}

// WithoutProperty creates a new ObjectType without the specified field
func (objectType *ObjectType) WithoutProperty(name PropertyName) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	delete(result.properties, name)

	return result
}

// WithoutEmbeddedProperties creates a new ObjectType from this one but
// without any embedded properties.
func (objectType *ObjectType) WithoutEmbeddedProperties() *ObjectType {
	result := objectType.copy()
	result.embedded = make(map[TypeName]*PropertyDefinition)
	return result
}

func (objectType *ObjectType) checkEmbeddedProperty(property *PropertyDefinition) error {
	// There are certain expectations about the shape of the provided property -- namely
	// it must not have a name and its Type must be TypeName or Optional[TypeName]
	if property.PropertyName() != "" {
		return errors.Errorf("embedded property name must be empty, was: %s", property.PropertyName())
	}

	return nil
}

// WithEmbeddedProperty creates a new ObjectType with another embedded property
func (objectType *ObjectType) WithEmbeddedProperty(property *PropertyDefinition) (*ObjectType, error) {
	err := objectType.checkEmbeddedProperty(property)
	if err != nil {
		return objectType, err
	}

	typeName, err := extractEmbeddedTypeName(property.PropertyType())
	if err != nil {
		return objectType, err
	}

	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.embedded[typeName] = property

	return result, nil
}

// WithFunction creates a new ObjectType with a function (method) attached to it
func (objectType *ObjectType) WithFunction(function Function) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.functions[function.Name()] = function

	return result
}

// WithInterface creates a new ObjectType with a function (method) attached to it
func (objectType *ObjectType) WithInterface(iface *InterfaceImplementation) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.InterfaceImplementer = result.InterfaceImplementer.WithInterface(iface)
	return result
}

func (objectType *ObjectType) copy() *ObjectType {
	result := NewObjectType()

	for key, value := range objectType.embedded {
		result.embedded[key] = value
	}

	for key, value := range objectType.properties {
		result.properties[key] = value
	}

	for key, value := range objectType.functions {
		result.functions[key] = value
	}

	result.InterfaceImplementer = objectType.InterfaceImplementer.copy()

	return result
}

// String implements fmt.Stringer
func (objectType *ObjectType) String() string {
	return "(object)"
}

// IsObjectType returns true if the passed type is an object type OR if it is a wrapper type containing an object type
func IsObjectType(t Type) bool {
	_, ok := t.(*ObjectType)
	return ok
}

// IsObjectDefinition returns true if the passed definition is for a Arm type; false otherwise.
func IsObjectDefinition(definition TypeDefinition) bool {
	return IsObjectType(definition.theType)
}

func extractEmbeddedTypeName(t Type) (TypeName, error) {
	typeName, isTypeName := t.(TypeName)
	optionalType, isOptionalType := t.(*OptionalType)
	if isOptionalType {
		typeName, isTypeName = optionalType.Element().(TypeName)
	}

	if !isTypeName {
		return TypeName{}, errors.Errorf("embedded property type must be TypeName or Optional[TypeName], was: %T", t)
	}

	return typeName, nil
}
