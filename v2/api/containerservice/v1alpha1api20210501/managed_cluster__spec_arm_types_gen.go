// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type ManagedCluster_SpecARM struct {
	AzureName string `json:"azureName"`

	//ExtendedLocation: The extended location of the Virtual Machine.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	//Identity: The identity of the managed cluster, if configured.
	Identity *ManagedClusterIdentityARM `json:"identity,omitempty"`

	//Location: Resource location
	Location string `json:"location"`
	Name     string `json:"name"`

	//Properties: Properties of a managed cluster.
	Properties *ManagedClusterPropertiesARM `json:"properties,omitempty"`

	//Sku: The managed cluster SKU.
	Sku *ManagedClusterSKUARM `json:"sku,omitempty"`

	//Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ManagedCluster_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (cluster ManagedCluster_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (cluster ManagedCluster_SpecARM) GetName() string {
	return cluster.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters"
func (cluster ManagedCluster_SpecARM) GetType() string {
	return "Microsoft.ContainerService/managedClusters"
}

type ExtendedLocationARM struct {
	//Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	//Type: The type of the extended location.
	Type *ExtendedLocationType `json:"type,omitempty"`
}

type ManagedClusterIdentityARM struct {
	//Type: For more information see [use managed identities in
	//AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).
	Type *ManagedClusterIdentityType `json:"type,omitempty"`
}

