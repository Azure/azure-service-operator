/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
)

// State is an immutable instance that captures the information being passed along the pipeline
type State struct {
	definitions     astmodel.TypeDefinitionSet // set of type definitions generated so far
	conversionGraph *storage.ConversionGraph   // graph of transitions between packages in our conversion graph
	stagesSeen      set.Set[string]            // set of ids of the stages already run
	stagesExpected  map[string]set.Set[string] // set of ids of expected stages, each with a set of ids for the stages expecting them
}

/*
 *  TODO: Future extension (suggested by @matthchr):
 *  Instead of hard coding specific knowledge in the state type, implement a generic solution where stages can stash
 *  information in a map indexed by their unique identifier; later stages can then retrieve that information using
 *  that identifier.
 */

// NewState returns a new empty state
// definitions is a (possibly empty) sequence of types to combine for the initial state
func NewState(definitions ...astmodel.TypeDefinitionSet) *State {
	defs := make(astmodel.TypeDefinitionSet)
	for _, ts := range definitions {
		defs = defs.OverlayWith(ts)
	}

	return &State{
		definitions:     defs,
		conversionGraph: nil,
		stagesSeen:      set.Make[string](),
		stagesExpected:  make(map[string]set.Set[string]),
	}
}

// WithDefinitions returns a new independentState with the given type definitions instead
func (s *State) WithDefinitions(definitions astmodel.TypeDefinitionSet) *State {
	result := s.copy()
	result.definitions = definitions
	return result
}

// WithConversionGraph returns a new independent State with the given conversion graph instead
func (s *State) WithConversionGraph(graph *storage.ConversionGraph) *State {
	if s.conversionGraph != nil {
		panic("may only set the conversion graph once")
	}

	result := s.copy()
	result.conversionGraph = graph
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
	if set, ok := result.stagesExpected[laterStage]; ok {
		set.Add(earlierStage)
		return result
	}

	set := set.Make(earlierStage)
	result.stagesExpected[laterStage] = set

	return result
}

// Definitions returns the set of type definitions contained by the state
func (s *State) Definitions() astmodel.TypeDefinitionSet {
	return s.definitions
}

// ConversionGraph returns the conversion graph included in our state (may be null)
func (s *State) ConversionGraph() *storage.ConversionGraph {
	return s.conversionGraph
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
		definitions:     s.definitions.Copy(),
		conversionGraph: s.conversionGraph,
		stagesSeen:      s.stagesSeen,
		stagesExpected:  s.stagesExpected,
	}
}
