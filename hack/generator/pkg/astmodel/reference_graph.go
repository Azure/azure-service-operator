/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// ReferenceGraph is a graph of references between types
type ReferenceGraph struct {
	roots      TypeNameSet
	references map[TypeName]TypeNameSet
}

// CollectResourceDefinitions returns a TypeNameSet of all of the
// resource definitions in the definitions passed in.
func CollectResourceDefinitions(definitions Types) TypeNameSet {
	resources := make(TypeNameSet)
	for _, def := range definitions {
		if _, ok := def.Type().(*ResourceType); ok {
			resources.Add(def.Name())
		}
	}
	return resources
}

// NewReferenceGraph produces a new ReferenceGraph with the given roots and references
func NewReferenceGraph(roots TypeNameSet, references map[TypeName]TypeNameSet) ReferenceGraph {
	return ReferenceGraph{
		roots:      roots,
		references: references,
	}
}

// NewReferenceGraphWithResourcesAsRoots produces a ReferenceGraph for the given set of
// types, where the Resource types are the roots.
func NewReferenceGraphWithResourcesAsRoots(types Types) ReferenceGraph {
	roots := CollectResourceDefinitions(types)

	references := make(map[TypeName]TypeNameSet)
	for _, def := range types {
		references[def.Name()] = def.References()
	}

	return NewReferenceGraph(roots, references)
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
