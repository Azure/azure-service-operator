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
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/interfaces"
)

const ApplyDefaulterAndValidatorInterfaceStageID = "applyDefaulterAndValidatorInterfaces"

// ApplyDefaulterAndValidatorInterfaces add the admission.Defaulter and admission.Validator interfaces to each resource that requires them
func ApplyDefaulterAndValidatorInterfaces(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		ApplyDefaulterAndValidatorInterfaceStageID,
		"Add the admission.Defaulter and admission.Validator interfaces to each resource that requires them",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			updatedDefs := make(astmodel.TypeDefinitionSet)

			for _, resourceDef := range astmodel.FindResourceDefinitions(defs) {
				defaults, err := getDefaults(configuration, resourceDef, idFactory, state.Definitions())
				if err != nil {
					return nil, errors.Wrap(err, "failed to get defaults")
				}

				resource, err := interfaces.AddDefaulterInterface(resourceDef, idFactory, defaults)
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

			err := configuration.ObjectModelConfiguration.DefaultAzureName.VerifyConsumed()
			if err != nil {
				return nil, err
			}

			return state.WithDefinitions(defs.OverlayWith(updatedDefs)), nil
		})

	stage.RequiresPrerequisiteStages(ApplyKubernetesResourceInterfaceStageID, AddOperatorSpecStageID)
	return stage
}

func getDefaults(
	configuration *config.Configuration,
	resourceDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	defs astmodel.TypeDefinitionSet,
) ([]*functions.ResourceFunction, error) {
	var result []*functions.ResourceFunction

	resolved, err := defs.ResolveResourceSpecAndStatus(resourceDef)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to resolve resource %s", resourceDef.Name())
	}

	defaultAzureName, err := configuration.ObjectModelConfiguration.DefaultAzureName.Lookup(resourceDef.Name())
	if err != nil {
		if config.IsNotConfiguredError(err) {
			// Default to true if we have no explicit configuration
			defaultAzureName = true
		} else {
			return nil, err
		}
	}

	// Determine if the resource has a SetName function
	if resolved.SpecType.HasFunctionWithName(astmodel.SetAzureNameFunc) && defaultAzureName {
		result = append(result, functions.NewDefaultAzureNameFunction(resolved.ResourceType, idFactory))
	}

	return result, nil
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
			functions.NewValidateWriteOncePropertiesFunction(resource, idFactory),
		},
	}

	secrets, err := getOperatorSpecSubType(defs, resource, astmodel.OperatorSpecSecretsProperty)
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

	configMaps, err := getOperatorSpecSubType(defs, resource, astmodel.OperatorSpecConfigMapsProperty)
	if err != nil {
		return nil, err
	}
	if configMaps != nil {
		validations[functions.ValidationKindCreate] = append(
			validations[functions.ValidationKindCreate],
			NewValidateConfigMapDestinationsFunction(resource, idFactory))
		validations[functions.ValidationKindUpdate] = append(
			validations[functions.ValidationKindUpdate],
			NewValidateConfigMapDestinationsFunction(resource, idFactory))
	}

	hasConfigMapReferencePairs, err := hasOptionalConfigMapReferencePairs(resourceDef, defs)
	if err != nil {
		return nil, err
	}
	if hasConfigMapReferencePairs {
		validations[functions.ValidationKindCreate] = append(
			validations[functions.ValidationKindCreate],
			functions.NewValidateOptionalConfigMapReferenceFunction(resource, idFactory))
		validations[functions.ValidationKindUpdate] = append(
			validations[functions.ValidationKindUpdate],
			functions.NewValidateOptionalConfigMapReferenceFunction(resource, idFactory))
	}

	return validations, nil
}

// Note: This isn't defined in the functions package because it has a dependency on getResourceSecretsType which
// doesn't make a lot of sense to put into astmodel. Functions can't import code from pipelines though, so we just
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
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Returns: []*dst.Field{
			{
				Type: dst.NewIdent("error"),
			},
		},
		Body: validateOperatorSpecSliceBody(
			codeGenerationContext,
			k.Resource(),
			receiverIdent,
			astmodel.OperatorSpecSecretsProperty,
			astmodel.NewOptionalType(astmodel.SecretDestinationType),
			"ValidateSecretDestinations"),
	}

	fn.AddComments("validates there are no colliding genruntime.SecretDestination's")
	return fn.DefineFunc()
}

func NewValidateConfigMapDestinationsFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *functions.ResourceFunction {
	return functions.NewResourceFunction(
		"validateConfigMapDestinations",
		resource,
		idFactory,
		validateConfigMapDestinations,
		astmodel.NewPackageReferenceSet(astmodel.GenRuntimeReference))
}

