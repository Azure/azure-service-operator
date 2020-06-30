/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
)

// SimpleTypeDefiner is a TypeDefiner for simple cases (not structs or enums)
type SimpleTypeDefiner struct {
	name        *TypeName
	description *string
	theType     Type
}

func NewSimpleTypeDefiner(name *TypeName, theType Type) *SimpleTypeDefiner {
	return &SimpleTypeDefiner{name: name, theType: theType}
}

// SimpleTypeDefiner is a TypeDefiner
var _ TypeDefiner = (*SimpleTypeDefiner)(nil)

func (std *SimpleTypeDefiner) Name() *TypeName {
	return std.name
}

func (std *SimpleTypeDefiner) Type() Type {
	return std.theType
}

func (std *SimpleTypeDefiner) WithDescription(desc *string) TypeDefiner {
	result := *std
	result.description = desc
	return &result
}

func (std *SimpleTypeDefiner) AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	var docComments *ast.CommentGroup
	if std.description != nil {
		docComments = &ast.CommentGroup{}
		addDocComment(&docComments.List, *std.description, 120)
	}

	return []ast.Decl{
		&ast.GenDecl{
			Doc: docComments,
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: ast.NewIdent(std.name.name),
					Type: std.theType.AsType(codeGenerationContext),
				},
			},
		},
	}
}

// RequiredImports returns a list of packages required by this type
func (std *SimpleTypeDefiner) RequiredImports() []*PackageReference {
	return std.theType.RequiredImports()
}
