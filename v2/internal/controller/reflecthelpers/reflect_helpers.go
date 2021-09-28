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
	t := reflect.TypeOf(metaObject).Elem()

	statusField, ok := t.FieldByName("Status")
	if !ok {
		return nil, errors.Errorf("couldn't find status field on type %T", metaObject)
	}

	statusPtr := reflect.New(statusField.Type)
	status, ok := statusPtr.Interface().(genruntime.FromARMConverter)
	if !ok {
		return nil, errors.Errorf("status did not implement genruntime.ArmTransformer")
	}

	return status, nil
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

// TODO: Consider moving this into genruntime.Resolver - need to fix package hierarchy to make that work though
func resolve(
	ctx context.Context,
	resolver *genruntime.Resolver,
	metaObject genruntime.MetaObject) (genruntime.ResourceHierarchy, genruntime.ConvertToARMResolvedDetails, error) {

	resourceHierarchy, err := resolver.ResolveResourceHierarchy(ctx, metaObject)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	// Find all of the references
	refs, err := FindResourceReferences(metaObject)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, errors.Wrapf(err, "finding references on %q", metaObject.GetName())
	}

	// resolve them
	resolvedRefs, err := resolver.ResolveReferencesToARMIDs(ctx, refs)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, errors.Wrapf(err, "failed resolving ARM IDs for references")
	}

	resolvedDetails := genruntime.ConvertToARMResolvedDetails{
		Name:               resourceHierarchy.FullAzureName(),
		ResolvedReferences: resolvedRefs,
	}

	// Augment with Scope information if the resource is an extension resource
	if metaObject.GetResourceKind() == genruntime.ResourceKindExtension {
		scope, err := resourceHierarchy.Scope()
		if err != nil {
			return nil, genruntime.ConvertToARMResolvedDetails{}, errors.Wrapf(err, "couldn't get resource scope")
		}
		resolvedDetails.Scope = &scope
	}

	return resourceHierarchy, resolvedDetails, nil
}

func extractARMTransformer(metaObject genruntime.MetaObject) (genruntime.ARMTransformer, error) {
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

	armTransformer, ok := spec.(genruntime.ARMTransformer)
	if !ok {
		return nil, errors.Errorf("spec was of type %T which doesn't implement genruntime.ArmTransformer", spec)
	}

	return armTransformer, nil
}
