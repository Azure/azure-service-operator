/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"sort"
)

// ObjectType represents an (unnamed) object type
type ObjectType struct {
	properties map[PropertyName]*PropertyDefinition
	functions  map[string]Function
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
	}
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
func (objectType *ObjectType) RequiredImports() []*PackageReference {
	var result []*PackageReference
	for _, property := range objectType.properties {
		result = append(result, property.PropertyType().RequiredImports()...)
	}

	for _, function := range objectType.functions {
		result = append(result, function.RequiredImports()...)
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

// CreateInternalDefinitions defines a named type for this object and returns that type to be used in-place
// of the anonymous object type. This is needed for controller-gen to work correctly:
func (objectType *ObjectType) CreateInternalDefinitions(name *TypeName, idFactory IdentifierFactory) (Type, []TypeDefiner) {
	// an internal object must always be named:
	definedObject, otherTypes := objectType.CreateDefinitions(name, idFactory)
	return definedObject.Name(), append(otherTypes, definedObject)
}

// CreateDefinitions defines a named type for this object and invokes CreateInternalDefinitions for each property type
// to instantiate any definitions required by internal types.
func (objectType *ObjectType) CreateDefinitions(name *TypeName, idFactory IdentifierFactory) (TypeDefiner, []TypeDefiner) {

	var otherTypes []TypeDefiner
	var newProperties []*PropertyDefinition

	for _, property := range objectType.properties {

		// create definitions for nested types
		nestedName := name.Name() + string(property.propertyName)
		nameHint := NewTypeName(name.PackageReference, nestedName)
		newPropertyType, moreTypes := property.propertyType.CreateInternalDefinitions(nameHint, idFactory)

		otherTypes = append(otherTypes, moreTypes...)
		newProperties = append(newProperties, property.WithType(newPropertyType))
	}

	newObjectType := NewObjectType().WithProperties(newProperties...)
	for functionName, function := range objectType.functions {
		newObjectType.functions[functionName] = function
	}

	return NewObjectDefinition(name, newObjectType), otherTypes
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

// WithFunction creates a new ObjectType with a function (method) attached to it
func (objectType *ObjectType) WithFunction(name string, function Function) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.functions[name] = function

	return result
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
