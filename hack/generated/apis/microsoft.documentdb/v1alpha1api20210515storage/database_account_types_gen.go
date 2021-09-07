// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.DatabaseAccount
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts
type DatabaseAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_Spec            `json:"spec,omitempty"`
	Status            DatabaseAccountGetResults_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DatabaseAccount{}

// GetConditions returns the conditions of the resource
func (databaseAccount *DatabaseAccount) GetConditions() conditions.Conditions {
	return databaseAccount.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (databaseAccount *DatabaseAccount) SetConditions(conditions conditions.Conditions) {
	databaseAccount.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DatabaseAccount{}

// AzureName returns the Azure name of the resource
func (databaseAccount *DatabaseAccount) AzureName() string {
	return databaseAccount.Spec.AzureName
}

// GetSpec returns the specification of this resource
func (databaseAccount *DatabaseAccount) GetSpec() genruntime.ConvertibleSpec {
	return &databaseAccount.Spec
}

// GetStatus returns the status of this resource
func (databaseAccount *DatabaseAccount) GetStatus() genruntime.ConvertibleStatus {
	return &databaseAccount.Status
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (databaseAccount *DatabaseAccount) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(databaseAccount.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: databaseAccount.Namespace, Name: databaseAccount.Spec.Owner.Name}
}

// SetStatus sets the status of this resource
func (databaseAccount *DatabaseAccount) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccountGetResults_Status); ok {
		databaseAccount.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccountGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	databaseAccount.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (databaseAccount *DatabaseAccount) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: databaseAccount.Spec.OriginalVersion,
		Kind:    "DatabaseAccount",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.DatabaseAccount
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts
type DatabaseAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseAccount `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountGetResults_Status
//Generated from:
type DatabaseAccountGetResults_Status struct {
	AnalyticalStorageConfiguration     *AnalyticalStorageConfiguration_Status                 `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                      *ApiProperties_Status                                  `json:"apiProperties,omitempty"`
	BackupPolicy                       *BackupPolicy_Status                                   `json:"backupPolicy,omitempty"`
	Capabilities                       []Capability_Status                                    `json:"capabilities,omitempty"`
	Conditions                         []conditions.Condition                                 `json:"conditions,omitempty"`
	ConnectorOffer                     *string                                                `json:"connectorOffer,omitempty"`
	ConsistencyPolicy                  *ConsistencyPolicy_Status                              `json:"consistencyPolicy,omitempty"`
	Cors                               []CorsPolicy_Status                                    `json:"cors,omitempty"`
	DatabaseAccountOfferType           *string                                                `json:"databaseAccountOfferType,omitempty"`
	DefaultIdentity                    *string                                                `json:"defaultIdentity,omitempty"`
	DisableKeyBasedMetadataWriteAccess *bool                                                  `json:"disableKeyBasedMetadataWriteAccess,omitempty"`
	DocumentEndpoint                   *string                                                `json:"documentEndpoint,omitempty"`
	EnableAnalyticalStorage            *bool                                                  `json:"enableAnalyticalStorage,omitempty"`
	EnableAutomaticFailover            *bool                                                  `json:"enableAutomaticFailover,omitempty"`
	EnableCassandraConnector           *bool                                                  `json:"enableCassandraConnector,omitempty"`
	EnableFreeTier                     *bool                                                  `json:"enableFreeTier,omitempty"`
	EnableMultipleWriteLocations       *bool                                                  `json:"enableMultipleWriteLocations,omitempty"`
	FailoverPolicies                   []FailoverPolicy_Status                                `json:"failoverPolicies,omitempty"`
	Id                                 *string                                                `json:"id,omitempty"`
	Identity                           *ManagedServiceIdentity_Status                         `json:"identity,omitempty"`
	IpRules                            []IpAddressOrRange_Status                              `json:"ipRules,omitempty"`
	IsVirtualNetworkFilterEnabled      *bool                                                  `json:"isVirtualNetworkFilterEnabled,omitempty"`
	KeyVaultKeyUri                     *string                                                `json:"keyVaultKeyUri,omitempty"`
	Kind                               *string                                                `json:"kind,omitempty"`
	Location                           *string                                                `json:"location,omitempty"`
	Locations                          []Location_Status                                      `json:"locations,omitempty"`
	Name                               *string                                                `json:"name,omitempty"`
	NetworkAclBypass                   *string                                                `json:"networkAclBypass,omitempty"`
	NetworkAclBypassResourceIds        []string                                               `json:"networkAclBypassResourceIds,omitempty"`
	PrivateEndpointConnections         []PrivateEndpointConnection_Status_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                        genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	ProvisioningState                  *string                                                `json:"provisioningState,omitempty"`
	PublicNetworkAccess                *string                                                `json:"publicNetworkAccess,omitempty"`
	ReadLocations                      []Location_Status                                      `json:"readLocations,omitempty"`
	Tags                               map[string]string                                      `json:"tags,omitempty"`
	Type                               *string                                                `json:"type,omitempty"`
	VirtualNetworkRules                []VirtualNetworkRule_Status                            `json:"virtualNetworkRules,omitempty"`
	WriteLocations                     []Location_Status                                      `json:"writeLocations,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccountGetResults_Status{}

// ConvertStatusFrom populates our DatabaseAccountGetResults_Status from the provided source
func (databaseAccountGetResultsStatus *DatabaseAccountGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == databaseAccountGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(databaseAccountGetResultsStatus)
}

// ConvertStatusTo populates the provided destination from our DatabaseAccountGetResults_Status
func (databaseAccountGetResultsStatus *DatabaseAccountGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == databaseAccountGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(databaseAccountGetResultsStatus)
}

//Storage version of v1alpha1api20210515.DatabaseAccounts_Spec
type DatabaseAccounts_Spec struct {
	AnalyticalStorageConfiguration *AnalyticalStorageConfiguration `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                  *ApiProperties                  `json:"apiProperties,omitempty"`

	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=3
	// +kubebuilder:validation:Pattern="^[a-z0-9]+(-[a-z0-9]+)*"
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                          string                  `json:"azureName"`
	BackupPolicy                       *BackupPolicy           `json:"backupPolicy,omitempty"`
	Capabilities                       []Capability            `json:"capabilities,omitempty"`
	ConnectorOffer                     *string                 `json:"connectorOffer,omitempty"`
	ConsistencyPolicy                  *ConsistencyPolicy      `json:"consistencyPolicy,omitempty"`
	Cors                               []CorsPolicy            `json:"cors,omitempty"`
	DatabaseAccountOfferType           *string                 `json:"databaseAccountOfferType,omitempty"`
	DefaultIdentity                    *string                 `json:"defaultIdentity,omitempty"`
	DisableKeyBasedMetadataWriteAccess *bool                   `json:"disableKeyBasedMetadataWriteAccess,omitempty"`
	EnableAnalyticalStorage            *bool                   `json:"enableAnalyticalStorage,omitempty"`
	EnableAutomaticFailover            *bool                   `json:"enableAutomaticFailover,omitempty"`
	EnableCassandraConnector           *bool                   `json:"enableCassandraConnector,omitempty"`
	EnableFreeTier                     *bool                   `json:"enableFreeTier,omitempty"`
	EnableMultipleWriteLocations       *bool                   `json:"enableMultipleWriteLocations,omitempty"`
	Identity                           *ManagedServiceIdentity `json:"identity,omitempty"`
	IpRules                            []IpAddressOrRange      `json:"ipRules,omitempty"`
	IsVirtualNetworkFilterEnabled      *bool                   `json:"isVirtualNetworkFilterEnabled,omitempty"`
	KeyVaultKeyUri                     *string                 `json:"keyVaultKeyUri,omitempty"`
	Kind                               *string                 `json:"kind,omitempty"`
	Location                           *string                 `json:"location,omitempty"`
	Locations                          []Location              `json:"locations,omitempty"`
	NetworkAclBypass                   *string                 `json:"networkAclBypass,omitempty"`
	NetworkAclBypassResourceIds        []string                `json:"networkAclBypassResourceIds,omitempty"`
	OriginalVersion                    string                  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner               genruntime.KnownResourceReference `group:"microsoft.resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                           `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                 `json:"tags,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule              `json:"virtualNetworkRules,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_Spec from the provided source
func (databaseAccountsSpec *DatabaseAccounts_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databaseAccountsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databaseAccountsSpec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_Spec
func (databaseAccountsSpec *DatabaseAccounts_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databaseAccountsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databaseAccountsSpec)
}

//Storage version of v1alpha1api20210515.AnalyticalStorageConfiguration
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AnalyticalStorageConfiguration
type AnalyticalStorageConfiguration struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SchemaType  *string                `json:"schemaType,omitempty"`
}

//Storage version of v1alpha1api20210515.AnalyticalStorageConfiguration_Status
//Generated from:
type AnalyticalStorageConfiguration_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SchemaType  *string                `json:"schemaType,omitempty"`
}

