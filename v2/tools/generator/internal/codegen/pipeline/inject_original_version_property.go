/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// InjectOriginalVersionPropertyStageID is the unique identifier for this pipeline stage
const InjectOriginalVersionPropertyStageID = "injectOriginalVersionProperty"

// InjectOriginalVersionProperty injects the property OriginalVersion into each Storage Spec type
// This property gets populated by reading from the OriginalVersion() function previously injected into the API Spec
// types, allowing us to recover the original version used to create each custom resource, and giving the operator the
// information needed to interact with ARM using the correct API version.
func InjectOriginalVersionProperty() *Stage {
	stage := NewLegacyStage(
		InjectOriginalVersionPropertyStageID,
		"Inject the property OriginalVersion into each Storage Spec type",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			injector := astmodel.NewPropertyInjector()
			result := definitions.Copy()

			doesNotHaveOriginalVersionFunction := func(definition astmodel.TypeDefinition) bool {
				ot, ok := astmodel.AsObjectType(definition.Type())
				if !ok {
					// Not an object
					return false
				}

				// Skip objects that have OriginalVersion() functions
				return !ot.HasFunctionWithName("OriginalVersion")
			}

			storageSpecs := astmodel.FindSpecDefinitions(definitions).Where(doesNotHaveOriginalVersionFunction)

			for name, def := range storageSpecs {
				prop := astmodel.NewPropertyDefinition("OriginalVersion", "originalVersion", astmodel.StringType)
				prop.WithDescription("returns the original API version used to create the resource")
				defWithProp, err := injector.Inject(def, prop)
				if err != nil {
					return nil, errors.Wrapf(err, "injecting OriginalVersion into %s", name)
				}

				result[defWithProp.Name()] = defWithProp
			}

			return result, nil
		})

	return stage
}
