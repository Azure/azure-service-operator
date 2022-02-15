/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

// configurableStringSlice represents a string slice that may be configured.
// Includes tracking for whether we consume the configured slice or not, allowing us to flag unnecessary configuration.
type configurableStringSlice struct {
	value    []string
	consumed bool
}

// Read returns the value configured and true, if configured; otherwise returns nil and false.
// If present, the value is flagged as consumed.
func (cs *configurableStringSlice) read() ([]string, bool) {
	if cs.value != nil {
		cs.consumed = true
		return cs.value, true
	}

	return nil, false
}

// Write sets the value configured and marks it as unconsumed
func (cs *configurableStringSlice) write(v []string) {
	cs.value = v
	cs.consumed = false
}

// isUnconsumed returns true if we have a configured value that hasn't been consumed
func (cs *configurableStringSlice) isUnconsumed() bool {
	return cs.value != nil && !cs.consumed
}
