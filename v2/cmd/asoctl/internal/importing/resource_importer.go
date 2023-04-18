/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"fmt"
	"sync"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"golang.org/x/sync/errgroup"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

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
	log       logr.Logger                     // Logger to use for logging
	progress  *mpb.Progress                   // Progress bar to use for showing progress
	report    *resourceImportReport           // Report to summarise the import
}

// NewResourceImporter creates a new factory with the scheme baked in
func NewResourceImporter(
	scheme *runtime.Scheme,
	client *genericarmclient.GenericClient,
	log logr.Logger,
	progress *mpb.Progress) *ResourceImporter {
	return &ResourceImporter{
		scheme:    scheme,
		client:    client,
		pending:   make(map[string]ImportableResource),
		completed: make(map[string]ImportableResource),
		log:       log,
		progress:  progress,
		report:    newResourceImportReport(),
	}
}

// Add adds an importer to the list of resources to import.
func (ri *ResourceImporter) Add(importer ImportableResource) {
	// Lock while we're modifying the maps
	ri.lock.Lock()
	defer ri.lock.Unlock()

	ri.addImpl(importer)
}

// AddARMID adds an ARM ID to the list of resources to import.
func (ri *ResourceImporter) AddARMID(armID string) error {
	importer, err := NewImportableARMResource(armID, nil /* no owner */, ri.client, ri.scheme)
	if err != nil {
		return errors.Wrapf(err, "failed to create importer for %q", armID)
	}

	ri.Add(importer)
	return nil
}

// addImpl is the internal implementation of Add that assumes the lock is already held
// This allows us to call it from other methods that already have the lock
func (ri *ResourceImporter) addImpl(importer ImportableResource) {
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
func (ri *ResourceImporter) Import(
	ctx context.Context,
) (*ResourceImportResult, error) {
	var processed int64 = 0

	globalBar := ri.progress.AddBar(
		0, // zero total because we don't know how much work will be needed,
		mpb.PrependDecorators(
			decor.Name("Import Azure Resources", decor.WCSyncSpaceR)),
		mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})))
	globalBar.SetPriority(-1) // must display at the top

	// Why the nested loop?
	// We can potentially trigger all the pending resources listed in ri.queue to be imported in parallel
	// and end up with nothing on the queue because they're still running.
	// The outer loop gives us a last look at the queue and allows us to keep going if there are new items
	// requiring import.
	var errs []error
	for len(ri.queue) > 0 {
		var eg errgroup.Group

		// Limit concurrent processing, so we don't do too much at once
		// The exact value isn't too important
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

			eg.Go(func() error {
				err := ri.ImportResource(ctx, rsrc)

				// Update our global progress too
				processed++
				globalBar.SetTotal(processed+int64(len(ri.queue)), false)
				globalBar.SetCurrent(processed)

				return err
			})
		}

		// Wait for all the running imports to complete
		err := eg.Wait()
		if err != nil {
			errs = append(errs, err)
		}
	}

	globalBar.SetTotal(processed, true)

	// Now we've imported everything, return the resources
	// We do this even if there's an error so that we can return partial results
	resources := make([]genruntime.MetaObject, 0, len(ri.completed))
	for _, importer := range ri.completed {
		resources = append(resources, importer.Resource())
	}

	ri.report.WriteToLog(ri.log)

	return &ResourceImportResult{
		resources: resources,
	}, kerrors.NewAggregate(errs)
}

func (ri *ResourceImporter) ImportResource(ctx context.Context, rsrc ImportableResource) error {
	// Import it
	gk := rsrc.GroupKind()
	name := fmt.Sprintf("%s %s", gk, rsrc.Name())
	bar := ri.progress.AddBar(
		0, // zero total because we don't know how much work will be needed
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(name, decor.WCSyncSpaceR)),
		mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})))

	defer func() {
		// Ensure the progress bar is always updated when we're done
		// We're done with this resource, so remove it from the progress bar
		bar.Increment()
		bar.SetTotal(bar.Current(), true)
	}()

	pending, err := rsrc.Import(ctx, bar)
	if err != nil {
		var skipped *ImportSkippedError
		if errors.As(err, &skipped) {
			ri.log.V(1).Info(
				"Skipped",
				"kind", gk,
				"name", rsrc.Name(),
				"because", skipped.Because)
			ri.report.AddSkippedImport(rsrc, skipped.Because)
			return nil
		}

		ri.log.Error(err,
			"Failed",
			"kind", gk,
			"name", rsrc.Name())

		ri.report.AddFailedImport(rsrc, err.Error())

		// Don't need to wrap the error, it's already been logged (we don't want to log it twice)
		return errors.Errorf("failed during import of %s", rsrc.Name())
	}

	ri.log.Info(
		"Imported",
		"kind", gk,
		"name", rsrc.Name())

	ri.report.AddSuccessfulImport(rsrc)

	ri.Complete(rsrc, pending)
	return nil
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
	for _, p := range pending {
		ri.addImpl(p)
	}
}
