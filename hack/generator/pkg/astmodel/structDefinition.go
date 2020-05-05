/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
	"path/filepath"
)

// Package refernece indicates which package
// a struct belongs to.
type PackageReference struct {
	groupName   string
	packageName string
}

func (pr *PackageReference) PackagePath() string {
	return filepath.Join(pr.GroupName(), pr.PackageName())
}

func (pr *PackageReference) GroupName() string {
	return pr.groupName
}

func (pr *PackageReference) PackageName() string {
	return pr.packageName
}

// StructReference is the (versioned) name of a struct
// that can be used as a type
type StructReference struct {
	PackageReference
	name string
}

func NewStructReference(name string, group string, version string) StructReference {
	return StructReference{PackageReference{group, version}, name}
}

// assert that we implemented Type correctly
var _ Type = (*StructReference)(nil)

// AsType implements Type for StructReference
func (sr *StructReference) AsType() ast.Expr {
	return ast.NewIdent(sr.name)
}

func (sr *StructReference) RequiredImports() []PackageReference {
	return []PackageReference{sr.PackageReference}
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
func NewStructDefinition(ref StructReference, fields ...*FieldDefinition) *StructDefinition {
	return &StructDefinition{
		StructReference: ref,
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

func (definition *StructDefinition) RequiredImports() []PackageReference {
	var result []PackageReference
	for _, field := range definition.fields {
		for _, requiredImport := range field.FieldType().RequiredImports() {
			result = append(result, requiredImport)
		}
	}

	return result
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
