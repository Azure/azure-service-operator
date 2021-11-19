/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
)

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

// GetVersionedSpec returns a versioned spec for the provided resource; the original API version used when the
// resource was first created is used to identify the version to return
func GetVersionedSpec(metaObject MetaObject, scheme *runtime.Scheme) (ConvertibleSpec, error) {
	rsrc, err := NewEmptyVersionedResource(metaObject, scheme)
	if err != nil {
		return nil, errors.Wrap(err, "getting versioned spec")
	}

	if rsrc.GetObjectKind().GroupVersionKind() == metaObject.GetObjectKind().GroupVersionKind() {
		// No conversion needed, empty resource is the same GVK that we already have
		return metaObject.GetSpec(), nil
	}

	// Get a blank spec and populate it
	spec := rsrc.GetSpec()
	err = spec.ConvertSpecFrom(metaObject.GetSpec())
	if err != nil {
		return nil, errors.Wrap(err, "failed conversion of spec")
	}

	return spec, nil
}

// GetVersionedARMSpec returns a spec object ready for serialization to ARM; the original API version used when the
// resource was first created is used to create the appropriate version for submission.
func GetVersionedARMSpec(metaObject MetaObject, resolved ConvertToARMResolvedDetails, scheme *runtime.Scheme) (interface{}, error) {
	spec, err := GetVersionedSpec(metaObject, scheme)
	if err != nil {
		return nil, errors.Wrap(err, "creating ARM spec")
	}

	converter, ok := spec.(ToARMConverter)
	if !ok {
		return nil, errors.Errorf("expected %T to implement genruntime.ToARMConverter", spec)
	}

	result, err := converter.ConvertToARM(resolved)
	if err != nil {
		return nil, errors.Wrapf(err, "creating ARM spec from %T", spec)
	}

	return result, err
}
