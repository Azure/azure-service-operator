/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package interfaces

import (
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

func AddValidatorInterface(
	resourceDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinition, error) {

	resolved, err := definitions.ResolveResourceSpecAndStatus(resourceDef)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "unable to resolve resource %s", resourceDef.Name())
	}

	validatorBuilder := functions.NewValidatorBuilder(resourceDef.Name(), resolved.ResourceType, idFactory)
	validatorBuilder.AddValidation(functions.ValidationKindCreate, functions.NewValidateResourceReferencesFunction(resolved.ResourceType, idFactory))
	validatorBuilder.AddValidation(functions.ValidationKindUpdate, functions.NewValidateResourceReferencesFunction(resolved.ResourceType, idFactory))

	resourceType := resolved.ResourceType.WithInterface(validatorBuilder.ToInterfaceImplementation())

	return resourceDef.WithType(resourceType), nil
}
