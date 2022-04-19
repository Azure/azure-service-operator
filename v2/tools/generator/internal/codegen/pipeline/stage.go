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

	"github.com/Azure/azure-service-operator/v2/internal/set"
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

	if !knownLegacyStages.Contains(id) {
		msg := fmt.Sprintf(
			"No new legacy stages (use NewStage instead): %s is not the id of a known legacy stage",
			id)
		panic(msg)
	}

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

var knownLegacyStages = set.MakeStringSet(
	"addCrossResourceReferences",
	"addCrossplaneAtProviderProperty",
	"addCrossplaneEmbeddedResourceSpec",
	"addCrossplaneEmbeddedResourceStatus",
	"addCrossplaneForProviderProperty",
	"addCrossplaneOwnerProperties",
	"addStatusFromSwagger",
	"allof-anyof-objects",
	"applyArmConversionInterface",
	"assertTypesStructureValid",
	"augmentSpecWithStatus",
	"collapseCrossGroupReferences",
	"createArmTypes",
	"deleteGenerated",
	"determineResourceOwnership",
	"ensureArmTypeExistsForEveryType",
	"exportControllerResourceRegistrations",
	"exportPackages",
	"flattenProperties",
	"flattenResources",
	"injectHubFunction",
	"injectOriginalGVKFunction",
	"injectOriginalVersionFunction",
	"injectOriginalVersionProperty",
	"loadSchema",
	"markStorageVersion",
	"nameTypes",
	"pluralizeNames",
	"propertyRewrites",
	"removeAliases",
	"removeEmbeddedResources",
	"replaceAnyTypeWithJSON",
	"reportTypesAndVersions",
	"rogueCheck",
	"simplifyDefinitions",
	"stripUnreferenced")

// HasId returns true if this stage has the specified id, false otherwise
func (stage *Stage) HasId(id string) bool {
	return stage.id == id
}

// RequiresPrerequisiteStages declares which stages must have completed before this one is executed.
// Use prerequisites to specify stages that must be present for this stage to work - typically this means that the
// earlier stage is responsible for creating the preconditions required for this stage to operate correctly.
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
// Use postrequisites when it is necessary for that later stage to act on the results of this one.
//
// For example, InjectJsonSerializationTests creates round trip serialization tests for any object types that have
// properties. It's not correct to give InjectJsonSerializationTests a prerequisite on every earlier stage that creates
// new object types becauses it isn't concerned with where those object came from. However, those earlier stages DO want
// their new object types to be tested, so they declare a post-requisite on InjectJsonSerializationTests to ensure this
// happens.
//
// Post-requisites are thus not completely isomorphic with RequiresPrerequisiteStages  as there may be supporting stages that are
// sometimes omitted from execution when targeting different outcomes. Having both pre- and post-requisites allows the
// dependencies to drop out cleanly when different stages are present.
//
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
	if err := stage.checkPreconditions(state); err != nil {
		return nil, errors.Wrapf(err, "preconditions of stage %s not met", stage.id)
	}

	resultState, err := stage.action(ctx, state)
	if err != nil {
		return nil, err
	}

	return resultState.WithSeenStage(stage.id), nil
}

// checkPreconditions checks to ensure the preconditions of this stage have been satisfied
func (stage *Stage) checkPreconditions(state *State) error {
	klog.V(3).Infof("Checking requisites of %s", stage.Id())

	if err := stage.checkPrerequisites(state); err != nil {
		return err
	}

	if err := stage.checkPostrequisites(state); err != nil {
		return err
	}

	return nil
}

// checkPrerequisites returns an error if the prerequisites of this stage have not been met
func (stage *Stage) checkPrerequisites(state *State) error {

	var errs []error
	for _, prereq := range stage.prerequisites {
		satisfied := state.stagesSeen.Contains(prereq)
		if satisfied {
			klog.V(3).Infof("[✓] Required prerequisite %s satisfied", prereq)
		} else {
			klog.V(3).Infof("[✗] Required prerequisite %s NOT satisfied", prereq)
		}

		if !satisfied {
			errs = append(errs, errors.Errorf("prerequisite %q of stage %q NOT satisfied.", prereq, stage.Id()))
		}
	}

	return kerrors.NewAggregate(errs)
}

// checkPostrequisites returns an error if the postrequisites of this stage have been satisfied early
func (stage *Stage) checkPostrequisites(state *State) error {
	var errs []error
	for _, postreq := range stage.postrequisites {
		early := state.stagesSeen.Contains(postreq)
		if early {
			klog.V(3).Infof("[✗] Required postrequisite %s satisfied early", postreq)
			errs = append(errs, errors.Errorf("postrequisite %q satisfied of stage %q early.", postreq, stage.Id()))
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
