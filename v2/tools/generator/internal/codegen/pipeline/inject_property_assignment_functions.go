/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/dave/dst"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// InjectPropertyAssignmentFunctionsStageID is the unique identifier for this pipeline stage
const InjectPropertyAssignmentFunctionsStageID = "injectPropertyAssignmentFunctions"

// InjectPropertyAssignmentFunctions injects property assignment functions AssignTo*() and AssignFrom*() into both
// resources and object types. These functions do the heavy lifting of the conversions between versions of each type and
// are the building blocks of the main CovertTo*() and ConvertFrom*() methods.
func InjectPropertyAssignmentFunctions(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectPropertyAssignmentFunctionsStageID,
		"Inject property assignment functions AssignFrom() and AssignTo() into resources and objects",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			result := defs.Copy()
			factory := newPropertyAssignmentFunctionsFactory(state.ConversionGraph(), idFactory, configuration, defs)

			for name, def := range defs {
				_, ok := astmodel.AsFunctionContainer(def.Type())
				if !ok {
					// just skip it - not a resource nor an object
					klog.V(4).Infof("Skipping %s as no conversion functions needed", name)
					continue
				}

				klog.V(3).Infof("Injecting conversion functions into %s", name)

				// Find the definition we want to convert to/from
				nextName, err := state.ConversionGraph().FindNextType(name, state.Definitions())
				if err != nil {
					return nil, errors.Wrapf(err, "finding next type after %s", name)
				}

				if nextName.IsEmpty() {
					// No next type, so nothing to do (this is expected if we have the hub storage package)
					continue
				}

				nextDef, ok := defs[nextName]
				if !ok {
					// No next type so nothing to do
					// (this is expected if the type is discontinued, or we're looking at the hub type)
					continue
				}

				var augmentationInterface *astmodel.TypeName
				if astmodel.IsStoragePackageReference(def.Name().PackageReference) {
					ifaceType := astmodel.NewInterfaceType()

					assignPropertiesToFunc := functions.NewObjectFunction(
						"AssignPropertiesTo",
						idFactory,
						createAssignPropertiesOverrideStub("dst", astmodel.NewOptionalType(nextDef.Name())))
					assignPropertiesToFunc.AddPackageReference(nextDef.Name().PackageReference)

					assignPropertiesFromFunc := functions.NewObjectFunction(
						"AssignPropertiesFrom",
						idFactory,
						createAssignPropertiesOverrideStub("src", astmodel.NewOptionalType(nextDef.Name())))
					assignPropertiesFromFunc.AddPackageReference(nextDef.Name().PackageReference)

					ifaceType = ifaceType.WithFunction(assignPropertiesToFunc).WithFunction(assignPropertiesFromFunc)

					augmentationInterfaceName := "augmentConversionFor" + idFactory.CreateIdentifier(def.Name().Name(), astmodel.Exported)
					augmentationInterfaceTypeName := def.Name().WithName(augmentationInterfaceName)
					augmentationInterface = &augmentationInterfaceTypeName
					ifaceDef := astmodel.MakeTypeDefinition(
						augmentationInterfaceTypeName,
						ifaceType)
					result.Add(ifaceDef)
				}

				modified, err := factory.injectBetween(def, nextDef, augmentationInterface)
				if err != nil {
					return nil, errors.Wrapf(err, "injecting property assignment functions into %s", name)
				}

				result[modified.Name()] = modified
			}

			return state.WithDefinitions(result), nil
		})

	// Needed to populate the conversion graph
	stage.RequiresPrerequisiteStages(CreateStorageTypesStageID)
	return stage
}

type propertyAssignmentFunctionsFactory struct {
	graph            *storage.ConversionGraph
	idFactory        astmodel.IdentifierFactory
	configuration    *config.Configuration
	definitions      astmodel.TypeDefinitionSet
	functionInjector *astmodel.FunctionInjector
}

func newPropertyAssignmentFunctionsFactory(
	graph *storage.ConversionGraph,
	idFactory astmodel.IdentifierFactory,
	configuration *config.Configuration,
	definitions astmodel.TypeDefinitionSet) *propertyAssignmentFunctionsFactory {
	return &propertyAssignmentFunctionsFactory{
		graph:            graph,
		idFactory:        idFactory,
		configuration:    configuration,
		definitions:      definitions,
		functionInjector: astmodel.NewFunctionInjector(),
	}
}

// injectBetween injects conversion methods between the two specified definitions
// upstreamDef is the definition further away from our hub type in our directed conversion graph
// downstreamDef is the definition closer to our hub type in our directed conversion graph
func (f propertyAssignmentFunctionsFactory) injectBetween(
	upstreamDef astmodel.TypeDefinition,
	downstreamDef astmodel.TypeDefinition,
	augmentationInterface *astmodel.TypeName) (astmodel.TypeDefinition, error) {

	assignmentContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, f.definitions, f.idFactory).
		WithConfiguration(f.configuration.ObjectModelConfiguration).
		WithConversionGraph(f.graph)

	// Create conversion functions
	assignFromBuilder := functions.NewPropertyAssignmentFunctionBuilder(upstreamDef, downstreamDef, conversions.ConvertFrom)
	assignFromFn, err := assignFromBuilder.Build(assignmentContext)
	upstreamName := upstreamDef.Name()
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "creating AssignFrom() function for %q", upstreamName)
	}

	if augmentationInterface != nil {
		assignFromFn = assignFromFn.WithAugmentationInterface(*augmentationInterface)
	}

	assignToBuilder := functions.NewPropertyAssignmentFunctionBuilder(upstreamDef, downstreamDef, conversions.ConvertTo)
	assignToFn, err := assignToBuilder.Build(assignmentContext)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "creating AssignTo() function for %q", upstreamName)
	}
	if augmentationInterface != nil {
		assignToFn = assignToFn.WithAugmentationInterface(*augmentationInterface)
	}

	updatedDefinition, err := f.functionInjector.Inject(upstreamDef, assignFromFn)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "failed to inject %s function into %q", assignFromFn.Name(), upstreamName)
	}

	updatedDefinition, err = f.functionInjector.Inject(updatedDefinition, assignToFn)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "failed to inject %s function into %q", assignToFn.Name(), upstreamName)
	}

	return updatedDefinition, nil
}

func createAssignPropertiesOverrideStub(
	paramName string,
	paramType astmodel.Type) func(f *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	return func(f *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		funcDetails := &astbuilder.FuncDetails{
			Name: methodName,
		}
		funcDetails.AddParameter(paramName, paramType.AsType(codeGenerationContext))
		funcDetails.AddReturn(dst.NewIdent("error"))

		return funcDetails.DefineFuncHeader()
	}
}
