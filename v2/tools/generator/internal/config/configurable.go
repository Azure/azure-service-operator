/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"github.com/pkg/errors"
)

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

// Lookup returns the value configured, or an error if not configured.
func (c *configurable[T]) Lookup() (T, error) {
	if v, ok := c.read(); ok {
		return v, nil
	}

	msg := fmt.Sprintf("%s not specified for %s", c.tag, c.scope)
	return *new(T), NewNotConfiguredError(msg)
}

// VerifyConsumed returns an error if the value is configured but not consumed.
func (c *configurable[T]) VerifyConsumed() error {
	if c.isUnconsumed() {
		return errors.Errorf("%s specified for %s but not consumed", c.tag, c.scope)
	}

	return nil
}

// MarkUnconsumed resets the consumption flag so that the value can be reused
func (c *configurable[T]) MarkUnconsumed() {
	c.markUnconsumed()
}

// Set sets the value configured (used in tests)
func (c *configurable[T]) Set(v T) {
	c.write(v)
}

// Read returns the value configured and true, if configured; otherwise returns default(T) and false.
// If present, the value is flagged as consumed.
func (c *configurable[T]) read() (T, bool) {
	if c.value != nil {
		c.consumed = true
		return *c.value, true
	}

	return *new(T), false
}

// Write sets the value configured and marks it as unconsumed
func (c *configurable[T]) write(v T) {
	c.value = &v
	c.consumed = false
}

// isUnconsumed returns true if we have a configured value that hasn't been consumed
func (c *configurable[T]) isUnconsumed() bool {
	return c.value != nil && !c.consumed
}

// markUnconsumed marks this as unconsumed
func (c *configurable[T]) markUnconsumed() {
	c.consumed = false
}
