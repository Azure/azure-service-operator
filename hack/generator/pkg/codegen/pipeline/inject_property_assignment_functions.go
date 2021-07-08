/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
)

// injectPropertyAssignmentFunctionsStageId is the unique identifier for this pipeline stage
const injectPropertyAssignmentFunctionsStageId = "injectPropertyAssignmentFunctions"

// InjectPropertyAssignmentFunctions injects property assignment functions AssignTo*() and AssignFrom*() into both
// resources and object types. These functions do the heavy lifting of the conversions between versions of each type and
// are the building blocks of the main CovertTo*() and ConvertFrom*() methods.
func InjectPropertyAssignmentFunctions(graph *storage.ConversionGraph, idFactory astmodel.IdentifierFactory) Stage {

	stage := MakeStage(
		injectPropertyAssignmentFunctionsStageId,
		"Inject property assignment functions AssignFrom() and AssignTo() into resources and objects",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			result := types.Copy()
			factory := NewPropertyAssignmentFunctionsFactory(graph, idFactory, types)

			for name, def := range types {
				_, ok := astmodel.AsFunctionContainer(def.Type())
				if !ok {
					// just skip it - not a resource nor an object
					klog.V(4).Infof("Skipping %s as no conversion functions needed", name)
					continue
				}

				klog.V(3).Infof("Injecting conversion functions into %s", name)

				// Find the definition we want to convert to/from
				nextPackage, ok := graph.LookupTransition(name.PackageReference)
				if !ok {
					// No next package, so nothing to do
					// (this is expected if we have the hub storage package)
					continue
				}

				nextName := astmodel.MakeTypeName(nextPackage, name.Name())
				nextDef, ok := types[nextName]
				if !ok {
					// No next type so nothing to do
					// (this is expected if the type is discontinued or we're looking at the hub type)
					continue
				}

				modified, err := factory.injectBetween(def, nextDef)
				if err != nil {
					return nil, errors.Wrapf(err, "injecting property assignment functions into %s", name)
				}

				result[modified.Name()] = modified
			}

			return result, nil
		})

	// Needed to populate the conversion graph
	stage.RequiresPrerequisiteStages(CreateStorageTypesStageId)
	return stage
}

type propertyAssignmentFunctionsFactory struct {
	graph            *storage.ConversionGraph
	idFactory        astmodel.IdentifierFactory
	types            astmodel.Types
	functionInjector *storage.FunctionInjector
}

func NewPropertyAssignmentFunctionsFactory(
	graph *storage.ConversionGraph,
	idFactory astmodel.IdentifierFactory,
	types astmodel.Types) *propertyAssignmentFunctionsFactory {
	return &propertyAssignmentFunctionsFactory{
		graph:            graph,
		idFactory:        idFactory,
		types:            types,
		functionInjector: storage.NewFunctionInjector(),
	}
}

// injectBetween injects conversion methods between the two specified definitions
// upstreamDef is the definition further away from our hub type in our directed conversion graph
// downstreamDef is the definition closer to our hub type in our directed conversion graph
func (f propertyAssignmentFunctionsFactory) injectBetween(
	upstreamDef astmodel.TypeDefinition, downstreamDef astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {

	// Create conversion functions
	conversionContext := conversions.NewPropertyConversionContext(f.types, f.idFactory)

	assignFromFn, err := functions.NewPropertyAssignmentFromFunction(upstreamDef, downstreamDef, f.idFactory, conversionContext)
	upstreamName := upstreamDef.Name()
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "creating AssignFrom() function for %q", upstreamName)
	}

	assignToFn, err := functions.NewPropertyAssignmentToFunction(upstreamDef, downstreamDef, f.idFactory, conversionContext)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "creating AssignTo() function for %q", upstreamName)
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
