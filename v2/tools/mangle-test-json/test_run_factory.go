/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

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
			Package: action.Package,
			Test:    action.Test,
		testRun = &TestRun{
		}

		f.runs[action.key()] = testrun
		packageRuns[action.Test] = testRun
	}

	switch action.Action {
	case "run":
		testrun.run(action.Time)
		testRun.run(action.Time)

	case "pause":
		testrun.pause(action.Time)
		testRun.pause(action.Time)

	case "cont":
		testrun.resume(action.Time)
		testRun.resume(action.Time)

	case "output":
		testrun.output(action.Output)
		testRun.output(action.Output)

	case "pass", "fail", "skip":
		testrun.complete(action.Action, action.Time)
		testRun.complete(action.Action, action.Time)
	}
}
