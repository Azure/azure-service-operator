/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/readonly"
)

// TypeDefinitionSet is a map of TypeName to TypeDefinition, representing a set of type definitions.
type TypeDefinitionSet map[TypeName]TypeDefinition

// A restricted interface to indicate that the
// consumer won’t modify the contained types.
type ReadonlyTypeDefinitions interface {
	MustGetDefinition(name TypeName) TypeDefinition
	GetDefinition(name TypeName) (TypeDefinition, error)
}

var _ ReadonlyTypeDefinitions = TypeDefinitionSet{}

// MakeTypeDefinitionSet makes it easier to declare a TypeDefinitionSet from a map
func MakeTypeDefinitionSet(types map[TypeName]Type) TypeDefinitionSet {
	result := make(TypeDefinitionSet, len(types))
	for name, ty := range types {
		result.Add(MakeTypeDefinition(name, ty))
	}

	return result
}

// MustGetDefinition looks up a type definition and panics if it cannot be found
func (set TypeDefinitionSet) MustGetDefinition(name TypeName) TypeDefinition {
	result, ok := set[name]
	if !ok {
		panic(fmt.Sprintf("couldn't find type %q", name))
	}

	return result
}

// GetDefinition attempts to look up a type definition based on the name. An error is
// returned if it cannot be found
func (set TypeDefinitionSet) GetDefinition(name TypeName) (TypeDefinition, error) {
	result, ok := set[name]
	if !ok {
		return TypeDefinition{}, errors.Errorf("couldn't find type %q", name)
	}

	return result, nil
}

// Add adds a type to the set, with safety check that it has not already been defined
func (set TypeDefinitionSet) Add(def TypeDefinition) {
	key := def.Name()
	if _, ok := set[key]; ok {
		panic(fmt.Sprintf("type already defined: %s", key))
	}

	set[key] = def
}

// FullyResolve turns something that might be a TypeName into something that isn't
func (set TypeDefinitionSet) FullyResolve(t Type) (Type, error) {
	tName, ok := t.(TypeName)
	for ok {
		tDef, found := set[tName]
		if !found {
			return nil, errors.Errorf("couldn't find definition for %s", tName)
		}

		t = tDef.Type()
		tName, ok = t.(TypeName)
	}

	return t, nil
}

// AddAll adds multiple definitions to the set, with the same safety check as Add() to panic if a duplicate is included
func (set TypeDefinitionSet) AddAll(otherDefinitions ...TypeDefinition) {
	for _, t := range otherDefinitions {
		set.Add(t)
	}
}

// AddTypes adds multiple types to the set, with the same safety check as Add() to panic if a duplicate is included
func (set TypeDefinitionSet) AddTypes(definitions TypeDefinitionSet) {
	for _, t := range definitions {
		set.Add(t)
	}
}

