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
// PrimitiveType and true, or nil and false.
func AsPrimitiveType(aType Type) (*PrimitiveType, bool) {
	if primitive, ok := aType.(*PrimitiveType); ok {
		return primitive, true
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsPrimitiveType(wrapper.Unwrap())
	}

	return nil, false
}

// AsObjectType unwraps any wrappers around the provided type and returns either the underlying
// ObjectType and true, or nil and false.
func AsObjectType(aType Type) (*ObjectType, bool) {
	if obj, ok := aType.(*ObjectType); ok {
		return obj, true
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsObjectType(wrapper.Unwrap())
	}

	return nil, false
}

// AsArrayType unwraps any wrappers around the provided type and returns either the underlying
// ArrayType and true, or nil and false.
func AsArrayType(aType Type) (*ArrayType, bool) {
	if arr, ok := aType.(*ArrayType); ok {
		return arr, true
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsArrayType(wrapper.Unwrap())
	}

	return nil, false
}

// AsMapType unwraps any wrappers around the provided type and returns either the underlying
// MapType and true, or nil and false.
func AsMapType(aType Type) (*MapType, bool) {
	if mt, ok := aType.(*MapType); ok {
		return mt, true
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsMapType(wrapper.Unwrap())
	}

	return nil, false
}

// AsOptionalType unwraps any wrappers around the provided type and returns either the underlying
// OptionalType and true, or nil and false.
func AsOptionalType(aType Type) (*OptionalType, bool) {
	if opt, ok := aType.(*OptionalType); ok {
		return opt, true
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsOptionalType(wrapper.Unwrap())
	}

	return nil, false
}

// AsEnumType unwraps any wrappers around the provided type and returns either the underlying EnumType and true, or nil and false.
func AsEnumType(aType Type) (*EnumType, bool) {
	if enm, ok := aType.(*EnumType); ok {
		return enm, true
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsEnumType(wrapper.Unwrap())
	}

	return nil, false
}

// AsTypeName unwraps any wrappers around the provided type and returns either the underlying TypeName and true, or a blank and false.
func AsTypeName(aType Type) (TypeName, bool) {
	if name, ok := aType.(TypeName); ok {
		return name, true
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsTypeName(wrapper.Unwrap())
	}

	return TypeName{}, false
}
