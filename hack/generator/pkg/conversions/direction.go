/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

// Direction specifies the direction of conversion we're implementing with this function
type Direction interface {
	SelectString(from string, to string) string
	WhenFrom(fn func()) Direction
	WhenTo(fn func()) Direction
}

var (
	// ConvertFrom indicates the conversion is from the passed 'other', populating the receiver with properties from the other
	ConvertFrom Direction = &ConvertFromDirection{}
	// ConvertTo indicates the conversion is to the passed 'other', populating the other with properties from the receiver
	ConvertTo Direction = &ConvertToDirection{}
)

type ConvertFromDirection struct{}

var _ Direction = &ConvertFromDirection{}

func (dir *ConvertFromDirection) SelectString(fromValue string, _ string) string {
	return fromValue
}

func (dir *ConvertFromDirection) WhenFrom(fn func()) Direction {
	fn()
	return dir
}

func (dir *ConvertFromDirection) WhenTo(_ func()) Direction {
	// Nothing
	return dir
}

type ConvertToDirection struct{}

var _ Direction = &ConvertToDirection{}

func (dir *ConvertToDirection) SelectString(_ string, toValue string) string {
	return toValue
}

func (dir *ConvertToDirection) WhenFrom(_ func()) Direction {
	// Nothing
	return dir
}

func (dir *ConvertToDirection) WhenTo(fn func()) Direction {
	fn()
	return dir
}
