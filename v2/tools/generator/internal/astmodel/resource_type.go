/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/token"
	"sort"
	"strings"

	"github.com/dave/dst"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

type ResourceKind string

const (
	// ResourceKindNormal is a standard ARM resource.
	ResourceKindNormal = ResourceKind("normal")
	// ResourceKindExtension is an extension resource. Extension resources can have any resource as their parent.
	ResourceKindExtension = ResourceKind("extension")
)

// ResourceType represents a Kubernetes CRD resource which has both
// spec (the user-requested state) and status (the current state)
type ResourceType struct {
	spec                Type
	status              Type
	isStorageVersion    bool
	owner               *TypeName
	properties          PropertySet
	functions           map[string]Function
	testcases           map[string]TestCase
	annotations         []string // TODO: Consider ensuring that these are actually kubebuilder annotations.
	kind                ResourceKind
	armType             string
	apiVersionTypeName  TypeName
	apiVersionEnumValue EnumValue
	InterfaceImplementer
}

// NewResourceType defines a new resource type
func NewResourceType(specType Type, statusType Type) *ResourceType {
	result := &ResourceType{
		isStorageVersion:     false,
		owner:                nil,
		properties:           make(PropertySet),
		functions:            make(map[string]Function),
		testcases:            make(map[string]TestCase),
		kind:                 ResourceKindNormal,
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
func NewAzureResourceType(specType Type, statusType Type, typeName TypeName, kind ResourceKind) *ResourceType {
	if objectType, ok := specType.(*ObjectType); ok {
		// We have certain expectations about structure for resources
		var nameProperty *PropertyDefinition
		var typeProperty *PropertyDefinition
		var apiVersionProperty *PropertyDefinition

		isNameOptional := false
		isTypeOptional := false
		for _, property := range objectType.Properties() {
			// force this string because otherwise linter complains thinking it's an enum without an exhaustive switch...
			// It turns out there are other reasons to alias string than just to make an enum, seems like the linter doesn't
			// realize that
			switch string(property.PropertyName()) {
			case NameProperty:
				nameProperty = property
				if _, ok := property.PropertyType().(*OptionalType); ok {
					isNameOptional = true
				}
			case TypeProperty:
				typeProperty = property
				if _, ok := property.PropertyType().(*OptionalType); ok {
					isTypeOptional = true
				}
			case APIVersionProperty:
				apiVersionProperty = property
			}
		}

		if typeProperty == nil {
			panic(fmt.Sprintf("Resource %s is missing type property", typeName))
		}

		if nameProperty == nil {
			klog.V(1).Infof("resource %s is missing field 'Name', fabricating one...", typeName)

			nameProperty = NewPropertyDefinition("Name", "name", StringType)
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

	return NewResourceType(specType, statusType).WithKind(kind)
}

// Ensure ResourceType implements the Type interface correctly
var _ Type = &ResourceType{}

// Ensure ResourceType implements the PropertyContainer interface correctly
var _ PropertyContainer = &ResourceType{}

// Ensure ResourceType implements the FunctionContainer interface correctly
var _ FunctionContainer = &ResourceType{}

// Ensure ResourceType implements the TestCaseContainer interface correctly
var _ TestCaseContainer = &ResourceType{}

// SpecType returns the type used for specification
func (resource *ResourceType) SpecType() Type {
	return resource.spec
}

// StatusType returns the type used for current status
func (resource *ResourceType) StatusType() Type {
	return resource.status
}

// IsStorageVersion returns true if the resource is a storage version
func (resource *ResourceType) IsStorageVersion() bool {
	return resource.isStorageVersion
}

// WithSpec returns a new resource that has the specified spec type
func (resource *ResourceType) WithSpec(specType Type) *ResourceType {
	if TypeEquals(resource.spec, specType) {
		return resource // short-circuit
	}

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
	if TypeEquals(resource.status, statusType) {
		return resource // short-circuit
	}

	if specResource, ok := statusType.(*ResourceType); ok {
		// type is a resource, take its StatusType instead, so we don't nest resources
		return resource.WithStatus(specResource.StatusType())
	}

	result := resource.copy()
	result.status = statusType
	return result
}

func (resource *ResourceType) WithoutInterface(name TypeName) *ResourceType {
	if _, found := resource.InterfaceImplementer.FindInterface(name); !found {
		return resource
	}

	result := *resource
	result.InterfaceImplementer = result.InterfaceImplementer.WithoutInterface(name)
	return &result
}

// WithInterface creates a new Resource with a function (method) attached to it
func (resource *ResourceType) WithInterface(iface *InterfaceImplementation) *ResourceType {
	result := resource.copy()
	result.InterfaceImplementer = result.InterfaceImplementer.WithInterface(iface)
	return result
}

// WithFunction creates a new Resource with a function (method) attached to it
func (resource *ResourceType) WithFunction(function Function) *ResourceType {
	// Create a copy to preserve immutability
	result := resource.copy()
	result.functions[function.Name()] = function

	return result
}

// WithTestCase creates a new Resource that's a copy with an additional test case included
func (resource *ResourceType) WithTestCase(testcase TestCase) *ResourceType {
	result := resource.copy()
	result.testcases[testcase.Name()] = testcase
	return result
}

// WithKind returns a new ResourceType with the specified kind
func (resource *ResourceType) WithKind(kind ResourceKind) *ResourceType {
	result := resource.copy()
	result.kind = kind
	return result
}

// WithARMType returns a new ResourceType containing the type of the ARM resource, such as Microsoft.Batch/batchAccounts
func (resource *ResourceType) WithARMType(armType string) *ResourceType {
	result := resource.copy()
	result.armType = armType
	return result
}

// WithAPIVersion returns a new ResourceType with the specified API version (type and value).
func (resource *ResourceType) WithAPIVersion(apiVersionTypeName TypeName, apiVersionEnumValue EnumValue) *ResourceType {
	result := resource.copy()
	result.apiVersionTypeName = apiVersionTypeName
	result.apiVersionEnumValue = apiVersionEnumValue
	return result
}

// TestCases returns a new slice containing all the test cases associated with this resource
func (resource *ResourceType) TestCases() []TestCase {
	var result []TestCase
	for _, tc := range resource.testcases {
		result = append(result, tc)
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i].Name() < result[j].Name()
	})

	return result
}

// AsType always panics because a resource has no direct AST representation
func (resource *ResourceType) AsType(_ *CodeGenerationContext) dst.Expr {
	panic("a resource cannot be used directly as a type")
}

// AsZero always panics because a resource has no direct AST representation
func (resource *ResourceType) AsZero(_ Types, _ *CodeGenerationContext) dst.Expr {
	panic("a resource cannot be used directly as a type")
}

// Equals returns true if the other type is also a ResourceType and has Equal fields
func (resource *ResourceType) Equals(other Type, override EqualityOverrides) bool {
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
		len(resource.functions) != len(otherResource.functions) ||
		!TypeEquals(resource.spec, otherResource.spec, override) ||
		!TypeEquals(resource.status, otherResource.status, override) ||
		len(resource.annotations) != len(otherResource.annotations) ||
		resource.kind != otherResource.kind ||
		resource.armType != otherResource.armType ||
		!TypeEquals(resource.apiVersionTypeName, otherResource.apiVersionTypeName) ||
		resource.apiVersionEnumValue.Equals(&otherResource.apiVersionEnumValue) ||
		!resource.InterfaceImplementer.Equals(otherResource.InterfaceImplementer, override) {
		return false
	}

	// Check same functions present
	for name, fn := range otherResource.functions {
		ourFn, ok := resource.functions[name]
		if !ok {
			return false
		}

		if !ourFn.Equals(fn, override) {
			return false
		}
	}

	// Check same test cases present
	for name, testcase := range otherResource.testcases {
		ourCase, ok := resource.testcases[name]
		if !ok {
			// Didn't find the func, not equal
			return false
		}

		if !ourCase.Equals(testcase, override) {
			// Different testcase, even though same name; not-equal
			return false
		}
	}

	// Check same annotations present in the same order
	for i, ourAnnotation := range resource.annotations {
		otherAnnotation := otherResource.annotations[i]
		if ourAnnotation != otherAnnotation {
			return false
		}
	}

	return true
}

// EmbeddedProperties returns all the embedded properties for this resource type
// An ordered slice is returned to preserve immutability and provide determinism
func (resource *ResourceType) EmbeddedProperties() []*PropertyDefinition {
	typeMetaType := MakeTypeName(MetaV1PackageReference, "TypeMeta")
	typeMetaProperty := NewPropertyDefinition("", "", typeMetaType).
		WithTag("json", "inline")

	objectMetaType := MakeTypeName(MetaV1PackageReference, "ObjectMeta")
	objectMetaProperty := NewPropertyDefinition("", "metadata", objectMetaType).
		WithTag("json", "omitempty")

	return []*PropertyDefinition{
		typeMetaProperty,
		objectMetaProperty,
	}
}

// WithProperty creates a new ResourceType with another property attached to it
// Properties are unique by name, so this can be used to Add and Replace a property
func (resource *ResourceType) WithProperty(property *PropertyDefinition) *ResourceType {
	if property.HasName("Status") || property.HasName("Spec") {
		panic(fmt.Sprintf("May not modify property %q on a resource", property.PropertyName()))
	}

	// Create a copy to preserve immutability
	result := resource.copy()
	result.properties[property.propertyName] = property

	return result
}

// WithoutProperty creates a new ResourceType that's a copy without the specified property
func (resource *ResourceType) WithoutProperty(name PropertyName) *ResourceType {
	if name == "Status" || name == "Spec" {
		panic(fmt.Sprintf("May not remove property %q from a resource", name))
	}

	// Create a copy to preserve immutability
	result := resource.copy()
	delete(result.properties, name)

	return result
}

// Properties returns all the properties from this resource type
func (resource *ResourceType) Properties() PropertySet {
	result := resource.properties.Copy()

	result.Add(resource.createSpecProperty())
	if resource.status != nil {
		result.Add(resource.createStatusProperty())
	}

	return result
}

func (resource *ResourceType) createStatusProperty() *PropertyDefinition {
	statusProperty := NewPropertyDefinition("Status", "status", resource.status).
		WithTag("json", "omitempty")
	return statusProperty
}

func (resource *ResourceType) createSpecProperty() *PropertyDefinition {
	return NewPropertyDefinition("Spec", "spec", resource.spec).
		WithTag("json", "omitempty")
}

// Property returns the property and true if the named property is found, nil and false otherwise
func (resource *ResourceType) Property(name PropertyName) (*PropertyDefinition, bool) {
	if name == "Spec" {
		return resource.createSpecProperty(), true
	}

	if name == "Status" {
		return resource.createStatusProperty(), true
	}

	prop, ok := resource.properties[name]
	return prop, ok
}

// Kind returns the ResourceKind of the resource
func (resource *ResourceType) Kind() ResourceKind {
	return resource.kind
}

// ARMType returns the ARM Type of the resource. The ARM type is something like Microsoft.Batch/batchAccounts
func (resource *ResourceType) ARMType() string {
	return resource.armType
}

// TODO: It's possible we could do this at the package level?
// APIVersionTypeName returns the type name of the API version
func (resource *ResourceType) APIVersionTypeName() TypeName {
	return resource.apiVersionTypeName
}

// APIVersionEnumValue returns the enum value representing the ARM API version of the resource.
func (resource *ResourceType) APIVersionEnumValue() EnumValue {
	return resource.apiVersionEnumValue
}

// Functions returns all the function implementations
// A sorted slice is returned to preserve immutability and provide determinism
func (resource *ResourceType) Functions() []Function {
	var functions []Function
	for _, f := range resource.functions {
		functions = append(functions, f)
	}

	sort.Slice(functions, func(i int, j int) bool {
		return functions[i].Name() < functions[j].Name()
	})

	return functions
}

// HasFunctionWithName determines if this resource has a function with the given name
func (resource *ResourceType) HasFunctionWithName(name string) bool {
	_, ok := resource.functions[name]
	return ok
}

// References returns the types referenced by Status or Spec parts of the resource
func (resource *ResourceType) References() TypeNameSet {
	spec := resource.spec.References()

	var status TypeNameSet
	if resource.status != nil {
		status = resource.status.References()
	}

	result := SetUnion(spec, status)
	// It's a bit awkward to have to do this, but it doesn't exist as a reference
	// anywhere else
	if !resource.APIVersionTypeName().Equals(TypeName{}, EqualityOverrides{}) {
		result.Add(resource.APIVersionTypeName())
	}

	return result
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

// WithAnnotation adds the annotation to the resource and returns a copy of the resource
func (resource *ResourceType) WithAnnotation(annotation string) *ResourceType {
	result := resource.copy()
	result.annotations = append(result.annotations, annotation)
	return result
}

// RequiredPackageReferences returns a list of packages required by this
func (resource *ResourceType) RequiredPackageReferences() *PackageReferenceSet {
	references := NewPackageReferenceSet(MetaV1PackageReference)
	references.Merge(resource.spec.RequiredPackageReferences())

	if resource.status != nil {
		references.Merge(resource.status.RequiredPackageReferences())
	}

	for _, fn := range resource.functions {
		references.Merge(fn.RequiredPackageReferences())
	}

	// Interface imports
	references.Merge(resource.InterfaceImplementer.RequiredPackageReferences())

	return references
}

// AsDeclarations converts the resource type to a set of go declarations
func (resource *ResourceType) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	/*
		start off with:
			metav1.TypeMeta   `json:",inline"`
			metav1.ObjectMeta `json:"metadata,omitempty"`

		then the Spec/Status properties
	*/
	var fields []*dst.Field
	for _, property := range resource.EmbeddedProperties() {
		f := property.AsField(codeGenerationContext)
		if f != nil {
			fields = append(fields, f)
		}
	}

	for _, property := range resource.Properties().AsSlice() {
		f := property.AsField(codeGenerationContext)
		if f != nil {
			fields = append(fields, f)
		}
	}

	if len(fields) > 0 {
		// A Before:EmptyLine decoration on the first field looks odd, so we force it to Before:NewLine
		// This makes the output look nicer 🙂
		fields[0].Decs.Before = dst.NewLine
	}

	resourceTypeSpec := &dst.TypeSpec{
		Name: dst.NewIdent(declContext.Name.Name()),
		Type: &dst.StructType{
			Fields: &dst.FieldList{List: fields},
		},
	}

	var comments dst.Decorations

	// Add required RBAC annotations, only on storage version
	if resource.isStorageVersion {
		lpr, ok := declContext.Name.PackageReference.AsLocalPackage()
		if !ok {
			panic(fmt.Sprintf("expected resource package reference to be local: %q", declContext.Name))
		}
		group := strings.ToLower(lpr.Group() + GroupSuffix)
		resourceName := strings.ToLower(declContext.Name.Plural().Name())

		astbuilder.AddComment(&comments, fmt.Sprintf("// +kubebuilder:rbac:groups=%s,resources=%s,verbs=get;list;watch;create;update;patch;delete", group, resourceName))
		astbuilder.AddComment(&comments, fmt.Sprintf("// +kubebuilder:rbac:groups=%s,resources={%s/status,%s/finalizers},verbs=get;update;patch", group, resourceName, resourceName))

		// This newline is REQUIRED for controller-gen to realize these comments are here. Without it they are silently ignored, see:
		// https://github.com/kubernetes-sigs/controller-tools/issues/436
		comments = append(comments, "\n")
	}

	astbuilder.AddComment(&comments, "// +kubebuilder:object:root=true")
	if resource.status != nil {
		astbuilder.AddComment(&comments, "// +kubebuilder:subresource:status")
	}

	if resource.isStorageVersion {
		astbuilder.AddComment(&comments, "// +kubebuilder:storageversion")
	}

	// Add any custom kubebuilder annotations
	if len(resource.annotations) > 0 {
		astbuilder.AddComments(&comments, resource.annotations)
	}

	astbuilder.AddWrappedComments(&comments, declContext.Description, 200)
	AddValidationComments(&comments, declContext.Validations)

	resourceDeclaration := &dst.GenDecl{
		Tok:   token.TYPE,
		Specs: []dst.Spec{resourceTypeSpec},
		Decs: dst.GenDeclDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
				After:  dst.EmptyLine,
				Start:  comments,
			},
		},
	}

	var declarations []dst.Decl
	declarations = append(declarations, resourceDeclaration)
	declarations = append(declarations, resource.InterfaceImplementer.AsDeclarations(codeGenerationContext, declContext.Name, nil)...)
	declarations = append(declarations, resource.generateMethodDecls(codeGenerationContext, declContext.Name)...)
	declarations = append(declarations, resource.resourceListTypeDecls(codeGenerationContext, declContext.Name, declContext.Description)...)

	return declarations
}

