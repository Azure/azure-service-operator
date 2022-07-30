/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

// LocatableResource represents a resource with a location.
// TODO: The generated types should impl this. Currently only ResourceGroup has this interface and it's written manually
type LocatableResource interface {
	Location() string
}