// AddTypesAllowDuplicates adds multiple types to the set.
// Multiple adds of a type with the same shape are allowed, but attempting to add two
// types with the same name but different shape will trigger an error.
func (set TypeDefinitionSet) AddTypesAllowDuplicates(definitions TypeDefinitionSet) error {
	for _, t := range definitions {
		err := set.AddAllowDuplicates(t)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddAllowDuplicates attempts to add the specified definition to the types collection.
// Multiple adds of a type with the same shape are allowed, but attempting to add two
// types with the same name but different shape will trigger an error.
func (set TypeDefinitionSet) AddAllowDuplicates(def TypeDefinition) error {
	if !set.Contains(def.Name()) {
		set.Add(def)
		return nil
	}

	existing := set[def.Name()]
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
		readonly.Map[string, Function]{},
		readonly.Map[string, TestCase]{},
		readonly.Map[string, []string]{},
	)

	return cmp.Diff(x, y, allowAll)
}

// AddAllAllowDuplicates adds multiple definitions to the set.
// Multiple adds of a type with the same shape are allowed, but attempting to add two
// types with the same name but different shape will trigger an error.
func (set TypeDefinitionSet) AddAllAllowDuplicates(otherDefinitions []TypeDefinition) error {
	for _, def := range otherDefinitions {
		err := set.AddAllowDuplicates(def)
		if err != nil {
			return err
		}
	}

	return nil
}

// Where returns a new set of types including only those that satisfy the predicate
func (set TypeDefinitionSet) Where(predicate func(definition TypeDefinition) bool) TypeDefinitionSet {
	result := make(TypeDefinitionSet)
	for _, t := range set {
		if predicate(t) {
			result[t.Name()] = t
		}
	}

	return result
}

// Intersect returns a new set of types including only those defined in both types and otherTypes.
func (set TypeDefinitionSet) Intersect(definitions TypeDefinitionSet) TypeDefinitionSet {
	return set.Where(func(def TypeDefinition) bool {
		return definitions.Contains(def.Name())
	})
}

// Except returns a new set of types including only those not defined in otherTypes
func (set TypeDefinitionSet) Except(definitions TypeDefinitionSet) TypeDefinitionSet {
	return set.Where(func(def TypeDefinition) bool {
		return !definitions.Contains(def.Name())
	})
}

// Contains returns true if the set contains a definition for the specified name
func (set TypeDefinitionSet) Contains(name TypeName) bool {
	_, ok := set[name]
	return ok
}

// OverlayWith creates a new set containing all the type definitions from both this and the provided set. Any name
// collisions are resolved in favour of the provided set. Returns a new independent set, leaving the original unmodified.
func (set TypeDefinitionSet) OverlayWith(t TypeDefinitionSet) TypeDefinitionSet {
	result := t.Copy()
	result.AddTypes(set.Except(t))
	return result
}

// Names returns the names of all of the types in the set
func (set TypeDefinitionSet) Names() TypeNameSet {
	result := NewTypeNameSet()
	for name := range set {
		result.Add(name)
	}

	return result
}

// TypesDisjointUnion merges this and other, with a safety check that no type is overwritten.
// If an attempt is made to overwrite a type, this function panics
func TypesDisjointUnion(s1 TypeDefinitionSet, s2 TypeDefinitionSet) TypeDefinitionSet {
	result := s1.Copy()
	result.AddTypes(s2)
	return result
}

// Copy makes an independent copy of this set of types
func (set TypeDefinitionSet) Copy() TypeDefinitionSet {
	result := make(TypeDefinitionSet)
	result.AddTypes(set)
	return result
}

// ResolveResourceType returns the underlying resource type if the definition contains one or names one
func (set TypeDefinitionSet) ResolveResourceType(aType Type) (*ResourceType, bool) {
	switch t := aType.(type) {

	case *ResourceType:
		return t, true

	case TypeName:
		if def, ok := set[t]; ok {
			return set.ResolveResourceType(def.theType)
		}
		return nil, false

	default:
		return nil, false
	}
}

// ResolveEnumType returns true if the passed type is an enum type or names an enum type; false otherwise.
func (set TypeDefinitionSet) ResolveEnumType(aType Type) (EnumType, bool) {
	switch t := aType.(type) {
	case *EnumType:
		return *t, true

	case TypeName:
		if def, ok := set[t]; ok {
			return set.ResolveEnumDefinition(&def)
		}
		return EnumType{}, false

	default:
		return EnumType{}, false
	}
}

// ResolveObjectType returns the underlying resource type if the definition contains one or names one
func (set TypeDefinitionSet) ResolveObjectType(aType Type) (*ObjectType, bool) {
	switch t := aType.(type) {

	case *ObjectType:
		return t, true

	case TypeName:
		if def, ok := set[t]; ok {
			return set.ResolveObjectType(def.theType)
		}
		return nil, false

	default:
		return nil, false
	}
}

// ResolveEnumDefinition returns true if the passed definition is for an Enum type or names an Enum type; false otherwise.
func (set TypeDefinitionSet) ResolveEnumDefinition(definition *TypeDefinition) (EnumType, bool) {
	return set.ResolveEnumType(definition.Type())
}

// ResolveResourceSpecDefinition finds the TypeDefinition associated with the resource Spec.
func (set TypeDefinitionSet) ResolveResourceSpecDefinition(resourceType *ResourceType) (TypeDefinition, error) {
	return ResolveResourceSpecDefinition(set, resourceType)
}

// ResolveResourceStatusDefinition finds the TypeDefinition associated with the resource Status.
func (set TypeDefinitionSet) ResolveResourceStatusDefinition(resourceType *ResourceType) (TypeDefinition, error) {
	return ResolveResourceStatusDefinition(set, resourceType)
}

// AsSlice creates a new slice containing all the definitions
func (set TypeDefinitionSet) AsSlice() []TypeDefinition {
	result := make([]TypeDefinition, 0, len(set))
	for _, def := range set {
		result = append(result, def)
	}

	return result
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
func (set TypeDefinitionSet) ResolveResourceSpecAndStatus(resourceDef TypeDefinition) (*ResolvedResourceDefinition, error) {
	return ResolveResourceSpecAndStatus(set, resourceDef)
}

// Process applies a func to transform all members of this set of type definitions, returning a new set of type
// definitions containing the results of the transformation, or possibly an error
// Only definitions returned by the func will be included in the results of the function. The func may return a nil
// TypeDefinition if it doesn't want to include anything in the output set.
func (set TypeDefinitionSet) Process(transformation func(definition TypeDefinition) (*TypeDefinition, error)) (TypeDefinitionSet, error) {
	result := make(TypeDefinitionSet)

	for _, def := range set {
		d, err := transformation(def)
		if err != nil {
			return nil, err
		} else if d != nil {
			result.Add(*d)
		}
	}

	return result, nil
}

// ResolveResourceSpecDefinition finds the TypeDefinition associated with the resource Spec.
func ResolveResourceSpecDefinition(defs ReadonlyTypeDefinitions, resourceType *ResourceType) (TypeDefinition, error) {
	// The expectation is that the spec type is just a name
	specName, ok := resourceType.SpecType().(TypeName)
	if !ok {
		return TypeDefinition{}, errors.Errorf("spec was not of type TypeName, instead: %T", resourceType.SpecType())
	}

	resourceSpecDef, err := defs.GetDefinition(specName)
	if !ok {
		return TypeDefinition{}, errors.Wrapf(err, "couldn't find spec")
	}

	return resourceSpecDef, nil
}

// ResolveResourceStatusDefinition finds the TypeDefinition associated with the resource Status.
func ResolveResourceStatusDefinition(defs ReadonlyTypeDefinitions, resourceType *ResourceType) (TypeDefinition, error) {
	statusName, ok := resourceType.StatusType().(TypeName)
	if !ok {
		return TypeDefinition{}, errors.Errorf("status was not of type TypeName, instead: %T", resourceType.StatusType())
	}

	resourceStatusDef, err := defs.GetDefinition(statusName)
	if !ok {
		return TypeDefinition{}, errors.Wrapf(err, "couldn't find status")
	}

	// preserve outer spec name
	return resourceStatusDef.WithName(statusName), nil
}

// ResolveResourceSpecAndStatus takes a TypeDefinition that is a ResourceType and looks up its Spec and Status (as well as
// the TypeDefinition's corresponding to them) and returns a ResolvedResourceDefinition
func ResolveResourceSpecAndStatus(defs ReadonlyTypeDefinitions, resourceDef TypeDefinition) (*ResolvedResourceDefinition, error) {
	resource, ok := AsResourceType(resourceDef.Type())
	if !ok {
		return nil, errors.Errorf("expected %q to be a Resource but instead it was a %T", resourceDef.Name(), resourceDef.Type())
	}

	// Resolve the spec
	specDef, err := ResolveResourceSpecDefinition(defs, resource)
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
		statusDef, err = ResolveResourceStatusDefinition(defs, resource)
		if err != nil {
			return nil, err
		}
		status, ok = AsObjectType(statusDef.Type())
		if !ok {
			return nil, errors.Errorf("resource status %q did not contain an object, instead %s", resource.StatusType().String(), statusDef.Type())
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

// FindResourceDefinitions walks the provided set of TypeDefinitions and returns all the resource definitions
func FindResourceDefinitions(definitions TypeDefinitionSet) TypeDefinitionSet {
	result := make(TypeDefinitionSet)

	// Find all our resources and extract all their Specs
	for _, def := range definitions {
		_, ok := AsResourceType(def.Type())
		if !ok {
			continue
		}

		// We have a resource type
		result.Add(def)
	}

	return result
}

// FindSpecDefinitions walks the provided set of TypeDefinitions and returns all the spec definitions
func FindSpecDefinitions(definitions TypeDefinitionSet) TypeDefinitionSet {
	result := make(TypeDefinitionSet)

	// Find all our resources and extract all their Specs
	for _, def := range definitions {
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
		if spec, err := definitions.GetDefinition(tn); err == nil {
			// Use AddAllowDuplicates here because some resources share the same spec
			// across multiple resources, which can trigger multiple adds of the same type
			err = result.AddAllowDuplicates(spec)
			if err != nil {
				panic(err)
			}
		}
	}

	return result
}

// FindStatusDefinitions walks the provided set of TypeDefinitions and returns all the status definitions
func FindStatusDefinitions(definitions TypeDefinitionSet) TypeDefinitionSet {
	result := make(TypeDefinitionSet)

	// Find all our resources and extract all their Statuses
	for _, def := range definitions {
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
		if status, err := definitions.GetDefinition(tn); err == nil {
			// Use AddAllowDuplicates here because some resources share the same status
			// across multiple resources, which can trigger multiple adds of the same type
			err = result.AddAllowDuplicates(status)
			if err != nil {
				panic(err)
			}
		}
	}

	return result
}

// FindConnectedDefinitions finds all types reachable from the provided definitions
// TODO: This is very similar to ReferenceGraph.Connected.
func FindConnectedDefinitions(definitions TypeDefinitionSet, roots TypeDefinitionSet) (TypeDefinitionSet, error) {
	walker := NewTypeWalker(
		definitions,
		TypeVisitorBuilder{}.Build())

	result := make(TypeDefinitionSet)
	for _, def := range roots {
		types, err := walker.Walk(def)
		if err != nil {
			return nil, errors.Wrapf(err, "failed walking types")
		}

		err = result.AddTypesAllowDuplicates(types)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// FindSpecConnectedDefinitions finds all spec definitions and all types referenced by those spec definitions.
// This differs from FindSpecDefinitions in that it finds not only the top level spec definitions but
// also the types which the top level types are built out of.
func FindSpecConnectedDefinitions(definitions TypeDefinitionSet) (TypeDefinitionSet, error) {
	specDefs := FindSpecDefinitions(definitions)
	return FindConnectedDefinitions(definitions, specDefs)
}

// FindStatusConnectedDefinitions finds all status definitions and all types referenced by those spec definitions.
// This differs from FindStatusDefinitions in that it finds not only the top level status definitions but
// also the types which the top level types are built out of.
func FindStatusConnectedDefinitions(definitions TypeDefinitionSet) (TypeDefinitionSet, error) {
	statusDefs := FindStatusDefinitions(definitions)
	return FindConnectedDefinitions(definitions, statusDefs)
}
