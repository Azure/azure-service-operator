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
	group string                                                  // Name of the group needing conversions
	links map[astmodel.PackageReference]astmodel.PackageReference // All the directed links in our conversion graph
}

// LookupTransition a link and find out where it ends, given the starting reference.
// Returns the end and true if it's found, or nil and false if not.
func (graph *GroupConversionGraph) LookupTransition(start astmodel.PackageReference) (astmodel.PackageReference, bool) {
	end, found := graph.links[start]
	return end, found
}

// TransitionCount returns the number of transitions in the graph
func (graph *GroupConversionGraph) TransitionCount() int {
	return len(graph.links)
}
