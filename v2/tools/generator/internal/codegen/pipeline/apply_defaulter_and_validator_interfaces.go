/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"go/token"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/interfaces"
)

const ApplyDefaulterAndValidatorInterfaceStageID = "applyDefaulterAndValidatorInterfaces"

// We need to:
// - Move location of validators to internal folder?
// - change how we attach validators to objects validators, see https://github.com/kubernetes-sigs/kubebuilder/blob/v4.3.0/docs/book/src/cronjob-tutorial/testdata/project/internal/webhook/v1/cronjob_webhook.go#L51
// - Validators are no longer

// ApplyDefaulterAndValidatorInterfaces add the webhook.CustomDefaulter and webhook.CustomValidator interfaces to each resource that requires them
func ApplyDefaulterAndValidatorInterfaces(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		ApplyDefaulterAndValidatorInterfaceStageID,
		"Add the webhook.CustomDefaulter and webhook.CustomValidator interfaces to a Resource webhook type for each resource that requires them",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			updatedDefs := make(astmodel.TypeDefinitionSet)

			for _, resourceDef := range defs.AllResources() {
				// Create an object to hold the implementation of the validator and defaulter interfaces
				name := astmodel.CreateWebhookTypeName(resourceDef.Name())
				webhookDef := astmodel.MakeTypeDefinition(name, astmodel.NewObjectType())

				defaults, err := getDefaults(configuration, resourceDef, idFactory, state.Definitions())
				if err != nil {
					return nil, eris.Wrap(err, "failed to get defaults")
				}

				webhookDef, err = interfaces.AddDefaulterInterface(resourceDef.Name(), webhookDef, idFactory, defaults)
				if err != nil {
					return nil, err
				}

				validations, err := getValidations(resourceDef, idFactory, state.Definitions())
				if err != nil {
					return nil, eris.Wrapf(err, "error getting validation functions")
				}
				webhookDef, err = interfaces.AddValidatorInterface(resourceDef.Name(), webhookDef, idFactory, validations)
				if err != nil {
					return nil, err
				}

				updatedDefs.Add(webhookDef)
			}

			err := configuration.ObjectModelConfiguration.DefaultAzureName.VerifyConsumed()
			if err != nil {
				return nil, err
			}

			return state.WithOverlaidDefinitions(updatedDefs), nil
		})

	stage.RequiresPrerequisiteStages(ApplyKubernetesResourceInterfaceStageID, AddOperatorSpecStageID)
	return stage
}

func getDefaults(
	configuration *config.Configuration,
	resourceDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	defs astmodel.TypeDefinitionSet,
) ([]*functions.DefaultFunction, error) {
	var result []*functions.DefaultFunction

	resolved, err := defs.ResolveResourceSpecAndStatus(resourceDef)
	if err != nil {
		return nil, eris.Wrapf(err, "unable to resolve resource %s", resourceDef.Name())
	}

	defaultAzureName := true
	if configuredDefaultAzureName, ok := configuration.ObjectModelConfiguration.DefaultAzureName.Lookup(resourceDef.Name()); ok {
		defaultAzureName = configuredDefaultAzureName
	}

	// Determine if the resource has a SetName function
	if resolved.SpecType.HasFunctionWithName(astmodel.SetAzureNameFunc) && defaultAzureName {
		result = append(result, functions.NewDefaultAzureNameFunction(resolved.ResourceDef, idFactory))
	}

	return result, nil
}

