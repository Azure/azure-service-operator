/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

type testRunFactory struct {
	runs map[string]*TestRun
}

func newTestRunFactory() *testRunFactory {
	return &testRunFactory{
		runs: make(map[string]*TestRun),
	}
}

func (f *testRunFactory) apply(action JSONFormat) {
	// Find (or create) the test run for this item of data
	testrun, found := f.runs[action.key()]
	if !found {
		testrun = &TestRun{
			Package: action.Package,
			Test:    action.Test,
		}

		f.runs[action.key()] = testrun
	}

	switch action.Action {
	case "run":
		testrun.run(action.Time)

	case "pause":
		testrun.pause(action.Time)

	case "cont":
		testrun.resume(action.Time)

	case "output":
		testrun.output(action.Output)

	case "pass", "fail", "skip":
		testrun.complete(action.Action, action.Time)
	}
}
