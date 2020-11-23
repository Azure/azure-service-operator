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
	ast "github.com/dave/dst"
)

// PropertyName is a semantic type
type PropertyName string

// PropertyDefinition encapsulates the definition of a property
type PropertyDefinition struct {
	propertyName PropertyName
	propertyType Type
	description  string
	validations  []Validation
	tags         map[string][]string
}

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

// WithValidation adds the given validation to the property's set of validations
func (property *PropertyDefinition) WithValidation(validation Validation) *PropertyDefinition {
	result := property.copy()
	result.validations = append(result.validations, validation)
	return result
}

// WithoutValidation removes all validations from the field
func (property *PropertyDefinition) WithoutValidation() *PropertyDefinition {
	result := property.copy()
	result.validations = nil
	return result
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
	if !property.hasOptionalType() && property.HasRequiredValidation() && !property.hasJsonOmitEmpty() {
		return property
	}
	result := property.copy()

	if property.hasOptionalType() {
		// Need to remove the optionality
		ot := property.propertyType.(*OptionalType)
		result.propertyType = ot.BaseType()
	}

	if !property.HasRequiredValidation() {
		result = result.WithValidation(ValidateRequired())
	}

	result = result.withoutJsonOmitEmpty()

	return result
}

// MakeOptional returns a new PropertyDefinition that has an optional value
func (property *PropertyDefinition) MakeOptional() *PropertyDefinition {
	if isTypeOptional(property.propertyType) && !property.HasRequiredValidation() && property.hasJsonOmitEmpty() {
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

	if property.HasRequiredValidation() {
		// Need to remove the Required validation
		var validations []Validation
		for _, v := range result.validations {
			if !v.HasName(RequiredValidationName) {
				validations = append(validations, v)
			}
		}

		result.validations = validations
	}

	result = result.withJsonOmitEmpty()

	return result
}

// HasRequiredValidation returns true if the property has validation specifying that it is required;
// returns false otherwise.
func (property *PropertyDefinition) HasRequiredValidation() bool {
	required := ValidateRequired()
	for _, v := range property.validations {
		if v == required {
			return true
		}
	}

	return false
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
func (property *PropertyDefinition) AsField(codeGenerationContext *CodeGenerationContext) *ast.Field {
	tags := property.renderedTags()

	var names []*ast.Ident
	if property.propertyName != "" {
		names = []*ast.Ident{ast.NewIdent(string(property.propertyName))}
	}

	// We don't use StringLiteral() for the tag as it adds extra quotes
	result := &ast.Field{
		Decs: ast.FieldDecorations{
			NodeDecs: ast.NodeDecs{
				Before: ast.NewLine,
			},
		},
		Names: names,
		Type:  property.PropertyType().AsType(codeGenerationContext),
		Tag:   astbuilder.TextLiteralf("`%s`", tags),
	}

	// generate validation comments:
	for _, validation := range property.validations {
		result.Decs.Before = ast.EmptyLine
		// these are not doc comments but they must go here to be emitted before the property
		astbuilder.AddComment(&result.Decs.Start, GenerateKubebuilderComment(validation))
	}

	// generate comment:
	if property.description != "" {
		result.Decs.Before = ast.EmptyLine
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

func (property *PropertyDefinition) validationsEqual(f *PropertyDefinition) bool {
	if len(property.validations) != len(f.validations) {
		return false
	}

	for i := 0; i < len(property.validations); i++ {
		if !property.validations[i].Equals(f.validations[i]) {
			return false
		}
	}

	return true
}

// Equals tests to see if the specified PropertyDefinition specifies the same property
func (property *PropertyDefinition) Equals(f *PropertyDefinition) bool {
	return property == f || (property.propertyName == f.propertyName &&
		property.propertyType.Equals(f.propertyType) &&
		property.tagsEqual(f) &&
		property.validationsEqual(f) &&
		property.description == f.description)
}

func (property *PropertyDefinition) copy() *PropertyDefinition {
	result := *property

	// Copy ptr fields
	result.tags = make(map[string][]string, len(property.tags))
	for key, value := range property.tags {
		result.tags[key] = append([]string(nil), value...)
	}

	result.validations = nil
	result.validations = append([]Validation(nil), property.validations...)

	return &result
}
