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

	// Property returns the property and true if the named property is found, nil and false otherwise
	Property(name PropertyName) (*PropertyDefinition, bool)
}

// AsPropertyContainer converts a type into a property container
func AsPropertyContainer(theType Type) (PropertyContainer, bool) {
	switch t := theType.(type) {
	case PropertyContainer:
		return t, true
	case MetaType:
		return AsPropertyContainer(t.Unwrap())
	default:
		return nil, false
	}
}
