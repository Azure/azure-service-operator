/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dave/dst"
	"github.com/pkg/errors"
)

// OneOfType represents something that can be any one of a number of selected types
type OneOfType struct {
	swaggerName           string                        // Name of the OneOf as defined in the Swagger file
	propertyObjects       []*ObjectType                 // Object definitions used to specify the properties held by this OneOf. May be empty.
	options               TypeNameSet[InternalTypeName] // References to the type names of the options for this OneOf. May be empty.
	types                 TypeSet                       // Set of all possible types
	discriminatorProperty string                        // Identifies the discriminatorProperty property
	discriminatorValue    string                        // Discriminator value used to identify this subtype
}

var _ Type = &OneOfType{}

// NewOneOfType creates a new instance of a OneOfType
func NewOneOfType(name string, types ...Type) *OneOfType {
	return &OneOfType{
		swaggerName: name,
		types:       MakeTypeSet(types...),
		options:     NewTypeNameSet[InternalTypeName](),
	}
}

// Name returns the internal (swagger) name of the OneOf
func (oneOf *OneOfType) Name() string {
	return oneOf.swaggerName
}

// DiscriminatorProperty returns the name of the discriminatorProperty property (if any)
func (oneOf *OneOfType) DiscriminatorProperty() string {
	return oneOf.discriminatorProperty
}

// HasDiscriminatorProperty returns true if the OneOf has a discriminator property
func (oneOf *OneOfType) HasDiscriminatorProperty() bool {
	return oneOf.discriminatorProperty != ""
}

// WithDiscriminatorProperty returns a new OneOf object with the specified discriminatorProperty property
func (oneOf *OneOfType) WithDiscriminatorProperty(discriminator string) *OneOfType {
	if oneOf.discriminatorProperty == discriminator {
		return oneOf
	}

	result := oneOf.copy()
	result.discriminatorProperty = discriminator
	return result
}

// DiscriminatorValue returns the discriminator value used to identify this subtype
func (oneOf *OneOfType) DiscriminatorValue() string {
	return oneOf.discriminatorValue
}

// HasDiscriminatorValue returns true if the OneOf has a discriminator value
func (oneOf *OneOfType) HasDiscriminatorValue() bool {
	return oneOf.discriminatorValue != ""
}

func (oneOf *OneOfType) WithDiscriminatorValue(value string) *OneOfType {
	if oneOf.discriminatorValue == value {
		return oneOf
	}

	result := oneOf.copy()
	result.discriminatorValue = value
	return result
}

// WithType returns a new OneOf with the specified type included
func (oneOf *OneOfType) WithType(t Type) *OneOfType {
	if oneOf.types.Contains(t, EqualityOverrides{}) {
		// Already present
		return oneOf
	}

	result := oneOf.copy()
	result.types = result.types.Copy()
	result.types.Add(t)
	return result
}

// WithoutType returns a new OneOf with the specified type removed
func (oneOf *OneOfType) WithoutType(t Type) *OneOfType {
	if !oneOf.types.Contains(t, EqualityOverrides{}) {
		// Nothing to remove
		return oneOf
	}

	result := oneOf.copy()
	result.types = result.types.Copy()
	result.types.Remove(t)
	return result
}

// WithTypes returns a new OneOf with only the specified types
func (oneOf *OneOfType) WithTypes(types []Type) *OneOfType {
	result := oneOf.copy()
	result.types = MakeTypeSet(types...)
	return result
}

// Types returns what subtypes the OneOf may be.
// Exposed as ReadonlyTypeSet so caller cannot break invariants.
func (oneOf *OneOfType) Types() ReadonlyTypeSet {
	return oneOf.types
}

// PropertyObjects returns all the ObjectTypes that define the properties of this OneOf
func (oneOf *OneOfType) PropertyObjects() []*ObjectType {
	return oneOf.propertyObjects
}

// WithAdditionalPropertyObject returns a new OneOf that includes the specified properties as well as those already present
func (oneOf *OneOfType) WithAdditionalPropertyObject(propertyObject *ObjectType) *OneOfType {
	result := oneOf.copy()
	result.propertyObjects = append(result.propertyObjects, propertyObject)
	return result
}

