// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "encoding/json"

// An object that represents a machine learning workspace.
type Workspace_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the resource.
	Identity *ManagedServiceIdentity_STATUS `json:"identity,omitempty"`
	Kind     *string                        `json:"kind,omitempty"`

	// Location: Specifies the location of the resource.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the machine learning workspace.
	Properties *WorkspaceProperties_STATUS `json:"properties,omitempty"`

	// Sku: The sku of the workspace.
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Contains resource tags defined as key/value pairs.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of a machine learning workspace.
type WorkspaceProperties_STATUS struct {
	// AllowPublicAccessWhenBehindVnet: The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet *bool `json:"allowPublicAccessWhenBehindVnet,omitempty"`

	// ApplicationInsights: ARM id of the application insights associated with this workspace.
	ApplicationInsights  *string  `json:"applicationInsights,omitempty"`
	AssociatedWorkspaces []string `json:"associatedWorkspaces,omitempty"`

	// ContainerRegistry: ARM id of the container registry associated with this workspace.
	ContainerRegistry *string `json:"containerRegistry,omitempty"`

	// Description: The description of this workspace.
	Description *string `json:"description,omitempty"`

	// DiscoveryUrl: Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl        *string `json:"discoveryUrl,omitempty"`
	EnableDataIsolation *bool   `json:"enableDataIsolation,omitempty"`

	// Encryption: The encryption settings of Azure ML workspace.
	Encryption *EncryptionProperty_STATUS `json:"encryption,omitempty"`

	// FeatureStoreSettings: Settings for feature store type workspace.
	FeatureStoreSettings *FeatureStoreSettings_STATUS `json:"featureStoreSettings,omitempty"`

	// FriendlyName: The friendly name for this workspace. This name in mutable
	FriendlyName *string `json:"friendlyName,omitempty"`

	// HbiWorkspace: The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace  *bool   `json:"hbiWorkspace,omitempty"`
	HubResourceId *string `json:"hubResourceId,omitempty"`

	// ImageBuildCompute: The compute name for image build
	ImageBuildCompute *string `json:"imageBuildCompute,omitempty"`

	// KeyVault: ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been
	// created
	KeyVault *string `json:"keyVault,omitempty"`

	// ManagedNetwork: Managed Network settings for a machine learning workspace.
	ManagedNetwork *ManagedNetworkSettings_STATUS `json:"managedNetwork,omitempty"`

	// MlFlowTrackingUri: The URI associated with this workspace that machine learning flow must point at to set up tracking.
	MlFlowTrackingUri *string `json:"mlFlowTrackingUri,omitempty"`

	// NotebookInfo: The notebook info of Azure ML workspace.
	NotebookInfo *NotebookResourceInfo_STATUS `json:"notebookInfo,omitempty"`

	// PrimaryUserAssignedIdentity: The user assigned identity resource id that represents the workspace identity.
	PrimaryUserAssignedIdentity *string `json:"primaryUserAssignedIdentity,omitempty"`

	// PrivateEndpointConnections: The list of private endpoint connections in the workspace.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// PrivateLinkCount: Count of private connections in the workspace
	PrivateLinkCount *int `json:"privateLinkCount,omitempty"`

	// ProvisioningState: The current deployment state of workspace resource. The provisioningState is to indicate states for
	// resource provisioning.
	ProvisioningState *WorkspaceProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether requests from Public Network are allowed.
	PublicNetworkAccess *WorkspaceProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// ServerlessComputeSettings: Settings for serverless compute created in the workspace
	ServerlessComputeSettings *ServerlessComputeSettings_STATUS `json:"serverlessComputeSettings,omitempty"`

	// ServiceManagedResourcesSettings: The service managed resource settings.
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings_STATUS `json:"serviceManagedResourcesSettings,omitempty"`

	// ServiceProvisionedResourceGroup: The name of the managed resource group created by workspace RP in customer subscription
	// if the workspace is CMK workspace
	ServiceProvisionedResourceGroup *string `json:"serviceProvisionedResourceGroup,omitempty"`

	// SharedPrivateLinkResources: The list of shared private link resources in this workspace.
	SharedPrivateLinkResources []SharedPrivateLinkResource_STATUS `json:"sharedPrivateLinkResources,omitempty"`

	// StorageAccount: ARM id of the storage account associated with this workspace. This cannot be changed once the workspace
	// has been created
	StorageAccount *string `json:"storageAccount,omitempty"`

	// StorageHnsEnabled: If the storage associated with the workspace has hierarchical namespace(HNS) enabled.
	StorageHnsEnabled *bool `json:"storageHnsEnabled,omitempty"`

	// TenantId: The tenant id associated with this workspace.
	TenantId *string `json:"tenantId,omitempty"`

	// V1LegacyMode: Enabling v1_legacy_mode may prevent you from using features provided by the v2 API.
	V1LegacyMode *bool `json:"v1LegacyMode,omitempty"`

	// WorkspaceHubConfig: WorkspaceHub's configuration object.
	WorkspaceHubConfig *WorkspaceHubConfig_STATUS `json:"workspaceHubConfig,omitempty"`

	// WorkspaceId: The immutable id associated with this workspace.
	WorkspaceId *string `json:"workspaceId,omitempty"`
}

