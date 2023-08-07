/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
)

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

// AsTypeName unwraps any wrappers around the provided type and returns either the underlying TypeName and true, or a
// blank and false.
func AsTypeName(aType Type) (TypeName, bool) {
	if name, ok := aType.(TypeName); ok {
		return name, true
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsTypeName(wrapper.Unwrap())
	}

	return nil, false
}

// AsResourceType unwraps any wrappers around the provided type and returns either the underlying ResourceType and true,
// or a nil and false.
func AsResourceType(aType Type) (*ResourceType, bool) {
	if name, ok := aType.(*ResourceType); ok {
		return name, true
	}

	if wrapper, ok := aType.(MetaType); ok {
		return AsResourceType(wrapper.Unwrap())
	}

	return nil, false
}

func MustBeResourceType(aType Type) *ResourceType {
	result, ok := AsResourceType(aType)
	if !ok {
		panic(fmt.Sprintf("must have ResourceType, but received %T", aType))
	}

	return result
}

// Unwrap unwraps the type and returns the underlying unwrapped type
func Unwrap(aType Type) Type {
	if wrapper, ok := aType.(MetaType); ok {
		return Unwrap(wrapper.Unwrap())
	}

	return aType
}

// ExtractTypeName extracts a TypeName from the specified type if possible. This includes unwrapping
// MetaType's like ValidatedType as well as checking the element type of types such as ArrayType and MapType.
func ExtractTypeName(aType Type) (TypeName, bool) {
	if typeName, ok := AsTypeName(aType); ok {
		return typeName, ok
	}

	if arrayType, ok := AsArrayType(aType); ok {
		return ExtractTypeName(arrayType.Element())
	}

	if mapType, ok := AsMapType(aType); ok {
		return ExtractTypeName(mapType.ValueType())
	}

	return nil, false
}
