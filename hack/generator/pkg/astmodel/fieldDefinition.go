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

// FieldDefinition encapsulates the definition of a field
type FieldDefinition struct {
	fieldName   string
	fieldType   Type
	jsonName    string
	description string
}

// NewFieldDefinition is a factory method for creating a new FieldDefinition
// name is the name for the new field (mandatory)
// fieldType is the type for the new field (mandatory)
func NewFieldDefinition(fieldName string, jsonName string, fieldType Type) *FieldDefinition {
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
func (field *FieldDefinition) FieldName() string {
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

// AsField generates an AST field node representing this field definition
func (field *FieldDefinition) AsField() *ast.Field {

	// TODO: add field tags for api hints / json binding
	result := &ast.Field{
		Names: []*ast.Ident{ast.NewIdent(field.fieldName)},
		Type:  field.FieldType().AsType(),
		Tag: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("`json:%q`", field.jsonName),
		},
	}

	if field.description != "" {
		result.Doc = &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: fmt.Sprintf("\n/* %s: %s */", field.fieldName, field.description),
				},
			},
		}
	}

	return result
}
