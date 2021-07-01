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

// InjectHubFunctionStageId is the unique identifier for this pipeline stage
const InjectHubFunctionStageId = "injectHubFunction"

// InjectHubFunction modifies the nominates storage version (aka hub version) of each resource by injecting a Hub()
// function so that it satisfies the required interface.
func InjectHubFunction(idFactory astmodel.IdentifierFactory) Stage {

	result := MakeStage(
		InjectHubFunctionStageId,
		"Inject the function Hub() into each hub resource",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			injector := storage.NewFunctionInjector()
			result := types.Copy()

			resources := storage.FindResourceTypes(types)
			for name, def := range resources {
				rt, ok := astmodel.AsResourceType(def.Type())
				if !ok {
					return nil, errors.Errorf("expected %s to be a resource type (should never happen)", name)
				}

				if rt.IsStorageVersion() {
					fn := functions.NewHubFunction(idFactory)
					defWithFn, err := injector.Inject(def, fn)
					if err != nil {
						return nil, errors.Wrapf(err, "injecting Hub() into %s", name)
					}

					result[name] = defWithFn
				}
			}

			return result, nil
		})

	result.RequiresPrerequisiteStages(MarkStorageVersionStageId)
	return result
}
