/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reflecthelpers

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// ValueOfPtr dereferences a pointer and returns the value the pointer points to.
// Use this as carefully as you would the * operator
// TODO: Can we delete this helper later when we have some better code generated functions?
func ValueOfPtr(ptr interface{}) interface{} {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("Can't get value of pointer for non-pointer type %T", ptr))
	}
	val := reflect.Indirect(v)

	return val.Interface()
}

// DeepCopyInto calls in.DeepCopyInto(out)
func DeepCopyInto(in client.Object, out client.Object) {
	inVal := reflect.ValueOf(in)

	method := inVal.MethodByName("DeepCopyInto")
	method.Call([]reflect.Value{reflect.ValueOf(out)})
}

// FindReferences finds references of the given type on the provided object
func FindReferences(obj interface{}, t reflect.Type) (map[interface{}]struct{}, error) {
	result := make(map[interface{}]struct{})

	visitor := NewReflectVisitor()
	visitor.VisitStruct = func(this *ReflectVisitor, it interface{}, ctx interface{}) error {
		if reflect.TypeOf(it) == t {
			val := reflect.ValueOf(it)
			if val.CanInterface() {
				result[val.Interface()] = struct{}{}
			}
			return nil
		}

		return IdentityVisitStruct(this, it, ctx)
	}

	err := visitor.Visit(obj, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "scanning for references of type %s", t.String())
	}

	return result, nil
}

// FindResourceReferences finds all the genruntime.ResourceReference's on the provided object
func FindResourceReferences(obj interface{}) (map[genruntime.ResourceReference]struct{}, error) {
	untypedResult, err := FindReferences(obj, reflect.TypeOf(genruntime.ResourceReference{}))
	if err != nil {
		return nil, err
	}

	result := make(map[genruntime.ResourceReference]struct{})
	for k := range untypedResult {
		result[k.(genruntime.ResourceReference)] = struct{}{}
	}

	return result, nil
}
