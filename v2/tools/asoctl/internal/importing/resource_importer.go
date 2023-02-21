/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/versions"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"

	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ResourceImporter is the entry point for importing resources.
// Factory methods here provide ways to instantiate importers for different kinds of resources.
type ResourceImporter struct {
	// scheme is a reference to the scheme used by asoctl
	scheme *runtime.Scheme
}

// NewResourceImporter creates a new factory with the scheme baked in
func NewResourceImporter(scheme *runtime.Scheme) *ResourceImporter {
	return &ResourceImporter{
		scheme: scheme,
	}
}

// CreateARMImporter creates an ARMResourceImporter with the given client and service configuration.
func (f *ResourceImporter) CreateARMImporter(
	client *azruntime.Pipeline,
	serviceConfig cloud.ServiceConfiguration,
) *ARMResourceImporter {
	return &ARMResourceImporter{
		ResourceImporter:     *f,
		client:               client,
		serviceConfiguration: serviceConfig,
	}
}

// Scheme returns the scheme used by the importer.
func (f *ResourceImporter) Scheme() *runtime.Scheme {
	return f.scheme
}

// createBlankObjectFromGVK is a helper function to create a blank object of from a given GVK.
func (f *ResourceImporter) createBlankObjectFromGVK(gvk schema.GroupVersionKind) (runtime.Object, error) {
	obj, err := f.scheme.New(gvk)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	obj.GetObjectKind().SetGroupVersionKind(gvk)
	return obj, nil
}

// selectVersionFromGK is a helper function to select the latest version of a given GroupKind.
// The latest stable version will be selected if it exists, otherwise the latest preview version will be selected.
func (f *ResourceImporter) selectVersionFromGK(gk schema.GroupKind) (schema.GroupVersionKind, error) {
	knownVersions := f.scheme.VersionsForGroupKind(gk)
	if len(knownVersions) == 0 {
		return schema.GroupVersionKind{},
			errors.Errorf(
				"no known versions for Group %s, Kind %s",
				gk.Group,
				gk.Kind)
	}

	return f.selectLatestVersion(gk, knownVersions), nil
}

// selectLatestVersion is a helper function to select the latest from a slice of GroupVersions
// The latest stable version will be selected if it exists, otherwise the latest preview version will be selected.
func (*ResourceImporter) selectLatestVersion(
	gk schema.GroupKind,
	knownVersions []schema.GroupVersion,
) schema.GroupVersionKind {
	// Sort the versions the same way we do in the generator, for consistency.
	// The versions.Compare() function used to sort the versions is the one exported by the generator.
	slices.SortFunc(
		knownVersions,
		func(left schema.GroupVersion, right schema.GroupVersion) bool {
			return versions.Compare(left.Version, right.Version)
		})

	isStorageVersion := func(version string) bool {
		return strings.HasSuffix(version, "storage")
	}

	// Ideally we want to find the latest stable version, but if there isn't one we'll take the latest preview.
	// Preview versions might introduce odd behaviour, so we err on the side of caution.
	// Storage versions need to be skipped though, as they don't have a fixed OriginalVersion()
	var previewVersion schema.GroupVersion
	var stableVersion schema.GroupVersion
	for _, gv := range knownVersions {
		if isStorageVersion(gv.Version) {
			// Skip storage versions
			continue
		}

		if versions.IsPreview(gv.Version) {
			previewVersion = gv
		} else {
			stableVersion = gv
		}
	}

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
