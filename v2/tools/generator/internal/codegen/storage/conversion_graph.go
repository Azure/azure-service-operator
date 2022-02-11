/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

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
	group, _, ok := ref.GroupVersion()
	if !ok {
		panic(fmt.Sprintf("cannot use external package reference %s with a conversion graph", ref))
	}

	subgraph, ok := graph.subGraphs[group]
	if !ok {
		return nil, false
	}

	return subgraph.LookupTransition(ref)
}

// FindNextType returns the type name of the next closest type on the path to the hub type.
// Returns the type name and no error if the next type is found; nil and no error if not; nil and an error if something
// goes wrong.
// If the name passed in is for the hub type for the given resource, no next type will be found.
// This is used to identify the next type needed for property assignment functions, and is a building block for
// identification of hub types.
func (graph *ConversionGraph) FindNextType(name astmodel.TypeName, types astmodel.Types) (astmodel.TypeName, error) {

	// Find the next package to consider
	nextPackage, ok := graph.LookupTransition(name.PackageReference)
	if !ok {
		// No next package, return nil. This is not an error condition.
		return astmodel.EmptyTypeName, nil
	}

	// Look to see if we have a type-rename that's directing us to use a different type
	//
	// We only look for type renames if we're starting from a storage package - no renames are needed when converting
	// from an api package to a storage package because the storage versions are always synthesized with an exact match
	// on type names.
	//
	haveRename := false
	rename := ""
	if graph.configuration != nil && astmodel.IsStoragePackageReference(name.PackageReference) {
		var err error
		rename, err = graph.configuration.LookupNameInNextVersion(name)

		// If we have any error other than a NotConfiguredError, something went wrong, and we must abort
		if err != nil && !config.IsNotConfiguredError(err) {
			return astmodel.EmptyTypeName, errors.Wrapf(
				err,
				"finding next type after %s",
				name)
		}

		haveRename = err == nil
	}

	// Our usual convention is to use the same name as this type, but found in nextPackage
	nextTypeName := astmodel.MakeTypeName(nextPackage, name.Name())

	// With no configured rename, we can return nextTypeName, as long as it exists
	if !haveRename {
		if _, found := types.TryGet(nextTypeName); found {
			return nextTypeName, nil
		}

		return astmodel.EmptyTypeName, nil
	}

	// Validity check on the type-rename we found - ensure it specifies a type that's known
	// If we don't find the type, the configured rename is invalid
	renamedTypeName := astmodel.MakeTypeName(nextPackage, rename)
	if _, found := types.TryGet(renamedTypeName); !found {
		return astmodel.EmptyTypeName, errors.Errorf(
			"rename of %s invalid because specified type %s does not exist",
			name,
			renamedTypeName)
	}

	// Validity check that the type-rename doesn't conflict
	// if v1.Foo and v2.Foo both exist, it's illegal to specify a type-rename on v1.Foo
	if _, found := types.TryGet(nextTypeName); found {
		return astmodel.EmptyTypeName, errors.Errorf(
			"configured rename of %s to %s conflicts with existing type %s",
			name,
			renamedTypeName,
			nextTypeName)
	}

	// We don't have a conflict, so we can return this rename
	return renamedTypeName, nil
}

// FindHub returns the type name of the hub resource, given the type name of one of the resources that is
// persisted using that hub type. This is done by following links in the conversion graph until we either reach the end
// or we find that a newer version of the type does not exist.
// Returns the hub type and true if found; an empty name and false if not.
func (graph *ConversionGraph) FindHub(name astmodel.TypeName, types astmodel.Types) (astmodel.TypeName, error) {
	// Look for the hub step
	result := name
	for {
		hub, err := graph.FindNextType(result, types)
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
// types is a set of known types.
//
func (graph *ConversionGraph) FindNextProperty(
	ref astmodel.PropertyReference,
	types astmodel.Types) (astmodel.PropertyReference, error) {
	nextType, err := graph.FindNextType(ref.DeclaringType(), types)
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
