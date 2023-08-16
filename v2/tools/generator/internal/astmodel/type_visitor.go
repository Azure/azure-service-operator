/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/pkg/errors"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// TypeVisitor represents a visitor for a tree of types.
// The `ctx` argument can be used to “smuggle” additional data down the call-chain.
type TypeVisitor[C any] struct {
	visitTypeNameIsIdentity bool // performance optimization to avoid reboxing TypeNames constantly
	visitTypeName           func(this *TypeVisitor[C], it TypeName, ctx C) (Type, error)

	visitOneOfType     func(this *TypeVisitor[C], it *OneOfType, ctx C) (Type, error)
	visitAllOfType     func(this *TypeVisitor[C], it *AllOfType, ctx C) (Type, error)
	visitArrayType     func(this *TypeVisitor[C], it *ArrayType, ctx C) (Type, error)
	visitPrimitive     func(this *TypeVisitor[C], it *PrimitiveType, ctx C) (Type, error)
	visitObjectType    func(this *TypeVisitor[C], it *ObjectType, ctx C) (Type, error)
	visitMapType       func(this *TypeVisitor[C], it *MapType, ctx C) (Type, error)
	visitOptionalType  func(this *TypeVisitor[C], it *OptionalType, ctx C) (Type, error)
	visitEnumType      func(this *TypeVisitor[C], it *EnumType, ctx C) (Type, error)
	visitResourceType  func(this *TypeVisitor[C], it *ResourceType, ctx C) (Type, error)
	visitFlaggedType   func(this *TypeVisitor[C], it *FlaggedType, ctx C) (Type, error)
	visitValidatedType func(this *TypeVisitor[C], it *ValidatedType, ctx C) (Type, error)
	visitErroredType   func(this *TypeVisitor[C], it *ErroredType, ctx C) (Type, error)
	visitInterfaceType func(this *TypeVisitor[C], it *InterfaceType, ctx C) (Type, error)
}

// Visit invokes the appropriate VisitX on TypeVisitor
func (tv *TypeVisitor[C]) Visit(t Type, ctx C) (Type, error) {
	if t == nil {
		return nil, nil
	}

	switch it := t.(type) {
	case TypeName:
		// IdentityVisitOfTypeName will re-box the TypeName
		// avoid this allocation if possible by short-cutting
		if tv.visitTypeNameIsIdentity {
			return t, nil
		}
		return tv.visitTypeName(tv, it, ctx)
	case *OneOfType:
		return tv.visitOneOfType(tv, it, ctx)
	case *AllOfType:
		return tv.visitAllOfType(tv, it, ctx)
	case *ArrayType:
		return tv.visitArrayType(tv, it, ctx)
	case *PrimitiveType:
		return tv.visitPrimitive(tv, it, ctx)
	case *ObjectType:
		return tv.visitObjectType(tv, it, ctx)
	case *MapType:
		return tv.visitMapType(tv, it, ctx)
	case *OptionalType:
		return tv.visitOptionalType(tv, it, ctx)
	case *EnumType:
		return tv.visitEnumType(tv, it, ctx)
	case *ResourceType:
		return tv.visitResourceType(tv, it, ctx)
	case *FlaggedType:
		return tv.visitFlaggedType(tv, it, ctx)
	case *ValidatedType:
		return tv.visitValidatedType(tv, it, ctx)
	case *ErroredType:
		return tv.visitErroredType(tv, it, ctx)
	case *InterfaceType:
		return tv.visitInterfaceType(tv, it, ctx)
	}

	panic(fmt.Sprintf("unhandled type: (%T) %s", t, t))
}

// VisitTypeName invokes the TypeVisitor on the TypeName, returning a TypeName.
// name is the TypeName to visit.
// This is a convenience method for when a TypeName is expected as the result.
func (tv *TypeVisitor[C]) VisitTypeName(name TypeName, ctx C) (TypeName, error) {
	t, err := tv.visitTypeName(tv, name, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "visit of TypeName %q failed", name)
	}

	n, ok := t.(TypeName)
	if !ok {
		return nil, errors.Errorf("expected visit of TypeName %q to return TypeName, not %T", name, t)
	}

	return n, nil
}

