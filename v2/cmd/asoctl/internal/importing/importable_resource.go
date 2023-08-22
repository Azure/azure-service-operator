/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vbauerster/mpb/v8"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// ImportableResource is an interface that wraps a Kubernetes resource that can be imported.
// Different implementations of this interface will be used for different types of resources.
type ImportableResource interface {
	// GroupKind returns the GroupKind of the resource being imported.
	// (may be empty if the GK can't be determined)
	GroupKind() schema.GroupKind

	// Name is a human readable identifier for this resource
	Name() string

	// Id is a unique identifier for this resource.
	// The Id of a resource unique within the import operation; the easiest way to achive this is
	// to make it globally unique.
	Id() string

	// Resource returns the actual resource that has been imported.
	// Only available after the import is complete (nil otherwise).
	Resource() genruntime.MetaObject

	// Import does the actual import, updating the Spec on the wrapped resource.
	// ctx allows for cancellation of the import.
	// If there are any additional resources that also need to be imported, they should be returned.
	Import(ctx context.Context, bar *mpb.Bar) ([]ImportableResource, error)
}

// importableResource is a core of common data and support methods for implementing ImportableResource
type importableResource struct {
	scheme *runtime.Scheme
}

// createBlankObjectFromGVK is a helper function to create a blank object of from a given GVK.
func (i *importableResource) createBlankObjectFromGVK(gvk schema.GroupVersionKind) (runtime.Object, error) {
	obj, err := i.scheme.New(gvk)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	obj.GetObjectKind().SetGroupVersionKind(gvk)
	return obj, nil
}

// selectVersionFromGK is a helper function to select the latest version of a given GroupKind.
// The latest stable version will be selected if it exists, otherwise the latest preview version will be selected.
func (i *importableResource) selectVersionFromGK(gk schema.GroupKind) (schema.GroupVersionKind, error) {
	knownVersions := i.scheme.VersionsForGroupKind(gk)
	if len(knownVersions) == 0 {
		return schema.GroupVersionKind{},
			errors.Errorf(
				"no known versions for Group %s, Kind %s",
				gk.Group,
				gk.Kind)
	}

	// Scan for the GVK that implements genruntime.ImportableResource
	// We expect there to be exactly one
	var result *schema.GroupVersionKind
	for _, gv := range knownVersions {
		gvk := gk.WithVersion(gv.Version)
		obj, err := i.createBlankObjectFromGVK(gvk)
		if err != nil {
			return schema.GroupVersionKind{}, errors.Wrapf(err, "unable to create blank resource for GVK %s", gvk)
		}

		if _, ok := obj.(genruntime.ImportableResource); ok {
			if result != nil {
				return schema.GroupVersionKind{},
					errors.Errorf(
						"multiple known versions for Group %s, Kind %s implement genruntime.ImportableResource",
						gk.Group,
						gk.Kind)
			}

			result = &gvk
		}
	}

	if result == nil {
		return schema.GroupVersionKind{},
			errors.Errorf(
				"no known versions for Group %s, Kind %s implement genruntime.ImportableResource",
				gk.Group,
				gk.Kind)
	}

	return *result, nil
}
