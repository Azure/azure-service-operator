/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"sort"

	"github.com/pkg/errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// markStorageVersion creates a PipelineStage to mark a particular version as a storage version
func markStorageVersion() PipelineStage {
	return MakePipelineStage(
		"markStorageVersion",
		"Marking the latest version of each resource as the storage version",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			updatedDefs, err := MarkLatestResourceVersionsForStorage(types)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to mark latest resource version as storage version")
			}

			return updatedDefs, nil
		})
}

// MarkLatestResourceVersionsForStorage marks the latest version of each resource as the storage version
func MarkLatestResourceVersionsForStorage(types astmodel.Types) (astmodel.Types, error) {

	result := make(astmodel.Types)
	resourceLookup, err := groupResourcesByVersion(types)
	if err != nil {
		return nil, err
	}

	for _, def := range types {
		// see if it is a resource
		if resourceType, ok := def.Type().(*astmodel.ResourceType); ok {

			unversionedName, err := getUnversionedName(def.Name())
			if err != nil {
				// should never happen as all resources have versioned names
				return nil, err
			}

			allVersionsOfResource := resourceLookup[unversionedName]
			latestVersionOfResource := allVersionsOfResource[len(allVersionsOfResource)-1]

			thisPackagePath := def.Name().PackageReference.PackagePath()
			latestPackagePath := latestVersionOfResource.Name().PackageReference.PackagePath()

			// mark as storage version if it's the latest version
			isLatestVersion := thisPackagePath == latestPackagePath
			if isLatestVersion {
				def = astmodel.MakeTypeDefinition(def.Name(), resourceType.MarkAsStorageVersion()).
					WithDescription(def.Description())
			}
		}
		result.Add(def)
	}

	return result, nil
}

func groupResourcesByVersion(types astmodel.Types) (map[unversionedName][]astmodel.TypeDefinition, error) {

	result := make(map[unversionedName][]astmodel.TypeDefinition)

	for _, def := range types {
		if astmodel.IsResourceDefinition(def) {
			name, err := getUnversionedName(def.Name())
			if err != nil {
				// this should never happen as resources will all have versioned names
				return nil, errors.Wrapf(err, "Unable to extract unversioned name in groupResources")
			}

			result[name] = append(result[name], def)
		}
	}

	// order each set of resources by package name (== by version as these are sortable dates)
	for _, slice := range result {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i].Name().PackageReference.PackageName() < slice[j].Name().PackageReference.PackageName()
		})
	}

	return result, nil
}

func getUnversionedName(name astmodel.TypeName) (unversionedName, error) {

	if localRef, ok := name.PackageReference.AsLocalPackage(); ok {
		group := localRef.Group()
		return unversionedName{group, name.Name()}, nil
	}

	return unversionedName{}, errors.New("cannot get unversioned name from non-local package")
}

type unversionedName struct {
	group string
	name  string
}
