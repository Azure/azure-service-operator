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

func AddDefaulterInterface(
	resourceDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	defaultFunctions []*functions.ResourceFunction,
) (astmodel.TypeDefinition, error) {

	resourceType, ok := resourceDef.Type().(*astmodel.ResourceType)
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("cannot add defaulter interface to non-resource type: %s %T", resourceDef.Name(), resourceDef.Type())
	}

	defaulterBuilder := functions.NewDefaulterBuilder(resourceDef.Name(), resourceType, idFactory)
	for _, f := range defaultFunctions {
		defaulterBuilder.AddDefault(f)
	}

	resourceType = resourceType.WithInterface(defaulterBuilder.ToInterfaceImplementation())

	return resourceDef.WithType(resourceType), nil
}
