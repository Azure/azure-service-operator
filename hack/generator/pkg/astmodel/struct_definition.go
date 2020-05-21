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

// NewStructReference creates a new StructReference
func NewStructReference(name string, group string, version string, isResource bool) *StructReference {
	return &StructReference{DefinitionName{PackageReference{group, version}, name}, isResource}
}

// IsResource indicates that the struct is an Azure resource
func (sr *StructReference) IsResource() bool {
	return sr.isResource
}

// StructDefinition encapsulates the definition of a struct
type StructDefinition struct {
	StructReference *StructReference
	StructType      *StructType

	description string
}

// Ensure StructDefinition implements Definition interface correctly
var _ Definition = (*StructDefinition)(nil)

// Reference provides the definition name
func (definition *StructDefinition) Reference() *DefinitionName {
	return &definition.StructReference.DefinitionName
}

// Type provides the type of the struct
func (definition *StructDefinition) Type() Type {
	return definition.StructType
}

// NewStructDefinition is a factory method for creating a new StructDefinition
func NewStructDefinition(ref *StructReference, fields ...*FieldDefinition) *StructDefinition {
	return &StructDefinition{ref, NewStructType(fields...), ""}
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
	return *definition.StructType.fields[index]
}

// FieldCount indicates how many fields are contained
func (definition *StructDefinition) FieldCount() int {
	return len(definition.StructType.fields)
}

// RequiredImports returns a list of package required by this
func (definition *StructDefinition) RequiredImports() []PackageReference {
	var result []PackageReference
	for _, field := range definition.StructType.fields {
		result = append(result, field.FieldType().RequiredImports()...)
	}

	return result
}

// FileNameHint is a hint of what to name the file
func (definition *StructDefinition) FileNameHint() string {
	return definition.StructReference.Name()
}

// AsDeclarations generates an AST node representing this struct definition
func (definition *StructDefinition) AsDeclarations() []ast.Decl {
	var identifier *ast.Ident
	if definition.StructReference.IsResource() {
		// if it's a resource then this is the Spec type and we will generate
		// the non-spec type later:
		identifier = ast.NewIdent(definition.StructReference.name + "Spec")
	} else {
		identifier = ast.NewIdent(definition.StructReference.name)
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

	if definition.StructReference.IsResource() {
		resourceIdentifier := ast.NewIdent(definition.StructReference.name)

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

// CreateRelatedDefinitions implements the HasRelatedDefinitions interface for StructType
func (definition *StructDefinition) CreateRelatedDefinitions(ref PackageReference, namehint string, idFactory IdentifierFactory) []Definition {
	return definition.StructType.CreateRelatedDefinitions(ref, namehint, idFactory)
}

// TODO: metav1 import should be added via RequiredImports?
var typeMetaField = defineField("", "metav1.TypeMeta", "`json:\",inline\"`")
var objectMetaField = defineField("", "metav1.ObjectMeta", "`json:\"metadata,omitempty\"`")
