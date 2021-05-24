/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import "github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"

// StorageConversionContext captures additional supporting information that may be needed when a
// storage conversion factory creates a conversion
type StorageConversionContext struct {
	// types is a map of all known type definitions, used to resolve TypeNames to actual types
	types astmodel.Types
	// functionName is the name of the function we're currently generating
	functionName string
	// knownLocals is a reference to our set of local variables
	// (Pointer because it's a reference type, not because it's optional)
	knownLocals *astmodel.KnownLocalsSet
}

// NewStorageConversionContext creates a new instance of a StorageConversionContext
func NewStorageConversionContext(types astmodel.Types, idFactory astmodel.IdentifierFactory) *StorageConversionContext {
	return &StorageConversionContext{
		types:       types,
		knownLocals: astmodel.NewKnownLocalsSet(idFactory),
	}
}

//TODO: Remove this?
func (c *StorageConversionContext) FunctionName() string {
	return c.functionName
}

//TODO: Remove this?
// Types returns the set of types available in this context
func (c *StorageConversionContext) Types() astmodel.Types {
	return c.types
}

// WithFunctionName returns a new context with the specified function name included
func (c *StorageConversionContext) WithFunctionName(name string) *StorageConversionContext {
	result := c.clone()
	result.functionName = name
	return result
}

// WithKnownLocals returns a new context with the specified known locals set referenced
func (c *StorageConversionContext) WithKnownLocals(knownLocals *astmodel.KnownLocalsSet) *StorageConversionContext {
	result := c.clone()
	result.knownLocals = knownLocals
	return result
}

// NestedContext returns a new context with a cloned knownLocals so that nested blocks
// can declare and reuse locals independently of each other.
func (c *StorageConversionContext) NestedContext() *StorageConversionContext {
	result := c.clone()
	return result
}

// ResolveType resolves a type that might be a type name into both the name and the actual
// type it references, returning true iff it was a TypeName that could be resolved
func (c *StorageConversionContext) ResolveType(t astmodel.Type) (astmodel.TypeName, astmodel.Type, bool) {
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
func (c *StorageConversionContext) TryCreateLocal(local string) bool {
	if c.knownLocals.HasName(local) {
		return false
	}

	c.knownLocals.Add(local)
	return true
}

// clone returns a new independent copy of this context
func (c *StorageConversionContext) clone() *StorageConversionContext {
	return &StorageConversionContext{
		types:        c.types,
		functionName: c.functionName,
		knownLocals:  c.knownLocals.Clone(),
	}
}