type ManagedClusterPropertiesARM struct {
	//AadProfile: The Azure Active Directory configuration.
	AadProfile *ManagedClusterAADProfileARM `json:"aadProfile,omitempty"`

	//AddonProfiles: The profile of managed cluster add-on.
	AddonProfiles *v1.JSON `json:"addonProfiles,omitempty"`

	//AgentPoolProfiles: The agent pool properties.
	AgentPoolProfiles []ManagedClusterAgentPoolProfileARM `json:"agentPoolProfiles,omitempty"`

	//ApiServerAccessProfile: The access profile for managed cluster API server.
	ApiServerAccessProfile *ManagedClusterAPIServerAccessProfileARM `json:"apiServerAccessProfile,omitempty"`

	//AutoScalerProfile: Parameters to be applied to the cluster-autoscaler when
	//enabled
	AutoScalerProfile *ManagedClusterPropertiesAutoScalerProfileARM `json:"autoScalerProfile,omitempty"`

	//AutoUpgradeProfile: The auto upgrade configuration.
	AutoUpgradeProfile *ManagedClusterAutoUpgradeProfileARM `json:"autoUpgradeProfile,omitempty"`

	//DisableLocalAccounts: If set to true, getting static credentials will be
	//disabled for this cluster. This must only be used on Managed Clusters that are
	//AAD enabled. For more details see [disable local
	//accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).
	DisableLocalAccounts *bool   `json:"disableLocalAccounts,omitempty"`
	DiskEncryptionSetID  *string `json:"diskEncryptionSetID,omitempty"`

	//DnsPrefix: This cannot be updated once the Managed Cluster has been created.
	DnsPrefix *string `json:"dnsPrefix,omitempty"`

	//EnablePodSecurityPolicy: (DEPRECATING) Whether to enable Kubernetes pod security
	//policy (preview). This feature is set for removal on October 15th, 2020. Learn
	//more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy *bool `json:"enablePodSecurityPolicy,omitempty"`

	//EnableRBAC: Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC *bool `json:"enableRBAC,omitempty"`

	//FqdnSubdomain: This cannot be updated once the Managed Cluster has been created.
	FqdnSubdomain *string `json:"fqdnSubdomain,omitempty"`

	//HttpProxyConfig: Configurations for provisioning the cluster with HTTP proxy
	//servers.
	HttpProxyConfig *ManagedClusterHTTPProxyConfigARM `json:"httpProxyConfig,omitempty"`

	//IdentityProfile: Identities associated with the cluster.
	IdentityProfile *v1.JSON `json:"identityProfile,omitempty"`

	//KubernetesVersion: When you upgrade a supported AKS cluster, Kubernetes minor
	//versions cannot be skipped. All upgrades must be performed sequentially by major
	//version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x ->
	//1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an
	//AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more
	//details.
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`

	//LinuxProfile: The profile for Linux VMs in the Managed Cluster.
	LinuxProfile *ContainerServiceLinuxProfileARM `json:"linuxProfile,omitempty"`

	//NetworkProfile: The network configuration profile.
	NetworkProfile *ContainerServiceNetworkProfileARM `json:"networkProfile,omitempty"`

	//NodeResourceGroup: The name of the resource group containing agent pool nodes.
	NodeResourceGroup *string `json:"nodeResourceGroup,omitempty"`

	//PodIdentityProfile: See [use AAD pod
	//identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for
	//more details on AAD pod identity integration.
	PodIdentityProfile *ManagedClusterPodIdentityProfileARM `json:"podIdentityProfile,omitempty"`

	//PrivateLinkResources: Private link resources associated with the cluster.
	PrivateLinkResources []genruntime.ResourceReference `json:"privateLinkResources,omitempty"`

	//ServicePrincipalProfile: Information about a service principal identity for the
	//cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfileARM `json:"servicePrincipalProfile,omitempty"`

	//WindowsProfile: The profile for Windows VMs in the Managed Cluster.
	WindowsProfile *ManagedClusterWindowsProfileARM `json:"windowsProfile,omitempty"`
}

type ManagedClusterSKUARM struct {
	//Name: The name of a managed cluster SKU.
	Name *ManagedClusterSKUName `json:"name,omitempty"`

	//Tier: If not specified, the default is 'Free'. See [uptime
	//SLA](https://docs.microsoft.com/azure/aks/uptime-sla) for more details.
	Tier *ManagedClusterSKUTier `json:"tier,omitempty"`
}

type ContainerServiceLinuxProfileARM struct {
	//AdminUsername: The administrator username to use for Linux VMs.
	AdminUsername string `json:"adminUsername"`

	//Ssh: The SSH configuration for Linux-based VMs running on Azure.
	Ssh ContainerServiceSshConfigurationARM `json:"ssh"`
}

type ContainerServiceNetworkProfileARM struct {
	//DnsServiceIP: An IP address assigned to the Kubernetes DNS service. It must be
	//within the Kubernetes service address range specified in serviceCidr.
	DnsServiceIP *string `json:"dnsServiceIP,omitempty"`

	//DockerBridgeCidr: A CIDR notation IP range assigned to the Docker bridge
	//network. It must not overlap with any Subnet IP ranges or the Kubernetes service
	//address range.
	DockerBridgeCidr *string `json:"dockerBridgeCidr,omitempty"`

	//LoadBalancerProfile: Profile of the cluster load balancer.
	LoadBalancerProfile *ManagedClusterLoadBalancerProfileARM `json:"loadBalancerProfile,omitempty"`

	//LoadBalancerSku: The default is 'standard'. See [Azure Load Balancer
	//SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information
	//about the differences between load balancer SKUs.
	LoadBalancerSku *ContainerServiceNetworkProfileLoadBalancerSku `json:"loadBalancerSku,omitempty"`

	//NetworkMode: This cannot be specified if networkPlugin is anything other than
	//'azure'.
	NetworkMode *ContainerServiceNetworkProfileNetworkMode `json:"networkMode,omitempty"`

	//NetworkPlugin: Network plugin used for building the Kubernetes network.
	NetworkPlugin *ContainerServiceNetworkProfileNetworkPlugin `json:"networkPlugin,omitempty"`

	//NetworkPolicy: Network policy used for building the Kubernetes network.
	NetworkPolicy *ContainerServiceNetworkProfileNetworkPolicy `json:"networkPolicy,omitempty"`

	//OutboundType: This can only be set at cluster creation time and cannot be
	//changed later. For more information see [egress outbound
	//type](https://docs.microsoft.com/azure/aks/egress-outboundtype).
	OutboundType *ContainerServiceNetworkProfileOutboundType `json:"outboundType,omitempty"`

	//PodCidr: A CIDR notation IP range from which to assign pod IPs when kubenet is
	//used.
	PodCidr *string `json:"podCidr,omitempty"`

	//ServiceCidr: A CIDR notation IP range from which to assign service cluster IPs.
	//It must not overlap with any Subnet IP ranges.
	ServiceCidr *string `json:"serviceCidr,omitempty"`
}

// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationTypeEdgeZone = ExtendedLocationType("EdgeZone")

type ManagedClusterAADProfileARM struct {
	//AdminGroupObjectIDs: The list of AAD group object IDs that will have admin role
	//of the cluster.
	AdminGroupObjectIDs []string `json:"adminGroupObjectIDs,omitempty"`

	//ClientAppID: The client AAD application ID.
	ClientAppID *string `json:"clientAppID,omitempty"`

	//EnableAzureRBAC: Whether to enable Azure RBAC for Kubernetes authorization.
	EnableAzureRBAC *bool `json:"enableAzureRBAC,omitempty"`

	//Managed: Whether to enable managed AAD.
	Managed *bool `json:"managed,omitempty"`

	//ServerAppID: The server AAD application ID.
	ServerAppID *string `json:"serverAppID,omitempty"`

	//ServerAppSecret: The server AAD application secret.
	ServerAppSecret *string `json:"serverAppSecret,omitempty"`

	//TenantID: The AAD tenant ID to use for authentication. If not specified, will
	//use the tenant of the deployment subscription.
	TenantID *string `json:"tenantID,omitempty"`
}

type ManagedClusterAPIServerAccessProfileARM struct {
	//AuthorizedIPRanges: IP ranges are specified in CIDR format, e.g.
	//137.117.106.88/29. This feature is not compatible with clusters that use Public
	//IP Per Node, or clusters that are using a Basic Load Balancer. For more
	//information see [API server authorized IP
	//ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges).
	AuthorizedIPRanges []string `json:"authorizedIPRanges,omitempty"`

	//EnablePrivateCluster: For more details, see [Creating a private AKS
	//cluster](https://docs.microsoft.com/azure/aks/private-clusters).
	EnablePrivateCluster *bool `json:"enablePrivateCluster,omitempty"`

	//EnablePrivateClusterPublicFQDN: Whether to create additional public FQDN for
	//private cluster or not.
	EnablePrivateClusterPublicFQDN *bool `json:"enablePrivateClusterPublicFQDN,omitempty"`

	//PrivateDNSZone: The default is System. For more details see [configure private
	//DNS
	//zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone).
	//Allowed values are 'system' and 'none'.
	PrivateDNSZone *string `json:"privateDNSZone,omitempty"`
}

type ManagedClusterAgentPoolProfileARM struct {
	//AvailabilityZones: The list of Availability zones to use for nodes. This can
	//only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.
	AvailabilityZones []string `json:"availabilityZones,omitempty"`

	//Count: Number of agents (VMs) to host docker containers. Allowed values must be
	//in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to
	//1000 (inclusive) for system pools. The default value is 1.
	Count *int `json:"count,omitempty"`

	//EnableAutoScaling: Whether to enable auto-scaler
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty"`

	//EnableEncryptionAtHost: This is only supported on certain VM sizes and in
	//certain Azure regions. For more information, see:
	//https://docs.microsoft.com/azure/aks/enable-host-encryption
	EnableEncryptionAtHost *bool `json:"enableEncryptionAtHost,omitempty"`

	//EnableFIPS: See [Add a FIPS-enabled node
	//pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview)
	//for more details.
	EnableFIPS *bool `json:"enableFIPS,omitempty"`

	//EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive
	//their own dedicated public IP addresses. A common scenario is for gaming
	//workloads, where a console needs to make a direct connection to a cloud virtual
	//machine to minimize hops. For more information see [assigning a public IP per
	//node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools).
	//The default is false.
	EnableNodePublicIP *bool `json:"enableNodePublicIP,omitempty"`

	//EnableUltraSSD: Whether to enable UltraSSD
	EnableUltraSSD *bool `json:"enableUltraSSD,omitempty"`

	//GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance
	//profile for supported GPU VM SKU.
	GpuInstanceProfile *GPUInstanceProfile `json:"gpuInstanceProfile,omitempty"`

	//KubeletConfig: The Kubelet configuration on the agent pool nodes.
	KubeletConfig   *KubeletConfigARM `json:"kubeletConfig,omitempty"`
	KubeletDiskType *KubeletDiskType  `json:"kubeletDiskType,omitempty"`

	//LinuxOSConfig: The OS configuration of Linux agent nodes.
	LinuxOSConfig *LinuxOSConfigARM `json:"linuxOSConfig,omitempty"`

	//MaxCount: The maximum number of nodes for auto-scaling
	MaxCount *int `json:"maxCount,omitempty"`

	//MaxPods: The maximum number of pods that can run on a node.
	MaxPods *int `json:"maxPods,omitempty"`

	//MinCount: The minimum number of nodes for auto-scaling
	MinCount *int           `json:"minCount,omitempty"`
	Mode     *AgentPoolMode `json:"mode,omitempty"`

	//Name: Windows agent pool names must be 6 characters or less.
	Name *string `json:"name,omitempty"`

	//NodeLabels: The node labels to be persisted across all nodes in agent pool.
	NodeLabels           map[string]string `json:"nodeLabels,omitempty"`
	NodePublicIPPrefixID *string           `json:"nodePublicIPPrefixID,omitempty"`

	//NodeTaints: The taints added to new nodes during node pool create and scale. For
	//example, key=value:NoSchedule.
	NodeTaints []string `json:"nodeTaints,omitempty"`

	//OrchestratorVersion: As a best practice, you should upgrade all node pools in an
	//AKS cluster to the same Kubernetes version. The node pool version must have the
	//same major version as the control plane. The node pool minor version must be
	//within two minor versions of the control plane version. The node pool version
	//cannot be greater than the control plane version. For more information see
	//[upgrading a node
	//pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
	OrchestratorVersion *string                 `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB        *ContainerServiceOSDisk `json:"osDiskSizeGB,omitempty"`
	OsDiskType          *OSDiskType             `json:"osDiskType,omitempty"`
	OsSKU               *OSSKU                  `json:"osSKU,omitempty"`
	OsType              *OSType                 `json:"osType,omitempty"`
	PodSubnetID         *string                 `json:"podSubnetID,omitempty"`

	//ProximityPlacementGroupID: The ID for Proximity Placement Group.
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupID,omitempty"`

	//ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is
	//'Spot'. If not specified, the default is 'Delete'.
	ScaleSetEvictionPolicy *ScaleSetEvictionPolicy `json:"scaleSetEvictionPolicy,omitempty"`

	//ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the
	//default is 'Regular'.
	ScaleSetPriority *ScaleSetPriority `json:"scaleSetPriority,omitempty"`

	//SpotMaxPrice: Possible values are any decimal value greater than zero or -1
	//which indicates the willingness to pay any on-demand price. For more details on
	//spot pricing, see [spot VMs
	//pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
	SpotMaxPrice *float64 `json:"spotMaxPrice,omitempty"`

	//Tags: The tags to be persisted on the agent pool virtual machine scale set.
	Tags map[string]string `json:"tags,omitempty"`
	Type *AgentPoolType    `json:"type,omitempty"`

	//UpgradeSettings: Settings for upgrading the agentpool
	UpgradeSettings *AgentPoolUpgradeSettingsARM `json:"upgradeSettings,omitempty"`

	//VmSize: VM size availability varies by region. If a node contains insufficient
	//compute resources (memory, cpu, etc) pods might fail to run correctly. For more
	//details on restricted VM sizes, see:
	//https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VmSize       *string `json:"vmSize,omitempty"`
	VnetSubnetID *string `json:"vnetSubnetID,omitempty"`
}

