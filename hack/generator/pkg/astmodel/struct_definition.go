/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
)

// StructDefinition encapsulates the definition of a struct
type StructDefinition struct {
	TypeName    *TypeName
	StructType  *StructType
	isResource  bool
	description *string
}

// IsResource indicates if this is a ARM resource and should be a kubebuilder root
func (definition *StructDefinition) IsResource() bool {
	return definition.isResource
}

// Ensure StructDefinition implements TypeDefiner interface correctly
var _ TypeDefiner = (*StructDefinition)(nil)

// Name provides the struct name
func (definition *StructDefinition) Name() *TypeName {
	return definition.TypeName
}

// Type provides the type of the struct
func (definition *StructDefinition) Type() Type {
	return definition.StructType
}

// NewStructDefinition is a factory method for creating a new StructDefinition
func NewStructDefinition(name *TypeName, structType *StructType, isResource bool) *StructDefinition {
	return &StructDefinition{name, structType, isResource, nil}
}

// WithDescription adds a description (doc-comment) to the struct
func (definition *StructDefinition) WithDescription(description *string) TypeDefiner {
	result := *definition
	result.description = description
	return &result
}

// Field provides indexed access to our fields
func (definition *StructDefinition) Field(index int) FieldDefinition {
	return *definition.StructType.fields[index]
}

// FieldCount indicates how many fields are contained
func (definition *StructDefinition) FieldCount() int {
	return len(definition.StructType.fields)
}

// AsDeclarations generates an AST node representing this struct definition
func (definition *StructDefinition) AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	var identifier *ast.Ident
	if definition.IsResource() {
		// if it's a resource then this is the Spec type and we will generate
		// the non-spec type later:
		identifier = ast.NewIdent(definition.Name().name + "Spec")
	} else {
		identifier = ast.NewIdent(definition.Name().name)
	}

	typeSpecification := &ast.TypeSpec{
		Name: identifier,
		Type: definition.StructType.AsType(codeGenerationContext),
	}

	declaration := &ast.GenDecl{
		Tok: token.TYPE,
		Doc: &ast.CommentGroup{},
		Specs: []ast.Spec{
			typeSpecification,
		},
	}

	if definition.description != nil {
		declaration.Doc.List = append(declaration.Doc.List,
			&ast.Comment{Text: "\n/*" + *definition.description + "*/"})
	}

	declarations := []ast.Decl{declaration}

	if definition.IsResource() {
		resourceIdentifier := ast.NewIdent(definition.Name().name)

		/*
			start off with:
				metav1.TypeMeta   `json:",inline"`
				metav1.ObjectMeta `json:"metadata,omitempty"`

			then the Spec field
		*/
		resourceTypeSpec := &ast.TypeSpec{
			Name: resourceIdentifier,
			Type: &ast.StructType{
				Fields: &ast.FieldList{
					List: []*ast.Field{
						typeMetaField,
						objectMetaField,
						defineField("Spec", identifier.Name, "`json:\"spec,omitempty\"`"),
					},
				},
			},
		}

		resourceDeclaration := &ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{resourceTypeSpec},
			Doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "// +kubebuilder:object:root=true\n",
					},
				},
			},
		}

		declarations = append(declarations, resourceDeclaration)
	}

	// Append the methods
	declarations = append(declarations, definition.generateMethodDecls(codeGenerationContext)...)

	return declarations
}

func (definition *StructDefinition) generateMethodDecls(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	var result []ast.Decl
	for methodName, function := range definition.StructType.functions {
		funcDef := function.AsFunc(codeGenerationContext, definition.Name(), methodName)
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

// TODO: metav1 import should be added via RequiredImports?
var typeMetaField = defineField("", "metav1.TypeMeta", "`json:\",inline\"`")
var objectMetaField = defineField("", "metav1.ObjectMeta", "`json:\"metadata,omitempty\"`")
