/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/pkg/errors"
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
	stateInfo      map[StateInfo]any          // map of state information
}

/*
 *  TODO: Future extension (suggested by @matthchr):
 *  Instead of hard coding specific knowledge in the state type, implement a generic solution where stages can stash
 *  information in a map indexed by their unique identifier; later stages can then retrieve that information using
 *  that identifier.
 */

type StateInfo string

const (
	ConversionGraphInfo StateInfo = "ConversionGraph"
	ExportedConfigMaps  StateInfo = "ExportedConfigMaps"
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
		for stageId := range requiredBy {
			errs = append(errs, errors.Errorf("postrequisite %q of stage %q not satisfied", required, stageId))
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

// StateWithInfo returns a new state with the given information included.
// Any existing information with the same key is replaced.
// Has to be written as a standalone function because methods can't introduce new generic variation.
func StateWithInfo[I any](
	state *State,
	key StateInfo,
	info I,
) *State {
	result := state.copy()
	if result.stateInfo == nil {
		result.stateInfo = make(map[StateInfo]any)
	}

	result.stateInfo[key] = info
	return result
}

// GetStateInfo returns the information stored in the state under the given key.
// If no information is stored under the key, or if it doesn't have the expected type, the second
// return value is false.
// Has to be written as a standalone function because methods can't introduce new generic variation.
func GetStateInfo[I any](
	state *State,
	key StateInfo,
) (I, bool) {
	if state.stateInfo == nil {
		var result I
		return result, false
	}

	info, ok := state.stateInfo[key].(I)
	return info, ok
}
