/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// MarkLatestAPIVersionAsStorageVersionId is the unique identifier for this pipeline stage
const MarkLatestAPIVersionAsStorageVersionId = "markStorageVersion"

// MarkLatestAPIVersionAsStorageVersion creates a Stage to mark a particular version as a storage version
func MarkLatestAPIVersionAsStorageVersion() *Stage {
	return NewLegacyStage(
		MarkLatestAPIVersionAsStorageVersionId,
		"Mark the latest API version of each resource as the storage version",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			updatedDefs, err := MarkLatestResourceVersionsForStorage(definitions)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to mark latest resource version as storage version")
			}

			return updatedDefs, nil
		})
}

// MarkLatestResourceVersionsForStorage marks the latest version of each resource as the storage version
func MarkLatestResourceVersionsForStorage(definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)
	resourceLookup := groupResourcesByVersion(definitions)
	for _, def := range definitions {
		// see if it is a resource
		if resourceType, ok := def.Type().(*astmodel.ResourceType); ok {

			unversionedName := getUnversionedName(def.Name())
			allVersionsOfResource := resourceLookup[unversionedName]
			latestVersionOfResource := allVersionsOfResource[len(allVersionsOfResource)-1]

			thisPackagePath := def.Name().PackageReference().ImportPath()
			latestPackagePath := latestVersionOfResource.Name().PackageReference().ImportPath()

			// mark as storage version if it's the latest version
			isLatestVersion := thisPackagePath == latestPackagePath
			if isLatestVersion {
				def = astmodel.MakeTypeDefinition(def.Name(), resourceType.MarkAsStorageVersion()).
					WithDescription(def.Description()...)
			}
		}
		result.Add(def)
	}

	return result, nil
}

func groupResourcesByVersion(definitions astmodel.TypeDefinitionSet) map[unversionedName][]astmodel.TypeDefinition {
	result := make(map[unversionedName][]astmodel.TypeDefinition)

	for _, def := range definitions {

		// We want to explicitly avoid storage definitions, as this approach for flagging the hub version is
		// used when we aren't leveraging the conversions between storage versions.
		if astmodel.IsStoragePackageReference(def.Name().PackageReference()) {
			continue
		}

		if astmodel.IsResourceDefinition(def) {
			name := getUnversionedName(def.Name())
			result[name] = append(result[name], def)
		}
	}

	// order each set of resources by package name (== by version as these are sortable dates)
	for _, slice := range result {
		slices.SortFunc(
			slice,
			func(left astmodel.TypeDefinition, right astmodel.TypeDefinition) int {
				return astmodel.ComparePathAndVersion(
					left.Name().PackageReference().ImportPath(),
					right.Name().PackageReference().ImportPath())
			})
	}

	return result
}

func getUnversionedName(name astmodel.InternalTypeName) unversionedName {
	ref := name.InternalPackageReference()
	group := ref.Group()
	return unversionedName{
		group: group,
		name:  name.Name(),
	}
}

type unversionedName struct {
	group string
	name  string
}
