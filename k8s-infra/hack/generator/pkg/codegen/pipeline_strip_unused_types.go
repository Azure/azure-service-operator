/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

func stripUnreferencedTypeDefinitions() PipelineStage {
	return MakePipelineStage(
		"stripUnreferenced",
		"Strip unreferenced types",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			roots := astmodel.CollectResourceDefinitions(defs)

			return StripUnusedDefinitions(roots, defs)
		})
}

// StripUnusedDefinitions removes all types that aren't in roots or
// referred to by the types in roots, for example types that are
// generated as a byproduct of an allOf element.
func StripUnusedDefinitions(
	roots astmodel.TypeNameSet,
	definitions astmodel.Types) (astmodel.Types, error) {

	// Collect all the reference sets for each type.
	references := make(map[astmodel.TypeName]astmodel.TypeNameSet)
	for _, def := range definitions {
		references[def.Name()] = def.References()
	}

	graph := astmodel.NewReferenceGraph(roots, references)
	connectedTypes := graph.Connected()

	usedDefinitions := make(astmodel.Types)
	for _, def := range definitions {
		if connectedTypes.Contains(def.Name()) {
			usedDefinitions.Add(def)
		}
	}

	return usedDefinitions, nil
}
