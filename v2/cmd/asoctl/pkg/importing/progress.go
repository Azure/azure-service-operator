/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

// Progress is an interface for reporting the progress of an import operation.
type Progress interface {
	// AddPending adds a number of pending steps
	AddPending(pending int)

	// Completed marks a number of steps as completed
	Completed(completed int)

	// Create creates a new Progress instance that is nested within the current one.
	// This is useful for reporting progress of a part of the operation.
	Create(name string) Progress
}
