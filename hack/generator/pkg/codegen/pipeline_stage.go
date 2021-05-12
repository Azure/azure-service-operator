/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// PipelineStage represents a composable stage of processing that can transform or process the set
// of generated types
type PipelineStage struct {
	// Unique identifier used to manipulate the pipeline from code
	id string
	// Description of the stage to use when logging
	description string
	// Stage implementation
	action func(context.Context, astmodel.Types) (astmodel.Types, error)
	// Tag used for filtering
	targets []PipelineTarget
	// Identifiers for other stages that must be completed first
	prerequisites []string
}

// MakePipelineStage creates a new pipeline stage that's ready for execution
func MakePipelineStage(
	id string,
	description string,
	action func(context.Context, astmodel.Types) (astmodel.Types, error)) PipelineStage {
	return PipelineStage{
		id:          id,
		description: description,
		action:      action,
	}
}

// HasId returns true if this stage has the specified id, false otherwise
func (stage *PipelineStage) HasId(id string) bool {
	return stage.id == id
}

func (stage PipelineStage) RequiresPrerequisiteStages(prerequisites ...string) PipelineStage {
	if len(stage.prerequisites) > 0 {
		panic(fmt.Sprintf(
			"Prerequisites of stage '%s' already set to '%s'; cannot modify to '%s'.",
			stage.id,
			strings.Join(stage.prerequisites, "; "),
			strings.Join(prerequisites, "; ")))
	}

	stage.prerequisites = prerequisites

	return stage
}

// UsedFor specifies that this stage should be used for only the specified targets
func (stage PipelineStage) UsedFor(targets ...PipelineTarget) PipelineStage {
	stage.targets = targets
	return stage
}

// IsUsedFor returns true if this stage should be used for the specified target
func (stage *PipelineStage) IsUsedFor(target PipelineTarget) bool {

	if len(stage.targets) == 0 {
		// Stages without specific targeting are always used
		return true
	}

	for _, t := range stage.targets {
		if t == target {
			// Stage should be used for this target
			return true
		}
	}

	return false
}
