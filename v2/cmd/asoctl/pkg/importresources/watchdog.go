/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"sync/atomic"
)

// watchdog keeps track of the number of inflight operations and triggers an action when they all complete.
type watchdog struct {
	inflight atomic.Int32
	action   func()
}

// newWatchdog returns a new watchdog instance
func newWatchdog(
	action func(),
) *watchdog {
	return &watchdog{
		action: action,
	}
}

// starting increments the number of inflight operations
func (w *watchdog) starting() {
	w.inflight.Add(1)
}

// stopped decrements the number of inflight operations and triggers the action if there are no
// remaining inflight operations
func (w *watchdog) stopped() {
	count := w.inflight.Add(-1)
	if count == 0 {
		w.action()
	}
}
