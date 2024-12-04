/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// importFactory is a helper struct that's used to create things.
type importFactory struct {
	scheme *runtime.Scheme
}

func newImportFactory(
	scheme *runtime.Scheme,
) *importFactory {
	return &importFactory{
		scheme: scheme,
	}
}

// createBlankObjectFromGVK creates a blank object of from a given GVK.
func (f *importFactory) createBlankObjectFromGVK(gvk schema.GroupVersionKind) (runtime.Object, error) {
	obj, err := f.scheme.New(gvk)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	obj.GetObjectKind().SetGroupVersionKind(gvk)
	return obj, nil
}

// selectVersionFromGK selects the latest version of a given GroupKind.
// The latest stable version will be selected if it exists, otherwise the latest preview version will be selected.
func (f *importFactory) selectVersionFromGK(gk schema.GroupKind) (schema.GroupVersionKind, error) {
	knownVersions := f.scheme.VersionsForGroupKind(gk)
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
		obj, err := f.createBlankObjectFromGVK(gvk)
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
