/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/dave/dst"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/interfaces"
)

const ApplyDefaulterAndValidatorInterfaceStageID = "applyDefaulterAndValidatorInterfaces"

// ApplyDefaulterAndValidatorInterfaces add the admission.Defaulter and admission.Validator interfaces to each resource that requires them
func ApplyDefaulterAndValidatorInterfaces(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		ApplyDefaulterAndValidatorInterfaceStageID,
		"Add the admission.Defaulter and admission.Validator interfaces to each resource that requires them",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			updatedDefs := make(astmodel.TypeDefinitionSet)

			for _, typeDef := range astmodel.FindResourceDefinitions(defs) {
				resource, err := interfaces.AddDefaulterInterface(typeDef, idFactory, defs)
				if err != nil {
					return nil, err
				}

				validations, err := getValidations(resource, idFactory, state.Definitions())
				if err != nil {
					return nil, errors.Wrapf(err, "error getting validation functions")
				}
				resource, err = interfaces.AddValidatorInterface(resource, idFactory, defs, validations)
				if err != nil {
					return nil, err
				}

				updatedDefs.Add(resource)
			}

			return state.WithDefinitions(defs.OverlayWith(updatedDefs)), nil
		})

	stage.RequiresPrerequisiteStages(ApplyKubernetesResourceInterfaceStageID, AddOperatorSpecStageID)
	return stage
}

func getValidations(
	resourceDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	defs astmodel.TypeDefinitionSet) (map[functions.ValidationKind][]*functions.ResourceFunction, error) {

	resource, ok := resourceDef.Type().(*astmodel.ResourceType)
	if !ok {
		return nil, errors.Errorf("resource %s did not have type of kind *astmodel.ResourceType, instead %T", resourceDef.Name(), resourceDef.Type())
	}

	validations := map[functions.ValidationKind][]*functions.ResourceFunction{
		functions.ValidationKindCreate: {
			functions.NewValidateResourceReferencesFunction(resource, idFactory),
		},
		functions.ValidationKindUpdate: {
			functions.NewValidateResourceReferencesFunction(resource, idFactory),
			functions.NewValidateImmutablePropertiesFunction(resource, idFactory),
		},
	}

	secrets, err := getResourceSecretsType(defs, resource)
	if err != nil {
		return nil, err
	}

	if secrets != nil {
		validations[functions.ValidationKindCreate] = append(
			validations[functions.ValidationKindCreate],
			NewValidateSecretDestinationsFunction(resource, idFactory))
		validations[functions.ValidationKindUpdate] = append(
			validations[functions.ValidationKindUpdate],
			NewValidateSecretDestinationsFunction(resource, idFactory))
	}

	return validations, nil
}

// Note: This isn't defined in the functions package because it has a dependency on getResourceSecretsType which
// doesn't make a lot of sense to put into astmodel. Functions can't import code from pipelines though so we just
// define this function here.

func NewValidateSecretDestinationsFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *functions.ResourceFunction {
	return functions.NewResourceFunction(
		"validateSecretDestinations",
		resource,
		idFactory,
		validateSecretDestinations,
		astmodel.NewPackageReferenceSet(astmodel.GenRuntimeReference))
}

func validateSecretDestinations(k *functions.ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Returns: []*dst.Field{
			{
				Type: dst.NewIdent("error"),
			},
		},
		Body: validateSecretDestinationsBody(codeGenerationContext, k.Resource(), receiverIdent),
	}

	fn.AddComments("validates there are no colliding genruntime.SecretDestination's")
	return fn.DefineFunc()
}

