// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210702

type IotHub_STATUS_ARM struct {
	// Etag: The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per
	// the normal ETag convention.
	Etag *string `json:"etag,omitempty"`

	// Id: The resource identifier.
	Id *string `json:"id,omitempty"`

	// Identity: The managed identities for the IotHub.
	Identity *ArmIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Location: The resource location.
	Location *string `json:"location,omitempty"`

	// Name: The resource name.
	Name *string `json:"name,omitempty"`

	// Properties: IotHub properties
	Properties *IotHubProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: IotHub SKU info
	Sku *IotHubSkuInfo_STATUS_ARM `json:"sku,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: The resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The resource type.
	Type *string `json:"type,omitempty"`
}

type ArmIdentity_STATUS_ARM struct {
	// PrincipalId: Principal Id
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an implicitly
	// created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
	Type                   *ArmIdentity_Type_STATUS              `json:"type,omitempty"`
	UserAssignedIdentities map[string]ArmUserIdentity_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// The properties of an IoT hub.
type IotHubProperties_STATUS_ARM struct {
	// AllowedFqdnList: List of allowed FQDNs(Fully Qualified Domain Name) for egress from Iot Hub.
	AllowedFqdnList []string `json:"allowedFqdnList,omitempty"`

	// AuthorizationPolicies: The shared access policies you can use to secure a connection to the IoT hub.
	AuthorizationPolicies []SharedAccessSignatureAuthorizationRule_STATUS_ARM `json:"authorizationPolicies,omitempty"`

	// CloudToDevice: The IoT hub cloud-to-device messaging properties.
	CloudToDevice *CloudToDeviceProperties_STATUS_ARM `json:"cloudToDevice,omitempty"`

	// Comments: IoT hub comments.
	Comments *string `json:"comments,omitempty"`

	// DisableDeviceSAS: If true, all device(including Edge devices but excluding modules) scoped SAS keys cannot be used for
	// authentication.
	DisableDeviceSAS *bool `json:"disableDeviceSAS,omitempty"`

	// DisableLocalAuth: If true, SAS tokens with Iot hub scoped SAS keys cannot be used for authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// DisableModuleSAS: If true, all module scoped SAS keys cannot be used for authentication.
	DisableModuleSAS *bool `json:"disableModuleSAS,omitempty"`

	// EnableDataResidency: This property when set to true, will enable data residency, thus, disabling disaster recovery.
	EnableDataResidency *bool `json:"enableDataResidency,omitempty"`

	// EnableFileUploadNotifications: If True, file upload notifications are enabled.
	EnableFileUploadNotifications *bool `json:"enableFileUploadNotifications,omitempty"`

	// EventHubEndpoints: The Event Hub-compatible endpoint properties. The only possible keys to this dictionary is events.
	// This key has to be present in the dictionary while making create or update calls for the IoT hub.
	EventHubEndpoints map[string]EventHubProperties_STATUS_ARM `json:"eventHubEndpoints,omitempty"`

	// Features: The capabilities and features enabled for the IoT hub.
	Features *IotHubProperties_Features_STATUS `json:"features,omitempty"`

	// HostName: The name of the host.
	HostName *string `json:"hostName,omitempty"`

	// IpFilterRules: The IP filter rules.
	IpFilterRules []IpFilterRule_STATUS_ARM `json:"ipFilterRules,omitempty"`

	// Locations: Primary and secondary location for iot hub
	Locations []IotHubLocationDescription_STATUS_ARM `json:"locations,omitempty"`

	// MessagingEndpoints: The messaging endpoint properties for the file upload notification queue.
	MessagingEndpoints map[string]MessagingEndpointProperties_STATUS_ARM `json:"messagingEndpoints,omitempty"`

	// MinTlsVersion: Specifies the minimum TLS version to support for this hub. Can be set to "1.2" to have clients that use a
	// TLS version below 1.2 to be rejected.
	MinTlsVersion *string `json:"minTlsVersion,omitempty"`

	// NetworkRuleSets: Network Rule Set Properties of IotHub
	NetworkRuleSets *NetworkRuleSetProperties_STATUS_ARM `json:"networkRuleSets,omitempty"`

	// PrivateEndpointConnections: Private endpoint connections created on this IotHub
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: The provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether requests from Public Network are allowed
	PublicNetworkAccess *IotHubProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// RestrictOutboundNetworkAccess: If true, egress from IotHub will be restricted to only the allowed FQDNs that are
	// configured via allowedFqdnList.
	RestrictOutboundNetworkAccess *bool `json:"restrictOutboundNetworkAccess,omitempty"`

	// Routing: The routing related properties of the IoT hub. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging
	Routing *RoutingProperties_STATUS_ARM `json:"routing,omitempty"`

	// State: The hub state.
	State *string `json:"state,omitempty"`

	// StorageEndpoints: The list of Azure Storage endpoints where you can upload files. Currently you can configure only one
	// Azure Storage account and that MUST have its key as $default. Specifying more than one storage account causes an error
	// to be thrown. Not specifying a value for this property when the enableFileUploadNotifications property is set to True,
	// causes an error to be thrown.
	StorageEndpoints map[string]StorageEndpointProperties_STATUS_ARM `json:"storageEndpoints,omitempty"`
}

// Information about the SKU of the IoT hub.
type IotHubSkuInfo_STATUS_ARM struct {
	// Capacity: The number of provisioned IoT Hub units. See:
	// https://docs.microsoft.com/azure/azure-subscription-service-limits#iot-hub-limits.
	Capacity *int `json:"capacity,omitempty"`

	// Name: The name of the SKU.
	Name *IotHubSkuInfo_Name_STATUS `json:"name,omitempty"`

	// Tier: The billing tier for the IoT hub.
	Tier *IotHubSkuInfo_Tier_STATUS `json:"tier,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

type ArmIdentity_Type_STATUS string

const (
	ArmIdentity_Type_STATUS_None                       = ArmIdentity_Type_STATUS("None")
	ArmIdentity_Type_STATUS_SystemAssigned             = ArmIdentity_Type_STATUS("SystemAssigned")
	ArmIdentity_Type_STATUS_SystemAssignedUserAssigned = ArmIdentity_Type_STATUS("SystemAssigned, UserAssigned")
	ArmIdentity_Type_STATUS_UserAssigned               = ArmIdentity_Type_STATUS("UserAssigned")
)

type ArmUserIdentity_STATUS_ARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// The IoT hub cloud-to-device messaging properties.
type CloudToDeviceProperties_STATUS_ARM struct {
	// DefaultTtlAsIso8601: The default time to live for cloud-to-device messages in the device queue. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
	DefaultTtlAsIso8601 *string `json:"defaultTtlAsIso8601,omitempty"`

	// Feedback: The properties of the feedback queue for cloud-to-device messages.
	Feedback *FeedbackProperties_STATUS_ARM `json:"feedback,omitempty"`

	// MaxDeliveryCount: The max delivery count for cloud-to-device messages in the device queue. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
	MaxDeliveryCount *int `json:"maxDeliveryCount,omitempty"`
}

// The properties of the provisioned Event Hub-compatible endpoint used by the IoT hub.
type EventHubProperties_STATUS_ARM struct {
	// Endpoint: The Event Hub-compatible endpoint.
	Endpoint *string `json:"endpoint,omitempty"`

	// PartitionCount: The number of partitions for receiving device-to-cloud messages in the Event Hub-compatible endpoint.
	// See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#device-to-cloud-messages.
	PartitionCount *int `json:"partitionCount,omitempty"`

	// PartitionIds: The partition ids in the Event Hub-compatible endpoint.
	PartitionIds []string `json:"partitionIds,omitempty"`

	// Path: The Event Hub-compatible name.
	Path *string `json:"path,omitempty"`

	// RetentionTimeInDays: The retention time for device-to-cloud messages in days. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#device-to-cloud-messages
	RetentionTimeInDays *int `json:"retentionTimeInDays,omitempty"`
}

// Public representation of one of the locations where a resource is provisioned.
type IotHubLocationDescription_STATUS_ARM struct {
	// Location: The name of the Azure region
	Location *string `json:"location,omitempty"`

	// Role: The role of the region, can be either primary or secondary. The primary region is where the IoT hub is currently
	// provisioned. The secondary region is the Azure disaster recovery (DR) paired region and also the region where the IoT
	// hub can failover to.
	Role *IotHubLocationDescription_Role_STATUS `json:"role,omitempty"`
}

type IotHubProperties_Features_STATUS string

const (
	IotHubProperties_Features_STATUS_DeviceManagement = IotHubProperties_Features_STATUS("DeviceManagement")
	IotHubProperties_Features_STATUS_None             = IotHubProperties_Features_STATUS("None")
)

type IotHubProperties_PublicNetworkAccess_STATUS string

const (
	IotHubProperties_PublicNetworkAccess_STATUS_Disabled = IotHubProperties_PublicNetworkAccess_STATUS("Disabled")
	IotHubProperties_PublicNetworkAccess_STATUS_Enabled  = IotHubProperties_PublicNetworkAccess_STATUS("Enabled")
)

type IotHubSkuInfo_Name_STATUS string

const (
	IotHubSkuInfo_Name_STATUS_B1 = IotHubSkuInfo_Name_STATUS("B1")
	IotHubSkuInfo_Name_STATUS_B2 = IotHubSkuInfo_Name_STATUS("B2")
	IotHubSkuInfo_Name_STATUS_B3 = IotHubSkuInfo_Name_STATUS("B3")
	IotHubSkuInfo_Name_STATUS_F1 = IotHubSkuInfo_Name_STATUS("F1")
	IotHubSkuInfo_Name_STATUS_S1 = IotHubSkuInfo_Name_STATUS("S1")
	IotHubSkuInfo_Name_STATUS_S2 = IotHubSkuInfo_Name_STATUS("S2")
	IotHubSkuInfo_Name_STATUS_S3 = IotHubSkuInfo_Name_STATUS("S3")
)

type IotHubSkuInfo_Tier_STATUS string

const (
	IotHubSkuInfo_Tier_STATUS_Basic    = IotHubSkuInfo_Tier_STATUS("Basic")
	IotHubSkuInfo_Tier_STATUS_Free     = IotHubSkuInfo_Tier_STATUS("Free")
	IotHubSkuInfo_Tier_STATUS_Standard = IotHubSkuInfo_Tier_STATUS("Standard")
)

// The IP filter rules for the IoT hub.
type IpFilterRule_STATUS_ARM struct {
	// Action: The desired action for requests captured by this rule.
	Action *IpFilterRule_Action_STATUS `json:"action,omitempty"`

	// FilterName: The name of the IP filter rule.
	FilterName *string `json:"filterName,omitempty"`

	// IpMask: A string that contains the IP address range in CIDR notation for the rule.
	IpMask *string `json:"ipMask,omitempty"`
}

// The properties of the messaging endpoints used by this IoT hub.
type MessagingEndpointProperties_STATUS_ARM struct {
	// LockDurationAsIso8601: The lock duration. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-file-upload.
	LockDurationAsIso8601 *string `json:"lockDurationAsIso8601,omitempty"`

	// MaxDeliveryCount: The number of times the IoT hub attempts to deliver a message. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-file-upload.
	MaxDeliveryCount *int `json:"maxDeliveryCount,omitempty"`

	// TtlAsIso8601: The period of time for which a message is available to consume before it is expired by the IoT hub. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-file-upload.
	TtlAsIso8601 *string `json:"ttlAsIso8601,omitempty"`
}

// Network Rule Set Properties of IotHub
type NetworkRuleSetProperties_STATUS_ARM struct {
	// ApplyToBuiltInEventHubEndpoint: If True, then Network Rule Set is also applied to BuiltIn EventHub EndPoint of IotHub
	ApplyToBuiltInEventHubEndpoint *bool `json:"applyToBuiltInEventHubEndpoint,omitempty"`

	// DefaultAction: Default Action for Network Rule Set
	DefaultAction *NetworkRuleSetProperties_DefaultAction_STATUS `json:"defaultAction,omitempty"`

	// IpRules: List of IP Rules
	IpRules []NetworkRuleSetIpRule_STATUS_ARM `json:"ipRules,omitempty"`
}

// The private endpoint connection of an IotHub
type PrivateEndpointConnection_STATUS_ARM struct {
	// Id: The resource identifier.
	Id *string `json:"id,omitempty"`
}

// The routing related properties of the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging
type RoutingProperties_STATUS_ARM struct {
	// Endpoints: The properties related to the custom endpoints to which your IoT hub routes messages based on the routing
	// rules. A  maximum of 10 custom endpoints are allowed across all endpoint types for paid hubs and only 1 custom endpoint
	// is allowed  across all endpoint types for free hubs.
	Endpoints *RoutingEndpoints_STATUS_ARM `json:"endpoints,omitempty"`

	// Enrichments: The list of user-provided enrichments that the IoT hub applies to messages to be delivered to built-in and
	// custom endpoints. See: https://aka.ms/telemetryoneventgrid
	Enrichments []EnrichmentProperties_STATUS_ARM `json:"enrichments,omitempty"`

	// FallbackRoute: The properties of the route that is used as a fall-back route when none of the conditions specified in
	// the 'routes' section are met. This is an optional parameter. When this property is not present in the template, the
	// fallback route is disabled by default.
	FallbackRoute *FallbackRouteProperties_STATUS_ARM `json:"fallbackRoute,omitempty"`

	// Routes: The list of user-provided routing rules that the IoT hub uses to route messages to built-in and custom
	// endpoints. A maximum of 100 routing rules are allowed for paid hubs and a maximum of 5 routing rules are allowed for
	// free hubs.
	Routes []RouteProperties_STATUS_ARM `json:"routes,omitempty"`
}

// The properties of an IoT hub shared access policy.
type SharedAccessSignatureAuthorizationRule_STATUS_ARM struct {
	// KeyName: The name of the shared access policy.
	KeyName *string `json:"keyName,omitempty"`

	// Rights: The permissions assigned to the shared access policy.
	Rights *SharedAccessSignatureAuthorizationRule_Rights_STATUS `json:"rights,omitempty"`
}

// The properties of the Azure Storage endpoint for file upload.
type StorageEndpointProperties_STATUS_ARM struct {
	// AuthenticationType: Specifies authentication type being used for connecting to the storage account.
	AuthenticationType *StorageEndpointProperties_AuthenticationType_STATUS `json:"authenticationType,omitempty"`

	// ConnectionString: The connection string for the Azure Storage account to which files are uploaded.
	ConnectionString *string `json:"connectionString,omitempty"`

	// ContainerName: The name of the root container where you upload files. The container need not exist but should be
	// creatable using the connectionString specified.
	ContainerName *string `json:"containerName,omitempty"`

	// Identity: Managed identity properties of storage endpoint for file upload.
	Identity *ManagedIdentity_STATUS_ARM `json:"identity,omitempty"`

	// SasTtlAsIso8601: The period of time for which the SAS URI generated by IoT Hub for file upload is valid. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-file-upload#file-upload-notification-configuration-options.
	SasTtlAsIso8601 *string `json:"sasTtlAsIso8601,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// The properties of an enrichment that your IoT hub applies to messages delivered to endpoints.
type EnrichmentProperties_STATUS_ARM struct {
	// EndpointNames: The list of endpoints for which the enrichment is applied to the message.
	EndpointNames []string `json:"endpointNames,omitempty"`

	// Key: The key or name for the enrichment property.
	Key *string `json:"key,omitempty"`

	// Value: The value for the enrichment property.
	Value *string `json:"value,omitempty"`
}

// The properties of the fallback route. IoT Hub uses these properties when it routes messages to the fallback endpoint.
type FallbackRouteProperties_STATUS_ARM struct {
	// Condition: The condition which is evaluated in order to apply the fallback route. If the condition is not provided it
	// will evaluate to true by default. For grammar, See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language
	Condition *string `json:"condition,omitempty"`

	// EndpointNames: The list of endpoints to which the messages that satisfy the condition are routed to. Currently only 1
	// endpoint is allowed.
	EndpointNames []string `json:"endpointNames,omitempty"`

	// IsEnabled: Used to specify whether the fallback route is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Name: The name of the route. The name can only include alphanumeric characters, periods, underscores, hyphens, has a
	// maximum length of 64 characters, and must be unique.
	Name *string `json:"name,omitempty"`

	// Source: The source to which the routing rule is to be applied to. For example, DeviceMessages
	Source *FallbackRouteProperties_Source_STATUS `json:"source,omitempty"`
}

// The properties of the feedback queue for cloud-to-device messages.
type FeedbackProperties_STATUS_ARM struct {
	// LockDurationAsIso8601: The lock duration for the feedback queue. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
	LockDurationAsIso8601 *string `json:"lockDurationAsIso8601,omitempty"`

	// MaxDeliveryCount: The number of times the IoT hub attempts to deliver a message on the feedback queue. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
	MaxDeliveryCount *int `json:"maxDeliveryCount,omitempty"`

	// TtlAsIso8601: The period of time for which a message is available to consume before it is expired by the IoT hub. See:
	// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
	TtlAsIso8601 *string `json:"ttlAsIso8601,omitempty"`
}

type IotHubLocationDescription_Role_STATUS string

const (
	IotHubLocationDescription_Role_STATUS_Primary   = IotHubLocationDescription_Role_STATUS("primary")
	IotHubLocationDescription_Role_STATUS_Secondary = IotHubLocationDescription_Role_STATUS("secondary")
)

type IpFilterRule_Action_STATUS string

const (
	IpFilterRule_Action_STATUS_Accept = IpFilterRule_Action_STATUS("Accept")
	IpFilterRule_Action_STATUS_Reject = IpFilterRule_Action_STATUS("Reject")
)

// The properties of the Managed identity.
type ManagedIdentity_STATUS_ARM struct {
	// UserAssignedIdentity: The user assigned identity.
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// IP Rule to be applied as part of Network Rule Set
type NetworkRuleSetIpRule_STATUS_ARM struct {
	// Action: IP Filter Action
	Action *NetworkRuleSetIpRule_Action_STATUS `json:"action,omitempty"`

	// FilterName: Name of the IP filter rule.
	FilterName *string `json:"filterName,omitempty"`

	// IpMask: A string that contains the IP address range in CIDR notation for the rule.
	IpMask *string `json:"ipMask,omitempty"`
}

type NetworkRuleSetProperties_DefaultAction_STATUS string

const (
	NetworkRuleSetProperties_DefaultAction_STATUS_Allow = NetworkRuleSetProperties_DefaultAction_STATUS("Allow")
	NetworkRuleSetProperties_DefaultAction_STATUS_Deny  = NetworkRuleSetProperties_DefaultAction_STATUS("Deny")
)

// The properties of a routing rule that your IoT hub uses to route messages to endpoints.
type RouteProperties_STATUS_ARM struct {
	// Condition: The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true
	// by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language
	Condition *string `json:"condition,omitempty"`

	// EndpointNames: The list of endpoints to which messages that satisfy the condition are routed. Currently only one
	// endpoint is allowed.
	EndpointNames []string `json:"endpointNames,omitempty"`

	// IsEnabled: Used to specify whether a route is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Name: The name of the route. The name can only include alphanumeric characters, periods, underscores, hyphens, has a
	// maximum length of 64 characters, and must be unique.
	Name *string `json:"name,omitempty"`

	// Source: The source that the routing rule is to be applied to, such as DeviceMessages.
	Source *RouteProperties_Source_STATUS `json:"source,omitempty"`
}

// The properties related to the custom endpoints to which your IoT hub routes messages based on the routing rules. A
// maximum of 10 custom endpoints are allowed across all endpoint types for paid hubs and only 1 custom endpoint is allowed
// across all endpoint types for free hubs.
type RoutingEndpoints_STATUS_ARM struct {
	// EventHubs: The list of Event Hubs endpoints that IoT hub routes messages to, based on the routing rules. This list does
	// not include the built-in Event Hubs endpoint.
	EventHubs []RoutingEventHubProperties_STATUS_ARM `json:"eventHubs,omitempty"`

	// ServiceBusQueues: The list of Service Bus queue endpoints that IoT hub routes the messages to, based on the routing
	// rules.
	ServiceBusQueues []RoutingServiceBusQueueEndpointProperties_STATUS_ARM `json:"serviceBusQueues,omitempty"`

	// ServiceBusTopics: The list of Service Bus topic endpoints that the IoT hub routes the messages to, based on the routing
	// rules.
	ServiceBusTopics []RoutingServiceBusTopicEndpointProperties_STATUS_ARM `json:"serviceBusTopics,omitempty"`

	// StorageContainers: The list of storage container endpoints that IoT hub routes messages to, based on the routing rules.
	StorageContainers []RoutingStorageContainerProperties_STATUS_ARM `json:"storageContainers,omitempty"`
}

type SharedAccessSignatureAuthorizationRule_Rights_STATUS string

const (
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_DeviceConnect                                        = SharedAccessSignatureAuthorizationRule_Rights_STATUS("DeviceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryRead                                         = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryRead")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryReadDeviceConnect                            = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryRead, DeviceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryReadRegistryWrite                            = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryRead, RegistryWrite")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryReadRegistryWriteDeviceConnect               = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryRead, RegistryWrite, DeviceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryReadRegistryWriteServiceConnect              = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryRead, RegistryWrite, ServiceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryReadRegistryWriteServiceConnectDeviceConnect = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryReadServiceConnect                           = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryRead, ServiceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryReadServiceConnectDeviceConnect              = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryRead, ServiceConnect, DeviceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryWrite                                        = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryWrite")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryWriteDeviceConnect                           = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryWrite, DeviceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryWriteServiceConnect                          = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryWrite, ServiceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_RegistryWriteServiceConnectDeviceConnect             = SharedAccessSignatureAuthorizationRule_Rights_STATUS("RegistryWrite, ServiceConnect, DeviceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_ServiceConnect                                       = SharedAccessSignatureAuthorizationRule_Rights_STATUS("ServiceConnect")
	SharedAccessSignatureAuthorizationRule_Rights_STATUS_ServiceConnectDeviceConnect                          = SharedAccessSignatureAuthorizationRule_Rights_STATUS("ServiceConnect, DeviceConnect")
)

type StorageEndpointProperties_AuthenticationType_STATUS string

const (
	StorageEndpointProperties_AuthenticationType_STATUS_IdentityBased = StorageEndpointProperties_AuthenticationType_STATUS("identityBased")
	StorageEndpointProperties_AuthenticationType_STATUS_KeyBased      = StorageEndpointProperties_AuthenticationType_STATUS("keyBased")
)

type FallbackRouteProperties_Source_STATUS string

const FallbackRouteProperties_Source_STATUS_DeviceMessages = FallbackRouteProperties_Source_STATUS("DeviceMessages")

type NetworkRuleSetIpRule_Action_STATUS string

const NetworkRuleSetIpRule_Action_STATUS_Allow = NetworkRuleSetIpRule_Action_STATUS("Allow")

type RouteProperties_Source_STATUS string

const (
	RouteProperties_Source_STATUS_DeviceConnectionStateEvents = RouteProperties_Source_STATUS("DeviceConnectionStateEvents")
	RouteProperties_Source_STATUS_DeviceJobLifecycleEvents    = RouteProperties_Source_STATUS("DeviceJobLifecycleEvents")
	RouteProperties_Source_STATUS_DeviceLifecycleEvents       = RouteProperties_Source_STATUS("DeviceLifecycleEvents")
	RouteProperties_Source_STATUS_DeviceMessages              = RouteProperties_Source_STATUS("DeviceMessages")
	RouteProperties_Source_STATUS_Invalid                     = RouteProperties_Source_STATUS("Invalid")
	RouteProperties_Source_STATUS_TwinChangeEvents            = RouteProperties_Source_STATUS("TwinChangeEvents")
)

// The properties related to an event hub endpoint.
type RoutingEventHubProperties_STATUS_ARM struct {
	// AuthenticationType: Method used to authenticate against the event hub endpoint
	AuthenticationType *RoutingEventHubProperties_AuthenticationType_STATUS `json:"authenticationType,omitempty"`

	// ConnectionString: The connection string of the event hub endpoint.
	ConnectionString *string `json:"connectionString,omitempty"`

	// EndpointUri: The url of the event hub endpoint. It must include the protocol sb://
	EndpointUri *string `json:"endpointUri,omitempty"`

	// EntityPath: Event hub name on the event hub namespace
	EntityPath *string `json:"entityPath,omitempty"`

	// Id: Id of the event hub endpoint
	Id *string `json:"id,omitempty"`

	// Identity: Managed identity properties of routing event hub endpoint.
	Identity *ManagedIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Name: The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores,
	// hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications,
	// $default. Endpoint names must be unique across endpoint types.
	Name *string `json:"name,omitempty"`

	// ResourceGroup: The name of the resource group of the event hub endpoint.
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// SubscriptionId: The subscription identifier of the event hub endpoint.
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// The properties related to service bus queue endpoint types.
type RoutingServiceBusQueueEndpointProperties_STATUS_ARM struct {
	// AuthenticationType: Method used to authenticate against the service bus queue endpoint
	AuthenticationType *RoutingServiceBusQueueEndpointProperties_AuthenticationType_STATUS `json:"authenticationType,omitempty"`

	// ConnectionString: The connection string of the service bus queue endpoint.
	ConnectionString *string `json:"connectionString,omitempty"`

	// EndpointUri: The url of the service bus queue endpoint. It must include the protocol sb://
	EndpointUri *string `json:"endpointUri,omitempty"`

	// EntityPath: Queue name on the service bus namespace
	EntityPath *string `json:"entityPath,omitempty"`

	// Id: Id of the service bus queue endpoint
	Id *string `json:"id,omitempty"`

	// Identity: Managed identity properties of routing service bus queue endpoint.
	Identity *ManagedIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Name: The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores,
	// hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications,
	// $default. Endpoint names must be unique across endpoint types. The name need not be the same as the actual queue name.
	Name *string `json:"name,omitempty"`

	// ResourceGroup: The name of the resource group of the service bus queue endpoint.
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// SubscriptionId: The subscription identifier of the service bus queue endpoint.
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// The properties related to service bus topic endpoint types.
type RoutingServiceBusTopicEndpointProperties_STATUS_ARM struct {
	// AuthenticationType: Method used to authenticate against the service bus topic endpoint
	AuthenticationType *RoutingServiceBusTopicEndpointProperties_AuthenticationType_STATUS `json:"authenticationType,omitempty"`

	// ConnectionString: The connection string of the service bus topic endpoint.
	ConnectionString *string `json:"connectionString,omitempty"`

	// EndpointUri: The url of the service bus topic endpoint. It must include the protocol sb://
	EndpointUri *string `json:"endpointUri,omitempty"`

	// EntityPath: Queue name on the service bus topic
	EntityPath *string `json:"entityPath,omitempty"`

	// Id: Id of the service bus topic endpoint
	Id *string `json:"id,omitempty"`

	// Identity: Managed identity properties of routing service bus topic endpoint.
	Identity *ManagedIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Name: The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores,
	// hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications,
	// $default. Endpoint names must be unique across endpoint types.  The name need not be the same as the actual topic name.
	Name *string `json:"name,omitempty"`

	// ResourceGroup: The name of the resource group of the service bus topic endpoint.
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// SubscriptionId: The subscription identifier of the service bus topic endpoint.
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// The properties related to a storage container endpoint.
type RoutingStorageContainerProperties_STATUS_ARM struct {
	// AuthenticationType: Method used to authenticate against the storage endpoint
	AuthenticationType *RoutingStorageContainerProperties_AuthenticationType_STATUS `json:"authenticationType,omitempty"`

	// BatchFrequencyInSeconds: Time interval at which blobs are written to storage. Value should be between 60 and 720
	// seconds. Default value is 300 seconds.
	BatchFrequencyInSeconds *int `json:"batchFrequencyInSeconds,omitempty"`

	// ConnectionString: The connection string of the storage account.
	ConnectionString *string `json:"connectionString,omitempty"`

	// ContainerName: The name of storage container in the storage account.
	ContainerName *string `json:"containerName,omitempty"`

	// Encoding: Encoding that is used to serialize messages to blobs. Supported values are 'avro', 'avrodeflate', and 'JSON'.
	// Default value is 'avro'.
	Encoding *RoutingStorageContainerProperties_Encoding_STATUS `json:"encoding,omitempty"`

	// EndpointUri: The url of the storage endpoint. It must include the protocol https://
	EndpointUri *string `json:"endpointUri,omitempty"`

	// FileNameFormat: File name format for the blob. Default format is {iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}. All
	// parameters are mandatory but can be reordered.
	FileNameFormat *string `json:"fileNameFormat,omitempty"`

	// Id: Id of the storage container endpoint
	Id *string `json:"id,omitempty"`

	// Identity: Managed identity properties of routing storage endpoint.
	Identity *ManagedIdentity_STATUS_ARM `json:"identity,omitempty"`

	// MaxChunkSizeInBytes: Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB)
	// and 524288000(500MB). Default value is 314572800(300MB).
	MaxChunkSizeInBytes *int `json:"maxChunkSizeInBytes,omitempty"`

	// Name: The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores,
	// hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications,
	// $default. Endpoint names must be unique across endpoint types.
	Name *string `json:"name,omitempty"`

	// ResourceGroup: The name of the resource group of the storage account.
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// SubscriptionId: The subscription identifier of the storage account.
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

type RoutingEventHubProperties_AuthenticationType_STATUS string

const (
	RoutingEventHubProperties_AuthenticationType_STATUS_IdentityBased = RoutingEventHubProperties_AuthenticationType_STATUS("identityBased")
	RoutingEventHubProperties_AuthenticationType_STATUS_KeyBased      = RoutingEventHubProperties_AuthenticationType_STATUS("keyBased")
)

type RoutingServiceBusQueueEndpointProperties_AuthenticationType_STATUS string

const (
	RoutingServiceBusQueueEndpointProperties_AuthenticationType_STATUS_IdentityBased = RoutingServiceBusQueueEndpointProperties_AuthenticationType_STATUS("identityBased")
	RoutingServiceBusQueueEndpointProperties_AuthenticationType_STATUS_KeyBased      = RoutingServiceBusQueueEndpointProperties_AuthenticationType_STATUS("keyBased")
)

type RoutingServiceBusTopicEndpointProperties_AuthenticationType_STATUS string

const (
	RoutingServiceBusTopicEndpointProperties_AuthenticationType_STATUS_IdentityBased = RoutingServiceBusTopicEndpointProperties_AuthenticationType_STATUS("identityBased")
	RoutingServiceBusTopicEndpointProperties_AuthenticationType_STATUS_KeyBased      = RoutingServiceBusTopicEndpointProperties_AuthenticationType_STATUS("keyBased")
)

type RoutingStorageContainerProperties_AuthenticationType_STATUS string

const (
	RoutingStorageContainerProperties_AuthenticationType_STATUS_IdentityBased = RoutingStorageContainerProperties_AuthenticationType_STATUS("identityBased")
	RoutingStorageContainerProperties_AuthenticationType_STATUS_KeyBased      = RoutingStorageContainerProperties_AuthenticationType_STATUS("keyBased")
)

type RoutingStorageContainerProperties_Encoding_STATUS string

const (
	RoutingStorageContainerProperties_Encoding_STATUS_Avro        = RoutingStorageContainerProperties_Encoding_STATUS("Avro")
	RoutingStorageContainerProperties_Encoding_STATUS_AvroDeflate = RoutingStorageContainerProperties_Encoding_STATUS("AvroDeflate")
	RoutingStorageContainerProperties_Encoding_STATUS_JSON        = RoutingStorageContainerProperties_Encoding_STATUS("JSON")
)
