/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// A MetaType is a type wrapper that provide additional information/context about another type
type MetaType interface {
	// Unwrap returns the type contained within the wrapper type
	Unwrap() Type
}

// AsPrimitiveType unwraps any wrappers around the provided type and returns either the underlying
// PrimitiveType or nil
func AsPrimitiveType(aType Type) *PrimitiveType {
	if primitive, ok := aType.(*PrimitiveType); ok {
		return primitive
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsPrimitiveType(wrapper.Unwrap())
	}

	return nil
}

// AsObjectType unwraps any wrappers around the provided type and returns either the underlying
// ObjectType or nil
func AsObjectType(aType Type) *ObjectType {
	if obj, ok := aType.(*ObjectType); ok {
		return obj
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsObjectType(wrapper.Unwrap())
	}

	return nil
}

// AsArrayType unwraps any wrappers the provided type and returns either the underlying ArrayType or nil
func AsArrayType(aType Type) *ArrayType {
	if arr, ok := aType.(*ArrayType); ok {
		return arr
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsArrayType(wrapper.Unwrap())
	}

	return nil
}

// AsMapType unwraps any wrappers around the provided type and returns either the underlying MapType or nil
func AsMapType(aType Type) *MapType {
	if mt, ok := aType.(*MapType); ok {
		return mt
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsMapType(wrapper.Unwrap())
	}

	return nil
}

// AsOptionalType unwraps any wrappers around the provided type and returns either the underlying OptionalType or nil
func AsOptionalType(aType Type) *OptionalType {
	if opt, ok := aType.(*OptionalType); ok {
		return opt
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsOptionalType(wrapper.Unwrap())
	}

	return nil
}