func (resource *ResourceType) generateMethodDecls(codeGenerationContext *CodeGenerationContext, typeName TypeName) []dst.Decl {
	var result []dst.Decl

	for _, f := range resource.Functions() {
		funcDef := f.AsFunc(codeGenerationContext, typeName)
		result = append(result, funcDef)
	}

	return result
}

func (resource *ResourceType) makeResourceListTypeName(name TypeName) TypeName {
	return MakeTypeName(
		name.PackageReference,
		name.Name()+"List")
}

func (resource *ResourceType) resourceListTypeDecls(
	codeGenerationContext *CodeGenerationContext,
	resourceTypeName TypeName,
	description []string) []dst.Decl {

	typeName := resource.makeResourceListTypeName(resourceTypeName)

	packageName := codeGenerationContext.MustGetImportedPackageName(MetaV1PackageReference)

	typeMetaField := defineField("", dst.NewIdent(fmt.Sprintf("%s.TypeMeta", packageName)), "`json:\",inline\"`")
	objectMetaField := defineField("", dst.NewIdent(fmt.Sprintf("%s.ListMeta", packageName)), "`json:\"metadata,omitempty\"`")

	// We need an array of items
	items := NewArrayType(resourceTypeName)

	fields := []*dst.Field{
		typeMetaField,
		objectMetaField,
		defineField("Items", items.AsType(codeGenerationContext), "`json:\"items\"`"),
	}

	resourceTypeSpec := &dst.TypeSpec{
		Name: dst.NewIdent(typeName.Name()),
		Type: &dst.StructType{
			Fields: &dst.FieldList{List: fields},
		},
	}

	var comments dst.Decorations = []string{
		"// +kubebuilder:object:root=true\n",
	}

	astbuilder.AddWrappedComments(&comments, description, 200)

	return []dst.Decl{
		&dst.GenDecl{
			Tok:   token.TYPE,
			Specs: []dst.Spec{resourceTypeSpec},
			Decs:  dst.GenDeclDecorations{NodeDecs: dst.NodeDecs{Start: comments}},
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
		properties:           make(map[PropertyName]*PropertyDefinition),
		functions:            make(map[string]Function),
		testcases:            make(map[string]TestCase),
		annotations:          append([]string(nil), resource.annotations...),
		kind:                 resource.kind,
		armType:              resource.armType,
		apiVersionTypeName:   resource.apiVersionTypeName,
		apiVersionEnumValue:  resource.apiVersionEnumValue,
		InterfaceImplementer: resource.InterfaceImplementer.copy(),
	}

	for key, property := range resource.properties {
		result.properties[key] = property
	}

	for key, testcase := range resource.testcases {
		result.testcases[key] = testcase
	}

	for key, fn := range resource.functions {
		result.functions[key] = fn
	}

	return result
}

func (resource *ResourceType) HasTestCases() bool {
	return len(resource.testcases) > 0
}

// WriteDebugDescription adds a description of the current type to the passed builder
// builder receives the full description, including nested types
// types is a dictionary for resolving named types
func (resource *ResourceType) WriteDebugDescription(builder *strings.Builder, types Types) {
	if resource == nil {
		builder.WriteString("<nilResource>")
		return
	}

	builder.WriteString("Resource[spec:")
	resource.spec.WriteDebugDescription(builder, types)
	builder.WriteString("|status:")
	resource.status.WriteDebugDescription(builder, types)
	builder.WriteString("]")
}
