/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

type ResourceKind string

const (
	// ResourceKindNormal is a standard ARM resource.
	ResourceKindNormal = ResourceKind("normal")
	// ResourceKindExtension is an extension resource. Extension resources can have any resource as their parent.
	ResourceKindExtension = ResourceKind("extension")
)

// TODO: These should become Status properties at some point.
const (
	ResourceIDAnnotation = "resource-id.azure.com"
)

// MetaObject represents an arbitrary ASO custom resource
type MetaObject interface {
	runtime.Object
	metav1.Object
	KubernetesResource
}

type Reconciler interface {
	Reconcile(ctx context.Context) (ctrl.Result, error)
}

// TODO: We really want these methods to be on MetaObject itself -- should update code generator to make them at some point
func GetResourceID(obj MetaObject) (string, bool) {
	result, ok := obj.GetAnnotations()[ResourceIDAnnotation]
	return result, ok
}

func GetResourceIDOrDefault(obj MetaObject) string {
	return obj.GetAnnotations()[ResourceIDAnnotation]
}

func SetResourceID(obj MetaObject, id string) {
	AddAnnotation(obj, ResourceIDAnnotation, id)
}

func AddAnnotation(obj MetaObject, k string, v string) {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	// I think this is the behavior we want...
	if v == "" {
		delete(annotations, k)
	} else {
		annotations[k] = v
	}
	obj.SetAnnotations(annotations)
}

// ARMResourceSpec is an ARM resource specification. This interface contains
// methods to access properties common to all ARM Resource Specs. An Azure
// Deployment is made of these.
type ARMResourceSpec interface {
	GetAPIVersion() string

	GetType() string

	GetName() string
}

// ARMResourceStatus is an ARM resource status
type ARMResourceStatus interface { // TODO: Unsure what the actual content of this interface needs to be.
	// TODO: We need to define it and generate the code for it
	// GetId() string
}

type ARMResource interface {
	Spec() ARMResourceSpec
	Status() ARMResourceStatus

	GetID() string // TODO: Should this be on Status instead?
}

func NewArmResource(spec ARMResourceSpec, status ARMResourceStatus, id string) ARMResource {
	return &armResourceImpl{
		spec:   spec,
		status: status,
		Id:     id,
	}
}

type armResourceImpl struct {
	spec   ARMResourceSpec
	status ARMResourceStatus
	Id     string
}

var _ ARMResource = &armResourceImpl{}

func (resource *armResourceImpl) Spec() ARMResourceSpec {
	return resource.spec
}

func (resource *armResourceImpl) Status() ARMResourceStatus {
	return resource.status
}

func (resource *armResourceImpl) GetID() string {
	return resource.Id
}

type DeployableResource interface {
	Spec() ARMResourceSpec
}

func NewDeployableResourceGroupResource(resourceGroup string, spec ARMResourceSpec) *ResourceGroupResource {
	return &ResourceGroupResource{
		resourceGroup: resourceGroup,
		spec:          spec,
	}
}

func NewDeployableSubscriptionResource(location string, spec ARMResourceSpec) *SubscriptionResource {
	return &SubscriptionResource{
		location: location,
		spec:     spec,
	}
}

// ResourceGroupResource represents a resource which can be deployed to Azure inside of a
// resource group.
type ResourceGroupResource struct {
	resourceGroup string
	spec          ARMResourceSpec
}

var _ DeployableResource = &ResourceGroupResource{}

func (r *ResourceGroupResource) ResourceGroup() string {
	return r.resourceGroup
}

func (r *ResourceGroupResource) Spec() ARMResourceSpec {
	return r.spec
}

// SubscriptionResource represents a resource which can be deployed to Azure directly
// in a subscription (not inside of a resource group).
type SubscriptionResource struct {
	location string
	spec     ARMResourceSpec
}

var _ DeployableResource = &SubscriptionResource{}

func (r *SubscriptionResource) Location() string {
	return r.location
}

func (r *SubscriptionResource) Spec() ARMResourceSpec {
	return r.spec
}
