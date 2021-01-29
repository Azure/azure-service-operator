/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/dave/dst"
)

// PropertyName is a semantic type
type PropertyName string

// PropertyDefinition encapsulates the definition of a property
type PropertyDefinition struct {
	propertyName PropertyName
	propertyType Type
	description  string
	isRequired   bool
	tags         map[string][]string
}

var _ fmt.Stringer = &PropertyDefinition{}

// NewPropertyDefinition is a factory method for creating a new PropertyDefinition
// name is the name for the new property (mandatory)
// propertyType is the type for the new property (mandatory)
func NewPropertyDefinition(propertyName PropertyName, jsonName string, propertyType Type) *PropertyDefinition {
	tags := make(map[string][]string)
	tags["json"] = []string{jsonName}

	return &PropertyDefinition{
		propertyName: propertyName,
		propertyType: propertyType,
		description:  "",
		tags:         tags,
	}
}

// PropertyName returns the name of the property
func (property *PropertyDefinition) PropertyName() PropertyName {
	return property.propertyName
}

// PropertyType returns the data type of the property
func (property *PropertyDefinition) PropertyType() Type {
	return property.propertyType
}

// SetRequired sets if the property is required or not
func (property *PropertyDefinition) SetRequired(required bool) *PropertyDefinition {
	result := *property
	result.isRequired = required
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

// WithType clones the property and returns it with a new type
func (property *PropertyDefinition) WithType(newType Type) *PropertyDefinition {

	if property.propertyType.Equals(newType) {
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
		//TODO: Do we want a generic helper that does this?
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
	if !property.hasOptionalType() && property.IsRequired() && !property.hasJsonOmitEmpty() {
		return property
	}
	result := property.copy()

	if property.hasOptionalType() {
		// Need to remove the optionality
		ot := property.propertyType.(*OptionalType)
		result.propertyType = ot.BaseType()
	}

	result.isRequired = true
	result = result.withoutJsonOmitEmpty()

	return result
}

// MakeOptional returns a new PropertyDefinition that has an optional value
func (property *PropertyDefinition) MakeOptional() *PropertyDefinition {
	if isTypeOptional(property.propertyType) && !property.IsRequired() && property.hasJsonOmitEmpty() {
		// No change required
		return property
	}

	result := property.copy()

	// Note that this function uses isTypeOptional while MakeRequired uses property.hasOptionalType
	// because in this direction we care if the type behaves optionally already (Map, Array included),
	// whereas in the MakeRequired direction we only care if the type is specifically *astmodel.Optional
	if !isTypeOptional(property.propertyType) {
		// Need to make the type optional
		result.propertyType = NewOptionalType(result.propertyType)
	}

	result.isRequired = false
	result = result.withJsonOmitEmpty()

	return result
}

// IsRequired returns true if the property is required;
// returns false otherwise.
func (property *PropertyDefinition) IsRequired() bool {
	return property.isRequired
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
	tags := property.renderedTags()

	var names []*dst.Ident
	if property.propertyName != "" {
		names = []*dst.Ident{dst.NewIdent(string(property.propertyName))}
	}

	var doc dst.Decorations
	if property.IsRequired() {
		AddValidationComments(&doc, []KubeBuilderValidation{ValidateRequired()})
	}

	// if we have validations, unwrap them
	propType := property.propertyType
	if validated, ok := propType.(*ValidatedType); ok {
		propType = validated.ElementType()
		AddValidationComments(&doc, validated.Validations().ToKubeBuilderValidations())
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
func (property *PropertyDefinition) Equals(f *PropertyDefinition) bool {
	return property == f || (property.propertyName == f.propertyName &&
		property.propertyType.Equals(f.propertyType) &&
		property.tagsEqual(f) &&
		property.isRequired == f.isRequired &&
		property.description == f.description)
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
