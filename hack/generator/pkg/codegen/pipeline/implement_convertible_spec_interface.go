package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
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
				convertible := createConvertibleSpecInterfaceImplementation(
					def, state.ConversionGraph(), idFactory)
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

func createConvertibleSpecInterfaceImplementation(
	spec astmodel.TypeDefinition,
	conversionGraph *storage.ConversionGraph,
	idFactory astmodel.IdentifierFactory) *astmodel.InterfaceImplementation {
	var specConversions []astmodel.Function

	obj, ok := astmodel.AsObjectType(spec.Type())
	if !ok {
		// This shouldn't happen due to earlier filtering
		return nil
	}

	for _, fn := range obj.Functions() {
		if propertyAssignmentFn, ok := fn.(*functions.PropertyAssignmentFunction); ok {
			hub := conversionGraph.FindHubTypeName(spec.Name())
			conversionFn := functions.NewSpecConversionFunction(hub, propertyAssignmentFn, idFactory)
			specConversions = append(specConversions, conversionFn)
		}
	}

	return astmodel.NewInterfaceImplementation(astmodel.ConvertibleSpecInterfaceType, specConversions...)
}
