/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

// configurable represents a value that may be configured.
// Includes tracking for whether we consume the configured value or not, allowing us to flag unnecessary configuration
type configurable[T any] struct {
	value    *T
	consumed bool
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
