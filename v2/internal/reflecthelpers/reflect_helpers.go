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

// ConvertResourceToDeployableResource converts a genruntime.MetaObject (a Kubernetes representation of a resource) into
// a genruntime.DeployableResource - a specification which can be submitted to Azure for deployment
func ConvertResourceToDeployableResource(
	ctx context.Context,
	resolver *genruntime.Resolver,
	metaObject genruntime.MetaObject) (genruntime.DeployableResource, error) {

	spec, err := genruntime.GetVersionedSpec(metaObject, resolver.Scheme())
	if err != nil {
		return nil, errors.Errorf("unable to get spec from %s", metaObject.GetObjectKind().GroupVersionKind())
	}

	armTransformer, ok := spec.(genruntime.ARMTransformer)
	if !ok {
		return nil, errors.Errorf("spec was of type %T which doesn't implement genruntime.ArmTransformer", spec)
	}

	resourceHierarchy, resolvedDetails, err := resolve(ctx, resolver, metaObject)
	if err != nil {
		return nil, err
	}

	armSpec, err := armTransformer.ConvertToARM(resolvedDetails)
	if err != nil {
		return nil, errors.Wrapf(err, "transforming resource %s to ARM", metaObject.GetName())
	}

	typedArmSpec, ok := armSpec.(genruntime.ARMResourceSpec)
	if !ok {
		return nil, errors.Errorf("casting armSpec of type %T to genruntime.ArmResourceSpec", armSpec)
	}

	// We have different deployment models for Subscription rooted vs ResourceGroup rooted resources
	rootKind := resourceHierarchy.RootKind()
	if rootKind == genruntime.ResourceHierarchyRootResourceGroup {
		rg, err := resourceHierarchy.ResourceGroup()
		if err != nil {
			return nil, errors.Wrapf(err, "getting resource group")
		}
		return genruntime.NewDeployableResourceGroupResource(rg, typedArmSpec), nil
	} else if rootKind == genruntime.ResourceHierarchyRootSubscription {
		location, err := resourceHierarchy.Location()
		if err != nil {
			return nil, errors.Wrapf(err, "getting location")
		}
		return genruntime.NewDeployableSubscriptionResource(location, typedArmSpec), nil
	} else {
		return nil, errors.Errorf("unknown resource hierarchy root kind %s", rootKind)
	}
}

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
//TODO: this no longer uses reflection, inline it where used
func NewEmptyStatus(metaObject genruntime.MetaObject) (genruntime.FromARMConverter, error) {
	status, ok := metaObject.NewEmptyStatus().(genruntime.FromARMConverter)
	if !ok {
		return nil, errors.Errorf(
			"status %s did not implement genruntime.ArmTransformer", metaObject.GetObjectKind().GroupVersionKind())
	}

	return status, nil
}

// SetStatus updates a genruntime.MetaObject with a new status
//TODO: this no longer uses reflection, inline it where used
func SetStatus(metaObj genruntime.MetaObject, status interface{}) error {

	s, ok := status.(genruntime.ConvertibleStatus)
	if !ok {
		return errors.Errorf("expected SetStatus() to be passed a genruntime.ConvertibleStatus but received %T", status)
	}

	return metaObj.SetStatus(s)
}

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
