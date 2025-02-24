v1api20231102preview
====================

| Metadata             | Value                                                                                |
|----------------------|--------------------------------------------------------------------------------------|
| Group                | containerservice.azure.com                                                           |
| Version              | v1api20231102preview                                                                 |
| Module               | github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231102preview |
| Property Optionality | Optional                                                                             |

<a id="APIVersion"></a>APIVersion
---------------------------------

| Value                | Description |
|----------------------|-------------|
| "2023-11-02-preview" |             |

<a id="ManagedCluster"></a>ManagedCluster
-----------------------------------------

Used by: [ManagedClusterList](#ManagedClusterList).

| Property                                                                                | Description | Type                                                                        |
|-----------------------------------------------------------------------------------------|-------------|-----------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                             |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                             |
| spec                                                                                    |             | [ManagedCluster_Spec](#ManagedCluster_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [ManagedCluster_STATUS](#ManagedCluster_STATUS)<br/><small>Optional</small> |

### <a id="ManagedCluster_Spec"></a>ManagedCluster_Spec

| Property                   | Description                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                 |
|----------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile                 |                                                                                                                                                                                                                                                                                              | [ManagedClusterAADProfile](#ManagedClusterAADProfile)<br/><small>Optional</small>                                                                                    |
| addonProfiles              |                                                                                                                                                                                                                                                                                              | [map[string]ManagedClusterAddonProfile](#ManagedClusterAddonProfile)<br/><small>Optional</small>                                                                     |
| agentPoolProfiles          |                                                                                                                                                                                                                                                                                              | [ManagedClusterAgentPoolProfile[]](#ManagedClusterAgentPoolProfile)<br/><small>Optional</small>                                                                      |
| aiToolchainOperatorProfile |                                                                                                                                                                                                                                                                                              | [ManagedClusterAIToolchainOperatorProfile](#ManagedClusterAIToolchainOperatorProfile)<br/><small>Optional</small>                                                    |
| apiServerAccessProfile     |                                                                                                                                                                                                                                                                                              | [ManagedClusterAPIServerAccessProfile](#ManagedClusterAPIServerAccessProfile)<br/><small>Optional</small>                                                            |
| autoScalerProfile          |                                                                                                                                                                                                                                                                                              | [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile)<br/><small>Optional</small>                                                |
| autoUpgradeProfile         |                                                                                                                                                                                                                                                                                              | [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile)<br/><small>Optional</small>                                                                    |
| azureMonitorProfile        |                                                                                                                                                                                                                                                                                              | [ManagedClusterAzureMonitorProfile](#ManagedClusterAzureMonitorProfile)<br/><small>Optional</small>                                                                  |
| azureName                  | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| creationData               |                                                                                                                                                                                                                                                                                              | [CreationData](#CreationData)<br/><small>Optional</small>                                                                                                            |
| disableLocalAccounts       |                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| diskEncryptionSetReference |                                                                                                                                                                                                                                                                                              | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| dnsPrefix                  |                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| enableNamespaceResources   |                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| enablePodSecurityPolicy    |                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| enableRBAC                 |                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| extendedLocation           |                                                                                                                                                                                                                                                                                              | [ExtendedLocation](#ExtendedLocation)<br/><small>Optional</small>                                                                                                    |
| fqdnSubdomain              |                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| httpProxyConfig            |                                                                                                                                                                                                                                                                                              | [ManagedClusterHTTPProxyConfig](#ManagedClusterHTTPProxyConfig)<br/><small>Optional</small>                                                                          |
| identity                   |                                                                                                                                                                                                                                                                                              | [ManagedClusterIdentity](#ManagedClusterIdentity)<br/><small>Optional</small>                                                                                        |
| identityProfile            |                                                                                                                                                                                                                                                                                              | [map[string]UserAssignedIdentity](#UserAssignedIdentity)<br/><small>Optional</small>                                                                                 |
| ingressProfile             |                                                                                                                                                                                                                                                                                              | [ManagedClusterIngressProfile](#ManagedClusterIngressProfile)<br/><small>Optional</small>                                                                            |
| kubernetesVersion          |                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| linuxProfile               |                                                                                                                                                                                                                                                                                              | [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile)<br/><small>Optional</small>                                                                            |
| location                   |                                                                                                                                                                                                                                                                                              | string<br/><small>Required</small>                                                                                                                                   |
| metricsProfile             |                                                                                                                                                                                                                                                                                              | [ManagedClusterMetricsProfile](#ManagedClusterMetricsProfile)<br/><small>Optional</small>                                                                            |
| networkProfile             |                                                                                                                                                                                                                                                                                              | [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile)<br/><small>Optional</small>                                                                        |
| nodeProvisioningProfile    |                                                                                                                                                                                                                                                                                              | [ManagedClusterNodeProvisioningProfile](#ManagedClusterNodeProvisioningProfile)<br/><small>Optional</small>                                                          |
| nodeResourceGroup          |                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| nodeResourceGroupProfile   |                                                                                                                                                                                                                                                                                              | [ManagedClusterNodeResourceGroupProfile](#ManagedClusterNodeResourceGroupProfile)<br/><small>Optional</small>                                                        |
| oidcIssuerProfile          |                                                                                                                                                                                                                                                                                              | [ManagedClusterOIDCIssuerProfile](#ManagedClusterOIDCIssuerProfile)<br/><small>Optional</small>                                                                      |
| operatorSpec               | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                              | [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec)<br/><small>Optional</small>                                                                                |
| owner                      | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| podIdentityProfile         |                                                                                                                                                                                                                                                                                              | [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile)<br/><small>Optional</small>                                                                    |
| privateLinkResources       |                                                                                                                                                                                                                                                                                              | [PrivateLinkResource[]](#PrivateLinkResource)<br/><small>Optional</small>                                                                                            |
| publicNetworkAccess        |                                                                                                                                                                                                                                                                                              | [ManagedClusterProperties_PublicNetworkAccess](#ManagedClusterProperties_PublicNetworkAccess)<br/><small>Optional</small>                                            |
| safeguardsProfile          |                                                                                                                                                                                                                                                                                              | [SafeguardsProfile](#SafeguardsProfile)<br/><small>Optional</small>                                                                                                  |
| securityProfile            |                                                                                                                                                                                                                                                                                              | [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile)<br/><small>Optional</small>                                                                          |
| serviceMeshProfile         |                                                                                                                                                                                                                                                                                              | [ServiceMeshProfile](#ServiceMeshProfile)<br/><small>Optional</small>                                                                                                |
| servicePrincipalProfile    |                                                                                                                                                                                                                                                                                              | [ManagedClusterServicePrincipalProfile](#ManagedClusterServicePrincipalProfile)<br/><small>Optional</small>                                                          |
| sku                        |                                                                                                                                                                                                                                                                                              | [ManagedClusterSKU](#ManagedClusterSKU)<br/><small>Optional</small>                                                                                                  |
| storageProfile             |                                                                                                                                                                                                                                                                                              | [ManagedClusterStorageProfile](#ManagedClusterStorageProfile)<br/><small>Optional</small>                                                                            |
| supportPlan                |                                                                                                                                                                                                                                                                                              | [KubernetesSupportPlan](#KubernetesSupportPlan)<br/><small>Optional</small>                                                                                          |
| tags                       |                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                                                                                                        |
| upgradeSettings            |                                                                                                                                                                                                                                                                                              | [ClusterUpgradeSettings](#ClusterUpgradeSettings)<br/><small>Optional</small>                                                                                        |
| windowsProfile             |                                                                                                                                                                                                                                                                                              | [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile)<br/><small>Optional</small>                                                                            |
| workloadAutoScalerProfile  |                                                                                                                                                                                                                                                                                              | [ManagedClusterWorkloadAutoScalerProfile](#ManagedClusterWorkloadAutoScalerProfile)<br/><small>Optional</small>                                                      |

### <a id="ManagedCluster_STATUS"></a>ManagedCluster_STATUS

| Property                   | Description                        | Type                                                                                                                                                    |
|----------------------------|------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile                 |                                    | [ManagedClusterAADProfile_STATUS](#ManagedClusterAADProfile_STATUS)<br/><small>Optional</small>                                                         |
| addonProfiles              |                                    | [map[string]ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS)<br/><small>Optional</small>                                          |
| agentPoolProfiles          |                                    | [ManagedClusterAgentPoolProfile_STATUS[]](#ManagedClusterAgentPoolProfile_STATUS)<br/><small>Optional</small>                                           |
| aiToolchainOperatorProfile |                                    | [ManagedClusterAIToolchainOperatorProfile_STATUS](#ManagedClusterAIToolchainOperatorProfile_STATUS)<br/><small>Optional</small>                         |
| apiServerAccessProfile     |                                    | [ManagedClusterAPIServerAccessProfile_STATUS](#ManagedClusterAPIServerAccessProfile_STATUS)<br/><small>Optional</small>                                 |
| autoScalerProfile          |                                    | [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS)<br/><small>Optional</small>                     |
| autoUpgradeProfile         |                                    | [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS)<br/><small>Optional</small>                                         |
| azureMonitorProfile        |                                    | [ManagedClusterAzureMonitorProfile_STATUS](#ManagedClusterAzureMonitorProfile_STATUS)<br/><small>Optional</small>                                       |
| azurePortalFQDN            |                                    | string<br/><small>Optional</small>                                                                                                                      |
| conditions                 | The observed state of the resource | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| creationData               |                                    | [CreationData_STATUS](#CreationData_STATUS)<br/><small>Optional</small>                                                                                 |
| currentKubernetesVersion   |                                    | string<br/><small>Optional</small>                                                                                                                      |
| disableLocalAccounts       |                                    | bool<br/><small>Optional</small>                                                                                                                        |
| diskEncryptionSetID        |                                    | string<br/><small>Optional</small>                                                                                                                      |
| dnsPrefix                  |                                    | string<br/><small>Optional</small>                                                                                                                      |
| enableNamespaceResources   |                                    | bool<br/><small>Optional</small>                                                                                                                        |
| enablePodSecurityPolicy    |                                    | bool<br/><small>Optional</small>                                                                                                                        |
| enableRBAC                 |                                    | bool<br/><small>Optional</small>                                                                                                                        |
| extendedLocation           |                                    | [ExtendedLocation_STATUS](#ExtendedLocation_STATUS)<br/><small>Optional</small>                                                                         |
| fqdn                       |                                    | string<br/><small>Optional</small>                                                                                                                      |
| fqdnSubdomain              |                                    | string<br/><small>Optional</small>                                                                                                                      |
| httpProxyConfig            |                                    | [ManagedClusterHTTPProxyConfig_STATUS](#ManagedClusterHTTPProxyConfig_STATUS)<br/><small>Optional</small>                                               |
| id                         |                                    | string<br/><small>Optional</small>                                                                                                                      |
| identity                   |                                    | [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS)<br/><small>Optional</small>                                                             |
| identityProfile            |                                    | [map[string]UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small>                                                      |
| ingressProfile             |                                    | [ManagedClusterIngressProfile_STATUS](#ManagedClusterIngressProfile_STATUS)<br/><small>Optional</small>                                                 |
| kubernetesVersion          |                                    | string<br/><small>Optional</small>                                                                                                                      |
| linuxProfile               |                                    | [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS)<br/><small>Optional</small>                                                 |
| location                   |                                    | string<br/><small>Optional</small>                                                                                                                      |
| maxAgentPools              |                                    | int<br/><small>Optional</small>                                                                                                                         |
| metricsProfile             |                                    | [ManagedClusterMetricsProfile_STATUS](#ManagedClusterMetricsProfile_STATUS)<br/><small>Optional</small>                                                 |
| name                       |                                    | string<br/><small>Optional</small>                                                                                                                      |
| networkProfile             |                                    | [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS)<br/><small>Optional</small>                                             |
| nodeProvisioningProfile    |                                    | [ManagedClusterNodeProvisioningProfile_STATUS](#ManagedClusterNodeProvisioningProfile_STATUS)<br/><small>Optional</small>                               |
| nodeResourceGroup          |                                    | string<br/><small>Optional</small>                                                                                                                      |
| nodeResourceGroupProfile   |                                    | [ManagedClusterNodeResourceGroupProfile_STATUS](#ManagedClusterNodeResourceGroupProfile_STATUS)<br/><small>Optional</small>                             |
| oidcIssuerProfile          |                                    | [ManagedClusterOIDCIssuerProfile_STATUS](#ManagedClusterOIDCIssuerProfile_STATUS)<br/><small>Optional</small>                                           |
| podIdentityProfile         |                                    | [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS)<br/><small>Optional</small>                                         |
| powerState                 |                                    | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                                     |
| privateFQDN                |                                    | string<br/><small>Optional</small>                                                                                                                      |
| privateLinkResources       |                                    | [PrivateLinkResource_STATUS[]](#PrivateLinkResource_STATUS)<br/><small>Optional</small>                                                                 |
| provisioningState          |                                    | string<br/><small>Optional</small>                                                                                                                      |
| publicNetworkAccess        |                                    | [ManagedClusterProperties_PublicNetworkAccess_STATUS](#ManagedClusterProperties_PublicNetworkAccess_STATUS)<br/><small>Optional</small>                 |
| resourceUID                |                                    | string<br/><small>Optional</small>                                                                                                                      |
| safeguardsProfile          |                                    | [SafeguardsProfile_STATUS](#SafeguardsProfile_STATUS)<br/><small>Optional</small>                                                                       |
| securityProfile            |                                    | [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS)<br/><small>Optional</small>                                               |
| serviceMeshProfile         |                                    | [ServiceMeshProfile_STATUS](#ServiceMeshProfile_STATUS)<br/><small>Optional</small>                                                                     |
| servicePrincipalProfile    |                                    | [ManagedClusterServicePrincipalProfile_STATUS](#ManagedClusterServicePrincipalProfile_STATUS)<br/><small>Optional</small>                               |
| sku                        |                                    | [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS)<br/><small>Optional</small>                                                                       |
| storageProfile             |                                    | [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS)<br/><small>Optional</small>                                                 |
| supportPlan                |                                    | [KubernetesSupportPlan_STATUS](#KubernetesSupportPlan_STATUS)<br/><small>Optional</small>                                                               |
| systemData                 |                                    | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags                       |                                    | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                       |                                    | string<br/><small>Optional</small>                                                                                                                      |
| upgradeSettings            |                                    | [ClusterUpgradeSettings_STATUS](#ClusterUpgradeSettings_STATUS)<br/><small>Optional</small>                                                             |
| windowsProfile             |                                    | [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS)<br/><small>Optional</small>                                                 |
| workloadAutoScalerProfile  |                                    | [ManagedClusterWorkloadAutoScalerProfile_STATUS](#ManagedClusterWorkloadAutoScalerProfile_STATUS)<br/><small>Optional</small>                           |

<a id="ManagedClusterList"></a>ManagedClusterList
-------------------------------------------------

| Property                                                                            | Description | Type                                                            |
|-------------------------------------------------------------------------------------|-------------|-----------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                 |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                 |
| items                                                                               |             | [ManagedCluster[]](#ManagedCluster)<br/><small>Optional</small> |

<a id="ManagedClustersAgentPool"></a>ManagedClustersAgentPool
-------------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-11-02-preview/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}

Used by: [ManagedClustersAgentPoolList](#ManagedClustersAgentPoolList).

| Property                                                                                | Description | Type                                                                                            |
|-----------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                                                 |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                                                 |
| spec                                                                                    |             | [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS)<br/><small>Optional</small> |

### <a id="ManagedClustersAgentPool_Spec"></a>ManagedClustersAgentPool_Spec

| Property                          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                 |
|-----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile          | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile](#AgentPoolArtifactStreamingProfile)<br/><small>Optional</small>                                                                  |
| availabilityZones                 | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                                 |
| azureName                         | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                                   |
| capacityReservationGroupReference | AKS will associate the specified agent pool with the Capacity Reservation Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| count                             | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | int<br/><small>Optional</small>                                                                                                                                      |
| creationData                      | CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [CreationData](#CreationData)<br/><small>Optional</small>                                                                                                            |
| enableAutoScaling                 | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                                     |
| enableCustomCATrust               | When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded certificates into node trust stores. Defaults to false.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                                     |
| enableEncryptionAtHost            | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                                     |
| enableFIPS                        | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | bool<br/><small>Optional</small>                                                                                                                                     |
| enableNodePublicIP                | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                                                                                                                                                                                                                                                                                                                                            | bool<br/><small>Optional</small>                                                                                                                                     |
| enableUltraSSD                    | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| gpuInstanceProfile                | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile](#GPUInstanceProfile)<br/><small>Optional</small>                                                                                                |
| gpuProfile                        | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile](#AgentPoolGPUProfile)<br/><small>Optional</small>                                                                                              |
| hostGroupReference                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/hostGroups/{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| kubeletConfig                     | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig](#KubeletConfig)<br/><small>Optional</small>                                                                                                          |
| kubeletDiskType                   | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType](#KubeletDiskType)<br/><small>Optional</small>                                                                                                      |
| linuxOSConfig                     | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig](#LinuxOSConfig)<br/><small>Optional</small>                                                                                                          |
| maxCount                          | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                                      |
| maxPods                           | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                                      |
| messageOfTheDay                   | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                                   |
| minCount                          | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                                      |
| mode                              | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode](#AgentPoolMode)<br/><small>Optional</small>                                                                                                          |
| networkProfile                    | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile](#AgentPoolNetworkProfile)<br/><small>Optional</small>                                                                                      |
| nodeInitializationTaints          | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                                 |
| nodeLabels                        | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                                                                        |
| nodePublicIPPrefixReference       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| nodeTaints                        | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                                                                                 |
| operatorSpec                      | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ManagedClustersAgentPoolOperatorSpec](#ManagedClustersAgentPoolOperatorSpec)<br/><small>Optional</small>                                                            |
| orchestratorVersion               | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                                                                   |
| osDiskSizeGB                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ContainerServiceOSDisk](#ContainerServiceOSDisk)<br/><small>Optional</small>                                                                                        |
| osDiskType                        | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType](#OSDiskType)<br/><small>Optional</small>                                                                                                                |
| osSKU                             | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU](#OSSKU)<br/><small>Optional</small>                                                                                                                          |
| osType                            | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType](#OSType)<br/><small>Optional</small>                                                                                                                        |
| owner                             | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| podSubnetReference                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                                                                                                                                                                                                                                                                                                                                                       | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| powerState                        | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState](#PowerState)<br/><small>Optional</small>                                                                                                                |
| proximityPlacementGroupReference  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| scaleDownMode                     | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode](#ScaleDownMode)<br/><small>Optional</small>                                                                                                          |
| scaleSetEvictionPolicy            | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)<br/><small>Optional</small>                                                                                        |
| scaleSetPriority                  | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority](#ScaleSetPriority)<br/><small>Optional</small>                                                                                                    |
| securityProfile                   | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile](#AgentPoolSecurityProfile)<br/><small>Optional</small>                                                                                    |
| spotMaxPrice                      | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                                                                                  |
| tags                              | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                                                                        |
| type                              | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType](#AgentPoolType)<br/><small>Optional</small>                                                                                                          |
| upgradeSettings                   | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings)<br/><small>Optional</small>                                                                                    |
| virtualMachineNodesStatus         |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes[]](#VirtualMachineNodes)<br/><small>Optional</small>                                                                                            |
| virtualMachinesProfile            | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile](#VirtualMachinesProfile)<br/><small>Optional</small>                                                                                        |
| vmSize                            | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                                   |
| vnetSubnetReference               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| windowsProfile                    | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile](#AgentPoolWindowsProfile)<br/><small>Optional</small>                                                                                      |
| workloadRuntime                   | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime](#WorkloadRuntime)<br/><small>Optional</small>                                                                                                      |

### <a id="ManagedClustersAgentPool_STATUS"></a>ManagedClustersAgentPool_STATUS

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                    |
|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile   | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile_STATUS](#AgentPoolArtifactStreamingProfile_STATUS)<br/><small>Optional</small>                                       |
| availabilityZones          | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                    |
| capacityReservationGroupID | AKS will associate the specified agent pool with the Capacity Reservation Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                                                                                      |
| conditions                 | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| count                      | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | int<br/><small>Optional</small>                                                                                                                         |
| creationData               | CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [CreationData_STATUS](#CreationData_STATUS)<br/><small>Optional</small>                                                                                 |
| currentOrchestratorVersion | If orchestratorVersion was a fully specified version <major.minor.patch>, this field will be exactly equal to it. If orchestratorVersion was <major.minor>, this field will contain the full <major.minor.patch> version being used.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| enableAutoScaling          | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                        |
| enableCustomCATrust        | When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded certificates into node trust stores. Defaults to false.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                        |
| enableEncryptionAtHost     | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                        |
| enableFIPS                 | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | bool<br/><small>Optional</small>                                                                                                                        |
| enableNodePublicIP         | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                                                                                                                                                                                                                                                                                                                                            | bool<br/><small>Optional</small>                                                                                                                        |
| enableUltraSSD             | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                        |
| gpuInstanceProfile         | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)<br/><small>Optional</small>                                                                     |
| gpuProfile                 | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile_STATUS](#AgentPoolGPUProfile_STATUS)<br/><small>Optional</small>                                                                   |
| hostGroupID                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/hostGroups/{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| id                         | Resource ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| kubeletConfig              | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig_STATUS](#KubeletConfig_STATUS)<br/><small>Optional</small>                                                                               |
| kubeletDiskType            | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)<br/><small>Optional</small>                                                                           |
| linuxOSConfig              | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)<br/><small>Optional</small>                                                                               |
| maxCount                   | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                         |
| maxPods                    | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                         |
| messageOfTheDay            | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                      |
| minCount                   | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                         |
| mode                       | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)<br/><small>Optional</small>                                                                               |
| name                       | The name of the resource that is unique within a resource group. This name can be used to access the resource.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                      |
| networkProfile             | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS)<br/><small>Optional</small>                                                           |
| nodeImageVersion           | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| nodeInitializationTaints   | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                    |
| nodeLabels                 | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                                                           |
| nodePublicIPPrefixID       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| nodeTaints                 | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                                                                    |
| orchestratorVersion        | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                                                      |
| osDiskSizeGB               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | int<br/><small>Optional</small>                                                                                                                         |
| osDiskType                 | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType_STATUS](#OSDiskType_STATUS)<br/><small>Optional</small>                                                                                     |
| osSKU                      | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU_STATUS](#OSSKU_STATUS)<br/><small>Optional</small>                                                                                               |
| osType                     | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType_STATUS](#OSType_STATUS)<br/><small>Optional</small>                                                                                             |
| podSubnetID                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                                                                                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| powerState                 | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                                     |
| properties_type            | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType_STATUS](#AgentPoolType_STATUS)<br/><small>Optional</small>                                                                               |
| provisioningState          | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| proximityPlacementGroupID  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| scaleDownMode              | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode_STATUS](#ScaleDownMode_STATUS)<br/><small>Optional</small>                                                                               |
| scaleSetEvictionPolicy     | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)<br/><small>Optional</small>                                                             |
| scaleSetPriority           | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)<br/><small>Optional</small>                                                                         |
| securityProfile            | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile_STATUS](#AgentPoolSecurityProfile_STATUS)<br/><small>Optional</small>                                                         |
| spotMaxPrice               | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                                                                     |
| tags                       | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                       | Resource type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| upgradeSettings            | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS)<br/><small>Optional</small>                                                         |
| virtualMachineNodesStatus  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes_STATUS[]](#VirtualMachineNodes_STATUS)<br/><small>Optional</small>                                                                 |
| virtualMachinesProfile     | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile_STATUS](#VirtualMachinesProfile_STATUS)<br/><small>Optional</small>                                                             |
| vmSize                     | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| vnetSubnetID               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| windowsProfile             | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile_STATUS](#AgentPoolWindowsProfile_STATUS)<br/><small>Optional</small>                                                           |
| workloadRuntime            | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime_STATUS](#WorkloadRuntime_STATUS)<br/><small>Optional</small>                                                                           |

<a id="ManagedClustersAgentPoolList"></a>ManagedClustersAgentPoolList
---------------------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-11-02-preview/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}

| Property                                                                            | Description | Type                                                                                |
|-------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                                     |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                                     |
| items                                                                               |             | [ManagedClustersAgentPool[]](#ManagedClustersAgentPool)<br/><small>Optional</small> |

<a id="ManagedCluster_Spec"></a>ManagedCluster_Spec
---------------------------------------------------

Used by: [ManagedCluster](#ManagedCluster).

| Property                   | Description                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                 |
|----------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile                 |                                                                                                                                                                                                                                                                                              | [ManagedClusterAADProfile](#ManagedClusterAADProfile)<br/><small>Optional</small>                                                                                    |
| addonProfiles              |                                                                                                                                                                                                                                                                                              | [map[string]ManagedClusterAddonProfile](#ManagedClusterAddonProfile)<br/><small>Optional</small>                                                                     |
| agentPoolProfiles          |                                                                                                                                                                                                                                                                                              | [ManagedClusterAgentPoolProfile[]](#ManagedClusterAgentPoolProfile)<br/><small>Optional</small>                                                                      |
| aiToolchainOperatorProfile |                                                                                                                                                                                                                                                                                              | [ManagedClusterAIToolchainOperatorProfile](#ManagedClusterAIToolchainOperatorProfile)<br/><small>Optional</small>                                                    |
| apiServerAccessProfile     |                                                                                                                                                                                                                                                                                              | [ManagedClusterAPIServerAccessProfile](#ManagedClusterAPIServerAccessProfile)<br/><small>Optional</small>                                                            |
| autoScalerProfile          |                                                                                                                                                                                                                                                                                              | [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile)<br/><small>Optional</small>                                                |
| autoUpgradeProfile         |                                                                                                                                                                                                                                                                                              | [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile)<br/><small>Optional</small>                                                                    |
| azureMonitorProfile        |                                                                                                                                                                                                                                                                                              | [ManagedClusterAzureMonitorProfile](#ManagedClusterAzureMonitorProfile)<br/><small>Optional</small>                                                                  |
| azureName                  | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| creationData               |                                                                                                                                                                                                                                                                                              | [CreationData](#CreationData)<br/><small>Optional</small>                                                                                                            |
| disableLocalAccounts       |                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| diskEncryptionSetReference |                                                                                                                                                                                                                                                                                              | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| dnsPrefix                  |                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| enableNamespaceResources   |                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| enablePodSecurityPolicy    |                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| enableRBAC                 |                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| extendedLocation           |                                                                                                                                                                                                                                                                                              | [ExtendedLocation](#ExtendedLocation)<br/><small>Optional</small>                                                                                                    |
| fqdnSubdomain              |                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| httpProxyConfig            |                                                                                                                                                                                                                                                                                              | [ManagedClusterHTTPProxyConfig](#ManagedClusterHTTPProxyConfig)<br/><small>Optional</small>                                                                          |
| identity                   |                                                                                                                                                                                                                                                                                              | [ManagedClusterIdentity](#ManagedClusterIdentity)<br/><small>Optional</small>                                                                                        |
| identityProfile            |                                                                                                                                                                                                                                                                                              | [map[string]UserAssignedIdentity](#UserAssignedIdentity)<br/><small>Optional</small>                                                                                 |
| ingressProfile             |                                                                                                                                                                                                                                                                                              | [ManagedClusterIngressProfile](#ManagedClusterIngressProfile)<br/><small>Optional</small>                                                                            |
| kubernetesVersion          |                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| linuxProfile               |                                                                                                                                                                                                                                                                                              | [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile)<br/><small>Optional</small>                                                                            |
| location                   |                                                                                                                                                                                                                                                                                              | string<br/><small>Required</small>                                                                                                                                   |
| metricsProfile             |                                                                                                                                                                                                                                                                                              | [ManagedClusterMetricsProfile](#ManagedClusterMetricsProfile)<br/><small>Optional</small>                                                                            |
| networkProfile             |                                                                                                                                                                                                                                                                                              | [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile)<br/><small>Optional</small>                                                                        |
| nodeProvisioningProfile    |                                                                                                                                                                                                                                                                                              | [ManagedClusterNodeProvisioningProfile](#ManagedClusterNodeProvisioningProfile)<br/><small>Optional</small>                                                          |
| nodeResourceGroup          |                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| nodeResourceGroupProfile   |                                                                                                                                                                                                                                                                                              | [ManagedClusterNodeResourceGroupProfile](#ManagedClusterNodeResourceGroupProfile)<br/><small>Optional</small>                                                        |
| oidcIssuerProfile          |                                                                                                                                                                                                                                                                                              | [ManagedClusterOIDCIssuerProfile](#ManagedClusterOIDCIssuerProfile)<br/><small>Optional</small>                                                                      |
| operatorSpec               | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                              | [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec)<br/><small>Optional</small>                                                                                |
| owner                      | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| podIdentityProfile         |                                                                                                                                                                                                                                                                                              | [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile)<br/><small>Optional</small>                                                                    |
| privateLinkResources       |                                                                                                                                                                                                                                                                                              | [PrivateLinkResource[]](#PrivateLinkResource)<br/><small>Optional</small>                                                                                            |
| publicNetworkAccess        |                                                                                                                                                                                                                                                                                              | [ManagedClusterProperties_PublicNetworkAccess](#ManagedClusterProperties_PublicNetworkAccess)<br/><small>Optional</small>                                            |
| safeguardsProfile          |                                                                                                                                                                                                                                                                                              | [SafeguardsProfile](#SafeguardsProfile)<br/><small>Optional</small>                                                                                                  |
| securityProfile            |                                                                                                                                                                                                                                                                                              | [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile)<br/><small>Optional</small>                                                                          |
| serviceMeshProfile         |                                                                                                                                                                                                                                                                                              | [ServiceMeshProfile](#ServiceMeshProfile)<br/><small>Optional</small>                                                                                                |
| servicePrincipalProfile    |                                                                                                                                                                                                                                                                                              | [ManagedClusterServicePrincipalProfile](#ManagedClusterServicePrincipalProfile)<br/><small>Optional</small>                                                          |
| sku                        |                                                                                                                                                                                                                                                                                              | [ManagedClusterSKU](#ManagedClusterSKU)<br/><small>Optional</small>                                                                                                  |
| storageProfile             |                                                                                                                                                                                                                                                                                              | [ManagedClusterStorageProfile](#ManagedClusterStorageProfile)<br/><small>Optional</small>                                                                            |
| supportPlan                |                                                                                                                                                                                                                                                                                              | [KubernetesSupportPlan](#KubernetesSupportPlan)<br/><small>Optional</small>                                                                                          |
| tags                       |                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                                                                                                        |
| upgradeSettings            |                                                                                                                                                                                                                                                                                              | [ClusterUpgradeSettings](#ClusterUpgradeSettings)<br/><small>Optional</small>                                                                                        |
| windowsProfile             |                                                                                                                                                                                                                                                                                              | [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile)<br/><small>Optional</small>                                                                            |
| workloadAutoScalerProfile  |                                                                                                                                                                                                                                                                                              | [ManagedClusterWorkloadAutoScalerProfile](#ManagedClusterWorkloadAutoScalerProfile)<br/><small>Optional</small>                                                      |

<a id="ManagedCluster_STATUS"></a>ManagedCluster_STATUS
-------------------------------------------------------

Used by: [ManagedCluster](#ManagedCluster).

| Property                   | Description                        | Type                                                                                                                                                    |
|----------------------------|------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile                 |                                    | [ManagedClusterAADProfile_STATUS](#ManagedClusterAADProfile_STATUS)<br/><small>Optional</small>                                                         |
| addonProfiles              |                                    | [map[string]ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS)<br/><small>Optional</small>                                          |
| agentPoolProfiles          |                                    | [ManagedClusterAgentPoolProfile_STATUS[]](#ManagedClusterAgentPoolProfile_STATUS)<br/><small>Optional</small>                                           |
| aiToolchainOperatorProfile |                                    | [ManagedClusterAIToolchainOperatorProfile_STATUS](#ManagedClusterAIToolchainOperatorProfile_STATUS)<br/><small>Optional</small>                         |
| apiServerAccessProfile     |                                    | [ManagedClusterAPIServerAccessProfile_STATUS](#ManagedClusterAPIServerAccessProfile_STATUS)<br/><small>Optional</small>                                 |
| autoScalerProfile          |                                    | [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS)<br/><small>Optional</small>                     |
| autoUpgradeProfile         |                                    | [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS)<br/><small>Optional</small>                                         |
| azureMonitorProfile        |                                    | [ManagedClusterAzureMonitorProfile_STATUS](#ManagedClusterAzureMonitorProfile_STATUS)<br/><small>Optional</small>                                       |
| azurePortalFQDN            |                                    | string<br/><small>Optional</small>                                                                                                                      |
| conditions                 | The observed state of the resource | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| creationData               |                                    | [CreationData_STATUS](#CreationData_STATUS)<br/><small>Optional</small>                                                                                 |
| currentKubernetesVersion   |                                    | string<br/><small>Optional</small>                                                                                                                      |
| disableLocalAccounts       |                                    | bool<br/><small>Optional</small>                                                                                                                        |
| diskEncryptionSetID        |                                    | string<br/><small>Optional</small>                                                                                                                      |
| dnsPrefix                  |                                    | string<br/><small>Optional</small>                                                                                                                      |
| enableNamespaceResources   |                                    | bool<br/><small>Optional</small>                                                                                                                        |
| enablePodSecurityPolicy    |                                    | bool<br/><small>Optional</small>                                                                                                                        |
| enableRBAC                 |                                    | bool<br/><small>Optional</small>                                                                                                                        |
| extendedLocation           |                                    | [ExtendedLocation_STATUS](#ExtendedLocation_STATUS)<br/><small>Optional</small>                                                                         |
| fqdn                       |                                    | string<br/><small>Optional</small>                                                                                                                      |
| fqdnSubdomain              |                                    | string<br/><small>Optional</small>                                                                                                                      |
| httpProxyConfig            |                                    | [ManagedClusterHTTPProxyConfig_STATUS](#ManagedClusterHTTPProxyConfig_STATUS)<br/><small>Optional</small>                                               |
| id                         |                                    | string<br/><small>Optional</small>                                                                                                                      |
| identity                   |                                    | [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS)<br/><small>Optional</small>                                                             |
| identityProfile            |                                    | [map[string]UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small>                                                      |
| ingressProfile             |                                    | [ManagedClusterIngressProfile_STATUS](#ManagedClusterIngressProfile_STATUS)<br/><small>Optional</small>                                                 |
| kubernetesVersion          |                                    | string<br/><small>Optional</small>                                                                                                                      |
| linuxProfile               |                                    | [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS)<br/><small>Optional</small>                                                 |
| location                   |                                    | string<br/><small>Optional</small>                                                                                                                      |
| maxAgentPools              |                                    | int<br/><small>Optional</small>                                                                                                                         |
| metricsProfile             |                                    | [ManagedClusterMetricsProfile_STATUS](#ManagedClusterMetricsProfile_STATUS)<br/><small>Optional</small>                                                 |
| name                       |                                    | string<br/><small>Optional</small>                                                                                                                      |
| networkProfile             |                                    | [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS)<br/><small>Optional</small>                                             |
| nodeProvisioningProfile    |                                    | [ManagedClusterNodeProvisioningProfile_STATUS](#ManagedClusterNodeProvisioningProfile_STATUS)<br/><small>Optional</small>                               |
| nodeResourceGroup          |                                    | string<br/><small>Optional</small>                                                                                                                      |
| nodeResourceGroupProfile   |                                    | [ManagedClusterNodeResourceGroupProfile_STATUS](#ManagedClusterNodeResourceGroupProfile_STATUS)<br/><small>Optional</small>                             |
| oidcIssuerProfile          |                                    | [ManagedClusterOIDCIssuerProfile_STATUS](#ManagedClusterOIDCIssuerProfile_STATUS)<br/><small>Optional</small>                                           |
| podIdentityProfile         |                                    | [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS)<br/><small>Optional</small>                                         |
| powerState                 |                                    | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                                     |
| privateFQDN                |                                    | string<br/><small>Optional</small>                                                                                                                      |
| privateLinkResources       |                                    | [PrivateLinkResource_STATUS[]](#PrivateLinkResource_STATUS)<br/><small>Optional</small>                                                                 |
| provisioningState          |                                    | string<br/><small>Optional</small>                                                                                                                      |
| publicNetworkAccess        |                                    | [ManagedClusterProperties_PublicNetworkAccess_STATUS](#ManagedClusterProperties_PublicNetworkAccess_STATUS)<br/><small>Optional</small>                 |
| resourceUID                |                                    | string<br/><small>Optional</small>                                                                                                                      |
| safeguardsProfile          |                                    | [SafeguardsProfile_STATUS](#SafeguardsProfile_STATUS)<br/><small>Optional</small>                                                                       |
| securityProfile            |                                    | [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS)<br/><small>Optional</small>                                               |
| serviceMeshProfile         |                                    | [ServiceMeshProfile_STATUS](#ServiceMeshProfile_STATUS)<br/><small>Optional</small>                                                                     |
| servicePrincipalProfile    |                                    | [ManagedClusterServicePrincipalProfile_STATUS](#ManagedClusterServicePrincipalProfile_STATUS)<br/><small>Optional</small>                               |
| sku                        |                                    | [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS)<br/><small>Optional</small>                                                                       |
| storageProfile             |                                    | [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS)<br/><small>Optional</small>                                                 |
| supportPlan                |                                    | [KubernetesSupportPlan_STATUS](#KubernetesSupportPlan_STATUS)<br/><small>Optional</small>                                                               |
| systemData                 |                                    | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags                       |                                    | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                       |                                    | string<br/><small>Optional</small>                                                                                                                      |
| upgradeSettings            |                                    | [ClusterUpgradeSettings_STATUS](#ClusterUpgradeSettings_STATUS)<br/><small>Optional</small>                                                             |
| windowsProfile             |                                    | [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS)<br/><small>Optional</small>                                                 |
| workloadAutoScalerProfile  |                                    | [ManagedClusterWorkloadAutoScalerProfile_STATUS](#ManagedClusterWorkloadAutoScalerProfile_STATUS)<br/><small>Optional</small>                           |

<a id="ManagedClustersAgentPool_Spec"></a>ManagedClustersAgentPool_Spec
-----------------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#ManagedClustersAgentPool).

| Property                          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                 |
|-----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile          | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile](#AgentPoolArtifactStreamingProfile)<br/><small>Optional</small>                                                                  |
| availabilityZones                 | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                                 |
| azureName                         | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                                   |
| capacityReservationGroupReference | AKS will associate the specified agent pool with the Capacity Reservation Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| count                             | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | int<br/><small>Optional</small>                                                                                                                                      |
| creationData                      | CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [CreationData](#CreationData)<br/><small>Optional</small>                                                                                                            |
| enableAutoScaling                 | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                                     |
| enableCustomCATrust               | When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded certificates into node trust stores. Defaults to false.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                                     |
| enableEncryptionAtHost            | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                                     |
| enableFIPS                        | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | bool<br/><small>Optional</small>                                                                                                                                     |
| enableNodePublicIP                | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                                                                                                                                                                                                                                                                                                                                            | bool<br/><small>Optional</small>                                                                                                                                     |
| enableUltraSSD                    | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| gpuInstanceProfile                | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile](#GPUInstanceProfile)<br/><small>Optional</small>                                                                                                |
| gpuProfile                        | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile](#AgentPoolGPUProfile)<br/><small>Optional</small>                                                                                              |
| hostGroupReference                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/hostGroups/{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| kubeletConfig                     | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig](#KubeletConfig)<br/><small>Optional</small>                                                                                                          |
| kubeletDiskType                   | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType](#KubeletDiskType)<br/><small>Optional</small>                                                                                                      |
| linuxOSConfig                     | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig](#LinuxOSConfig)<br/><small>Optional</small>                                                                                                          |
| maxCount                          | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                                      |
| maxPods                           | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                                      |
| messageOfTheDay                   | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                                   |
| minCount                          | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                                      |
| mode                              | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode](#AgentPoolMode)<br/><small>Optional</small>                                                                                                          |
| networkProfile                    | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile](#AgentPoolNetworkProfile)<br/><small>Optional</small>                                                                                      |
| nodeInitializationTaints          | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                                 |
| nodeLabels                        | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                                                                        |
| nodePublicIPPrefixReference       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| nodeTaints                        | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                                                                                 |
| operatorSpec                      | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ManagedClustersAgentPoolOperatorSpec](#ManagedClustersAgentPoolOperatorSpec)<br/><small>Optional</small>                                                            |
| orchestratorVersion               | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                                                                   |
| osDiskSizeGB                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ContainerServiceOSDisk](#ContainerServiceOSDisk)<br/><small>Optional</small>                                                                                        |
| osDiskType                        | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType](#OSDiskType)<br/><small>Optional</small>                                                                                                                |
| osSKU                             | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU](#OSSKU)<br/><small>Optional</small>                                                                                                                          |
| osType                            | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType](#OSType)<br/><small>Optional</small>                                                                                                                        |
| owner                             | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| podSubnetReference                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                                                                                                                                                                                                                                                                                                                                                       | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| powerState                        | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState](#PowerState)<br/><small>Optional</small>                                                                                                                |
| proximityPlacementGroupReference  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| scaleDownMode                     | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode](#ScaleDownMode)<br/><small>Optional</small>                                                                                                          |
| scaleSetEvictionPolicy            | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)<br/><small>Optional</small>                                                                                        |
| scaleSetPriority                  | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority](#ScaleSetPriority)<br/><small>Optional</small>                                                                                                    |
| securityProfile                   | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile](#AgentPoolSecurityProfile)<br/><small>Optional</small>                                                                                    |
| spotMaxPrice                      | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                                                                                  |
| tags                              | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                                                                        |
| type                              | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType](#AgentPoolType)<br/><small>Optional</small>                                                                                                          |
| upgradeSettings                   | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings)<br/><small>Optional</small>                                                                                    |
| virtualMachineNodesStatus         |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes[]](#VirtualMachineNodes)<br/><small>Optional</small>                                                                                            |
| virtualMachinesProfile            | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile](#VirtualMachinesProfile)<br/><small>Optional</small>                                                                                        |
| vmSize                            | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                                   |
| vnetSubnetReference               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| windowsProfile                    | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile](#AgentPoolWindowsProfile)<br/><small>Optional</small>                                                                                      |
| workloadRuntime                   | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime](#WorkloadRuntime)<br/><small>Optional</small>                                                                                                      |

<a id="ManagedClustersAgentPool_STATUS"></a>ManagedClustersAgentPool_STATUS
---------------------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#ManagedClustersAgentPool).

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                    |
|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile   | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile_STATUS](#AgentPoolArtifactStreamingProfile_STATUS)<br/><small>Optional</small>                                       |
| availabilityZones          | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                    |
| capacityReservationGroupID | AKS will associate the specified agent pool with the Capacity Reservation Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                                                                                      |
| conditions                 | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| count                      | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | int<br/><small>Optional</small>                                                                                                                         |
| creationData               | CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [CreationData_STATUS](#CreationData_STATUS)<br/><small>Optional</small>                                                                                 |
| currentOrchestratorVersion | If orchestratorVersion was a fully specified version <major.minor.patch>, this field will be exactly equal to it. If orchestratorVersion was <major.minor>, this field will contain the full <major.minor.patch> version being used.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| enableAutoScaling          | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                        |
| enableCustomCATrust        | When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded certificates into node trust stores. Defaults to false.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                        |
| enableEncryptionAtHost     | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                        |
| enableFIPS                 | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | bool<br/><small>Optional</small>                                                                                                                        |
| enableNodePublicIP         | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                                                                                                                                                                                                                                                                                                                                            | bool<br/><small>Optional</small>                                                                                                                        |
| enableUltraSSD             | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                        |
| gpuInstanceProfile         | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)<br/><small>Optional</small>                                                                     |
| gpuProfile                 | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile_STATUS](#AgentPoolGPUProfile_STATUS)<br/><small>Optional</small>                                                                   |
| hostGroupID                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/hostGroups/{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| id                         | Resource ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| kubeletConfig              | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig_STATUS](#KubeletConfig_STATUS)<br/><small>Optional</small>                                                                               |
| kubeletDiskType            | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)<br/><small>Optional</small>                                                                           |
| linuxOSConfig              | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)<br/><small>Optional</small>                                                                               |
| maxCount                   | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                         |
| maxPods                    | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                         |
| messageOfTheDay            | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                      |
| minCount                   | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                         |
| mode                       | A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)<br/><small>Optional</small>                                                                               |
| name                       | The name of the resource that is unique within a resource group. This name can be used to access the resource.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                      |
| networkProfile             | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS)<br/><small>Optional</small>                                                           |
| nodeImageVersion           | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| nodeInitializationTaints   | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                    |
| nodeLabels                 | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                                                           |
| nodePublicIPPrefixID       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| nodeTaints                 | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                                                                    |
| orchestratorVersion        | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                                                      |
| osDiskSizeGB               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | int<br/><small>Optional</small>                                                                                                                         |
| osDiskType                 | The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType_STATUS](#OSDiskType_STATUS)<br/><small>Optional</small>                                                                                     |
| osSKU                      | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU_STATUS](#OSSKU_STATUS)<br/><small>Optional</small>                                                                                               |
| osType                     | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType_STATUS](#OSType_STATUS)<br/><small>Optional</small>                                                                                             |
| podSubnetID                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                                                                                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| powerState                 | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                                     |
| properties_type            | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType_STATUS](#AgentPoolType_STATUS)<br/><small>Optional</small>                                                                               |
| provisioningState          | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| proximityPlacementGroupID  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| scaleDownMode              | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode_STATUS](#ScaleDownMode_STATUS)<br/><small>Optional</small>                                                                               |
| scaleSetEvictionPolicy     | This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)<br/><small>Optional</small>                                                             |
| scaleSetPriority           | The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)<br/><small>Optional</small>                                                                         |
| securityProfile            | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile_STATUS](#AgentPoolSecurityProfile_STATUS)<br/><small>Optional</small>                                                         |
| spotMaxPrice               | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                                                                     |
| tags                       | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                       | Resource type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| upgradeSettings            | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS)<br/><small>Optional</small>                                                         |
| virtualMachineNodesStatus  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes_STATUS[]](#VirtualMachineNodes_STATUS)<br/><small>Optional</small>                                                                 |
| virtualMachinesProfile     | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile_STATUS](#VirtualMachinesProfile_STATUS)<br/><small>Optional</small>                                                             |
| vmSize                     | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| vnetSubnetID               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| windowsProfile             | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile_STATUS](#AgentPoolWindowsProfile_STATUS)<br/><small>Optional</small>                                                           |
| workloadRuntime            | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime_STATUS](#WorkloadRuntime_STATUS)<br/><small>Optional</small>                                                                           |

<a id="AgentPoolArtifactStreamingProfile"></a>AgentPoolArtifactStreamingProfile
-------------------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="AgentPoolArtifactStreamingProfile_STATUS"></a>AgentPoolArtifactStreamingProfile_STATUS
---------------------------------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="AgentPoolGPUProfile"></a>AgentPoolGPUProfile
---------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property         | Description | Type                             |
|------------------|-------------|----------------------------------|
| installGPUDriver |             | bool<br/><small>Optional</small> |

<a id="AgentPoolGPUProfile_STATUS"></a>AgentPoolGPUProfile_STATUS
-----------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property         | Description | Type                             |
|------------------|-------------|----------------------------------|
| installGPUDriver |             | bool<br/><small>Optional</small> |

<a id="AgentPoolMode"></a>AgentPoolMode
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value    | Description |
|----------|-------------|
| "System" |             |
| "User"   |             |

<a id="AgentPoolMode_STATUS"></a>AgentPoolMode_STATUS
-----------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value    | Description |
|----------|-------------|
| "System" |             |
| "User"   |             |

<a id="AgentPoolNetworkProfile"></a>AgentPoolNetworkProfile
-----------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property                            | Description | Type                                                                                                                                                         |
|-------------------------------------|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| allowedHostPorts                    |             | [PortRange[]](#PortRange)<br/><small>Optional</small>                                                                                                        |
| applicationSecurityGroupsReferences |             | [genruntime.ResourceReference[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| nodePublicIPTags                    |             | [IPTag[]](#IPTag)<br/><small>Optional</small>                                                                                                                |

<a id="AgentPoolNetworkProfile_STATUS"></a>AgentPoolNetworkProfile_STATUS
-------------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property                  | Description | Type                                                                |
|---------------------------|-------------|---------------------------------------------------------------------|
| allowedHostPorts          |             | [PortRange_STATUS[]](#PortRange_STATUS)<br/><small>Optional</small> |
| applicationSecurityGroups |             | string[]<br/><small>Optional</small>                                |
| nodePublicIPTags          |             | [IPTag_STATUS[]](#IPTag_STATUS)<br/><small>Optional</small>         |

<a id="AgentPoolSecurityProfile"></a>AgentPoolSecurityProfile
-------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property         | Description | Type                                                                  |
|------------------|-------------|-----------------------------------------------------------------------|
| enableSecureBoot |             | bool<br/><small>Optional</small>                                      |
| enableVTPM       |             | bool<br/><small>Optional</small>                                      |
| sshAccess        |             | [AgentPoolSSHAccess](#AgentPoolSSHAccess)<br/><small>Optional</small> |

<a id="AgentPoolSecurityProfile_STATUS"></a>AgentPoolSecurityProfile_STATUS
---------------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property         | Description | Type                                                                                |
|------------------|-------------|-------------------------------------------------------------------------------------|
| enableSecureBoot |             | bool<br/><small>Optional</small>                                                    |
| enableVTPM       |             | bool<br/><small>Optional</small>                                                    |
| sshAccess        |             | [AgentPoolSSHAccess_STATUS](#AgentPoolSSHAccess_STATUS)<br/><small>Optional</small> |

<a id="AgentPoolType"></a>AgentPoolType
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value                     | Description |
|---------------------------|-------------|
| "AvailabilitySet"         |             |
| "VirtualMachineScaleSets" |             |
| "VirtualMachines"         |             |

<a id="AgentPoolType_STATUS"></a>AgentPoolType_STATUS
-----------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value                     | Description |
|---------------------------|-------------|
| "AvailabilitySet"         |             |
| "VirtualMachineScaleSets" |             |
| "VirtualMachines"         |             |

<a id="AgentPoolUpgradeSettings"></a>AgentPoolUpgradeSettings
-------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property                  | Description | Type                               |
|---------------------------|-------------|------------------------------------|
| drainTimeoutInMinutes     |             | int<br/><small>Optional</small>    |
| maxSurge                  |             | string<br/><small>Optional</small> |
| nodeSoakDurationInMinutes |             | int<br/><small>Optional</small>    |

<a id="AgentPoolUpgradeSettings_STATUS"></a>AgentPoolUpgradeSettings_STATUS
---------------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property                  | Description | Type                               |
|---------------------------|-------------|------------------------------------|
| drainTimeoutInMinutes     |             | int<br/><small>Optional</small>    |
| maxSurge                  |             | string<br/><small>Optional</small> |
| nodeSoakDurationInMinutes |             | int<br/><small>Optional</small>    |

<a id="AgentPoolWindowsProfile"></a>AgentPoolWindowsProfile
-----------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property           | Description | Type                             |
|--------------------|-------------|----------------------------------|
| disableOutboundNat |             | bool<br/><small>Optional</small> |

<a id="AgentPoolWindowsProfile_STATUS"></a>AgentPoolWindowsProfile_STATUS
-------------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property           | Description | Type                             |
|--------------------|-------------|----------------------------------|
| disableOutboundNat |             | bool<br/><small>Optional</small> |

<a id="ClusterUpgradeSettings"></a>ClusterUpgradeSettings
---------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property         | Description | Type                                                                            |
|------------------|-------------|---------------------------------------------------------------------------------|
| overrideSettings |             | [UpgradeOverrideSettings](#UpgradeOverrideSettings)<br/><small>Optional</small> |

<a id="ClusterUpgradeSettings_STATUS"></a>ClusterUpgradeSettings_STATUS
-----------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property         | Description | Type                                                                                          |
|------------------|-------------|-----------------------------------------------------------------------------------------------|
| overrideSettings |             | [UpgradeOverrideSettings_STATUS](#UpgradeOverrideSettings_STATUS)<br/><small>Optional</small> |

<a id="ContainerServiceLinuxProfile"></a>ContainerServiceLinuxProfile
---------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property      | Description | Type                                                                                              |
|---------------|-------------|---------------------------------------------------------------------------------------------------|
| adminUsername |             | string<br/><small>Required</small>                                                                |
| ssh           |             | [ContainerServiceSshConfiguration](#ContainerServiceSshConfiguration)<br/><small>Required</small> |

<a id="ContainerServiceLinuxProfile_STATUS"></a>ContainerServiceLinuxProfile_STATUS
-----------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property      | Description | Type                                                                                                            |
|---------------|-------------|-----------------------------------------------------------------------------------------------------------------|
| adminUsername |             | string<br/><small>Optional</small>                                                                              |
| ssh           |             | [ContainerServiceSshConfiguration_STATUS](#ContainerServiceSshConfiguration_STATUS)<br/><small>Optional</small> |

<a id="ContainerServiceNetworkProfile"></a>ContainerServiceNetworkProfile
-------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property            | Description | Type                                                                                                                          |
|---------------------|-------------|-------------------------------------------------------------------------------------------------------------------------------|
| dnsServiceIP        |             | string<br/><small>Optional</small>                                                                                            |
| ipFamilies          |             | [IpFamily[]](#IpFamily)<br/><small>Optional</small>                                                                           |
| kubeProxyConfig     |             | [ContainerServiceNetworkProfile_KubeProxyConfig](#ContainerServiceNetworkProfile_KubeProxyConfig)<br/><small>Optional</small> |
| loadBalancerProfile |             | [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile)<br/><small>Optional</small>                           |
| loadBalancerSku     |             | [LoadBalancerSku](#LoadBalancerSku)<br/><small>Optional</small>                                                               |
| monitoring          |             | [NetworkMonitoring](#NetworkMonitoring)<br/><small>Optional</small>                                                           |
| natGatewayProfile   |             | [ManagedClusterNATGatewayProfile](#ManagedClusterNATGatewayProfile)<br/><small>Optional</small>                               |
| networkDataplane    |             | [NetworkDataplane](#NetworkDataplane)<br/><small>Optional</small>                                                             |
| networkMode         |             | [NetworkMode](#NetworkMode)<br/><small>Optional</small>                                                                       |
| networkPlugin       |             | [NetworkPlugin](#NetworkPlugin)<br/><small>Optional</small>                                                                   |
| networkPluginMode   |             | [NetworkPluginMode](#NetworkPluginMode)<br/><small>Optional</small>                                                           |
| networkPolicy       |             | [NetworkPolicy](#NetworkPolicy)<br/><small>Optional</small>                                                                   |
| outboundType        |             | [ContainerServiceNetworkProfile_OutboundType](#ContainerServiceNetworkProfile_OutboundType)<br/><small>Optional</small>       |
| podCidr             |             | string<br/><small>Optional</small>                                                                                            |
| podCidrs            |             | string[]<br/><small>Optional</small>                                                                                          |
| serviceCidr         |             | string<br/><small>Optional</small>                                                                                            |
| serviceCidrs        |             | string[]<br/><small>Optional</small>                                                                                          |

<a id="ContainerServiceNetworkProfile_STATUS"></a>ContainerServiceNetworkProfile_STATUS
---------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property            | Description | Type                                                                                                                                        |
|---------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| dnsServiceIP        |             | string<br/><small>Optional</small>                                                                                                          |
| ipFamilies          |             | [IpFamily_STATUS[]](#IpFamily_STATUS)<br/><small>Optional</small>                                                                           |
| kubeProxyConfig     |             | [ContainerServiceNetworkProfile_KubeProxyConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_STATUS)<br/><small>Optional</small> |
| loadBalancerProfile |             | [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS)<br/><small>Optional</small>                           |
| loadBalancerSku     |             | [LoadBalancerSku_STATUS](#LoadBalancerSku_STATUS)<br/><small>Optional</small>                                                               |
| monitoring          |             | [NetworkMonitoring_STATUS](#NetworkMonitoring_STATUS)<br/><small>Optional</small>                                                           |
| natGatewayProfile   |             | [ManagedClusterNATGatewayProfile_STATUS](#ManagedClusterNATGatewayProfile_STATUS)<br/><small>Optional</small>                               |
| networkDataplane    |             | [NetworkDataplane_STATUS](#NetworkDataplane_STATUS)<br/><small>Optional</small>                                                             |
| networkMode         |             | [NetworkMode_STATUS](#NetworkMode_STATUS)<br/><small>Optional</small>                                                                       |
| networkPlugin       |             | [NetworkPlugin_STATUS](#NetworkPlugin_STATUS)<br/><small>Optional</small>                                                                   |
| networkPluginMode   |             | [NetworkPluginMode_STATUS](#NetworkPluginMode_STATUS)<br/><small>Optional</small>                                                           |
| networkPolicy       |             | [NetworkPolicy_STATUS](#NetworkPolicy_STATUS)<br/><small>Optional</small>                                                                   |
| outboundType        |             | [ContainerServiceNetworkProfile_OutboundType_STATUS](#ContainerServiceNetworkProfile_OutboundType_STATUS)<br/><small>Optional</small>       |
| podCidr             |             | string<br/><small>Optional</small>                                                                                                          |
| podCidrs            |             | string[]<br/><small>Optional</small>                                                                                                        |
| serviceCidr         |             | string<br/><small>Optional</small>                                                                                                          |
| serviceCidrs        |             | string[]<br/><small>Optional</small>                                                                                                        |

<a id="ContainerServiceOSDisk"></a>ContainerServiceOSDisk
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

<a id="CreationData"></a>CreationData
-------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec), [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property                | Description | Type                                                                                                                                                       |
|-------------------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| sourceResourceReference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

<a id="CreationData_STATUS"></a>CreationData_STATUS
---------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property         | Description | Type                               |
|------------------|-------------|------------------------------------|
| sourceResourceId |             | string<br/><small>Optional</small> |

<a id="ExtendedLocation"></a>ExtendedLocation
---------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description | Type                                                                      |
|----------|-------------|---------------------------------------------------------------------------|
| name     |             | string<br/><small>Optional</small>                                        |
| type     |             | [ExtendedLocationType](#ExtendedLocationType)<br/><small>Optional</small> |

<a id="ExtendedLocation_STATUS"></a>ExtendedLocation_STATUS
-----------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description | Type                                                                                    |
|----------|-------------|-----------------------------------------------------------------------------------------|
| name     |             | string<br/><small>Optional</small>                                                      |
| type     |             | [ExtendedLocationType_STATUS](#ExtendedLocationType_STATUS)<br/><small>Optional</small> |

<a id="GPUInstanceProfile"></a>GPUInstanceProfile
-------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value   | Description |
|---------|-------------|
| "MIG1g" |             |
| "MIG2g" |             |
| "MIG3g" |             |
| "MIG4g" |             |
| "MIG7g" |             |

<a id="GPUInstanceProfile_STATUS"></a>GPUInstanceProfile_STATUS
---------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value   | Description |
|---------|-------------|
| "MIG1g" |             |
| "MIG2g" |             |
| "MIG3g" |             |
| "MIG4g" |             |
| "MIG7g" |             |

<a id="KubeletConfig"></a>KubeletConfig
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property              | Description | Type                                 |
|-----------------------|-------------|--------------------------------------|
| allowedUnsafeSysctls  |             | string[]<br/><small>Optional</small> |
| containerLogMaxFiles  |             | int<br/><small>Optional</small>      |
| containerLogMaxSizeMB |             | int<br/><small>Optional</small>      |
| cpuCfsQuota           |             | bool<br/><small>Optional</small>     |
| cpuCfsQuotaPeriod     |             | string<br/><small>Optional</small>   |
| cpuManagerPolicy      |             | string<br/><small>Optional</small>   |
| failSwapOn            |             | bool<br/><small>Optional</small>     |
| imageGcHighThreshold  |             | int<br/><small>Optional</small>      |
| imageGcLowThreshold   |             | int<br/><small>Optional</small>      |
| podMaxPids            |             | int<br/><small>Optional</small>      |
| topologyManagerPolicy |             | string<br/><small>Optional</small>   |

<a id="KubeletConfig_STATUS"></a>KubeletConfig_STATUS
-----------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property              | Description | Type                                 |
|-----------------------|-------------|--------------------------------------|
| allowedUnsafeSysctls  |             | string[]<br/><small>Optional</small> |
| containerLogMaxFiles  |             | int<br/><small>Optional</small>      |
| containerLogMaxSizeMB |             | int<br/><small>Optional</small>      |
| cpuCfsQuota           |             | bool<br/><small>Optional</small>     |
| cpuCfsQuotaPeriod     |             | string<br/><small>Optional</small>   |
| cpuManagerPolicy      |             | string<br/><small>Optional</small>   |
| failSwapOn            |             | bool<br/><small>Optional</small>     |
| imageGcHighThreshold  |             | int<br/><small>Optional</small>      |
| imageGcLowThreshold   |             | int<br/><small>Optional</small>      |
| podMaxPids            |             | int<br/><small>Optional</small>      |
| topologyManagerPolicy |             | string<br/><small>Optional</small>   |

<a id="KubeletDiskType"></a>KubeletDiskType
-------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value       | Description |
|-------------|-------------|
| "OS"        |             |
| "Temporary" |             |

<a id="KubeletDiskType_STATUS"></a>KubeletDiskType_STATUS
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value       | Description |
|-------------|-------------|
| "OS"        |             |
| "Temporary" |             |

<a id="KubernetesSupportPlan"></a>KubernetesSupportPlan
-------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Value                | Description |
|----------------------|-------------|
| "AKSLongTermSupport" |             |
| "KubernetesOfficial" |             |

<a id="KubernetesSupportPlan_STATUS"></a>KubernetesSupportPlan_STATUS
---------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Value                | Description |
|----------------------|-------------|
| "AKSLongTermSupport" |             |
| "KubernetesOfficial" |             |

<a id="LinuxOSConfig"></a>LinuxOSConfig
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property                   | Description | Type                                                      |
|----------------------------|-------------|-----------------------------------------------------------|
| swapFileSizeMB             |             | int<br/><small>Optional</small>                           |
| sysctls                    |             | [SysctlConfig](#SysctlConfig)<br/><small>Optional</small> |
| transparentHugePageDefrag  |             | string<br/><small>Optional</small>                        |
| transparentHugePageEnabled |             | string<br/><small>Optional</small>                        |

<a id="LinuxOSConfig_STATUS"></a>LinuxOSConfig_STATUS
-----------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property                   | Description | Type                                                                    |
|----------------------------|-------------|-------------------------------------------------------------------------|
| swapFileSizeMB             |             | int<br/><small>Optional</small>                                         |
| sysctls                    |             | [SysctlConfig_STATUS](#SysctlConfig_STATUS)<br/><small>Optional</small> |
| transparentHugePageDefrag  |             | string<br/><small>Optional</small>                                      |
| transparentHugePageEnabled |             | string<br/><small>Optional</small>                                      |

<a id="ManagedClusterAADProfile"></a>ManagedClusterAADProfile
-------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property            | Description | Type                                 |
|---------------------|-------------|--------------------------------------|
| adminGroupObjectIDs |             | string[]<br/><small>Optional</small> |
| clientAppID         |             | string<br/><small>Optional</small>   |
| enableAzureRBAC     |             | bool<br/><small>Optional</small>     |
| managed             |             | bool<br/><small>Optional</small>     |
| serverAppID         |             | string<br/><small>Optional</small>   |
| serverAppSecret     |             | string<br/><small>Optional</small>   |
| tenantID            |             | string<br/><small>Optional</small>   |

<a id="ManagedClusterAADProfile_STATUS"></a>ManagedClusterAADProfile_STATUS
---------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property            | Description | Type                                 |
|---------------------|-------------|--------------------------------------|
| adminGroupObjectIDs |             | string[]<br/><small>Optional</small> |
| clientAppID         |             | string<br/><small>Optional</small>   |
| enableAzureRBAC     |             | bool<br/><small>Optional</small>     |
| managed             |             | bool<br/><small>Optional</small>     |
| serverAppID         |             | string<br/><small>Optional</small>   |
| serverAppSecret     |             | string<br/><small>Optional</small>   |
| tenantID            |             | string<br/><small>Optional</small>   |

<a id="ManagedClusterAddonProfile"></a>ManagedClusterAddonProfile
-----------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description | Type                                          |
|----------|-------------|-----------------------------------------------|
| config   |             | map[string]string<br/><small>Optional</small> |
| enabled  |             | bool<br/><small>Required</small>              |

<a id="ManagedClusterAddonProfile_STATUS"></a>ManagedClusterAddonProfile_STATUS
-------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description | Type                                                                                    |
|----------|-------------|-----------------------------------------------------------------------------------------|
| config   |             | map[string]string<br/><small>Optional</small>                                           |
| enabled  |             | bool<br/><small>Optional</small>                                                        |
| identity |             | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterAgentPoolProfile"></a>ManagedClusterAgentPoolProfile
-------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                          | Description | Type                                                                                                                                                       |
|-----------------------------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile          |             | [AgentPoolArtifactStreamingProfile](#AgentPoolArtifactStreamingProfile)<br/><small>Optional</small>                                                        |
| availabilityZones                 |             | string[]<br/><small>Optional</small>                                                                                                                       |
| capacityReservationGroupReference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| count                             |             | int<br/><small>Optional</small>                                                                                                                            |
| creationData                      |             | [CreationData](#CreationData)<br/><small>Optional</small>                                                                                                  |
| enableAutoScaling                 |             | bool<br/><small>Optional</small>                                                                                                                           |
| enableCustomCATrust               |             | bool<br/><small>Optional</small>                                                                                                                           |
| enableEncryptionAtHost            |             | bool<br/><small>Optional</small>                                                                                                                           |
| enableFIPS                        |             | bool<br/><small>Optional</small>                                                                                                                           |
| enableNodePublicIP                |             | bool<br/><small>Optional</small>                                                                                                                           |
| enableUltraSSD                    |             | bool<br/><small>Optional</small>                                                                                                                           |
| gpuInstanceProfile                |             | [GPUInstanceProfile](#GPUInstanceProfile)<br/><small>Optional</small>                                                                                      |
| gpuProfile                        |             | [AgentPoolGPUProfile](#AgentPoolGPUProfile)<br/><small>Optional</small>                                                                                    |
| hostGroupReference                |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| kubeletConfig                     |             | [KubeletConfig](#KubeletConfig)<br/><small>Optional</small>                                                                                                |
| kubeletDiskType                   |             | [KubeletDiskType](#KubeletDiskType)<br/><small>Optional</small>                                                                                            |
| linuxOSConfig                     |             | [LinuxOSConfig](#LinuxOSConfig)<br/><small>Optional</small>                                                                                                |
| maxCount                          |             | int<br/><small>Optional</small>                                                                                                                            |
| maxPods                           |             | int<br/><small>Optional</small>                                                                                                                            |
| messageOfTheDay                   |             | string<br/><small>Optional</small>                                                                                                                         |
| minCount                          |             | int<br/><small>Optional</small>                                                                                                                            |
| mode                              |             | [AgentPoolMode](#AgentPoolMode)<br/><small>Optional</small>                                                                                                |
| name                              |             | string<br/><small>Required</small>                                                                                                                         |
| networkProfile                    |             | [AgentPoolNetworkProfile](#AgentPoolNetworkProfile)<br/><small>Optional</small>                                                                            |
| nodeInitializationTaints          |             | string[]<br/><small>Optional</small>                                                                                                                       |
| nodeLabels                        |             | map[string]string<br/><small>Optional</small>                                                                                                              |
| nodePublicIPPrefixReference       |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| nodeTaints                        |             | string[]<br/><small>Optional</small>                                                                                                                       |
| orchestratorVersion               |             | string<br/><small>Optional</small>                                                                                                                         |
| osDiskSizeGB                      |             | [ContainerServiceOSDisk](#ContainerServiceOSDisk)<br/><small>Optional</small>                                                                              |
| osDiskType                        |             | [OSDiskType](#OSDiskType)<br/><small>Optional</small>                                                                                                      |
| osSKU                             |             | [OSSKU](#OSSKU)<br/><small>Optional</small>                                                                                                                |
| osType                            |             | [OSType](#OSType)<br/><small>Optional</small>                                                                                                              |
| podSubnetReference                |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| powerState                        |             | [PowerState](#PowerState)<br/><small>Optional</small>                                                                                                      |
| proximityPlacementGroupReference  |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| scaleDownMode                     |             | [ScaleDownMode](#ScaleDownMode)<br/><small>Optional</small>                                                                                                |
| scaleSetEvictionPolicy            |             | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)<br/><small>Optional</small>                                                                              |
| scaleSetPriority                  |             | [ScaleSetPriority](#ScaleSetPriority)<br/><small>Optional</small>                                                                                          |
| securityProfile                   |             | [AgentPoolSecurityProfile](#AgentPoolSecurityProfile)<br/><small>Optional</small>                                                                          |
| spotMaxPrice                      |             | float64<br/><small>Optional</small>                                                                                                                        |
| tags                              |             | map[string]string<br/><small>Optional</small>                                                                                                              |
| type                              |             | [AgentPoolType](#AgentPoolType)<br/><small>Optional</small>                                                                                                |
| upgradeSettings                   |             | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings)<br/><small>Optional</small>                                                                          |
| virtualMachineNodesStatus         |             | [VirtualMachineNodes[]](#VirtualMachineNodes)<br/><small>Optional</small>                                                                                  |
| virtualMachinesProfile            |             | [VirtualMachinesProfile](#VirtualMachinesProfile)<br/><small>Optional</small>                                                                              |
| vmSize                            |             | string<br/><small>Optional</small>                                                                                                                         |
| vnetSubnetReference               |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| windowsProfile                    |             | [AgentPoolWindowsProfile](#AgentPoolWindowsProfile)<br/><small>Optional</small>                                                                            |
| workloadRuntime                   |             | [WorkloadRuntime](#WorkloadRuntime)<br/><small>Optional</small>                                                                                            |

<a id="ManagedClusterAgentPoolProfile_STATUS"></a>ManagedClusterAgentPoolProfile_STATUS
---------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                   | Description | Type                                                                                                              |
|----------------------------|-------------|-------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile   |             | [AgentPoolArtifactStreamingProfile_STATUS](#AgentPoolArtifactStreamingProfile_STATUS)<br/><small>Optional</small> |
| availabilityZones          |             | string[]<br/><small>Optional</small>                                                                              |
| capacityReservationGroupID |             | string<br/><small>Optional</small>                                                                                |
| count                      |             | int<br/><small>Optional</small>                                                                                   |
| creationData               |             | [CreationData_STATUS](#CreationData_STATUS)<br/><small>Optional</small>                                           |
| currentOrchestratorVersion |             | string<br/><small>Optional</small>                                                                                |
| enableAutoScaling          |             | bool<br/><small>Optional</small>                                                                                  |
| enableCustomCATrust        |             | bool<br/><small>Optional</small>                                                                                  |
| enableEncryptionAtHost     |             | bool<br/><small>Optional</small>                                                                                  |
| enableFIPS                 |             | bool<br/><small>Optional</small>                                                                                  |
| enableNodePublicIP         |             | bool<br/><small>Optional</small>                                                                                  |
| enableUltraSSD             |             | bool<br/><small>Optional</small>                                                                                  |
| gpuInstanceProfile         |             | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)<br/><small>Optional</small>                               |
| gpuProfile                 |             | [AgentPoolGPUProfile_STATUS](#AgentPoolGPUProfile_STATUS)<br/><small>Optional</small>                             |
| hostGroupID                |             | string<br/><small>Optional</small>                                                                                |
| kubeletConfig              |             | [KubeletConfig_STATUS](#KubeletConfig_STATUS)<br/><small>Optional</small>                                         |
| kubeletDiskType            |             | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)<br/><small>Optional</small>                                     |
| linuxOSConfig              |             | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)<br/><small>Optional</small>                                         |
| maxCount                   |             | int<br/><small>Optional</small>                                                                                   |
| maxPods                    |             | int<br/><small>Optional</small>                                                                                   |
| messageOfTheDay            |             | string<br/><small>Optional</small>                                                                                |
| minCount                   |             | int<br/><small>Optional</small>                                                                                   |
| mode                       |             | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)<br/><small>Optional</small>                                         |
| name                       |             | string<br/><small>Optional</small>                                                                                |
| networkProfile             |             | [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS)<br/><small>Optional</small>                     |
| nodeImageVersion           |             | string<br/><small>Optional</small>                                                                                |
| nodeInitializationTaints   |             | string[]<br/><small>Optional</small>                                                                              |
| nodeLabels                 |             | map[string]string<br/><small>Optional</small>                                                                     |
| nodePublicIPPrefixID       |             | string<br/><small>Optional</small>                                                                                |
| nodeTaints                 |             | string[]<br/><small>Optional</small>                                                                              |
| orchestratorVersion        |             | string<br/><small>Optional</small>                                                                                |
| osDiskSizeGB               |             | int<br/><small>Optional</small>                                                                                   |
| osDiskType                 |             | [OSDiskType_STATUS](#OSDiskType_STATUS)<br/><small>Optional</small>                                               |
| osSKU                      |             | [OSSKU_STATUS](#OSSKU_STATUS)<br/><small>Optional</small>                                                         |
| osType                     |             | [OSType_STATUS](#OSType_STATUS)<br/><small>Optional</small>                                                       |
| podSubnetID                |             | string<br/><small>Optional</small>                                                                                |
| powerState                 |             | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                               |
| provisioningState          |             | string<br/><small>Optional</small>                                                                                |
| proximityPlacementGroupID  |             | string<br/><small>Optional</small>                                                                                |
| scaleDownMode              |             | [ScaleDownMode_STATUS](#ScaleDownMode_STATUS)<br/><small>Optional</small>                                         |
| scaleSetEvictionPolicy     |             | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)<br/><small>Optional</small>                       |
| scaleSetPriority           |             | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)<br/><small>Optional</small>                                   |
| securityProfile            |             | [AgentPoolSecurityProfile_STATUS](#AgentPoolSecurityProfile_STATUS)<br/><small>Optional</small>                   |
| spotMaxPrice               |             | float64<br/><small>Optional</small>                                                                               |
| tags                       |             | map[string]string<br/><small>Optional</small>                                                                     |
| type                       |             | [AgentPoolType_STATUS](#AgentPoolType_STATUS)<br/><small>Optional</small>                                         |
| upgradeSettings            |             | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS)<br/><small>Optional</small>                   |
| virtualMachineNodesStatus  |             | [VirtualMachineNodes_STATUS[]](#VirtualMachineNodes_STATUS)<br/><small>Optional</small>                           |
| virtualMachinesProfile     |             | [VirtualMachinesProfile_STATUS](#VirtualMachinesProfile_STATUS)<br/><small>Optional</small>                       |
| vmSize                     |             | string<br/><small>Optional</small>                                                                                |
| vnetSubnetID               |             | string<br/><small>Optional</small>                                                                                |
| windowsProfile             |             | [AgentPoolWindowsProfile_STATUS](#AgentPoolWindowsProfile_STATUS)<br/><small>Optional</small>                     |
| workloadRuntime            |             | [WorkloadRuntime_STATUS](#WorkloadRuntime_STATUS)<br/><small>Optional</small>                                     |

<a id="ManagedClusterAIToolchainOperatorProfile"></a>ManagedClusterAIToolchainOperatorProfile
---------------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterAIToolchainOperatorProfile_STATUS"></a>ManagedClusterAIToolchainOperatorProfile_STATUS
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterAPIServerAccessProfile"></a>ManagedClusterAPIServerAccessProfile
-------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                       | Description | Type                                 |
|--------------------------------|-------------|--------------------------------------|
| authorizedIPRanges             |             | string[]<br/><small>Optional</small> |
| disableRunCommand              |             | bool<br/><small>Optional</small>     |
| enablePrivateCluster           |             | bool<br/><small>Optional</small>     |
| enablePrivateClusterPublicFQDN |             | bool<br/><small>Optional</small>     |
| enableVnetIntegration          |             | bool<br/><small>Optional</small>     |
| privateDNSZone                 |             | string<br/><small>Optional</small>   |
| subnetId                       |             | string<br/><small>Optional</small>   |

<a id="ManagedClusterAPIServerAccessProfile_STATUS"></a>ManagedClusterAPIServerAccessProfile_STATUS
---------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                       | Description | Type                                 |
|--------------------------------|-------------|--------------------------------------|
| authorizedIPRanges             |             | string[]<br/><small>Optional</small> |
| disableRunCommand              |             | bool<br/><small>Optional</small>     |
| enablePrivateCluster           |             | bool<br/><small>Optional</small>     |
| enablePrivateClusterPublicFQDN |             | bool<br/><small>Optional</small>     |
| enableVnetIntegration          |             | bool<br/><small>Optional</small>     |
| privateDNSZone                 |             | string<br/><small>Optional</small>   |
| subnetId                       |             | string<br/><small>Optional</small>   |

<a id="ManagedClusterAutoUpgradeProfile"></a>ManagedClusterAutoUpgradeProfile
-----------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property             | Description | Type                                                                                                                                        |
|----------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| nodeOSUpgradeChannel |             | [ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel](#ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel)<br/><small>Optional</small> |
| upgradeChannel       |             | [ManagedClusterAutoUpgradeProfile_UpgradeChannel](#ManagedClusterAutoUpgradeProfile_UpgradeChannel)<br/><small>Optional</small>             |

<a id="ManagedClusterAutoUpgradeProfile_STATUS"></a>ManagedClusterAutoUpgradeProfile_STATUS
-------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property             | Description | Type                                                                                                                                                      |
|----------------------|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| nodeOSUpgradeChannel |             | [ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS](#ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS)<br/><small>Optional</small> |
| upgradeChannel       |             | [ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS](#ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS)<br/><small>Optional</small>             |

<a id="ManagedClusterAzureMonitorProfile"></a>ManagedClusterAzureMonitorProfile
-------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description | Type                                                                                                              |
|----------|-------------|-------------------------------------------------------------------------------------------------------------------|
| logs     |             | [ManagedClusterAzureMonitorProfileLogs](#ManagedClusterAzureMonitorProfileLogs)<br/><small>Optional</small>       |
| metrics  |             | [ManagedClusterAzureMonitorProfileMetrics](#ManagedClusterAzureMonitorProfileMetrics)<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfile_STATUS"></a>ManagedClusterAzureMonitorProfile_STATUS
---------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description | Type                                                                                                                            |
|----------|-------------|---------------------------------------------------------------------------------------------------------------------------------|
| logs     |             | [ManagedClusterAzureMonitorProfileLogs_STATUS](#ManagedClusterAzureMonitorProfileLogs_STATUS)<br/><small>Optional</small>       |
| metrics  |             | [ManagedClusterAzureMonitorProfileMetrics_STATUS](#ManagedClusterAzureMonitorProfileMetrics_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterHTTPProxyConfig"></a>ManagedClusterHTTPProxyConfig
-----------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property   | Description | Type                                 |
|------------|-------------|--------------------------------------|
| httpProxy  |             | string<br/><small>Optional</small>   |
| httpsProxy |             | string<br/><small>Optional</small>   |
| noProxy    |             | string[]<br/><small>Optional</small> |
| trustedCa  |             | string<br/><small>Optional</small>   |

<a id="ManagedClusterHTTPProxyConfig_STATUS"></a>ManagedClusterHTTPProxyConfig_STATUS
-------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property         | Description | Type                                 |
|------------------|-------------|--------------------------------------|
| effectiveNoProxy |             | string[]<br/><small>Optional</small> |
| httpProxy        |             | string<br/><small>Optional</small>   |
| httpsProxy       |             | string<br/><small>Optional</small>   |
| noProxy          |             | string[]<br/><small>Optional</small> |
| trustedCa        |             | string<br/><small>Optional</small>   |

<a id="ManagedClusterIdentity"></a>ManagedClusterIdentity
---------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property               | Description | Type                                                                                      |
|------------------------|-------------|-------------------------------------------------------------------------------------------|
| delegatedResources     |             | [map[string]DelegatedResource](#DelegatedResource)<br/><small>Optional</small>            |
| type                   |             | [ManagedClusterIdentity_Type](#ManagedClusterIdentity_Type)<br/><small>Optional</small>   |
| userAssignedIdentities |             | [UserAssignedIdentityDetails[]](#UserAssignedIdentityDetails)<br/><small>Optional</small> |

<a id="ManagedClusterIdentity_STATUS"></a>ManagedClusterIdentity_STATUS
-----------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property               | Description | Type                                                                                                                                                 |
|------------------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| delegatedResources     |             | [map[string]DelegatedResource_STATUS](#DelegatedResource_STATUS)<br/><small>Optional</small>                                                         |
| principalId            |             | string<br/><small>Optional</small>                                                                                                                   |
| tenantId               |             | string<br/><small>Optional</small>                                                                                                                   |
| type                   |             | [ManagedClusterIdentity_Type_STATUS](#ManagedClusterIdentity_Type_STATUS)<br/><small>Optional</small>                                                |
| userAssignedIdentities |             | [map[string]ManagedClusterIdentity_UserAssignedIdentities_STATUS](#ManagedClusterIdentity_UserAssignedIdentities_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterIngressProfile"></a>ManagedClusterIngressProfile
---------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property      | Description | Type                                                                                                                |
|---------------|-------------|---------------------------------------------------------------------------------------------------------------------|
| webAppRouting |             | [ManagedClusterIngressProfileWebAppRouting](#ManagedClusterIngressProfileWebAppRouting)<br/><small>Optional</small> |

<a id="ManagedClusterIngressProfile_STATUS"></a>ManagedClusterIngressProfile_STATUS
-----------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property      | Description | Type                                                                                                                              |
|---------------|-------------|-----------------------------------------------------------------------------------------------------------------------------------|
| webAppRouting |             | [ManagedClusterIngressProfileWebAppRouting_STATUS](#ManagedClusterIngressProfileWebAppRouting_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterMetricsProfile"></a>ManagedClusterMetricsProfile
---------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property     | Description | Type                                                                                  |
|--------------|-------------|---------------------------------------------------------------------------------------|
| costAnalysis |             | [ManagedClusterCostAnalysis](#ManagedClusterCostAnalysis)<br/><small>Optional</small> |

<a id="ManagedClusterMetricsProfile_STATUS"></a>ManagedClusterMetricsProfile_STATUS
-----------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property     | Description | Type                                                                                                |
|--------------|-------------|-----------------------------------------------------------------------------------------------------|
| costAnalysis |             | [ManagedClusterCostAnalysis_STATUS](#ManagedClusterCostAnalysis_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterNodeProvisioningProfile"></a>ManagedClusterNodeProvisioningProfile
---------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description | Type                                                                                                                  |
|----------|-------------|-----------------------------------------------------------------------------------------------------------------------|
| mode     |             | [ManagedClusterNodeProvisioningProfile_Mode](#ManagedClusterNodeProvisioningProfile_Mode)<br/><small>Optional</small> |

<a id="ManagedClusterNodeProvisioningProfile_STATUS"></a>ManagedClusterNodeProvisioningProfile_STATUS
-----------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description | Type                                                                                                                                |
|----------|-------------|-------------------------------------------------------------------------------------------------------------------------------------|
| mode     |             | [ManagedClusterNodeProvisioningProfile_Mode_STATUS](#ManagedClusterNodeProvisioningProfile_Mode_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterNodeResourceGroupProfile"></a>ManagedClusterNodeResourceGroupProfile
-----------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property         | Description | Type                                                                                                                                            |
|------------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| restrictionLevel |             | [ManagedClusterNodeResourceGroupProfile_RestrictionLevel](#ManagedClusterNodeResourceGroupProfile_RestrictionLevel)<br/><small>Optional</small> |

<a id="ManagedClusterNodeResourceGroupProfile_STATUS"></a>ManagedClusterNodeResourceGroupProfile_STATUS
-------------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property         | Description | Type                                                                                                                                                          |
|------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| restrictionLevel |             | [ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS](#ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterOIDCIssuerProfile"></a>ManagedClusterOIDCIssuerProfile
---------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterOIDCIssuerProfile_STATUS"></a>ManagedClusterOIDCIssuerProfile_STATUS
-----------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property  | Description | Type                               |
|-----------|-------------|------------------------------------|
| enabled   |             | bool<br/><small>Optional</small>   |
| issuerURL |             | string<br/><small>Optional</small> |

<a id="ManagedClusterOperatorSpec"></a>ManagedClusterOperatorSpec
-----------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| configMaps           | configures where to place operator written ConfigMaps.                                        | [ManagedClusterOperatorConfigMaps](#ManagedClusterOperatorConfigMaps)<br/><small>Optional</small>                                                                   |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secrets              | configures where to place Azure generated secrets.                                            | [ManagedClusterOperatorSecrets](#ManagedClusterOperatorSecrets)<br/><small>Optional</small>                                                                         |

<a id="ManagedClusterPodIdentityProfile"></a>ManagedClusterPodIdentityProfile
-----------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                       | Description | Type                                                                                                    |
|--------------------------------|-------------|---------------------------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      |             | bool<br/><small>Optional</small>                                                                        |
| enabled                        |             | bool<br/><small>Optional</small>                                                                        |
| userAssignedIdentities         |             | [ManagedClusterPodIdentity[]](#ManagedClusterPodIdentity)<br/><small>Optional</small>                   |
| userAssignedIdentityExceptions |             | [ManagedClusterPodIdentityException[]](#ManagedClusterPodIdentityException)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentityProfile_STATUS"></a>ManagedClusterPodIdentityProfile_STATUS
-------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                       | Description | Type                                                                                                                  |
|--------------------------------|-------------|-----------------------------------------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      |             | bool<br/><small>Optional</small>                                                                                      |
| enabled                        |             | bool<br/><small>Optional</small>                                                                                      |
| userAssignedIdentities         |             | [ManagedClusterPodIdentity_STATUS[]](#ManagedClusterPodIdentity_STATUS)<br/><small>Optional</small>                   |
| userAssignedIdentityExceptions |             | [ManagedClusterPodIdentityException_STATUS[]](#ManagedClusterPodIdentityException_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterProperties_AutoScalerProfile"></a>ManagedClusterProperties_AutoScalerProfile
-------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                              | Description | Type                                              |
|---------------------------------------|-------------|---------------------------------------------------|
| balance-similar-node-groups           |             | string<br/><small>Optional</small>                |
| daemonset-eviction-for-empty-nodes    |             | bool<br/><small>Optional</small>                  |
| daemonset-eviction-for-occupied-nodes |             | bool<br/><small>Optional</small>                  |
| expander                              |             | [Expander](#Expander)<br/><small>Optional</small> |
| ignore-daemonsets-utilization         |             | bool<br/><small>Optional</small>                  |
| max-empty-bulk-delete                 |             | string<br/><small>Optional</small>                |
| max-graceful-termination-sec          |             | string<br/><small>Optional</small>                |
| max-node-provision-time               |             | string<br/><small>Optional</small>                |
| max-total-unready-percentage          |             | string<br/><small>Optional</small>                |
| new-pod-scale-up-delay                |             | string<br/><small>Optional</small>                |
| ok-total-unready-count                |             | string<br/><small>Optional</small>                |
| scale-down-delay-after-add            |             | string<br/><small>Optional</small>                |
| scale-down-delay-after-delete         |             | string<br/><small>Optional</small>                |
| scale-down-delay-after-failure        |             | string<br/><small>Optional</small>                |
| scale-down-unneeded-time              |             | string<br/><small>Optional</small>                |
| scale-down-unready-time               |             | string<br/><small>Optional</small>                |
| scale-down-utilization-threshold      |             | string<br/><small>Optional</small>                |
| scan-interval                         |             | string<br/><small>Optional</small>                |
| skip-nodes-with-local-storage         |             | string<br/><small>Optional</small>                |
| skip-nodes-with-system-pods           |             | string<br/><small>Optional</small>                |

<a id="ManagedClusterProperties_AutoScalerProfile_STATUS"></a>ManagedClusterProperties_AutoScalerProfile_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                              | Description | Type                                                            |
|---------------------------------------|-------------|-----------------------------------------------------------------|
| balance-similar-node-groups           |             | string<br/><small>Optional</small>                              |
| daemonset-eviction-for-empty-nodes    |             | bool<br/><small>Optional</small>                                |
| daemonset-eviction-for-occupied-nodes |             | bool<br/><small>Optional</small>                                |
| expander                              |             | [Expander_STATUS](#Expander_STATUS)<br/><small>Optional</small> |
| ignore-daemonsets-utilization         |             | bool<br/><small>Optional</small>                                |
| max-empty-bulk-delete                 |             | string<br/><small>Optional</small>                              |
| max-graceful-termination-sec          |             | string<br/><small>Optional</small>                              |
| max-node-provision-time               |             | string<br/><small>Optional</small>                              |
| max-total-unready-percentage          |             | string<br/><small>Optional</small>                              |
| new-pod-scale-up-delay                |             | string<br/><small>Optional</small>                              |
| ok-total-unready-count                |             | string<br/><small>Optional</small>                              |
| scale-down-delay-after-add            |             | string<br/><small>Optional</small>                              |
| scale-down-delay-after-delete         |             | string<br/><small>Optional</small>                              |
| scale-down-delay-after-failure        |             | string<br/><small>Optional</small>                              |
| scale-down-unneeded-time              |             | string<br/><small>Optional</small>                              |
| scale-down-unready-time               |             | string<br/><small>Optional</small>                              |
| scale-down-utilization-threshold      |             | string<br/><small>Optional</small>                              |
| scan-interval                         |             | string<br/><small>Optional</small>                              |
| skip-nodes-with-local-storage         |             | string<br/><small>Optional</small>                              |
| skip-nodes-with-system-pods           |             | string<br/><small>Optional</small>                              |

<a id="ManagedClusterProperties_PublicNetworkAccess"></a>ManagedClusterProperties_PublicNetworkAccess
-----------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Value                | Description |
|----------------------|-------------|
| "Disabled"           |             |
| "Enabled"            |             |
| "SecuredByPerimeter" |             |

<a id="ManagedClusterProperties_PublicNetworkAccess_STATUS"></a>ManagedClusterProperties_PublicNetworkAccess_STATUS
-------------------------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Value                | Description |
|----------------------|-------------|
| "Disabled"           |             |
| "Enabled"            |             |
| "SecuredByPerimeter" |             |

<a id="ManagedClustersAgentPoolOperatorSpec"></a>ManagedClustersAgentPoolOperatorSpec
-------------------------------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfile"></a>ManagedClusterSecurityProfile
-----------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                  | Description | Type                                                                                                                        |
|---------------------------|-------------|-----------------------------------------------------------------------------------------------------------------------------|
| azureKeyVaultKms          |             | [AzureKeyVaultKms](#AzureKeyVaultKms)<br/><small>Optional</small>                                                           |
| customCATrustCertificates |             | ManagedClusterSecurityProfileCustomCATrustCertificates<br/><small>Optional</small>                                          |
| defender                  |             | [ManagedClusterSecurityProfileDefender](#ManagedClusterSecurityProfileDefender)<br/><small>Optional</small>                 |
| imageCleaner              |             | [ManagedClusterSecurityProfileImageCleaner](#ManagedClusterSecurityProfileImageCleaner)<br/><small>Optional</small>         |
| imageIntegrity            |             | [ManagedClusterSecurityProfileImageIntegrity](#ManagedClusterSecurityProfileImageIntegrity)<br/><small>Optional</small>     |
| nodeRestriction           |             | [ManagedClusterSecurityProfileNodeRestriction](#ManagedClusterSecurityProfileNodeRestriction)<br/><small>Optional</small>   |
| workloadIdentity          |             | [ManagedClusterSecurityProfileWorkloadIdentity](#ManagedClusterSecurityProfileWorkloadIdentity)<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfile_STATUS"></a>ManagedClusterSecurityProfile_STATUS
-------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                  | Description | Type                                                                                                                                      |
|---------------------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| azureKeyVaultKms          |             | [AzureKeyVaultKms_STATUS](#AzureKeyVaultKms_STATUS)<br/><small>Optional</small>                                                           |
| customCATrustCertificates |             | string[]<br/><small>Optional</small>                                                                                                      |
| defender                  |             | [ManagedClusterSecurityProfileDefender_STATUS](#ManagedClusterSecurityProfileDefender_STATUS)<br/><small>Optional</small>                 |
| imageCleaner              |             | [ManagedClusterSecurityProfileImageCleaner_STATUS](#ManagedClusterSecurityProfileImageCleaner_STATUS)<br/><small>Optional</small>         |
| imageIntegrity            |             | [ManagedClusterSecurityProfileImageIntegrity_STATUS](#ManagedClusterSecurityProfileImageIntegrity_STATUS)<br/><small>Optional</small>     |
| nodeRestriction           |             | [ManagedClusterSecurityProfileNodeRestriction_STATUS](#ManagedClusterSecurityProfileNodeRestriction_STATUS)<br/><small>Optional</small>   |
| workloadIdentity          |             | [ManagedClusterSecurityProfileWorkloadIdentity_STATUS](#ManagedClusterSecurityProfileWorkloadIdentity_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterServicePrincipalProfile"></a>ManagedClusterServicePrincipalProfile
---------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description | Type                                                                                                                                                   |
|----------|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| clientId |             | string<br/><small>Required</small>                                                                                                                     |
| secret   |             | [genruntime.SecretReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference)<br/><small>Optional</small> |

<a id="ManagedClusterServicePrincipalProfile_STATUS"></a>ManagedClusterServicePrincipalProfile_STATUS
-----------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description | Type                               |
|----------|-------------|------------------------------------|
| clientId |             | string<br/><small>Optional</small> |

<a id="ManagedClusterSKU"></a>ManagedClusterSKU
-----------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description | Type                                                                          |
|----------|-------------|-------------------------------------------------------------------------------|
| name     |             | [ManagedClusterSKU_Name](#ManagedClusterSKU_Name)<br/><small>Optional</small> |
| tier     |             | [ManagedClusterSKU_Tier](#ManagedClusterSKU_Tier)<br/><small>Optional</small> |

<a id="ManagedClusterSKU_STATUS"></a>ManagedClusterSKU_STATUS
-------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description | Type                                                                                        |
|----------|-------------|---------------------------------------------------------------------------------------------|
| name     |             | [ManagedClusterSKU_Name_STATUS](#ManagedClusterSKU_Name_STATUS)<br/><small>Optional</small> |
| tier     |             | [ManagedClusterSKU_Tier_STATUS](#ManagedClusterSKU_Tier_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterStorageProfile"></a>ManagedClusterStorageProfile
---------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property           | Description | Type                                                                                                                          |
|--------------------|-------------|-------------------------------------------------------------------------------------------------------------------------------|
| blobCSIDriver      |             | [ManagedClusterStorageProfileBlobCSIDriver](#ManagedClusterStorageProfileBlobCSIDriver)<br/><small>Optional</small>           |
| diskCSIDriver      |             | [ManagedClusterStorageProfileDiskCSIDriver](#ManagedClusterStorageProfileDiskCSIDriver)<br/><small>Optional</small>           |
| fileCSIDriver      |             | [ManagedClusterStorageProfileFileCSIDriver](#ManagedClusterStorageProfileFileCSIDriver)<br/><small>Optional</small>           |
| snapshotController |             | [ManagedClusterStorageProfileSnapshotController](#ManagedClusterStorageProfileSnapshotController)<br/><small>Optional</small> |

<a id="ManagedClusterStorageProfile_STATUS"></a>ManagedClusterStorageProfile_STATUS
-----------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property           | Description | Type                                                                                                                                        |
|--------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| blobCSIDriver      |             | [ManagedClusterStorageProfileBlobCSIDriver_STATUS](#ManagedClusterStorageProfileBlobCSIDriver_STATUS)<br/><small>Optional</small>           |
| diskCSIDriver      |             | [ManagedClusterStorageProfileDiskCSIDriver_STATUS](#ManagedClusterStorageProfileDiskCSIDriver_STATUS)<br/><small>Optional</small>           |
| fileCSIDriver      |             | [ManagedClusterStorageProfileFileCSIDriver_STATUS](#ManagedClusterStorageProfileFileCSIDriver_STATUS)<br/><small>Optional</small>           |
| snapshotController |             | [ManagedClusterStorageProfileSnapshotController_STATUS](#ManagedClusterStorageProfileSnapshotController_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterWindowsProfile"></a>ManagedClusterWindowsProfile
---------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property       | Description | Type                                                                                                                                                   |
|----------------|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| adminPassword  |             | [genruntime.SecretReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference)<br/><small>Optional</small> |
| adminUsername  |             | string<br/><small>Required</small>                                                                                                                     |
| enableCSIProxy |             | bool<br/><small>Optional</small>                                                                                                                       |
| gmsaProfile    |             | [WindowsGmsaProfile](#WindowsGmsaProfile)<br/><small>Optional</small>                                                                                  |
| licenseType    |             | [ManagedClusterWindowsProfile_LicenseType](#ManagedClusterWindowsProfile_LicenseType)<br/><small>Optional</small>                                      |

<a id="ManagedClusterWindowsProfile_STATUS"></a>ManagedClusterWindowsProfile_STATUS
-----------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property       | Description | Type                                                                                                                            |
|----------------|-------------|---------------------------------------------------------------------------------------------------------------------------------|
| adminUsername  |             | string<br/><small>Optional</small>                                                                                              |
| enableCSIProxy |             | bool<br/><small>Optional</small>                                                                                                |
| gmsaProfile    |             | [WindowsGmsaProfile_STATUS](#WindowsGmsaProfile_STATUS)<br/><small>Optional</small>                                             |
| licenseType    |             | [ManagedClusterWindowsProfile_LicenseType_STATUS](#ManagedClusterWindowsProfile_LicenseType_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterWorkloadAutoScalerProfile"></a>ManagedClusterWorkloadAutoScalerProfile
-------------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property              | Description | Type                                                                                                                                                      |
|-----------------------|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| keda                  |             | [ManagedClusterWorkloadAutoScalerProfileKeda](#ManagedClusterWorkloadAutoScalerProfileKeda)<br/><small>Optional</small>                                   |
| verticalPodAutoscaler |             | [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler)<br/><small>Optional</small> |

<a id="ManagedClusterWorkloadAutoScalerProfile_STATUS"></a>ManagedClusterWorkloadAutoScalerProfile_STATUS
---------------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property              | Description | Type                                                                                                                                                                    |
|-----------------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| keda                  |             | [ManagedClusterWorkloadAutoScalerProfileKeda_STATUS](#ManagedClusterWorkloadAutoScalerProfileKeda_STATUS)<br/><small>Optional</small>                                   |
| verticalPodAutoscaler |             | [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS)<br/><small>Optional</small> |

<a id="OSDiskType"></a>OSDiskType
---------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value       | Description |
|-------------|-------------|
| "Ephemeral" |             |
| "Managed"   |             |

<a id="OSDiskType_STATUS"></a>OSDiskType_STATUS
-----------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value       | Description |
|-------------|-------------|
| "Ephemeral" |             |
| "Managed"   |             |

<a id="OSSKU"></a>OSSKU
-----------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value           | Description |
|-----------------|-------------|
| "AzureLinux"    |             |
| "CBLMariner"    |             |
| "Mariner"       |             |
| "Ubuntu"        |             |
| "Windows2019"   |             |
| "Windows2022"   |             |
| "WindowsAnnual" |             |

<a id="OSSKU_STATUS"></a>OSSKU_STATUS
-------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value           | Description |
|-----------------|-------------|
| "AzureLinux"    |             |
| "CBLMariner"    |             |
| "Mariner"       |             |
| "Ubuntu"        |             |
| "Windows2019"   |             |
| "Windows2022"   |             |
| "WindowsAnnual" |             |

<a id="OSType"></a>OSType
-------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value     | Description |
|-----------|-------------|
| "Linux"   |             |
| "Windows" |             |

<a id="OSType_STATUS"></a>OSType_STATUS
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value     | Description |
|-----------|-------------|
| "Linux"   |             |
| "Windows" |             |

<a id="PowerState"></a>PowerState
---------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property | Description | Type                                                            |
|----------|-------------|-----------------------------------------------------------------|
| code     |             | [PowerState_Code](#PowerState_Code)<br/><small>Optional</small> |

<a id="PowerState_STATUS"></a>PowerState_STATUS
-----------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property | Description | Type                                                                          |
|----------|-------------|-------------------------------------------------------------------------------|
| code     |             | [PowerState_Code_STATUS](#PowerState_Code_STATUS)<br/><small>Optional</small> |

<a id="PrivateLinkResource"></a>PrivateLinkResource
---------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property        | Description | Type                                                                                                                                                       |
|-----------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| groupId         |             | string<br/><small>Optional</small>                                                                                                                         |
| name            |             | string<br/><small>Optional</small>                                                                                                                         |
| reference       |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| requiredMembers |             | string[]<br/><small>Optional</small>                                                                                                                       |
| type            |             | string<br/><small>Optional</small>                                                                                                                         |

<a id="PrivateLinkResource_STATUS"></a>PrivateLinkResource_STATUS
-----------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property             | Description | Type                                 |
|----------------------|-------------|--------------------------------------|
| groupId              |             | string<br/><small>Optional</small>   |
| id                   |             | string<br/><small>Optional</small>   |
| name                 |             | string<br/><small>Optional</small>   |
| privateLinkServiceID |             | string<br/><small>Optional</small>   |
| requiredMembers      |             | string[]<br/><small>Optional</small> |
| type                 |             | string<br/><small>Optional</small>   |

<a id="SafeguardsProfile"></a>SafeguardsProfile
-----------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property           | Description | Type                                                                            |
|--------------------|-------------|---------------------------------------------------------------------------------|
| excludedNamespaces |             | string[]<br/><small>Optional</small>                                            |
| level              |             | [SafeguardsProfile_Level](#SafeguardsProfile_Level)<br/><small>Required</small> |
| version            |             | string<br/><small>Optional</small>                                              |

<a id="SafeguardsProfile_STATUS"></a>SafeguardsProfile_STATUS
-------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                 | Description | Type                                                                                          |
|--------------------------|-------------|-----------------------------------------------------------------------------------------------|
| excludedNamespaces       |             | string[]<br/><small>Optional</small>                                                          |
| level                    |             | [SafeguardsProfile_Level_STATUS](#SafeguardsProfile_Level_STATUS)<br/><small>Optional</small> |
| systemExcludedNamespaces |             | string[]<br/><small>Optional</small>                                                          |
| version                  |             | string<br/><small>Optional</small>                                                            |

<a id="ScaleDownMode"></a>ScaleDownMode
---------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

<a id="ScaleDownMode_STATUS"></a>ScaleDownMode_STATUS
-----------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

<a id="ScaleSetEvictionPolicy"></a>ScaleSetEvictionPolicy
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

<a id="ScaleSetEvictionPolicy_STATUS"></a>ScaleSetEvictionPolicy_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

<a id="ScaleSetPriority"></a>ScaleSetPriority
---------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value     | Description |
|-----------|-------------|
| "Regular" |             |
| "Spot"    |             |

<a id="ScaleSetPriority_STATUS"></a>ScaleSetPriority_STATUS
-----------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value     | Description |
|-----------|-------------|
| "Regular" |             |
| "Spot"    |             |

<a id="ServiceMeshProfile"></a>ServiceMeshProfile
-------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description | Type                                                                            |
|----------|-------------|---------------------------------------------------------------------------------|
| istio    |             | [IstioServiceMesh](#IstioServiceMesh)<br/><small>Optional</small>               |
| mode     |             | [ServiceMeshProfile_Mode](#ServiceMeshProfile_Mode)<br/><small>Required</small> |

<a id="ServiceMeshProfile_STATUS"></a>ServiceMeshProfile_STATUS
---------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description | Type                                                                                          |
|----------|-------------|-----------------------------------------------------------------------------------------------|
| istio    |             | [IstioServiceMesh_STATUS](#IstioServiceMesh_STATUS)<br/><small>Optional</small>               |
| mode     |             | [ServiceMeshProfile_Mode_STATUS](#ServiceMeshProfile_Mode_STATUS)<br/><small>Optional</small> |

<a id="SystemData_STATUS"></a>SystemData_STATUS
-----------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property           | Description | Type                                                                                                      |
|--------------------|-------------|-----------------------------------------------------------------------------------------------------------|
| createdAt          |             | string<br/><small>Optional</small>                                                                        |
| createdBy          |             | string<br/><small>Optional</small>                                                                        |
| createdByType      |             | [SystemData_CreatedByType_STATUS](#SystemData_CreatedByType_STATUS)<br/><small>Optional</small>           |
| lastModifiedAt     |             | string<br/><small>Optional</small>                                                                        |
| lastModifiedBy     |             | string<br/><small>Optional</small>                                                                        |
| lastModifiedByType |             | [SystemData_LastModifiedByType_STATUS](#SystemData_LastModifiedByType_STATUS)<br/><small>Optional</small> |

<a id="UserAssignedIdentity"></a>UserAssignedIdentity
-----------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec), and [ManagedClusterPodIdentity](#ManagedClusterPodIdentity).

| Property          | Description | Type                                                                                                                                                       |
|-------------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| clientId          |             | string<br/><small>Optional</small>                                                                                                                         |
| objectId          |             | string<br/><small>Optional</small>                                                                                                                         |
| resourceReference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

<a id="UserAssignedIdentity_STATUS"></a>UserAssignedIdentity_STATUS
-------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS), [ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS), [ManagedClusterIngressProfileWebAppRouting_STATUS](#ManagedClusterIngressProfileWebAppRouting_STATUS), and [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Property   | Description | Type                               |
|------------|-------------|------------------------------------|
| clientId   |             | string<br/><small>Optional</small> |
| objectId   |             | string<br/><small>Optional</small> |
| resourceId |             | string<br/><small>Optional</small> |

<a id="VirtualMachineNodes"></a>VirtualMachineNodes
---------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property | Description | Type                               |
|----------|-------------|------------------------------------|
| count    |             | int<br/><small>Optional</small>    |
| size     |             | string<br/><small>Optional</small> |

<a id="VirtualMachineNodes_STATUS"></a>VirtualMachineNodes_STATUS
-----------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property | Description | Type                               |
|----------|-------------|------------------------------------|
| count    |             | int<br/><small>Optional</small>    |
| size     |             | string<br/><small>Optional</small> |

<a id="VirtualMachinesProfile"></a>VirtualMachinesProfile
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property | Description | Type                                                      |
|----------|-------------|-----------------------------------------------------------|
| scale    |             | [ScaleProfile](#ScaleProfile)<br/><small>Optional</small> |

<a id="VirtualMachinesProfile_STATUS"></a>VirtualMachinesProfile_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property | Description | Type                                                                    |
|----------|-------------|-------------------------------------------------------------------------|
| scale    |             | [ScaleProfile_STATUS](#ScaleProfile_STATUS)<br/><small>Optional</small> |

<a id="WorkloadRuntime"></a>WorkloadRuntime
-------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value                 | Description |
|-----------------------|-------------|
| "KataMshvVmIsolation" |             |
| "OCIContainer"        |             |
| "WasmWasi"            |             |

<a id="WorkloadRuntime_STATUS"></a>WorkloadRuntime_STATUS
---------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value                 | Description |
|-----------------------|-------------|
| "KataMshvVmIsolation" |             |
| "OCIContainer"        |             |
| "WasmWasi"            |             |

<a id="AgentPoolSSHAccess"></a>AgentPoolSSHAccess
-------------------------------------------------

Used by: [AgentPoolSecurityProfile](#AgentPoolSecurityProfile).

| Value       | Description |
|-------------|-------------|
| "Disabled"  |             |
| "LocalUser" |             |

<a id="AgentPoolSSHAccess_STATUS"></a>AgentPoolSSHAccess_STATUS
---------------------------------------------------------------

Used by: [AgentPoolSecurityProfile_STATUS](#AgentPoolSecurityProfile_STATUS).

| Value       | Description |
|-------------|-------------|
| "Disabled"  |             |
| "LocalUser" |             |

<a id="AzureKeyVaultKms"></a>AzureKeyVaultKms
---------------------------------------------

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property                  | Description | Type                                                                                                                                                       |
|---------------------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| enabled                   |             | bool<br/><small>Optional</small>                                                                                                                           |
| keyId                     |             | string<br/><small>Optional</small>                                                                                                                         |
| keyVaultNetworkAccess     |             | [AzureKeyVaultKms_KeyVaultNetworkAccess](#AzureKeyVaultKms_KeyVaultNetworkAccess)<br/><small>Optional</small>                                              |
| keyVaultResourceReference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

<a id="AzureKeyVaultKms_STATUS"></a>AzureKeyVaultKms_STATUS
-----------------------------------------------------------

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property              | Description | Type                                                                                                                        |
|-----------------------|-------------|-----------------------------------------------------------------------------------------------------------------------------|
| enabled               |             | bool<br/><small>Optional</small>                                                                                            |
| keyId                 |             | string<br/><small>Optional</small>                                                                                          |
| keyVaultNetworkAccess |             | [AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS](#AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS)<br/><small>Optional</small> |
| keyVaultResourceId    |             | string<br/><small>Optional</small>                                                                                          |

<a id="ContainerServiceNetworkProfile_KubeProxyConfig"></a>ContainerServiceNetworkProfile_KubeProxyConfig
---------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property   | Description | Type                                                                                                                                                |
|------------|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| enabled    |             | bool<br/><small>Optional</small>                                                                                                                    |
| ipvsConfig |             | [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig)<br/><small>Optional</small> |
| mode       |             | [ContainerServiceNetworkProfile_KubeProxyConfig_Mode](#ContainerServiceNetworkProfile_KubeProxyConfig_Mode)<br/><small>Optional</small>             |

<a id="ContainerServiceNetworkProfile_KubeProxyConfig_STATUS"></a>ContainerServiceNetworkProfile_KubeProxyConfig_STATUS
-----------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property   | Description | Type                                                                                                                                                              |
|------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| enabled    |             | bool<br/><small>Optional</small>                                                                                                                                  |
| ipvsConfig |             | [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS)<br/><small>Optional</small> |
| mode       |             | [ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS)<br/><small>Optional</small>             |

<a id="ContainerServiceNetworkProfile_OutboundType"></a>ContainerServiceNetworkProfile_OutboundType
---------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value                    | Description |
|--------------------------|-------------|
| "loadBalancer"           |             |
| "managedNATGateway"      |             |
| "userAssignedNATGateway" |             |
| "userDefinedRouting"     |             |

<a id="ContainerServiceNetworkProfile_OutboundType_STATUS"></a>ContainerServiceNetworkProfile_OutboundType_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value                    | Description |
|--------------------------|-------------|
| "loadBalancer"           |             |
| "managedNATGateway"      |             |
| "userAssignedNATGateway" |             |
| "userDefinedRouting"     |             |

<a id="ContainerServiceSshConfiguration"></a>ContainerServiceSshConfiguration
-----------------------------------------------------------------------------

Used by: [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile).

| Property   | Description | Type                                                                                        |
|------------|-------------|---------------------------------------------------------------------------------------------|
| publicKeys |             | [ContainerServiceSshPublicKey[]](#ContainerServiceSshPublicKey)<br/><small>Required</small> |

<a id="ContainerServiceSshConfiguration_STATUS"></a>ContainerServiceSshConfiguration_STATUS
-------------------------------------------------------------------------------------------

Used by: [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS).

| Property   | Description | Type                                                                                                      |
|------------|-------------|-----------------------------------------------------------------------------------------------------------|
| publicKeys |             | [ContainerServiceSshPublicKey_STATUS[]](#ContainerServiceSshPublicKey_STATUS)<br/><small>Optional</small> |

<a id="DelegatedResource"></a>DelegatedResource
-----------------------------------------------

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity).

| Property          | Description | Type                                                                                                                                                       |
|-------------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| location          |             | string<br/><small>Optional</small>                                                                                                                         |
| referralResource  |             | string<br/><small>Optional</small>                                                                                                                         |
| resourceReference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| tenantId          |             | string<br/><small>Optional</small>                                                                                                                         |

<a id="DelegatedResource_STATUS"></a>DelegatedResource_STATUS
-------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS).

| Property         | Description | Type                               |
|------------------|-------------|------------------------------------|
| location         |             | string<br/><small>Optional</small> |
| referralResource |             | string<br/><small>Optional</small> |
| resourceId       |             | string<br/><small>Optional</small> |
| tenantId         |             | string<br/><small>Optional</small> |

<a id="Expander"></a>Expander
-----------------------------

Used by: [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile).

| Value         | Description |
|---------------|-------------|
| "least-waste" |             |
| "most-pods"   |             |
| "priority"    |             |
| "random"      |             |

<a id="Expander_STATUS"></a>Expander_STATUS
-------------------------------------------

Used by: [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS).

| Value         | Description |
|---------------|-------------|
| "least-waste" |             |
| "most-pods"   |             |
| "priority"    |             |
| "random"      |             |

<a id="ExtendedLocationType"></a>ExtendedLocationType
-----------------------------------------------------

Used by: [ExtendedLocation](#ExtendedLocation).

| Value      | Description |
|------------|-------------|
| "EdgeZone" |             |

<a id="ExtendedLocationType_STATUS"></a>ExtendedLocationType_STATUS
-------------------------------------------------------------------

Used by: [ExtendedLocation_STATUS](#ExtendedLocation_STATUS).

| Value      | Description |
|------------|-------------|
| "EdgeZone" |             |

<a id="IpFamily"></a>IpFamily
-----------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value  | Description |
|--------|-------------|
| "IPv4" |             |
| "IPv6" |             |

<a id="IpFamily_STATUS"></a>IpFamily_STATUS
-------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value  | Description |
|--------|-------------|
| "IPv4" |             |
| "IPv6" |             |

<a id="IPTag"></a>IPTag
-----------------------

Used by: [AgentPoolNetworkProfile](#AgentPoolNetworkProfile).

| Property  | Description | Type                               |
|-----------|-------------|------------------------------------|
| ipTagType |             | string<br/><small>Optional</small> |
| tag       |             | string<br/><small>Optional</small> |

<a id="IPTag_STATUS"></a>IPTag_STATUS
-------------------------------------

Used by: [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS).

| Property  | Description | Type                               |
|-----------|-------------|------------------------------------|
| ipTagType |             | string<br/><small>Optional</small> |
| tag       |             | string<br/><small>Optional</small> |

<a id="IstioServiceMesh"></a>IstioServiceMesh
---------------------------------------------

Used by: [ServiceMeshProfile](#ServiceMeshProfile).

| Property             | Description | Type                                                                                |
|----------------------|-------------|-------------------------------------------------------------------------------------|
| certificateAuthority |             | [IstioCertificateAuthority](#IstioCertificateAuthority)<br/><small>Optional</small> |
| components           |             | [IstioComponents](#IstioComponents)<br/><small>Optional</small>                     |
| revisions            |             | string[]<br/><small>Optional</small>                                                |

<a id="IstioServiceMesh_STATUS"></a>IstioServiceMesh_STATUS
-----------------------------------------------------------

Used by: [ServiceMeshProfile_STATUS](#ServiceMeshProfile_STATUS).

| Property             | Description | Type                                                                                              |
|----------------------|-------------|---------------------------------------------------------------------------------------------------|
| certificateAuthority |             | [IstioCertificateAuthority_STATUS](#IstioCertificateAuthority_STATUS)<br/><small>Optional</small> |
| components           |             | [IstioComponents_STATUS](#IstioComponents_STATUS)<br/><small>Optional</small>                     |
| revisions            |             | string[]<br/><small>Optional</small>                                                              |

<a id="LoadBalancerSku"></a>LoadBalancerSku
-------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value      | Description |
|------------|-------------|
| "basic"    |             |
| "standard" |             |

<a id="LoadBalancerSku_STATUS"></a>LoadBalancerSku_STATUS
---------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value      | Description |
|------------|-------------|
| "basic"    |             |
| "standard" |             |

<a id="ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel"></a>ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel
-----------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile).

| Value           | Description |
|-----------------|-------------|
| "NodeImage"     |             |
| "None"          |             |
| "SecurityPatch" |             |
| "Unmanaged"     |             |

<a id="ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS"></a>ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS
-------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS).

| Value           | Description |
|-----------------|-------------|
| "NodeImage"     |             |
| "None"          |             |
| "SecurityPatch" |             |
| "Unmanaged"     |             |

<a id="ManagedClusterAutoUpgradeProfile_UpgradeChannel"></a>ManagedClusterAutoUpgradeProfile_UpgradeChannel
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile).

| Value        | Description |
|--------------|-------------|
| "node-image" |             |
| "none"       |             |
| "patch"      |             |
| "rapid"      |             |
| "stable"     |             |

<a id="ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS"></a>ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS
-------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS).

| Value        | Description |
|--------------|-------------|
| "node-image" |             |
| "none"       |             |
| "patch"      |             |
| "rapid"      |             |
| "stable"     |             |

<a id="ManagedClusterAzureMonitorProfileLogs"></a>ManagedClusterAzureMonitorProfileLogs
---------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfile](#ManagedClusterAzureMonitorProfile).

| Property          | Description | Type                                                                                                                                  |
|-------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------|
| appMonitoring     |             | [ManagedClusterAzureMonitorProfileAppMonitoring](#ManagedClusterAzureMonitorProfileAppMonitoring)<br/><small>Optional</small>         |
| containerInsights |             | [ManagedClusterAzureMonitorProfileContainerInsights](#ManagedClusterAzureMonitorProfileContainerInsights)<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileLogs_STATUS"></a>ManagedClusterAzureMonitorProfileLogs_STATUS
-----------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfile_STATUS](#ManagedClusterAzureMonitorProfile_STATUS).

| Property          | Description | Type                                                                                                                                                |
|-------------------|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| appMonitoring     |             | [ManagedClusterAzureMonitorProfileAppMonitoring_STATUS](#ManagedClusterAzureMonitorProfileAppMonitoring_STATUS)<br/><small>Optional</small>         |
| containerInsights |             | [ManagedClusterAzureMonitorProfileContainerInsights_STATUS](#ManagedClusterAzureMonitorProfileContainerInsights_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileMetrics"></a>ManagedClusterAzureMonitorProfileMetrics
---------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfile](#ManagedClusterAzureMonitorProfile).

| Property                          | Description | Type                                                                                                                                                                  |
|-----------------------------------|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| appMonitoringOpenTelemetryMetrics |             | [ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics](#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics)<br/><small>Optional</small> |
| enabled                           |             | bool<br/><small>Required</small>                                                                                                                                      |
| kubeStateMetrics                  |             | [ManagedClusterAzureMonitorProfileKubeStateMetrics](#ManagedClusterAzureMonitorProfileKubeStateMetrics)<br/><small>Optional</small>                                   |

<a id="ManagedClusterAzureMonitorProfileMetrics_STATUS"></a>ManagedClusterAzureMonitorProfileMetrics_STATUS
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfile_STATUS](#ManagedClusterAzureMonitorProfile_STATUS).

| Property                          | Description | Type                                                                                                                                                                                |
|-----------------------------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| appMonitoringOpenTelemetryMetrics |             | [ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics_STATUS](#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics_STATUS)<br/><small>Optional</small> |
| enabled                           |             | bool<br/><small>Optional</small>                                                                                                                                                    |
| kubeStateMetrics                  |             | [ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS](#ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS)<br/><small>Optional</small>                                   |

<a id="ManagedClusterCostAnalysis"></a>ManagedClusterCostAnalysis
-----------------------------------------------------------------

Used by: [ManagedClusterMetricsProfile](#ManagedClusterMetricsProfile).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterCostAnalysis_STATUS"></a>ManagedClusterCostAnalysis_STATUS
-------------------------------------------------------------------------------

Used by: [ManagedClusterMetricsProfile_STATUS](#ManagedClusterMetricsProfile_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterIdentity_Type"></a>ManagedClusterIdentity_Type
-------------------------------------------------------------------

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

<a id="ManagedClusterIdentity_Type_STATUS"></a>ManagedClusterIdentity_Type_STATUS
---------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

<a id="ManagedClusterIdentity_UserAssignedIdentities_STATUS"></a>ManagedClusterIdentity_UserAssignedIdentities_STATUS
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS).

| Property    | Description | Type                               |
|-------------|-------------|------------------------------------|
| clientId    |             | string<br/><small>Optional</small> |
| principalId |             | string<br/><small>Optional</small> |

<a id="ManagedClusterIngressProfileWebAppRouting"></a>ManagedClusterIngressProfileWebAppRouting
-----------------------------------------------------------------------------------------------

Used by: [ManagedClusterIngressProfile](#ManagedClusterIngressProfile).

| Property                  | Description | Type                                                                                                                                                         |
|---------------------------|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| dnsZoneResourceReferences |             | [genruntime.ResourceReference[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| enabled                   |             | bool<br/><small>Optional</small>                                                                                                                             |

<a id="ManagedClusterIngressProfileWebAppRouting_STATUS"></a>ManagedClusterIngressProfileWebAppRouting_STATUS
-------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterIngressProfile_STATUS](#ManagedClusterIngressProfile_STATUS).

| Property           | Description | Type                                                                                    |
|--------------------|-------------|-----------------------------------------------------------------------------------------|
| dnsZoneResourceIds |             | string[]<br/><small>Optional</small>                                                    |
| enabled            |             | bool<br/><small>Optional</small>                                                        |
| identity           |             | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile"></a>ManagedClusterLoadBalancerProfile
-------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property                            | Description | Type                                                                                                                                      |
|-------------------------------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts              |             | int<br/><small>Optional</small>                                                                                                           |
| backendPoolType                     |             | [ManagedClusterLoadBalancerProfile_BackendPoolType](#ManagedClusterLoadBalancerProfile_BackendPoolType)<br/><small>Optional</small>       |
| effectiveOutboundIPs                |             | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small>                                                                     |
| enableMultipleStandardLoadBalancers |             | bool<br/><small>Optional</small>                                                                                                          |
| idleTimeoutInMinutes                |             | int<br/><small>Optional</small>                                                                                                           |
| managedOutboundIPs                  |             | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs)<br/><small>Optional</small> |
| outboundIPPrefixes                  |             | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes)<br/><small>Optional</small> |
| outboundIPs                         |             | [ManagedClusterLoadBalancerProfile_OutboundIPs](#ManagedClusterLoadBalancerProfile_OutboundIPs)<br/><small>Optional</small>               |

<a id="ManagedClusterLoadBalancerProfile_STATUS"></a>ManagedClusterLoadBalancerProfile_STATUS
---------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property                            | Description | Type                                                                                                                                                    |
|-------------------------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts              |             | int<br/><small>Optional</small>                                                                                                                         |
| backendPoolType                     |             | [ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS](#ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS)<br/><small>Optional</small>       |
| effectiveOutboundIPs                |             | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small>                                                                     |
| enableMultipleStandardLoadBalancers |             | bool<br/><small>Optional</small>                                                                                                                        |
| idleTimeoutInMinutes                |             | int<br/><small>Optional</small>                                                                                                                         |
| managedOutboundIPs                  |             | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS)<br/><small>Optional</small> |
| outboundIPPrefixes                  |             | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS)<br/><small>Optional</small> |
| outboundIPs                         |             | [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS)<br/><small>Optional</small>               |

<a id="ManagedClusterNATGatewayProfile"></a>ManagedClusterNATGatewayProfile
---------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property                 | Description | Type                                                                                                          |
|--------------------------|-------------|---------------------------------------------------------------------------------------------------------------|
| effectiveOutboundIPs     |             | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small>                                         |
| idleTimeoutInMinutes     |             | int<br/><small>Optional</small>                                                                               |
| managedOutboundIPProfile |             | [ManagedClusterManagedOutboundIPProfile](#ManagedClusterManagedOutboundIPProfile)<br/><small>Optional</small> |

<a id="ManagedClusterNATGatewayProfile_STATUS"></a>ManagedClusterNATGatewayProfile_STATUS
-----------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property                 | Description | Type                                                                                                                        |
|--------------------------|-------------|-----------------------------------------------------------------------------------------------------------------------------|
| effectiveOutboundIPs     |             | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small>                                         |
| idleTimeoutInMinutes     |             | int<br/><small>Optional</small>                                                                                             |
| managedOutboundIPProfile |             | [ManagedClusterManagedOutboundIPProfile_STATUS](#ManagedClusterManagedOutboundIPProfile_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterNodeProvisioningProfile_Mode"></a>ManagedClusterNodeProvisioningProfile_Mode
-------------------------------------------------------------------------------------------------

Used by: [ManagedClusterNodeProvisioningProfile](#ManagedClusterNodeProvisioningProfile).

| Value    | Description |
|----------|-------------|
| "Auto"   |             |
| "Manual" |             |

<a id="ManagedClusterNodeProvisioningProfile_Mode_STATUS"></a>ManagedClusterNodeProvisioningProfile_Mode_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterNodeProvisioningProfile_STATUS](#ManagedClusterNodeProvisioningProfile_STATUS).

| Value    | Description |
|----------|-------------|
| "Auto"   |             |
| "Manual" |             |

<a id="ManagedClusterNodeResourceGroupProfile_RestrictionLevel"></a>ManagedClusterNodeResourceGroupProfile_RestrictionLevel
---------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterNodeResourceGroupProfile](#ManagedClusterNodeResourceGroupProfile).

| Value          | Description |
|----------------|-------------|
| "ReadOnly"     |             |
| "Unrestricted" |             |

<a id="ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS"></a>ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS
-----------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterNodeResourceGroupProfile_STATUS](#ManagedClusterNodeResourceGroupProfile_STATUS).

| Value          | Description |
|----------------|-------------|
| "ReadOnly"     |             |
| "Unrestricted" |             |

<a id="ManagedClusterOperatorConfigMaps"></a>ManagedClusterOperatorConfigMaps
-----------------------------------------------------------------------------

Used by: [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec).

| Property          | Description                                                                                                   | Type                                                                                                                                                             |
|-------------------|---------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| oidcIssuerProfile | indicates where the OIDCIssuerProfile config map should be placed. If omitted, no config map will be created. | [genruntime.ConfigMapDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapDestination)<br/><small>Optional</small> |

<a id="ManagedClusterOperatorSecrets"></a>ManagedClusterOperatorSecrets
-----------------------------------------------------------------------

Used by: [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec).

| Property         | Description                                                                                                            | Type                                                                                                                                                       |
|------------------|------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| adminCredentials | indicates where the AdminCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure. | [genruntime.SecretDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination)<br/><small>Optional</small> |
| userCredentials  | indicates where the UserCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure.  | [genruntime.SecretDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentity"></a>ManagedClusterPodIdentity
---------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile).

| Property        | Description | Type                                                                      |
|-----------------|-------------|---------------------------------------------------------------------------|
| bindingSelector |             | string<br/><small>Optional</small>                                        |
| identity        |             | [UserAssignedIdentity](#UserAssignedIdentity)<br/><small>Required</small> |
| name            |             | string<br/><small>Required</small>                                        |
| namespace       |             | string<br/><small>Required</small>                                        |

<a id="ManagedClusterPodIdentity_STATUS"></a>ManagedClusterPodIdentity_STATUS
-----------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS).

| Property          | Description | Type                                                                                                                                  |
|-------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------|
| bindingSelector   |             | string<br/><small>Optional</small>                                                                                                    |
| identity          |             | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small>                                               |
| name              |             | string<br/><small>Optional</small>                                                                                                    |
| namespace         |             | string<br/><small>Optional</small>                                                                                                    |
| provisioningInfo  |             | [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS)<br/><small>Optional</small>   |
| provisioningState |             | [ManagedClusterPodIdentity_ProvisioningState_STATUS](#ManagedClusterPodIdentity_ProvisioningState_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentityException"></a>ManagedClusterPodIdentityException
---------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile).

| Property  | Description | Type                                          |
|-----------|-------------|-----------------------------------------------|
| name      |             | string<br/><small>Required</small>            |
| namespace |             | string<br/><small>Required</small>            |
| podLabels |             | map[string]string<br/><small>Required</small> |

<a id="ManagedClusterPodIdentityException_STATUS"></a>ManagedClusterPodIdentityException_STATUS
-----------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS).

| Property  | Description | Type                                          |
|-----------|-------------|-----------------------------------------------|
| name      |             | string<br/><small>Optional</small>            |
| namespace |             | string<br/><small>Optional</small>            |
| podLabels |             | map[string]string<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfileDefender"></a>ManagedClusterSecurityProfileDefender
---------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property                               | Description | Type                                                                                                                                                       |
|----------------------------------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| logAnalyticsWorkspaceResourceReference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| securityMonitoring                     |             | [ManagedClusterSecurityProfileDefenderSecurityMonitoring](#ManagedClusterSecurityProfileDefenderSecurityMonitoring)<br/><small>Optional</small>            |

<a id="ManagedClusterSecurityProfileDefender_STATUS"></a>ManagedClusterSecurityProfileDefender_STATUS
-----------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property                        | Description | Type                                                                                                                                                          |
|---------------------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| logAnalyticsWorkspaceResourceId |             | string<br/><small>Optional</small>                                                                                                                            |
| securityMonitoring              |             | [ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS](#ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfileImageCleaner"></a>ManagedClusterSecurityProfileImageCleaner
-----------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property      | Description | Type                             |
|---------------|-------------|----------------------------------|
| enabled       |             | bool<br/><small>Optional</small> |
| intervalHours |             | int<br/><small>Optional</small>  |

<a id="ManagedClusterSecurityProfileImageCleaner_STATUS"></a>ManagedClusterSecurityProfileImageCleaner_STATUS
-------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property      | Description | Type                             |
|---------------|-------------|----------------------------------|
| enabled       |             | bool<br/><small>Optional</small> |
| intervalHours |             | int<br/><small>Optional</small>  |

<a id="ManagedClusterSecurityProfileImageIntegrity"></a>ManagedClusterSecurityProfileImageIntegrity
---------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfileImageIntegrity_STATUS"></a>ManagedClusterSecurityProfileImageIntegrity_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfileNodeRestriction"></a>ManagedClusterSecurityProfileNodeRestriction
-----------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfileNodeRestriction_STATUS"></a>ManagedClusterSecurityProfileNodeRestriction_STATUS
-------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfileWorkloadIdentity"></a>ManagedClusterSecurityProfileWorkloadIdentity
-------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfileWorkloadIdentity_STATUS"></a>ManagedClusterSecurityProfileWorkloadIdentity_STATUS
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterSKU_Name"></a>ManagedClusterSKU_Name
---------------------------------------------------------

Used by: [ManagedClusterSKU](#ManagedClusterSKU).

| Value  | Description |
|--------|-------------|
| "Base" |             |

<a id="ManagedClusterSKU_Name_STATUS"></a>ManagedClusterSKU_Name_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS).

| Value  | Description |
|--------|-------------|
| "Base" |             |

<a id="ManagedClusterSKU_Tier"></a>ManagedClusterSKU_Tier
---------------------------------------------------------

Used by: [ManagedClusterSKU](#ManagedClusterSKU).

| Value      | Description |
|------------|-------------|
| "Free"     |             |
| "Premium"  |             |
| "Standard" |             |

<a id="ManagedClusterSKU_Tier_STATUS"></a>ManagedClusterSKU_Tier_STATUS
-----------------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS).

| Value      | Description |
|------------|-------------|
| "Free"     |             |
| "Premium"  |             |
| "Standard" |             |

<a id="ManagedClusterStorageProfileBlobCSIDriver"></a>ManagedClusterStorageProfileBlobCSIDriver
-----------------------------------------------------------------------------------------------

Used by: [ManagedClusterStorageProfile](#ManagedClusterStorageProfile).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterStorageProfileBlobCSIDriver_STATUS"></a>ManagedClusterStorageProfileBlobCSIDriver_STATUS
-------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterStorageProfileDiskCSIDriver"></a>ManagedClusterStorageProfileDiskCSIDriver
-----------------------------------------------------------------------------------------------

Used by: [ManagedClusterStorageProfile](#ManagedClusterStorageProfile).

| Property | Description | Type                               |
|----------|-------------|------------------------------------|
| enabled  |             | bool<br/><small>Optional</small>   |
| version  |             | string<br/><small>Optional</small> |

<a id="ManagedClusterStorageProfileDiskCSIDriver_STATUS"></a>ManagedClusterStorageProfileDiskCSIDriver_STATUS
-------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS).

| Property | Description | Type                               |
|----------|-------------|------------------------------------|
| enabled  |             | bool<br/><small>Optional</small>   |
| version  |             | string<br/><small>Optional</small> |

<a id="ManagedClusterStorageProfileFileCSIDriver"></a>ManagedClusterStorageProfileFileCSIDriver
-----------------------------------------------------------------------------------------------

Used by: [ManagedClusterStorageProfile](#ManagedClusterStorageProfile).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterStorageProfileFileCSIDriver_STATUS"></a>ManagedClusterStorageProfileFileCSIDriver_STATUS
-------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterStorageProfileSnapshotController"></a>ManagedClusterStorageProfileSnapshotController
---------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterStorageProfile](#ManagedClusterStorageProfile).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterStorageProfileSnapshotController_STATUS"></a>ManagedClusterStorageProfileSnapshotController_STATUS
-----------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterWindowsProfile_LicenseType"></a>ManagedClusterWindowsProfile_LicenseType
---------------------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "Windows_Server" |             |

<a id="ManagedClusterWindowsProfile_LicenseType_STATUS"></a>ManagedClusterWindowsProfile_LicenseType_STATUS
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "Windows_Server" |             |

<a id="ManagedClusterWorkloadAutoScalerProfileKeda"></a>ManagedClusterWorkloadAutoScalerProfileKeda
---------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfile](#ManagedClusterWorkloadAutoScalerProfile).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Required</small> |

<a id="ManagedClusterWorkloadAutoScalerProfileKeda_STATUS"></a>ManagedClusterWorkloadAutoScalerProfileKeda_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfile_STATUS](#ManagedClusterWorkloadAutoScalerProfile_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler"></a>ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler
-------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfile](#ManagedClusterWorkloadAutoScalerProfile).

| Property         | Description | Type                                                                                                                                                                                        |
|------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| addonAutoscaling |             | [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling)<br/><small>Optional</small> |
| enabled          |             | bool<br/><small>Required</small>                                                                                                                                                            |

<a id="ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS"></a>ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS
---------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfile_STATUS](#ManagedClusterWorkloadAutoScalerProfile_STATUS).

| Property         | Description | Type                                                                                                                                                                                                      |
|------------------|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| addonAutoscaling |             | [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling_STATUS](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling_STATUS)<br/><small>Optional</small> |
| enabled          |             | bool<br/><small>Optional</small>                                                                                                                                                                          |

<a id="NetworkDataplane"></a>NetworkDataplane
---------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "cilium" |             |

<a id="NetworkDataplane_STATUS"></a>NetworkDataplane_STATUS
-----------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "cilium" |             |

<a id="NetworkMode"></a>NetworkMode
-----------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value         | Description |
|---------------|-------------|
| "bridge"      |             |
| "transparent" |             |

<a id="NetworkMode_STATUS"></a>NetworkMode_STATUS
-------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value         | Description |
|---------------|-------------|
| "bridge"      |             |
| "transparent" |             |

<a id="NetworkMonitoring"></a>NetworkMonitoring
-----------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="NetworkMonitoring_STATUS"></a>NetworkMonitoring_STATUS
-------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="NetworkPlugin"></a>NetworkPlugin
---------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value     | Description |
|-----------|-------------|
| "azure"   |             |
| "kubenet" |             |
| "none"    |             |

<a id="NetworkPlugin_STATUS"></a>NetworkPlugin_STATUS
-----------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value     | Description |
|-----------|-------------|
| "azure"   |             |
| "kubenet" |             |
| "none"    |             |

<a id="NetworkPluginMode"></a>NetworkPluginMode
-----------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value     | Description |
|-----------|-------------|
| "overlay" |             |

<a id="NetworkPluginMode_STATUS"></a>NetworkPluginMode_STATUS
-------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value     | Description |
|-----------|-------------|
| "overlay" |             |

<a id="NetworkPolicy"></a>NetworkPolicy
---------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "calico" |             |
| "cilium" |             |
| "none"   |             |

<a id="NetworkPolicy_STATUS"></a>NetworkPolicy_STATUS
-----------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "calico" |             |
| "cilium" |             |
| "none"   |             |

<a id="PortRange"></a>PortRange
-------------------------------

Used by: [AgentPoolNetworkProfile](#AgentPoolNetworkProfile).

| Property  | Description | Type                                                                  |
|-----------|-------------|-----------------------------------------------------------------------|
| portEnd   |             | int<br/><small>Optional</small>                                       |
| portStart |             | int<br/><small>Optional</small>                                       |
| protocol  |             | [PortRange_Protocol](#PortRange_Protocol)<br/><small>Optional</small> |

<a id="PortRange_STATUS"></a>PortRange_STATUS
---------------------------------------------

Used by: [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS).

| Property  | Description | Type                                                                                |
|-----------|-------------|-------------------------------------------------------------------------------------|
| portEnd   |             | int<br/><small>Optional</small>                                                     |
| portStart |             | int<br/><small>Optional</small>                                                     |
| protocol  |             | [PortRange_Protocol_STATUS](#PortRange_Protocol_STATUS)<br/><small>Optional</small> |

<a id="PowerState_Code"></a>PowerState_Code
-------------------------------------------

Used by: [PowerState](#PowerState).

| Value     | Description |
|-----------|-------------|
| "Running" |             |
| "Stopped" |             |

<a id="PowerState_Code_STATUS"></a>PowerState_Code_STATUS
---------------------------------------------------------

Used by: [PowerState_STATUS](#PowerState_STATUS).

| Value     | Description |
|-----------|-------------|
| "Running" |             |
| "Stopped" |             |

<a id="SafeguardsProfile_Level"></a>SafeguardsProfile_Level
-----------------------------------------------------------

Used by: [SafeguardsProfile](#SafeguardsProfile).

| Value         | Description |
|---------------|-------------|
| "Enforcement" |             |
| "Off"         |             |
| "Warning"     |             |

<a id="SafeguardsProfile_Level_STATUS"></a>SafeguardsProfile_Level_STATUS
-------------------------------------------------------------------------

Used by: [SafeguardsProfile_STATUS](#SafeguardsProfile_STATUS).

| Value         | Description |
|---------------|-------------|
| "Enforcement" |             |
| "Off"         |             |
| "Warning"     |             |

<a id="ScaleProfile"></a>ScaleProfile
-------------------------------------

Used by: [VirtualMachinesProfile](#VirtualMachinesProfile).

| Property | Description | Type                                                                    |
|----------|-------------|-------------------------------------------------------------------------|
| manual   |             | [ManualScaleProfile[]](#ManualScaleProfile)<br/><small>Optional</small> |

<a id="ScaleProfile_STATUS"></a>ScaleProfile_STATUS
---------------------------------------------------

Used by: [VirtualMachinesProfile_STATUS](#VirtualMachinesProfile_STATUS).

| Property | Description | Type                                                                                  |
|----------|-------------|---------------------------------------------------------------------------------------|
| manual   |             | [ManualScaleProfile_STATUS[]](#ManualScaleProfile_STATUS)<br/><small>Optional</small> |

<a id="ServiceMeshProfile_Mode"></a>ServiceMeshProfile_Mode
-----------------------------------------------------------

Used by: [ServiceMeshProfile](#ServiceMeshProfile).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Istio"    |             |

<a id="ServiceMeshProfile_Mode_STATUS"></a>ServiceMeshProfile_Mode_STATUS
-------------------------------------------------------------------------

Used by: [ServiceMeshProfile_STATUS](#ServiceMeshProfile_STATUS).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Istio"    |             |

<a id="SysctlConfig"></a>SysctlConfig
-------------------------------------

Used by: [LinuxOSConfig](#LinuxOSConfig).

| Property                       | Description | Type                               |
|--------------------------------|-------------|------------------------------------|
| fsAioMaxNr                     |             | int<br/><small>Optional</small>    |
| fsFileMax                      |             | int<br/><small>Optional</small>    |
| fsInotifyMaxUserWatches        |             | int<br/><small>Optional</small>    |
| fsNrOpen                       |             | int<br/><small>Optional</small>    |
| kernelThreadsMax               |             | int<br/><small>Optional</small>    |
| netCoreNetdevMaxBacklog        |             | int<br/><small>Optional</small>    |
| netCoreOptmemMax               |             | int<br/><small>Optional</small>    |
| netCoreRmemDefault             |             | int<br/><small>Optional</small>    |
| netCoreRmemMax                 |             | int<br/><small>Optional</small>    |
| netCoreSomaxconn               |             | int<br/><small>Optional</small>    |
| netCoreWmemDefault             |             | int<br/><small>Optional</small>    |
| netCoreWmemMax                 |             | int<br/><small>Optional</small>    |
| netIpv4IpLocalPortRange        |             | string<br/><small>Optional</small> |
| netIpv4NeighDefaultGcThresh1   |             | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh2   |             | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh3   |             | int<br/><small>Optional</small>    |
| netIpv4TcpFinTimeout           |             | int<br/><small>Optional</small>    |
| netIpv4TcpkeepaliveIntvl       |             | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveProbes      |             | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveTime        |             | int<br/><small>Optional</small>    |
| netIpv4TcpMaxSynBacklog        |             | int<br/><small>Optional</small>    |
| netIpv4TcpMaxTwBuckets         |             | int<br/><small>Optional</small>    |
| netIpv4TcpTwReuse              |             | bool<br/><small>Optional</small>   |
| netNetfilterNfConntrackBuckets |             | int<br/><small>Optional</small>    |
| netNetfilterNfConntrackMax     |             | int<br/><small>Optional</small>    |
| vmMaxMapCount                  |             | int<br/><small>Optional</small>    |
| vmSwappiness                   |             | int<br/><small>Optional</small>    |
| vmVfsCachePressure             |             | int<br/><small>Optional</small>    |

<a id="SysctlConfig_STATUS"></a>SysctlConfig_STATUS
---------------------------------------------------

Used by: [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS).

| Property                       | Description | Type                               |
|--------------------------------|-------------|------------------------------------|
| fsAioMaxNr                     |             | int<br/><small>Optional</small>    |
| fsFileMax                      |             | int<br/><small>Optional</small>    |
| fsInotifyMaxUserWatches        |             | int<br/><small>Optional</small>    |
| fsNrOpen                       |             | int<br/><small>Optional</small>    |
| kernelThreadsMax               |             | int<br/><small>Optional</small>    |
| netCoreNetdevMaxBacklog        |             | int<br/><small>Optional</small>    |
| netCoreOptmemMax               |             | int<br/><small>Optional</small>    |
| netCoreRmemDefault             |             | int<br/><small>Optional</small>    |
| netCoreRmemMax                 |             | int<br/><small>Optional</small>    |
| netCoreSomaxconn               |             | int<br/><small>Optional</small>    |
| netCoreWmemDefault             |             | int<br/><small>Optional</small>    |
| netCoreWmemMax                 |             | int<br/><small>Optional</small>    |
| netIpv4IpLocalPortRange        |             | string<br/><small>Optional</small> |
| netIpv4NeighDefaultGcThresh1   |             | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh2   |             | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh3   |             | int<br/><small>Optional</small>    |
| netIpv4TcpFinTimeout           |             | int<br/><small>Optional</small>    |
| netIpv4TcpkeepaliveIntvl       |             | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveProbes      |             | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveTime        |             | int<br/><small>Optional</small>    |
| netIpv4TcpMaxSynBacklog        |             | int<br/><small>Optional</small>    |
| netIpv4TcpMaxTwBuckets         |             | int<br/><small>Optional</small>    |
| netIpv4TcpTwReuse              |             | bool<br/><small>Optional</small>   |
| netNetfilterNfConntrackBuckets |             | int<br/><small>Optional</small>    |
| netNetfilterNfConntrackMax     |             | int<br/><small>Optional</small>    |
| vmMaxMapCount                  |             | int<br/><small>Optional</small>    |
| vmSwappiness                   |             | int<br/><small>Optional</small>    |
| vmVfsCachePressure             |             | int<br/><small>Optional</small>    |

<a id="SystemData_CreatedByType_STATUS"></a>SystemData_CreatedByType_STATUS
---------------------------------------------------------------------------

Used by: [SystemData_STATUS](#SystemData_STATUS).

| Value             | Description |
|-------------------|-------------|
| "Application"     |             |
| "Key"             |             |
| "ManagedIdentity" |             |
| "User"            |             |

<a id="SystemData_LastModifiedByType_STATUS"></a>SystemData_LastModifiedByType_STATUS
-------------------------------------------------------------------------------------

Used by: [SystemData_STATUS](#SystemData_STATUS).

| Value             | Description |
|-------------------|-------------|
| "Application"     |             |
| "Key"             |             |
| "ManagedIdentity" |             |
| "User"            |             |

<a id="UpgradeOverrideSettings"></a>UpgradeOverrideSettings
-----------------------------------------------------------

Used by: [ClusterUpgradeSettings](#ClusterUpgradeSettings).

| Property     | Description | Type                               |
|--------------|-------------|------------------------------------|
| forceUpgrade |             | bool<br/><small>Optional</small>   |
| until        |             | string<br/><small>Optional</small> |

<a id="UpgradeOverrideSettings_STATUS"></a>UpgradeOverrideSettings_STATUS
-------------------------------------------------------------------------

Used by: [ClusterUpgradeSettings_STATUS](#ClusterUpgradeSettings_STATUS).

| Property     | Description | Type                               |
|--------------|-------------|------------------------------------|
| forceUpgrade |             | bool<br/><small>Optional</small>   |
| until        |             | string<br/><small>Optional</small> |

<a id="UserAssignedIdentityDetails"></a>UserAssignedIdentityDetails
-------------------------------------------------------------------

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity).

| Property  | Description | Type                                                                                                                                                       |
|-----------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| reference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

<a id="WindowsGmsaProfile"></a>WindowsGmsaProfile
-------------------------------------------------

Used by: [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile).

| Property       | Description | Type                               |
|----------------|-------------|------------------------------------|
| dnsServer      |             | string<br/><small>Optional</small> |
| enabled        |             | bool<br/><small>Optional</small>   |
| rootDomainName |             | string<br/><small>Optional</small> |

<a id="WindowsGmsaProfile_STATUS"></a>WindowsGmsaProfile_STATUS
---------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS).

| Property       | Description | Type                               |
|----------------|-------------|------------------------------------|
| dnsServer      |             | string<br/><small>Optional</small> |
| enabled        |             | bool<br/><small>Optional</small>   |
| rootDomainName |             | string<br/><small>Optional</small> |

<a id="AzureKeyVaultKms_KeyVaultNetworkAccess"></a>AzureKeyVaultKms_KeyVaultNetworkAccess
-----------------------------------------------------------------------------------------

Used by: [AzureKeyVaultKms](#AzureKeyVaultKms).

| Value     | Description |
|-----------|-------------|
| "Private" |             |
| "Public"  |             |

<a id="AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS"></a>AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS
-------------------------------------------------------------------------------------------------------

Used by: [AzureKeyVaultKms_STATUS](#AzureKeyVaultKms_STATUS).

| Value     | Description |
|-----------|-------------|
| "Private" |             |
| "Public"  |             |

<a id="ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig"></a>ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig
-------------------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig](#ContainerServiceNetworkProfile_KubeProxyConfig).

| Property             | Description | Type                                                                                                                                                                    |
|----------------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| scheduler            |             | [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler)<br/><small>Optional</small> |
| tcpFinTimeoutSeconds |             | int<br/><small>Optional</small>                                                                                                                                         |
| tcpTimeoutSeconds    |             | int<br/><small>Optional</small>                                                                                                                                         |
| udpTimeoutSeconds    |             | int<br/><small>Optional</small>                                                                                                                                         |

<a id="ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS"></a>ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS
---------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_STATUS).

| Property             | Description | Type                                                                                                                                                                                  |
|----------------------|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| scheduler            |             | [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS)<br/><small>Optional</small> |
| tcpFinTimeoutSeconds |             | int<br/><small>Optional</small>                                                                                                                                                       |
| tcpTimeoutSeconds    |             | int<br/><small>Optional</small>                                                                                                                                                       |
| udpTimeoutSeconds    |             | int<br/><small>Optional</small>                                                                                                                                                       |

<a id="ContainerServiceNetworkProfile_KubeProxyConfig_Mode"></a>ContainerServiceNetworkProfile_KubeProxyConfig_Mode
-------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig](#ContainerServiceNetworkProfile_KubeProxyConfig).

| Value      | Description |
|------------|-------------|
| "IPTABLES" |             |
| "IPVS"     |             |

<a id="ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS"></a>ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS
---------------------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_STATUS).

| Value      | Description |
|------------|-------------|
| "IPTABLES" |             |
| "IPVS"     |             |

<a id="ContainerServiceSshPublicKey"></a>ContainerServiceSshPublicKey
---------------------------------------------------------------------

Used by: [ContainerServiceSshConfiguration](#ContainerServiceSshConfiguration).

| Property | Description | Type                               |
|----------|-------------|------------------------------------|
| keyData  |             | string<br/><small>Required</small> |

<a id="ContainerServiceSshPublicKey_STATUS"></a>ContainerServiceSshPublicKey_STATUS
-----------------------------------------------------------------------------------

Used by: [ContainerServiceSshConfiguration_STATUS](#ContainerServiceSshConfiguration_STATUS).

| Property | Description | Type                               |
|----------|-------------|------------------------------------|
| keyData  |             | string<br/><small>Optional</small> |

<a id="IstioCertificateAuthority"></a>IstioCertificateAuthority
---------------------------------------------------------------

Used by: [IstioServiceMesh](#IstioServiceMesh).

| Property | Description | Type                                                                                            |
|----------|-------------|-------------------------------------------------------------------------------------------------|
| plugin   |             | [IstioPluginCertificateAuthority](#IstioPluginCertificateAuthority)<br/><small>Optional</small> |

<a id="IstioCertificateAuthority_STATUS"></a>IstioCertificateAuthority_STATUS
-----------------------------------------------------------------------------

Used by: [IstioServiceMesh_STATUS](#IstioServiceMesh_STATUS).

| Property | Description | Type                                                                                                          |
|----------|-------------|---------------------------------------------------------------------------------------------------------------|
| plugin   |             | [IstioPluginCertificateAuthority_STATUS](#IstioPluginCertificateAuthority_STATUS)<br/><small>Optional</small> |

<a id="IstioComponents"></a>IstioComponents
-------------------------------------------

Used by: [IstioServiceMesh](#IstioServiceMesh).

| Property        | Description | Type                                                                      |
|-----------------|-------------|---------------------------------------------------------------------------|
| egressGateways  |             | [IstioEgressGateway[]](#IstioEgressGateway)<br/><small>Optional</small>   |
| ingressGateways |             | [IstioIngressGateway[]](#IstioIngressGateway)<br/><small>Optional</small> |

<a id="IstioComponents_STATUS"></a>IstioComponents_STATUS
---------------------------------------------------------

Used by: [IstioServiceMesh_STATUS](#IstioServiceMesh_STATUS).

| Property        | Description | Type                                                                                    |
|-----------------|-------------|-----------------------------------------------------------------------------------------|
| egressGateways  |             | [IstioEgressGateway_STATUS[]](#IstioEgressGateway_STATUS)<br/><small>Optional</small>   |
| ingressGateways |             | [IstioIngressGateway_STATUS[]](#IstioIngressGateway_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileAppMonitoring"></a>ManagedClusterAzureMonitorProfileAppMonitoring
---------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileLogs](#ManagedClusterAzureMonitorProfileLogs).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileAppMonitoring_STATUS"></a>ManagedClusterAzureMonitorProfileAppMonitoring_STATUS
-----------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileLogs_STATUS](#ManagedClusterAzureMonitorProfileLogs_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics"></a>ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics
-------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileMetrics](#ManagedClusterAzureMonitorProfileMetrics).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics_STATUS"></a>ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics_STATUS
---------------------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileMetrics_STATUS](#ManagedClusterAzureMonitorProfileMetrics_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileContainerInsights"></a>ManagedClusterAzureMonitorProfileContainerInsights
-----------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileLogs](#ManagedClusterAzureMonitorProfileLogs).

| Property                               | Description | Type                                                                                                                                                       |
|----------------------------------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| enabled                                |             | bool<br/><small>Optional</small>                                                                                                                           |
| logAnalyticsWorkspaceResourceReference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| windowsHostLogs                        |             | [ManagedClusterAzureMonitorProfileWindowsHostLogs](#ManagedClusterAzureMonitorProfileWindowsHostLogs)<br/><small>Optional</small>                          |

<a id="ManagedClusterAzureMonitorProfileContainerInsights_STATUS"></a>ManagedClusterAzureMonitorProfileContainerInsights_STATUS
-------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileLogs_STATUS](#ManagedClusterAzureMonitorProfileLogs_STATUS).

| Property                        | Description | Type                                                                                                                                            |
|---------------------------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| enabled                         |             | bool<br/><small>Optional</small>                                                                                                                |
| logAnalyticsWorkspaceResourceId |             | string<br/><small>Optional</small>                                                                                                              |
| windowsHostLogs                 |             | [ManagedClusterAzureMonitorProfileWindowsHostLogs_STATUS](#ManagedClusterAzureMonitorProfileWindowsHostLogs_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileKubeStateMetrics"></a>ManagedClusterAzureMonitorProfileKubeStateMetrics
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileMetrics](#ManagedClusterAzureMonitorProfileMetrics).

| Property                   | Description | Type                               |
|----------------------------|-------------|------------------------------------|
| metricAnnotationsAllowList |             | string<br/><small>Optional</small> |
| metricLabelsAllowlist      |             | string<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS"></a>ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS
-----------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileMetrics_STATUS](#ManagedClusterAzureMonitorProfileMetrics_STATUS).

| Property                   | Description | Type                               |
|----------------------------|-------------|------------------------------------|
| metricAnnotationsAllowList |             | string<br/><small>Optional</small> |
| metricLabelsAllowlist      |             | string<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_BackendPoolType"></a>ManagedClusterLoadBalancerProfile_BackendPoolType
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Value                 | Description |
|-----------------------|-------------|
| "NodeIP"              |             |
| "NodeIPConfiguration" |             |

<a id="ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS"></a>ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS
-----------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Value                 | Description |
|-----------------------|-------------|
| "NodeIP"              |             |
| "NodeIPConfiguration" |             |

<a id="ManagedClusterLoadBalancerProfile_ManagedOutboundIPs"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property  | Description | Type                            |
|-----------|-------------|---------------------------------|
| count     |             | int<br/><small>Optional</small> |
| countIPv6 |             | int<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS"></a>ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS
-----------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property  | Description | Type                            |
|-----------|-------------|---------------------------------|
| count     |             | int<br/><small>Optional</small> |
| countIPv6 |             | int<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPPrefixes"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property         | Description | Type                                                                  |
|------------------|-------------|-----------------------------------------------------------------------|
| publicIPPrefixes |             | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS"></a>ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS
-----------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property         | Description | Type                                                                                |
|------------------|-------------|-------------------------------------------------------------------------------------|
| publicIPPrefixes |             | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPs"></a>ManagedClusterLoadBalancerProfile_OutboundIPs
-------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property  | Description | Type                                                                  |
|-----------|-------------|-----------------------------------------------------------------------|
| publicIPs |             | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small> |

<a id="ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS"></a>ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS
---------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property  | Description | Type                                                                                |
|-----------|-------------|-------------------------------------------------------------------------------------|
| publicIPs |             | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterManagedOutboundIPProfile"></a>ManagedClusterManagedOutboundIPProfile
-----------------------------------------------------------------------------------------

Used by: [ManagedClusterNATGatewayProfile](#ManagedClusterNATGatewayProfile).

| Property | Description | Type                            |
|----------|-------------|---------------------------------|
| count    |             | int<br/><small>Optional</small> |

<a id="ManagedClusterManagedOutboundIPProfile_STATUS"></a>ManagedClusterManagedOutboundIPProfile_STATUS
-------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterNATGatewayProfile_STATUS](#ManagedClusterNATGatewayProfile_STATUS).

| Property | Description | Type                            |
|----------|-------------|---------------------------------|
| count    |             | int<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentity_ProvisioningInfo_STATUS"></a>ManagedClusterPodIdentity_ProvisioningInfo_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Property | Description | Type                                                                                                                                |
|----------|-------------|-------------------------------------------------------------------------------------------------------------------------------------|
| error    |             | [ManagedClusterPodIdentityProvisioningError_STATUS](#ManagedClusterPodIdentityProvisioningError_STATUS)<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentity_ProvisioningState_STATUS"></a>ManagedClusterPodIdentity_ProvisioningState_STATUS
-----------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Value       | Description |
|-------------|-------------|
| "Assigned"  |             |
| "Canceled"  |             |
| "Deleting"  |             |
| "Failed"    |             |
| "Succeeded" |             |
| "Updating"  |             |

<a id="ManagedClusterSecurityProfileDefenderSecurityMonitoring"></a>ManagedClusterSecurityProfileDefenderSecurityMonitoring
---------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfileDefender](#ManagedClusterSecurityProfileDefender).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS"></a>ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS
-----------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterSecurityProfileDefender_STATUS](#ManagedClusterSecurityProfileDefender_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling"></a>ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Enabled"  |             |

<a id="ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling_STATUS"></a>ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling_STATUS
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Enabled"  |             |

<a id="ManualScaleProfile"></a>ManualScaleProfile
-------------------------------------------------

Used by: [ScaleProfile](#ScaleProfile).

| Property | Description | Type                                 |
|----------|-------------|--------------------------------------|
| count    |             | int<br/><small>Optional</small>      |
| sizes    |             | string[]<br/><small>Optional</small> |

<a id="ManualScaleProfile_STATUS"></a>ManualScaleProfile_STATUS
---------------------------------------------------------------

Used by: [ScaleProfile_STATUS](#ScaleProfile_STATUS).

| Property | Description | Type                                 |
|----------|-------------|--------------------------------------|
| count    |             | int<br/><small>Optional</small>      |
| sizes    |             | string[]<br/><small>Optional</small> |

<a id="PortRange_Protocol"></a>PortRange_Protocol
-------------------------------------------------

Used by: [PortRange](#PortRange).

| Value | Description |
|-------|-------------|
| "TCP" |             |
| "UDP" |             |

<a id="PortRange_Protocol_STATUS"></a>PortRange_Protocol_STATUS
---------------------------------------------------------------

Used by: [PortRange_STATUS](#PortRange_STATUS).

| Value | Description |
|-------|-------------|
| "TCP" |             |
| "UDP" |             |

<a id="ResourceReference"></a>ResourceReference
-----------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile), [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes), [ManagedClusterLoadBalancerProfile_OutboundIPs](#ManagedClusterLoadBalancerProfile_OutboundIPs), and [ManagedClusterNATGatewayProfile](#ManagedClusterNATGatewayProfile).

| Property  | Description | Type                                                                                                                                                       |
|-----------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| reference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

<a id="ResourceReference_STATUS"></a>ResourceReference_STATUS
-------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS), [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS), [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS), and [ManagedClusterNATGatewayProfile_STATUS](#ManagedClusterNATGatewayProfile_STATUS).

| Property | Description | Type                               |
|----------|-------------|------------------------------------|
| id       |             | string<br/><small>Optional</small> |

<a id="ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler"></a>ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler
---------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig).

| Value             | Description |
|-------------------|-------------|
| "LeastConnection" |             |
| "RoundRobin"      |             |

<a id="ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS"></a>ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS
-----------------------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS).

| Value             | Description |
|-------------------|-------------|
| "LeastConnection" |             |
| "RoundRobin"      |             |

<a id="IstioEgressGateway"></a>IstioEgressGateway
-------------------------------------------------

Used by: [IstioComponents](#IstioComponents).

| Property     | Description | Type                                          |
|--------------|-------------|-----------------------------------------------|
| enabled      |             | bool<br/><small>Required</small>              |
| nodeSelector |             | map[string]string<br/><small>Optional</small> |

<a id="IstioEgressGateway_STATUS"></a>IstioEgressGateway_STATUS
---------------------------------------------------------------

Used by: [IstioComponents_STATUS](#IstioComponents_STATUS).

| Property     | Description | Type                                          |
|--------------|-------------|-----------------------------------------------|
| enabled      |             | bool<br/><small>Optional</small>              |
| nodeSelector |             | map[string]string<br/><small>Optional</small> |

<a id="IstioIngressGateway"></a>IstioIngressGateway
---------------------------------------------------

Used by: [IstioComponents](#IstioComponents).

| Property | Description | Type                                                                              |
|----------|-------------|-----------------------------------------------------------------------------------|
| enabled  |             | bool<br/><small>Required</small>                                                  |
| mode     |             | [IstioIngressGateway_Mode](#IstioIngressGateway_Mode)<br/><small>Required</small> |

<a id="IstioIngressGateway_STATUS"></a>IstioIngressGateway_STATUS
-----------------------------------------------------------------

Used by: [IstioComponents_STATUS](#IstioComponents_STATUS).

| Property | Description | Type                                                                                            |
|----------|-------------|-------------------------------------------------------------------------------------------------|
| enabled  |             | bool<br/><small>Optional</small>                                                                |
| mode     |             | [IstioIngressGateway_Mode_STATUS](#IstioIngressGateway_Mode_STATUS)<br/><small>Optional</small> |

<a id="IstioPluginCertificateAuthority"></a>IstioPluginCertificateAuthority
---------------------------------------------------------------------------

Used by: [IstioCertificateAuthority](#IstioCertificateAuthority).

| Property            | Description | Type                                                                                                                                                       |
|---------------------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| certChainObjectName |             | string<br/><small>Optional</small>                                                                                                                         |
| certObjectName      |             | string<br/><small>Optional</small>                                                                                                                         |
| keyObjectName       |             | string<br/><small>Optional</small>                                                                                                                         |
| keyVaultReference   |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| rootCertObjectName  |             | string<br/><small>Optional</small>                                                                                                                         |

<a id="IstioPluginCertificateAuthority_STATUS"></a>IstioPluginCertificateAuthority_STATUS
-----------------------------------------------------------------------------------------

Used by: [IstioCertificateAuthority_STATUS](#IstioCertificateAuthority_STATUS).

| Property            | Description | Type                               |
|---------------------|-------------|------------------------------------|
| certChainObjectName |             | string<br/><small>Optional</small> |
| certObjectName      |             | string<br/><small>Optional</small> |
| keyObjectName       |             | string<br/><small>Optional</small> |
| keyVaultId          |             | string<br/><small>Optional</small> |
| rootCertObjectName  |             | string<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileWindowsHostLogs"></a>ManagedClusterAzureMonitorProfileWindowsHostLogs
-------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileContainerInsights](#ManagedClusterAzureMonitorProfileContainerInsights).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterAzureMonitorProfileWindowsHostLogs_STATUS"></a>ManagedClusterAzureMonitorProfileWindowsHostLogs_STATUS
---------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAzureMonitorProfileContainerInsights_STATUS](#ManagedClusterAzureMonitorProfileContainerInsights_STATUS).

| Property | Description | Type                             |
|----------|-------------|----------------------------------|
| enabled  |             | bool<br/><small>Optional</small> |

<a id="ManagedClusterPodIdentityProvisioningError_STATUS"></a>ManagedClusterPodIdentityProvisioningError_STATUS
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS).

| Property | Description | Type                                                                                                                                        |
|----------|-------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| error    |             | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS)<br/><small>Optional</small> |

<a id="IstioIngressGateway_Mode"></a>IstioIngressGateway_Mode
-------------------------------------------------------------

Used by: [IstioIngressGateway](#IstioIngressGateway).

| Value      | Description |
|------------|-------------|
| "External" |             |
| "Internal" |             |

<a id="IstioIngressGateway_Mode_STATUS"></a>IstioIngressGateway_Mode_STATUS
---------------------------------------------------------------------------

Used by: [IstioIngressGateway_STATUS](#IstioIngressGateway_STATUS).

| Value      | Description |
|------------|-------------|
| "External" |             |
| "Internal" |             |

<a id="ManagedClusterPodIdentityProvisioningErrorBody_STATUS"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS
-----------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProvisioningError_STATUS](#ManagedClusterPodIdentityProvisioningError_STATUS).

| Property | Description | Type                                                                                                                                                            |
|----------|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| code     |             | string<br/><small>Optional</small>                                                                                                                              |
| details  |             | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled[]](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled)<br/><small>Optional</small> |
| message  |             | string<br/><small>Optional</small>                                                                                                                              |
| target   |             | string<br/><small>Optional</small>                                                                                                                              |

<a id="ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled"></a>ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled
-----------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS).

| Property | Description | Type                               |
|----------|-------------|------------------------------------|
| code     |             | string<br/><small>Optional</small> |
| message  |             | string<br/><small>Optional</small> |
| target   |             | string<br/><small>Optional</small> |
