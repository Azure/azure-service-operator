/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
)

// ImplementConvertibleStatusInterfaceStageId is the unique identifier for this pipeline stage
const ImplementConvertibleStatusInterfaceStageId = "implementConvertibleStatusInterface"

func ImplementConvertibleStatusInterface(idFactory astmodel.IdentifierFactory) Stage {
	stage := MakeStage(
		ImplementConvertibleStatusInterfaceStageId,
		"Inject ConvertStatusTo() and ConvertStatusFrom() to implement genruntime.ConvertibleStatus",
		func(ctx context.Context, state *State) (*State, error) {
			injector := astmodel.NewInterfaceInjector()

			modifiedTypes := make(astmodel.Types)
			statuses := astmodel.FindStatusTypes(state.Types())
			for name, def := range statuses {
				convertible := createConvertibleStatusInterfaceImplementation(def, idFactory)
				modified, err := injector.Inject(def, convertible)
				if err != nil {
					return nil, errors.Wrapf(err, "injecting Convertible interface into %s", name)
				}

				modifiedTypes.Add(modified)
			}

			newTypes := state.Types().OverlayWith(modifiedTypes)
			return state.WithTypes(newTypes), nil
		})

	stage.RequiresPrerequisiteStages(InjectPropertyAssignmentFunctionsStageID)
	return stage
}

func createConvertibleStatusInterfaceImplementation(
	status astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory) *astmodel.InterfaceImplementation {
	container, ok := astmodel.AsFunctionContainer(status.Type())
	if !ok {
		// This shouldn't happen due to earlier filtering
		return nil
	}

	fnFrom := createConvertibleStatusFunction(conversions.ConvertFrom, container, idFactory)
	fnTo := createConvertibleStatusFunction(conversions.ConvertTo, container, idFactory)

	return astmodel.NewInterfaceImplementation(astmodel.ConvertibleStatusInterfaceType, fnFrom, fnTo)
}

// createConvertibleStatusFunction creates a conversion function for the specified direction. If a suitable property
// assignment function can be found, we create a chained conversion function, otherwise we create a pivot (assuming that
// that we have the hub status type)
// direction is the direction of the required conversion
// container is the function container we iterate to look for a property assignment function
// idFactory is a reference to our shared identifier factory
func createConvertibleStatusFunction(
	direction conversions.Direction,
	container astmodel.FunctionContainer,
	idFactory astmodel.IdentifierFactory) astmodel.Function {

	for _, fn := range container.Functions() {
		if propertyAssignmentFn, ok := fn.(*functions.PropertyAssignmentFunction); ok {
			if propertyAssignmentFn.Direction() != direction {
				continue
			}

			return functions.NewStatusChainedConversionFunction(propertyAssignmentFn, idFactory)
		}
	}

	return functions.NewStatusPivotConversionFunction(direction, idFactory)
}
