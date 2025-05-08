/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// StripUnreferencedTypeDefinitionsStageID is the unique identifier for this pipeline stage
const StripUnreferencedTypeDefinitionsStageID = "stripUnreferenced"

func StripUnreferencedTypeDefinitions() *Stage {
	return NewStage(
		StripUnreferencedTypeDefinitionsStageID,
		"Strip unreferenced types",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			resources := astmodel.FindResourceDefinitions(defs)
			armSpecAndStatus := astmodel.CollectARMSpecAndStatusDefinitions(defs)
			roots := astmodel.SetUnion(resources.Names(), armSpecAndStatus)

			result, err := StripUnusedDefinitions(roots, defs)
			if err != nil {
				return nil, err
			}

			return state.WithDefinitions(result), nil
		})
}

// StripUnusedDefinitions removes all types that aren't in roots or
// referred to by the types in roots, for example types that are
// generated as a byproduct of an allOf element.
func StripUnusedDefinitions(
	roots astmodel.TypeNameSet,
	defs astmodel.TypeDefinitionSet,
) (astmodel.TypeDefinitionSet, error) {
	graph := astmodel.MakeReferenceGraphWithRoots(roots, defs)
	connectedDefinitions := graph.Connected()

	usedDefinitions := make(astmodel.TypeDefinitionSet)
	for _, def := range defs {
		if connectedDefinitions.Contains(def.Name()) {
			usedDefinitions.Add(def)
		}
	}

	return usedDefinitions, nil
}
