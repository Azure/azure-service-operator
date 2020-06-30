/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
	"go/token"
)

// CreateResourceDefinitions creates definitions for a resource
func CreateResourceDefinitions(name *TypeName, specType *StructType, statusType *StructType, idFactory IdentifierFactory) (TypeDefiner, []TypeDefiner) {

	var others []TypeDefiner

	defineStruct := func(suffix string, structType *StructType) *TypeName {
		definedName := NewTypeName(name.PackageReference, name.Name()+suffix)
		defined, definedOthers := structType.CreateDefinitions(definedName, idFactory)
		others = append(append(others, defined), definedOthers...)
		return definedName
	}

	var specName *TypeName
	if specType == nil {
		panic("spec must always be provided")
	} else {
		specName = defineStruct("Spec", specType)
	}

	var statusName *TypeName
	if statusType != nil {
		statusName = defineStruct("Status", statusType)
	}

	this := &ResourceDefinition{typeName: name, spec: specName, status: statusName, isStorageVersion: false}

	return this, others
}

// ResourceDefinition represents an ARM resource
type ResourceDefinition struct {
	typeName         *TypeName
	spec             *TypeName
	status           *TypeName
	isStorageVersion bool
	description      *string
}

// assert that ResourceDefinition implements TypeDefiner
var _ TypeDefiner = &ResourceDefinition{}

// Name returns the name of the type being defined
func (definition *ResourceDefinition) Name() *TypeName {
	return definition.typeName
}

// Type returns the type to be associated with the name
func (definition *ResourceDefinition) Type() Type {
	return definition.spec // TODO BUG: the status is not considered here
	// TO FIX: consider lifting up the two methods used on the result of this method
	// (which are References/RequiredImports) into the TypeDefiner interface
}

// MarkAsStorageVersion marks the resource as the Kubebuilder storage version
func (definition *ResourceDefinition) MarkAsStorageVersion() *ResourceDefinition {
	result := *definition
	result.isStorageVersion = true
	return &result
}

// WithDescription replaces the description of the resource
func (definition *ResourceDefinition) WithDescription(description *string) TypeDefiner {
	result := *definition
	result.description = description
	return &result
}

// RequiredImports returns a list of packages required by this
func (definition *ResourceDefinition) RequiredImports() []*PackageReference {
	typeImports := definition.spec.RequiredImports()
	// TODO BUG: the status is not considered here
	typeImports = append(typeImports, MetaV1PackageReference)

	return typeImports
}

// AsDeclarations converts the ResourceDefinition to a go declaration
func (definition *ResourceDefinition) AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl {

	packageName, err := codeGenerationContext.GetImportedPackageName(MetaV1PackageReference)
	if err != nil {
		panic(fmt.Errorf("resource definition for %s failed to import package: %w", definition.typeName, err))
	}
	typeMetaField := defineField("", fmt.Sprintf("%s.TypeMeta", packageName), "`json:\",inline\"`")
	objectMetaField := defineField("", fmt.Sprintf("%s.ObjectMeta", packageName), "`json:\"metadata,omitempty\"`")

	/*
		start off with:
			metav1.TypeMeta   `json:",inline"`
			metav1.ObjectMeta `json:"metadata,omitempty"`

		then the Spec/Status fields
	*/
	fields := []*ast.Field{
		typeMetaField,
		objectMetaField,
		defineField("Spec", definition.spec.name, "`json:\"spec,omitempty\"`"),
	}

	if definition.status != nil {
		fields = append(fields, defineField("Status", definition.status.name, "`json:\"spec,omitempty\"`"))
	}

	resourceIdentifier := ast.NewIdent(definition.typeName.name)
	resourceTypeSpec := &ast.TypeSpec{
		Name: resourceIdentifier,
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: fields},
		},
	}

	comments :=
		[]*ast.Comment{
			{
				Text: "// +kubebuilder:object:root=true\n",
			},
		}

	if definition.isStorageVersion {
		comments = append(comments, &ast.Comment{
			Text: "// +kubebuilder:storageversion\n",
		})
	}

	if definition.description != nil {
		comments = append(comments, &ast.Comment{
			Text: "/*" + *definition.description + "*/",
		})
	}

	return []ast.Decl{
		&ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{resourceTypeSpec},
			Doc:   &ast.CommentGroup{List: comments},
		},
	}
}
