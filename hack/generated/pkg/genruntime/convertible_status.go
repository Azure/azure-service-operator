/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

type ConvertibleStatus interface {
	// ConvertStatusTo will populate the passed Status by copying over all available information from this one
	ConvertStatusTo(destination ConvertibleStatus) error

	// ConvertStatusFrom will populate this status by copying over all available information from the passed one
	ConvertStatusFrom(source ConvertibleStatus) error
}
