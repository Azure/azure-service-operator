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

// AddCrossplaneEmbeddedResourceStatus puts an embedded runtimev1alpha1.ResourceStatus on every spec type
func AddCrossplaneEmbeddedResourceStatus(idFactory astmodel.IdentifierFactory) *Stage {
	return NewLegacyStage(
		"addCrossplaneEmbeddedResourceStatus",
		"Add an embedded runtimev1alpha1.ResourceStatus to every status type",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			statusTypeName := astmodel.MakeTypeName(
				CrossplaneRuntimeV1Alpha1Package,
				idFactory.CreateIdentifier("ResourceStatus", astmodel.Exported))
			embeddedStatus := astmodel.NewPropertyDefinition("", ",inline", statusTypeName)

			result := make(astmodel.TypeDefinitionSet)
			for _, typeDef := range definitions {
				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {

					if astmodel.IgnoringErrors(resource.StatusType()) == nil {
						continue
					}

					statusDef, err := definitions.ResolveResourceStatusDefinition(resource)
					if err != nil {
						return nil, errors.Wrapf(err, "getting resource status definition")
					}

					// The assumption here is that specs are all Objects
					updatedDef, err := statusDef.ApplyObjectTransformation(func(o *astmodel.ObjectType) (astmodel.Type, error) {
						return o.WithEmbeddedProperty(embeddedStatus)
					})
					if err != nil {
						return nil, errors.Wrapf(err, "adding embedded crossplane status")
					}

					result.Add(typeDef)
					// Allow duplicates here because some resources share the same _Status type
					// which means it'll get processed multiple times. That's OK as long as it looks
					// the same though.
					err = result.AddAllowDuplicates(updatedDef)
					if err != nil {
						return nil, err
					}
				}
			}

			for _, typeDef := range definitions {
				if !result.Contains(typeDef.Name()) {
					result.Add(typeDef)
				}
			}

			return result, nil
		})
}