type EncryptionProperty_STATUS struct {
	// Identity: The identity that will be used to access the key vault for encryption at rest.
	Identity *IdentityForCmk_STATUS `json:"identity,omitempty"`

	// KeyVaultProperties: Customer Key vault properties.
	KeyVaultProperties *EncryptionKeyVaultProperties_STATUS `json:"keyVaultProperties,omitempty"`

	// Status: Indicates whether or not the encryption is enabled for the workspace.
	Status *EncryptionProperty_Status_STATUS `json:"status,omitempty"`
}

// Settings for feature store type workspace.
type FeatureStoreSettings_STATUS struct {
	// ComputeRuntime: Compute runtime config for feature store type workspace.
	ComputeRuntime             *ComputeRuntimeDto_STATUS `json:"computeRuntime,omitempty"`
	OfflineStoreConnectionName *string                   `json:"offlineStoreConnectionName,omitempty"`
	OnlineStoreConnectionName  *string                   `json:"onlineStoreConnectionName,omitempty"`
}

// Managed Network settings for a machine learning workspace.
type ManagedNetworkSettings_STATUS struct {
	// IsolationMode: Isolation mode for the managed network of a machine learning workspace.
	IsolationMode *IsolationMode_STATUS          `json:"isolationMode,omitempty"`
	NetworkId     *string                        `json:"networkId,omitempty"`
	OutboundRules map[string]OutboundRule_STATUS `json:"outboundRules,omitempty"`

	// Status: Status of the Provisioning for the managed network of a machine learning workspace.
	Status *ManagedNetworkProvisionStatus_STATUS `json:"status,omitempty"`
}

type NotebookResourceInfo_STATUS struct {
	Fqdn *string `json:"fqdn,omitempty"`

	// NotebookPreparationError: The error that occurs when preparing notebook.
	NotebookPreparationError *NotebookPreparationError_STATUS `json:"notebookPreparationError,omitempty"`

	// ResourceId: the data plane resourceId that used to initialize notebook component
	ResourceId *string `json:"resourceId,omitempty"`
}

// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`
}

type ServerlessComputeSettings_STATUS struct {
	// ServerlessComputeCustomSubnet: The resource ID of an existing virtual network subnet in which serverless compute nodes
	// should be deployed
	ServerlessComputeCustomSubnet *string `json:"serverlessComputeCustomSubnet,omitempty"`

	// ServerlessComputeNoPublicIP: The flag to signal if serverless compute nodes deployed in custom vNet would have no public
	// IP addresses for a workspace with private endpoint
	ServerlessComputeNoPublicIP *bool `json:"serverlessComputeNoPublicIP,omitempty"`
}

type ServiceManagedResourcesSettings_STATUS struct {
	// CosmosDb: The settings for the service managed cosmosdb account.
	CosmosDb *CosmosDbSettings_STATUS `json:"cosmosDb,omitempty"`
}

type SharedPrivateLinkResource_STATUS struct {
	// Name: Unique name of the private link.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *SharedPrivateLinkResourceProperty_STATUS `json:"properties,omitempty"`
}

