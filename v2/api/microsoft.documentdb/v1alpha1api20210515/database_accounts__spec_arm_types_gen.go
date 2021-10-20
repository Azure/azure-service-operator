// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DatabaseAccounts_SpecARM struct {
	//APIVersion: API Version of the resource type, optional when apiProfile is used
	//on the template
	APIVersion DatabaseAccountsSpecAPIVersion `json:"apiVersion"`

	//Identity: Identity for the resource.
	Identity *ManagedServiceIdentityARM `json:"identity,omitempty"`

	//Kind: Indicates the type of database account. This can only be set at database
	//account creation.
	Kind *DatabaseAccountsSpecKind `json:"kind,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: Cosmos DB database account name.
	Name string `json:"name"`

	//Properties: Properties to create and update Azure Cosmos DB database accounts.
	Properties DatabaseAccountCreateUpdatePropertiesARM `json:"properties"`

	//Tags: Tags are a list of key-value pairs that describe the resource. These tags
	//can be used in viewing and grouping this resource (across resource groups). A
	//maximum of 15 tags can be provided for a resource. Each tag must have a key no
	//greater than 128 characters and value no greater than 256 characters. For
	//example, the default experience for a template type is set with
	//"defaultExperience": "Cassandra". Current "defaultExperience" values also
	//include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type
	Type DatabaseAccountsSpecType `json:"type"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SpecARM{}

// GetAPIVersion returns the APIVersion of the resource
func (databaseAccountsSpecARM DatabaseAccounts_SpecARM) GetAPIVersion() string {
	return string(databaseAccountsSpecARM.APIVersion)
}

// GetName returns the Name of the resource
func (databaseAccountsSpecARM DatabaseAccounts_SpecARM) GetName() string {
	return databaseAccountsSpecARM.Name
}

