/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package progress

import (
	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/pkg/importing"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

type barProgress struct {
	progress *mpb.Progress      // MultiProgressBar for displaying progress
	updates  chan progressDelta // Channel for receiving updates (avoids concurrent access to progress)
	done     chan struct{}      // Channel for knowing when we're done
}

var _ importing.Progress = &barProgress{}

// NewProgressBar creates a new progress reporter using an on-screen progress bar on the console.
func NewProgressBar(
	name string,
	progress *mpb.Progress,
	done chan struct{},
) importing.Progress {
	return newBarProgress(name, progress, done, nil)
}

func newBarProgress(
	name string,
	progress *mpb.Progress,
	done chan struct{},
	parent importing.Progress,
) *barProgress {
	bar := progress.AddBar(
		0, // zero total because we don't know how much work will be needed
		mpb.PrependDecorators(
			decor.Name(name, decor.WCSyncSpaceR)),
		mpb.AppendDecorators(
			decor.CountersNoUnit("%d/%d", decor.WCSyncSpaceR)))

	result := &barProgress{
		updates:  make(chan progressDelta),
		progress: progress,
		done:     done,
	}

	// Monitor for progress updates
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
			// We're the root progress bar, so we need to handle the done channel once all the work is complete
			close(done)
		}
	}()

	return result
}

func (bp *barProgress) AddPending(pending int) {
	bp.updates <- progressDelta{
		complete: 0,
		pending:  pending,
	}
}

func (bp *barProgress) Completed(completed int) {
	bp.updates <- progressDelta{
		complete: completed,
		pending:  0,
	}
}

func (bp *barProgress) Create(name string) importing.Progress {
	return newBarProgress(name, bp.progress, bp.done, bp)
}

type progressDelta struct {
	complete int // Change to count of completed steps
	pending  int // Change to count of pending steps
}
