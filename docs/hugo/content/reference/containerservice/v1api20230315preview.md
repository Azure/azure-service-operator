---

title: containerservice.azure.com/v1api20230315preview

linktitle: v1api20230315preview
-------------------------------

APIVersion{#APIVersion}
-----------------------

| Value                | Description |
|----------------------|-------------|
| "2023-03-15-preview" |             |

Fleet{#Fleet}
-------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;fleets/&ZeroWidthSpace;{fleetName}

Used by: [FleetList](#FleetList).

| Property                                                                                | Description | Type                                                      |
|-----------------------------------------------------------------------------------------|-------------|-----------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                           |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                           |
| spec                                                                                    |             | [Fleet_Spec](#Fleet_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [Fleet_STATUS](#Fleet_STATUS)<br/><small>Optional</small> |

### Fleet_Spec {#Fleet_Spec}

| Property     | Description                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                 |
|--------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| hubProfile   | The FleetHubProfile configures the Fleet's hub.                                                                                                                                                                                                                                              | [FleetHubProfile](#FleetHubProfile)<br/><small>Optional</small>                                                                                                      |
| location     | The geo-location where the resource lives                                                                                                                                                                                                                                                    | string<br/><small>Required</small>                                                                                                                                   |
| operatorSpec | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                              | [FleetOperatorSpec](#FleetOperatorSpec)<br/><small>Optional</small>                                                                                                  |
| owner        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| tags         | Resource tags.                                                                                                                                                                                                                                                                               | map[string]string<br/><small>Optional</small>                                                                                                                        |

### Fleet_STATUS{#Fleet_STATUS}

| Property          | Description                                                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                    |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions        | The observed state of the resource                                                                                                                                                                                                                                                                                                                                  | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| eTag              | If eTag is provided in the response body, it may also be provided as a header per the normal etag convention. Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. | string<br/><small>Optional</small>                                                                                                                      |
| hubProfile        | The FleetHubProfile configures the Fleet's hub.                                                                                                                                                                                                                                                                                                                     | [FleetHubProfile_STATUS](#FleetHubProfile_STATUS)<br/><small>Optional</small>                                                                           |
| id                | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}                                           | string<br/><small>Optional</small>                                                                                                                      |
| location          | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| name              | The name of the resource                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState | The status of the last operation.                                                                                                                                                                                                                                                                                                                                   | [FleetProvisioningState_STATUS](#FleetProvisioningState_STATUS)<br/><small>Optional</small>                                                             |
| systemData        | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                                                                    | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags              | Resource tags.                                                                                                                                                                                                                                                                                                                                                      | map[string]string<br/><small>Optional</small>                                                                                                           |
| type              | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |

FleetList{#FleetList}
---------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;fleets/&ZeroWidthSpace;{fleetName}

| Property                                                                            | Description | Type                                          |
|-------------------------------------------------------------------------------------|-------------|-----------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                               |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                               |
| items                                                                               |             | [Fleet[]](#Fleet)<br/><small>Optional</small> |

FleetsMember{#FleetsMember}
---------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;fleets/&ZeroWidthSpace;{fleetName}/&ZeroWidthSpace;members/&ZeroWidthSpace;{fleetMemberName}

Used by: [FleetsMemberList](#FleetsMemberList).

| Property                                                                                | Description | Type                                                                    |
|-----------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                         |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                         |
| spec                                                                                    |             | [FleetsMember_Spec](#FleetsMember_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [FleetsMember_STATUS](#FleetsMember_STATUS)<br/><small>Optional</small> |

### FleetsMember_Spec {#FleetsMember_Spec}

| Property                 | Description                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                 |
|--------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName                | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                                   |
| clusterResourceReference | The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{clusterName}'. | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Required</small>           |
| group                    | The group this member belongs to for multi-cluster update management.                                                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec             | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                            | [FleetsMemberOperatorSpec](#FleetsMemberOperatorSpec)<br/><small>Optional</small>                                                                                    |
| owner                    | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/Fleet resource                                                                                | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |

### FleetsMember_STATUS{#FleetsMember_STATUS}

| Property          | Description                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                    |
|-------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| clusterResourceId | The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{clusterName}'. | string<br/><small>Optional</small>                                                                                                                      |
| conditions        | The observed state of the resource                                                                                                                                                                                                                                                                                                                                         | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| eTag              | If eTag is provided in the response body, it may also be provided as a header per the normal etag convention. Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.        | string<br/><small>Optional</small>                                                                                                                      |
| group             | The group this member belongs to for multi-cluster update management.                                                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                      |
| id                | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}                                                  | string<br/><small>Optional</small>                                                                                                                      |
| name              | The name of the resource                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState | The status of the last operation.                                                                                                                                                                                                                                                                                                                                          | [FleetMemberProvisioningState_STATUS](#FleetMemberProvisioningState_STATUS)<br/><small>Optional</small>                                                 |
| systemData        | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                                                                           | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type              | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |

FleetsMemberList{#FleetsMemberList}
-----------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;fleets/&ZeroWidthSpace;{fleetName}/&ZeroWidthSpace;members/&ZeroWidthSpace;{fleetMemberName}

| Property                                                                            | Description | Type                                                        |
|-------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                             |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                             |
| items                                                                               |             | [FleetsMember[]](#FleetsMember)<br/><small>Optional</small> |

FleetsUpdateRun{#FleetsUpdateRun}
---------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;fleets/&ZeroWidthSpace;{fleetName}/&ZeroWidthSpace;updateRuns/&ZeroWidthSpace;{updateRunName}

Used by: [FleetsUpdateRunList](#FleetsUpdateRunList).

| Property                                                                                | Description | Type                                                                          |
|-----------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                               |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                               |
| spec                                                                                    |             | [FleetsUpdateRun_Spec](#FleetsUpdateRun_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [FleetsUpdateRun_STATUS](#FleetsUpdateRun_STATUS)<br/><small>Optional</small> |

### FleetsUpdateRun_Spec {#FleetsUpdateRun_Spec}

| Property             | Description                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                 |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName            | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                                   |
| managedClusterUpdate | The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be modified until the run is started.                                                                                                                                                                           | [ManagedClusterUpdate](#ManagedClusterUpdate)<br/><small>Required</small>                                                                                            |
| operatorSpec         | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                         | [FleetsUpdateRunOperatorSpec](#FleetsUpdateRunOperatorSpec)<br/><small>Optional</small>                                                                              |
| owner                | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/Fleet resource             | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| strategy             | The strategy defines the order in which the clusters will be updated. If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single UpdateGroup targeting all members. The strategy of the UpdateRun can be modified until the run is started. | [UpdateRunStrategy](#UpdateRunStrategy)<br/><small>Optional</small>                                                                                                  |

### FleetsUpdateRun_STATUS{#FleetsUpdateRun_STATUS}

| Property             | Description                                                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                    |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions           | The observed state of the resource                                                                                                                                                                                                                                                                                                                                  | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| eTag                 | If eTag is provided in the response body, it may also be provided as a header per the normal etag convention. Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. | string<br/><small>Optional</small>                                                                                                                      |
| id                   | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}                                           | string<br/><small>Optional</small>                                                                                                                      |
| managedClusterUpdate | The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be modified until the run is started.                                                                                                                                                                                                                                       | [ManagedClusterUpdate_STATUS](#ManagedClusterUpdate_STATUS)<br/><small>Optional</small>                                                                 |
| name                 | The name of the resource                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState    | The provisioning state of the UpdateRun resource.                                                                                                                                                                                                                                                                                                                   | [UpdateRunProvisioningState_STATUS](#UpdateRunProvisioningState_STATUS)<br/><small>Optional</small>                                                     |
| status               | The status of the UpdateRun.                                                                                                                                                                                                                                                                                                                                        | [UpdateRunStatus_STATUS](#UpdateRunStatus_STATUS)<br/><small>Optional</small>                                                                           |
| strategy             | The strategy defines the order in which the clusters will be updated. If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single UpdateGroup targeting all members. The strategy of the UpdateRun can be modified until the run is started.                                                             | [UpdateRunStrategy_STATUS](#UpdateRunStrategy_STATUS)<br/><small>Optional</small>                                                                       |
| systemData           | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                                                                    | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type                 | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |

FleetsUpdateRunList{#FleetsUpdateRunList}
-----------------------------------------

Generator information: - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;fleets/&ZeroWidthSpace;{fleetName}/&ZeroWidthSpace;updateRuns/&ZeroWidthSpace;{updateRunName}

| Property                                                                            | Description | Type                                                              |
|-------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                   |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                   |
| items                                                                               |             | [FleetsUpdateRun[]](#FleetsUpdateRun)<br/><small>Optional</small> |

Fleet_Spec{#Fleet_Spec}
-----------------------

Used by: [Fleet](#Fleet).

| Property     | Description                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                 |
|--------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| hubProfile   | The FleetHubProfile configures the Fleet's hub.                                                                                                                                                                                                                                              | [FleetHubProfile](#FleetHubProfile)<br/><small>Optional</small>                                                                                                      |
| location     | The geo-location where the resource lives                                                                                                                                                                                                                                                    | string<br/><small>Required</small>                                                                                                                                   |
| operatorSpec | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                              | [FleetOperatorSpec](#FleetOperatorSpec)<br/><small>Optional</small>                                                                                                  |
| owner        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| tags         | Resource tags.                                                                                                                                                                                                                                                                               | map[string]string<br/><small>Optional</small>                                                                                                                        |

Fleet_STATUS{#Fleet_STATUS}
---------------------------

The Fleet resource.

Used by: [Fleet](#Fleet).

| Property          | Description                                                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                    |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions        | The observed state of the resource                                                                                                                                                                                                                                                                                                                                  | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| eTag              | If eTag is provided in the response body, it may also be provided as a header per the normal etag convention. Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. | string<br/><small>Optional</small>                                                                                                                      |
| hubProfile        | The FleetHubProfile configures the Fleet's hub.                                                                                                                                                                                                                                                                                                                     | [FleetHubProfile_STATUS](#FleetHubProfile_STATUS)<br/><small>Optional</small>                                                                           |
| id                | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}                                           | string<br/><small>Optional</small>                                                                                                                      |
| location          | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| name              | The name of the resource                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState | The status of the last operation.                                                                                                                                                                                                                                                                                                                                   | [FleetProvisioningState_STATUS](#FleetProvisioningState_STATUS)<br/><small>Optional</small>                                                             |
| systemData        | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                                                                    | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags              | Resource tags.                                                                                                                                                                                                                                                                                                                                                      | map[string]string<br/><small>Optional</small>                                                                                                           |
| type              | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |

FleetsMember_Spec{#FleetsMember_Spec}
-------------------------------------

Used by: [FleetsMember](#FleetsMember).

| Property                 | Description                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                 |
|--------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName                | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                                   |
| clusterResourceReference | The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{clusterName}'. | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Required</small>           |
| group                    | The group this member belongs to for multi-cluster update management.                                                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec             | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                            | [FleetsMemberOperatorSpec](#FleetsMemberOperatorSpec)<br/><small>Optional</small>                                                                                    |
| owner                    | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/Fleet resource                                                                                | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |

FleetsMember_STATUS{#FleetsMember_STATUS}
-----------------------------------------

Used by: [FleetsMember](#FleetsMember).

| Property          | Description                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                    |
|-------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| clusterResourceId | The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ContainerService/&ZeroWidthSpace;managedClusters/&ZeroWidthSpace;{clusterName}'. | string<br/><small>Optional</small>                                                                                                                      |
| conditions        | The observed state of the resource                                                                                                                                                                                                                                                                                                                                         | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| eTag              | If eTag is provided in the response body, it may also be provided as a header per the normal etag convention. Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.        | string<br/><small>Optional</small>                                                                                                                      |
| group             | The group this member belongs to for multi-cluster update management.                                                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                      |
| id                | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}                                                  | string<br/><small>Optional</small>                                                                                                                      |
| name              | The name of the resource                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState | The status of the last operation.                                                                                                                                                                                                                                                                                                                                          | [FleetMemberProvisioningState_STATUS](#FleetMemberProvisioningState_STATUS)<br/><small>Optional</small>                                                 |
| systemData        | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                                                                           | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type              | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |

FleetsUpdateRun_Spec{#FleetsUpdateRun_Spec}
-------------------------------------------

Used by: [FleetsUpdateRun](#FleetsUpdateRun).

| Property             | Description                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                 |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName            | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                                   |
| managedClusterUpdate | The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be modified until the run is started.                                                                                                                                                                           | [ManagedClusterUpdate](#ManagedClusterUpdate)<br/><small>Required</small>                                                                                            |
| operatorSpec         | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                         | [FleetsUpdateRunOperatorSpec](#FleetsUpdateRunOperatorSpec)<br/><small>Optional</small>                                                                              |
| owner                | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a containerservice.azure.com/Fleet resource             | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| strategy             | The strategy defines the order in which the clusters will be updated. If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single UpdateGroup targeting all members. The strategy of the UpdateRun can be modified until the run is started. | [UpdateRunStrategy](#UpdateRunStrategy)<br/><small>Optional</small>                                                                                                  |

FleetsUpdateRun_STATUS{#FleetsUpdateRun_STATUS}
-----------------------------------------------

Used by: [FleetsUpdateRun](#FleetsUpdateRun).

| Property             | Description                                                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                    |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions           | The observed state of the resource                                                                                                                                                                                                                                                                                                                                  | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| eTag                 | If eTag is provided in the response body, it may also be provided as a header per the normal etag convention. Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. | string<br/><small>Optional</small>                                                                                                                      |
| id                   | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}                                           | string<br/><small>Optional</small>                                                                                                                      |
| managedClusterUpdate | The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be modified until the run is started.                                                                                                                                                                                                                                       | [ManagedClusterUpdate_STATUS](#ManagedClusterUpdate_STATUS)<br/><small>Optional</small>                                                                 |
| name                 | The name of the resource                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState    | The provisioning state of the UpdateRun resource.                                                                                                                                                                                                                                                                                                                   | [UpdateRunProvisioningState_STATUS](#UpdateRunProvisioningState_STATUS)<br/><small>Optional</small>                                                     |
| status               | The status of the UpdateRun.                                                                                                                                                                                                                                                                                                                                        | [UpdateRunStatus_STATUS](#UpdateRunStatus_STATUS)<br/><small>Optional</small>                                                                           |
| strategy             | The strategy defines the order in which the clusters will be updated. If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single UpdateGroup targeting all members. The strategy of the UpdateRun can be modified until the run is started.                                                             | [UpdateRunStrategy_STATUS](#UpdateRunStrategy_STATUS)<br/><small>Optional</small>                                                                       |
| systemData           | Azure Resource Manager metadata containing createdBy and modifiedBy information.                                                                                                                                                                                                                                                                                    | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type                 | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |

FleetHubProfile{#FleetHubProfile}
---------------------------------

The FleetHubProfile configures the fleet hub.

Used by: [Fleet_Spec](#Fleet_Spec).

| Property  | Description                                           | Type                               |
|-----------|-------------------------------------------------------|------------------------------------|
| dnsPrefix | DNS prefix used to create the FQDN for the Fleet hub. | string<br/><small>Optional</small> |

FleetHubProfile_STATUS{#FleetHubProfile_STATUS}
-----------------------------------------------

The FleetHubProfile configures the fleet hub.

Used by: [Fleet_STATUS](#Fleet_STATUS).

| Property          | Description                                           | Type                               |
|-------------------|-------------------------------------------------------|------------------------------------|
| dnsPrefix         | DNS prefix used to create the FQDN for the Fleet hub. | string<br/><small>Optional</small> |
| fqdn              | The FQDN of the Fleet hub.                            | string<br/><small>Optional</small> |
| kubernetesVersion | The Kubernetes version of the Fleet hub.              | string<br/><small>Optional</small> |

FleetMemberProvisioningState_STATUS{#FleetMemberProvisioningState_STATUS}
-------------------------------------------------------------------------

The provisioning state of the last accepted operation.

Used by: [FleetsMember_STATUS](#FleetsMember_STATUS).

| Value       | Description |
|-------------|-------------|
| "Canceled"  |             |
| "Failed"    |             |
| "Joining"   |             |
| "Leaving"   |             |
| "Succeeded" |             |
| "Updating"  |             |

FleetOperatorSpec{#FleetOperatorSpec}
-------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [Fleet_Spec](#Fleet_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secrets              | configures where to place Azure generated secrets.                                            | [FleetOperatorSecrets](#FleetOperatorSecrets)<br/><small>Optional</small>                                                                                           |

FleetProvisioningState_STATUS{#FleetProvisioningState_STATUS}
-------------------------------------------------------------

The provisioning state of the last accepted operation.

Used by: [Fleet_STATUS](#Fleet_STATUS).

| Value       | Description |
|-------------|-------------|
| "Canceled"  |             |
| "Creating"  |             |
| "Deleting"  |             |
| "Failed"    |             |
| "Succeeded" |             |
| "Updating"  |             |

FleetsMemberOperatorSpec{#FleetsMemberOperatorSpec}
---------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [FleetsMember_Spec](#FleetsMember_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

FleetsUpdateRunOperatorSpec{#FleetsUpdateRunOperatorSpec}
---------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [FleetsUpdateRun_Spec](#FleetsUpdateRun_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

ManagedClusterUpdate{#ManagedClusterUpdate}
-------------------------------------------

The update to be applied to the ManagedClusters.

Used by: [FleetsUpdateRun_Spec](#FleetsUpdateRun_Spec).

| Property | Description                                  | Type                                                                                |
|----------|----------------------------------------------|-------------------------------------------------------------------------------------|
| upgrade  | The upgrade to apply to the ManagedClusters. | [ManagedClusterUpgradeSpec](#ManagedClusterUpgradeSpec)<br/><small>Required</small> |

ManagedClusterUpdate_STATUS{#ManagedClusterUpdate_STATUS}
---------------------------------------------------------

The update to be applied to the ManagedClusters.

Used by: [FleetsUpdateRun_STATUS](#FleetsUpdateRun_STATUS).

| Property | Description                                  | Type                                                                                              |
|----------|----------------------------------------------|---------------------------------------------------------------------------------------------------|
| upgrade  | The upgrade to apply to the ManagedClusters. | [ManagedClusterUpgradeSpec_STATUS](#ManagedClusterUpgradeSpec_STATUS)<br/><small>Optional</small> |

SystemData_STATUS{#SystemData_STATUS}
-------------------------------------

Metadata pertaining to creation and last modification of the resource.

Used by: [Fleet_STATUS](#Fleet_STATUS), [FleetsMember_STATUS](#FleetsMember_STATUS), and [FleetsUpdateRun_STATUS](#FleetsUpdateRun_STATUS).

| Property           | Description                                           | Type                                                                                                      |
|--------------------|-------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| createdAt          | The timestamp of resource creation (UTC).             | string<br/><small>Optional</small>                                                                        |
| createdBy          | The identity that created the resource.               | string<br/><small>Optional</small>                                                                        |
| createdByType      | The type of identity that created the resource.       | [SystemData_CreatedByType_STATUS](#SystemData_CreatedByType_STATUS)<br/><small>Optional</small>           |
| lastModifiedAt     | The timestamp of resource last modification (UTC)     | string<br/><small>Optional</small>                                                                        |
| lastModifiedBy     | The identity that last modified the resource.         | string<br/><small>Optional</small>                                                                        |
| lastModifiedByType | The type of identity that last modified the resource. | [SystemData_LastModifiedByType_STATUS](#SystemData_LastModifiedByType_STATUS)<br/><small>Optional</small> |

UpdateRunProvisioningState_STATUS{#UpdateRunProvisioningState_STATUS}
---------------------------------------------------------------------

The provisioning state of the UpdateRun resource.

Used by: [FleetsUpdateRun_STATUS](#FleetsUpdateRun_STATUS).

| Value       | Description |
|-------------|-------------|
| "Canceled"  |             |
| "Failed"    |             |
| "Succeeded" |             |

UpdateRunStatus_STATUS{#UpdateRunStatus_STATUS}
-----------------------------------------------

The status of a UpdateRun.

Used by: [FleetsUpdateRun_STATUS](#FleetsUpdateRun_STATUS).

| Property | Description                                                                           | Type                                                                                |
|----------|---------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| stages   | The stages composing an update run. Stages are run sequentially withing an UpdateRun. | [UpdateStageStatus_STATUS[]](#UpdateStageStatus_STATUS)<br/><small>Optional</small> |
| status   | The status of the UpdateRun.                                                          | [UpdateStatus_STATUS](#UpdateStatus_STATUS)<br/><small>Optional</small>             |

UpdateRunStrategy{#UpdateRunStrategy}
-------------------------------------

Defines the update sequence of the clusters via stages and groups. Stages within a run are executed sequentially one after another. Groups within a stage are executed in parallel. Member clusters within a group are updated sequentially one after another. A valid strategy contains no duplicate groups within or across stages.

Used by: [FleetsUpdateRun_Spec](#FleetsUpdateRun_Spec).

| Property | Description                                                   | Type                                                      |
|----------|---------------------------------------------------------------|-----------------------------------------------------------|
| stages   | The list of stages that compose this update run. Min size: 1. | [UpdateStage[]](#UpdateStage)<br/><small>Required</small> |

UpdateRunStrategy_STATUS{#UpdateRunStrategy_STATUS}
---------------------------------------------------

Defines the update sequence of the clusters via stages and groups. Stages within a run are executed sequentially one after another. Groups within a stage are executed in parallel. Member clusters within a group are updated sequentially one after another. A valid strategy contains no duplicate groups within or across stages.

Used by: [FleetsUpdateRun_STATUS](#FleetsUpdateRun_STATUS).

| Property | Description                                                   | Type                                                                    |
|----------|---------------------------------------------------------------|-------------------------------------------------------------------------|
| stages   | The list of stages that compose this update run. Min size: 1. | [UpdateStage_STATUS[]](#UpdateStage_STATUS)<br/><small>Optional</small> |

FleetOperatorSecrets{#FleetOperatorSecrets}
-------------------------------------------

Used by: [FleetOperatorSpec](#FleetOperatorSpec).

| Property        | Description                                                                                                           | Type                                                                                                                                                       |
|-----------------|-----------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| userCredentials | indicates where the UserCredentials secret should be placed. If omitted, the secret will not be retrieved from Azure. | [genruntime.SecretDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination)<br/><small>Optional</small> |

ManagedClusterUpgradeSpec{#ManagedClusterUpgradeSpec}
-----------------------------------------------------

The upgrade to apply to a ManagedCluster.

Used by: [ManagedClusterUpdate](#ManagedClusterUpdate).

| Property          | Description                                                     | Type                                                                                |
|-------------------|-----------------------------------------------------------------|-------------------------------------------------------------------------------------|
| kubernetesVersion | The Kubernetes version to upgrade the member clusters to.       | string<br/><small>Optional</small>                                                  |
| type              | ManagedClusterUpgradeType is the type of upgrade to be applied. | [ManagedClusterUpgradeType](#ManagedClusterUpgradeType)<br/><small>Required</small> |

ManagedClusterUpgradeSpec_STATUS{#ManagedClusterUpgradeSpec_STATUS}
-------------------------------------------------------------------

The upgrade to apply to a ManagedCluster.

Used by: [ManagedClusterUpdate_STATUS](#ManagedClusterUpdate_STATUS).

| Property          | Description                                                     | Type                                                                                              |
|-------------------|-----------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| kubernetesVersion | The Kubernetes version to upgrade the member clusters to.       | string<br/><small>Optional</small>                                                                |
| type              | ManagedClusterUpgradeType is the type of upgrade to be applied. | [ManagedClusterUpgradeType_STATUS](#ManagedClusterUpgradeType_STATUS)<br/><small>Optional</small> |

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

UpdateStage{#UpdateStage}
-------------------------

Defines a stage which contains the groups to update and the steps to take (e.g., wait for a time period) before starting the next stage.

Used by: [UpdateRunStrategy](#UpdateRunStrategy).

| Property                | Description                                                                                                              | Type                                                      |
|-------------------------|--------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------|
| afterStageWaitInSeconds | The time in seconds to wait at the end of this stage before starting the next one. Defaults to 0 seconds if unspecified. | int<br/><small>Optional</small>                           |
| groups                  | Defines the groups to be executed in parallel in this stage. Duplicate groups are not allowed. Min size: 1.              | [UpdateGroup[]](#UpdateGroup)<br/><small>Optional</small> |
| name                    | The name of the stage. Must be unique within the UpdateRun.                                                              | string<br/><small>Required</small>                        |

UpdateStage_STATUS{#UpdateStage_STATUS}
---------------------------------------

Defines a stage which contains the groups to update and the steps to take (e.g., wait for a time period) before starting the next stage.

Used by: [UpdateRunStrategy_STATUS](#UpdateRunStrategy_STATUS).

| Property                | Description                                                                                                              | Type                                                                    |
|-------------------------|--------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| afterStageWaitInSeconds | The time in seconds to wait at the end of this stage before starting the next one. Defaults to 0 seconds if unspecified. | int<br/><small>Optional</small>                                         |
| groups                  | Defines the groups to be executed in parallel in this stage. Duplicate groups are not allowed. Min size: 1.              | [UpdateGroup_STATUS[]](#UpdateGroup_STATUS)<br/><small>Optional</small> |
| name                    | The name of the stage. Must be unique within the UpdateRun.                                                              | string<br/><small>Optional</small>                                      |

UpdateStageStatus_STATUS{#UpdateStageStatus_STATUS}
---------------------------------------------------

The status of a UpdateStage.

Used by: [UpdateRunStatus_STATUS](#UpdateRunStatus_STATUS).

| Property             | Description                                                   | Type                                                                                |
|----------------------|---------------------------------------------------------------|-------------------------------------------------------------------------------------|
| afterStageWaitStatus | The status of the wait period configured on the UpdateStage.  | [WaitStatus_STATUS](#WaitStatus_STATUS)<br/><small>Optional</small>                 |
| groups               | The list of groups to be updated as part of this UpdateStage. | [UpdateGroupStatus_STATUS[]](#UpdateGroupStatus_STATUS)<br/><small>Optional</small> |
| name                 | The name of the UpdateStage.                                  | string<br/><small>Optional</small>                                                  |
| status               | The status of the UpdateStage.                                | [UpdateStatus_STATUS](#UpdateStatus_STATUS)<br/><small>Optional</small>             |

UpdateStatus_STATUS{#UpdateStatus_STATUS}
-----------------------------------------

The status for an operation or group of operations.

Used by: [MemberUpdateStatus_STATUS](#MemberUpdateStatus_STATUS), [UpdateGroupStatus_STATUS](#UpdateGroupStatus_STATUS), [UpdateRunStatus_STATUS](#UpdateRunStatus_STATUS), [UpdateStageStatus_STATUS](#UpdateStageStatus_STATUS), and [WaitStatus_STATUS](#WaitStatus_STATUS).

| Property      | Description                                      | Type                                                                  |
|---------------|--------------------------------------------------|-----------------------------------------------------------------------|
| completedTime | The time the operation or group was completed.   | string<br/><small>Optional</small>                                    |
| error         | The error details when a failure is encountered. | [ErrorDetail_STATUS](#ErrorDetail_STATUS)<br/><small>Optional</small> |
| startTime     | The time the operation or group was started.     | string<br/><small>Optional</small>                                    |
| state         | The State of the operation or group.             | [UpdateState_STATUS](#UpdateState_STATUS)<br/><small>Optional</small> |

ErrorDetail_STATUS{#ErrorDetail_STATUS}
---------------------------------------

The error detail.

Used by: [UpdateStatus_STATUS](#UpdateStatus_STATUS).

| Property       | Description                | Type                                                                                      |
|----------------|----------------------------|-------------------------------------------------------------------------------------------|
| additionalInfo | The error additional info. | [ErrorAdditionalInfo_STATUS[]](#ErrorAdditionalInfo_STATUS)<br/><small>Optional</small>   |
| code           | The error code.            | string<br/><small>Optional</small>                                                        |
| details        | The error details.         | [ErrorDetail_STATUS_Unrolled[]](#ErrorDetail_STATUS_Unrolled)<br/><small>Optional</small> |
| message        | The error message.         | string<br/><small>Optional</small>                                                        |
| target         | The error target.          | string<br/><small>Optional</small>                                                        |

ManagedClusterUpgradeType{#ManagedClusterUpgradeType}
-----------------------------------------------------

The type of upgrade to perform when targeting ManagedClusters.

Used by: [ManagedClusterUpgradeSpec](#ManagedClusterUpgradeSpec).

| Value           | Description |
|-----------------|-------------|
| "Full"          |             |
| "NodeImageOnly" |             |

ManagedClusterUpgradeType_STATUS{#ManagedClusterUpgradeType_STATUS}
-------------------------------------------------------------------

The type of upgrade to perform when targeting ManagedClusters.

Used by: [ManagedClusterUpgradeSpec_STATUS](#ManagedClusterUpgradeSpec_STATUS).

| Value           | Description |
|-----------------|-------------|
| "Full"          |             |
| "NodeImageOnly" |             |

UpdateGroup{#UpdateGroup}
-------------------------

A group to be updated.

Used by: [UpdateStage](#UpdateStage).

| Property | Description                                                                | Type                               |
|----------|----------------------------------------------------------------------------|------------------------------------|
| name     | Name of the group. It must match a group name of an existing fleet member. | string<br/><small>Required</small> |

UpdateGroup_STATUS{#UpdateGroup_STATUS}
---------------------------------------

A group to be updated.

Used by: [UpdateStage_STATUS](#UpdateStage_STATUS).

| Property | Description                                                                | Type                               |
|----------|----------------------------------------------------------------------------|------------------------------------|
| name     | Name of the group. It must match a group name of an existing fleet member. | string<br/><small>Optional</small> |

UpdateGroupStatus_STATUS{#UpdateGroupStatus_STATUS}
---------------------------------------------------

The status of a UpdateGroup.

Used by: [UpdateStageStatus_STATUS](#UpdateStageStatus_STATUS).

| Property | Description                                  | Type                                                                                  |
|----------|----------------------------------------------|---------------------------------------------------------------------------------------|
| members  | The list of member this UpdateGroup updates. | [MemberUpdateStatus_STATUS[]](#MemberUpdateStatus_STATUS)<br/><small>Optional</small> |
| name     | The name of the UpdateGroup.                 | string<br/><small>Optional</small>                                                    |
| status   | The status of the UpdateGroup.               | [UpdateStatus_STATUS](#UpdateStatus_STATUS)<br/><small>Optional</small>               |

UpdateState_STATUS{#UpdateState_STATUS}
---------------------------------------

The state of the UpdateRun, UpdateStage, UpdateGroup, or MemberUpdate.

Used by: [UpdateStatus_STATUS](#UpdateStatus_STATUS).

| Value        | Description |
|--------------|-------------|
| "Completed"  |             |
| "Failed"     |             |
| "NotStarted" |             |
| "Running"    |             |
| "Stopped"    |             |
| "Stopping"   |             |

WaitStatus_STATUS{#WaitStatus_STATUS}
-------------------------------------

The status of the wait duration.

Used by: [UpdateStageStatus_STATUS](#UpdateStageStatus_STATUS).

| Property              | Description                              | Type                                                                    |
|-----------------------|------------------------------------------|-------------------------------------------------------------------------|
| status                | The status of the wait duration.         | [UpdateStatus_STATUS](#UpdateStatus_STATUS)<br/><small>Optional</small> |
| waitDurationInSeconds | The wait duration configured in seconds. | int<br/><small>Optional</small>                                         |

ErrorAdditionalInfo_STATUS{#ErrorAdditionalInfo_STATUS}
-------------------------------------------------------

The resource management error additional info.

Used by: [ErrorDetail_STATUS](#ErrorDetail_STATUS), and [ErrorDetail_STATUS_Unrolled](#ErrorDetail_STATUS_Unrolled).

| Property | Description               | Type                                           |
|----------|---------------------------|------------------------------------------------|
| info     | The additional info.      | map[string]v1.JSON<br/><small>Optional</small> |
| type     | The additional info type. | string<br/><small>Optional</small>             |

ErrorDetail_STATUS_Unrolled{#ErrorDetail_STATUS_Unrolled}
---------------------------------------------------------

Used by: [ErrorDetail_STATUS](#ErrorDetail_STATUS).

| Property       | Description                | Type                                                                                    |
|----------------|----------------------------|-----------------------------------------------------------------------------------------|
| additionalInfo | The error additional info. | [ErrorAdditionalInfo_STATUS[]](#ErrorAdditionalInfo_STATUS)<br/><small>Optional</small> |
| code           | The error code.            | string<br/><small>Optional</small>                                                      |
| message        | The error message.         | string<br/><small>Optional</small>                                                      |
| target         | The error target.          | string<br/><small>Optional</small>                                                      |

MemberUpdateStatus_STATUS{#MemberUpdateStatus_STATUS}
-----------------------------------------------------

The status of a member update operation.

Used by: [UpdateGroupStatus_STATUS](#UpdateGroupStatus_STATUS).

| Property          | Description                                                               | Type                                                                    |
|-------------------|---------------------------------------------------------------------------|-------------------------------------------------------------------------|
| clusterResourceId | The Azure resource id of the target Kubernetes cluster.                   | string<br/><small>Optional</small>                                      |
| name              | The name of the FleetMember.                                              | string<br/><small>Optional</small>                                      |
| operationId       | The operation resource id of the latest attempt to perform the operation. | string<br/><small>Optional</small>                                      |
| status            | The status of the MemberUpdate operation.                                 | [UpdateStatus_STATUS](#UpdateStatus_STATUS)<br/><small>Optional</small> |
