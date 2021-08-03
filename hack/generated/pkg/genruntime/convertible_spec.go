/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

// ConvertibleSpec is implemented by Spec types to allow conversion among the different versions of a given spec
//
// Why do we need both directions of conversion?
//
// Each version of a resource is in a different package, so the implementations of this interface will necessarily be
// referencing types from other packages. If we tried to use an interface with a single method, we'd inevitably end up
// with circular package references:
//
//      +----------------+                    +----------------+
//      |       v1       |                    |       v2       |
//      |   PersonSpec   | --- import v2 ---> |   PersonSpec   |
//      |                |                    |                |
//      | ConvertTo()    | <--- import v1 --- | ConvertTo()    |
//      +----------------+                    +----------------+
//
// Instead, we have to have support for both directions, so that we can always operate from one side of the package
// reference chain:
//
//      +----------------+                    +----------------+
//      |       v1       |                    |       v2       |
//      |   PersonSpec   |                    |   PersonSpec   |
//      |                |                    |                |
//      | ConvertTo()    | --- import v2 ---> |                |
//      | ConvertFrom()  |                    |                |
//      +----------------+                    +----------------+
//
type ConvertibleSpec interface {
	// ConvertSpecTo will populate the passed Spec by copying over all available information from this one
	ConvertSpecTo(destination ConvertibleSpec) error

	// ConvertSpecFrom will populate this spec by copying over all available information from the passed one
	ConvertSpecFrom(source ConvertibleSpec) error
}
