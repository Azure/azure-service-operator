/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/pkg/errors"
	"io"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// GroupConversionGraph represents the directed graph of conversions between versions for a single group
type GroupConversionGraph struct {
	group           string                              // Name of the group of the resources needing conversions
	configuration   *config.ObjectModelConfiguration    // Configuration used to look up renames
	subGraphs       map[string]*ResourceConversionGraph // Nested graphs, one for each resource type, keyed by resource name
	storagePackages []astmodel.InternalPackageReference // Sorted list of known storage packages in this group
}

// LookupTransition accepts a type name and looks up the transition to the next version in the graph
// Returns the next version if it's found, or an empty type name if not.
func (graph *GroupConversionGraph) LookupTransition(name astmodel.InternalTypeName) astmodel.InternalTypeName {
	subgraph, ok := graph.subGraphs[name.Name()]
	if !ok {
		return astmodel.InternalTypeName{}
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

// searchForRenamedType walks through the conversion graph looking for a match to our configured rename.
// If no type-rename is configured, returns an empty type name and no error.
// If a type-rename is configured, either returns the found type, or an error if not found.
// We only look for type renames if we're starting from a storage package - no renames are needed when converting
// from an api package to a storage package because the storage versions are always synthesized with an exact match
// on type names.
func (graph *GroupConversionGraph) searchForRenamedType(
	name astmodel.InternalTypeName,
	definitions astmodel.TypeDefinitionSet,
) (astmodel.InternalTypeName, error) {

	// No configuration, or we're not looking at a storage package
	if graph.configuration == nil || !astmodel.IsStoragePackageReference(name.PackageReference()) {
		return astmodel.InternalTypeName{}, nil
	}

	rename, ok := graph.configuration.TypeNameInNextVersion.Lookup(name)
	if !ok {
		// We found no configured rename, nothing to do
		return astmodel.InternalTypeName{}, nil
	}

	// We have a configured rename, need to search through packages to find the type with that name
	foundStart := false
	for _, pkg := range graph.storagePackages {
		// Look for the package that contains the type we're converting from
		if pkg.Equals(name.PackageReference()) {
			foundStart = true
			continue
		}

		// Skip any packages before the one we're converting from
		if !foundStart {
			continue
		}

		// Does our target type exist in this package?
		newType := name.WithPackageReference(pkg).WithName(rename)
		if definitions.Contains(newType) {
			return newType, nil
		}
	}

	// Didn't find the type we're looking for
	return astmodel.InternalTypeName{},
		errors.Errorf("rename of %s invalid because no type with name %s was found in any later version", name, rename)
}

// WriteTo gives a debug dump of the conversion graph for a particular type name
func (graph *GroupConversionGraph) WriteTo(kind string, writer io.Writer) error {
	subgraph, ok := graph.subGraphs[kind]
	if !ok {
		return nil
	}

	return subgraph.WriteTo(writer)
}
