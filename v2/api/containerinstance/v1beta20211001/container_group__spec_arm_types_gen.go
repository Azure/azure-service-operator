// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type ContainerGroup_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Identity: The identity of the container group, if configured.
	Identity *ContainerGroupIdentityARM `json:"identity,omitempty"`

	// Location: The resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The container group properties
	Properties *ContainerGroup_Spec_PropertiesARM `json:"properties,omitempty"`

	// Tags: The resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: The zones for the container group.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ContainerGroup_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (group ContainerGroup_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (group *ContainerGroup_SpecARM) GetName() string {
	return group.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerInstance/containerGroups"
func (group *ContainerGroup_SpecARM) GetType() string {
	return "Microsoft.ContainerInstance/containerGroups"
}

type ContainerGroup_Spec_PropertiesARM struct {
	// Containers: The containers within the container group.
	Containers []ContainerARM `json:"containers,omitempty"`

	// Diagnostics: The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnosticsARM `json:"diagnostics,omitempty"`

	// DnsConfig: The DNS config information for a container group.
	DnsConfig *DnsConfigurationARM `json:"dnsConfig,omitempty"`

	// EncryptionProperties: The encryption properties for a container group.
	EncryptionProperties *EncryptionPropertiesARM `json:"encryptionProperties,omitempty"`

	// ImageRegistryCredentials: The image registry credentials by which the container group is created from.
	ImageRegistryCredentials []ImageRegistryCredentialARM `json:"imageRegistryCredentials,omitempty"`

	// InitContainers: The init containers for a container group.
	InitContainers []InitContainerDefinitionARM `json:"initContainers,omitempty"`

	// IpAddress: The IP address type of the container group.
	IpAddress *IpAddressARM `json:"ipAddress,omitempty"`

	// OsType: The operating system type required by the containers in the container group.
	OsType *ContainerGroup_Spec_Properties_OsType `json:"osType,omitempty"`

	// RestartPolicy: Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy *ContainerGroup_Spec_Properties_RestartPolicy `json:"restartPolicy,omitempty"`

	// Sku: The SKU for a container group.
	Sku *ContainerGroupSku `json:"sku,omitempty"`

	// SubnetIds: The subnet resource IDs for a container group.
	SubnetIds []ContainerGroupSubnetIdARM `json:"subnetIds,omitempty"`

	// Volumes: The list of volumes that can be mounted by containers in this container group.
	Volumes []VolumeARM `json:"volumes,omitempty"`
}

type ContainerGroupIdentityARM struct {
	// Type: The type of identity used for the container group. The type 'SystemAssigned, UserAssigned' includes both an
	// implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the
	// container group.
	Type *ContainerGroupIdentity_Type `json:"type,omitempty"`
}

type ContainerARM struct {
	// Name: The user-provided name of the container instance.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the container instance.
	Properties *ContainerPropertiesARM `json:"properties,omitempty"`
}

