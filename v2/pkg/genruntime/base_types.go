/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type ResourceKind string

const (
	// ResourceKindNormal is a standard ARM resource.
	ResourceKindNormal = ResourceKind("normal") // TODO: Rename this to ResourceGroup?
	// ResourceKindExtension is an extension resource. Extension resources can have any resource as their parent.
	ResourceKindExtension = ResourceKind("extension")
	// ResourceKindTenant is an Azure resource rooted to the tenant (subscription, managementGroup, etc)
	ResourceKindTenant = ResourceKind("tenant")
)

// TODO: It's weird that this is isn't with the other annotations
// TODO: Should we move them all here (so they're exported?) Or shold we move them
// TODO: to serviceoperator-internal.azure.com to signify they are internal?
const (
	ResourceIDAnnotation = "serviceoperator.azure.com/resource-id"
)

// MetaObject represents an arbitrary ASO custom resource
type MetaObject interface {
	runtime.Object
	metav1.Object
	conditions.Conditioner
}

// ARMMetaObject represents an arbitrary ASO resource that is an ARM resource
type ARMMetaObject interface {
	MetaObject
	KubernetesResource
}

// ARMOwnedMetaObject represents an arbitrary ASO resource that is owned by an ARM resource
type ARMOwnedMetaObject interface {
	MetaObject
	ARMOwned
}

// AddAnnotation adds the specified annotation to the object.
// Empty string annotations are not allowed. Attempting to add an annotation with a value
// of empty string will result in the removal of that annotation.
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

// RemoveAnnotation removes the specified annotation from the object
func RemoveAnnotation(obj MetaObject, k string) {
	AddAnnotation(obj, k, "")
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

func NewARMResource(spec ARMResourceSpec, status ARMResourceStatus, id string) ARMResource {
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

// GetReadyCondition gets the ready condition from the object
func GetReadyCondition(obj conditions.Conditioner) *conditions.Condition {
	for _, c := range obj.GetConditions() {
		if c.Type == conditions.ConditionTypeReady {
			return &c
		}
	}

	return nil
}
