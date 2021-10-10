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

func StripUnreferencedTypeDefinitions() Stage {
	return MakeLegacyStage(
		StripUnreferencedTypeDefinitionsStageID,
		"Strip unreferenced types",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			resources := astmodel.CollectResourceDefinitions(defs)
			armSpecAndStatus := astmodel.CollectARMSpecAndStatusDefinitions(defs)
			roots := astmodel.SetUnion(resources, armSpecAndStatus)

			return StripUnusedDefinitions(roots, defs)
		})
}

// StripUnusedDefinitions removes all types that aren't in roots or
// referred to by the types in roots, for example types that are
// generated as a byproduct of an allOf element.
func StripUnusedDefinitions(
	roots astmodel.TypeNameSet,
	defs astmodel.Types) (astmodel.Types, error) {
	graph := astmodel.MakeReferenceGraphWithRoots(roots, defs)
	connectedTypes := graph.Connected()

	usedDefinitions := make(astmodel.Types)
	for _, def := range defs {
		if connectedTypes.Contains(def.Name()) {
			usedDefinitions.Add(def)
		}
	}

	return usedDefinitions, nil
}
