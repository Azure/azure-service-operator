/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// Direction specifies the direction of conversion we're implementing with this function
type Direction interface {
	// SelectString returns one of the provided strings, depending on the direction of conversion
	SelectString(from string, to string) string
	// SelectType returns one of the provided types, depending on the direction of conversion
	SelectType(from astmodel.Type, to astmodel.Type) astmodel.Type
	// SelectExpr returns on one of the provided expressions, depending on the direction of conversion
	SelectExpr(from dst.Expr, to dst.Expr) dst.Expr
	// WhenFrom will run the specified function only if the direction is "From", returning the current direction for chaining
	WhenFrom(fn func()) Direction
	// WhenTo will run the specified function only if the direction is "To", returning the current direction for chaining
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

// SelectString returns the string for conversion FROM
func (dir *ConvertFromDirection) SelectString(fromString string, _ string) string {
	return fromString
}

// SelectType returns the type for conversion FROM
func (dir *ConvertFromDirection) SelectType(fromType astmodel.Type, _ astmodel.Type) astmodel.Type {
	return fromType
}

// SelectExpr returns the expression for conversion FROM
func (dir *ConvertFromDirection) SelectExpr(fromExpr dst.Expr, _ dst.Expr) dst.Expr {
	return fromExpr
}

// WhenFrom will run the supplied function, returning this FROM direction for chaining
func (dir *ConvertFromDirection) WhenFrom(fn func()) Direction {
	fn()
	return dir
}

// WhenTo will skip the supplied function, returning this FROM direction for chaining
func (dir *ConvertFromDirection) WhenTo(_ func()) Direction {
	// Nothing
	return dir
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

// SelectExpr returns the expression for conversion TO
func (dir *ConvertToDirection) SelectExpr(_ dst.Expr, toExpr dst.Expr) dst.Expr {
	return toExpr
}

// WhenFrom will skip the supplied function, returning this TO direction for chaining
func (dir *ConvertToDirection) WhenFrom(_ func()) Direction {
	return dir
}

// WhenTo will run the supplied function, returning this TO direction for chaining
func (dir *ConvertToDirection) WhenTo(fn func()) Direction {
	fn()
	return dir
}
