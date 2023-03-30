/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/names"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"

	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/versions"
)

// ImportableResource is an interface that wraps a Kubernetes resource that can be imported.
// Different implementations of this interface will be used for different types of resources.
type ImportableResource interface {
	// Name is a unique identifier for this resource
	Name() string

	// Resource returns the actual resource that has been imported.
	// Only available after the import is complete (nil otherwise).
	Resource() genruntime.MetaObject

	// Import does the actual import, updating the Spec on the wrapped resource.
	// ctx allows for cancellation of the import.
	// If there are any additional resources that also need to be imported, they should be returned.
	Import(ctx context.Context) ([]ImportableResource, error)
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

	return i.selectLatestVersion(gk, knownVersions), nil
}

// selectLatestVersion is a helper function to select the latest from a slice of GroupVersions
// The latest stable version will be selected if it exists, otherwise the latest preview version will be selected.
func (i *importableResource) selectLatestVersion(
	gk schema.GroupKind,
	knownVersions []schema.GroupVersion,
) schema.GroupVersionKind {
	// Sort the versions the same way we do in the generator, for consistency.
	// In the generator, we compare `astmodel.PackageReference` paths using versions.Compare() to order them
	// consistently. Here, we're comparing `schema.GroupVersion` values instead, but we want consistent results.
	slices.SortFunc(
		knownVersions,
		func(left schema.GroupVersion, right schema.GroupVersion) bool {
			return versions.Compare(left.Version, right.Version)
		})

	// Ideally we want to find the latest stable version, but if there isn't one we'll take the latest preview.
	// Preview versions might introduce odd behaviour, so we err on the side of caution.
	// Storage versions need to be skipped though, as they don't have a fixed OriginalVersion()
	// This logic is similar to the way the storage/hub version is selected in the generator, but here we are
	// looking for an API version (not a storage version), and we're dealing with a slice of GroupVersions
	// instead of a slice of PackageReferences.
	var previewVersion schema.GroupVersion
	var stableVersion schema.GroupVersion
	for _, gv := range knownVersions {
		// IsStorageVersion() is exported from the generator to ensure we use the same logic here
		if names.IsStorageVersion(gv.Version) {
			// Skip storage versions
			continue
		}

		if versions.IsPreview(gv.Version) {
			previewVersion = gv
		} else {
			stableVersion = gv
		}
	}

	// If we found a stable version, use that, otherwise use the preview version
	var result schema.GroupVersionKind
	if !stableVersion.Empty() {
		result = stableVersion.WithKind(gk.Kind)
	} else {
		result = previewVersion.WithKind(gk.Kind)
	}

	// Only need to log version as Group and Kind will have been logged elsewhere
	klog.V(3).Infof("Version: %s", result.Version)

	return result
}
