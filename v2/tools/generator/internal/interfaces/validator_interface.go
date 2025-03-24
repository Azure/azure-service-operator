/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package interfaces

import (
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

func AddValidatorInterface(
	resourceName astmodel.InternalTypeName,
	webhookDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	validations map[functions.ValidationKind][]*functions.ValidateFunction,
) (astmodel.TypeDefinition, error) {
	webhhookObject, ok := webhookDef.Type().(*astmodel.ObjectType)
	if !ok {
		return astmodel.TypeDefinition{}, eris.Errorf("cannot add validator interface to non-object type: %s %T", webhookDef.Name(), webhookDef.Type())
	}

	validatorBuilder := functions.NewValidatorBuilder(resourceName, idFactory)
	for validationKind, vs := range validations {
		for _, validation := range vs {
			validatorBuilder.AddValidation(validationKind, validation)
		}
	}

	webhhookObject = webhhookObject.WithInterface(validatorBuilder.ToInterfaceImplementation())
	return webhookDef.WithType(webhhookObject), nil
}
