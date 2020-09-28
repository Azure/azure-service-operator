/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// MetaObject represents an arbitrary k8s-infra custom resource
type MetaObject interface {
	runtime.Object
	metav1.Object
	KubernetesResource
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

	// TODO: GetAPIVersion here?

	// TODO: I think we need this
	// SetStatus(status interface{})
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

// DeployableResource represents a resource which can be deployed to Azure. In addition
// to specification information, this includes the target resource group of the resource
type DeployableResource struct {
	resourceGroup string
	spec          ArmResourceSpec
}

func NewDeployableResource(resourceGroup string, spec ArmResourceSpec) *DeployableResource {
	return &DeployableResource{
		resourceGroup: resourceGroup,
		spec:          spec,
	}
}

func (r *DeployableResource) ResourceGroup() string {
	return r.resourceGroup
}

func (r *DeployableResource) Spec() ArmResourceSpec {
	return r.spec
}
