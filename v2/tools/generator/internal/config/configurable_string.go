/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

type configurableString struct {
	value    *string
	consumed bool
}

// Read returns the value configured and true, if configured; otherwise returns empty string and false.
// If present, the value is flagged as consumed.
func (cs *configurableString) read() (string, bool) {
	if cs.value != nil {
		cs.consumed = true
		return *cs.value, true
	}

	return "", false
}

// Write sets the value configured and marks it as unconsumed
func (cs *configurableString) write(v string) {
	cs.value = &v
	cs.consumed = false
}

// isUnconsumed returns true if we have a configured value that hasn't been consumed
func (cs *configurableString) isUnconsumed() bool {
	return cs.value != nil && !cs.consumed
}
