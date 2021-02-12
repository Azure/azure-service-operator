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
}

// NewStorageConversionContext creates a new instance of a StorageConversionContext
func NewStorageConversionContext(types Types) *StorageConversionContext {
	return &StorageConversionContext{
		types: types,
	}
}

// ResolveEnum takes a Type and resolves it into the name and definition of an enumeration,
// returning true if found or false if not.
func (c *StorageConversionContext) ResolveEnum(t Type) (TypeName, *EnumType, bool) {
	name, ok := AsTypeName(t)
	if !ok {
		// Source is not identified by name
		return TypeName{}, nil, false
	}

	aType, err := c.types.FullyResolve(name)
	if err != nil {
		// Can't resolve source
		return TypeName{}, nil, false
	}

	enumType, isEnum := AsEnumType(aType)
	if !isEnum {
		// Source is not an enum
		return TypeName{}, nil, false
	}

	return name, enumType, true
}
