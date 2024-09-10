/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importreporter

type discardReporter struct{}

var _ Interface = &barReporter{}

// NewDiscard creates a new import reporter that discards all updates.
func NewDiscard() Interface {
	return &discardReporter{}
}

// AddPending does nothing
func (*discardReporter) AddPending(_ int) {}

// Completed does nothing
func (*discardReporter) Completed(_ int) {}

// Create returns a discarding reporter
func (d *discardReporter) Create(_ string) Interface {
	return d
}
