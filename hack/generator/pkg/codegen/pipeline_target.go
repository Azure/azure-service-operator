/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"fmt"
)

// PipelineTarget is used to classify what kind of pipeline we have
// Deliberately wraps a string because we *do* *not* want type compatibility with literal strings
type PipelineTarget struct {
	name string
}

var _ fmt.Stringer = PipelineTarget{}

var (
	// ARMTarget is used to tag stages that are required when generating types for working directly with Azure
	ARMTarget PipelineTarget = MakePipelineTarget("azure")

	// CrossplaneTarget is used to tag stages that are required when generating types for working with Crossplane
	CrossplaneTarget PipelineTarget = MakePipelineTarget("crossplane")
)

func MakePipelineTarget(tag string) PipelineTarget {
	return PipelineTarget{
		name: tag,
	}
}

func (t PipelineTarget) String() string {
	return t.name
}
