/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"fmt"
)

// GitCommit and GitTreeState are populated by the Makefile.
var (
	GitCommit    string
	GitTreeState string
)

func CombinedVersion() string {
	result := GitCommit
	if GitTreeState != "clean" {
		result += fmt.Sprintf(" (tree is %s)", GitTreeState)
	}
	return result
}
