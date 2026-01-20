/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"fmt"
	"strings"
	"time"
)

type TestAction byte

const (
	Running    = TestAction("Running")
	Paused     = TestAction("Paused")
	Continuing = TestAction("Continuing")
	Output     = TestAction("Output")
	Passed     = TestAction("Passed")
	Failed     = TestAction("Failed")
	Skipped    = TestAction("Skipped")
)

// TestRun captures the details of an individual test run
type TestRun struct {
	Action   TestAction
	Package  string
	Test     string
	Output   []string
	RunTime  time.Duration
	started  *time.Time
	hasPanic bool
}

// run is used to capture the time when a test started execution
func (tr *TestRun) run(started time.Time) {
	// Only record the start if we're currently running
	if tr.started == nil {
		tr.started = &started
	}

	tr.Action = Running
}

// pause is used to capture the instant when a test was paused and stopped running
func (tr *TestRun) pause(stopped time.Time) {
	// If we're not running, panic because this is unexpected
	if tr.started == nil {
		msg := fmt.Sprintf(
			"Test %s in package %s was paused when it was not running",
			tr.Test,
			tr.Package)
		panic(msg)
	}

	tr.RunTime += stopped.Sub(*tr.started)
	tr.started = nil
	tr.Action = Paused
}

// resume indicates the test is continuing to run
// If we don't think we're currently running, capture the start time.
// Yes, this is very similar implementation to run(), but the semantics are
// different and I suspect we'll want to diverge them in the future.
func (tr *TestRun) resume(continued time.Time) {
	if tr.started == nil {
		tr.started = &continued
	}

	tr.Action = Running
}

// output adds a line of output to this test
func (tr *TestRun) output(line string) {
	tr.Output = append(tr.Output, line)

	if strings.HasPrefix(line, "panic") {
		tr.hasPanic = true
	}
}

// complete captures the final result of the test
func (tr *TestRun) complete(result string, completed time.Time) {
	if tr.started != nil {
		tr.RunTime += completed.Sub(*tr.started)
		tr.started = nil
	}

	tr.RunTime = sensitiveRound(tr.RunTime)
	switch strings.ToLower(result) {
	case "pass":
		tr.Action = Passed
	case "fail":
		tr.Action = Failed
	case "skip":
		tr.Action = Skipped
	default:
		tr.Action = TestAction(result) // capture unknown actions as-is
	}
}

func (tr *TestRun) IsInteresting() bool {
	result := true

	// Tests that pass, skip, or pause are not interesting
	switch tr.Action {
	case Passed:
		// Tests that pass aren't interesting
		result = false
	case Skipped:
		// Tests that are skipped aren't interesting
		result = false
	case Paused:
		// Tests that are paused aren't interesting (another test will have been responsible for the terminating the
		// test suite before they can resume)
		result = false
	case Running:
		// Tests that are running aren't interesting (another test will be responsible for the terminating the test
		// suite while they're executing)
		result = false
	}

	// Tests that have a panic are interesting, regardless of the result
	if tr.hasPanic {
		result = true
	}

	return result
}

func (d TestRun) actionSymbol() string {
	switch d.Action {
	case Passed:
		return "✅"
	case Failed:
		return "❌"
	case Skipped:
		return "⏭️"
	default:
		panic(fmt.Sprintf("unhandled action: %s", d.Action))
	}
}
