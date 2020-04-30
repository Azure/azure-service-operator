/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
)

// StructReference is the (versioned) name of a struct
// that can be used as a type
type StructReference struct {
	name    string
	version string
}

// AsType implements Type for StructReference
func (sr *StructReference) AsType() ast.Expr {
	// TODO: namespaces/versions
	return ast.NewIdent(sr.name)
}

// StructDefinition encapsulates the definition of a struct
type StructDefinition struct {
	StructReference
	StructType

	description string
}

// StructDefinition must implement Definition
var _ Definition = &StructDefinition{}

// NewStructDefinition is a factory method for creating a new StructDefinition
func NewStructDefinition(name string, version string, fields ...*FieldDefinition) *StructDefinition {
	return &StructDefinition{
		StructReference: StructReference{name, version},
		StructType:      StructType{fields},
	}
}

// WithDescription adds a description (doc-comment) to the struct
func (definition *StructDefinition) WithDescription(description *string) *StructDefinition {
	if description == nil {
		return definition
	}

	result := *definition
	result.description = *description
	return &result
}

// Name returns the name of the struct
func (definition *StructDefinition) Name() string {
	return definition.name
}

// Version returns the version of this struct
func (definition *StructDefinition) Version() string {
	return definition.version
}

// Field provides indexed access to our fields
func (definition *StructDefinition) Field(index int) FieldDefinition {
	return *definition.fields[index]
}

// FieldCount indicates how many fields are contained
func (definition *StructDefinition) FieldCount() int {
	return len(definition.fields)
}

// AsAst generates an AST node representing this field definition
func (definition *StructDefinition) AsAst() ast.Node {
	return definition.AsDeclaration()
}

// AsDeclaration generates an AST node representing this struct definition
func (definition *StructDefinition) AsDeclaration() *ast.GenDecl {

	identifier := ast.NewIdent(definition.name)

	typeSpecification := &ast.TypeSpec{
		Name: identifier,
		Type: definition.StructType.AsType(),
	}

	declaration := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			typeSpecification,
		},
	}

	if definition.description != "" {
		declaration.Doc = &ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "\n/* " + definition.description + " */"},
			},
		}
	}

	return declaration
}
