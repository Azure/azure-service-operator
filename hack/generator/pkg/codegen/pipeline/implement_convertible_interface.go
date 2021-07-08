/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
)

// ImplementConvertibleInterfaceStageId is the unique identifier for this pipeline stage
const ImplementConvertibleInterfaceStageId = "implementConvertibleInterface"

// ImplementConvertibleInterface injects the functions ConvertTo() and ConvertFrom() into each non-hub Resource
// Type, providing the required implementation of the Convertible interface needed by the controller
func ImplementConvertibleInterface(idFactory astmodel.IdentifierFactory) Stage {

	stage := MakeStage(
		ImplementConvertibleInterfaceStageId,
		"Implement the Convertible interface on each non-hub Resource type",
		func(ctx context.Context, state *State) (*State, error) {
			injector := astmodel.NewInterfaceInjector()

			modifiedTypes := make(astmodel.Types)
			resources := storage.FindResourceTypes(state.Types())
			for name, def := range resources {
				resource, ok := astmodel.AsResourceType(def.Type())
				if !ok {
					// Skip non-resources (though, they should be filtered out, above)
					continue
				}

				if resource.IsStorageVersion() {
					// The hub storage version doesn't implement Convertible
					continue
				}

				convertible := createConvertibleInterfaceImplementation(
					name, resource, state.ConversionGraph(), idFactory)
				if convertible.FunctionCount() > 0 {
					modified, err := injector.Inject(def, convertible)
					if err != nil {
						return nil, errors.Wrapf(err, "injecting Convertible interface into %s", name)
					}

					modifiedTypes.Add(modified)
				}
			}

			newTypes := state.Types().OverlayWith(modifiedTypes)
			return state.WithTypes(newTypes), nil
		})

	stage.RequiresPrerequisiteStages(InjectPropertyAssignmentFunctionsStageID)
	return stage
}

// createConvertibleInterfaceImplementation creates the required implementation of conversion.Convertible, ready for
// injection onto the resource. The ConvertTo() and ConvertFrom() methods chain the required conversion between resource
// versions, but are dependent upon previously injected AssignPropertiesTo() and AssignPropertiesFrom() methods to
// actually copy information across. See resource_conversion_function.go for more information.
func createConvertibleInterfaceImplementation(
	name astmodel.TypeName,
	resource *astmodel.ResourceType,
	conversionGraph *storage.ConversionGraph,
	idFactory astmodel.IdentifierFactory) *astmodel.InterfaceImplementation {
	var conversionFunctions []astmodel.Function

	for _, fn := range resource.Functions() {
		if propertyAssignmentFn, ok := fn.(*functions.PropertyAssignmentFunction); ok {
			hub := conversionGraph.FindHubTypeName(name)
			conversionFn := functions.NewResourceConversionFunction(hub, propertyAssignmentFn, idFactory)
			conversionFunctions = append(conversionFunctions, conversionFn)
		}
	}

	return astmodel.NewInterfaceImplementation(astmodel.ConvertibleInterface, conversionFunctions...)
}
