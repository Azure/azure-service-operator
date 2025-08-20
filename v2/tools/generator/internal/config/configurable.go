/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"reflect"
	
	"github.com/rotisserie/eris"
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

// Merge merges the configuration from 'other' into this configurable.
// For primitive types, only overwrites if the current value is nil.
// For slices, appends new values to the end.
// For maps, adds new key-value pairs but errors if attempting to overwrite existing keys.
// Returns an error if there are conflicts during merging.
func (c *configurable[T]) Merge(other *configurable[T]) error {
	if other == nil || other.value == nil {
		return nil // Nothing to merge
	}

	if c.value == nil {
		// No existing value, just copy from other
		c.Set(*other.value)
		return nil
	}

	// Both have values - handle different types
	baseVal := reflect.ValueOf(c.value).Elem()
	otherVal := reflect.ValueOf(other.value).Elem()
	
	switch baseVal.Kind() {
	case reflect.Slice:
		// Generic slice handling - append other slice to base slice
		if otherVal.Len() > 0 {
			result := reflect.AppendSlice(baseVal, otherVal)
			baseVal.Set(result)
		}
		return nil

	case reflect.Map:
		// Generic map handling - merge keys from other map to base map
		if otherVal.Len() > 0 {
			if baseVal.IsNil() {
				baseVal.Set(reflect.MakeMap(baseVal.Type()))
			}
			for _, key := range otherVal.MapKeys() {
				newVal := otherVal.MapIndex(key)
				if existingVal := baseVal.MapIndex(key); existingVal.IsValid() {
					// Key exists, check for conflict
					if !reflect.DeepEqual(existingVal.Interface(), newVal.Interface()) {
						return eris.Errorf("conflict in %s for %s: key %v already exists with value %v, cannot overwrite with %v", 
							c.tag, c.scope, key.Interface(), existingVal.Interface(), newVal.Interface())
					}
				} else {
					// Key doesn't exist, add it
					baseVal.SetMapIndex(key, newVal)
				}
			}
		}
		return nil

	default:
		// For primitive types (string, bool) and enum types (PayloadType, ImportConfigMapMode, ReferenceType)
		// Both values are configured (non-nil), check if they are the same
		if any(*c.value) == any(*other.value) {
			return nil // Same value, no conflict
		}

		// Both have configured values that differ - conflict
		return eris.Errorf("conflict in %s for %s: base value %v cannot be overwritten with %v", c.tag, c.scope, *c.value, *other.value)
	}
}
