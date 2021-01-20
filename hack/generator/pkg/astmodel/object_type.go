/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"
	"sort"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/dave/dst"
	"github.com/pkg/errors"
)

// ObjectType represents an (unnamed) object type
type ObjectType struct {
	embedded   map[TypeName]*PropertyDefinition
	properties map[PropertyName]*PropertyDefinition
	functions  map[string]Function
	testcases  map[string]TestCase
	InterfaceImplementer
}

// EmptyObjectType is an empty object
var EmptyObjectType = NewObjectType()

// Ensure ObjectType implements the Type interface correctly
var _ Type = &ObjectType{}

// NewObjectType is a factory method for creating a new ObjectType
func NewObjectType() *ObjectType {
	return &ObjectType{
		embedded:             make(map[TypeName]*PropertyDefinition),
		properties:           make(map[PropertyName]*PropertyDefinition),
		functions:            make(map[string]Function),
		testcases:            make(map[string]TestCase),
		InterfaceImplementer: MakeInterfaceImplementer(),
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

	astbuilder.AddWrappedComments(&declaration.Decs.Start, declContext.Description, 200)
	AddValidationComments(&declaration.Decs.Start, declContext.Validations)

	result := []dst.Decl{declaration}
	result = append(result, objectType.InterfaceImplementer.AsDeclarations(codeGenerationContext, declContext.Name, nil)...)
	result = append(result, objectType.generateMethodDecls(codeGenerationContext, declContext.Name)...)
	return result
}

func (objectType *ObjectType) generateMethodDecls(codeGenerationContext *CodeGenerationContext, typeName TypeName) []dst.Decl {
	var result []dst.Decl

	// Functions must be ordered by name for deterministic output
	var functions []Function
	for _, f := range objectType.functions {
		functions = append(functions, f)
	}

	sort.Slice(functions, func(i int, j int) bool {
		return functions[i].Name() < functions[j].Name()
	})

	for _, f := range functions {
		funcDef := f.AsFunc(codeGenerationContext, typeName)
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
// A sorted slice is returned to preserve immutability and provide determinism
func (objectType *ObjectType) Properties() []*PropertyDefinition {
	var result []*PropertyDefinition
	for _, property := range objectType.properties {
		result = append(result, property)
	}

	// Sorted so that it's always consistent
	sort.Slice(result, func(left int, right int) bool {
		return result[left].propertyName < result[right].propertyName
	})

	return result
}

// Property returns the details of a specific property based on its unique case sensitive name
func (objectType *ObjectType) Property(name PropertyName) (*PropertyDefinition, bool) {
	prop, ok := objectType.properties[name]
	return prop, ok
}

// EmbeddedProperties returns all the embedded properties
// A sorted slice is returned to preserve immutability and provide determinism
func (objectType *ObjectType) EmbeddedProperties() []*PropertyDefinition {
	var result []*PropertyDefinition
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

// HasFunctionWithName determines if this object has a function with the given name
func (objectType *ObjectType) HasFunctionWithName(name string) bool {
	_, ok := objectType.functions[name]
	return ok
}

// AsType implements Type for ObjectType
func (objectType *ObjectType) AsType(codeGenerationContext *CodeGenerationContext) dst.Expr {
	embedded := objectType.EmbeddedProperties()
	properties := objectType.Properties()
	var fields []*dst.Field

	for _, f := range embedded {
		fields = append(fields, f.AsField(codeGenerationContext))
	}

	for _, f := range properties {
		fields = append(fields, f.AsField(codeGenerationContext))
	}

	if len(fields) > 0 {
		// if first field has Before:EmptyLine decoration, switch it to NewLine
		// this makes the output look nicer 🙂
		fields[0].Decs.Before = dst.NewLine
	}

	return &dst.StructType{
		Fields: &dst.FieldList{
			List: fields,
		},
	}
}

// RequiredPackageReferences returns a list of packages required by this
func (objectType *ObjectType) RequiredPackageReferences() *PackageReferenceSet {
	result := NewPackageReferenceSet()

	for _, property := range objectType.embedded {
		propertyType := property.PropertyType()
		result.Merge(propertyType.RequiredPackageReferences())
	}

	for _, property := range objectType.properties {
		propertyType := property.PropertyType()
		result.Merge(propertyType.RequiredPackageReferences())
	}

	for _, function := range objectType.functions {
		result.Merge(function.RequiredPackageReferences())
	}

	result.Merge(objectType.InterfaceImplementer.RequiredPackageReferences())

	return result
}

// References returns the set of all the types referred to by any property.
func (objectType *ObjectType) References() TypeNameSet {
	var results TypeNameSet
	for _, property := range objectType.properties {
		for ref := range property.PropertyType().References() {
			results = results.Add(ref)
		}
	}

	// Not collecting types from functions deliberately.
	return results
}

// Equals returns true if the passed type is a object type with the same properties, false otherwise
// The order of the properties is not relevant
func (objectType *ObjectType) Equals(t Type) bool {
	if objectType == t {
		return true
	}

	other, ok := t.(*ObjectType)
	if !ok {
		return false
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

		if !ourProperty.Equals(f) {
			// Different property, even though same name; not-equal
			return false
		}
	}

	if len(objectType.properties) != len(other.properties) {
		// Different number of properties, not equal
		return false
	}

	for n, f := range other.properties {
		ourProperty, ok := objectType.properties[n]
		if !ok {
			// Didn't find the property, not equal
			return false
		}

		if !ourProperty.Equals(f) {
			// Different property, even though same name; not-equal
			return false
		}
	}

	if len(objectType.functions) != len(other.functions) {
		// Different number of functions, not equal
		return false
	}

	for name, function := range other.functions {
		ourFunction, ok := objectType.functions[name]
		if !ok {
			// Didn't find the func, not equal
			return false
		}

		if !ourFunction.Equals(function) {
			// Different testcase, even though same name; not-equal
			return false
		}
	}

	if len(objectType.testcases) != len(other.testcases) {
		// Different number of test cases, not equal
		return false
	}

	for name, testcase := range other.testcases {
		ourCase, ok := objectType.testcases[name]
		if !ok {
			// Didn't find the func, not equal
			return false
		}

		if !ourCase.Equals(testcase) {
			// Different testcase, even though same name; not-equal
			return false
		}
	}

	if len(objectType.testcases) != len(other.testcases) {
		// Different number of test cases, not equal
		return false
	}

	for name, testcase := range other.testcases {
		ourCase, ok := objectType.testcases[name]
		if !ok {
			// Didn't find the func, not equal
			return false
		}

		if !ourCase.Equals(testcase) {
			// Different testcase, even though same name; not-equal
			return false
		}
	}

	return objectType.InterfaceImplementer.Equals(other.InterfaceImplementer)
}

// WithProperty creates a new ObjectType with another property attached to it
// Properties are unique by name, so this can be used to Add and Replace a property
func (objectType *ObjectType) WithProperty(property *PropertyDefinition) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	result.properties[property.propertyName] = property

	return result
}

// WithProperties creates a new ObjectType that's a copy with additional properties included
// Properties are unique by name, so this can be used to both Add and Replace properties.
func (objectType *ObjectType) WithProperties(properties ...*PropertyDefinition) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	for _, f := range properties {
		result.properties[f.propertyName] = f
	}

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
	result.properties = make(map[PropertyName]*PropertyDefinition)
	return result
}

// WithoutProperty creates a new ObjectType that's a copy without the specified property
func (objectType *ObjectType) WithoutProperty(name PropertyName) *ObjectType {
	// Create a copy of objectType to preserve immutability
	result := objectType.copy()
	delete(result.properties, name)

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
	result.functions[function.Name()] = function

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
	if !objectType.InterfaceImplementer.HasInterface(name) {
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
	result.testcases[testcase.Name()] = testcase
	return result
}

func (objectType *ObjectType) copy() *ObjectType {
	result := NewObjectType()

	for key, value := range objectType.embedded {
		result.embedded[key] = value
	}

	for key, value := range objectType.properties {
		result.properties[key] = value
	}

	for key, value := range objectType.functions {
		result.functions[key] = value
	}

	for key, value := range objectType.testcases {
		result.testcases[key] = value
	}

	result.InterfaceImplementer = objectType.InterfaceImplementer.copy()

	return result
}

// String implements fmt.Stringer
func (objectType *ObjectType) String() string {
	return "(object)"
}

func (objectType *ObjectType) HasTestCases() bool {
	return len(objectType.testcases) > 0
}

func (objectType *ObjectType) TestCases() []TestCase {
	var result []TestCase
	for _, tc := range objectType.testcases {
		result = append(result, tc)
	}

	return result
}

// IsObjectType returns true if the passed type is an object type OR if it is a wrapper type containing an object type
func IsObjectType(t Type) bool {
	_, ok := t.(*ObjectType)
	return ok
}

// IsObjectDefinition returns true if the passed definition is for a Arm type; false otherwise.
func IsObjectDefinition(definition TypeDefinition) bool {
	return IsObjectType(definition.theType)
}

func extractEmbeddedTypeName(t Type) (TypeName, error) {
	typeName, isTypeName := t.(TypeName)
	optionalType, isOptionalType := t.(*OptionalType)
	if isOptionalType {
		typeName, isTypeName = optionalType.Element().(TypeName)
	}

	if !isTypeName {
		return TypeName{}, errors.Errorf("embedded property type must be TypeName or Optional[TypeName], was: %T", t)
	}

	return typeName, nil
}
