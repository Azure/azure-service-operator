/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"sort"
)

// StructType represents an (unnamed) struct type
type StructType struct {
	fields    map[FieldName]*FieldDefinition
	functions map[string]Function
}

// EmptyStructType is an empty struct
var EmptyStructType = NewStructType()

// Ensure StructType implements the Type interface correctly
var _ Type = (*StructType)(nil)

// NewStructType is a factory method for creating a new StructTypeDefinition
func NewStructType() *StructType {
	return &StructType{
		fields:    make(map[FieldName]*FieldDefinition),
		functions: make(map[string]Function),
	}
}

// Fields returns all our field definitions
// A sorted slice is returned to preserve immutability and provide determinism
func (structType *StructType) Fields() []*FieldDefinition {
	var result []*FieldDefinition
	for _, field := range structType.fields {
		result = append(result, field)
	}

	sort.Slice(result, func(left int, right int) bool {
		return result[left].fieldName < result[right].fieldName
	})

	return result
}

// AsType implements Type for StructType
func (structType *StructType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {

	// Copy the slice of fields and sort it
	fields := structType.Fields()
	sort.Slice(fields, func(i int, j int) bool {
		return fields[i].fieldName < fields[j].fieldName
	})

	fieldDefinitions := make([]*ast.Field, len(fields))
	for i, f := range fields {
		fieldDefinitions[i] = f.AsField(codeGenerationContext)
	}

	return &ast.StructType{
		Fields: &ast.FieldList{
			List: fieldDefinitions,
		},
	}
}

// RequiredImports returns a list of packages required by this
func (structType *StructType) RequiredImports() []*PackageReference {
	var result []*PackageReference
	for _, field := range structType.fields {
		result = append(result, field.FieldType().RequiredImports()...)
	}

	for _, function := range structType.functions {
		result = append(result, function.RequiredImports()...)
	}

	return result
}

// References this type has to the given type
func (structType *StructType) References(d *TypeName) bool {

	for _, field := range structType.fields {
		if field.FieldType().References(d) {
			return true
		}
	}

	// For now, not considering functions in references on purpose

	return false
}

// Equals returns true if the passed type is a struct type with the same fields, false otherwise
// The order of the fields is not relevant
func (structType *StructType) Equals(t Type) bool {
	if structType == t {
		return true
	}

	if st, ok := t.(*StructType); ok {
		if len(structType.fields) != len(st.fields) {
			// Different number of fields, not equal
			return false
		}

		for n, f := range st.fields {
			ourField, ok := structType.fields[n]
			if !ok {
				// Didn't find the field, not equal
				return false
			}

			if !ourField.Equals(f) {
				// Different field, even though same name; not-equal
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

		// All fields match, equal
		return true
	}

	return false
}

// CreateInternalDefinitions defines a named type for this struct and returns that type to be used in-place
// of the anonymous struct type. This is needed for controller-gen to work correctly:
func (structType *StructType) CreateInternalDefinitions(name *TypeName, idFactory IdentifierFactory) (Type, []TypeDefiner) {
	// an internal struct must always be named:
	definedStruct, otherTypes := structType.CreateDefinitions(name, idFactory)
	return definedStruct.Name(), append(otherTypes, definedStruct)
}

// CreateDefinitions defines a named type for this struct and invokes CreateInternalDefinitions for each field type
// to instantiate any definitions required by internal types.
func (structType *StructType) CreateDefinitions(name *TypeName, idFactory IdentifierFactory) (TypeDefiner, []TypeDefiner) {

	var otherTypes []TypeDefiner
	var newFields []*FieldDefinition

	for _, field := range structType.fields {

		// create definitions for nested types
		nestedName := name.Name() + string(field.fieldName)
		nameHint := NewTypeName(name.PackageReference, nestedName)
		newFieldType, moreTypes := field.fieldType.CreateInternalDefinitions(nameHint, idFactory)

		otherTypes = append(otherTypes, moreTypes...)
		newFields = append(newFields, field.WithType(newFieldType))
	}

	newStructType := NewStructType().WithFields(newFields...)
	for functionName, function := range structType.functions {
		newStructType.functions[functionName] = function
	}

	return NewStructDefinition(name, newStructType), otherTypes
}

// WithField creates a new StructType with another field attached to it
func (structType *StructType) WithField(field *FieldDefinition) *StructType {
	// Create a copy of structType to preserve immutability
	result := structType.copy()
	result.fields[field.fieldName] = field

	return result
}

// WithFields creates a new StructType with additional fields attached to it
func (structType *StructType) WithFields(fields ...*FieldDefinition) *StructType {
	// Create a copy of structType to preserve immutability
	result := structType.copy()
	for _, f := range fields {
		result.fields[f.fieldName] = f
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

	for key, value := range structType.fields {
		result.fields[key] = value
	}

	for key, value := range structType.functions {
		result.functions[key] = value
	}

	return result
}
