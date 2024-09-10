/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importreporter

// Interface is an interface for reporting the importreporter of an import operation.
type Interface interface {
	// AddPending adds a number of pending steps
	AddPending(pending int)

	// Completed marks a number of steps as completed
	Completed(completed int)

	// Create creates a new Interface instance that is nested within the current one.
	// This is useful for reporting importreporter of a part of the operation.
	Create(name string) Interface
}
