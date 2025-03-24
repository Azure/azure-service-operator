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

func AddDefaulterInterface(
	resourceName astmodel.InternalTypeName,
	webhookDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	defaultFunctions []*functions.DefaultFunction,
) (astmodel.TypeDefinition, error) {
	webhhookObject, ok := webhookDef.Type().(*astmodel.ObjectType)
	if !ok {
		return astmodel.TypeDefinition{}, eris.Errorf("cannot add defaulter interface to non-object type: %s %T", webhookDef.Name(), webhookDef.Type())
	}

	defaulterBuilder := functions.NewDefaulterBuilder(resourceName, idFactory)
	for _, f := range defaultFunctions {
		defaulterBuilder.AddDefault(f)
	}

	webhhookObject = webhhookObject.WithInterface(defaulterBuilder.ToInterfaceImplementation())

	return webhookDef.WithType(webhhookObject), nil
}
