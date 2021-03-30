/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/pkg/errors"
)

// Types is a map of TypeName to TypeDefinition, representing a set of types.
type Types map[TypeName]TypeDefinition

// Add adds a type to the set, with safety check that it has not already been defined
func (types Types) Add(def TypeDefinition) {
	key := def.Name()
	if _, ok := types[key]; ok {
		panic(fmt.Sprintf("type already defined: %v", key))
	}

	types[key] = def
}

// FullyResolve turns something that might be a TypeName into something that isn't
func (types Types) FullyResolve(t Type) (Type, error) {
	tName, ok := t.(TypeName)
	for ok {
		tDef, found := types[tName]
		if !found {
			return nil, errors.Errorf("couldn't find definition for %v", tName)
		}

		t = tDef.Type()
		tName, ok = t.(TypeName)
	}

	return t, nil
}

// AddAll adds multiple definitions to the set, with the same safety check as Add() to panic if a duplicate is included
func (types Types) AddAll(otherDefinitions []TypeDefinition) {
	for _, t := range otherDefinitions {
		types.Add(t)
	}
}

// AddTypes adds multiple types to the set, with the same safety check as Add() to panic if a duplicate is included
func (types Types) AddTypes(otherTypes Types) {
	for _, t := range otherTypes {
		types.Add(t)
	}
}

// AddWithEqualityCheck attempts to add the specified definition to the types collection.
// Multiple adds of a type with the same shape are allowed, but attempting to add two
// types with the same name but different shape will trigger an error.
func (types Types) AddWithEqualityCheck(def TypeDefinition) error {
	if !types.Contains(def.Name()) {
		types.Add(def)
		return nil
	}

	existing := types[def.Name()]
	if !def.Type().Equals(existing.Type()) {
		return errors.Errorf("type definition for %q has two shapes", existing.Name())
	}

	// Can safely skip this add
	return nil
}

// Where returns a new set of types including only those that satisfy the predicate
func (types Types) Where(predicate func(definition TypeDefinition) bool) Types {
	result := make(Types)
	for _, t := range types {
		if predicate(t) {
			result[t.Name()] = t
		}
	}

	return result
}

// Except returns a new set of types including only those not defined in otherTypes
func (types Types) Except(otherTypes Types) Types {
	return types.Where(func(def TypeDefinition) bool {
		return !otherTypes.Contains(def.Name())
	})
}

// Contains returns true if the set contains a definition for the specified name
func (types Types) Contains(name TypeName) bool {
	_, ok := types[name]
	return ok
}

// TypesDisjointUnion merges this and other, with a safety check that no type is overwritten.
// If an attempt is made to overwrite a type, this function panics
func TypesDisjointUnion(s1 Types, s2 Types) Types {
	result := s1.Copy()
	result.AddTypes(s2)
	return result
}

// Copy makes an independent copy of this set of types
func (types Types) Copy() Types {
	result := make(Types)
	result.AddTypes(types)
	return result
}

// ResolveResourceType returns the underlying resource type if the definition contains one or names one
func (types Types) ResolveResourceType(aType Type) (*ResourceType, bool) {
	switch t := aType.(type) {

	case *ResourceType:
		return t, true

	case TypeName:
		if def, ok := types[t]; ok {
			return types.ResolveResourceType(def.theType)
		}
		return nil, false

	default:
		return nil, false
	}
}

// ResolveEnumType returns true if the passed type is an enum type or names an enum type; false otherwise.
func (types Types) ResolveEnumType(aType Type) (EnumType, bool) {
	switch t := aType.(type) {
	case *EnumType:
		return *t, true

	case TypeName:
		if def, ok := types[t]; ok {
			return types.ResolveEnumDefinition(&def)
		}
		return EnumType{}, false

	default:
		return EnumType{}, false
	}
}

// ResolveEnumDefinition returns true if the passed definition is for an Enum type or names an Enum type; false otherwise.
func (types Types) ResolveEnumDefinition(definition *TypeDefinition) (EnumType, bool) {
	return types.ResolveEnumType(definition.Type())
}

// ResolveResourceSpecDefinition finds the TypeDefinition associated with the resource Spec.
func (types Types) ResolveResourceSpecDefinition(
	resourceType *ResourceType) (TypeDefinition, error) {

	// The expectation is that the spec type is just a name
	specName, ok := resourceType.SpecType().(TypeName)
	if !ok {
		return TypeDefinition{}, errors.Errorf("spec was not of type TypeName, instead: %T", resourceType.SpecType())
	}

	resourceSpecDef, ok := types[specName]
	if !ok {
		return TypeDefinition{}, errors.Errorf("couldn't find spec %v", specName)
	}

	return resourceSpecDef, nil
}

func (types Types) ResolveResourceStatusDefinition(
	resourceType *ResourceType) (TypeDefinition, error) {

	statusName, ok := resourceType.StatusType().(TypeName)
	if !ok {
		return TypeDefinition{}, errors.Errorf("status was not of type TypeName, instead: %T", resourceType.StatusType())
	}

	resourceStatusDef, ok := types[statusName]
	if !ok {
		return TypeDefinition{}, errors.Errorf("couldn't find status %v", statusName)
	}

	// preserve outer spec name
	return resourceStatusDef.WithName(statusName), nil
}
