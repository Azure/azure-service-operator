// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DatabaseAccount_Spec struct {
	// Identity: Identity for the resource.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// Kind: Indicates the type of database account. This can only be set at database account creation.
	Kind *DatabaseAccount_Kind_Spec `json:"kind,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB database accounts.
	Properties *DatabaseAccountCreateUpdateProperties `json:"properties,omitempty"`
	Tags       map[string]string                      `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccount_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (account DatabaseAccount_Spec) GetAPIVersion() string {
	return "2021-05-15"
}

// GetName returns the Name of the resource
func (account *DatabaseAccount_Spec) GetName() string {
	return account.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts"
func (account *DatabaseAccount_Spec) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts"
}

// +kubebuilder:validation:Enum={"GlobalDocumentDB","MongoDB","Parse"}
type DatabaseAccount_Kind_Spec string

const (
	DatabaseAccount_Kind_Spec_GlobalDocumentDB = DatabaseAccount_Kind_Spec("GlobalDocumentDB")
	DatabaseAccount_Kind_Spec_MongoDB          = DatabaseAccount_Kind_Spec("MongoDB")
	DatabaseAccount_Kind_Spec_Parse            = DatabaseAccount_Kind_Spec("Parse")
)

// Mapping from string to DatabaseAccount_Kind_Spec
var databaseAccount_Kind_Spec_Values = map[string]DatabaseAccount_Kind_Spec{
	"globaldocumentdb": DatabaseAccount_Kind_Spec_GlobalDocumentDB,
	"mongodb":          DatabaseAccount_Kind_Spec_MongoDB,
	"parse":            DatabaseAccount_Kind_Spec_Parse,
}

// Properties to create and update Azure Cosmos DB database accounts.
type DatabaseAccountCreateUpdateProperties struct {
	// AnalyticalStorageConfiguration: Analytical storage specific properties.
	AnalyticalStorageConfiguration *AnalyticalStorageConfiguration `json:"analyticalStorageConfiguration,omitempty"`

	// ApiProperties: API specific properties. Currently, supported only for MongoDB API.
	ApiProperties *ApiProperties `json:"apiProperties,omitempty"`

	// BackupPolicy: The object representing the policy for taking backups on an account.
	BackupPolicy *BackupPolicy `json:"backupPolicy,omitempty"`

	// Capabilities: List of Cosmos DB capabilities for the account
	Capabilities []Capability `json:"capabilities,omitempty"`

	// ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer *ConnectorOffer `json:"connectorOffer,omitempty"`

	// ConsistencyPolicy: The consistency policy for the Cosmos DB account.
	ConsistencyPolicy *ConsistencyPolicy `json:"consistencyPolicy,omitempty"`

	// Cors: The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicy `json:"cors,omitempty"`

	// DatabaseAccountOfferType: The offer type for the database
	DatabaseAccountOfferType *DatabaseAccountOfferType `json:"databaseAccountOfferType,omitempty"`

	// DefaultIdentity: The default identity for accessing key vault used in features like customer managed keys. The default
	// identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity *string `json:"defaultIdentity,omitempty"`

	// DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata resources (databases, containers, throughput)
	// via account keys
	DisableKeyBasedMetadataWriteAccess *bool `json:"disableKeyBasedMetadataWriteAccess,omitempty"`

	// EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage *bool `json:"enableAnalyticalStorage,omitempty"`

	// EnableAutomaticFailover: Enables automatic failover of the write region in the rare event that the region is unavailable
	// due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the
	// failover priorities configured for the account.
	EnableAutomaticFailover *bool `json:"enableAutomaticFailover,omitempty"`

	// EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector *bool `json:"enableCassandraConnector,omitempty"`

	// EnableFreeTier: Flag to indicate whether Free Tier is enabled.
	EnableFreeTier *bool `json:"enableFreeTier,omitempty"`

	// EnableMultipleWriteLocations: Enables the account to write in multiple locations
	EnableMultipleWriteLocations *bool `json:"enableMultipleWriteLocations,omitempty"`

	// IpRules: List of IpRules.
	IpRules []IpAddressOrRange `json:"ipRules,omitempty"`

	// IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `json:"isVirtualNetworkFilterEnabled,omitempty"`

	// KeyVaultKeyUri: The URI of the key vault
	KeyVaultKeyUri *string `json:"keyVaultKeyUri,omitempty"`

	// Locations: An array that contains the georeplication locations enabled for the Cosmos DB account.
	Locations []Location `json:"locations,omitempty"`

	// NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass *NetworkAclBypass `json:"networkAclBypass,omitempty"`

	// NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds []string `json:"networkAclBypassResourceIds,omitempty"`

	// PublicNetworkAccess: Whether requests from Public Network are allowed
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules []VirtualNetworkRule `json:"virtualNetworkRules,omitempty"`
}

