/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reflecthelpers

import (
	"context"
	"fmt"
	"reflect"

	"github.com/pkg/errors"

	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/armresourceresolver"
)

// ResourceSpecToArmResourceSpec converts a genruntime.MetaObject (a Kubernetes representation of a resource) into
// a genruntime.ArmResourceSpec - a specification which can be submitted to Azure for deployment
func ConvertResourceToDeployableResource(
	ctx context.Context,
	resolver *armresourceresolver.Resolver,
	metaObject genruntime.MetaObject) (genruntime.DeployableResource, error) {

	metaObjReflector := reflect.Indirect(reflect.ValueOf(metaObject))
	if !metaObjReflector.IsValid() {
		return nil, errors.Errorf("couldn't indirect %T", metaObject)
	}

	specField := metaObjReflector.FieldByName("Spec")
	if !specField.IsValid() {
		return nil, errors.Errorf("couldn't find spec field on type %T", metaObject)
	}

	// Spec fields are values, we want a ptr
	specFieldPtr := reflect.New(specField.Type())
	specFieldPtr.Elem().Set(specField)

	spec := specFieldPtr.Interface()

	armTransformer, ok := spec.(genruntime.ArmTransformer)
	if !ok {
		return nil, errors.Errorf("spec was of type %T which doesn't implement genruntime.ArmTransformer", spec)
	}

	resourceHierarchy, err := resolver.ResolveResourceHierarchy(ctx, metaObject)
	if err != nil {
		return nil, err
	}

	armSpec, err := armTransformer.ConvertToArm(resourceHierarchy.FullAzureName())
	if err != nil {
		return nil, errors.Wrapf(err, "transforming resource %s to ARM", metaObject.GetName())
	}

	typedArmSpec, ok := armSpec.(genruntime.ArmResourceSpec)
	if !ok {
		return nil, errors.Errorf("casting armSpec of type %T to genruntime.ArmResourceSpec", armSpec)
	}

	// We have different deployment models for Subscription rooted vs ResourceGroup rooted resources
	rootKind := resourceHierarchy.RootKind()
	if rootKind == armresourceresolver.ResourceHierarchyRootResourceGroup {
		rg, err := resourceHierarchy.ResourceGroup()
		if err != nil {
			return nil, errors.Wrapf(err, "getting resource group")
		}
		return genruntime.NewDeployableResourceGroupResource(rg, typedArmSpec), nil
	} else if rootKind == armresourceresolver.ResourceHierarchyRootSubscription {
		location, err := resourceHierarchy.Location()
		if err != nil {
			return nil, errors.Wrapf(err, "getting location")
		}
		return genruntime.NewDeployableSubscriptionResource(location, typedArmSpec), nil
	} else {
		return nil, errors.Errorf("unknown resource hierarchy root kind %s", rootKind)
	}
}

// NewEmptyArmResourceStatus creates an empty genruntime.ArmResourceStatus from a genruntime.MetaObject
// (a Kubernetes representation of a resource), which can be filled by a call to Azure
func NewEmptyArmResourceStatus(metaObject genruntime.MetaObject) (genruntime.ArmResourceStatus, error) {
	kubeStatus, err := NewEmptyStatus(metaObject)
	if err != nil {
		return nil, err
	}

	// TODO: Do we actually want to return a ptr here, not a value?
	armStatus := kubeStatus.CreateEmptyArmValue()

	// TODO: Some reflect hackery here to make sure that this is a ptr not a value
	armStatusPtr := NewPtrFromValue(armStatus)
	castArmStatus, ok := armStatusPtr.(genruntime.ArmResourceStatus)
	if !ok {
		// TODO: Should these be panics instead - they aren't really recoverable?
		return nil, errors.Errorf("resource status %T did not implement genruntime.ArmResourceStatus", armStatus)
	}

	return castArmStatus, nil
}

// NewEmptyStatus creates a new empty Status object (which implements FromArmConverter) from
// a genruntime.MetaObject.
func NewEmptyStatus(metaObject genruntime.MetaObject) (genruntime.FromArmConverter, error) {
	t := reflect.TypeOf(metaObject).Elem()

	statusField, ok := t.FieldByName("Status")
	if !ok {
		return nil, errors.Errorf("couldn't find status field on type %T", metaObject)
	}

	statusPtr := reflect.New(statusField.Type)
	status, ok := statusPtr.Interface().(genruntime.FromArmConverter)
	if !ok {
		return nil, errors.Errorf("status did not implement genruntime.ArmTransformer")
	}

	return status, nil
}

// TODO: hacking this for now -- replace with a code-generated method later
func SetStatus(metaObj genruntime.MetaObject, status interface{}) error {
	ptr := reflect.ValueOf(metaObj)
	val := ptr.Elem()

	if val.Kind() != reflect.Struct {
		return errors.Errorf("metaObj kind was not struct")
	}

	field := val.FieldByName("Status")
	statusVal := reflect.ValueOf(status).Elem()
	field.Set(statusVal)

	return nil
}

func HasStatus(metaObj genruntime.MetaObject) (bool, error) {
	ptr := reflect.ValueOf(metaObj)
	val := ptr.Elem()

	if val.Kind() != reflect.Struct {
		return false, errors.Errorf("metaObj kind was not struct")
	}

	field := val.FieldByName("Status")
	emptyStatusPtr := reflect.New(field.Type())
	emptyStatus := emptyStatusPtr.Elem().Interface()

	return !reflect.DeepEqual(field.Interface(), emptyStatus), nil
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