type ContainerGroupDiagnosticsARM struct {
	// LogAnalytics: Container group log analytics information.
	LogAnalytics *LogAnalyticsARM `json:"logAnalytics,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type ContainerGroupIdentity_Type string

const (
	ContainerGroupIdentity_Type_None                       = ContainerGroupIdentity_Type("None")
	ContainerGroupIdentity_Type_SystemAssigned             = ContainerGroupIdentity_Type("SystemAssigned")
	ContainerGroupIdentity_Type_SystemAssignedUserAssigned = ContainerGroupIdentity_Type("SystemAssigned, UserAssigned")
	ContainerGroupIdentity_Type_UserAssigned               = ContainerGroupIdentity_Type("UserAssigned")
)

type ContainerGroupSubnetIdARM struct {
	Id *string `json:"id,omitempty"`

	// Name: Friendly name for the subnet.
	Name *string `json:"name,omitempty"`
}

type DnsConfigurationARM struct {
	// NameServers: The DNS servers for the container group.
	NameServers []string `json:"nameServers,omitempty"`

	// Options: The DNS options for the container group.
	Options *string `json:"options,omitempty"`

	// SearchDomains: The DNS search domains for hostname lookup in the container group.
	SearchDomains *string `json:"searchDomains,omitempty"`
}

type EncryptionPropertiesARM struct {
	// KeyName: The encryption key name.
	KeyName *string `json:"keyName,omitempty"`

	// KeyVersion: The encryption key version.
	KeyVersion *string `json:"keyVersion,omitempty"`

	// VaultBaseUrl: The keyvault base url.
	VaultBaseUrl *string `json:"vaultBaseUrl,omitempty"`
}

type ImageRegistryCredentialARM struct {
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

type InitContainerDefinitionARM struct {
	// Name: The name for the init container.
	Name *string `json:"name,omitempty"`

	// Properties: The properties for the init container.
	Properties *InitContainerPropertiesDefinitionARM `json:"properties,omitempty"`
}

type IpAddressARM struct {
	// DnsNameLabel: The Dns name label for the IP.
	DnsNameLabel *string `json:"dnsNameLabel,omitempty"`

	// DnsNameLabelReusePolicy: The value representing the security enum.
	DnsNameLabelReusePolicy *IpAddress_DnsNameLabelReusePolicy `json:"dnsNameLabelReusePolicy,omitempty"`

	// Ip: The IP exposed to the public internet.
	Ip *string `json:"ip,omitempty"`

	// Ports: The list of ports exposed on the container group.
	Ports []PortARM `json:"ports,omitempty"`

	// Type: Specifies if the IP is exposed to the public internet or private VNET.
	Type *IpAddress_Type `json:"type,omitempty"`
}

type VolumeARM struct {
	// AzureFile: The Azure File volume.
	AzureFile *AzureFileVolumeARM `json:"azureFile,omitempty"`

	// EmptyDir: The empty directory volume.
	EmptyDir map[string]v1.JSON `json:"emptyDir,omitempty"`

	// GitRepo: The git repo volume.
	GitRepo *GitRepoVolumeARM `json:"gitRepo,omitempty"`

	// Name: The name of the volume.
	Name *string `json:"name,omitempty"`

	// Secret: The secret volume.
	Secret map[string]string `json:"secret,omitempty"`
}

type AzureFileVolumeARM struct {
	// ReadOnly: The flag indicating whether the Azure File shared mounted as a volume is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`

	// ShareName: The name of the Azure File share to be mounted as a volume.
	ShareName *string `json:"shareName,omitempty"`

	// StorageAccountKey: The storage account access key used to access the Azure File share.
	StorageAccountKey *string `json:"storageAccountKey,omitempty"`

	// StorageAccountName: The name of the storage account that contains the Azure File share.
	StorageAccountName *string `json:"storageAccountName,omitempty"`
}

type ContainerPropertiesARM struct {
	// Command: The commands to execute within the container instance in exec form.
	Command []string `json:"command,omitempty"`

	// EnvironmentVariables: The environment variables to set in the container instance.
	EnvironmentVariables []EnvironmentVariableARM `json:"environmentVariables,omitempty"`

	// Image: The name of the image used to create the container instance.
	Image *string `json:"image,omitempty"`

	// LivenessProbe: The liveness probe.
	LivenessProbe *ContainerProbeARM `json:"livenessProbe,omitempty"`

	// Ports: The exposed ports on the container instance.
	Ports []ContainerPortARM `json:"ports,omitempty"`

	// ReadinessProbe: The readiness probe.
	ReadinessProbe *ContainerProbeARM `json:"readinessProbe,omitempty"`

	// Resources: The resource requirements of the container instance.
	Resources *ResourceRequirementsARM `json:"resources,omitempty"`

	// VolumeMounts: The volume mounts available to the container instance.
	VolumeMounts []VolumeMountARM `json:"volumeMounts,omitempty"`
}

type GitRepoVolumeARM struct {
	// Directory: Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be
	// the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the
	// given name.
	Directory *string `json:"directory,omitempty"`

	// Repository: Repository URL
	Repository *string `json:"repository,omitempty"`

	// Revision: Commit hash for the specified revision.
	Revision *string `json:"revision,omitempty"`
}

type InitContainerPropertiesDefinitionARM struct {
	// Command: The command to execute within the init container in exec form.
	Command []string `json:"command,omitempty"`

	// EnvironmentVariables: The environment variables to set in the init container.
	EnvironmentVariables []EnvironmentVariableARM `json:"environmentVariables,omitempty"`

	// Image: The image of the init container.
	Image *string `json:"image,omitempty"`

	// VolumeMounts: The volume mounts available to the init container.
	VolumeMounts []VolumeMountARM `json:"volumeMounts,omitempty"`
}

