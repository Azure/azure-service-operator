/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// Stage represents a composable stage of processing that can transform or process the set
// of generated types
type Stage struct {
	// Unique identifier used to manipulate the pipeline from code
	id string
	// Description of the stage to use when logging
	description string
	// Stage implementation
	action func(context.Context, *State) (*State, error)
	// Tag used for filtering
	targets []Target
	// Identifiers for other stages that must be completed before this one
	prerequisites []string
	// Identifiers for other stages that must be completed after this one
	postrequisites []string
}

// NewStage creates a new pipeline stage that's ready for execution
func NewStage(
	id string,
	description string,
	action func(context.Context, *State) (*State, error)) *Stage {
	return &Stage{
		id:          id,
		description: description,
		action:      action,
	}
}

// NewLegacyStage is a legacy constructor for creating a new pipeline stage that's ready for execution
// DO NOT USE THIS FOR ANY NEW STAGES - it's kept for compatibility with an older style of pipeline stages that will be
// migrated to the new style over time.
func NewLegacyStage(
	id string,
	description string,
	action func(context.Context, astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error)) *Stage {

	klog.Warning(id)

	return NewStage(
		id,
		description,
		func(ctx context.Context, state *State) (*State, error) {
			types, err := action(ctx, state.Definitions())
			if err != nil {
				return nil, err
			}

			return state.WithDefinitions(types), nil
		})
}

// HasId returns true if this stage has the specified id, false otherwise
func (stage *Stage) HasId(id string) bool {
	return stage.id == id
}

// RequiresPrerequisiteStages declares which stages must have completed before this one is executed
func (stage *Stage) RequiresPrerequisiteStages(prerequisites ...string) {
	if len(stage.prerequisites) > 0 {
		panic(fmt.Sprintf(
			"Prerequisites of stage '%s' already set to '%s'; cannot modify to '%s'.",
			stage.id,
			strings.Join(stage.prerequisites, "; "),
			strings.Join(prerequisites, "; ")))
	}

	stage.prerequisites = prerequisites
}

// RequiresPostrequisiteStages declares which stages must be executed after this one has completed
// This is not completely isomorphic with RequiresPrerequisiteStages as there may be supporting stages that are
// sometimes omitted from execution when targeting different outcomes. Having both pre- and post-requisites allows the
// dependencies to drop out cleanly when different stages are present.
func (stage *Stage) RequiresPostrequisiteStages(postrequisites ...string) {
	if len(stage.postrequisites) > 0 {
		panic(fmt.Sprintf(
			"Postrequisites of stage '%s' already set to '%s'; cannot modify to '%s'.",
			stage.id,
			strings.Join(stage.postrequisites, "; "),
			strings.Join(postrequisites, "; ")))
	}

	stage.postrequisites = postrequisites
}

// UsedFor specifies that this stage should be used for only the specified targets
func (stage *Stage) UsedFor(targets ...Target) *Stage {
	stage.targets = targets
	return stage
}

// IsUsedFor returns true if this stage should be used for the specified target
func (stage *Stage) IsUsedFor(target Target) bool {
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

// Id returns the unique identifier for this stage
func (stage *Stage) Id() string {
	return stage.id
}

// Description returns a human-readable description of this stage
func (stage *Stage) Description() string {
	return stage.description
}

// Run is used to execute the action associated with this stage
func (stage *Stage) Run(ctx context.Context, state *State) (*State, error) {
	return stage.action(ctx, state)
}

// CheckPrerequisites returns an error if the prerequisites of this stage have not been met
func (stage *Stage) CheckPrerequisites(priorStages astmodel.StringSet) error {
	var errs []error
	for _, prereq := range stage.prerequisites {
		satisfied := priorStages.Contains(prereq)
		if satisfied {
			klog.V(0).Infof("[✓] Required prerequisite %s satisfied", prereq)
		} else {
			klog.V(0).Infof("[✗] Required prerequisite %s NOT satisfied", prereq)
		}

		if !satisfied {
			errs = append(errs, errors.Errorf("prerequisite %q of stage %q not satisfied.", prereq, stage.id))
		}
	}

	return kerrors.NewAggregate(errs)
}

// Postrequisites returns the unique ids of stages that must run after this stage
func (stage *Stage) Postrequisites() []string {
	return stage.postrequisites
}

// Targets returns the targets this stage should be used for
// If no targets are returned, this stage should always be used
func (stage *Stage) Targets() []Target {
	return stage.targets
}
