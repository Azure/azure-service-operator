/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
)

// State is an immutable instance that captures the information being passed along the pipeline
type State struct {
	definitions     astmodel.TypeDefinitionSet // set of type definitions generated so far
	conversionGraph *storage.ConversionGraph   // graph of transitions between packages in our conversion graph
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
	}
}

// WithDefinitions returns a new independentState with the given type definitions instead
func (s *State) WithDefinitions(definitions astmodel.TypeDefinitionSet) *State {
	return &State{
		definitions:     definitions,
		conversionGraph: s.conversionGraph,
	}
}

// WithConversionGraph returns a new independent State with the given conversion graph instead
func (s *State) WithConversionGraph(graph *storage.ConversionGraph) *State {
	if s.conversionGraph != nil {
		panic("may only set the conversion graph once")
	}

	return &State{
		definitions:     s.definitions.Copy(),
		conversionGraph: graph,
	}
}

// Definitions returns the set of type definitions contained by the state
func (s *State) Definitions() astmodel.TypeDefinitionSet {
	return s.definitions
}

// ConversionGraph returns the conversion graph included in our state (may be null)
func (s *State) ConversionGraph() *storage.ConversionGraph {
	return s.conversionGraph
}
