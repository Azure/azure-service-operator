/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const AddAPIVersionEnumsStageId = "add-api-version-enums"

func AddAPIVersionEnums() *Stage {
	stage := NewStage(
		AddAPIVersionEnumsStageId,
		"Add enums for API Versions in each package",
		func(ctx context.Context, state *State) (*State, error) {
			newDefs := make(astmodel.TypeDefinitionSet)
			for _, def := range state.Definitions() {
				if rt, ok := astmodel.AsResourceType(def.Type()); ok && rt.HasAPIVersion() {
					apiVersionName := rt.APIVersionTypeName()
					if !newDefs.Contains(apiVersionName) {
						klog.Infof("GENERATED ENUM: %s", rt.APIVersionTypeName())
						apiVersionValue := rt.APIVersionEnumValue()
						enumType := astmodel.NewEnumType(astmodel.StringType, apiVersionValue)
						newDefs.Add(astmodel.MakeTypeDefinition(apiVersionName, enumType))
					}
				}
			}

			state.definitions.AddTypes(newDefs)
			return state, nil
		},
	)

	stage.RequiresPrerequisiteStages(RemoveAPIVersionPropertyStageID)
	return stage
}
