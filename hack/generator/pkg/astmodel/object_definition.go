/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
)

// ObjectDefinition encapsulates a complex (object) schema type
type ObjectDefinition struct {
	typeName    *TypeName
	objectType  *ObjectType
	description *string
}

// NewObjectDefinition creates a new ObjectDefinition
func NewObjectDefinition(typeName *TypeName, objectType *ObjectType) *ObjectDefinition {
	return &ObjectDefinition{typeName, objectType, nil}
}

// ensure ObjectDefinition is a TypeDefiner
var _ TypeDefiner = &ObjectDefinition{}

// Name provides the type name
func (definition *ObjectDefinition) Name() *TypeName {
	return definition.typeName
}

// Type provides the type being linked to the name
func (definition *ObjectDefinition) Type() Type {
	return definition.objectType
}

// Description returns the description of the object
func (definition *ObjectDefinition) Description() string {
	return *definition.description
}

// References returns the types referenced by the object type
func (definition *ObjectDefinition) References() TypeNameSet {
	return definition.objectType.References()
}

// WithDescription adds a description (doc-comment) to the definition
func (definition *ObjectDefinition) WithDescription(description *string) TypeDefiner {
	result := *definition
	result.description = description
	return &result
}

// RequiredImports returns a list of packages required by this
func (definition *ObjectDefinition) RequiredImports() []*PackageReference {
	return definition.objectType.RequiredImports()
}

// AsDeclarations returns the Go AST declarations for this object
func (definition *ObjectDefinition) AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	identifier := ast.NewIdent(definition.typeName.name)
	declaration := &ast.GenDecl{
		Tok: token.TYPE,
		Doc: &ast.CommentGroup{},
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: identifier,
				Type: definition.objectType.AsType(codeGenerationContext),
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

func (definition *ObjectDefinition) generateMethodDecls(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	var result []ast.Decl
	for methodName, function := range definition.objectType.functions {
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
