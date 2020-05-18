/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
	"go/token"
)

// FieldName is a semantic type
type FieldName string

// FieldDefinition encapsulates the definition of a field
type FieldDefinition struct {
	fieldName   FieldName
	fieldType   Type
	jsonName    string
	description string
	validations []Validation
}

// NewFieldDefinition is a factory method for creating a new FieldDefinition
// name is the name for the new field (mandatory)
// fieldType is the type for the new field (mandatory)
func NewFieldDefinition(fieldName FieldName, jsonName string, fieldType Type) *FieldDefinition {
	return &FieldDefinition{
		fieldName:   fieldName,
		fieldType:   fieldType,
		jsonName:    jsonName,
		description: "",
	}
}

// NewEmbeddedStructDefinition is a factory method for defining an embedding
// of another struct type.
func NewEmbeddedStructDefinition(structType Type) *FieldDefinition {
	// in Go, this is just a field without a name:
	return &FieldDefinition{
		fieldName:   "",
		fieldType:   structType,
		jsonName:    "",
		description: "",
	}
}

// FieldName returns the name of the field
func (field *FieldDefinition) FieldName() FieldName {
	return field.fieldName
}

// FieldType returns the data type of the field
func (field *FieldDefinition) FieldType() Type {
	return field.fieldType
}

// WithDescription returns a new FieldDefinition with the specified description
func (field *FieldDefinition) WithDescription(description *string) *FieldDefinition {
	if description == nil {
		return field
	}

	result := *field
	result.description = *description
	return &result
}

// WithValidation adds the given validation to the field's set of validations
func (field *FieldDefinition) WithValidation(validation Validation) *FieldDefinition {
	result := *field
	result.validations = append(result.validations, validation)
	return &result
}

// MakeRequired returns a new FieldDefinition that is marked as required
func (field *FieldDefinition) MakeRequired() *FieldDefinition {
	return field.WithValidation(ValidateRequired())
}

// MakeOptional returns a new FieldDefinition that has an optional value
func (field *FieldDefinition) MakeOptional() *FieldDefinition {
	result := *field
	result.fieldType = NewOptionalType(result.fieldType)
	return &result
}

// AsField generates an AST field node representing this field definition
func (field *FieldDefinition) AsField() *ast.Field {

	result := &ast.Field{
		Doc:   &ast.CommentGroup{},
		Names: []*ast.Ident{ast.NewIdent(string(field.fieldName))},
		Type:  field.FieldType().AsType(),
		Tag: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("`json:%q`", field.jsonName),
		},
	}

	addDocComment := func(comment string) {
		newLine := ""
		if result.Doc.List == nil {
			// if first comment, add a newline
			newLine = "\n"
		}

		result.Doc.List = append(result.Doc.List, &ast.Comment{
			Text: newLine + comment,
		})
	}

	// generate validation comments:
	for _, validation := range field.validations {
		// these are not doc comments but they must go here to be emitted before the field
		addDocComment(GenerateKubebuilderComment(validation))
	}

	// generate doc comment:
	if field.description != "" {
		addDocComment(fmt.Sprintf("/* %s: %s */", field.fieldName, field.description))
	}

	return result
}

// Equals tests to see if the specified FieldDefinition specifies the same field
func (field *FieldDefinition) Equals(f *FieldDefinition) bool {
	return field == f || (field.fieldName == f.fieldName && field.fieldType.Equals(f.fieldType))
}