// WorkspaceHub's configuration object.
type WorkspaceHubConfig_STATUS struct {
	AdditionalWorkspaceStorageAccounts []string `json:"additionalWorkspaceStorageAccounts,omitempty"`
	DefaultWorkspaceResourceGroup      *string  `json:"defaultWorkspaceResourceGroup,omitempty"`
}

type WorkspaceProperties_ProvisioningState_STATUS string

const (
	WorkspaceProperties_ProvisioningState_STATUS_Canceled  = WorkspaceProperties_ProvisioningState_STATUS("Canceled")
	WorkspaceProperties_ProvisioningState_STATUS_Creating  = WorkspaceProperties_ProvisioningState_STATUS("Creating")
	WorkspaceProperties_ProvisioningState_STATUS_Deleting  = WorkspaceProperties_ProvisioningState_STATUS("Deleting")
	WorkspaceProperties_ProvisioningState_STATUS_Failed    = WorkspaceProperties_ProvisioningState_STATUS("Failed")
	WorkspaceProperties_ProvisioningState_STATUS_Succeeded = WorkspaceProperties_ProvisioningState_STATUS("Succeeded")
	WorkspaceProperties_ProvisioningState_STATUS_Unknown   = WorkspaceProperties_ProvisioningState_STATUS("Unknown")
	WorkspaceProperties_ProvisioningState_STATUS_Updating  = WorkspaceProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to WorkspaceProperties_ProvisioningState_STATUS
var workspaceProperties_ProvisioningState_STATUS_Values = map[string]WorkspaceProperties_ProvisioningState_STATUS{
	"canceled":  WorkspaceProperties_ProvisioningState_STATUS_Canceled,
	"creating":  WorkspaceProperties_ProvisioningState_STATUS_Creating,
	"deleting":  WorkspaceProperties_ProvisioningState_STATUS_Deleting,
	"failed":    WorkspaceProperties_ProvisioningState_STATUS_Failed,
	"succeeded": WorkspaceProperties_ProvisioningState_STATUS_Succeeded,
	"unknown":   WorkspaceProperties_ProvisioningState_STATUS_Unknown,
	"updating":  WorkspaceProperties_ProvisioningState_STATUS_Updating,
}

type WorkspaceProperties_PublicNetworkAccess_STATUS string

const (
	WorkspaceProperties_PublicNetworkAccess_STATUS_Disabled = WorkspaceProperties_PublicNetworkAccess_STATUS("Disabled")
	WorkspaceProperties_PublicNetworkAccess_STATUS_Enabled  = WorkspaceProperties_PublicNetworkAccess_STATUS("Enabled")
)

// Mapping from string to WorkspaceProperties_PublicNetworkAccess_STATUS
var workspaceProperties_PublicNetworkAccess_STATUS_Values = map[string]WorkspaceProperties_PublicNetworkAccess_STATUS{
	"disabled": WorkspaceProperties_PublicNetworkAccess_STATUS_Disabled,
	"enabled":  WorkspaceProperties_PublicNetworkAccess_STATUS_Enabled,
}

// Compute runtime config for feature store type workspace.
type ComputeRuntimeDto_STATUS struct {
	SparkRuntimeVersion *string `json:"sparkRuntimeVersion,omitempty"`
}

type CosmosDbSettings_STATUS struct {
	// CollectionsThroughput: The throughput of the collections in cosmosdb database
	CollectionsThroughput *int `json:"collectionsThroughput,omitempty"`
}

type EncryptionKeyVaultProperties_STATUS struct {
	// IdentityClientId: For future use - The client id of the identity which will be used to access key vault.
	IdentityClientId *string `json:"identityClientId,omitempty"`

	// KeyIdentifier: Key vault uri to access the encryption key.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`

	// KeyVaultArmId: The ArmId of the keyVault where the customer owned encryption key is present.
	KeyVaultArmId *string `json:"keyVaultArmId,omitempty"`
}

type EncryptionProperty_Status_STATUS string

const (
	EncryptionProperty_Status_STATUS_Disabled = EncryptionProperty_Status_STATUS("Disabled")
	EncryptionProperty_Status_STATUS_Enabled  = EncryptionProperty_Status_STATUS("Enabled")
)

// Mapping from string to EncryptionProperty_Status_STATUS
var encryptionProperty_Status_STATUS_Values = map[string]EncryptionProperty_Status_STATUS{
	"disabled": EncryptionProperty_Status_STATUS_Disabled,
	"enabled":  EncryptionProperty_Status_STATUS_Enabled,
}

// Identity that will be used to access key vault for encryption at rest
type IdentityForCmk_STATUS struct {
	// UserAssignedIdentity: The ArmId of the user assigned identity that will be used to access the customer managed key vault
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// Isolation mode for the managed network of a machine learning workspace.
type IsolationMode_STATUS string

const (
	IsolationMode_STATUS_AllowInternetOutbound     = IsolationMode_STATUS("AllowInternetOutbound")
	IsolationMode_STATUS_AllowOnlyApprovedOutbound = IsolationMode_STATUS("AllowOnlyApprovedOutbound")
	IsolationMode_STATUS_Disabled                  = IsolationMode_STATUS("Disabled")
)

// Mapping from string to IsolationMode_STATUS
var isolationMode_STATUS_Values = map[string]IsolationMode_STATUS{
	"allowinternetoutbound":     IsolationMode_STATUS_AllowInternetOutbound,
	"allowonlyapprovedoutbound": IsolationMode_STATUS_AllowOnlyApprovedOutbound,
	"disabled":                  IsolationMode_STATUS_Disabled,
}

// Status of the Provisioning for the managed network of a machine learning workspace.
type ManagedNetworkProvisionStatus_STATUS struct {
	SparkReady *bool `json:"sparkReady,omitempty"`

	// Status: Status for the managed network of a machine learning workspace.
	Status *ManagedNetworkStatus_STATUS `json:"status,omitempty"`
}

type NotebookPreparationError_STATUS struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	StatusCode   *int    `json:"statusCode,omitempty"`
}

type OutboundRule_STATUS struct {
	// FQDN: Mutually exclusive with all other properties
	FQDN *FqdnOutboundRule_STATUS `json:"fqdn,omitempty"`

	// PrivateEndpoint: Mutually exclusive with all other properties
	PrivateEndpoint *PrivateEndpointOutboundRule_STATUS `json:"privateEndpoint,omitempty"`

	// ServiceTag: Mutually exclusive with all other properties
	ServiceTag *ServiceTagOutboundRule_STATUS `json:"serviceTag,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because OutboundRule_STATUS represents a discriminated union (JSON OneOf)
func (rule OutboundRule_STATUS) MarshalJSON() ([]byte, error) {
	if rule.FQDN != nil {
		return json.Marshal(rule.FQDN)
	}

	if rule.PrivateEndpoint != nil {
		return json.Marshal(rule.PrivateEndpoint)
	}

	if rule.ServiceTag != nil {
		return json.Marshal(rule.ServiceTag)
	}

	return nil, nil
}

// UnmarshalJSON unmarshals the OutboundRule_STATUS
func (rule *OutboundRule_STATUS) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "FQDN" {
		rule.FQDN = &FqdnOutboundRule_STATUS{}
		return json.Unmarshal(data, rule.FQDN)
	}
	if discriminator == "PrivateEndpoint" {
		rule.PrivateEndpoint = &PrivateEndpointOutboundRule_STATUS{}
		return json.Unmarshal(data, rule.PrivateEndpoint)
	}
	if discriminator == "ServiceTag" {
		rule.ServiceTag = &ServiceTagOutboundRule_STATUS{}
		return json.Unmarshal(data, rule.ServiceTag)
	}

	// No error
	return nil
}

// Properties of a shared private link resource.
type SharedPrivateLinkResourceProperty_STATUS struct {
	// GroupId: The private link resource group id.
	GroupId *string `json:"groupId,omitempty"`

	// PrivateLinkResourceId: The resource id that private link links to.
	PrivateLinkResourceId *string `json:"privateLinkResourceId,omitempty"`

	// RequestMessage: Request message.
	RequestMessage *string `json:"requestMessage,omitempty"`

	// Status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus_STATUS `json:"status,omitempty"`
}

type FqdnOutboundRule_STATUS struct {
	// Category: Category of a managed network Outbound Rule of a machine learning workspace.
	Category    *RuleCategory_STATUS `json:"category,omitempty"`
	Destination *string              `json:"destination,omitempty"`

	// Status: Type of a managed network Outbound Rule of a machine learning workspace.
	Status *RuleStatus_STATUS           `json:"status,omitempty"`
	Type   FqdnOutboundRule_Type_STATUS `json:"type,omitempty"`
}

// Status for the managed network of a machine learning workspace.
type ManagedNetworkStatus_STATUS string

const (
	ManagedNetworkStatus_STATUS_Active   = ManagedNetworkStatus_STATUS("Active")
	ManagedNetworkStatus_STATUS_Inactive = ManagedNetworkStatus_STATUS("Inactive")
)

// Mapping from string to ManagedNetworkStatus_STATUS
var managedNetworkStatus_STATUS_Values = map[string]ManagedNetworkStatus_STATUS{
	"active":   ManagedNetworkStatus_STATUS_Active,
	"inactive": ManagedNetworkStatus_STATUS_Inactive,
}

type PrivateEndpointOutboundRule_STATUS struct {
	// Category: Category of a managed network Outbound Rule of a machine learning workspace.
	Category *RuleCategory_STATUS `json:"category,omitempty"`

	// Destination: Private Endpoint destination for a Private Endpoint Outbound Rule for the managed network of a machine
	// learning  workspace.
	Destination *PrivateEndpointDestination_STATUS `json:"destination,omitempty"`

	// Status: Type of a managed network Outbound Rule of a machine learning workspace.
	Status *RuleStatus_STATUS                      `json:"status,omitempty"`
	Type   PrivateEndpointOutboundRule_Type_STATUS `json:"type,omitempty"`
}

// The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus_STATUS string

const (
	PrivateEndpointServiceConnectionStatus_STATUS_Approved     = PrivateEndpointServiceConnectionStatus_STATUS("Approved")
	PrivateEndpointServiceConnectionStatus_STATUS_Disconnected = PrivateEndpointServiceConnectionStatus_STATUS("Disconnected")
	PrivateEndpointServiceConnectionStatus_STATUS_Pending      = PrivateEndpointServiceConnectionStatus_STATUS("Pending")
	PrivateEndpointServiceConnectionStatus_STATUS_Rejected     = PrivateEndpointServiceConnectionStatus_STATUS("Rejected")
	PrivateEndpointServiceConnectionStatus_STATUS_Timeout      = PrivateEndpointServiceConnectionStatus_STATUS("Timeout")
)

// Mapping from string to PrivateEndpointServiceConnectionStatus_STATUS
var privateEndpointServiceConnectionStatus_STATUS_Values = map[string]PrivateEndpointServiceConnectionStatus_STATUS{
	"approved":     PrivateEndpointServiceConnectionStatus_STATUS_Approved,
	"disconnected": PrivateEndpointServiceConnectionStatus_STATUS_Disconnected,
	"pending":      PrivateEndpointServiceConnectionStatus_STATUS_Pending,
	"rejected":     PrivateEndpointServiceConnectionStatus_STATUS_Rejected,
	"timeout":      PrivateEndpointServiceConnectionStatus_STATUS_Timeout,
}

type ServiceTagOutboundRule_STATUS struct {
	// Category: Category of a managed network Outbound Rule of a machine learning workspace.
	Category *RuleCategory_STATUS `json:"category,omitempty"`

	// Destination: Service Tag destination for a Service Tag Outbound Rule for the managed network of a machine learning
	// workspace.
	Destination *ServiceTagDestination_STATUS `json:"destination,omitempty"`

	// Status: Type of a managed network Outbound Rule of a machine learning workspace.
	Status *RuleStatus_STATUS                 `json:"status,omitempty"`
	Type   ServiceTagOutboundRule_Type_STATUS `json:"type,omitempty"`
}

type FqdnOutboundRule_Type_STATUS string

const FqdnOutboundRule_Type_STATUS_FQDN = FqdnOutboundRule_Type_STATUS("FQDN")

// Mapping from string to FqdnOutboundRule_Type_STATUS
var fqdnOutboundRule_Type_STATUS_Values = map[string]FqdnOutboundRule_Type_STATUS{
	"fqdn": FqdnOutboundRule_Type_STATUS_FQDN,
}

// Private Endpoint destination for a Private Endpoint Outbound Rule for the managed network of a machine learning
// workspace.
type PrivateEndpointDestination_STATUS struct {
	ServiceResourceId *string `json:"serviceResourceId,omitempty"`
	SparkEnabled      *bool   `json:"sparkEnabled,omitempty"`

	// SparkStatus: Type of a managed network Outbound Rule of a machine learning workspace.
	SparkStatus       *RuleStatus_STATUS `json:"sparkStatus,omitempty"`
	SubresourceTarget *string            `json:"subresourceTarget,omitempty"`
}

type PrivateEndpointOutboundRule_Type_STATUS string

const PrivateEndpointOutboundRule_Type_STATUS_PrivateEndpoint = PrivateEndpointOutboundRule_Type_STATUS("PrivateEndpoint")

// Mapping from string to PrivateEndpointOutboundRule_Type_STATUS
var privateEndpointOutboundRule_Type_STATUS_Values = map[string]PrivateEndpointOutboundRule_Type_STATUS{
	"privateendpoint": PrivateEndpointOutboundRule_Type_STATUS_PrivateEndpoint,
}

// Category of a managed network Outbound Rule of a machine learning workspace.
type RuleCategory_STATUS string

const (
	RuleCategory_STATUS_Dependency  = RuleCategory_STATUS("Dependency")
	RuleCategory_STATUS_Recommended = RuleCategory_STATUS("Recommended")
	RuleCategory_STATUS_Required    = RuleCategory_STATUS("Required")
	RuleCategory_STATUS_UserDefined = RuleCategory_STATUS("UserDefined")
)

// Mapping from string to RuleCategory_STATUS
var ruleCategory_STATUS_Values = map[string]RuleCategory_STATUS{
	"dependency":  RuleCategory_STATUS_Dependency,
	"recommended": RuleCategory_STATUS_Recommended,
	"required":    RuleCategory_STATUS_Required,
	"userdefined": RuleCategory_STATUS_UserDefined,
}

// Type of a managed network Outbound Rule of a machine learning workspace.
type RuleStatus_STATUS string

const (
	RuleStatus_STATUS_Active   = RuleStatus_STATUS("Active")
	RuleStatus_STATUS_Inactive = RuleStatus_STATUS("Inactive")
)

// Mapping from string to RuleStatus_STATUS
var ruleStatus_STATUS_Values = map[string]RuleStatus_STATUS{
	"active":   RuleStatus_STATUS_Active,
	"inactive": RuleStatus_STATUS_Inactive,
}

// Service Tag destination for a Service Tag Outbound Rule for the managed network of a machine learning workspace.
type ServiceTagDestination_STATUS struct {
	// Action: The action enum for networking rule.
	Action *RuleAction_STATUS `json:"action,omitempty"`

	// AddressPrefixes: Optional, if provided, the ServiceTag property will be ignored.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`
	PortRanges      *string  `json:"portRanges,omitempty"`
	Protocol        *string  `json:"protocol,omitempty"`
	ServiceTag      *string  `json:"serviceTag,omitempty"`
}

type ServiceTagOutboundRule_Type_STATUS string

const ServiceTagOutboundRule_Type_STATUS_ServiceTag = ServiceTagOutboundRule_Type_STATUS("ServiceTag")

// Mapping from string to ServiceTagOutboundRule_Type_STATUS
var serviceTagOutboundRule_Type_STATUS_Values = map[string]ServiceTagOutboundRule_Type_STATUS{
	"servicetag": ServiceTagOutboundRule_Type_STATUS_ServiceTag,
}

// The action enum for networking rule.
type RuleAction_STATUS string

const (
	RuleAction_STATUS_Allow = RuleAction_STATUS("Allow")
	RuleAction_STATUS_Deny  = RuleAction_STATUS("Deny")
)

// Mapping from string to RuleAction_STATUS
var ruleAction_STATUS_Values = map[string]RuleAction_STATUS{
	"allow": RuleAction_STATUS_Allow,
	"deny":  RuleAction_STATUS_Deny,
}
