// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cdn.azure.com,resources=profiles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cdn.azure.com,resources={profiles/status,profiles/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210601.Profile
// Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.Cdn.json#/resourceDefinitions/profiles
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Profile_Spec   `json:"spec,omitempty"`
	Status            Profile_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Profile{}

// GetConditions returns the conditions of the resource
func (profile *Profile) GetConditions() conditions.Conditions {
	return profile.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (profile *Profile) SetConditions(conditions conditions.Conditions) {
	profile.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Profile{}

// AzureName returns the Azure name of the resource
func (profile *Profile) AzureName() string {
	return profile.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (profile Profile) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (profile *Profile) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (profile *Profile) GetSpec() genruntime.ConvertibleSpec {
	return &profile.Spec
}

// GetStatus returns the status of this resource
func (profile *Profile) GetStatus() genruntime.ConvertibleStatus {
	return &profile.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles"
func (profile *Profile) GetType() string {
	return "Microsoft.Cdn/profiles"
}

// NewEmptyStatus returns a new empty (blank) status
func (profile *Profile) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Profile_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (profile *Profile) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(profile.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  profile.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (profile *Profile) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Profile_STATUS); ok {
		profile.Status = *st
		return nil
	}

	// Convert status to required version
	var st Profile_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	profile.Status = st
	return nil
}

// Hub marks that this Profile is the hub type for conversion
func (profile *Profile) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (profile *Profile) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: profile.Spec.OriginalVersion,
		Kind:    "Profile",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210601.Profile
// Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.Cdn.json#/resourceDefinitions/profiles
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// Storage version of v1beta20210601.APIVersion
// +kubebuilder:validation:Enum={"2021-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-06-01")

// Storage version of v1beta20210601.Profile_Spec
type Profile_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                    string  `json:"azureName,omitempty"`
	Location                     *string `json:"location,omitempty"`
	OriginResponseTimeoutSeconds *int    `json:"originResponseTimeoutSeconds,omitempty"`
	OriginalVersion              string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *Sku                               `json:"sku,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Profile_Spec{}

// ConvertSpecFrom populates our Profile_Spec from the provided source
func (profile *Profile_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == profile {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(profile)
}

// ConvertSpecTo populates the provided destination from our Profile_Spec
func (profile *Profile_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == profile {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(profile)
}

// Storage version of v1beta20210601.Profile_STATUS
type Profile_STATUS struct {
	Conditions                   []conditions.Condition `json:"conditions,omitempty"`
	FrontDoorId                  *string                `json:"frontDoorId,omitempty"`
	Id                           *string                `json:"id,omitempty"`
	Kind                         *string                `json:"kind,omitempty"`
	Location                     *string                `json:"location,omitempty"`
	Name                         *string                `json:"name,omitempty"`
	OriginResponseTimeoutSeconds *int                   `json:"originResponseTimeoutSeconds,omitempty"`
	PropertyBag                  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                `json:"provisioningState,omitempty"`
	ResourceState                *string                `json:"resourceState,omitempty"`
	Sku                          *Sku_STATUS            `json:"sku,omitempty"`
	SystemData                   *SystemData_STATUS     `json:"systemData,omitempty"`
	Tags                         map[string]string      `json:"tags,omitempty"`
	Type                         *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Profile_STATUS{}

// ConvertStatusFrom populates our Profile_STATUS from the provided source
func (profile *Profile_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == profile {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(profile)
}

// ConvertStatusTo populates the provided destination from our Profile_STATUS
func (profile *Profile_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == profile {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(profile)
}

// Storage version of v1beta20210601.Sku
// Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.Cdn.json#/definitions/Sku
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.Sku_STATUS
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.SystemData_STATUS
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}
