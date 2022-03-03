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
	definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinition, error) {

	resolved, err := definitions.ResolveResourceSpecAndStatus(resourceDef)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "unable to resolve resource %s", resourceDef.Name())
	}

	defaulterBuilder := functions.NewDefaulterBuilder(resourceDef.Name(), resolved.ResourceType, idFactory)

	// Determine if the resource has a SetName function
	if resolved.SpecType.HasFunctionWithName(astmodel.SetAzureNameFunc) {
		defaulterBuilder.AddDefault(functions.NewDefaultAzureNameFunction(resolved.ResourceType, idFactory))
	}

	resourceType := resolved.ResourceType.WithInterface(defaulterBuilder.ToInterfaceImplementation())

	return resourceDef.WithType(resourceType), nil
}
