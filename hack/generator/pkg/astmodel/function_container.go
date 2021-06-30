/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// FunctionContainer is implemented by Types that contain functions
// Can't include the withers for modification until we have generics
type FunctionContainer interface {
	// Functions returns all the function implementations
	// A sorted slice is returned to preserve immutability and provide determinism
	Functions() []Function

	// HasFunctionWithName determines if this resource has a function with the given name
	HasFunctionWithName(name string) bool
}

// AsFunctionContainer converts a type into a function container
func AsFunctionContainer(theType Type) (FunctionContainer, bool) {
	switch t := theType.(type) {
	case FunctionContainer:
		return t, true
	case MetaType:
		return AsFunctionContainer(t.Unwrap())
	default:
		return nil, false
	}
}
