/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

func CreateResourceExtensions(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) Stage {
	return MakeLegacyStage(
		"createResourceExtensions",
		"create Resource Extensions for each resource type",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			extendedResourceTypes := make(astmodel.Types)
			resourceTypes := astmodel.FindResourceTypes(types)

			for _, typeDef := range resourceTypes {

				if !extendedResourceTypes.Contains(typeDef.Name()) {

					group, _, _ := typeDef.Name().PackageReference.GroupVersion()
					packageRef := astmodel.MakeLocalPackageReference(configuration.LocalPathPrefix(), group, "extensions")
					name := astmodel.MakeTypeName(packageRef, typeDef.Name().Name()+"Extension")

					newType := astmodel.MakeTypeDefinition(
						name,
						astmodel.NewObjectType().WithFunction(functions.NewGetExtendedResourcesFunction(idFactory)))

					extendedResourceTypes.AddAllowDuplicates(newType)
				}
			}
			types.AddTypes(extendedResourceTypes)
			return types, nil
		})
}
