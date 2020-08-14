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

// ResourceType represents a Kubernetes CRD resource which has both
// spec (the user-requested state) and status (the current state)
type ResourceType struct {
	spec             Type
	status           Type
	isStorageVersion bool
}

// NewResourceType defines a new resource type
func NewResourceType(specType Type, statusType Type) *ResourceType {
	return &ResourceType{specType, statusType, false}
}

// assert that ResourceType implements Type
var _ Type = &ResourceType{}

// SpecType returns the type used for specificiation
func (definition *ResourceType) SpecType() Type {
	return definition.spec
}

// StatusType returns the type used for current status
func (definition *ResourceType) StatusType() Type {
	return definition.status
}

// AsType converts the ResourceType to go AST Expr
func (definition *ResourceType) AsType(_ *CodeGenerationContext) ast.Expr {
	panic("a resource cannot be used directly as a type")
}

// Equals returns true if the other type is also a ResourceType and has Equal fields
func (definition *ResourceType) Equals(other Type) bool {
	if definition == other {
		return true
	}

	if otherResource, ok := other.(*ResourceType); ok {
		return TypeEquals(definition.spec, otherResource.spec) &&
			TypeEquals(definition.status, otherResource.status) &&
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
func (definition *ResourceType) RequiredImports() []PackageReference {
	typeImports := definition.spec.RequiredImports()

	if definition.status != nil {
		typeImports = append(typeImports, definition.status.RequiredImports()...)
	}

	typeImports = append(typeImports, MetaV1PackageReference)

	return typeImports
}

// AsDeclarations converts the resource type to a set of go declarations
func (definition *ResourceType) AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description []string) []ast.Decl {

	packageName, err := codeGenerationContext.GetImportedPackageName(MetaV1PackageReference)
	if err != nil {
		panic(errors.Wrapf(err, "resource definition for %s failed to import package", name))
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

	resourceIdentifier := ast.NewIdent(name.Name())
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

	addDocComments(&comments, description, 200)

	return []ast.Decl{
		&ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{resourceTypeSpec},
			Doc:   &ast.CommentGroup{List: comments},
		},
	}
}
