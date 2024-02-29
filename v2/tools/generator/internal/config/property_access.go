/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import "github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"

// propertyAccess provides access to a specific property's configuration.
type propertyAccess[T any] struct {
	model    *ObjectModelConfiguration
	accessor func(*PropertyConfiguration) *configurable[T]
	fallback *typeAccess[T]
}

// makePropertyAccess creates a new propertyAccess[T] for the given model and accessor function
func makePropertyAccess[T any](
	model *ObjectModelConfiguration,
	accessor func(*PropertyConfiguration,
	) *configurable[T],
) propertyAccess[T] {
	return propertyAccess[T]{
		model:    model,
		accessor: accessor,
	}
}

// Lookup returns the configured value for the given type name and property name
func (a *propertyAccess[T]) Lookup(
	name astmodel.InternalTypeName,
	property astmodel.PropertyName,
) (T, bool) {
	result, ok := a.lookupCore(name, property)
	if ok {
		return result, true
	}

	if a.fallback != nil {
		return a.fallback.Lookup(name)
	}

	return result, false
}

// lookupCore is the core implementation of lookup
func (a *propertyAccess[T]) lookupCore(
	name astmodel.InternalTypeName,
	property astmodel.PropertyName,
) (T, bool) {
	var c *configurable[T]
	visitor := newSinglePropertyConfigurationVisitor(
		name,
		property,
		func(configuration *PropertyConfiguration) error {
			c = a.accessor(configuration)
			return nil
		})

	err := visitor.visit(a.model)
	if err != nil {
		// Something went wrong; this is unexpected and shouldn't happen
		panic(err)
	}

	if c == nil {
		var zero T
		return zero, false
	}

	return c.Lookup()
}

// VerifyConsumed ensures that all configured values have been consumed
func (a *propertyAccess[T]) VerifyConsumed() error {
	visitor := newEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			c := a.accessor(configuration)
			return c.VerifyConsumed()
		})

	err := visitor.visit(a.model)
	if err != nil {
		return err
	}

	if a.fallback != nil {
		return a.fallback.VerifyConsumed()
	}

	return nil
}

// MarkUnconsumed marks all configured values as unconsumed
func (a *propertyAccess[T]) MarkUnconsumed() error {
	visitor := newEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			c := a.accessor(configuration)
			c.MarkUnconsumed()
			return nil
		})

	return visitor.visit(a.model)
}
