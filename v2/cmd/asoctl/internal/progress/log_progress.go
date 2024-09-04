/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package progress

import (
	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/pkg/importing"
	"github.com/go-logr/logr"
)

type logProgress struct {
	updates chan progressDelta // Channel for receiving updates (avoids concurrent access to progress)
	log     logr.Logger        // Logger for updating progress
	done    chan struct{}      // Channel for knowing when we're done
}

var _ importing.Progress = &logProgress{}

// NewLog creates a new progress reporter using a log to report progress.
func NewLog(
	name string,
	log logr.Logger,
	done chan struct{},
) importing.Progress {
	return newLogProgress(name, log, done, nil)
}

func newLogProgress(
	name string,
	log logr.Logger,
	done chan struct{},
	parent importing.Progress,
) *logProgress {

	result := &logProgress{
		updates: make(chan progressDelta),
		log:     log,
		done:    done,
	}

	// Monitor for progress updates
	go func() {
		var completed int64
		var total int64
		for delta := range result.updates {
			completed += int64(delta.complete)
			total += int64(delta.pending)

			// If we're done, finish up
			if parent == nil && total > 0 && completed >= total {
				// We're the root progress log, so we also need to handle the done channel
				close(done)
				break
			}

			// Pass updates to our parent bar (if any)
			if parent != nil {
				if delta.pending > 0 {
					parent.AddPending(delta.pending)
				}

				if delta.complete > 0 {
					parent.Completed(delta.complete)
				}
			}
		}

		if total == 0 {
			total = 1
		}

		// Final step, log that we're done.
		result.log.Info("Import complete", "name", name)
	}()

	log.Info("Starting import", "name", name)

	return result
}

func (lp *logProgress) AddPending(pending int) {
	lp.updates <- progressDelta{
		complete: 0,
		pending:  pending,
	}
}

func (lp *logProgress) Completed(completed int) {
	lp.updates <- progressDelta{
		complete: completed,
		pending:  0,
	}
}

func (lp *logProgress) Create(name string) importing.Progress {
	return newLogProgress(name, lp.log, lp.done, lp)
}
