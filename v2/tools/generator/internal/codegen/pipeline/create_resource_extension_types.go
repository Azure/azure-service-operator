/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// CreateResourceExtensionsStageID is the unique identifier of this stage
const (
	CreateResourceExtensionsStageID = "createResourceExtensions"
)

func CreateResourceExtensions(localPath string, idFactory astmodel.IdentifierFactory) Stage {
	stage := MakeStage(
		CreateResourceExtensionsStageID,
		"Create Resource Extensions for each resource type",
		func(ctx context.Context, state *State) (*State, error) {

			// Map of the new extension types, to all the resource types names on which the extension applies to
			extendedResourceTypesMapping := make(map[astmodel.TypeName][]astmodel.TypeName)
			extendedResourceDefs := make(astmodel.TypeDefinitionSet)
			resourceDefs := astmodel.FindResourceDefinitions(state.definitions)

			// Iterate through resource types and aggregate the resource types that share the same extension type in a map.
			for _, typeDef := range resourceDefs {
				group, _, _ := typeDef.Name().PackageReference.GroupVersion()
				packageRef := astmodel.MakeLocalPackageReference(localPath, group, "customizations")
				extensionTypeName := astmodel.MakeTypeName(packageRef, typeDef.Name().Name()+"Extension")
				extendedResourceTypesMapping[extensionTypeName] = append(extendedResourceTypesMapping[extensionTypeName], typeDef.Name())
			}

			// For each resource extension type, create a registration function
			for extensionName, extendedResources := range extendedResourceTypesMapping {
				fn := functions.NewGetExtendedResourcesFunction(idFactory, extendedResources)

				newExtensionType := astmodel.MakeTypeDefinition(
					extensionName,
					astmodel.NewObjectType().WithFunction(fn))

				if err := extendedResourceDefs.AddAllowDuplicates(newExtensionType); err != nil {
					return nil, err
				}

			}
			state.definitions.AddTypes(extendedResourceDefs)
			return state, nil
		})

	// We don't want tests to be generated for resourceExtensions, since these are not the actual resource types.
	// Therefore, we want to make sure that 'createResourceExtensions' stage only runs when these prerequisite
	// stages have completed.
	return stage.RequiresPrerequisiteStages(InjectJsonSerializationTestsID, InjectPropertyAssignmentTestsID)
}
