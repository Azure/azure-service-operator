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
	name        TypeName
	description []string
	theType     Type
}

func MakeTypeDefinition(name TypeName, theType Type) TypeDefinition {
	return TypeDefinition{name: name, theType: theType}
}

// Name returns the name being associated with the type
func (def TypeDefinition) Name() TypeName {
	return def.name
}

// Type returns the type being associated with the name
func (def TypeDefinition) Type() Type {
	return def.theType
}

// Description returns the description to be attached to this type definition (as a comment)
// We return a new slice to preserve immutability
func (def TypeDefinition) Description() []string {
	var result []string
	result = append(result, def.description...)
	return result
}

func (def TypeDefinition) References() TypeNameSet {
	return def.theType.References()
}

// WithDescription replaces the description of the definition with a new one (if any)
func (def TypeDefinition) WithDescription(desc []string) TypeDefinition {
	var d []string
	def.description = append(d, desc...)
	return def
}

// WithType returns an updated TypeDefinition with the specified type
func (def TypeDefinition) WithType(t Type) TypeDefinition {
	result := def
	result.theType = t
	return result
}

// WithName returns an updated TypeDefinition with the specified name
func (def TypeDefinition) WithName(typeName TypeName) TypeDefinition {
	result := def
	result.name = typeName
	return result
}

func (def TypeDefinition) AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	return def.theType.AsDeclarations(codeGenerationContext, def.name, def.description)
}

// AsSimpleDeclarations is a helper for types that only require a simple name/alias to be defined
func AsSimpleDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description []string, theType Type) []ast.Decl {
	var docComments *ast.CommentGroup
	if len(description) > 0 {
		docComments = &ast.CommentGroup{}
		addDocComments(&docComments.List, description, 120)
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
func (def TypeDefinition) RequiredImports() []PackageReference {
	return def.theType.RequiredImports()
}

// FileNameHint returns what a file that contains this name (if any) should be called
// this is not always used as we often combine multiple definitions into one file
func FileNameHint(name TypeName) string {
	return transformToSnakeCase(name.name)
}
