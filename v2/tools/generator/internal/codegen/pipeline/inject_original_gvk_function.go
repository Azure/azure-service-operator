/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// InjectOriginalGVKFunctionStageID is the unique identifier for this pipeline stage
const InjectOriginalGVKFunctionStageID = "injectOriginalGVKFunction"

// InjectOriginalGVKFunction injects the function OriginalGVK() into each Resource type
// This function allows us to recover the original version used to create each custom resource, giving the operator the
// information needed to interact with ARM using the correct API version.
func InjectOriginalGVKFunction(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewLegacyStage(
		InjectOriginalGVKFunctionStageID,
		"Inject the function OriginalGVK() into each Resource type",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			injector := astmodel.NewFunctionInjector()
			result := definitions.Copy()

			resources := astmodel.FindResourceDefinitions(definitions)
			for name, def := range resources {
				var fn *functions.OriginalGVKFunction
				if astmodel.IsStoragePackageReference(name.PackageReference) {
					fn = functions.NewOriginalGVKFunction(functions.ReadProperty, idFactory)
				} else {
					fn = functions.NewOriginalGVKFunction(functions.ReadFunction, idFactory)
				}

				defWithFn, err := injector.Inject(def, fn)
				if err != nil {
					return nil, errors.Wrapf(err, "injecting OriginalGVK() into %s", name)
				}

				result[defWithFn.Name()] = defWithFn
			}

			return result, nil
		})

	stage.RequiresPrerequisiteStages(InjectOriginalVersionFunctionStageID)
	return stage
}
