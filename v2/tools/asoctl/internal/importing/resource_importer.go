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
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

// ResourceImporter is the entry point for importing resources.
// Factory methods here provide ways to instantiate importers for different kinds of resources.
type ResourceImporter struct {
	scheme    *runtime.Scheme                 // a reference to the scheme used by asoctl
	client    *genericarmclient.GenericClient // Client to use when talking to ARM
	pending   map[string]ImportableResource   // A set of importers that are pending import
	completed map[string]ImportableResource   // A set of importers that have been successfully imported
	queue     []string                        // Queue of names of resources to import (so we do things in a reasonable order)
}

// NewResourceImporter creates a new factory with the scheme baked in
func NewResourceImporter(
	scheme *runtime.Scheme,
	client *genericarmclient.GenericClient,
) *ResourceImporter {
	return &ResourceImporter{
		scheme:    scheme,
		client:    client,
		pending:   make(map[string]ImportableResource),
		completed: make(map[string]ImportableResource),
	}
}

// Add adds an importer to the list of resources to import.
func (ri *ResourceImporter) Add(importer ImportableResource) {
	// If we've already handled this resource, skip it
	if _, ok := ri.completed[importer.Name()]; ok {
		return
	}

	// If we're already pending import of this resource, skip it
	if _, ok := ri.pending[importer.Name()]; ok {
		return
	}

	// Add it to our map and our queue
	ri.queue = append(ri.queue, importer.Name())
	ri.pending[importer.Name()] = importer
}

// AddARMID adds an ARM ID to the list of resources to import.
func (ri *ResourceImporter) AddARMID(armID string) error {
	importer, err := newImportableARMResource(armID, nil /* no owner */, ri.client, ri.scheme)
	if err != nil {
		return errors.Wrapf(err, "failed to create importer for %q", armID)
	}

	ri.Add(importer)
	return nil
}

// Import imports all the resources that have been added to the importer.
// Partial results are returned even in the case of an error.
func (ri *ResourceImporter) Import(ctx context.Context) (*ResourceImportResult, error) {
	var errs []error
	processed := 0
	for len(ri.pending) > 0 {
		// Remove the first pending importer
		head := ri.queue[0]
		ri.queue = ri.queue[1:]

		importer := ri.pending[head]
		delete(ri.pending, head)

		processed++
		pendingCount := len(ri.pending)

		// If we've already handled this resource, skip it
		if _, ok := ri.completed[importer.Name()]; ok {
			klog.Infof(
				"Already imported %d/%d: %s",
				processed,
				processed+pendingCount,
				importer.Name())
			continue
		}

		klog.Infof(
			"Importing %d/%d: %s",
			processed,
			processed+pendingCount,
			importer.Name())

		// Import it
		pending, err := importer.Import(ctx)
		if err != nil {
			var notImportable NotImportableError
			if errors.As(err, &notImportable) {
				klog.Infof(err.Error())
				continue
			}

			errs = append(errs, errors.Wrapf(err, "failed during import of %s", importer.Name()))
			continue
		}

		ri.completed[importer.Name()] = importer
		start := len(ri.pending)
		for _, p := range pending {
			ri.Add(p)
		}

		if len(ri.pending) > start {
			klog.Infof("Queued %d new resources", len(ri.pending)-start)
		}
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
