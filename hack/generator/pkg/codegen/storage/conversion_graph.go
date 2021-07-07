/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// ConversionGraph builds up a set of graphs of the required conversions between versions
// For each group (e.g. microsoft.storage or microsoft.batch) we have a separate subgraph of directed conversions
type ConversionGraph struct {
	subgraphs map[string]*GroupConversionGraph
}

// NewConversionGraph creates a new ConversionGraph
func NewConversionGraph() *ConversionGraph {
	return &ConversionGraph{
		subgraphs: make(map[string]*GroupConversionGraph),
	}
}

// AddLink adds a directed link into our graph
func (graph *ConversionGraph) AddLink(source astmodel.PackageReference, destination astmodel.PackageReference) {
	subgraph := graph.getSubGraph(source)
	subgraph.AddLink(source, destination)
}

// LookupTransition looks for a link and find out where it ends, given the starting reference.
// Returns the end and true if it's found, or nil and false if not.
func (graph *ConversionGraph) LookupTransition(start astmodel.PackageReference) (astmodel.PackageReference, bool) {
	subgraph := graph.getSubGraph(start)
	return subgraph.LookupTransition(start)
}

// getSubGraph finds the relevant subgraph for the group of the provided reference, creating one if necessary
func (graph *ConversionGraph) getSubGraph(ref astmodel.PackageReference) *GroupConversionGraph {
	// Expect to get either a local or a storage reference, not an external one
	local, ok := ref.AsLocalPackage()
	if !ok {
		panic(fmt.Sprintf("cannot use external package reference %s with a conversion graph", ref))
	}

	group := local.Group()
	subgraph, ok := graph.subgraphs[group]
	if !ok {
		subgraph = NewGroupConversionGraph(group)
		graph.subgraphs[group] = subgraph
	}

	return subgraph
}
