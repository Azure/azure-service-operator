/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
)

// ConversionGraph builds up a set of graphs of the required conversions between versions
// For each group (e.g. microsoft.storage or microsoft.batch) we have a separate subgraph of directed conversions
type ConversionGraph struct {
	subGraphs map[string]*GroupConversionGraph
}

// LookupTransition looks for a link and find out where it ends, given the starting reference.
// Returns the end and true if it's found, or nil and false if not.
func (graph *ConversionGraph) LookupTransition(start astmodel.PackageReference) (astmodel.PackageReference, bool) {
	// Expect to get either a local or a storage reference, not an external one
	local, ok := start.AsLocalPackage()
	if !ok {
		panic(fmt.Sprintf("cannot use external package reference %s with a conversion graph", start))
	}

	group := local.Group()
	subgraph, ok := graph.subGraphs[group]
	if !ok {
		return nil, false
	}

	return subgraph.LookupTransition(start)
}

// FindHubTypeName returns the type name of the hub resource, given the type name of one of the resources that is
// persisted using that hub type. This is done by following links in the conversion graph until we reach the end
func (graph *ConversionGraph) FindHubTypeName(name astmodel.TypeName) astmodel.TypeName {
	ref := name.PackageReference
	for {
		r, ok := graph.LookupTransition(ref)
		if !ok {
			break
		}
		ref = r
	}

	return astmodel.MakeTypeName(ref, name.Name())
}

// TransitionCount returns the number of transitions in the graph
func (graph *ConversionGraph) TransitionCount() int {
	result := 0
	for _, g := range graph.subGraphs {
		result += g.TransitionCount()
	}

	return result
}
