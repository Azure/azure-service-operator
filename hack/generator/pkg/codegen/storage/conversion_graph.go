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

// FindNext returns the type name of the next closest type on the path to the hub type.
// Returns the type name and true if the next type is found; an empty name and false if not.
// If the name passed in is for the hub type for the given resource, no next type will be found.
func (graph *ConversionGraph) FindNext(name astmodel.TypeName, types astmodel.Types) (astmodel.TypeName, bool) {
	ref := name.PackageReference
	for {
		// Find the next package to consider
		r, ok := graph.LookupTransition(ref)
		if !ok {
			// No next reference
			return astmodel.TypeName{}, false
		}

		// Look up to see if the resource exists in this package
		n := astmodel.MakeTypeName(r, name.Name())
		if _, ok := types.TryGet(n); ok {
			// found the next type
			return n, true
		}

		// Didn't find it, check the next package
		// We do this to allow for gaps where a type is dropped and then restored across versions
		ref = r
	}

	// Didn't find it
	return astmodel.TypeName{}, false
}


// FindHub returns the type name of the hub resource, given the type name of one of the resources that is
// persisted using that hub type. This is done by following links in the conversion graph until we either reach the end
// or we find that a newer version of the type does not exist.
// Returns the hub type and true if found; an empty name and false if not.
// If the name passed in is for the hub type for the given resource, no hub type will be found.
func (graph *ConversionGraph) FindHub(name astmodel.TypeName, types astmodel.Types) (astmodel.TypeName, bool) {
	// Look for the hub step
	hub, ok := graph.FindNext(name, types)
	if !ok {
		// Not found, name is already have the hub type
		return astmodel.TypeName{}, false
	}

	// Look for a hub further on and return that if we found one
	if h, ok := graph.FindHub(hub, types); ok {
		return h, true
	}

	// Return the hub we found earlier
	return hub, true
}

// TransitionCount returns the number of transitions in the graph
func (graph *ConversionGraph) TransitionCount() int {
	result := 0
	for _, g := range graph.subGraphs {
		result += g.TransitionCount()
	}

	return result
}
