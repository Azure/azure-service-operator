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

// AddAll adds multiple types to the set, with the same safety check as Add() to panic if a duplicate is included
func (types Types) AddAll(otherTypes []TypeDefinition) {
	for _, t := range otherTypes {
		types.Add(t)
	}
}

// AddTypes adds multiple types to the set, with the same safety check as Add() to panic if a duplicate is included
func (types Types) AddTypes(otherTypes Types) {
	for _, t := range otherTypes {
		types.Add(t)
	}
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

// IsArmType returns true if the passed type is an Arm type or names an Arm type; false otherwise.
func (types Types) IsArmType(aType Type) bool {
	switch t := aType.(type) {
	case *ArmType:
		return true

	case *ResourceType:
		return types.IsArmResource(t)

	case TypeName:
		if def, ok := types[t]; ok {
			return types.IsArmDefinition(&def)
		}
		return false

	default:
		return false
	}
}

// IsArmDefinition returns true if the passed definition is for an Arm type or names an Arm type; false otherwise.
func (types Types) IsArmDefinition(definition *TypeDefinition) bool {
	return types.IsArmType(definition.Type())
}

// IsArmDefinition returns true if the passed resource contains Arm Types or names Arm types; false otherwise.
func (types Types) IsArmResource(resource *ResourceType) bool {
	return types.IsArmType(resource.SpecType()) || types.IsArmType(resource.StatusType())
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
