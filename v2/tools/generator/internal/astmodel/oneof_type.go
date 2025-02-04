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
	"github.com/rotisserie/eris"
)

// OneOfType represents something that can be any one of a number of selected types
type OneOfType struct {
	swaggerName           string        // Name of the OneOf as defined in the Swagger file
	propertyObjects       []*ObjectType // Object definitions used to specify the properties held by this OneOf. May be empty.
	types                 TypeSet       // Set of all possible types
	discriminatorProperty string        // Identifies the discriminatorProperty property
	discriminatorValue    string        // Discriminator value used to identify this subtype
}

var _ Type = &OneOfType{}

// NewOneOfType creates a new instance of a OneOfType
func NewOneOfType(name string, types ...Type) *OneOfType {
	return &OneOfType{
		swaggerName: name,
		types:       MakeTypeSet(types...),
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
	return &oneOf.types
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
func (oneOf *OneOfType) References() TypeNameSet {
	result := NewTypeNameSet()
	oneOf.types.ForEach(func(t Type, _ int) {
		result = SetUnion(result, t.References())
	})

	return result
}

var oneOFailureMsg = "OneOfType should have been replaced by generation time by 'convertAllOfAndOneOf' phase"

// AsType always panics; OneOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf *OneOfType) AsTypeExpr(codeGenerationContext *CodeGenerationContext) (dst.Expr, error) {
	panic(eris.New(oneOFailureMsg))
}

// AsDeclarations always panics; OneOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf *OneOfType) AsDeclarations(
	codeGenerationContext *CodeGenerationContext,
	declContext DeclarationContext,
) ([]dst.Decl, error) {
	panic(eris.New(oneOFailureMsg))
}

// AsZero always panics; OneOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf *OneOfType) AsZero(_ TypeDefinitionSet, _ *CodeGenerationContext) dst.Expr {
	panic(eris.New(oneOFailureMsg))
}

// RequiredPackageReferences returns the union of the required imports of all the oneOf types
func (oneOf *OneOfType) RequiredPackageReferences() *PackageReferenceSet {
	panic(eris.New(oneOFailureMsg))
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

	// Check for different properties
	if oneOf.swaggerName != other.swaggerName {
		return false
	}

	if oneOf.discriminatorProperty != other.discriminatorProperty {
		return false
	}

	if oneOf.discriminatorValue != other.discriminatorValue {
		return false
	}

	// Check for different options to select from
	if !oneOf.types.Equals(other.types, overrides) {
		return false
	}

	// Check for different common properties
	if len(oneOf.propertyObjects) != len(other.propertyObjects) {
		return false
	}

	// Requiring exactly the same property objects in the same order is overly
	// strict as they're actually all merged together into a single object
	// and the order is not significant. Moreover, two one-of types would be the
	// same if the merge is the same, regardless of how many object types were there
	// to start with. This is all too complex to handle here though, so we'll just
	// use the strict check.
	for i := range oneOf.propertyObjects {
		if !oneOf.propertyObjects[i].Equals(other.propertyObjects[i], overrides) {
			return false
		}
	}

	return true
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
func (oneOf *OneOfType) WriteDebugDescription(builder *strings.Builder, currentPackage InternalPackageReference) {
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
