// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210901

type Registry_StatusARM struct {
	//Id: The resource ID.
	Id *string `json:"id,omitempty"`

	//Identity: The identity of the container registry.
	Identity *IdentityProperties_StatusARM `json:"identity,omitempty"`

	//Location: The location of the resource. This cannot be changed after the
	//resource is created.
	Location *string `json:"location,omitempty"`

	//Name: The name of the resource.
	Name *string `json:"name,omitempty"`

	//Properties: The properties of the container registry.
	Properties *RegistryProperties_StatusARM `json:"properties,omitempty"`

	//Sku: The SKU of the container registry.
	Sku *Sku_StatusARM `json:"sku,omitempty"`

	//SystemData: Metadata pertaining to creation and last modification of the
	//resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Tags: The tags of the resource.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}

type IdentityProperties_StatusARM struct {
	//PrincipalId: The principal ID of resource identity.
	PrincipalId *string `json:"principalId,omitempty"`

	//TenantId: The tenant ID of resource.
	TenantId *string `json:"tenantId,omitempty"`

	//Type: The identity type.
	Type *IdentityPropertiesStatusType `json:"type,omitempty"`

	//UserAssignedIdentities: The list of user identities associated with the
	//resource. The user identity
	//dictionary key references will be ARM resource ids in the form:
	//'/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/
	//providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]UserIdentityProperties_StatusARM `json:"userAssignedIdentities,omitempty"`
}

type RegistryProperties_StatusARM struct {
	//AdminUserEnabled: The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`

	//CreationDate: The creation date of the container registry in ISO8601 format.
	CreationDate *string `json:"creationDate,omitempty"`

	//DataEndpointEnabled: Enable a single data endpoint per region for serving data.
	DataEndpointEnabled *bool `json:"dataEndpointEnabled,omitempty"`

	//DataEndpointHostNames: List of host names that will serve data when
	//dataEndpointEnabled is true.
	DataEndpointHostNames []string `json:"dataEndpointHostNames,omitempty"`

	//Encryption: The encryption settings of container registry.
	Encryption *EncryptionProperty_StatusARM `json:"encryption,omitempty"`

	//LoginServer: The URL that can be used to log into the container registry.
	LoginServer *string `json:"loginServer,omitempty"`

	//NetworkRuleBypassOptions: Whether to allow trusted Azure services to access a
	//network restricted registry.
	NetworkRuleBypassOptions *RegistryPropertiesStatusNetworkRuleBypassOptions `json:"networkRuleBypassOptions,omitempty"`

	//NetworkRuleSet: The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSet_StatusARM `json:"networkRuleSet,omitempty"`

	//Policies: The policies for a container registry.
	Policies *Policies_StatusARM `json:"policies,omitempty"`

	//PrivateEndpointConnections: List of private endpoint connections for a container
	//registry.
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`

	//ProvisioningState: The provisioning state of the container registry at the time
	//the operation was called.
	ProvisioningState *RegistryPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	//PublicNetworkAccess: Whether or not public network access is allowed for the
	//container registry.
	PublicNetworkAccess *RegistryPropertiesStatusPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	//Status: The status of the container registry at the time the operation was
	//called.
	Status *Status_StatusARM `json:"status,omitempty"`

	//ZoneRedundancy: Whether or not zone redundancy is enabled for this container
	//registry
	ZoneRedundancy *RegistryPropertiesStatusZoneRedundancy `json:"zoneRedundancy,omitempty"`
}

type Sku_StatusARM struct {
	//Name: The SKU name of the container registry. Required for registry creation.
	Name SkuStatusName `json:"name"`

	//Tier: The SKU tier based on the SKU name.
	Tier *SkuStatusTier `json:"tier,omitempty"`
}

type SystemData_StatusARM struct {
	//CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	//CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	//CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemDataStatusCreatedByType `json:"createdByType,omitempty"`

	//LastModifiedAt: The timestamp of resource modification (UTC).
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	//LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	//LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type EncryptionProperty_StatusARM struct {
	//KeyVaultProperties: Key vault properties.
	KeyVaultProperties *KeyVaultProperties_StatusARM `json:"keyVaultProperties,omitempty"`

	//Status: Indicates whether or not the encryption is enabled for container
	//registry.
	Status *EncryptionPropertyStatusStatus `json:"status,omitempty"`
}

type IdentityPropertiesStatusType string

const (
	IdentityPropertiesStatusTypeNone                       = IdentityPropertiesStatusType("None")
	IdentityPropertiesStatusTypeSystemAssigned             = IdentityPropertiesStatusType("SystemAssigned")
	IdentityPropertiesStatusTypeSystemAssignedUserAssigned = IdentityPropertiesStatusType("SystemAssigned, UserAssigned")
	IdentityPropertiesStatusTypeUserAssigned               = IdentityPropertiesStatusType("UserAssigned")
)

type NetworkRuleSet_StatusARM struct {
	//DefaultAction: The default action of allow or deny when no other rules match.
	DefaultAction NetworkRuleSetStatusDefaultAction `json:"defaultAction"`

	//IpRules: The IP ACL rules.
	IpRules []IPRule_StatusARM `json:"ipRules,omitempty"`
}

type Policies_StatusARM struct {
	//ExportPolicy: The export policy for a container registry.
	ExportPolicy *ExportPolicy_StatusARM `json:"exportPolicy,omitempty"`

	//QuarantinePolicy: The quarantine policy for a container registry.
	QuarantinePolicy *QuarantinePolicy_StatusARM `json:"quarantinePolicy,omitempty"`

	//RetentionPolicy: The retention policy for a container registry.
	RetentionPolicy *RetentionPolicy_StatusARM `json:"retentionPolicy,omitempty"`

	//TrustPolicy: The content trust policy for a container registry.
	TrustPolicy *TrustPolicy_StatusARM `json:"trustPolicy,omitempty"`
}

type PrivateEndpointConnection_Status_SubResourceEmbeddedARM struct {
	//Id: The resource ID.
	Id *string `json:"id,omitempty"`

	//SystemData: Metadata pertaining to creation and last modification of the
	//resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`
}

type SkuStatusName string

const (
	SkuStatusNameBasic    = SkuStatusName("Basic")
	SkuStatusNameClassic  = SkuStatusName("Classic")
	SkuStatusNamePremium  = SkuStatusName("Premium")
	SkuStatusNameStandard = SkuStatusName("Standard")
)

type SkuStatusTier string

const (
	SkuStatusTierBasic    = SkuStatusTier("Basic")
	SkuStatusTierClassic  = SkuStatusTier("Classic")
	SkuStatusTierPremium  = SkuStatusTier("Premium")
	SkuStatusTierStandard = SkuStatusTier("Standard")
)

type Status_StatusARM struct {
	//DisplayStatus: The short label for the status.
	DisplayStatus *string `json:"displayStatus,omitempty"`

	//Message: The detailed message for the status, including alerts and error
	//messages.
	Message *string `json:"message,omitempty"`

	//Timestamp: The timestamp when the status was changed to the current value.
	Timestamp *string `json:"timestamp,omitempty"`
}

type SystemDataStatusCreatedByType string

const (
	SystemDataStatusCreatedByTypeApplication     = SystemDataStatusCreatedByType("Application")
	SystemDataStatusCreatedByTypeKey             = SystemDataStatusCreatedByType("Key")
	SystemDataStatusCreatedByTypeManagedIdentity = SystemDataStatusCreatedByType("ManagedIdentity")
	SystemDataStatusCreatedByTypeUser            = SystemDataStatusCreatedByType("User")
)

type SystemDataStatusLastModifiedByType string

const (
	SystemDataStatusLastModifiedByTypeApplication     = SystemDataStatusLastModifiedByType("Application")
	SystemDataStatusLastModifiedByTypeKey             = SystemDataStatusLastModifiedByType("Key")
	SystemDataStatusLastModifiedByTypeManagedIdentity = SystemDataStatusLastModifiedByType("ManagedIdentity")
	SystemDataStatusLastModifiedByTypeUser            = SystemDataStatusLastModifiedByType("User")
)

type UserIdentityProperties_StatusARM struct {
	//ClientId: The client id of user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	//PrincipalId: The principal id of user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

type ExportPolicy_StatusARM struct {
	//Status: The value that indicates whether the policy is enabled or not.
	Status *ExportPolicyStatusStatus `json:"status,omitempty"`
}

type IPRule_StatusARM struct {
	//Action: The action of IP ACL rule.
	Action *IPRuleStatusAction `json:"action,omitempty"`

	//Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	Value string `json:"value"`
}

type KeyVaultProperties_StatusARM struct {
	//Identity: The client id of the identity which will be used to access key vault.
	Identity *string `json:"identity,omitempty"`

	//KeyIdentifier: Key vault uri to access the encryption key.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`

	//KeyRotationEnabled: Auto key rotation status for a CMK enabled registry.
	KeyRotationEnabled *bool `json:"keyRotationEnabled,omitempty"`

	//LastKeyRotationTimestamp: Timestamp of the last successful key rotation.
	LastKeyRotationTimestamp *string `json:"lastKeyRotationTimestamp,omitempty"`

	//VersionedKeyIdentifier: The fully qualified key identifier that includes the
	//version of the key that is actually used for encryption.
	VersionedKeyIdentifier *string `json:"versionedKeyIdentifier,omitempty"`
}

type QuarantinePolicy_StatusARM struct {
	//Status: The value that indicates whether the policy is enabled or not.
	Status *QuarantinePolicyStatusStatus `json:"status,omitempty"`
}

type RetentionPolicy_StatusARM struct {
	//Days: The number of days to retain an untagged manifest after which it gets
	//purged.
	Days *int `json:"days,omitempty"`

	//LastUpdatedTime: The timestamp when the policy was last updated.
	LastUpdatedTime *string `json:"lastUpdatedTime,omitempty"`

	//Status: The value that indicates whether the policy is enabled or not.
	Status *RetentionPolicyStatusStatus `json:"status,omitempty"`
}

type TrustPolicy_StatusARM struct {
	//Status: The value that indicates whether the policy is enabled or not.
	Status *TrustPolicyStatusStatus `json:"status,omitempty"`

	//Type: The type of trust policy.
	Type *TrustPolicyStatusType `json:"type,omitempty"`
}
