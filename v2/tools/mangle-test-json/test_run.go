/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import "time"

// TestRun captures the details of an individual test run
type TestRun struct {
	Action  string
	Package string
	Test    string
	Output  []string
	RunTime time.Duration
}
