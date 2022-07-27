// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130storage

import (
	"fmt"
	v20181130s "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20181130storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1alpha1api20181130.UserAssignedIdentity
// Deprecated version of UserAssignedIdentity. Use v1beta20181130.UserAssignedIdentity instead
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

var _ conversion.Convertible = &UserAssignedIdentity{}

// ConvertFrom populates our UserAssignedIdentity from the provided hub UserAssignedIdentity
func (identity *UserAssignedIdentity) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20181130s.UserAssignedIdentity)
	if !ok {
		return fmt.Errorf("expected managedidentity/v1beta20181130storage/UserAssignedIdentity but received %T instead", hub)
	}

	return identity.AssignPropertiesFromUserAssignedIdentity(source)
}

// ConvertTo populates the provided hub UserAssignedIdentity from our UserAssignedIdentity
func (identity *UserAssignedIdentity) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20181130s.UserAssignedIdentity)
	if !ok {
		return fmt.Errorf("expected managedidentity/v1beta20181130storage/UserAssignedIdentity but received %T instead", hub)
	}

	return identity.AssignPropertiesToUserAssignedIdentity(destination)
}

var _ genruntime.KubernetesResource = &UserAssignedIdentity{}

// AzureName returns the Azure name of the resource
func (identity *UserAssignedIdentity) AzureName() string {
	return identity.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-11-30"
func (identity UserAssignedIdentity) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceScope returns the scope of the resource
func (identity *UserAssignedIdentity) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
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

// AssignPropertiesFromUserAssignedIdentity populates our UserAssignedIdentity from the provided source UserAssignedIdentity
func (identity *UserAssignedIdentity) AssignPropertiesFromUserAssignedIdentity(source *v20181130s.UserAssignedIdentity) error {

	// ObjectMeta
	identity.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec UserAssignedIdentities_Spec
	err := spec.AssignPropertiesFromUserAssignedIdentitiesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromUserAssignedIdentitiesSpec() to populate field Spec")
	}
	identity.Spec = spec

	// Status
	var status Identity_Status
	err = status.AssignPropertiesFromIdentityStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromIdentityStatus() to populate field Status")
	}
	identity.Status = status

	// No error
	return nil
}

// AssignPropertiesToUserAssignedIdentity populates the provided destination UserAssignedIdentity from our UserAssignedIdentity
func (identity *UserAssignedIdentity) AssignPropertiesToUserAssignedIdentity(destination *v20181130s.UserAssignedIdentity) error {

	// ObjectMeta
	destination.ObjectMeta = *identity.ObjectMeta.DeepCopy()

	// Spec
	var spec v20181130s.UserAssignedIdentities_Spec
	err := identity.Spec.AssignPropertiesToUserAssignedIdentitiesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToUserAssignedIdentitiesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20181130s.Identity_Status
	err = identity.Status.AssignPropertiesToIdentityStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToIdentityStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (identity *UserAssignedIdentity) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: identity.Spec.OriginalVersion,
		Kind:    "UserAssignedIdentity",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20181130.UserAssignedIdentity
// Deprecated version of UserAssignedIdentity. Use v1beta20181130.UserAssignedIdentity instead
type UserAssignedIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserAssignedIdentity `json:"items"`
}

// Storage version of v1alpha1api20181130.APIVersion
// Deprecated version of APIVersion. Use v1beta20181130.APIVersion instead
// +kubebuilder:validation:Enum={"2018-11-30"}
type APIVersion string

const APIVersionValue = APIVersion("2018-11-30")

// Storage version of v1alpha1api20181130.Identity_Status
// Deprecated version of Identity_Status. Use v1beta20181130.Identity_Status instead
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
	src, ok := source.(*v20181130s.Identity_Status)
	if ok {
		// Populate our instance from source
		return identity.AssignPropertiesFromIdentityStatus(src)
	}

	// Convert to an intermediate form
	src = &v20181130s.Identity_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = identity.AssignPropertiesFromIdentityStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Identity_Status
func (identity *Identity_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20181130s.Identity_Status)
	if ok {
		// Populate destination from our instance
		return identity.AssignPropertiesToIdentityStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20181130s.Identity_Status{}
	err := identity.AssignPropertiesToIdentityStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignPropertiesFromIdentityStatus populates our Identity_Status from the provided source Identity_Status
func (identity *Identity_Status) AssignPropertiesFromIdentityStatus(source *v20181130s.Identity_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ClientId
	identity.ClientId = genruntime.ClonePointerToString(source.ClientId)

	// Conditions
	identity.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	identity.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	identity.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	identity.Name = genruntime.ClonePointerToString(source.Name)

	// PrincipalId
	identity.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// Tags
	identity.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// TenantId
	identity.TenantId = genruntime.ClonePointerToString(source.TenantId)

	// Type
	identity.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		identity.PropertyBag = propertyBag
	} else {
		identity.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToIdentityStatus populates the provided destination Identity_Status from our Identity_Status
func (identity *Identity_Status) AssignPropertiesToIdentityStatus(destination *v20181130s.Identity_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(identity.PropertyBag)

	// ClientId
	destination.ClientId = genruntime.ClonePointerToString(identity.ClientId)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(identity.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(identity.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(identity.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(identity.Name)

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(identity.PrincipalId)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(identity.Tags)

	// TenantId
	destination.TenantId = genruntime.ClonePointerToString(identity.TenantId)

	// Type
	destination.Type = genruntime.ClonePointerToString(identity.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20181130.UserAssignedIdentities_Spec
type UserAssignedIdentities_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &UserAssignedIdentities_Spec{}

// ConvertSpecFrom populates our UserAssignedIdentities_Spec from the provided source
func (identities *UserAssignedIdentities_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20181130s.UserAssignedIdentities_Spec)
	if ok {
		// Populate our instance from source
		return identities.AssignPropertiesFromUserAssignedIdentitiesSpec(src)
	}

	// Convert to an intermediate form
	src = &v20181130s.UserAssignedIdentities_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = identities.AssignPropertiesFromUserAssignedIdentitiesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our UserAssignedIdentities_Spec
func (identities *UserAssignedIdentities_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20181130s.UserAssignedIdentities_Spec)
	if ok {
		// Populate destination from our instance
		return identities.AssignPropertiesToUserAssignedIdentitiesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20181130s.UserAssignedIdentities_Spec{}
	err := identities.AssignPropertiesToUserAssignedIdentitiesSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromUserAssignedIdentitiesSpec populates our UserAssignedIdentities_Spec from the provided source UserAssignedIdentities_Spec
func (identities *UserAssignedIdentities_Spec) AssignPropertiesFromUserAssignedIdentitiesSpec(source *v20181130s.UserAssignedIdentities_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	identities.AzureName = source.AzureName

	// Location
	identities.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	identities.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		identities.Owner = &owner
	} else {
		identities.Owner = nil
	}

	// Tags
	identities.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		identities.PropertyBag = propertyBag
	} else {
		identities.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToUserAssignedIdentitiesSpec populates the provided destination UserAssignedIdentities_Spec from our UserAssignedIdentities_Spec
func (identities *UserAssignedIdentities_Spec) AssignPropertiesToUserAssignedIdentitiesSpec(destination *v20181130s.UserAssignedIdentities_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(identities.PropertyBag)

	// AzureName
	destination.AzureName = identities.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(identities.Location)

	// OriginalVersion
	destination.OriginalVersion = identities.OriginalVersion

	// Owner
	if identities.Owner != nil {
		owner := identities.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(identities.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&UserAssignedIdentity{}, &UserAssignedIdentityList{})
}
