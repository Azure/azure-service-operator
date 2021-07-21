/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

type ConvertibleSpec interface {
	// ConvertToSpec will populate the passed Spec by copying over all available information from this one
	ConvertToSpec(destination ConvertibleSpec) error

	// ConvertFromSpec will populate this spec by copying over all available information from the passed one
	ConvertFromSpec(source ConvertibleSpec) error
}
