/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reflecthelpers

import (
	"reflect"
	"slices"

	"github.com/rotisserie/eris"
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
	return slices.Contains(primitiveKinds, k)
}

// ReflectVisitor allows traversing an arbitrary object graph.
type ReflectVisitor struct {
	VisitPrimitive func(this *ReflectVisitor, it reflect.Value, ctx any) error
	VisitStruct    func(this *ReflectVisitor, it reflect.Value, ctx any) error
	VisitPtr       func(this *ReflectVisitor, it reflect.Value, ctx any) error
	VisitSlice     func(this *ReflectVisitor, it reflect.Value, ctx any) error
	VisitMap       func(this *ReflectVisitor, it reflect.Value, ctx any) error
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
func (r *ReflectVisitor) Visit(val any, ctx any) error {
	if val == nil {
		return nil
	}

	// This can happen because an interface holding nil is not itself nil
	v := reflect.ValueOf(val)
	return r.visit(v, ctx)
}

func (r *ReflectVisitor) visit(val reflect.Value, ctx any) error {
	if val.IsZero() {
		return nil
	}

	kind := val.Type().Kind()
	if IsPrimitiveKind(kind) {
		return r.VisitPrimitive(r, val, ctx)
	}

	switch kind {
	case reflect.Pointer:
		return r.VisitPtr(r, val, ctx)
	case reflect.Slice:
		return r.VisitSlice(r, val, ctx)
	case reflect.Map:
		return r.VisitMap(r, val, ctx)
	case reflect.Struct:
		return r.VisitStruct(r, val, ctx)
	default:
		return eris.Errorf("unknown reflect.Kind: %s", kind)
	}
}

// IdentityVisitPrimitive is the identity visit function for primitive types.
func IdentityVisitPrimitive(this *ReflectVisitor, it reflect.Value, ctx any) error {
	return nil
}

// IdentityVisitPtr is the identity visit function for pointer types. It dereferences the pointer and visits the type
// pointed to.
func IdentityVisitPtr(this *ReflectVisitor, it reflect.Value, ctx any) error {
	elem := it.Elem()

	if elem.IsZero() {
		return nil
	}

	return this.visit(elem, ctx)
}

// IdentityVisitSlice is the identity visit function for slices. It visits each element of the slice.
func IdentityVisitSlice(this *ReflectVisitor, it reflect.Value, ctx any) error {
	for i := 0; i < it.Len(); i++ {
		err := this.visit(it.Index(i), ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// IdentityVisitMap is the identity visit function for maps. It visits each key and value in the map.
func IdentityVisitMap(this *ReflectVisitor, it reflect.Value, ctx any) error {
	for _, key := range it.MapKeys() {

		err := this.visit(key, ctx)
		if err != nil {
			return err
		}

		err = this.visit(it.MapIndex(key), ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// IdentityVisitStruct is the identity visit function for structs. It visits each exported field of the struct.
func IdentityVisitStruct(this *ReflectVisitor, it reflect.Value, ctx any) error {
	for _, fieldVal := range it.Fields() {
		if !fieldVal.CanInterface() {
			// Bypass unexported fields
			continue
		}

		err := this.visit(fieldVal, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
