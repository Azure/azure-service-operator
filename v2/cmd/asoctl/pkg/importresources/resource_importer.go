/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/pkg/importreporter"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

// ResourceImporter is the entry point for importing resources.
// Factory methods here provide ways to instantiate importers for different kinds of resources.
type ResourceImporter struct {
	scheme    *runtime.Scheme                 // a reference to the scheme used by asoctl
	client    *genericarmclient.GenericClient // Client to use when talking to ARM
	resources []ImportableResource            // A slice of resources to be imported
	imported  map[string]ImportedResource     // A set of importers that have been successfully imported
	log       logr.Logger                     // Logger to use for logging
	reporter  importreporter.Interface        // Reporter to use for reporter updates
	options   ResourceImporterOptions         // Options for the importer
}

// ResourceImporterOptions are optional configuration items for the importer
type ResourceImporterOptions struct {
	// Workers is the number of concurrent imports to run at the same time. If not specified, a default of 4 is used.
	Workers int
}

type ImportResourceResult struct {
	resource ImportedResource
	pending  []ImportableResource
	err      error
}

// New creates a new factory with the scheme baked in
func New(
	scheme *runtime.Scheme,
	client *genericarmclient.GenericClient,
	log logr.Logger,
	reporter importreporter.Interface,
	options ResourceImporterOptions,
) *ResourceImporter {
	return &ResourceImporter{
		scheme:   scheme,
		client:   client,
		imported: make(map[string]ImportedResource),
		log:      log,
		reporter: reporter,
		options:  options,
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
	done chan struct{},
) (*Result, error) {
	workersRequired := ri.desiredWorkers()
	candidates := make(chan ImportableResource)  // candidates that need to be deduped
	pending := make(chan ImportableResource)     // importers that are pending import
	completed := make(chan ImportResourceResult) // importers that have been executed successfully
	report := make(chan *resourceImportReport)   // summary report of the import

	// Dedupe candidates so we import each distinct resource only once
	go ri.queueUniqueImporters(candidates, pending, ri.reporter)

	// Create workers to run the import
	for i := 0; i < workersRequired; i++ {
		go ri.importWorker(ctx, pending, completed, ri.reporter)
	}

	// Collate the results
	go ri.collateResults(completed, candidates, ri.reporter, report)

	// Set up by adding our initial resources; these will be completed when we collate their results
	for _, rsrc := range ri.resources {
		candidates <- rsrc
		ri.reporter.AddPending(1)
	}

	// Wait while everything runs
	<-done

	// Check for an abort, and return the error if so
	if ctx.Err() != nil {
		ri.log.Error(ctx.Err(), "Cancelling import.")
		return nil, ctx.Err()
	}

	// Close channels so final reporting and other cleanup occurs
	close(candidates)
	close(pending)
	close(completed)

	// Get the summary report and write it
	rpt := <-report
	rpt.WriteToLog(ri.log)

	// Now we've imported everything, return the resources
	// We do this even if there's an error so that we can return partial results
	resources := make([]ImportedResource, 0, len(ri.imported))
	for _, imported := range ri.imported {
		resources = append(resources, imported)
	}

	return &Result{
		imported: resources,
	}, nil
}

// queueUniqueImporters reads from the candidates channel, putting each importer onto pending exactly once.
// This ensures each distinct resource is only imported once, regardless of how many times we
// encounter it. We also proactively buffer the pending channel to avoid blocking. Any fixed size channel
// would risk deadlock for a sufficiently large resource graph.
func (ri *ResourceImporter) queueUniqueImporters(
	candidates <-chan ImportableResource,
	pending chan<- ImportableResource,
	progress importreporter.Interface,
) {
	seen := set.Make[string]()
	var queue []ImportableResource
	var current ImportableResource = nil

	running := true
	for running {
		// Dequeue from our internal buffer if needed
		if current == nil && len(queue) > 0 {
			current = queue[0]
			queue = queue[1:]
		}

		// If we have a current importable to send, use the pending queue
		var upstream chan<- ImportableResource = nil
		if current != nil {
			upstream = pending
		}

		select {
		case rsrc, ok := <-candidates:
			if !ok {
				// Channel closed
				running = false
			} else if seen.Contains(rsrc.ID()) {
				// We've already seen this resource (we've already queued it for import)
				// So remove it from our count of work to be done
				ri.log.V(2).Info("Skipping duplicate import", "resource", rsrc.ID())
				progress.AddPending(-1)
			} else {
				// Remember we've seen this, and add it to our queue
				seen.Add(rsrc.ID())
				queue = append(queue, rsrc)
				ri.log.V(2).Info("Buffering import", "resource", rsrc.ID())
			}
		case upstream <- current:
			// We've sent the current importable, so clear it
			ri.log.V(2).Info("Queued import", "resource", current.ID())
			current = nil
		}
	}
}

func (ri *ResourceImporter) importWorker(
	ctx context.Context,
	pending <-chan ImportableResource,
	completed chan<- ImportResourceResult,
	progress importreporter.Interface,
) {
	for rsrc := range pending {
		if ctx.Err() != nil {
			// If we're aborting, just remove everything from the queue until it's all gone
			progress.AddPending(-1)
			continue
		}

		// We have a resource to import
		ri.log.V(1).Info("Importing", "resource", rsrc.ID())
		result := ri.importResource(ctx, rsrc, progress)
		completed <- result
	}
}

func (ri *ResourceImporter) collateResults(
	completed <-chan ImportResourceResult, // completed imports for us to collate
	candidates chan<- ImportableResource, // additional candidates for importing
	progress importreporter.Interface, // importreporter tracking
	publish chan<- *resourceImportReport, // publishing our final summary
) {
	report := newResourceImportReport()
	for importResult := range completed {
		rsrc := importResult.resource
		gk := rsrc.GroupKind()

		// Enqueue any child resources we found; these will be marked as completed when we collate their results
		progress.AddPending(len(importResult.pending))
		for _, p := range importResult.pending {
			candidates <- p
		}

		if importResult.err != nil {
			var skipped *SkippedError
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
			ri.imported[rsrc.ID()] = rsrc
		}

		// Flag the main resource as complete
		// We do this after everything else because it might indicate we're finished
		progress.Completed(1)
	}

	publish <- report
}

func (ri *ResourceImporter) importResource(
	ctx context.Context,
	rsrc ImportableResource,
	parent importreporter.Interface,
) ImportResourceResult {
	// Import it
	gk := rsrc.GroupKind()
	name := fmt.Sprintf("%s %s", gk, rsrc.Name())

	// Create our importreporter indicator
	progress := parent.Create(name)

	// Our main resource is pending
	progress.AddPending(1)

	// Import the resource itself
	imported, err := rsrc.Import(ctx, ri.log)
	result := ImportResourceResult{
		resource: imported,
		err:      err,
	}

	// If the main resource was imported ok, look for any children
	if result.err == nil {
		result.pending, result.err = rsrc.FindChildren(ctx, progress)
	}

	// Indicate the main resource is complete
	// (we must only do this after checking for children, to ensure we don't appear complete too early)
	progress.Completed(1)

	return result
}

// desiredWorkers returns the number of workers to use for importing resources.
func (ri *ResourceImporter) desiredWorkers() int {
	if ri.options.Workers > 0 {
		return ri.options.Workers
	}

	return 4
}
