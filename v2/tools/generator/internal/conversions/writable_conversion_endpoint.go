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

// WritableConversionEndpoint is an endpoint that includes a way to generate statements to write the final value
type WritableConversionEndpoint struct {
	endpoint *TypedConversionEndpoint
	// writer is a function that accepts an expression for the containing instance, another expression for the value to
	// write, and returns a set of statements for writing the value to that instance
	writer func(dst.Expr, dst.Expr) []dst.Stmt
	// description is a human-readable string for what is being written
	description string
}

var _ fmt.Stringer = &WritableConversionEndpoint{}

// NewWritableConversionEndpointWritingProperty creates a WritableConversionEndpoint for a specific property
func NewWritableConversionEndpointWritingProperty(
	propertyName astmodel.PropertyName,
	propertyType astmodel.Type,
) *WritableConversionEndpoint {
	name := string(propertyName)
	return &WritableConversionEndpoint{
		endpoint: NewTypedConversionEndpoint(propertyType, name),
		writer: func(destination dst.Expr, value dst.Expr) []dst.Stmt {
			return []dst.Stmt{
				astbuilder.SimpleAssignment(
					astbuilder.Selector(destination, name),
					value),
			}
		},
		description: fmt.Sprintf("write to property %s", propertyName),
	}
}

// NewWritableConversionEndpointWritingPropertyBagMember creates a WritableConversionEndpoint that writes an item into a
// property bag
func NewWritableConversionEndpointWritingPropertyBagMember(
	itemName string,
	itemType astmodel.Type,
) *WritableConversionEndpoint {
	return &WritableConversionEndpoint{
		endpoint: NewTypedConversionEndpoint(NewPropertyBagMemberType(itemType), itemName),
		writer: func(destination dst.Expr, value dst.Expr) []dst.Stmt {
			return []dst.Stmt{
				astbuilder.SimpleAssignment(
					astbuilder.Selector(destination, itemName),
					value),
			}
		},
		description: fmt.Sprintf("write %s to property bag", itemName),
	}
}

func (w *WritableConversionEndpoint) String() string {
	return w.description
}

// Write generates a series of statements to write the specified value to our destination endpoint
func (w *WritableConversionEndpoint) Write(destination dst.Expr, value dst.Expr) []dst.Stmt {
	return w.writer(destination, value)
}

// Endpoint provides access to the endpoint we write
func (w *WritableConversionEndpoint) Endpoint() *TypedConversionEndpoint {
	return w.endpoint
}
