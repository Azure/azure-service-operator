/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TransformValidatedFloatsStageID is the unique identifier of this stage
const (
	TransformValidatedFloatsStageID = "transformValidatedFloats"
)

func TransformValidatedFloats() *Stage {
	stage := NewStage(
		TransformValidatedFloatsStageID,
		"Transform validated 'spec' float type values to validated integer types for compatibility with controller-gen",
		func(ctx context.Context, state *State) (*State, error) {

			definitions := state.Definitions()

			result, err := getFloatTransformations(definitions)
			if err != nil {
				return nil, err
			}

			remaining := definitions.Except(result)
			result.AddTypes(remaining)

			return state.WithDefinitions(result), nil
		})

	stage.RequiresPrerequisiteStages(RemoveStatusPropertyValidationsStageID)

	return stage
}

func getFloatTransformations(definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	visitor := astmodel.TypeVisitorBuilder[any]{
		VisitValidatedType: visitValidatedType,
		VisitPrimitive:     transformFloatToInt,
	}.Build()

	return visitor.VisitDefinitions(astmodel.FindSpecConnectedDefinitions(definitions))
}

func visitValidatedType(this *astmodel.TypeVisitor[any], validated *astmodel.ValidatedType, ctx any) (astmodel.Type, error) {
	return astmodel.IdentityVisitOfValidatedType(this, validated, true) // Pass ctx so that transformFloatToInt can use it
}

// transformFloatToInt transforms all the validated FloatTypes to IntegerTypes
func transformFloatToInt(this *astmodel.TypeVisitor[any], prim *astmodel.PrimitiveType, ctx any) (astmodel.Type, error) {
	validated, ok := ctx.(bool)
	if prim == astmodel.FloatType && ok && validated {
		return astmodel.IntType, nil
	}
	return astmodel.IdentityVisitOfPrimitiveType(this, prim, ctx)
}