// VisitDefinition invokes the TypeVisitor on both the name and type of the definition
// NB: this is only valid if visitTypeName returns a TypeName and not generally a Type
func (tv *TypeVisitor[C]) VisitDefinition(td TypeDefinition, ctx C) (TypeDefinition, error) {
	visitedName, err := tv.visitTypeName(tv, td.Name(), ctx)
	if err != nil {
		return TypeDefinition{}, errors.Wrapf(err, "visit of %q failed", td.Name())
	}

	name, ok := visitedName.(InternalTypeName)
	if !ok {
		return TypeDefinition{}, errors.Errorf("expected visit of %q to return TypeName, not %T", td.Name(), visitedName)
	}

	visitedType, err := tv.Visit(td.Type(), ctx)
	if err != nil {
		return TypeDefinition{}, errors.Wrapf(err, "visit of type of %q failed", td.Name())
	}

	def := td.WithName(name).WithType(visitedType)
	return def, nil
}

func (tv *TypeVisitor[C]) VisitDefinitions(definitions TypeDefinitionSet, ctx C) (TypeDefinitionSet, error) {
	result := make(TypeDefinitionSet)
	var errs []error
	for _, d := range definitions {
		def, err := tv.VisitDefinition(d, ctx)
		if err != nil {
			errs = append(errs, err)
		} else {
			result[def.Name()] = def
		}
	}

	if len(errs) > 0 {
		err := kerrors.NewAggregate(errs)
		return nil, err
	}

	return result, nil
}

func IdentityVisitOfTypeName[C any](_ *TypeVisitor[C], it TypeName, _ C) (Type, error) {
	return it, nil
}

func IdentityVisitOfArrayType[C any](this *TypeVisitor[C], it *ArrayType, ctx C) (Type, error) {
	newElement, err := this.Visit(it.element, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to visit type of array")
	}

	return it.WithElement(newElement), nil
}

func IdentityVisitOfPrimitiveType[C any](_ *TypeVisitor[C], it *PrimitiveType, _ C) (Type, error) {
	return it, nil
}

func IdentityVisitOfObjectType[C any](this *TypeVisitor[C], it *ObjectType, ctx C) (Type, error) {
	return identityVisitOfObjectTypeWithPerPropertyContext(
		this,
		it,
		ctx,
		func(_ *ObjectType, _ *PropertyDefinition, ctx C) (C, error) {
			return ctx, nil
		})
}

func OrderedIdentityVisitOfObjectType[C any](this *TypeVisitor[C], it *ObjectType, ctx C) (Type, error) {
	return orderedIdentityVisitOfObjectTypeWithPerPropertyContext(
		this,
		it,
		ctx,
		func(_ *ObjectType, _ *PropertyDefinition, ctx C) (C, error) {
			return ctx, nil
		})
}

type MakePerPropertyContext[C any] func(ot *ObjectType, prop *PropertyDefinition, ctx C) (C, error)

// MakeIdentityVisitOfObjectType creates a visitor function which creates a per-property context before visiting each
// property of the ObjectType
func MakeIdentityVisitOfObjectType[C any](
	makeCtx MakePerPropertyContext[C],
) func(this *TypeVisitor[C], it *ObjectType, ctx C) (Type, error) {
	return func(this *TypeVisitor[C], it *ObjectType, ctx C) (Type, error) {
		return identityVisitOfObjectTypeWithPerPropertyContext(this, it, ctx, makeCtx)
	}
}

