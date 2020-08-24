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

// ResourceListType represents a Kubernetes CRD resource list
type ResourceListType struct {
	resource TypeName
}

// NewResourceListType defines a new resource list type
func NewResourceListType(resource TypeName) *ResourceListType {
	return &ResourceListType{
		resource: resource,
	}
}

// assert that ResourceType implements Type
var _ Type = &ResourceListType{}

// SpecType returns the type used for specification
func (list *ResourceListType) Resource() TypeName {
	return list.resource
}

// String implements fmt.Stringer
func (list *ResourceListType) String() string {
	return fmt.Sprintf("(ResourceList<%s>)", list.resource)
}

// AsType converts the ResourceType to go AST Expr
func (list *ResourceListType) AsType(_ *CodeGenerationContext) ast.Expr {
	panic("a resource list cannot be used directly as a type")
}

// Equals returns true if the other type is also a ResourceType and has Equal fields
func (list *ResourceListType) Equals(other Type) bool {
	if list == other {
		return true
	}

	if otherList, ok := other.(*ResourceListType); ok {
		return TypeEquals(list.resource, otherList.resource)
	}

	return false
}

// References returns the types referenced by Status or Spec parts of the resource
func (list *ResourceListType) References() TypeNameSet {
	return list.resource.References()
}

// RequiredImports returns a list of packages required by this
func (list *ResourceListType) RequiredImports() []PackageReference {
	imports := list.resource.RequiredImports()
	imports = append(imports, MetaV1PackageReference)
	return imports
}

// AsDeclarations converts the resource type to a set of go declarations
func (list *ResourceListType) AsDeclarations(codeGenerationContext *CodeGenerationContext, typeName TypeName, description []string) []ast.Decl {

	packageName, err := codeGenerationContext.GetImportedPackageName(MetaV1PackageReference)
	if err != nil {
		panic(errors.Wrapf(err, "resource list definition for %s failed to import package", typeName))
	}

	typeMetaField := defineField("", ast.NewIdent(fmt.Sprintf("%s.TypeMeta", packageName)), "`json:\",inline\"`")
	objectMetaField := defineField("", ast.NewIdent(fmt.Sprintf("%s.ListMeta", packageName)), "`json:\"metadata,omitempty\"`")

	// We need an array of items
	items := NewArrayType(list.resource)

	fields := []*ast.Field{
		typeMetaField,
		objectMetaField,
		defineField("Items", items.AsType(codeGenerationContext), "`json:\"items\"`"),
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

	addDocComments(&comments, description, 200)

	return []ast.Decl{
		&ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{resourceTypeSpec},
			Doc:   &ast.CommentGroup{List: comments},
		},
	}
}
