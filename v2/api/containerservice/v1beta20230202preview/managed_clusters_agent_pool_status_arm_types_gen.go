// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20230202preview

type ManagedClusters_AgentPool_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of an agent pool.
	Properties *ManagedClusterAgentPoolProfileProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

// Properties for the container service agent pool profile.
type ManagedClusterAgentPoolProfileProperties_STATUS_ARM struct {
	// AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
	// property is 'VirtualMachineScaleSets'.
	AvailabilityZones []string `json:"availabilityZones,omitempty"`

	// CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.
	CapacityReservationGroupID *string `json:"capacityReservationGroupID,omitempty"`

	// Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
	// for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
	Count *int `json:"count,omitempty"`

	// CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
	// a snapshot.
	CreationData *CreationData_STATUS_ARM `json:"creationData,omitempty"`

	// CurrentOrchestratorVersion: If orchestratorVersion was a fully specified version <major.minor.patch>, this field will be
	// exactly equal to it. If orchestratorVersion was <major.minor>, this field will contain the full <major.minor.patch>
	// version being used.
	CurrentOrchestratorVersion *string `json:"currentOrchestratorVersion,omitempty"`

	// EnableAutoScaling: Whether to enable auto-scaler
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty"`

	// EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
	// daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
	// certificates into node trust stores. Defaults to false.
	EnableCustomCATrust *bool `json:"enableCustomCATrust,omitempty"`

	// EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
	// see: https://docs.microsoft.com/azure/aks/enable-host-encryption
	EnableEncryptionAtHost *bool `json:"enableEncryptionAtHost,omitempty"`

	// EnableFIPS: See [Add a FIPS-enabled node
	// pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more
	// details.
	EnableFIPS *bool `json:"enableFIPS,omitempty"`

	// EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
	// A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
	// to minimize hops. For more information see [assigning a public IP per
	// node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The
	// default is false.
	EnableNodePublicIP *bool `json:"enableNodePublicIP,omitempty"`

	// EnableUltraSSD: Whether to enable UltraSSD
	EnableUltraSSD *bool `json:"enableUltraSSD,omitempty"`

	// GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
	GpuInstanceProfile *GPUInstanceProfile_STATUS `json:"gpuInstanceProfile,omitempty"`

	// HostGroupID: This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}.
	// For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).
	HostGroupID *string `json:"hostGroupID,omitempty"`

	// KubeletConfig: The Kubelet configuration on the agent pool nodes.
	KubeletConfig *KubeletConfig_STATUS_ARM `json:"kubeletConfig,omitempty"`

	// KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
	// storage.
	KubeletDiskType *KubeletDiskType_STATUS `json:"kubeletDiskType,omitempty"`

	// LinuxOSConfig: The OS configuration of Linux agent nodes.
	LinuxOSConfig *LinuxOSConfig_STATUS_ARM `json:"linuxOSConfig,omitempty"`

	// MaxCount: The maximum number of nodes for auto-scaling
	MaxCount *int `json:"maxCount,omitempty"`

	// MaxPods: The maximum number of pods that can run on a node.
	MaxPods *int `json:"maxPods,omitempty"`

	// MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
	// the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
	// will be printed raw and not be executed as a script).
	MessageOfTheDay *string `json:"messageOfTheDay,omitempty"`

	// MinCount: The minimum number of nodes for auto-scaling
	MinCount *int `json:"minCount,omitempty"`

	// Mode: A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool
	// restrictions  and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
	Mode *AgentPoolMode_STATUS `json:"mode,omitempty"`

	// NetworkProfile: Network-related settings of an agent pool.
	NetworkProfile *AgentPoolNetworkProfile_STATUS_ARM `json:"networkProfile,omitempty"`

	// NodeImageVersion: The version of node image
	NodeImageVersion *string `json:"nodeImageVersion,omitempty"`

	// NodeLabels: The node labels to be persisted across all nodes in agent pool.
	NodeLabels map[string]string `json:"nodeLabels,omitempty"`

	// NodePublicIPPrefixID: This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
	NodePublicIPPrefixID *string `json:"nodePublicIPPrefixID,omitempty"`

	// NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `json:"nodeTaints,omitempty"`

	// OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
	// specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
	// <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
	// best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
	// must have the same major version as the control plane. The node pool minor version must be within two minor versions of
	// the control plane version. The node pool version cannot be greater than the control plane version. For more information
	// see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
	OrchestratorVersion *string `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB        *int    `json:"osDiskSizeGB,omitempty"`

	// OsDiskType: The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested
	// OSDiskSizeGB. Otherwise,  defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral
	// OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
	OsDiskType *OSDiskType_STATUS `json:"osDiskType,omitempty"`

	// OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
	// Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
	// deprecated.
	OsSKU *OSSKU_STATUS `json:"osSKU,omitempty"`

	// OsType: The operating system type. The default is Linux.
	OsType *OSType_STATUS `json:"osType,omitempty"`

	// PodSubnetID: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is
	// of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	PodSubnetID *string `json:"podSubnetID,omitempty"`

	// PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
	// field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
	// be stopped if it is Running and provisioning state is Succeeded
	PowerState *PowerState_STATUS_ARM `json:"powerState,omitempty"`

	// ProvisioningState: The current deployment or provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// ProximityPlacementGroupID: The ID for Proximity Placement Group.
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupID,omitempty"`

	// ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
	ScaleDownMode *ScaleDownMode_STATUS `json:"scaleDownMode,omitempty"`

	// ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is
	// 'Delete'.
	ScaleSetEvictionPolicy *ScaleSetEvictionPolicy_STATUS `json:"scaleSetEvictionPolicy,omitempty"`

	// ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
	ScaleSetPriority *ScaleSetPriority_STATUS `json:"scaleSetPriority,omitempty"`

	// SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
	// on-demand price. For more details on spot pricing, see [spot VMs
	// pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
	SpotMaxPrice *float64 `json:"spotMaxPrice,omitempty"`

	// Tags: The tags to be persisted on the agent pool virtual machine scale set.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of Agent Pool.
	Type *AgentPoolType_STATUS `json:"type,omitempty"`

	// UpgradeSettings: Settings for upgrading the agentpool
	UpgradeSettings *AgentPoolUpgradeSettings_STATUS_ARM `json:"upgradeSettings,omitempty"`

	// VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
	// might fail to run correctly. For more details on restricted VM sizes, see:
	// https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VmSize *string `json:"vmSize,omitempty"`

	// VnetSubnetID: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified,
	// this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	VnetSubnetID *string `json:"vnetSubnetID,omitempty"`

	// WindowsProfile: The Windows agent pool's specific profile.
	WindowsProfile *AgentPoolWindowsProfile_STATUS_ARM `json:"windowsProfile,omitempty"`

	// WorkloadRuntime: Determines the type of workload a node can run.
	WorkloadRuntime *WorkloadRuntime_STATUS `json:"workloadRuntime,omitempty"`
}

