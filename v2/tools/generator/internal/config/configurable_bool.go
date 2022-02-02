/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

type configurableBool struct {
	value    *bool
	consumed bool
}

// Read returns the value configured, if any.
// The second result indicates whether a value is available.
// If present, the value is flagged as consumed.
func (cb *configurableBool) read() (bool, bool) {
	if cb.value != nil {
		cb.consumed = true
		return *cb.value, true
	}

	return false, false
}

// Write sets the value configured and marks it as unconsumed
func (cb *configurableBool) write(v bool) {
	cb.value = &v
	cb.consumed = false
}

// isUnconsumed returns true if we have a configured value that hasn't been consumed
func (cb *configurableBool) isUnconsumed() bool {
	return cb.value != nil && !cb.consumed
}
