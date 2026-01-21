/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import "unique"

type testRunFactory struct {
	// testRuns contains all tests, keyed by package then test name
	testRuns map[string]map[string]*TestRun
}

func newTestRunFactory() *testRunFactory {
	return &testRunFactory{
		testRuns: make(map[string]map[string]*TestRun),
	}
}

func (f *testRunFactory) apply(action JSONFormat) {
	// Find (or create) the test run for this item of data
	packageRuns, found := f.testRuns[action.Package]
	if !found {
		packageRuns = make(map[string]*TestRun)
		f.testRuns[action.Package] = packageRuns
	}

	testRun, found := packageRuns[action.Test]
	if !found {
		testRun = &TestRun{
			Package: unique.Make(action.Package),
			Test:    unique.Make(action.Test),
		}

		packageRuns[action.Test] = testRun
	}

	switch action.Action {
	case "run":
		testRun.run(action.Time)

	case "pause":
		testRun.pause(action.Time)

	case "cont":
		testRun.resume(action.Time)

	case "output":
		testRun.output(action.Output)

	case "pass", "fail", "skip":
		testRun.complete(action.Action, action.Time)
	}
}