type ManagedClusterAutoUpgradeProfileARM struct {
	//UpgradeChannel: For more information see [setting the AKS cluster auto-upgrade
	//channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel).
	UpgradeChannel *ManagedClusterAutoUpgradeProfileUpgradeChannel `json:"upgradeChannel,omitempty"`
}

type ManagedClusterHTTPProxyConfigARM struct {
	//HttpProxy: The HTTP proxy server endpoint to use.
	HttpProxy *string `json:"httpProxy,omitempty"`

	//HttpsProxy: The HTTPS proxy server endpoint to use.
	HttpsProxy *string `json:"httpsProxy,omitempty"`

	//NoProxy: The endpoints that should not go through proxy.
	NoProxy []string `json:"noProxy,omitempty"`

	//TrustedCa: Alternative CA cert to use for connecting to proxy servers.
	TrustedCa *string `json:"trustedCa,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type ManagedClusterIdentityType string

const (
	ManagedClusterIdentityTypeNone           = ManagedClusterIdentityType("None")
	ManagedClusterIdentityTypeSystemAssigned = ManagedClusterIdentityType("SystemAssigned")
	ManagedClusterIdentityTypeUserAssigned   = ManagedClusterIdentityType("UserAssigned")
)

type ManagedClusterPodIdentityProfileARM struct {
	//AllowNetworkPluginKubenet: Running in Kubenet is disabled by default due to the
	//security related nature of AAD Pod Identity and the risks of IP spoofing. See
	//[using Kubenet network plugin with AAD Pod
	//Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities)
	//for more information.
	AllowNetworkPluginKubenet *bool `json:"allowNetworkPluginKubenet,omitempty"`

	//Enabled: Whether the pod identity addon is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	//UserAssignedIdentities: The pod identities to use in the cluster.
	UserAssignedIdentities []ManagedClusterPodIdentityARM `json:"userAssignedIdentities,omitempty"`

	//UserAssignedIdentityExceptions: The pod identity exceptions to allow.
	UserAssignedIdentityExceptions []ManagedClusterPodIdentityExceptionARM `json:"userAssignedIdentityExceptions,omitempty"`
}

type ManagedClusterPropertiesAutoScalerProfileARM struct {
	//BalanceSimilarNodeGroups: Valid values are 'true' and 'false'
	BalanceSimilarNodeGroups *string `json:"balance-similar-node-groups,omitempty"`

	//Expander: If not specified, the default is 'random'. See
	//[expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders)
	//for more information.
	Expander *ManagedClusterPropertiesAutoScalerProfileExpander `json:"expander,omitempty"`

	//MaxEmptyBulkDelete: The default is 10.
	MaxEmptyBulkDelete *string `json:"max-empty-bulk-delete,omitempty"`

	//MaxGracefulTerminationSec: The default is 600.
	MaxGracefulTerminationSec *string `json:"max-graceful-termination-sec,omitempty"`

	//MaxNodeProvisionTime: The default is '15m'. Values must be an integer followed
	//by an 'm'. No unit of time other than minutes (m) is supported.
	MaxNodeProvisionTime *string `json:"max-node-provision-time,omitempty"`

	//MaxTotalUnreadyPercentage: The default is 45. The maximum is 100 and the minimum
	//is 0.
	MaxTotalUnreadyPercentage *string `json:"max-total-unready-percentage,omitempty"`

	//NewPodScaleUpDelay: For scenarios like burst/batch scale where you don't want CA
	//to act before the kubernetes scheduler could schedule all the pods, you can tell
	//CA to ignore unscheduled pods before they're a certain age. The default is '0s'.
	//Values must be an integer followed by a unit ('s' for seconds, 'm' for minutes,
	//'h' for hours, etc).
	NewPodScaleUpDelay *string `json:"new-pod-scale-up-delay,omitempty"`

	//OkTotalUnreadyCount: This must be an integer. The default is 3.
	OkTotalUnreadyCount *string `json:"ok-total-unready-count,omitempty"`

	//ScaleDownDelayAfterAdd: The default is '10m'. Values must be an integer followed
	//by an 'm'. No unit of time other than minutes (m) is supported.
	ScaleDownDelayAfterAdd *string `json:"scale-down-delay-after-add,omitempty"`

	//ScaleDownDelayAfterDelete: The default is the scan-interval. Values must be an
	//integer followed by an 'm'. No unit of time other than minutes (m) is supported.
	ScaleDownDelayAfterDelete *string `json:"scale-down-delay-after-delete,omitempty"`

	//ScaleDownDelayAfterFailure: The default is '3m'. Values must be an integer
	//followed by an 'm'. No unit of time other than minutes (m) is supported.
	ScaleDownDelayAfterFailure *string `json:"scale-down-delay-after-failure,omitempty"`

	//ScaleDownUnneededTime: The default is '10m'. Values must be an integer followed
	//by an 'm'. No unit of time other than minutes (m) is supported.
	ScaleDownUnneededTime *string `json:"scale-down-unneeded-time,omitempty"`

	//ScaleDownUnreadyTime: The default is '20m'. Values must be an integer followed
	//by an 'm'. No unit of time other than minutes (m) is supported.
	ScaleDownUnreadyTime *string `json:"scale-down-unready-time,omitempty"`

	//ScaleDownUtilizationThreshold: The default is '0.5'.
	ScaleDownUtilizationThreshold *string `json:"scale-down-utilization-threshold,omitempty"`

	//ScanInterval: The default is '10'. Values must be an integer number of seconds.
	ScanInterval *string `json:"scan-interval,omitempty"`

	//SkipNodesWithLocalStorage: The default is true.
	SkipNodesWithLocalStorage *string `json:"skip-nodes-with-local-storage,omitempty"`

	//SkipNodesWithSystemPods: The default is true.
	SkipNodesWithSystemPods *string `json:"skip-nodes-with-system-pods,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic"}
