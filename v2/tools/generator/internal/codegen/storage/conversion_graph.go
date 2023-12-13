/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	"os"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// ConversionGraph builds up a set of graphs of the required conversions between versions
// For each group (e.g. microsoft.storage or microsoft.batch) we have a separate subgraph of directed conversions
type ConversionGraph struct {
	configuration *config.ObjectModelConfiguration
	subGraphs     map[string]*GroupConversionGraph // Map of group name to subgraph
}

// LookupTransition accepts a type name and looks up the transition to the next version in the graph
// Returns the next version and true if it's found, or an empty type name and false if not.
func (graph *ConversionGraph) LookupTransition(name astmodel.InternalTypeName) astmodel.InternalTypeName {
	// Expect to get either a local or a storage reference, not an external one
	group := name.InternalPackageReference().Group()
	subgraph, ok := graph.subGraphs[group]
	if !ok {
		return astmodel.InternalTypeName{}
	}

	return subgraph.LookupTransition(name)
}

// FindNextType returns the type name of the next closest type on the path to the hub type.
// Returns the type name and no error if the next type is found; empty name and no error if not; empty name and an error if
// something goes wrong.
// If the name passed in is for the hub type for the given resource, no next type will be found.
// This is used to identify the next type needed for property assignment functions, and is a building block for
// identification of hub definitions.
func (graph *ConversionGraph) FindNextType(
	name astmodel.InternalTypeName,
	definitions astmodel.TypeDefinitionSet,
) (astmodel.InternalTypeName, error) {
	group := name.InternalPackageReference().Group()
	subgraph, ok := graph.subGraphs[group]
	if !ok {
		return astmodel.InternalTypeName{}, nil
	}

	// Look for a next type with the same name
	nextType := subgraph.LookupTransition(name)

	// Look for a renamed type with the same name
	renamedType, err := subgraph.searchForRenamedType(name, definitions)
	if err != nil {
		// Something went wrong
		return astmodel.InternalTypeName{}, errors.Wrapf(err, "searching for type renamed from %s", name)
	}

	// If we have no renamed type, return the next type (if any)
	if renamedType.IsEmpty() {
		return nextType, nil
	}

	// If we have no next type, return the renamed type (if any)
	if nextType.IsEmpty() {
		return renamedType, nil
	}

	// We have both a next type and a renamed type
	// If they're in the same package, the type-rename has been configured on the wrong version (or the wrong type)
	if nextType.PackageReference().Equals(renamedType.PackageReference()) {
		return astmodel.InternalTypeName{},
			errors.Errorf("confict between rename of %s to %s and existing type %s", name, renamedType, nextType)
	}

	// Now we need to return the earlier type. We can do this by comparing the package paths.
	// (this be needed if a different type is introduced with the same name in a later version, or if a type is
	// renamed in one version and renamed back in a later one)
	if astmodel.ComparePathAndVersion(
		nextType.PackageReference().ImportPath(),
		renamedType.PackageReference().ImportPath()) < 0 {
		// nextType came first
		return nextType, nil
	}

	return renamedType, nil
}

// FindHub returns the type name of the hub resource, given the type name of one of the resources that is
// persisted using that hub type. This is done by following links in the conversion graph until we either reach the end
// or we find that a newer version of the type does not exist.
// Returns the hub type if found; an empty name and an error if not.
func (graph *ConversionGraph) FindHub(
	name astmodel.InternalTypeName,
	definitions astmodel.TypeDefinitionSet,
) (astmodel.InternalTypeName, error) {
	result, _, err := graph.FindHubAndDistance(name, definitions)
	return result, err
}

// FindHubAndDistance returns the type name of the hub resource, given the type name of one of the resources that is
// persisted using that hub type. This is done by following links in the conversion graph until we either reach the end
// or we find that a newer version of the type does not exist.
// Returns the distance if a hub was found, an error if not.
func (graph *ConversionGraph) FindHubAndDistance(
	name astmodel.InternalTypeName,
	definitions astmodel.TypeDefinitionSet,
) (astmodel.InternalTypeName, int, error) {
	// Look for the hub step
	result := name
	distance := 0
	for {
		hub, err := graph.FindNextType(result, definitions)
		if err != nil {
			return astmodel.InternalTypeName{},
				-1,
				errors.Wrapf(err, "finding hub for %s",
					name)
		}

		if hub.IsEmpty() {
			break
		}

		result = hub
		distance++
	}

	return result, distance, nil
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

func (graph *ConversionGraph) String(group string, kind string) (string, error) {
	var content bytes.Buffer
	err := graph.WriteTo(group, kind, &content)
	if err != nil {
		return "", errors.Wrapf(err, "writing conversion graph")
	}

	return content.String(), nil
}

func (graph *ConversionGraph) SaveTo(group string, kind string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return errors.Wrapf(err, "creating %s", filename)
	}

	defer file.Close()

	err = graph.WriteTo(group, kind, file)
	if err != nil {
		return errors.Wrapf(err, "writing conversion graph to %s", filename)
	}

	return nil
}

// WriteTo gives a debug dump of the conversion graph for a particular type name
func (graph *ConversionGraph) WriteTo(group string, kind string, writer io.Writer) error {
	subgraph, ok := graph.subGraphs[group]
	if !ok {
		return nil
	}

	return subgraph.WriteTo(kind, writer)
}