// GetType returns the Type of the resource
func (databaseAccountsSpecARM DatabaseAccounts_SpecARM) GetType() string {
	return string(databaseAccountsSpecARM.Type)
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/DatabaseAccountCreateUpdateProperties
type DatabaseAccountCreateUpdatePropertiesARM struct {
	//AnalyticalStorageConfiguration: Analytical storage specific properties.
	AnalyticalStorageConfiguration *AnalyticalStorageConfigurationARM `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                  *ApiPropertiesARM                  `json:"apiProperties,omitempty"`

	//BackupPolicy: The object representing the policy for taking backups on an
	//account.
	BackupPolicy *BackupPolicyARM `json:"backupPolicy,omitempty"`

	//Capabilities: List of Cosmos DB capabilities for the account
	Capabilities []CapabilityARM `json:"capabilities,omitempty"`

	//ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C*
	//account.
	ConnectorOffer *DatabaseAccountCreateUpdatePropertiesConnectorOffer `json:"connectorOffer,omitempty"`

	//ConsistencyPolicy: The consistency policy for the Cosmos DB database account.
	ConsistencyPolicy *ConsistencyPolicyARM `json:"consistencyPolicy,omitempty"`

	//Cors: The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicyARM `json:"cors,omitempty"`

	//DatabaseAccountOfferType: The offer type for the database
	DatabaseAccountOfferType DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType `json:"databaseAccountOfferType"`

	//DefaultIdentity: The default identity for accessing key vault used in features
	//like customer managed keys. The default identity needs to be explicitly set by
	//the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity *string `json:"defaultIdentity,omitempty"`

	//DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata
	//resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess *bool `json:"disableKeyBasedMetadataWriteAccess,omitempty"`

	//EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage *bool `json:"enableAnalyticalStorage,omitempty"`

	//EnableAutomaticFailover: Enables automatic failover of the write region in the
	//rare event that the region is unavailable due to an outage. Automatic failover
	//will result in a new write region for the account and is chosen based on the
	//failover priorities configured for the account.
	EnableAutomaticFailover *bool `json:"enableAutomaticFailover,omitempty"`

	//EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C*
	//account
	EnableCassandraConnector *bool `json:"enableCassandraConnector,omitempty"`

	//EnableFreeTier: Flag to indicate whether Free Tier is enabled.
	EnableFreeTier *bool `json:"enableFreeTier,omitempty"`

	//EnableMultipleWriteLocations: Enables the account to write in multiple locations
	EnableMultipleWriteLocations *bool `json:"enableMultipleWriteLocations,omitempty"`

	//IpRules: Array of IpAddressOrRange objects.
	IpRules []IpAddressOrRangeARM `json:"ipRules,omitempty"`

	//IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable
	//Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `json:"isVirtualNetworkFilterEnabled,omitempty"`

	//KeyVaultKeyUri: The URI of the key vault
	KeyVaultKeyUri *string `json:"keyVaultKeyUri,omitempty"`

	//Locations: An array that contains the georeplication locations enabled for the
	//Cosmos DB account.
	Locations []LocationARM `json:"locations"`

	//NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass *DatabaseAccountCreateUpdatePropertiesNetworkAclBypass `json:"networkAclBypass,omitempty"`

	//NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network
	//Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds []string `json:"networkAclBypassResourceIds,omitempty"`

	//PublicNetworkAccess: Whether requests from Public Network are allowed.
	PublicNetworkAccess *DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	//VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos
	//DB account.
	VirtualNetworkRules []VirtualNetworkRuleARM `json:"virtualNetworkRules,omitempty"`
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type DatabaseAccountsSpecAPIVersion string

const DatabaseAccountsSpecAPIVersion20210515 = DatabaseAccountsSpecAPIVersion("2021-05-15")

// +kubebuilder:validation:Enum={"GlobalDocumentDB","MongoDB","Parse"}
type DatabaseAccountsSpecKind string

const (
	DatabaseAccountsSpecKindGlobalDocumentDB = DatabaseAccountsSpecKind("GlobalDocumentDB")
	DatabaseAccountsSpecKindMongoDB          = DatabaseAccountsSpecKind("MongoDB")
	DatabaseAccountsSpecKindParse            = DatabaseAccountsSpecKind("Parse")
)

// +kubebuilder:validation:Enum={"Microsoft.DocumentDB/databaseAccounts"}
type DatabaseAccountsSpecType string

const DatabaseAccountsSpecTypeMicrosoftDocumentDBDatabaseAccounts = DatabaseAccountsSpecType("Microsoft.DocumentDB/databaseAccounts")

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ManagedServiceIdentity
type ManagedServiceIdentityARM struct {
	//Type: The type of identity used for the resource. The type
	//'SystemAssigned,UserAssigned' includes both an implicitly created identity and a
	//set of user assigned identities. The type 'None' will remove any identities from
	//the service.
	Type *ManagedServiceIdentityType `json:"type,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AnalyticalStorageConfiguration
type AnalyticalStorageConfigurationARM struct {
	SchemaType *AnalyticalStorageConfigurationSchemaType `json:"schemaType,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ApiProperties
type ApiPropertiesARM struct {
	//ServerVersion: Describes the ServerVersion of an a MongoDB account.
	ServerVersion *ApiPropertiesServerVersion `json:"serverVersion,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/BackupPolicy
type BackupPolicyARM struct {
	//ContinuousModeBackupPolicy: Mutually exclusive with all other properties
	ContinuousModeBackupPolicy *ContinuousModeBackupPolicyARM `json:"continuousModeBackupPolicy,omitempty"`

	//PeriodicModeBackupPolicy: Mutually exclusive with all other properties
	PeriodicModeBackupPolicy *PeriodicModeBackupPolicyARM `json:"periodicModeBackupPolicy,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupPolicyARM represents a discriminated union (JSON OneOf)
func (backupPolicyARM BackupPolicyARM) MarshalJSON() ([]byte, error) {
	if backupPolicyARM.ContinuousModeBackupPolicy != nil {
		return json.Marshal(backupPolicyARM.ContinuousModeBackupPolicy)
	}
	if backupPolicyARM.PeriodicModeBackupPolicy != nil {
		return json.Marshal(backupPolicyARM.PeriodicModeBackupPolicy)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupPolicyARM
func (backupPolicyARM *BackupPolicyARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "Continuous" {
		backupPolicyARM.ContinuousModeBackupPolicy = &ContinuousModeBackupPolicyARM{}
		return json.Unmarshal(data, backupPolicyARM.ContinuousModeBackupPolicy)
	}
	if discriminator == "Periodic" {
		backupPolicyARM.PeriodicModeBackupPolicy = &PeriodicModeBackupPolicyARM{}
		return json.Unmarshal(data, backupPolicyARM.PeriodicModeBackupPolicy)
	}

	// No error
	return nil
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Capability
type CapabilityARM struct {
	//Name: Name of the Cosmos DB capability. For example, "name": "EnableCassandra".
	//Current values also include "EnableTable" and "EnableGremlin".
	Name *string `json:"name,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConsistencyPolicy
type ConsistencyPolicyARM struct {
	//DefaultConsistencyLevel: The default consistency level and configuration
	//settings of the Cosmos DB account.
	DefaultConsistencyLevel ConsistencyPolicyDefaultConsistencyLevel `json:"defaultConsistencyLevel"`

	//MaxIntervalInSeconds: When used with the Bounded Staleness consistency level,
	//this value represents the time amount of staleness (in seconds) tolerated.
	//Accepted range for this value is 5 - 86400. Required when
	//defaultConsistencyPolicy is set to 'BoundedStaleness'.
	MaxIntervalInSeconds *int `json:"maxIntervalInSeconds,omitempty"`

	//MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this
	//value represents the number of stale requests tolerated. Accepted range for this
	//value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set to
	//'BoundedStaleness'.
	MaxStalenessPrefix *int `json:"maxStalenessPrefix,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CorsPolicy
type CorsPolicyARM struct {
	//AllowedHeaders: The request headers that the origin domain may specify on the
	//CORS request.
	AllowedHeaders *string `json:"allowedHeaders,omitempty"`

	//AllowedMethods: The methods (HTTP request verbs) that the origin domain may use
	//for a CORS request.
	AllowedMethods *string `json:"allowedMethods,omitempty"`

	//AllowedOrigins: The origin domains that are permitted to make a request against
	//the service via CORS.
	AllowedOrigins string `json:"allowedOrigins"`

	//ExposedHeaders: The response headers that may be sent in the response to the
	//CORS request and exposed by the browser to the request issuer.
	ExposedHeaders *string `json:"exposedHeaders,omitempty"`

	//MaxAgeInSeconds: The maximum amount time that a browser should cache the
	//preflight OPTIONS request.
	MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IpAddressOrRange
type IpAddressOrRangeARM struct {
	//IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR
	//format. Provided IPs must be well-formatted and cannot be contained in one of
	//the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12, 192.168.0.0/16,
	//since these are not enforceable by the IP address filter. Example of valid
	//inputs: “23.40.210.245” or “23.40.210.0/8”.
	IpAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Location
type LocationARM struct {
	//FailoverPriority: The failover priority of the region. A failover priority of 0
	//indicates a write region. The maximum value for a failover priority = (total
	//number of regions - 1). Failover priority values must be unique for each of the
	//regions in which the database account exists.
	FailoverPriority *int `json:"failoverPriority,omitempty"`

	//IsZoneRedundant: Flag to indicate whether or not this region is an
	//AvailabilityZone region
	IsZoneRedundant *bool `json:"isZoneRedundant,omitempty"`

	//LocationName: The name of the region.
	LocationName *string `json:"locationName,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned             = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeSystemAssignedUserAssigned = ManagedServiceIdentityType("SystemAssigned,UserAssigned")
	ManagedServiceIdentityTypeUserAssigned               = ManagedServiceIdentityType("UserAssigned")
)

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/VirtualNetworkRule
type VirtualNetworkRuleARM struct {
	Id *string `json:"id,omitempty"`

	//IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual
	//network has vnet service endpoint enabled.
	IgnoreMissingVNetServiceEndpoint *bool `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContinuousModeBackupPolicy
type ContinuousModeBackupPolicyARM struct {
	Type ContinuousModeBackupPolicyType `json:"type"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeBackupPolicy
type PeriodicModeBackupPolicyARM struct {
	//PeriodicModeProperties: Configuration values for periodic mode backup
	PeriodicModeProperties *PeriodicModePropertiesARM   `json:"periodicModeProperties,omitempty"`
	Type                   PeriodicModeBackupPolicyType `json:"type"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeProperties
type PeriodicModePropertiesARM struct {
	//BackupIntervalInMinutes: An integer representing the interval in minutes between
	//two backups
	BackupIntervalInMinutes *int `json:"backupIntervalInMinutes,omitempty"`

	//BackupRetentionIntervalInHours: An integer representing the time (in hours) that
	//each backup is retained
	BackupRetentionIntervalInHours *int `json:"backupRetentionIntervalInHours,omitempty"`
}