func getValidations(
	resourceDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	defs astmodel.TypeDefinitionSet,
) (map[functions.ValidationKind][]*functions.ValidateFunction, error) {
	resource, ok := resourceDef.Type().(*astmodel.ResourceType)
	if !ok {
		return nil, eris.Errorf("resource %s did not have type of kind *astmodel.ResourceType, instead %T", resourceDef.Name(), resourceDef.Type())
	}

	validations := map[functions.ValidationKind][]*functions.ValidateFunction{
		functions.ValidationKindCreate: {
			functions.NewValidateResourceReferencesFunction(resourceDef, idFactory),
		},
		functions.ValidationKindUpdate: {
			functions.NewValidateResourceReferencesFunction(resourceDef, idFactory),
			functions.NewValidateWriteOncePropertiesFunction(resourceDef, idFactory),
		},
	}

	if !resource.Owner().IsEmpty() {
		validations[functions.ValidationKindCreate] = append(
			validations[functions.ValidationKindCreate],
			functions.NewValidateOwnerReferenceFunction(resourceDef, idFactory))
		validations[functions.ValidationKindUpdate] = append(
			validations[functions.ValidationKindUpdate],
			functions.NewValidateOwnerReferenceFunction(resourceDef, idFactory))
	}

	// The expectation is that every resource has an Spec.OperatorSpec.SecretExpressions and
	// Spec.OperatorSpec.ConfigMapExpressions field, so we always include their validations.
	// If this assumption has been violated, generating the validation function will raise an error.
	validations[functions.ValidationKindCreate] = append(
		validations[functions.ValidationKindCreate],
		NewValidateSecretDestinationsFunction(resourceDef, idFactory))
	validations[functions.ValidationKindUpdate] = append(
		validations[functions.ValidationKindUpdate],
		NewValidateSecretDestinationsFunction(resourceDef, idFactory))
	validations[functions.ValidationKindCreate] = append(
		validations[functions.ValidationKindCreate],
		NewValidateConfigMapDestinationsFunction(resourceDef, idFactory))
	validations[functions.ValidationKindUpdate] = append(
		validations[functions.ValidationKindUpdate],
		NewValidateConfigMapDestinationsFunction(resourceDef, idFactory))

	hasConfigMapReferencePairs, err := hasOptionalConfigMapReferencePairs(resourceDef, defs)
	if err != nil {
		return nil, err
	}
	if hasConfigMapReferencePairs {
		validations[functions.ValidationKindCreate] = append(
			validations[functions.ValidationKindCreate],
			functions.NewValidateOptionalConfigMapReferenceFunction(resourceDef, idFactory))
		validations[functions.ValidationKindUpdate] = append(
			validations[functions.ValidationKindUpdate],
			functions.NewValidateOptionalConfigMapReferenceFunction(resourceDef, idFactory))
	}

	return validations, nil
}

// Note: This isn't defined in the functions package because it has a dependency on getResourceSecretsType which
// doesn't make a lot of sense to put into astmodel. Functions can't import code from pipelines though, so we just
// define this function here.

// NewValidateSecretDestinationsFunction creates a function for validating secret destinations:
//
//	func (account *<obj>) validateSecretDestinations(ctx context.Context, obj *<obj>) (admission.Warnings, error) {
//		if obj.Spec.OperatorSpec == nil {
//			return nil, nil
//		}
//		return secrets.ValidateDestinations(obj, <operatorSpecSecrets>, <operatorSpecCELSecrets>)
//	}
func NewValidateSecretDestinationsFunction(resourceDef astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) *functions.ValidateFunction {
	return functions.NewValidateFunction(
		"validateSecretDestinations",
		resourceDef.Name(),
		idFactory,
		validateSecretDestinations(resourceDef),
		astmodel.GenRuntimeSecretsReference)
}

func validateSecretDestinations(resourceDef astmodel.TypeDefinition) functions.DataFunctionHandler[astmodel.InternalTypeName] {
	resourceType, ok := astmodel.AsResourceType(resourceDef.Type())
	if !ok {
		// This is not expected
		panic("resource type was not a ResourceType")
	}

	return func(k *functions.ValidateFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) (*dst.FuncDecl, error) {
		objIdent := "obj"
		contextIdent := "ctx"

		receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
		receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating receiver expression")
		}

		body, err := validateOperatorSpecSliceBody(
			codeGenerationContext,
			resourceType,
			objIdent,
			astmodel.OperatorSpecSecretsProperty,
			astmodel.OperatorSpecSecretExpressionsProperty,
			astmodel.NewOptionalType(astmodel.SecretDestinationType),
			astmodel.GenRuntimeSecretsReference,
			"ValidateDestinations")
		if err != nil {
			return nil, eris.Wrapf(err, "creating body of method %s", methodName)
		}

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  astbuilder.PointerTo(receiverExpr),
			Body:          body,
		}

		contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating context type expression")
		}
		fn.AddParameter(contextIdent, contextTypeExpr)

		resourceTypeExpr, err := k.Data().AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating resource type expression")
		}
		fn.AddParameter(objIdent, astbuilder.PointerTo(resourceTypeExpr))

		fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
		fn.AddReturn(dst.NewIdent("error"))
		fn.AddComments("validates there are no colliding genruntime.SecretDestination's")

		return fn.DefineFunc(), nil
	}
}

// NewValidateConfigMapDestinationsFunction creates a function for validating configmap destinations
//
//	func (endpoint *<obj>) validateConfigMapDestinations(ctx context.Context, obj *<obj>) (admission.Warnings, error) {
//		if obj.Spec.OperatorSpec == nil {
//			return nil, nil
//		}
//		return configmaps.ValidateDestinations(obj, <operatorSpecSecrets>, <operatorSpecCELSecrets>)
//	}
func NewValidateConfigMapDestinationsFunction(resourceDef astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) *functions.ValidateFunction {
	return functions.NewValidateFunction(
		"validateConfigMapDestinations",
		resourceDef.Name(),
		idFactory,
		validateConfigMapDestinations(resourceDef),
		astmodel.GenRuntimeConfigMapsReference)
}

