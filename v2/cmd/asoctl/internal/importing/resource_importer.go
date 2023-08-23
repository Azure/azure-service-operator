/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

// ResourceImporter is the entry point for importing resources.
// Factory methods here provide ways to instantiate importers for different kinds of resources.
type ResourceImporter struct {
	scheme    *runtime.Scheme                 // a reference to the scheme used by asoctl
	client    *genericarmclient.GenericClient // Client to use when talking to ARM
	resources []ImportableResource            // A slice of resources to be imported

	imported map[string]ImportableResource // A set of importers that have been successfully imported
	log      logr.Logger                   // Logger to use for logging
	progress *mpb.Progress                 // Progress bar to use for showing progress
}

type ImportResourceResult struct {
	resource ImportableResource
	pending  []ImportableResource
	err      error
}

type progressDelta struct {
	complete int
	total    int
}

// NewResourceImporter creates a new factory with the scheme baked in
func NewResourceImporter(
	scheme *runtime.Scheme,
	client *genericarmclient.GenericClient,
	log logr.Logger,
	progress *mpb.Progress) *ResourceImporter {
	return &ResourceImporter{
		scheme:   scheme,
		client:   client,
		imported: make(map[string]ImportableResource),
		log:      log,
		progress: progress,
	}
}

// Add adds an importer to the list of resources to import.
func (ri *ResourceImporter) Add(importer ImportableResource) {
	ri.resources = append(ri.resources, importer)
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

// Import imports all the resources that have been added to the importer.
// Partial results are returned even in the case of an error.
func (ri *ResourceImporter) Import(
	ctx context.Context,
) (*ResourceImportResult, error) {

	workers := 1
	candidates := make(chan ImportableResource)  // candidates that need to be deduped
	pending := make(chan ImportableResource)     // importers that are pending import
	completed := make(chan ImportResourceResult) // importers that have been executed successfully
	done := make(chan struct{})                  // signal that we're done

	progress := ri.createProgressBar("Import Azure Resources", done)

	// Dedupe candidates so we import each distinct resource only once
	go ri.queueUniqueImporters(candidates, pending, progress)

	// Create workers to run the import
	for i := 0; i < workers; i++ {
		go ri.importWorker(ctx, pending, completed, progress)
	}

	// Collate the results
	go ri.collateResults(completed, candidates, progress)

	// Set up by adding our initial resources
	for _, rsrc := range ri.resources {
		candidates <- rsrc
		progress <- progressDelta{total: 1}
	}

	// Wait while everything runs
	<-done

	// Check for an abort, and return the error if so
	if ctx.Err() != nil {
		ri.log.Error(ctx.Err(), "Cancelling import.")
		return nil, ctx.Err()
	}

	// Now we've imported everything, return the resources
	// We do this even if there's an error so that we can return partial results
	resources := make([]genruntime.MetaObject, 0, len(ri.imported))
	for _, importer := range ri.imported {
		resources = append(resources, importer.Resource())
	}

	return &ResourceImportResult{
		resources: resources,
	}, nil
}

// queueUniqueImporters reads from the candicates channel, putting each importer onto pending exactly once.
// This ensures each distinct resource is only imported once, regardless of how many times we
// encounter it. We also proactively buffer the pending channel to avoid blocking. Any fixed size channel
// would risk deadlock for a sufficiently large resource graph.
func (ri *ResourceImporter) queueUniqueImporters(
	candidates chan ImportableResource,
	pending chan ImportableResource,
	progress chan progressDelta,
) {
	seen := set.Make[string]()
	var queue []ImportableResource
	var current ImportableResource = nil
	for {
		// Dequeue from our internal buffer if needed
		if current == nil && len(queue) > 0 {
			current = queue[0]
			queue = queue[1:]
		}

		// If we have a current importable to send, use the pending queue
		var upstream chan ImportableResource = nil
		if current != nil {
			upstream = pending
		}

		select {
		case rsrc := <-candidates:
			if seen.Contains(rsrc.Id()) {
				// We've already seen this resource (we've already queued it for import)
				// So remove it from our count of work to be done
				ri.log.V(2).Info("Skipping duplicate import", "resource", rsrc.Id())
				progress <- progressDelta{total: -1}
			} else {
				// Remember we've seen this, and add it to our queue
				seen.Add(rsrc.Id())
				queue = append(queue, rsrc)
				ri.log.V(2).Info("Buffering import", "resource", rsrc.Id())
			}
		case upstream <- current:
			// We've sent the current importable, so clear it
			ri.log.V(2).Info("Queued import", "resource", current.Id())
			current = nil
		}
	}
}

func (ri *ResourceImporter) importWorker(
	ctx context.Context,
	pending chan ImportableResource,
	completed chan ImportResourceResult,
	progress chan progressDelta,
) {
	for rsrc := range pending {
		if ctx.Err() != nil {
			// If we're aborting, just discard everything until it's all gone
			progress <- progressDelta{total: -1}
			continue
		}

		// We have a resource to import
		ri.log.V(1).Info("Importing", "resource", rsrc.Id())
		result := ri.ImportResource(ctx, rsrc, progress)
		completed <- result

	}
}

func (ri *ResourceImporter) collateResults(
	completed chan ImportResourceResult,
	candidates chan ImportableResource,
	progress chan progressDelta,
) {
	report := newResourceImportReport()

	for importResult := range completed {
		rsrc := importResult.resource
		gk := rsrc.GroupKind()

		// Enqueue any child resources we found
		progress <- progressDelta{total: len(importResult.pending)}
		for _, p := range importResult.pending {
			candidates <- p
		}
		progress <- progressDelta{complete: 1}

		if importResult.err != nil {
			var skipped *ImportSkippedError
			if errors.As(importResult.err, &skipped) {
				ri.log.V(1).Info(
					"Skipped",
					"kind", gk,
					"name", rsrc.Name(),
					"because", skipped.Because)
				report.AddSkippedImport(rsrc, skipped.Because)
			} else {
				ri.log.Error(importResult.err,
					"Failed",
					"kind", gk,
					"name", rsrc.Name())

				report.AddFailedImport(rsrc, importResult.err.Error())
			}
		} else {
			ri.log.Info(
				"Imported",
				"kind", gk,
				"name", rsrc.Name())

			report.AddSuccessfulImport(rsrc)
			ri.imported[rsrc.Id()] = rsrc
		}
	}

	report.WriteToLog(ri.log)
}

func (ri *ResourceImporter) ImportResource(
	ctx context.Context,
	rsrc ImportableResource,
	parent chan progressDelta,
) ImportResourceResult {
	// Import it
	gk := rsrc.GroupKind()
	name := fmt.Sprintf("%s %s", gk, rsrc.Name())
	progress := ri.createSubProgressBar(name, parent)
	pending, err := rsrc.Import(ctx, progress, ri.log)
	result := ImportResourceResult{
		resource: rsrc,
		pending:  pending,
		err:      err,
	}

	return result
}

func (ri *ResourceImporter) createProgressBar(
	name string,
	done chan struct{},
) chan progressDelta {
	result := make(chan progressDelta)

	// Create a new progress bar
	bar := ri.progress.AddBar(
		0, // zero total because we don't know how much work will be needed
		mpb.PrependDecorators(
			decor.Name(name, decor.WCSyncSpaceR)),
		mpb.AppendDecorators(
			decor.Percentage(decor.WC{W: 5})))

	// Monitor for progress updates
	go func() {
		var complete int64
		var total int64
		for delta := range result {
			complete += int64(delta.complete)
			total += int64(delta.total)
			bar.SetTotal(total, false)
			bar.SetCurrent(complete)

			if total > 0 && complete >= total {
				close(done)
				break
			}
		}

		if total == 0 {
			total = 1
		}

		bar.SetCurrent(total)
		bar.SetTotal(total, true)
	}()

	return result
}

func (ri *ResourceImporter) createSubProgressBar(
	name string,
	parent chan progressDelta,
) chan progressDelta {
	result := make(chan progressDelta)

	// Create a new progress bar
	bar := ri.progress.AddBar(
		0, // zero total because we don't know how much work will be needed
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(name, decor.WCSyncSpaceR)),
		mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})))

	// Monitor for progress updates
	go func() {
		var complete int64
		var total int64
		for delta := range result {
			complete += int64(delta.complete)
			total += int64(delta.total)
			bar.SetTotal(total, false)
			bar.SetCurrent(complete)
			parent <- delta
		}

		if total == 0 {
			total = 1
		}

		bar.SetCurrent(total)
		bar.SetTotal(total, true)
	}()

	return result
}
