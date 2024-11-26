/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import "github.com/rotisserie/eris"

// configurable represents a value that may be configured.
// Includes tracking for whether we consume the configured value or not, allowing us to flag unnecessary configuration
type configurable[T any] struct {
	value    *T
	consumed bool
	tag      string
	scope    string
}

// makeConfigurable creates a new configurable[T] with the given tag and scope
func makeConfigurable[T any](tag string, scope string) configurable[T] {
	return configurable[T]{
		tag:   tag,
		scope: scope,
	}
}

// Lookup returns the value configured and true, or false if not configured.
func (c *configurable[T]) Lookup() (T, bool) {
	return c.read()
}

// VerifyConsumed returns an error if the value is configured but not consumed.
func (c *configurable[T]) VerifyConsumed() error {
	if c.isUnconsumed() {
		return eris.Errorf("%s specified for %s but not consumed", c.tag, c.scope)
	}

	return nil
}

// MarkUnconsumed resets the consumption flag so that the value can be reused
func (c *configurable[T]) MarkUnconsumed() {
	c.consumed = false
}

// Set writes the value configured and marks it as unconsumed
func (c *configurable[T]) Set(v T) {
	c.value = &v
	c.consumed = false
}

// Read returns the value configured and true, if configured; otherwise returns default(T) and false.
// If present, the value is flagged as consumed.
func (c *configurable[T]) read() (T, bool) {
	if c.value != nil {
		c.consumed = true
		return *c.value, true
	}

	var zero T
	return zero, false
}

// isUnconsumed returns true if we have a configured value that hasn't been consumed
func (c *configurable[T]) isUnconsumed() bool {
	return c.value != nil && !c.consumed
}