// validateSecretDestinationsBody helps generate the body of the validateResourceReferences function:
// func (account *DatabaseAccount) validateSecretDestinations() error {
//     if <receiver>.Spec.OperatorSpec == nil {
//         return nil
//     }
//     if <receiver>.Spec.OperatorSpec.Secrets == nil {
//         return nil
//     }
//     secrets := []*genruntime.SecretDestination{
//         account.Spec.OperatorSpec.Secrets.PrimaryReadonlyMasterKey,
//         account.Spec.OperatorSpec.Secrets.SecondaryReadonlyMasterKey,
//         ...
//     }
//     return genruntime.ValidateSecretDestinations(secrets)
// }
func validateSecretDestinationsBody(codeGenerationContext *astmodel.CodeGenerationContext, resource *astmodel.ResourceType, receiverIdent string) []dst.Stmt {
	genRuntime := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)

	resourceSecrets, err := getResourceSecretsType(codeGenerationContext, resource)
	if err != nil {
		panic(err)
	}

	var body []dst.Stmt

	specSelector := astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")
	// if <receiver>.Spec.OperatorSpec == nil {
	//     return nil
	// }
	operatorSpecSelector := astbuilder.Selector(specSelector, astmodel.OperatorSpecProperty)
	body = append(body, astbuilder.ReturnIfNil(operatorSpecSelector, astbuilder.Nil()))

	// if <receiver>.Spec.OperatorSpec.Secrets == nil {
	//     return nil
	// }
	secretsSelector := astbuilder.Selector(operatorSpecSelector, astmodel.OperatorSpecSecretsProperty)
	body = append(body, astbuilder.ReturnIfNil(secretsSelector, astbuilder.Nil()))

	// secrets := []*genruntime.SecretDestination{
	//     account.Spec.OperatorSpec.Secrets.PrimaryReadonlyMasterKey,
	//     account.Spec.OperatorSpec.Secrets.SecondaryReadonlyMasterKey,
	//     ...
	// }
	sliceBuilder := astbuilder.NewSliceLiteralBuilder(
		astmodel.NewOptionalType(astmodel.SecretDestinationType).AsType(codeGenerationContext),
		true)
	for _, prop := range resourceSecrets.Properties().AsSlice() {
		propSelector := astbuilder.Selector(secretsSelector, prop.PropertyName().String())
		sliceBuilder.AddElement(propSelector)
	}
	secretsVar := "secrets"
	body = append(body, astbuilder.ShortDeclaration(secretsVar, sliceBuilder.Build()))

	// return genruntime.ValidateSecretDestinations(secrets)
	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				genRuntime,
				"ValidateSecretDestinations",
				dst.NewIdent(secretsVar))))

	return body
}

func getResourceSecretsType(defs astmodel.ReadonlyTypeDefinitions, resource *astmodel.ResourceType) (*astmodel.ObjectType, error) {
	spec, err := astmodel.ResolveResourceSpecDefinition(defs, resource)
	if err != nil {
		return nil, err
	}

	typedSpec, ok := astmodel.AsObjectType(spec.Type())
	if !ok {
		return nil, errors.Errorf("resource spec was not of expected type *astmodel.ObjectType, instead %T", spec.Type())
	}

	operatorSpecProp, ok := typedSpec.Property(astmodel.OperatorSpecProperty)
	if !ok {
		// No OperatorSpec property - this means no secrets
		return nil, nil
	}
	operatorSpecTypeName, ok := astmodel.AsTypeName(operatorSpecProp.PropertyType())
	if !ok {
		return nil, errors.Errorf(
			"expected %s to be an astmodel.TypeName, but it was %T",
			astmodel.OperatorSpecProperty,
			operatorSpecProp.PropertyType())
	}

	operatorSpecDef, err := defs.GetDefinition(operatorSpecTypeName)
	if err != nil {
		return nil, err
	}
	operatorSpecType, ok := astmodel.AsObjectType(operatorSpecDef.Type())
	if !ok {
		return nil, errors.Errorf(
			"expected %s to be an astmodel.ObjectType but it was %T",
			operatorSpecTypeName,
			operatorSpecDef.Type())
	}

	secretsProp, ok := operatorSpecType.Property(astmodel.OperatorSpecSecretsProperty)
	if !ok {
		// No secrets property
		return nil, nil
	}

	secretsTypeName, ok := astmodel.AsTypeName(secretsProp.PropertyType())
	if !ok {
		return nil, errors.Errorf(
			"expected %s to be an astmodel.TypeName, but it was %T",
			astmodel.OperatorSpecSecretsProperty,
			secretsProp.PropertyType())
	}
	secretsDef, err := defs.GetDefinition(secretsTypeName)
	if err != nil {
		return nil, err
	}
	secretsType, ok := astmodel.AsObjectType(secretsDef.Type())
	if !ok {
		panic(fmt.Sprintf("expected %s to be an astmodel.ObjectType but it was %T", secretsTypeName, secretsDef.Type()))
	}

	return secretsType, nil
}
