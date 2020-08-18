/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"fmt"
)

// GitCommit and GitTreeState are populated by the Makefile.
var (
	GitCommit    string
	GitTreeState string
)

func combinedVersion() string {
	result := GitCommit
	if GitTreeState != "clean" {
		result += fmt.Sprintf(" (tree is %s)", GitTreeState)
	}
	return result
}