type ManagedClusterSKUName string

const ManagedClusterSKUNameBasic = ManagedClusterSKUName("Basic")

// +kubebuilder:validation:Enum={"Free","Paid"}
type ManagedClusterSKUTier string

const (
	ManagedClusterSKUTierFree = ManagedClusterSKUTier("Free")
	ManagedClusterSKUTierPaid = ManagedClusterSKUTier("Paid")
)

type ManagedClusterServicePrincipalProfileARM struct {
	//ClientId: The ID for the service principal.
	ClientId string `json:"clientId"`

	//Secret: The secret password associated with the service principal in plain text.
	Secret *string `json:"secret,omitempty"`
}

type ManagedClusterWindowsProfileARM struct {
	//AdminPassword: Specifies the password of the administrator account.
	//Minimum-length: 8 characters
	//Max-length: 123 characters
	//Complexity requirements: 3 out of 4 conditions below need to be fulfilled
	//Has lower characters
	//Has upper characters
	//Has a digit
	//Has a special character (Regex match [\W_])
	//Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word",
	//"pass@word1", "Password!", "Password1", "Password22", "iloveyou!"
	AdminPassword *string `json:"adminPassword,omitempty"`

	//AdminUsername: Specifies the name of the administrator account.
	//Restriction: Cannot end in "."
	//Disallowed values: "administrator", "admin", "user", "user1", "test", "user2",
	//"test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2",
	//"aspnet", "backup", "console", "david", "guest", "john", "owner", "root",
	//"server", "sql", "support", "support_388945a0", "sys", "test2", "test3",
	//"user4", "user5".
	//Minimum-length: 1 character
	//Max-length: 20 characters
	AdminUsername string `json:"adminUsername"`

	//EnableCSIProxy: For more details on CSI proxy, see the [CSI proxy GitHub
	//repo](https://github.com/kubernetes-csi/csi-proxy).
	EnableCSIProxy *bool `json:"enableCSIProxy,omitempty"`

	//LicenseType: The license type to use for Windows VMs. See [Azure Hybrid User
	//Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more
	//details.
	LicenseType *ManagedClusterWindowsProfileLicenseType `json:"licenseType,omitempty"`
}

