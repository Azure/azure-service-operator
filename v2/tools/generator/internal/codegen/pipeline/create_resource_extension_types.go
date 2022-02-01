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
const CreateResourceExtensionsStageID = "createResourceExtensions"

func CreateResourceExtensions(localPath string, idFactory astmodel.IdentifierFactory) Stage {
	return MakeLegacyStage(
		CreateResourceExtensionsStageID,
		"create Resource Extensions for each resource type",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			extendedResourceTypes := make(astmodel.Types)
			extendedResourceTypesMapping := make(map[astmodel.TypeName][]astmodel.TypeName)
			resourcePackageRef := make(map[astmodel.TypeName][]astmodel.PackageReference)
			resourceTypes := astmodel.FindResourceTypes(types)

			//iterate through resource types and aggregate the similar resource types and package refs in a map.
			for _, typeDef := range resourceTypes {
				group, _, _ := typeDef.Name().PackageReference.GroupVersion()
				packageRef := astmodel.MakeLocalPackageReference(localPath, group, "extensions")
				name := astmodel.MakeTypeName(packageRef, typeDef.Name().Name()+"Extension")

				extendedResourceTypesMapping[name] = append(extendedResourceTypesMapping[name], typeDef.Name())
				//putting package references in map for each extension
				resourcePackageRef[name] = append(resourcePackageRef[name], typeDef.Name().PackageReference)
			}

			//iterate through the extendedResources map and create a ResourceExtension type
			for extensionName, extendedResources := range extendedResourceTypesMapping {
				fn := functions.NewGetExtendedResourcesFunction(idFactory, extendedResources, resourcePackageRef[extensionName])

				newType := astmodel.MakeTypeDefinition(
					extensionName,
					astmodel.NewObjectType().WithFunction(fn))

				extendedResourceTypes.AddAllowDuplicates(newType)
			}
			types.AddTypes(extendedResourceTypes)
			return types, nil
		})
}
