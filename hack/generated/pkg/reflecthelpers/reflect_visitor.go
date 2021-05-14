/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reflecthelpers

import (
	"reflect"

	"github.com/pkg/errors"
)

var primitiveKinds = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Uintptr,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
	reflect.String,
}

// IsPrimitiveKind returns true if the provided reflect.Kind is for a primitive type, otherwise false.
func IsPrimitiveKind(k reflect.Kind) bool {
	for _, kind := range primitiveKinds {
		if k == kind {
			return true
		}
	}

	return false
}

// ReflectVisitor allows traversing an arbitrary object graph.
type ReflectVisitor struct {
	VisitPrimitive func(this *ReflectVisitor, it interface{}, ctx interface{}) error
	VisitStruct    func(this *ReflectVisitor, it interface{}, ctx interface{}) error
	VisitPtr       func(this *ReflectVisitor, it interface{}, ctx interface{}) error
	VisitSlice     func(this *ReflectVisitor, it interface{}, ctx interface{}) error
	VisitMap       func(this *ReflectVisitor, it interface{}, ctx interface{}) error
}

// NewReflectVisitor creates an identity ReflectVisitor.
func NewReflectVisitor() *ReflectVisitor {
	return &ReflectVisitor{
		VisitPrimitive: IdentityVisitPrimitive,
		VisitStruct:    IdentityVisitStruct,
		VisitPtr:       IdentityVisitPtr,
		VisitSlice:     IdentityVisitSlice,
		VisitMap:       IdentityVisitMap,
	}
}

// Visit visits the provided value. The ctx parameter can be used to pass data through the visit hierarchy.
func (r *ReflectVisitor) Visit(val interface{}, ctx interface{}) error {
	if val == nil {
		return nil
	}

	t := reflect.TypeOf(val)

	// This can happen because an interface holding nil is not itself nil
	v := reflect.ValueOf(val)
	if v.IsZero() {
		return nil
	}

	kind := t.Kind()
	if IsPrimitiveKind(kind) {
		return r.VisitPrimitive(r, val, ctx)
	}

	switch kind { // nolint: exhaustive
	case reflect.Ptr:
		return r.VisitPtr(r, val, ctx)
	case reflect.Slice:
		return r.VisitSlice(r, val, ctx)
	case reflect.Map:
		return r.VisitMap(r, val, ctx)
	case reflect.Struct:
		return r.VisitStruct(r, val, ctx)
	default:
		return errors.Errorf("unknown reflect.Kind: %s", kind)
	}
}

// IdentityVisitPrimitive is the identity visit function for primitive types.
func IdentityVisitPrimitive(this *ReflectVisitor, it interface{}, ctx interface{}) error {
	return nil
}

// IdentityVisitPtr is the identity visit function for pointer types. It dereferences the pointer and visits the type
// pointed to.
func IdentityVisitPtr(this *ReflectVisitor, it interface{}, ctx interface{}) error {
	val := reflect.ValueOf(it)
	elem := val.Elem()

	if elem.IsZero() {
		return nil
	}

	return this.Visit(elem.Interface(), ctx)
}

// IdentityVisitSlice is the identity visit function for slices. It visits each element of the slice.
func IdentityVisitSlice(this *ReflectVisitor, it interface{}, ctx interface{}) error {
	val := reflect.ValueOf(it)

	for i := 0; i < val.Len(); i++ {
		err := this.Visit(val.Index(i).Interface(), ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// IdentityVisitMap is the identity visit function for maps. It visits each key and value in the map.
func IdentityVisitMap(this *ReflectVisitor, it interface{}, ctx interface{}) error {
	val := reflect.ValueOf(it)

	for _, key := range val.MapKeys() {

		err := this.Visit(key.Interface(), ctx)
		if err != nil {
			return err
		}

		err = this.Visit(val.MapIndex(key).Interface(), ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// IdentityVisitStruct is the identity visit function for structs. It visits each exported field of the struct.
func IdentityVisitStruct(this *ReflectVisitor, it interface{}, ctx interface{}) error {
	val := reflect.ValueOf(it)

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		if !fieldVal.CanInterface() {
			// Bypass unexported fields
			continue
		}

		err := this.Visit(val.Field(i).Interface(), ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
