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

type ResourceImporterFactory interface {
	// CreateARMFactory creates a new factory specifically for ARM resources
	CreateARMFactory(client *azruntime.Pipeline, serviceConfig cloud.ServiceConfiguration) ARMResourceImporter
	// Scheme returns the scheme used by the factory
	Scheme() *runtime.Scheme
}

// a resourceImporterFactory is used to create resourceImporter instances for specific resources
type resourceImporterFactory struct {
	// scheme is a reference to the scheme used by asoctl
	scheme *runtime.Scheme
}

var _ ResourceImporterFactory = &resourceImporterFactory{}

// newResourceImporterFactory creates a new factory with the scheme baked in
func newResourceImporterFactory(scheme *runtime.Scheme) ResourceImporterFactory {
	return &resourceImporterFactory{
		scheme: scheme,
	}
}

func (f *resourceImporterFactory) CreateARMFactory(
	client *azruntime.Pipeline,
	serviceConfig cloud.ServiceConfiguration,
) ARMResourceImporter {
	return &armResourceImporter{
		resourceImporterFactory: *f,
		armClient:               client,
		armConfig:               serviceConfig,
	}
}

func (f *resourceImporterFactory) Scheme() *runtime.Scheme {
	return f.scheme
}

func (f *resourceImporterFactory) createBlankObjectFromGVK(gvk schema.GroupVersionKind) (runtime.Object, error) {
	obj, err := f.scheme.New(gvk)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	obj.GetObjectKind().SetGroupVersionKind(gvk)
	return obj, nil
}

func (f *resourceImporterFactory) selectVersionFromGK(gk schema.GroupKind) (schema.GroupVersionKind, error) {
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

func (*resourceImporterFactory) selectLatestVersion(
	gk schema.GroupKind,
	knownVersions []schema.GroupVersion,
) schema.GroupVersionKind {
	// Sort the versions the same way we do in the generator, for consistency
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
