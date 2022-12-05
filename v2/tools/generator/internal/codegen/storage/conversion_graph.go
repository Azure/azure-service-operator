/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// ConversionGraph builds up a set of graphs of the required conversions between versions
// For each group (e.g. microsoft.storage or microsoft.batch) we have a separate subgraph of directed conversions
type ConversionGraph struct {
	configuration *config.ObjectModelConfiguration
	subGraphs     map[string]*GroupConversionGraph
}

// LookupTransition looks for a link and find out where it ends, given the starting reference.
// Returns the end and true if it's found, or nil and false if not.
func (graph *ConversionGraph) LookupTransition(ref astmodel.PackageReference) (astmodel.PackageReference, bool) {
	// Expect to get either a local or a storage reference, not an external one
	group, _ := ref.GroupVersion()
	subgraph, ok := graph.subGraphs[group]
	if !ok {
		return nil, false
	}

	return subgraph.LookupTransition(ref)
}

// FindNextType returns the type name of the next closest type on the path to the hub type.
// Returns the type name and no error if the next type is found; empty name and no error if not; empty name and an error if
// something goes wrong.
// If the name passed in is for the hub type for the given resource, no next type will be found.
// This is used to identify the next type needed for property assignment functions, and is a building block for
// identification of hub definitions.
func (graph *ConversionGraph) FindNextType(name astmodel.TypeName, definitions astmodel.TypeDefinitionSet) (astmodel.TypeName, error) {

	// Look for a next type with the same name
	nextType, stepsUntilNextFound := graph.searchForMatchingType(name, definitions)

	// Look for a renamed type with the same name
	renamedType, stepsUntilRenameFound, err := graph.searchForRenamedType(name, definitions)
	if err != nil {
		return astmodel.EmptyTypeName, errors.Wrapf(err, "searching for type renamed from %s", name)
	}

	if renamedType.IsEmpty() {
		// No rename active, return whatever we found from the original name
		return nextType, nil
	}

	if nextType.IsEmpty() {
		// Have a renamed type, and no match on original name, return the renamed type
		return renamedType, nil
	}

	if stepsUntilRenameFound < stepsUntilNextFound {
		// Rename matched earlier type, no conflict
		// (this might happen if a different type is introduced with the same name in a later version, or if a type
		// is renamed in one version and renamed back in a later one)
		return renamedType, nil
	}

	// We have a conflict between the renamed type and the original type
	return astmodel.EmptyTypeName, errors.Errorf(
		"configured rename of %s to %s conflicts with existing type %s",
		name,
		renamedType,
		nextType)
}

// FindHub returns the type name of the hub resource, given the type name of one of the resources that is
// persisted using that hub type. This is done by following links in the conversion graph until we either reach the end
// or we find that a newer version of the type does not exist.
// Returns the hub type and true if found; an empty name and false if not.
func (graph *ConversionGraph) FindHub(name astmodel.TypeName, definitions astmodel.TypeDefinitionSet) (astmodel.TypeName, error) {
	// Look for the hub step
	result := name
	for {
		hub, err := graph.FindNextType(result, definitions)
		if err != nil {
			return astmodel.EmptyTypeName, errors.Wrapf(
				err,
				"finding hub for %s",
				name)
		}

		if hub.IsEmpty() {
			break
		}

		result = hub
	}

	return result, nil
}

// TransitionCount returns the number of transitions in the graph
func (graph *ConversionGraph) TransitionCount() int {
	result := 0
	for _, g := range graph.subGraphs {
		result += g.TransitionCount()
	}

	return result
}

// FindNextProperty finds what a given property would be called on the next type in our conversion graph.
// Type renames are respected.
// When implemented, property renames need to be respected as well (this is why the method has been implemented here).
// declaringType is the type containing the property.
// property is the name of the property.
// definitions is a set of known definitions.
func (graph *ConversionGraph) FindNextProperty(
	ref astmodel.PropertyReference,
	definitions astmodel.TypeDefinitionSet) (astmodel.PropertyReference, error) {
	nextType, err := graph.FindNextType(ref.DeclaringType(), definitions)
	if err != nil {
		// Something went wrong
		return astmodel.EmptyPropertyReference,
			errors.Wrapf(err, "finding next property for %s", ref)
	}

	// If no next type, no next property either
	if nextType.IsEmpty() {
		return astmodel.EmptyPropertyReference, nil
	}

	//TODO: property renaming support goes here (when implemented)

	return astmodel.MakePropertyReference(nextType, ref.Property()), nil
}

// searchForRenamedType walks through the conversion graph looking for a match to our configured rename.
// If no rename is configured, returns an empty type name and -1. If one is configured, either returns the found type
// and the number of steps, or an error if not found.
// We only look for type renames if we're starting from a storage package - no renames are needed when converting
// from an api package to a storage package because the storage versions are always synthesized with an exact match
// on type names.
func (graph *ConversionGraph) searchForRenamedType(
	name astmodel.TypeName,
	definitions astmodel.TypeDefinitionSet) (astmodel.TypeName, int, error) {

	// No configuration, or we're not looking at a storage package
	if graph.configuration == nil || !astmodel.IsStoragePackageReference(name.PackageReference) {
		return astmodel.EmptyTypeName, -1, nil
	}

	rename, err := graph.configuration.LookupNameInNextVersion(name)
	if config.IsNotConfiguredError(err) {
		// We found no configured rename, nothing to do
		return astmodel.EmptyTypeName, -1, nil
	}

	// If we have any error other than a NotConfiguredError, something went wrong, and we must abort
	if err != nil {
		return astmodel.EmptyTypeName, -1, errors.Wrapf(
			err,
			"finding next type after %s",
			name)
	}

	newType := name.WithName(rename)
	result, stepsUntilFound := graph.searchForMatchingType(newType, definitions)

	// Validity check on the type-rename to verify that it specifies a type that exists.
	// If we didn't find the type, the configured rename is invalid
	if result.IsEmpty() {
		group, version := name.PackageReference.GroupVersion()
		return astmodel.EmptyTypeName, -1, errors.Errorf(
			"rename of %s/%s/%s invalid because no type with name %s was found in any later version",
			group,
			version,
			name.Name(),
			rename)
	}

	return result, stepsUntilFound, nil
}

// searchForMatchingType walks through the conversion graph looking for the next type with the same name as the one
// provided. Returns the found type and the number of steps, if found; otherwise returns an empty type name and -1.
func (graph *ConversionGraph) searchForMatchingType(
	typeName astmodel.TypeName,
	definitions astmodel.TypeDefinitionSet) (astmodel.TypeName, int) {
	return graph.searchForMatchingTypeImpl(typeName, definitions, 1)
}

// searchForMatchingTypeImpl is the recursive implementation of searchForMatchingType.
func (graph *ConversionGraph) searchForMatchingTypeImpl(
	typeName astmodel.TypeName,
	definitions astmodel.TypeDefinitionSet,
	stepsSoFar int) (astmodel.TypeName, int) {
	nextPackage, ok := graph.LookupTransition(typeName.PackageReference)
	if !ok {
		// No next package, we've fallen off the end of the graph
		return astmodel.EmptyTypeName, -1
	}

	nextTypeName := astmodel.MakeTypeName(nextPackage, typeName.Name())
	if definitions.Contains(nextTypeName) {
		// Found the type we're looking for
		return nextTypeName, stepsSoFar
	}

	// Keep walking to find the next type, if any
	return graph.searchForMatchingTypeImpl(nextTypeName, definitions, stepsSoFar+1)
}
