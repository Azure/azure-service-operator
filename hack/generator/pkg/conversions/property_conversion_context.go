/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import "github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"

// PropertyConversionContext captures additional supporting information that may be needed when a
// storage conversion factory creates a conversion
type PropertyConversionContext struct {
	// types is a map of all known type definitions, used to resolve TypeNames to actual types
	types astmodel.Types
	// functionName is the name of the function we're currently generating
	functionName string
	// direction is the direction of the conversion we're generating
	direction Direction
	// knownLocals is a reference to our set of local variables
	// (Pointer because it's a reference type, not because it's optional)
	knownLocals *astmodel.KnownLocalsSet
	// idFactory is used for generating method names
	idFactory astmodel.IdentifierFactory
}

// NewPropertyConversionContext creates a new instance of a PropertyConversionContext
func NewPropertyConversionContext(types astmodel.Types, idFactory astmodel.IdentifierFactory) *PropertyConversionContext {
	return &PropertyConversionContext{
		types:       types,
		idFactory:   idFactory,
		knownLocals: astmodel.NewKnownLocalsSet(idFactory),
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

// WithKnownLocals returns a new context with the specified known locals set referenced
func (c *PropertyConversionContext) WithKnownLocals(knownLocals *astmodel.KnownLocalsSet) *PropertyConversionContext {
	result := c.clone()
	result.knownLocals = knownLocals
	return result
}

// WithDirection returns a new context with the specified direction
func (c *PropertyConversionContext) WithDirection(dir Direction) *PropertyConversionContext {
	result := c.clone()
	result.direction = dir
	return result
}

// NestedContext returns a new context with a cloned knownLocals so that nested blocks
// can declare and reuse locals independently of each other.
func (c *PropertyConversionContext) NestedContext() *PropertyConversionContext {
	result := c.clone()
	return result
}

// ResolveType resolves a type that might be a type name into both the name and the actual
// type it references, returning true iff it was a TypeName that could be resolved
func (c *PropertyConversionContext) ResolveType(t astmodel.Type) (astmodel.TypeName, astmodel.Type, bool) {
	name, ok := astmodel.AsTypeName(t)
	if !ok {
		return astmodel.TypeName{}, nil, false
	}

	actualType, err := c.types.FullyResolve(name)
	if err != nil {
		return astmodel.TypeName{}, nil, false
	}

	return name, actualType, true
}

// TryCreateLocal returns true if the specified local was successfully created, false if it already existed
func (c *PropertyConversionContext) TryCreateLocal(local string) bool {
	if c.knownLocals.HasName(local) {
		return false
	}

	c.knownLocals.Add(local)
	return true
}

// clone returns a new independent copy of this context
func (c *PropertyConversionContext) clone() *PropertyConversionContext {
	return &PropertyConversionContext{
		types:        c.types,
		functionName: c.functionName,
		direction:    c.direction,
		idFactory:    c.idFactory,
		knownLocals:  c.knownLocals.Clone(),
	}
}
