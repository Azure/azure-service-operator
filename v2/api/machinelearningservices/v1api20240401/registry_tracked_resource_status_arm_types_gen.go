// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240401

type RegistryTrackedResource_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Kind: Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `json:"kind,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: [Required] Additional attributes of the entity.
	Properties *Registry_STATUS_ARM `json:"properties,omitempty"`

	// Sku: Sku details required for ARM contract for Autoscaling.
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity_STATUS_ARM struct {
	// PrincipalId: The service principal ID of the system assigned identity. This property will only be provided for a system
	// assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type                   *ManagedServiceIdentityType_STATUS_ARM     `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// Details of the Registry
type Registry_STATUS_ARM struct {
	// DiscoveryUrl: Discovery URL for the Registry
	DiscoveryUrl *string `json:"discoveryUrl,omitempty"`

	// IntellectualPropertyPublisher: IntellectualPropertyPublisher for the registry
	IntellectualPropertyPublisher *string `json:"intellectualPropertyPublisher,omitempty"`

	// ManagedResourceGroup: ResourceId of the managed RG if the registry has system created resources
	ManagedResourceGroup *ArmResourceId_STATUS_ARM `json:"managedResourceGroup,omitempty"`

	// MlFlowRegistryUri: MLFlow Registry URI for the Registry
	MlFlowRegistryUri *string `json:"mlFlowRegistryUri,omitempty"`

	// PublicNetworkAccess: Is the Registry accessible from the internet?
	// Possible values: "Enabled" or "Disabled"
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty"`

	// RegionDetails: Details of each region the registry is in
	RegionDetails []RegistryRegionArmDetails_STATUS_ARM `json:"regionDetails,omitempty"`

	// RegistryPrivateEndpointConnections: Private endpoint connections info used for pending connections in private link portal
	RegistryPrivateEndpointConnections []RegistryPrivateEndpointConnection_STATUS_ARM `json:"registryPrivateEndpointConnections,omitempty"`
}

// The resource model definition representing SKU
type Sku_STATUS_ARM struct {
	// Capacity: If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible
	// for the resource this may be omitted.
	Capacity *int `json:"capacity,omitempty"`

	// Family: If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`

	// Name: The name of the SKU. Ex - P3. It is typically a letter+number code
	Name *string `json:"name,omitempty"`

	// Size: The SKU size. When the name field is the combination of tier and some other value, this would be the standalone
	// code.
	Size *string `json:"size,omitempty"`

	// Tier: This field is required to be implemented by the Resource Provider if the service has more than one tier, but is
	// not  required on a PUT.
	Tier *SkuTier_STATUS_ARM `json:"tier,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS_ARM `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS_ARM `json:"lastModifiedByType,omitempty"`
}

