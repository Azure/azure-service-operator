/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
)

// PropertyReference is a fully qualified reference to where a property may be found
type PropertyReference struct {
	declaringType InternalTypeName // which type declares the property
	property      PropertyName     // the name of the property
}

// EmptyPropertyReference is a convenience constant for when we have no value
var EmptyPropertyReference = PropertyReference{}

// MakePropertyReference creates a new property reference identifying both the containing type and the property name.
func MakePropertyReference(declaringType InternalTypeName, property PropertyName) PropertyReference {
	return PropertyReference{
		declaringType: declaringType,
		property:      property,
	}
}

// DeclaringType returns the type name of the type declaring the property
func (ref PropertyReference) DeclaringType() InternalTypeName {
	return ref.declaringType
}

// Property returns the actual name of the property
func (ref PropertyReference) Property() PropertyName {
	return ref.property
}

// IsEmpty returns true if the reference is empty, false otherwise.
func (ref PropertyReference) IsEmpty() bool {
	return ref == EmptyPropertyReference
}

// String returns a string representation of this property reference
func (ref PropertyReference) String() string {
	declaringType := ref.declaringType
	return fmt.Sprintf(
		"%s/%s.%s",
		declaringType.InternalPackageReference().FolderPath(),
		declaringType.Name(),
		ref.property)
}