func validateConfigMapDestinations(k *functions.ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Returns: []*dst.Field{
			{
				Type: dst.NewIdent("error"),
			},
		},
		Body: validateOperatorSpecSliceBody(
			codeGenerationContext,
			k.Resource(),
			receiverIdent,
			astmodel.OperatorSpecConfigMapsProperty,
			astmodel.NewOptionalType(astmodel.ConfigMapDestinationType),
			"ValidateConfigMapDestinations"),
	}

	fn.AddComments("validates there are no colliding genruntime.ConfigMapDestinations's")
	return fn.DefineFunc()
}

// validateOperatorSpecSliceBody helps generate the body of the validateResourceReferences function:
//
//	func (account *DatabaseAccount) validateConfigMapDestinations() error {
//	    if <receiver>.Spec.OperatorSpec == nil {
//	        return nil
//	    }
//	    if <receiver>.Spec.OperatorSpec.<operatorSpecProperty> == nil {
//	        return nil
//	    }
//	    toValidate := []*<validateType>{
//	        account.Spec.OperatorSpec.ConfigMaps.ClientId,
//	        account.Spec.OperatorSpec.ConfigMaps.PrincipalId,
//	        ...
//	    }
//	    return genruntime.<validateFunctionName>(toValidate)
//	}
func validateOperatorSpecSliceBody(
	codeGenerationContext *astmodel.CodeGenerationContext,
	resource *astmodel.ResourceType,
	receiverIdent string,
	operatorSpecProperty string,
	validateType astmodel.Type,
	validateFunctionName string,
) []dst.Stmt {
	genRuntime := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)

	operatorSpecPropertyObj, err := getOperatorSpecSubType(codeGenerationContext, resource, operatorSpecProperty)
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

	// if <receiver>.Spec.OperatorSpec.<operatorSpecProperty> == nil {
	//     return nil
	// }
	specPropertySelector := astbuilder.Selector(operatorSpecSelector, operatorSpecProperty)
	body = append(body, astbuilder.ReturnIfNil(specPropertySelector, astbuilder.Nil()))

	// secrets := []<validateType>{
	//     account.Spec.OperatorSpec.Secrets.PrimaryReadonlyMasterKey,
	//     account.Spec.OperatorSpec.Secrets.SecondaryReadonlyMasterKey,
	//     ...
	// }
	sliceBuilder := astbuilder.NewSliceLiteralBuilder(
		validateType.AsType(codeGenerationContext),
		true)
	for _, prop := range operatorSpecPropertyObj.Properties().AsSlice() {
		propSelector := astbuilder.Selector(specPropertySelector, prop.PropertyName().String())
		sliceBuilder.AddElement(propSelector)
	}
	toValidateVar := "toValidate"
	body = append(body, astbuilder.ShortDeclaration(toValidateVar, sliceBuilder.Build()))

	// return genruntime.<validateFunctionName>(secrets)
	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				genRuntime,
				validateFunctionName,
				dst.NewIdent(toValidateVar))))

	return body
}

func getOperatorSpecType(defs astmodel.ReadonlyTypeDefinitions, resource *astmodel.ResourceType) (*astmodel.ObjectType, error) {
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

	return operatorSpecType, nil
}

func getOperatorSpecSubType(defs astmodel.ReadonlyTypeDefinitions, resource *astmodel.ResourceType, name string) (*astmodel.ObjectType, error) {
	operatorSpecType, err := getOperatorSpecType(defs, resource)
	if err != nil {
		return nil, err
	}
	if operatorSpecType == nil {
		// Not found, just return
		return nil, nil
	}

	secretsProp, ok := operatorSpecType.Property(astmodel.PropertyName(name))
	if !ok {
		// No secrets property
		return nil, nil
	}

	secretsTypeName, ok := astmodel.AsTypeName(secretsProp.PropertyType())
	if !ok {
		return nil, errors.Errorf(
			"expected %s to be an astmodel.TypeName, but it was %T",
			name,
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

// hasOptionalConfigMapReferencePairs returns true if the type has optional genruntime.ConfigMapReference pairs
func hasOptionalConfigMapReferencePairs(resourceDef astmodel.TypeDefinition, defs astmodel.TypeDefinitionSet) (bool, error) {
	result := false
	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: astmodel.MakeIdentityVisitOfObjectType(func(ot *astmodel.ObjectType, prop *astmodel.PropertyDefinition, ctx interface{}) (interface{}, error) {
			if prop.HasTag(astmodel.OptionalConfigMapPairTag) {
				result = true
			}

			return ctx, nil
		}),
	}.Build()

	walker := astmodel.NewTypeWalker(defs, visitor)
	_, err := walker.Walk(resourceDef)
	if err != nil {
		return false, err
	}

	return result, nil
}
