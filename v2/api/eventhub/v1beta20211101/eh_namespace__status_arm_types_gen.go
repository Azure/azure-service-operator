// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

type EHNamespace_StatusARM struct {
	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Identity: Properties of BYOK Identity description
	Identity *Identity_StatusARM `json:"identity,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Properties: Namespace properties supplied for create namespace operation.
	Properties *EHNamespace_Status_PropertiesARM `json:"properties,omitempty"`

	//Sku: Properties of sku resource
	Sku *Sku_StatusARM `json:"sku,omitempty"`

	//SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type EHNamespace_Status_PropertiesARM struct {
	//AlternateName: Alternate name specified when alias and namespace names are same.
	AlternateName *string `json:"alternateName,omitempty"`

	//ClusterArmId: Cluster ARM ID of the Namespace.
	ClusterArmId *string `json:"clusterArmId,omitempty"`

	//CreatedAt: The time the Namespace was created.
	CreatedAt *string `json:"createdAt,omitempty"`

	//DisableLocalAuth: This property disables SAS authentication for the Event Hubs namespace.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	//Encryption: Properties of BYOK Encryption description
	Encryption *Encryption_StatusARM `json:"encryption,omitempty"`

	//IsAutoInflateEnabled: Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled *bool `json:"isAutoInflateEnabled,omitempty"`

	//KafkaEnabled: Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled *bool `json:"kafkaEnabled,omitempty"`

	//MaximumThroughputUnits: Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20
	//throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits *int `json:"maximumThroughputUnits,omitempty"`

	//MetricId: Identifier for Azure Insights metrics.
	MetricId *string `json:"metricId,omitempty"`

	//PrivateEndpointConnections: List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`

	//ProvisioningState: Provisioning state of the Namespace.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	//ServiceBusEndpoint: Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string `json:"serviceBusEndpoint,omitempty"`

	//Status: Status of the Namespace.
	Status *string `json:"status,omitempty"`

	//UpdatedAt: The time the Namespace was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`

	//ZoneRedundant: Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

type Identity_StatusARM struct {
	//PrincipalId: ObjectId from the KeyVault
	PrincipalId *string `json:"principalId,omitempty"`

	//TenantId: TenantId from the KeyVault
	TenantId *string `json:"tenantId,omitempty"`

	//Type: Type of managed service identity.
	Type *IdentityStatusType `json:"type,omitempty"`

	//UserAssignedIdentities: Properties for User Assigned Identities
	UserAssignedIdentities map[string]UserAssignedIdentity_StatusARM `json:"userAssignedIdentities,omitempty"`
}

type Sku_StatusARM struct {
	//Capacity: The Event Hubs throughput units for Basic or Standard tiers, where value should be 0 to 20 throughput units.
	//The Event Hubs premium units for Premium tier, where value should be 0 to 10 premium units.
	Capacity *int `json:"capacity,omitempty"`

	//Name: Name of this SKU.
	Name *SkuStatusName `json:"name,omitempty"`

	//Tier: The billing tier of this particular SKU.
	Tier *SkuStatusTier `json:"tier,omitempty"`
}

type Encryption_StatusARM struct {
	//KeySource: Enumerates the possible value of keySource for Encryption
	KeySource *EncryptionStatusKeySource `json:"keySource,omitempty"`

	//KeyVaultProperties: Properties of KeyVault
	KeyVaultProperties []KeyVaultProperties_StatusARM `json:"keyVaultProperties,omitempty"`

	//RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`
}

type IdentityStatusType string

const (
	IdentityStatusTypeNone                       = IdentityStatusType("None")
	IdentityStatusTypeSystemAssigned             = IdentityStatusType("SystemAssigned")
	IdentityStatusTypeSystemAssignedUserAssigned = IdentityStatusType("SystemAssigned, UserAssigned")
	IdentityStatusTypeUserAssigned               = IdentityStatusType("UserAssigned")
)

type PrivateEndpointConnection_Status_SubResourceEmbeddedARM struct {
	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`
}

type SkuStatusName string

const (
	SkuStatusNameBasic    = SkuStatusName("Basic")
	SkuStatusNamePremium  = SkuStatusName("Premium")
	SkuStatusNameStandard = SkuStatusName("Standard")
)

type SkuStatusTier string

const (
	SkuStatusTierBasic    = SkuStatusTier("Basic")
	SkuStatusTierPremium  = SkuStatusTier("Premium")
	SkuStatusTierStandard = SkuStatusTier("Standard")
)

type UserAssignedIdentity_StatusARM struct {
	//ClientId: Client Id of user assigned identity
	ClientId *string `json:"clientId,omitempty"`

	//PrincipalId: Principal Id of user assigned identity
	PrincipalId *string `json:"principalId,omitempty"`
}

type EncryptionStatusKeySource string

const EncryptionStatusKeySourceMicrosoftKeyVault = EncryptionStatusKeySource("Microsoft.KeyVault")

type KeyVaultProperties_StatusARM struct {
	Identity *UserAssignedIdentityProperties_StatusARM `json:"identity,omitempty"`

	//KeyName: Name of the Key from KeyVault
	KeyName *string `json:"keyName,omitempty"`

	//KeyVaultUri: Uri of KeyVault
	KeyVaultUri *string `json:"keyVaultUri,omitempty"`

	//KeyVersion: Key Version
	KeyVersion *string `json:"keyVersion,omitempty"`
}

type UserAssignedIdentityProperties_StatusARM struct {
	//UserAssignedIdentity: ARM ID of user Identity selected for encryption
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
