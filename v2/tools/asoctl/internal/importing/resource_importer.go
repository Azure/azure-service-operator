/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

// ResourceImporter is the entry point for importing resources.
// Factory methods here provide ways to instantiate importers for different kinds of resources.
type ResourceImporter struct {
	scheme    *runtime.Scheme                 // a reference to the scheme used by asoctl
	client    *genericarmclient.GenericClient // Client to use when talking to ARM
	pending   []ImportableResource            // A set of importers that are pending import
	completed map[string]ImportableResource   // A set of importers that have been imported
}

// NewResourceImporter creates a new factory with the scheme baked in
func NewResourceImporter(
	scheme *runtime.Scheme,
	client *genericarmclient.GenericClient,
) *ResourceImporter {
	return &ResourceImporter{
		scheme:    scheme,
		client:    client,
		completed: make(map[string]ImportableResource),
	}
}

// Add adds an importer to the list of resources to import.
func (ri *ResourceImporter) Add(importer ImportableResource) {
	ri.pending = append(ri.pending, importer)
}

// AddARMID adds an ARM ID to the list of resources to import.
func (ri *ResourceImporter) AddARMID(armID string) {
	importer := NewImportableARMResource(armID, ri.client, ri.scheme)
	ri.Add(importer)
}

// Import imports all the resources that have been added to the importer.
// Partial results are returned even in the case of an error.
func (ri *ResourceImporter) Import(ctx context.Context) (*ResourceImportResult, error) {
	var errs []error
	for len(ri.pending) > 0 {
		// Remove the first pending importer
		importer := ri.pending[0]
		ri.pending = ri.pending[1:]

		// If we've already handled this resource, skip it
		if _, ok := ri.completed[importer.Name()]; ok {
			continue
		}

		thisResource := len(ri.completed) + 1
		pendingResources := len(ri.pending)
		klog.Infof("Importing %d/%d: %s", thisResource, thisResource+pendingResources, importer.Name())

		// Import it
		pending, err := importer.Import(ctx)
		if err != nil {
			var notImportable NotImportableError
			if errors.As(err, &notImportable) {
				// This is a resource we don't know how to import, but that's ok (details will already have been logged)
				continue
			}

			errs = append(errs, errors.Wrapf(err, "failed during import of %s", importer.Name()))
			continue
		}

		ri.completed[importer.Name()] = importer
		ri.pending = append(ri.pending, pending...)
	}

	// Now we've imported everything, return the resources
	resources := make([]genruntime.MetaObject, 0, len(ri.completed))
	for _, importer := range ri.completed {
		resources = append(resources, importer.Resource())
	}

	return &ResourceImportResult{
		resources: resources,
	}, kerrors.NewAggregate(errs)
}
