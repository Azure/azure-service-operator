// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// +kubebuilder:rbac:groups=entra.azure.com,resources=users,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=entra.azure.com,resources={users/status,users/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// +kubebuilder:storageversion
// SecurityGroup is an Entra Security Group.
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec,omitempty"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SecurityGroup{}

// GetConditions returns the conditions of the resource
func (user *SecurityGroup) GetConditions() conditions.Conditions {
	return user.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (user *SecurityGroup) SetConditions(conditions conditions.Conditions) {
	user.Status.Conditions = conditions
}

// +kubebuilder:webhook:path=/mutate-entra-azure-com-v1-user,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=entra.azure.com,resources=users,verbs=create;update,versions=v1,name=default.v1.securitygroup.entra.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &SecurityGroup{}

// Default applies defaults to the FlexibleServer resource
func (user *SecurityGroup) Default() {
	user.defaultImpl()
	var temp interface{} = user
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (user *SecurityGroup) defaultAzureName() {
	if user.Spec.AzureName == "" {
		user.Spec.AzureName = user.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServer resource
func (user *SecurityGroup) defaultImpl() { user.defaultAzureName() }

var _ genruntime.ARMOwned = &SecurityGroup{}

// AzureName returns the Azure name of the resource
func (user *SecurityGroup) AzureName() string {
	return user.Spec.AzureName
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (user *SecurityGroup) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(user.Spec)
	return user.Spec.Owner.AsResourceReference(group, kind)
}

// +kubebuilder:webhook:path=/validate-entra-azure-com-v1-user,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=entra.azure.com,resources=users,verbs=create;update,versions=v1,name=validate.v1.securitygroup.entra.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &SecurityGroup{}

// ValidateCreate validates the creation of the resource
func (user *SecurityGroup) ValidateCreate() (admission.Warnings, error) {
	validations := user.createValidations()
	var temp interface{} = user
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (user *SecurityGroup) ValidateDelete() (admission.Warnings, error) {
	validations := user.deleteValidations()
	var temp interface{} = user
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (user *SecurityGroup) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := user.updateValidations()
	var temp interface{} = user
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (user *SecurityGroup) createValidations() []func() (admission.Warnings, error) {
	return nil
}

// deleteValidations validates the deletion of the resource
func (user *SecurityGroup) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (user *SecurityGroup) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return nil
}

var _ conversion.Hub = &SecurityGroup{}

// Hub marks that this userSpec is the hub type for conversion
func (user *SecurityGroup) Hub() {}

// +kubebuilder:object:root=true
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

type SecurityGroupSpec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner *genruntime.KubernetesOwnerReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
}

// OriginalVersion returns the original API version used to create the resource.
func (userSpec *SecurityGroupSpec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (userSpec *SecurityGroupSpec) SetAzureName(azureName string) {
	userSpec.AzureName = azureName
}

type SecurityGroupStatus struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
