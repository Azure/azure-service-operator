/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importreporter

import (
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

type barReporter struct {
	progress *mpb.Progress      // MultiProgressBar for displaying progress
	updates  chan progressDelta // Channel for receiving updates (avoids concurrent access to progress)
	done     chan struct{}      // Channel for knowing when we're done
}

var _ Interface = &barReporter{}

// NewBar creates a new import reporter that displays an on-screen progress bar on the console.
func NewBar(
	name string,
	progress *mpb.Progress,
	done chan struct{},
) Interface {
	return newBarProgress(name, progress, done, nil)
}

func newBarProgress(
	name string,
	progress *mpb.Progress,
	done chan struct{},
	parent Interface,
) *barReporter {
	bar := progress.AddBar(
		0, // zero total because we don't know how much work will be needed
		mpb.PrependDecorators(
			decor.Name(name, decor.WCSyncSpaceR)),
		mpb.AppendDecorators(
			decor.CountersNoUnit("%d/%d", decor.WCSyncSpaceR)),
		mpb.BarRemoveOnComplete())

	result := &barReporter{
		updates:  make(chan progressDelta),
		progress: progress,
		done:     done,
	}

	// Monitor for importreporter updates
	go func() {
		var completed int64
		var total int64
		for delta := range result.updates {
			completed += int64(delta.complete)
			total += int64(delta.pending)
			bar.SetTotal(total, false)
			bar.SetCurrent(completed)

			if parent != nil {
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

			if total > 0 && completed >= total {
				// Yay, we're finished!
				break
			}
		}

		if total == 0 {
			total = 1
		}

		bar.SetCurrent(total)
		bar.SetTotal(total, true)

		if parent == nil {
			// We're the root importreporter bar, so we need to handle the done channel once all the work is complete
			close(done)
		}
	}()

	return result
}

func (bp *barReporter) AddPending(pending int) {
	bp.updates <- progressDelta{
		complete: 0,
		pending:  pending,
	}
}

func (bp *barReporter) Completed(completed int) {
	bp.updates <- progressDelta{
		complete: completed,
		pending:  0,
	}
}

func (bp *barReporter) Create(name string) Interface {
	return newBarProgress(name, bp.progress, bp.done, bp)
}

type progressDelta struct {
	complete int // Change to count of completed steps
	pending  int // Change to count of pending steps
}