func identityVisitOfObjectTypeWithPerPropertyContext[C any](
	this *TypeVisitor[C],
	it *ObjectType,
	ctx C,
	makeCtx MakePerPropertyContext[C],
) (Type, error) {
	// just map the property types

	var errs []error
	var newProps []*PropertyDefinition
	it.Properties().ForEach(func(prop *PropertyDefinition) {
		newCtx, err := makeCtx(it, prop, ctx)
		if err != nil {
			errs = append(errs, err)
			return // continue
		}

		p, err := this.Visit(prop.propertyType, newCtx)
		if err != nil {
			errs = append(errs, err)
		} else {
			// only replace property if the type was changed;
			// this allows short-circuiting below
			if !TypeEquals(p, prop.propertyType) {
				newProps = append(newProps, prop.WithType(p))
			}
		}
	})

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	// map the embedded types too
	embeddedPropsChanged := false
	var newEmbeddedProps []*PropertyDefinition
	for _, prop := range it.EmbeddedProperties() {
		newCtx, err := makeCtx(it, prop, ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		p, err := this.Visit(prop.propertyType, newCtx)
		if err != nil {
			errs = append(errs, err)
		} else {
			if !TypeEquals(p, prop.propertyType) {
				embeddedPropsChanged = true
			}

			newEmbeddedProps = append(newEmbeddedProps, prop.WithType(p))
		}
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	result := it.WithProperties(newProps...)

	var err error
	if embeddedPropsChanged {
		// Since it's possible that the type was renamed we need to clear the old embedded properties
		result = result.WithoutEmbeddedProperties()
		result, err = result.WithEmbeddedProperties(newEmbeddedProps...)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func orderedIdentityVisitOfObjectTypeWithPerPropertyContext[C any](
	this *TypeVisitor[C],
	it *ObjectType,
	ctx C,
	makeCtx MakePerPropertyContext[C],
) (Type, error) {
	// just map the property types

	var errs []error
	var newProps []*PropertyDefinition
	for _, prop := range it.Properties().AsSlice() {
		newCtx, err := makeCtx(it, prop, ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		p, err := this.Visit(prop.propertyType, newCtx)
		if err != nil {
			errs = append(errs, err)
		} else {
			// only replace property if the type was changed;
			// this allows short-circuiting below
			if !TypeEquals(p, prop.propertyType) {
				newProps = append(newProps, prop.WithType(p))
			}
		}
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	// map the embedded types too
	embeddedPropsChanged := false
	var newEmbeddedProps []*PropertyDefinition
	for _, prop := range it.EmbeddedProperties() {
		newCtx, err := makeCtx(it, prop, ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		p, err := this.Visit(prop.propertyType, newCtx)
		if err != nil {
			errs = append(errs, err)
		} else {
			if !TypeEquals(p, prop.propertyType) {
				embeddedPropsChanged = true
			}

			newEmbeddedProps = append(newEmbeddedProps, prop.WithType(p))
		}
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	result := it.WithProperties(newProps...)

	var err error
	if embeddedPropsChanged {
		// Since it's possible that the type was renamed we need to clear the old embedded properties
		result = result.WithoutEmbeddedProperties()
		result, err = result.WithEmbeddedProperties(newEmbeddedProps...)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func IdentityVisitOfMapType[C any](this *TypeVisitor[C], it *MapType, ctx C) (Type, error) {
	visitedKey, err := this.Visit(it.key, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit map key type %q", it.key)
	}

	visitedValue, err := this.Visit(it.value, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit map value type %q", it.value)
	}

	return it.WithKeyType(visitedKey).WithValueType(visitedValue), nil
}

func IdentityVisitOfEnumType[C any](_ *TypeVisitor[C], it *EnumType, _ C) (Type, error) {
	// if we visit the enum base type then we will also have to do something
	// about the values. so by default don't do anything with the enum base
	return it, nil
}

func IdentityVisitOfOptionalType[C any](this *TypeVisitor[C], it *OptionalType, ctx C) (Type, error) {
	visitedElement, err := this.Visit(it.element, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit optional element type %q", it.element)
	}

	return it.WithElement(visitedElement), nil
}

func IdentityVisitOfResourceType[C any](this *TypeVisitor[C], it *ResourceType, ctx C) (Type, error) {
	visitedSpec, err := this.Visit(it.spec, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit resource spec type %q", it.spec)
	}

	visitedStatus, err := this.Visit(it.status, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit resource status type %q", it.status)
	}

	changedAPIVersionName := false
	if it.HasAPIVersion() {
		originalAPIVersionTypeName := it.APIVersionTypeName()
		newAPIVersion, err := this.visitTypeName(this, originalAPIVersionTypeName, ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to visit resource API version name %q", originalAPIVersionTypeName)
		}

		if !TypeEquals(originalAPIVersionTypeName, newAPIVersion) {
			newAPIVersionName, ok := newAPIVersion.(TypeName)
			if !ok {
				return nil, errors.Wrapf(err, "attempted to change API Version type name into non-type name %q", newAPIVersion)
			}

			changedAPIVersionName = true
			it = it.WithAPIVersion(newAPIVersionName, it.APIVersionEnumValue())
		}
	}

	if visitedSpec == it.spec &&
		visitedStatus == it.status &&
		!changedAPIVersionName {
		return it, nil // short-circuit
	}

	return it.WithSpec(visitedSpec).WithStatus(visitedStatus), nil
}

func IdentityVisitOfOneOfType[C any](this *TypeVisitor[C], it *OneOfType, ctx C) (Type, error) {

	result := it.WithoutAnyPropertyObjects()

	propertyObjects := it.PropertyObjects()
	for _, obj := range propertyObjects {
		newObj, err := this.Visit(obj, ctx)
		if err != nil {
			return nil, err
		}

		obj, ok := newObj.(*ObjectType)
		if !ok {
			return nil, errors.Errorf("expected to visit oneof property object to result in object type, instead got %T", newObj)
		}

		result = result.WithAdditionalPropertyObject(obj)
	}

	newTypes := make([]Type, 0, it.Types().Len())
	err := it.Types().ForEachError(func(oneOf Type, _ int) error {
		newType, err := this.Visit(oneOf, ctx)
		if err != nil {
			return errors.Wrapf(err, "failed to visit oneOf")
		}

		newTypes = append(newTypes, newType)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result.WithTypes(newTypes), nil
}

func IdentityVisitOfAllOfType[C any](this *TypeVisitor[C], it *AllOfType, ctx C) (Type, error) {
	var newTypes []Type
	err := it.Types().ForEachError(func(allOf Type, _ int) error {
		newType, err := this.Visit(allOf, ctx)
		if err != nil {
			return errors.Wrapf(err, "failed to visit allOf")
		}

		newTypes = append(newTypes, newType)
		return nil
	})
	if err != nil {
		return nil, err
	}

	if typeSlicesFastEqual(newTypes, it.types.types) {
		return it, nil // short-circuit
	}

	return BuildAllOfType(newTypes...), nil
}

func IdentityVisitOfInterfaceType[C any](_ *TypeVisitor[C], it *InterfaceType, _ C) (Type, error) {
	// We don't visit the functions here to match ObjectType visit behavior
	return it, nil
}

// just checks reference equality of types
// this is used to short-circuit when we don't need to make new types,
// in a fast manner than invoking Equals and walking the whole tree
func typeSlicesFastEqual(t1 []Type, t2 []Type) bool {
	if len(t1) != len(t2) {
		return false
	}

	for ix := range t1 {
		if t1[ix] != t2[ix] {
			return false
		}
	}

	return true
}

func IdentityVisitOfFlaggedType[C any](this *TypeVisitor[C], ft *FlaggedType, ctx C) (Type, error) {
	nt, err := this.Visit(ft.element, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit flagged type %q", ft.element)
	}

	if nt == ft.element {
		return ft, nil // short-circuit
	}

	return ft.WithElement(nt), nil
}

func IdentityVisitOfValidatedType[C any](this *TypeVisitor[C], v *ValidatedType, ctx C) (Type, error) {
	nt, err := this.Visit(v.element, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit validated type %q", v.element)
	}

	if nt == v.element {
		return v, nil // short-circuit
	}

	return v.WithType(nt), nil
}

func IdentityVisitOfErroredType[C any](this *TypeVisitor[C], e *ErroredType, ctx C) (Type, error) {
	nt, err := this.Visit(e.inner, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit errored type %q", e.inner)
	}

	if nt == e.inner {
		return e, nil // short-circuit
	}

	return e.WithType(nt), nil
}
