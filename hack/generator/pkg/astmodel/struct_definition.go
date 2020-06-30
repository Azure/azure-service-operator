/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
)

// StructDefinition encapsulates a complex (object) schema type
type StructDefinition struct {
	typeName    *TypeName
	structType  *StructType
	description *string
}

// NewStructDefinition creates a new StructDefinition
func NewStructDefinition(typeName *TypeName, structType *StructType) *StructDefinition {
	return &StructDefinition{typeName, structType, nil}
}

// ensure StructDefinition is a TypeDefiner
var _ TypeDefiner = &StructDefinition{}

// Name provides the type name
func (definition *StructDefinition) Name() *TypeName {
	return definition.typeName
}

// Type provides the type being linked to the name
func (definition *StructDefinition) Type() Type {
	return definition.structType
}

// WithDescription adds a description (doc-comment) to the definition
func (definition *StructDefinition) WithDescription(description *string) TypeDefiner {
	result := *definition
	result.description = description
	return &result
}

// RequiredImports returns a list of packages required by this
func (definition *StructDefinition) RequiredImports() []*PackageReference {
	return definition.structType.RequiredImports()
}

// AsDeclarations returns the Go AST declarations for this struct
func (definition *StructDefinition) AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	identifier := ast.NewIdent(definition.typeName.name)
	declaration := &ast.GenDecl{
		Tok: token.TYPE,
		Doc: &ast.CommentGroup{},
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: identifier,
				Type: definition.structType.AsType(codeGenerationContext),
			},
		},
	}

	if definition.description != nil {
		addDocComment(&declaration.Doc.List, *definition.description, 200)
	}

	result := []ast.Decl{declaration}
	result = append(result, definition.generateMethodDecls(codeGenerationContext)...)
	return result
}

func (definition *StructDefinition) generateMethodDecls(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	var result []ast.Decl
	for methodName, function := range definition.structType.functions {
		funcDef := function.AsFunc(codeGenerationContext, definition.typeName, methodName)
		result = append(result, funcDef)
	}

	return result
}

func defineField(fieldName string, typeName string, tag string) *ast.Field {

	result := &ast.Field{
		Type: ast.NewIdent(typeName),
		Tag:  &ast.BasicLit{Kind: token.STRING, Value: tag},
	}

	if fieldName != "" {
		result.Names = []*ast.Ident{ast.NewIdent(fieldName)}
	}

	return result
}