// Identity for the resource.
type ManagedServiceIdentity struct {
	// Type: The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both an implicitly
	// created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
	Type                   *ManagedServiceIdentity_Type           `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Analytical storage specific properties.
type AnalyticalStorageConfiguration struct {
	// SchemaType: Describes the types of schema for analytical storage.
	SchemaType *AnalyticalStorageSchemaType `json:"schemaType,omitempty"`
}

type ApiProperties struct {
	// ServerVersion: Describes the ServerVersion of an a MongoDB account.
	ServerVersion *ApiProperties_ServerVersion `json:"serverVersion,omitempty"`
}

type BackupPolicy struct {
	// Continuous: Mutually exclusive with all other properties
	Continuous *ContinuousModeBackupPolicy `json:"continuous,omitempty"`

	// Periodic: Mutually exclusive with all other properties
	Periodic *PeriodicModeBackupPolicy `json:"periodic,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupPolicy represents a discriminated union (JSON OneOf)
func (policy BackupPolicy) MarshalJSON() ([]byte, error) {
	if policy.Continuous != nil {
		return json.Marshal(policy.Continuous)
	}
	if policy.Periodic != nil {
		return json.Marshal(policy.Periodic)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupPolicy
func (policy *BackupPolicy) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "Continuous" {
		policy.Continuous = &ContinuousModeBackupPolicy{}
		return json.Unmarshal(data, policy.Continuous)
	}
	if discriminator == "Periodic" {
		policy.Periodic = &PeriodicModeBackupPolicy{}
		return json.Unmarshal(data, policy.Periodic)
	}

	// No error
	return nil
}

// Cosmos DB capability object
type Capability struct {
	// Name: Name of the Cosmos DB capability. For example, "name": "EnableCassandra". Current values also include
	// "EnableTable" and "EnableGremlin".
	Name *string `json:"name,omitempty"`
}

// The cassandra connector offer type for the Cosmos DB C* database account.
// +kubebuilder:validation:Enum={"Small"}
type ConnectorOffer string

const ConnectorOffer_Small = ConnectorOffer("Small")

// Mapping from string to ConnectorOffer
var connectorOffer_Values = map[string]ConnectorOffer{
	"small": ConnectorOffer_Small,
}

// The consistency policy for the Cosmos DB database account.
type ConsistencyPolicy struct {
	// DefaultConsistencyLevel: The default consistency level and configuration settings of the Cosmos DB account.
	DefaultConsistencyLevel *ConsistencyPolicy_DefaultConsistencyLevel `json:"defaultConsistencyLevel,omitempty"`

	// MaxIntervalInSeconds: When used with the Bounded Staleness consistency level, this value represents the time amount of
	// staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is
	// set to 'BoundedStaleness'.
	MaxIntervalInSeconds *int `json:"maxIntervalInSeconds,omitempty"`

	// MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this value represents the number of stale
	// requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set
	// to 'BoundedStaleness'.
	MaxStalenessPrefix *int `json:"maxStalenessPrefix,omitempty"`
}

// The CORS policy for the Cosmos DB database account.
type CorsPolicy struct {
	// AllowedHeaders: The request headers that the origin domain may specify on the CORS request.
	AllowedHeaders *string `json:"allowedHeaders,omitempty"`

	// AllowedMethods: The methods (HTTP request verbs) that the origin domain may use for a CORS request.
	AllowedMethods *string `json:"allowedMethods,omitempty"`

	// AllowedOrigins: The origin domains that are permitted to make a request against the service via CORS.
	AllowedOrigins *string `json:"allowedOrigins,omitempty"`

	// ExposedHeaders: The response headers that may be sent in the response to the CORS request and exposed by the browser to
	// the request issuer.
	ExposedHeaders *string `json:"exposedHeaders,omitempty"`

	// MaxAgeInSeconds: The maximum amount time that a browser should cache the preflight OPTIONS request.
	MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
}

// The offer type for the Cosmos DB database account.
// +kubebuilder:validation:Enum={"Standard"}
type DatabaseAccountOfferType string

const DatabaseAccountOfferType_Standard = DatabaseAccountOfferType("Standard")

// Mapping from string to DatabaseAccountOfferType
var databaseAccountOfferType_Values = map[string]DatabaseAccountOfferType{
	"standard": DatabaseAccountOfferType_Standard,
}

// IpAddressOrRange object
type IpAddressOrRange struct {
	// IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR format. Provided IPs must be
	// well-formatted and cannot be contained in one of the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12,
	// 192.168.0.0/16, since these are not enforceable by the IP address filter. Example of valid inputs: “23.40.210.245”
	// or “23.40.210.0/8”.
	IpAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

// A region in which the Azure Cosmos DB database account is deployed.
type Location struct {
	// FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
	// value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
	// regions in which the database account exists.
	FailoverPriority *int `json:"failoverPriority,omitempty"`

	// IsZoneRedundant: Flag to indicate whether or not this region is an AvailabilityZone region
	IsZoneRedundant *bool `json:"isZoneRedundant,omitempty"`

	// LocationName: The name of the region.
	LocationName *string `json:"locationName,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ManagedServiceIdentity_Type string

const (
	ManagedServiceIdentity_Type_None                       = ManagedServiceIdentity_Type("None")
	ManagedServiceIdentity_Type_SystemAssigned             = ManagedServiceIdentity_Type("SystemAssigned")
	ManagedServiceIdentity_Type_SystemAssignedUserAssigned = ManagedServiceIdentity_Type("SystemAssigned,UserAssigned")
	ManagedServiceIdentity_Type_UserAssigned               = ManagedServiceIdentity_Type("UserAssigned")
)

// Mapping from string to ManagedServiceIdentity_Type
var managedServiceIdentity_Type_Values = map[string]ManagedServiceIdentity_Type{
	"none":                        ManagedServiceIdentity_Type_None,
	"systemassigned":              ManagedServiceIdentity_Type_SystemAssigned,
	"systemassigned,userassigned": ManagedServiceIdentity_Type_SystemAssignedUserAssigned,
	"userassigned":                ManagedServiceIdentity_Type_UserAssigned,
}

// Indicates what services are allowed to bypass firewall checks.
// +kubebuilder:validation:Enum={"AzureServices","None"}
type NetworkAclBypass string

const (
	NetworkAclBypass_AzureServices = NetworkAclBypass("AzureServices")
	NetworkAclBypass_None          = NetworkAclBypass("None")
)

// Mapping from string to NetworkAclBypass
var networkAclBypass_Values = map[string]NetworkAclBypass{
	"azureservices": NetworkAclBypass_AzureServices,
	"none":          NetworkAclBypass_None,
}

// Whether requests from Public Network are allowed
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type PublicNetworkAccess string

const (
	PublicNetworkAccess_Disabled = PublicNetworkAccess("Disabled")
	PublicNetworkAccess_Enabled  = PublicNetworkAccess("Enabled")
)

// Mapping from string to PublicNetworkAccess
var publicNetworkAccess_Values = map[string]PublicNetworkAccess{
	"disabled": PublicNetworkAccess_Disabled,
	"enabled":  PublicNetworkAccess_Enabled,
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
}

// Virtual Network ACL Rule object
type VirtualNetworkRule struct {
	Id *string `json:"id,omitempty"`

	// IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVNetServiceEndpoint *bool `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}

// Describes the types of schema for analytical storage.
// +kubebuilder:validation:Enum={"FullFidelity","WellDefined"}
type AnalyticalStorageSchemaType string

const (
	AnalyticalStorageSchemaType_FullFidelity = AnalyticalStorageSchemaType("FullFidelity")
	AnalyticalStorageSchemaType_WellDefined  = AnalyticalStorageSchemaType("WellDefined")
)

// Mapping from string to AnalyticalStorageSchemaType
var analyticalStorageSchemaType_Values = map[string]AnalyticalStorageSchemaType{
	"fullfidelity": AnalyticalStorageSchemaType_FullFidelity,
	"welldefined":  AnalyticalStorageSchemaType_WellDefined,
}

// +kubebuilder:validation:Enum={"3.2","3.6","4.0"}
type ApiProperties_ServerVersion string

const (
	ApiProperties_ServerVersion_32 = ApiProperties_ServerVersion("3.2")
	ApiProperties_ServerVersion_36 = ApiProperties_ServerVersion("3.6")
	ApiProperties_ServerVersion_40 = ApiProperties_ServerVersion("4.0")
)

// Mapping from string to ApiProperties_ServerVersion
var apiProperties_ServerVersion_Values = map[string]ApiProperties_ServerVersion{
	"3.2": ApiProperties_ServerVersion_32,
	"3.6": ApiProperties_ServerVersion_36,
	"4.0": ApiProperties_ServerVersion_40,
}

// +kubebuilder:validation:Enum={"BoundedStaleness","ConsistentPrefix","Eventual","Session","Strong"}
type ConsistencyPolicy_DefaultConsistencyLevel string

const (
	ConsistencyPolicy_DefaultConsistencyLevel_BoundedStaleness = ConsistencyPolicy_DefaultConsistencyLevel("BoundedStaleness")
	ConsistencyPolicy_DefaultConsistencyLevel_ConsistentPrefix = ConsistencyPolicy_DefaultConsistencyLevel("ConsistentPrefix")
	ConsistencyPolicy_DefaultConsistencyLevel_Eventual         = ConsistencyPolicy_DefaultConsistencyLevel("Eventual")
	ConsistencyPolicy_DefaultConsistencyLevel_Session          = ConsistencyPolicy_DefaultConsistencyLevel("Session")
	ConsistencyPolicy_DefaultConsistencyLevel_Strong           = ConsistencyPolicy_DefaultConsistencyLevel("Strong")
)

// Mapping from string to ConsistencyPolicy_DefaultConsistencyLevel
var consistencyPolicy_DefaultConsistencyLevel_Values = map[string]ConsistencyPolicy_DefaultConsistencyLevel{
	"boundedstaleness": ConsistencyPolicy_DefaultConsistencyLevel_BoundedStaleness,
	"consistentprefix": ConsistencyPolicy_DefaultConsistencyLevel_ConsistentPrefix,
	"eventual":         ConsistencyPolicy_DefaultConsistencyLevel_Eventual,
	"session":          ConsistencyPolicy_DefaultConsistencyLevel_Session,
	"strong":           ConsistencyPolicy_DefaultConsistencyLevel_Strong,
}

type ContinuousModeBackupPolicy struct {
	Type ContinuousModeBackupPolicy_Type `json:"type,omitempty"`
}

type PeriodicModeBackupPolicy struct {
	// PeriodicModeProperties: Configuration values for periodic mode backup
	PeriodicModeProperties *PeriodicModeProperties       `json:"periodicModeProperties,omitempty"`
	Type                   PeriodicModeBackupPolicy_Type `json:"type,omitempty"`
}

// +kubebuilder:validation:Enum={"Continuous"}
type ContinuousModeBackupPolicy_Type string

const ContinuousModeBackupPolicy_Type_Continuous = ContinuousModeBackupPolicy_Type("Continuous")

// Mapping from string to ContinuousModeBackupPolicy_Type
var continuousModeBackupPolicy_Type_Values = map[string]ContinuousModeBackupPolicy_Type{
	"continuous": ContinuousModeBackupPolicy_Type_Continuous,
}

// +kubebuilder:validation:Enum={"Periodic"}
type PeriodicModeBackupPolicy_Type string

const PeriodicModeBackupPolicy_Type_Periodic = PeriodicModeBackupPolicy_Type("Periodic")

// Mapping from string to PeriodicModeBackupPolicy_Type
var periodicModeBackupPolicy_Type_Values = map[string]PeriodicModeBackupPolicy_Type{
	"periodic": PeriodicModeBackupPolicy_Type_Periodic,
}

// Configuration values for periodic mode backup
type PeriodicModeProperties struct {
	// BackupIntervalInMinutes: An integer representing the interval in minutes between two backups
	BackupIntervalInMinutes *int `json:"backupIntervalInMinutes,omitempty"`

	// BackupRetentionIntervalInHours: An integer representing the time (in hours) that each backup is retained
	BackupRetentionIntervalInHours *int `json:"backupRetentionIntervalInHours,omitempty"`
}
