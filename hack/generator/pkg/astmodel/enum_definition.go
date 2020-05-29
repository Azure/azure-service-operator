/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
)

// EnumDefinition generates the full definition of an enumeration
type EnumDefinition struct {
	typeName    *TypeName
	baseType    *EnumType
	description *string
}

var _ TypeDefiner = (*EnumDefinition)(nil)

// NewEnumDefinition is a factory method for creating new Enum Definitions
func NewEnumDefinition(name *TypeName, t *EnumType) *EnumDefinition {
	return &EnumDefinition{typeName: name, baseType: t}
}

// Name returns the unique name to use for specifying this enumeration
func (enum *EnumDefinition) Name() *TypeName {
	return enum.typeName
}

// Type returns the underlying EnumerationType for this enum
func (enum *EnumDefinition) Type() Type {
	return enum.baseType
}

func (enum *EnumDefinition) WithDescription(description *string) TypeDefiner {
	result := *enum
	result.description = description
	return &result
}

// AsDeclarations generates the Go code representing this definition
func (enum *EnumDefinition) AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	var specs []ast.Spec
	for _, v := range enum.baseType.Options() {
		s := enum.createValueDeclaration(v)
		specs = append(specs, s)
	}

	declaration := &ast.GenDecl{
		Tok:   token.CONST,
		Doc:   &ast.CommentGroup{},
		Specs: specs,
	}

	result := []ast.Decl{
		enum.createBaseDeclaration(codeGenerationContext),
		declaration}

	return result
}

func (enum *EnumDefinition) createBaseDeclaration(codeGenerationContext *CodeGenerationContext) ast.Decl {
	identifier := ast.NewIdent(enum.typeName.name)

	typeSpecification := &ast.TypeSpec{
		Name: identifier,
		Type: enum.baseType.BaseType.AsType(codeGenerationContext),
	}

	declaration := &ast.GenDecl{
		Tok: token.TYPE,
		Doc: &ast.CommentGroup{},
		Specs: []ast.Spec{
			typeSpecification,
		},
	}

	return declaration
}

func (enum *EnumDefinition) createValueDeclaration(value EnumValue) ast.Spec {

	enumIdentifier := ast.NewIdent(enum.typeName.name)
	valueIdentifier := ast.NewIdent(enum.typeName.name + value.Identifier)

	valueLiteral := ast.BasicLit{
		Kind:  token.STRING,
		Value: value.Value,
	}

	valueSpec := &ast.ValueSpec{
		Names: []*ast.Ident{valueIdentifier},
		Values: []ast.Expr{
			&ast.CallExpr{
				Fun:  enumIdentifier,
				Args: []ast.Expr{&valueLiteral},
			},
		},
	}

	return valueSpec
}
