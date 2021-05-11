/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// StorageConversionContext captures additional supporting information that may be needed when a
// storage conversion factory creates a conversion
type StorageConversionContext struct {
	// types is a map of all known type definitions, used to resolve TypeNames to actual types
	types Types
	// functionName is the name of the function we're currently generating
	functionName string
	// knownLocals is a reference to our set of local variables
	knownLocals *KnownLocalsSet
}

// NewStorageConversionContext creates a new instance of a StorageConversionContext
func NewStorageConversionContext(types Types) *StorageConversionContext {
	return &StorageConversionContext{
		types:       types,
		knownLocals: nil,
	}
}

// WithFunctionName returns a new context with the specified function name included
func (c *StorageConversionContext) WithFunctionName(name string) *StorageConversionContext {
	result := c.clone()
	result.functionName = name
	return result
}

// WithKnownLocals returns a new context with the specified known locals set referenced
func (c *StorageConversionContext) WithKnownLocals(knownLocals *KnownLocalsSet) *StorageConversionContext {
	result := c.clone()
	result.knownLocals = knownLocals
	return result
}

// NestedContext returns a nested context with independent knownLocals so that independent blocks
// can declare and reuse locals independently
func (c *StorageConversionContext) NestedContext() *StorageConversionContext {
	result := c.clone()
	result.knownLocals = c.knownLocals.Clone()
	return result
}

// ResolveType resolves a type that might be a type name into both the name and the actual
// type it references, returning true iff it was a TypeName that could be resolved
func (c *StorageConversionContext) ResolveType(t Type) (TypeName, Type, bool) {
	name, ok := AsTypeName(t)
	if !ok {
		return TypeName{}, nil, false
	}

	actualType, err := c.types.FullyResolve(name)
	if err != nil {
		return TypeName{}, nil, false
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
		knownLocals:  c.knownLocals,
	}
}
