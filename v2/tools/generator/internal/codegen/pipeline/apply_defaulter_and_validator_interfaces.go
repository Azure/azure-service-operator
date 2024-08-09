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
) ([]*functions.ResourceFunction, error) {
	var result []*functions.ResourceFunction

	resolved, err := defs.ResolveResourceSpecAndStatus(resourceDef)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to resolve resource %s", resourceDef.Name())
	}

	defaultAzureName := true
	if configuredDefaultAzureName, ok := configuration.ObjectModelConfiguration.DefaultAzureName.Lookup(resourceDef.Name()); ok {
		defaultAzureName = configuredDefaultAzureName
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
	defs astmodel.TypeDefinitionSet,
) (map[functions.ValidationKind][]*functions.ResourceFunction, error) {
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

	if !resource.Owner().IsEmpty() {
		validations[functions.ValidationKindCreate] = append(
			validations[functions.ValidationKindCreate],
			functions.NewValidateOwnerReferenceFunction(resource, idFactory))
		validations[functions.ValidationKindUpdate] = append(
			validations[functions.ValidationKindUpdate],
			functions.NewValidateOwnerReferenceFunction(resource, idFactory))
	}

	// The expectation is that every resource has an Spec.OperatorSpec.SecretExpressions and
	// Spec.OperatorSpec.ConfigMapExpressions field, so we always include their validations.
	// If this assumption has been violated, generating the validation function will raise an error.
	validations[functions.ValidationKindCreate] = append(
		validations[functions.ValidationKindCreate],
		NewValidateSecretDestinationsFunction(resource, idFactory))
	validations[functions.ValidationKindUpdate] = append(
		validations[functions.ValidationKindUpdate],
		NewValidateSecretDestinationsFunction(resource, idFactory))
	validations[functions.ValidationKindCreate] = append(
		validations[functions.ValidationKindCreate],
		NewValidateConfigMapDestinationsFunction(resource, idFactory))
	validations[functions.ValidationKindUpdate] = append(
		validations[functions.ValidationKindUpdate],
		NewValidateConfigMapDestinationsFunction(resource, idFactory))

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
		astmodel.NewPackageReferenceSet(astmodel.GenRuntimeSecretsReference))
}

func validateSecretDestinations(
	k *functions.ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, errors.Wrapf(err, "creating receiver expression")
	}

	body, err := validateOperatorSpecSliceBody(
		codeGenerationContext,
		k.Resource(),
		receiverIdent,
		astmodel.OperatorSpecSecretsProperty,
		astmodel.OperatorSpecSecretExpressionsProperty,
		astmodel.NewOptionalType(astmodel.SecretDestinationType),
		astmodel.GenRuntimeSecretsReference,
		"ValidateDestinations")
	if err != nil {
		return nil, errors.Wrapf(err, "creating body of method %s", methodName)
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body:          body,
	}

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates there are no colliding genruntime.SecretDestination's")

	return fn.DefineFunc(), nil
}

func NewValidateConfigMapDestinationsFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *functions.ResourceFunction {
	return functions.NewResourceFunction(
		"validateConfigMapDestinations",
		resource,
		idFactory,
		validateConfigMapDestinations,
		astmodel.NewPackageReferenceSet(astmodel.GenRuntimeConfigMapsReference))
}

func validateConfigMapDestinations(
	k *functions.ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, errors.Wrapf(err, "creating receiver expression")
	}

	body, err := validateOperatorSpecSliceBody(
		codeGenerationContext,
		k.Resource(),
		receiverIdent,
		astmodel.OperatorSpecConfigMapsProperty,
		astmodel.OperatorSpecConfigMapExpressionsProperty,
		astmodel.NewOptionalType(astmodel.ConfigMapDestinationType),
		astmodel.GenRuntimeConfigMapsReference,
		"ValidateDestinations")
	if err != nil {
		return nil, errors.Wrapf(err, "creating body of method %s", methodName)
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body:          body,
	}

	runtimeAdmission := codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission)
	fn.AddReturn(astbuilder.QualifiedTypeName(runtimeAdmission, "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates there are no colliding genruntime.ConfigMapDestinations")

	return fn.DefineFunc(), nil
}

// validateOperatorSpecSliceBody helps generate the body of the validateResourceReferences function:
//
//	func (account *DatabaseAccount) validateConfigMapDestinations() error {
//	    if <receiver>.Spec.OperatorSpec == nil {
//		       return nil
//		}
//
//	    var toValidate []<validateType>
//	    if <receiver>.Spec.OperatorSpec.<property> != nil {
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
	receiverIdent string,
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
		return nil, errors.Wrapf(err, "getting operator spec sub type for %s", property)
	}
	expressionsOperatorSpecPropertySlice, err := getOperatorSpecExpressionType(codeGenerationContext, resource, expressionsProperty)
	if err != nil {
		return nil, errors.Wrapf(err, "getting operator spec sub type for %s", expressionsProperty)
	}

	if operatorSpecPropertyObj == nil && expressionsOperatorSpecPropertySlice == nil {
		return nil, errors.Errorf(
			"can't generate method for validating OperatorSpec %s and %s if both fields don't exist",
			property,
			expressionsProperty)
	}

	var body []dst.Stmt

	specSelector := astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")
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
			return nil, errors.Wrapf(err, "creating type expression for %s", validateType)
		}
		validateVarSliceExpr, err := astmodel.NewArrayType(validateType).AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, errors.Wrapf(err, "creating type expression for %s", astmodel.NewArrayType(validateType))
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

	selfParameter := dst.NewIdent(receiverIdent)

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
		return nil, errors.Errorf("resource spec was not of expected type *astmodel.ObjectType, instead %T", spec.Type())
	}

	operatorSpecProp, ok := typedSpec.Property(astmodel.OperatorSpecProperty)
	if !ok {
		// No OperatorSpec property - this means no secrets
		return nil, nil
	}
	operatorSpecTypeName, ok := astmodel.AsInternalTypeName(operatorSpecProp.PropertyType())
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
		return nil, errors.Errorf(
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
