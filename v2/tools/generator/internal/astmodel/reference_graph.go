/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/pkg/errors"
)

// ReferenceGraph is a graph of references between types
type ReferenceGraph struct {
	roots      TypeNameSet
	references map[TypeName]TypeNameSet
}

// CollectARMSpecAndStatusDefinitions returns a TypeNameSet of all of the ARM spec definitions
// passed in.
func CollectARMSpecAndStatusDefinitions(definitions TypeDefinitionSet) TypeNameSet {
	findARMType := func(t Type) (TypeName, error) {
		name, ok := t.(TypeName)
		if !ok {
			return nil, errors.Errorf("type was not of type TypeName, instead %T", t)
		}

		armName := CreateARMTypeName(name)

		if _, ok = definitions[armName]; !ok {
			return nil, errors.Errorf("couldn't find ARM type %q", armName)
		}

		return armName, nil
	}

	armSpecAndStatus := NewTypeNameSet()
	resources := FindResourceDefinitions(definitions)
	for _, def := range resources {
		resourceType, ok := AsResourceType(def.Type())
		if !ok {
			panic(fmt.Sprintf("FindResourceDefinitions() returned non-resource type %T", def.Type()))
		}

		armSpecName, err := findARMType(resourceType.spec)
		if err != nil {
			continue // No ARM type here - skip
		}
		armSpecAndStatus.Add(armSpecName)

		statusType := IgnoringErrors(resourceType.status)
		if statusType != nil {
			armStatusName, err := findARMType(statusType)
			if err != nil {
				continue // No ARM type here - skip
			}
			armSpecAndStatus.Add(armStatusName)
		}
	}
	return armSpecAndStatus
}

// MakeReferenceGraph produces a new ReferenceGraph with the given roots and references
func MakeReferenceGraph(roots TypeNameSet, references map[TypeName]TypeNameSet) ReferenceGraph {
	return ReferenceGraph{
		roots:      roots,
		references: references,
	}
}

// MakeReferenceGraphWithRoots produces a ReferenceGraph with the given roots, and references
// derived from the provided types collection.
func MakeReferenceGraphWithRoots(roots TypeNameSet, definitions TypeDefinitionSet) ReferenceGraph {
	references := make(map[TypeName]TypeNameSet, len(definitions))
	for _, def := range definitions {
		references[def.Name()] = def.References()
	}

	return MakeReferenceGraph(roots, references)
}

// MakeReferenceGraphWithResourcesAsRoots produces a ReferenceGraph for the given set of
// definitions, where the Resource types (and their ARM spec/status) are the roots.
func MakeReferenceGraphWithResourcesAsRoots(definitions TypeDefinitionSet) ReferenceGraph {
	resources := FindResourceDefinitions(definitions)
	armSpecAndStatus := CollectARMSpecAndStatusDefinitions(definitions)
	roots := SetUnion(resources.Names(), armSpecAndStatus)

	return MakeReferenceGraphWithRoots(roots, definitions)
}

type ReachableTypes map[TypeName]int // int = depth

func (reachable ReachableTypes) Contains(tn TypeName) bool {
	_, ok := reachable[tn]
	return ok
}

// Connected returns the set of types that are reachable from the roots.
func (c ReferenceGraph) Connected() ReachableTypes {
	// Make a non-nil set so we don't need to worry about passing it back down.
	connectedTypes := make(ReachableTypes)
	for node := range c.roots {
		c.collectTypes(0, node, connectedTypes)
	}

	return connectedTypes
}

// collectTypes returns the set of types that are reachable from the given
// 'node' (and places them in 'collected').
func (c ReferenceGraph) collectTypes(depth int, node TypeName, collected ReachableTypes) {
	if currentDepth, ok := collected[node]; ok {
		// We can stop here - we've already visited this node.
		// But first, see if we are at a lower depth:
		if depth >= currentDepth {
			return
		}

		// if the depth is lower than what we already had
		// we must continue to reassign depths
	}

	collected[node] = depth
	for child := range c.references[node] {
		c.collectTypes(depth+1, child, collected)
	}
}
