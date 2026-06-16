/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// AddAzureNameFromConfigStageID is the unique identifier for this pipeline stage
const AddAzureNameFromConfigStageID = "addAzureNameFromConfig"

// AddAzureNameFromConfig creates a Stage that injects the AzureNameFromConfig property into
// Spec types for resources that have $azureNameFromConfig: true in their configuration.
// This allows the Azure name of a resource to be resolved from a ConfigMap at reconciliation time.
func AddAzureNameFromConfig(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		AddAzureNameFromConfigStageID,
		"Injects the AzureNameFromConfig property into Spec types for resources that support it",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			result := make(astmodel.TypeDefinitionSet)

			for _, resource := range defs.AllResources() {
				azureNameFromConfig, ok := configuration.ObjectModelConfiguration.AzureNameFromConfig.Lookup(resource.Name())
				if !ok || !azureNameFromConfig {
					// Not opted in, skip
					continue
				}

				resolved, err := defs.ResolveResourceSpecAndStatus(resource)
				if err != nil {
					return nil, eris.Wrapf(err, "resolving spec and status for %s", resource.Name())
				}

				// Inject the AzureNameFromConfig property into the Spec
				updatedSpecDef, err := injectAzureNameFromConfigProperty(resolved.SpecDef, idFactory)
				if err != nil {
					return nil, eris.Wrapf(err, "injecting AzureNameFromConfig property into %s", resolved.SpecDef.Name())
				}
				result.Add(updatedSpecDef)
			}

			err := configuration.ObjectModelConfiguration.AzureNameFromConfig.VerifyConsumed()
			if err != nil {
				return nil, err
			}

			return state.WithOverlaidDefinitions(result), nil
		},
	)

	stage.RequiresPostrequisiteStages(ApplyDefaulterAndValidatorInterfaceStageID)
	return stage
}

// injectAzureNameFromConfigProperty adds the AzureNameFromConfig property to the provided Spec type definition.
// The GetAzureNameFromConfig method is injected separately on storage specs by InjectAzureNameFromConfigMethod.
func injectAzureNameFromConfigProperty(
	specDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
) (astmodel.TypeDefinition, error) {
	// Create the AzureNameFromConfig property
	propName := idFactory.CreatePropertyName(astmodel.AzureNameFromConfigProperty, astmodel.Exported)
	jsonName := idFactory.CreateStringIdentifier(astmodel.AzureNameFromConfigProperty, astmodel.NotExported)

	prop := astmodel.NewPropertyDefinition(propName, jsonName, astmodel.OptionalConfigMapReferenceType).
		WithDescription("AzureNameFromConfig: if specified, the Azure name of the resource is resolved from this ConfigMap key at reconciliation time, overriding any AzureName specified directly.")

	propInjector := astmodel.NewPropertyInjector()
	updatedDef, err := propInjector.Inject(specDef, prop)
	if err != nil {
		return astmodel.TypeDefinition{}, eris.Wrapf(err, "adding AzureNameFromConfig property to %s", specDef.Name())
	}

	return updatedDef, nil
}
