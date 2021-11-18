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

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
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

	tags map[string][]string
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
	tags := make(map[string][]string)
	tags["json"] = []string{jsonName}

	return &PropertyDefinition{
		propertyName:  propertyName,
		propertyType:  propertyType,
		description:   "",
		flattenedFrom: []PropertyName{propertyName},
		tags:          tags,
	}
}

// WithName returns a new PropertyDefinition with the given name
func (property *PropertyDefinition) WithName(name PropertyName) *PropertyDefinition {
	result := *property
	result.propertyName = name
	return &result
}

// WithName returns a new PropertyDefinition with the JSON name
func (property *PropertyDefinition) WithJsonName(jsonName string) *PropertyDefinition {
	result := *property
	// TODO post-alpha: replace result.tags with structured types
	if jsonName != result.tags["json"][0] {
		// copy tags and update
		newTags := cloneMapOfStringToSliceOfString(result.tags)
		jsonTagValues := cloneSliceOfString(newTags["json"])
		jsonTagValues[0] = jsonName
		newTags["json"] = jsonTagValues

		result.tags = newTags
	}

	return &result
}

func cloneMapOfStringToSliceOfString(m map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for k, v := range m {
		result[k] = v
	}

	return result
}

func cloneSliceOfString(s []string) []string {
	result := make([]string, len(s))
	copy(result, s)
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

// WithKubebuilderRequiredValidation returns a new PropertyDefinition with the given Kubebuilder required annotation.
// Note that this does not in any way change the underlying type that this PropertyDefinition points to, unlike
// MakeRequired and MakeOptional.
func (property *PropertyDefinition) WithKubebuilderRequiredValidation(required bool) *PropertyDefinition {
	if required == property.hasKubebuilderRequiredValidation {
		return property
	}

	result := *property
	result.hasKubebuilderRequiredValidation = required
	return &result
}

// SetFlatten sets if the property should be flattened or not
func (property *PropertyDefinition) SetFlatten(flatten bool) *PropertyDefinition {
	if flatten == property.flatten {
		return property
	}

	result := *property
	result.flatten = flatten
	return &result
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
	// Check if exists
	for _, v := range property.tags[key] {
		if v == value {
			return property
		}
	}

	result := property.copy()
	result.tags[key] = append(result.tags[key], value)

	return result
}

// WithoutTag removes the given tag (or value of tag) from the field. If value is empty, remove the entire tag.
// if value is not empty, remove just that value.
func (property *PropertyDefinition) WithoutTag(key string, value string) *PropertyDefinition {
	result := property.copy()

	if value != "" {
		// Find the value and remove it
		// TODO: Do we want a generic helper that does this?
		var tagsWithoutValue []string
		for _, item := range result.tags[key] {
			if item == value {
				continue
			}
			tagsWithoutValue = append(tagsWithoutValue, item)
		}

		if len(tagsWithoutValue) == 0 {
			delete(result.tags, key)
		} else {
			result.tags[key] = tagsWithoutValue
		}
	} else {
		delete(result.tags, key)
	}

	return result
}

// HasTag returns true if the property has the specified tag
func (property *PropertyDefinition) HasTag(key string) bool {
	_, ok := property.tags[key]
	return ok
}

// HasTagValue returns true if the property has the specified tag value
func (property *PropertyDefinition) HasTagValue(key string, value string) bool {
	if values, ok := property.tags[key]; ok {
		for _, item := range values {
			if item == value {
				return true
			}
		}
	}

	return false
}

// Tag returns the tag values of a given tag key
func (property *PropertyDefinition) Tag(key string) ([]string, bool) {
	result, ok := property.tags[key]
	return result, ok
}

func (property *PropertyDefinition) hasJsonOmitEmpty() bool {
	return property.HasTagValue("json", "omitempty")
}

func (property *PropertyDefinition) withJsonOmitEmpty() *PropertyDefinition {
	return property.WithTag("json", "omitempty")
}

func (property *PropertyDefinition) withoutJsonOmitEmpty() *PropertyDefinition {
	return property.WithoutTag("json", "omitempty")
}

// MakeRequired returns a new PropertyDefinition that is marked as required
func (property *PropertyDefinition) MakeRequired() *PropertyDefinition {
	if !property.hasOptionalType() && property.HasKubebuilderRequiredValidation() && !property.hasJsonOmitEmpty() {
		return property
	}
	result := property.copy()

	if property.hasOptionalType() {
		// Need to remove the optionality
		ot := property.propertyType.(*OptionalType)
		result.propertyType = ot.BaseType()
	}

	result.hasKubebuilderRequiredValidation = true
	result = result.withoutJsonOmitEmpty()

	return result
}

// MakeOptional returns a new PropertyDefinition that has an optional value
func (property *PropertyDefinition) MakeOptional() *PropertyDefinition {
	if isTypeOptional(property.propertyType) && !property.HasKubebuilderRequiredValidation() && property.hasJsonOmitEmpty() {
		// No change required
		return property
	}

	result := property.copy()

	// Note that this function uses isTypeOptional while MakeRequired uses property.hasOptionalType
	// because in this direction we care if the type behaves optionally already (Map, Array included),
	// whereas in the MakeRequired direction we only care if the type is specifically *astmodel.Optional
	if !isTypeOptional(property.propertyType) {
		// Need to make the type optional

		// we must check if the type is validated to preserve the validation invariant
		// that all validations are *directly* under either a named type or a property
		if vType, ok := result.propertyType.(*ValidatedType); ok {
			result.propertyType = NewValidatedType(NewOptionalType(vType.ElementType()), vType.validations)
		} else {
			result.propertyType = NewOptionalType(result.propertyType)
		}
	}

	result.hasKubebuilderRequiredValidation = false
	result = result.withJsonOmitEmpty()

	return result
}

// HasKubebuilderRequiredValidation returns true if the property is required;
// returns false otherwise.
func (property *PropertyDefinition) HasKubebuilderRequiredValidation() bool {
	return property.hasKubebuilderRequiredValidation
}

// Flatten returns true iff the property is marked with 'flatten'.
func (property *PropertyDefinition) Flatten() bool {
	return property.flatten
}

// hasOptionalType returns true if the type of this property is an optional reference to a value
// (and might therefore be nil).
func (property *PropertyDefinition) hasOptionalType() bool {
	_, ok := property.propertyType.(*OptionalType)
	return ok
}

func (property *PropertyDefinition) renderedTags() string {
	var orderedKeys []string
	for key := range property.tags {
		orderedKeys = append(orderedKeys, key)
	}

	sort.Slice(orderedKeys, func(i, j int) bool {
		return orderedKeys[i] < orderedKeys[j]
	})

	var tags []string
	for _, key := range orderedKeys {
		tagString := strings.Join(property.tags[key], ",")
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
	if property.HasKubebuilderRequiredValidation() {
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
		astbuilder.AddWrappedComment(&result.Decs.Start, fmt.Sprintf("%s: %s", property.propertyName, property.description), 80)
	}

	return result
}

func (property *PropertyDefinition) tagsEqual(f *PropertyDefinition) bool {
	if len(property.tags) != len(f.tags) {
		return false
	}

	for key, value := range property.tags {
		otherValue, ok := f.tags[key]
		if !ok || len(value) != len(otherValue) {
			return false
		}
		// Comparison here takes ordering into account because for tags like
		// json, ordering matters - `json:"foo,omitempty"` is different than
		// `json:"omitempty,foo`
		for i := 0; i < len(value); i++ {
			if value[i] != otherValue[i] {
				return false
			}
		}
	}

	return true
}

// Equals tests to see if the specified PropertyDefinition specifies the same property
func (property *PropertyDefinition) Equals(o *PropertyDefinition, overrides EqualityOverrides) bool {
	return property == o || (property.propertyName == o.propertyName &&
		property.propertyType.Equals(o.propertyType, overrides) &&
		property.flatten == o.flatten &&
		propertyNameSlicesEqual(property.flattenedFrom, o.flattenedFrom) &&
		property.tagsEqual(o) &&
		property.hasKubebuilderRequiredValidation == o.hasKubebuilderRequiredValidation &&
		property.description == o.description)
}

func (property *PropertyDefinition) copy() *PropertyDefinition {
	result := *property

	// Copy ptr fields
	result.tags = make(map[string][]string, len(property.tags))
	for key, value := range property.tags {
		result.tags[key] = append([]string(nil), value...)
	}

	return &result
}

func (property *PropertyDefinition) String() string {
	return fmt.Sprintf("%s: %s %s", property.propertyName, property.propertyType, property.renderedTags())
}
