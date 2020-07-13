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
	return PipelineStage{
		"Strip unreferenced types",
		func(ctx context.Context, defs map[astmodel.TypeName]astmodel.TypeDefiner) (map[astmodel.TypeName]astmodel.TypeDefiner, error) {
			roots := collectResourceDefinitions(defs)
			return StripUnusedDefinitions(roots, defs)
		},
	}
}

// stripUnusedDefinitions removes all types that aren't in roots or
// referred to by the types in roots, for example types that are
// generated as a byproduct of an allOf element.
func StripUnusedDefinitions(
	roots astmodel.TypeNameSet,
	definitions map[astmodel.TypeName]astmodel.TypeDefiner,
) (map[astmodel.TypeName]astmodel.TypeDefiner, error) {
	// Collect all the reference sets for each type.
	references := make(map[astmodel.TypeName]astmodel.TypeNameSet)

	for _, def := range definitions {
		references[*def.Name()] = def.Type().References()
	}

	graph := newReferenceGraph(roots, references)
	connectedTypes := graph.connected()
	usedDefinitions := make(map[astmodel.TypeName]astmodel.TypeDefiner)
	for _, def := range definitions {
		if connectedTypes.Contains(*def.Name()) {
			usedDefinitions[*def.Name()] = def
		}
	}

	return usedDefinitions, nil
}

// collectResourceDefinitions returns a TypeNameSet of all of the
// resource definitions in the definitions passed in.
func collectResourceDefinitions(definitions map[astmodel.TypeName]astmodel.TypeDefiner) astmodel.TypeNameSet {
	resources := make(astmodel.TypeNameSet)
	for _, def := range definitions {
		if _, ok := def.(*astmodel.ResourceDefinition); ok {
			resources.Add(*def.Name())
		}
	}
	return resources
}

func newReferenceGraph(roots astmodel.TypeNameSet, references map[astmodel.TypeName]astmodel.TypeNameSet) *referenceGraph {
	return &referenceGraph{
		roots:      roots,
		references: references,
	}
}

type referenceGraph struct {
	roots      astmodel.TypeNameSet
	references map[astmodel.TypeName]astmodel.TypeNameSet
}

// connected returns the set of types that are reachable from the
// roots.
func (c referenceGraph) connected() astmodel.TypeNameSet {
	// Make a non-nil set so we don't need to worry about passing it back down.
	connectedTypes := make(astmodel.TypeNameSet)
	for node := range c.roots {
		c.collectTypes(node, connectedTypes)
	}
	return connectedTypes
}

func (c referenceGraph) collectTypes(node astmodel.TypeName, collected astmodel.TypeNameSet) {
	if collected.Contains(node) {
		// We can stop here - we've already visited this node.
		return
	}
	collected.Add(node)
	for child := range c.references[node] {
		c.collectTypes(child, collected)
	}
}
