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
}

// NewStorageConversionContext creates a new instance of a StorageConversionContext
func NewStorageConversionContext(types Types) *StorageConversionContext {
	return &StorageConversionContext{
		types: types,
	}
}

func (c *StorageConversionContext) WithFunctionName(name string) *StorageConversionContext {
	result := NewStorageConversionContext(c.types)
	result.functionName = name
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
