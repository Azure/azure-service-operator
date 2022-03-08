/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const CreateTypesForBackwardCompatibilityID = "createTypesForBackwardCompatibility"

// CreateTypesForBackwardCompatibility returns a pipeline stage that creates copies of storage types in other packages to
// provide backward compatibility with previous releases of Azure Service Operator.
// Backward compatibility versions are created for all storage versions to allow users of older versions of the operator
// to easily upgrade. We lean on the fact that storage versions have already been modified to increase compatibility -
// everything is optional, all enums have been rendered into their underlying primitive types, and so on.
func CreateTypesForBackwardCompatibility(prefix string) *Stage {
	stage := NewStage(
		CreateTypesForBackwardCompatibilityID,
		"Create clones of CRD storage types for backward compatibility with prior ASO versions",
		func(ctx context.Context, state *State) (*State, error) {
			// Predicate to filter to only storage types
			isStorageType := func(def astmodel.TypeDefinition) bool {
				return astmodel.IsStoragePackageReference(def.Name().PackageReference)
			}

			// Filter to the types we want to process
			typesToConvert := state.Definitions().Where(isStorageType)

			// Update the description of each to reflect purpose
			typesWithDescriptions := addCompatibilityComments(typesToConvert)

			// Work out the new names for all our new types
			renames := createBackwardCompatibilityRenameMap(typesToConvert, prefix)

			// Rename all the types
			visitor := astmodel.NewRenamingVisitor(renames)

			renamed, err := visitor.RenameAll(typesWithDescriptions)
			if err != nil {
				return nil, errors.Wrap(err, "creating types for backward compatibility")
			}

			// Add the new types into our state
			defs := state.Definitions()
			defs.AddTypes(renamed)

			return state.WithDefinitions(defs), nil
		})

	stage.RequiresPrerequisiteStages(CreateStorageTypesStageID)
	stage.RequiresPostrequisiteStages(CreateConversionGraphStageId)
	return stage
}

func addCompatibilityComments(defs astmodel.TypeDefinitionSet) astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)

	for _, def := range defs {
		name := def.Name()
		desc := []string{
			fmt.Sprintf(
				"Backward compatibility type for %s.%s",
				name.PackageReference.PackageName(),
				name.Name()),
		}

		result.Add(def.WithDescription(desc))
	}

	return result
}

func createBackwardCompatibilityRenameMap(
	set astmodel.TypeDefinitionSet,
	versionPrefix string) map[astmodel.TypeName]astmodel.TypeName {
	result := make(map[astmodel.TypeName]astmodel.TypeName)

	for name := range set {
		if _, ok := result[name]; !ok {
			newName := createBackwardCompatibilityRename(name, versionPrefix)
			result[name] = newName
		}
	}

	return result
}

func createBackwardCompatibilityRename(name astmodel.TypeName, versionPrefix string) astmodel.TypeName {
	ref := name.PackageReference.(astmodel.StoragePackageReference)
	local := ref.Local().WithVersionPrefix(versionPrefix)
	newRef := astmodel.MakeStoragePackageReference(local)
	return name.WithPackageReference(newRef)
}
