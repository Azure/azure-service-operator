/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

const AddOperatorSpecStageID = "addOperatorSpec"

func AddOperatorSpec(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) *Stage {
	return NewStage(
		AddOperatorSpecStageID,
		"Adds the property 'OperatorSpec' to all Spec types that require it",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			result := make(astmodel.TypeDefinitionSet)

			for _, resource := range astmodel.FindResourceDefinitions(defs) {
				resolved, err := defs.ResolveResourceSpecAndStatus(resource)
				if err != nil {
					return nil, errors.Wrapf(err, "resolving resource spec and status for %s", resource.Name())
				}

				newDefs, err := createOperatorSpecIfNeeded(configuration, idFactory, resolved)
				if err != nil {
					return nil, err
				}
				result.AddTypes(newDefs)
			}

			// confirm that all the Azure generated secrets were used. Note that this also indirectly confirms that
			// this property was only used on resources, since that's the only place we try to check it from. If it's
			// set on anything else it will be labeled unconsumed.
			err := configuration.ObjectModelConfiguration.VerifyAzureGeneratedSecretsConsumed()
			if err != nil {
				return nil, err
			}

			result = defs.OverlayWith(result)

			return state.WithDefinitions(result), nil
		})
}

func createOperatorSpecIfNeeded(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory,
	resolved *astmodel.ResolvedResourceDefinition) (astmodel.TypeDefinitionSet, error) {

	secrets, err := configuration.ObjectModelConfiguration.AzureGeneratedSecrets(resolved.ResourceDef.Name())
	if err != nil {
		if config.IsNotConfiguredError(err) {
			// In this case, error is OK and just means we don't need to make an OperatorSpec type
			return nil, nil
		}
		return nil, errors.Wrapf(err, "reading azureGeneratedSecrets for %s", resolved.ResourceDef.Name())
	}

	builder := newOperatorSpecBuilder(configuration, idFactory, resolved.ResourceDef)
	operatorSpec := builder.newEmptyOperatorSpec()
	operatorSpec = builder.addSecretsToOperatorSpec(operatorSpec, secrets)

	propInjector := astmodel.NewPropertyInjector()
	updatedDef, err := propInjector.Inject(resolved.SpecDef, builder.newOperatorSpecProperty(operatorSpec))
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't add OperatorSpec to spec %q", resolved.SpecDef.Name())
	}

	result := make(astmodel.TypeDefinitionSet)

	result.Add(updatedDef)
	result.Add(operatorSpec)
	result.AddTypes(builder.definitions) // Add any other types that were needed as well

	return result, nil
}

type operatorSpecBuilder struct {
	idFactory     astmodel.IdentifierFactory
	configuration *config.Configuration
	resource      astmodel.TypeDefinition
	definitions   astmodel.TypeDefinitionSet
}

func newOperatorSpecBuilder(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory,
	resource astmodel.TypeDefinition) *operatorSpecBuilder {
	return &operatorSpecBuilder{
		idFactory:     idFactory,
		configuration: configuration,
		resource:      resource,
		definitions:   make(astmodel.TypeDefinitionSet),
	}
}

func (b *operatorSpecBuilder) newEmptyOperatorSpec() astmodel.TypeDefinition {
	name := b.idFactory.CreateIdentifier(b.resource.Name().Name()+"OperatorSpec", astmodel.Exported)
	operatorSpecTypeName := b.resource.Name().WithName(name)
	operatorSpec := astmodel.NewObjectType()

	operatorSpecDefinition := astmodel.MakeTypeDefinition(operatorSpecTypeName, operatorSpec)
	description := "Details for configuring operator behavior. Fields in this struct are " +
		"interpreted by the operator directly rather than being passed to Azure"
	operatorSpecDefinition = operatorSpecDefinition.WithDescription([]string{description})

	return operatorSpecDefinition
}

func (b *operatorSpecBuilder) newOperatorSpecProperty(operatorSpec astmodel.TypeDefinition) *astmodel.PropertyDefinition {
	prop := astmodel.NewPropertyDefinition(
		astmodel.OperatorSpecProperty,
		b.idFactory.CreateIdentifier(astmodel.OperatorSpecProperty, astmodel.NotExported),
		operatorSpec.Name()).MakeTypeOptional()
	desc := "The specification for configuring operator behavior. " +
		"This field is interpreted by the operator and not passed directly to Azure"
	prop = prop.WithDescription(desc)

	return prop
}

func (b *operatorSpecBuilder) newSecretsProperty(secretsTypeName astmodel.TypeName) *astmodel.PropertyDefinition {
	secretProp := astmodel.NewPropertyDefinition(
		b.idFactory.CreatePropertyName(astmodel.OperatorSpecSecretsProperty, astmodel.Exported),
		b.idFactory.CreateIdentifier(astmodel.OperatorSpecSecretsProperty, astmodel.NotExported),
		secretsTypeName)
	secretProp = secretProp.WithDescription("configures where to place Azure generated secrets.")
	secretProp = secretProp.MakeTypeOptional()

	return secretProp
}

func (b *operatorSpecBuilder) addSecretsToOperatorSpec(
	operatorSpecDef astmodel.TypeDefinition,
	azureGeneratedSecrets []string) astmodel.TypeDefinition {

	operatorSpec, ok := astmodel.AsObjectType(operatorSpecDef.Type())
	if !ok {
		panic(fmt.Sprintf("OperatorSpec %q was not an ObjectType, which is impossible", operatorSpecDef.Name()))
	}

	// Create a new "secrets" type to hold the secrets
	resourceName := b.resource.Name()
	secretsTypeName := resourceName.WithName(
		b.idFactory.CreateIdentifier(
			resourceName.Name()+"OperatorSecrets",
			astmodel.Exported))
	secretsType := astmodel.NewObjectType()

	// Add the "secrets" property to the operator spec
	secretProp := b.newSecretsProperty(secretsTypeName)
	operatorSpec = operatorSpec.WithProperty(secretProp)

	for _, secret := range azureGeneratedSecrets {
		prop := astmodel.NewPropertyDefinition(
			b.idFactory.CreatePropertyName(secret, astmodel.Exported),
			b.idFactory.CreateIdentifier(secret, astmodel.NotExported),
			astmodel.SecretDestinationType).MakeTypeOptional()
		desc := fmt.Sprintf(
			"indicates where the %s secret should be placed. If omitted, the secret will not be retrieved from Azure.",
			secret)
		prop = prop.WithDescription(desc)
		prop = prop.MakeOptional()
		secretsType = secretsType.WithProperty(prop)
	}

	secretsTypeDef := astmodel.MakeTypeDefinition(secretsTypeName, secretsType)
	b.definitions.Add(secretsTypeDef)

	return operatorSpecDef.WithType(operatorSpec)
}
