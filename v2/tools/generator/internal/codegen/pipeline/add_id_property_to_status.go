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

const AddIDPropertyToStatusID = "addIdPropToStatus"

// Note: we expect that everything returned from ARM has an Id property even if it is not document in Swagger.
func AddIDPropertyToStatus() *Stage {
	return NewStage(
		AddIDPropertyToStatusID,
		"Adds an 'Id' property to all top-level status types",
		addIDPropertyToStatus)
}

func addIDPropertyToStatus(ctx context.Context, state *State) (*State, error) {
	newDefs := make(astmodel.TypeDefinitionSet)
	for _, def := range state.Definitions() {
		rt, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			continue
		}

		statusType, err := state.Definitions().FullyResolve(rt.StatusType())
		if err != nil {
			return nil, err
		}

		ot, ok := astmodel.AsObjectType(statusType)
		if !ok {
			return nil, errors.New("expected status type to be object type")
		}

		_, ok = ot.Property(idPropName)
		if ok {
			continue // already has "id" prop
		}

		ot = ot.WithProperty(idProp)
		newDefs.Add(def.WithType(rt.WithStatus(ot)))
	}

	return state.WithDefinitions(state.Definitions().OverlayWith(newDefs)), nil
}

var idProp = astmodel.NewPropertyDefinition("Id", "id", astmodel.StringType).WithDescription("The ARM Id for this resource.")
var idPropName = astmodel.PropertyName("Id")
