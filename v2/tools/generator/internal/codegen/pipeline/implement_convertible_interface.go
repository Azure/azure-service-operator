/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// ImplementConvertibleInterfaceStageID is the unique identifier for this pipeline stage
const ImplementConvertibleInterfaceStageID = "implementConvertibleInterface"

// ImplementConvertibleInterface injects the functions ConvertTo() and ConvertFrom() into each non-hub Resource
// Type, providing the required implementation of the Convertible interface needed by the controller
func ImplementConvertibleInterface(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		ImplementConvertibleInterfaceStageID,
		"Implement the Convertible interface on each non-hub Resource type",
		func(ctx context.Context, state *State) (*State, error) {
			injector := astmodel.NewInterfaceInjector()

			modifiedTypes, err := astmodel.FindResourceDefinitions(state.Definitions()).Process(
				func(def astmodel.TypeDefinition) (*astmodel.TypeDefinition, error) {
					graph, err := GetStateData[*storage.ConversionGraph](state, ConversionGraphInfo)
					if err != nil {
						return nil, eris.Wrapf(err, "couldn't find conversion graph")
					}

					hub, err := graph.FindHub(def.Name(), state.Definitions())
					if err != nil {
						return nil, eris.Wrapf(
							err,
							"finding hub for %s",
							def.Name())
					}

					if astmodel.TypeEquals(def.Name(), hub) {
						// The hub storage version doesn't implement Convertible
						return nil, nil
					}

					// For each PropertyAssignmentFunction, create a conversion function that uses it
					rsrc := astmodel.MustBeResourceType(def.Type())
					var conversionFunctions []astmodel.Function
					for _, fn := range rsrc.Functions() {
						if propertyAssignmentFn, ok := fn.(*functions.PropertyAssignmentFunction); ok {

							conversionFn := functions.NewResourceConversionFunction(hub, propertyAssignmentFn, idFactory)
							conversionFunctions = append(conversionFunctions, conversionFn)
						}
					}

					// Create the interface implementation and inject into the rsrc
					impl := astmodel.NewInterfaceImplementation(astmodel.ConvertibleInterface, conversionFunctions...)

					modified, err := injector.Inject(def, impl)
					if err != nil {
						return nil, eris.Wrapf(err, "injecting conversions.Convertible interface into %s", def.Name())
					}

					return &modified, nil
				})
			if err != nil {
				return nil, eris.Wrap(err, "injecting conversions.Convertible implementations")
			}

			return state.WithOverlaidDefinitions(modifiedTypes), nil
		})

	stage.RequiresPrerequisiteStages(InjectPropertyAssignmentFunctionsStageID)
	return stage
}
