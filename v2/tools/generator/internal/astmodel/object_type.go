/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"
	"sort"
	"strings"

	"github.com/dave/dst"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/readonly"
)

// ObjectType represents an (unnamed) object type
type ObjectType struct {
	embedded   map[TypeName]*PropertyDefinition
	properties ReadOnlyPropertySet
	functions  readonly.Map[string, Function]
	testcases  readonly.Map[string, TestCase]
	InterfaceImplementer
	isResource bool
	resources  TypeNameSet
}

// for want of a better place for this to live…
var AdditionalPropertiesPropertyName = PropertyName("AdditionalProperties")
var AdditionalPropertiesJsonName = "additionalProperties"

// EmptyObjectType is an empty object
var EmptyObjectType = NewObjectType()

// Ensure ObjectType implements the Type interface correctly
var _ Type = &ObjectType{}

// Ensure ObjectType implements the PropertyContainer interface correctly
var _ PropertyContainer = &ObjectType{}

// Ensure ObjectType implements the FunctionContainer interface correctly
var _ FunctionContainer = &ObjectType{}

// Ensure ObjectType implements the TestCaseContainer interface correctly
var _ TestCaseContainer = &ObjectType{}

// NewObjectType is a factory method for creating a new ObjectType
func NewObjectType() *ObjectType {
	return &ObjectType{
		embedded:             make(map[TypeName]*PropertyDefinition),
		properties:           make(PropertySet),
		functions:            readonly.EmptyMap[string, Function](),
		testcases:            readonly.EmptyMap[string, TestCase](),
		InterfaceImplementer: MakeInterfaceImplementer(),
		isResource:           false,
	}
}

func (objectType *ObjectType) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	declaration := &dst.GenDecl{
		Decs: dst.GenDeclDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
				After:  dst.EmptyLine,
			},
		},
		Tok: token.TYPE,
		Specs: []dst.Spec{
			&dst.TypeSpec{
				Name: dst.NewIdent(declContext.Name.Name()),
				Type: objectType.AsType(codeGenerationContext),
			},
		},
	}

	astbuilder.AddUnwrappedComments(&declaration.Decs.Start, declContext.Description)
	AddValidationComments(&declaration.Decs.Start, declContext.Validations)

	result := []dst.Decl{declaration}
	result = append(result, objectType.InterfaceImplementer.AsDeclarations(codeGenerationContext, declContext.Name, nil)...)
	result = append(result, objectType.generateMethodDecls(codeGenerationContext, declContext.Name)...)
	return result
}

func (objectType *ObjectType) generateMethodDecls(codeGenerationContext *CodeGenerationContext, typeName TypeName) []dst.Decl {
	funcs := objectType.Functions()
	result := make([]dst.Decl, 0, len(funcs))
	for _, f := range funcs {
		funcDef := generateMethodDeclForFunction(typeName, f, codeGenerationContext)
		result = append(result, funcDef)
	}

	return result
}

func defineField(fieldName string, fieldType dst.Expr, tag string) *dst.Field {
	result := &dst.Field{
		Type: fieldType,
		Tag:  astbuilder.TextLiteral(tag),
	}

	if fieldName != "" {
		result.Names = []*dst.Ident{dst.NewIdent(fieldName)}
	}

	return result
}

// Properties returns all the property definitions
func (objectType *ObjectType) Properties() ReadOnlyPropertySet {
	return objectType.properties
}

// Property returns the details of a specific property based on its unique case-sensitive name
func (objectType *ObjectType) Property(name PropertyName) (*PropertyDefinition, bool) {
	return objectType.properties.Get(name)
}

// EmbeddedProperties returns all the embedded properties
// A sorted slice is returned to preserve immutability and provide determinism
func (objectType *ObjectType) EmbeddedProperties() []*PropertyDefinition {
	result := make([]*PropertyDefinition, 0, len(objectType.embedded))
	for _, embedded := range objectType.embedded {
		result = append(result, embedded)
	}

	sort.Slice(result, func(left int, right int) bool {
		lTypeName, err := extractEmbeddedTypeName(result[left].PropertyType())
		if err != nil {
			// It shouldn't be possible to get an invalid embedded type here, as we should have
			// failed to add it.
			panic(err)
		}
		rTypeName, err := extractEmbeddedTypeName(result[right].PropertyType())
		if err != nil {
			// It shouldn't be possible to get an invalid embedded type here, as we should have
			// failed to add it.
			panic(err)
		}
		return lTypeName.Name() < rTypeName.Name()
	})

	return result
}

