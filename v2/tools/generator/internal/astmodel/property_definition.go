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
	"golang.org/x/exp/slices"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/readonly"
)

// PropertyName is a semantic type
type PropertyName string

// Implement Stringer for easy fmt printing
func (pn PropertyName) String() string {
	return string(pn)
}

// PropertyDefinition encapsulates the definition of a property
type PropertyDefinition struct {
	propertyName                     PropertyName
	propertyType                     Type
	description                      string
	hasKubebuilderRequiredValidation bool

	// Two properties to handle flattening:
	// - flatten is set when a property should be flattened and its inner properties
	//   moved out into the parent object.
	// - flattenedFrom is set once an ‘inner’ property has been flattened, to record
	//   which property(/ies) (that had flatten:true) it was flattened from. this
	//   is stored with most-recent first so that `[0].[1].[2]` would form the
	//   correct nested property name.
	flatten bool // maps to x-ms-client-flatten: should the propertyType be flattened into the parent?

	// a list of property names from whence this property was flattened
	// the last entry is always the original property name, so it will look like:
	// x,y,z,propName
	flattenedFrom []PropertyName

	isSecret bool
	readOnly bool

	tags readonly.Map[string, []string] // Note: have to be careful about not mutating inner []string
}

func (property *PropertyDefinition) AddFlattenedFrom(name PropertyName) *PropertyDefinition {
	result := *property
	result.flattenedFrom = append([]PropertyName{name}, result.flattenedFrom...)
	return &result
}

func (property *PropertyDefinition) WasFlattened() bool {
	return len(property.flattenedFrom) > 1
}

func (property *PropertyDefinition) WasFlattenedFrom(name PropertyName) bool {
	return property.WasFlattened() && property.flattenedFrom[0] == name
}

func (property *PropertyDefinition) FlattenedFrom() []PropertyName {
	return append([]PropertyName{}, property.flattenedFrom...)
}

var _ fmt.Stringer = &PropertyDefinition{}

// NewPropertyDefinition is a factory method for creating a new PropertyDefinition
// name is the name for the new property (mandatory)
// propertyType is the type for the new property (mandatory)
func NewPropertyDefinition(propertyName PropertyName, jsonName string, propertyType Type) *PropertyDefinition {
	return &PropertyDefinition{
		propertyName:  propertyName,
		propertyType:  propertyType,
		description:   "",
		flattenedFrom: []PropertyName{propertyName},
		tags: readonly.CreateMap(map[string][]string{
			"json": {jsonName, "omitempty"},
		}),
	}
}

// WithName returns a new PropertyDefinition with the given name
func (property *PropertyDefinition) WithName(name PropertyName) *PropertyDefinition {
	result := property.copy()
	result.propertyName = name
	return result
}

// WithName returns a new PropertyDefinition with the JSON name
func (property *PropertyDefinition) WithJsonName(jsonName string) *PropertyDefinition {
	result := property.copy()
	// TODO post-alpha: replace result.tags with structured types
	jsonTag, _ := result.tags.Get("json")
	if jsonName != jsonTag[0] {
		// copy tags and update
		jsonTagValues := slices.Clone(jsonTag)
		jsonTagValues[0] = jsonName
		result.tags = result.tags.With("json", jsonTagValues)
	}

	return result
}

// PropertyName returns the name of the property
func (property *PropertyDefinition) PropertyName() PropertyName {
	return property.propertyName
}

// PropertyType returns the data type of the property
func (property *PropertyDefinition) PropertyType() Type {
	return property.propertyType
}

// SetFlatten sets if the property should be flattened or not
func (property *PropertyDefinition) SetFlatten(flatten bool) *PropertyDefinition {
	if flatten == property.flatten {
		return property
	}

	result := property.copy()
	result.flatten = flatten
	return result
}

// WithReadOnly returns a new PropertyDefinition with ReadOnly set to the specified value
func (property *PropertyDefinition) WithReadOnly(readOnly bool) *PropertyDefinition {
	if property.readOnly == readOnly {
		return property
	}

	result := property.copy()
	result.readOnly = readOnly
	return result
}

// WithIsSecret returns a new PropertyDefinition with IsSecret set to the specified value
func (property *PropertyDefinition) WithIsSecret(secret bool) *PropertyDefinition {
	if secret == property.isSecret {
		return property
	}

	result := property.copy()
	result.isSecret = secret
	return result
}

// WithDescription returns a new PropertyDefinition with the specified description
func (property *PropertyDefinition) WithDescription(description string) *PropertyDefinition {
	if description == property.description {
		return property
	}

	result := property.copy()
	result.description = description
	return result
}

// Description returns the property description
func (property *PropertyDefinition) Description() string {
	return property.description
}

