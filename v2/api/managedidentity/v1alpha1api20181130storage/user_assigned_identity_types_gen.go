// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=managedidentity.azure.com,resources=userassignedidentities,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=managedidentity.azure.com,resources={userassignedidentities/status,userassignedidentities/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20181130.UserAssignedIdentity
//Generated from: https://schema.management.azure.com/schemas/2018-11-30/Microsoft.ManagedIdentity.json#/resourceDefinitions/userAssignedIdentities
type UserAssignedIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserAssignedIdentities_Spec `json:"spec,omitempty"`
	Status            Identity_Status             `json:"status,omitempty"`
}

var _ conditions.Conditioner = &UserAssignedIdentity{}

// GetConditions returns the conditions of the resource
func (identity *UserAssignedIdentity) GetConditions() conditions.Conditions {
	return identity.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (identity *UserAssignedIdentity) SetConditions(conditions conditions.Conditions) {
	identity.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &UserAssignedIdentity{}

// AzureName returns the Azure name of the resource
func (identity *UserAssignedIdentity) AzureName() string {
	return identity.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-11-30"
func (identity UserAssignedIdentity) GetAPIVersion() string {
	return "2018-11-30"
}

// GetResourceKind returns the kind of the resource
func (identity *UserAssignedIdentity) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (identity *UserAssignedIdentity) GetSpec() genruntime.ConvertibleSpec {
	return &identity.Spec
}

// GetStatus returns the status of this resource
func (identity *UserAssignedIdentity) GetStatus() genruntime.ConvertibleStatus {
	return &identity.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities"
func (identity *UserAssignedIdentity) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities"
}

// NewEmptyStatus returns a new empty (blank) status
func (identity *UserAssignedIdentity) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Identity_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (identity *UserAssignedIdentity) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(identity.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  identity.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (identity *UserAssignedIdentity) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Identity_Status); ok {
		identity.Status = *st
		return nil
	}

	// Convert status to required version
	var st Identity_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	identity.Status = st
	return nil
}

// Hub marks that this UserAssignedIdentity is the hub type for conversion
func (identity *UserAssignedIdentity) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (identity *UserAssignedIdentity) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: identity.Spec.OriginalVersion,
		Kind:    "UserAssignedIdentity",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20181130.UserAssignedIdentity
//Generated from: https://schema.management.azure.com/schemas/2018-11-30/Microsoft.ManagedIdentity.json#/resourceDefinitions/userAssignedIdentities
type UserAssignedIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserAssignedIdentity `json:"items"`
}

//Storage version of v1alpha1api20181130.Identity_Status
type Identity_Status struct {
	ClientId    *string                `json:"clientId,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tags        map[string]string      `json:"tags,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Identity_Status{}

// ConvertStatusFrom populates our Identity_Status from the provided source
func (identity *Identity_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == identity {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(identity)
}

// ConvertStatusTo populates the provided destination from our Identity_Status
func (identity *Identity_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == identity {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(identity)
}

//Storage version of v1alpha1api20181130.UserAssignedIdentities_Spec
type UserAssignedIdentities_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName       string  `json:"azureName"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a resources.azure.com/ResourceGroup resource
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &UserAssignedIdentities_Spec{}

// ConvertSpecFrom populates our UserAssignedIdentities_Spec from the provided source
func (identities *UserAssignedIdentities_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == identities {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(identities)
}

// ConvertSpecTo populates the provided destination from our UserAssignedIdentities_Spec
func (identities *UserAssignedIdentities_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == identities {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(identities)
}

func init() {
	SchemeBuilder.Register(&UserAssignedIdentity{}, &UserAssignedIdentityList{})
}