// ARM ResourceId of a resource
type ArmResourceId_STATUS_ARM struct {
	// ResourceId: Arm ResourceId is in the format
	// "/subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/providers/Microsoft.Storage/storageAccounts/{StorageAccountName}"
	// or
	// "/subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{AcrName}"
	ResourceId *string `json:"resourceId,omitempty"`
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType_STATUS_ARM string

const (
	ManagedServiceIdentityType_STATUS_ARM_None                       = ManagedServiceIdentityType_STATUS_ARM("None")
	ManagedServiceIdentityType_STATUS_ARM_SystemAssigned             = ManagedServiceIdentityType_STATUS_ARM("SystemAssigned")
	ManagedServiceIdentityType_STATUS_ARM_SystemAssignedUserAssigned = ManagedServiceIdentityType_STATUS_ARM("SystemAssigned,UserAssigned")
	ManagedServiceIdentityType_STATUS_ARM_UserAssigned               = ManagedServiceIdentityType_STATUS_ARM("UserAssigned")
)

// Mapping from string to ManagedServiceIdentityType_STATUS_ARM
var managedServiceIdentityType_STATUS_ARM_Values = map[string]ManagedServiceIdentityType_STATUS_ARM{
	"none":                        ManagedServiceIdentityType_STATUS_ARM_None,
	"systemassigned":              ManagedServiceIdentityType_STATUS_ARM_SystemAssigned,
	"systemassigned,userassigned": ManagedServiceIdentityType_STATUS_ARM_SystemAssignedUserAssigned,
	"userassigned":                ManagedServiceIdentityType_STATUS_ARM_UserAssigned,
}

// Private endpoint connection definition.
type RegistryPrivateEndpointConnection_STATUS_ARM struct {
	// Id: This is the private endpoint connection name created on SRP
	// Full resource id:
	// /subscriptions/{subId}/resourceGroups/{rgName}/providers/Microsoft.MachineLearningServices/{resourceType}/{resourceName}/registryPrivateEndpointConnections/{peConnectionName}
	Id *string `json:"id,omitempty"`

	// Location: Same as workspace location.
	Location *string `json:"location,omitempty"`

	// Properties: Properties of the Private Endpoint Connection
	Properties *RegistryPrivateEndpointConnectionProperties_STATUS_ARM `json:"properties,omitempty"`
}

// Details for each region the registry is in
type RegistryRegionArmDetails_STATUS_ARM struct {
	// AcrDetails: List of ACR accounts
	AcrDetails []AcrDetails_STATUS_ARM `json:"acrDetails,omitempty"`

	// Location: The location where the registry exists
	Location *string `json:"location,omitempty"`

	// StorageAccountDetails: List of storage accounts
	StorageAccountDetails []StorageAccountDetails_STATUS_ARM `json:"storageAccountDetails,omitempty"`
}

// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not
// required on a PUT.
type SkuTier_STATUS_ARM string

const (
	SkuTier_STATUS_ARM_Basic    = SkuTier_STATUS_ARM("Basic")
	SkuTier_STATUS_ARM_Free     = SkuTier_STATUS_ARM("Free")
	SkuTier_STATUS_ARM_Premium  = SkuTier_STATUS_ARM("Premium")
	SkuTier_STATUS_ARM_Standard = SkuTier_STATUS_ARM("Standard")
)

// Mapping from string to SkuTier_STATUS_ARM
var skuTier_STATUS_ARM_Values = map[string]SkuTier_STATUS_ARM{
	"basic":    SkuTier_STATUS_ARM_Basic,
	"free":     SkuTier_STATUS_ARM_Free,
	"premium":  SkuTier_STATUS_ARM_Premium,
	"standard": SkuTier_STATUS_ARM_Standard,
}

type SystemData_CreatedByType_STATUS_ARM string

const (
	SystemData_CreatedByType_STATUS_ARM_Application     = SystemData_CreatedByType_STATUS_ARM("Application")
	SystemData_CreatedByType_STATUS_ARM_Key             = SystemData_CreatedByType_STATUS_ARM("Key")
	SystemData_CreatedByType_STATUS_ARM_ManagedIdentity = SystemData_CreatedByType_STATUS_ARM("ManagedIdentity")
	SystemData_CreatedByType_STATUS_ARM_User            = SystemData_CreatedByType_STATUS_ARM("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS_ARM
var systemData_CreatedByType_STATUS_ARM_Values = map[string]SystemData_CreatedByType_STATUS_ARM{
	"application":     SystemData_CreatedByType_STATUS_ARM_Application,
	"key":             SystemData_CreatedByType_STATUS_ARM_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ARM_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_ARM_User,
}

type SystemData_LastModifiedByType_STATUS_ARM string

const (
	SystemData_LastModifiedByType_STATUS_ARM_Application     = SystemData_LastModifiedByType_STATUS_ARM("Application")
	SystemData_LastModifiedByType_STATUS_ARM_Key             = SystemData_LastModifiedByType_STATUS_ARM("Key")
	SystemData_LastModifiedByType_STATUS_ARM_ManagedIdentity = SystemData_LastModifiedByType_STATUS_ARM("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_ARM_User            = SystemData_LastModifiedByType_STATUS_ARM("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS_ARM
var systemData_LastModifiedByType_STATUS_ARM_Values = map[string]SystemData_LastModifiedByType_STATUS_ARM{
	"application":     SystemData_LastModifiedByType_STATUS_ARM_Application,
	"key":             SystemData_LastModifiedByType_STATUS_ARM_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ARM_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_ARM_User,
}

// User assigned identity properties
type UserAssignedIdentity_STATUS_ARM struct {
	// ClientId: The client ID of the assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal ID of the assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

// Details of ACR account to be used for the Registry
type AcrDetails_STATUS_ARM struct {
	// SystemCreatedAcrAccount: Details of system created ACR account to be used for the Registry
	SystemCreatedAcrAccount *SystemCreatedAcrAccount_STATUS_ARM `json:"systemCreatedAcrAccount,omitempty"`

	// UserCreatedAcrAccount: Details of user created ACR account to be used for the Registry
	UserCreatedAcrAccount *UserCreatedAcrAccount_STATUS_ARM `json:"userCreatedAcrAccount,omitempty"`
}

// Properties of the Private Endpoint Connection
type RegistryPrivateEndpointConnectionProperties_STATUS_ARM struct {
	// GroupIds: The group ids
	GroupIds []string `json:"groupIds,omitempty"`

	// PrivateEndpoint: The PE network resource that is linked to this PE connection.
	PrivateEndpoint *PrivateEndpointResource_STATUS_ARM `json:"privateEndpoint,omitempty"`

	// ProvisioningState: One of null, "Succeeded", "Provisioning", "Failed". While not approved, it's null.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// RegistryPrivateLinkServiceConnectionState: The connection state.
	RegistryPrivateLinkServiceConnectionState *RegistryPrivateLinkServiceConnectionState_STATUS_ARM `json:"registryPrivateLinkServiceConnectionState,omitempty"`
}

// Details of storage account to be used for the Registry
type StorageAccountDetails_STATUS_ARM struct {
	// SystemCreatedStorageAccount: Details of system created storage account to be used for the registry
	SystemCreatedStorageAccount *SystemCreatedStorageAccount_STATUS_ARM `json:"systemCreatedStorageAccount,omitempty"`

	// UserCreatedStorageAccount: Details of user created storage account to be used for the registry
	UserCreatedStorageAccount *UserCreatedStorageAccount_STATUS_ARM `json:"userCreatedStorageAccount,omitempty"`
}

// The PE network resource that is linked to this PE connection.
type PrivateEndpointResource_STATUS_ARM struct {
	// Id: The ARM identifier for Private Endpoint
	Id *string `json:"id,omitempty"`

	// SubnetArmId: The subnetId that the private endpoint is connected to.
	SubnetArmId *string `json:"subnetArmId,omitempty"`
}

// The connection state.
type RegistryPrivateLinkServiceConnectionState_STATUS_ARM struct {
	// ActionsRequired: Some RP chose "None". Other RPs use this for region expansion.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// Description: User-defined message that, per NRP doc, may be used for approval-related message.
	Description *string `json:"description,omitempty"`

	// Status: Connection status of the service consumer with the service provider
	Status *EndpointServiceConnectionStatus_STATUS_ARM `json:"status,omitempty"`
}

type SystemCreatedAcrAccount_STATUS_ARM struct {
	// AcrAccountName: Name of the ACR account
	AcrAccountName *string `json:"acrAccountName,omitempty"`

	// AcrAccountSku: SKU of the ACR account
	AcrAccountSku *string `json:"acrAccountSku,omitempty"`

	// ArmResourceId: This is populated once the ACR account is created.
	ArmResourceId *ArmResourceId_STATUS_ARM `json:"armResourceId,omitempty"`
}

type SystemCreatedStorageAccount_STATUS_ARM struct {
	// AllowBlobPublicAccess: Public blob access allowed
	AllowBlobPublicAccess *bool `json:"allowBlobPublicAccess,omitempty"`

	// ArmResourceId: This is populated once the storage account is created.
	ArmResourceId *ArmResourceId_STATUS_ARM `json:"armResourceId,omitempty"`

	// StorageAccountHnsEnabled: HNS enabled for storage account
	StorageAccountHnsEnabled *bool `json:"storageAccountHnsEnabled,omitempty"`

	// StorageAccountName: Name of the storage account
	StorageAccountName *string `json:"storageAccountName,omitempty"`

	// StorageAccountType: Allowed values:
	// "Standard_LRS",
	// "Standard_GRS",
	// "Standard_RAGRS",
	// "Standard_ZRS",
	// "Standard_GZRS",
	// "Standard_RAGZRS",
	// "Premium_LRS",
	// "Premium_ZRS"
	StorageAccountType *string `json:"storageAccountType,omitempty"`
}

type UserCreatedAcrAccount_STATUS_ARM struct {
	// ArmResourceId: ARM ResourceId of a resource
	ArmResourceId *ArmResourceId_STATUS_ARM `json:"armResourceId,omitempty"`
}

type UserCreatedStorageAccount_STATUS_ARM struct {
	// ArmResourceId: ARM ResourceId of a resource
	ArmResourceId *ArmResourceId_STATUS_ARM `json:"armResourceId,omitempty"`
}

// Connection status of the service consumer with the service provider
type EndpointServiceConnectionStatus_STATUS_ARM string

const (
	EndpointServiceConnectionStatus_STATUS_ARM_Approved     = EndpointServiceConnectionStatus_STATUS_ARM("Approved")
	EndpointServiceConnectionStatus_STATUS_ARM_Disconnected = EndpointServiceConnectionStatus_STATUS_ARM("Disconnected")
	EndpointServiceConnectionStatus_STATUS_ARM_Pending      = EndpointServiceConnectionStatus_STATUS_ARM("Pending")
	EndpointServiceConnectionStatus_STATUS_ARM_Rejected     = EndpointServiceConnectionStatus_STATUS_ARM("Rejected")
)

// Mapping from string to EndpointServiceConnectionStatus_STATUS_ARM
var endpointServiceConnectionStatus_STATUS_ARM_Values = map[string]EndpointServiceConnectionStatus_STATUS_ARM{
	"approved":     EndpointServiceConnectionStatus_STATUS_ARM_Approved,
	"disconnected": EndpointServiceConnectionStatus_STATUS_ARM_Disconnected,
	"pending":      EndpointServiceConnectionStatus_STATUS_ARM_Pending,
	"rejected":     EndpointServiceConnectionStatus_STATUS_ARM_Rejected,
}
