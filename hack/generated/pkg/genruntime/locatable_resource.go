/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

// TODO: The generated types should impl this
type LocatableResource interface {
	Location() string
}
