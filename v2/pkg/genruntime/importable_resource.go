/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

// ImportableResource is implemented by any resource that can be imported into the operator via asoctl
type ImportableResource interface {
	// InitializeSpec initializes the Spec of the resource from the provided Status.
	// (In other words, using ConvertibleStatus is not an error here.)
	InitializeSpec(status ConvertibleStatus) error
}