// WithoutAnyPropertyObjects returns a new OneOf that has no Object properties
func (oneOf *OneOfType) WithoutAnyPropertyObjects() *OneOfType {
	result := oneOf.copy()
	result.propertyObjects = nil
	return result
}

// References returns any type referenced by the OneOf types
func (oneOf *OneOfType) References() TypeNameSet[TypeName] {
	result := NewTypeNameSet[TypeName]()
	oneOf.types.ForEach(func(t Type, _ int) {
		result = SetUnion(result, t.References())
	})

	return result
}

var oneOfPanicMsg = "OneOfType should have been replaced by generation time by 'convertAllOfAndOneOf' phase"

// AsType always panics; OneOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf *OneOfType) AsType(_ *CodeGenerationContext) dst.Expr {
	panic(errors.New(oneOfPanicMsg))
}

// AsDeclarations always panics; OneOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf *OneOfType) AsDeclarations(_ *CodeGenerationContext, _ DeclarationContext) []dst.Decl {
	panic(errors.New(oneOfPanicMsg))
}

// AsZero always panics; OneOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf *OneOfType) AsZero(_ TypeDefinitionSet, _ *CodeGenerationContext) dst.Expr {
	panic(errors.New(oneOfPanicMsg))
}

// RequiredPackageReferences returns the union of the required imports of all the oneOf types
func (oneOf *OneOfType) RequiredPackageReferences() *PackageReferenceSet {
	panic(errors.New(oneOfPanicMsg))
}

// Equals returns true if the other Type is a OneOfType that contains
// the same set of types
func (oneOf *OneOfType) Equals(t Type, overrides EqualityOverrides) bool {
	if oneOf == t {
		return true // short-circuit
	}

	other, ok := t.(*OneOfType)
	if !ok {
		return false
	}

	return oneOf.types.Equals(other.types, overrides)
}

// String implements fmt.Stringer
func (oneOf *OneOfType) String() string {
	var subStrings []string
	oneOf.types.ForEach(func(t Type, _ int) {
		subStrings = append(subStrings, t.String())
	})

	sort.Slice(subStrings, func(i, j int) bool {
		return subStrings[i] < subStrings[j]
	})

	return fmt.Sprintf("(oneOf: %s)", strings.Join(subStrings, ", "))
}

// WriteDebugDescription adds a description of the current type to the passed builder.
// builder receives the full description, including nested types.
// definitions is a dictionary for resolving named types.
func (oneOf *OneOfType) WriteDebugDescription(builder *strings.Builder, currentPackage PackageReference) {
	if oneOf == nil {
		builder.WriteString("<nilOneOf>")
		return
	}

	builder.WriteString("OneOf[")

	if oneOf.swaggerName != "" {
		builder.WriteString(oneOf.swaggerName)
		builder.WriteRune(';')
	}

	if oneOf.discriminatorProperty != "" {
		builder.WriteString("d:")
		builder.WriteString(oneOf.discriminatorProperty)
		builder.WriteRune(';')
	}

	if oneOf.discriminatorValue != "" {
		builder.WriteString("v:")
		builder.WriteString(oneOf.discriminatorValue)
		builder.WriteRune(';')
	}

	builder.WriteString("types:")
	oneOf.types.ForEach(func(t Type, ix int) {
		if ix > 0 {
			builder.WriteString("|")
		}
		t.WriteDebugDescription(builder, currentPackage)
	})
	builder.WriteString("]")
}

func (oneOf *OneOfType) copy() *OneOfType {
	// We can share internal sets and slices as we use copy-on-write semantics
	result := *oneOf
	return &result
}

// AsOneOfType unwraps any wrappers around the provided type and returns either the underlying OneOfType and true,
// or nil and false.
func AsOneOfType(t Type) (*OneOfType, bool) {
	if oneOf, ok := t.(*OneOfType); ok {
		return oneOf, true
	}

	if wrapper, ok := t.(MetaType); ok {
		return AsOneOfType(wrapper.Unwrap())
	}

	return nil, false
}
