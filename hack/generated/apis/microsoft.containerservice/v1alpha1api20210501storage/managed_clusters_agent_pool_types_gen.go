// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501storage

import (
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210501.ManagedClustersAgentPool
//Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/resourceDefinitions/managedClusters_agentPools
type ManagedClustersAgentPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedClustersAgentPools_Spec `json:"spec,omitempty"`
	Status            AgentPool_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ManagedClustersAgentPool{}

// GetConditions returns the conditions of the resource
func (managedClustersAgentPool *ManagedClustersAgentPool) GetConditions() conditions.Conditions {
	return managedClustersAgentPool.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (managedClustersAgentPool *ManagedClustersAgentPool) SetConditions(conditions conditions.Conditions) {
	managedClustersAgentPool.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ManagedClustersAgentPool{}

// AzureName returns the Azure name of the resource
func (managedClustersAgentPool *ManagedClustersAgentPool) AzureName() string {
	return managedClustersAgentPool.Spec.AzureName
}

// GetSpec returns the specification of this resource
func (managedClustersAgentPool *ManagedClustersAgentPool) GetSpec() genruntime.ConvertibleSpec {
	return &managedClustersAgentPool.Spec
}

// GetStatus returns the status of this resource
func (managedClustersAgentPool *ManagedClustersAgentPool) GetStatus() genruntime.ConvertibleStatus {
	return &managedClustersAgentPool.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters/agentPools"
func (managedClustersAgentPool *ManagedClustersAgentPool) GetType() string {
	return "Microsoft.ContainerService/managedClusters/agentPools"
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (managedClustersAgentPool *ManagedClustersAgentPool) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(managedClustersAgentPool.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: managedClustersAgentPool.Namespace, Name: managedClustersAgentPool.Spec.Owner.Name}
}

// SetStatus sets the status of this resource
func (managedClustersAgentPool *ManagedClustersAgentPool) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AgentPool_Status); ok {
		managedClustersAgentPool.Status = *st
		return nil
	}

	// Convert status to required version
	var st AgentPool_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	managedClustersAgentPool.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (managedClustersAgentPool *ManagedClustersAgentPool) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: managedClustersAgentPool.Spec.OriginalVersion,
		Kind:    "ManagedClustersAgentPool",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210501.ManagedClustersAgentPool
//Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/resourceDefinitions/managedClusters_agentPools
type ManagedClustersAgentPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedClustersAgentPool `json:"items"`
}

//Storage version of v1alpha1api20210501.AgentPool_Status
//Generated from:
type AgentPool_Status struct {
	AvailabilityZones         []string                         `json:"availabilityZones,omitempty"`
	Conditions                []conditions.Condition           `json:"conditions,omitempty"`
	Count                     *int                             `json:"count,omitempty"`
	EnableAutoScaling         *bool                            `json:"enableAutoScaling,omitempty"`
	EnableEncryptionAtHost    *bool                            `json:"enableEncryptionAtHost,omitempty"`
	EnableFIPS                *bool                            `json:"enableFIPS,omitempty"`
	EnableNodePublicIP        *bool                            `json:"enableNodePublicIP,omitempty"`
	EnableUltraSSD            *bool                            `json:"enableUltraSSD,omitempty"`
	GpuInstanceProfile        *string                          `json:"gpuInstanceProfile,omitempty"`
	Id                        *string                          `json:"id,omitempty"`
	KubeletConfig             *KubeletConfig_Status            `json:"kubeletConfig,omitempty"`
	KubeletDiskType           *string                          `json:"kubeletDiskType,omitempty"`
	LinuxOSConfig             *LinuxOSConfig_Status            `json:"linuxOSConfig,omitempty"`
	MaxCount                  *int                             `json:"maxCount,omitempty"`
	MaxPods                   *int                             `json:"maxPods,omitempty"`
	MinCount                  *int                             `json:"minCount,omitempty"`
	Mode                      *string                          `json:"mode,omitempty"`
	Name                      *string                          `json:"name,omitempty"`
	NodeImageVersion          *string                          `json:"nodeImageVersion,omitempty"`
	NodeLabels                map[string]string                `json:"nodeLabels,omitempty"`
	NodePublicIPPrefixID      *string                          `json:"nodePublicIPPrefixID,omitempty"`
	NodeTaints                []string                         `json:"nodeTaints,omitempty"`
	OrchestratorVersion       *string                          `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB              *int                             `json:"osDiskSizeGB,omitempty"`
	OsDiskType                *string                          `json:"osDiskType,omitempty"`
	OsSKU                     *string                          `json:"osSKU,omitempty"`
	OsType                    *string                          `json:"osType,omitempty"`
	PodSubnetID               *string                          `json:"podSubnetID,omitempty"`
	PowerState                *PowerState_Status               `json:"powerState,omitempty"`
	PropertiesType            *string                          `json:"properties_type,omitempty"`
	PropertyBag               genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                          `json:"provisioningState,omitempty"`
	ProximityPlacementGroupID *string                          `json:"proximityPlacementGroupID,omitempty"`
	ScaleSetEvictionPolicy    *string                          `json:"scaleSetEvictionPolicy,omitempty"`
	ScaleSetPriority          *string                          `json:"scaleSetPriority,omitempty"`
	SpotMaxPrice              *float64                         `json:"spotMaxPrice,omitempty"`
	Tags                      map[string]string                `json:"tags,omitempty"`
	Type                      *string                          `json:"type,omitempty"`
	UpgradeSettings           *AgentPoolUpgradeSettings_Status `json:"upgradeSettings,omitempty"`
	VmSize                    *string                          `json:"vmSize,omitempty"`
	VnetSubnetID              *string                          `json:"vnetSubnetID,omitempty"`
}

var _ genruntime.ConvertibleStatus = &AgentPool_Status{}

// ConvertStatusFrom populates our AgentPool_Status from the provided source
func (agentPoolStatus *AgentPool_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == agentPoolStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(agentPoolStatus)
}

// ConvertStatusTo populates the provided destination from our AgentPool_Status
func (agentPoolStatus *AgentPool_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == agentPoolStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(agentPoolStatus)
}

//Storage version of v1alpha1api20210501.ManagedClustersAgentPools_Spec
type ManagedClustersAgentPools_Spec struct {
	AvailabilityZones []string `json:"availabilityZones,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName              string            `json:"azureName"`
	Count                  *int              `json:"count,omitempty"`
	EnableAutoScaling      *bool             `json:"enableAutoScaling,omitempty"`
	EnableEncryptionAtHost *bool             `json:"enableEncryptionAtHost,omitempty"`
	EnableFIPS             *bool             `json:"enableFIPS,omitempty"`
	EnableNodePublicIP     *bool             `json:"enableNodePublicIP,omitempty"`
	EnableUltraSSD         *bool             `json:"enableUltraSSD,omitempty"`
	GpuInstanceProfile     *string           `json:"gpuInstanceProfile,omitempty"`
	KubeletConfig          *KubeletConfig    `json:"kubeletConfig,omitempty"`
	KubeletDiskType        *string           `json:"kubeletDiskType,omitempty"`
	LinuxOSConfig          *LinuxOSConfig    `json:"linuxOSConfig,omitempty"`
	Location               *string           `json:"location,omitempty"`
	MaxCount               *int              `json:"maxCount,omitempty"`
	MaxPods                *int              `json:"maxPods,omitempty"`
	MinCount               *int              `json:"minCount,omitempty"`
	Mode                   *string           `json:"mode,omitempty"`
	NodeLabels             map[string]string `json:"nodeLabels,omitempty"`

	//NodePublicIPPrefixIDReference: This is of the form:
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
	NodePublicIPPrefixIDReference *genruntime.ResourceReference `armReference:"NodePublicIPPrefixID" json:"nodePublicIPPrefixIDReference,omitempty"`
	NodeTaints                    []string                      `json:"nodeTaints,omitempty"`
	OrchestratorVersion           *string                       `json:"orchestratorVersion,omitempty"`
	OriginalVersion               string                        `json:"originalVersion"`
	OsDiskSizeGB                  *int                          `json:"osDiskSizeGB,omitempty"`
	OsDiskType                    *string                       `json:"osDiskType,omitempty"`
	OsSKU                         *string                       `json:"osSKU,omitempty"`
	OsType                        *string                       `json:"osType,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"microsoft.containerservice.azure.com" json:"owner" kind:"ManagedCluster"`

	//PodSubnetIDReference: If omitted, pod IPs are statically assigned on the node
	//subnet (see vnetSubnetID for more details). This is of the form:
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	PodSubnetIDReference      *genruntime.ResourceReference `armReference:"PodSubnetID" json:"podSubnetIDReference,omitempty"`
	PropertyBag               genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	ProximityPlacementGroupID *string                       `json:"proximityPlacementGroupID,omitempty"`
	ScaleSetEvictionPolicy    *string                       `json:"scaleSetEvictionPolicy,omitempty"`
	ScaleSetPriority          *string                       `json:"scaleSetPriority,omitempty"`
	SpotMaxPrice              *float64                      `json:"spotMaxPrice,omitempty"`
	Tags                      map[string]string             `json:"tags,omitempty"`
	Type                      *string                       `json:"type,omitempty"`
	UpgradeSettings           *AgentPoolUpgradeSettings     `json:"upgradeSettings,omitempty"`
	VmSize                    *string                       `json:"vmSize,omitempty"`

	//VnetSubnetIDReference: If this is not specified, a VNET and subnet will be
	//generated and used. If no podSubnetID is specified, this applies to nodes and
	//pods, otherwise it applies to just nodes. This is of the form:
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	VnetSubnetIDReference *genruntime.ResourceReference `armReference:"VnetSubnetID" json:"vnetSubnetIDReference,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ManagedClustersAgentPools_Spec{}

// ConvertSpecFrom populates our ManagedClustersAgentPools_Spec from the provided source
func (managedClustersAgentPoolsSpec *ManagedClustersAgentPools_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == managedClustersAgentPoolsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(managedClustersAgentPoolsSpec)
}

// ConvertSpecTo populates the provided destination from our ManagedClustersAgentPools_Spec
func (managedClustersAgentPoolsSpec *ManagedClustersAgentPools_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == managedClustersAgentPoolsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(managedClustersAgentPoolsSpec)
}

//Storage version of v1alpha1api20210501.AgentPoolUpgradeSettings
//Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/AgentPoolUpgradeSettings
type AgentPoolUpgradeSettings struct {
	MaxSurge    *string                `json:"maxSurge,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210501.AgentPoolUpgradeSettings_Status
//Generated from:
type AgentPoolUpgradeSettings_Status struct {
	MaxSurge    *string                `json:"maxSurge,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210501.KubeletConfig
//Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/KubeletConfig
type KubeletConfig struct {
	AllowedUnsafeSysctls  []string               `json:"allowedUnsafeSysctls,omitempty"`
	ContainerLogMaxFiles  *int                   `json:"containerLogMaxFiles,omitempty"`
	ContainerLogMaxSizeMB *int                   `json:"containerLogMaxSizeMB,omitempty"`
	CpuCfsQuota           *bool                  `json:"cpuCfsQuota,omitempty"`
	CpuCfsQuotaPeriod     *string                `json:"cpuCfsQuotaPeriod,omitempty"`
	CpuManagerPolicy      *string                `json:"cpuManagerPolicy,omitempty"`
	FailSwapOn            *bool                  `json:"failSwapOn,omitempty"`
	ImageGcHighThreshold  *int                   `json:"imageGcHighThreshold,omitempty"`
	ImageGcLowThreshold   *int                   `json:"imageGcLowThreshold,omitempty"`
	PodMaxPids            *int                   `json:"podMaxPids,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TopologyManagerPolicy *string                `json:"topologyManagerPolicy,omitempty"`
}

//Storage version of v1alpha1api20210501.KubeletConfig_Status
//Generated from:
type KubeletConfig_Status struct {
	AllowedUnsafeSysctls  []string               `json:"allowedUnsafeSysctls,omitempty"`
	ContainerLogMaxFiles  *int                   `json:"containerLogMaxFiles,omitempty"`
	ContainerLogMaxSizeMB *int                   `json:"containerLogMaxSizeMB,omitempty"`
	CpuCfsQuota           *bool                  `json:"cpuCfsQuota,omitempty"`
	CpuCfsQuotaPeriod     *string                `json:"cpuCfsQuotaPeriod,omitempty"`
	CpuManagerPolicy      *string                `json:"cpuManagerPolicy,omitempty"`
	FailSwapOn            *bool                  `json:"failSwapOn,omitempty"`
	ImageGcHighThreshold  *int                   `json:"imageGcHighThreshold,omitempty"`
	ImageGcLowThreshold   *int                   `json:"imageGcLowThreshold,omitempty"`
	PodMaxPids            *int                   `json:"podMaxPids,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TopologyManagerPolicy *string                `json:"topologyManagerPolicy,omitempty"`
}

//Storage version of v1alpha1api20210501.LinuxOSConfig
//Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/LinuxOSConfig
type LinuxOSConfig struct {
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SwapFileSizeMB             *int                   `json:"swapFileSizeMB,omitempty"`
	Sysctls                    *SysctlConfig          `json:"sysctls,omitempty"`
	TransparentHugePageDefrag  *string                `json:"transparentHugePageDefrag,omitempty"`
	TransparentHugePageEnabled *string                `json:"transparentHugePageEnabled,omitempty"`
}

//Storage version of v1alpha1api20210501.LinuxOSConfig_Status
//Generated from:
type LinuxOSConfig_Status struct {
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SwapFileSizeMB             *int                   `json:"swapFileSizeMB,omitempty"`
	Sysctls                    *SysctlConfig_Status   `json:"sysctls,omitempty"`
	TransparentHugePageDefrag  *string                `json:"transparentHugePageDefrag,omitempty"`
	TransparentHugePageEnabled *string                `json:"transparentHugePageEnabled,omitempty"`
}

//Storage version of v1alpha1api20210501.SysctlConfig
//Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/SysctlConfig
type SysctlConfig struct {
	FsAioMaxNr                     *int                   `json:"fsAioMaxNr,omitempty"`
	FsFileMax                      *int                   `json:"fsFileMax,omitempty"`
	FsInotifyMaxUserWatches        *int                   `json:"fsInotifyMaxUserWatches,omitempty"`
	FsNrOpen                       *int                   `json:"fsNrOpen,omitempty"`
	KernelThreadsMax               *int                   `json:"kernelThreadsMax,omitempty"`
	NetCoreNetdevMaxBacklog        *int                   `json:"netCoreNetdevMaxBacklog,omitempty"`
	NetCoreOptmemMax               *int                   `json:"netCoreOptmemMax,omitempty"`
	NetCoreRmemDefault             *int                   `json:"netCoreRmemDefault,omitempty"`
	NetCoreRmemMax                 *int                   `json:"netCoreRmemMax,omitempty"`
	NetCoreSomaxconn               *int                   `json:"netCoreSomaxconn,omitempty"`
	NetCoreWmemDefault             *int                   `json:"netCoreWmemDefault,omitempty"`
	NetCoreWmemMax                 *int                   `json:"netCoreWmemMax,omitempty"`
	NetIpv4IpLocalPortRange        *string                `json:"netIpv4IpLocalPortRange,omitempty"`
	NetIpv4NeighDefaultGcThresh1   *int                   `json:"netIpv4NeighDefaultGcThresh1,omitempty"`
	NetIpv4NeighDefaultGcThresh2   *int                   `json:"netIpv4NeighDefaultGcThresh2,omitempty"`
	NetIpv4NeighDefaultGcThresh3   *int                   `json:"netIpv4NeighDefaultGcThresh3,omitempty"`
	NetIpv4TcpFinTimeout           *int                   `json:"netIpv4TcpFinTimeout,omitempty"`
	NetIpv4TcpKeepaliveProbes      *int                   `json:"netIpv4TcpKeepaliveProbes,omitempty"`
	NetIpv4TcpKeepaliveTime        *int                   `json:"netIpv4TcpKeepaliveTime,omitempty"`
	NetIpv4TcpMaxSynBacklog        *int                   `json:"netIpv4TcpMaxSynBacklog,omitempty"`
	NetIpv4TcpMaxTwBuckets         *int                   `json:"netIpv4TcpMaxTwBuckets,omitempty"`
	NetIpv4TcpTwReuse              *bool                  `json:"netIpv4TcpTwReuse,omitempty"`
	NetIpv4TcpkeepaliveIntvl       *int                   `json:"netIpv4TcpkeepaliveIntvl,omitempty"`
	NetNetfilterNfConntrackBuckets *int                   `json:"netNetfilterNfConntrackBuckets,omitempty"`
	NetNetfilterNfConntrackMax     *int                   `json:"netNetfilterNfConntrackMax,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VmMaxMapCount                  *int                   `json:"vmMaxMapCount,omitempty"`
	VmSwappiness                   *int                   `json:"vmSwappiness,omitempty"`
	VmVfsCachePressure             *int                   `json:"vmVfsCachePressure,omitempty"`
}

//Storage version of v1alpha1api20210501.SysctlConfig_Status
//Generated from:
type SysctlConfig_Status struct {
	FsAioMaxNr                     *int                   `json:"fsAioMaxNr,omitempty"`
	FsFileMax                      *int                   `json:"fsFileMax,omitempty"`
	FsInotifyMaxUserWatches        *int                   `json:"fsInotifyMaxUserWatches,omitempty"`
	FsNrOpen                       *int                   `json:"fsNrOpen,omitempty"`
	KernelThreadsMax               *int                   `json:"kernelThreadsMax,omitempty"`
	NetCoreNetdevMaxBacklog        *int                   `json:"netCoreNetdevMaxBacklog,omitempty"`
	NetCoreOptmemMax               *int                   `json:"netCoreOptmemMax,omitempty"`
	NetCoreRmemDefault             *int                   `json:"netCoreRmemDefault,omitempty"`
	NetCoreRmemMax                 *int                   `json:"netCoreRmemMax,omitempty"`
	NetCoreSomaxconn               *int                   `json:"netCoreSomaxconn,omitempty"`
	NetCoreWmemDefault             *int                   `json:"netCoreWmemDefault,omitempty"`
	NetCoreWmemMax                 *int                   `json:"netCoreWmemMax,omitempty"`
	NetIpv4IpLocalPortRange        *string                `json:"netIpv4IpLocalPortRange,omitempty"`
	NetIpv4NeighDefaultGcThresh1   *int                   `json:"netIpv4NeighDefaultGcThresh1,omitempty"`
	NetIpv4NeighDefaultGcThresh2   *int                   `json:"netIpv4NeighDefaultGcThresh2,omitempty"`
	NetIpv4NeighDefaultGcThresh3   *int                   `json:"netIpv4NeighDefaultGcThresh3,omitempty"`
	NetIpv4TcpFinTimeout           *int                   `json:"netIpv4TcpFinTimeout,omitempty"`
	NetIpv4TcpKeepaliveProbes      *int                   `json:"netIpv4TcpKeepaliveProbes,omitempty"`
	NetIpv4TcpKeepaliveTime        *int                   `json:"netIpv4TcpKeepaliveTime,omitempty"`
	NetIpv4TcpMaxSynBacklog        *int                   `json:"netIpv4TcpMaxSynBacklog,omitempty"`
	NetIpv4TcpMaxTwBuckets         *int                   `json:"netIpv4TcpMaxTwBuckets,omitempty"`
	NetIpv4TcpTwReuse              *bool                  `json:"netIpv4TcpTwReuse,omitempty"`
	NetIpv4TcpkeepaliveIntvl       *int                   `json:"netIpv4TcpkeepaliveIntvl,omitempty"`
	NetNetfilterNfConntrackBuckets *int                   `json:"netNetfilterNfConntrackBuckets,omitempty"`
	NetNetfilterNfConntrackMax     *int                   `json:"netNetfilterNfConntrackMax,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VmMaxMapCount                  *int                   `json:"vmMaxMapCount,omitempty"`
	VmSwappiness                   *int                   `json:"vmSwappiness,omitempty"`
	VmVfsCachePressure             *int                   `json:"vmVfsCachePressure,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ManagedClustersAgentPool{}, &ManagedClustersAgentPoolList{})
}