func (objectType *ObjectType) IsResource() bool {
	return objectType.isResource
}

func (objectType *ObjectType) WithIsResource(isResource bool) *ObjectType {
	if objectType.isResource == isResource {
		return objectType
	}

	result := objectType.copy()
	result.isResource = isResource
	return result
}

func (objectType *ObjectType) Resources() TypeNameSet {
	return objectType.resources
}

func (objectType *ObjectType) WithResource(resource TypeName) *ObjectType {
	if objectType.resources.Contains(resource) {
		return objectType
	}

	result := objectType.copy()
	result.resources.Add(resource)
	return result
}

func (objectType *ObjectType) WithResources(resources TypeNameSet) *ObjectType {
	if objectType.resources.ContainsAll(resources) {
		return objectType
	}

	result := objectType.copy()
	result.resources.AddAll(resources)
	return result
}

func (objectType *ObjectType) ClearResources() *ObjectType {
	if len(objectType.resources) == 0 {
		return objectType
	}

	result := objectType.copy()
	result.resources = make(TypeNameSet)
	return result
}

// Functions returns all the function implementations
// A sorted slice is returned to preserve immutability and provide determinism
func (objectType *ObjectType) Functions() []Function {
	functions := objectType.functions.Values()

	sort.Slice(functions, func(i int, j int) bool {
		return functions[i].Name() < functions[j].Name()
	})

	return functions
}

// HasFunctionWithName determines if this object has a function with the given name
func (objectType *ObjectType) HasFunctionWithName(name string) bool {
	return objectType.functions.ContainsKey(name)
}

// AsType implements Type for ObjectType
func (objectType *ObjectType) AsType(codeGenerationContext *CodeGenerationContext) dst.Expr {
	embedded := objectType.EmbeddedProperties()
	fields := make([]*dst.Field, 0, len(embedded))
	for _, f := range embedded {
		fields = append(fields, f.AsField(codeGenerationContext))
	}

	for _, f := range objectType.properties.AsSlice() {
		fields = append(fields, f.AsField(codeGenerationContext))
	}

	if len(fields) > 0 {
		// A Before:EmptyLine decoration on the first field looks odd, so we force it to Before:NewLine
		// This makes the output look nicer 🙂
		fields[0].Decs.Before = dst.NewLine
	}

	return &dst.StructType{
		Fields: &dst.FieldList{
			List: fields,
		},
	}
}

// AsZero renders an expression for the "zero" value of the type
// We can only generate a zero value for a named object type (and that's handled by TypeName, so
// we'll only end up here if it is an anonymous type.)
func (objectType *ObjectType) AsZero(_ TypeDefinitionSet, _ *CodeGenerationContext) dst.Expr {
	panic("cannot create a zero value for an object type")
}

// RequiredPackageReferences returns a list of packages required by this
func (objectType *ObjectType) RequiredPackageReferences() *PackageReferenceSet {
	result := NewPackageReferenceSet()

	for _, property := range objectType.embedded {
		propertyType := property.PropertyType()
		result.Merge(propertyType.RequiredPackageReferences())
	}

	objectType.properties.ForEach(func(property *PropertyDefinition) {
		propertyType := property.PropertyType()
		result.Merge(propertyType.RequiredPackageReferences())
	})

	objectType.functions.ForEach(func(_ string, function Function) {
		result.Merge(function.RequiredPackageReferences())
	})

	result.Merge(objectType.InterfaceImplementer.RequiredPackageReferences())

	return result
}

// References returns the set of all the types referred to by any property.
func (objectType *ObjectType) References() TypeNameSet {
	results := NewTypeNameSet()
	objectType.properties.ForEach(func(property *PropertyDefinition) {
		results.AddAll(property.PropertyType().References())
	})

	for _, property := range objectType.embedded {
		results.AddAll(property.PropertyType().References())
	}

	objectType.functions.ForEach(func(_ string, fn Function) {
		results.AddAll(fn.References())
	})

	return results
}

