/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import "github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"

// groupAccess provides access to a specific group's configuration.
type groupAccess[T any] struct {
	model    *ObjectModelConfiguration
	accessor func(*GroupConfiguration) *configurable[T]
}

/*
 * groupAccess
 */

func makeGroupAccess[T any](
	model *ObjectModelConfiguration,
	accessor func(g *GroupConfiguration) *configurable[T],
) groupAccess[T] {
	return groupAccess[T]{
		model:    model,
		accessor: accessor,
	}
}

func (a groupAccess[T]) withTypeOverride(
	accessor func(t *TypeConfiguration) *configurable[T],
) typeAccess[T] {
	ta := makeTypeAccess[T](a.model, accessor)
	ta.fallback = &a
	return ta
}

func (a *groupAccess[T]) Lookup(
	ref astmodel.InternalPackageReference,
) (T, bool) {
	var c *configurable[T]
	visitor := newSingleGroupConfigurationVisitor(
		ref,
		func(configuration *GroupConfiguration) error {
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

func (a *groupAccess[T]) VerifyConsumed() error {
	visitor := newEveryGroupConfigurationVisitor(
		func(configuration *GroupConfiguration) error {
			c := a.accessor(configuration)
			return c.VerifyConsumed()
		})
	return visitor.visit(a.model)
}

func (a *groupAccess[T]) MarkUnconsumed() error {
	visitor := newEveryGroupConfigurationVisitor(
		func(configuration *GroupConfiguration) error {
			c := a.accessor(configuration)
			c.MarkUnconsumed()
			return nil
		})

	return visitor.visit(a.model)
}
