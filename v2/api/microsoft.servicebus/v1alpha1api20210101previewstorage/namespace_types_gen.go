// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210101preview.Namespace
//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces
type Namespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Spec    `json:"spec,omitempty"`
	Status            SBNamespace_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Namespace{}

// GetConditions returns the conditions of the resource
func (namespace *Namespace) GetConditions() conditions.Conditions {
	return namespace.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (namespace *Namespace) SetConditions(conditions conditions.Conditions) {
	namespace.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Namespace{}

// AzureName returns the Azure name of the resource
func (namespace *Namespace) AzureName() string {
	return namespace.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (namespace *Namespace) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (namespace *Namespace) GetSpec() genruntime.ConvertibleSpec {
	return &namespace.Spec
}

// GetStatus returns the status of this resource
func (namespace *Namespace) GetStatus() genruntime.ConvertibleStatus {
	return &namespace.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces"
func (namespace *Namespace) GetType() string {
	return "Microsoft.ServiceBus/namespaces"
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (namespace *Namespace) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(namespace.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: namespace.Namespace, Name: namespace.Spec.Owner.Name}
}

// SetStatus sets the status of this resource
func (namespace *Namespace) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SBNamespace_Status); ok {
		namespace.Status = *st
		return nil
	}

	// Convert status to required version
	var st SBNamespace_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	namespace.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (namespace *Namespace) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: namespace.Spec.OriginalVersion,
		Kind:    "Namespace",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210101preview.Namespace
//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces
type NamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Namespace `json:"items"`
}

//Storage version of v1alpha1api20210101preview.Namespaces_Spec
type Namespaces_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string      `json:"azureName"`
	Encryption      *Encryption `json:"encryption,omitempty"`
	Identity        *Identity   `json:"identity,omitempty"`
	Location        *string     `json:"location,omitempty"`
	OriginalVersion string      `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner         genruntime.KnownResourceReference `group:"microsoft.resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag   genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Sku           *SBSku                            `json:"sku,omitempty"`
	Tags          map[string]string                 `json:"tags,omitempty"`
	ZoneRedundant *bool                             `json:"zoneRedundant,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Namespaces_Spec{}

// ConvertSpecFrom populates our Namespaces_Spec from the provided source
func (namespacesSpec *Namespaces_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == namespacesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(namespacesSpec)
}

// ConvertSpecTo populates the provided destination from our Namespaces_Spec
func (namespacesSpec *Namespaces_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == namespacesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(namespacesSpec)
}

//Storage version of v1alpha1api20210101preview.SBNamespace_Status
//Generated from:
type SBNamespace_Status struct {
	Conditions                 []conditions.Condition                                 `json:"conditions,omitempty"`
	CreatedAt                  *string                                                `json:"createdAt,omitempty"`
	Encryption                 *Encryption_Status                                     `json:"encryption,omitempty"`
	Id                         *string                                                `json:"id,omitempty"`
	Identity                   *Identity_Status                                       `json:"identity,omitempty"`
	Location                   *string                                                `json:"location,omitempty"`
	MetricId                   *string                                                `json:"metricId,omitempty"`
	Name                       *string                                                `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                `json:"provisioningState,omitempty"`
	ServiceBusEndpoint         *string                                                `json:"serviceBusEndpoint,omitempty"`
	Sku                        *SBSku_Status                                          `json:"sku,omitempty"`
	Status                     *string                                                `json:"status,omitempty"`
	SystemData                 *SystemData_Status                                     `json:"systemData,omitempty"`
	Tags                       map[string]string                                      `json:"tags,omitempty"`
	Type                       *string                                                `json:"type,omitempty"`
	UpdatedAt                  *string                                                `json:"updatedAt,omitempty"`
	ZoneRedundant              *bool                                                  `json:"zoneRedundant,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SBNamespace_Status{}

// ConvertStatusFrom populates our SBNamespace_Status from the provided source
func (sbNamespaceStatus *SBNamespace_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == sbNamespaceStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(sbNamespaceStatus)
}

// ConvertStatusTo populates the provided destination from our SBNamespace_Status
func (sbNamespaceStatus *SBNamespace_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == sbNamespaceStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(sbNamespaceStatus)
}

//Storage version of v1alpha1api20210101preview.Encryption
//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Encryption
type Encryption struct {
	KeySource                       *string                `json:"keySource,omitempty"`
	KeyVaultProperties              []KeyVaultProperties   `json:"keyVaultProperties,omitempty"`
	PropertyBag                     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequireInfrastructureEncryption *bool                  `json:"requireInfrastructureEncryption,omitempty"`
}

//Storage version of v1alpha1api20210101preview.Encryption_Status
//Generated from:
type Encryption_Status struct {
	KeySource                       *string                     `json:"keySource,omitempty"`
	KeyVaultProperties              []KeyVaultProperties_Status `json:"keyVaultProperties,omitempty"`
	PropertyBag                     genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequireInfrastructureEncryption *bool                       `json:"requireInfrastructureEncryption,omitempty"`
}

//Storage version of v1alpha1api20210101preview.Identity
//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Identity
type Identity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210101preview.Identity_Status
//Generated from:
type Identity_Status struct {
	PrincipalId            *string                           `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	TenantId               *string                           `json:"tenantId,omitempty"`
	Type                   *string                           `json:"type,omitempty"`
	UserAssignedIdentities map[string]DictionaryValue_Status `json:"userAssignedIdentities,omitempty"`
}

//Storage version of v1alpha1api20210101preview.PrivateEndpointConnection_Status_SubResourceEmbedded
//Generated from:
type PrivateEndpointConnection_Status_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_Status     `json:"systemData,omitempty"`
}

//Storage version of v1alpha1api20210101preview.SBSku
//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBSku
type SBSku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20210101preview.SBSku_Status
//Generated from:
type SBSku_Status struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20210101preview.SystemData_Status
//Generated from:
type SystemData_Status struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210101preview.DictionaryValue_Status
//Generated from:
type DictionaryValue_Status struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210101preview.KeyVaultProperties
//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/KeyVaultProperties
type KeyVaultProperties struct {
	Identity    *UserAssignedIdentityProperties `json:"identity,omitempty"`
	KeyName     *string                         `json:"keyName,omitempty"`
	KeyVaultUri *string                         `json:"keyVaultUri,omitempty"`
	KeyVersion  *string                         `json:"keyVersion,omitempty"`
	PropertyBag genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210101preview.KeyVaultProperties_Status
//Generated from:
type KeyVaultProperties_Status struct {
	Identity    *UserAssignedIdentityProperties_Status `json:"identity,omitempty"`
	KeyName     *string                                `json:"keyName,omitempty"`
	KeyVaultUri *string                                `json:"keyVaultUri,omitempty"`
	KeyVersion  *string                                `json:"keyVersion,omitempty"`
	PropertyBag genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210101preview.UserAssignedIdentityProperties
//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/UserAssignedIdentityProperties
type UserAssignedIdentityProperties struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//UserAssignedIdentityReference: ARM ID of user Identity selected for encryption
	UserAssignedIdentityReference *genruntime.ResourceReference `armReference:"UserAssignedIdentity" json:"userAssignedIdentityReference,omitempty"`
}

//Storage version of v1alpha1api20210101preview.UserAssignedIdentityProperties_Status
//Generated from:
type UserAssignedIdentityProperties_Status struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Namespace{}, &NamespaceList{})
}
