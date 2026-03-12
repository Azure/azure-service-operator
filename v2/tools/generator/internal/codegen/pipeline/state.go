/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/rotisserie/eris"
	"golang.org/x/exp/maps"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// State is an immutable instance that captures the information being passed along the pipeline
type State struct {
	definitions    astmodel.TypeDefinitionSet // set of type definitions generated so far
	stagesSeen     set.Set[string]            // set of ids of the stages already run
	stagesExpected map[string]set.Set[string] // set of ids of expected stages, each with a set of ids for the stages expecting them
	stateInfo      map[StateDataKey]any       // map of state information
}

// StateDataKey defines a unique key used to store/recall information in the pipeline state.
// This allows one stage to store information for consumption by another later stage.
type StateDataKey string

const (
	AllKnownResources   StateDataKey = "AllKnownResources"
	ConversionGraphInfo StateDataKey = "ConversionGraph"
	ExportedConfigMaps  StateDataKey = "ExportedConfigMaps"
)

// NewState returns a new empty state
// definitions is a (possibly empty) sequence of types to combine for the initial state
func NewState(definitions ...astmodel.TypeDefinitionSet) *State {
	defs := make(astmodel.TypeDefinitionSet)
	for _, ts := range definitions {
		defs = defs.OverlayWith(ts)
	}

	return &State{
		definitions:    defs,
		stagesSeen:     set.Make[string](),
		stagesExpected: make(map[string]set.Set[string]),
	}
}

// WithDefinitions returns a new independentState with the given type definitions instead
func (s *State) WithDefinitions(definitions astmodel.TypeDefinitionSet) *State {
	result := s.copy()
	result.definitions = definitions
	return result
}

// WithOverlaidDefinitions returns a new independent State with the given type definitions overlaid on the existing ones.
// Any new definitions are added, and any existing definitions are replaced.
func (s *State) WithOverlaidDefinitions(definitions astmodel.TypeDefinitionSet) *State {
	result := s.copy()
	result.definitions = result.definitions.OverlayWith(definitions)
	return result
}

// WithSeenStage records that the passed stage has been seen
func (s *State) WithSeenStage(id string) *State {
	result := s.copy()
	result.stagesSeen.Add(id)         // Record that we saw this stage
	delete(result.stagesExpected, id) // Discard expectations as they are satisfied
	return result
}

// WithExpectation records our expectation that the later stage is coming
func (s *State) WithExpectation(earlierStage string, laterStage string) *State {
	result := s.copy()
	if expected, ok := result.stagesExpected[laterStage]; ok {
		expected.Add(earlierStage)
		return result
	}

	expected := set.Make(earlierStage)
	result.stagesExpected[laterStage] = expected

	return result
}

// Definitions returns the set of type definitions contained by the state
func (s *State) Definitions() astmodel.TypeDefinitionSet {
	return s.definitions
}

// CheckFinalState checks that our final state is valid, returning an error if not
func (s *State) CheckFinalState() error {
	var errs []error
	for required, requiredBy := range s.stagesExpected {
		for stageID := range requiredBy {
			errs = append(errs, eris.Errorf("postrequisite %q of stage %q not satisfied", required, stageID))
		}
	}

	return kerrors.NewAggregate(errs)
}

// copy creates a new independent copy of the state
func (s *State) copy() *State {
	return &State{
		definitions:    s.definitions.Copy(),
		stagesSeen:     s.stagesSeen,
		stagesExpected: s.stagesExpected,
		stateInfo:      maps.Clone(s.stateInfo),
	}
}

// StateWithData returns a new state with the given information included.
// Any existing information with the same key is replaced.
// Has to be written as a standalone function because methods can't introduce new generic variation.
func StateWithData[I any](
	state *State,
	key StateDataKey,
	info I,
) *State {
	result := state.copy()
	if result.stateInfo == nil {
		result.stateInfo = make(map[StateDataKey]any)
	}

	result.stateInfo[key] = info
	return result
}

// GetStateData returns the information stored in the state under the given key.
// If no information is stored under the key, or if it doesn't have the expected type, the second
// return value is false.
// Has to be written as a standalone function because methods can't introduce new generic variables.
func GetStateData[I any](
	state *State,
	key StateDataKey,
) (I, error) {
	if state.stateInfo == nil {
		var zero I
		return zero, eris.Errorf("no state information available")
	}

	value, ok := state.stateInfo[key]
	if !ok {
		var zero I
		return zero, eris.Errorf("no state information found for key %s", key)
	}

	info, ok := value.(I)
	if !ok {
		var zero I
		return zero,
			eris.Errorf(
				"state information found for key %s of type %T, expected %T",
				key,
				value,
				zero)

	}

	return info, nil
}
