/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importreporter

import (
	"github.com/go-logr/logr"
)

type logReporter struct {
	updates chan progressDelta // Channel for receiving updates (avoids concurrent access to importreporter)
	log     logr.Logger        // Logger for updating importreporter
	done    chan struct{}      // Channel for knowing when we're done
}

var _ Interface = &logReporter{}

// NewLog creates a new import reporter using a log to report progress.
func NewLog(
	name string,
	log logr.Logger,
	done chan struct{},
) Interface {
	return newLogProgress(name, log, done, nil)
}

func newLogProgress(
	name string,
	log logr.Logger,
	done chan struct{},
	parent Interface,
) *logReporter {
	result := &logReporter{
		updates: make(chan progressDelta),
		log:     log,
		done:    done,
	}

	// Monitor for importreporter updates
	go func() {
		var completed int64
		var total int64
		for delta := range result.updates {
			completed += int64(delta.complete)
			total += int64(delta.pending)

			// If we're done, finish up
			if total > 0 && completed >= total {
				close(result.updates)

				if parent == nil {
					// We're the root importreporter log, so we also need to handle the done channel
					close(done)
				}

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

		// Final step, log that we're done.
		result.log.Info("Import complete", "name", name)
	}()

	log.Info("Starting import", "name", name)

	return result
}

func (lp *logReporter) AddPending(pending int) {
	lp.updates <- progressDelta{
		complete: 0,
		pending:  pending,
	}
}

func (lp *logReporter) Completed(completed int) {
	lp.updates <- progressDelta{
		complete: completed,
		pending:  0,
	}
}

func (lp *logReporter) Create(name string) Interface {
	return newLogProgress(name, lp.log, lp.done, lp)
}
