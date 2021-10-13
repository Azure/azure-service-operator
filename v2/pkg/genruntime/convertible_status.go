/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
)

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

// GetVersionedStatus returns a versioned status for the provided resource; the original API version used when the
// resource was first created is used to identify the version to return
func GetVersionedStatus(metaObject MetaObject, scheme *runtime.Scheme) (ConvertibleStatus, error) {
	rsrc, err := NewEmptyVersionedResource(metaObject, scheme)
	if err != nil {
		return nil, errors.Wrap(err, "creating new empty versioned status")
	}

	if rsrc == nil {
		// No conversion needed
		return metaObject.GetStatus(), nil
	}

	// Get a blank status and populate it
	status := rsrc.GetStatus()
	err = status.ConvertStatusFrom(metaObject.GetStatus())
	if err != nil {
		return nil, errors.Wrapf(err, "failed conversion of status")
	}

	return status, nil
}
