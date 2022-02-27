/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"context"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type ResourceKind string

const (
	// ResourceKindNormal is a standard ARM resource.
	ResourceKindNormal = ResourceKind("normal")
	// ResourceKindExtension is an extension resource. Extension resources can have any resource as their parent.
	ResourceKindExtension = ResourceKind("extension")
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
	KubernetesResource
}

type Reconciler interface {
	Reconcile(ctx context.Context, log logr.Logger, obj MetaObject) (ctrl.Result, error)
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
func GetReadyCondition(obj MetaObject) *conditions.Condition {
	for _, c := range obj.GetConditions() {
		if c.Type == conditions.ConditionTypeReady {
			return &c
		}
	}

	return nil
}