// Equals returns true if the passed type is an object type with the same properties, false otherwise
// The order of the properties is not relevant
func (objectType *ObjectType) Equals(t Type, overrides EqualityOverrides) bool {
	if objectType == t {
		return true // short circuit
	}

	other, ok := t.(*ObjectType)
	if !ok {
		return false
	}

	if overrides.ObjectType != nil {
		return overrides.ObjectType(objectType, other)
	}

	if len(objectType.embedded) != len(other.embedded) {
		// Different number of embedded properties, not equal
		return false
	}

	for n, f := range other.embedded {
		ourProperty, ok := objectType.embedded[n]
		if !ok {
			// Didn't find the property, not equal
			return false
		}

		if !ourProperty.Equals(f, overrides) {
			// Different property, even though same name; not-equal
			return false
		}
	}

	if !objectType.properties.Equals(other.properties, overrides) {
		return false
	}

	if !objectType.functions.Equals(other.functions, func(l, r Function) bool {
		return l.Equals(r, overrides)
	}) {
		return false
	}

	if !objectType.testcases.Equals(other.testcases, func(l, r TestCase) bool {
		return l.Equals(r, overrides)
	}) {
		return false
	}

	if objectType.isResource != other.isResource {
		return false
	}

	if !objectType.resources.Equals(other.resources) {
		return false
	}

	return objectType.InterfaceImplementer.Equals(other.InterfaceImplementer, overrides)
}

// WithProperty creates a new ObjectType with another property attached to it
// Properties are unique by name, so this can be used to Add and to Replace a property
func (objectType *ObjectType) WithProperty(property *PropertyDefinition) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()

	props := result.properties.Copy()
	props[property.propertyName] = property
	result.properties = props

	return result
}

// WithProperties creates a new ObjectType that's a copy with additional properties included
// Properties are unique by name, so this can be used to both Add and Replace properties.
func (objectType *ObjectType) WithProperties(properties ...*PropertyDefinition) *ObjectType {
	if len(properties) == 0 {
		// short-circuit
		return objectType
	}

	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	props := result.properties.Copy()

	for _, f := range properties {
		props[f.propertyName] = f
	}

	result.properties = props

	return result
}

// WithEmbeddedProperties creates a new ObjectType with the additional embedded properties included
func (objectType *ObjectType) WithEmbeddedProperties(properties ...*PropertyDefinition) (*ObjectType, error) {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()

	for _, p := range properties {
		err := objectType.checkEmbeddedProperty(p)
		if err != nil {
			return objectType, err
		}

		typeName, err := extractEmbeddedTypeName(p.PropertyType())
		if err != nil {
			return objectType, err
		}

		result.embedded[typeName] = p
	}

	return result, nil
}

// WithoutProperties creates a new ObjectType from this one but
// without any properties.
func (objectType *ObjectType) WithoutProperties() *ObjectType {
	result := objectType.copy()
	result.properties = PropertySet(make(map[PropertyName]*PropertyDefinition))
	return result
}

// WithoutSpecificProperties creates a new ObjectType from this one but
// without any properties.
func (objectType *ObjectType) WithoutSpecificProperties(props ...PropertyName) *ObjectType {
	if len(props) == 0 {
		return objectType
	}

	result := objectType.copy()

	resultProps := result.properties.Copy()
	for _, name := range props {
		delete(resultProps, name)
	}
	result.properties = resultProps

	return result
}

// WithoutProperty creates a new ObjectType that's a copy without the specified property
func (objectType *ObjectType) WithoutProperty(name PropertyName) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()

	resultProps := result.properties.Copy()
	delete(resultProps, name)
	result.properties = resultProps

	return result
}

// WithoutEmbeddedProperty creates a new ObjectType that's a copy without the specified typeName embedded
func (objectType *ObjectType) WithoutEmbeddedProperty(name TypeName) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	delete(result.embedded, name)

	return result
}

// WithoutEmbeddedProperties creates a new ObjectType from this one but
// without any embedded properties.
func (objectType *ObjectType) WithoutEmbeddedProperties() *ObjectType {
	result := objectType.copy()
	result.embedded = make(map[TypeName]*PropertyDefinition)
	return result
}

func (objectType *ObjectType) checkEmbeddedProperty(property *PropertyDefinition) error {
	// There are certain expectations about the shape of the provided property -- namely
	// it must not have a name and its Type must be TypeName or Optional[TypeName]
	if property.PropertyName() != "" {
		return errors.Errorf("embedded property name must be empty, was: %s", property.PropertyName())
	}

	return nil
}

