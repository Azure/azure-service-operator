/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// FunctionContainer is implemented by Types that contain functions
// Provides readonly access as we need to use a TypeVisitor for modifications to preserve type wrapping
type FunctionContainer interface {
	// Functions returns all the function implementations
	// A sorted slice is returned to preserve immutability and provide determinism
	Functions() []Function

	// HasFunctionWithName determines if this resource has a function with the given name
	HasFunctionWithName(name string) bool
}

// AsFunctionContainer converts a type into a function container
// Only use this readonly access as we must use a TypeVisitor for modifications to preserve type wrapping
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
