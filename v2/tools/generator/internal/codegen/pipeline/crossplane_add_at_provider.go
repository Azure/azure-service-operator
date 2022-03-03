/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// AddCrossplaneAtProvider adds an "AtProvider" property as the sole property in every resource status
func AddCrossplaneAtProvider(idFactory astmodel.IdentifierFactory) *Stage {
	return NewLegacyStage(
		"addCrossplaneAtProviderProperty",
		"Add an 'AtProvider' property on every status",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			result := make(astmodel.TypeDefinitionSet)
			for _, typeDef := range definitions {
				if _, ok := astmodel.AsResourceType(typeDef.Type()); ok {
					atProviderTypes, err := nestStatusIntoAtProvider(
						idFactory, definitions, typeDef)
					if err != nil {
						return nil, errors.Wrapf(err, "creating AtProvider definitions")
					}

					// Allow duplicates here because some resources share the same _Status type
					// which means it'll get processed multiple times. That's OK as long as it looks
					// the same though.
					err = result.AddAllAllowDuplicates(atProviderTypes)
					if err != nil {
						return nil, err
					}
				}
			}

			unmodified := definitions.Except(result)
			result.AddTypes(unmodified)

			return result, nil
		})
}

// nestStatusIntoAtProvider returns the type definitions required to nest the contents of the "Status" type
// into a property named "AtProvider" whose type is "<name>Observation"
func nestStatusIntoAtProvider(
	idFactory astmodel.IdentifierFactory,
	definitions astmodel.TypeDefinitionSet,
	typeDef astmodel.TypeDefinition) ([]astmodel.TypeDefinition, error) {

	resource, ok := astmodel.AsResourceType(typeDef.Type())
	if !ok {
		return nil, errors.Errorf("provided typeDef was not a resourceType, instead %T", typeDef.Type())
	}
	resourceName := typeDef.Name()

	statusType := astmodel.IgnoringErrors(resource.StatusType())
	if statusType == nil {
		return nil, nil // TODO: Some definitions don't have status yet
	}

	statusName, ok := astmodel.AsTypeName(resource.StatusType())
	if !ok {
		return nil, errors.Errorf("resource %q status was not of type TypeName, instead: %T", resourceName, resource.StatusType())
	}

	// In the case where a status type is reused across multiple resource definitions, we need to make sure
	// to generate the same names for all of their nested properties, so base the nested type name off the
	// status type name
	nestedTypeName := strings.Split(statusName.Name(), "_")[0] + "Observation"
	nestedPropertyName := "AtProvider"
	return nestType(idFactory, definitions, statusName, nestedTypeName, nestedPropertyName)
}
