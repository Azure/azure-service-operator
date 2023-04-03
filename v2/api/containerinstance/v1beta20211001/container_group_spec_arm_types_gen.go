// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Deprecated version of ContainerGroup_Spec. Use v1api20211001.ContainerGroup_Spec instead
type ContainerGroup_Spec_ARM struct {
	Identity   *ContainerGroupIdentity_ARM         `json:"identity,omitempty"`
	Location   *string                             `json:"location,omitempty"`
	Name       string                              `json:"name,omitempty"`
	Properties *ContainerGroup_Properties_Spec_ARM `json:"properties,omitempty"`
	Tags       map[string]string                   `json:"tags,omitempty"`
	Zones      []string                            `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ContainerGroup_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (group ContainerGroup_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (group *ContainerGroup_Spec_ARM) GetName() string {
	return group.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerInstance/containerGroups"
func (group *ContainerGroup_Spec_ARM) GetType() string {
	return "Microsoft.ContainerInstance/containerGroups"
}

// Deprecated version of ContainerGroup_Properties_Spec. Use v1api20211001.ContainerGroup_Properties_Spec instead
type ContainerGroup_Properties_Spec_ARM struct {
	Containers               []Container_ARM                               `json:"containers,omitempty"`
	Diagnostics              *ContainerGroupDiagnostics_ARM                `json:"diagnostics,omitempty"`
	DnsConfig                *DnsConfiguration_ARM                         `json:"dnsConfig,omitempty"`
	EncryptionProperties     *EncryptionProperties_ARM                     `json:"encryptionProperties,omitempty"`
	ImageRegistryCredentials []ImageRegistryCredential_ARM                 `json:"imageRegistryCredentials,omitempty"`
	InitContainers           []InitContainerDefinition_ARM                 `json:"initContainers,omitempty"`
	IpAddress                *IpAddress_ARM                                `json:"ipAddress,omitempty"`
	OsType                   *ContainerGroup_Properties_OsType_Spec        `json:"osType,omitempty"`
	RestartPolicy            *ContainerGroup_Properties_RestartPolicy_Spec `json:"restartPolicy,omitempty"`
	Sku                      *ContainerGroupSku                            `json:"sku,omitempty"`
	SubnetIds                []ContainerGroupSubnetId_ARM                  `json:"subnetIds,omitempty"`
	Volumes                  []Volume_ARM                                  `json:"volumes,omitempty"`
}

// Deprecated version of ContainerGroupIdentity. Use v1api20211001.ContainerGroupIdentity instead
type ContainerGroupIdentity_ARM struct {
	Type *ContainerGroupIdentity_Type `json:"type,omitempty"`
}

// Deprecated version of Container. Use v1api20211001.Container instead
type Container_ARM struct {
	Name       *string                  `json:"name,omitempty"`
	Properties *ContainerProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of ContainerGroupDiagnostics. Use v1api20211001.ContainerGroupDiagnostics instead
type ContainerGroupDiagnostics_ARM struct {
	LogAnalytics *LogAnalytics_ARM `json:"logAnalytics,omitempty"`
}

// Deprecated version of ContainerGroupIdentity_Type. Use v1api20211001.ContainerGroupIdentity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type ContainerGroupIdentity_Type string

const (
	ContainerGroupIdentity_Type_None                       = ContainerGroupIdentity_Type("None")
	ContainerGroupIdentity_Type_SystemAssigned             = ContainerGroupIdentity_Type("SystemAssigned")
	ContainerGroupIdentity_Type_SystemAssignedUserAssigned = ContainerGroupIdentity_Type("SystemAssigned, UserAssigned")
	ContainerGroupIdentity_Type_UserAssigned               = ContainerGroupIdentity_Type("UserAssigned")
)

// Deprecated version of ContainerGroupSubnetId. Use v1api20211001.ContainerGroupSubnetId instead
type ContainerGroupSubnetId_ARM struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Deprecated version of DnsConfiguration. Use v1api20211001.DnsConfiguration instead
type DnsConfiguration_ARM struct {
	NameServers   []string `json:"nameServers,omitempty"`
	Options       *string  `json:"options,omitempty"`
	SearchDomains *string  `json:"searchDomains,omitempty"`
}

// Deprecated version of EncryptionProperties. Use v1api20211001.EncryptionProperties instead
type EncryptionProperties_ARM struct {
	KeyName      *string `json:"keyName,omitempty"`
	KeyVersion   *string `json:"keyVersion,omitempty"`
	VaultBaseUrl *string `json:"vaultBaseUrl,omitempty"`
}

// Deprecated version of ImageRegistryCredential. Use v1api20211001.ImageRegistryCredential instead
type ImageRegistryCredential_ARM struct {
	Identity    *string `json:"identity,omitempty"`
	IdentityUrl *string `json:"identityUrl,omitempty"`
	Password    *string `json:"password,omitempty"`
	Server      *string `json:"server,omitempty"`
	Username    *string `json:"username,omitempty"`
}

// Deprecated version of InitContainerDefinition. Use v1api20211001.InitContainerDefinition instead
type InitContainerDefinition_ARM struct {
	Name       *string                                `json:"name,omitempty"`
	Properties *InitContainerPropertiesDefinition_ARM `json:"properties,omitempty"`
}

// Deprecated version of IpAddress. Use v1api20211001.IpAddress instead
type IpAddress_ARM struct {
	AutoGeneratedDomainNameLabelScope *IpAddress_AutoGeneratedDomainNameLabelScope `json:"autoGeneratedDomainNameLabelScope,omitempty"`
	DnsNameLabel                      *string                                      `json:"dnsNameLabel,omitempty"`
	Ip                                *string                                      `json:"ip,omitempty"`
	Ports                             []Port_ARM                                   `json:"ports,omitempty"`
	Type                              *IpAddress_Type                              `json:"type,omitempty"`
}

// Deprecated version of Volume. Use v1api20211001.Volume instead
type Volume_ARM struct {
	AzureFile *AzureFileVolume_ARM `json:"azureFile,omitempty"`
	EmptyDir  map[string]v1.JSON   `json:"emptyDir,omitempty"`
	GitRepo   *GitRepoVolume_ARM   `json:"gitRepo,omitempty"`
	Name      *string              `json:"name,omitempty"`
	Secret    map[string]string    `json:"secret,omitempty"`
}

// Deprecated version of AzureFileVolume. Use v1api20211001.AzureFileVolume instead
type AzureFileVolume_ARM struct {
	ReadOnly           *bool   `json:"readOnly,omitempty"`
	ShareName          *string `json:"shareName,omitempty"`
	StorageAccountKey  *string `json:"storageAccountKey,omitempty"`
	StorageAccountName *string `json:"storageAccountName,omitempty"`
}

// Deprecated version of ContainerProperties. Use v1api20211001.ContainerProperties instead
type ContainerProperties_ARM struct {
	Command              []string                  `json:"command,omitempty"`
	EnvironmentVariables []EnvironmentVariable_ARM `json:"environmentVariables,omitempty"`
	Image                *string                   `json:"image,omitempty"`
	LivenessProbe        *ContainerProbe_ARM       `json:"livenessProbe,omitempty"`
	Ports                []ContainerPort_ARM       `json:"ports,omitempty"`
	ReadinessProbe       *ContainerProbe_ARM       `json:"readinessProbe,omitempty"`
	Resources            *ResourceRequirements_ARM `json:"resources,omitempty"`
	VolumeMounts         []VolumeMount_ARM         `json:"volumeMounts,omitempty"`
}

// Deprecated version of GitRepoVolume. Use v1api20211001.GitRepoVolume instead
type GitRepoVolume_ARM struct {
	Directory  *string `json:"directory,omitempty"`
	Repository *string `json:"repository,omitempty"`
	Revision   *string `json:"revision,omitempty"`
}

// Deprecated version of InitContainerPropertiesDefinition. Use v1api20211001.InitContainerPropertiesDefinition instead
type InitContainerPropertiesDefinition_ARM struct {
	Command              []string                  `json:"command,omitempty"`
	EnvironmentVariables []EnvironmentVariable_ARM `json:"environmentVariables,omitempty"`
	Image                *string                   `json:"image,omitempty"`
	VolumeMounts         []VolumeMount_ARM         `json:"volumeMounts,omitempty"`
}

// Deprecated version of LogAnalytics. Use v1api20211001.LogAnalytics instead
type LogAnalytics_ARM struct {
	LogType             *LogAnalytics_LogType `json:"logType,omitempty"`
	Metadata            map[string]string     `json:"metadata,omitempty"`
	WorkspaceId         *string               `json:"workspaceId,omitempty"`
	WorkspaceKey        string                `json:"workspaceKey,omitempty"`
	WorkspaceResourceId *string               `json:"workspaceResourceId,omitempty"`
}

// Deprecated version of Port. Use v1api20211001.Port instead
type Port_ARM struct {
	Port     *int           `json:"port,omitempty"`
	Protocol *Port_Protocol `json:"protocol,omitempty"`
}

// Deprecated version of ContainerPort. Use v1api20211001.ContainerPort instead
type ContainerPort_ARM struct {
	Port     *int                    `json:"port,omitempty"`
	Protocol *ContainerPort_Protocol `json:"protocol,omitempty"`
}

// Deprecated version of ContainerProbe. Use v1api20211001.ContainerProbe instead
type ContainerProbe_ARM struct {
	Exec                *ContainerExec_ARM    `json:"exec,omitempty"`
	FailureThreshold    *int                  `json:"failureThreshold,omitempty"`
	HttpGet             *ContainerHttpGet_ARM `json:"httpGet,omitempty"`
	InitialDelaySeconds *int                  `json:"initialDelaySeconds,omitempty"`
	PeriodSeconds       *int                  `json:"periodSeconds,omitempty"`
	SuccessThreshold    *int                  `json:"successThreshold,omitempty"`
	TimeoutSeconds      *int                  `json:"timeoutSeconds,omitempty"`
}

// Deprecated version of EnvironmentVariable. Use v1api20211001.EnvironmentVariable instead
type EnvironmentVariable_ARM struct {
	Name        *string `json:"name,omitempty"`
	SecureValue *string `json:"secureValue,omitempty"`
	Value       *string `json:"value,omitempty"`
}

// Deprecated version of ResourceRequirements. Use v1api20211001.ResourceRequirements instead
type ResourceRequirements_ARM struct {
	Limits   *ResourceLimits_ARM   `json:"limits,omitempty"`
	Requests *ResourceRequests_ARM `json:"requests,omitempty"`
}

// Deprecated version of VolumeMount. Use v1api20211001.VolumeMount instead
type VolumeMount_ARM struct {
	MountPath *string `json:"mountPath,omitempty"`
	Name      *string `json:"name,omitempty"`
	ReadOnly  *bool   `json:"readOnly,omitempty"`
}

// Deprecated version of ContainerExec. Use v1api20211001.ContainerExec instead
type ContainerExec_ARM struct {
	Command []string `json:"command,omitempty"`
}

// Deprecated version of ContainerHttpGet. Use v1api20211001.ContainerHttpGet instead
type ContainerHttpGet_ARM struct {
	HttpHeaders []HttpHeader_ARM         `json:"httpHeaders,omitempty"`
	Path        *string                  `json:"path,omitempty"`
	Port        *int                     `json:"port,omitempty"`
	Scheme      *ContainerHttpGet_Scheme `json:"scheme,omitempty"`
}

// Deprecated version of ResourceLimits. Use v1api20211001.ResourceLimits instead
type ResourceLimits_ARM struct {
	Cpu        *float64         `json:"cpu,omitempty"`
	Gpu        *GpuResource_ARM `json:"gpu,omitempty"`
	MemoryInGB *float64         `json:"memoryInGB,omitempty"`
}

// Deprecated version of ResourceRequests. Use v1api20211001.ResourceRequests instead
type ResourceRequests_ARM struct {
	Cpu        *float64         `json:"cpu,omitempty"`
	Gpu        *GpuResource_ARM `json:"gpu,omitempty"`
	MemoryInGB *float64         `json:"memoryInGB,omitempty"`
}

// Deprecated version of GpuResource. Use v1api20211001.GpuResource instead
type GpuResource_ARM struct {
	Count *int             `json:"count,omitempty"`
	Sku   *GpuResource_Sku `json:"sku,omitempty"`
}

// Deprecated version of HttpHeader. Use v1api20211001.HttpHeader instead
type HttpHeader_ARM struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
