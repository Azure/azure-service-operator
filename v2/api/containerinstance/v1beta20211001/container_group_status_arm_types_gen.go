// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

type ContainerGroup_STATUSARM struct {
	// Id: The resource id.
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the container group, if configured.
	Identity *ContainerGroupIdentity_STATUSARM `json:"identity,omitempty"`

	// Location: The resource location.
	Location *string `json:"location,omitempty"`

	// Name: The resource name.
	Name *string `json:"name,omitempty"`

	// Properties: The container group properties
	Properties *ContainerGroup_Properties_STATUSARM `json:"properties,omitempty"`

	// Tags: The resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The resource type.
	Type *string `json:"type,omitempty"`

	// Zones: The zones for the container group.
	Zones []string `json:"zones,omitempty"`
}

type ContainerGroup_Properties_STATUSARM struct {
	// Containers: The containers within the container group.
	Containers []Container_STATUSARM `json:"containers,omitempty"`

	// Diagnostics: The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnostics_STATUSARM `json:"diagnostics,omitempty"`

	// DnsConfig: The DNS config information for a container group.
	DnsConfig *DnsConfiguration_STATUSARM `json:"dnsConfig,omitempty"`

	// EncryptionProperties: The encryption properties for a container group.
	EncryptionProperties *EncryptionProperties_STATUSARM `json:"encryptionProperties,omitempty"`

	// ImageRegistryCredentials: The image registry credentials by which the container group is created from.
	ImageRegistryCredentials []ImageRegistryCredential_STATUSARM `json:"imageRegistryCredentials,omitempty"`

	// InitContainers: The init containers for a container group.
	InitContainers []InitContainerDefinition_STATUSARM `json:"initContainers,omitempty"`

	// InstanceView: The instance view of the container group. Only valid in response.
	InstanceView *ContainerGroup_Properties_InstanceView_STATUSARM `json:"instanceView,omitempty"`

	// IpAddress: The IP address type of the container group.
	IpAddress *IpAddress_STATUSARM `json:"ipAddress,omitempty"`

	// OsType: The operating system type required by the containers in the container group.
	OsType *ContainerGroup_Properties_OsType_STATUS `json:"osType,omitempty"`

	// ProvisioningState: The provisioning state of the container group. This only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// RestartPolicy: Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy *ContainerGroup_Properties_RestartPolicy_STATUS `json:"restartPolicy,omitempty"`

	// Sku: The SKU for a container group.
	Sku *ContainerGroupSku_STATUS `json:"sku,omitempty"`

	// SubnetIds: The subnet resource IDs for a container group.
	SubnetIds []ContainerGroupSubnetId_STATUSARM `json:"subnetIds,omitempty"`

	// Volumes: The list of volumes that can be mounted by containers in this container group.
	Volumes []Volume_STATUSARM `json:"volumes,omitempty"`
}

type ContainerGroupIdentity_STATUSARM struct {
	// PrincipalId: The principal id of the container group identity. This property will only be provided for a system assigned
	// identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant id associated with the container group. This property will only be provided for a system assigned
	// identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of identity used for the container group. The type 'SystemAssigned, UserAssigned' includes both an
	// implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the
	// container group.
	Type *ContainerGroupIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user identities associated with the container group. The user identity dictionary
	// key references will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]ContainerGroupIdentity_UserAssignedIdentities_STATUSARM `json:"userAssignedIdentities,omitempty"`
}

type Container_STATUSARM struct {
	// Name: The user-provided name of the container instance.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the container instance.
	Properties *ContainerProperties_STATUSARM `json:"properties,omitempty"`
}

type ContainerGroup_Properties_InstanceView_STATUSARM struct {
	// Events: The events of this container group.
	Events []Event_STATUSARM `json:"events,omitempty"`

	// State: The state of the container group. Only valid in response.
	State *string `json:"state,omitempty"`
}

type ContainerGroupDiagnostics_STATUSARM struct {
	// LogAnalytics: Container group log analytics information.
	LogAnalytics *LogAnalytics_STATUSARM `json:"logAnalytics,omitempty"`
}

type ContainerGroupIdentity_Type_STATUS string

const (
	ContainerGroupIdentity_Type_None_STATUS                       = ContainerGroupIdentity_Type_STATUS("None")
	ContainerGroupIdentity_Type_SystemAssigned_STATUS             = ContainerGroupIdentity_Type_STATUS("SystemAssigned")
	ContainerGroupIdentity_Type_SystemAssignedUserAssigned_STATUS = ContainerGroupIdentity_Type_STATUS("SystemAssigned, UserAssigned")
	ContainerGroupIdentity_Type_UserAssigned_STATUS               = ContainerGroupIdentity_Type_STATUS("UserAssigned")
)

