/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// ReadableConversionEndpoint is an endpoint that includes a way to generate statements to write the final value
type WritableConversionEndpoint struct {
	endpoint *TypedConversionEndpoint
	// writer is a function that accepts an expression for the containing instance, another expression for the value to
	// write, and returns a set of statements for writing the value to that instance
	writer func(dst.Expr, dst.Expr) []dst.Stmt
	// description is a human readable string for what is being written
	description string
}

var _ fmt.Stringer = WritableConversionEndpoint{}

// MakeWritableConversionEndpointForProperty creates a WritableConversionEndpoint for a specific property
func MakeWritableConversionEndpointForProperty(
	prop *astmodel.PropertyDefinition,
	knownLocals *astmodel.KnownLocalsSet,
) WritableConversionEndpoint {
	propertyName := string(prop.PropertyName())
	return WritableConversionEndpoint{
		endpoint: NewStorageConversionEndpoint(prop.PropertyType(), propertyName, knownLocals),
		writer: func(destination dst.Expr, value dst.Expr) []dst.Stmt {
			return []dst.Stmt{
				astbuilder.SimpleAssignment(
					astbuilder.Selector(destination, propertyName),
					token.ASSIGN,
					value),
			}
		},
		description: fmt.Sprintf("write property %s", propertyName),
	}
}

func (w WritableConversionEndpoint) String() string {
	return w.description
}

// Write generates a series of statements to write the specified value to our destination endpoint
func (w WritableConversionEndpoint) Write(destination dst.Expr, value dst.Expr) []dst.Stmt {
	return w.writer(destination, value)
}

// Endpoint() provides access to the endpoint we write
func (w WritableConversionEndpoint) Endpoint() *TypedConversionEndpoint {
	return w.endpoint
}

// CreateWritableEndpoints creates a map of all the writable endpoints found on a type
func CreateWritableEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) map[string]WritableConversionEndpoint {
	result := make(map[string]WritableConversionEndpoint)

	propContainer, ok := astmodel.AsPropertyContainer(instance)
	if ok {
		for _, prop := range propContainer.Properties() {
			endpoint := MakeWritableConversionEndpointForProperty(prop, knownLocals)
			result[string(prop.PropertyName())] = endpoint
		}
	}

	return result
}
