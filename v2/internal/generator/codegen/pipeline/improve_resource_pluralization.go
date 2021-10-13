/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
)

// ImproveResourcePluralization improves pluralization for resources
func ImproveResourcePluralization() Stage {
	stage := MakeLegacyStage(
		"pluralizeNames",
		"Improve resource pluralization",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			result := make(astmodel.Types)

			renames := make(map[astmodel.TypeName]astmodel.TypeName)

			for _, typeDef := range types {
				if resourceType, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					newTypeName := typeDef.Name().Singular()
					// check if there is already a resource with this name
					if _, ok := types[newTypeName]; !ok {
						// not found: rename the resource
						renames[typeDef.Name()] = newTypeName
						typeDef = typeDef.WithName(newTypeName)
					}

					// Need to update owner ref too if applicable
					if resourceType.Owner() != nil {
						owner := resourceType.Owner().Singular()
						resourceType = resourceType.WithOwner(&owner)
						typeDef = typeDef.WithType(resourceType)
					}

					result.Add(typeDef)
				} else {
					result.Add(typeDef)
				}
			}

			// On the off chance that one of the names we changed is referenced someplace, fix it up. This is pretty
			// rare since usually resources don't refer directly to other resources, but there are a few places it does happen.
			renamingVisitor := astmodel.NewRenamingVisitor(renames)
			return renamingVisitor.RenameAll(result)
		})

	stage.RequiresPrerequisiteStages(RemoveTypeAliasesStageID)
	return stage
}
