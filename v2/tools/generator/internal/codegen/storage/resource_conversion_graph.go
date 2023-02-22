/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ResourceConversionGraph represents the directed graph of conversions between versions for a single group
type ResourceConversionGraph struct {
	name  string                                  // Name of the resource needing conversions
	links map[astmodel.TypeName]astmodel.TypeName // All the directed links in our conversion graph
}

// LookupTransition accepts a type name and looks up the transition to the next version in the graph
// Returns the next version and true if it's found, or an empty type name and false if not.
func (graph *ResourceConversionGraph) LookupTransition(name astmodel.TypeName) (astmodel.TypeName, bool) {
	next, found := graph.links[name]
	return next, found
}

// TransitionCount returns the number of transitions in the graph
func (graph *ResourceConversionGraph) TransitionCount() int {
	return len(graph.links)
}
