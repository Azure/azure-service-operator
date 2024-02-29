/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import "github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"

// typeAccess provides access to a specific type's configuration.
type typeAccess[T any] struct {
	model    *ObjectModelConfiguration
	accessor func(*TypeConfiguration) *configurable[T]
	fallback *groupAccess[T]
}

/*
 * typeAccess
 */

// makeTypeAccess creates a new typeAccess[T] for the given model and accessor function
func makeTypeAccess[T any](
	model *ObjectModelConfiguration,
	accessor func(*TypeConfiguration) *configurable[T],
) typeAccess[T] {
	return typeAccess[T]{
		model:    model,
		accessor: accessor,
	}
}

func (a typeAccess[T]) withPropertyOverride(
	accessor func(p *PropertyConfiguration) *configurable[T],
) propertyAccess[T] {
	pa := makePropertyAccess[T](a.model, accessor)
	pa.fallback = &a
	return pa
}

// Lookup returns the configured value for the given type name
func (a *typeAccess[T]) Lookup(
	name astmodel.InternalTypeName,
) (T, bool) {
	result, ok := a.lookupCore(name)
	if ok {
		return result, true
	}

	if a.fallback != nil {
		return a.fallback.Lookup(name.InternalPackageReference())
	}

	return result, false
}

// lookupCore is the core implementation of lookup
func (a *typeAccess[T]) lookupCore(
	name astmodel.InternalTypeName,
) (T, bool) {
	var c *configurable[T]
	visitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
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
func (a *typeAccess[T]) VerifyConsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
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
func (a *typeAccess[T]) MarkUnconsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			c := a.accessor(configuration)
			c.MarkUnconsumed()
			return nil
		})

	return visitor.visit(a.model)
}
