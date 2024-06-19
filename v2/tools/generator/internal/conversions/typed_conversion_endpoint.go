/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"strings"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TypedConversionEndpoint represents either a source or a destination target for a storage conversion
// In simple cases these will be fields, but can also represent indexed members of slices and maps.
type TypedConversionEndpoint struct {
	// theType is the Type of the value accessible via this endpoint
	theType astmodel.Type
	// name is the name of the underlying property, used to generate useful local identifiers
	name string
	// path is a dot qualified path that leads to the property, used to match properties that may have been flattened
	path string
}

func NewTypedConversionEndpoint(theType astmodel.Type, name string) *TypedConversionEndpoint {
	return &TypedConversionEndpoint{
		theType: theType,
		name:    name,
	}
}

// Name returns the actual property name of this endpoint
func (endpoint *TypedConversionEndpoint) Name() string {
	return endpoint.name
}

// Path returns a dot qualified path for this endpoint (might be empty)
func (endpoint *TypedConversionEndpoint) Path() string {
	return endpoint.path
}

// Type returns the type of this endpoint
func (endpoint *TypedConversionEndpoint) Type() astmodel.Type {
	return endpoint.theType
}

// WithType creates a new endpoint with a different type
func (endpoint *TypedConversionEndpoint) WithType(theType astmodel.Type) *TypedConversionEndpoint {
	result := *endpoint
	result.theType = theType
	return &result
}

// IsOptional returns true if the endpoint contains an optional type, false otherwise
func (endpoint *TypedConversionEndpoint) IsOptional() bool {
	_, result := astmodel.AsOptionalType(endpoint.Type())
	return result
}

// IsBagItem returns true if the endpoint contains a property bag item, false otherwise
func (endpoint *TypedConversionEndpoint) IsBagItem() bool {
	_, result := AsPropertyBagMemberType(endpoint.Type())
	return result
}

func (endpoint *TypedConversionEndpoint) WithPath(
	namePath []astmodel.PropertyName,
) *TypedConversionEndpoint {
	path := make([]string, len(namePath))
	for i, name := range namePath {
		path[i] = string(name)
	}

	result := *endpoint
	result.path = strings.Join(path, ".")
	return &result
}
