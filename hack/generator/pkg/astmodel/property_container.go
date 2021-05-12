/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// PropertyContainer is implemented by Types that contain properties
// Can't include the withers for modification until we have generics
type PropertyContainer interface {
	// Properties returns all the properties from this container
	// A sorted slice is returned to preserve immutability and provide determinism
	Properties() []*PropertyDefinition
}