// WithType clones the property and returns it with a new type
func (property *PropertyDefinition) WithType(newType Type) *PropertyDefinition {
	if newType == nil {
		panic("nil type provided to WithType")
	}

	if TypeEquals(property.propertyType, newType) {
		return property
	}

	result := property.copy()
	result.propertyType = newType
	return result
}

// HasName returns true if the property has the given name
func (property *PropertyDefinition) HasName(name PropertyName) bool {
	return property.propertyName == name
}

// WithTag adds the given tag to the field
func (property *PropertyDefinition) WithTag(key string, value string) *PropertyDefinition {
	values, ok := property.tags.Get(key)
	if ok {
		if slices.Contains(values, value) {
			return property
		} else {
			result := property.copy()
			result.tags = result.tags.With(key, append(slices.Clone(values), value))
			return result
		}
	} else {
		result := property.copy()
		result.tags = result.tags.With(key, []string{value})
		return result
	}
}

// WithoutTag removes the given tag (or value of tag) from the field. If value is empty, remove the entire tag.
// if value is not empty, remove just that value.
func (property *PropertyDefinition) WithoutTag(key string, value string) *PropertyDefinition {
	if value != "" {
		// Find the value and remove it
		// TODO: Do we want a generic helper that does this?
		values, ok := property.tags.Get(key)
		if !ok {
			return property
		}

		result := property.copy()
		tagsWithoutValue := make([]string, 0, len(values))
		for _, item := range values {
			if item == value {
				continue
			}
			tagsWithoutValue = append(tagsWithoutValue, item)
		}

		if len(tagsWithoutValue) == 0 {
			result.tags = result.tags.Without(key)
		} else {
			result.tags = result.tags.With(key, tagsWithoutValue)
		}
		return result
	} else {
		result := property.copy()
		result.tags = result.tags.Without(key)
		return result
	}
}

// HasTag returns true if the property has the specified tag
func (property *PropertyDefinition) HasTag(key string) bool {
	return property.tags.ContainsKey(key)
}

// HasTagValue returns true if the property has the specified tag value
func (property *PropertyDefinition) HasTagValue(key string, value string) bool {
	if values, ok := property.tags.Get(key); ok {
		return slices.Contains(values, value)
	}

	return false
}

// Tag returns the tag values of a given tag key
func (property *PropertyDefinition) Tag(key string) ([]string, bool) {
	if result, ok := property.tags.Get(key); ok {
		return slices.Clone(result), true
	}

	return nil, false
}

// JSONName returns the JSON name of the property, or false if there is no json name
func (property *PropertyDefinition) JSONName() (string, bool) {
	jsonTag, ok := property.Tag("json")
	if !ok {
		return "", false
	}

	return jsonTag[0], true
}

// MakeRequired returns a new PropertyDefinition that is marked as required
func (property *PropertyDefinition) MakeRequired() *PropertyDefinition {
	if property.IsRequired() {
		return property
	}

	if !isTypeOptional(property.PropertyType()) {
		panic(fmt.Sprintf(
			"property %s with non-optional type %T cannot be marked kubebuilder:validation:Required.",
			property.PropertyName(),
			property.PropertyType()))
	}

	result := property.copy()
	result.hasKubebuilderRequiredValidation = true

	return result
}

// MakeOptional returns a new PropertyDefinition that has an optional value
func (property *PropertyDefinition) MakeOptional() *PropertyDefinition {
	if !property.IsRequired() {
		// No change required
		return property
	}

	result := property.copy()
	result.hasKubebuilderRequiredValidation = false

	return result
}

// MakeTypeRequired makes this properties Type required (non-optional). Note that this does
// NOT impact the kubebuilder:validation:Required annotation. For that, see MakeRequired.
func (property *PropertyDefinition) MakeTypeRequired() *PropertyDefinition {
	if !canTypeBeMadeRequired(property.propertyType) {
		return property
	}

	result := property.copy()
	result.propertyType = makePropertyTypeRequired(property.PropertyType())

	return result
}

// MakeTypeOptional makes this properties Type optional. Note that this does
// NOT impact the kubebuilder:validation:Required annotation. For that, see MakeOptional.
func (property *PropertyDefinition) MakeTypeOptional() *PropertyDefinition {
	if isTypeOptional(property.propertyType) {
		return property
	}

	result := property.copy()
	result.propertyType = makePropertyTypeOptional(property.PropertyType())

	return result
}

func makePropertyTypeRequired(t Type) Type {
	switch typ := t.(type) {
	case *OptionalType:
		return typ.Element()
	case *ValidatedType:
		return typ.WithType(makePropertyTypeRequired(typ.ElementType()))
	case MetaType:
		// We didn't use a visitor here for two reasons: Reason 1 is that visitor has an err case that we don't want.
		// Reason 2 is that visitor by default visits types inside Maps and Lists, which we also do not want. We could
		// have overridden those visit methods but this is cleaner than that and fails in a safe way if new types
		// are ever added.
		panic(fmt.Sprintf("cannot make %T required", t))
	}

	return t
}