// WithEmbeddedProperty creates a new ObjectType with another embedded property
func (objectType *ObjectType) WithEmbeddedProperty(property *PropertyDefinition) (*ObjectType, error) {
	err := objectType.checkEmbeddedProperty(property)
	if err != nil {
		return objectType, err
	}

	typeName, err := extractEmbeddedTypeName(property.PropertyType())
	if err != nil {
		return objectType, err
	}

	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.embedded[typeName] = property

	return result, nil
}

// WithFunction creates a new ObjectType with a function (method) attached to it
func (objectType *ObjectType) WithFunction(function Function) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.functions = result.functions.With(function.Name(), function)

	return result
}

// WithoutFunctions creates a new ObjectType with no functions (useful for testing)
func (objectType *ObjectType) WithoutFunctions() *ObjectType {
	// Create a copy to preserve immutability
	result := objectType.copy()
	result.functions = readonly.EmptyMap[string, Function]()

	return result
}

// WithInterface creates a new ObjectType that's a copy with an interface implementation attached
func (objectType *ObjectType) WithInterface(iface *InterfaceImplementation) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.InterfaceImplementer = result.InterfaceImplementer.WithInterface(iface)
	return result
}

// WithoutInterface removes the specified interface
func (objectType *ObjectType) WithoutInterface(name TypeName) *ObjectType {
	if _, found := objectType.InterfaceImplementer.FindInterface(name); !found {
		return objectType
	}

	result := objectType.copy()
	result.InterfaceImplementer = result.InterfaceImplementer.WithoutInterface(name)
	return result
}

// WithTestCase creates a new ObjectType that's a copy with an additional test case included
func (objectType *ObjectType) WithTestCase(testcase TestCase) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.testcases = result.testcases.With(testcase.Name(), testcase)
	return result
}

func (objectType *ObjectType) copy() *ObjectType {
	result := &ObjectType{
		// no need to clone these, they are all readonly
		properties: objectType.properties,
		functions:  objectType.functions,
		testcases:  objectType.testcases,
	}

	result.embedded = make(map[TypeName]*PropertyDefinition, len(objectType.embedded))
	for k, v := range objectType.embedded {
		result.embedded[k] = v
	}

	result.InterfaceImplementer = objectType.InterfaceImplementer.copy()
	result.isResource = objectType.isResource
	result.resources = objectType.resources.Copy()

	return result
}

// String implements fmt.Stringer
func (objectType *ObjectType) String() string {
	return "(object)"
}

func (objectType *ObjectType) HasTestCases() bool {
	return !objectType.testcases.IsEmpty()
}

func (objectType *ObjectType) TestCases() []TestCase {
	result := objectType.testcases.Values()

	sort.Slice(result, func(i int, j int) bool {
		return result[i].Name() < result[j].Name()
	})

	return result
}

// IsObjectDefinition returns true if the passed definition is for an ARM type; false otherwise.
func IsObjectDefinition(definition TypeDefinition) bool {
	_, ok := AsObjectType(definition.theType)
	return ok
}

func extractEmbeddedTypeName(t Type) (TypeName, error) {
	typeName, isTypeName := AsTypeName(t)
	if isTypeName {
		return typeName, nil
	}

	return nil, errors.Errorf("embedded property type must be TypeName, was: %T", t)
}

// WriteDebugDescription adds a description of the current type to the passed builder.
// builder receives the full description, including nested types.
func (objectType *ObjectType) WriteDebugDescription(builder *strings.Builder, _ PackageReference) {
	if objectType == nil {
		builder.WriteString("<nilObject>")
	} else {
		builder.WriteString("Object")
	}
}

// FindPropertyWithTagValue finds the property with the given tag and value if it exists. The boolean return
// is false if no match can be found.
func (objectType *ObjectType) FindPropertyWithTagValue(tag string, value string) (*PropertyDefinition, bool) {
	return objectType.properties.Find(func(prop *PropertyDefinition) bool {
		values, ok := prop.Tag(tag)
		return ok && slices.Contains(values, value)
	})
}

// FindAllPropertiesWithTagValue finds all the properties with the given tag and value if they exist.
func (objectType *ObjectType) FindAllPropertiesWithTagValue(tag string, value string) []*PropertyDefinition {
	return objectType.properties.FindAll(func(prop *PropertyDefinition) bool {
		values, ok := prop.Tag(tag)
		return ok && slices.Contains(values, value)
	})
}
