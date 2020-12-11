/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/token"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	ast "github.com/dave/dst"
	"k8s.io/klog/v2"
)

// ResourceType represents a Kubernetes CRD resource which has both
// spec (the user-requested state) and status (the current state)
type ResourceType struct {
	spec             Type
	status           Type
	isStorageVersion bool
	owner            *TypeName
	testcases        map[string]TestCase
	InterfaceImplementer
}

// NewResourceType defines a new resource type
func NewResourceType(specType Type, statusType Type) *ResourceType {
	result := &ResourceType{
		isStorageVersion:     false,
		owner:                nil,
		testcases:            make(map[string]TestCase),
		InterfaceImplementer: MakeInterfaceImplementer(),
	}

	return result.WithSpec(specType).WithStatus(statusType)
}

func IsResourceType(t Type) bool {
	_, ok := t.(*ResourceType)
	return ok
}

func IsResourceDefinition(def TypeDefinition) bool {
	return IsResourceType(def.Type())
}

// NewAzureResourceType defines a new resource type for Azure. It ensures that
// the resource has certain expected properties such as type and name.
// The typeName parameter is just used for logging.
func NewAzureResourceType(specType Type, statusType Type, typeName TypeName) *ResourceType {
	if objectType, ok := specType.(*ObjectType); ok {
		// We have certain expectations about structure for resources
		var nameProperty *PropertyDefinition
		var typeProperty *PropertyDefinition
		var apiVersionProperty *PropertyDefinition
		isNameOptional := false
		isTypeOptional := false
		for _, property := range objectType.Properties() {
			if property.HasName(NameProperty) {
				nameProperty = property
				if _, ok := property.PropertyType().(*OptionalType); ok {
					isNameOptional = true
				}
			}

			if property.HasName(TypeProperty) {
				typeProperty = property
				if _, ok := property.PropertyType().(*OptionalType); ok {
					isTypeOptional = true
				}
			}

			if property.HasName(ApiVersionProperty) {
				apiVersionProperty = property
			}
		}

		if typeProperty == nil {
			panic(fmt.Sprintf("Resource %s is missing type property", typeName))
		}

		if nameProperty == nil {
			klog.V(1).Infof("resource %s is missing field 'Name', fabricating one...", typeName)

			nameProperty = NewPropertyDefinition(PropertyName("Name"), "name", StringType)
			nameProperty.WithDescription("The name of the resource")
			isNameOptional = true
		}

		if apiVersionProperty == nil {
			panic(fmt.Sprintf("Resource %s is missing apiVersion property", typeName))
		}

		if isNameOptional {
			// Fix name to be required -- again this is an artifact of bad spec more than anything
			nameProperty = nameProperty.MakeRequired()
			objectType = objectType.WithProperty(nameProperty)
		}

		// Fix APIVersion to be required. Technically this isn't due to a bad specification, but in our
		// case forcing it to required makes our lives simpler (and the vast majority of resources specify
		// it as required anyway). The only time it's allowed to be optional is if you set apiProfile on
		// the ARM template instead, which we never do.
		apiVersionProperty = apiVersionProperty.MakeRequired()
		objectType = objectType.WithProperty(apiVersionProperty)

		if isTypeOptional {
			typeProperty = typeProperty.MakeRequired()
			objectType = objectType.WithProperty(typeProperty)
		}

		specType = objectType
	}

	return NewResourceType(specType, statusType)
}

// assert that ResourceType implements Type
var _ Type = &ResourceType{}

// SpecType returns the type used for specification
func (resource *ResourceType) SpecType() Type {
	return resource.spec
}

// StatusType returns the type used for current status
func (resource *ResourceType) StatusType() Type {
	return resource.status
}

// WithSpec returns a new resource that has the specified spec type
func (resource *ResourceType) WithSpec(specType Type) *ResourceType {

	if specResource, ok := specType.(*ResourceType); ok {
		// type is a resource, take its SpecType instead
		// so we don't nest resources
		return resource.WithSpec(specResource.SpecType())
	}

	result := resource.copy()
	result.spec = specType
	return result
}

// WithStatus returns a new resource that has the specified status type
func (resource *ResourceType) WithStatus(statusType Type) *ResourceType {

	if specResource, ok := statusType.(*ResourceType); ok {
		// type is a resource, take its StatusType instead
		// so we don't nest resources
		return resource.WithStatus(specResource.StatusType())
	}

	result := resource.copy()
	result.status = statusType
	return result
}

// WithInterface creates a new Resource with a function (method) attached to it
func (resource *ResourceType) WithInterface(iface *InterfaceImplementation) *ResourceType {
	result := resource.copy()
	result.InterfaceImplementer = result.InterfaceImplementer.WithInterface(iface)
	return result
}

// WithTestCase creates a new Resource that's a copy with an additional test case included
func (resource *ResourceType) WithTestCase(testcase TestCase) *ResourceType {
	result := resource.copy()
	result.testcases[testcase.Name()] = testcase
	return result
}

// AsType converts the ResourceType to go AST Expr
func (resource *ResourceType) AsType(_ *CodeGenerationContext) ast.Expr {
	panic("a resource cannot be used directly as a type")
}

// Equals returns true if the other type is also a ResourceType and has Equal fields
func (resource *ResourceType) Equals(other Type) bool {
	if resource == other {
		// Same reference
		return true
	}

	otherResource, ok := other.(*ResourceType)
	if !ok {
		return false
	}

	// Do cheap tests earlier
	if resource.isStorageVersion != otherResource.isStorageVersion ||
		len(resource.testcases) != len(otherResource.testcases) ||
		!TypeEquals(resource.spec, otherResource.spec) ||
		!TypeEquals(resource.status, otherResource.status) ||
		!resource.InterfaceImplementer.Equals(otherResource.InterfaceImplementer) {
		return false
	}

	for name, testcase := range otherResource.testcases {
		ourCase, ok := resource.testcases[name]
		if !ok {
			// Didn't find the func, not equal
			return false
		}

		if !ourCase.Equals(testcase) {
			// Different testcase, even though same name; not-equal
			return false
		}
	}

	return true
}

