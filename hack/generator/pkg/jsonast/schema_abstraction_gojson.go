/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"math/big"
	"net/url"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

// GoJSONSchema implements the Schema abstraction for gojsonschema
type GoJSONSchema struct {
	inner *gojsonschema.SubSchema
}

// MakeGoJSONSchema wrapes a gojsonschema.SubSchema to conform to the Schema abstraction
func MakeGoJSONSchema(schema *gojsonschema.SubSchema) Schema {
	return GoJSONSchema{schema}
}

var _ Schema = GoJSONSchema{}

func transformGoJSONSlice(slice []*gojsonschema.SubSchema) []Schema {
	result := make([]Schema, len(slice))
	for i := range slice {
		result[i] = GoJSONSchema{slice[i]}
	}

	return result
}

func (schema GoJSONSchema) url() *url.URL {
	return schema.inner.ID.GetUrl()
}

func (schema GoJSONSchema) title() *string {
	return schema.inner.Title
}

func (schema GoJSONSchema) extensions(key string) interface{} {
	return nil
}

func (schema GoJSONSchema) hasType(schemaType SchemaType) bool {
	return schema.inner.Types.Contains(string(schemaType))
}

func (schema GoJSONSchema) requiredProperties() []string {
	return schema.inner.Required
}

func (schema GoJSONSchema) hasAllOf() bool {
	return len(schema.inner.AllOf) > 0
}

func (schema GoJSONSchema) allOf() []Schema {
	return transformGoJSONSlice(schema.inner.AllOf)
}

func (schema GoJSONSchema) hasAnyOf() bool {
	return len(schema.inner.AnyOf) > 0
}

func (schema GoJSONSchema) anyOf() []Schema {
	return transformGoJSONSlice(schema.inner.AnyOf)
}

func (schema GoJSONSchema) hasOneOf() bool {
	return len(schema.inner.OneOf) > 0
}

func (schema GoJSONSchema) oneOf() []Schema {
	return transformGoJSONSlice(schema.inner.OneOf)
}

func (schema GoJSONSchema) properties() map[string]Schema {
	result := make(map[string]Schema)
	for _, prop := range schema.inner.PropertiesChildren {
		result[prop.Property] = GoJSONSchema{prop}
	}

	return result
}

func (schema GoJSONSchema) maxLength() *int64 {
	return intPointerToInt64Pointer(schema.inner.MaxLength)
}

func (schema GoJSONSchema) minLength() *int64 {
	return intPointerToInt64Pointer(schema.inner.MinLength)
}

func intPointerToInt64Pointer(ptr *int) *int64 {
	if ptr == nil {
		return nil
	}

	value := int64(*ptr)
	return &value
}

func (schema GoJSONSchema) pattern() *regexp.Regexp {
	return schema.inner.Pattern
}

func (schema GoJSONSchema) maxItems() *int64 {
	return intPointerToInt64Pointer(schema.inner.MaxItems)
}

func (schema GoJSONSchema) minItems() *int64 {
	return intPointerToInt64Pointer(schema.inner.MinItems)
}

func (schema GoJSONSchema) uniqueItems() bool {
	return schema.inner.UniqueItems
}

func (schema GoJSONSchema) minValue() *big.Rat {
	r := schema.inner.Minimum
	if r != nil {
		return r
	}

	return schema.inner.ExclusiveMinimum
}

func (schema GoJSONSchema) minValueExclusive() bool {
	return schema.inner.ExclusiveMinimum != nil
}

func (schema GoJSONSchema) maxValue() *big.Rat {
	r := schema.inner.Maximum
	if r != nil {
		return r
	}

	return schema.inner.ExclusiveMaximum
}

func (schema GoJSONSchema) maxValueExclusive() bool {
	return schema.inner.ExclusiveMaximum != nil
}

func (schema GoJSONSchema) multipleOf() *big.Rat {
	return schema.inner.MultipleOf
}

func (schema GoJSONSchema) description() *string {
	return schema.inner.Description
}

func (schema GoJSONSchema) items() []Schema {
	return transformGoJSONSlice(schema.inner.ItemsChildren)
}

func (schema GoJSONSchema) additionalPropertiesAllowed() bool {
	aps := schema.inner.AdditionalProperties

	return aps == nil || aps != false
}

func (schema GoJSONSchema) additionalPropertiesSchema() Schema {
	result := schema.inner.AdditionalProperties

	if result == nil {
		return nil
	}

	return GoJSONSchema{result.(*gojsonschema.SubSchema)}
}

func (schema GoJSONSchema) enumValues() []string {
	return schema.inner.Enum
}

func (schema GoJSONSchema) isRef() bool {
	return schema.inner.RefSchema != nil
}

func (schema GoJSONSchema) refSchema() Schema {
	return GoJSONSchema{schema.inner.RefSchema}
}

func isURLPathSeparator(c rune) bool {
	return c == '/'
}

func (schema GoJSONSchema) refObjectName() (string, error) {
	return objectTypeOf(schema.inner.Ref.GetUrl())
}

func objectTypeOf(url *url.URL) (string, error) {
	fragmentParts := strings.FieldsFunc(url.Fragment, isURLPathSeparator)

	if len(fragmentParts) == 0 {
		return "", errors.Errorf("unexpected URL format: no fragment parts extracted from %q", url.String())
	}

	return fragmentParts[len(fragmentParts)-1], nil
}

func (schema GoJSONSchema) refGroupName() (string, error) {
	return groupOf(schema.inner.Ref.GetUrl())
}

func groupOf(url *url.URL) (string, error) {
	pathParts := strings.FieldsFunc(url.Path, isURLPathSeparator)

	file := pathParts[len(pathParts)-1]
	if !strings.HasSuffix(file, ".json") {
		return "", errors.Errorf("Unexpected URL format (doesn't point to .json file)")
	}

	return strings.TrimSuffix(file, ".json"), nil
}

var versionRegex = regexp.MustCompile(`\d{4}-\d{2}-\d{2}(-preview)?`)

func (schema GoJSONSchema) refVersion() (string, error) {
	return versionOf(schema.inner.Ref.GetUrl()), nil
}

func versionOf(url *url.URL) string {
	pathParts := strings.FieldsFunc(url.Path, isURLPathSeparator)

	for _, p := range pathParts {
		if versionRegex.MatchString(p) {
			return p
		}
	}

	// No version found, that's fine
	return ""
}