// Network settings of an agent pool.
type AgentPoolNetworkProfile_STATUS_ARM struct {
	// AllowedHostPorts: The port ranges that are allowed to access. The specified ranges are allowed to overlap.
	AllowedHostPorts []PortRange_STATUS_ARM `json:"allowedHostPorts,omitempty"`

	// ApplicationSecurityGroups: The IDs of the application security groups which agent pool will associate when created.
	ApplicationSecurityGroups []string `json:"applicationSecurityGroups,omitempty"`

	// NodePublicIPTags: IPTags of instance-level public IPs.
	NodePublicIPTags []IPTag_STATUS_ARM `json:"nodePublicIPTags,omitempty"`
}

// Settings for upgrading an agentpool
type AgentPoolUpgradeSettings_STATUS_ARM struct {
	// MaxSurge: This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it
	// is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded
	// up. If not specified, the default is 1. For more information, including best practices, see:
	// https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade
	MaxSurge *string `json:"maxSurge,omitempty"`
}

// The Windows agent pool's specific profile.
type AgentPoolWindowsProfile_STATUS_ARM struct {
	// DisableOutboundNat: The default value is false. Outbound NAT can only be disabled if the cluster outboundType is NAT
	// Gateway and the Windows agent pool does not have node public IP enabled.
	DisableOutboundNat *bool `json:"disableOutboundNat,omitempty"`
}

// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type KubeletConfig_STATUS_ARM struct {
	// AllowedUnsafeSysctls: Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in `*`).
	AllowedUnsafeSysctls []string `json:"allowedUnsafeSysctls,omitempty"`

	// ContainerLogMaxFiles: The maximum number of container log files that can be present for a container. The number must be
	// ≥ 2.
	ContainerLogMaxFiles *int `json:"containerLogMaxFiles,omitempty"`

	// ContainerLogMaxSizeMB: The maximum size (e.g. 10Mi) of container log file before it is rotated.
	ContainerLogMaxSizeMB *int `json:"containerLogMaxSizeMB,omitempty"`

	// CpuCfsQuota: The default is true.
	CpuCfsQuota *bool `json:"cpuCfsQuota,omitempty"`

	// CpuCfsQuotaPeriod: The default is '100ms.' Valid values are a sequence of decimal numbers with an optional fraction and
	// a unit suffix. For example: '300ms', '2h45m'. Supported units are 'ns', 'us', 'ms', 's', 'm', and 'h'.
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`

	// CpuManagerPolicy: The default is 'none'. See [Kubernetes CPU management
	// policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies) for more
	// information. Allowed values are 'none' and 'static'.
	CpuManagerPolicy *string `json:"cpuManagerPolicy,omitempty"`

	// FailSwapOn: If set to true it will make the Kubelet fail to start if swap is enabled on the node.
	FailSwapOn *bool `json:"failSwapOn,omitempty"`

	// ImageGcHighThreshold: To disable image garbage collection, set to 100. The default is 85%
	ImageGcHighThreshold *int `json:"imageGcHighThreshold,omitempty"`

	// ImageGcLowThreshold: This cannot be set higher than imageGcHighThreshold. The default is 80%
	ImageGcLowThreshold *int `json:"imageGcLowThreshold,omitempty"`

	// PodMaxPids: The maximum number of processes per pod.
	PodMaxPids *int `json:"podMaxPids,omitempty"`

	// TopologyManagerPolicy: For more information see [Kubernetes Topology
	// Manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager). The default is 'none'. Allowed values
	// are 'none', 'best-effort', 'restricted', and 'single-numa-node'.
	TopologyManagerPolicy *string `json:"topologyManagerPolicy,omitempty"`
}

// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type LinuxOSConfig_STATUS_ARM struct {
	// SwapFileSizeMB: The size in MB of a swap file that will be created on each node.
	SwapFileSizeMB *int `json:"swapFileSizeMB,omitempty"`

	// Sysctls: Sysctl settings for Linux agent nodes.
	Sysctls *SysctlConfig_STATUS_ARM `json:"sysctls,omitempty"`

	// TransparentHugePageDefrag: Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is
	// 'madvise'. For more information see [Transparent
	// Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).
	TransparentHugePageDefrag *string `json:"transparentHugePageDefrag,omitempty"`

	// TransparentHugePageEnabled: Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more
	// information see [Transparent
	// Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).
	TransparentHugePageEnabled *string `json:"transparentHugePageEnabled,omitempty"`
}

// Contains the IPTag associated with the object.
type IPTag_STATUS_ARM struct {
	// IpTagType: The IP tag type. Example: RoutingPreference.
	IpTagType *string `json:"ipTagType,omitempty"`

	// Tag: The value of the IP tag associated with the public IP. Example: Internet.
	Tag *string `json:"tag,omitempty"`
}

// The port range.
type PortRange_STATUS_ARM struct {
	// PortEnd: The maximum port that is included in the range. It should be ranged from 1 to 65535, and be greater than or
	// equal to portStart.
	PortEnd *int `json:"portEnd,omitempty"`

	// PortStart: The minimum port that is included in the range. It should be ranged from 1 to 65535, and be less than or
	// equal to portEnd.
	PortStart *int `json:"portStart,omitempty"`

	// Protocol: The network protocol of the port.
	Protocol *PortRange_Protocol_STATUS `json:"protocol,omitempty"`
}

// Sysctl settings for Linux agent nodes.
type SysctlConfig_STATUS_ARM struct {
	// FsAioMaxNr: Sysctl setting fs.aio-max-nr.
	FsAioMaxNr *int `json:"fsAioMaxNr,omitempty"`

	// FsFileMax: Sysctl setting fs.file-max.
	FsFileMax *int `json:"fsFileMax,omitempty"`

	// FsInotifyMaxUserWatches: Sysctl setting fs.inotify.max_user_watches.
	FsInotifyMaxUserWatches *int `json:"fsInotifyMaxUserWatches,omitempty"`

	// FsNrOpen: Sysctl setting fs.nr_open.
	FsNrOpen *int `json:"fsNrOpen,omitempty"`

	// KernelThreadsMax: Sysctl setting kernel.threads-max.
	KernelThreadsMax *int `json:"kernelThreadsMax,omitempty"`

	// NetCoreNetdevMaxBacklog: Sysctl setting net.core.netdev_max_backlog.
	NetCoreNetdevMaxBacklog *int `json:"netCoreNetdevMaxBacklog,omitempty"`

	// NetCoreOptmemMax: Sysctl setting net.core.optmem_max.
	NetCoreOptmemMax *int `json:"netCoreOptmemMax,omitempty"`

	// NetCoreRmemDefault: Sysctl setting net.core.rmem_default.
	NetCoreRmemDefault *int `json:"netCoreRmemDefault,omitempty"`

	// NetCoreRmemMax: Sysctl setting net.core.rmem_max.
	NetCoreRmemMax *int `json:"netCoreRmemMax,omitempty"`

	// NetCoreSomaxconn: Sysctl setting net.core.somaxconn.
	NetCoreSomaxconn *int `json:"netCoreSomaxconn,omitempty"`

	// NetCoreWmemDefault: Sysctl setting net.core.wmem_default.
	NetCoreWmemDefault *int `json:"netCoreWmemDefault,omitempty"`

	// NetCoreWmemMax: Sysctl setting net.core.wmem_max.
	NetCoreWmemMax *int `json:"netCoreWmemMax,omitempty"`

	// NetIpv4IpLocalPortRange: Sysctl setting net.ipv4.ip_local_port_range.
	NetIpv4IpLocalPortRange *string `json:"netIpv4IpLocalPortRange,omitempty"`

	// NetIpv4NeighDefaultGcThresh1: Sysctl setting net.ipv4.neigh.default.gc_thresh1.
	NetIpv4NeighDefaultGcThresh1 *int `json:"netIpv4NeighDefaultGcThresh1,omitempty"`

	// NetIpv4NeighDefaultGcThresh2: Sysctl setting net.ipv4.neigh.default.gc_thresh2.
	NetIpv4NeighDefaultGcThresh2 *int `json:"netIpv4NeighDefaultGcThresh2,omitempty"`

	// NetIpv4NeighDefaultGcThresh3: Sysctl setting net.ipv4.neigh.default.gc_thresh3.
	NetIpv4NeighDefaultGcThresh3 *int `json:"netIpv4NeighDefaultGcThresh3,omitempty"`

	// NetIpv4TcpFinTimeout: Sysctl setting net.ipv4.tcp_fin_timeout.
	NetIpv4TcpFinTimeout *int `json:"netIpv4TcpFinTimeout,omitempty"`

	// NetIpv4TcpKeepaliveProbes: Sysctl setting net.ipv4.tcp_keepalive_probes.
	NetIpv4TcpKeepaliveProbes *int `json:"netIpv4TcpKeepaliveProbes,omitempty"`

	// NetIpv4TcpKeepaliveTime: Sysctl setting net.ipv4.tcp_keepalive_time.
	NetIpv4TcpKeepaliveTime *int `json:"netIpv4TcpKeepaliveTime,omitempty"`

	// NetIpv4TcpMaxSynBacklog: Sysctl setting net.ipv4.tcp_max_syn_backlog.
	NetIpv4TcpMaxSynBacklog *int `json:"netIpv4TcpMaxSynBacklog,omitempty"`

	// NetIpv4TcpMaxTwBuckets: Sysctl setting net.ipv4.tcp_max_tw_buckets.
	NetIpv4TcpMaxTwBuckets *int `json:"netIpv4TcpMaxTwBuckets,omitempty"`

	// NetIpv4TcpTwReuse: Sysctl setting net.ipv4.tcp_tw_reuse.
	NetIpv4TcpTwReuse *bool `json:"netIpv4TcpTwReuse,omitempty"`

	// NetIpv4TcpkeepaliveIntvl: Sysctl setting net.ipv4.tcp_keepalive_intvl.
	NetIpv4TcpkeepaliveIntvl *int `json:"netIpv4TcpkeepaliveIntvl,omitempty"`

	// NetNetfilterNfConntrackBuckets: Sysctl setting net.netfilter.nf_conntrack_buckets.
	NetNetfilterNfConntrackBuckets *int `json:"netNetfilterNfConntrackBuckets,omitempty"`

	// NetNetfilterNfConntrackMax: Sysctl setting net.netfilter.nf_conntrack_max.
	NetNetfilterNfConntrackMax *int `json:"netNetfilterNfConntrackMax,omitempty"`

	// VmMaxMapCount: Sysctl setting vm.max_map_count.
	VmMaxMapCount *int `json:"vmMaxMapCount,omitempty"`

	// VmSwappiness: Sysctl setting vm.swappiness.
	VmSwappiness *int `json:"vmSwappiness,omitempty"`

	// VmVfsCachePressure: Sysctl setting vm.vfs_cache_pressure.
	VmVfsCachePressure *int `json:"vmVfsCachePressure,omitempty"`
}
