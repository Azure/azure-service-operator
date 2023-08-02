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

var CrossplaneRuntimeV1Package = astmodel.MakeExternalPackageReference("github.com/crossplane/crossplane-runtime/apis/common/v1")

// AddCrossplaneEmbeddedResourceSpec puts an embedded runtimev1alpha1.ResourceSpec on every spec type
func AddCrossplaneEmbeddedResourceSpec(idFactory astmodel.IdentifierFactory) *Stage {
	return NewLegacyStage(
		"addCrossplaneEmbeddedResourceSpec",
		"Add an embedded runtimev1alpha1.ResourceSpec to every spec type",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			specTypeName := astmodel.MakeInternalTypeName(
				CrossplaneRuntimeV1Package,
				idFactory.CreateIdentifier("ResourceSpec", astmodel.Exported))
			embeddedSpec := astmodel.NewPropertyDefinition("", ",inline", specTypeName)

			result := make(astmodel.TypeDefinitionSet)
			for _, typeDef := range definitions {
				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {

					specDef, err := definitions.ResolveResourceSpecDefinition(resource)
					if err != nil {
						return nil, errors.Wrapf(err, "getting resource spec definition")
					}

					// The assumption here is that specs are all Objects
					updatedDef, err := specDef.ApplyObjectTransformation(func(o *astmodel.ObjectType) (astmodel.Type, error) {
						return o.WithEmbeddedProperty(embeddedSpec)
					})
					if err != nil {
						return nil, errors.Wrapf(err, "adding embedded crossplane spec")
					}

					result.Add(typeDef)
					result.Add(updatedDef)
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
