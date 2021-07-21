/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

type ConvertibleStatus interface {
	// ConvertToStatus will populate the passed Status by copying over all available information from this one
	ConvertToStatus(destination ConvertibleStatus) error

	// ConvertFromStatus will populate this status by copying over all available information from the passed one
	ConvertFromStatus(source ConvertibleStatus) error
}