type ContainerServiceSshConfigurationARM struct {
	//PublicKeys: The list of SSH public keys used to authenticate with Linux-based
	//VMs. A maximum of 1 key may be specified.
	PublicKeys []ContainerServiceSshPublicKeyARM `json:"publicKeys"`
}

type ManagedClusterLoadBalancerProfileARM struct {
	//AllocatedOutboundPorts: The desired number of allocated SNAT ports per VM.
	//Allowed values are in the range of 0 to 64000 (inclusive). The default value is
	//0 which results in Azure dynamically allocating ports.
	AllocatedOutboundPorts *int `json:"allocatedOutboundPorts,omitempty"`

	//EffectiveOutboundIPs: The effective outbound IP resources of the cluster load
	//balancer.
	EffectiveOutboundIPs []ResourceReferenceARM `json:"effectiveOutboundIPs,omitempty"`

	//IdleTimeoutInMinutes: Desired outbound flow idle timeout in minutes. Allowed
	//values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	//ManagedOutboundIPs: Desired managed outbound IPs for the cluster load balancer.
	ManagedOutboundIPs *ManagedClusterLoadBalancerProfileManagedOutboundIPsARM `json:"managedOutboundIPs,omitempty"`

	//OutboundIPPrefixes: Desired outbound IP Prefix resources for the cluster load
	//balancer.
	OutboundIPPrefixes *ManagedClusterLoadBalancerProfileOutboundIPPrefixesARM `json:"outboundIPPrefixes,omitempty"`

	//OutboundIPs: Desired outbound IP resources for the cluster load balancer.
	OutboundIPs *ManagedClusterLoadBalancerProfileOutboundIPsARM `json:"outboundIPs,omitempty"`
}

