// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccounts/status,storageaccounts/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210401.StorageAccount
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts
type StorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccounts_Spec  `json:"spec,omitempty"`
	Status            StorageAccount_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccount{}

// GetConditions returns the conditions of the resource
func (account *StorageAccount) GetConditions() conditions.Conditions {
	return account.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (account *StorageAccount) SetConditions(conditions conditions.Conditions) {
	account.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &StorageAccount{}

// AzureName returns the Azure name of the resource
func (account *StorageAccount) AzureName() string {
	return account.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (account StorageAccount) GetAPIVersion() string {
	return "2021-04-01"
}

// GetResourceKind returns the kind of the resource
func (account *StorageAccount) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (account *StorageAccount) GetSpec() genruntime.ConvertibleSpec {
	return &account.Spec
}

// GetStatus returns the status of this resource
func (account *StorageAccount) GetStatus() genruntime.ConvertibleStatus {
	return &account.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts"
func (account *StorageAccount) GetType() string {
	return "Microsoft.Storage/storageAccounts"
}

// NewEmptyStatus returns a new empty (blank) status
func (account *StorageAccount) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccount_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (account *StorageAccount) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(account.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  account.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (account *StorageAccount) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccount_Status); ok {
		account.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccount_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	account.Status = st
	return nil
}

// Hub marks that this StorageAccount is the hub type for conversion
func (account *StorageAccount) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (account *StorageAccount) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: account.Spec.OriginalVersion,
		Kind:    "StorageAccount",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210401.StorageAccount
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts
type StorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccount `json:"items"`
}

//Storage version of v1alpha1api20210401.StorageAccount_Status
type StorageAccount_Status struct {
	AccessTier                            *string                                                `json:"accessTier,omitempty"`
	AllowBlobPublicAccess                 *bool                                                  `json:"allowBlobPublicAccess,omitempty"`
	AllowCrossTenantReplication           *bool                                                  `json:"allowCrossTenantReplication,omitempty"`
	AllowSharedKeyAccess                  *bool                                                  `json:"allowSharedKeyAccess,omitempty"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication_Status          `json:"azureFilesIdentityBasedAuthentication,omitempty"`
	BlobRestoreStatus                     *BlobRestoreStatus_Status                              `json:"blobRestoreStatus,omitempty"`
	Conditions                            []conditions.Condition                                 `json:"conditions,omitempty"`
	CreationTime                          *string                                                `json:"creationTime,omitempty"`
	CustomDomain                          *CustomDomain_Status                                   `json:"customDomain,omitempty"`
	Encryption                            *Encryption_Status                                     `json:"encryption,omitempty"`
	ExtendedLocation                      *ExtendedLocation_Status                               `json:"extendedLocation,omitempty"`
	FailoverInProgress                    *bool                                                  `json:"failoverInProgress,omitempty"`
	GeoReplicationStats                   *GeoReplicationStats_Status                            `json:"geoReplicationStats,omitempty"`
	Id                                    *string                                                `json:"id,omitempty"`
	Identity                              *Identity_Status                                       `json:"identity,omitempty"`
	IsHnsEnabled                          *bool                                                  `json:"isHnsEnabled,omitempty"`
	IsNfsV3Enabled                        *bool                                                  `json:"isNfsV3Enabled,omitempty"`
	KeyCreationTime                       *KeyCreationTime_Status                                `json:"keyCreationTime,omitempty"`
	KeyPolicy                             *KeyPolicy_Status                                      `json:"keyPolicy,omitempty"`
	Kind                                  *string                                                `json:"kind,omitempty"`
	LargeFileSharesState                  *string                                                `json:"largeFileSharesState,omitempty"`
	LastGeoFailoverTime                   *string                                                `json:"lastGeoFailoverTime,omitempty"`
	Location                              *string                                                `json:"location,omitempty"`
	MinimumTlsVersion                     *string                                                `json:"minimumTlsVersion,omitempty"`
	Name                                  *string                                                `json:"name,omitempty"`
	NetworkAcls                           *NetworkRuleSet_Status                                 `json:"networkAcls,omitempty"`
	PrimaryEndpoints                      *Endpoints_Status                                      `json:"primaryEndpoints,omitempty"`
	PrimaryLocation                       *string                                                `json:"primaryLocation,omitempty"`
	PrivateEndpointConnections            []PrivateEndpointConnection_Status_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                           genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	ProvisioningState                     *string                                                `json:"provisioningState,omitempty"`
	RoutingPreference                     *RoutingPreference_Status                              `json:"routingPreference,omitempty"`
	SasPolicy                             *SasPolicy_Status                                      `json:"sasPolicy,omitempty"`
	SecondaryEndpoints                    *Endpoints_Status                                      `json:"secondaryEndpoints,omitempty"`
	SecondaryLocation                     *string                                                `json:"secondaryLocation,omitempty"`
	Sku                                   *Sku_Status                                            `json:"sku,omitempty"`
	StatusOfPrimary                       *string                                                `json:"statusOfPrimary,omitempty"`
	StatusOfSecondary                     *string                                                `json:"statusOfSecondary,omitempty"`
	SupportsHttpsTrafficOnly              *bool                                                  `json:"supportsHttpsTrafficOnly,omitempty"`
	Tags                                  map[string]string                                      `json:"tags,omitempty"`
	Type                                  *string                                                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccount_Status{}

// ConvertStatusFrom populates our StorageAccount_Status from the provided source
func (account *StorageAccount_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(account)
}

// ConvertStatusTo populates the provided destination from our StorageAccount_Status
func (account *StorageAccount_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(account)
}

//Storage version of v1alpha1api20210401.StorageAccounts_Spec
type StorageAccounts_Spec struct {
	AccessTier                            *string                                `json:"accessTier,omitempty"`
	AllowBlobPublicAccess                 *bool                                  `json:"allowBlobPublicAccess,omitempty"`
	AllowCrossTenantReplication           *bool                                  `json:"allowCrossTenantReplication,omitempty"`
	AllowSharedKeyAccess                  *bool                                  `json:"allowSharedKeyAccess,omitempty"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication `json:"azureFilesIdentityBasedAuthentication,omitempty"`

	// +kubebuilder:validation:MaxLength=24
	// +kubebuilder:validation:MinLength=3
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName            string                      `json:"azureName"`
	CustomDomain         *CustomDomain               `json:"customDomain,omitempty"`
	Encryption           *Encryption                 `json:"encryption,omitempty"`
	ExtendedLocation     *ExtendedLocation           `json:"extendedLocation,omitempty"`
	Identity             *Identity                   `json:"identity,omitempty"`
	IsHnsEnabled         *bool                       `json:"isHnsEnabled,omitempty"`
	IsNfsV3Enabled       *bool                       `json:"isNfsV3Enabled,omitempty"`
	KeyPolicy            *KeyPolicy                  `json:"keyPolicy,omitempty"`
	Kind                 *string                     `json:"kind,omitempty"`
	LargeFileSharesState *string                     `json:"largeFileSharesState,omitempty"`
	Location             *string                     `json:"location,omitempty"`
	MinimumTlsVersion    *string                     `json:"minimumTlsVersion,omitempty"`
	NetworkAcls          *NetworkRuleSet             `json:"networkAcls,omitempty"`
	OperatorSpec         *StorageAccountOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion      string                      `json:"originalVersion"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a resources.azure.com/ResourceGroup resource
	Owner                    genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag              genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RoutingPreference        *RoutingPreference                `json:"routingPreference,omitempty"`
	SasPolicy                *SasPolicy                        `json:"sasPolicy,omitempty"`
	Sku                      *Sku                              `json:"sku,omitempty"`
	SupportsHttpsTrafficOnly *bool                             `json:"supportsHttpsTrafficOnly,omitempty"`
	Tags                     map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccounts_Spec{}

// ConvertSpecFrom populates our StorageAccounts_Spec from the provided source
func (accounts *StorageAccounts_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == accounts {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(accounts)
}

// ConvertSpecTo populates the provided destination from our StorageAccounts_Spec
func (accounts *StorageAccounts_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == accounts {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(accounts)
}

//Storage version of v1alpha1api20210401.AzureFilesIdentityBasedAuthentication
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/AzureFilesIdentityBasedAuthentication
type AzureFilesIdentityBasedAuthentication struct {
	ActiveDirectoryProperties *ActiveDirectoryProperties `json:"activeDirectoryProperties,omitempty"`
	DefaultSharePermission    *string                    `json:"defaultSharePermission,omitempty"`
	DirectoryServiceOptions   *string                    `json:"directoryServiceOptions,omitempty"`
	PropertyBag               genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.AzureFilesIdentityBasedAuthentication_Status
type AzureFilesIdentityBasedAuthentication_Status struct {
	ActiveDirectoryProperties *ActiveDirectoryProperties_Status `json:"activeDirectoryProperties,omitempty"`
	DefaultSharePermission    *string                           `json:"defaultSharePermission,omitempty"`
	DirectoryServiceOptions   *string                           `json:"directoryServiceOptions,omitempty"`
	PropertyBag               genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.BlobRestoreStatus_Status
type BlobRestoreStatus_Status struct {
	FailureReason *string                       `json:"failureReason,omitempty"`
	Parameters    *BlobRestoreParameters_Status `json:"parameters,omitempty"`
	PropertyBag   genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RestoreId     *string                       `json:"restoreId,omitempty"`
	Status        *string                       `json:"status,omitempty"`
}

//Storage version of v1alpha1api20210401.CustomDomain
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/CustomDomain
type CustomDomain struct {
	Name             *string                `json:"name,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UseSubDomainName *bool                  `json:"useSubDomainName,omitempty"`
}

//Storage version of v1alpha1api20210401.CustomDomain_Status
type CustomDomain_Status struct {
	Name             *string                `json:"name,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UseSubDomainName *bool                  `json:"useSubDomainName,omitempty"`
}

//Storage version of v1alpha1api20210401.Encryption
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/Encryption
type Encryption struct {
	Identity                        *EncryptionIdentity    `json:"identity,omitempty"`
	KeySource                       *string                `json:"keySource,omitempty"`
	Keyvaultproperties              *KeyVaultProperties    `json:"keyvaultproperties,omitempty"`
	PropertyBag                     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequireInfrastructureEncryption *bool                  `json:"requireInfrastructureEncryption,omitempty"`
	Services                        *EncryptionServices    `json:"services,omitempty"`
}

//Storage version of v1alpha1api20210401.Encryption_Status
type Encryption_Status struct {
	Identity                        *EncryptionIdentity_Status `json:"identity,omitempty"`
	KeySource                       *string                    `json:"keySource,omitempty"`
	Keyvaultproperties              *KeyVaultProperties_Status `json:"keyvaultproperties,omitempty"`
	PropertyBag                     genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	RequireInfrastructureEncryption *bool                      `json:"requireInfrastructureEncryption,omitempty"`
	Services                        *EncryptionServices_Status `json:"services,omitempty"`
}

//Storage version of v1alpha1api20210401.Endpoints_Status
type Endpoints_Status struct {
	Blob               *string                                  `json:"blob,omitempty"`
	Dfs                *string                                  `json:"dfs,omitempty"`
	File               *string                                  `json:"file,omitempty"`
	InternetEndpoints  *StorageAccountInternetEndpoints_Status  `json:"internetEndpoints,omitempty"`
	MicrosoftEndpoints *StorageAccountMicrosoftEndpoints_Status `json:"microsoftEndpoints,omitempty"`
	PropertyBag        genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	Queue              *string                                  `json:"queue,omitempty"`
	Table              *string                                  `json:"table,omitempty"`
	Web                *string                                  `json:"web,omitempty"`
}

//Storage version of v1alpha1api20210401.ExtendedLocation
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ExtendedLocation
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210401.ExtendedLocation_Status
type ExtendedLocation_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210401.GeoReplicationStats_Status
type GeoReplicationStats_Status struct {
	CanFailover  *bool                  `json:"canFailover,omitempty"`
	LastSyncTime *string                `json:"lastSyncTime,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status       *string                `json:"status,omitempty"`
}

//Storage version of v1alpha1api20210401.Identity
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/Identity
type Identity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210401.Identity_Status
type Identity_Status struct {
	PrincipalId            *string                                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_Status `json:"userAssignedIdentities,omitempty"`
}

//Storage version of v1alpha1api20210401.KeyCreationTime_Status
type KeyCreationTime_Status struct {
	Key1        *string                `json:"key1,omitempty"`
	Key2        *string                `json:"key2,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.KeyPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/KeyPolicy
type KeyPolicy struct {
	KeyExpirationPeriodInDays *int                   `json:"keyExpirationPeriodInDays,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.KeyPolicy_Status
type KeyPolicy_Status struct {
	KeyExpirationPeriodInDays *int                   `json:"keyExpirationPeriodInDays,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.NetworkRuleSet
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/NetworkRuleSet
type NetworkRuleSet struct {
	Bypass              *string                `json:"bypass,omitempty"`
	DefaultAction       *string                `json:"defaultAction,omitempty"`
	IpRules             []IPRule               `json:"ipRules,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceAccessRules []ResourceAccessRule   `json:"resourceAccessRules,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule   `json:"virtualNetworkRules,omitempty"`
}

//Storage version of v1alpha1api20210401.NetworkRuleSet_Status
type NetworkRuleSet_Status struct {
	Bypass              *string                     `json:"bypass,omitempty"`
	DefaultAction       *string                     `json:"defaultAction,omitempty"`
	IpRules             []IPRule_Status             `json:"ipRules,omitempty"`
	PropertyBag         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	ResourceAccessRules []ResourceAccessRule_Status `json:"resourceAccessRules,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule_Status `json:"virtualNetworkRules,omitempty"`
}

//Storage version of v1alpha1api20210401.PrivateEndpointConnection_Status_SubResourceEmbedded
type PrivateEndpointConnection_Status_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.RoutingPreference
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/RoutingPreference
type RoutingPreference struct {
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublishInternetEndpoints  *bool                  `json:"publishInternetEndpoints,omitempty"`
	PublishMicrosoftEndpoints *bool                  `json:"publishMicrosoftEndpoints,omitempty"`
	RoutingChoice             *string                `json:"routingChoice,omitempty"`
}

//Storage version of v1alpha1api20210401.RoutingPreference_Status
type RoutingPreference_Status struct {
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublishInternetEndpoints  *bool                  `json:"publishInternetEndpoints,omitempty"`
	PublishMicrosoftEndpoints *bool                  `json:"publishMicrosoftEndpoints,omitempty"`
	RoutingChoice             *string                `json:"routingChoice,omitempty"`
}

//Storage version of v1alpha1api20210401.SasPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/SasPolicy
type SasPolicy struct {
	ExpirationAction    *string                `json:"expirationAction,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SasExpirationPeriod *string                `json:"sasExpirationPeriod,omitempty"`
}

//Storage version of v1alpha1api20210401.SasPolicy_Status
type SasPolicy_Status struct {
	ExpirationAction    *string                `json:"expirationAction,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SasExpirationPeriod *string                `json:"sasExpirationPeriod,omitempty"`
}

//Storage version of v1alpha1api20210401.Sku
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/Sku
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20210401.Sku_Status
type Sku_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20210401.StorageAccountOperatorSpec
//Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type StorageAccountOperatorSpec struct {
	PropertyBag genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	Secrets     *StorageAccountOperatorSecrets `json:"secrets,omitempty"`
}

//Storage version of v1alpha1api20210401.ActiveDirectoryProperties
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ActiveDirectoryProperties
type ActiveDirectoryProperties struct {
	AzureStorageSid   *string                `json:"azureStorageSid,omitempty"`
	DomainGuid        *string                `json:"domainGuid,omitempty"`
	DomainName        *string                `json:"domainName,omitempty"`
	DomainSid         *string                `json:"domainSid,omitempty"`
	ForestName        *string                `json:"forestName,omitempty"`
	NetBiosDomainName *string                `json:"netBiosDomainName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.ActiveDirectoryProperties_Status
type ActiveDirectoryProperties_Status struct {
	AzureStorageSid   *string                `json:"azureStorageSid,omitempty"`
	DomainGuid        *string                `json:"domainGuid,omitempty"`
	DomainName        *string                `json:"domainName,omitempty"`
	DomainSid         *string                `json:"domainSid,omitempty"`
	ForestName        *string                `json:"forestName,omitempty"`
	NetBiosDomainName *string                `json:"netBiosDomainName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.BlobRestoreParameters_Status
type BlobRestoreParameters_Status struct {
	BlobRanges    []BlobRestoreRange_Status `json:"blobRanges,omitempty"`
	PropertyBag   genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	TimeToRestore *string                   `json:"timeToRestore,omitempty"`
}

//Storage version of v1alpha1api20210401.EncryptionIdentity
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/EncryptionIdentity
type EncryptionIdentity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//UserAssignedIdentityReference: Resource identifier of the UserAssigned identity to be associated with server-side
	//encryption on the storage account.
	UserAssignedIdentityReference *genruntime.ResourceReference `armReference:"UserAssignedIdentity" json:"userAssignedIdentityReference,omitempty"`
}

//Storage version of v1alpha1api20210401.EncryptionIdentity_Status
type EncryptionIdentity_Status struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

//Storage version of v1alpha1api20210401.EncryptionServices
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/EncryptionServices
type EncryptionServices struct {
	Blob        *EncryptionService     `json:"blob,omitempty"`
	File        *EncryptionService     `json:"file,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Queue       *EncryptionService     `json:"queue,omitempty"`
	Table       *EncryptionService     `json:"table,omitempty"`
}

//Storage version of v1alpha1api20210401.EncryptionServices_Status
type EncryptionServices_Status struct {
	Blob        *EncryptionService_Status `json:"blob,omitempty"`
	File        *EncryptionService_Status `json:"file,omitempty"`
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Queue       *EncryptionService_Status `json:"queue,omitempty"`
	Table       *EncryptionService_Status `json:"table,omitempty"`
}

//Storage version of v1alpha1api20210401.IPRule
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/IPRule
type IPRule struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

//Storage version of v1alpha1api20210401.IPRule_Status
type IPRule_Status struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

//Storage version of v1alpha1api20210401.KeyVaultProperties
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/KeyVaultProperties
type KeyVaultProperties struct {
	Keyname     *string                `json:"keyname,omitempty"`
	Keyvaulturi *string                `json:"keyvaulturi,omitempty"`
	Keyversion  *string                `json:"keyversion,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.KeyVaultProperties_Status
type KeyVaultProperties_Status struct {
	CurrentVersionedKeyIdentifier *string                `json:"currentVersionedKeyIdentifier,omitempty"`
	Keyname                       *string                `json:"keyname,omitempty"`
	Keyvaulturi                   *string                `json:"keyvaulturi,omitempty"`
	Keyversion                    *string                `json:"keyversion,omitempty"`
	LastKeyRotationTimestamp      *string                `json:"lastKeyRotationTimestamp,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.ResourceAccessRule
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ResourceAccessRule
type ResourceAccessRule struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//ResourceReference: Resource Id
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
	TenantId          *string                       `json:"tenantId,omitempty"`
}

//Storage version of v1alpha1api20210401.ResourceAccessRule_Status
type ResourceAccessRule_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId  *string                `json:"resourceId,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
}

//Storage version of v1alpha1api20210401.StorageAccountInternetEndpoints_Status
type StorageAccountInternetEndpoints_Status struct {
	Blob        *string                `json:"blob,omitempty"`
	Dfs         *string                `json:"dfs,omitempty"`
	File        *string                `json:"file,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Web         *string                `json:"web,omitempty"`
}

//Storage version of v1alpha1api20210401.StorageAccountMicrosoftEndpoints_Status
type StorageAccountMicrosoftEndpoints_Status struct {
	Blob        *string                `json:"blob,omitempty"`
	Dfs         *string                `json:"dfs,omitempty"`
	File        *string                `json:"file,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Queue       *string                `json:"queue,omitempty"`
	Table       *string                `json:"table,omitempty"`
	Web         *string                `json:"web,omitempty"`
}

//Storage version of v1alpha1api20210401.StorageAccountOperatorSecrets
type StorageAccountOperatorSecrets struct {
	BlobEndpoint  *genruntime.SecretDestination `json:"blobEndpoint,omitempty"`
	DfsEndpoint   *genruntime.SecretDestination `json:"dfsEndpoint,omitempty"`
	FileEndpoint  *genruntime.SecretDestination `json:"fileEndpoint,omitempty"`
	Key1          *genruntime.SecretDestination `json:"key1,omitempty"`
	Key2          *genruntime.SecretDestination `json:"key2,omitempty"`
	PropertyBag   genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	QueueEndpoint *genruntime.SecretDestination `json:"queueEndpoint,omitempty"`
	TableEndpoint *genruntime.SecretDestination `json:"tableEndpoint,omitempty"`
	WebEndpoint   *genruntime.SecretDestination `json:"webEndpoint,omitempty"`
}

//Storage version of v1alpha1api20210401.UserAssignedIdentity_Status
type UserAssignedIdentity_Status struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.VirtualNetworkRule
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/VirtualNetworkRule
type VirtualNetworkRule struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	//Reference: Resource ID of a subnet, for example:
	///subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
	Reference genruntime.ResourceReference `armReference:"Id" json:"reference"`
	State     *string                      `json:"state,omitempty"`
}

//Storage version of v1alpha1api20210401.VirtualNetworkRule_Status
type VirtualNetworkRule_Status struct {
	Action      *string                `json:"action,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
}

//Storage version of v1alpha1api20210401.BlobRestoreRange_Status
type BlobRestoreRange_Status struct {
	EndRange    *string                `json:"endRange,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartRange  *string                `json:"startRange,omitempty"`
}

//Storage version of v1alpha1api20210401.EncryptionService
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/EncryptionService
type EncryptionService struct {
	Enabled     *bool                  `json:"enabled,omitempty"`
	KeyType     *string                `json:"keyType,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.EncryptionService_Status
type EncryptionService_Status struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	KeyType         *string                `json:"keyType,omitempty"`
	LastEnabledTime *string                `json:"lastEnabledTime,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&StorageAccount{}, &StorageAccountList{})
}
