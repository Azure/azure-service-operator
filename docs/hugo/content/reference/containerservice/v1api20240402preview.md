---

title: containerservice.azure.com/v1api20240402preview

linktitle: v1api20240402preview
-------------------------------

APIVersion{#APIVersion}
-----------------------

| Value                | Description |
|----------------------|-------------|
| "2024-04-02-preview" |             |

ManagedCluster{#ManagedCluster}
-------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{resourceName}

Used by: [ManagedClusterList](#ManagedClusterList).

| Property                                                                                | Description | Type                                                                        |
|-----------------------------------------------------------------------------------------|-------------|-----------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                             |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                             |
| spec                                                                                    |             | [ManagedCluster_Spec](#ManagedCluster_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [ManagedCluster_STATUS](#ManagedCluster_STATUS)<br/><small>Optional</small> |

### ManagedCluster_Spec {#ManagedCluster_Spec}

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                                 |
|----------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile                 | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile](#ManagedClusterAADProfile)<br/><small>Optional</small>                                                                                    |
| addonProfiles              | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile](#ManagedClusterAddonProfile)<br/><small>Optional</small>                                                                     |
| agentPoolProfiles          | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile[]](#ManagedClusterAgentPoolProfile)<br/><small>Optional</small>                                                                      |
| aiToolchainOperatorProfile | AI toolchain operator settings that apply to the whole cluster.                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAIToolchainOperatorProfile](#ManagedClusterAIToolchainOperatorProfile)<br/><small>Optional</small>                                                    |
| apiServerAccessProfile     | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile](#ManagedClusterAPIServerAccessProfile)<br/><small>Optional</small>                                                            |
| autoScalerProfile          | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile)<br/><small>Optional</small>                                                |
| autoUpgradeProfile         | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile)<br/><small>Optional</small>                                                                    |
| azureMonitorProfile        | Prometheus addon profile for the container service cluster                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAzureMonitorProfile](#ManagedClusterAzureMonitorProfile)<br/><small>Optional</small>                                                                  |
| azureName                  | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| bootstrapProfile           | Profile of the cluster bootstrap configuration.                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterBootstrapProfile](#ManagedClusterBootstrapProfile)<br/><small>Optional</small>                                                                        |
| creationData               | CreationData to be used to specify the source Snapshot ID if the cluster will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                         | [CreationData](#CreationData)<br/><small>Optional</small>                                                                                                            |
| disableLocalAccounts       | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| diskEncryptionSetReference | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;diskEncryptionSets/&ZeroWidthSpace;{encryptionSetName}'                                                                                                | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| dnsPrefix                  | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                                   |
| enableNamespaceResources   | The default value is false. It can be enabled/disabled on creation and updating of the managed cluster. See [https://aka.ms/NamespaceARMResource](https://aka.ms/NamespaceARMResource) for more details on Namespace as a ARM Resource.                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                                                     |
| enablePodSecurityPolicy    | (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at https://aka.ms/k8s/psp and https://aka.ms/aks/psp.                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                                                                     |
| enableRBAC                 | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                                                     |
| extendedLocation           | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation](#ExtendedLocation)<br/><small>Optional</small>                                                                                                    |
| fqdnSubdomain              | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                                   |
| httpProxyConfig            | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig](#ManagedClusterHTTPProxyConfig)<br/><small>Optional</small>                                                                          |
| identity                   | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity](#ManagedClusterIdentity)<br/><small>Optional</small>                                                                                        |
| identityProfile            | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity](#UserAssignedIdentity)<br/><small>Optional</small>                                                                                 |
| ingressProfile             | Ingress profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterIngressProfile](#ManagedClusterIngressProfile)<br/><small>Optional</small>                                                                            |
| kind                       | This is primarily used to expose different UI experiences in the portal for different kinds                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |
| kubernetesVersion          | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string<br/><small>Optional</small>                                                                                                                                   |
| linuxProfile               | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile)<br/><small>Optional</small>                                                                            |
| location                   | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Required</small>                                                                                                                                   |
| metricsProfile             | Optional cluster metrics configuration.                                                                                                                                                                                                                                                                                                                                                     | [ManagedClusterMetricsProfile](#ManagedClusterMetricsProfile)<br/><small>Optional</small>                                                                            |
| networkProfile             | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile)<br/><small>Optional</small>                                                                        |
| nodeProvisioningProfile    | Node provisioning settings that apply to the whole cluster.                                                                                                                                                                                                                                                                                                                                 | [ManagedClusterNodeProvisioningProfile](#ManagedClusterNodeProvisioningProfile)<br/><small>Optional</small>                                                          |
| nodeResourceGroup          | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |
| nodeResourceGroupProfile   | The node resource group configuration profile.                                                                                                                                                                                                                                                                                                                                              | [ManagedClusterNodeResourceGroupProfile](#ManagedClusterNodeResourceGroupProfile)<br/><small>Optional</small>                                                        |
| oidcIssuerProfile          | The OIDC issuer profile of the Managed Cluster.                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterOIDCIssuerProfile](#ManagedClusterOIDCIssuerProfile)<br/><small>Optional</small>                                                                      |
| operatorSpec               | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                             | [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec)<br/><small>Optional</small>                                                                                |
| owner                      | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource                                                                                                | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| podIdentityProfile         | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile)<br/><small>Optional</small>                                                                    |
| privateLinkResources       | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource[]](#PrivateLinkResource)<br/><small>Optional</small>                                                                                            |
| publicNetworkAccess        | Allow or deny public network access for AKS                                                                                                                                                                                                                                                                                                                                                 | [ManagedClusterProperties_PublicNetworkAccess](#ManagedClusterProperties_PublicNetworkAccess)<br/><small>Optional</small>                                            |
| safeguardsProfile          | The Safeguards profile holds all the safeguards information for a given cluster                                                                                                                                                                                                                                                                                                             | [SafeguardsProfile](#SafeguardsProfile)<br/><small>Optional</small>                                                                                                  |
| securityProfile            | Security profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile)<br/><small>Optional</small>                                                                          |
| serviceMeshProfile         | Service mesh profile for a managed cluster.                                                                                                                                                                                                                                                                                                                                                 | [ServiceMeshProfile](#ServiceMeshProfile)<br/><small>Optional</small>                                                                                                |
| servicePrincipalProfile    | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile](#ManagedClusterServicePrincipalProfile)<br/><small>Optional</small>                                                          |
| sku                        | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU](#ManagedClusterSKU)<br/><small>Optional</small>                                                                                                  |
| storageProfile             | Storage profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterStorageProfile](#ManagedClusterStorageProfile)<br/><small>Optional</small>                                                                            |
| supportPlan                | The support plan for the Managed Cluster. If unspecified, the default is `KubernetesOfficial`.                                                                                                                                                                                                                                                                                              | [KubernetesSupportPlan](#KubernetesSupportPlan)<br/><small>Optional</small>                                                                                          |
| tags                       | Resource tags.                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                                                                                                        |
| upgradeSettings            | Settings for upgrading a cluster.                                                                                                                                                                                                                                                                                                                                                           | [ClusterUpgradeSettings](#ClusterUpgradeSettings)<br/><small>Optional</small>                                                                                        |
| windowsProfile             | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile)<br/><small>Optional</small>                                                                            |
| workloadAutoScalerProfile  | Workload Auto-scaler profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                       | [ManagedClusterWorkloadAutoScalerProfile](#ManagedClusterWorkloadAutoScalerProfile)<br/><small>Optional</small>                                                      |

### ManagedCluster_STATUS{#ManagedCluster_STATUS}

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                    |
|----------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile                 | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile_STATUS](#ManagedClusterAADProfile_STATUS)<br/><small>Optional</small>                                                         |
| addonProfiles              | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS)<br/><small>Optional</small>                                          |
| agentPoolProfiles          | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile_STATUS[]](#ManagedClusterAgentPoolProfile_STATUS)<br/><small>Optional</small>                                           |
| aiToolchainOperatorProfile | AI toolchain operator settings that apply to the whole cluster.                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAIToolchainOperatorProfile_STATUS](#ManagedClusterAIToolchainOperatorProfile_STATUS)<br/><small>Optional</small>                         |
| apiServerAccessProfile     | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile_STATUS](#ManagedClusterAPIServerAccessProfile_STATUS)<br/><small>Optional</small>                                 |
| autoScalerProfile          | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS)<br/><small>Optional</small>                     |
| autoUpgradeProfile         | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS)<br/><small>Optional</small>                                         |
| azureMonitorProfile        | Prometheus addon profile for the container service cluster                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAzureMonitorProfile_STATUS](#ManagedClusterAzureMonitorProfile_STATUS)<br/><small>Optional</small>                                       |
| azurePortalFQDN            | The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                      |
| bootstrapProfile           | Profile of the cluster bootstrap configuration.                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterBootstrapProfile_STATUS](#ManagedClusterBootstrapProfile_STATUS)<br/><small>Optional</small>                                             |
| conditions                 | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                          | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| creationData               | CreationData to be used to specify the source Snapshot ID if the cluster will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                         | [CreationData_STATUS](#CreationData_STATUS)<br/><small>Optional</small>                                                                                 |
| currentKubernetesVersion   | The version of Kubernetes the Managed Cluster is running.                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| disableLocalAccounts       | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool<br/><small>Optional</small>                                                                                                                        |
| diskEncryptionSetID        | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;diskEncryptionSets/&ZeroWidthSpace;{encryptionSetName}'                                                                                                | string<br/><small>Optional</small>                                                                                                                      |
| dnsPrefix                  | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| enableNamespaceResources   | The default value is false. It can be enabled/disabled on creation and updating of the managed cluster. See [https://aka.ms/NamespaceARMResource](https://aka.ms/NamespaceARMResource) for more details on Namespace as a ARM Resource.                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                                        |
| enablePodSecurityPolicy    | (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at https://aka.ms/k8s/psp and https://aka.ms/aks/psp.                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                                                        |
| enableRBAC                 | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                                        |
| eTag                       | Unique read-only string used to implement optimistic concurrency. The eTag value will change when the resource is updated. Specify an if-match or if-none-match header with the eTag value for a subsequent request to enable optimistic concurrency per the normal etag convention.                                                                                                        | string<br/><small>Optional</small>                                                                                                                      |
| extendedLocation           | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation_STATUS](#ExtendedLocation_STATUS)<br/><small>Optional</small>                                                                         |
| fqdn                       | The FQDN of the master pool.                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                      |
| fqdnSubdomain              | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| httpProxyConfig            | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig_STATUS](#ManagedClusterHTTPProxyConfig_STATUS)<br/><small>Optional</small>                                               |
| id                         | Fully qualified resource ID for the resource. E.g. "/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}"                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| identity                   | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS)<br/><small>Optional</small>                                                             |
| identityProfile            | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small>                                                      |
| ingressProfile             | Ingress profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterIngressProfile_STATUS](#ManagedClusterIngressProfile_STATUS)<br/><small>Optional</small>                                                 |
| kind                       | This is primarily used to expose different UI experiences in the portal for different kinds                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| kubernetesVersion          | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string<br/><small>Optional</small>                                                                                                                      |
| linuxProfile               | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS)<br/><small>Optional</small>                                                 |
| location                   | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| maxAgentPools              | The max number of agent pools for the managed cluster.                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                         |
| metricsProfile             | Optional cluster metrics configuration.                                                                                                                                                                                                                                                                                                                                                     | [ManagedClusterMetricsProfile_STATUS](#ManagedClusterMetricsProfile_STATUS)<br/><small>Optional</small>                                                 |
| name                       | The name of the resource                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| networkProfile             | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS)<br/><small>Optional</small>                                             |
| nodeProvisioningProfile    | Node provisioning settings that apply to the whole cluster.                                                                                                                                                                                                                                                                                                                                 | [ManagedClusterNodeProvisioningProfile_STATUS](#ManagedClusterNodeProvisioningProfile_STATUS)<br/><small>Optional</small>                               |
| nodeResourceGroup          | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| nodeResourceGroupProfile   | The node resource group configuration profile.                                                                                                                                                                                                                                                                                                                                              | [ManagedClusterNodeResourceGroupProfile_STATUS](#ManagedClusterNodeResourceGroupProfile_STATUS)<br/><small>Optional</small>                             |
| oidcIssuerProfile          | The OIDC issuer profile of the Managed Cluster.                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterOIDCIssuerProfile_STATUS](#ManagedClusterOIDCIssuerProfile_STATUS)<br/><small>Optional</small>                                           |
| podIdentityProfile         | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS)<br/><small>Optional</small>                                         |
| powerState                 | The Power State of the cluster.                                                                                                                                                                                                                                                                                                                                                             | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                                     |
| privateFQDN                | The FQDN of private cluster.                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                      |
| privateLinkResources       | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource_STATUS[]](#PrivateLinkResource_STATUS)<br/><small>Optional</small>                                                                 |
| provisioningState          | The current provisioning state.                                                                                                                                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                      |
| publicNetworkAccess        | Allow or deny public network access for AKS                                                                                                                                                                                                                                                                                                                                                 | [ManagedClusterProperties_PublicNetworkAccess_STATUS](#ManagedClusterProperties_PublicNetworkAccess_STATUS)<br/><small>Optional</small>                 |
| resourceUID                | The resourceUID uniquely identifies ManagedClusters that reuse ARM ResourceIds (i.e: create, delete, create sequence)                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| safeguardsProfile          | The Safeguards profile holds all the safeguards information for a given cluster                                                                                                                                                                                                                                                                                                             | [SafeguardsProfile_STATUS](#SafeguardsProfile_STATUS)<br/><small>Optional</small>                                                                       |
| securityProfile            | Security profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS)<br/><small>Optional</small>                                               |
| serviceMeshProfile         | Service mesh profile for a managed cluster.                                                                                                                                                                                                                                                                                                                                                 | [ServiceMeshProfile_STATUS](#ServiceMeshProfile_STATUS)<br/><small>Optional</small>                                                                     |
| servicePrincipalProfile    | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile_STATUS](#ManagedClusterServicePrincipalProfile_STATUS)<br/><small>Optional</small>                               |
| sku                        | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS)<br/><small>Optional</small>                                                                       |
| storageProfile             | Storage profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS)<br/><small>Optional</small>                                                 |
| supportPlan                | The support plan for the Managed Cluster. If unspecified, the default is `KubernetesOfficial`.                                                                                                                                                                                                                                                                                              | [KubernetesSupportPlan_STATUS](#KubernetesSupportPlan_STATUS)<br/><small>Optional</small>                                                               |
| systemData                 | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags                       | Resource tags.                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                       | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| upgradeSettings            | Settings for upgrading a cluster.                                                                                                                                                                                                                                                                                                                                                           | [ClusterUpgradeSettings_STATUS](#ClusterUpgradeSettings_STATUS)<br/><small>Optional</small>                                                             |
| windowsProfile             | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS)<br/><small>Optional</small>                                                 |
| workloadAutoScalerProfile  | Workload Auto-scaler profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                       | [ManagedClusterWorkloadAutoScalerProfile_STATUS](#ManagedClusterWorkloadAutoScalerProfile_STATUS)<br/><small>Optional</small>                           |

ManagedClusterList{#ManagedClusterList}
---------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{resourceName}

| Property                                                                            | Description | Type                                                            |
|-------------------------------------------------------------------------------------|-------------|-----------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                 |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                 |
| items                                                                               |             | [ManagedCluster[]](#ManagedCluster)<br/><small>Optional</small> |

ManagedClustersAgentPool{#ManagedClustersAgentPool}
---------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{resourceName}/&ZeroWidthSpace;agentPools/&ZeroWidthSpace;{agentPoolName}

Used by: [ManagedClustersAgentPoolList](#ManagedClustersAgentPoolList).

| Property                                                                                | Description | Type                                                                                            |
|-----------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                                                 |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                                                 |
| spec                                                                                    |             | [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS)<br/><small>Optional</small> |

### ManagedClustersAgentPool_Spec {#ManagedClustersAgentPool_Spec}

| Property                          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                 |
|-----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile          | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile](#AgentPoolArtifactStreamingProfile)<br/><small>Optional</small>                                                                  |
| availabilityZones                 | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is `VirtualMachineScaleSets`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                                 |
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
| gatewayProfile                    | Profile specific to a managed agent pool in Gateway mode. This field cannot be set if agent pool mode is not Gateway.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [AgentPoolGatewayProfile](#AgentPoolGatewayProfile)<br/><small>Optional</small>                                                                                      |
| gpuInstanceProfile                | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile](#GPUInstanceProfile)<br/><small>Optional</small>                                                                                                |
| gpuProfile                        | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile](#AgentPoolGPUProfile)<br/><small>Optional</small>                                                                                              |
| hostGroupReference                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;hostGroups/&ZeroWidthSpace;{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                    | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| kubeletConfig                     | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig](#KubeletConfig)<br/><small>Optional</small>                                                                                                          |
| kubeletDiskType                   | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType](#KubeletDiskType)<br/><small>Optional</small>                                                                                                      |
| linuxOSConfig                     | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig](#LinuxOSConfig)<br/><small>Optional</small>                                                                                                          |
| maxCount                          | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                                      |
| maxPods                           | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                                      |
| messageOfTheDay                   | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                                   |
| minCount                          | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                                      |
| mode                              | A cluster must have at least one `System` Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode](#AgentPoolMode)<br/><small>Optional</small>                                                                                                          |
| networkProfile                    | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile](#AgentPoolNetworkProfile)<br/><small>Optional</small>                                                                                      |
| nodeInitializationTaints          | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                                 |
| nodeLabels                        | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                                                                        |
| nodePublicIPPrefixReference       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;publicIPPrefixes/&ZeroWidthSpace;{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| nodeTaints                        | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                                                                                 |
| operatorSpec                      | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ManagedClustersAgentPoolOperatorSpec](#ManagedClustersAgentPoolOperatorSpec)<br/><small>Optional</small>                                                            |
| orchestratorVersion               | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                                                                   |
| osDiskSizeGB                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ContainerServiceOSDisk](#ContainerServiceOSDisk)<br/><small>Optional</small>                                                                                        |
| osDiskType                        | The default is `Ephemeral` if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to `Managed`. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType](#OSDiskType)<br/><small>Optional</small>                                                                                                                |
| osSKU                             | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU](#OSSKU)<br/><small>Optional</small>                                                                                                                          |
| osType                            | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType](#OSType)<br/><small>Optional</small>                                                                                                                        |
| owner                             | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| podIPAllocationMode               | The IP allocation mode for pods in the agent pool. Must be used with podSubnetId. The default is `DynamicIndividual`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [PodIPAllocationMode](#PodIPAllocationMode)<br/><small>Optional</small>                                                                                              |
| podSubnetReference                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                                                                                                       | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| powerState                        | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState](#PowerState)<br/><small>Optional</small>                                                                                                                |
| proximityPlacementGroupReference  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| scaleDownMode                     | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode](#ScaleDownMode)<br/><small>Optional</small>                                                                                                          |
| scaleSetEvictionPolicy            | This cannot be specified unless the scaleSetPriority is `Spot`. If not specified, the default is `Delete`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)<br/><small>Optional</small>                                                                                        |
| scaleSetPriority                  | The Virtual Machine Scale Set priority. If not specified, the default is `Regular`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority](#ScaleSetPriority)<br/><small>Optional</small>                                                                                                    |
| securityProfile                   | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile](#AgentPoolSecurityProfile)<br/><small>Optional</small>                                                                                    |
| spotMaxPrice                      | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                                                                                  |
| tags                              | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                                                                        |
| type                              | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType](#AgentPoolType)<br/><small>Optional</small>                                                                                                          |
| upgradeSettings                   | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings)<br/><small>Optional</small>                                                                                    |
| virtualMachineNodesStatus         |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes[]](#VirtualMachineNodes)<br/><small>Optional</small>                                                                                            |
| virtualMachinesProfile            | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile](#VirtualMachinesProfile)<br/><small>Optional</small>                                                                                        |
| vmSize                            | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                                   |
| vnetSubnetReference               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| windowsProfile                    | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile](#AgentPoolWindowsProfile)<br/><small>Optional</small>                                                                                      |
| workloadRuntime                   | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime](#WorkloadRuntime)<br/><small>Optional</small>                                                                                                      |

### ManagedClustersAgentPool_STATUS{#ManagedClustersAgentPool_STATUS}

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                    |
|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile   | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile_STATUS](#AgentPoolArtifactStreamingProfile_STATUS)<br/><small>Optional</small>                                       |
| availabilityZones          | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is `VirtualMachineScaleSets`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                    |
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
| eTag                       | Unique read-only string used to implement optimistic concurrency. The eTag value will change when the resource is updated. Specify an if-match or if-none-match header with the eTag value for a subsequent request to enable optimistic concurrency per the normal etag convention.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| gatewayProfile             | Profile specific to a managed agent pool in Gateway mode. This field cannot be set if agent pool mode is not Gateway.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [AgentPoolGatewayProfile_STATUS](#AgentPoolGatewayProfile_STATUS)<br/><small>Optional</small>                                                           |
| gpuInstanceProfile         | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)<br/><small>Optional</small>                                                                     |
| gpuProfile                 | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile_STATUS](#AgentPoolGPUProfile_STATUS)<br/><small>Optional</small>                                                                   |
| hostGroupID                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;hostGroups/&ZeroWidthSpace;{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| id                         | Resource ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| kubeletConfig              | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig_STATUS](#KubeletConfig_STATUS)<br/><small>Optional</small>                                                                               |
| kubeletDiskType            | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)<br/><small>Optional</small>                                                                           |
| linuxOSConfig              | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)<br/><small>Optional</small>                                                                               |
| maxCount                   | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                         |
| maxPods                    | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                         |
| messageOfTheDay            | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                      |
| minCount                   | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                         |
| mode                       | A cluster must have at least one `System` Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)<br/><small>Optional</small>                                                                               |
| name                       | The name of the resource that is unique within a resource group. This name can be used to access the resource.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                      |
| networkProfile             | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS)<br/><small>Optional</small>                                                           |
| nodeImageVersion           | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| nodeInitializationTaints   | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                    |
| nodeLabels                 | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                                                           |
| nodePublicIPPrefixID       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;publicIPPrefixes/&ZeroWidthSpace;{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| nodeTaints                 | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                                                                    |
| orchestratorVersion        | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                                                      |
| osDiskSizeGB               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | int<br/><small>Optional</small>                                                                                                                         |
| osDiskType                 | The default is `Ephemeral` if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to `Managed`. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType_STATUS](#OSDiskType_STATUS)<br/><small>Optional</small>                                                                                     |
| osSKU                      | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU_STATUS](#OSSKU_STATUS)<br/><small>Optional</small>                                                                                               |
| osType                     | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType_STATUS](#OSType_STATUS)<br/><small>Optional</small>                                                                                             |
| podIPAllocationMode        | The IP allocation mode for pods in the agent pool. Must be used with podSubnetId. The default is `DynamicIndividual`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [PodIPAllocationMode_STATUS](#PodIPAllocationMode_STATUS)<br/><small>Optional</small>                                                                   |
| podSubnetID                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| powerState                 | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                                     |
| properties_type            | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType_STATUS](#AgentPoolType_STATUS)<br/><small>Optional</small>                                                                               |
| provisioningState          | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| proximityPlacementGroupID  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| scaleDownMode              | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode_STATUS](#ScaleDownMode_STATUS)<br/><small>Optional</small>                                                                               |
| scaleSetEvictionPolicy     | This cannot be specified unless the scaleSetPriority is `Spot`. If not specified, the default is `Delete`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)<br/><small>Optional</small>                                                             |
| scaleSetPriority           | The Virtual Machine Scale Set priority. If not specified, the default is `Regular`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)<br/><small>Optional</small>                                                                         |
| securityProfile            | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile_STATUS](#AgentPoolSecurityProfile_STATUS)<br/><small>Optional</small>                                                         |
| spotMaxPrice               | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                                                                     |
| tags                       | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                       | Resource type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| upgradeSettings            | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS)<br/><small>Optional</small>                                                         |
| virtualMachineNodesStatus  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes_STATUS[]](#VirtualMachineNodes_STATUS)<br/><small>Optional</small>                                                                 |
| virtualMachinesProfile     | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile_STATUS](#VirtualMachinesProfile_STATUS)<br/><small>Optional</small>                                                             |
| vmSize                     | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| vnetSubnetID               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| windowsProfile             | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile_STATUS](#AgentPoolWindowsProfile_STATUS)<br/><small>Optional</small>                                                           |
| workloadRuntime            | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime_STATUS](#WorkloadRuntime_STATUS)<br/><small>Optional</small>                                                                           |

ManagedClustersAgentPoolList{#ManagedClustersAgentPoolList}
-----------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{resourceName}/&ZeroWidthSpace;agentPools/&ZeroWidthSpace;{agentPoolName}

| Property                                                                            | Description | Type                                                                                |
|-------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                                     |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                                     |
| items                                                                               |             | [ManagedClustersAgentPool[]](#ManagedClustersAgentPool)<br/><small>Optional</small> |

TrustedAccessRoleBinding{#TrustedAccessRoleBinding}
---------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{resourceName}/&ZeroWidthSpace;trustedAccessRoleBindings/&ZeroWidthSpace;{trustedAccessRoleBindingName}

Used by: [TrustedAccessRoleBindingList](#TrustedAccessRoleBindingList).

| Property                                                                                | Description | Type                                                                                            |
|-----------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                                                 |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                                                 |
| spec                                                                                    |             | [TrustedAccessRoleBinding_Spec](#TrustedAccessRoleBinding_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [TrustedAccessRoleBinding_STATUS](#TrustedAccessRoleBinding_STATUS)<br/><small>Optional</small> |

### TrustedAccessRoleBinding_Spec {#TrustedAccessRoleBinding_Spec}

| Property                | Description                                                                                                                                                                                                                                                                                          | Type                                                                                                                                                                 |
|-------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName               | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec            | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                      | [TrustedAccessRoleBindingOperatorSpec](#TrustedAccessRoleBindingOperatorSpec)<br/><small>Optional</small>                                                            |
| owner                   | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| roles                   | A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.                                                                                                                                                       | string[]<br/><small>Required</small>                                                                                                                                 |
| sourceResourceReference | The ARM resource ID of source resource that trusted access is configured for.                                                                                                                                                                                                                        | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Required</small>           |

### TrustedAccessRoleBinding_STATUS{#TrustedAccessRoleBinding_STATUS}

| Property          | Description                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                    |
|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions        | The observed state of the resource                                                                                                                                                                                                                                                                                          | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| id                | Fully qualified resource ID for the resource. E.g. "/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}" | string<br/><small>Optional</small>                                                                                                                      |
| name              | The name of the resource                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState | The current provisioning state of trusted access role binding.                                                                                                                                                                                                                                                              | [TrustedAccessRoleBindingProperties_ProvisioningState_STATUS](#TrustedAccessRoleBindingProperties_ProvisioningState_STATUS)<br/><small>Optional</small> |
| roles             | A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.                                                                                                                                                                              | string[]<br/><small>Optional</small>                                                                                                                    |
| sourceResourceId  | The ARM resource ID of source resource that trusted access is configured for.                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| systemData        | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type              | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |

TrustedAccessRoleBindingList{#TrustedAccessRoleBindingList}
-----------------------------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/managedClusters.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{resourceName}/&ZeroWidthSpace;trustedAccessRoleBindings/&ZeroWidthSpace;{trustedAccessRoleBindingName}

| Property                                                                            | Description | Type                                                                                |
|-------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                                     |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                                     |
| items                                                                               |             | [TrustedAccessRoleBinding[]](#TrustedAccessRoleBinding)<br/><small>Optional</small> |

ManagedCluster_Spec{#ManagedCluster_Spec}
-----------------------------------------

Used by: [ManagedCluster](#ManagedCluster).

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                                 |
|----------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile                 | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile](#ManagedClusterAADProfile)<br/><small>Optional</small>                                                                                    |
| addonProfiles              | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile](#ManagedClusterAddonProfile)<br/><small>Optional</small>                                                                     |
| agentPoolProfiles          | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile[]](#ManagedClusterAgentPoolProfile)<br/><small>Optional</small>                                                                      |
| aiToolchainOperatorProfile | AI toolchain operator settings that apply to the whole cluster.                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAIToolchainOperatorProfile](#ManagedClusterAIToolchainOperatorProfile)<br/><small>Optional</small>                                                    |
| apiServerAccessProfile     | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile](#ManagedClusterAPIServerAccessProfile)<br/><small>Optional</small>                                                            |
| autoScalerProfile          | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile)<br/><small>Optional</small>                                                |
| autoUpgradeProfile         | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile)<br/><small>Optional</small>                                                                    |
| azureMonitorProfile        | Prometheus addon profile for the container service cluster                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAzureMonitorProfile](#ManagedClusterAzureMonitorProfile)<br/><small>Optional</small>                                                                  |
| azureName                  | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                                   |
| bootstrapProfile           | Profile of the cluster bootstrap configuration.                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterBootstrapProfile](#ManagedClusterBootstrapProfile)<br/><small>Optional</small>                                                                        |
| creationData               | CreationData to be used to specify the source Snapshot ID if the cluster will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                         | [CreationData](#CreationData)<br/><small>Optional</small>                                                                                                            |
| disableLocalAccounts       | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| diskEncryptionSetReference | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;diskEncryptionSets/&ZeroWidthSpace;{encryptionSetName}'                                                                                                | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| dnsPrefix                  | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                                   |
| enableNamespaceResources   | The default value is false. It can be enabled/disabled on creation and updating of the managed cluster. See [https://aka.ms/NamespaceARMResource](https://aka.ms/NamespaceARMResource) for more details on Namespace as a ARM Resource.                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                                                     |
| enablePodSecurityPolicy    | (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at https://aka.ms/k8s/psp and https://aka.ms/aks/psp.                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                                                                     |
| enableRBAC                 | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                                                     |
| extendedLocation           | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation](#ExtendedLocation)<br/><small>Optional</small>                                                                                                    |
| fqdnSubdomain              | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                                   |
| httpProxyConfig            | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig](#ManagedClusterHTTPProxyConfig)<br/><small>Optional</small>                                                                          |
| identity                   | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity](#ManagedClusterIdentity)<br/><small>Optional</small>                                                                                        |
| identityProfile            | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity](#UserAssignedIdentity)<br/><small>Optional</small>                                                                                 |
| ingressProfile             | Ingress profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterIngressProfile](#ManagedClusterIngressProfile)<br/><small>Optional</small>                                                                            |
| kind                       | This is primarily used to expose different UI experiences in the portal for different kinds                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |
| kubernetesVersion          | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string<br/><small>Optional</small>                                                                                                                                   |
| linuxProfile               | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile)<br/><small>Optional</small>                                                                            |
| location                   | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Required</small>                                                                                                                                   |
| metricsProfile             | Optional cluster metrics configuration.                                                                                                                                                                                                                                                                                                                                                     | [ManagedClusterMetricsProfile](#ManagedClusterMetricsProfile)<br/><small>Optional</small>                                                                            |
| networkProfile             | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile)<br/><small>Optional</small>                                                                        |
| nodeProvisioningProfile    | Node provisioning settings that apply to the whole cluster.                                                                                                                                                                                                                                                                                                                                 | [ManagedClusterNodeProvisioningProfile](#ManagedClusterNodeProvisioningProfile)<br/><small>Optional</small>                                                          |
| nodeResourceGroup          | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |
| nodeResourceGroupProfile   | The node resource group configuration profile.                                                                                                                                                                                                                                                                                                                                              | [ManagedClusterNodeResourceGroupProfile](#ManagedClusterNodeResourceGroupProfile)<br/><small>Optional</small>                                                        |
| oidcIssuerProfile          | The OIDC issuer profile of the Managed Cluster.                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterOIDCIssuerProfile](#ManagedClusterOIDCIssuerProfile)<br/><small>Optional</small>                                                                      |
| operatorSpec               | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                             | [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec)<br/><small>Optional</small>                                                                                |
| owner                      | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource                                                                                                | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| podIdentityProfile         | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile)<br/><small>Optional</small>                                                                    |
| privateLinkResources       | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource[]](#PrivateLinkResource)<br/><small>Optional</small>                                                                                            |
| publicNetworkAccess        | Allow or deny public network access for AKS                                                                                                                                                                                                                                                                                                                                                 | [ManagedClusterProperties_PublicNetworkAccess](#ManagedClusterProperties_PublicNetworkAccess)<br/><small>Optional</small>                                            |
| safeguardsProfile          | The Safeguards profile holds all the safeguards information for a given cluster                                                                                                                                                                                                                                                                                                             | [SafeguardsProfile](#SafeguardsProfile)<br/><small>Optional</small>                                                                                                  |
| securityProfile            | Security profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile)<br/><small>Optional</small>                                                                          |
| serviceMeshProfile         | Service mesh profile for a managed cluster.                                                                                                                                                                                                                                                                                                                                                 | [ServiceMeshProfile](#ServiceMeshProfile)<br/><small>Optional</small>                                                                                                |
| servicePrincipalProfile    | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile](#ManagedClusterServicePrincipalProfile)<br/><small>Optional</small>                                                          |
| sku                        | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU](#ManagedClusterSKU)<br/><small>Optional</small>                                                                                                  |
| storageProfile             | Storage profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterStorageProfile](#ManagedClusterStorageProfile)<br/><small>Optional</small>                                                                            |
| supportPlan                | The support plan for the Managed Cluster. If unspecified, the default is `KubernetesOfficial`.                                                                                                                                                                                                                                                                                              | [KubernetesSupportPlan](#KubernetesSupportPlan)<br/><small>Optional</small>                                                                                          |
| tags                       | Resource tags.                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                                                                                                        |
| upgradeSettings            | Settings for upgrading a cluster.                                                                                                                                                                                                                                                                                                                                                           | [ClusterUpgradeSettings](#ClusterUpgradeSettings)<br/><small>Optional</small>                                                                                        |
| windowsProfile             | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile)<br/><small>Optional</small>                                                                            |
| workloadAutoScalerProfile  | Workload Auto-scaler profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                       | [ManagedClusterWorkloadAutoScalerProfile](#ManagedClusterWorkloadAutoScalerProfile)<br/><small>Optional</small>                                                      |

ManagedCluster_STATUS{#ManagedCluster_STATUS}
---------------------------------------------

Managed cluster.

Used by: [ManagedCluster](#ManagedCluster).

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                    |
|----------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| aadProfile                 | The Azure Active Directory configuration.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterAADProfile_STATUS](#ManagedClusterAADProfile_STATUS)<br/><small>Optional</small>                                                         |
| addonProfiles              | The profile of managed cluster add-on.                                                                                                                                                                                                                                                                                                                                                      | [map[string]ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS)<br/><small>Optional</small>                                          |
| agentPoolProfiles          | The agent pool properties.                                                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAgentPoolProfile_STATUS[]](#ManagedClusterAgentPoolProfile_STATUS)<br/><small>Optional</small>                                           |
| aiToolchainOperatorProfile | AI toolchain operator settings that apply to the whole cluster.                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAIToolchainOperatorProfile_STATUS](#ManagedClusterAIToolchainOperatorProfile_STATUS)<br/><small>Optional</small>                         |
| apiServerAccessProfile     | The access profile for managed cluster API server.                                                                                                                                                                                                                                                                                                                                          | [ManagedClusterAPIServerAccessProfile_STATUS](#ManagedClusterAPIServerAccessProfile_STATUS)<br/><small>Optional</small>                                 |
| autoScalerProfile          | Parameters to be applied to the cluster-autoscaler when enabled                                                                                                                                                                                                                                                                                                                             | [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS)<br/><small>Optional</small>                     |
| autoUpgradeProfile         | The auto upgrade configuration.                                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS)<br/><small>Optional</small>                                         |
| azureMonitorProfile        | Prometheus addon profile for the container service cluster                                                                                                                                                                                                                                                                                                                                  | [ManagedClusterAzureMonitorProfile_STATUS](#ManagedClusterAzureMonitorProfile_STATUS)<br/><small>Optional</small>                                       |
| azurePortalFQDN            | The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                      |
| bootstrapProfile           | Profile of the cluster bootstrap configuration.                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterBootstrapProfile_STATUS](#ManagedClusterBootstrapProfile_STATUS)<br/><small>Optional</small>                                             |
| conditions                 | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                          | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| creationData               | CreationData to be used to specify the source Snapshot ID if the cluster will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                         | [CreationData_STATUS](#CreationData_STATUS)<br/><small>Optional</small>                                                                                 |
| currentKubernetesVersion   | The version of Kubernetes the Managed Cluster is running.                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| disableLocalAccounts       | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).                                                                                                              | bool<br/><small>Optional</small>                                                                                                                        |
| diskEncryptionSetID        | This is of the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;diskEncryptionSets/&ZeroWidthSpace;{encryptionSetName}'                                                                                                | string<br/><small>Optional</small>                                                                                                                      |
| dnsPrefix                  | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| enableNamespaceResources   | The default value is false. It can be enabled/disabled on creation and updating of the managed cluster. See [https://aka.ms/NamespaceARMResource](https://aka.ms/NamespaceARMResource) for more details on Namespace as a ARM Resource.                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                                        |
| enablePodSecurityPolicy    | (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at https://aka.ms/k8s/psp and https://aka.ms/aks/psp.                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                                                        |
| enableRBAC                 | Whether to enable Kubernetes Role-Based Access Control.                                                                                                                                                                                                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                                        |
| eTag                       | Unique read-only string used to implement optimistic concurrency. The eTag value will change when the resource is updated. Specify an if-match or if-none-match header with the eTag value for a subsequent request to enable optimistic concurrency per the normal etag convention.                                                                                                        | string<br/><small>Optional</small>                                                                                                                      |
| extendedLocation           | The extended location of the Virtual Machine.                                                                                                                                                                                                                                                                                                                                               | [ExtendedLocation_STATUS](#ExtendedLocation_STATUS)<br/><small>Optional</small>                                                                         |
| fqdn                       | The FQDN of the master pool.                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                      |
| fqdnSubdomain              | This cannot be updated once the Managed Cluster has been created.                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| httpProxyConfig            | Configurations for provisioning the cluster with HTTP proxy servers.                                                                                                                                                                                                                                                                                                                        | [ManagedClusterHTTPProxyConfig_STATUS](#ManagedClusterHTTPProxyConfig_STATUS)<br/><small>Optional</small>                                               |
| id                         | Fully qualified resource ID for the resource. E.g. "/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}"                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| identity                   | The identity of the managed cluster, if configured.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS)<br/><small>Optional</small>                                                             |
| identityProfile            | Identities associated with the cluster.                                                                                                                                                                                                                                                                                                                                                     | [map[string]UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small>                                                      |
| ingressProfile             | Ingress profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterIngressProfile_STATUS](#ManagedClusterIngressProfile_STATUS)<br/><small>Optional</small>                                                 |
| kind                       | This is primarily used to expose different UI experiences in the portal for different kinds                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| kubernetesVersion          | When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details. | string<br/><small>Optional</small>                                                                                                                      |
| linuxProfile               | The profile for Linux VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                           | [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS)<br/><small>Optional</small>                                                 |
| location                   | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| maxAgentPools              | The max number of agent pools for the managed cluster.                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                         |
| metricsProfile             | Optional cluster metrics configuration.                                                                                                                                                                                                                                                                                                                                                     | [ManagedClusterMetricsProfile_STATUS](#ManagedClusterMetricsProfile_STATUS)<br/><small>Optional</small>                                                 |
| name                       | The name of the resource                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| networkProfile             | The network configuration profile.                                                                                                                                                                                                                                                                                                                                                          | [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS)<br/><small>Optional</small>                                             |
| nodeProvisioningProfile    | Node provisioning settings that apply to the whole cluster.                                                                                                                                                                                                                                                                                                                                 | [ManagedClusterNodeProvisioningProfile_STATUS](#ManagedClusterNodeProvisioningProfile_STATUS)<br/><small>Optional</small>                               |
| nodeResourceGroup          | The name of the resource group containing agent pool nodes.                                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| nodeResourceGroupProfile   | The node resource group configuration profile.                                                                                                                                                                                                                                                                                                                                              | [ManagedClusterNodeResourceGroupProfile_STATUS](#ManagedClusterNodeResourceGroupProfile_STATUS)<br/><small>Optional</small>                             |
| oidcIssuerProfile          | The OIDC issuer profile of the Managed Cluster.                                                                                                                                                                                                                                                                                                                                             | [ManagedClusterOIDCIssuerProfile_STATUS](#ManagedClusterOIDCIssuerProfile_STATUS)<br/><small>Optional</small>                                           |
| podIdentityProfile         | See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.                                                                                                                                                                                                                                                | [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS)<br/><small>Optional</small>                                         |
| powerState                 | The Power State of the cluster.                                                                                                                                                                                                                                                                                                                                                             | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                                     |
| privateFQDN                | The FQDN of private cluster.                                                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                      |
| privateLinkResources       | Private link resources associated with the cluster.                                                                                                                                                                                                                                                                                                                                         | [PrivateLinkResource_STATUS[]](#PrivateLinkResource_STATUS)<br/><small>Optional</small>                                                                 |
| provisioningState          | The current provisioning state.                                                                                                                                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                      |
| publicNetworkAccess        | Allow or deny public network access for AKS                                                                                                                                                                                                                                                                                                                                                 | [ManagedClusterProperties_PublicNetworkAccess_STATUS](#ManagedClusterProperties_PublicNetworkAccess_STATUS)<br/><small>Optional</small>                 |
| resourceUID                | The resourceUID uniquely identifies ManagedClusters that reuse ARM ResourceIds (i.e: create, delete, create sequence)                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| safeguardsProfile          | The Safeguards profile holds all the safeguards information for a given cluster                                                                                                                                                                                                                                                                                                             | [SafeguardsProfile_STATUS](#SafeguardsProfile_STATUS)<br/><small>Optional</small>                                                                       |
| securityProfile            | Security profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                   | [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS)<br/><small>Optional</small>                                               |
| serviceMeshProfile         | Service mesh profile for a managed cluster.                                                                                                                                                                                                                                                                                                                                                 | [ServiceMeshProfile_STATUS](#ServiceMeshProfile_STATUS)<br/><small>Optional</small>                                                                     |
| servicePrincipalProfile    | Information about a service principal identity for the cluster to use for manipulating Azure APIs.                                                                                                                                                                                                                                                                                          | [ManagedClusterServicePrincipalProfile_STATUS](#ManagedClusterServicePrincipalProfile_STATUS)<br/><small>Optional</small>                               |
| sku                        | The managed cluster SKU.                                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS)<br/><small>Optional</small>                                                                       |
| storageProfile             | Storage profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                                    | [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS)<br/><small>Optional</small>                                                 |
| supportPlan                | The support plan for the Managed Cluster. If unspecified, the default is `KubernetesOfficial`.                                                                                                                                                                                                                                                                                              | [KubernetesSupportPlan_STATUS](#KubernetesSupportPlan_STATUS)<br/><small>Optional</small>                                                               |
| systemData                 | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags                       | Resource tags.                                                                                                                                                                                                                                                                                                                                                                              | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                       | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| upgradeSettings            | Settings for upgrading a cluster.                                                                                                                                                                                                                                                                                                                                                           | [ClusterUpgradeSettings_STATUS](#ClusterUpgradeSettings_STATUS)<br/><small>Optional</small>                                                             |
| windowsProfile             | The profile for Windows VMs in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                         | [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS)<br/><small>Optional</small>                                                 |
| workloadAutoScalerProfile  | Workload Auto-scaler profile for the managed cluster.                                                                                                                                                                                                                                                                                                                                       | [ManagedClusterWorkloadAutoScalerProfile_STATUS](#ManagedClusterWorkloadAutoScalerProfile_STATUS)<br/><small>Optional</small>                           |

ManagedClustersAgentPool_Spec{#ManagedClustersAgentPool_Spec}
-------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#ManagedClustersAgentPool).

| Property                          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                 |
|-----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile          | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile](#AgentPoolArtifactStreamingProfile)<br/><small>Optional</small>                                                                  |
| availabilityZones                 | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is `VirtualMachineScaleSets`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                                 |
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
| gatewayProfile                    | Profile specific to a managed agent pool in Gateway mode. This field cannot be set if agent pool mode is not Gateway.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [AgentPoolGatewayProfile](#AgentPoolGatewayProfile)<br/><small>Optional</small>                                                                                      |
| gpuInstanceProfile                | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile](#GPUInstanceProfile)<br/><small>Optional</small>                                                                                                |
| gpuProfile                        | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile](#AgentPoolGPUProfile)<br/><small>Optional</small>                                                                                              |
| hostGroupReference                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;hostGroups/&ZeroWidthSpace;{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                    | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| kubeletConfig                     | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig](#KubeletConfig)<br/><small>Optional</small>                                                                                                          |
| kubeletDiskType                   | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType](#KubeletDiskType)<br/><small>Optional</small>                                                                                                      |
| linuxOSConfig                     | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig](#LinuxOSConfig)<br/><small>Optional</small>                                                                                                          |
| maxCount                          | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                                      |
| maxPods                           | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                                      |
| messageOfTheDay                   | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                                   |
| minCount                          | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                                      |
| mode                              | A cluster must have at least one `System` Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode](#AgentPoolMode)<br/><small>Optional</small>                                                                                                          |
| networkProfile                    | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile](#AgentPoolNetworkProfile)<br/><small>Optional</small>                                                                                      |
| nodeInitializationTaints          | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                                 |
| nodeLabels                        | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                                                                        |
| nodePublicIPPrefixReference       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;publicIPPrefixes/&ZeroWidthSpace;{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| nodeTaints                        | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                                                                                 |
| operatorSpec                      | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ManagedClustersAgentPoolOperatorSpec](#ManagedClustersAgentPoolOperatorSpec)<br/><small>Optional</small>                                                            |
| orchestratorVersion               | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                                                                   |
| osDiskSizeGB                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ContainerServiceOSDisk](#ContainerServiceOSDisk)<br/><small>Optional</small>                                                                                        |
| osDiskType                        | The default is `Ephemeral` if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to `Managed`. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType](#OSDiskType)<br/><small>Optional</small>                                                                                                                |
| osSKU                             | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU](#OSSKU)<br/><small>Optional</small>                                                                                                                          |
| osType                            | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType](#OSType)<br/><small>Optional</small>                                                                                                                        |
| owner                             | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| podIPAllocationMode               | The IP allocation mode for pods in the agent pool. Must be used with podSubnetId. The default is `DynamicIndividual`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [PodIPAllocationMode](#PodIPAllocationMode)<br/><small>Optional</small>                                                                                              |
| podSubnetReference                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                                                                                                       | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| powerState                        | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState](#PowerState)<br/><small>Optional</small>                                                                                                                |
| proximityPlacementGroupReference  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| scaleDownMode                     | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode](#ScaleDownMode)<br/><small>Optional</small>                                                                                                          |
| scaleSetEvictionPolicy            | This cannot be specified unless the scaleSetPriority is `Spot`. If not specified, the default is `Delete`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)<br/><small>Optional</small>                                                                                        |
| scaleSetPriority                  | The Virtual Machine Scale Set priority. If not specified, the default is `Regular`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority](#ScaleSetPriority)<br/><small>Optional</small>                                                                                                    |
| securityProfile                   | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile](#AgentPoolSecurityProfile)<br/><small>Optional</small>                                                                                    |
| spotMaxPrice                      | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                                                                                  |
| tags                              | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                                                                        |
| type                              | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType](#AgentPoolType)<br/><small>Optional</small>                                                                                                          |
| upgradeSettings                   | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings)<br/><small>Optional</small>                                                                                    |
| virtualMachineNodesStatus         |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes[]](#VirtualMachineNodes)<br/><small>Optional</small>                                                                                            |
| virtualMachinesProfile            | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile](#VirtualMachinesProfile)<br/><small>Optional</small>                                                                                        |
| vmSize                            | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                                   |
| vnetSubnetReference               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| windowsProfile                    | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile](#AgentPoolWindowsProfile)<br/><small>Optional</small>                                                                                      |
| workloadRuntime                   | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime](#WorkloadRuntime)<br/><small>Optional</small>                                                                                                      |

ManagedClustersAgentPool_STATUS{#ManagedClustersAgentPool_STATUS}
-----------------------------------------------------------------

Used by: [ManagedClustersAgentPool](#ManagedClustersAgentPool).

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                    |
|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile   | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile_STATUS](#AgentPoolArtifactStreamingProfile_STATUS)<br/><small>Optional</small>                                       |
| availabilityZones          | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is `VirtualMachineScaleSets`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                    |
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
| eTag                       | Unique read-only string used to implement optimistic concurrency. The eTag value will change when the resource is updated. Specify an if-match or if-none-match header with the eTag value for a subsequent request to enable optimistic concurrency per the normal etag convention.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| gatewayProfile             | Profile specific to a managed agent pool in Gateway mode. This field cannot be set if agent pool mode is not Gateway.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [AgentPoolGatewayProfile_STATUS](#AgentPoolGatewayProfile_STATUS)<br/><small>Optional</small>                                                           |
| gpuInstanceProfile         | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)<br/><small>Optional</small>                                                                     |
| gpuProfile                 | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile_STATUS](#AgentPoolGPUProfile_STATUS)<br/><small>Optional</small>                                                                   |
| hostGroupID                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;hostGroups/&ZeroWidthSpace;{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| id                         | Resource ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| kubeletConfig              | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig_STATUS](#KubeletConfig_STATUS)<br/><small>Optional</small>                                                                               |
| kubeletDiskType            | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)<br/><small>Optional</small>                                                                           |
| linuxOSConfig              | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)<br/><small>Optional</small>                                                                               |
| maxCount                   | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                         |
| maxPods                    | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                         |
| messageOfTheDay            | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                      |
| minCount                   | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                         |
| mode                       | A cluster must have at least one `System` Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)<br/><small>Optional</small>                                                                               |
| name                       | The name of the resource that is unique within a resource group. This name can be used to access the resource.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                      |
| networkProfile             | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS)<br/><small>Optional</small>                                                           |
| nodeImageVersion           | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| nodeInitializationTaints   | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                    |
| nodeLabels                 | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                                                           |
| nodePublicIPPrefixID       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;publicIPPrefixes/&ZeroWidthSpace;{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| nodeTaints                 | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                                                                    |
| orchestratorVersion        | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                                                      |
| osDiskSizeGB               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | int<br/><small>Optional</small>                                                                                                                         |
| osDiskType                 | The default is `Ephemeral` if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to `Managed`. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType_STATUS](#OSDiskType_STATUS)<br/><small>Optional</small>                                                                                     |
| osSKU                      | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU_STATUS](#OSSKU_STATUS)<br/><small>Optional</small>                                                                                               |
| osType                     | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType_STATUS](#OSType_STATUS)<br/><small>Optional</small>                                                                                             |
| podIPAllocationMode        | The IP allocation mode for pods in the agent pool. Must be used with podSubnetId. The default is `DynamicIndividual`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [PodIPAllocationMode_STATUS](#PodIPAllocationMode_STATUS)<br/><small>Optional</small>                                                                   |
| podSubnetID                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| powerState                 | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                                                                     |
| properties_type            | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType_STATUS](#AgentPoolType_STATUS)<br/><small>Optional</small>                                                                               |
| provisioningState          | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| proximityPlacementGroupID  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| scaleDownMode              | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode_STATUS](#ScaleDownMode_STATUS)<br/><small>Optional</small>                                                                               |
| scaleSetEvictionPolicy     | This cannot be specified unless the scaleSetPriority is `Spot`. If not specified, the default is `Delete`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)<br/><small>Optional</small>                                                             |
| scaleSetPriority           | The Virtual Machine Scale Set priority. If not specified, the default is `Regular`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)<br/><small>Optional</small>                                                                         |
| securityProfile            | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile_STATUS](#AgentPoolSecurityProfile_STATUS)<br/><small>Optional</small>                                                         |
| spotMaxPrice               | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                                                                     |
| tags                       | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                       | Resource type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| upgradeSettings            | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS)<br/><small>Optional</small>                                                         |
| virtualMachineNodesStatus  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes_STATUS[]](#VirtualMachineNodes_STATUS)<br/><small>Optional</small>                                                                 |
| virtualMachinesProfile     | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile_STATUS](#VirtualMachinesProfile_STATUS)<br/><small>Optional</small>                                                             |
| vmSize                     | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| vnetSubnetID               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| windowsProfile             | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile_STATUS](#AgentPoolWindowsProfile_STATUS)<br/><small>Optional</small>                                                           |
| workloadRuntime            | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime_STATUS](#WorkloadRuntime_STATUS)<br/><small>Optional</small>                                                                           |

TrustedAccessRoleBinding_Spec{#TrustedAccessRoleBinding_Spec}
-------------------------------------------------------------

Used by: [TrustedAccessRoleBinding](#TrustedAccessRoleBinding).

| Property                | Description                                                                                                                                                                                                                                                                                          | Type                                                                                                                                                                 |
|-------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName               | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec            | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                      | [TrustedAccessRoleBindingOperatorSpec](#TrustedAccessRoleBindingOperatorSpec)<br/><small>Optional</small>                                                            |
| owner                   | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/ManagedCluster resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| roles                   | A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.                                                                                                                                                       | string[]<br/><small>Required</small>                                                                                                                                 |
| sourceResourceReference | The ARM resource ID of source resource that trusted access is configured for.                                                                                                                                                                                                                        | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Required</small>           |

TrustedAccessRoleBinding_STATUS{#TrustedAccessRoleBinding_STATUS}
-----------------------------------------------------------------

Used by: [TrustedAccessRoleBinding](#TrustedAccessRoleBinding).

| Property          | Description                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                    |
|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions        | The observed state of the resource                                                                                                                                                                                                                                                                                          | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| id                | Fully qualified resource ID for the resource. E.g. "/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}" | string<br/><small>Optional</small>                                                                                                                      |
| name              | The name of the resource                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState | The current provisioning state of trusted access role binding.                                                                                                                                                                                                                                                              | [TrustedAccessRoleBindingProperties_ProvisioningState_STATUS](#TrustedAccessRoleBindingProperties_ProvisioningState_STATUS)<br/><small>Optional</small> |
| roles             | A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.                                                                                                                                                                              | string[]<br/><small>Optional</small>                                                                                                                    |
| sourceResourceId  | The ARM resource ID of source resource that trusted access is configured for.                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| systemData        | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type              | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |

AgentPoolArtifactStreamingProfile{#AgentPoolArtifactStreamingProfile}
---------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property | Description                                                                                                                                                                                                                    | Type                             |
|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | Artifact streaming speeds up the cold-start of containers on a node through on-demand image loading. To use this feature, container images must also enable artifact streaming on ACR. If not specified, the default is false. | bool<br/><small>Optional</small> |

AgentPoolArtifactStreamingProfile_STATUS{#AgentPoolArtifactStreamingProfile_STATUS}
-----------------------------------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property | Description                                                                                                                                                                                                                    | Type                             |
|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | Artifact streaming speeds up the cold-start of containers on a node through on-demand image loading. To use this feature, container images must also enable artifact streaming on ACR. If not specified, the default is false. | bool<br/><small>Optional</small> |

AgentPoolGatewayProfile{#AgentPoolGatewayProfile}
-------------------------------------------------

Profile of the managed cluster gateway agent pool.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property           | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Type                            |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| publicIPPrefixSize | The Gateway agent pool associates one public IPPrefix for each static egress gateway to provide public egress. The size of Public IPPrefix should be selected by the user. Each node in the agent pool is assigned with one IP from the IPPrefix. The IPPrefix size thus serves as a cap on the size of the Gateway agent pool. Due to Azure public IPPrefix size limitation, the valid value range is [28, 31](/31 = 2 nodes/IPs, /30 = 4 nodes/IPs, /29 = 8 nodes/IPs, /28 = 16 nodes/IPs). The default value is 31. | int<br/><small>Optional</small> |

AgentPoolGatewayProfile_STATUS{#AgentPoolGatewayProfile_STATUS}
---------------------------------------------------------------

Profile of the managed cluster gateway agent pool.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property           | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Type                            |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| publicIPPrefixSize | The Gateway agent pool associates one public IPPrefix for each static egress gateway to provide public egress. The size of Public IPPrefix should be selected by the user. Each node in the agent pool is assigned with one IP from the IPPrefix. The IPPrefix size thus serves as a cap on the size of the Gateway agent pool. Due to Azure public IPPrefix size limitation, the valid value range is [28, 31](/31 = 2 nodes/IPs, /30 = 4 nodes/IPs, /29 = 8 nodes/IPs, /28 = 16 nodes/IPs). The default value is 31. | int<br/><small>Optional</small> |

AgentPoolGPUProfile{#AgentPoolGPUProfile}
-----------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property         | Description                                                                                                                                                                                                                                                                                                                                                          | Type                             |
|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| installGPUDriver | The default value is true when the vmSize of the agent pool contains a GPU, false otherwise. GPU Driver Installation can only be set true when VM has an associated GPU resource. Setting this field to false prevents automatic GPU driver installation. In that case, in order for the GPU to be usable, the user must perform GPU driver installation themselves. | bool<br/><small>Optional</small> |

AgentPoolGPUProfile_STATUS{#AgentPoolGPUProfile_STATUS}
-------------------------------------------------------

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property         | Description                                                                                                                                                                                                                                                                                                                                                          | Type                             |
|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| installGPUDriver | The default value is true when the vmSize of the agent pool contains a GPU, false otherwise. GPU Driver Installation can only be set true when VM has an associated GPU resource. Setting this field to false prevents automatic GPU driver installation. In that case, in order for the GPU to be usable, the user must perform GPU driver installation themselves. | bool<br/><small>Optional</small> |

AgentPoolMode{#AgentPoolMode}
-----------------------------

A cluster must have at least one `System` Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value     | Description |
|-----------|-------------|
| "Gateway" |             |
| "System"  |             |
| "User"    |             |

AgentPoolMode_STATUS{#AgentPoolMode_STATUS}
-------------------------------------------

A cluster must have at least one `System` Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value     | Description |
|-----------|-------------|
| "Gateway" |             |
| "System"  |             |
| "User"    |             |

AgentPoolNetworkProfile{#AgentPoolNetworkProfile}
-------------------------------------------------

Network settings of an agent pool.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property                            | Description                                                                              | Type                                                                                                                                                         |
|-------------------------------------|------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| allowedHostPorts                    | The port ranges that are allowed to access. The specified ranges are allowed to overlap. | [PortRange[]](#PortRange)<br/><small>Optional</small>                                                                                                        |
| applicationSecurityGroupsReferences | The IDs of the application security groups which agent pool will associate when created. | [genruntime.ResourceReference[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| nodePublicIPTags                    | IPTags of instance-level public IPs.                                                     | [IPTag[]](#IPTag)<br/><small>Optional</small>                                                                                                                |

AgentPoolNetworkProfile_STATUS{#AgentPoolNetworkProfile_STATUS}
---------------------------------------------------------------

Network settings of an agent pool.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property                  | Description                                                                              | Type                                                                |
|---------------------------|------------------------------------------------------------------------------------------|---------------------------------------------------------------------|
| allowedHostPorts          | The port ranges that are allowed to access. The specified ranges are allowed to overlap. | [PortRange_STATUS[]](#PortRange_STATUS)<br/><small>Optional</small> |
| applicationSecurityGroups | The IDs of the application security groups which agent pool will associate when created. | string[]<br/><small>Optional</small>                                |
| nodePublicIPTags          | IPTags of instance-level public IPs.                                                     | [IPTag_STATUS[]](#IPTag_STATUS)<br/><small>Optional</small>         |

AgentPoolSecurityProfile{#AgentPoolSecurityProfile}
---------------------------------------------------

The security settings of an agent pool.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property         | Description                                                                                                                                                                                                           | Type                                                                  |
|------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------|
| enableSecureBoot | Secure Boot is a feature of Trusted Launch which ensures that only signed operating systems and drivers can boot. For more details, see aka.ms/aks/trustedlaunch. If not specified, the default is false.             | bool<br/><small>Optional</small>                                      |
| enableVTPM       | vTPM is a Trusted Launch feature for configuring a dedicated secure vault for keys and measurements held locally on the node. For more details, see aka.ms/aks/trustedlaunch. If not specified, the default is false. | bool<br/><small>Optional</small>                                      |
| sshAccess        | SSH access method of an agent pool.                                                                                                                                                                                   | [AgentPoolSSHAccess](#AgentPoolSSHAccess)<br/><small>Optional</small> |

AgentPoolSecurityProfile_STATUS{#AgentPoolSecurityProfile_STATUS}
-----------------------------------------------------------------

The security settings of an agent pool.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property         | Description                                                                                                                                                                                                           | Type                                                                                |
|------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| enableSecureBoot | Secure Boot is a feature of Trusted Launch which ensures that only signed operating systems and drivers can boot. For more details, see aka.ms/aks/trustedlaunch. If not specified, the default is false.             | bool<br/><small>Optional</small>                                                    |
| enableVTPM       | vTPM is a Trusted Launch feature for configuring a dedicated secure vault for keys and measurements held locally on the node. For more details, see aka.ms/aks/trustedlaunch. If not specified, the default is false. | bool<br/><small>Optional</small>                                                    |
| sshAccess        | SSH access method of an agent pool.                                                                                                                                                                                   | [AgentPoolSSHAccess_STATUS](#AgentPoolSSHAccess_STATUS)<br/><small>Optional</small> |

AgentPoolType{#AgentPoolType}
-----------------------------

The type of Agent Pool.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value                     | Description |
|---------------------------|-------------|
| "AvailabilitySet"         |             |
| "VirtualMachineScaleSets" |             |
| "VirtualMachines"         |             |

AgentPoolType_STATUS{#AgentPoolType_STATUS}
-------------------------------------------

The type of Agent Pool.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value                     | Description |
|---------------------------|-------------|
| "AvailabilitySet"         |             |
| "VirtualMachineScaleSets" |             |
| "VirtualMachines"         |             |

AgentPoolUpgradeSettings{#AgentPoolUpgradeSettings}
---------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type                                                                                                                              |
|---------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------|
| drainTimeoutInMinutes     | The amount of time (in minutes) to wait on eviction of pods and graceful termination per node. This eviction wait time honors waiting on pod disruption budgets. If this time is exceeded, the upgrade fails. If not specified, the default is 30 minutes.                                                                                                                                                             | int<br/><small>Optional</small>                                                                                                   |
| maxSurge                  | This can either be set to an integer (e.g. `5`) or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string<br/><small>Optional</small>                                                                                                |
| nodeSoakDurationInMinutes | The amount of time (in minutes) to wait after draining a node and before reimaging it and moving on to next node. If not specified, the default is 0 minutes.                                                                                                                                                                                                                                                          | int<br/><small>Optional</small>                                                                                                   |
| undrainableNodeBehavior   | Defines the behavior for undrainable nodes during upgrade. The most common cause of undrainable nodes is Pod Disruption Budgets (PDBs), but other issues, such as pod termination grace period is exceeding the remaining per-node drain timeout or pod is still being in a running state, can also cause undrainable nodes.                                                                                           | [AgentPoolUpgradeSettings_UndrainableNodeBehavior](#AgentPoolUpgradeSettings_UndrainableNodeBehavior)<br/><small>Optional</small> |

AgentPoolUpgradeSettings_STATUS{#AgentPoolUpgradeSettings_STATUS}
-----------------------------------------------------------------

Settings for upgrading an agentpool

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Type                                                                                                                                            |
|---------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| drainTimeoutInMinutes     | The amount of time (in minutes) to wait on eviction of pods and graceful termination per node. This eviction wait time honors waiting on pod disruption budgets. If this time is exceeded, the upgrade fails. If not specified, the default is 30 minutes.                                                                                                                                                             | int<br/><small>Optional</small>                                                                                                                 |
| maxSurge                  | This can either be set to an integer (e.g. `5`) or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | string<br/><small>Optional</small>                                                                                                              |
| nodeSoakDurationInMinutes | The amount of time (in minutes) to wait after draining a node and before reimaging it and moving on to next node. If not specified, the default is 0 minutes.                                                                                                                                                                                                                                                          | int<br/><small>Optional</small>                                                                                                                 |
| undrainableNodeBehavior   | Defines the behavior for undrainable nodes during upgrade. The most common cause of undrainable nodes is Pod Disruption Budgets (PDBs), but other issues, such as pod termination grace period is exceeding the remaining per-node drain timeout or pod is still being in a running state, can also cause undrainable nodes.                                                                                           | [AgentPoolUpgradeSettings_UndrainableNodeBehavior_STATUS](#AgentPoolUpgradeSettings_UndrainableNodeBehavior_STATUS)<br/><small>Optional</small> |

AgentPoolWindowsProfile{#AgentPoolWindowsProfile}
-------------------------------------------------

The Windows agent pool's specific profile.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property           | Description                                                                                                                                                               | Type                             |
|--------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| disableOutboundNat | The default value is false. Outbound NAT can only be disabled if the cluster outboundType is NAT Gateway and the Windows agent pool does not have node public IP enabled. | bool<br/><small>Optional</small> |

AgentPoolWindowsProfile_STATUS{#AgentPoolWindowsProfile_STATUS}
---------------------------------------------------------------

The Windows agent pool's specific profile.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property           | Description                                                                                                                                                               | Type                             |
|--------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| disableOutboundNat | The default value is false. Outbound NAT can only be disabled if the cluster outboundType is NAT Gateway and the Windows agent pool does not have node public IP enabled. | bool<br/><small>Optional</small> |

ClusterUpgradeSettings{#ClusterUpgradeSettings}
-----------------------------------------------

Settings for upgrading a cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property         | Description             | Type                                                                            |
|------------------|-------------------------|---------------------------------------------------------------------------------|
| overrideSettings | Settings for overrides. | [UpgradeOverrideSettings](#UpgradeOverrideSettings)<br/><small>Optional</small> |

ClusterUpgradeSettings_STATUS{#ClusterUpgradeSettings_STATUS}
-------------------------------------------------------------

Settings for upgrading a cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property         | Description             | Type                                                                                          |
|------------------|-------------------------|-----------------------------------------------------------------------------------------------|
| overrideSettings | Settings for overrides. | [UpgradeOverrideSettings_STATUS](#UpgradeOverrideSettings_STATUS)<br/><small>Optional</small> |

ContainerServiceLinuxProfile{#ContainerServiceLinuxProfile}
-----------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property      | Description                                                 | Type                                                                                              |
|---------------|-------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string<br/><small>Required</small>                                                                |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration](#ContainerServiceSshConfiguration)<br/><small>Required</small> |

ContainerServiceLinuxProfile_STATUS{#ContainerServiceLinuxProfile_STATUS}
-------------------------------------------------------------------------

Profile for Linux VMs in the container service cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property      | Description                                                 | Type                                                                                                            |
|---------------|-------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| adminUsername | The administrator username to use for Linux VMs.            | string<br/><small>Optional</small>                                                                              |
| ssh           | The SSH configuration for Linux-based VMs running on Azure. | [ContainerServiceSshConfiguration_STATUS](#ContainerServiceSshConfiguration_STATUS)<br/><small>Optional</small> |

ContainerServiceNetworkProfile{#ContainerServiceNetworkProfile}
---------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                   | Description                                                                                                                                                                                                                                                                                                                                    | Type                                                                                                                          |
|----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| advancedNetworking         | Advanced Networking profile for enabling observability on a cluster. Note that enabling advanced networking features may incur additional costs. For more information see aka.ms/aksadvancednetworking.                                                                                                                                        | [AdvancedNetworking](#AdvancedNetworking)<br/><small>Optional</small>                                                         |
| dnsServiceIP               | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                            |
| ipFamilies                 | IP families are used to determine single-stack or dual-stack clusters. For single-stack, the expected value is IPv4. For dual-stack, the expected values are IPv4 and IPv6.                                                                                                                                                                    | [IpFamily[]](#IpFamily)<br/><small>Optional</small>                                                                           |
| kubeProxyConfig            | Holds configuration customizations for kube-proxy. Any values not defined will use the kube-proxy defaulting behavior. See https://v<version>.docs.kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/ where <version> is represented by a <major version>\-<minor version> string. Kubernetes version 1.23 would be '1-23'. | [ContainerServiceNetworkProfile_KubeProxyConfig](#ContainerServiceNetworkProfile_KubeProxyConfig)<br/><small>Optional</small> |
| loadBalancerProfile        | Profile of the cluster load balancer.                                                                                                                                                                                                                                                                                                          | [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile)<br/><small>Optional</small>                           |
| loadBalancerSku            | The default is `standard`. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs.                                                                                                                                                          | [LoadBalancerSku](#LoadBalancerSku)<br/><small>Optional</small>                                                               |
| natGatewayProfile          | Profile of the cluster NAT gateway.                                                                                                                                                                                                                                                                                                            | [ManagedClusterNATGatewayProfile](#ManagedClusterNATGatewayProfile)<br/><small>Optional</small>                               |
| networkDataplane           | Network dataplane used in the Kubernetes cluster.                                                                                                                                                                                                                                                                                              | [NetworkDataplane](#NetworkDataplane)<br/><small>Optional</small>                                                             |
| networkMode                | This cannot be specified if networkPlugin is anything other than `azure`.                                                                                                                                                                                                                                                                      | [NetworkMode](#NetworkMode)<br/><small>Optional</small>                                                                       |
| networkPlugin              | Network plugin used for building the Kubernetes network.                                                                                                                                                                                                                                                                                       | [NetworkPlugin](#NetworkPlugin)<br/><small>Optional</small>                                                                   |
| networkPluginMode          | Network plugin mode used for building the Kubernetes network.                                                                                                                                                                                                                                                                                  | [NetworkPluginMode](#NetworkPluginMode)<br/><small>Optional</small>                                                           |
| networkPolicy              | Network policy used for building the Kubernetes network.                                                                                                                                                                                                                                                                                       | [NetworkPolicy](#NetworkPolicy)<br/><small>Optional</small>                                                                   |
| outboundType               | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype).                                                                                                                                                          | [ContainerServiceNetworkProfile_OutboundType](#ContainerServiceNetworkProfile_OutboundType)<br/><small>Optional</small>       |
| podCidr                    | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                            |
| podCidrs                   | One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking.                                                                                                                                                                                                   | string[]<br/><small>Optional</small>                                                                                          |
| podLinkLocalAccess         | Defines access to special link local addresses (Azure Instance Metadata Service, aka IMDS) for pods with hostNetwork=false. if not specified, the default is `IMDS`.                                                                                                                                                                           | [PodLinkLocalAccess](#PodLinkLocalAccess)<br/><small>Optional</small>                                                         |
| serviceCidr                | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                            |
| serviceCidrs               | One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking. They must not overlap with any Subnet IP ranges.                                                                                                                                                  | string[]<br/><small>Optional</small>                                                                                          |
| staticEgressGatewayProfile | The profile for Static Egress Gateway addon. For more details about Static Egress Gateway, see https://aka.ms/aks/static-egress-gateway.                                                                                                                                                                                                       | [ManagedClusterStaticEgressGatewayProfile](#ManagedClusterStaticEgressGatewayProfile)<br/><small>Optional</small>             |

ContainerServiceNetworkProfile_STATUS{#ContainerServiceNetworkProfile_STATUS}
-----------------------------------------------------------------------------

Profile of network configuration.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                   | Description                                                                                                                                                                                                                                                                                                                                    | Type                                                                                                                                        |
|----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| advancedNetworking         | Advanced Networking profile for enabling observability on a cluster. Note that enabling advanced networking features may incur additional costs. For more information see aka.ms/aksadvancednetworking.                                                                                                                                        | [AdvancedNetworking_STATUS](#AdvancedNetworking_STATUS)<br/><small>Optional</small>                                                         |
| dnsServiceIP               | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                          |
| ipFamilies                 | IP families are used to determine single-stack or dual-stack clusters. For single-stack, the expected value is IPv4. For dual-stack, the expected values are IPv4 and IPv6.                                                                                                                                                                    | [IpFamily_STATUS[]](#IpFamily_STATUS)<br/><small>Optional</small>                                                                           |
| kubeProxyConfig            | Holds configuration customizations for kube-proxy. Any values not defined will use the kube-proxy defaulting behavior. See https://v<version>.docs.kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/ where <version> is represented by a <major version>\-<minor version> string. Kubernetes version 1.23 would be '1-23'. | [ContainerServiceNetworkProfile_KubeProxyConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_STATUS)<br/><small>Optional</small> |
| loadBalancerProfile        | Profile of the cluster load balancer.                                                                                                                                                                                                                                                                                                          | [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS)<br/><small>Optional</small>                           |
| loadBalancerSku            | The default is `standard`. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs.                                                                                                                                                          | [LoadBalancerSku_STATUS](#LoadBalancerSku_STATUS)<br/><small>Optional</small>                                                               |
| natGatewayProfile          | Profile of the cluster NAT gateway.                                                                                                                                                                                                                                                                                                            | [ManagedClusterNATGatewayProfile_STATUS](#ManagedClusterNATGatewayProfile_STATUS)<br/><small>Optional</small>                               |
| networkDataplane           | Network dataplane used in the Kubernetes cluster.                                                                                                                                                                                                                                                                                              | [NetworkDataplane_STATUS](#NetworkDataplane_STATUS)<br/><small>Optional</small>                                                             |
| networkMode                | This cannot be specified if networkPlugin is anything other than `azure`.                                                                                                                                                                                                                                                                      | [NetworkMode_STATUS](#NetworkMode_STATUS)<br/><small>Optional</small>                                                                       |
| networkPlugin              | Network plugin used for building the Kubernetes network.                                                                                                                                                                                                                                                                                       | [NetworkPlugin_STATUS](#NetworkPlugin_STATUS)<br/><small>Optional</small>                                                                   |
| networkPluginMode          | Network plugin mode used for building the Kubernetes network.                                                                                                                                                                                                                                                                                  | [NetworkPluginMode_STATUS](#NetworkPluginMode_STATUS)<br/><small>Optional</small>                                                           |
| networkPolicy              | Network policy used for building the Kubernetes network.                                                                                                                                                                                                                                                                                       | [NetworkPolicy_STATUS](#NetworkPolicy_STATUS)<br/><small>Optional</small>                                                                   |
| outboundType               | This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype).                                                                                                                                                          | [ContainerServiceNetworkProfile_OutboundType_STATUS](#ContainerServiceNetworkProfile_OutboundType_STATUS)<br/><small>Optional</small>       |
| podCidr                    | A CIDR notation IP range from which to assign pod IPs when kubenet is used.                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                          |
| podCidrs                   | One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking.                                                                                                                                                                                                   | string[]<br/><small>Optional</small>                                                                                                        |
| podLinkLocalAccess         | Defines access to special link local addresses (Azure Instance Metadata Service, aka IMDS) for pods with hostNetwork=false. if not specified, the default is `IMDS`.                                                                                                                                                                           | [PodLinkLocalAccess_STATUS](#PodLinkLocalAccess_STATUS)<br/><small>Optional</small>                                                         |
| serviceCidr                | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                          |
| serviceCidrs               | One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking. They must not overlap with any Subnet IP ranges.                                                                                                                                                  | string[]<br/><small>Optional</small>                                                                                                        |
| staticEgressGatewayProfile | The profile for Static Egress Gateway addon. For more details about Static Egress Gateway, see https://aka.ms/aks/static-egress-gateway.                                                                                                                                                                                                       | [ManagedClusterStaticEgressGatewayProfile_STATUS](#ManagedClusterStaticEgressGatewayProfile_STATUS)<br/><small>Optional</small>             |

ContainerServiceOSDisk{#ContainerServiceOSDisk}
-----------------------------------------------

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

CreationData{#CreationData}
---------------------------

Data used when creating a target resource from a source resource.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec), [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property                | Description                                                                     | Type                                                                                                                                                       |
|-------------------------|---------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| sourceResourceReference | This is the ARM ID of the source object to be used to create the target object. | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

CreationData_STATUS{#CreationData_STATUS}
-----------------------------------------

Data used when creating a target resource from a source resource.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property         | Description                                                                     | Type                               |
|------------------|---------------------------------------------------------------------------------|------------------------------------|
| sourceResourceId | This is the ARM ID of the source object to be used to create the target object. | string<br/><small>Optional</small> |

ExtendedLocation{#ExtendedLocation}
-----------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                        | Type                                                                      |
|----------|------------------------------------|---------------------------------------------------------------------------|
| name     | The name of the extended location. | string<br/><small>Optional</small>                                        |
| type     | The type of the extended location. | [ExtendedLocationType](#ExtendedLocationType)<br/><small>Optional</small> |

ExtendedLocation_STATUS{#ExtendedLocation_STATUS}
-------------------------------------------------

The complex type of the extended location.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                        | Type                                                                                    |
|----------|------------------------------------|-----------------------------------------------------------------------------------------|
| name     | The name of the extended location. | string<br/><small>Optional</small>                                                      |
| type     | The type of the extended location. | [ExtendedLocationType_STATUS](#ExtendedLocationType_STATUS)<br/><small>Optional</small> |

GPUInstanceProfile{#GPUInstanceProfile}
---------------------------------------

GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value   | Description |
|---------|-------------|
| "MIG1g" |             |
| "MIG2g" |             |
| "MIG3g" |             |
| "MIG4g" |             |
| "MIG7g" |             |

GPUInstanceProfile_STATUS{#GPUInstanceProfile_STATUS}
-----------------------------------------------------

GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value   | Description |
|---------|-------------|
| "MIG1g" |             |
| "MIG2g" |             |
| "MIG3g" |             |
| "MIG4g" |             |
| "MIG7g" |             |

KubeletConfig{#KubeletConfig}
-----------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property              | Description                                                                                                                                                                                                                          | Type                                 |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| allowedUnsafeSysctls  | Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in `*`).                                                                                                                                                            | string[]<br/><small>Optional</small> |
| containerLogMaxFiles  | The maximum number of container log files that can be present for a container. The number must be ≥ 2.                                                                                                                               | int<br/><small>Optional</small>      |
| containerLogMaxSizeMB | The maximum size (e.g. 10Mi) of container log file before it is rotated.                                                                                                                                                             | int<br/><small>Optional</small>      |
| cpuCfsQuota           | The default is true.                                                                                                                                                                                                                 | bool<br/><small>Optional</small>     |
| cpuCfsQuotaPeriod     | The default is `100ms.` Valid values are a sequence of decimal numbers with an optional fraction and a unit suffix. For example: `300ms`, `2h45m`. Supported units are `ns`, `us`, `ms`, `s`, `m`, and `h`.                          | string<br/><small>Optional</small>   |
| cpuManagerPolicy      | The default is `none`. See [Kubernetes CPU management policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies) for more information. Allowed values are `none` and `static`.  | string<br/><small>Optional</small>   |
| failSwapOn            | If set to true it will make the Kubelet fail to start if swap is enabled on the node.                                                                                                                                                | bool<br/><small>Optional</small>     |
| imageGcHighThreshold  | To disable image garbage collection, set to 100. The default is 85%                                                                                                                                                                  | int<br/><small>Optional</small>      |
| imageGcLowThreshold   | This cannot be set higher than imageGcHighThreshold. The default is 80%                                                                                                                                                              | int<br/><small>Optional</small>      |
| podMaxPids            | The maximum number of processes per pod.                                                                                                                                                                                             | int<br/><small>Optional</small>      |
| topologyManagerPolicy | For more information see [Kubernetes Topology Manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager). The default is `none`. Allowed values are `none`, 'best-effort', `restricted`, and 'single-numa-node'. | string<br/><small>Optional</small>   |

KubeletConfig_STATUS{#KubeletConfig_STATUS}
-------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property              | Description                                                                                                                                                                                                                          | Type                                 |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| allowedUnsafeSysctls  | Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in `*`).                                                                                                                                                            | string[]<br/><small>Optional</small> |
| containerLogMaxFiles  | The maximum number of container log files that can be present for a container. The number must be ≥ 2.                                                                                                                               | int<br/><small>Optional</small>      |
| containerLogMaxSizeMB | The maximum size (e.g. 10Mi) of container log file before it is rotated.                                                                                                                                                             | int<br/><small>Optional</small>      |
| cpuCfsQuota           | The default is true.                                                                                                                                                                                                                 | bool<br/><small>Optional</small>     |
| cpuCfsQuotaPeriod     | The default is `100ms.` Valid values are a sequence of decimal numbers with an optional fraction and a unit suffix. For example: `300ms`, `2h45m`. Supported units are `ns`, `us`, `ms`, `s`, `m`, and `h`.                          | string<br/><small>Optional</small>   |
| cpuManagerPolicy      | The default is `none`. See [Kubernetes CPU management policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies) for more information. Allowed values are `none` and `static`.  | string<br/><small>Optional</small>   |
| failSwapOn            | If set to true it will make the Kubelet fail to start if swap is enabled on the node.                                                                                                                                                | bool<br/><small>Optional</small>     |
| imageGcHighThreshold  | To disable image garbage collection, set to 100. The default is 85%                                                                                                                                                                  | int<br/><small>Optional</small>      |
| imageGcLowThreshold   | This cannot be set higher than imageGcHighThreshold. The default is 80%                                                                                                                                                              | int<br/><small>Optional</small>      |
| podMaxPids            | The maximum number of processes per pod.                                                                                                                                                                                             | int<br/><small>Optional</small>      |
| topologyManagerPolicy | For more information see [Kubernetes Topology Manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager). The default is `none`. Allowed values are `none`, 'best-effort', `restricted`, and 'single-numa-node'. | string<br/><small>Optional</small>   |

KubeletDiskType{#KubeletDiskType}
---------------------------------

Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value       | Description |
|-------------|-------------|
| "OS"        |             |
| "Temporary" |             |

KubeletDiskType_STATUS{#KubeletDiskType_STATUS}
-----------------------------------------------

Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value       | Description |
|-------------|-------------|
| "OS"        |             |
| "Temporary" |             |

KubernetesSupportPlan{#KubernetesSupportPlan}
---------------------------------------------

Different support tiers for AKS managed clusters

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Value                | Description |
|----------------------|-------------|
| "AKSLongTermSupport" |             |
| "KubernetesOfficial" |             |

KubernetesSupportPlan_STATUS{#KubernetesSupportPlan_STATUS}
-----------------------------------------------------------

Different support tiers for AKS managed clusters

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Value                | Description |
|----------------------|-------------|
| "AKSLongTermSupport" |             |
| "KubernetesOfficial" |             |

LinuxOSConfig{#LinuxOSConfig}
-----------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property                   | Description                                                                                                                                                                                                                                         | Type                                                      |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int<br/><small>Optional</small>                           |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig](#SysctlConfig)<br/><small>Optional</small> |
| transparentHugePageDefrag  | Valid values are `always`, `defer`, 'defer+madvise', `madvise` and `never`. The default is `madvise`. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string<br/><small>Optional</small>                        |
| transparentHugePageEnabled | Valid values are `always`, `madvise`, and `never`. The default is `always`. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string<br/><small>Optional</small>                        |

LinuxOSConfig_STATUS{#LinuxOSConfig_STATUS}
-------------------------------------------

See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property                   | Description                                                                                                                                                                                                                                         | Type                                                                    |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| swapFileSizeMB             | The size in MB of a swap file that will be created on each node.                                                                                                                                                                                    | int<br/><small>Optional</small>                                         |
| sysctls                    | Sysctl settings for Linux agent nodes.                                                                                                                                                                                                              | [SysctlConfig_STATUS](#SysctlConfig_STATUS)<br/><small>Optional</small> |
| transparentHugePageDefrag  | Valid values are `always`, `defer`, 'defer+madvise', `madvise` and `never`. The default is `madvise`. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge). | string<br/><small>Optional</small>                                      |
| transparentHugePageEnabled | Valid values are `always`, `madvise`, and `never`. The default is `always`. For more information see [Transparent Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).                           | string<br/><small>Optional</small>                                      |

ManagedClusterAADProfile{#ManagedClusterAADProfile}
---------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property            | Description                                                                                                        | Type                                 |
|---------------------|--------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[]<br/><small>Optional</small> |
| clientAppID         | (DEPRECATED) The client AAD application ID. Learn more at https://aka.ms/aks/aad-legacy.                           | string<br/><small>Optional</small>   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool<br/><small>Optional</small>     |
| managed             | Whether to enable managed AAD.                                                                                     | bool<br/><small>Optional</small>     |
| serverAppID         | (DEPRECATED) The server AAD application ID. Learn more at https://aka.ms/aks/aad-legacy.                           | string<br/><small>Optional</small>   |
| serverAppSecret     | (DEPRECATED) The server AAD application secret. Learn more at https://aka.ms/aks/aad-legacy.                       | string<br/><small>Optional</small>   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string<br/><small>Optional</small>   |

ManagedClusterAADProfile_STATUS{#ManagedClusterAADProfile_STATUS}
-----------------------------------------------------------------

For more details see [managed AAD on AKS](https://docs.microsoft.com/azure/aks/managed-aad).

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property            | Description                                                                                                        | Type                                 |
|---------------------|--------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| adminGroupObjectIDs | The list of AAD group object IDs that will have admin role of the cluster.                                         | string[]<br/><small>Optional</small> |
| clientAppID         | (DEPRECATED) The client AAD application ID. Learn more at https://aka.ms/aks/aad-legacy.                           | string<br/><small>Optional</small>   |
| enableAzureRBAC     | Whether to enable Azure RBAC for Kubernetes authorization.                                                         | bool<br/><small>Optional</small>     |
| managed             | Whether to enable managed AAD.                                                                                     | bool<br/><small>Optional</small>     |
| serverAppID         | (DEPRECATED) The server AAD application ID. Learn more at https://aka.ms/aks/aad-legacy.                           | string<br/><small>Optional</small>   |
| serverAppSecret     | (DEPRECATED) The server AAD application secret. Learn more at https://aka.ms/aks/aad-legacy.                       | string<br/><small>Optional</small>   |
| tenantID            | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | string<br/><small>Optional</small>   |

ManagedClusterAddonProfile{#ManagedClusterAddonProfile}
-------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                | Type                                          |
|----------|--------------------------------------------|-----------------------------------------------|
| config   | Key-value pairs for configuring an add-on. | map[string]string<br/><small>Optional</small> |
| enabled  | Whether the add-on is enabled or not.      | bool<br/><small>Required</small>              |

ManagedClusterAddonProfile_STATUS{#ManagedClusterAddonProfile_STATUS}
---------------------------------------------------------------------

A Kubernetes add-on profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                                                | Type                                                                                    |
|----------|------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| config   | Key-value pairs for configuring an add-on.                 | map[string]string<br/><small>Optional</small>                                           |
| enabled  | Whether the add-on is enabled or not.                      | bool<br/><small>Optional</small>                                                        |
| identity | Information of user assigned identity used by this add-on. | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small> |

ManagedClusterAgentPoolProfile{#ManagedClusterAgentPoolProfile}
---------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                       |
|-----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile          | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile](#AgentPoolArtifactStreamingProfile)<br/><small>Optional</small>                                                        |
| availabilityZones                 | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is `VirtualMachineScaleSets`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                       |
| capacityReservationGroupReference | AKS will associate the specified agent pool with the Capacity Reservation Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| count                             | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | int<br/><small>Optional</small>                                                                                                                            |
| creationData                      | CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [CreationData](#CreationData)<br/><small>Optional</small>                                                                                                  |
| enableAutoScaling                 | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                           |
| enableCustomCATrust               | When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded certificates into node trust stores. Defaults to false.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                           |
| enableEncryptionAtHost            | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                           |
| enableFIPS                        | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | bool<br/><small>Optional</small>                                                                                                                           |
| enableNodePublicIP                | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                                                                                                                                                                                                                                                                                                                                            | bool<br/><small>Optional</small>                                                                                                                           |
| enableUltraSSD                    | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                           |
| gatewayProfile                    | Profile specific to a managed agent pool in Gateway mode. This field cannot be set if agent pool mode is not Gateway.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [AgentPoolGatewayProfile](#AgentPoolGatewayProfile)<br/><small>Optional</small>                                                                            |
| gpuInstanceProfile                | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile](#GPUInstanceProfile)<br/><small>Optional</small>                                                                                      |
| gpuProfile                        | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile](#AgentPoolGPUProfile)<br/><small>Optional</small>                                                                                    |
| hostGroupReference                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;hostGroups/&ZeroWidthSpace;{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                    | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| kubeletConfig                     | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig](#KubeletConfig)<br/><small>Optional</small>                                                                                                |
| kubeletDiskType                   | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType](#KubeletDiskType)<br/><small>Optional</small>                                                                                            |
| linuxOSConfig                     | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig](#LinuxOSConfig)<br/><small>Optional</small>                                                                                                |
| maxCount                          | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                            |
| maxPods                           | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                                                            |
| messageOfTheDay                   | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                         |
| minCount                          | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                                                            |
| mode                              | A cluster must have at least one `System` Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode](#AgentPoolMode)<br/><small>Optional</small>                                                                                                |
| name                              | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | string<br/><small>Required</small>                                                                                                                         |
| networkProfile                    | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile](#AgentPoolNetworkProfile)<br/><small>Optional</small>                                                                            |
| nodeInitializationTaints          | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                                                                       |
| nodeLabels                        | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                                                              |
| nodePublicIPPrefixReference       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;publicIPPrefixes/&ZeroWidthSpace;{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| nodeTaints                        | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                                                                       |
| orchestratorVersion               | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                                                         |
| osDiskSizeGB                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [ContainerServiceOSDisk](#ContainerServiceOSDisk)<br/><small>Optional</small>                                                                              |
| osDiskType                        | The default is `Ephemeral` if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to `Managed`. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType](#OSDiskType)<br/><small>Optional</small>                                                                                                      |
| osSKU                             | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU](#OSSKU)<br/><small>Optional</small>                                                                                                                |
| osType                            | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType](#OSType)<br/><small>Optional</small>                                                                                                              |
| podIPAllocationMode               | The IP allocation mode for pods in the agent pool. Must be used with podSubnetId. The default is `DynamicIndividual`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [PodIPAllocationMode](#PodIPAllocationMode)<br/><small>Optional</small>                                                                                    |
| podSubnetReference                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                                                                                                       | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| powerState                        | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState](#PowerState)<br/><small>Optional</small>                                                                                                      |
| proximityPlacementGroupReference  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| scaleDownMode                     | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode](#ScaleDownMode)<br/><small>Optional</small>                                                                                                |
| scaleSetEvictionPolicy            | This cannot be specified unless the scaleSetPriority is `Spot`. If not specified, the default is `Delete`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy](#ScaleSetEvictionPolicy)<br/><small>Optional</small>                                                                              |
| scaleSetPriority                  | The Virtual Machine Scale Set priority. If not specified, the default is `Regular`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority](#ScaleSetPriority)<br/><small>Optional</small>                                                                                          |
| securityProfile                   | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile](#AgentPoolSecurityProfile)<br/><small>Optional</small>                                                                          |
| spotMaxPrice                      | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                                                                        |
| tags                              | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                                                              |
| type                              | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType](#AgentPoolType)<br/><small>Optional</small>                                                                                                |
| upgradeSettings                   | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings)<br/><small>Optional</small>                                                                          |
| virtualMachineNodesStatus         |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes[]](#VirtualMachineNodes)<br/><small>Optional</small>                                                                                  |
| virtualMachinesProfile            | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile](#VirtualMachinesProfile)<br/><small>Optional</small>                                                                              |
| vmSize                            | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                         |
| vnetSubnetReference               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                               | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| windowsProfile                    | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile](#AgentPoolWindowsProfile)<br/><small>Optional</small>                                                                            |
| workloadRuntime                   | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime](#WorkloadRuntime)<br/><small>Optional</small>                                                                                            |

ManagedClusterAgentPoolProfile_STATUS{#ManagedClusterAgentPoolProfile_STATUS}
-----------------------------------------------------------------------------

Profile for the container service agent pool.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                              |
|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| artifactStreamingProfile   | Configuration for using artifact streaming on AKS.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolArtifactStreamingProfile_STATUS](#AgentPoolArtifactStreamingProfile_STATUS)<br/><small>Optional</small> |
| availabilityZones          | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is `VirtualMachineScaleSets`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                              |
| capacityReservationGroupID | AKS will associate the specified agent pool with the Capacity Reservation Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                                                |
| count                      | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | int<br/><small>Optional</small>                                                                                   |
| creationData               | CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using a snapshot.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [CreationData_STATUS](#CreationData_STATUS)<br/><small>Optional</small>                                           |
| currentOrchestratorVersion | If orchestratorVersion was a fully specified version <major.minor.patch>, this field will be exactly equal to it. If orchestratorVersion was <major.minor>, this field will contain the full <major.minor.patch> version being used.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                |
| enableAutoScaling          | Whether to enable auto-scaler                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                  |
| enableCustomCATrust        | When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded certificates into node trust stores. Defaults to false.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                  |
| enableEncryptionAtHost     | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                  |
| enableFIPS                 | See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | bool<br/><small>Optional</small>                                                                                  |
| enableNodePublicIP         | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.                                                                                                                                                                                                                                                                                                                                                                            | bool<br/><small>Optional</small>                                                                                  |
| enableUltraSSD             | Whether to enable UltraSSD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                  |
| eTag                       | Unique read-only string used to implement optimistic concurrency. The eTag value will change when the resource is updated. Specify an if-match or if-none-match header with the eTag value for a subsequent request to enable optimistic concurrency per the normal etag convention.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                |
| gatewayProfile             | Profile specific to a managed agent pool in Gateway mode. This field cannot be set if agent pool mode is not Gateway.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [AgentPoolGatewayProfile_STATUS](#AgentPoolGatewayProfile_STATUS)<br/><small>Optional</small>                     |
| gpuInstanceProfile         | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [GPUInstanceProfile_STATUS](#GPUInstanceProfile_STATUS)<br/><small>Optional</small>                               |
| gpuProfile                 | The GPU settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [AgentPoolGPUProfile_STATUS](#AgentPoolGPUProfile_STATUS)<br/><small>Optional</small>                             |
| hostGroupID                | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Compute/&ZeroWidthSpace;hostGroups/&ZeroWidthSpace;{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                |
| kubeletConfig              | The Kubelet configuration on the agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [KubeletConfig_STATUS](#KubeletConfig_STATUS)<br/><small>Optional</small>                                         |
| kubeletDiskType            | Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [KubeletDiskType_STATUS](#KubeletDiskType_STATUS)<br/><small>Optional</small>                                     |
| linuxOSConfig              | The OS configuration of Linux agent nodes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS)<br/><small>Optional</small>                                         |
| maxCount                   | The maximum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                   |
| maxPods                    | The maximum number of pods that can run on a node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | int<br/><small>Optional</small>                                                                                   |
| messageOfTheDay            | A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e., will be printed raw and not be executed as a script).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                |
| minCount                   | The minimum number of nodes for auto-scaling                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | int<br/><small>Optional</small>                                                                                   |
| mode                       | A cluster must have at least one `System` Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolMode_STATUS](#AgentPoolMode_STATUS)<br/><small>Optional</small>                                         |
| name                       | Windows agent pool names must be 6 characters or less.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                |
| networkProfile             | Network-related settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS)<br/><small>Optional</small>                     |
| nodeImageVersion           | The version of node image                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                |
| nodeInitializationTaints   | These taints will not be reconciled by AKS and can be removed with a kubectl call. This field can be modified after node pool is created, but nodes will not be recreated with new taints until another operation that requires recreation (e.g. node image upgrade) happens. These taints allow for required configuration to run before the node is ready to accept workloads, for example 'key1=value1:NoSchedule' that then can be removed with `kubectl taint nodes node1 key1=value1:NoSchedule-`                                                                                                                                                                                                                                                                                                                 | string[]<br/><small>Optional</small>                                                                              |
| nodeLabels                 | The node labels to be persisted across all nodes in agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | map[string]string<br/><small>Optional</small>                                                                     |
| nodePublicIPPrefixID       | This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;publicIPPrefixes/&ZeroWidthSpace;{publicIPPrefixName}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                |
| nodeTaints                 | The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | string[]<br/><small>Optional</small>                                                                              |
| orchestratorVersion        | Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same <major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool). | string<br/><small>Optional</small>                                                                                |
| osDiskSizeGB               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | int<br/><small>Optional</small>                                                                                   |
| osDiskType                 | The default is `Ephemeral` if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to `Managed`. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSDiskType_STATUS](#OSDiskType_STATUS)<br/><small>Optional</small>                                               |
| osSKU                      | Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [OSSKU_STATUS](#OSSKU_STATUS)<br/><small>Optional</small>                                                         |
| osType                     | The operating system type. The default is Linux.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [OSType_STATUS](#OSType_STATUS)<br/><small>Optional</small>                                                       |
| podIPAllocationMode        | The IP allocation mode for pods in the agent pool. Must be used with podSubnetId. The default is `DynamicIndividual`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [PodIPAllocationMode_STATUS](#PodIPAllocationMode_STATUS)<br/><small>Optional</small>                             |
| podSubnetID                | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                |
| powerState                 | When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [PowerState_STATUS](#PowerState_STATUS)<br/><small>Optional</small>                                               |
| provisioningState          | The current deployment or provisioning state.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                |
| proximityPlacementGroupID  | The ID for Proximity Placement Group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                |
| scaleDownMode              | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [ScaleDownMode_STATUS](#ScaleDownMode_STATUS)<br/><small>Optional</small>                                         |
| scaleSetEvictionPolicy     | This cannot be specified unless the scaleSetPriority is `Spot`. If not specified, the default is `Delete`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [ScaleSetEvictionPolicy_STATUS](#ScaleSetEvictionPolicy_STATUS)<br/><small>Optional</small>                       |
| scaleSetPriority           | The Virtual Machine Scale Set priority. If not specified, the default is `Regular`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [ScaleSetPriority_STATUS](#ScaleSetPriority_STATUS)<br/><small>Optional</small>                                   |
| securityProfile            | The security settings of an agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolSecurityProfile_STATUS](#AgentPoolSecurityProfile_STATUS)<br/><small>Optional</small>                   |
| spotMaxPrice               | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | float64<br/><small>Optional</small>                                                                               |
| tags                       | The tags to be persisted on the agent pool virtual machine scale set.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | map[string]string<br/><small>Optional</small>                                                                     |
| type                       | The type of Agent Pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [AgentPoolType_STATUS](#AgentPoolType_STATUS)<br/><small>Optional</small>                                         |
| upgradeSettings            | Settings for upgrading the agentpool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS)<br/><small>Optional</small>                   |
| virtualMachineNodesStatus  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [VirtualMachineNodes_STATUS[]](#VirtualMachineNodes_STATUS)<br/><small>Optional</small>                           |
| virtualMachinesProfile     | Specifications on VirtualMachines agent pool.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [VirtualMachinesProfile_STATUS](#VirtualMachinesProfile_STATUS)<br/><small>Optional</small>                       |
| vmSize                     | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                |
| vnetSubnetID               | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Network/&ZeroWidthSpace;virtualNetworks/&ZeroWidthSpace;{virtualNetworkName}/&ZeroWidthSpace;subnets/&ZeroWidthSpace;{subnetName}                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                |
| windowsProfile             | The Windows agent pool's specific profile.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [AgentPoolWindowsProfile_STATUS](#AgentPoolWindowsProfile_STATUS)<br/><small>Optional</small>                     |
| workloadRuntime            | Determines the type of workload a node can run.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | [WorkloadRuntime_STATUS](#WorkloadRuntime_STATUS)<br/><small>Optional</small>                                     |

ManagedClusterAIToolchainOperatorProfile{#ManagedClusterAIToolchainOperatorProfile}
-----------------------------------------------------------------------------------

When enabling the operator, a set of AKS managed CRDs and controllers will be installed in the cluster. The operator automates the deployment of OSS models for inference and/or training purposes. It provides a set of preset models and enables distributed inference against them.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                        | Type                             |
|----------|----------------------------------------------------|----------------------------------|
| enabled  | Indicates if AI toolchain operator enabled or not. | bool<br/><small>Optional</small> |

ManagedClusterAIToolchainOperatorProfile_STATUS{#ManagedClusterAIToolchainOperatorProfile_STATUS}
-------------------------------------------------------------------------------------------------

When enabling the operator, a set of AKS managed CRDs and controllers will be installed in the cluster. The operator automates the deployment of OSS models for inference and/or training purposes. It provides a set of preset models and enables distributed inference against them.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                                        | Type                             |
|----------|----------------------------------------------------|----------------------------------|
| enabled  | Indicates if AI toolchain operator enabled or not. | bool<br/><small>Optional</small> |

ManagedClusterAPIServerAccessProfile{#ManagedClusterAPIServerAccessProfile}
---------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type                                 |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[]<br/><small>Optional</small> |
| disableRunCommand              | Whether to disable run command for the cluster or not.                                                                                                                                                                                                                                                                        | bool<br/><small>Optional</small>     |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool<br/><small>Optional</small>     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>     |
| enableVnetIntegration          | Whether to enable apiserver vnet integration for the cluster or not.                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are `system` and `none`.                                                                                                                           | string<br/><small>Optional</small>   |
| subnetId                       | It is required when: 1. creating a new cluster with BYO Vnet; 2. updating an existing cluster to enable apiserver vnet integration.                                                                                                                                                                                           | string<br/><small>Optional</small>   |

ManagedClusterAPIServerAccessProfile_STATUS{#ManagedClusterAPIServerAccessProfile_STATUS}
-----------------------------------------------------------------------------------------

Access profile for managed cluster API server.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                       | Description                                                                                                                                                                                                                                                                                                                   | Type                                 |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| authorizedIPRanges             | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see [API server authorized IP ranges](https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges). | string[]<br/><small>Optional</small> |
| disableRunCommand              | Whether to disable run command for the cluster or not.                                                                                                                                                                                                                                                                        | bool<br/><small>Optional</small>     |
| enablePrivateCluster           | For more details, see [Creating a private AKS cluster](https://docs.microsoft.com/azure/aks/private-clusters).                                                                                                                                                                                                                | bool<br/><small>Optional</small>     |
| enablePrivateClusterPublicFQDN | Whether to create additional public FQDN for private cluster or not.                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>     |
| enableVnetIntegration          | Whether to enable apiserver vnet integration for the cluster or not.                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>     |
| privateDNSZone                 | The default is System. For more details see [configure private DNS zone](https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone). Allowed values are `system` and `none`.                                                                                                                           | string<br/><small>Optional</small>   |
| subnetId                       | It is required when: 1. creating a new cluster with BYO Vnet; 2. updating an existing cluster to enable apiserver vnet integration.                                                                                                                                                                                           | string<br/><small>Optional</small>   |

ManagedClusterAutoUpgradeProfile{#ManagedClusterAutoUpgradeProfile}
-------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property             | Description                                                                                                                                             | Type                                                                                                                                        |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| nodeOSUpgradeChannel | The default is Unmanaged, but may change to either NodeImage or SecurityPatch at GA.                                                                    | [ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel](#ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel)<br/><small>Optional</small> |
| upgradeChannel       | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel](#ManagedClusterAutoUpgradeProfile_UpgradeChannel)<br/><small>Optional</small>             |

ManagedClusterAutoUpgradeProfile_STATUS{#ManagedClusterAutoUpgradeProfile_STATUS}
---------------------------------------------------------------------------------

Auto upgrade profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property             | Description                                                                                                                                             | Type                                                                                                                                                      |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| nodeOSUpgradeChannel | The default is Unmanaged, but may change to either NodeImage or SecurityPatch at GA.                                                                    | [ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS](#ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS)<br/><small>Optional</small> |
| upgradeChannel       | For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel). | [ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS](#ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS)<br/><small>Optional</small>             |

ManagedClusterAzureMonitorProfile{#ManagedClusterAzureMonitorProfile}
---------------------------------------------------------------------

Prometheus addon profile for the container service cluster

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property          | Description                                                                                                                                                                                                                                                                    | Type                                                                                                                                  |
|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| appMonitoring     | Application Monitoring Profile for Kubernetes Application Container. Collects application logs, metrics and traces through auto-instrumentation of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview. | [ManagedClusterAzureMonitorProfileAppMonitoring](#ManagedClusterAzureMonitorProfileAppMonitoring)<br/><small>Optional</small>         |
| containerInsights | Azure Monitor Container Insights Profile for Kubernetes Events, Inventory and Container stdout & stderr logs etc. See aka.ms/AzureMonitorContainerInsights for an overview.                                                                                                    | [ManagedClusterAzureMonitorProfileContainerInsights](#ManagedClusterAzureMonitorProfileContainerInsights)<br/><small>Optional</small> |
| metrics           | Metrics profile for the prometheus service addon                                                                                                                                                                                                                               | [ManagedClusterAzureMonitorProfileMetrics](#ManagedClusterAzureMonitorProfileMetrics)<br/><small>Optional</small>                     |

ManagedClusterAzureMonitorProfile_STATUS{#ManagedClusterAzureMonitorProfile_STATUS}
-----------------------------------------------------------------------------------

Prometheus addon profile for the container service cluster

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property          | Description                                                                                                                                                                                                                                                                    | Type                                                                                                                                                |
|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| appMonitoring     | Application Monitoring Profile for Kubernetes Application Container. Collects application logs, metrics and traces through auto-instrumentation of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview. | [ManagedClusterAzureMonitorProfileAppMonitoring_STATUS](#ManagedClusterAzureMonitorProfileAppMonitoring_STATUS)<br/><small>Optional</small>         |
| containerInsights | Azure Monitor Container Insights Profile for Kubernetes Events, Inventory and Container stdout & stderr logs etc. See aka.ms/AzureMonitorContainerInsights for an overview.                                                                                                    | [ManagedClusterAzureMonitorProfileContainerInsights_STATUS](#ManagedClusterAzureMonitorProfileContainerInsights_STATUS)<br/><small>Optional</small> |
| metrics           | Metrics profile for the prometheus service addon                                                                                                                                                                                                                               | [ManagedClusterAzureMonitorProfileMetrics_STATUS](#ManagedClusterAzureMonitorProfileMetrics_STATUS)<br/><small>Optional</small>                     |

ManagedClusterBootstrapProfile{#ManagedClusterBootstrapProfile}
---------------------------------------------------------------

The bootstrap profile.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                   | Description                                                                                                                  | Type                                                                                                                                                       |
|----------------------------|------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| artifactSource             | The source where the artifacts are downloaded from.                                                                          | [ManagedClusterBootstrapProfile_ArtifactSource](#ManagedClusterBootstrapProfile_ArtifactSource)<br/><small>Optional</small>                                |
| containerRegistryReference | The resource Id of Azure Container Registry. The registry must have private network access, premium SKU and zone redundancy. | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

ManagedClusterBootstrapProfile_STATUS{#ManagedClusterBootstrapProfile_STATUS}
-----------------------------------------------------------------------------

The bootstrap profile.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property            | Description                                                                                                                  | Type                                                                                                                                      |
|---------------------|------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| artifactSource      | The source where the artifacts are downloaded from.                                                                          | [ManagedClusterBootstrapProfile_ArtifactSource_STATUS](#ManagedClusterBootstrapProfile_ArtifactSource_STATUS)<br/><small>Optional</small> |
| containerRegistryId | The resource Id of Azure Container Registry. The registry must have private network access, premium SKU and zone redundancy. | string<br/><small>Optional</small>                                                                                                        |

ManagedClusterHTTPProxyConfig{#ManagedClusterHTTPProxyConfig}
-------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property   | Description                                                 | Type                                 |
|------------|-------------------------------------------------------------|--------------------------------------|
| httpProxy  | The HTTP proxy server endpoint to use.                      | string<br/><small>Optional</small>   |
| httpsProxy | The HTTPS proxy server endpoint to use.                     | string<br/><small>Optional</small>   |
| noProxy    | The endpoints that should not go through proxy.             | string[]<br/><small>Optional</small> |
| trustedCa  | Alternative CA cert to use for connecting to proxy servers. | string<br/><small>Optional</small>   |

ManagedClusterHTTPProxyConfig_STATUS{#ManagedClusterHTTPProxyConfig_STATUS}
---------------------------------------------------------------------------

Cluster HTTP proxy configuration.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property         | Description                                                                                                                                         | Type                                 |
|------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| effectiveNoProxy | A read-only list of all endpoints for which traffic should not be sent to the proxy. This list is a superset of noProxy and values injected by AKS. | string[]<br/><small>Optional</small> |
| httpProxy        | The HTTP proxy server endpoint to use.                                                                                                              | string<br/><small>Optional</small>   |
| httpsProxy       | The HTTPS proxy server endpoint to use.                                                                                                             | string<br/><small>Optional</small>   |
| noProxy          | The endpoints that should not go through proxy.                                                                                                     | string[]<br/><small>Optional</small> |
| trustedCa        | Alternative CA cert to use for connecting to proxy servers.                                                                                         | string<br/><small>Optional</small>   |

ManagedClusterIdentity{#ManagedClusterIdentity}
-----------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property               | Description                                                                                                                                                                                                                                                                                                                    | Type                                                                                      |
|------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| delegatedResources     | The delegated identity resources assigned to this managed cluster. This can only be set by another Azure Resource Provider, and managed cluster only accept one delegated identity resource. Internal use only.                                                                                                                | [map[string]DelegatedResource](#DelegatedResource)<br/><small>Optional</small>            |
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).                                                                                                                                                                                                           | [ManagedClusterIdentity_Type](#ManagedClusterIdentity_Type)<br/><small>Optional</small>   |
| userAssignedIdentities | The keys must be ARM resource IDs in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/&ZeroWidthSpace;userAssignedIdentities/&ZeroWidthSpace;{identityName}'. | [UserAssignedIdentityDetails[]](#UserAssignedIdentityDetails)<br/><small>Optional</small> |

ManagedClusterIdentity_STATUS{#ManagedClusterIdentity_STATUS}
-------------------------------------------------------------

Identity for the managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property               | Description                                                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                 |
|------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| delegatedResources     | The delegated identity resources assigned to this managed cluster. This can only be set by another Azure Resource Provider, and managed cluster only accept one delegated identity resource. Internal use only.                                                                                                                | [map[string]DelegatedResource_STATUS](#DelegatedResource_STATUS)<br/><small>Optional</small>                                                         |
| principalId            | The principal id of the system assigned identity which is used by master components.                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                   |
| tenantId               | The tenant id of the system assigned identity which is used by master components.                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                   |
| type                   | For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).                                                                                                                                                                                                           | [ManagedClusterIdentity_Type_STATUS](#ManagedClusterIdentity_Type_STATUS)<br/><small>Optional</small>                                                |
| userAssignedIdentities | The keys must be ARM resource IDs in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/&ZeroWidthSpace;userAssignedIdentities/&ZeroWidthSpace;{identityName}'. | [map[string]ManagedClusterIdentity_UserAssignedIdentities_STATUS](#ManagedClusterIdentity_UserAssignedIdentities_STATUS)<br/><small>Optional</small> |

ManagedClusterIngressProfile{#ManagedClusterIngressProfile}
-----------------------------------------------------------

Ingress profile for the container service cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property      | Description                                       | Type                                                                                                                |
|---------------|---------------------------------------------------|---------------------------------------------------------------------------------------------------------------------|
| webAppRouting | Web App Routing settings for the ingress profile. | [ManagedClusterIngressProfileWebAppRouting](#ManagedClusterIngressProfileWebAppRouting)<br/><small>Optional</small> |

ManagedClusterIngressProfile_STATUS{#ManagedClusterIngressProfile_STATUS}
-------------------------------------------------------------------------

Ingress profile for the container service cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property      | Description                                       | Type                                                                                                                              |
|---------------|---------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------|
| webAppRouting | Web App Routing settings for the ingress profile. | [ManagedClusterIngressProfileWebAppRouting_STATUS](#ManagedClusterIngressProfileWebAppRouting_STATUS)<br/><small>Optional</small> |

ManagedClusterMetricsProfile{#ManagedClusterMetricsProfile}
-----------------------------------------------------------

The metrics profile for the ManagedCluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property     | Description                                     | Type                                                                                  |
|--------------|-------------------------------------------------|---------------------------------------------------------------------------------------|
| costAnalysis | The cost analysis configuration for the cluster | [ManagedClusterCostAnalysis](#ManagedClusterCostAnalysis)<br/><small>Optional</small> |

ManagedClusterMetricsProfile_STATUS{#ManagedClusterMetricsProfile_STATUS}
-------------------------------------------------------------------------

The metrics profile for the ManagedCluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property     | Description                                     | Type                                                                                                |
|--------------|-------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| costAnalysis | The cost analysis configuration for the cluster | [ManagedClusterCostAnalysis_STATUS](#ManagedClusterCostAnalysis_STATUS)<br/><small>Optional</small> |

ManagedClusterNodeProvisioningProfile{#ManagedClusterNodeProvisioningProfile}
-----------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                                        | Type                                                                                                                  |
|----------|--------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| mode     | Once the mode it set to Auto, it cannot be changed back to Manual. | [ManagedClusterNodeProvisioningProfile_Mode](#ManagedClusterNodeProvisioningProfile_Mode)<br/><small>Optional</small> |

ManagedClusterNodeProvisioningProfile_STATUS{#ManagedClusterNodeProvisioningProfile_STATUS}
-------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                                                        | Type                                                                                                                                |
|----------|--------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| mode     | Once the mode it set to Auto, it cannot be changed back to Manual. | [ManagedClusterNodeProvisioningProfile_Mode_STATUS](#ManagedClusterNodeProvisioningProfile_Mode_STATUS)<br/><small>Optional</small> |

ManagedClusterNodeResourceGroupProfile{#ManagedClusterNodeResourceGroupProfile}
-------------------------------------------------------------------------------

Node resource group lockdown profile for a managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property         | Description                                                        | Type                                                                                                                                            |
|------------------|--------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| restrictionLevel | The restriction level applied to the cluster's node resource group | [ManagedClusterNodeResourceGroupProfile_RestrictionLevel](#ManagedClusterNodeResourceGroupProfile_RestrictionLevel)<br/><small>Optional</small> |

ManagedClusterNodeResourceGroupProfile_STATUS{#ManagedClusterNodeResourceGroupProfile_STATUS}
---------------------------------------------------------------------------------------------

Node resource group lockdown profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property         | Description                                                        | Type                                                                                                                                                          |
|------------------|--------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| restrictionLevel | The restriction level applied to the cluster's node resource group | [ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS](#ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS)<br/><small>Optional</small> |

ManagedClusterOIDCIssuerProfile{#ManagedClusterOIDCIssuerProfile}
-----------------------------------------------------------------

The OIDC issuer profile of the Managed Cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                         | Type                             |
|----------|-------------------------------------|----------------------------------|
| enabled  | Whether the OIDC issuer is enabled. | bool<br/><small>Optional</small> |

ManagedClusterOIDCIssuerProfile_STATUS{#ManagedClusterOIDCIssuerProfile_STATUS}
-------------------------------------------------------------------------------

The OIDC issuer profile of the Managed Cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property  | Description                                 | Type                               |
|-----------|---------------------------------------------|------------------------------------|
| enabled   | Whether the OIDC issuer is enabled.         | bool<br/><small>Optional</small>   |
| issuerURL | The OIDC issuer url of the Managed Cluster. | string<br/><small>Optional</small> |

ManagedClusterOperatorSpec{#ManagedClusterOperatorSpec}
-------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| configMaps           | configures where to place operator written ConfigMaps.                                        | [ManagedClusterOperatorConfigMaps](#ManagedClusterOperatorConfigMaps)<br/><small>Optional</small>                                                                   |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secrets              | configures where to place Azure generated secrets.                                            | [ManagedClusterOperatorSecrets](#ManagedClusterOperatorSecrets)<br/><small>Optional</small>                                                                         |

ManagedClusterPodIdentityProfile{#ManagedClusterPodIdentityProfile}
-------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                                                    |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool<br/><small>Optional</small>                                                                        |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool<br/><small>Optional</small>                                                                        |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity[]](#ManagedClusterPodIdentity)<br/><small>Optional</small>                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException[]](#ManagedClusterPodIdentityException)<br/><small>Optional</small> |

ManagedClusterPodIdentityProfile_STATUS{#ManagedClusterPodIdentityProfile_STATUS}
---------------------------------------------------------------------------------

See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on pod identity integration.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                       | Description                                                                                                                                                                                                                                                                                                                                                   | Type                                                                                                                  |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| allowNetworkPluginKubenet      | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See [using Kubenet network plugin with AAD Pod Identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities) for more information. | bool<br/><small>Optional</small>                                                                                      |
| enabled                        | Whether the pod identity addon is enabled.                                                                                                                                                                                                                                                                                                                    | bool<br/><small>Optional</small>                                                                                      |
| userAssignedIdentities         | The pod identities to use in the cluster.                                                                                                                                                                                                                                                                                                                     | [ManagedClusterPodIdentity_STATUS[]](#ManagedClusterPodIdentity_STATUS)<br/><small>Optional</small>                   |
| userAssignedIdentityExceptions | The pod identity exceptions to allow.                                                                                                                                                                                                                                                                                                                         | [ManagedClusterPodIdentityException_STATUS[]](#ManagedClusterPodIdentityException_STATUS)<br/><small>Optional</small> |

ManagedClusterProperties_AutoScalerProfile{#ManagedClusterProperties_AutoScalerProfile}
---------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                              | Description                                                                                                                                                                                                                                                                                                                         | Type                                              |
|---------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------|
| balance-similar-node-groups           | Valid values are `true` and `false`                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                |
| daemonset-eviction-for-empty-nodes    | If set to true, all daemonset pods on empty nodes will be evicted before deletion of the node. If the daemonset pod cannot be evicted another node will be chosen for scaling. If set to false, the node will be deleted without ensuring that daemonset pods are deleted or evicted.                                               | bool<br/><small>Optional</small>                  |
| daemonset-eviction-for-occupied-nodes | If set to true, all daemonset pods on occupied nodes will be evicted before deletion of the node. If the daemonset pod cannot be evicted another node will be chosen for scaling. If set to false, the node will be deleted without ensuring that daemonset pods are deleted or evicted.                                            | bool<br/><small>Optional</small>                  |
| expander                              | Available values are: 'least-waste', 'most-pods', `priority`, `random`.                                                                                                                                                                                                                                                             | [Expander](#Expander)<br/><small>Optional</small> |
| ignore-daemonsets-utilization         | If set to true, the resources used by daemonset will be taken into account when making scaling down decisions.                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                  |
| max-empty-bulk-delete                 | The default is 10.                                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                |
| max-graceful-termination-sec          | The default is 600.                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                |
| max-node-provision-time               | The default is `15m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                |
| max-total-unready-percentage          | The default is 45. The maximum is 100 and the minimum is 0.                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                |
| new-pod-scale-up-delay                | For scenarios like burst/batch scale where you don't want CA to act before the kubernetes scheduler could schedule all the pods, you can tell CA to ignore unscheduled pods before they're a certain age. The default is `0s`. Values must be an integer followed by a unit (`s` for seconds, `m` for minutes, `h` for hours, etc). | string<br/><small>Optional</small>                |
| ok-total-unready-count                | This must be an integer. The default is 3.                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                |
| scale-down-delay-after-add            | The default is `10m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                |
| scale-down-delay-after-delete         | The default is the scan-interval. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                | string<br/><small>Optional</small>                |
| scale-down-delay-after-failure        | The default is `3m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                             | string<br/><small>Optional</small>                |
| scale-down-unneeded-time              | The default is `10m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                |
| scale-down-unready-time               | The default is `20m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                |
| scale-down-utilization-threshold      | The default is `0.5`.                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                |
| scan-interval                         | The default is `10`. Values must be an integer number of seconds.                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                |
| skip-nodes-with-local-storage         | The default is true.                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                |
| skip-nodes-with-system-pods           | The default is true.                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                |

ManagedClusterProperties_AutoScalerProfile_STATUS{#ManagedClusterProperties_AutoScalerProfile_STATUS}
-----------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                              | Description                                                                                                                                                                                                                                                                                                                         | Type                                                            |
|---------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------|
| balance-similar-node-groups           | Valid values are `true` and `false`                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                              |
| daemonset-eviction-for-empty-nodes    | If set to true, all daemonset pods on empty nodes will be evicted before deletion of the node. If the daemonset pod cannot be evicted another node will be chosen for scaling. If set to false, the node will be deleted without ensuring that daemonset pods are deleted or evicted.                                               | bool<br/><small>Optional</small>                                |
| daemonset-eviction-for-occupied-nodes | If set to true, all daemonset pods on occupied nodes will be evicted before deletion of the node. If the daemonset pod cannot be evicted another node will be chosen for scaling. If set to false, the node will be deleted without ensuring that daemonset pods are deleted or evicted.                                            | bool<br/><small>Optional</small>                                |
| expander                              | Available values are: 'least-waste', 'most-pods', `priority`, `random`.                                                                                                                                                                                                                                                             | [Expander_STATUS](#Expander_STATUS)<br/><small>Optional</small> |
| ignore-daemonsets-utilization         | If set to true, the resources used by daemonset will be taken into account when making scaling down decisions.                                                                                                                                                                                                                      | bool<br/><small>Optional</small>                                |
| max-empty-bulk-delete                 | The default is 10.                                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                              |
| max-graceful-termination-sec          | The default is 600.                                                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                              |
| max-node-provision-time               | The default is `15m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                              |
| max-total-unready-percentage          | The default is 45. The maximum is 100 and the minimum is 0.                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                              |
| new-pod-scale-up-delay                | For scenarios like burst/batch scale where you don't want CA to act before the kubernetes scheduler could schedule all the pods, you can tell CA to ignore unscheduled pods before they're a certain age. The default is `0s`. Values must be an integer followed by a unit (`s` for seconds, `m` for minutes, `h` for hours, etc). | string<br/><small>Optional</small>                              |
| ok-total-unready-count                | This must be an integer. The default is 3.                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                              |
| scale-down-delay-after-add            | The default is `10m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                              |
| scale-down-delay-after-delete         | The default is the scan-interval. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                | string<br/><small>Optional</small>                              |
| scale-down-delay-after-failure        | The default is `3m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                             | string<br/><small>Optional</small>                              |
| scale-down-unneeded-time              | The default is `10m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                              |
| scale-down-unready-time               | The default is `20m`. Values must be an integer followed by an `m`. No unit of time other than minutes (m) is supported.                                                                                                                                                                                                            | string<br/><small>Optional</small>                              |
| scale-down-utilization-threshold      | The default is `0.5`.                                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                              |
| scan-interval                         | The default is `10`. Values must be an integer number of seconds.                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                              |
| skip-nodes-with-local-storage         | The default is true.                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                              |
| skip-nodes-with-system-pods           | The default is true.                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                              |

ManagedClusterProperties_PublicNetworkAccess{#ManagedClusterProperties_PublicNetworkAccess}
-------------------------------------------------------------------------------------------

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Value                | Description |
|----------------------|-------------|
| "Disabled"           |             |
| "Enabled"            |             |
| "SecuredByPerimeter" |             |

ManagedClusterProperties_PublicNetworkAccess_STATUS{#ManagedClusterProperties_PublicNetworkAccess_STATUS}
---------------------------------------------------------------------------------------------------------

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Value                | Description |
|----------------------|-------------|
| "Disabled"           |             |
| "Enabled"            |             |
| "SecuredByPerimeter" |             |

ManagedClustersAgentPoolOperatorSpec{#ManagedClustersAgentPoolOperatorSpec}
---------------------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

ManagedClusterSecurityProfile{#ManagedClusterSecurityProfile}
-------------------------------------------------------------

Security profile for the container service cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property                  | Description                                                                                                                                                                                                                                                                | Type                                                                                                                        |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| azureKeyVaultKms          | Azure Key Vault [key management service](https://kubernetes.io/docs/tasks/administer-cluster/kms-provider/) settings for the security profile.                                                                                                                             | [AzureKeyVaultKms](#AzureKeyVaultKms)<br/><small>Optional</small>                                                           |
| customCATrustCertificates | A list of up to 10 base64 encoded CAs that will be added to the trust store on nodes with the Custom CA Trust feature enabled. For more information see [Custom CA Trust Certificates](https://learn.microsoft.com/en-us/azure/aks/custom-certificate-authority)           | ManagedClusterSecurityProfileCustomCATrustCertificates<br/><small>Optional</small>                                          |
| defender                  | Microsoft Defender settings for the security profile.                                                                                                                                                                                                                      | [ManagedClusterSecurityProfileDefender](#ManagedClusterSecurityProfileDefender)<br/><small>Optional</small>                 |
| imageCleaner              | Image Cleaner settings for the security profile.                                                                                                                                                                                                                           | [ManagedClusterSecurityProfileImageCleaner](#ManagedClusterSecurityProfileImageCleaner)<br/><small>Optional</small>         |
| imageIntegrity            | Image integrity is a feature that works with Azure Policy to verify image integrity by signature. This will not have any effect unless Azure Policy is applied to enforce image signatures. See https://aka.ms/aks/image-integrity for how to use this feature via policy. | [ManagedClusterSecurityProfileImageIntegrity](#ManagedClusterSecurityProfileImageIntegrity)<br/><small>Optional</small>     |
| nodeRestriction           | [Node Restriction](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#noderestriction) settings for the security profile.                                                                                                                      | [ManagedClusterSecurityProfileNodeRestriction](#ManagedClusterSecurityProfileNodeRestriction)<br/><small>Optional</small>   |
| workloadIdentity          | Workload identity settings for the security profile. Workload identity enables Kubernetes applications to access Azure cloud resources securely with Azure AD. See https://aka.ms/aks/wi for more details.                                                                 | [ManagedClusterSecurityProfileWorkloadIdentity](#ManagedClusterSecurityProfileWorkloadIdentity)<br/><small>Optional</small> |

ManagedClusterSecurityProfile_STATUS{#ManagedClusterSecurityProfile_STATUS}
---------------------------------------------------------------------------

Security profile for the container service cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                  | Description                                                                                                                                                                                                                                                                | Type                                                                                                                                      |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| azureKeyVaultKms          | Azure Key Vault [key management service](https://kubernetes.io/docs/tasks/administer-cluster/kms-provider/) settings for the security profile.                                                                                                                             | [AzureKeyVaultKms_STATUS](#AzureKeyVaultKms_STATUS)<br/><small>Optional</small>                                                           |
| customCATrustCertificates | A list of up to 10 base64 encoded CAs that will be added to the trust store on nodes with the Custom CA Trust feature enabled. For more information see [Custom CA Trust Certificates](https://learn.microsoft.com/en-us/azure/aks/custom-certificate-authority)           | string[]<br/><small>Optional</small>                                                                                                      |
| defender                  | Microsoft Defender settings for the security profile.                                                                                                                                                                                                                      | [ManagedClusterSecurityProfileDefender_STATUS](#ManagedClusterSecurityProfileDefender_STATUS)<br/><small>Optional</small>                 |
| imageCleaner              | Image Cleaner settings for the security profile.                                                                                                                                                                                                                           | [ManagedClusterSecurityProfileImageCleaner_STATUS](#ManagedClusterSecurityProfileImageCleaner_STATUS)<br/><small>Optional</small>         |
| imageIntegrity            | Image integrity is a feature that works with Azure Policy to verify image integrity by signature. This will not have any effect unless Azure Policy is applied to enforce image signatures. See https://aka.ms/aks/image-integrity for how to use this feature via policy. | [ManagedClusterSecurityProfileImageIntegrity_STATUS](#ManagedClusterSecurityProfileImageIntegrity_STATUS)<br/><small>Optional</small>     |
| nodeRestriction           | [Node Restriction](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#noderestriction) settings for the security profile.                                                                                                                      | [ManagedClusterSecurityProfileNodeRestriction_STATUS](#ManagedClusterSecurityProfileNodeRestriction_STATUS)<br/><small>Optional</small>   |
| workloadIdentity          | Workload identity settings for the security profile. Workload identity enables Kubernetes applications to access Azure cloud resources securely with Azure AD. See https://aka.ms/aks/wi for more details.                                                                 | [ManagedClusterSecurityProfileWorkloadIdentity_STATUS](#ManagedClusterSecurityProfileWorkloadIdentity_STATUS)<br/><small>Optional</small> |

ManagedClusterServicePrincipalProfile{#ManagedClusterServicePrincipalProfile}
-----------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                                              | Type                                                                                                                                                   |
|----------|--------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| clientId | The ID for the service principal.                                        | string<br/><small>Required</small>                                                                                                                     |
| secret   | The secret password associated with the service principal in plain text. | [genruntime.SecretReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference)<br/><small>Optional</small> |

ManagedClusterServicePrincipalProfile_STATUS{#ManagedClusterServicePrincipalProfile_STATUS}
-------------------------------------------------------------------------------------------

Information about a service principal identity for the cluster to use for manipulating Azure APIs.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                       | Type                               |
|----------|-----------------------------------|------------------------------------|
| clientId | The ID for the service principal. | string<br/><small>Optional</small> |

ManagedClusterSKU{#ManagedClusterSKU}
-------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                                                                                                                                          | Type                                                                          |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                                                   | [ManagedClusterSKU_Name](#ManagedClusterSKU_Name)<br/><small>Optional</small> |
| tier     | If not specified, the default is `Free`. See [AKS Pricing Tier](https://learn.microsoft.com/azure/aks/free-standard-pricing-tiers) for more details. | [ManagedClusterSKU_Tier](#ManagedClusterSKU_Tier)<br/><small>Optional</small> |

ManagedClusterSKU_STATUS{#ManagedClusterSKU_STATUS}
---------------------------------------------------

The SKU of a Managed Cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                                                                                                                                          | Type                                                                                        |
|----------|------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| name     | The name of a managed cluster SKU.                                                                                                                   | [ManagedClusterSKU_Name_STATUS](#ManagedClusterSKU_Name_STATUS)<br/><small>Optional</small> |
| tier     | If not specified, the default is `Free`. See [AKS Pricing Tier](https://learn.microsoft.com/azure/aks/free-standard-pricing-tiers) for more details. | [ManagedClusterSKU_Tier_STATUS](#ManagedClusterSKU_Tier_STATUS)<br/><small>Optional</small> |

ManagedClusterStorageProfile{#ManagedClusterStorageProfile}
-----------------------------------------------------------

Storage profile for the container service cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property           | Description                                            | Type                                                                                                                          |
|--------------------|--------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| blobCSIDriver      | AzureBlob CSI Driver settings for the storage profile. | [ManagedClusterStorageProfileBlobCSIDriver](#ManagedClusterStorageProfileBlobCSIDriver)<br/><small>Optional</small>           |
| diskCSIDriver      | AzureDisk CSI Driver settings for the storage profile. | [ManagedClusterStorageProfileDiskCSIDriver](#ManagedClusterStorageProfileDiskCSIDriver)<br/><small>Optional</small>           |
| fileCSIDriver      | AzureFile CSI Driver settings for the storage profile. | [ManagedClusterStorageProfileFileCSIDriver](#ManagedClusterStorageProfileFileCSIDriver)<br/><small>Optional</small>           |
| snapshotController | Snapshot Controller settings for the storage profile.  | [ManagedClusterStorageProfileSnapshotController](#ManagedClusterStorageProfileSnapshotController)<br/><small>Optional</small> |

ManagedClusterStorageProfile_STATUS{#ManagedClusterStorageProfile_STATUS}
-------------------------------------------------------------------------

Storage profile for the container service cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property           | Description                                            | Type                                                                                                                                        |
|--------------------|--------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| blobCSIDriver      | AzureBlob CSI Driver settings for the storage profile. | [ManagedClusterStorageProfileBlobCSIDriver_STATUS](#ManagedClusterStorageProfileBlobCSIDriver_STATUS)<br/><small>Optional</small>           |
| diskCSIDriver      | AzureDisk CSI Driver settings for the storage profile. | [ManagedClusterStorageProfileDiskCSIDriver_STATUS](#ManagedClusterStorageProfileDiskCSIDriver_STATUS)<br/><small>Optional</small>           |
| fileCSIDriver      | AzureFile CSI Driver settings for the storage profile. | [ManagedClusterStorageProfileFileCSIDriver_STATUS](#ManagedClusterStorageProfileFileCSIDriver_STATUS)<br/><small>Optional</small>           |
| snapshotController | Snapshot Controller settings for the storage profile.  | [ManagedClusterStorageProfileSnapshotController_STATUS](#ManagedClusterStorageProfileSnapshotController_STATUS)<br/><small>Optional</small> |

ManagedClusterWindowsProfile{#ManagedClusterWindowsProfile}
-----------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                                                                   |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| adminPassword  | Specifies the password of the administrator account. Minimum-length: 8 characters Max-length: 123 characters Complexity requirements: 3 out of 4 conditions below need to be fulfilled Has lower characters Has upper characters Has a digit Has a special character (Regex match \[\W_]) Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"                       | [genruntime.SecretReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference)<br/><small>Optional</small> |
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string<br/><small>Required</small>                                                                                                                     |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                                                       |
| gmsaProfile    | The Windows gMSA Profile in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                                                                                                   | [WindowsGmsaProfile](#WindowsGmsaProfile)<br/><small>Optional</small>                                                                                  |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType](#ManagedClusterWindowsProfile_LicenseType)<br/><small>Optional</small>                                      |

ManagedClusterWindowsProfile_STATUS{#ManagedClusterWindowsProfile_STATUS}
-------------------------------------------------------------------------

Profile for Windows VMs in the managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                                            |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------|
| adminUsername  | Specifies the name of the administrator account. Restriction: Cannot end in "." Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". Minimum-length: 1 character Max-length: 20 characters | string<br/><small>Optional</small>                                                                                              |
| enableCSIProxy | For more details on CSI proxy, see the [CSI proxy GitHub repo](https://github.com/kubernetes-csi/csi-proxy).                                                                                                                                                                                                                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                                |
| gmsaProfile    | The Windows gMSA Profile in the Managed Cluster.                                                                                                                                                                                                                                                                                                                                                                                                                   | [WindowsGmsaProfile_STATUS](#WindowsGmsaProfile_STATUS)<br/><small>Optional</small>                                             |
| licenseType    | The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.                                                                                                                                                                                                                                                                                                               | [ManagedClusterWindowsProfile_LicenseType_STATUS](#ManagedClusterWindowsProfile_LicenseType_STATUS)<br/><small>Optional</small> |

ManagedClusterWorkloadAutoScalerProfile{#ManagedClusterWorkloadAutoScalerProfile}
---------------------------------------------------------------------------------

Workload Auto-scaler profile for the managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property              | Description                                                                               | Type                                                                                                                                                      |
|-----------------------|-------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| keda                  | KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile. | [ManagedClusterWorkloadAutoScalerProfileKeda](#ManagedClusterWorkloadAutoScalerProfileKeda)<br/><small>Optional</small>                                   |
| verticalPodAutoscaler |                                                                                           | [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler)<br/><small>Optional</small> |

ManagedClusterWorkloadAutoScalerProfile_STATUS{#ManagedClusterWorkloadAutoScalerProfile_STATUS}
-----------------------------------------------------------------------------------------------

Workload Auto-scaler profile for the managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property              | Description                                                                               | Type                                                                                                                                                                    |
|-----------------------|-------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| keda                  | KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile. | [ManagedClusterWorkloadAutoScalerProfileKeda_STATUS](#ManagedClusterWorkloadAutoScalerProfileKeda_STATUS)<br/><small>Optional</small>                                   |
| verticalPodAutoscaler |                                                                                           | [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS)<br/><small>Optional</small> |

OSDiskType{#OSDiskType}
-----------------------

The default is `Ephemeral` if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to `Managed`. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value       | Description |
|-------------|-------------|
| "Ephemeral" |             |
| "Managed"   |             |

OSDiskType_STATUS{#OSDiskType_STATUS}
-------------------------------------

The default is `Ephemeral` if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to `Managed`. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value       | Description |
|-------------|-------------|
| "Ephemeral" |             |
| "Managed"   |             |

OSSKU{#OSSKU}
-------------

Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.

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

OSSKU_STATUS{#OSSKU_STATUS}
---------------------------

Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.

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

OSType{#OSType}
---------------

The operating system type. The default is Linux.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value     | Description |
|-----------|-------------|
| "Linux"   |             |
| "Windows" |             |

OSType_STATUS{#OSType_STATUS}
-----------------------------

The operating system type. The default is Linux.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value     | Description |
|-----------|-------------|
| "Linux"   |             |
| "Windows" |             |

PodIPAllocationMode{#PodIPAllocationMode}
-----------------------------------------

The IP allocation mode for pods in the agent pool. Must be used with podSubnetId. The default is `DynamicIndividual`.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value               | Description |
|---------------------|-------------|
| "DynamicIndividual" |             |
| "StaticBlock"       |             |

PodIPAllocationMode_STATUS{#PodIPAllocationMode_STATUS}
-------------------------------------------------------

The IP allocation mode for pods in the agent pool. Must be used with podSubnetId. The default is `DynamicIndividual`.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value               | Description |
|---------------------|-------------|
| "DynamicIndividual" |             |
| "StaticBlock"       |             |

PowerState{#PowerState}
-----------------------

Describes the Power State of the cluster

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property | Description                                     | Type                                                            |
|----------|-------------------------------------------------|-----------------------------------------------------------------|
| code     | Tells whether the cluster is Running or Stopped | [PowerState_Code](#PowerState_Code)<br/><small>Optional</small> |

PowerState_STATUS{#PowerState_STATUS}
-------------------------------------

Describes the Power State of the cluster

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS), [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property | Description                                     | Type                                                                          |
|----------|-------------------------------------------------|-------------------------------------------------------------------------------|
| code     | Tells whether the cluster is Running or Stopped | [PowerState_Code_STATUS](#PowerState_Code_STATUS)<br/><small>Optional</small> |

PrivateLinkResource{#PrivateLinkResource}
-----------------------------------------

A private link resource

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property        | Description                            | Type                                                                                                                                                       |
|-----------------|----------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| groupId         | The group ID of the resource.          | string<br/><small>Optional</small>                                                                                                                         |
| name            | The name of the private link resource. | string<br/><small>Optional</small>                                                                                                                         |
| reference       | The ID of the private link resource.   | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| requiredMembers | The RequiredMembers of the resource    | string[]<br/><small>Optional</small>                                                                                                                       |
| type            | The resource type.                     | string<br/><small>Optional</small>                                                                                                                         |

PrivateLinkResource_STATUS{#PrivateLinkResource_STATUS}
-------------------------------------------------------

A private link resource

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property             | Description                                                                                | Type                                 |
|----------------------|--------------------------------------------------------------------------------------------|--------------------------------------|
| groupId              | The group ID of the resource.                                                              | string<br/><small>Optional</small>   |
| id                   | The ID of the private link resource.                                                       | string<br/><small>Optional</small>   |
| name                 | The name of the private link resource.                                                     | string<br/><small>Optional</small>   |
| privateLinkServiceID | The private link service ID of the resource, this field is exposed only to NRP internally. | string<br/><small>Optional</small>   |
| requiredMembers      | The RequiredMembers of the resource                                                        | string[]<br/><small>Optional</small> |
| type                 | The resource type.                                                                         | string<br/><small>Optional</small>   |

SafeguardsProfile{#SafeguardsProfile}
-------------------------------------

The Safeguards profile.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property           | Description                                                                                                                                       | Type                                                                            |
|--------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------|
| excludedNamespaces | List of namespaces excluded from Safeguards checks                                                                                                | string[]<br/><small>Optional</small>                                            |
| level              | The Safeguards level to be used. By default, Safeguards is enabled for all namespaces except those that AKS excludes via systemExcludedNamespaces | [SafeguardsProfile_Level](#SafeguardsProfile_Level)<br/><small>Required</small> |
| version            | The version of constraints to use                                                                                                                 | string<br/><small>Optional</small>                                              |

SafeguardsProfile_STATUS{#SafeguardsProfile_STATUS}
---------------------------------------------------

The Safeguards profile.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property                 | Description                                                                                                                                       | Type                                                                                          |
|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|
| excludedNamespaces       | List of namespaces excluded from Safeguards checks                                                                                                | string[]<br/><small>Optional</small>                                                          |
| level                    | The Safeguards level to be used. By default, Safeguards is enabled for all namespaces except those that AKS excludes via systemExcludedNamespaces | [SafeguardsProfile_Level_STATUS](#SafeguardsProfile_Level_STATUS)<br/><small>Optional</small> |
| systemExcludedNamespaces | List of namespaces specified by AKS to be excluded from Safeguards                                                                                | string[]<br/><small>Optional</small>                                                          |
| version                  | The version of constraints to use                                                                                                                 | string<br/><small>Optional</small>                                                            |

ScaleDownMode{#ScaleDownMode}
-----------------------------

Describes how VMs are added to or removed from Agent Pools. See [billing states](https://docs.microsoft.com/azure/virtual-machines/states-billing).

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

ScaleDownMode_STATUS{#ScaleDownMode_STATUS}
-------------------------------------------

Describes how VMs are added to or removed from Agent Pools. See [billing states](https://docs.microsoft.com/azure/virtual-machines/states-billing).

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

ScaleSetEvictionPolicy{#ScaleSetEvictionPolicy}
-----------------------------------------------

The eviction policy specifies what to do with the VM when it is evicted. The default is Delete. For more information about eviction see [spot VMs](https://docs.microsoft.com/azure/virtual-machines/spot-vms)

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

ScaleSetEvictionPolicy_STATUS{#ScaleSetEvictionPolicy_STATUS}
-------------------------------------------------------------

The eviction policy specifies what to do with the VM when it is evicted. The default is Delete. For more information about eviction see [spot VMs](https://docs.microsoft.com/azure/virtual-machines/spot-vms)

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value        | Description |
|--------------|-------------|
| "Deallocate" |             |
| "Delete"     |             |

ScaleSetPriority{#ScaleSetPriority}
-----------------------------------

The Virtual Machine Scale Set priority.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value     | Description |
|-----------|-------------|
| "Regular" |             |
| "Spot"    |             |

ScaleSetPriority_STATUS{#ScaleSetPriority_STATUS}
-------------------------------------------------

The Virtual Machine Scale Set priority.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value     | Description |
|-----------|-------------|
| "Regular" |             |
| "Spot"    |             |

ServiceMeshProfile{#ServiceMeshProfile}
---------------------------------------

Service mesh profile for a managed cluster.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec).

| Property | Description                       | Type                                                                            |
|----------|-----------------------------------|---------------------------------------------------------------------------------|
| istio    | Istio service mesh configuration. | [IstioServiceMesh](#IstioServiceMesh)<br/><small>Optional</small>               |
| mode     | Mode of the service mesh.         | [ServiceMeshProfile_Mode](#ServiceMeshProfile_Mode)<br/><small>Required</small> |

ServiceMeshProfile_STATUS{#ServiceMeshProfile_STATUS}
-----------------------------------------------------

Service mesh profile for a managed cluster.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS).

| Property | Description                       | Type                                                                                          |
|----------|-----------------------------------|-----------------------------------------------------------------------------------------------|
| istio    | Istio service mesh configuration. | [IstioServiceMesh_STATUS](#IstioServiceMesh_STATUS)<br/><small>Optional</small>               |
| mode     | Mode of the service mesh.         | [ServiceMeshProfile_Mode_STATUS](#ServiceMeshProfile_Mode_STATUS)<br/><small>Optional</small> |

SystemData_STATUS{#SystemData_STATUS}
-------------------------------------

Metadata pertaining to creation and last modification of the resource.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS), and [TrustedAccessRoleBinding_STATUS](#TrustedAccessRoleBinding_STATUS).

| Property           | Description                                           | Type                                                                                                      |
|--------------------|-------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| createdAt          | The timestamp of resource creation (UTC).             | string<br/><small>Optional</small>                                                                        |
| createdBy          | The identity that created the resource.               | string<br/><small>Optional</small>                                                                        |
| createdByType      | The type of identity that created the resource.       | [SystemData_CreatedByType_STATUS](#SystemData_CreatedByType_STATUS)<br/><small>Optional</small>           |
| lastModifiedAt     | The timestamp of resource last modification (UTC)     | string<br/><small>Optional</small>                                                                        |
| lastModifiedBy     | The identity that last modified the resource.         | string<br/><small>Optional</small>                                                                        |
| lastModifiedByType | The type of identity that last modified the resource. | [SystemData_LastModifiedByType_STATUS](#SystemData_LastModifiedByType_STATUS)<br/><small>Optional</small> |

TrustedAccessRoleBindingOperatorSpec{#TrustedAccessRoleBindingOperatorSpec}
---------------------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [TrustedAccessRoleBinding_Spec](#TrustedAccessRoleBinding_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

TrustedAccessRoleBindingProperties_ProvisioningState_STATUS{#TrustedAccessRoleBindingProperties_ProvisioningState_STATUS}
-------------------------------------------------------------------------------------------------------------------------

Used by: [TrustedAccessRoleBinding_STATUS](#TrustedAccessRoleBinding_STATUS).

| Value       | Description |
|-------------|-------------|
| "Canceled"  |             |
| "Deleting"  |             |
| "Failed"    |             |
| "Succeeded" |             |
| "Updating"  |             |

UserAssignedIdentity{#UserAssignedIdentity}
-------------------------------------------

Details about a user assigned identity.

Used by: [ManagedCluster_Spec](#ManagedCluster_Spec), and [ManagedClusterPodIdentity](#ManagedClusterPodIdentity).

| Property           | Description                                    | Type                                                                                                                                                         |
|--------------------|------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| clientId           | The client ID of the user assigned identity.   | string<br/><small>Optional</small>                                                                                                                           |
| clientIdFromConfig | The client ID of the user assigned identity.   | [genruntime.ConfigMapReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference)<br/><small>Optional</small> |
| objectId           | The object ID of the user assigned identity.   | string<br/><small>Optional</small>                                                                                                                           |
| objectIdFromConfig | The object ID of the user assigned identity.   | [genruntime.ConfigMapReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference)<br/><small>Optional</small> |
| resourceReference  | The resource ID of the user assigned identity. | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>   |

UserAssignedIdentity_STATUS{#UserAssignedIdentity_STATUS}
---------------------------------------------------------

Details about a user assigned identity.

Used by: [ManagedCluster_STATUS](#ManagedCluster_STATUS), [ManagedClusterAddonProfile_STATUS](#ManagedClusterAddonProfile_STATUS), [ManagedClusterIngressProfileWebAppRouting_STATUS](#ManagedClusterIngressProfileWebAppRouting_STATUS), and [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Property   | Description                                    | Type                               |
|------------|------------------------------------------------|------------------------------------|
| clientId   | The client ID of the user assigned identity.   | string<br/><small>Optional</small> |
| objectId   | The object ID of the user assigned identity.   | string<br/><small>Optional</small> |
| resourceId | The resource ID of the user assigned identity. | string<br/><small>Optional</small> |

VirtualMachineNodes{#VirtualMachineNodes}
-----------------------------------------

Current status on a group of nodes of the same vm size.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property | Description                                                 | Type                               |
|----------|-------------------------------------------------------------|------------------------------------|
| count    | Number of nodes.                                            | int<br/><small>Optional</small>    |
| size     | The VM size of the agents used to host this group of nodes. | string<br/><small>Optional</small> |

VirtualMachineNodes_STATUS{#VirtualMachineNodes_STATUS}
-------------------------------------------------------

Current status on a group of nodes of the same vm size.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property | Description                                                 | Type                               |
|----------|-------------------------------------------------------------|------------------------------------|
| count    | Number of nodes.                                            | int<br/><small>Optional</small>    |
| size     | The VM size of the agents used to host this group of nodes. | string<br/><small>Optional</small> |

VirtualMachinesProfile{#VirtualMachinesProfile}
-----------------------------------------------

Specifications on VirtualMachines agent pool.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Property | Description                                                  | Type                                                      |
|----------|--------------------------------------------------------------|-----------------------------------------------------------|
| scale    | Specifications on how to scale a VirtualMachines agent pool. | [ScaleProfile](#ScaleProfile)<br/><small>Optional</small> |

VirtualMachinesProfile_STATUS{#VirtualMachinesProfile_STATUS}
-------------------------------------------------------------

Specifications on VirtualMachines agent pool.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Property | Description                                                  | Type                                                                    |
|----------|--------------------------------------------------------------|-------------------------------------------------------------------------|
| scale    | Specifications on how to scale a VirtualMachines agent pool. | [ScaleProfile_STATUS](#ScaleProfile_STATUS)<br/><small>Optional</small> |

WorkloadRuntime{#WorkloadRuntime}
---------------------------------

Determines the type of workload a node can run.

Used by: [ManagedClusterAgentPoolProfile](#ManagedClusterAgentPoolProfile), and [ManagedClustersAgentPool_Spec](#ManagedClustersAgentPool_Spec).

| Value                 | Description |
|-----------------------|-------------|
| "KataMshvVmIsolation" |             |
| "OCIContainer"        |             |
| "WasmWasi"            |             |

WorkloadRuntime_STATUS{#WorkloadRuntime_STATUS}
-----------------------------------------------

Determines the type of workload a node can run.

Used by: [ManagedClusterAgentPoolProfile_STATUS](#ManagedClusterAgentPoolProfile_STATUS), and [ManagedClustersAgentPool_STATUS](#ManagedClustersAgentPool_STATUS).

| Value                 | Description |
|-----------------------|-------------|
| "KataMshvVmIsolation" |             |
| "OCIContainer"        |             |
| "WasmWasi"            |             |

AdvancedNetworking{#AdvancedNetworking}
---------------------------------------

Advanced Networking profile for enabling observability on a cluster. Note that enabling advanced networking features may incur additional costs. For more information see aka.ms/aksadvancednetworking.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property      | Description                                                                                      | Type                                                                                            |
|---------------|--------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| observability | Observability profile to enable advanced network metrics and flow logs with historical contexts. | [AdvancedNetworkingObservability](#AdvancedNetworkingObservability)<br/><small>Optional</small> |

AdvancedNetworking_STATUS{#AdvancedNetworking_STATUS}
-----------------------------------------------------

Advanced Networking profile for enabling observability on a cluster. Note that enabling advanced networking features may incur additional costs. For more information see aka.ms/aksadvancednetworking.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property      | Description                                                                                      | Type                                                                                                          |
|---------------|--------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| observability | Observability profile to enable advanced network metrics and flow logs with historical contexts. | [AdvancedNetworkingObservability_STATUS](#AdvancedNetworkingObservability_STATUS)<br/><small>Optional</small> |

AgentPoolSSHAccess{#AgentPoolSSHAccess}
---------------------------------------

SSH access method of an agent pool.

Used by: [AgentPoolSecurityProfile](#AgentPoolSecurityProfile).

| Value       | Description |
|-------------|-------------|
| "Disabled"  |             |
| "LocalUser" |             |

AgentPoolSSHAccess_STATUS{#AgentPoolSSHAccess_STATUS}
-----------------------------------------------------

SSH access method of an agent pool.

Used by: [AgentPoolSecurityProfile_STATUS](#AgentPoolSecurityProfile_STATUS).

| Value       | Description |
|-------------|-------------|
| "Disabled"  |             |
| "LocalUser" |             |

AgentPoolUpgradeSettings_UndrainableNodeBehavior{#AgentPoolUpgradeSettings_UndrainableNodeBehavior}
---------------------------------------------------------------------------------------------------

Used by: [AgentPoolUpgradeSettings](#AgentPoolUpgradeSettings).

| Value      | Description |
|------------|-------------|
| "Cordon"   |             |
| "Schedule" |             |

AgentPoolUpgradeSettings_UndrainableNodeBehavior_STATUS{#AgentPoolUpgradeSettings_UndrainableNodeBehavior_STATUS}
-----------------------------------------------------------------------------------------------------------------

Used by: [AgentPoolUpgradeSettings_STATUS](#AgentPoolUpgradeSettings_STATUS).

| Value      | Description |
|------------|-------------|
| "Cordon"   |             |
| "Schedule" |             |

AzureKeyVaultKms{#AzureKeyVaultKms}
-----------------------------------

Azure Key Vault key management service settings for the security profile.

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property                  | Description                                                                                                                                                                                                                                                                                                                                                                                              | Type                                                                                                                                                       |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| enabled                   | Whether to enable Azure Key Vault key management service. The default is false.                                                                                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>                                                                                                                           |
| keyId                     | Identifier of Azure Key Vault key. See [key identifier format](https://docs.microsoft.com/en-us/azure/key-vault/general/about-keys-secrets-certificates#vault-name-and-object-name) for more details. When Azure Key Vault key management service is enabled, this field is required and must be a valid key identifier. When Azure Key Vault key management service is disabled, leave the field empty. | string<br/><small>Optional</small>                                                                                                                         |
| keyVaultNetworkAccess     | Network access of key vault. The possible values are `Public` and `Private`. `Public` means the key vault allows public access from all networks. `Private` means the key vault disables public access and enables private link. The default value is `Public`.                                                                                                                                          | [AzureKeyVaultKms_KeyVaultNetworkAccess](#AzureKeyVaultKms_KeyVaultNetworkAccess)<br/><small>Optional</small>                                              |
| keyVaultResourceReference | Resource ID of key vault. When keyVaultNetworkAccess is `Private`, this field is required and must be a valid resource ID. When keyVaultNetworkAccess is `Public`, leave the field empty.                                                                                                                                                                                                                | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

AzureKeyVaultKms_STATUS{#AzureKeyVaultKms_STATUS}
-------------------------------------------------

Azure Key Vault key management service settings for the security profile.

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property              | Description                                                                                                                                                                                                                                                                                                                                                                                              | Type                                                                                                                        |
|-----------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| enabled               | Whether to enable Azure Key Vault key management service. The default is false.                                                                                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>                                                                                            |
| keyId                 | Identifier of Azure Key Vault key. See [key identifier format](https://docs.microsoft.com/en-us/azure/key-vault/general/about-keys-secrets-certificates#vault-name-and-object-name) for more details. When Azure Key Vault key management service is enabled, this field is required and must be a valid key identifier. When Azure Key Vault key management service is disabled, leave the field empty. | string<br/><small>Optional</small>                                                                                          |
| keyVaultNetworkAccess | Network access of key vault. The possible values are `Public` and `Private`. `Public` means the key vault allows public access from all networks. `Private` means the key vault disables public access and enables private link. The default value is `Public`.                                                                                                                                          | [AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS](#AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS)<br/><small>Optional</small> |
| keyVaultResourceId    | Resource ID of key vault. When keyVaultNetworkAccess is `Private`, this field is required and must be a valid resource ID. When keyVaultNetworkAccess is `Public`, leave the field empty.                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                          |

ContainerServiceNetworkProfile_KubeProxyConfig{#ContainerServiceNetworkProfile_KubeProxyConfig}
-----------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property   | Description                                                                                                                                            | Type                                                                                                                                                |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| enabled    | Whether to enable on kube-proxy on the cluster (if no `kubeProxyConfig` exists, kube-proxy is enabled in AKS by default without these customizations). | bool<br/><small>Optional</small>                                                                                                                    |
| ipvsConfig | Holds configuration customizations for IPVS. May only be specified if `mode` is set to `IPVS`.                                                         | [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig)<br/><small>Optional</small> |
| mode       | Specify which proxy mode to use (`IPTABLES` or `IPVS`\)                                                                                                | [ContainerServiceNetworkProfile_KubeProxyConfig_Mode](#ContainerServiceNetworkProfile_KubeProxyConfig_Mode)<br/><small>Optional</small>             |

ContainerServiceNetworkProfile_KubeProxyConfig_STATUS{#ContainerServiceNetworkProfile_KubeProxyConfig_STATUS}
-------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property   | Description                                                                                                                                            | Type                                                                                                                                                              |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| enabled    | Whether to enable on kube-proxy on the cluster (if no `kubeProxyConfig` exists, kube-proxy is enabled in AKS by default without these customizations). | bool<br/><small>Optional</small>                                                                                                                                  |
| ipvsConfig | Holds configuration customizations for IPVS. May only be specified if `mode` is set to `IPVS`.                                                         | [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS)<br/><small>Optional</small> |
| mode       | Specify which proxy mode to use (`IPTABLES` or `IPVS`\)                                                                                                | [ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS)<br/><small>Optional</small>             |

ContainerServiceNetworkProfile_OutboundType{#ContainerServiceNetworkProfile_OutboundType}
-----------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value                    | Description |
|--------------------------|-------------|
| "loadBalancer"           |             |
| "managedNATGateway"      |             |
| "none"                   |             |
| "userAssignedNATGateway" |             |
| "userDefinedRouting"     |             |

ContainerServiceNetworkProfile_OutboundType_STATUS{#ContainerServiceNetworkProfile_OutboundType_STATUS}
-------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value                    | Description |
|--------------------------|-------------|
| "loadBalancer"           |             |
| "managedNATGateway"      |             |
| "none"                   |             |
| "userAssignedNATGateway" |             |
| "userDefinedRouting"     |             |

ContainerServiceSshConfiguration{#ContainerServiceSshConfiguration}
-------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile](#ContainerServiceLinuxProfile).

| Property   | Description                                                                                                 | Type                                                                                        |
|------------|-------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey[]](#ContainerServiceSshPublicKey)<br/><small>Required</small> |

ContainerServiceSshConfiguration_STATUS{#ContainerServiceSshConfiguration_STATUS}
---------------------------------------------------------------------------------

SSH configuration for Linux-based VMs running on Azure.

Used by: [ContainerServiceLinuxProfile_STATUS](#ContainerServiceLinuxProfile_STATUS).

| Property   | Description                                                                                                 | Type                                                                                                      |
|------------|-------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| publicKeys | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [ContainerServiceSshPublicKey_STATUS[]](#ContainerServiceSshPublicKey_STATUS)<br/><small>Optional</small> |

DelegatedResource{#DelegatedResource}
-------------------------------------

Delegated resource properties - internal use only.

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity).

| Property          | Description                                                                  | Type                                                                                                                                                       |
|-------------------|------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| location          | The source resource location - internal use only.                            | string<br/><small>Optional</small>                                                                                                                         |
| referralResource  | The delegation id of the referral delegation (optional) - internal use only. | string<br/><small>Optional</small>                                                                                                                         |
| resourceReference | The ARM resource id of the delegated resource - internal use only.           | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| tenantId          | The tenant id of the delegated resource - internal use only.                 | string<br/><small>Optional</small>                                                                                                                         |

DelegatedResource_STATUS{#DelegatedResource_STATUS}
---------------------------------------------------

Delegated resource properties - internal use only.

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS).

| Property         | Description                                                                  | Type                               |
|------------------|------------------------------------------------------------------------------|------------------------------------|
| location         | The source resource location - internal use only.                            | string<br/><small>Optional</small> |
| referralResource | The delegation id of the referral delegation (optional) - internal use only. | string<br/><small>Optional</small> |
| resourceId       | The ARM resource id of the delegated resource - internal use only.           | string<br/><small>Optional</small> |
| tenantId         | The tenant id of the delegated resource - internal use only.                 | string<br/><small>Optional</small> |

Expander{#Expander}
-------------------

If not specified, the default is `random`. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.

Used by: [ManagedClusterProperties_AutoScalerProfile](#ManagedClusterProperties_AutoScalerProfile).

| Value         | Description |
|---------------|-------------|
| "least-waste" |             |
| "most-pods"   |             |
| "priority"    |             |
| "random"      |             |

Expander_STATUS{#Expander_STATUS}
---------------------------------

If not specified, the default is `random`. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.

Used by: [ManagedClusterProperties_AutoScalerProfile_STATUS](#ManagedClusterProperties_AutoScalerProfile_STATUS).

| Value         | Description |
|---------------|-------------|
| "least-waste" |             |
| "most-pods"   |             |
| "priority"    |             |
| "random"      |             |

ExtendedLocationType{#ExtendedLocationType}
-------------------------------------------

The type of extendedLocation.

Used by: [ExtendedLocation](#ExtendedLocation).

| Value      | Description |
|------------|-------------|
| "EdgeZone" |             |

ExtendedLocationType_STATUS{#ExtendedLocationType_STATUS}
---------------------------------------------------------

The type of extendedLocation.

Used by: [ExtendedLocation_STATUS](#ExtendedLocation_STATUS).

| Value      | Description |
|------------|-------------|
| "EdgeZone" |             |

IpFamily{#IpFamily}
-------------------

To determine if address belongs IPv4 or IPv6 family.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value  | Description |
|--------|-------------|
| "IPv4" |             |
| "IPv6" |             |

IpFamily_STATUS{#IpFamily_STATUS}
---------------------------------

To determine if address belongs IPv4 or IPv6 family.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value  | Description |
|--------|-------------|
| "IPv4" |             |
| "IPv6" |             |

IPTag{#IPTag}
-------------

Contains the IPTag associated with the object.

Used by: [AgentPoolNetworkProfile](#AgentPoolNetworkProfile).

| Property  | Description                                                               | Type                               |
|-----------|---------------------------------------------------------------------------|------------------------------------|
| ipTagType | The IP tag type. Example: RoutingPreference.                              | string<br/><small>Optional</small> |
| tag       | The value of the IP tag associated with the public IP. Example: Internet. | string<br/><small>Optional</small> |

IPTag_STATUS{#IPTag_STATUS}
---------------------------

Contains the IPTag associated with the object.

Used by: [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS).

| Property  | Description                                                               | Type                               |
|-----------|---------------------------------------------------------------------------|------------------------------------|
| ipTagType | The IP tag type. Example: RoutingPreference.                              | string<br/><small>Optional</small> |
| tag       | The value of the IP tag associated with the public IP. Example: Internet. | string<br/><small>Optional</small> |

IstioServiceMesh{#IstioServiceMesh}
-----------------------------------

Istio service mesh configuration.

Used by: [ServiceMeshProfile](#ServiceMeshProfile).

| Property             | Description                                                                                                                                                                                                                                                                     | Type                                                                                |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| certificateAuthority | Istio Service Mesh Certificate Authority (CA) configuration. For now, we only support plugin certificates as described here https://aka.ms/asm-plugin-ca                                                                                                                        | [IstioCertificateAuthority](#IstioCertificateAuthority)<br/><small>Optional</small> |
| components           | Istio components configuration.                                                                                                                                                                                                                                                 | [IstioComponents](#IstioComponents)<br/><small>Optional</small>                     |
| revisions            | The list of revisions of the Istio control plane. When an upgrade is not in progress, this holds one value. When canary upgrade is in progress, this can only hold two consecutive values. For more information, see: https://learn.microsoft.com/en-us/azure/aks/istio-upgrade | string[]<br/><small>Optional</small>                                                |

IstioServiceMesh_STATUS{#IstioServiceMesh_STATUS}
-------------------------------------------------

Istio service mesh configuration.

Used by: [ServiceMeshProfile_STATUS](#ServiceMeshProfile_STATUS).

| Property             | Description                                                                                                                                                                                                                                                                     | Type                                                                                              |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| certificateAuthority | Istio Service Mesh Certificate Authority (CA) configuration. For now, we only support plugin certificates as described here https://aka.ms/asm-plugin-ca                                                                                                                        | [IstioCertificateAuthority_STATUS](#IstioCertificateAuthority_STATUS)<br/><small>Optional</small> |
| components           | Istio components configuration.                                                                                                                                                                                                                                                 | [IstioComponents_STATUS](#IstioComponents_STATUS)<br/><small>Optional</small>                     |
| revisions            | The list of revisions of the Istio control plane. When an upgrade is not in progress, this holds one value. When canary upgrade is in progress, this can only hold two consecutive values. For more information, see: https://learn.microsoft.com/en-us/azure/aks/istio-upgrade | string[]<br/><small>Optional</small>                                                              |

LoadBalancerSku{#LoadBalancerSku}
---------------------------------

The default is `standard`. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value      | Description |
|------------|-------------|
| "basic"    |             |
| "standard" |             |

LoadBalancerSku_STATUS{#LoadBalancerSku_STATUS}
-----------------------------------------------

The default is `standard`. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value      | Description |
|------------|-------------|
| "basic"    |             |
| "standard" |             |

ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel{#ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel}
-------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile).

| Value           | Description |
|-----------------|-------------|
| "NodeImage"     |             |
| "None"          |             |
| "SecurityPatch" |             |
| "Unmanaged"     |             |

ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS{#ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS}
---------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS).

| Value           | Description |
|-----------------|-------------|
| "NodeImage"     |             |
| "None"          |             |
| "SecurityPatch" |             |
| "Unmanaged"     |             |

ManagedClusterAutoUpgradeProfile_UpgradeChannel{#ManagedClusterAutoUpgradeProfile_UpgradeChannel}
-------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile](#ManagedClusterAutoUpgradeProfile).

| Value        | Description |
|--------------|-------------|
| "node-image" |             |
| "none"       |             |
| "patch"      |             |
| "rapid"      |             |
| "stable"     |             |

ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS{#ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS}
---------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterAutoUpgradeProfile_STATUS](#ManagedClusterAutoUpgradeProfile_STATUS).

| Value        | Description |
|--------------|-------------|
| "node-image" |             |
| "none"       |             |
| "patch"      |             |
| "rapid"      |             |
| "stable"     |             |

ManagedClusterAzureMonitorProfileAppMonitoring{#ManagedClusterAzureMonitorProfileAppMonitoring}
-----------------------------------------------------------------------------------------------

Application Monitoring Profile for Kubernetes Application Container. Collects application logs, metrics and traces through auto-instrumentation of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.

Used by: [ManagedClusterAzureMonitorProfile](#ManagedClusterAzureMonitorProfile).

| Property             | Description                                                                                                                                                                                                                                                                                      | Type                                                                                                                                                                  |
|----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| autoInstrumentation  | Application Monitoring Auto Instrumentation for Kubernetes Application Container. Deploys web hook to auto-instrument Azure Monitor OpenTelemetry based SDKs to collect OpenTelemetry metrics, logs and traces of the application. See aka.ms/AzureMonitorApplicationMonitoring for an overview. | [ManagedClusterAzureMonitorProfileAppMonitoringAutoInstrumentation](#ManagedClusterAzureMonitorProfileAppMonitoringAutoInstrumentation)<br/><small>Optional</small>   |
| openTelemetryLogs    | Application Monitoring Open Telemetry Metrics Profile for Kubernetes Application Container Logs and Traces. Collects OpenTelemetry logs and traces of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.                | [ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogs](#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogs)<br/><small>Optional</small>       |
| openTelemetryMetrics | Application Monitoring Open Telemetry Metrics Profile for Kubernetes Application Container Metrics. Collects OpenTelemetry metrics of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.                                | [ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics](#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics)<br/><small>Optional</small> |

ManagedClusterAzureMonitorProfileAppMonitoring_STATUS{#ManagedClusterAzureMonitorProfileAppMonitoring_STATUS}
-------------------------------------------------------------------------------------------------------------

Application Monitoring Profile for Kubernetes Application Container. Collects application logs, metrics and traces through auto-instrumentation of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.

Used by: [ManagedClusterAzureMonitorProfile_STATUS](#ManagedClusterAzureMonitorProfile_STATUS).

| Property             | Description                                                                                                                                                                                                                                                                                      | Type                                                                                                                                                                                |
|----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| autoInstrumentation  | Application Monitoring Auto Instrumentation for Kubernetes Application Container. Deploys web hook to auto-instrument Azure Monitor OpenTelemetry based SDKs to collect OpenTelemetry metrics, logs and traces of the application. See aka.ms/AzureMonitorApplicationMonitoring for an overview. | [ManagedClusterAzureMonitorProfileAppMonitoringAutoInstrumentation_STATUS](#ManagedClusterAzureMonitorProfileAppMonitoringAutoInstrumentation_STATUS)<br/><small>Optional</small>   |
| openTelemetryLogs    | Application Monitoring Open Telemetry Metrics Profile for Kubernetes Application Container Logs and Traces. Collects OpenTelemetry logs and traces of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.                | [ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogs_STATUS](#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogs_STATUS)<br/><small>Optional</small>       |
| openTelemetryMetrics | Application Monitoring Open Telemetry Metrics Profile for Kubernetes Application Container Metrics. Collects OpenTelemetry metrics of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.                                | [ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics_STATUS](#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics_STATUS)<br/><small>Optional</small> |

ManagedClusterAzureMonitorProfileContainerInsights{#ManagedClusterAzureMonitorProfileContainerInsights}
-------------------------------------------------------------------------------------------------------

Azure Monitor Container Insights Profile for Kubernetes Events, Inventory and Container stdout & stderr logs etc. See aka.ms/AzureMonitorContainerInsights for an overview.

Used by: [ManagedClusterAzureMonitorProfile](#ManagedClusterAzureMonitorProfile).

| Property                               | Description                                                                                                                                                                                                              | Type                                                                                                                                                       |
|----------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| disableCustomMetrics                   | Indicates whether custom metrics collection has to be disabled or not. If not specified the default is false. No custom metrics will be emitted if this field is false but the container insights enabled field is false | bool<br/><small>Optional</small>                                                                                                                           |
| disablePrometheusMetricsScraping       | Indicates whether prometheus metrics scraping is disabled or not. If not specified the default is false. No prometheus metrics will be emitted if this field is false but the container insights enabled field is false  | bool<br/><small>Optional</small>                                                                                                                           |
| enabled                                | Indicates if Azure Monitor Container Insights Logs Addon is enabled or not.                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                           |
| logAnalyticsWorkspaceResourceReference | Fully Qualified ARM Resource Id of Azure Log Analytics Workspace for storing Azure Monitor Container Insights Logs.                                                                                                      | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| syslogPort                             | The syslog host port. If not specified, the default port is 28330.                                                                                                                                                       | int<br/><small>Optional</small>                                                                                                                            |

ManagedClusterAzureMonitorProfileContainerInsights_STATUS{#ManagedClusterAzureMonitorProfileContainerInsights_STATUS}
---------------------------------------------------------------------------------------------------------------------

Azure Monitor Container Insights Profile for Kubernetes Events, Inventory and Container stdout & stderr logs etc. See aka.ms/AzureMonitorContainerInsights for an overview.

Used by: [ManagedClusterAzureMonitorProfile_STATUS](#ManagedClusterAzureMonitorProfile_STATUS).

| Property                         | Description                                                                                                                                                                                                              | Type                               |
|----------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| disableCustomMetrics             | Indicates whether custom metrics collection has to be disabled or not. If not specified the default is false. No custom metrics will be emitted if this field is false but the container insights enabled field is false | bool<br/><small>Optional</small>   |
| disablePrometheusMetricsScraping | Indicates whether prometheus metrics scraping is disabled or not. If not specified the default is false. No prometheus metrics will be emitted if this field is false but the container insights enabled field is false  | bool<br/><small>Optional</small>   |
| enabled                          | Indicates if Azure Monitor Container Insights Logs Addon is enabled or not.                                                                                                                                              | bool<br/><small>Optional</small>   |
| logAnalyticsWorkspaceResourceId  | Fully Qualified ARM Resource Id of Azure Log Analytics Workspace for storing Azure Monitor Container Insights Logs.                                                                                                      | string<br/><small>Optional</small> |
| syslogPort                       | The syslog host port. If not specified, the default port is 28330.                                                                                                                                                       | int<br/><small>Optional</small>    |

ManagedClusterAzureMonitorProfileMetrics{#ManagedClusterAzureMonitorProfileMetrics}
-----------------------------------------------------------------------------------

Metrics profile for the prometheus service addon

Used by: [ManagedClusterAzureMonitorProfile](#ManagedClusterAzureMonitorProfile).

| Property         | Description                                                                       | Type                                                                                                                                |
|------------------|-----------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| enabled          | Whether to enable the Prometheus collector                                        | bool<br/><small>Required</small>                                                                                                    |
| kubeStateMetrics | Kube State Metrics for prometheus addon profile for the container service cluster | [ManagedClusterAzureMonitorProfileKubeStateMetrics](#ManagedClusterAzureMonitorProfileKubeStateMetrics)<br/><small>Optional</small> |

ManagedClusterAzureMonitorProfileMetrics_STATUS{#ManagedClusterAzureMonitorProfileMetrics_STATUS}
-------------------------------------------------------------------------------------------------

Metrics profile for the prometheus service addon

Used by: [ManagedClusterAzureMonitorProfile_STATUS](#ManagedClusterAzureMonitorProfile_STATUS).

| Property         | Description                                                                       | Type                                                                                                                                              |
|------------------|-----------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| enabled          | Whether to enable the Prometheus collector                                        | bool<br/><small>Optional</small>                                                                                                                  |
| kubeStateMetrics | Kube State Metrics for prometheus addon profile for the container service cluster | [ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS](#ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS)<br/><small>Optional</small> |

ManagedClusterBootstrapProfile_ArtifactSource{#ManagedClusterBootstrapProfile_ArtifactSource}
---------------------------------------------------------------------------------------------

Used by: [ManagedClusterBootstrapProfile](#ManagedClusterBootstrapProfile).

| Value    | Description |
|----------|-------------|
| "Cache"  |             |
| "Direct" |             |

ManagedClusterBootstrapProfile_ArtifactSource_STATUS{#ManagedClusterBootstrapProfile_ArtifactSource_STATUS}
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterBootstrapProfile_STATUS](#ManagedClusterBootstrapProfile_STATUS).

| Value    | Description |
|----------|-------------|
| "Cache"  |             |
| "Direct" |             |

ManagedClusterCostAnalysis{#ManagedClusterCostAnalysis}
-------------------------------------------------------

The cost analysis configuration for the cluster

Used by: [ManagedClusterMetricsProfile](#ManagedClusterMetricsProfile).

| Property | Description                                                                                                                                                                                                                                                                                                    | Type                             |
|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | The Managed Cluster sku.tier must be set to `Standard` or `Premium` to enable this feature. Enabling this will add Kubernetes Namespace and Deployment details to the Cost Analysis views in the Azure portal. If not specified, the default is false. For more information see aka.ms/aks/docs/cost-analysis. | bool<br/><small>Optional</small> |

ManagedClusterCostAnalysis_STATUS{#ManagedClusterCostAnalysis_STATUS}
---------------------------------------------------------------------

The cost analysis configuration for the cluster

Used by: [ManagedClusterMetricsProfile_STATUS](#ManagedClusterMetricsProfile_STATUS).

| Property | Description                                                                                                                                                                                                                                                                                                    | Type                             |
|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | The Managed Cluster sku.tier must be set to `Standard` or `Premium` to enable this feature. Enabling this will add Kubernetes Namespace and Deployment details to the Cost Analysis views in the Azure portal. If not specified, the default is false. For more information see aka.ms/aks/docs/cost-analysis. | bool<br/><small>Optional</small> |

ManagedClusterIdentity_Type{#ManagedClusterIdentity_Type}
---------------------------------------------------------

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

ManagedClusterIdentity_Type_STATUS{#ManagedClusterIdentity_Type_STATUS}
-----------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

ManagedClusterIdentity_UserAssignedIdentities_STATUS{#ManagedClusterIdentity_UserAssignedIdentities_STATUS}
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterIdentity_STATUS](#ManagedClusterIdentity_STATUS).

| Property    | Description                                 | Type                               |
|-------------|---------------------------------------------|------------------------------------|
| clientId    | The client id of user assigned identity.    | string<br/><small>Optional</small> |
| principalId | The principal id of user assigned identity. | string<br/><small>Optional</small> |

ManagedClusterIngressProfileWebAppRouting{#ManagedClusterIngressProfileWebAppRouting}
-------------------------------------------------------------------------------------

Web App Routing settings for the ingress profile.

Used by: [ManagedClusterIngressProfile](#ManagedClusterIngressProfile).

| Property                  | Description                                                                                                                                                                                                                                                                                                            | Type                                                                                                                                                         |
|---------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| dnsZoneResourceReferences | Resource IDs of the DNS zones to be associated with the Web App Routing add-on. Used only when Web App Routing is enabled. Public and private DNS zones can be in different resource groups, but all public DNS zones must be in the same resource group and all private DNS zones must be in the same resource group. | [genruntime.ResourceReference[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| enabled                   | Whether to enable Web App Routing.                                                                                                                                                                                                                                                                                     | bool<br/><small>Optional</small>                                                                                                                             |

ManagedClusterIngressProfileWebAppRouting_STATUS{#ManagedClusterIngressProfileWebAppRouting_STATUS}
---------------------------------------------------------------------------------------------------

Web App Routing settings for the ingress profile.

Used by: [ManagedClusterIngressProfile_STATUS](#ManagedClusterIngressProfile_STATUS).

| Property           | Description                                                                                                                                                                                                                                                                                                                                              | Type                                                                                    |
|--------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| dnsZoneResourceIds | Resource IDs of the DNS zones to be associated with the Web App Routing add-on. Used only when Web App Routing is enabled. Public and private DNS zones can be in different resource groups, but all public DNS zones must be in the same resource group and all private DNS zones must be in the same resource group.                                   | string[]<br/><small>Optional</small>                                                    |
| enabled            | Whether to enable Web App Routing.                                                                                                                                                                                                                                                                                                                       | bool<br/><small>Optional</small>                                                        |
| identity           | Managed identity of the Web Application Routing add-on. This is the identity that should be granted permissions, for example, to manage the associated Azure DNS resource and get certificates from Azure Key Vault. See [this overview of the add-on](https://learn.microsoft.com/en-us/azure/aks/web-app-routing?tabs=with-osm) for more instructions. | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small> |

ManagedClusterLoadBalancerProfile{#ManagedClusterLoadBalancerProfile}
---------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property                                  | Description                                                                                                                                                                               | Type                                                                                                                                                                                    |
|-------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts                    | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int<br/><small>Optional</small>                                                                                                                                                         |
| backendPoolType                           | The type of the managed inbound Load Balancer BackendPool.                                                                                                                                | [ManagedClusterLoadBalancerProfile_BackendPoolType](#ManagedClusterLoadBalancerProfile_BackendPoolType)<br/><small>Optional</small>                                                     |
| clusterServiceLoadBalancerHealthProbeMode | The health probing behavior for External Traffic Policy Cluster services.                                                                                                                 | [ManagedClusterLoadBalancerProfile_ClusterServiceLoadBalancerHealthProbeMode](#ManagedClusterLoadBalancerProfile_ClusterServiceLoadBalancerHealthProbeMode)<br/><small>Optional</small> |
| effectiveOutboundIPs                      | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small>                                                                                                                   |
| enableMultipleStandardLoadBalancers       | Enable multiple standard load balancers per AKS cluster or not.                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                                                        |
| idleTimeoutInMinutes                      | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int<br/><small>Optional</small>                                                                                                                                                         |
| managedOutboundIPs                        | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs)<br/><small>Optional</small>                                               |
| outboundIPPrefixes                        | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes)<br/><small>Optional</small>                                               |
| outboundIPs                               | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs](#ManagedClusterLoadBalancerProfile_OutboundIPs)<br/><small>Optional</small>                                                             |

ManagedClusterLoadBalancerProfile_STATUS{#ManagedClusterLoadBalancerProfile_STATUS}
-----------------------------------------------------------------------------------

Profile of the managed cluster load balancer.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property                                  | Description                                                                                                                                                                               | Type                                                                                                                                                                                                  |
|-------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| allocatedOutboundPorts                    | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | int<br/><small>Optional</small>                                                                                                                                                                       |
| backendPoolType                           | The type of the managed inbound Load Balancer BackendPool.                                                                                                                                | [ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS](#ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS)<br/><small>Optional</small>                                                     |
| clusterServiceLoadBalancerHealthProbeMode | The health probing behavior for External Traffic Policy Cluster services.                                                                                                                 | [ManagedClusterLoadBalancerProfile_ClusterServiceLoadBalancerHealthProbeMode_STATUS](#ManagedClusterLoadBalancerProfile_ClusterServiceLoadBalancerHealthProbeMode_STATUS)<br/><small>Optional</small> |
| effectiveOutboundIPs                      | The effective outbound IP resources of the cluster load balancer.                                                                                                                         | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small>                                                                                                                   |
| enableMultipleStandardLoadBalancers       | Enable multiple standard load balancers per AKS cluster or not.                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                                                                      |
| idleTimeoutInMinutes                      | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes.                                                  | int<br/><small>Optional</small>                                                                                                                                                                       |
| managedOutboundIPs                        | Desired managed outbound IPs for the cluster load balancer.                                                                                                                               | [ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS)<br/><small>Optional</small>                                               |
| outboundIPPrefixes                        | Desired outbound IP Prefix resources for the cluster load balancer.                                                                                                                       | [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS)<br/><small>Optional</small>                                               |
| outboundIPs                               | Desired outbound IP resources for the cluster load balancer.                                                                                                                              | [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS)<br/><small>Optional</small>                                                             |

ManagedClusterNATGatewayProfile{#ManagedClusterNATGatewayProfile}
-----------------------------------------------------------------

Profile of the managed cluster NAT gateway.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property                 | Description                                                                                                                             | Type                                                                                                          |
|--------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| effectiveOutboundIPs     | The effective outbound IP resources of the cluster NAT gateway.                                                                         | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small>                                         |
| idleTimeoutInMinutes     | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 4 minutes. | int<br/><small>Optional</small>                                                                               |
| managedOutboundIPProfile | Profile of the managed outbound IP resources of the cluster NAT gateway.                                                                | [ManagedClusterManagedOutboundIPProfile](#ManagedClusterManagedOutboundIPProfile)<br/><small>Optional</small> |

ManagedClusterNATGatewayProfile_STATUS{#ManagedClusterNATGatewayProfile_STATUS}
-------------------------------------------------------------------------------

Profile of the managed cluster NAT gateway.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property                 | Description                                                                                                                             | Type                                                                                                                        |
|--------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| effectiveOutboundIPs     | The effective outbound IP resources of the cluster NAT gateway.                                                                         | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small>                                         |
| idleTimeoutInMinutes     | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 4 minutes. | int<br/><small>Optional</small>                                                                                             |
| managedOutboundIPProfile | Profile of the managed outbound IP resources of the cluster NAT gateway.                                                                | [ManagedClusterManagedOutboundIPProfile_STATUS](#ManagedClusterManagedOutboundIPProfile_STATUS)<br/><small>Optional</small> |

ManagedClusterNodeProvisioningProfile_Mode{#ManagedClusterNodeProvisioningProfile_Mode}
---------------------------------------------------------------------------------------

Used by: [ManagedClusterNodeProvisioningProfile](#ManagedClusterNodeProvisioningProfile).

| Value    | Description |
|----------|-------------|
| "Auto"   |             |
| "Manual" |             |

ManagedClusterNodeProvisioningProfile_Mode_STATUS{#ManagedClusterNodeProvisioningProfile_Mode_STATUS}
-----------------------------------------------------------------------------------------------------

Used by: [ManagedClusterNodeProvisioningProfile_STATUS](#ManagedClusterNodeProvisioningProfile_STATUS).

| Value    | Description |
|----------|-------------|
| "Auto"   |             |
| "Manual" |             |

ManagedClusterNodeResourceGroupProfile_RestrictionLevel{#ManagedClusterNodeResourceGroupProfile_RestrictionLevel}
-----------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterNodeResourceGroupProfile](#ManagedClusterNodeResourceGroupProfile).

| Value          | Description |
|----------------|-------------|
| "ReadOnly"     |             |
| "Unrestricted" |             |

ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS{#ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS}
-------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterNodeResourceGroupProfile_STATUS](#ManagedClusterNodeResourceGroupProfile_STATUS).

| Value          | Description |
|----------------|-------------|
| "ReadOnly"     |             |
| "Unrestricted" |             |

ManagedClusterOperatorConfigMaps{#ManagedClusterOperatorConfigMaps}
-------------------------------------------------------------------

Used by: [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec).

| Property          | Description                                                                                                   | Type                                                                                                                                                             |
|-------------------|---------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| oidcIssuerProfile | indicates where the OIDCIssuerProfile config map should be placed. If omitted, no config map will be created. | [genruntime.ConfigMapDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapDestination)<br/><small>Optional</small> |
| principalId       | indicates where the PrincipalId config map should be placed. If omitted, no config map will be created.       | [genruntime.ConfigMapDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapDestination)<br/><small>Optional</small> |

ManagedClusterOperatorSecrets{#ManagedClusterOperatorSecrets}
-------------------------------------------------------------

Used by: [ManagedClusterOperatorSpec](#ManagedClusterOperatorSpec).

| Property         | Description                                                                                                            | Type                                                                                                                                                       |
|------------------|------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| adminCredentials | indicates where the AdminCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure. | [genruntime.SecretDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination)<br/><small>Optional</small> |
| userCredentials  | indicates where the UserCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure.  | [genruntime.SecretDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination)<br/><small>Optional</small> |

ManagedClusterPodIdentity{#ManagedClusterPodIdentity}
-----------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile).

| Property        | Description                                                        | Type                                                                      |
|-----------------|--------------------------------------------------------------------|---------------------------------------------------------------------------|
| bindingSelector | The binding selector to use for the AzureIdentityBinding resource. | string<br/><small>Optional</small>                                        |
| identity        | The user assigned identity details.                                | [UserAssignedIdentity](#UserAssignedIdentity)<br/><small>Required</small> |
| name            | The name of the pod identity.                                      | string<br/><small>Required</small>                                        |
| namespace       | The namespace of the pod identity.                                 | string<br/><small>Required</small>                                        |

ManagedClusterPodIdentity_STATUS{#ManagedClusterPodIdentity_STATUS}
-------------------------------------------------------------------

Details about the pod identity assigned to the Managed Cluster.

Used by: [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS).

| Property          | Description                                                        | Type                                                                                                                                  |
|-------------------|--------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| bindingSelector   | The binding selector to use for the AzureIdentityBinding resource. | string<br/><small>Optional</small>                                                                                                    |
| identity          | The user assigned identity details.                                | [UserAssignedIdentity_STATUS](#UserAssignedIdentity_STATUS)<br/><small>Optional</small>                                               |
| name              | The name of the pod identity.                                      | string<br/><small>Optional</small>                                                                                                    |
| namespace         | The namespace of the pod identity.                                 | string<br/><small>Optional</small>                                                                                                    |
| provisioningInfo  |                                                                    | [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS)<br/><small>Optional</small>   |
| provisioningState | The current provisioning state of the pod identity.                | [ManagedClusterPodIdentity_ProvisioningState_STATUS](#ManagedClusterPodIdentity_ProvisioningState_STATUS)<br/><small>Optional</small> |

ManagedClusterPodIdentityException{#ManagedClusterPodIdentityException}
-----------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile](#ManagedClusterPodIdentityProfile).

| Property  | Description                                  | Type                                          |
|-----------|----------------------------------------------|-----------------------------------------------|
| name      | The name of the pod identity exception.      | string<br/><small>Required</small>            |
| namespace | The namespace of the pod identity exception. | string<br/><small>Required</small>            |
| podLabels | The pod labels to match.                     | map[string]string<br/><small>Required</small> |

ManagedClusterPodIdentityException_STATUS{#ManagedClusterPodIdentityException_STATUS}
-------------------------------------------------------------------------------------

See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.

Used by: [ManagedClusterPodIdentityProfile_STATUS](#ManagedClusterPodIdentityProfile_STATUS).

| Property  | Description                                  | Type                                          |
|-----------|----------------------------------------------|-----------------------------------------------|
| name      | The name of the pod identity exception.      | string<br/><small>Optional</small>            |
| namespace | The namespace of the pod identity exception. | string<br/><small>Optional</small>            |
| podLabels | The pod labels to match.                     | map[string]string<br/><small>Optional</small> |

ManagedClusterSecurityProfileDefender{#ManagedClusterSecurityProfileDefender}
-----------------------------------------------------------------------------

Microsoft Defender settings for the security profile.

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property                               | Description                                                                                                                                                                                                                                            | Type                                                                                                                                                       |
|----------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| logAnalyticsWorkspaceResourceReference | Resource ID of the Log Analytics workspace to be associated with Microsoft Defender. When Microsoft Defender is enabled, this field is required and must be a valid workspace resource ID. When Microsoft Defender is disabled, leave the field empty. | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| securityMonitoring                     | Microsoft Defender threat detection for Cloud settings for the security profile.                                                                                                                                                                       | [ManagedClusterSecurityProfileDefenderSecurityMonitoring](#ManagedClusterSecurityProfileDefenderSecurityMonitoring)<br/><small>Optional</small>            |

ManagedClusterSecurityProfileDefender_STATUS{#ManagedClusterSecurityProfileDefender_STATUS}
-------------------------------------------------------------------------------------------

Microsoft Defender settings for the security profile.

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property                        | Description                                                                                                                                                                                                                                            | Type                                                                                                                                                          |
|---------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| logAnalyticsWorkspaceResourceId | Resource ID of the Log Analytics workspace to be associated with Microsoft Defender. When Microsoft Defender is enabled, this field is required and must be a valid workspace resource ID. When Microsoft Defender is disabled, leave the field empty. | string<br/><small>Optional</small>                                                                                                                            |
| securityMonitoring              | Microsoft Defender threat detection for Cloud settings for the security profile.                                                                                                                                                                       | [ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS](#ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS)<br/><small>Optional</small> |

ManagedClusterSecurityProfileImageCleaner{#ManagedClusterSecurityProfileImageCleaner}
-------------------------------------------------------------------------------------

Image Cleaner removes unused images from nodes, freeing up disk space and helping to reduce attack surface area. Here are settings for the security profile.

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property      | Description                                     | Type                             |
|---------------|-------------------------------------------------|----------------------------------|
| enabled       | Whether to enable Image Cleaner on AKS cluster. | bool<br/><small>Optional</small> |
| intervalHours | Image Cleaner scanning interval in hours.       | int<br/><small>Optional</small>  |

ManagedClusterSecurityProfileImageCleaner_STATUS{#ManagedClusterSecurityProfileImageCleaner_STATUS}
---------------------------------------------------------------------------------------------------

Image Cleaner removes unused images from nodes, freeing up disk space and helping to reduce attack surface area. Here are settings for the security profile.

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property      | Description                                     | Type                             |
|---------------|-------------------------------------------------|----------------------------------|
| enabled       | Whether to enable Image Cleaner on AKS cluster. | bool<br/><small>Optional</small> |
| intervalHours | Image Cleaner scanning interval in hours.       | int<br/><small>Optional</small>  |

ManagedClusterSecurityProfileImageIntegrity{#ManagedClusterSecurityProfileImageIntegrity}
-----------------------------------------------------------------------------------------

Image integrity related settings for the security profile.

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property | Description                                                    | Type                             |
|----------|----------------------------------------------------------------|----------------------------------|
| enabled  | Whether to enable image integrity. The default value is false. | bool<br/><small>Optional</small> |

ManagedClusterSecurityProfileImageIntegrity_STATUS{#ManagedClusterSecurityProfileImageIntegrity_STATUS}
-------------------------------------------------------------------------------------------------------

Image integrity related settings for the security profile.

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property | Description                                                    | Type                             |
|----------|----------------------------------------------------------------|----------------------------------|
| enabled  | Whether to enable image integrity. The default value is false. | bool<br/><small>Optional</small> |

ManagedClusterSecurityProfileNodeRestriction{#ManagedClusterSecurityProfileNodeRestriction}
-------------------------------------------------------------------------------------------

Node Restriction settings for the security profile.

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property | Description                        | Type                             |
|----------|------------------------------------|----------------------------------|
| enabled  | Whether to enable Node Restriction | bool<br/><small>Optional</small> |

ManagedClusterSecurityProfileNodeRestriction_STATUS{#ManagedClusterSecurityProfileNodeRestriction_STATUS}
---------------------------------------------------------------------------------------------------------

Node Restriction settings for the security profile.

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property | Description                        | Type                             |
|----------|------------------------------------|----------------------------------|
| enabled  | Whether to enable Node Restriction | bool<br/><small>Optional</small> |

ManagedClusterSecurityProfileWorkloadIdentity{#ManagedClusterSecurityProfileWorkloadIdentity}
---------------------------------------------------------------------------------------------

Workload identity settings for the security profile.

Used by: [ManagedClusterSecurityProfile](#ManagedClusterSecurityProfile).

| Property | Description                          | Type                             |
|----------|--------------------------------------|----------------------------------|
| enabled  | Whether to enable workload identity. | bool<br/><small>Optional</small> |

ManagedClusterSecurityProfileWorkloadIdentity_STATUS{#ManagedClusterSecurityProfileWorkloadIdentity_STATUS}
-----------------------------------------------------------------------------------------------------------

Workload identity settings for the security profile.

Used by: [ManagedClusterSecurityProfile_STATUS](#ManagedClusterSecurityProfile_STATUS).

| Property | Description                          | Type                             |
|----------|--------------------------------------|----------------------------------|
| enabled  | Whether to enable workload identity. | bool<br/><small>Optional</small> |

ManagedClusterSKU_Name{#ManagedClusterSKU_Name}
-----------------------------------------------

Used by: [ManagedClusterSKU](#ManagedClusterSKU).

| Value       | Description |
|-------------|-------------|
| "Automatic" |             |
| "Base"      |             |

ManagedClusterSKU_Name_STATUS{#ManagedClusterSKU_Name_STATUS}
-------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS).

| Value       | Description |
|-------------|-------------|
| "Automatic" |             |
| "Base"      |             |

ManagedClusterSKU_Tier{#ManagedClusterSKU_Tier}
-----------------------------------------------

Used by: [ManagedClusterSKU](#ManagedClusterSKU).

| Value      | Description |
|------------|-------------|
| "Free"     |             |
| "Premium"  |             |
| "Standard" |             |

ManagedClusterSKU_Tier_STATUS{#ManagedClusterSKU_Tier_STATUS}
-------------------------------------------------------------

Used by: [ManagedClusterSKU_STATUS](#ManagedClusterSKU_STATUS).

| Value      | Description |
|------------|-------------|
| "Free"     |             |
| "Premium"  |             |
| "Standard" |             |

ManagedClusterStaticEgressGatewayProfile{#ManagedClusterStaticEgressGatewayProfile}
-----------------------------------------------------------------------------------

The Static Egress Gateway addon configuration for the cluster.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Property | Description                                                 | Type                             |
|----------|-------------------------------------------------------------|----------------------------------|
| enabled  | Indicates if Static Egress Gateway addon is enabled or not. | bool<br/><small>Optional</small> |

ManagedClusterStaticEgressGatewayProfile_STATUS{#ManagedClusterStaticEgressGatewayProfile_STATUS}
-------------------------------------------------------------------------------------------------

The Static Egress Gateway addon configuration for the cluster.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Property | Description                                                 | Type                             |
|----------|-------------------------------------------------------------|----------------------------------|
| enabled  | Indicates if Static Egress Gateway addon is enabled or not. | bool<br/><small>Optional</small> |

ManagedClusterStorageProfileBlobCSIDriver{#ManagedClusterStorageProfileBlobCSIDriver}
-------------------------------------------------------------------------------------

AzureBlob CSI Driver settings for the storage profile.

Used by: [ManagedClusterStorageProfile](#ManagedClusterStorageProfile).

| Property | Description                                                         | Type                             |
|----------|---------------------------------------------------------------------|----------------------------------|
| enabled  | Whether to enable AzureBlob CSI Driver. The default value is false. | bool<br/><small>Optional</small> |

ManagedClusterStorageProfileBlobCSIDriver_STATUS{#ManagedClusterStorageProfileBlobCSIDriver_STATUS}
---------------------------------------------------------------------------------------------------

AzureBlob CSI Driver settings for the storage profile.

Used by: [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS).

| Property | Description                                                         | Type                             |
|----------|---------------------------------------------------------------------|----------------------------------|
| enabled  | Whether to enable AzureBlob CSI Driver. The default value is false. | bool<br/><small>Optional</small> |

ManagedClusterStorageProfileDiskCSIDriver{#ManagedClusterStorageProfileDiskCSIDriver}
-------------------------------------------------------------------------------------

AzureDisk CSI Driver settings for the storage profile.

Used by: [ManagedClusterStorageProfile](#ManagedClusterStorageProfile).

| Property | Description                                                        | Type                               |
|----------|--------------------------------------------------------------------|------------------------------------|
| enabled  | Whether to enable AzureDisk CSI Driver. The default value is true. | bool<br/><small>Optional</small>   |
| version  | The version of AzureDisk CSI Driver. The default value is v1.      | string<br/><small>Optional</small> |

ManagedClusterStorageProfileDiskCSIDriver_STATUS{#ManagedClusterStorageProfileDiskCSIDriver_STATUS}
---------------------------------------------------------------------------------------------------

AzureDisk CSI Driver settings for the storage profile.

Used by: [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS).

| Property | Description                                                        | Type                               |
|----------|--------------------------------------------------------------------|------------------------------------|
| enabled  | Whether to enable AzureDisk CSI Driver. The default value is true. | bool<br/><small>Optional</small>   |
| version  | The version of AzureDisk CSI Driver. The default value is v1.      | string<br/><small>Optional</small> |

ManagedClusterStorageProfileFileCSIDriver{#ManagedClusterStorageProfileFileCSIDriver}
-------------------------------------------------------------------------------------

AzureFile CSI Driver settings for the storage profile.

Used by: [ManagedClusterStorageProfile](#ManagedClusterStorageProfile).

| Property | Description                                                        | Type                             |
|----------|--------------------------------------------------------------------|----------------------------------|
| enabled  | Whether to enable AzureFile CSI Driver. The default value is true. | bool<br/><small>Optional</small> |

ManagedClusterStorageProfileFileCSIDriver_STATUS{#ManagedClusterStorageProfileFileCSIDriver_STATUS}
---------------------------------------------------------------------------------------------------

AzureFile CSI Driver settings for the storage profile.

Used by: [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS).

| Property | Description                                                        | Type                             |
|----------|--------------------------------------------------------------------|----------------------------------|
| enabled  | Whether to enable AzureFile CSI Driver. The default value is true. | bool<br/><small>Optional</small> |

ManagedClusterStorageProfileSnapshotController{#ManagedClusterStorageProfileSnapshotController}
-----------------------------------------------------------------------------------------------

Snapshot Controller settings for the storage profile.

Used by: [ManagedClusterStorageProfile](#ManagedClusterStorageProfile).

| Property | Description                                                       | Type                             |
|----------|-------------------------------------------------------------------|----------------------------------|
| enabled  | Whether to enable Snapshot Controller. The default value is true. | bool<br/><small>Optional</small> |

ManagedClusterStorageProfileSnapshotController_STATUS{#ManagedClusterStorageProfileSnapshotController_STATUS}
-------------------------------------------------------------------------------------------------------------

Snapshot Controller settings for the storage profile.

Used by: [ManagedClusterStorageProfile_STATUS](#ManagedClusterStorageProfile_STATUS).

| Property | Description                                                       | Type                             |
|----------|-------------------------------------------------------------------|----------------------------------|
| enabled  | Whether to enable Snapshot Controller. The default value is true. | bool<br/><small>Optional</small> |

ManagedClusterWindowsProfile_LicenseType{#ManagedClusterWindowsProfile_LicenseType}
-----------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "Windows_Server" |             |

ManagedClusterWindowsProfile_LicenseType_STATUS{#ManagedClusterWindowsProfile_LicenseType_STATUS}
-------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "Windows_Server" |             |

ManagedClusterWorkloadAutoScalerProfileKeda{#ManagedClusterWorkloadAutoScalerProfileKeda}
-----------------------------------------------------------------------------------------

KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.

Used by: [ManagedClusterWorkloadAutoScalerProfile](#ManagedClusterWorkloadAutoScalerProfile).

| Property | Description             | Type                             |
|----------|-------------------------|----------------------------------|
| enabled  | Whether to enable KEDA. | bool<br/><small>Required</small> |

ManagedClusterWorkloadAutoScalerProfileKeda_STATUS{#ManagedClusterWorkloadAutoScalerProfileKeda_STATUS}
-------------------------------------------------------------------------------------------------------

KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.

Used by: [ManagedClusterWorkloadAutoScalerProfile_STATUS](#ManagedClusterWorkloadAutoScalerProfile_STATUS).

| Property | Description             | Type                             |
|----------|-------------------------|----------------------------------|
| enabled  | Whether to enable KEDA. | bool<br/><small>Optional</small> |

ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler{#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler}
---------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfile](#ManagedClusterWorkloadAutoScalerProfile).

| Property         | Description                                                                | Type                                                                                                                                                                                        |
|------------------|----------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| addonAutoscaling | Whether VPA add-on is enabled and configured to scale AKS-managed add-ons. | [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling)<br/><small>Optional</small> |
| enabled          | Whether to enable VPA add-on in cluster. Default value is false.           | bool<br/><small>Required</small>                                                                                                                                                            |

ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS{#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS}
-----------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfile_STATUS](#ManagedClusterWorkloadAutoScalerProfile_STATUS).

| Property         | Description                                                                | Type                                                                                                                                                                                                      |
|------------------|----------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| addonAutoscaling | Whether VPA add-on is enabled and configured to scale AKS-managed add-ons. | [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling_STATUS](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling_STATUS)<br/><small>Optional</small> |
| enabled          | Whether to enable VPA add-on in cluster. Default value is false.           | bool<br/><small>Optional</small>                                                                                                                                                                          |

NetworkDataplane{#NetworkDataplane}
-----------------------------------

Network dataplane used in the Kubernetes cluster.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "cilium" |             |

NetworkDataplane_STATUS{#NetworkDataplane_STATUS}
-------------------------------------------------

Network dataplane used in the Kubernetes cluster.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "cilium" |             |

NetworkMode{#NetworkMode}
-------------------------

This cannot be specified if networkPlugin is anything other than `azure`.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value         | Description |
|---------------|-------------|
| "bridge"      |             |
| "transparent" |             |

NetworkMode_STATUS{#NetworkMode_STATUS}
---------------------------------------

This cannot be specified if networkPlugin is anything other than `azure`.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value         | Description |
|---------------|-------------|
| "bridge"      |             |
| "transparent" |             |

NetworkPlugin{#NetworkPlugin}
-----------------------------

Network plugin used for building the Kubernetes network.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value     | Description |
|-----------|-------------|
| "azure"   |             |
| "kubenet" |             |
| "none"    |             |

NetworkPlugin_STATUS{#NetworkPlugin_STATUS}
-------------------------------------------

Network plugin used for building the Kubernetes network.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value     | Description |
|-----------|-------------|
| "azure"   |             |
| "kubenet" |             |
| "none"    |             |

NetworkPluginMode{#NetworkPluginMode}
-------------------------------------

The mode the network plugin should use.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value     | Description |
|-----------|-------------|
| "overlay" |             |

NetworkPluginMode_STATUS{#NetworkPluginMode_STATUS}
---------------------------------------------------

The mode the network plugin should use.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value     | Description |
|-----------|-------------|
| "overlay" |             |

NetworkPolicy{#NetworkPolicy}
-----------------------------

Network policy used for building the Kubernetes network.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "calico" |             |
| "cilium" |             |
| "none"   |             |

NetworkPolicy_STATUS{#NetworkPolicy_STATUS}
-------------------------------------------

Network policy used for building the Kubernetes network.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value    | Description |
|----------|-------------|
| "azure"  |             |
| "calico" |             |
| "cilium" |             |
| "none"   |             |

PodLinkLocalAccess{#PodLinkLocalAccess}
---------------------------------------

Defines access to special link local addresses (Azure Instance Metadata Service, aka IMDS) for pods with hostNetwork=false. If not specified, the default is `IMDS`.

Used by: [ContainerServiceNetworkProfile](#ContainerServiceNetworkProfile).

| Value  | Description |
|--------|-------------|
| "IMDS" |             |
| "None" |             |

PodLinkLocalAccess_STATUS{#PodLinkLocalAccess_STATUS}
-----------------------------------------------------

Defines access to special link local addresses (Azure Instance Metadata Service, aka IMDS) for pods with hostNetwork=false. If not specified, the default is `IMDS`.

Used by: [ContainerServiceNetworkProfile_STATUS](#ContainerServiceNetworkProfile_STATUS).

| Value  | Description |
|--------|-------------|
| "IMDS" |             |
| "None" |             |

PortRange{#PortRange}
---------------------

The port range.

Used by: [AgentPoolNetworkProfile](#AgentPoolNetworkProfile).

| Property  | Description                                                                                                                     | Type                                                                  |
|-----------|---------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------|
| portEnd   | The maximum port that is included in the range. It should be ranged from 1 to 65535, and be greater than or equal to portStart. | int<br/><small>Optional</small>                                       |
| portStart | The minimum port that is included in the range. It should be ranged from 1 to 65535, and be less than or equal to portEnd.      | int<br/><small>Optional</small>                                       |
| protocol  | The network protocol of the port.                                                                                               | [PortRange_Protocol](#PortRange_Protocol)<br/><small>Optional</small> |

PortRange_STATUS{#PortRange_STATUS}
-----------------------------------

The port range.

Used by: [AgentPoolNetworkProfile_STATUS](#AgentPoolNetworkProfile_STATUS).

| Property  | Description                                                                                                                     | Type                                                                                |
|-----------|---------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| portEnd   | The maximum port that is included in the range. It should be ranged from 1 to 65535, and be greater than or equal to portStart. | int<br/><small>Optional</small>                                                     |
| portStart | The minimum port that is included in the range. It should be ranged from 1 to 65535, and be less than or equal to portEnd.      | int<br/><small>Optional</small>                                                     |
| protocol  | The network protocol of the port.                                                                                               | [PortRange_Protocol_STATUS](#PortRange_Protocol_STATUS)<br/><small>Optional</small> |

PowerState_Code{#PowerState_Code}
---------------------------------

Used by: [PowerState](#PowerState).

| Value     | Description |
|-----------|-------------|
| "Running" |             |
| "Stopped" |             |

PowerState_Code_STATUS{#PowerState_Code_STATUS}
-----------------------------------------------

Used by: [PowerState_STATUS](#PowerState_STATUS).

| Value     | Description |
|-----------|-------------|
| "Running" |             |
| "Stopped" |             |

SafeguardsProfile_Level{#SafeguardsProfile_Level}
-------------------------------------------------

Used by: [SafeguardsProfile](#SafeguardsProfile).

| Value         | Description |
|---------------|-------------|
| "Enforcement" |             |
| "Off"         |             |
| "Warning"     |             |

SafeguardsProfile_Level_STATUS{#SafeguardsProfile_Level_STATUS}
---------------------------------------------------------------

Used by: [SafeguardsProfile_STATUS](#SafeguardsProfile_STATUS).

| Value         | Description |
|---------------|-------------|
| "Enforcement" |             |
| "Off"         |             |
| "Warning"     |             |

ScaleProfile{#ScaleProfile}
---------------------------

Specifications on how to scale a VirtualMachines agent pool.

Used by: [VirtualMachinesProfile](#VirtualMachinesProfile).

| Property  | Description                                                                                                                                            | Type                                                                    |
|-----------|--------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| autoscale | Specifications on how to auto-scale the VirtualMachines agent pool within a predefined size range. Currently, at most one AutoScaleProfile is allowed. | [AutoScaleProfile[]](#AutoScaleProfile)<br/><small>Optional</small>     |
| manual    | Specifications on how to scale the VirtualMachines agent pool to a fixed size.                                                                         | [ManualScaleProfile[]](#ManualScaleProfile)<br/><small>Optional</small> |

ScaleProfile_STATUS{#ScaleProfile_STATUS}
-----------------------------------------

Specifications on how to scale a VirtualMachines agent pool.

Used by: [VirtualMachinesProfile_STATUS](#VirtualMachinesProfile_STATUS).

| Property  | Description                                                                                                                                            | Type                                                                                  |
|-----------|--------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| autoscale | Specifications on how to auto-scale the VirtualMachines agent pool within a predefined size range. Currently, at most one AutoScaleProfile is allowed. | [AutoScaleProfile_STATUS[]](#AutoScaleProfile_STATUS)<br/><small>Optional</small>     |
| manual    | Specifications on how to scale the VirtualMachines agent pool to a fixed size.                                                                         | [ManualScaleProfile_STATUS[]](#ManualScaleProfile_STATUS)<br/><small>Optional</small> |

ServiceMeshProfile_Mode{#ServiceMeshProfile_Mode}
-------------------------------------------------

Used by: [ServiceMeshProfile](#ServiceMeshProfile).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Istio"    |             |

ServiceMeshProfile_Mode_STATUS{#ServiceMeshProfile_Mode_STATUS}
---------------------------------------------------------------

Used by: [ServiceMeshProfile_STATUS](#ServiceMeshProfile_STATUS).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Istio"    |             |

SysctlConfig{#SysctlConfig}
---------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig](#LinuxOSConfig).

| Property                       | Description                                        | Type                               |
|--------------------------------|----------------------------------------------------|------------------------------------|
| fsAioMaxNr                     | Sysctl setting fs.aio-max-nr.                      | int<br/><small>Optional</small>    |
| fsFileMax                      | Sysctl setting fs.file-max.                        | int<br/><small>Optional</small>    |
| fsInotifyMaxUserWatches        | Sysctl setting fs.inotify.max_user_watches.        | int<br/><small>Optional</small>    |
| fsNrOpen                       | Sysctl setting fs.nr_open.                         | int<br/><small>Optional</small>    |
| kernelThreadsMax               | Sysctl setting kernel.threads-max.                 | int<br/><small>Optional</small>    |
| netCoreNetdevMaxBacklog        | Sysctl setting net.core.netdev_max_backlog.        | int<br/><small>Optional</small>    |
| netCoreOptmemMax               | Sysctl setting net.core.optmem_max.                | int<br/><small>Optional</small>    |
| netCoreRmemDefault             | Sysctl setting net.core.rmem_default.              | int<br/><small>Optional</small>    |
| netCoreRmemMax                 | Sysctl setting net.core.rmem_max.                  | int<br/><small>Optional</small>    |
| netCoreSomaxconn               | Sysctl setting net.core.somaxconn.                 | int<br/><small>Optional</small>    |
| netCoreWmemDefault             | Sysctl setting net.core.wmem_default.              | int<br/><small>Optional</small>    |
| netCoreWmemMax                 | Sysctl setting net.core.wmem_max.                  | int<br/><small>Optional</small>    |
| netIpv4IpLocalPortRange        | Sysctl setting net.ipv4.ip_local_port_range.       | string<br/><small>Optional</small> |
| netIpv4NeighDefaultGcThresh1   | Sysctl setting net.ipv4.neigh.default.gc_thresh1.  | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh2   | Sysctl setting net.ipv4.neigh.default.gc_thresh2.  | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh3   | Sysctl setting net.ipv4.neigh.default.gc_thresh3.  | int<br/><small>Optional</small>    |
| netIpv4TcpFinTimeout           | Sysctl setting net.ipv4.tcp_fin_timeout.           | int<br/><small>Optional</small>    |
| netIpv4TcpkeepaliveIntvl       | Sysctl setting net.ipv4.tcp_keepalive_intvl.       | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveProbes      | Sysctl setting net.ipv4.tcp_keepalive_probes.      | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveTime        | Sysctl setting net.ipv4.tcp_keepalive_time.        | int<br/><small>Optional</small>    |
| netIpv4TcpMaxSynBacklog        | Sysctl setting net.ipv4.tcp_max_syn_backlog.       | int<br/><small>Optional</small>    |
| netIpv4TcpMaxTwBuckets         | Sysctl setting net.ipv4.tcp_max_tw_buckets.        | int<br/><small>Optional</small>    |
| netIpv4TcpTwReuse              | Sysctl setting net.ipv4.tcp_tw_reuse.              | bool<br/><small>Optional</small>   |
| netNetfilterNfConntrackBuckets | Sysctl setting net.netfilter.nf_conntrack_buckets. | int<br/><small>Optional</small>    |
| netNetfilterNfConntrackMax     | Sysctl setting net.netfilter.nf_conntrack_max.     | int<br/><small>Optional</small>    |
| vmMaxMapCount                  | Sysctl setting vm.max_map_count.                   | int<br/><small>Optional</small>    |
| vmSwappiness                   | Sysctl setting vm.swappiness.                      | int<br/><small>Optional</small>    |
| vmVfsCachePressure             | Sysctl setting vm.vfs_cache_pressure.              | int<br/><small>Optional</small>    |

SysctlConfig_STATUS{#SysctlConfig_STATUS}
-----------------------------------------

Sysctl settings for Linux agent nodes.

Used by: [LinuxOSConfig_STATUS](#LinuxOSConfig_STATUS).

| Property                       | Description                                        | Type                               |
|--------------------------------|----------------------------------------------------|------------------------------------|
| fsAioMaxNr                     | Sysctl setting fs.aio-max-nr.                      | int<br/><small>Optional</small>    |
| fsFileMax                      | Sysctl setting fs.file-max.                        | int<br/><small>Optional</small>    |
| fsInotifyMaxUserWatches        | Sysctl setting fs.inotify.max_user_watches.        | int<br/><small>Optional</small>    |
| fsNrOpen                       | Sysctl setting fs.nr_open.                         | int<br/><small>Optional</small>    |
| kernelThreadsMax               | Sysctl setting kernel.threads-max.                 | int<br/><small>Optional</small>    |
| netCoreNetdevMaxBacklog        | Sysctl setting net.core.netdev_max_backlog.        | int<br/><small>Optional</small>    |
| netCoreOptmemMax               | Sysctl setting net.core.optmem_max.                | int<br/><small>Optional</small>    |
| netCoreRmemDefault             | Sysctl setting net.core.rmem_default.              | int<br/><small>Optional</small>    |
| netCoreRmemMax                 | Sysctl setting net.core.rmem_max.                  | int<br/><small>Optional</small>    |
| netCoreSomaxconn               | Sysctl setting net.core.somaxconn.                 | int<br/><small>Optional</small>    |
| netCoreWmemDefault             | Sysctl setting net.core.wmem_default.              | int<br/><small>Optional</small>    |
| netCoreWmemMax                 | Sysctl setting net.core.wmem_max.                  | int<br/><small>Optional</small>    |
| netIpv4IpLocalPortRange        | Sysctl setting net.ipv4.ip_local_port_range.       | string<br/><small>Optional</small> |
| netIpv4NeighDefaultGcThresh1   | Sysctl setting net.ipv4.neigh.default.gc_thresh1.  | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh2   | Sysctl setting net.ipv4.neigh.default.gc_thresh2.  | int<br/><small>Optional</small>    |
| netIpv4NeighDefaultGcThresh3   | Sysctl setting net.ipv4.neigh.default.gc_thresh3.  | int<br/><small>Optional</small>    |
| netIpv4TcpFinTimeout           | Sysctl setting net.ipv4.tcp_fin_timeout.           | int<br/><small>Optional</small>    |
| netIpv4TcpkeepaliveIntvl       | Sysctl setting net.ipv4.tcp_keepalive_intvl.       | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveProbes      | Sysctl setting net.ipv4.tcp_keepalive_probes.      | int<br/><small>Optional</small>    |
| netIpv4TcpKeepaliveTime        | Sysctl setting net.ipv4.tcp_keepalive_time.        | int<br/><small>Optional</small>    |
| netIpv4TcpMaxSynBacklog        | Sysctl setting net.ipv4.tcp_max_syn_backlog.       | int<br/><small>Optional</small>    |
| netIpv4TcpMaxTwBuckets         | Sysctl setting net.ipv4.tcp_max_tw_buckets.        | int<br/><small>Optional</small>    |
| netIpv4TcpTwReuse              | Sysctl setting net.ipv4.tcp_tw_reuse.              | bool<br/><small>Optional</small>   |
| netNetfilterNfConntrackBuckets | Sysctl setting net.netfilter.nf_conntrack_buckets. | int<br/><small>Optional</small>    |
| netNetfilterNfConntrackMax     | Sysctl setting net.netfilter.nf_conntrack_max.     | int<br/><small>Optional</small>    |
| vmMaxMapCount                  | Sysctl setting vm.max_map_count.                   | int<br/><small>Optional</small>    |
| vmSwappiness                   | Sysctl setting vm.swappiness.                      | int<br/><small>Optional</small>    |
| vmVfsCachePressure             | Sysctl setting vm.vfs_cache_pressure.              | int<br/><small>Optional</small>    |

SystemData_CreatedByType_STATUS{#SystemData_CreatedByType_STATUS}
-----------------------------------------------------------------

Used by: [SystemData_STATUS](#SystemData_STATUS).

| Value             | Description |
|-------------------|-------------|
| "Application"     |             |
| "Key"             |             |
| "ManagedIdentity" |             |
| "User"            |             |

SystemData_LastModifiedByType_STATUS{#SystemData_LastModifiedByType_STATUS}
---------------------------------------------------------------------------

Used by: [SystemData_STATUS](#SystemData_STATUS).

| Value             | Description |
|-------------------|-------------|
| "Application"     |             |
| "Key"             |             |
| "ManagedIdentity" |             |
| "User"            |             |

UpgradeOverrideSettings{#UpgradeOverrideSettings}
-------------------------------------------------

Settings for overrides when upgrading a cluster.

Used by: [ClusterUpgradeSettings](#ClusterUpgradeSettings).

| Property     | Description                                                                                                                                                                                                                                                                                     | Type                               |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| forceUpgrade | Whether to force upgrade the cluster. Note that this option instructs upgrade operation to bypass upgrade protections such as checking for deprecated API usage. Enable this option only with caution.                                                                                          | bool<br/><small>Optional</small>   |
| until        | Until when the overrides are effective. Note that this only matches the start time of an upgrade, and the effectiveness won't change once an upgrade starts even if the `until` expires as upgrade proceeds. This field is not set by default. It must be set for the overrides to take effect. | string<br/><small>Optional</small> |

UpgradeOverrideSettings_STATUS{#UpgradeOverrideSettings_STATUS}
---------------------------------------------------------------

Settings for overrides when upgrading a cluster.

Used by: [ClusterUpgradeSettings_STATUS](#ClusterUpgradeSettings_STATUS).

| Property     | Description                                                                                                                                                                                                                                                                                     | Type                               |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| forceUpgrade | Whether to force upgrade the cluster. Note that this option instructs upgrade operation to bypass upgrade protections such as checking for deprecated API usage. Enable this option only with caution.                                                                                          | bool<br/><small>Optional</small>   |
| until        | Until when the overrides are effective. Note that this only matches the start time of an upgrade, and the effectiveness won't change once an upgrade starts even if the `until` expires as upgrade proceeds. This field is not set by default. It must be set for the overrides to take effect. | string<br/><small>Optional</small> |

UserAssignedIdentityDetails{#UserAssignedIdentityDetails}
---------------------------------------------------------

Information about the user assigned identity for the resource

Used by: [ManagedClusterIdentity](#ManagedClusterIdentity).

| Property  | Description | Type                                                                                                                                                       |
|-----------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| reference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

WindowsGmsaProfile{#WindowsGmsaProfile}
---------------------------------------

Windows gMSA Profile in the managed cluster.

Used by: [ManagedClusterWindowsProfile](#ManagedClusterWindowsProfile).

| Property       | Description                                                                                                                                                     | Type                               |
|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| dnsServer      | Specifies the DNS server for Windows gMSA. Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.       | string<br/><small>Optional</small> |
| enabled        | Specifies whether to enable Windows gMSA in the managed cluster.                                                                                                | bool<br/><small>Optional</small>   |
| rootDomainName | Specifies the root domain name for Windows gMSA. Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster. | string<br/><small>Optional</small> |

WindowsGmsaProfile_STATUS{#WindowsGmsaProfile_STATUS}
-----------------------------------------------------

Windows gMSA Profile in the managed cluster.

Used by: [ManagedClusterWindowsProfile_STATUS](#ManagedClusterWindowsProfile_STATUS).

| Property       | Description                                                                                                                                                     | Type                               |
|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| dnsServer      | Specifies the DNS server for Windows gMSA. Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.       | string<br/><small>Optional</small> |
| enabled        | Specifies whether to enable Windows gMSA in the managed cluster.                                                                                                | bool<br/><small>Optional</small>   |
| rootDomainName | Specifies the root domain name for Windows gMSA. Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster. | string<br/><small>Optional</small> |

AdvancedNetworkingObservability{#AdvancedNetworkingObservability}
-----------------------------------------------------------------

Observability profile to enable advanced network metrics and flow logs with historical contexts.

Used by: [AdvancedNetworking](#AdvancedNetworking).

| Property | Description                                                                                | Type                             |
|----------|--------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | Indicates the enablement of Advanced Networking observability functionalities on clusters. | bool<br/><small>Optional</small> |

AdvancedNetworkingObservability_STATUS{#AdvancedNetworkingObservability_STATUS}
-------------------------------------------------------------------------------

Observability profile to enable advanced network metrics and flow logs with historical contexts.

Used by: [AdvancedNetworking_STATUS](#AdvancedNetworking_STATUS).

| Property | Description                                                                                | Type                             |
|----------|--------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | Indicates the enablement of Advanced Networking observability functionalities on clusters. | bool<br/><small>Optional</small> |

AutoScaleProfile{#AutoScaleProfile}
-----------------------------------

Specifications on auto-scaling.

Used by: [ScaleProfile](#ScaleProfile).

| Property | Description                                                                                                                                                                                                                                                       | Type                                 |
|----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| maxCount | The maximum number of nodes of the specified sizes.                                                                                                                                                                                                               | int<br/><small>Optional</small>      |
| minCount | The minimum number of nodes of the specified sizes.                                                                                                                                                                                                               | int<br/><small>Optional</small>      |
| sizes    | The list of allowed vm sizes e.g. \[`Standard_E4s_v3`, `Standard_E16s_v3`, `Standard_D16s_v5`]. AKS will use the first available one when auto scaling. If a VM size is unavailable (e.g. due to quota or regional capacity reasons), AKS will use the next size. | string[]<br/><small>Optional</small> |

AutoScaleProfile_STATUS{#AutoScaleProfile_STATUS}
-------------------------------------------------

Specifications on auto-scaling.

Used by: [ScaleProfile_STATUS](#ScaleProfile_STATUS).

| Property | Description                                                                                                                                                                                                                                                       | Type                                 |
|----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| maxCount | The maximum number of nodes of the specified sizes.                                                                                                                                                                                                               | int<br/><small>Optional</small>      |
| minCount | The minimum number of nodes of the specified sizes.                                                                                                                                                                                                               | int<br/><small>Optional</small>      |
| sizes    | The list of allowed vm sizes e.g. \[`Standard_E4s_v3`, `Standard_E16s_v3`, `Standard_D16s_v5`]. AKS will use the first available one when auto scaling. If a VM size is unavailable (e.g. due to quota or regional capacity reasons), AKS will use the next size. | string[]<br/><small>Optional</small> |

AzureKeyVaultKms_KeyVaultNetworkAccess{#AzureKeyVaultKms_KeyVaultNetworkAccess}
-------------------------------------------------------------------------------

Used by: [AzureKeyVaultKms](#AzureKeyVaultKms).

| Value     | Description |
|-----------|-------------|
| "Private" |             |
| "Public"  |             |

AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS{#AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS}
---------------------------------------------------------------------------------------------

Used by: [AzureKeyVaultKms_STATUS](#AzureKeyVaultKms_STATUS).

| Value     | Description |
|-----------|-------------|
| "Private" |             |
| "Public"  |             |

ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig{#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig}
---------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig](#ContainerServiceNetworkProfile_KubeProxyConfig).

| Property             | Description                                                                                                      | Type                                                                                                                                                                    |
|----------------------|------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| scheduler            | IPVS scheduler, for more information please see http://www.linuxvirtualserver.org/docs/scheduling.html.          | [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler)<br/><small>Optional</small> |
| tcpFinTimeoutSeconds | The timeout value used for IPVS TCP sessions after receiving a FIN in seconds. Must be a positive integer value. | int<br/><small>Optional</small>                                                                                                                                         |
| tcpTimeoutSeconds    | The timeout value used for idle IPVS TCP sessions in seconds. Must be a positive integer value.                  | int<br/><small>Optional</small>                                                                                                                                         |
| udpTimeoutSeconds    | The timeout value used for IPVS UDP packets in seconds. Must be a positive integer value.                        | int<br/><small>Optional</small>                                                                                                                                         |

ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS{#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS}
-----------------------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_STATUS).

| Property             | Description                                                                                                      | Type                                                                                                                                                                                  |
|----------------------|------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| scheduler            | IPVS scheduler, for more information please see http://www.linuxvirtualserver.org/docs/scheduling.html.          | [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS)<br/><small>Optional</small> |
| tcpFinTimeoutSeconds | The timeout value used for IPVS TCP sessions after receiving a FIN in seconds. Must be a positive integer value. | int<br/><small>Optional</small>                                                                                                                                                       |
| tcpTimeoutSeconds    | The timeout value used for idle IPVS TCP sessions in seconds. Must be a positive integer value.                  | int<br/><small>Optional</small>                                                                                                                                                       |
| udpTimeoutSeconds    | The timeout value used for IPVS UDP packets in seconds. Must be a positive integer value.                        | int<br/><small>Optional</small>                                                                                                                                                       |

ContainerServiceNetworkProfile_KubeProxyConfig_Mode{#ContainerServiceNetworkProfile_KubeProxyConfig_Mode}
---------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig](#ContainerServiceNetworkProfile_KubeProxyConfig).

| Value      | Description |
|------------|-------------|
| "IPTABLES" |             |
| "IPVS"     |             |

ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS{#ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS}
-----------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_STATUS).

| Value      | Description |
|------------|-------------|
| "IPTABLES" |             |
| "IPVS"     |             |

ContainerServiceSshPublicKey{#ContainerServiceSshPublicKey}
-----------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration](#ContainerServiceSshConfiguration).

| Property | Description                                                                                                                      | Type                               |
|----------|----------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string<br/><small>Required</small> |

ContainerServiceSshPublicKey_STATUS{#ContainerServiceSshPublicKey_STATUS}
-------------------------------------------------------------------------

Contains information about SSH certificate public key data.

Used by: [ContainerServiceSshConfiguration_STATUS](#ContainerServiceSshConfiguration_STATUS).

| Property | Description                                                                                                                      | Type                               |
|----------|----------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| keyData  | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | string<br/><small>Optional</small> |

IstioCertificateAuthority{#IstioCertificateAuthority}
-----------------------------------------------------

Istio Service Mesh Certificate Authority (CA) configuration. For now, we only support plugin certificates as described here https://aka.ms/asm-plugin-ca

Used by: [IstioServiceMesh](#IstioServiceMesh).

| Property | Description                                       | Type                                                                                            |
|----------|---------------------------------------------------|-------------------------------------------------------------------------------------------------|
| plugin   | Plugin certificates information for Service Mesh. | [IstioPluginCertificateAuthority](#IstioPluginCertificateAuthority)<br/><small>Optional</small> |

IstioCertificateAuthority_STATUS{#IstioCertificateAuthority_STATUS}
-------------------------------------------------------------------

Istio Service Mesh Certificate Authority (CA) configuration. For now, we only support plugin certificates as described here https://aka.ms/asm-plugin-ca

Used by: [IstioServiceMesh_STATUS](#IstioServiceMesh_STATUS).

| Property | Description                                       | Type                                                                                                          |
|----------|---------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| plugin   | Plugin certificates information for Service Mesh. | [IstioPluginCertificateAuthority_STATUS](#IstioPluginCertificateAuthority_STATUS)<br/><small>Optional</small> |

IstioComponents{#IstioComponents}
---------------------------------

Istio components configuration.

Used by: [IstioServiceMesh](#IstioServiceMesh).

| Property        | Description             | Type                                                                      |
|-----------------|-------------------------|---------------------------------------------------------------------------|
| egressGateways  | Istio egress gateways.  | [IstioEgressGateway[]](#IstioEgressGateway)<br/><small>Optional</small>   |
| ingressGateways | Istio ingress gateways. | [IstioIngressGateway[]](#IstioIngressGateway)<br/><small>Optional</small> |

IstioComponents_STATUS{#IstioComponents_STATUS}
-----------------------------------------------

Istio components configuration.

Used by: [IstioServiceMesh_STATUS](#IstioServiceMesh_STATUS).

| Property        | Description             | Type                                                                                    |
|-----------------|-------------------------|-----------------------------------------------------------------------------------------|
| egressGateways  | Istio egress gateways.  | [IstioEgressGateway_STATUS[]](#IstioEgressGateway_STATUS)<br/><small>Optional</small>   |
| ingressGateways | Istio ingress gateways. | [IstioIngressGateway_STATUS[]](#IstioIngressGateway_STATUS)<br/><small>Optional</small> |

ManagedClusterAzureMonitorProfileAppMonitoringAutoInstrumentation{#ManagedClusterAzureMonitorProfileAppMonitoringAutoInstrumentation}
-------------------------------------------------------------------------------------------------------------------------------------

Application Monitoring Auto Instrumentation for Kubernetes Application Container. Deploys web hook to auto-instrument Azure Monitor OpenTelemetry based SDKs to collect OpenTelemetry metrics, logs and traces of the application. See aka.ms/AzureMonitorApplicationMonitoring for an overview.

Used by: [ManagedClusterAzureMonitorProfileAppMonitoring](#ManagedClusterAzureMonitorProfileAppMonitoring).

| Property | Description                                                                 | Type                             |
|----------|-----------------------------------------------------------------------------|----------------------------------|
| enabled  | Indicates if Application Monitoring Auto Instrumentation is enabled or not. | bool<br/><small>Optional</small> |

ManagedClusterAzureMonitorProfileAppMonitoringAutoInstrumentation_STATUS{#ManagedClusterAzureMonitorProfileAppMonitoringAutoInstrumentation_STATUS}
---------------------------------------------------------------------------------------------------------------------------------------------------

Application Monitoring Auto Instrumentation for Kubernetes Application Container. Deploys web hook to auto-instrument Azure Monitor OpenTelemetry based SDKs to collect OpenTelemetry metrics, logs and traces of the application. See aka.ms/AzureMonitorApplicationMonitoring for an overview.

Used by: [ManagedClusterAzureMonitorProfileAppMonitoring_STATUS](#ManagedClusterAzureMonitorProfileAppMonitoring_STATUS).

| Property | Description                                                                 | Type                             |
|----------|-----------------------------------------------------------------------------|----------------------------------|
| enabled  | Indicates if Application Monitoring Auto Instrumentation is enabled or not. | bool<br/><small>Optional</small> |

ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogs{#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogs}
---------------------------------------------------------------------------------------------------------------------------------

Application Monitoring Open Telemetry Metrics Profile for Kubernetes Application Container Logs and Traces. Collects OpenTelemetry logs and traces of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.

Used by: [ManagedClusterAzureMonitorProfileAppMonitoring](#ManagedClusterAzureMonitorProfileAppMonitoring).

| Property | Description                                                                                                   | Type                             |
|----------|---------------------------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | Indicates if Application Monitoring Open Telemetry Logs and traces is enabled or not.                         | bool<br/><small>Optional</small> |
| port     | The Open Telemetry host port for Open Telemetry logs and traces. If not specified, the default port is 28331. | int<br/><small>Optional</small>  |

ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogs_STATUS{#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogs_STATUS}
-----------------------------------------------------------------------------------------------------------------------------------------------

Application Monitoring Open Telemetry Metrics Profile for Kubernetes Application Container Logs and Traces. Collects OpenTelemetry logs and traces of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.

Used by: [ManagedClusterAzureMonitorProfileAppMonitoring_STATUS](#ManagedClusterAzureMonitorProfileAppMonitoring_STATUS).

| Property | Description                                                                                                   | Type                             |
|----------|---------------------------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | Indicates if Application Monitoring Open Telemetry Logs and traces is enabled or not.                         | bool<br/><small>Optional</small> |
| port     | The Open Telemetry host port for Open Telemetry logs and traces. If not specified, the default port is 28331. | int<br/><small>Optional</small>  |

ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics{#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics}
---------------------------------------------------------------------------------------------------------------------------------------

Application Monitoring Open Telemetry Metrics Profile for Kubernetes Application Container Metrics. Collects OpenTelemetry metrics of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.

Used by: [ManagedClusterAzureMonitorProfileAppMonitoring](#ManagedClusterAzureMonitorProfileAppMonitoring).

| Property | Description                                                                                           | Type                             |
|----------|-------------------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | Indicates if Application Monitoring Open Telemetry Metrics is enabled or not.                         | bool<br/><small>Optional</small> |
| port     | The Open Telemetry host port for Open Telemetry metrics. If not specified, the default port is 28333. | int<br/><small>Optional</small>  |

ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics_STATUS{#ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics_STATUS}
-----------------------------------------------------------------------------------------------------------------------------------------------------

Application Monitoring Open Telemetry Metrics Profile for Kubernetes Application Container Metrics. Collects OpenTelemetry metrics of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.

Used by: [ManagedClusterAzureMonitorProfileAppMonitoring_STATUS](#ManagedClusterAzureMonitorProfileAppMonitoring_STATUS).

| Property | Description                                                                                           | Type                             |
|----------|-------------------------------------------------------------------------------------------------------|----------------------------------|
| enabled  | Indicates if Application Monitoring Open Telemetry Metrics is enabled or not.                         | bool<br/><small>Optional</small> |
| port     | The Open Telemetry host port for Open Telemetry metrics. If not specified, the default port is 28333. | int<br/><small>Optional</small>  |

ManagedClusterAzureMonitorProfileKubeStateMetrics{#ManagedClusterAzureMonitorProfileKubeStateMetrics}
-----------------------------------------------------------------------------------------------------

Kube State Metrics for prometheus addon profile for the container service cluster

Used by: [ManagedClusterAzureMonitorProfileMetrics](#ManagedClusterAzureMonitorProfileMetrics).

| Property                   | Description                                                                                                 | Type                               |
|----------------------------|-------------------------------------------------------------------------------------------------------------|------------------------------------|
| metricAnnotationsAllowList | Comma-separated list of additional Kubernetes label keys that will be used in the resource's labels metric. | string<br/><small>Optional</small> |
| metricLabelsAllowlist      | Comma-separated list of Kubernetes annotations keys that will be used in the resource's labels metric.      | string<br/><small>Optional</small> |

ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS{#ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS}
-------------------------------------------------------------------------------------------------------------------

Kube State Metrics for prometheus addon profile for the container service cluster

Used by: [ManagedClusterAzureMonitorProfileMetrics_STATUS](#ManagedClusterAzureMonitorProfileMetrics_STATUS).

| Property                   | Description                                                                                                 | Type                               |
|----------------------------|-------------------------------------------------------------------------------------------------------------|------------------------------------|
| metricAnnotationsAllowList | Comma-separated list of additional Kubernetes label keys that will be used in the resource's labels metric. | string<br/><small>Optional</small> |
| metricLabelsAllowlist      | Comma-separated list of Kubernetes annotations keys that will be used in the resource's labels metric.      | string<br/><small>Optional</small> |

ManagedClusterLoadBalancerProfile_BackendPoolType{#ManagedClusterLoadBalancerProfile_BackendPoolType}
-----------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Value                 | Description |
|-----------------------|-------------|
| "NodeIP"              |             |
| "NodeIPConfiguration" |             |

ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS{#ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS}
-------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Value                 | Description |
|-----------------------|-------------|
| "NodeIP"              |             |
| "NodeIPConfiguration" |             |

ManagedClusterLoadBalancerProfile_ClusterServiceLoadBalancerHealthProbeMode{#ManagedClusterLoadBalancerProfile_ClusterServiceLoadBalancerHealthProbeMode}
---------------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Value             | Description |
|-------------------|-------------|
| "ServiceNodePort" |             |
| "Shared"          |             |

ManagedClusterLoadBalancerProfile_ClusterServiceLoadBalancerHealthProbeMode_STATUS{#ManagedClusterLoadBalancerProfile_ClusterServiceLoadBalancerHealthProbeMode_STATUS}
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Value             | Description |
|-------------------|-------------|
| "ServiceNodePort" |             |
| "Shared"          |             |

ManagedClusterLoadBalancerProfile_ManagedOutboundIPs{#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs}
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property  | Description                                                                                                                                                                                                                | Type                            |
|-----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| count     | The desired number of IPv4 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.                                       | int<br/><small>Optional</small> |
| countIPv6 | The desired number of IPv6 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 0 for single-stack and 1 for dual-stack. | int<br/><small>Optional</small> |

ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS{#ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS}
-------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property  | Description                                                                                                                                                                                                                | Type                            |
|-----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| count     | The desired number of IPv4 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.                                       | int<br/><small>Optional</small> |
| countIPv6 | The desired number of IPv6 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 0 for single-stack and 1 for dual-stack. | int<br/><small>Optional</small> |

ManagedClusterLoadBalancerProfile_OutboundIPPrefixes{#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes}
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property         | Description                           | Type                                                                  |
|------------------|---------------------------------------|-----------------------------------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small> |

ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS{#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS}
-------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property         | Description                           | Type                                                                                |
|------------------|---------------------------------------|-------------------------------------------------------------------------------------|
| publicIPPrefixes | A list of public IP prefix resources. | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small> |

ManagedClusterLoadBalancerProfile_OutboundIPs{#ManagedClusterLoadBalancerProfile_OutboundIPs}
---------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile).

| Property  | Description                    | Type                                                                  |
|-----------|--------------------------------|-----------------------------------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference[]](#ResourceReference)<br/><small>Optional</small> |

ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS{#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS}
-----------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS).

| Property  | Description                    | Type                                                                                |
|-----------|--------------------------------|-------------------------------------------------------------------------------------|
| publicIPs | A list of public IP resources. | [ResourceReference_STATUS[]](#ResourceReference_STATUS)<br/><small>Optional</small> |

ManagedClusterManagedOutboundIPProfile{#ManagedClusterManagedOutboundIPProfile}
-------------------------------------------------------------------------------

Profile of the managed outbound IP resources of the managed cluster.

Used by: [ManagedClusterNATGatewayProfile](#ManagedClusterNATGatewayProfile).

| Property | Description                                                                                                                                      | Type                            |
|----------|--------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| count    | The desired number of outbound IPs created/managed by Azure. Allowed values must be in the range of 1 to 16 (inclusive). The default value is 1. | int<br/><small>Optional</small> |

ManagedClusterManagedOutboundIPProfile_STATUS{#ManagedClusterManagedOutboundIPProfile_STATUS}
---------------------------------------------------------------------------------------------

Profile of the managed outbound IP resources of the managed cluster.

Used by: [ManagedClusterNATGatewayProfile_STATUS](#ManagedClusterNATGatewayProfile_STATUS).

| Property | Description                                                                                                                                      | Type                            |
|----------|--------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| count    | The desired number of outbound IPs created/managed by Azure. Allowed values must be in the range of 1 to 16 (inclusive). The default value is 1. | int<br/><small>Optional</small> |

ManagedClusterPodIdentity_ProvisioningInfo_STATUS{#ManagedClusterPodIdentity_ProvisioningInfo_STATUS}
-----------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Property | Description                             | Type                                                                                                                                |
|----------|-----------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| error    | Pod identity assignment error (if any). | [ManagedClusterPodIdentityProvisioningError_STATUS](#ManagedClusterPodIdentityProvisioningError_STATUS)<br/><small>Optional</small> |

ManagedClusterPodIdentity_ProvisioningState_STATUS{#ManagedClusterPodIdentity_ProvisioningState_STATUS}
-------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentity_STATUS](#ManagedClusterPodIdentity_STATUS).

| Value       | Description |
|-------------|-------------|
| "Assigned"  |             |
| "Canceled"  |             |
| "Deleting"  |             |
| "Failed"    |             |
| "Succeeded" |             |
| "Updating"  |             |

ManagedClusterSecurityProfileDefenderSecurityMonitoring{#ManagedClusterSecurityProfileDefenderSecurityMonitoring}
-----------------------------------------------------------------------------------------------------------------

Microsoft Defender settings for the security profile threat detection.

Used by: [ManagedClusterSecurityProfileDefender](#ManagedClusterSecurityProfileDefender).

| Property | Description                                 | Type                             |
|----------|---------------------------------------------|----------------------------------|
| enabled  | Whether to enable Defender threat detection | bool<br/><small>Optional</small> |

ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS{#ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS}
-------------------------------------------------------------------------------------------------------------------------------

Microsoft Defender settings for the security profile threat detection.

Used by: [ManagedClusterSecurityProfileDefender_STATUS](#ManagedClusterSecurityProfileDefender_STATUS).

| Property | Description                                 | Type                             |
|----------|---------------------------------------------|----------------------------------|
| enabled  | Whether to enable Defender threat detection | bool<br/><small>Optional</small> |

ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling{#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling}
-------------------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Enabled"  |             |

ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling_STATUS{#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_AddonAutoscaling_STATUS}
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS](#ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Enabled"  |             |

ManualScaleProfile{#ManualScaleProfile}
---------------------------------------

Specifications on number of machines.

Used by: [ScaleProfile](#ScaleProfile).

| Property | Description                                                                                                                                                                                                                                                  | Type                                 |
|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| count    | Number of nodes.                                                                                                                                                                                                                                             | int<br/><small>Optional</small>      |
| sizes    | The list of allowed vm sizes e.g. \[`Standard_E4s_v3`, `Standard_E16s_v3`, `Standard_D16s_v5`]. AKS will use the first available one when scaling. If a VM size is unavailable (e.g. due to quota or regional capacity reasons), AKS will use the next size. | string[]<br/><small>Optional</small> |

ManualScaleProfile_STATUS{#ManualScaleProfile_STATUS}
-----------------------------------------------------

Specifications on number of machines.

Used by: [ScaleProfile_STATUS](#ScaleProfile_STATUS).

| Property | Description                                                                                                                                                                                                                                                  | Type                                 |
|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| count    | Number of nodes.                                                                                                                                                                                                                                             | int<br/><small>Optional</small>      |
| sizes    | The list of allowed vm sizes e.g. \[`Standard_E4s_v3`, `Standard_E16s_v3`, `Standard_D16s_v5`]. AKS will use the first available one when scaling. If a VM size is unavailable (e.g. due to quota or regional capacity reasons), AKS will use the next size. | string[]<br/><small>Optional</small> |

PortRange_Protocol{#PortRange_Protocol}
---------------------------------------

Used by: [PortRange](#PortRange).

| Value | Description |
|-------|-------------|
| "TCP" |             |
| "UDP" |             |

PortRange_Protocol_STATUS{#PortRange_Protocol_STATUS}
-----------------------------------------------------

Used by: [PortRange_STATUS](#PortRange_STATUS).

| Value | Description |
|-------|-------------|
| "TCP" |             |
| "UDP" |             |

ResourceReference{#ResourceReference}
-------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile](#ManagedClusterLoadBalancerProfile), [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes), [ManagedClusterLoadBalancerProfile_OutboundIPs](#ManagedClusterLoadBalancerProfile_OutboundIPs), and [ManagedClusterNATGatewayProfile](#ManagedClusterNATGatewayProfile).

| Property  | Description                            | Type                                                                                                                                                       |
|-----------|----------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| reference | The fully qualified Azure resource id. | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

ResourceReference_STATUS{#ResourceReference_STATUS}
---------------------------------------------------

A reference to an Azure resource.

Used by: [ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS), [ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS](#ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS), [ManagedClusterLoadBalancerProfile_STATUS](#ManagedClusterLoadBalancerProfile_STATUS), and [ManagedClusterNATGatewayProfile_STATUS](#ManagedClusterNATGatewayProfile_STATUS).

| Property | Description                            | Type                               |
|----------|----------------------------------------|------------------------------------|
| id       | The fully qualified Azure resource id. | string<br/><small>Optional</small> |

ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler{#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler}
-----------------------------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig).

| Value             | Description |
|-------------------|-------------|
| "LeastConnection" |             |
| "RoundRobin"      |             |

ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS{#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS}
-------------------------------------------------------------------------------------------------------------------------------------------------------

Used by: [ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS](#ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS).

| Value             | Description |
|-------------------|-------------|
| "LeastConnection" |             |
| "RoundRobin"      |             |

IstioEgressGateway{#IstioEgressGateway}
---------------------------------------

Istio egress gateway configuration.

Used by: [IstioComponents](#IstioComponents).

| Property | Description                           | Type                             |
|----------|---------------------------------------|----------------------------------|
| enabled  | Whether to enable the egress gateway. | bool<br/><small>Required</small> |

IstioEgressGateway_STATUS{#IstioEgressGateway_STATUS}
-----------------------------------------------------

Istio egress gateway configuration.

Used by: [IstioComponents_STATUS](#IstioComponents_STATUS).

| Property | Description                           | Type                             |
|----------|---------------------------------------|----------------------------------|
| enabled  | Whether to enable the egress gateway. | bool<br/><small>Optional</small> |

IstioIngressGateway{#IstioIngressGateway}
-----------------------------------------

Istio ingress gateway configuration. For now, we support up to one external ingress gateway named `aks-istio-ingressgateway-external` and one internal ingress gateway named `aks-istio-ingressgateway-internal`.

Used by: [IstioComponents](#IstioComponents).

| Property | Description                            | Type                                                                              |
|----------|----------------------------------------|-----------------------------------------------------------------------------------|
| enabled  | Whether to enable the ingress gateway. | bool<br/><small>Required</small>                                                  |
| mode     | Mode of an ingress gateway.            | [IstioIngressGateway_Mode](#IstioIngressGateway_Mode)<br/><small>Required</small> |

IstioIngressGateway_STATUS{#IstioIngressGateway_STATUS}
-------------------------------------------------------

Istio ingress gateway configuration. For now, we support up to one external ingress gateway named `aks-istio-ingressgateway-external` and one internal ingress gateway named `aks-istio-ingressgateway-internal`.

Used by: [IstioComponents_STATUS](#IstioComponents_STATUS).

| Property | Description                            | Type                                                                                            |
|----------|----------------------------------------|-------------------------------------------------------------------------------------------------|
| enabled  | Whether to enable the ingress gateway. | bool<br/><small>Optional</small>                                                                |
| mode     | Mode of an ingress gateway.            | [IstioIngressGateway_Mode_STATUS](#IstioIngressGateway_Mode_STATUS)<br/><small>Optional</small> |

IstioPluginCertificateAuthority{#IstioPluginCertificateAuthority}
-----------------------------------------------------------------

Plugin certificates information for Service Mesh.

Used by: [IstioCertificateAuthority](#IstioCertificateAuthority).

| Property            | Description                                                          | Type                                                                                                                                                       |
|---------------------|----------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| certChainObjectName | Certificate chain object name in Azure Key Vault.                    | string<br/><small>Optional</small>                                                                                                                         |
| certObjectName      | Intermediate certificate object name in Azure Key Vault.             | string<br/><small>Optional</small>                                                                                                                         |
| keyObjectName       | Intermediate certificate private key object name in Azure Key Vault. | string<br/><small>Optional</small>                                                                                                                         |
| keyVaultReference   | The resource ID of the Key Vault.                                    | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| rootCertObjectName  | Root certificate object name in Azure Key Vault.                     | string<br/><small>Optional</small>                                                                                                                         |

IstioPluginCertificateAuthority_STATUS{#IstioPluginCertificateAuthority_STATUS}
-------------------------------------------------------------------------------

Plugin certificates information for Service Mesh.

Used by: [IstioCertificateAuthority_STATUS](#IstioCertificateAuthority_STATUS).

| Property            | Description                                                          | Type                               |
|---------------------|----------------------------------------------------------------------|------------------------------------|
| certChainObjectName | Certificate chain object name in Azure Key Vault.                    | string<br/><small>Optional</small> |
| certObjectName      | Intermediate certificate object name in Azure Key Vault.             | string<br/><small>Optional</small> |
| keyObjectName       | Intermediate certificate private key object name in Azure Key Vault. | string<br/><small>Optional</small> |
| keyVaultId          | The resource ID of the Key Vault.                                    | string<br/><small>Optional</small> |
| rootCertObjectName  | Root certificate object name in Azure Key Vault.                     | string<br/><small>Optional</small> |

ManagedClusterPodIdentityProvisioningError_STATUS{#ManagedClusterPodIdentityProvisioningError_STATUS}
-----------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentity_ProvisioningInfo_STATUS](#ManagedClusterPodIdentity_ProvisioningInfo_STATUS).

| Property | Description              | Type                                                                                                                                        |
|----------|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| error    | Details about the error. | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS)<br/><small>Optional</small> |

IstioIngressGateway_Mode{#IstioIngressGateway_Mode}
---------------------------------------------------

Used by: [IstioIngressGateway](#IstioIngressGateway).

| Value      | Description |
|------------|-------------|
| "External" |             |
| "Internal" |             |

IstioIngressGateway_Mode_STATUS{#IstioIngressGateway_Mode_STATUS}
-----------------------------------------------------------------

Used by: [IstioIngressGateway_STATUS](#IstioIngressGateway_STATUS).

| Value      | Description |
|------------|-------------|
| "External" |             |
| "Internal" |             |

ManagedClusterPodIdentityProvisioningErrorBody_STATUS{#ManagedClusterPodIdentityProvisioningErrorBody_STATUS}
-------------------------------------------------------------------------------------------------------------

An error response from the pod identity provisioning.

Used by: [ManagedClusterPodIdentityProvisioningError_STATUS](#ManagedClusterPodIdentityProvisioningError_STATUS).

| Property | Description                                                                                        | Type                                                                                                                                                            |
|----------|----------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string<br/><small>Optional</small>                                                                                                                              |
| details  | A list of additional details about the error.                                                      | [ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled[]](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled)<br/><small>Optional</small> |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string<br/><small>Optional</small>                                                                                                                              |
| target   | The target of the particular error. For example, the name of the property in error.                | string<br/><small>Optional</small>                                                                                                                              |

ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled{#ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled}
-------------------------------------------------------------------------------------------------------------------------------

Used by: [ManagedClusterPodIdentityProvisioningErrorBody_STATUS](#ManagedClusterPodIdentityProvisioningErrorBody_STATUS).

| Property | Description                                                                                        | Type                               |
|----------|----------------------------------------------------------------------------------------------------|------------------------------------|
| code     | An identifier for the error. Codes are invariant and are intended to be consumed programmatically. | string<br/><small>Optional</small> |
| message  | A message describing the error, intended to be suitable for display in a user interface.           | string<br/><small>Optional</small> |
| target   | The target of the particular error. For example, the name of the property in error.                | string<br/><small>Optional</small> |
