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

// NewEmptyArmResourceStatus creates an empty genruntime.ARMResourceStatus from a genruntime.MetaObject
// (a Kubernetes representation of a resource), which can be filled by a call to Azure
func NewEmptyArmResourceStatus(metaObject genruntime.MetaObject) (genruntime.ARMResourceStatus, error) {
	kubeStatus, err := NewEmptyStatus(metaObject)
	if err != nil {
		return nil, err
	}

	armStatus := kubeStatus.CreateEmptyARMValue()
	return armStatus, nil
}

// NewEmptyStatus creates a new empty Status object (which implements FromArmConverter) from
// a genruntime.MetaObject.
func NewEmptyStatus(metaObject genruntime.MetaObject) (genruntime.FromARMConverter, error) {

	// This needs to return the right kind of status
	status := metaObject.GetStatus()

	statusType := reflect.TypeOf(status)
	if statusType.Kind() != reflect.Ptr {
		return nil, errors.Errorf(
			"expected GetStatus() on %T to return a pointer implementation of genruntime.ConvertibleStatus", 
			metaObject.GetName())
	}

	statusPtr := reflect.New(statusType.Elem())
	armStatus, ok := statusPtr.Interface().(genruntime.FromARMConverter)
	if !ok {
		return nil, errors.Errorf("status of type %T does not implement genruntime.FromARMConverter", statusPtr.Interface())
	}

	return armStatus, nil
}

// NewPtrFromValue creates a new pointer type from a value
func NewPtrFromValue(value interface{}) interface{} {
	v := reflect.ValueOf(value)

	// Spec fields are values, we want a ptr
	ptr := reflect.New(v.Type())
	ptr.Elem().Set(v)

	return ptr.Interface()
}

// TODO: Can we delete this helper later when we have some better code generated functions?
// ValueOfPtr dereferences a pointer and returns the value the pointer points to.
// Use this as carefully as you would the * operator
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

// FindResourceReferences finds all ResourceReferences specified by a given genruntime.ARMTransformer (resource spec)
func FindResourceReferences(transformer interface{}) (map[genruntime.ResourceReference]struct{}, error) {
	result := make(map[genruntime.ResourceReference]struct{})

	visitor := NewReflectVisitor()
	visitor.VisitStruct = func(this *ReflectVisitor, it interface{}, ctx interface{}) error {
		if reflect.TypeOf(it) == reflect.TypeOf(genruntime.ResourceReference{}) {
			val := reflect.ValueOf(it)
			if val.CanInterface() {
				result[val.Interface().(genruntime.ResourceReference)] = struct{}{}
			}
			return nil
		}

		return IdentityVisitStruct(this, it, ctx)
	}

	err := visitor.Visit(transformer, nil)
	if err != nil {
		return nil, errors.Wrap(err, "scanning for genruntime.ResourceReference")
	}

	return result, nil
}

