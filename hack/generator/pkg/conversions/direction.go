/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// Direction specifies the direction of conversion we're implementing with this function
type Direction interface {
	// SelectString returns one of the provided strings, depending on the direction of conversion
	SelectString(from string, to string) string
	// SelectType returns one of the provided types, depending on the direction of conversion
	SelectType(from astmodel.Type, to astmodel.Type) astmodel.Type
}

var (
	// ConvertFrom indicates the conversion is from the passed 'other', populating the receiver with properties from the other
	ConvertFrom Direction = &ConvertFromDirection{}
	// ConvertTo indicates the conversion is to the passed 'other', populating the other with properties from the receiver
	ConvertTo Direction = &ConvertToDirection{}
)

type ConvertFromDirection struct{}

var _ Direction = &ConvertFromDirection{}

// SelectString returns the string for conversion FROM
func (dir *ConvertFromDirection) SelectString(fromString string, _ string) string {
	return fromString
}

// SelectType returns the type for conversion FROM
func (dir *ConvertFromDirection) SelectType(fromType astmodel.Type, _ astmodel.Type) astmodel.Type {
	return fromType
}

type ConvertToDirection struct{}

var _ Direction = &ConvertToDirection{}

// SelectString returns the string for conversion TO
func (dir *ConvertToDirection) SelectString(_ string, toValue string) string {
	return toValue
}

// SelectType returns the type for conversion TO
func (dir *ConvertToDirection) SelectType(_ astmodel.Type, toType astmodel.Type) astmodel.Type {
	return toType
}