type ManagedClusterPodIdentityARM struct {
	//BindingSelector: The binding selector to use for the AzureIdentityBinding
	//resource.
	BindingSelector *string `json:"bindingSelector,omitempty"`

	//Identity: The user assigned identity details.
	Identity UserAssignedIdentityARM `json:"identity"`

	//Name: The name of the pod identity.
	Name string `json:"name"`

	//Namespace: The namespace of the pod identity.
	Namespace string `json:"namespace"`
}

type ManagedClusterPodIdentityExceptionARM struct {
	//Name: The name of the pod identity exception.
	Name string `json:"name"`

	//Namespace: The namespace of the pod identity exception.
	Namespace string `json:"namespace"`

	//PodLabels: The pod labels to match.
	PodLabels map[string]string `json:"podLabels"`
}

type ContainerServiceSshPublicKeyARM struct {
	//KeyData: Certificate public key used to authenticate with VMs through SSH. The
	//certificate must be in PEM format with or without headers.
	KeyData string `json:"keyData"`
}

type ManagedClusterLoadBalancerProfileManagedOutboundIPsARM struct {
	//Count: The desired number of outbound IPs created/managed by Azure for the
	//cluster load balancer. Allowed values must be in the range of 1 to 100
	//(inclusive). The default value is 1.
	Count *int `json:"count,omitempty"`
}

type ManagedClusterLoadBalancerProfileOutboundIPPrefixesARM struct {
	//PublicIPPrefixes: A list of public IP prefix resources.
	PublicIPPrefixes []ResourceReferenceARM `json:"publicIPPrefixes,omitempty"`
}

type ManagedClusterLoadBalancerProfileOutboundIPsARM struct {
	//PublicIPs: A list of public IP resources.
	PublicIPs []ResourceReferenceARM `json:"publicIPs,omitempty"`
}

type ResourceReferenceARM struct {
	Id *string `json:"id,omitempty"`
}

type UserAssignedIdentityARM struct {
	//ClientId: The client ID of the user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	//ObjectId: The object ID of the user assigned identity.
	ObjectId   *string `json:"objectId,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}
