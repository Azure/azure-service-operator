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
type TypeVisitor struct {
	VisitTypeName     func(this *TypeVisitor, it TypeName, ctx interface{}) (Type, error)
	VisitArrayType    func(this *TypeVisitor, it *ArrayType, ctx interface{}) (Type, error)
	VisitPrimitive    func(this *TypeVisitor, it *PrimitiveType, ctx interface{}) (Type, error)
	VisitObjectType   func(this *TypeVisitor, it *ObjectType, ctx interface{}) (Type, error)
	VisitMapType      func(this *TypeVisitor, it *MapType, ctx interface{}) (Type, error)
	VisitOptionalType func(this *TypeVisitor, it *OptionalType, ctx interface{}) (Type, error)
	VisitEnumType     func(this *TypeVisitor, it *EnumType, ctx interface{}) (Type, error)
	VisitResourceType func(this *TypeVisitor, it *ResourceType, ctx interface{}) (Type, error)
}

// Visit invokes the appropriate VisitX on TypeVisitor
func (tv *TypeVisitor) Visit(t Type, ctx interface{}) (Type, error) {
	if t == nil {
		return nil, nil
	}

	switch it := t.(type) {
	case TypeName:
		return tv.VisitTypeName(tv, it, ctx)
	case *ArrayType:
		return tv.VisitArrayType(tv, it, ctx)
	case *PrimitiveType:
		return tv.VisitPrimitive(tv, it, ctx)
	case *ObjectType:
		return tv.VisitObjectType(tv, it, ctx)
	case *MapType:
		return tv.VisitMapType(tv, it, ctx)
	case *OptionalType:
		return tv.VisitOptionalType(tv, it, ctx)
	case *EnumType:
		return tv.VisitEnumType(tv, it, ctx)
	case *ResourceType:
		return tv.VisitResourceType(tv, it, ctx)
	}

	panic(fmt.Sprintf("unhandled type: (%T) %v", t, t))
}

// VisitDefinition invokes the TypeVisitor on both the name and type of the definition
// NB: this is only valid if VisitTypeName returns a TypeName and not generally a Type
func (tv *TypeVisitor) VisitDefinition(td TypeDefinition, ctx interface{}) (*TypeDefinition, error) {
	visitedName, err := tv.VisitTypeName(tv, td.Name(), ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "visit of %q failed", td.Name())
	}

	name, ok := visitedName.(TypeName)
	if !ok {
		return nil, errors.Errorf("expected visit of %q to return TypeName, not %T", td.Name(), visitedName)
	}

	visitedType, err := tv.Visit(td.Type(), ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "visit of type of %q failed", td.Name())
	}

	def := MakeTypeDefinition(name, visitedType)
	return &def, nil
}

// MakeTypeVisitor returns a default (identity transform) visitor, which
// visits every type in the tree. If you want to actually do something you will
// need to override the properties on the returned TypeVisitor.
func MakeTypeVisitor() TypeVisitor {
	// TODO [performance]: we can do reference comparisons on the results of
	// recursive invocations of Visit to avoid having to rebuild the tree if the
	// leaf nodes do not actually change.
	return TypeVisitor{
		VisitTypeName: func(_ *TypeVisitor, it TypeName, _ interface{}) (Type, error) {
			return it, nil
		},
		VisitArrayType: func(this *TypeVisitor, it *ArrayType, ctx interface{}) (Type, error) {
			newElement, err := this.Visit(it.element, ctx)
			if err != nil {
				return nil, errors.Wrap(err, "failed to visit type of array")
			}

			return NewArrayType(newElement), nil
		},
		VisitPrimitive: func(_ *TypeVisitor, it *PrimitiveType, _ interface{}) (Type, error) {
			return it, nil
		},
		VisitObjectType: func(this *TypeVisitor, it *ObjectType, ctx interface{}) (Type, error) {
			// just map the property types
			var errs []error
			var newProps []*PropertyDefinition
			for _, prop := range it.properties {
				p, err := this.Visit(prop.propertyType, ctx)
				if err != nil {
					errs = append(errs, err)
				} else {
					newProps = append(newProps, prop.WithType(p))
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return it.WithProperties(newProps...), nil
		},
		VisitMapType: func(this *TypeVisitor, it *MapType, ctx interface{}) (Type, error) {
			visitedKey, err := this.Visit(it.key, ctx)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to visit map key type %v", it.key)
			}

			visitedValue, err := this.Visit(it.value, ctx)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to visit map value type %v", it.value)
			}

			return NewMapType(visitedKey, visitedValue), nil
		},
		VisitEnumType: func(_ *TypeVisitor, it *EnumType, _ interface{}) (Type, error) {
			// if we visit the enum base type then we will also have to do something
			// about the values. so by default don't do anything with the enum base
			return it, nil
		},
		VisitOptionalType: func(this *TypeVisitor, it *OptionalType, ctx interface{}) (Type, error) {
			visitedElement, err := this.Visit(it.element, ctx)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to visit optional element type %v", it.element)
			}

			return NewOptionalType(visitedElement), nil
		},
		VisitResourceType: func(this *TypeVisitor, it *ResourceType, ctx interface{}) (Type, error) {
			visitedSpec, err := this.Visit(it.spec, ctx)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to visit resource spec type %v", it.spec)
			}

			visitedStatus, err := this.Visit(it.status, ctx)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to visit resource status type %v", it.status)
			}

			return NewResourceType(visitedSpec, visitedStatus).WithOwner(it.Owner()), nil
		},
	}
}