func validateConfigMapDestinations(resourceDef astmodel.TypeDefinition) functions.DataFunctionHandler[astmodel.InternalTypeName] {
	resourceType, ok := astmodel.AsResourceType(resourceDef.Type())
	if !ok {
		// This is not expected
		panic("resource type was not a ResourceType")
	}

	return func(k *functions.ValidateFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) (*dst.FuncDecl, error) {
		objIdent := "obj"
		contextIdent := "ctx"

		receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
		receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating receiver expression")
		}

		body, err := validateOperatorSpecSliceBody(
			codeGenerationContext,
			resourceType,
			objIdent,
			astmodel.OperatorSpecConfigMapsProperty,
			astmodel.OperatorSpecConfigMapExpressionsProperty,
			astmodel.NewOptionalType(astmodel.ConfigMapDestinationType),
			astmodel.GenRuntimeConfigMapsReference,
			"ValidateDestinations")
		if err != nil {
			return nil, eris.Wrapf(err, "creating body of method %s", methodName)
		}

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  astbuilder.PointerTo(receiverExpr),
			Body:          body,
		}

		contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating context type expression")
		}
		fn.AddParameter(contextIdent, contextTypeExpr)

		resourceTypeExpr, err := k.Data().AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating resource type expression")
		}
		fn.AddParameter(objIdent, astbuilder.PointerTo(resourceTypeExpr))

		runtimeAdmission := codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission)
		fn.AddReturn(astbuilder.QualifiedTypeName(runtimeAdmission, "Warnings"))
		fn.AddReturn(dst.NewIdent("error"))
		fn.AddComments("validates there are no colliding genruntime.ConfigMapDestinations")

		return fn.DefineFunc(), nil
	}
}

// validateOperatorSpecSliceBody helps generate the body of the validateResourceReferences function:
//
//	func (account *DatabaseAccount) validateConfigMapDestinations(ctx context.Context, obj *T) error {
//	    if <obj>.Spec.OperatorSpec == nil {
//		       return nil
//		}
//
//	    var toValidate []<validateType>
//	    if <obj>.Spec.OperatorSpec.<property> != nil {
//	        toValidate = []<validateType>{
//	            account.Spec.OperatorSpec.Secrets.PrimaryReadonlyMasterKey,
//	            account.Spec.OperatorSpec.Secrets.SecondaryReadonlyMasterKey,
//	            ...
//	    }
//	    return <validatePkg>.<validateFunctionName>(account, toValidate, <receiver>.Spec.OperatorSpec.<expressionsProperty>)
//	}
func validateOperatorSpecSliceBody(
	codeGenerationContext *astmodel.CodeGenerationContext,
	resource *astmodel.ResourceType,
	objIdent string,
	property string,
	expressionsProperty string,
	validateType astmodel.Type,
	validatePackage astmodel.ExternalPackageReference,
	validateFunctionName string,
) ([]dst.Stmt, error) {
	pkg := codeGenerationContext.MustGetImportedPackageName(validatePackage)

	// operatorSpecPropertyObj and expressionsOperatorSpecPropertyObj may be nil. If BOTH are nil this method shouldn't
	// have ever been called, but one can be nil.
	operatorSpecPropertyObj, err := getOperatorSpecSubType(codeGenerationContext, resource, property)
	if err != nil {
		return nil, eris.Wrapf(err, "getting operator spec sub type for %s", property)
	}
	expressionsOperatorSpecPropertySlice, err := getOperatorSpecExpressionType(codeGenerationContext, resource, expressionsProperty)
	if err != nil {
		return nil, eris.Wrapf(err, "getting operator spec sub type for %s", expressionsProperty)
	}

	if operatorSpecPropertyObj == nil && expressionsOperatorSpecPropertySlice == nil {
		return nil, eris.Errorf(
			"can't generate method for validating OperatorSpec %s and %s if both fields don't exist",
			property,
			expressionsProperty)
	}

	var body []dst.Stmt

	specSelector := astbuilder.Selector(dst.NewIdent(objIdent), "Spec")
	// if <receiver>.Spec.OperatorSpec == nil {
	//     return nil, nil
	// }
	operatorSpecSelector := astbuilder.Selector(specSelector, astmodel.OperatorSpecProperty)
	body = append(body, astbuilder.ReturnIfNil(operatorSpecSelector, astbuilder.Nil(), astbuilder.Nil()))

	propertySelector := astbuilder.Selector(operatorSpecSelector, property)
	expressionsPropertySelector := astbuilder.Selector(operatorSpecSelector, expressionsProperty)

	toValidateVar := "toValidate"
	if operatorSpecPropertyObj != nil {
		// var toValidate []<validateType>
		validateTypeExpr, err := validateType.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating type expression for %s", validateType)
		}
		validateVarSliceExpr, err := astmodel.NewArrayType(validateType).AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating type expression for %s", astmodel.NewArrayType(validateType))
		}
		body = append(body, astbuilder.LocalVariableDeclaration(toValidateVar, validateVarSliceExpr, ""))

		// if <receiver>.Spec.OperatorSpec.<property> != nil {
		//     toValidate = []<validateType>{
		//         account.Spec.OperatorSpec.Secrets.PrimaryReadonlyMasterKey,
		//         account.Spec.OperatorSpec.Secrets.SecondaryReadonlyMasterKey,
		//         ...
		//     }
		// }

		sliceBuilder := astbuilder.NewSliceLiteralBuilder(validateTypeExpr, true)
		for _, prop := range operatorSpecPropertyObj.Properties().AsSlice() {
			propSelector := astbuilder.Selector(propertySelector, prop.PropertyName().String())
			sliceBuilder.AddElement(propSelector)
		}
		body = append(
			body,
			astbuilder.IfNotNil(
				propertySelector,
				astbuilder.AssignmentStatement(dst.NewIdent(toValidateVar), token.ASSIGN, sliceBuilder.Build())))
	}

	selfParameter := dst.NewIdent(objIdent)

	toValidateParameter := astbuilder.Nil()
	if operatorSpecPropertyObj != nil {
		toValidateParameter = dst.NewIdent(toValidateVar)
	}

	var expressionsToValidateParameter dst.Expr
	expressionsToValidateParameter = astbuilder.Nil()
	if expressionsOperatorSpecPropertySlice != nil {
		expressionsToValidateParameter = expressionsPropertySelector
	}

	// return <validatePkg>.<validateFunctionName>(self, toValidate, expressionsToValidate)
	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				pkg,
				validateFunctionName,
				selfParameter,
				toValidateParameter,
				expressionsToValidateParameter)))

	return body, nil
}

