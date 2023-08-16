/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// FlattenResources flattens any resources directly inside other resources
func FlattenResources() *Stage {
	return NewLegacyStage(
		"flattenResources",
		"Flatten nested resource types",
		func(ctx context.Context, defs astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			flattenEachResource := func(
				this *astmodel.TypeVisitor[any],
				it *astmodel.ResourceType,
				ctx any,
			) (astmodel.Type, error) {
				// visit inner types:
				visited, err := astmodel.IdentityVisitOfResourceType(this, it, ctx)
				if err != nil {
					return nil, err
				}

				newResource, err := flattenResource(defs, visited, 0)
				if err != nil {
					return nil, err
				}

				return newResource, nil
			}

			v := astmodel.TypeVisitorBuilder[any]{
				VisitResourceType: flattenEachResource,
			}.Build()

			results := make(astmodel.TypeDefinitionSet)

			for _, def := range defs {

				result, err := v.VisitDefinition(def, nil)
				if err != nil {
					return nil, errors.Wrapf(err, "error processing type %s", def.Name())
				}

				results.Add(result)
			}

			return results, nil
		})
}

func flattenResource(defs astmodel.TypeDefinitionSet, t astmodel.Type, depth int) (astmodel.Type, error) {
	if depth > 10 {
		return nil, errors.Errorf("too many levels of nesting while flattening resource")
	}

	if resource, ok := t.(*astmodel.ResourceType); ok {

		changed := false

		// resolve spec
		{
			specType, err := defs.FullyResolve(resource.SpecType())
			if err != nil {
				return nil, err
			}

			if specResource, ok := specType.(*astmodel.ResourceType); ok {
				resource = resource.WithSpec(specResource.SpecType())
				changed = true
			}
		}

		// resolve status
		{
			statusType, err := defs.FullyResolve(resource.StatusType())
			if err != nil {
				return nil, err
			}

			if statusResource, ok := statusType.(*astmodel.ResourceType); ok {
				resource = resource.WithStatus(statusResource.StatusType())
				changed = true
			}
		}

		if changed {
			t = resource
			// do it again, can be multiply nested
			return flattenResource(defs, t, depth+1)
		}
	}

	return t, nil
}
