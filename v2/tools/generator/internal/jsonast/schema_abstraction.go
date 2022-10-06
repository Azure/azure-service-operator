/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"math/big"
	"net/url"
	"regexp"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// Schema abstracts over the exact implementation of
// JSON schema we are using as we need to be able to
// process inputs provided by both 'gojsonschema' and
// 'go-openapi'. It is possible we could switch to
// 'go-openapi' for everything because it could handle
// both JSON Schema and Swagger; but this is not yet done.
type Schema interface {
	url() *url.URL
	Id() string // Unique ID of this Schema within the document (if any)
	title() *string
	description() *string

	// for extensions like x-ms-...
	extensionAsString(key string) (string, bool)
	extensionAsBool(key string) bool
	hasExtension(key string) bool

	hasType(schemaType SchemaType) bool

	// number things
	maxValue() *big.Rat
	minValue() *big.Rat
	maxValueExclusive() bool
	minValueExclusive() bool
	multipleOf() *big.Rat

	// string things
	maxLength() *int64
	minLength() *int64
	pattern() *regexp.Regexp
	format() string

	// complex things
	hasOneOf() bool
	oneOf() []Schema // Returns any directly embedded definitions held within a OneOf
	discriminator() string                // Returns the name of the discriminator field, or "" if it has none
	discriminatorValues() set.Set[string] // Finds the set of expected discriminators, or nil if this has none

	hasAnyOf() bool
	anyOf() []Schema

	hasAllOf() bool
	allOf() []Schema

	// enum things
	enumValues() []string

	// array things
	items() []Schema
	maxItems() *int64
	minItems() *int64
	uniqueItems() bool

	// object things
	requiredProperties() []string
	properties() map[string]Schema
	additionalPropertiesAllowed() bool
	additionalPropertiesSchema() Schema
	readOnly() bool

	// ref things
	isRef() bool
	refTypeName() (astmodel.TypeName, error)
	refSchema() Schema
}

// SchemaType defines the type of JSON schema node we are currently processing
type SchemaType string

// Definitions for different kinds of JSON schema
const (
	AnyOf   SchemaType = "anyOf"
	AllOf   SchemaType = "allOf"
	OneOf   SchemaType = "oneOf"
	Ref     SchemaType = "ref"
	Array   SchemaType = "array"
	Bool    SchemaType = "boolean"
	Int     SchemaType = "integer"
	Number  SchemaType = "number"
	Object  SchemaType = "object"
	String  SchemaType = "string"
	Enum    SchemaType = "enum"
	Unknown SchemaType = "unknown"
)