func getOperatorSpecType(defs astmodel.ReadonlyTypeDefinitions, resource *astmodel.ResourceType) (*astmodel.ObjectType, error) {
	spec, err := astmodel.ResolveResourceSpecDefinition(defs, resource)
	if err != nil {
		return nil, err
	}

	typedSpec, ok := astmodel.AsObjectType(spec.Type())
	if !ok {
		return nil, eris.Errorf("resource spec was not of expected type *astmodel.ObjectType, instead %T", spec.Type())
	}

	operatorSpecProp, ok := typedSpec.Property(astmodel.OperatorSpecProperty)
	if !ok {
		// No OperatorSpec property - this means no secrets
		return nil, nil
	}
	operatorSpecTypeName, ok := astmodel.AsInternalTypeName(operatorSpecProp.PropertyType())
	if !ok {
		return nil, eris.Errorf(
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
		return nil, eris.Errorf(
			"expected %s to be an astmodel.ObjectType but it was %T",
			operatorSpecTypeName,
			operatorSpecDef.Type())
	}

	return operatorSpecType, nil
}

func getOperatorSpecExpressionType(defs astmodel.ReadonlyTypeDefinitions, resource *astmodel.ResourceType, name string) (astmodel.Type, error) {
	operatorSpecType, err := getOperatorSpecType(defs, resource)
	if err != nil {
		return nil, err
	}
	if operatorSpecType == nil {
		// Not found, just return
		return nil, nil
	}

	prop, ok := operatorSpecType.Property(astmodel.PropertyName(name))
	if !ok {
		// No property
		return nil, nil
	}

	propType := astmodel.Unwrap(prop.PropertyType())
	return propType, nil
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

	prop, ok := operatorSpecType.Property(astmodel.PropertyName(name))
	if !ok {
		// No property
		return nil, nil
	}

	typeName, ok := astmodel.AsInternalTypeName(prop.PropertyType())
	if !ok {
		return nil, eris.Errorf(
			"expected %s to be an astmodel.InternalTypeName, but it was %T",
			name,
			prop.PropertyType())
	}
	def, err := defs.GetDefinition(typeName)
	if err != nil {
		return nil, err
	}
	defType, ok := astmodel.AsObjectType(def.Type())
	if !ok {
		panic(fmt.Sprintf("expected %s to be an astmodel.ObjectType but it was %T", typeName, def.Type()))
	}

	return defType, nil
}

// hasOptionalConfigMapReferencePairs returns true if the type has optional genruntime.ConfigMapReference pairs
func hasOptionalConfigMapReferencePairs(resourceDef astmodel.TypeDefinition, defs astmodel.TypeDefinitionSet) (bool, error) {
	result := false
	visitor := astmodel.TypeVisitorBuilder[any]{
		VisitObjectType: astmodel.MakeIdentityVisitOfObjectType(
			func(ot *astmodel.ObjectType, prop *astmodel.PropertyDefinition, ctx any) (any, error) {
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
