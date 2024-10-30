// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type ContainerGroup_Spec struct {
	// Identity: The identity of the container group, if configured.
	Identity *ContainerGroupIdentity `json:"identity,omitempty"`

	// Location: The resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The container group properties
	Properties *ContainerGroup_Properties_Spec `json:"properties,omitempty"`

	// Tags: The resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: The zones for the container group.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ContainerGroup_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (group ContainerGroup_Spec) GetAPIVersion() string {
	return "2021-10-01"
}

// GetName returns the Name of the resource
func (group *ContainerGroup_Spec) GetName() string {
	return group.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerInstance/containerGroups"
func (group *ContainerGroup_Spec) GetType() string {
	return "Microsoft.ContainerInstance/containerGroups"
}

type ContainerGroup_Properties_Spec struct {
	// Containers: The containers within the container group.
	Containers []Container `json:"containers,omitempty"`

	// Diagnostics: The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnostics `json:"diagnostics,omitempty"`

	// DnsConfig: The DNS config information for a container group.
	DnsConfig *DnsConfiguration `json:"dnsConfig,omitempty"`

	// EncryptionProperties: The encryption properties for a container group.
	EncryptionProperties *EncryptionProperties `json:"encryptionProperties,omitempty"`

	// ImageRegistryCredentials: The image registry credentials by which the container group is created from.
	ImageRegistryCredentials []ImageRegistryCredential `json:"imageRegistryCredentials,omitempty"`

	// InitContainers: The init containers for a container group.
	InitContainers []InitContainerDefinition `json:"initContainers,omitempty"`

	// IpAddress: The IP address type of the container group.
	IpAddress *IpAddress `json:"ipAddress,omitempty"`

	// OsType: The operating system type required by the containers in the container group.
	OsType *ContainerGroup_Properties_OsType_Spec `json:"osType,omitempty"`

	// RestartPolicy: Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy *ContainerGroup_Properties_RestartPolicy_Spec `json:"restartPolicy,omitempty"`

	// Sku: The SKU for a container group.
	Sku *ContainerGroupSku `json:"sku,omitempty"`

	// SubnetIds: The subnet resource IDs for a container group.
	SubnetIds []ContainerGroupSubnetId `json:"subnetIds,omitempty"`

	// Volumes: The list of volumes that can be mounted by containers in this container group.
	Volumes []Volume `json:"volumes,omitempty"`
}

// Identity for the container group.
type ContainerGroupIdentity struct {
	// Type: The type of identity used for the container group. The type 'SystemAssigned, UserAssigned' includes both an
	// implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the
	// container group.
	Type                   *ContainerGroupIdentity_Type           `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// A container instance.
type Container struct {
	// Name: The user-provided name of the container instance.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the container instance.
	Properties *ContainerProperties `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Linux","Windows"}
type ContainerGroup_Properties_OsType_Spec string

const (
	ContainerGroup_Properties_OsType_Spec_Linux   = ContainerGroup_Properties_OsType_Spec("Linux")
	ContainerGroup_Properties_OsType_Spec_Windows = ContainerGroup_Properties_OsType_Spec("Windows")
)

// Mapping from string to ContainerGroup_Properties_OsType_Spec
var containerGroup_Properties_OsType_Spec_Values = map[string]ContainerGroup_Properties_OsType_Spec{
	"linux":   ContainerGroup_Properties_OsType_Spec_Linux,
	"windows": ContainerGroup_Properties_OsType_Spec_Windows,
}

// +kubebuilder:validation:Enum={"Always","Never","OnFailure"}
type ContainerGroup_Properties_RestartPolicy_Spec string

const (
	ContainerGroup_Properties_RestartPolicy_Spec_Always    = ContainerGroup_Properties_RestartPolicy_Spec("Always")
	ContainerGroup_Properties_RestartPolicy_Spec_Never     = ContainerGroup_Properties_RestartPolicy_Spec("Never")
	ContainerGroup_Properties_RestartPolicy_Spec_OnFailure = ContainerGroup_Properties_RestartPolicy_Spec("OnFailure")
)

// Mapping from string to ContainerGroup_Properties_RestartPolicy_Spec
var containerGroup_Properties_RestartPolicy_Spec_Values = map[string]ContainerGroup_Properties_RestartPolicy_Spec{
	"always":    ContainerGroup_Properties_RestartPolicy_Spec_Always,
	"never":     ContainerGroup_Properties_RestartPolicy_Spec_Never,
	"onfailure": ContainerGroup_Properties_RestartPolicy_Spec_OnFailure,
}

// Container group diagnostic information.
type ContainerGroupDiagnostics struct {
	// LogAnalytics: Container group log analytics information.
	LogAnalytics *LogAnalytics `json:"logAnalytics,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type ContainerGroupIdentity_Type string

const (
	ContainerGroupIdentity_Type_None                       = ContainerGroupIdentity_Type("None")
	ContainerGroupIdentity_Type_SystemAssigned             = ContainerGroupIdentity_Type("SystemAssigned")
	ContainerGroupIdentity_Type_SystemAssignedUserAssigned = ContainerGroupIdentity_Type("SystemAssigned, UserAssigned")
	ContainerGroupIdentity_Type_UserAssigned               = ContainerGroupIdentity_Type("UserAssigned")
)

// Mapping from string to ContainerGroupIdentity_Type
var containerGroupIdentity_Type_Values = map[string]ContainerGroupIdentity_Type{
	"none":                         ContainerGroupIdentity_Type_None,
	"systemassigned":               ContainerGroupIdentity_Type_SystemAssigned,
	"systemassigned, userassigned": ContainerGroupIdentity_Type_SystemAssignedUserAssigned,
	"userassigned":                 ContainerGroupIdentity_Type_UserAssigned,
}

// The container group SKU.
// +kubebuilder:validation:Enum={"Dedicated","Standard"}
type ContainerGroupSku string

const (
	ContainerGroupSku_Dedicated = ContainerGroupSku("Dedicated")
	ContainerGroupSku_Standard  = ContainerGroupSku("Standard")
)

// Mapping from string to ContainerGroupSku
var containerGroupSku_Values = map[string]ContainerGroupSku{
	"dedicated": ContainerGroupSku_Dedicated,
	"standard":  ContainerGroupSku_Standard,
}

// Container group subnet information.
type ContainerGroupSubnetId struct {
	Id *string `json:"id,omitempty"`

	// Name: Friendly name for the subnet.
	Name *string `json:"name,omitempty"`
}

// DNS configuration for the container group.
type DnsConfiguration struct {
	// NameServers: The DNS servers for the container group.
	NameServers []string `json:"nameServers,omitempty"`

	// Options: The DNS options for the container group.
	Options *string `json:"options,omitempty"`

	// SearchDomains: The DNS search domains for hostname lookup in the container group.
	SearchDomains *string `json:"searchDomains,omitempty"`
}

// The container group encryption properties.
type EncryptionProperties struct {
	// KeyName: The encryption key name.
	KeyName *string `json:"keyName,omitempty"`

	// KeyVersion: The encryption key version.
	KeyVersion *string `json:"keyVersion,omitempty"`

	// VaultBaseUrl: The keyvault base url.
	VaultBaseUrl *string `json:"vaultBaseUrl,omitempty"`
}

// Image registry credential.
type ImageRegistryCredential struct {
	// Identity: The identity for the private registry.
	Identity *string `json:"identity,omitempty"`

	// IdentityUrl: The identity URL for the private registry.
	IdentityUrl *string `json:"identityUrl,omitempty"`

	// Password: The password for the private registry.
	Password *string `json:"password,omitempty"`

	// Server: The Docker image registry server without a protocol such as "http" and "https".
	Server *string `json:"server,omitempty"`

	// Username: The username for the private registry.
	Username *string `json:"username,omitempty"`
}

// The init container definition.
type InitContainerDefinition struct {
	// Name: The name for the init container.
	Name *string `json:"name,omitempty"`

	// Properties: The properties for the init container.
	Properties *InitContainerPropertiesDefinition `json:"properties,omitempty"`
}

// IP address for the container group.
type IpAddress struct {
	// AutoGeneratedDomainNameLabelScope: The value representing the security enum. The 'Unsecure' value is the default value
	// if not selected and means the object's domain name label is not secured against subdomain takeover. The 'TenantReuse'
	// value is the default value if selected and means the object's domain name label can be reused within the same tenant.
	// The 'SubscriptionReuse' value means the object's domain name label can be reused within the same subscription. The
	// 'ResourceGroupReuse' value means the object's domain name label can be reused within the same resource group. The
	// 'NoReuse' value means the object's domain name label cannot be reused within the same resource group, subscription, or
	// tenant.
	AutoGeneratedDomainNameLabelScope *IpAddress_AutoGeneratedDomainNameLabelScope `json:"autoGeneratedDomainNameLabelScope,omitempty"`

	// DnsNameLabel: The Dns name label for the IP.
	DnsNameLabel *string `json:"dnsNameLabel,omitempty"`

	// Ip: The IP exposed to the public internet.
	Ip *string `json:"ip,omitempty"`

	// Ports: The list of ports exposed on the container group.
	Ports []Port `json:"ports,omitempty"`

	// Type: Specifies if the IP is exposed to the public internet or private VNET.
	Type *IpAddress_Type `json:"type,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
}

// The properties of the volume.
type Volume struct {
	// AzureFile: The Azure File volume.
	AzureFile *AzureFileVolume `json:"azureFile,omitempty"`

	// EmptyDir: The empty directory volume.
	EmptyDir map[string]v1.JSON `json:"emptyDir,omitempty"`

	// GitRepo: The git repo volume.
	GitRepo *GitRepoVolume `json:"gitRepo,omitempty"`

	// Name: The name of the volume.
	Name *string `json:"name,omitempty"`

	// Secret: The secret volume.
	Secret map[string]string `json:"secret,omitempty"`
}

// The properties of the Azure File volume. Azure File shares are mounted as volumes.
type AzureFileVolume struct {
	// ReadOnly: The flag indicating whether the Azure File shared mounted as a volume is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`

	// ShareName: The name of the Azure File share to be mounted as a volume.
	ShareName *string `json:"shareName,omitempty"`

	// StorageAccountKey: The storage account access key used to access the Azure File share.
	StorageAccountKey *string `json:"storageAccountKey,omitempty"`

	// StorageAccountName: The name of the storage account that contains the Azure File share.
	StorageAccountName *string `json:"storageAccountName,omitempty"`
}

// The container instance properties.
type ContainerProperties struct {
	// Command: The commands to execute within the container instance in exec form.
	Command []string `json:"command,omitempty"`

	// EnvironmentVariables: The environment variables to set in the container instance.
	EnvironmentVariables []EnvironmentVariable `json:"environmentVariables,omitempty"`

	// Image: The name of the image used to create the container instance.
	Image *string `json:"image,omitempty"`

	// LivenessProbe: The liveness probe.
	LivenessProbe *ContainerProbe `json:"livenessProbe,omitempty"`

	// Ports: The exposed ports on the container instance.
	Ports []ContainerPort `json:"ports,omitempty"`

	// ReadinessProbe: The readiness probe.
	ReadinessProbe *ContainerProbe `json:"readinessProbe,omitempty"`

	// Resources: The resource requirements of the container instance.
	Resources *ResourceRequirements `json:"resources,omitempty"`

	// VolumeMounts: The volume mounts available to the container instance.
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`
}

// Represents a volume that is populated with the contents of a git repository
type GitRepoVolume struct {
	// Directory: Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be
	// the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the
	// given name.
	Directory *string `json:"directory,omitempty"`

	// Repository: Repository URL
	Repository *string `json:"repository,omitempty"`

	// Revision: Commit hash for the specified revision.
	Revision *string `json:"revision,omitempty"`
}

// The init container definition properties.
type InitContainerPropertiesDefinition struct {
	// Command: The command to execute within the init container in exec form.
	Command []string `json:"command,omitempty"`

	// EnvironmentVariables: The environment variables to set in the init container.
	EnvironmentVariables []EnvironmentVariable `json:"environmentVariables,omitempty"`

	// Image: The image of the init container.
	Image *string `json:"image,omitempty"`

	// VolumeMounts: The volume mounts available to the init container.
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`
}

// +kubebuilder:validation:Enum={"Noreuse","ResourceGroupReuse","SubscriptionReuse","TenantReuse","Unsecure"}
type IpAddress_AutoGeneratedDomainNameLabelScope string

const (
	IpAddress_AutoGeneratedDomainNameLabelScope_Noreuse            = IpAddress_AutoGeneratedDomainNameLabelScope("Noreuse")
	IpAddress_AutoGeneratedDomainNameLabelScope_ResourceGroupReuse = IpAddress_AutoGeneratedDomainNameLabelScope("ResourceGroupReuse")
	IpAddress_AutoGeneratedDomainNameLabelScope_SubscriptionReuse  = IpAddress_AutoGeneratedDomainNameLabelScope("SubscriptionReuse")
	IpAddress_AutoGeneratedDomainNameLabelScope_TenantReuse        = IpAddress_AutoGeneratedDomainNameLabelScope("TenantReuse")
	IpAddress_AutoGeneratedDomainNameLabelScope_Unsecure           = IpAddress_AutoGeneratedDomainNameLabelScope("Unsecure")
)

// Mapping from string to IpAddress_AutoGeneratedDomainNameLabelScope
var ipAddress_AutoGeneratedDomainNameLabelScope_Values = map[string]IpAddress_AutoGeneratedDomainNameLabelScope{
	"noreuse":            IpAddress_AutoGeneratedDomainNameLabelScope_Noreuse,
	"resourcegroupreuse": IpAddress_AutoGeneratedDomainNameLabelScope_ResourceGroupReuse,
	"subscriptionreuse":  IpAddress_AutoGeneratedDomainNameLabelScope_SubscriptionReuse,
	"tenantreuse":        IpAddress_AutoGeneratedDomainNameLabelScope_TenantReuse,
	"unsecure":           IpAddress_AutoGeneratedDomainNameLabelScope_Unsecure,
}

// +kubebuilder:validation:Enum={"Private","Public"}
type IpAddress_Type string

const (
	IpAddress_Type_Private = IpAddress_Type("Private")
	IpAddress_Type_Public  = IpAddress_Type("Public")
)

// Mapping from string to IpAddress_Type
var ipAddress_Type_Values = map[string]IpAddress_Type{
	"private": IpAddress_Type_Private,
	"public":  IpAddress_Type_Public,
}

// Container group log analytics information.
type LogAnalytics struct {
	// LogType: The log type to be used.
	LogType *LogAnalytics_LogType `json:"logType,omitempty"`

	// Metadata: Metadata for log analytics.
	Metadata map[string]string `json:"metadata,omitempty"`

	// WorkspaceId: The workspace id for log analytics
	WorkspaceId *string `json:"workspaceId,omitempty"`

	// WorkspaceKey: The workspace key for log analytics
	WorkspaceKey        *string `json:"workspaceKey,omitempty"`
	WorkspaceResourceId *string `json:"workspaceResourceId,omitempty"`
}

// The port exposed on the container group.
type Port struct {
	// Port: The port number.
	Port *int `json:"port,omitempty"`

	// Protocol: The protocol associated with the port.
	Protocol *Port_Protocol `json:"protocol,omitempty"`
}

// The port exposed on the container instance.
type ContainerPort struct {
	// Port: The port number exposed within the container group.
	Port *int `json:"port,omitempty"`

	// Protocol: The protocol associated with the port.
	Protocol *ContainerPort_Protocol `json:"protocol,omitempty"`
}

// The container probe, for liveness or readiness
type ContainerProbe struct {
	// Exec: The execution command to probe
	Exec *ContainerExec `json:"exec,omitempty"`

	// FailureThreshold: The failure threshold.
	FailureThreshold *int `json:"failureThreshold,omitempty"`

	// HttpGet: The Http Get settings to probe
	HttpGet *ContainerHttpGet `json:"httpGet,omitempty"`

	// InitialDelaySeconds: The initial delay seconds.
	InitialDelaySeconds *int `json:"initialDelaySeconds,omitempty"`

	// PeriodSeconds: The period seconds.
	PeriodSeconds *int `json:"periodSeconds,omitempty"`

	// SuccessThreshold: The success threshold.
	SuccessThreshold *int `json:"successThreshold,omitempty"`

	// TimeoutSeconds: The timeout seconds.
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
}

// The environment variable to set within the container instance.
type EnvironmentVariable struct {
	// Name: The name of the environment variable.
	Name *string `json:"name,omitempty"`

	// SecureValue: The value of the secure environment variable.
	SecureValue *string `json:"secureValue,omitempty"`

	// Value: The value of the environment variable.
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:validation:Enum={"ContainerInsights","ContainerInstanceLogs"}
type LogAnalytics_LogType string

const (
	LogAnalytics_LogType_ContainerInsights     = LogAnalytics_LogType("ContainerInsights")
	LogAnalytics_LogType_ContainerInstanceLogs = LogAnalytics_LogType("ContainerInstanceLogs")
)

// Mapping from string to LogAnalytics_LogType
var logAnalytics_LogType_Values = map[string]LogAnalytics_LogType{
	"containerinsights":     LogAnalytics_LogType_ContainerInsights,
	"containerinstancelogs": LogAnalytics_LogType_ContainerInstanceLogs,
}

// +kubebuilder:validation:Enum={"TCP","UDP"}
type Port_Protocol string

const (
	Port_Protocol_TCP = Port_Protocol("TCP")
	Port_Protocol_UDP = Port_Protocol("UDP")
)

// Mapping from string to Port_Protocol
var port_Protocol_Values = map[string]Port_Protocol{
	"tcp": Port_Protocol_TCP,
	"udp": Port_Protocol_UDP,
}

// The resource requirements.
type ResourceRequirements struct {
	// Limits: The resource limits of this container instance.
	Limits *ResourceLimits `json:"limits,omitempty"`

	// Requests: The resource requests of this container instance.
	Requests *ResourceRequests `json:"requests,omitempty"`
}

// The properties of the volume mount.
type VolumeMount struct {
	// MountPath: The path within the container where the volume should be mounted. Must not contain colon (:).
	MountPath *string `json:"mountPath,omitempty"`

	// Name: The name of the volume mount.
	Name *string `json:"name,omitempty"`

	// ReadOnly: The flag indicating whether the volume mount is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// The container execution command, for liveness or readiness probe
type ContainerExec struct {
	// Command: The commands to execute within the container.
	Command []string `json:"command,omitempty"`
}

// The container Http Get settings, for liveness or readiness probe
type ContainerHttpGet struct {
	// HttpHeaders: The HTTP headers.
	HttpHeaders []HttpHeader `json:"httpHeaders,omitempty"`

	// Path: The path to probe.
	Path *string `json:"path,omitempty"`

	// Port: The port number to probe.
	Port *int `json:"port,omitempty"`

	// Scheme: The scheme.
	Scheme *ContainerHttpGet_Scheme `json:"scheme,omitempty"`
}

// +kubebuilder:validation:Enum={"TCP","UDP"}
type ContainerPort_Protocol string

const (
	ContainerPort_Protocol_TCP = ContainerPort_Protocol("TCP")
	ContainerPort_Protocol_UDP = ContainerPort_Protocol("UDP")
)

// Mapping from string to ContainerPort_Protocol
var containerPort_Protocol_Values = map[string]ContainerPort_Protocol{
	"tcp": ContainerPort_Protocol_TCP,
	"udp": ContainerPort_Protocol_UDP,
}

// The resource limits.
type ResourceLimits struct {
	// Cpu: The CPU limit of this container instance.
	Cpu *float64 `json:"cpu,omitempty"`

	// Gpu: The GPU limit of this container instance.
	Gpu *GpuResource `json:"gpu,omitempty"`

	// MemoryInGB: The memory limit in GB of this container instance.
	MemoryInGB *float64 `json:"memoryInGB,omitempty"`
}

// The resource requests.
type ResourceRequests struct {
	// Cpu: The CPU request of this container instance.
	Cpu *float64 `json:"cpu,omitempty"`

	// Gpu: The GPU request of this container instance.
	Gpu *GpuResource `json:"gpu,omitempty"`

	// MemoryInGB: The memory request in GB of this container instance.
	MemoryInGB *float64 `json:"memoryInGB,omitempty"`
}

// +kubebuilder:validation:Enum={"http","https"}
type ContainerHttpGet_Scheme string

const (
	ContainerHttpGet_Scheme_Http  = ContainerHttpGet_Scheme("http")
	ContainerHttpGet_Scheme_Https = ContainerHttpGet_Scheme("https")
)

// Mapping from string to ContainerHttpGet_Scheme
var containerHttpGet_Scheme_Values = map[string]ContainerHttpGet_Scheme{
	"http":  ContainerHttpGet_Scheme_Http,
	"https": ContainerHttpGet_Scheme_Https,
}

// The GPU resource.
type GpuResource struct {
	// Count: The count of the GPU resource.
	Count *int `json:"count,omitempty"`

	// Sku: The SKU of the GPU resource.
	Sku *GpuResource_Sku `json:"sku,omitempty"`
}

// The HTTP header.
type HttpHeader struct {
	// Name: The header name.
	Name *string `json:"name,omitempty"`

	// Value: The header value.
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:validation:Enum={"K80","P100","V100"}
type GpuResource_Sku string

const (
	GpuResource_Sku_K80  = GpuResource_Sku("K80")
	GpuResource_Sku_P100 = GpuResource_Sku("P100")
	GpuResource_Sku_V100 = GpuResource_Sku("V100")
)

// Mapping from string to GpuResource_Sku
var gpuResource_Sku_Values = map[string]GpuResource_Sku{
	"k80":  GpuResource_Sku_K80,
	"p100": GpuResource_Sku_P100,
	"v100": GpuResource_Sku_V100,
}