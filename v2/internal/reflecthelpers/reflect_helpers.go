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
func FindResourceReferences(obj interface{}) (set.Set[genruntime.ResourceReference], error) {
	untypedResult, err := FindReferences(obj, reflect.TypeOf(genruntime.ResourceReference{}))
	if err != nil {
		return nil, err
	}

	result := set.Make[genruntime.ResourceReference]()
	for k := range untypedResult {
		result.Add(k.(genruntime.ResourceReference))
	}

	return result, nil
}

// FindSecretReferences finds all of the genruntime.SecretReference's on the provided object
func FindSecretReferences(obj interface{}) (set.Set[genruntime.SecretReference], error) {
	untypedResult, err := FindReferences(obj, reflect.TypeOf(genruntime.SecretReference{}))
	if err != nil {
		return nil, err
	}

	result := set.Make[genruntime.SecretReference]()
	for k := range untypedResult {
		result.Add(k.(genruntime.SecretReference)) 
	}

	return result, nil
}

// GetObjectListItems gets the list of items from an ObjectList
func GetObjectListItems(listPtr client.ObjectList) ([]client.Object, error) {
	itemsField, err := getItemsField(listPtr)
	if err != nil {
		return nil, err
	}

	var result []client.Object
	for i := 0; i < itemsField.Len(); i++ {
		item := itemsField.Index(i)

		if item.Kind() == reflect.Struct {
			if !item.CanAddr() {
				return nil, errors.Errorf("provided list elements were not pointers, but cannot be addressed")
			}
			item = item.Addr()
		}

		typedItem, ok := item.Interface().(client.Object)
		if !ok {
			return nil, errors.Errorf("provided list elements did not implement client.Object interface")
		}

		result = append(result, typedItem)
	}

	return result, nil
}

// SetObjectListItems gets the list of items from an ObjectList
func SetObjectListItems(listPtr client.ObjectList, items []client.Object) (returnErr error) {
	itemsField, err := getItemsField(listPtr)
	if err != nil {
		return err
	}

	if !itemsField.CanSet() {
		return errors.Errorf("cannot set items field of %T", listPtr)
	}
	defer func() {
		if recovered := recover(); recovered != nil {
			returnErr = errors.Errorf("failed to set items field of %T: %s", listPtr, recovered)
		}
	}()

	slice := reflect.MakeSlice(itemsField.Type(), 0, 0)
	for _, item := range items {
		val := reflect.ValueOf(item)

		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		slice = reflect.Append(slice, val)
	}

	itemsField.Set(slice)
	return nil
}

func getItemsField(listPtr client.ObjectList) (reflect.Value, error) {
	val := reflect.ValueOf(listPtr)
	if val.Kind() != reflect.Ptr {
		return reflect.Value{}, errors.Errorf("provided list was not a pointer, was %s", val.Kind())
	}

	list := val.Elem()

	if list.Kind() != reflect.Struct {
		return reflect.Value{}, errors.Errorf("provided list was not a struct, was %s", val.Kind())
	}

	itemsField := list.FieldByName("Items")
	if (itemsField == reflect.Value{}) {
		return reflect.Value{}, errors.Errorf("provided list has no field \"Items\"")
	}
	if itemsField.Kind() != reflect.Slice {
		return reflect.Value{}, errors.Errorf("provided list \"Items\" field was not of type slice")
	}

	return itemsField, nil
}