//Storage version of v1alpha1api20210515.ApiProperties
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ApiProperties
type ApiProperties struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServerVersion *string                `json:"serverVersion,omitempty"`
}

//Storage version of v1alpha1api20210515.ApiProperties_Status
//Generated from:
type ApiProperties_Status struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServerVersion *string                `json:"serverVersion,omitempty"`
}

//Storage version of v1alpha1api20210515.BackupPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/BackupPolicy
type BackupPolicy struct {
	ContinuousModeBackupPolicy *ContinuousModeBackupPolicy `json:"continuousModeBackupPolicy,omitempty"`
	PeriodicModeBackupPolicy   *PeriodicModeBackupPolicy   `json:"periodicModeBackupPolicy,omitempty"`
	PropertyBag                genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.BackupPolicy_Status
//Generated from:
type BackupPolicy_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210515.Capability
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Capability
type Capability struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.Capability_Status
//Generated from:
type Capability_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ConsistencyPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConsistencyPolicy
type ConsistencyPolicy struct {
	DefaultConsistencyLevel *string                `json:"defaultConsistencyLevel,omitempty"`
	MaxIntervalInSeconds    *int                   `json:"maxIntervalInSeconds,omitempty"`
	MaxStalenessPrefix      *int                   `json:"maxStalenessPrefix,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ConsistencyPolicy_Status
//Generated from:
type ConsistencyPolicy_Status struct {
	DefaultConsistencyLevel *string                `json:"defaultConsistencyLevel,omitempty"`
	MaxIntervalInSeconds    *int                   `json:"maxIntervalInSeconds,omitempty"`
	MaxStalenessPrefix      *int                   `json:"maxStalenessPrefix,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.CorsPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CorsPolicy
type CorsPolicy struct {
	AllowedHeaders  *string                `json:"allowedHeaders,omitempty"`
	AllowedMethods  *string                `json:"allowedMethods,omitempty"`
	AllowedOrigins  *string                `json:"allowedOrigins,omitempty"`
	ExposedHeaders  *string                `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.CorsPolicy_Status
//Generated from:
type CorsPolicy_Status struct {
	AllowedHeaders  *string                `json:"allowedHeaders,omitempty"`
	AllowedMethods  *string                `json:"allowedMethods,omitempty"`
	AllowedOrigins  *string                `json:"allowedOrigins,omitempty"`
	ExposedHeaders  *string                `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.FailoverPolicy_Status
//Generated from:
type FailoverPolicy_Status struct {
	FailoverPriority *int                   `json:"failoverPriority,omitempty"`
	Id               *string                `json:"id,omitempty"`
	LocationName     *string                `json:"locationName,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.IpAddressOrRange
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IpAddressOrRange
type IpAddressOrRange struct {
	IpAddressOrRange *string                `json:"ipAddressOrRange,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.IpAddressOrRange_Status
//Generated from:
type IpAddressOrRange_Status struct {
	IpAddressOrRange *string                `json:"ipAddressOrRange,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.Location
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Location
type Location struct {
	FailoverPriority *int                   `json:"failoverPriority,omitempty"`
	IsZoneRedundant  *bool                  `json:"isZoneRedundant,omitempty"`
	LocationName     *string                `json:"locationName,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.Location_Status
//Generated from:
type Location_Status struct {
	DocumentEndpoint  *string                `json:"documentEndpoint,omitempty"`
	FailoverPriority  *int                   `json:"failoverPriority,omitempty"`
	Id                *string                `json:"id,omitempty"`
	IsZoneRedundant   *bool                  `json:"isZoneRedundant,omitempty"`
	LocationName      *string                `json:"locationName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
}

//Storage version of v1alpha1api20210515.ManagedServiceIdentity
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ManagedServiceIdentity
type ManagedServiceIdentity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210515.ManagedServiceIdentity_Status
//Generated from:
type ManagedServiceIdentity_Status struct {
	PrincipalId            *string                                                         `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                                          `json:"$propertyBag,omitempty"`
	TenantId               *string                                                         `json:"tenantId,omitempty"`
	Type                   *string                                                         `json:"type,omitempty"`
	UserAssignedIdentities map[string]ManagedServiceIdentity_Status_UserAssignedIdentities `json:"userAssignedIdentities,omitempty"`
}

//Storage version of v1alpha1api20210515.PrivateEndpointConnection_Status_SubResourceEmbedded
//Generated from:
type PrivateEndpointConnection_Status_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.VirtualNetworkRule
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/VirtualNetworkRule
type VirtualNetworkRule struct {
	IgnoreMissingVNetServiceEndpoint *bool                  `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//Reference: Resource ID of a subnet, for example:
	///subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20210515.VirtualNetworkRule_Status
//Generated from:
type VirtualNetworkRule_Status struct {
	Id                               *string                `json:"id,omitempty"`
	IgnoreMissingVNetServiceEndpoint *bool                  `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ContinuousModeBackupPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContinuousModeBackupPolicy
type ContinuousModeBackupPolicy struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210515.ManagedServiceIdentity_Status_UserAssignedIdentities
type ManagedServiceIdentity_Status_UserAssignedIdentities struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.PeriodicModeBackupPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeBackupPolicy
type PeriodicModeBackupPolicy struct {
	PeriodicModeProperties *PeriodicModeProperties `json:"periodicModeProperties,omitempty"`
	PropertyBag            genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	Type                   *string                 `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210515.PeriodicModeProperties
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeProperties
type PeriodicModeProperties struct {
	BackupIntervalInMinutes        *int                   `json:"backupIntervalInMinutes,omitempty"`
	BackupRetentionIntervalInHours *int                   `json:"backupRetentionIntervalInHours,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DatabaseAccount{}, &DatabaseAccountList{})
}
