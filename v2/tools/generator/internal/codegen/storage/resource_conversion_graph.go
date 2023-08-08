/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ResourceConversionGraph represents the directed graph of conversions between versions for a single resource/type
type ResourceConversionGraph struct {
	name  string                                                  // Name of the resource needing conversions
	links map[astmodel.InternalTypeName]astmodel.InternalTypeName // All the directed links in our conversion graph
}

// LookupTransition accepts a type name and looks up the transition to the next version in the graph
// Returns the next version, or an empty type name if not.
func (graph *ResourceConversionGraph) LookupTransition(name astmodel.TypeName) astmodel.TypeName {
	next, _ := graph.links[name]
	return next
}

// TransitionCount returns the number of transitions in the graph
func (graph *ResourceConversionGraph) TransitionCount() int {
	return len(graph.links)
}
