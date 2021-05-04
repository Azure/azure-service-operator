/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/pkg/errors"
)

// flattenResources flattens any resources directly inside other resources
func flattenResources() PipelineStage {
	return MakePipelineStage(
		"flatten-resources",
		"Flatten nested resource types",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {

			flattenEachResource := func(this *astmodel.TypeVisitor, it *astmodel.ResourceType, ctx interface{}) (astmodel.Type, error) {
				// visit inner types:
				visited, err := astmodel.IdentityVisitOfResourceType(this, it, ctx)
				if err != nil {
					return nil, err
				}

				newResource, err := flattenResource(defs, visited)
				if err != nil {
					return nil, err
				}

				return newResource, nil
			}

			v := astmodel.TypeVisitorBuilder{
				VisitResourceType: flattenEachResource,
			}.Build()

			results := make(astmodel.Types)

			for _, def := range defs {

				result, err := v.VisitDefinition(def, nil)
				if err != nil {
					return nil, errors.Wrapf(err, "error processing type %v", def.Name())
				}

				results.Add(result)
			}

			return results, nil
		})
}

func flattenResource(defs astmodel.Types, t astmodel.Type) (astmodel.Type, error) {
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
			return flattenResource(defs, t)
		}
	}

	return t, nil
}
