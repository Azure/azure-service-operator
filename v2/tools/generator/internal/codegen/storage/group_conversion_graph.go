/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// GroupConversionGraph represents the directed graph of conversions between versions for a single group
type GroupConversionGraph struct {
	group     string                              // Group of the resource needing conversions
	subGraphs map[string]*ResourceConversionGraph // Map of subgraphs, one for each resource type
}

// LookupTransition accepts a type name and looks up the transition to the next version in the graph
// Returns the next version and true if it's found, or an empty type name and false if not.
func (graph *GroupConversionGraph) LookupTransition(name astmodel.TypeName) (astmodel.TypeName, bool) {
	subgraph, ok := graph.subGraphs[name.Name()]
	if !ok {
		return astmodel.EmptyTypeName, false
	}

	return subgraph.LookupTransition(name)
}

// TransitionCount returns the number of transitions in the graph
func (graph *GroupConversionGraph) TransitionCount() int {
	result := 0
	for _, g := range graph.subGraphs {
		result += g.TransitionCount()
	}

	return result
}
