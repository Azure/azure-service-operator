/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// TypeNameSet stores type names in no particular order without
// duplicates.
type TypeNameSet map[TypeName]struct{}

// NewTypeNameSet makes a TypeNameSet containing the specified
// names. If no elements are passed it might be nil.
func NewTypeNameSet(initial ...TypeName) TypeNameSet {
	var result TypeNameSet
	for _, name := range initial {
		result = result.Add(name)
	}
	return result
}

// Add includes the passed name in the set and returns the updated
// set, so that adding can work for a nil set - this makes it more
// convenient to add to sets kept in a map (in the way you might with
// a map of slices).
func (ts TypeNameSet) Add(val TypeName) TypeNameSet {
	if ts == nil {
		ts = make(TypeNameSet)
	}
	ts[val] = struct{}{}
	return ts
}

// Contains returns whether this name is in the set. Works for nil
// sets too.
func (ts TypeNameSet) Contains(val TypeName) bool {
	if ts == nil {
		return false
	}
	_, found := ts[val]
	return found
}

// SetUnion returns a new set with all of the names in s1 or s2.
func SetUnion(s1, s2 TypeNameSet) TypeNameSet {
	var result TypeNameSet
	for val := range s1 {
		result = result.Add(val)
	}
	for val := range s2 {
		result = result.Add(val)
	}
	return result
}

// StripUnusedDefinitions removes all types that aren't in roots or
// referred to by the types in roots, for example types that are
// generated as a byproduct of an allOf element.
func StripUnusedDefinitions(
	roots TypeNameSet,
	definitions []TypeDefiner,
) ([]TypeDefiner, error) {
	// Collect all the reference sets for each type.
	references := make(map[TypeName]TypeNameSet)

	for _, def := range definitions {
		references[*def.Name()] = def.Type().References()
	}

	graph := newReferenceGraph(roots, references)
	connectedTypes := graph.connected()
	var usedDefinitions []TypeDefiner
	for _, def := range definitions {
		if connectedTypes.Contains(*def.Name()) {
			usedDefinitions = append(usedDefinitions, def)
		}
	}
	return usedDefinitions, nil
}

// CollectResourceDefinitions returns a TypeNameSet of all of the
// resource definitions in the definitions passed in.
func CollectResourceDefinitions(definitions []TypeDefiner) TypeNameSet {
	resources := make(TypeNameSet)
	for _, def := range definitions {
		if _, ok := def.(*ResourceDefinition); ok {
			resources.Add(*def.Name())
		}
	}
	return resources
}

func newReferenceGraph(roots TypeNameSet, references map[TypeName]TypeNameSet) *referenceGraph {
	return &referenceGraph{
		roots:      roots,
		references: references,
	}
}

type referenceGraph struct {
	roots      TypeNameSet
	references map[TypeName]TypeNameSet
}

// connected returns the set of types that are reachable from the
// roots.
func (c referenceGraph) connected() TypeNameSet {
	// Make a non-nil set so we don't need to worry about passing it back down.
	connectedTypes := make(TypeNameSet)
	for node := range c.roots {
		c.collectTypes(node, connectedTypes)
	}
	return connectedTypes
}

func (c referenceGraph) collectTypes(node TypeName, collected TypeNameSet) {
	if collected.Contains(node) {
		// We can stop here - we've already visited this node.
		return
	}
	collected.Add(node)
	for child := range c.references[node] {
		c.collectTypes(child, collected)
	}
}
