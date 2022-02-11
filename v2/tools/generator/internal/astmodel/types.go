/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

// Types is a map of TypeName to TypeDefinition, representing a set of types.
type Types map[TypeName]TypeDefinition

// A restricted interface to indicate that the
// consumer won’t modify the contained types.
type ReadonlyTypes interface {
	FullyResolve(t Type) (Type, error)
	Get(t TypeName) TypeDefinition
	TryGet(t TypeName) (TypeDefinition, bool)
}

// MakeTypes makes it easier to declare a Types from a map
func MakeTypes(tys map[TypeName]Type) Types {
	result := make(Types, len(tys))
	for name, ty := range tys {
		result.Add(MakeTypeDefinition(name, ty))
	}

	return result
}

func (types Types) Get(t TypeName) TypeDefinition {
	return types[t]
}

func (types Types) TryGet(t TypeName) (TypeDefinition, bool) {
	result, ok := types[t]
	return result, ok
}

// Add adds a type to the set, with safety check that it has not already been defined
func (types Types) Add(def TypeDefinition) {
	key := def.Name()
	if _, ok := types[key]; ok {
		panic(fmt.Sprintf("type already defined: %s", key))
	}

	types[key] = def
}

// FullyResolve turns something that might be a TypeName into something that isn't
func (types Types) FullyResolve(t Type) (Type, error) {
	tName, ok := t.(TypeName)
	for ok {
		tDef, found := types[tName]
		if !found {
			return nil, errors.Errorf("couldn't find definition for %s", tName)
		}

		t = tDef.Type()
		tName, ok = t.(TypeName)
	}

	return t, nil
}

