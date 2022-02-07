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
	return MakeLegacyStage(
		CreateResourceExtensionsStageID,
		"Create Resource Extensions for each resource type",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			// Map of the new extension types, to all the resource types names on which the extension applies to
			extendedResourceTypesMapping := make(map[astmodel.TypeName][]astmodel.TypeName)
			extendedResourceTypes := make(astmodel.Types)
			resourceTypes := astmodel.FindResourceTypes(types)

			// Iterate through resource types and aggregate the resource types that share the same extension in a map.
			for _, typeDef := range resourceTypes {
				group, _, _ := typeDef.Name().PackageReference.GroupVersion()
				packageRef := astmodel.MakeLocalPackageReference(localPath, group, "extensions")
				extensionTypeName := astmodel.MakeTypeName(packageRef, typeDef.Name().Name()+"Extension")
				extendedResourceTypesMapping[extensionTypeName] = append(extendedResourceTypesMapping[extensionTypeName], typeDef.Name())
			}

			// Iterate through the extendedResources map and create a ResourceExtension type
			for extensionName, extendedResources := range extendedResourceTypesMapping {
				fn := functions.NewGetExtendedResourcesFunction(idFactory, extendedResources)

				newType := astmodel.MakeTypeDefinition(
					extensionName,
					astmodel.NewObjectType().WithFunction(fn))

				if err := extendedResourceTypes.AddAllowDuplicates(newType); err != nil {
					return nil, err
				}

			}
			types.AddTypes(extendedResourceTypes)
			return types, nil
		})
}
