// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftCluster_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The cluster properties.
	Properties *OpenShiftClusterProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// OpenShiftClusterProperties represents an OpenShift cluster's properties.
type OpenShiftClusterProperties_STATUS struct {
	// ApiserverProfile: The cluster API server profile.
	ApiserverProfile *APIServerProfile_STATUS `json:"apiserverProfile,omitempty"`

	// ClusterProfile: The cluster profile.
	ClusterProfile *ClusterProfile_STATUS `json:"clusterProfile,omitempty"`

	// ConsoleProfile: The console profile.
	ConsoleProfile *ConsoleProfile_STATUS `json:"consoleProfile,omitempty"`

	// IngressProfiles: The cluster ingress profiles.
	IngressProfiles []IngressProfile_STATUS `json:"ingressProfiles,omitempty"`

	// MasterProfile: The cluster master profile.
	MasterProfile *MasterProfile_STATUS `json:"masterProfile,omitempty"`

	// NetworkProfile: The cluster network profile.
	NetworkProfile *NetworkProfile_STATUS `json:"networkProfile,omitempty"`

	// ProvisioningState: The cluster provisioning state.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ServicePrincipalProfile: The cluster service principal profile.
	ServicePrincipalProfile *ServicePrincipalProfile_STATUS `json:"servicePrincipalProfile,omitempty"`

	// WorkerProfiles: The cluster worker profiles.
	WorkerProfiles []WorkerProfile_STATUS `json:"workerProfiles,omitempty"`

	// WorkerProfilesStatus: The cluster worker profiles status.
	WorkerProfilesStatus []WorkerProfile_STATUS `json:"workerProfilesStatus,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
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

// APIServerProfile represents an API server profile.
type APIServerProfile_STATUS struct {
	// Ip: The IP of the cluster API server.
	Ip *string `json:"ip,omitempty"`

	// Url: The URL to access the cluster API server.
	Url *string `json:"url,omitempty"`

	// Visibility: API server visibility.
	Visibility *Visibility_STATUS `json:"visibility,omitempty"`
}

// ClusterProfile represents a cluster profile.
type ClusterProfile_STATUS struct {
	// Domain: The domain for the cluster.
	Domain *string `json:"domain,omitempty"`

	// FipsValidatedModules: If FIPS validated crypto modules are used
	FipsValidatedModules *FipsValidatedModules_STATUS `json:"fipsValidatedModules,omitempty"`

	// ResourceGroupId: The ID of the cluster resource group.
	ResourceGroupId *string `json:"resourceGroupId,omitempty"`

	// Version: The version of the cluster.
	Version *string `json:"version,omitempty"`
}

// ConsoleProfile represents a console profile.
type ConsoleProfile_STATUS struct {
	// Url: The URL to access the cluster console.
	Url *string `json:"url,omitempty"`
}

// IngressProfile represents an ingress profile.
type IngressProfile_STATUS struct {
	// Ip: The IP of the ingress.
	Ip *string `json:"ip,omitempty"`

	// Name: The ingress profile name.
	Name *string `json:"name,omitempty"`

	// Visibility: Ingress visibility.
	Visibility *Visibility_STATUS `json:"visibility,omitempty"`
}

// MasterProfile represents a master profile.
type MasterProfile_STATUS struct {
	// DiskEncryptionSetId: The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetId *string `json:"diskEncryptionSetId,omitempty"`

	// EncryptionAtHost: Whether master virtual machines are encrypted at host.
	EncryptionAtHost *EncryptionAtHost_STATUS `json:"encryptionAtHost,omitempty"`

	// SubnetId: The Azure resource ID of the master subnet.
	SubnetId *string `json:"subnetId,omitempty"`

	// VmSize: The size of the master VMs.
	VmSize *string `json:"vmSize,omitempty"`
}

// NetworkProfile represents a network profile.
type NetworkProfile_STATUS struct {
	// LoadBalancerProfile: The cluster load balancer profile.
	LoadBalancerProfile *LoadBalancerProfile_STATUS `json:"loadBalancerProfile,omitempty"`

	// OutboundType: The OutboundType used for egress traffic.
	OutboundType *OutboundType_STATUS `json:"outboundType,omitempty"`

	// PodCidr: The CIDR used for OpenShift/Kubernetes Pods.
	PodCidr *string `json:"podCidr,omitempty"`

	// PreconfiguredNSG: Specifies whether subnets are pre-attached with an NSG
	PreconfiguredNSG *PreconfiguredNSG_STATUS `json:"preconfiguredNSG,omitempty"`

	// ServiceCidr: The CIDR used for OpenShift/Kubernetes Services.
	ServiceCidr *string `json:"serviceCidr,omitempty"`
}

// ProvisioningState represents a provisioning state.
type ProvisioningState_STATUS string

const (
	ProvisioningState_STATUS_AdminUpdating = ProvisioningState_STATUS("AdminUpdating")
	ProvisioningState_STATUS_Canceled      = ProvisioningState_STATUS("Canceled")
	ProvisioningState_STATUS_Creating      = ProvisioningState_STATUS("Creating")
	ProvisioningState_STATUS_Deleting      = ProvisioningState_STATUS("Deleting")
	ProvisioningState_STATUS_Failed        = ProvisioningState_STATUS("Failed")
	ProvisioningState_STATUS_Succeeded     = ProvisioningState_STATUS("Succeeded")
	ProvisioningState_STATUS_Updating      = ProvisioningState_STATUS("Updating")
)

// Mapping from string to ProvisioningState_STATUS
var provisioningState_STATUS_Values = map[string]ProvisioningState_STATUS{
	"adminupdating": ProvisioningState_STATUS_AdminUpdating,
	"canceled":      ProvisioningState_STATUS_Canceled,
	"creating":      ProvisioningState_STATUS_Creating,
	"deleting":      ProvisioningState_STATUS_Deleting,
	"failed":        ProvisioningState_STATUS_Failed,
	"succeeded":     ProvisioningState_STATUS_Succeeded,
	"updating":      ProvisioningState_STATUS_Updating,
}

// ServicePrincipalProfile represents a service principal profile.
type ServicePrincipalProfile_STATUS struct {
	// ClientId: The client ID used for the cluster.
	ClientId *string `json:"clientId,omitempty"`
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

// WorkerProfile represents a worker profile.
type WorkerProfile_STATUS struct {
	// Count: The number of worker VMs.
	Count *int `json:"count,omitempty"`

	// DiskEncryptionSetId: The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetId *string `json:"diskEncryptionSetId,omitempty"`

	// DiskSizeGB: The disk size of the worker VMs.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// EncryptionAtHost: Whether master virtual machines are encrypted at host.
	EncryptionAtHost *EncryptionAtHost_STATUS `json:"encryptionAtHost,omitempty"`

	// Name: The worker profile name.
	Name *string `json:"name,omitempty"`

	// SubnetId: The Azure resource ID of the worker subnet.
	SubnetId *string `json:"subnetId,omitempty"`

	// VmSize: The size of the worker VMs.
	VmSize *string `json:"vmSize,omitempty"`
}

// EncryptionAtHost represents encryption at host state
type EncryptionAtHost_STATUS string

const (
	EncryptionAtHost_STATUS_Disabled = EncryptionAtHost_STATUS("Disabled")
	EncryptionAtHost_STATUS_Enabled  = EncryptionAtHost_STATUS("Enabled")
)

// Mapping from string to EncryptionAtHost_STATUS
var encryptionAtHost_STATUS_Values = map[string]EncryptionAtHost_STATUS{
	"disabled": EncryptionAtHost_STATUS_Disabled,
	"enabled":  EncryptionAtHost_STATUS_Enabled,
}

// FipsValidatedModules determines if FIPS is used.
type FipsValidatedModules_STATUS string

const (
	FipsValidatedModules_STATUS_Disabled = FipsValidatedModules_STATUS("Disabled")
	FipsValidatedModules_STATUS_Enabled  = FipsValidatedModules_STATUS("Enabled")
)

// Mapping from string to FipsValidatedModules_STATUS
var fipsValidatedModules_STATUS_Values = map[string]FipsValidatedModules_STATUS{
	"disabled": FipsValidatedModules_STATUS_Disabled,
	"enabled":  FipsValidatedModules_STATUS_Enabled,
}

// LoadBalancerProfile represents the profile of the cluster public load balancer.
type LoadBalancerProfile_STATUS struct {
	// EffectiveOutboundIps: The list of effective outbound IP addresses of the public load balancer.
	EffectiveOutboundIps []EffectiveOutboundIP_STATUS `json:"effectiveOutboundIps,omitempty"`

	// ManagedOutboundIps: The desired managed outbound IPs for the cluster public load balancer.
	ManagedOutboundIps *ManagedOutboundIPs_STATUS `json:"managedOutboundIps,omitempty"`
}

// The outbound routing strategy used to provide your cluster egress to the internet.
type OutboundType_STATUS string

const (
	OutboundType_STATUS_Loadbalancer       = OutboundType_STATUS("Loadbalancer")
	OutboundType_STATUS_UserDefinedRouting = OutboundType_STATUS("UserDefinedRouting")
)

// Mapping from string to OutboundType_STATUS
var outboundType_STATUS_Values = map[string]OutboundType_STATUS{
	"loadbalancer":       OutboundType_STATUS_Loadbalancer,
	"userdefinedrouting": OutboundType_STATUS_UserDefinedRouting,
}

// PreconfiguredNSG represents whether customers want to use their own NSG attached to the subnets
type PreconfiguredNSG_STATUS string

const (
	PreconfiguredNSG_STATUS_Disabled = PreconfiguredNSG_STATUS("Disabled")
	PreconfiguredNSG_STATUS_Enabled  = PreconfiguredNSG_STATUS("Enabled")
)

// Mapping from string to PreconfiguredNSG_STATUS
var preconfiguredNSG_STATUS_Values = map[string]PreconfiguredNSG_STATUS{
	"disabled": PreconfiguredNSG_STATUS_Disabled,
	"enabled":  PreconfiguredNSG_STATUS_Enabled,
}

// Visibility represents visibility.
type Visibility_STATUS string

const (
	Visibility_STATUS_Private = Visibility_STATUS("Private")
	Visibility_STATUS_Public  = Visibility_STATUS("Public")
)

// Mapping from string to Visibility_STATUS
var visibility_STATUS_Values = map[string]Visibility_STATUS{
	"private": Visibility_STATUS_Private,
	"public":  Visibility_STATUS_Public,
}

// EffectiveOutboundIP represents an effective outbound IP resource of the cluster public load balancer.
type EffectiveOutboundIP_STATUS struct {
	// Id: The fully qualified Azure resource id of an IP address resource.
	Id *string `json:"id,omitempty"`
}

// ManagedOutboundIPs represents the desired managed outbound IPs for the cluster public load balancer.
type ManagedOutboundIPs_STATUS struct {
	// Count: Count represents the desired number of IPv4 outbound IPs created and managed by Azure for the cluster public load
	// balancer.  Allowed values are in the range of 1 - 20.  The default value is 1.
	Count *int `json:"count,omitempty"`
}