// References returns the types referenced by Status or Spec parts of the resource
func (resource *ResourceType) References() TypeNameSet {
	spec := resource.spec.References()

	var status TypeNameSet
	if resource.status != nil {
		status = resource.status.References()
	}

	return SetUnion(spec, status)
}

// Owner returns the name of the owner type
func (resource *ResourceType) Owner() *TypeName {
	return resource.owner
}

// MarkAsStorageVersion marks the resource as the Kubebuilder storage version
func (resource *ResourceType) MarkAsStorageVersion() *ResourceType {
	result := resource.copy()
	result.isStorageVersion = true
	return result
}

// WithOwner updates the owner of the resource and returns a copy of the resource
func (resource *ResourceType) WithOwner(owner *TypeName) *ResourceType {
	result := resource.copy()
	result.owner = owner
	return result
}

// RequiredPackageReferences returns a list of packages required by this
func (resource *ResourceType) RequiredPackageReferences() *PackageReferenceSet {
	references := NewPackageReferenceSet(MetaV1PackageReference)
	references.Merge(resource.spec.RequiredPackageReferences())

	if resource.status != nil {
		references.Merge(resource.status.RequiredPackageReferences())
	}

	// Interface imports
	references.Merge(resource.InterfaceImplementer.RequiredPackageReferences())

	return references
}

// AsDeclarations converts the resource type to a set of go declarations
func (resource *ResourceType) AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description []string) []ast.Decl {

	packageName := codeGenerationContext.MustGetImportedPackageName(MetaV1PackageReference)

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
		defineField("Spec", resource.spec.AsType(codeGenerationContext), "`json:\"spec,omitempty\"`"),
	}

	if resource.status != nil {
		fields = append(fields, defineField("Status", resource.status.AsType(codeGenerationContext), "`json:\"status,omitempty\"`"))
	}

	resourceTypeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(name.Name()),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: fields},
		},
	}

	var comments ast.Decorations

	astbuilder.AddComment(&comments, "// +kubebuilder:object:root=true")
	if resource.status != nil {
		astbuilder.AddComment(&comments, "// +kubebuilder:subresource:status")
	}

	if resource.isStorageVersion {
		astbuilder.AddComment(&comments, "// +kubebuilder:storageversion")
	}

	astbuilder.AddWrappedComments(&comments, description, 200)

	resourceDeclaration := &ast.GenDecl{
		Tok:   token.TYPE,
		Specs: []ast.Spec{resourceTypeSpec},
		Decs: ast.GenDeclDecorations{
			NodeDecs: ast.NodeDecs{
				Before: ast.EmptyLine,
				After:  ast.EmptyLine,
				Start:  comments,
			},
		},
	}

	var declarations []ast.Decl
	declarations = append(declarations, resourceDeclaration)
	declarations = append(declarations, resource.InterfaceImplementer.AsDeclarations(codeGenerationContext, name, nil)...)
	declarations = append(declarations, resource.resourceListTypeDecls(codeGenerationContext, name, description)...)

	return declarations
}

func (resource *ResourceType) makeResourceListTypeName(name TypeName) TypeName {
	return MakeTypeName(
		name.PackageReference,
		name.Name()+"List")
}

func (resource *ResourceType) resourceListTypeDecls(
	codeGenerationContext *CodeGenerationContext,
	resourceTypeName TypeName,
	description []string) []ast.Decl {

	typeName := resource.makeResourceListTypeName(resourceTypeName)

	packageName := codeGenerationContext.MustGetImportedPackageName(MetaV1PackageReference)

	typeMetaField := defineField("", ast.NewIdent(fmt.Sprintf("%s.TypeMeta", packageName)), "`json:\",inline\"`")
	objectMetaField := defineField("", ast.NewIdent(fmt.Sprintf("%s.ListMeta", packageName)), "`json:\"metadata,omitempty\"`")

	// We need an array of items
	items := NewArrayType(resourceTypeName)

	fields := []*ast.Field{
		typeMetaField,
		objectMetaField,
		defineField("Items", items.AsType(codeGenerationContext), "`json:\"items\"`"),
	}

	resourceTypeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(typeName.Name()),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: fields},
		},
	}

	var comments ast.Decorations = []string{
		"// +kubebuilder:object:root=true\n",
	}

	astbuilder.AddWrappedComments(&comments, description, 200)

	return []ast.Decl{
		&ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{resourceTypeSpec},
			Decs:  ast.GenDeclDecorations{NodeDecs: ast.NodeDecs{Start: comments}},
		},
	}
}

// SchemeTypes returns the types represented by this resource which must be registered
// with the controller Scheme
func (resource *ResourceType) SchemeTypes(name TypeName) []TypeName {
	return []TypeName{
		name,
		resource.makeResourceListTypeName(name),
	}
}

// String implements fmt.Stringer
func (*ResourceType) String() string {
	return "(resource)"
}

func (resource *ResourceType) copy() *ResourceType {
	result := &ResourceType{
		spec:                 resource.spec,
		status:               resource.status,
		isStorageVersion:     resource.isStorageVersion,
		owner:                resource.owner,
		testcases:            make(map[string]TestCase),
		InterfaceImplementer: resource.InterfaceImplementer.copy(),
	}

	for key, value := range resource.testcases {
		result.testcases[key] = value
	}

	return result
}

func (resource *ResourceType) HasTestCases() bool {
	return len(resource.testcases) > 0
}
