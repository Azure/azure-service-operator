// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// An object that represents a container registry.
type Registry_STATUS struct {
	// Id: The resource ID.
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the container registry.
	Identity *IdentityProperties_STATUS `json:"identity,omitempty"`

	// Location: The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the container registry.
	Properties *RegistryProperties_STATUS `json:"properties,omitempty"`

	// Sku: The SKU of the container registry.
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: The tags of the resource.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}

// Managed identity for the resource.
type IdentityProperties_STATUS struct {
	// PrincipalId: The principal ID of resource identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of resource.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type.
	Type *IdentityProperties_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user identities associated with the resource. The user identity
	// dictionary key references will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/
	// providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]UserIdentityProperties_STATUS `json:"userAssignedIdentities,omitempty"`
}

// The properties of a container registry.
type RegistryProperties_STATUS struct {
	// AdminUserEnabled: The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`

	// CreationDate: The creation date of the container registry in ISO8601 format.
	CreationDate *string `json:"creationDate,omitempty"`

	// DataEndpointEnabled: Enable a single data endpoint per region for serving data.
	DataEndpointEnabled *bool `json:"dataEndpointEnabled,omitempty"`

	// DataEndpointHostNames: List of host names that will serve data when dataEndpointEnabled is true.
	DataEndpointHostNames []string `json:"dataEndpointHostNames,omitempty"`

	// Encryption: The encryption settings of container registry.
	Encryption *EncryptionProperty_STATUS `json:"encryption,omitempty"`

	// LoginServer: The URL that can be used to log into the container registry.
	LoginServer *string `json:"loginServer,omitempty"`

	// NetworkRuleBypassOptions: Whether to allow trusted Azure services to access a network restricted registry.
	NetworkRuleBypassOptions *RegistryProperties_NetworkRuleBypassOptions_STATUS `json:"networkRuleBypassOptions,omitempty"`

	// NetworkRuleSet: The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSet_STATUS `json:"networkRuleSet,omitempty"`

	// Policies: The policies for a container registry.
	Policies *Policies_STATUS `json:"policies,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections for a container registry.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: The provisioning state of the container registry at the time the operation was called.
	ProvisioningState *RegistryProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess *RegistryProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// Status: The status of the container registry at the time the operation was called.
	Status *Status_STATUS `json:"status,omitempty"`

	// ZoneRedundancy: Whether or not zone redundancy is enabled for this container registry
	ZoneRedundancy *RegistryProperties_ZoneRedundancy_STATUS `json:"zoneRedundancy,omitempty"`
}

// The SKU of a container registry.
type Sku_STATUS struct {
	// Name: The SKU name of the container registry. Required for registry creation.
	Name *Sku_Name_STATUS `json:"name,omitempty"`

	// Tier: The SKU tier based on the SKU name.
	Tier *Sku_Tier_STATUS `json:"tier,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource modification (UTC).
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

type EncryptionProperty_STATUS struct {
	// KeyVaultProperties: Key vault properties.
	KeyVaultProperties *KeyVaultProperties_STATUS `json:"keyVaultProperties,omitempty"`

	// Status: Indicates whether or not the encryption is enabled for container registry.
	Status *EncryptionProperty_Status_STATUS `json:"status,omitempty"`
}

type IdentityProperties_Type_STATUS string

const (
	IdentityProperties_Type_STATUS_None                       = IdentityProperties_Type_STATUS("None")
	IdentityProperties_Type_STATUS_SystemAssigned             = IdentityProperties_Type_STATUS("SystemAssigned")
	IdentityProperties_Type_STATUS_SystemAssignedUserAssigned = IdentityProperties_Type_STATUS("SystemAssigned, UserAssigned")
	IdentityProperties_Type_STATUS_UserAssigned               = IdentityProperties_Type_STATUS("UserAssigned")
)

// Mapping from string to IdentityProperties_Type_STATUS
var identityProperties_Type_STATUS_Values = map[string]IdentityProperties_Type_STATUS{
	"none":                         IdentityProperties_Type_STATUS_None,
	"systemassigned":               IdentityProperties_Type_STATUS_SystemAssigned,
	"systemassigned, userassigned": IdentityProperties_Type_STATUS_SystemAssignedUserAssigned,
	"userassigned":                 IdentityProperties_Type_STATUS_UserAssigned,
}

// The network rule set for a container registry.
type NetworkRuleSet_STATUS struct {
	// DefaultAction: The default action of allow or deny when no other rules match.
	DefaultAction *NetworkRuleSet_DefaultAction_STATUS `json:"defaultAction,omitempty"`

	// IpRules: The IP ACL rules.
	IpRules []IPRule_STATUS `json:"ipRules,omitempty"`
}

// The policies for a container registry.
type Policies_STATUS struct {
	// ExportPolicy: The export policy for a container registry.
	ExportPolicy *ExportPolicy_STATUS `json:"exportPolicy,omitempty"`

	// QuarantinePolicy: The quarantine policy for a container registry.
	QuarantinePolicy *QuarantinePolicy_STATUS `json:"quarantinePolicy,omitempty"`

	// RetentionPolicy: The retention policy for a container registry.
	RetentionPolicy *RetentionPolicy_STATUS `json:"retentionPolicy,omitempty"`

	// TrustPolicy: The content trust policy for a container registry.
	TrustPolicy *TrustPolicy_STATUS `json:"trustPolicy,omitempty"`
}

// An object that represents a private endpoint connection for a container registry.
type PrivateEndpointConnection_STATUS struct {
	// Id: The resource ID.
	Id *string `json:"id,omitempty"`
}

type RegistryProperties_NetworkRuleBypassOptions_STATUS string

const (
	RegistryProperties_NetworkRuleBypassOptions_STATUS_AzureServices = RegistryProperties_NetworkRuleBypassOptions_STATUS("AzureServices")
	RegistryProperties_NetworkRuleBypassOptions_STATUS_None          = RegistryProperties_NetworkRuleBypassOptions_STATUS("None")
)

// Mapping from string to RegistryProperties_NetworkRuleBypassOptions_STATUS
var registryProperties_NetworkRuleBypassOptions_STATUS_Values = map[string]RegistryProperties_NetworkRuleBypassOptions_STATUS{
	"azureservices": RegistryProperties_NetworkRuleBypassOptions_STATUS_AzureServices,
	"none":          RegistryProperties_NetworkRuleBypassOptions_STATUS_None,
}

type RegistryProperties_ProvisioningState_STATUS string

const (
	RegistryProperties_ProvisioningState_STATUS_Canceled  = RegistryProperties_ProvisioningState_STATUS("Canceled")
	RegistryProperties_ProvisioningState_STATUS_Creating  = RegistryProperties_ProvisioningState_STATUS("Creating")
	RegistryProperties_ProvisioningState_STATUS_Deleting  = RegistryProperties_ProvisioningState_STATUS("Deleting")
	RegistryProperties_ProvisioningState_STATUS_Failed    = RegistryProperties_ProvisioningState_STATUS("Failed")
	RegistryProperties_ProvisioningState_STATUS_Succeeded = RegistryProperties_ProvisioningState_STATUS("Succeeded")
	RegistryProperties_ProvisioningState_STATUS_Updating  = RegistryProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to RegistryProperties_ProvisioningState_STATUS
var registryProperties_ProvisioningState_STATUS_Values = map[string]RegistryProperties_ProvisioningState_STATUS{
	"canceled":  RegistryProperties_ProvisioningState_STATUS_Canceled,
	"creating":  RegistryProperties_ProvisioningState_STATUS_Creating,
	"deleting":  RegistryProperties_ProvisioningState_STATUS_Deleting,
	"failed":    RegistryProperties_ProvisioningState_STATUS_Failed,
	"succeeded": RegistryProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  RegistryProperties_ProvisioningState_STATUS_Updating,
}

type RegistryProperties_PublicNetworkAccess_STATUS string

const (
	RegistryProperties_PublicNetworkAccess_STATUS_Disabled = RegistryProperties_PublicNetworkAccess_STATUS("Disabled")
	RegistryProperties_PublicNetworkAccess_STATUS_Enabled  = RegistryProperties_PublicNetworkAccess_STATUS("Enabled")
)

// Mapping from string to RegistryProperties_PublicNetworkAccess_STATUS
var registryProperties_PublicNetworkAccess_STATUS_Values = map[string]RegistryProperties_PublicNetworkAccess_STATUS{
	"disabled": RegistryProperties_PublicNetworkAccess_STATUS_Disabled,
	"enabled":  RegistryProperties_PublicNetworkAccess_STATUS_Enabled,
}

type RegistryProperties_ZoneRedundancy_STATUS string

const (
	RegistryProperties_ZoneRedundancy_STATUS_Disabled = RegistryProperties_ZoneRedundancy_STATUS("Disabled")
	RegistryProperties_ZoneRedundancy_STATUS_Enabled  = RegistryProperties_ZoneRedundancy_STATUS("Enabled")
)

// Mapping from string to RegistryProperties_ZoneRedundancy_STATUS
var registryProperties_ZoneRedundancy_STATUS_Values = map[string]RegistryProperties_ZoneRedundancy_STATUS{
	"disabled": RegistryProperties_ZoneRedundancy_STATUS_Disabled,
	"enabled":  RegistryProperties_ZoneRedundancy_STATUS_Enabled,
}

type Sku_Name_STATUS string

const (
	Sku_Name_STATUS_Basic    = Sku_Name_STATUS("Basic")
	Sku_Name_STATUS_Classic  = Sku_Name_STATUS("Classic")
	Sku_Name_STATUS_Premium  = Sku_Name_STATUS("Premium")
	Sku_Name_STATUS_Standard = Sku_Name_STATUS("Standard")
)

// Mapping from string to Sku_Name_STATUS
var sku_Name_STATUS_Values = map[string]Sku_Name_STATUS{
	"basic":    Sku_Name_STATUS_Basic,
	"classic":  Sku_Name_STATUS_Classic,
	"premium":  Sku_Name_STATUS_Premium,
	"standard": Sku_Name_STATUS_Standard,
}

type Sku_Tier_STATUS string

const (
	Sku_Tier_STATUS_Basic    = Sku_Tier_STATUS("Basic")
	Sku_Tier_STATUS_Classic  = Sku_Tier_STATUS("Classic")
	Sku_Tier_STATUS_Premium  = Sku_Tier_STATUS("Premium")
	Sku_Tier_STATUS_Standard = Sku_Tier_STATUS("Standard")
)

// Mapping from string to Sku_Tier_STATUS
var sku_Tier_STATUS_Values = map[string]Sku_Tier_STATUS{
	"basic":    Sku_Tier_STATUS_Basic,
	"classic":  Sku_Tier_STATUS_Classic,
	"premium":  Sku_Tier_STATUS_Premium,
	"standard": Sku_Tier_STATUS_Standard,
}

// The status of an Azure resource at the time the operation was called.
type Status_STATUS struct {
	// DisplayStatus: The short label for the status.
	DisplayStatus *string `json:"displayStatus,omitempty"`

	// Message: The detailed message for the status, including alerts and error messages.
	Message *string `json:"message,omitempty"`

	// Timestamp: The timestamp when the status was changed to the current value.
	Timestamp *string `json:"timestamp,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}

type UserIdentityProperties_STATUS struct {
	// ClientId: The client id of user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal id of user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

type EncryptionProperty_Status_STATUS string

const (
	EncryptionProperty_Status_STATUS_Disabled = EncryptionProperty_Status_STATUS("disabled")
	EncryptionProperty_Status_STATUS_Enabled  = EncryptionProperty_Status_STATUS("enabled")
)

// Mapping from string to EncryptionProperty_Status_STATUS
var encryptionProperty_Status_STATUS_Values = map[string]EncryptionProperty_Status_STATUS{
	"disabled": EncryptionProperty_Status_STATUS_Disabled,
	"enabled":  EncryptionProperty_Status_STATUS_Enabled,
}

// The export policy for a container registry.
type ExportPolicy_STATUS struct {
	// Status: The value that indicates whether the policy is enabled or not.
	Status *ExportPolicy_Status_STATUS `json:"status,omitempty"`
}

// IP rule with specific IP or IP range in CIDR format.
type IPRule_STATUS struct {
	// Action: The action of IP ACL rule.
	Action *IPRule_Action_STATUS `json:"action,omitempty"`

	// Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	Value *string `json:"value,omitempty"`
}

type KeyVaultProperties_STATUS struct {
	// Identity: The client id of the identity which will be used to access key vault.
	Identity *string `json:"identity,omitempty"`

	// KeyIdentifier: Key vault uri to access the encryption key.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`

	// KeyRotationEnabled: Auto key rotation status for a CMK enabled registry.
	KeyRotationEnabled *bool `json:"keyRotationEnabled,omitempty"`

	// LastKeyRotationTimestamp: Timestamp of the last successful key rotation.
	LastKeyRotationTimestamp *string `json:"lastKeyRotationTimestamp,omitempty"`

	// VersionedKeyIdentifier: The fully qualified key identifier that includes the version of the key that is actually used
	// for encryption.
	VersionedKeyIdentifier *string `json:"versionedKeyIdentifier,omitempty"`
}

type NetworkRuleSet_DefaultAction_STATUS string

const (
	NetworkRuleSet_DefaultAction_STATUS_Allow = NetworkRuleSet_DefaultAction_STATUS("Allow")
	NetworkRuleSet_DefaultAction_STATUS_Deny  = NetworkRuleSet_DefaultAction_STATUS("Deny")
)

// Mapping from string to NetworkRuleSet_DefaultAction_STATUS
var networkRuleSet_DefaultAction_STATUS_Values = map[string]NetworkRuleSet_DefaultAction_STATUS{
	"allow": NetworkRuleSet_DefaultAction_STATUS_Allow,
	"deny":  NetworkRuleSet_DefaultAction_STATUS_Deny,
}

// The quarantine policy for a container registry.
type QuarantinePolicy_STATUS struct {
	// Status: The value that indicates whether the policy is enabled or not.
	Status *QuarantinePolicy_Status_STATUS `json:"status,omitempty"`
}

// The retention policy for a container registry.
type RetentionPolicy_STATUS struct {
	// Days: The number of days to retain an untagged manifest after which it gets purged.
	Days *int `json:"days,omitempty"`

	// LastUpdatedTime: The timestamp when the policy was last updated.
	LastUpdatedTime *string `json:"lastUpdatedTime,omitempty"`

	// Status: The value that indicates whether the policy is enabled or not.
	Status *RetentionPolicy_Status_STATUS `json:"status,omitempty"`
}

// The content trust policy for a container registry.
type TrustPolicy_STATUS struct {
	// Status: The value that indicates whether the policy is enabled or not.
	Status *TrustPolicy_Status_STATUS `json:"status,omitempty"`

	// Type: The type of trust policy.
	Type *TrustPolicy_Type_STATUS `json:"type,omitempty"`
}

type ExportPolicy_Status_STATUS string

const (
	ExportPolicy_Status_STATUS_Disabled = ExportPolicy_Status_STATUS("disabled")
	ExportPolicy_Status_STATUS_Enabled  = ExportPolicy_Status_STATUS("enabled")
)

// Mapping from string to ExportPolicy_Status_STATUS
var exportPolicy_Status_STATUS_Values = map[string]ExportPolicy_Status_STATUS{
	"disabled": ExportPolicy_Status_STATUS_Disabled,
	"enabled":  ExportPolicy_Status_STATUS_Enabled,
}

type IPRule_Action_STATUS string

const IPRule_Action_STATUS_Allow = IPRule_Action_STATUS("Allow")

// Mapping from string to IPRule_Action_STATUS
var iPRule_Action_STATUS_Values = map[string]IPRule_Action_STATUS{
	"allow": IPRule_Action_STATUS_Allow,
}

type QuarantinePolicy_Status_STATUS string

const (
	QuarantinePolicy_Status_STATUS_Disabled = QuarantinePolicy_Status_STATUS("disabled")
	QuarantinePolicy_Status_STATUS_Enabled  = QuarantinePolicy_Status_STATUS("enabled")
)

// Mapping from string to QuarantinePolicy_Status_STATUS
var quarantinePolicy_Status_STATUS_Values = map[string]QuarantinePolicy_Status_STATUS{
	"disabled": QuarantinePolicy_Status_STATUS_Disabled,
	"enabled":  QuarantinePolicy_Status_STATUS_Enabled,
}

type RetentionPolicy_Status_STATUS string

const (
	RetentionPolicy_Status_STATUS_Disabled = RetentionPolicy_Status_STATUS("disabled")
	RetentionPolicy_Status_STATUS_Enabled  = RetentionPolicy_Status_STATUS("enabled")
)

// Mapping from string to RetentionPolicy_Status_STATUS
var retentionPolicy_Status_STATUS_Values = map[string]RetentionPolicy_Status_STATUS{
	"disabled": RetentionPolicy_Status_STATUS_Disabled,
	"enabled":  RetentionPolicy_Status_STATUS_Enabled,
}

type TrustPolicy_Status_STATUS string

const (
	TrustPolicy_Status_STATUS_Disabled = TrustPolicy_Status_STATUS("disabled")
	TrustPolicy_Status_STATUS_Enabled  = TrustPolicy_Status_STATUS("enabled")
)

// Mapping from string to TrustPolicy_Status_STATUS
var trustPolicy_Status_STATUS_Values = map[string]TrustPolicy_Status_STATUS{
	"disabled": TrustPolicy_Status_STATUS_Disabled,
	"enabled":  TrustPolicy_Status_STATUS_Enabled,
}

type TrustPolicy_Type_STATUS string

const TrustPolicy_Type_STATUS_Notary = TrustPolicy_Type_STATUS("Notary")

// Mapping from string to TrustPolicy_Type_STATUS
var trustPolicy_Type_STATUS_Values = map[string]TrustPolicy_Type_STATUS{
	"notary": TrustPolicy_Type_STATUS_Notary,
}