func makePropertyTypeOptional(t Type) Type {
	// we must check if the type is validated to preserve the validation invariant
	// that all validations are *directly* under either a named type or a property
	if vType, ok := t.(*ValidatedType); ok {
		return NewValidatedType(NewOptionalType(vType.ElementType()), vType.validations)
	} else {
		return NewOptionalType(t)
	}
}

// IsRequired returns true if the property is required;
// returns false otherwise.
func (property *PropertyDefinition) IsRequired() bool {
	return property.hasKubebuilderRequiredValidation
}

// Flatten returns true iff the property is marked with 'flatten'.
func (property *PropertyDefinition) Flatten() bool {
	return property.flatten
}

// ReadOnly returns true iff the property is a secret.
func (property *PropertyDefinition) ReadOnly() bool {
	return property.readOnly
}

// IsSecret returns true iff the property is a secret.
func (property *PropertyDefinition) IsSecret() bool {
	return property.isSecret
}

func (property *PropertyDefinition) renderedTags() string {
	orderedKeys := property.tags.Keys()

	sort.Slice(orderedKeys, func(i, j int) bool {
		return orderedKeys[i] < orderedKeys[j]
	})

	tags := make([]string, 0, len(orderedKeys))
	for _, key := range orderedKeys {
		values, _ := property.tags.Get(key)
		tagString := strings.Join(values, ",")
		tags = append(tags, fmt.Sprintf("%s:%q", key, tagString))
	}

	return strings.Join(tags, " ")
}

// AsField generates a Go AST field node representing this property definition
func (property *PropertyDefinition) AsField(codeGenerationContext *CodeGenerationContext) *dst.Field {
	if property.flatten {
		panic(fmt.Sprintf("property %s marked for flattening was not flattened", property.propertyName))
	}

	tags := property.renderedTags()

	var names []*dst.Ident
	if property.propertyName != "" {
		names = []*dst.Ident{dst.NewIdent(string(property.propertyName))}
	}

	var doc dst.Decorations
	if property.IsRequired() {
		AddValidationComments(&doc, []KubeBuilderValidation{MakeRequiredValidation()})
	}

	// if we have validations, unwrap them
	propType := property.propertyType
	if validated, ok := propType.(*ValidatedType); ok {
		propType = validated.ElementType()
		AddValidationComments(&doc, validated.Validations().ToKubeBuilderValidations())
	}

	// Some types opt out of codegen by returning nil
	typeDeclaration := propType.AsType(codeGenerationContext)
	if typeDeclaration == nil {
		return nil
	}

	before := dst.NewLine
	if len(doc) > 0 {
		before = dst.EmptyLine
	}

	// We don't use StringLiteral() for the tag as it adds extra quotes
	result := &dst.Field{
		Decs: dst.FieldDecorations{
			NodeDecs: dst.NodeDecs{
				Start:  doc,
				Before: before,
			},
		},
		Names: names,
		Type:  propType.AsType(codeGenerationContext),
		Tag:   astbuilder.TextLiteralf("`%s`", tags),
	}

	// generate comment:
	if property.description != "" {
		result.Decs.Before = dst.EmptyLine
		astbuilder.AddWrappedComment(&result.Decs.Start, fmt.Sprintf("%s: %s", property.propertyName, property.description))
	}

	return result
}

func (property *PropertyDefinition) tagsEqual(f *PropertyDefinition) bool {
	// Comparison here takes ordering into account because for tags like
	// json, ordering matters - `json:"foo,omitempty"` is different than
	// `json:"omitempty,foo`
	return property.tags.Equals(f.tags, slices.Equal[string])
}

// Equals tests to see if the specified PropertyDefinition specifies the same property
func (property *PropertyDefinition) Equals(o *PropertyDefinition, overrides EqualityOverrides) bool {
	return property == o || (property.propertyName == o.propertyName &&
		property.propertyType.Equals(o.propertyType, overrides) &&
		property.flatten == o.flatten &&
		propertyNameSlicesEqual(property.flattenedFrom, o.flattenedFrom) &&
		property.isSecret == o.isSecret &&
		property.tagsEqual(o) &&
		property.hasKubebuilderRequiredValidation == o.hasKubebuilderRequiredValidation &&
		property.description == o.description)
}

func (property *PropertyDefinition) copy() *PropertyDefinition {
	result := *property
	return &result
}

func (property *PropertyDefinition) String() string {
	return fmt.Sprintf("%s: %s %s", property.propertyName, property.propertyType, property.renderedTags())
}
