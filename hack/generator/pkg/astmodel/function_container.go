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
}
