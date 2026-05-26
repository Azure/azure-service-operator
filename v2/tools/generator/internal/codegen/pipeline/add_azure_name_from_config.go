/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
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

	stage.RequiresPrerequisiteStages(AddOperatorSpecStageID)
	stage.RequiresPostrequisiteStages(ApplyDefaulterAndValidatorInterfaceStageID)
	return stage
}

// injectAzureNameFromConfigProperty adds the AzureNameFromConfig property and GetAzureNameFromConfig method to
// the provided Spec type definition.
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

	// Create and inject the GetAzureNameFromConfig method
	getAzureNameFromConfigFn := functions.NewObjectFunction(
		astmodel.GetAzureNameFromConfigFunc,
		idFactory,
		getAzureNameFromConfigFunction,
		astmodel.GenRuntimeReference,
	)

	fnInjector := astmodel.NewFunctionInjector()
	updatedDef, err = fnInjector.Inject(updatedDef, getAzureNameFromConfigFn)
	if err != nil {
		return astmodel.TypeDefinition{}, eris.Wrapf(err, "adding GetAzureNameFromConfig method to %s", specDef.Name())
	}

	return updatedDef, nil
}

// getAzureNameFromConfigFunction generates the GetAzureNameFromConfig method body:
//
//	func (spec *MyResource_Spec) GetAzureNameFromConfig() *genruntime.ConfigMapReference {
//	    return spec.AzureNameFromConfig
//	}
func getAzureNameFromConfigFunction(
	k *functions.ObjectFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	configMapRefExpr, err := astmodel.OptionalConfigMapReferenceType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating ConfigMapReference type expression")
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body: astbuilder.Statements(
			astbuilder.Returns(
				astbuilder.Selector(dst.NewIdent(receiverIdent), astmodel.AzureNameFromConfigProperty),
			),
		),
	}

	fn.AddReturn(configMapRefExpr)
	fn.AddComments("returns the AzureNameFromConfig property of the spec, used to resolve the Azure name from a ConfigMap")

	return fn.DefineFunc(), nil
}
