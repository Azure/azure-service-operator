/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"github.com/sourcegraph/conc"
	"golang.org/x/exp/maps"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/pkg/importreporter"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// ResourceImporter is the entry point for importing resources.
// Factory methods here provide ways to instantiate importers for different kinds of resources.
type ResourceImporter struct {
	factory   *importFactory                  // Shared factory used to create resources and other things
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
	resource ImportedResource     // The resource that was imported
	pending  []ImportableResource // Any child resources that need to be imported next
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
		factory:  newImportFactory(scheme),
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
	owner, err := ri.createOwnerFor(armID)
	if err != nil {
		return eris.Wrapf(err, "adding ARMID %s to importer", armID)
	}

	importer, err := NewImportableARMResource(armID, owner, ri.client)
	if err != nil {
		return eris.Wrapf(err, "failed to create importer for %q", armID)
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
	resources := make(chan ImportableResource) // resources to import
	allstages := conc.NewWaitGroup()

	watchdog := ri.startWatchdog(resources)

	pendingResources := ri.startDeduplicator(resources, watchdog, allstages)
	successes, failures := ri.startWorkers(ctx, pendingResources, allstages)

	report := newResourceImportReport()
	ri.startCollationOfResults(successes, resources, watchdog, allstages, report)
	ri.startCollationOfErrors(failures, watchdog, allstages, report)

	// Set up by adding our initial resources; these will be completed when we collate their results
	for _, rsrc := range ri.resources {
		resources <- rsrc
		ri.reporter.AddPending(1)
	}

	allstages.Wait()

	// Check for an abort, and return the error if so
	if ctx.Err() != nil {
		ri.log.Error(ctx.Err(), "Cancelling import.")
		return nil, ctx.Err()
	}

	// Get the summary report and write it
	report.WriteToLog(ri.log)

	// Now we've imported everything, return the resources
	// We do this even if there's an error so that we can return partial results
	return &Result{
		imported: maps.Values(ri.imported),
	}, nil
}

// startWatchdog starts a watchdog goroutine that will initiate shutdown once all imports are complete.
// We keep track of the number of inflight imports, and once that reaches zero, we close the channel.
func (ri *ResourceImporter) startWatchdog(
	resources chan ImportableResource,
) *watchdog {
	watchdog := newWatchdog(func() {
		close(resources)
	})

	return watchdog
}

// startDeduplicator starts a deduplicator goroutine that will ensure each distinct resource is only imported once,
// regardless of how many times we encounter it. We also proactively buffer the pending channel to avoid blocking.
// Any fixed size channel would risk deadlock for a sufficiently large resource graph.
// resources is the channel of resources to deduplicate.
// watchdog is updated to track the number of inflight operations. We increase this every time we see a unique new resource.
// progress is used to report on progress as we go.
// Returns a new channel that will contain only unique resources.
// queueUniqueImporters reads from the candidates channel, putting each importer onto pending exactly once.
// This ensures each distinct resource is only imported once,
func (ri *ResourceImporter) startDeduplicator(
	resources <-chan ImportableResource,
	watchdog *watchdog,
	waitgroup *conc.WaitGroup,
) chan ImportableResource {
	seen := set.Make[string]()
	var queue []ImportableResource
	var current ImportableResource = nil

	uniqueResources := make(chan ImportableResource)

	waitgroup.Go(func() {
		// Close the channel when we're done, so that workers shut down too
		defer close(uniqueResources)

	run:
		for {
			// Dequeue from our internal buffer if needed
			if current == nil && len(queue) > 0 {
				current = queue[0]
				queue = queue[1:]
			}

			// If we have a current importable to send, use the pending queue
			var upstream chan<- ImportableResource = nil
			if current != nil {
				upstream = uniqueResources
			}

			select {
			case rsrc, ok := <-resources:
				if !ok {
					// Channel closed
					break run
				} else if seen.Contains(rsrc.ID()) {
					// We've already seen this resource (we've already queued it for import)
					// So remove it from our count of work to be done
					ri.log.V(2).Info("Skipping duplicate import", "resource", rsrc.ID())
					ri.reporter.AddPending(-1)
				} else {
					// Remember we've seen this, and add it to our queue
					watchdog.starting()
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
	})

	return uniqueResources
}

// startWorkers starts the worker goroutines that will import resources.
// ctx is used to check for cancellation.
// resources is the channel of unique resources to import.
// waitgroup is used to track the number of inflight goroutines.
// returns two channels: one for successful imports, and one for errors.
func (ri *ResourceImporter) startWorkers(
	ctx context.Context,
	resources <-chan ImportableResource,
	waitgroup *conc.WaitGroup,
) (chan ImportResourceResult, chan ImportError) {
	successes := make(chan ImportResourceResult) // importers that have been executed successfully
	failures := make(chan ImportError)           // errors from importers that failed

	wg := conc.NewWaitGroup()
	workersRequired := ri.desiredWorkers()
	for i := 0; i < workersRequired; i++ {
		wg.Go(func() {
			ri.importWorker(ctx, resources, successes, failures, ri.reporter)
		})
	}

	// Once all the workers are done, close the channels
	waitgroup.Go(func() {
		wg.Wait()
		close(successes)
		close(failures)
	})

	return successes, failures
}

// importerWorker is a goroutine for importing resources.
// It reads from the pending channel, imports the resource, and sends the result to the completed
// channel if it worked, or the failed channel if it didn't.
// ctx is used to check for cancellation.
// pending is a source of resources to import.
// completed is where we send the result of a successful import.
// failed is where we send the error from a failed import.
func (ri *ResourceImporter) importWorker(
	ctx context.Context,
	pending <-chan ImportableResource,
	completed chan<- ImportResourceResult,
	failed chan<- ImportError,
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

		if imported, err := ri.importResource(ctx, rsrc, progress); err != nil {
			failed <- MakeImportError(err, rsrc.GroupKind(), rsrc.Name())
		} else {
			completed <- imported
		}
	}
}

// startCollationOfResults starts a goroutine that will collate the results of imports.
// completed is the channel of completed imports to drain.
// candidates is the channel that receives any new resources to import.
// watchdog is updated to track the number of inflight operations.
// waitgroup is used to track running goroutines.
// report is the report to write to.
func (ri *ResourceImporter) startCollationOfResults(
	completed <-chan ImportResourceResult, // completed imports for us to collate
	candidates chan<- ImportableResource, // additional candidates for importing
	watchdog *watchdog,
	waitgroup *conc.WaitGroup,
	report *resourceImportReport, // report to write to
) {
	waitgroup.Go(func() {
		for importResult := range completed {
			rsrc := importResult.resource
			gk := rsrc.GroupKind()

			// Enqueue any child resources we found; these will be marked as completed when we collate their results
			ri.reporter.AddPending(len(importResult.pending))
			for _, p := range importResult.pending {
				candidates <- p
			}

			ri.log.Info(
				"Imported",
				"kind", gk,
				"name", rsrc.Name())

			report.AddSuccessfulImport(gk)
			ri.imported[rsrc.ID()] = rsrc

			ri.completed(watchdog)
		}
	})
}

// startCollationOfErrors starts a goroutine that will collage all the errors that occurred during import.
func (ri *ResourceImporter) startCollationOfErrors(
	failures <-chan ImportError,
	watchdog *watchdog,
	waitgroup *conc.WaitGroup,
	report *resourceImportReport,
) {
	waitgroup.Go(func() {
		for ie := range failures {
			var skipped *SkippedError
			if eris.As(ie.err, &skipped) {
				ri.log.V(1).Info(
					"Skipped",
					"kind", ie.gk,
					"name", ie.name,
					"because", skipped.Because)
				report.AddSkippedImport(ie.gk, skipped.Because)
			} else {
				ri.log.Error(ie.err,
					"Failed",
					"kind", ie.gk,
					"name", ie.name)

				report.AddFailedImport(ie.gk, ie.err.Error())
			}

			ri.completed(watchdog)
		}
	})
}

// completed is used to indicate a resource has been fully processed.
// We do this after everything else because it might indicate we're completed the entire process.
func (ri *ResourceImporter) completed(watchdog *watchdog) {
	ri.reporter.Completed(1)
	watchdog.stopped()
}

func (ri *ResourceImporter) importResource(
	ctx context.Context,
	rsrc ImportableResource,
	parent importreporter.Interface,
) (ImportResourceResult, error) {
	// Import it
	gk := rsrc.GroupKind()
	name := fmt.Sprintf("%s %s", gk, rsrc.Name())

	// Create our importreporter indicator
	progress := parent.Create(name)

	// Our main resource is pending
	progress.AddPending(1)

	// Indicate the main resource is complete
	// (we must only do this when we return, to ensure we don't appear complete too early)
	defer progress.Completed(1)

	// Import the resource itself
	if imported, err := rsrc.Import(ctx, progress, ri.factory, ri.log); err != nil {
		return ImportResourceResult{}, eris.Wrapf(err, "importing %s", name)
	} else {
		return imported, nil
	}
}

// desiredWorkers returns the number of workers to use for importing resources.
func (ri *ResourceImporter) desiredWorkers() int {
	if ri.options.Workers > 0 {
		return ri.options.Workers
	}

	return 4
}

func (ri *ResourceImporter) createOwnerFor(
	id string,
) (*genruntime.ResourceReference, error) {
	armID, err := arm.ParseResourceID(id)
	if err != nil {
		// Error is already detailed, no need to wrap
		return nil, err
	}

	if armID.Parent == nil {
		// there is no parent, so no owner
		return nil, nil
	}

	// Resource groups don't need owners
	if armID.ResourceType.String() == arm.ResourceGroupResourceType.String() {
		return nil, nil
	}

	parent := armID.Parent.String()
	return &genruntime.ResourceReference{
		ARMID: parent,
	}, nil
}
