/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
)

// State is an immutable instance that captures the information being passed along the pipeline
type State struct {
	types           astmodel.Types           // set of types generated so far
	conversionGraph *storage.ConversionGraph // graph of transitions between packages in our conversion graph
}

/*
 *  TODO: Future extension (suggested by @matthchr):
 *  Instead of hard coding specific knowledge in the state type, implement a generic solution where stages can stash
 *  information in a map indexed by their unique identifier; later stages can then retrieve that information using
 *  that identifier.
 */

// NewState returns a new empty state
func NewState() *State {
	return &State{
		types:           make(astmodel.Types),
		conversionGraph: nil,
	}
}

// WithTypes returns a new independentState with the given types instead
func (s *State) WithTypes(types astmodel.Types) *State {
	return &State{
		types:           types,
		conversionGraph: s.conversionGraph,
	}
}

// WithConversionGraph returns a new independent State with the given conversion graph instead
func (s *State) WithConversionGraph(graph *storage.ConversionGraph) *State {
	if s.conversionGraph != nil {
		panic("may only set the conversion graph once")
	}

	return &State{
		types:           s.types.Copy(),
		conversionGraph: graph,
	}
}

// Types returns the set of types contained by the state
func (s *State) Types() astmodel.Types {
	return s.types
}

// ConversionGraph returns the conversion graph included in our state (may be null)
func (s *State) ConversionGraph() *storage.ConversionGraph {
	return s.conversionGraph
}
