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

const AddStatusConditionsStageID = "addStatusConditions"

func AddStatusConditions(idFactory astmodel.IdentifierFactory) Stage {
	return MakeStage(
		AddStatusConditionsStageID,
		"Add the property 'Conditions' to all status types and implements genruntime.Conditioner on all resources",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Types()
			result := make(astmodel.Types)

			propInjector := astmodel.NewPropertyInjector()
			statusDefs := astmodel.FindStatusTypes(defs)
			for _, def := range statusDefs {
				conditionsProp := astmodel.NewPropertyDefinition(
					astmodel.ConditionsProperty,
					"conditions",
					astmodel.NewArrayType(astmodel.ConditionType))
				conditionsProp = conditionsProp.WithDescription("The observed state of the resource").MakeOptional()
				updatedDef, err := propInjector.Inject(def, conditionsProp)
				if err != nil {
					return nil, errors.Wrapf(err, "couldn't add Conditions condition to status %q", def.Name())
				}
				result.Add(updatedDef)
			}

			resourceDefs := astmodel.FindResourceTypes(defs)
			for _, def := range resourceDefs {
				resourceType := def.Type().(*astmodel.ResourceType)

				conditionerImpl, err := astmodel.NewConditionerInterfaceImpl(idFactory, resourceType)
				if err != nil {
					return nil, errors.Wrapf(err, "couldn't create genruntime.Conditioner implementation for %q", def.Name())
				}
				resourceType = resourceType.WithInterface(conditionerImpl)

				// Resources with the genruntime.Conditioner interface should also have kubebuilder:printcolumn set
				// so that the conditions are displayed
				resourceType = resourceType.WithAnnotation("// +kubebuilder:printcolumn:name=\"Ready\",type=\"string\",JSONPath=\".status.conditions[?(@.type=='Ready')].status\"")
				resourceType = resourceType.WithAnnotation("// +kubebuilder:printcolumn:name=\"Reason\",type=\"string\",JSONPath=\".status.conditions[?(@.type=='Ready')].reason\"")
				resourceType = resourceType.WithAnnotation("// +kubebuilder:printcolumn:name=\"Message\",type=\"string\",JSONPath=\".status.conditions[?(@.type=='Ready')].message\"")

				result.Add(def.WithType(resourceType))
			}
			result = defs.OverlayWith(result)

			return state.WithTypes(result), nil
		})
}
