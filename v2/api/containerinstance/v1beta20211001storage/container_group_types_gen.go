// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=containerinstance.azure.com,resources=containergroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=containerinstance.azure.com,resources={containergroups/status,containergroups/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20211001.ContainerGroup
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/resourceDefinitions/containerGroups
type ContainerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerGroups_Spec  `json:"spec,omitempty"`
	Status            ContainerGroup_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ContainerGroup{}

// GetConditions returns the conditions of the resource
func (group *ContainerGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *ContainerGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ContainerGroup{}

// AzureName returns the Azure name of the resource
func (group *ContainerGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (group ContainerGroup) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (group *ContainerGroup) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (group *ContainerGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *ContainerGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerInstance/containerGroups"
func (group *ContainerGroup) GetType() string {
	return "Microsoft.ContainerInstance/containerGroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *ContainerGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ContainerGroup_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (group *ContainerGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return &genruntime.ResourceReference{
		Group: ownerGroup,
		Kind:  ownerKind,
		Name:  group.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (group *ContainerGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ContainerGroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st ContainerGroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// Hub marks that this ContainerGroup is the hub type for conversion
func (group *ContainerGroup) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *ContainerGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion,
		Kind:    "ContainerGroup",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20211001.ContainerGroup
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/resourceDefinitions/containerGroups
type ContainerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerGroup `json:"items"`
}

// Storage version of v1beta20211001.APIVersion
// +kubebuilder:validation:Enum={"2021-10-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-10-01")

// Storage version of v1beta20211001.ContainerGroup_STATUS
type ContainerGroup_STATUS struct {
	Conditions               []conditions.Condition                         `json:"conditions,omitempty"`
	Containers               []Container_STATUS                             `json:"containers,omitempty"`
	Diagnostics              *ContainerGroupDiagnostics_STATUS              `json:"diagnostics,omitempty"`
	DnsConfig                *DnsConfiguration_STATUS                       `json:"dnsConfig,omitempty"`
	EncryptionProperties     *EncryptionProperties_STATUS                   `json:"encryptionProperties,omitempty"`
	Id                       *string                                        `json:"id,omitempty"`
	Identity                 *ContainerGroupIdentity_STATUS                 `json:"identity,omitempty"`
	ImageRegistryCredentials []ImageRegistryCredential_STATUS               `json:"imageRegistryCredentials,omitempty"`
	InitContainers           []InitContainerDefinition_STATUS               `json:"initContainers,omitempty"`
	InstanceView             *ContainerGroup_STATUS_Properties_InstanceView `json:"instanceView,omitempty"`
	IpAddress                *IpAddress_STATUS                              `json:"ipAddress,omitempty"`
	Location                 *string                                        `json:"location,omitempty"`
	Name                     *string                                        `json:"name,omitempty"`
	OsType                   *string                                        `json:"osType,omitempty"`
	PropertyBag              genruntime.PropertyBag                         `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                                        `json:"provisioningState,omitempty"`
	RestartPolicy            *string                                        `json:"restartPolicy,omitempty"`
	Sku                      *string                                        `json:"sku,omitempty"`
	SubnetIds                []ContainerGroupSubnetId_STATUS                `json:"subnetIds,omitempty"`
	Tags                     map[string]string                              `json:"tags,omitempty"`
	Type                     *string                                        `json:"type,omitempty"`
	Volumes                  []Volume_STATUS                                `json:"volumes,omitempty"`
	Zones                    []string                                       `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ContainerGroup_STATUS{}

// ConvertStatusFrom populates our ContainerGroup_STATUS from the provided source
func (group *ContainerGroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(group)
}

// ConvertStatusTo populates the provided destination from our ContainerGroup_STATUS
func (group *ContainerGroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(group)
}

// Storage version of v1beta20211001.ContainerGroups_Spec
type ContainerGroups_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                string                                                     `json:"azureName,omitempty"`
	Containers               []ContainerGroups_Spec_Properties_Containers               `json:"containers,omitempty"`
	Diagnostics              *ContainerGroupDiagnostics                                 `json:"diagnostics,omitempty"`
	DnsConfig                *DnsConfiguration                                          `json:"dnsConfig,omitempty"`
	EncryptionProperties     *EncryptionProperties                                      `json:"encryptionProperties,omitempty"`
	Identity                 *ContainerGroupIdentity                                    `json:"identity,omitempty"`
	ImageRegistryCredentials []ContainerGroups_Spec_Properties_ImageRegistryCredentials `json:"imageRegistryCredentials,omitempty"`
	InitContainers           []ContainerGroups_Spec_Properties_InitContainers           `json:"initContainers,omitempty"`
	IpAddress                *IpAddress                                                 `json:"ipAddress,omitempty"`
	Location                 *string                                                    `json:"location,omitempty"`
	OriginalVersion          string                                                     `json:"originalVersion,omitempty"`
	OsType                   *string                                                    `json:"osType,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner         *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag   genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RestartPolicy *string                            `json:"restartPolicy,omitempty"`
	Sku           *string                            `json:"sku,omitempty"`
	SubnetIds     []ContainerGroupSubnetId           `json:"subnetIds,omitempty"`
	Tags          map[string]string                  `json:"tags,omitempty"`
	Volumes       []Volume                           `json:"volumes,omitempty"`
	Zones         []string                           `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ContainerGroups_Spec{}

// ConvertSpecFrom populates our ContainerGroups_Spec from the provided source
func (groups *ContainerGroups_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == groups {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(groups)
}

// ConvertSpecTo populates the provided destination from our ContainerGroups_Spec
func (groups *ContainerGroups_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == groups {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(groups)
}

// Storage version of v1beta20211001.Container_STATUS
type Container_STATUS struct {
	Command              []string                                 `json:"command,omitempty"`
	EnvironmentVariables []EnvironmentVariable_STATUS             `json:"environmentVariables,omitempty"`
	Image                *string                                  `json:"image,omitempty"`
	InstanceView         *ContainerProperties_STATUS_InstanceView `json:"instanceView,omitempty"`
	LivenessProbe        *ContainerProbe_STATUS                   `json:"livenessProbe,omitempty"`
	Name                 *string                                  `json:"name,omitempty"`
	Ports                []ContainerPort_STATUS                   `json:"ports,omitempty"`
	PropertyBag          genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	ReadinessProbe       *ContainerProbe_STATUS                   `json:"readinessProbe,omitempty"`
	Resources            *ResourceRequirements_STATUS             `json:"resources,omitempty"`
	VolumeMounts         []VolumeMount_STATUS                     `json:"volumeMounts,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroup_STATUS_Properties_InstanceView
type ContainerGroup_STATUS_Properties_InstanceView struct {
	Events      []Event_STATUS         `json:"events,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroupDiagnostics
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ContainerGroupDiagnostics
type ContainerGroupDiagnostics struct {
	LogAnalytics *LogAnalytics          `json:"logAnalytics,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroupDiagnostics_STATUS
type ContainerGroupDiagnostics_STATUS struct {
	LogAnalytics *LogAnalytics_STATUS   `json:"logAnalytics,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroupIdentity
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ContainerGroupIdentity
type ContainerGroupIdentity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroupIdentity_STATUS
type ContainerGroupIdentity_STATUS struct {
	PrincipalId            *string                                                         `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                                          `json:"$propertyBag,omitempty"`
	TenantId               *string                                                         `json:"tenantId,omitempty"`
	Type                   *string                                                         `json:"type,omitempty"`
	UserAssignedIdentities map[string]ContainerGroupIdentity_STATUS_UserAssignedIdentities `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroups_Spec_Properties_Containers
type ContainerGroups_Spec_Properties_Containers struct {
	Command              []string               `json:"command,omitempty"`
	EnvironmentVariables []EnvironmentVariable  `json:"environmentVariables,omitempty"`
	Image                *string                `json:"image,omitempty"`
	LivenessProbe        *ContainerProbe        `json:"livenessProbe,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Ports                []ContainerPort        `json:"ports,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReadinessProbe       *ContainerProbe        `json:"readinessProbe,omitempty"`
	Resources            *ResourceRequirements  `json:"resources,omitempty"`
	VolumeMounts         []VolumeMount          `json:"volumeMounts,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroups_Spec_Properties_ImageRegistryCredentials
type ContainerGroups_Spec_Properties_ImageRegistryCredentials struct {
	Identity    *string                     `json:"identity,omitempty"`
	IdentityUrl *string                     `json:"identityUrl,omitempty"`
	Password    *genruntime.SecretReference `json:"password,omitempty"`
	PropertyBag genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	Server      *string                     `json:"server,omitempty"`
	Username    *string                     `json:"username,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroups_Spec_Properties_InitContainers
type ContainerGroups_Spec_Properties_InitContainers struct {
	Command              []string               `json:"command,omitempty"`
	EnvironmentVariables []EnvironmentVariable  `json:"environmentVariables,omitempty"`
	Image                *string                `json:"image,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VolumeMounts         []VolumeMount          `json:"volumeMounts,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroupSubnetId
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ContainerGroupSubnetId
type ContainerGroupSubnetId struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: Resource ID of virtual network and subnet.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroupSubnetId_STATUS
type ContainerGroupSubnetId_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.DnsConfiguration
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/DnsConfiguration
type DnsConfiguration struct {
	NameServers   []string               `json:"nameServers,omitempty"`
	Options       *string                `json:"options,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SearchDomains *string                `json:"searchDomains,omitempty"`
}

// Storage version of v1beta20211001.DnsConfiguration_STATUS
type DnsConfiguration_STATUS struct {
	NameServers   []string               `json:"nameServers,omitempty"`
	Options       *string                `json:"options,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SearchDomains *string                `json:"searchDomains,omitempty"`
}

// Storage version of v1beta20211001.EncryptionProperties
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/EncryptionProperties
type EncryptionProperties struct {
	KeyName      *string                `json:"keyName,omitempty"`
	KeyVersion   *string                `json:"keyVersion,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VaultBaseUrl *string                `json:"vaultBaseUrl,omitempty"`
}

// Storage version of v1beta20211001.EncryptionProperties_STATUS
type EncryptionProperties_STATUS struct {
	KeyName      *string                `json:"keyName,omitempty"`
	KeyVersion   *string                `json:"keyVersion,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VaultBaseUrl *string                `json:"vaultBaseUrl,omitempty"`
}

// Storage version of v1beta20211001.ImageRegistryCredential_STATUS
type ImageRegistryCredential_STATUS struct {
	Identity    *string                `json:"identity,omitempty"`
	IdentityUrl *string                `json:"identityUrl,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Server      *string                `json:"server,omitempty"`
	Username    *string                `json:"username,omitempty"`
}

// Storage version of v1beta20211001.InitContainerDefinition_STATUS
type InitContainerDefinition_STATUS struct {
	Command              []string                                               `json:"command,omitempty"`
	EnvironmentVariables []EnvironmentVariable_STATUS                           `json:"environmentVariables,omitempty"`
	Image                *string                                                `json:"image,omitempty"`
	InstanceView         *InitContainerPropertiesDefinition_STATUS_InstanceView `json:"instanceView,omitempty"`
	Name                 *string                                                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	VolumeMounts         []VolumeMount_STATUS                                   `json:"volumeMounts,omitempty"`
}

// Storage version of v1beta20211001.IpAddress
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/IpAddress
type IpAddress struct {
	AutoGeneratedDomainNameLabelScope *string                `json:"autoGeneratedDomainNameLabelScope,omitempty"`
	DnsNameLabel                      *string                `json:"dnsNameLabel,omitempty"`
	Ip                                *string                `json:"ip,omitempty"`
	Ports                             []Port                 `json:"ports,omitempty"`
	PropertyBag                       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                              *string                `json:"type,omitempty"`
}

// Storage version of v1beta20211001.IpAddress_STATUS
type IpAddress_STATUS struct {
	DnsNameLabel            *string                `json:"dnsNameLabel,omitempty"`
	DnsNameLabelReusePolicy *string                `json:"dnsNameLabelReusePolicy,omitempty"`
	Fqdn                    *string                `json:"fqdn,omitempty"`
	Ip                      *string                `json:"ip,omitempty"`
	Ports                   []Port_STATUS          `json:"ports,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

// Storage version of v1beta20211001.Volume
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/Volume
type Volume struct {
	AzureFile   *AzureFileVolume       `json:"azureFile,omitempty"`
	EmptyDir    map[string]v1.JSON     `json:"emptyDir,omitempty"`
	GitRepo     *GitRepoVolume         `json:"gitRepo,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Secret      map[string]string      `json:"secret,omitempty"`
}

// Storage version of v1beta20211001.Volume_STATUS
type Volume_STATUS struct {
	AzureFile   *AzureFileVolume_STATUS `json:"azureFile,omitempty"`
	EmptyDir    map[string]v1.JSON      `json:"emptyDir,omitempty"`
	GitRepo     *GitRepoVolume_STATUS   `json:"gitRepo,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	Secret      map[string]string       `json:"secret,omitempty"`
}

// Storage version of v1beta20211001.AzureFileVolume
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/AzureFileVolume
type AzureFileVolume struct {
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReadOnly           *bool                  `json:"readOnly,omitempty"`
	ShareName          *string                `json:"shareName,omitempty"`
	StorageAccountKey  *string                `json:"storageAccountKey,omitempty"`
	StorageAccountName *string                `json:"storageAccountName,omitempty"`
}

// Storage version of v1beta20211001.AzureFileVolume_STATUS
type AzureFileVolume_STATUS struct {
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReadOnly           *bool                  `json:"readOnly,omitempty"`
	ShareName          *string                `json:"shareName,omitempty"`
	StorageAccountKey  *string                `json:"storageAccountKey,omitempty"`
	StorageAccountName *string                `json:"storageAccountName,omitempty"`
}

// Storage version of v1beta20211001.ContainerGroupIdentity_STATUS_UserAssignedIdentities
type ContainerGroupIdentity_STATUS_UserAssignedIdentities struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ContainerPort
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ContainerPort
type ContainerPort struct {
	Port        *int                   `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol    *string                `json:"protocol,omitempty"`
}

// Storage version of v1beta20211001.ContainerPort_STATUS
type ContainerPort_STATUS struct {
	Port        *int                   `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol    *string                `json:"protocol,omitempty"`
}

// Storage version of v1beta20211001.ContainerProbe
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ContainerProbe
type ContainerProbe struct {
	Exec                *ContainerExec         `json:"exec,omitempty"`
	FailureThreshold    *int                   `json:"failureThreshold,omitempty"`
	HttpGet             *ContainerHttpGet      `json:"httpGet,omitempty"`
	InitialDelaySeconds *int                   `json:"initialDelaySeconds,omitempty"`
	PeriodSeconds       *int                   `json:"periodSeconds,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SuccessThreshold    *int                   `json:"successThreshold,omitempty"`
	TimeoutSeconds      *int                   `json:"timeoutSeconds,omitempty"`
}

// Storage version of v1beta20211001.ContainerProbe_STATUS
type ContainerProbe_STATUS struct {
	Exec                *ContainerExec_STATUS    `json:"exec,omitempty"`
	FailureThreshold    *int                     `json:"failureThreshold,omitempty"`
	HttpGet             *ContainerHttpGet_STATUS `json:"httpGet,omitempty"`
	InitialDelaySeconds *int                     `json:"initialDelaySeconds,omitempty"`
	PeriodSeconds       *int                     `json:"periodSeconds,omitempty"`
	PropertyBag         genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	SuccessThreshold    *int                     `json:"successThreshold,omitempty"`
	TimeoutSeconds      *int                     `json:"timeoutSeconds,omitempty"`
}

// Storage version of v1beta20211001.ContainerProperties_STATUS_InstanceView
type ContainerProperties_STATUS_InstanceView struct {
	CurrentState  *ContainerState_STATUS `json:"currentState,omitempty"`
	Events        []Event_STATUS         `json:"events,omitempty"`
	PreviousState *ContainerState_STATUS `json:"previousState,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RestartCount  *int                   `json:"restartCount,omitempty"`
}

// Storage version of v1beta20211001.EnvironmentVariable
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/EnvironmentVariable
type EnvironmentVariable struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SecureValue *string                `json:"secureValue,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20211001.EnvironmentVariable_STATUS
type EnvironmentVariable_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SecureValue *string                `json:"secureValue,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20211001.Event_STATUS
type Event_STATUS struct {
	Count          *int                   `json:"count,omitempty"`
	FirstTimestamp *string                `json:"firstTimestamp,omitempty"`
	LastTimestamp  *string                `json:"lastTimestamp,omitempty"`
	Message        *string                `json:"message,omitempty"`
	Name           *string                `json:"name,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type           *string                `json:"type,omitempty"`
}

// Storage version of v1beta20211001.GitRepoVolume
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/GitRepoVolume
type GitRepoVolume struct {
	Directory   *string                `json:"directory,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Repository  *string                `json:"repository,omitempty"`
	Revision    *string                `json:"revision,omitempty"`
}

// Storage version of v1beta20211001.GitRepoVolume_STATUS
type GitRepoVolume_STATUS struct {
	Directory   *string                `json:"directory,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Repository  *string                `json:"repository,omitempty"`
	Revision    *string                `json:"revision,omitempty"`
}

// Storage version of v1beta20211001.InitContainerPropertiesDefinition_STATUS_InstanceView
type InitContainerPropertiesDefinition_STATUS_InstanceView struct {
	CurrentState  *ContainerState_STATUS `json:"currentState,omitempty"`
	Events        []Event_STATUS         `json:"events,omitempty"`
	PreviousState *ContainerState_STATUS `json:"previousState,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RestartCount  *int                   `json:"restartCount,omitempty"`
}

// Storage version of v1beta20211001.LogAnalytics
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/LogAnalytics
type LogAnalytics struct {
	LogType      *string                `json:"logType,omitempty"`
	Metadata     map[string]string      `json:"metadata,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WorkspaceId  *string                `json:"workspaceId,omitempty"`
	WorkspaceKey *string                `json:"workspaceKey,omitempty"`

	// WorkspaceResourceReference: The workspace resource id for log analytics
	WorkspaceResourceReference *genruntime.ResourceReference `armReference:"WorkspaceResourceId" json:"workspaceResourceReference,omitempty"`
}

// Storage version of v1beta20211001.LogAnalytics_STATUS
type LogAnalytics_STATUS struct {
	LogType             *string                `json:"logType,omitempty"`
	Metadata            map[string]string      `json:"metadata,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WorkspaceId         *string                `json:"workspaceId,omitempty"`
	WorkspaceKey        *string                `json:"workspaceKey,omitempty"`
	WorkspaceResourceId *string                `json:"workspaceResourceId,omitempty"`
}

// Storage version of v1beta20211001.Port
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/Port
type Port struct {
	Port        *int                   `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol    *string                `json:"protocol,omitempty"`
}

// Storage version of v1beta20211001.Port_STATUS
type Port_STATUS struct {
	Port        *int                   `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol    *string                `json:"protocol,omitempty"`
}

// Storage version of v1beta20211001.ResourceRequirements
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ResourceRequirements
type ResourceRequirements struct {
	Limits      *ResourceLimits        `json:"limits,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Requests    *ResourceRequests      `json:"requests,omitempty"`
}

// Storage version of v1beta20211001.ResourceRequirements_STATUS
type ResourceRequirements_STATUS struct {
	Limits      *ResourceLimits_STATUS   `json:"limits,omitempty"`
	PropertyBag genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	Requests    *ResourceRequests_STATUS `json:"requests,omitempty"`
}

// Storage version of v1beta20211001.VolumeMount
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/VolumeMount
type VolumeMount struct {
	MountPath   *string                `json:"mountPath,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReadOnly    *bool                  `json:"readOnly,omitempty"`
}

// Storage version of v1beta20211001.VolumeMount_STATUS
type VolumeMount_STATUS struct {
	MountPath   *string                `json:"mountPath,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReadOnly    *bool                  `json:"readOnly,omitempty"`
}

// Storage version of v1beta20211001.ContainerExec
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ContainerExec
type ContainerExec struct {
	Command     []string               `json:"command,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ContainerExec_STATUS
type ContainerExec_STATUS struct {
	Command     []string               `json:"command,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ContainerHttpGet
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ContainerHttpGet
type ContainerHttpGet struct {
	HttpHeaders []HttpHeader           `json:"httpHeaders,omitempty"`
	Path        *string                `json:"path,omitempty"`
	Port        *int                   `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Scheme      *string                `json:"scheme,omitempty"`
}

// Storage version of v1beta20211001.ContainerHttpGet_STATUS
type ContainerHttpGet_STATUS struct {
	HttpHeaders []HttpHeader_STATUS    `json:"httpHeaders,omitempty"`
	Path        *string                `json:"path,omitempty"`
	Port        *int                   `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Scheme      *string                `json:"scheme,omitempty"`
}

// Storage version of v1beta20211001.ContainerState_STATUS
type ContainerState_STATUS struct {
	DetailStatus *string                `json:"detailStatus,omitempty"`
	ExitCode     *int                   `json:"exitCode,omitempty"`
	FinishTime   *string                `json:"finishTime,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartTime    *string                `json:"startTime,omitempty"`
	State        *string                `json:"state,omitempty"`
}

// Storage version of v1beta20211001.ResourceLimits
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ResourceLimits
type ResourceLimits struct {
	Cpu         *float64               `json:"cpu,omitempty"`
	Gpu         *GpuResource           `json:"gpu,omitempty"`
	MemoryInGB  *float64               `json:"memoryInGB,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ResourceLimits_STATUS
type ResourceLimits_STATUS struct {
	Cpu         *float64               `json:"cpu,omitempty"`
	Gpu         *GpuResource_STATUS    `json:"gpu,omitempty"`
	MemoryInGB  *float64               `json:"memoryInGB,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ResourceRequests
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/ResourceRequests
type ResourceRequests struct {
	Cpu         *float64               `json:"cpu,omitempty"`
	Gpu         *GpuResource           `json:"gpu,omitempty"`
	MemoryInGB  *float64               `json:"memoryInGB,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ResourceRequests_STATUS
type ResourceRequests_STATUS struct {
	Cpu         *float64               `json:"cpu,omitempty"`
	Gpu         *GpuResource_STATUS    `json:"gpu,omitempty"`
	MemoryInGB  *float64               `json:"memoryInGB,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.GpuResource
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/GpuResource
type GpuResource struct {
	Count       *int                   `json:"count,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sku         *string                `json:"sku,omitempty"`
}

// Storage version of v1beta20211001.GpuResource_STATUS
type GpuResource_STATUS struct {
	Count       *int                   `json:"count,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sku         *string                `json:"sku,omitempty"`
}

// Storage version of v1beta20211001.HttpHeader
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.ContainerInstance.json#/definitions/HttpHeader
type HttpHeader struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20211001.HttpHeader_STATUS
type HttpHeader_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ContainerGroup{}, &ContainerGroupList{})
}
