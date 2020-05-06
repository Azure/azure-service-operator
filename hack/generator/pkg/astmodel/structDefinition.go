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
	DefinitionName
	isResource bool // this might seem like a strange place to have this, but it affects how the struct is referenced
}

func NewStructReference(name string, group string, version string, isResource bool) StructReference {
	return StructReference{DefinitionName{PackageReference{group, version}, name}, isResource}
}

func (sr *StructReference) IsResource() bool {
	return sr.isResource
}

// StructDefinition encapsulates the definition of a struct
type StructDefinition struct {
	StructReference
	StructType

	description string
}

// StructDefinition must implement Definition
var _ Definition = (*StructDefinition)(nil)

func (definition *StructDefinition) Reference() *DefinitionName {
	return &definition.DefinitionName
}

func (definition *StructDefinition) Type() Type {
	return &definition.StructType
}

// NewStructDefinition is a factory method for creating a new StructDefinition
func NewStructDefinition(ref StructReference, fields ...*FieldDefinition) *StructDefinition {
	return &StructDefinition{ref, StructType{fields}, ""}
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

// Field provides indexed access to our fields
func (definition *StructDefinition) Field(index int) FieldDefinition {
	return *definition.fields[index]
}

// FieldCount indicates how many fields are contained
func (definition *StructDefinition) FieldCount() int {
	return len(definition.fields)
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

func (definition *StructDefinition) FileNameHint() string {
	return definition.Name()
}

// AsDeclaration generates an AST node representing this struct definition
func (definition *StructDefinition) AsDeclarations() []ast.Decl {

	var identifier *ast.Ident
	if definition.IsResource() {
		// if it's a resource then this is the Spec type and we will generate
		// the non-spec type later:
		identifier = ast.NewIdent(definition.name + "Spec")
	} else {
		identifier = ast.NewIdent(definition.name)
	}

	typeSpecification := &ast.TypeSpec{
		Name: identifier,
		Type: definition.StructType.AsType(),
	}

	declaration := &ast.GenDecl{
		Tok: token.TYPE,
		Doc: &ast.CommentGroup{},
		Specs: []ast.Spec{
			typeSpecification,
		},
	}

	if definition.description != "" {
		declaration.Doc.List = append(declaration.Doc.List,
			&ast.Comment{Text: "\n/* " + definition.description + " */"})
	}

	declarations := []ast.Decl{declaration}

	if definition.IsResource() {
		resourceIdentifier := ast.NewIdent(definition.name)

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

	return declarations
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
