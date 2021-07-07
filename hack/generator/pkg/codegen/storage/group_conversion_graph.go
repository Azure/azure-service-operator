/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// GroupConversionGraph represents the directed graph of conversions between versions for a single group
type GroupConversionGraph struct {
	group string                                                  // Name of the group needing conversions
	links map[astmodel.PackageReference]astmodel.PackageReference // All the directed links in our conversion graph
}

// NewGroupConversionGraph creates a new, empty, GroupConversionGraph ready for use
func NewGroupConversionGraph(group string) *GroupConversionGraph {
	return &GroupConversionGraph{
		group: group,
		links: make(map[astmodel.PackageReference]astmodel.PackageReference),
	}
}

// AddLink adds a directed link into our graph
func (graph *GroupConversionGraph) AddLink(source astmodel.PackageReference, destination astmodel.PackageReference) {
	graph.links[source] = destination
}

// LookupTransition a link and find out where it ends, given the starting reference.
// Returns the end and true if it's found, or nil and false if not.
func (graph *GroupConversionGraph) LookupTransition(start astmodel.PackageReference) (astmodel.PackageReference, bool) {
	end, found := graph.links[start]
	return end, found
}
