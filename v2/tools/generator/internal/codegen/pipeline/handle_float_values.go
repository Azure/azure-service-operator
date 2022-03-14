/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// HandleFloatValuesStageID is the unique identifier of this stage
const (
	HandleFloatValuesStageID = "handleFloatValues"
)

func HandleFloatValues() *Stage {
	stage := NewStage(
		HandleFloatValuesStageID,
		"Transform Float type values to Integer types",
		func(ctx context.Context, state *State) (*State, error) {

			definitions := state.Definitions()
			specDefinitions := astmodel.FindSpecDefinitions(definitions)

			result, errs := getFloatTransformations(definitions, specDefinitions)
			if err := kerrors.NewAggregate(errs); err != nil {
				return nil, err
			}

			remaining := definitions.Except(result)
			result.AddTypes(remaining)

			return state.WithDefinitions(result), nil
		})

	stage.RequiresPrerequisiteStages(RemoveStatusPropertyValidationsStageID)

	return stage
}

func getFloatTransformations(definitions astmodel.TypeDefinitionSet, validatedDefinitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, []error) {
	walker := astmodel.NewTypeWalker(
		definitions,
		astmodel.TypeVisitorBuilder{
			VisitObjectType: transformFloatToInt,
		}.Build())

	return walkTypes(validatedDefinitions, walker)
}

// transformFloatToInt transforms all the FloatTypes to IntergerTypes
func transformFloatToInt(this *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	for _, prop := range ot.Properties() {
		// check if the base type is Float64 Type
		if result, ok := astmodel.AsPrimitiveType(prop.PropertyType()); ok && result == astmodel.FloatType {
			// Check if the Float64 Type has validations. If yes, then we have to change it to IntType as code-gen doesn't
			// allow float types with validations.
			if validatedType, ok := prop.PropertyType().(*astmodel.ValidatedType); ok {
				ot = ot.WithProperty(prop.WithType(validatedType.WithType(astmodel.IntType)))

			}
		}
	}

	return astmodel.IdentityVisitOfObjectType(this, ot, ctx)
}

func walkTypes(definitions astmodel.TypeDefinitionSet, walker *astmodel.TypeWalker) (astmodel.TypeDefinitionSet, []error) {
	var errs []error
	result := make(astmodel.TypeDefinitionSet)

	for _, def := range definitions {
		visitedType, err := walker.Walk(def)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "failed walking definitions"))
		}

		err = result.AddTypesAllowDuplicates(visitedType)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return result, errs
}
