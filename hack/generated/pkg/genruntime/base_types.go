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

// TODO: These should become Status properties at some point.
const (
	ResourceIDAnnotation = "resource-id.infra.azure.com"
)

// MetaObject represents an arbitrary k8s-infra custom resource
type MetaObject interface {
	runtime.Object
	metav1.Object
	KubernetesResource
}

type Reconciler interface { // TODO: Sorta awkward interface name
	CreateOrUpdate(ctx context.Context) (ctrl.Result, error)

	Delete(ctx context.Context) (ctrl.Result, error)

	// TODO: Do we want a "Diff"? Or is the expectation that logic just exists in CreateOrUpdate?
}

// KubernetesResource is an Azure resource. This interface contains the common set of
// methods that apply to all k8s-infra resources.
type KubernetesResource interface {
	// Owner returns the ResourceReference of the owner, or nil if there is no owner
	Owner() *ResourceReference

	// TODO: I think we need this?
	// KnownOwner() *KnownResourceReference

	// AzureName returns the Azure name of the resource
	AzureName() string

	// Some types, but not all, have a corresponding:
	// 	SetAzureName(name string)
	// They do not if the name must be a fixed value (like 'default').

	// TODO: GetAPIVersion here?

	// TODO: I think we need this
	// SetStatus(status interface{})
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

// ArmResourceSpec is an ARM resource specification. This interface contains
// methods to access properties common to all ARM Resource Specs. An Azure
// Deployment is made of these.
type ArmResourceSpec interface {
	GetApiVersion() string

	GetType() string

	GetName() string
}

// ArmResourceStatus is an ARM resource status
type ArmResourceStatus interface {
	// TODO: Unsure what the actual content of this interface needs to be.
	// TODO: We need to define it and generate the code for it

	// GetId() string
}

type ArmResource interface {
	Spec() ArmResourceSpec
	Status() ArmResourceStatus

	// TODO: Golang wants this to be GetID
	GetId() string // TODO: Should this be on Status instead?
}

func NewArmResource(spec ArmResourceSpec, status ArmResourceStatus, id string) ArmResource {
	return &armResourceImpl{
		spec:   spec,
		status: status,
		Id:     id,
	}
}

type armResourceImpl struct {
	spec   ArmResourceSpec
	status ArmResourceStatus
	Id     string
}

var _ ArmResource = &armResourceImpl{}

func (resource *armResourceImpl) Spec() ArmResourceSpec {
	return resource.spec
}

func (resource *armResourceImpl) Status() ArmResourceStatus {
	return resource.status
}

func (resource *armResourceImpl) GetId() string {
	return resource.Id
}

type DeployableResource interface {
	Spec() ArmResourceSpec
}

func NewDeployableResourceGroupResource(resourceGroup string, spec ArmResourceSpec) *ResourceGroupResource {
	return &ResourceGroupResource{
		resourceGroup: resourceGroup,
		spec:          spec,
	}
}

func NewDeployableSubscriptionResource(location string, spec ArmResourceSpec) *SubscriptionResource {
	return &SubscriptionResource{
		location: location,
		spec:     spec,
	}
}

// ResourceGroupResource represents a resource which can be deployed to Azure inside of a
// resource group.
type ResourceGroupResource struct {
	resourceGroup string
	spec          ArmResourceSpec
}

var _ DeployableResource = &ResourceGroupResource{}

func (r *ResourceGroupResource) ResourceGroup() string {
	return r.resourceGroup
}

func (r *ResourceGroupResource) Spec() ArmResourceSpec {
	return r.spec
}

// SubscriptionResource represents a resource which can be deployed to Azure directly
// in a subscription (not inside of a resource group).
type SubscriptionResource struct {
	location string
	spec     ArmResourceSpec
}

var _ DeployableResource = &SubscriptionResource{}

func (r *SubscriptionResource) Location() string {
	return r.location
}

func (r *SubscriptionResource) Spec() ArmResourceSpec {
	return r.spec
}
