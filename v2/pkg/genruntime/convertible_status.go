/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

// ConvertibleStatus is implemented by status types to allow conversion among the different versions of a given status
//
// Why do we need both directions of conversion? See ConvertibleSpec for details.
//
type ConvertibleStatus interface {
	// ConvertStatusTo will populate the passed Status by copying over all available information from this one
	ConvertStatusTo(destination ConvertibleStatus) error

	// ConvertStatusFrom will populate this status by copying over all available information from the passed one
	ConvertStatusFrom(source ConvertibleStatus) error
}
