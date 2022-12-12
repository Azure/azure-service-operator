/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// EnumValue captures a single value of the enumeration
type EnumValue struct {
	// Identifier is a Go identifier for the value
	Identifier string
	// Value is the actual value expected by ARM
	Value string
}

// MakeEnumValue makes a new EnumValue with the given identifier and value
func MakeEnumValue(id string, value string) EnumValue {
	return EnumValue{
		Identifier: id,
		Value:      value,
	}
}

// Equals tests to see if the passed EnumValue has the same name and value
func (value *EnumValue) Equals(v *EnumValue) bool {
	if value == v {
		return true
	}

	return value.Identifier == v.Identifier && value.Value == v.Value
}

// String implements fmt.Stringer for debugging purposes
func (value *EnumValue) String() string {
	return value.Identifier
}
