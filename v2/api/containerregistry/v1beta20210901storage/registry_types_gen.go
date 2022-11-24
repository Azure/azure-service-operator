// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210901storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=containerregistry.azure.com,resources=registries,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=containerregistry.azure.com,resources={registries/status,registries/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210901.Registry
// Generator information:
// - Generated from: /containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/containerregistry.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Registry_Spec   `json:"spec,omitempty"`
	Status            Registry_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Registry{}

// GetConditions returns the conditions of the resource
func (registry *Registry) GetConditions() conditions.Conditions {
	return registry.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (registry *Registry) SetConditions(conditions conditions.Conditions) {
	registry.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Registry{}

// AzureName returns the Azure name of the resource
func (registry *Registry) AzureName() string {
	return registry.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-09-01"
func (registry Registry) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (registry *Registry) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (registry *Registry) GetSpec() genruntime.ConvertibleSpec {
	return &registry.Spec
}

// GetStatus returns the status of this resource
func (registry *Registry) GetStatus() genruntime.ConvertibleStatus {
	return &registry.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerRegistry/registries"
func (registry *Registry) GetType() string {
	return "Microsoft.ContainerRegistry/registries"
}

// NewEmptyStatus returns a new empty (blank) status
func (registry *Registry) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Registry_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (registry *Registry) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(registry.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  registry.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (registry *Registry) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Registry_STATUS); ok {
		registry.Status = *st
		return nil
	}

	// Convert status to required version
	var st Registry_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	registry.Status = st
	return nil
}

// Hub marks that this Registry is the hub type for conversion
func (registry *Registry) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (registry *Registry) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: registry.Spec.OriginalVersion,
		Kind:    "Registry",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210901.Registry
// Generator information:
// - Generated from: /containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/containerregistry.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}

// Storage version of v1beta20210901.APIVersion
// +kubebuilder:validation:Enum={"2021-09-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-09-01")

// Storage version of v1beta20210901.Registry_Spec
type Registry_Spec struct {
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`

	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=5
	// +kubebuilder:validation:Pattern="^[a-zA-Z0-9]*$"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                string              `json:"azureName,omitempty"`
	DataEndpointEnabled      *bool               `json:"dataEndpointEnabled,omitempty"`
	Encryption               *EncryptionProperty `json:"encryption,omitempty"`
	Identity                 *IdentityProperties `json:"identity,omitempty"`
	Location                 *string             `json:"location,omitempty"`
	NetworkRuleBypassOptions *string             `json:"networkRuleBypassOptions,omitempty"`
	NetworkRuleSet           *NetworkRuleSet     `json:"networkRuleSet,omitempty"`
	OriginalVersion          string              `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner               *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Policies            *Policies                          `json:"policies,omitempty"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                            `json:"publicNetworkAccess,omitempty"`
	Sku                 *Sku                               `json:"sku,omitempty"`
	Tags                map[string]string                  `json:"tags,omitempty"`
	ZoneRedundancy      *string                            `json:"zoneRedundancy,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Registry_Spec{}

// ConvertSpecFrom populates our Registry_Spec from the provided source
func (registry *Registry_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == registry {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(registry)
}

// ConvertSpecTo populates the provided destination from our Registry_Spec
func (registry *Registry_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == registry {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(registry)
}

// Storage version of v1beta20210901.Registry_STATUS
type Registry_STATUS struct {
	AdminUserEnabled           *bool                              `json:"adminUserEnabled,omitempty"`
	Conditions                 []conditions.Condition             `json:"conditions,omitempty"`
	CreationDate               *string                            `json:"creationDate,omitempty"`
	DataEndpointEnabled        *bool                              `json:"dataEndpointEnabled,omitempty"`
	DataEndpointHostNames      []string                           `json:"dataEndpointHostNames,omitempty"`
	Encryption                 *EncryptionProperty_STATUS         `json:"encryption,omitempty"`
	Id                         *string                            `json:"id,omitempty"`
	Identity                   *IdentityProperties_STATUS         `json:"identity,omitempty"`
	Location                   *string                            `json:"location,omitempty"`
	LoginServer                *string                            `json:"loginServer,omitempty"`
	Name                       *string                            `json:"name,omitempty"`
	NetworkRuleBypassOptions   *string                            `json:"networkRuleBypassOptions,omitempty"`
	NetworkRuleSet             *NetworkRuleSet_STATUS             `json:"networkRuleSet,omitempty"`
	Policies                   *Policies_STATUS                   `json:"policies,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                            `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                            `json:"publicNetworkAccess,omitempty"`
	Sku                        *Sku_STATUS                        `json:"sku,omitempty"`
	Status                     *Status_STATUS                     `json:"status,omitempty"`
	SystemData                 *SystemData_STATUS                 `json:"systemData,omitempty"`
	Tags                       map[string]string                  `json:"tags,omitempty"`
	Type                       *string                            `json:"type,omitempty"`
	ZoneRedundancy             *string                            `json:"zoneRedundancy,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Registry_STATUS{}

// ConvertStatusFrom populates our Registry_STATUS from the provided source
func (registry *Registry_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == registry {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(registry)
}

// ConvertStatusTo populates the provided destination from our Registry_STATUS
func (registry *Registry_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == registry {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(registry)
}

// Storage version of v1beta20210901.EncryptionProperty
type EncryptionProperty struct {
	KeyVaultProperties *KeyVaultProperties    `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status             *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210901.EncryptionProperty_STATUS
type EncryptionProperty_STATUS struct {
	KeyVaultProperties *KeyVaultProperties_STATUS `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Status             *string                    `json:"status,omitempty"`
}

// Storage version of v1beta20210901.IdentityProperties
type IdentityProperties struct {
	PrincipalId            *string                           `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	TenantId               *string                           `json:"tenantId,omitempty"`
	Type                   *string                           `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserIdentityProperties `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20210901.IdentityProperties_STATUS
type IdentityProperties_STATUS struct {
	PrincipalId            *string                                  `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	TenantId               *string                                  `json:"tenantId,omitempty"`
	Type                   *string                                  `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserIdentityProperties_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20210901.NetworkRuleSet
type NetworkRuleSet struct {
	DefaultAction *string                `json:"defaultAction,omitempty"`
	IpRules       []IPRule               `json:"ipRules,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210901.NetworkRuleSet_STATUS
type NetworkRuleSet_STATUS struct {
	DefaultAction *string                `json:"defaultAction,omitempty"`
	IpRules       []IPRule_STATUS        `json:"ipRules,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210901.Policies
type Policies struct {
	ExportPolicy     *ExportPolicy          `json:"exportPolicy,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QuarantinePolicy *QuarantinePolicy      `json:"quarantinePolicy,omitempty"`
	RetentionPolicy  *RetentionPolicy       `json:"retentionPolicy,omitempty"`
	TrustPolicy      *TrustPolicy           `json:"trustPolicy,omitempty"`
}

// Storage version of v1beta20210901.Policies_STATUS
type Policies_STATUS struct {
	ExportPolicy     *ExportPolicy_STATUS     `json:"exportPolicy,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	QuarantinePolicy *QuarantinePolicy_STATUS `json:"quarantinePolicy,omitempty"`
	RetentionPolicy  *RetentionPolicy_STATUS  `json:"retentionPolicy,omitempty"`
	TrustPolicy      *TrustPolicy_STATUS      `json:"trustPolicy,omitempty"`
}

// Storage version of v1beta20210901.PrivateEndpointConnection_STATUS
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210901.Sku
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210901.Sku_STATUS
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20210901.Status_STATUS
type Status_STATUS struct {
	DisplayStatus *string                `json:"displayStatus,omitempty"`
	Message       *string                `json:"message,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Timestamp     *string                `json:"timestamp,omitempty"`
}

// Storage version of v1beta20210901.SystemData_STATUS
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210901.ExportPolicy
type ExportPolicy struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210901.ExportPolicy_STATUS
type ExportPolicy_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210901.IPRule
type IPRule struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20210901.IPRule_STATUS
type IPRule_STATUS struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20210901.KeyVaultProperties
type KeyVaultProperties struct {
	Identity      *string                `json:"identity,omitempty"`
	KeyIdentifier *string                `json:"keyIdentifier,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210901.KeyVaultProperties_STATUS
type KeyVaultProperties_STATUS struct {
	Identity                 *string                `json:"identity,omitempty"`
	KeyIdentifier            *string                `json:"keyIdentifier,omitempty"`
	KeyRotationEnabled       *bool                  `json:"keyRotationEnabled,omitempty"`
	LastKeyRotationTimestamp *string                `json:"lastKeyRotationTimestamp,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VersionedKeyIdentifier   *string                `json:"versionedKeyIdentifier,omitempty"`
}

// Storage version of v1beta20210901.QuarantinePolicy
type QuarantinePolicy struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210901.QuarantinePolicy_STATUS
type QuarantinePolicy_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210901.RetentionPolicy
type RetentionPolicy struct {
	Days        *int                   `json:"days,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210901.RetentionPolicy_STATUS
type RetentionPolicy_STATUS struct {
	Days            *int                   `json:"days,omitempty"`
	LastUpdatedTime *string                `json:"lastUpdatedTime,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210901.TrustPolicy
type TrustPolicy struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20210901.TrustPolicy_STATUS
type TrustPolicy_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20210901.UserIdentityProperties
type UserIdentityProperties struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210901.UserIdentityProperties_STATUS
type UserIdentityProperties_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}
