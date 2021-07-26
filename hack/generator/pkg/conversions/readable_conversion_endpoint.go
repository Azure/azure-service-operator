/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"fmt"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// ReadableConversionEndpoint is an endpoint that includes a way to generate an expression for the source value
type ReadableConversionEndpoint struct {
	endpoint *TypedConversionEndpoint
	// reader is a function that accepts an expression for the containing instance and returns an expression for reading
	// the value from that instance
	reader func(expr dst.Expr) dst.Expr
	// description is a human readable string for what is being read
	description string
}

var _ fmt.Stringer = ReadableConversionEndpoint{}

// MakeReadableConversionEndpointForProperty creates a ReadableConversionEndpoint for a specific property
func MakeReadableConversionEndpointForProperty(
	prop *astmodel.PropertyDefinition,
	knownLocals *astmodel.KnownLocalsSet,
) ReadableConversionEndpoint {
	propertyName := string(prop.PropertyName())
	return ReadableConversionEndpoint{
		endpoint: NewStorageConversionEndpoint(prop.PropertyType(), propertyName, knownLocals),
		reader: func(source dst.Expr) dst.Expr {
			return astbuilder.Selector(source, propertyName)
		},
		description: fmt.Sprintf("read from property %s", propertyName),
	}
}

// MakeReadableConversionEndpointForValueFunction creates a ReadableConversionEndpoint for a specific value function
func MakeReadableConversionEndpointForValueFunction(
	fn astmodel.ValueFunction,
	knownLocals *astmodel.KnownLocalsSet,
) ReadableConversionEndpoint {
	functionName := fn.Name()
	return ReadableConversionEndpoint{
		endpoint: NewStorageConversionEndpoint(fn.ReturnType(), functionName, knownLocals),
		reader: func(source dst.Expr) dst.Expr {
			return astbuilder.CallExpr(source, functionName)
		},
		description: fmt.Sprintf("call function %s()", functionName),
	}
}

// String returns a human readable description of the endpoint
func (r ReadableConversionEndpoint) String() string {
	return r.description
}

// Read generates an expression to read our endpoint
func (r ReadableConversionEndpoint) Read(expr dst.Expr) dst.Expr {
	return r.reader(expr)
}

// Endpoint provides access to the end point we read
func (r ReadableConversionEndpoint) Endpoint() *TypedConversionEndpoint {
	return r.endpoint
}
