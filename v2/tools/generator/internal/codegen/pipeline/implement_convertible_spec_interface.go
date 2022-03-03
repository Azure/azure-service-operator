/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// ImplementConvertibleSpecInterfaceStageId is the unique identifier for this pipeline stage
const ImplementConvertibleSpecInterfaceStageId = "implementConvertibleSpecInterface"

func ImplementConvertibleSpecInterface(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		ImplementConvertibleSpecInterfaceStageId,
		"Inject ConvertSpecTo() and ConvertSpecFrom() to implement genruntime.ConvertibleSpec on each Spec type",
		func(ctx context.Context, state *State) (*State, error) {
			injector := astmodel.NewInterfaceInjector()

			modifiedDefinitions := make(astmodel.TypeDefinitionSet)
			specs := astmodel.FindSpecDefinitions(state.Definitions())
			for name, def := range specs {
				convertible := createConvertibleSpecInterfaceImplementation(def, idFactory)
				modified, err := injector.Inject(def, convertible)
				if err != nil {
					return nil, errors.Wrapf(err, "injecting Convertible interface into %s", name)
				}

				modifiedDefinitions.Add(modified)
			}

			defs := state.Definitions().OverlayWith(modifiedDefinitions)
			return state.WithDefinitions(defs), nil
		})

	stage.RequiresPrerequisiteStages(InjectPropertyAssignmentFunctionsStageID)
	return stage
}

// createConvertibleSpecInterfaceImplementation creates both of the funcs required for a given spec to implement the
// genruntime.ConvertibleSpec interface. See ChainedConversionFunction and PivotConversionFunction for details of the
// actual code generated.
func createConvertibleSpecInterfaceImplementation(
	spec astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory) *astmodel.InterfaceImplementation {
	container, ok := astmodel.AsFunctionContainer(spec.Type())
	if !ok {
		// This shouldn't happen due to earlier filtering
		return nil
	}

	fnFrom := createConvertibleSpecFunction(conversions.ConvertFrom, container, idFactory)
	fnTo := createConvertibleSpecFunction(conversions.ConvertTo, container, idFactory)

	return astmodel.NewInterfaceImplementation(astmodel.ConvertibleSpecInterfaceType, fnFrom, fnTo)
}

// createConvertibleSpecFunction creates a conversion function for the specified direction. If a suitable property
// assignment function can be found, we create a chained conversion function, otherwise we create a pivot (assuming that
// that we have the hub spec type)
// direction is the direction of the required conversion
// container is the function container we iterate to look for a property assignment function
// idFactory is a reference to our shared identifier factory
func createConvertibleSpecFunction(
	direction conversions.Direction,
	container astmodel.FunctionContainer,
	idFactory astmodel.IdentifierFactory) astmodel.Function {

	for _, fn := range container.Functions() {
		if propertyAssignmentFn, ok := fn.(*functions.PropertyAssignmentFunction); ok {
			if propertyAssignmentFn.Direction() != direction {
				continue
			}

			return functions.NewSpecChainedConversionFunction(propertyAssignmentFn, idFactory)
		}
	}

	return functions.NewSpecPivotConversionFunction(direction, idFactory)
}
