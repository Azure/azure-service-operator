/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
)

// TypeDefinition is a name paired with a type
type TypeDefinition struct {
	name        *TypeName
	description *string
	theType     Type
}

func MakeTypeDefinition(name *TypeName, theType Type) TypeDefinition {
	return TypeDefinition{name: name, theType: theType}
}

func (std *TypeDefinition) Name() *TypeName {
	return std.name
}

func (std *TypeDefinition) Type() Type {
	return std.theType
}

func (std *TypeDefinition) References() TypeNameSet {
	return std.theType.References()
}

func (std *TypeDefinition) WithDescription(desc *string) TypeDefinition {
	result := *std
	result.description = desc
	return result
}

func (std *TypeDefinition) AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	return std.theType.AsDeclarations(codeGenerationContext, std.name, std.description)
}

// AsSimpleDeclarations is a helper for types that only require a simple name/alias to be defined
func AsSimpleDeclarations(codeGenerationContext *CodeGenerationContext, name *TypeName, description *string, theType Type) []ast.Decl {
	var docComments *ast.CommentGroup
	if description != nil {
		docComments = &ast.CommentGroup{}
		addDocComment(&docComments.List, *description, 120)
	}

	return []ast.Decl{
		&ast.GenDecl{
			Doc: docComments,
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: ast.NewIdent(name.Name()),
					Type: theType.AsType(codeGenerationContext),
				},
			},
		},
	}
}

// RequiredImports returns a list of packages required by this type
func (std *TypeDefinition) RequiredImports() []*PackageReference {
	return std.theType.RequiredImports()
}

// FileNameHint returns what a file that contains this definition (if any) should be called
// this is not always used as we might combine multiple definitions into one file
func FileNameHint(def TypeDefinition) string {
	return transformToSnakeCase(def.Name().name)
}
