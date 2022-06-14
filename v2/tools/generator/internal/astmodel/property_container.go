/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// PropertyContainer is implemented by Types that contain properties
// Provides readonly access as we need to use a TypeVisitor for modifications to preserve type wrapping
type PropertyContainer interface {
	// Properties returns all the properties in this container
	Properties() ReadOnlyPropertySet

	// Property returns the property and true if the named property is found, nil and false otherwise
	Property(name PropertyName) (*PropertyDefinition, bool)

	// EmbeddedProperties returns all the embedded properties in this container
	EmbeddedProperties() []*PropertyDefinition
}

// AsPropertyContainer converts a type into a property container
// Only use this readonly access as we must use a TypeVisitor for modifications to preserve type wrapping
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
