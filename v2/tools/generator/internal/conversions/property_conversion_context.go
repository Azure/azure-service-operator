/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// PropertyConversionContext captures additional supporting information that may be needed when a
// storage conversion factory creates a conversion
type PropertyConversionContext struct {
	// types is a map of all known type definitions, used to resolve TypeNames to actual types
	types astmodel.Types
	// functionName is the name of the function we're currently generating
	functionName string
	// direction is the direction of the conversion we're generating
	direction Direction
	// propertyBagName is the name of the local variable used for a property bag, or "" if we don't have one
	propertyBagName string
	// Configuration containing additional metadata for generating conversions
	configuration *config.ObjectModelConfiguration
	// idFactory is used for generating method names
	idFactory astmodel.IdentifierFactory
}

// NewPropertyConversionContext creates a new instance of a PropertyConversionContext
func NewPropertyConversionContext(
	types astmodel.Types,
	idFactory astmodel.IdentifierFactory,
	configuration *config.ObjectModelConfiguration) *PropertyConversionContext {
	return &PropertyConversionContext{
		types:           types,
		idFactory:       idFactory,
		configuration:   configuration,
		propertyBagName: "",
	}
}

// FunctionName returns the name of this function, as it will be shown in the generated source
func (c *PropertyConversionContext) FunctionName() string {
	return c.functionName
}

// Types returns the set of types available in this context
func (c *PropertyConversionContext) Types() astmodel.Types {
	return c.types
}

// IDFactory returns a reference to our identifier factory
func (c *PropertyConversionContext) IDFactory() astmodel.IdentifierFactory {
	return c.idFactory
}

// WithFunctionName returns a new context with the specified function name included
func (c *PropertyConversionContext) WithFunctionName(name string) *PropertyConversionContext {
	result := c.clone()
	result.functionName = name
	return result
}

// WithDirection returns a new context with the specified direction
func (c *PropertyConversionContext) WithDirection(dir Direction) *PropertyConversionContext {
	result := c.clone()
	result.direction = dir
	return result
}

// WithPropertyBag returns a new context with the specified property bag name included
func (c *PropertyConversionContext) WithPropertyBag(name string) *PropertyConversionContext {
	result := c.clone()
	result.propertyBagName = name
	return result
}

// ResolveType resolves a type that might be a type name into both the name and the actual
// type it references, returning true iff it was a TypeName that could be resolved
func (c *PropertyConversionContext) ResolveType(t astmodel.Type) (astmodel.TypeName, astmodel.Type, bool) {
	name, ok := astmodel.AsTypeName(t)
	if !ok {
		return astmodel.EmptyTypeName, nil, false
	}

	actualType, err := c.types.FullyResolve(name)
	if err != nil {
		return astmodel.EmptyTypeName, nil, false
	}

	return name, actualType, true
}

// PropertyBagName returns the name to use for a local property bag variable
func (c *PropertyConversionContext) PropertyBagName() string {
	return c.propertyBagName
}

// TypeRename looks up a type-rename for the specified type, returning the new name and nil if found, or empty string
// and an error if not.
func (c *PropertyConversionContext) TypeRename(name astmodel.TypeName) (string, error) {
	return c.configuration.LookupNameInNextVersion(name)
}

// clone returns a new independent copy of this context
func (c *PropertyConversionContext) clone() *PropertyConversionContext {
	return &PropertyConversionContext{
		types:           c.types,
		functionName:    c.functionName,
		direction:       c.direction,
		propertyBagName: c.propertyBagName,
		idFactory:       c.idFactory,
		configuration:   c.configuration,
	}
}
