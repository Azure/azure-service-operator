/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// addCrossplaneAtProvider adds an "AtProvider" property as the sole property in every resource status
func addCrossplaneAtProvider(idFactory astmodel.IdentifierFactory) PipelineStage {

	return MakePipelineStage(
		"addCrossplaneAtProviderProperty",
		"Adds an 'AtProvider' property on every status",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			result := make(astmodel.Types)
			for _, typeDef := range types {
				if _, ok := astmodel.AsResourceType(typeDef.Type()); ok {
					atProviderTypes, err := nestStatusIntoAtProvider(
						idFactory, types, typeDef)
					if err != nil {
						return nil, errors.Wrapf(err, "creating AtProvider types")
					}
					result.AddAll(atProviderTypes)
				}
			}

			unmodified := types.Except(result)
			result.AddTypes(unmodified)

			return result, nil
		})
}

// nestStatusIntoAtProvider returns the type definitions required to nest the contents of the "Status" type
// into a property named "AtProvider" whose type is "<name>Observation"
func nestStatusIntoAtProvider(
	idFactory astmodel.IdentifierFactory,
	types astmodel.Types,
	typeDef astmodel.TypeDefinition) ([]astmodel.TypeDefinition, error) {

	resource, ok := astmodel.AsResourceType(typeDef.Type())
	if !ok {
		return nil, errors.Errorf("provided typeDef was not a resourceType, instead %T", typeDef.Type())
	}
	resourceName := typeDef.Name()

	statusType := astmodel.IgnoringErrors(resource.StatusType())
	if statusType == nil {
		return nil, nil // TODO: Some types don't have status yet
	}

	statusName, ok := astmodel.AsTypeName(resource.StatusType())
	if !ok {
		return nil, errors.Errorf("resource %q status was not of type TypeName, instead: %T", resourceName, resource.StatusType())
	}

	nestedTypeName := resourceName.Name() + "Observation"
	nestedPropertyName := "AtProvider"
	return nestType(idFactory, types, statusName, nestedTypeName, nestedPropertyName)
}
