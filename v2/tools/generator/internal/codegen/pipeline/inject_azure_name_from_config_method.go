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
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// InjectAzureNameFromConfigMethodStageID is the unique identifier for this pipeline stage
const InjectAzureNameFromConfigMethodStageID = "injectAzureNameFromConfigMethod"

// InjectAzureNameFromConfigMethod injects the GetAzureNameFromConfig() method onto storage spec types
// that have the AzureNameFromConfig property. This allows the reconciler to access the property via the
// AzureNameFromConfigProvider interface without needing to convert back to the API version.
func InjectAzureNameFromConfigMethod(idFactory astmodel.IdentifierFactory) *Stage {
	return NewStage(
		InjectAzureNameFromConfigMethodStageID,
		"Inject GetAzureNameFromConfig method on storage specs that have the AzureNameFromConfig property",
		func(ctx context.Context, state *State) (*State, error) {
			definitions := state.Definitions()
			result := definitions.Copy()

			propName := idFactory.CreatePropertyName(astmodel.AzureNameFromConfigProperty, astmodel.Exported)

			// Find all storage specs that have the AzureNameFromConfig property
			storageSpecs := astmodel.FindSpecDefinitions(definitions).Where(func(def astmodel.TypeDefinition) bool {
				if !astmodel.IsStoragePackageReference(def.Name().PackageReference()) {
					return false
				}

				ot, ok := astmodel.AsObjectType(def.Type())
				if !ok {
					return false
				}

				_, hasProp := ot.Property(propName)
				return hasProp
			})

			fnInjector := astmodel.NewFunctionInjector()
			for name, def := range storageSpecs {
				fn := functions.NewObjectFunction(
					astmodel.GetAzureNameFromConfigFunc,
					idFactory,
					getAzureNameFromConfigFunctionHandler,
					astmodel.GenRuntimeReference,
				)

				updated, err := fnInjector.Inject(def, fn)
				if err != nil {
					return nil, eris.Wrapf(err, "injecting GetAzureNameFromConfig into %s", name)
				}

				result[updated.Name()] = updated
			}

			return state.WithDefinitions(result), nil
		},
	)
}

// getAzureNameFromConfigFunctionHandler generates:
//
//	func (spec *MyResource_Spec) GetAzureNameFromConfig() *genruntime.ConfigMapReference {
//	    return spec.AzureNameFromConfig
//	}
func getAzureNameFromConfigFunctionHandler(
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
