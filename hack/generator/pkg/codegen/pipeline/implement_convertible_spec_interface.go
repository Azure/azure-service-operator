package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
)

// ImplementConvertibleSpecInterfaceStageId is the unique identifier for this pipeline stage
const ImplementConvertibleSpecInterfaceStageId = "implementConvertibleSpecInterface"

func ImplementConvertibleSpecInterface(idFactory astmodel.IdentifierFactory) Stage {
	stage := MakeStage(
		ImplementConvertibleSpecInterfaceStageId,
		"Inject ConvertSpecTo() and ConvertSpecFrom() to implement genruntime.ConvertibleSpec",
		func(ctx context.Context, state *State) (*State, error) {
			injector := astmodel.NewInterfaceInjector()

			modifiedTypes := make(astmodel.Types)
			specs := storage.FindSpecTypes(state.Types())
			for name, def := range specs {
				convertible := createConvertibleSpecInterfaceImplementation(def, idFactory)
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
