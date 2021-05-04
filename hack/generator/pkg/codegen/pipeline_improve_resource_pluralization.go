/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// improveResourcePluralization improves pluralization for resources
func improveResourcePluralization() PipelineStage {

	return MakePipelineStage(
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
			return fixNameReferences(result, renames)
		})
}

func fixNameReferences(types astmodel.Types, renames map[astmodel.TypeName]astmodel.TypeName) (astmodel.Types, error) {
	result := make(astmodel.Types)

	// On the off chance that something is referring to a top level resource type which was renamed (rare but possible) go fix up the references
	fixName := func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
		if newName, ok := renames[it]; ok {
			return astmodel.IdentityVisitOfTypeName(this, newName, ctx)
		}

		return astmodel.IdentityVisitOfTypeName(this, it, ctx)
	}

	visitor := astmodel.TypeVisitorBuilder{
		VisitTypeName: fixName,
	}.Build()

	for _, def := range types {
		updated, err := visitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, err
		}

		result.Add(updated)
	}

	return result, nil
}