type ContainerGroupIdentity_UserAssignedIdentities_STATUSARM struct {
	// ClientId: The client id of user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal id of user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

type ContainerGroupSubnetId_STATUSARM struct {
	// Id: Resource ID of virtual network and subnet.
	Id *string `json:"id,omitempty"`

	// Name: Friendly name for the subnet.
	Name *string `json:"name,omitempty"`
}

type DnsConfiguration_STATUSARM struct {
	// NameServers: The DNS servers for the container group.
	NameServers []string `json:"nameServers,omitempty"`

	// Options: The DNS options for the container group.
	Options *string `json:"options,omitempty"`

	// SearchDomains: The DNS search domains for hostname lookup in the container group.
	SearchDomains *string `json:"searchDomains,omitempty"`
}

type EncryptionProperties_STATUSARM struct {
	// KeyName: The encryption key name.
	KeyName *string `json:"keyName,omitempty"`

	// KeyVersion: The encryption key version.
	KeyVersion *string `json:"keyVersion,omitempty"`

	// VaultBaseUrl: The keyvault base url.
	VaultBaseUrl *string `json:"vaultBaseUrl,omitempty"`
}

type ImageRegistryCredential_STATUSARM struct {
	// Identity: The identity for the private registry.
	Identity *string `json:"identity,omitempty"`

	// IdentityUrl: The identity URL for the private registry.
	IdentityUrl *string `json:"identityUrl,omitempty"`

	// Server: The Docker image registry server without a protocol such as "http" and "https".
	Server *string `json:"server,omitempty"`

	// Username: The username for the private registry.
	Username *string `json:"username,omitempty"`
}

type InitContainerDefinition_STATUSARM struct {
	// Name: The name for the init container.
	Name *string `json:"name,omitempty"`

	// Properties: The properties for the init container.
	Properties *InitContainerPropertiesDefinition_STATUSARM `json:"properties,omitempty"`
}

type IpAddress_STATUSARM struct {
	// DnsNameLabel: The Dns name label for the IP.
	DnsNameLabel *string `json:"dnsNameLabel,omitempty"`

	// DnsNameLabelReusePolicy: The value representing the security enum.
	DnsNameLabelReusePolicy *IpAddress_DnsNameLabelReusePolicy_STATUS `json:"dnsNameLabelReusePolicy,omitempty"`

	// Fqdn: The FQDN for the IP.
	Fqdn *string `json:"fqdn,omitempty"`

	// Ip: The IP exposed to the public internet.
	Ip *string `json:"ip,omitempty"`

	// Ports: The list of ports exposed on the container group.
	Ports []Port_STATUSARM `json:"ports,omitempty"`

	// Type: Specifies if the IP is exposed to the public internet or private VNET.
	Type *IpAddress_Type_STATUS `json:"type,omitempty"`
}

type Volume_STATUSARM struct {
	// AzureFile: The Azure File volume.
	AzureFile *AzureFileVolume_STATUSARM `json:"azureFile,omitempty"`

	// EmptyDir: The empty directory volume.
	EmptyDir map[string]v1.JSON `json:"emptyDir,omitempty"`

	// GitRepo: The git repo volume.
	GitRepo *GitRepoVolume_STATUSARM `json:"gitRepo,omitempty"`

	// Name: The name of the volume.
	Name *string `json:"name,omitempty"`

	// Secret: The secret volume.
	Secret map[string]string `json:"secret,omitempty"`
}

type AzureFileVolume_STATUSARM struct {
	// ReadOnly: The flag indicating whether the Azure File shared mounted as a volume is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`

	// ShareName: The name of the Azure File share to be mounted as a volume.
	ShareName *string `json:"shareName,omitempty"`

	// StorageAccountKey: The storage account access key used to access the Azure File share.
	StorageAccountKey *string `json:"storageAccountKey,omitempty"`

	// StorageAccountName: The name of the storage account that contains the Azure File share.
	StorageAccountName *string `json:"storageAccountName,omitempty"`
}

type ContainerProperties_STATUSARM struct {
	// Command: The commands to execute within the container instance in exec form.
	Command []string `json:"command,omitempty"`

	// EnvironmentVariables: The environment variables to set in the container instance.
	EnvironmentVariables []EnvironmentVariable_STATUSARM `json:"environmentVariables,omitempty"`

	// Image: The name of the image used to create the container instance.
	Image *string `json:"image,omitempty"`

	// InstanceView: The instance view of the container instance. Only valid in response.
	InstanceView *ContainerProperties_InstanceView_STATUSARM `json:"instanceView,omitempty"`

	// LivenessProbe: The liveness probe.
	LivenessProbe *ContainerProbe_STATUSARM `json:"livenessProbe,omitempty"`

	// Ports: The exposed ports on the container instance.
	Ports []ContainerPort_STATUSARM `json:"ports,omitempty"`

	// ReadinessProbe: The readiness probe.
	ReadinessProbe *ContainerProbe_STATUSARM `json:"readinessProbe,omitempty"`

	// Resources: The resource requirements of the container instance.
	Resources *ResourceRequirements_STATUSARM `json:"resources,omitempty"`

	// VolumeMounts: The volume mounts available to the container instance.
	VolumeMounts []VolumeMount_STATUSARM `json:"volumeMounts,omitempty"`
}

type Event_STATUSARM struct {
	// Count: The count of the event.
	Count *int `json:"count,omitempty"`

	// FirstTimestamp: The date-time of the earliest logged event.
	FirstTimestamp *string `json:"firstTimestamp,omitempty"`

	// LastTimestamp: The date-time of the latest logged event.
	LastTimestamp *string `json:"lastTimestamp,omitempty"`

	// Message: The event message.
	Message *string `json:"message,omitempty"`

	// Name: The event name.
	Name *string `json:"name,omitempty"`

	// Type: The event type.
	Type *string `json:"type,omitempty"`
}

type GitRepoVolume_STATUSARM struct {
	// Directory: Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be
	// the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the
	// given name.
	Directory *string `json:"directory,omitempty"`

	// Repository: Repository URL
	Repository *string `json:"repository,omitempty"`

	// Revision: Commit hash for the specified revision.
	Revision *string `json:"revision,omitempty"`
}

type InitContainerPropertiesDefinition_STATUSARM struct {
	// Command: The command to execute within the init container in exec form.
	Command []string `json:"command,omitempty"`

	// EnvironmentVariables: The environment variables to set in the init container.
	EnvironmentVariables []EnvironmentVariable_STATUSARM `json:"environmentVariables,omitempty"`

	// Image: The image of the init container.
	Image *string `json:"image,omitempty"`

	// InstanceView: The instance view of the init container. Only valid in response.
	InstanceView *InitContainerPropertiesDefinition_InstanceView_STATUSARM `json:"instanceView,omitempty"`

	// VolumeMounts: The volume mounts available to the init container.
	VolumeMounts []VolumeMount_STATUSARM `json:"volumeMounts,omitempty"`
}

type LogAnalytics_STATUSARM struct {
	// LogType: The log type to be used.
	LogType *LogAnalytics_LogType_STATUS `json:"logType,omitempty"`

	// Metadata: Metadata for log analytics.
	Metadata map[string]string `json:"metadata,omitempty"`

	// WorkspaceId: The workspace id for log analytics
	WorkspaceId *string `json:"workspaceId,omitempty"`

	// WorkspaceKey: The workspace key for log analytics
	WorkspaceKey *string `json:"workspaceKey,omitempty"`

	// WorkspaceResourceId: The workspace resource id for log analytics
	WorkspaceResourceId *string `json:"workspaceResourceId,omitempty"`
}

type Port_STATUSARM struct {
	// Port: The port number.
	Port *int `json:"port,omitempty"`

	// Protocol: The protocol associated with the port.
	Protocol *Port_Protocol_STATUS `json:"protocol,omitempty"`
}

type ContainerPort_STATUSARM struct {
	// Port: The port number exposed within the container group.
	Port *int `json:"port,omitempty"`

	// Protocol: The protocol associated with the port.
	Protocol *ContainerPort_Protocol_STATUS `json:"protocol,omitempty"`
}

type ContainerProbe_STATUSARM struct {
	// Exec: The execution command to probe
	Exec *ContainerExec_STATUSARM `json:"exec,omitempty"`

	// FailureThreshold: The failure threshold.
	FailureThreshold *int `json:"failureThreshold,omitempty"`

	// HttpGet: The Http Get settings to probe
	HttpGet *ContainerHttpGet_STATUSARM `json:"httpGet,omitempty"`

	// InitialDelaySeconds: The initial delay seconds.
	InitialDelaySeconds *int `json:"initialDelaySeconds,omitempty"`

	// PeriodSeconds: The period seconds.
	PeriodSeconds *int `json:"periodSeconds,omitempty"`

	// SuccessThreshold: The success threshold.
	SuccessThreshold *int `json:"successThreshold,omitempty"`

	// TimeoutSeconds: The timeout seconds.
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
}

type ContainerProperties_InstanceView_STATUSARM struct {
	// CurrentState: Current container instance state.
	CurrentState *ContainerState_STATUSARM `json:"currentState,omitempty"`

	// Events: The events of the container instance.
	Events []Event_STATUSARM `json:"events,omitempty"`

	// PreviousState: Previous container instance state.
	PreviousState *ContainerState_STATUSARM `json:"previousState,omitempty"`

	// RestartCount: The number of times that the container instance has been restarted.
	RestartCount *int `json:"restartCount,omitempty"`
}

type EnvironmentVariable_STATUSARM struct {
	// Name: The name of the environment variable.
	Name *string `json:"name,omitempty"`

	// SecureValue: The value of the secure environment variable.
	SecureValue *string `json:"secureValue,omitempty"`

	// Value: The value of the environment variable.
	Value *string `json:"value,omitempty"`
}

type InitContainerPropertiesDefinition_InstanceView_STATUSARM struct {
	// CurrentState: The current state of the init container.
	CurrentState *ContainerState_STATUSARM `json:"currentState,omitempty"`

	// Events: The events of the init container.
	Events []Event_STATUSARM `json:"events,omitempty"`

	// PreviousState: The previous state of the init container.
	PreviousState *ContainerState_STATUSARM `json:"previousState,omitempty"`

	// RestartCount: The number of times that the init container has been restarted.
	RestartCount *int `json:"restartCount,omitempty"`
}

type ResourceRequirements_STATUSARM struct {
	// Limits: The resource limits of this container instance.
	Limits *ResourceLimits_STATUSARM `json:"limits,omitempty"`

	// Requests: The resource requests of this container instance.
	Requests *ResourceRequests_STATUSARM `json:"requests,omitempty"`
}

type VolumeMount_STATUSARM struct {
	// MountPath: The path within the container where the volume should be mounted. Must not contain colon (:).
	MountPath *string `json:"mountPath,omitempty"`

	// Name: The name of the volume mount.
	Name *string `json:"name,omitempty"`

	// ReadOnly: The flag indicating whether the volume mount is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

type ContainerExec_STATUSARM struct {
	// Command: The commands to execute within the container.
	Command []string `json:"command,omitempty"`
}

type ContainerHttpGet_STATUSARM struct {
	// HttpHeaders: The HTTP headers.
	HttpHeaders []HttpHeader_STATUSARM `json:"httpHeaders,omitempty"`

	// Path: The path to probe.
	Path *string `json:"path,omitempty"`

	// Port: The port number to probe.
	Port *int `json:"port,omitempty"`

	// Scheme: The scheme.
	Scheme *ContainerHttpGet_Scheme_STATUS `json:"scheme,omitempty"`
}

type ContainerState_STATUSARM struct {
	// DetailStatus: The human-readable status of the container instance state.
	DetailStatus *string `json:"detailStatus,omitempty"`

	// ExitCode: The container instance exit codes correspond to those from the `docker run` command.
	ExitCode *int `json:"exitCode,omitempty"`

	// FinishTime: The date-time when the container instance state finished.
	FinishTime *string `json:"finishTime,omitempty"`

	// StartTime: The date-time when the container instance state started.
	StartTime *string `json:"startTime,omitempty"`

	// State: The state of the container instance.
	State *string `json:"state,omitempty"`
}

type ResourceLimits_STATUSARM struct {
	// Cpu: The CPU limit of this container instance.
	Cpu *float64 `json:"cpu,omitempty"`

	// Gpu: The GPU limit of this container instance.
	Gpu *GpuResource_STATUSARM `json:"gpu,omitempty"`

	// MemoryInGB: The memory limit in GB of this container instance.
	MemoryInGB *float64 `json:"memoryInGB,omitempty"`
}

type ResourceRequests_STATUSARM struct {
	// Cpu: The CPU request of this container instance.
	Cpu *float64 `json:"cpu,omitempty"`

	// Gpu: The GPU request of this container instance.
	Gpu *GpuResource_STATUSARM `json:"gpu,omitempty"`

	// MemoryInGB: The memory request in GB of this container instance.
	MemoryInGB *float64 `json:"memoryInGB,omitempty"`
}

type GpuResource_STATUSARM struct {
	// Count: The count of the GPU resource.
	Count *int `json:"count,omitempty"`

	// Sku: The SKU of the GPU resource.
	Sku *GpuResource_Sku_STATUS `json:"sku,omitempty"`
}

type HttpHeader_STATUSARM struct {
	// Name: The header name.
	Name *string `json:"name,omitempty"`

	// Value: The header value.
	Value *string `json:"value,omitempty"`
}
