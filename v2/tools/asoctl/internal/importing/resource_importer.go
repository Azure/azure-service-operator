/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"k8s.io/apimachinery/pkg/runtime"
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
	lock      sync.Mutex                      // Lock to protect the above maps
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
	// Lock while we're modifying the maps
	ri.lock.Lock()
	defer ri.lock.Unlock()

	ri.addCore(importer)
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

// addCore is the internal implementation of Add that assumes the lock is already held
// This allows us to call it from other methods that already have the lock
func (ri *ResourceImporter) addCore(importer ImportableResource) {
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

// Import imports all the resources that have been added to the importer.
// Partial results are returned even in the case of an error.
func (ri *ResourceImporter) Import(ctx context.Context) (*ResourceImportResult, error) {
	processed := 0

	// Why the nested loop?
	// We can potentially trigger all the pending resources listed in ri.queue to be imported in parallel
	// and end up with nothing on the queue because they're still running.
	// The outer loop gives us a last look at the queue and allows us to keep going if there are new items
	// requiring import.
	for len(ri.queue) > 0 {
		var eg errgroup.Group
		eg.SetLimit(8)

		for {
			if ctx.Err() != nil {
				// Context cancelled, time to stop
				return nil, errors.Wrapf(ctx.Err(), "cancelled while importing")
			}

			rsrc, ok := ri.DequeueResource()
			if !ok {
				// No more resources to import
				break
			}

			processed++
			pendingCount := len(ri.pending)

			klog.Infof(
				"Importing %d/%d: %s",
				processed,
				processed+pendingCount,
				rsrc.Name())

			eg.Go(func() error {
				// Import it
				pending, err := rsrc.Import(ctx)
				if err != nil {
					var notImportable NotImportableError
					if errors.As(err, &notImportable) {
						klog.Infof(err.Error())
						return nil
					}

					return errors.Wrapf(err, "failed during import of %s", rsrc.Name())
				}

				ri.Complete(rsrc, pending)
				return nil
			})
		}

		// Wait for all the running imports to complete
		err := eg.Wait()
		if err != nil {
			return nil, err
		}
	}

	// Now we've imported everything, return the resources
	// We do this even if there's an error so that we can return partial results
	resources := make([]genruntime.MetaObject, 0, len(ri.completed))
	for _, importer := range ri.completed {
		resources = append(resources, importer.Resource())
	}

	return &ResourceImportResult{
		resources: resources,
	}, nil
}

// DequeueResource returns the next resource to import.
func (ri *ResourceImporter) DequeueResource() (ImportableResource, bool) {
	// Lock while we're modifying the maps
	ri.lock.Lock()
	defer ri.lock.Unlock()

	if len(ri.pending) == 0 {
		return nil, false
	}

	// Remove the first pending importer
	head := ri.queue[0]
	ri.queue = ri.queue[1:]

	importer := ri.pending[head]
	delete(ri.pending, head)

	return importer, true
}

func (ri *ResourceImporter) Complete(importer ImportableResource, pending []ImportableResource) {
	// Lock while we're modifying the maps
	ri.lock.Lock()
	defer ri.lock.Unlock()

	// Add it to our map and our queue
	ri.completed[importer.Name()] = importer
	start := len(ri.pending)
	for _, p := range pending {
		ri.addCore(p)
	}

	if len(ri.pending) > start {
		klog.Infof("Queued %d new resources", len(ri.pending)-start)
	}
}
