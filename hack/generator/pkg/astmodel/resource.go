/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/pkg/errors"
)

// ResourceType represents an ARM resource
type ResourceType struct {
	spec             Type
	status           Type
	isStorageVersion bool
}

// NewResourceType defines a new resource type
func NewResourceType(specType Type, statusType Type) *ResourceType {
	return &ResourceType{specType, statusType, false}
}

// assert that ResourceType implements TypeDefiner
var _ Type = &ResourceType{}

func (definition *ResourceType) AsType(_ *CodeGenerationContext) ast.Expr {
	panic("a resource cannot be used directly as a type")
}

func (definition *ResourceType) CreateInternalDefinitions(name *TypeName, idFactory IdentifierFactory) (Type, []TypeDefinition) {
	panic("a resource should never be nested")
}

func (definition *ResourceType) CreateDefinitions(name *TypeName, idFactory IdentifierFactory) (TypeDefinition, []TypeDefinition) {

	var others []TypeDefinition

	defineStruct := func(suffix string, theType Type) *TypeName {
		definedName := NewTypeName(name.PackageReference, name.Name()+suffix)
		defined, definedOthers := theType.CreateDefinitions(definedName, idFactory)
		others = append(append(others, defined), definedOthers...)
		return definedName
	}

	specName := defineStruct("Spec", definition.spec)

	var statusName *TypeName
	if definition.status != nil {
		statusName = defineStruct("Status", definition.status)
	}

	this := MakeTypeDefinition(name, &ResourceType{
		specName,
		statusName,
		definition.isStorageVersion,
	})

	return this, others
}

func (definition *ResourceType) Equals(other Type) bool {
	if definition == other {
		return true
	}

	if otherResource, ok := other.(*ResourceType); ok {
		return definition.spec.Equals(otherResource.spec) &&
			((definition.status == nil && otherResource.status == nil) ||
				definition.status.Equals(otherResource.status)) &&
			definition.isStorageVersion == otherResource.isStorageVersion
	}

	return false
}

// References returns the types referenced by Status or Spec parts of the resource
func (definition *ResourceType) References() TypeNameSet {
	spec := definition.spec.References()

	var status TypeNameSet
	if definition.status != nil {
		status = definition.status.References()
	}

	return SetUnion(spec, status)
}

// MarkAsStorageVersion marks the resource as the Kubebuilder storage version
func (definition *ResourceType) MarkAsStorageVersion() *ResourceType {
	result := *definition
	result.isStorageVersion = true
	return &result
}

// RequiredImports returns a list of packages required by this
func (definition *ResourceType) RequiredImports() []*PackageReference {
	typeImports := definition.spec.RequiredImports()
	// TODO BUG: the status is not considered here
	typeImports = append(typeImports, MetaV1PackageReference)

	return typeImports
}

// AsDeclarations converts the resource type to a set of go declarations
func (definition *ResourceType) AsDeclarations(codeGenerationContext *CodeGenerationContext, typeName *TypeName, description *string) []ast.Decl {

	packageName, err := codeGenerationContext.GetImportedPackageName(MetaV1PackageReference)
	if err != nil {
		panic(errors.Wrapf(err, "resource definition for %s failed to import package", typeName))
	}

	typeMetaField := defineField("", ast.NewIdent(fmt.Sprintf("%s.TypeMeta", packageName)), "`json:\",inline\"`")
	objectMetaField := defineField("", ast.NewIdent(fmt.Sprintf("%s.ObjectMeta", packageName)), "`json:\"metadata,omitempty\"`")

	/*
		start off with:
			metav1.TypeMeta   `json:",inline"`
			metav1.ObjectMeta `json:"metadata,omitempty"`

		then the Spec/Status properties
	*/
	fields := []*ast.Field{
		typeMetaField,
		objectMetaField,
		defineField("Spec", definition.spec.AsType(codeGenerationContext), "`json:\"spec,omitempty\"`"),
	}

	if definition.status != nil {
		fields = append(fields, defineField("Status", definition.status.AsType(codeGenerationContext), "`json:\"spec,omitempty\"`"))
	}

	resourceIdentifier := ast.NewIdent(typeName.Name())
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

	if description != nil {
		addDocComment(&comments, *description, 200)
	}

	return []ast.Decl{
		&ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{resourceTypeSpec},
			Doc:   &ast.CommentGroup{List: comments},
		},
	}
}