type LogAnalyticsARM struct {
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

type PortARM struct {
	// Port: The port number.
	Port *int `json:"port,omitempty"`

	// Protocol: The protocol associated with the port.
	Protocol *Port_Protocol `json:"protocol,omitempty"`
}

type ContainerPortARM struct {
	// Port: The port number exposed within the container group.
	Port *int `json:"port,omitempty"`

	// Protocol: The protocol associated with the port.
	Protocol *ContainerPort_Protocol `json:"protocol,omitempty"`
}

type ContainerProbeARM struct {
	// Exec: The execution command to probe
	Exec *ContainerExecARM `json:"exec,omitempty"`

	// FailureThreshold: The failure threshold.
	FailureThreshold *int `json:"failureThreshold,omitempty"`

	// HttpGet: The Http Get settings to probe
	HttpGet *ContainerHttpGetARM `json:"httpGet,omitempty"`

	// InitialDelaySeconds: The initial delay seconds.
	InitialDelaySeconds *int `json:"initialDelaySeconds,omitempty"`

	// PeriodSeconds: The period seconds.
	PeriodSeconds *int `json:"periodSeconds,omitempty"`

	// SuccessThreshold: The success threshold.
	SuccessThreshold *int `json:"successThreshold,omitempty"`

	// TimeoutSeconds: The timeout seconds.
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
}

type EnvironmentVariableARM struct {
	// Name: The name of the environment variable.
	Name *string `json:"name,omitempty"`

	// SecureValue: The value of the secure environment variable.
	SecureValue *string `json:"secureValue,omitempty"`

	// Value: The value of the environment variable.
	Value *string `json:"value,omitempty"`
}

type ResourceRequirementsARM struct {
	// Limits: The resource limits of this container instance.
	Limits *ResourceLimitsARM `json:"limits,omitempty"`

	// Requests: The resource requests of this container instance.
	Requests *ResourceRequestsARM `json:"requests,omitempty"`
}

type VolumeMountARM struct {
	// MountPath: The path within the container where the volume should be mounted. Must not contain colon (:).
	MountPath *string `json:"mountPath,omitempty"`

	// Name: The name of the volume mount.
	Name *string `json:"name,omitempty"`

	// ReadOnly: The flag indicating whether the volume mount is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

type ContainerExecARM struct {
	// Command: The commands to execute within the container.
	Command []string `json:"command,omitempty"`
}

type ContainerHttpGetARM struct {
	// HttpHeaders: The HTTP headers.
	HttpHeaders []HttpHeaderARM `json:"httpHeaders,omitempty"`

	// Path: The path to probe.
	Path *string `json:"path,omitempty"`

	// Port: The port number to probe.
	Port *int `json:"port,omitempty"`

	// Scheme: The scheme.
	Scheme *ContainerHttpGet_Scheme `json:"scheme,omitempty"`
}

type ResourceLimitsARM struct {
	// Cpu: The CPU limit of this container instance.
	Cpu *float64 `json:"cpu,omitempty"`

	// Gpu: The GPU limit of this container instance.
	Gpu *GpuResourceARM `json:"gpu,omitempty"`

	// MemoryInGB: The memory limit in GB of this container instance.
	MemoryInGB *float64 `json:"memoryInGB,omitempty"`
}

type ResourceRequestsARM struct {
	// Cpu: The CPU request of this container instance.
	Cpu *float64 `json:"cpu,omitempty"`

	// Gpu: The GPU request of this container instance.
	Gpu *GpuResourceARM `json:"gpu,omitempty"`

	// MemoryInGB: The memory request in GB of this container instance.
	MemoryInGB *float64 `json:"memoryInGB,omitempty"`
}

type GpuResourceARM struct {
	// Count: The count of the GPU resource.
	Count *int `json:"count,omitempty"`

	// Sku: The SKU of the GPU resource.
	Sku *GpuResource_Sku `json:"sku,omitempty"`
}

type HttpHeaderARM struct {
	// Name: The header name.
	Name *string `json:"name,omitempty"`

	// Value: The header value.
	Value *string `json:"value,omitempty"`
}