// AddAll adds multiple definitions to the set, with the same safety check as Add() to panic if a duplicate is included
func (types Types) AddAll(otherDefinitions ...TypeDefinition) {
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

// AddTypesAllowDuplicates adds multiple types to the set.
// Multiple adds of a type with the same shape are allowed, but attempting to add two
// types with the same name but different shape will trigger an error.
func (types Types) AddTypesAllowDuplicates(otherTypes Types) error {
	for _, t := range otherTypes {
		err := types.AddAllowDuplicates(t)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddAllowDuplicates attempts to add the specified definition to the types collection.
// Multiple adds of a type with the same shape are allowed, but attempting to add two
// types with the same name but different shape will trigger an error.
func (types Types) AddAllowDuplicates(def TypeDefinition) error {
	if !types.Contains(def.Name()) {
		types.Add(def)
		return nil
	}

	existing := types[def.Name()]
	if !TypeEquals(def.Type(), existing.Type()) {
		return errors.Errorf("type definition for %q has two shapes: %s", existing.Name(), DiffTypes(existing.Type(), def.Type()))
	}

	// Can safely skip this add
	return nil
}

func DiffTypes(x, y interface{}) string {
	allowAll := cmp.AllowUnexported(
		AllOfType{},
		ObjectType{},
		OneOfType{},
		PropertyDefinition{},
		OptionalType{},
		ArrayType{},
		PrimitiveType{},
		EnumType{},
		TypeName{},
		LocalPackageReference{},
		InterfaceImplementer{},
		TypeSet{},
	)

	return cmp.Diff(x, y, allowAll)
}

// AddAllAllowDuplicates adds multiple definitions to the set.
// Multiple adds of a type with the same shape are allowed, but attempting to add two
// types with the same name but different shape will trigger an error.
func (types Types) AddAllAllowDuplicates(otherDefinitions []TypeDefinition) error {
	for _, def := range otherDefinitions {
		err := types.AddAllowDuplicates(def)
		if err != nil {
			return err
		}
	}

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

// Intersect returns a new set of types including only those defined in both types and otherTypes.
func (types Types) Intersect(otherTypes Types) Types {
	return types.Where(func(def TypeDefinition) bool {
		return otherTypes.Contains(def.Name())
	})
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

// OverlayWith creates a new set containing all the type definitions from both this and the provided set. Any name
// collisions are resolved in favour of the provided set. Returns a new independent set, leaving the original unmodified.
func (types Types) OverlayWith(t Types) Types {
	result := t.Copy()
	result.AddTypes(types.Except(t))
	return result
}

// Names returns the names of all of the types in the set
func (types Types) Names() TypeNameSet {
	result := NewTypeNameSet()
	for name := range types {
		result.Add(name)
	}

	return result
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

// ResolveObjectType returns the underlying resource type if the definition contains one or names one
func (types Types) ResolveObjectType(aType Type) (*ObjectType, bool) {
	switch t := aType.(type) {

	case *ObjectType:
		return t, true

	case TypeName:
		if def, ok := types[t]; ok {
			return types.ResolveObjectType(def.theType)
		}
		return nil, false

	default:
		return nil, false
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
	specType := resourceType.SpecType()
	specName, ok := specType.(TypeName)
	if !ok {
		return TypeDefinition{}, errors.Errorf("spec was not of type TypeName, instead: %T (%s)", specType, specType)
	}

	resourceSpecDef, ok := types[specName]
	if !ok {
		return TypeDefinition{}, errors.Errorf("couldn't find spec %s", specName)
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
		return TypeDefinition{}, errors.Errorf("couldn't find status %s", statusName)
	}

	// preserve outer spec name
	return resourceStatusDef.WithName(statusName), nil
}

type ResolvedResourceDefinition struct {
	ResourceDef  TypeDefinition
	ResourceType *ResourceType

	SpecDef  TypeDefinition
	SpecType *ObjectType

	StatusDef  TypeDefinition
	StatusType *ObjectType
}

// ResolveResourceSpecAndStatus takes a TypeDefinition that is a ResourceType and looks up its Spec and Status (as well as
// the TypeDefinition's corresponding to them) and returns a ResolvedResourceDefinition
func (types Types) ResolveResourceSpecAndStatus(resourceDef TypeDefinition) (*ResolvedResourceDefinition, error) {
	resource, ok := AsResourceType(resourceDef.Type())
	if !ok {
		return nil, errors.Errorf("expected %q to be a Resource but instead it was a %T", resourceDef.Name(), resourceDef.Type())
	}

	// Resolve the spec
	specDef, err := types.ResolveResourceSpecDefinition(resource)
	if err != nil {
		return nil, err
	}
	spec, ok := AsObjectType(specDef.Type())
	if !ok {
		return nil, errors.Errorf("resource spec %q did not contain an object, instead %s", resource.SpecType().String(), specDef.Type())
	}

	// Resolve the status if it's there (we need this because our golden file tests don't have status currently)
	var statusDef TypeDefinition
	var status *ObjectType

	if IgnoringErrors(resource.StatusType()) != nil {
		statusDef, err = types.ResolveResourceStatusDefinition(resource)
		if err != nil {
			return nil, err
		}
		status, ok = AsObjectType(statusDef.Type())
		if !ok {
			return nil, errors.Errorf("resource status %q did not contain an object", resource.StatusType().String())
		}
	}

	return &ResolvedResourceDefinition{
		ResourceDef:  resourceDef,
		ResourceType: resource,
		SpecDef:      specDef,
		SpecType:     spec,
		StatusDef:    statusDef,
		StatusType:   status,
	}, nil
}

// Process applies a func to transform all members of this set of type definitions, returning a new set of type
// definitions containing the results of the transformation, or possibly an error
// Only definitions returned by the func will be included in the results of the function. The func may return a nil
// TypeDefinition if it doesn't want to include anything in the output set.
func (types Types) Process(transformation func(definition TypeDefinition) (*TypeDefinition, error)) (Types, error) {
	result := make(Types)

	for _, def := range types {
		d, err := transformation(def)
		if err != nil {
			return nil, err
		} else if d != nil {
			result.Add(*d)
		}
	}

	return result, nil
}

// FindResourceTypes walks the provided set of TypeDefinitions and returns all the resource types
func FindResourceTypes(types Types) Types {
	result := make(Types)

	// Find all our resources and extract all their Specs
	for _, def := range types {
		_, ok := AsResourceType(def.Type())
		if !ok {
			continue
		}

		// We have a resource type
		result.Add(def)
	}

	return result
}

// FindSpecTypes walks the provided set of TypeDefinitions and returns all the spec types
func FindSpecTypes(types Types) Types {
	result := make(Types)

	// Find all our resources and extract all their Specs
	for _, def := range types {
		rt, ok := AsResourceType(def.Type())
		if !ok {
			continue
		}

		// We have a resource type
		tn, ok := AsTypeName(rt.SpecType())
		if !ok {
			continue
		}

		// Add the named spec type to our results
		if spec, ok := types.TryGet(tn); ok {
			// Use AddAllowDuplicates here because some resources share the same spec
			// across multiple resources, which can trigger multiple adds of the same type
			err := result.AddAllowDuplicates(spec)
			if err != nil {
				panic(err)
			}
		}
	}

	return result
}

// FindStatusTypes walks the provided set of TypeDefinitions and returns all the status types
func FindStatusTypes(types Types) Types {
	result := make(Types)

	// Find all our resources and extract all their Statuses
	for _, def := range types {
		rt, ok := AsResourceType(def.Type())
		if !ok {
			continue
		}

		// We have a resource type
		tn, ok := AsTypeName(rt.StatusType())
		if !ok {
			continue
		}

		// Add the named status type to our results
		if status, ok := types.TryGet(tn); ok {
			// Use AddAllowDuplicates here because some resources share the same status
			// across multiple resources, which can trigger multiple adds of the same type
			err := result.AddAllowDuplicates(status)
			if err != nil {
				panic(err)
			}
		}
	}

	return result
}

// FindConnectedTypes finds all types reachable from the provided types
// TODO: This is very similar to ReferenceGraph.Connected.
func FindConnectedTypes(allTypes Types, roots Types) (Types, error) {
	walker := NewTypeWalker(
		allTypes,
		TypeVisitorBuilder{}.Build())

	result := make(Types)
	for _, def := range roots {
		types, err := walker.Walk(def)
		if err != nil {
			return nil, errors.Wrapf(err, "failed walking types")
		}

		// TODO: magic APIVersion handling
		for _, t := range types {
			if rt, ok := AsResourceType(t.Type()); ok {
				err = result.AddAllowDuplicates(allTypes[rt.APIVersionTypeName()])
				if err != nil {
					panic(err)
				}
			}
		}

		err = result.AddTypesAllowDuplicates(types)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// FindSpecConnectedTypes finds all spec types and all types referenced by those spec types.
// This differs from FindSpecTypes in that it finds not only the top level spec types but
// also the types which the top level types are built out of.
func FindSpecConnectedTypes(types Types) (Types, error) {
	specTypes := FindSpecTypes(types)
	return FindConnectedTypes(types, specTypes)
}

// FindStatusConnectedTypes finds all status types and all types referenced by those spec types.
// This differs from FindStatusTypes in that it finds not only the top level spec types but
// also the types which the top level types are built out of.
func FindStatusConnectedTypes(types Types) (Types, error) {
	statusTypes := FindStatusTypes(types)
	return FindConnectedTypes(types, statusTypes)
}
