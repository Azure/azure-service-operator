/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"fmt"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ReadableConversionEndpoint is an endpoint that includes a way to generate an expression for the source value
type ReadableConversionEndpoint struct {
	endpoint *TypedConversionEndpoint
	// reader is a function that accepts an expression for the containing instance and returns an expression for reading
	// the value from that instance
	reader func(expr dst.Expr) dst.Expr
	// description is a human-readable string for what is being read
	description string
}

var _ fmt.Stringer = &ReadableConversionEndpoint{}

// NewReadableConversionEndpointReadingProperty creates a ReadableConversionEndpoint that reads a value from a specific
// property
func NewReadableConversionEndpointReadingProperty(
	propertyName astmodel.PropertyName,
	propertyType astmodel.Type,
) *ReadableConversionEndpoint {
	name := string(propertyName)
	return &ReadableConversionEndpoint{
		endpoint: NewTypedConversionEndpoint(propertyType, name),
		reader: func(source dst.Expr) dst.Expr {
			return astbuilder.Selector(source, name)
		},
		description: fmt.Sprintf("read from property %s", name),
	}
}

// NewReadableConversionEndpointReadingValueFunction creates a ReadableConversionEndpoint that reads a value from a
// specific single valued function
func NewReadableConversionEndpointReadingValueFunction(
	fnName string,
	fnReturnType astmodel.Type,
) *ReadableConversionEndpoint {
	return &ReadableConversionEndpoint{
		endpoint: NewTypedConversionEndpoint(fnReturnType, fnName),
		reader: func(source dst.Expr) dst.Expr {
			return astbuilder.CallExpr(source, fnName)
		},
		description: fmt.Sprintf("call function %s()", fnName),
	}
}

// NewReadableConversionEndpointReadingPropertyBagMember creates a ReadableConversionEndpoint that reads an item from a
// property bag.
func NewReadableConversionEndpointReadingPropertyBagMember(
	itemName string,
	itemType astmodel.Type,
) *ReadableConversionEndpoint {
	return &ReadableConversionEndpoint{
		endpoint: NewTypedConversionEndpoint(NewPropertyBagMemberType(itemType), itemName),
		// We don't supply a reader function because we don't read the value from the source instance when dealing with
		// a property bag member; instead we read it from a property bag that's stashed in a local variable.
		// See AssignFromBagItem() for more details
		reader:      nil,
		description: fmt.Sprintf("read %s from property bag", itemName),
	}
}

// Name returns the name of the underlying endpoint
func (r *ReadableConversionEndpoint) Name() string {
	return r.endpoint.Name()
}

// String returns a human-readable description of the endpoint
func (r *ReadableConversionEndpoint) String() string {
	return r.description
}

// Read generates an expression to read our endpoint
func (r *ReadableConversionEndpoint) Read(expr dst.Expr) dst.Expr {
	if r.reader == nil {
		// If we don't have an expression to use, just return the original
		// (this can happen if this endpoint represents a source that doesn't directly read from our source instance)
		return expr
	}

	return r.reader(expr)
}

// Endpoint provides access to the end point we read
func (r *ReadableConversionEndpoint) Endpoint() *TypedConversionEndpoint {
	return r.endpoint
}
