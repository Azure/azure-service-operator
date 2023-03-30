/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcommon

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

var SimpleExtensionResourceGroupVersion = schema.GroupVersion{Group: "microsoft.test.azure.com", Version: "v1apitest"}

type SimpleExtensionResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SimpleExtensionResourceSpec   `json:"spec,omitempty"`
	Status            SimpleExtensionResourceStatus `json:"status,omitempty"`
}

func (r *SimpleExtensionResource) SetStatus(status genruntime.ConvertibleStatus) error {
	r.Status = status.(SimpleExtensionResourceStatus)
	return nil
}

func (r *SimpleExtensionResource) GetSpec() genruntime.ConvertibleSpec {
	return &r.Spec
}

func (r *SimpleExtensionResource) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SimpleExtensionResourceStatus{}
}

func (r *SimpleExtensionResource) GetStatus() genruntime.ConvertibleStatus {
	return &r.Status
}

func (r *SimpleExtensionResource) GetAPIVersion() string {
	return "2020-01-01"
}

var _ admission.Defaulter = &SimpleExtensionResource{}

// Default defaults the Azure name of the resource to the Kubernetes name
func (r *SimpleExtensionResource) Default() {
	if r.Spec.AzureName == "" {
		r.Spec.AzureName = r.Name
	}
}

var _ conditions.Conditioner = &SimpleExtensionResource{}

// GetConditions returns the conditions of the resource
func (r *SimpleExtensionResource) GetConditions() conditions.Conditions {
	return r.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (r *SimpleExtensionResource) SetConditions(conditions conditions.Conditions) {
	r.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SimpleExtensionResource{}

// AzureName returns the Azure name of the resource
func (r *SimpleExtensionResource) AzureName() string {
	return r.Spec.AzureName
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (r *SimpleExtensionResource) Owner() *genruntime.ResourceReference {
	return nil
}

func (r *SimpleExtensionResource) GetType() string {
	return "Microsoft.SimpleExtension/simpleExtensions"
}

// GetResourceScope returns the scope of the resource
func (r *SimpleExtensionResource) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeExtension
}

type SimpleExtensionResourceSpec struct {
	AzureName string `json:"azureName,omitempty"`

	Owner genruntime.ResourceReference `json:"owner"`
}

var _ genruntime.ConvertibleSpec = &SimpleExtensionResourceSpec{}

func (s SimpleExtensionResourceSpec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	panic("not expected to be called in this test resource")
}

func (s SimpleExtensionResourceSpec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	panic("not expected to be called in this test resource")
}

type SimpleExtensionResourceStatus struct {
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SimpleExtensionResourceStatus{}

func (s SimpleExtensionResourceStatus) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	panic("not expected to be called in this test resource")
}

func (s SimpleExtensionResourceStatus) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	panic("not expected to be called in this test resource")
}

func (r *SimpleExtensionResource) DeepCopyObject() runtime.Object {
	// Note: This obviously isn't a copy. We don't care because this is a test resource and we just need this method implemented to satisfy our interface.
	// We don't want to run controller-gen as it's added overhead for what should be a simple test process.
	return r
}
