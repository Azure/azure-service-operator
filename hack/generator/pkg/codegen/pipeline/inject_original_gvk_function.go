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

// injectOriginalGVKFunctionId is the unique identifier for this pipeline stage
const injectOriginalGVKFunctionId = "injectOriginalGVKFunction"

// InjectOriginalGVKFunction injects the function OriginalGVK() into each Resource type
// This function allows us to recover the original version used to create each custom resource, giving the operator the
// information needed to interact with ARM using the correct API version.
func InjectOriginalGVKFunction(idFactory astmodel.IdentifierFactory) Stage {

	stage := MakeStage(
		injectOriginalGVKFunctionId,
		"Inject the function OriginalGVK() into each Resource type",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			injector := storage.NewFunctionInjector()
			result := types.Copy()

			resources := storage.FindResourceTypes(types)
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

	stage.RequiresPrerequisiteStages(injectOriginalVersionFunctionStageId)
	return stage
}
