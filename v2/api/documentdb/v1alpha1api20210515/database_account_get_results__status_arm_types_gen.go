// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

type DatabaseAccountGetResults_StatusARM struct {
	//Id: The unique resource identifier of the ARM resource.
	Id       *string                           `json:"id,omitempty"`
	Identity *ManagedServiceIdentity_StatusARM `json:"identity,omitempty"`

	//Kind: Indicates the type of database account. This can only be set at database account creation.
	Kind *DatabaseAccountGetResultsStatusKind `json:"kind,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name       *string                                 `json:"name,omitempty"`
	Properties *DatabaseAccountGetProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                       `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

type DatabaseAccountGetProperties_StatusARM struct {
	//AnalyticalStorageConfiguration: Analytical storage specific properties.
	AnalyticalStorageConfiguration *AnalyticalStorageConfiguration_StatusARM `json:"analyticalStorageConfiguration,omitempty"`

	//ApiProperties: API specific properties.
	ApiProperties *ApiProperties_StatusARM `json:"apiProperties,omitempty"`

	//BackupPolicy: The object representing the policy for taking backups on an account.
	BackupPolicy *BackupPolicy_StatusARM `json:"backupPolicy,omitempty"`

	//Capabilities: List of Cosmos DB capabilities for the account
	Capabilities []Capability_StatusARM `json:"capabilities,omitempty"`

	//ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer *ConnectorOffer_Status `json:"connectorOffer,omitempty"`

	//ConsistencyPolicy: The consistency policy for the Cosmos DB database account.
	ConsistencyPolicy *ConsistencyPolicy_StatusARM `json:"consistencyPolicy,omitempty"`

	//Cors: The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicy_StatusARM `json:"cors,omitempty"`

	//DatabaseAccountOfferType: The offer type for the Cosmos DB database account. Default value: Standard.
	DatabaseAccountOfferType *DatabaseAccountOfferType_Status `json:"databaseAccountOfferType,omitempty"`

	//DefaultIdentity: The default identity for accessing key vault used in features like customer managed keys. The default
	//identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity *string `json:"defaultIdentity,omitempty"`

	//DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata resources (databases, containers, throughput)
	//via account keys
	DisableKeyBasedMetadataWriteAccess *bool `json:"disableKeyBasedMetadataWriteAccess,omitempty"`

	//DocumentEndpoint: The connection endpoint for the Cosmos DB database account.
	DocumentEndpoint *string `json:"documentEndpoint,omitempty"`

	//EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage *bool `json:"enableAnalyticalStorage,omitempty"`

	//EnableAutomaticFailover: Enables automatic failover of the write region in the rare event that the region is unavailable
	//due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the
	//failover priorities configured for the account.
	EnableAutomaticFailover *bool `json:"enableAutomaticFailover,omitempty"`

	//EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector *bool `json:"enableCassandraConnector,omitempty"`

	//EnableFreeTier: Flag to indicate whether Free Tier is enabled.
	EnableFreeTier *bool `json:"enableFreeTier,omitempty"`

	//EnableMultipleWriteLocations: Enables the account to write in multiple locations
	EnableMultipleWriteLocations *bool `json:"enableMultipleWriteLocations,omitempty"`

	//FailoverPolicies: An array that contains the regions ordered by their failover priorities.
	FailoverPolicies []FailoverPolicy_StatusARM `json:"failoverPolicies,omitempty"`

	//IpRules: List of IpRules.
	IpRules []IpAddressOrRange_StatusARM `json:"ipRules,omitempty"`

	//IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `json:"isVirtualNetworkFilterEnabled,omitempty"`

	//KeyVaultKeyUri: The URI of the key vault
	KeyVaultKeyUri *string `json:"keyVaultKeyUri,omitempty"`

	//Locations: An array that contains all of the locations enabled for the Cosmos DB account.
	Locations []Location_StatusARM `json:"locations,omitempty"`

	//NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass *NetworkAclBypass_Status `json:"networkAclBypass,omitempty"`

	//NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds []string `json:"networkAclBypassResourceIds,omitempty"`

	//PrivateEndpointConnections: List of Private Endpoint Connections configured for the Cosmos DB account.
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *string                                                   `json:"provisioningState,omitempty"`

	//PublicNetworkAccess: Whether requests from Public Network are allowed
	PublicNetworkAccess *PublicNetworkAccess_Status `json:"publicNetworkAccess,omitempty"`

	//ReadLocations: An array that contains of the read locations enabled for the Cosmos DB account.
	ReadLocations []Location_StatusARM `json:"readLocations,omitempty"`

	//VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules []VirtualNetworkRule_StatusARM `json:"virtualNetworkRules,omitempty"`

	//WriteLocations: An array that contains the write location for the Cosmos DB account.
	WriteLocations []Location_StatusARM `json:"writeLocations,omitempty"`
}

type DatabaseAccountGetResultsStatusKind string

const (
	DatabaseAccountGetResultsStatusKindGlobalDocumentDB = DatabaseAccountGetResultsStatusKind("GlobalDocumentDB")
	DatabaseAccountGetResultsStatusKindMongoDB          = DatabaseAccountGetResultsStatusKind("MongoDB")
	DatabaseAccountGetResultsStatusKindParse            = DatabaseAccountGetResultsStatusKind("Parse")
)

type ManagedServiceIdentity_StatusARM struct {
	//PrincipalId: The principal id of the system assigned identity. This property will only be provided for a system assigned
	//identity.
	PrincipalId *string `json:"principalId,omitempty"`

	//TenantId: The tenant id of the system assigned identity. This property will only be provided for a system assigned
	//identity.
	TenantId *string `json:"tenantId,omitempty"`

	//Type: The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both an implicitly
	//created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
	Type *ManagedServiceIdentityStatusType `json:"type,omitempty"`

	//UserAssignedIdentities: The list of user identities associated with resource. The user identity dictionary key
	//references will be ARM resource ids in the form:
	//'/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]ManagedServiceIdentity_Status_UserAssignedIdentitiesARM `json:"userAssignedIdentities,omitempty"`
}

type AnalyticalStorageConfiguration_StatusARM struct {
	SchemaType *AnalyticalStorageSchemaType_Status `json:"schemaType,omitempty"`
}

type ApiProperties_StatusARM struct {
	//ServerVersion: Describes the ServerVersion of an a MongoDB account.
	ServerVersion *ApiPropertiesStatusServerVersion `json:"serverVersion,omitempty"`
}

type BackupPolicy_StatusARM struct {
	Type *BackupPolicyType_Status `json:"type,omitempty"`
}

type Capability_StatusARM struct {
	//Name: Name of the Cosmos DB capability. For example, "name": "EnableCassandra". Current values also include
	//"EnableTable" and "EnableGremlin".
	Name *string `json:"name,omitempty"`
}

type ConsistencyPolicy_StatusARM struct {
	//DefaultConsistencyLevel: The default consistency level and configuration settings of the Cosmos DB account.
	DefaultConsistencyLevel *ConsistencyPolicyStatusDefaultConsistencyLevel `json:"defaultConsistencyLevel,omitempty"`

	//MaxIntervalInSeconds: When used with the Bounded Staleness consistency level, this value represents the time amount of
	//staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is
	//set to 'BoundedStaleness'.
	MaxIntervalInSeconds *int `json:"maxIntervalInSeconds,omitempty"`

	//MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this value represents the number of stale
	//requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set
	//to 'BoundedStaleness'.
	MaxStalenessPrefix *int `json:"maxStalenessPrefix,omitempty"`
}

type CorsPolicy_StatusARM struct {
	//AllowedHeaders: The request headers that the origin domain may specify on the CORS request.
	AllowedHeaders *string `json:"allowedHeaders,omitempty"`

	//AllowedMethods: The methods (HTTP request verbs) that the origin domain may use for a CORS request.
	AllowedMethods *string `json:"allowedMethods,omitempty"`

	//AllowedOrigins: The origin domains that are permitted to make a request against the service via CORS.
	AllowedOrigins *string `json:"allowedOrigins,omitempty"`

	//ExposedHeaders: The response headers that may be sent in the response to the CORS request and exposed by the browser to
	//the request issuer.
	ExposedHeaders *string `json:"exposedHeaders,omitempty"`

	//MaxAgeInSeconds: The maximum amount time that a browser should cache the preflight OPTIONS request.
	MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
}

type FailoverPolicy_StatusARM struct {
	//FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
	//value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
	//regions in which the database account exists.
	FailoverPriority *int `json:"failoverPriority,omitempty"`

	//Id: The unique identifier of the region in which the database account replicates to. Example:
	//&lt;accountName&gt;-&lt;locationName&gt;.
	Id *string `json:"id,omitempty"`

	//LocationName: The name of the region in which the database account exists.
	LocationName *string `json:"locationName,omitempty"`
}

type IpAddressOrRange_StatusARM struct {
	//IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR format. Provided IPs must be
	//well-formatted and cannot be contained in one of the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12,
	//192.168.0.0/16, since these are not enforceable by the IP address filter. Example of valid inputs: “23.40.210.245”
	//or “23.40.210.0/8”.
	IpAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

type Location_StatusARM struct {
	//DocumentEndpoint: The connection endpoint for the specific region. Example:
	//https://&lt;accountName&gt;-&lt;locationName&gt;.documents.azure.com:443/
	DocumentEndpoint *string `json:"documentEndpoint,omitempty"`

	//FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
	//value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
	//regions in which the database account exists.
	FailoverPriority *int `json:"failoverPriority,omitempty"`

	//Id: The unique identifier of the region within the database account. Example: &lt;accountName&gt;-&lt;locationName&gt;.
	Id *string `json:"id,omitempty"`

	//IsZoneRedundant: Flag to indicate whether or not this region is an AvailabilityZone region
	IsZoneRedundant *bool `json:"isZoneRedundant,omitempty"`

	//LocationName: The name of the region.
	LocationName      *string `json:"locationName,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

type ManagedServiceIdentityStatusType string

const (
	ManagedServiceIdentityStatusTypeNone                       = ManagedServiceIdentityStatusType("None")
	ManagedServiceIdentityStatusTypeSystemAssigned             = ManagedServiceIdentityStatusType("SystemAssigned")
	ManagedServiceIdentityStatusTypeSystemAssignedUserAssigned = ManagedServiceIdentityStatusType("SystemAssigned,UserAssigned")
	ManagedServiceIdentityStatusTypeUserAssigned               = ManagedServiceIdentityStatusType("UserAssigned")
)

type ManagedServiceIdentity_Status_UserAssignedIdentitiesARM struct {
	//ClientId: The client id of user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	//PrincipalId: The principal id of user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

type PrivateEndpointConnection_Status_SubResourceEmbeddedARM struct {
	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`
}

type VirtualNetworkRule_StatusARM struct {
	//Id: Resource ID of a subnet, for example:
	///subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.
	Id *string `json:"id,omitempty"`

	//IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVNetServiceEndpoint *bool `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}
