/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

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
	Action func(context.Context, astmodel.Types) (astmodel.Types, error)
}

// MakePipelineStage creates a new pipeline stage that's ready for execution
func MakePipelineStage(
	id string,
	description string,
	action func(context.Context, astmodel.Types) (astmodel.Types, error)) PipelineStage {
	return PipelineStage{
		id:          id,
		description: description,
		Action:      action,
	}
}

// HasId returns true if this stage has the specified id, false otherwise
func (stage *PipelineStage) HasId(id string) bool {
	return stage.id == id
}
