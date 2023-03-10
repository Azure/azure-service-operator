/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

// ImportableResource is implemented by any resource that can be imported into the operator via asoctl
type ImportableResource interface {
	// InitializeSpec is used to set up the spec for a resource based on the STATUS that's currently available
	// (In other words, using ConvertibleStatus is not an error here.)
	InitializeSpec(status ConvertibleStatus) error
}
