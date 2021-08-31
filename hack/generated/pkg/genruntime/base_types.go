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

	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
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

// KubernetesResource is an Azure resource. This interface contains the common set of
// methods that apply to all ASO resources.
type KubernetesResource interface {
	conditions.Conditioner

	// Owner returns the ResourceReference of the owner, or nil if there is no owner
	Owner() *ResourceReference

	// TODO: I think we need this?
	// KnownOwner() *KnownResourceReference

	// AzureName returns the Azure name of the resource
	AzureName() string

	// GetType returns the type of the resource according to Azure. For example Microsoft.Resources/resourceGroups or
	// Microsoft.Network/networkSecurityGroups/securityRules
	GetType() string

	// Some types, but not all, have a corresponding:
	// 	SetAzureName(name string)
	// They do not if the name must be a fixed value (like 'default').

	// TODO: GetAPIVersion here?

	// TODO: I think we need this
	// SetStatus(status interface{})

	// GetSpec returns the specification of the resource
	GetSpec() ConvertibleSpec

	// GetStatus returns the current status of the resource
	GetStatus() ConvertibleStatus
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
