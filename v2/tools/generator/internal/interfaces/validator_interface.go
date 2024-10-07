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
	definitions astmodel.TypeDefinitionSet,
	validations map[functions.ValidationKind][]*functions.ResourceFunction,
) (astmodel.TypeDefinition, error) {
	rt, ok := astmodel.AsResourceType(resourceDef.Type())
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("unable to resolve resource %s", resourceDef.Name())
	}

	validatorBuilder := functions.NewValidatorBuilder(resourceDef.Name(), rt, idFactory)
	for validationKind, vs := range validations {
		for _, validation := range vs {
			validatorBuilder.AddValidation(validationKind, validation)
		}
	}

	injector := astmodel.NewInterfaceInjector()
	def, err := injector.Inject(resourceDef, validatorBuilder.ToInterfaceImplementation())

	return def, err
}
