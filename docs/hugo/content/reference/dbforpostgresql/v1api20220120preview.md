---

title: dbforpostgresql.azure.com/v1api20220120preview

linktitle: v1api20220120preview
-------------------------------

APIVersion{#APIVersion}
-----------------------

| Value                | Description |
|----------------------|-------------|
| "2022-01-20-preview" |             |

FlexibleServer{#FlexibleServer}
-------------------------------

Generator information: - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.DBforPostgreSQL/&ZeroWidthSpace;flexibleServers/&ZeroWidthSpace;{serverName}

Used by: [FlexibleServerList](#FlexibleServerList).

| Property                                                                                | Description | Type                                                                        |
|-----------------------------------------------------------------------------------------|-------------|-----------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                             |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                             |
| spec                                                                                    |             | [FlexibleServer_Spec](#FlexibleServer_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [FlexibleServer_STATUS](#FlexibleServer_STATUS)<br/><small>Optional</small> |

### FlexibleServer_Spec {#FlexibleServer_Spec}

| Property                      | Description                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                 |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| administratorLogin            | The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                                   |
| administratorLoginPassword    | The administrator login password (required for server creation).                                                                                                                                                                                                                             | [genruntime.SecretReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference)<br/><small>Optional</small>               |
| availabilityZone              | availability zone information of the server.                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |
| azureName                     | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| backup                        | Backup properties of a server.                                                                                                                                                                                                                                                               | [Backup](#Backup)<br/><small>Optional</small>                                                                                                                        |
| createMode                    | The mode to create a new PostgreSQL server.                                                                                                                                                                                                                                                  | [ServerProperties_CreateMode](#ServerProperties_CreateMode)<br/><small>Optional</small>                                                                              |
| highAvailability              | High availability properties of a server.                                                                                                                                                                                                                                                    | [HighAvailability](#HighAvailability)<br/><small>Optional</small>                                                                                                    |
| location                      | The geo-location where the resource lives                                                                                                                                                                                                                                                    | string<br/><small>Required</small>                                                                                                                                   |
| maintenanceWindow             | Maintenance window properties of a server.                                                                                                                                                                                                                                                   | [MaintenanceWindow](#MaintenanceWindow)<br/><small>Optional</small>                                                                                                  |
| network                       | Network properties of a server.                                                                                                                                                                                                                                                              | [Network](#Network)<br/><small>Optional</small>                                                                                                                      |
| operatorSpec                  | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                              | [FlexibleServerOperatorSpec](#FlexibleServerOperatorSpec)<br/><small>Optional</small>                                                                                |
| owner                         | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| pointInTimeUTC                | Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when `createMode` is `PointInTimeRestore`.                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                                   |
| sku                           | The SKU (pricing tier) of the server.                                                                                                                                                                                                                                                        | [Sku](#Sku)<br/><small>Optional</small>                                                                                                                              |
| sourceServerResourceReference | The source server resource ID to restore from. It's required when `createMode` is `PointInTimeRestore`.                                                                                                                                                                                      | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| storage                       | Storage properties of a server.                                                                                                                                                                                                                                                              | [Storage](#Storage)<br/><small>Optional</small>                                                                                                                      |
| tags                          | Resource tags.                                                                                                                                                                                                                                                                               | map[string]string<br/><small>Optional</small>                                                                                                                        |
| version                       | PostgreSQL Server version.                                                                                                                                                                                                                                                                   | [ServerVersion](#ServerVersion)<br/><small>Optional</small>                                                                                                          |

### FlexibleServer_STATUS{#FlexibleServer_STATUS}

| Property                 | Description                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                    |
|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| administratorLogin       | The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).                                                                                                                                                                                        | string<br/><small>Optional</small>                                                                                                                      |
| availabilityZone         | availability zone information of the server.                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                      |
| backup                   | Backup properties of a server.                                                                                                                                                                                                                                                                                            | [Backup_STATUS](#Backup_STATUS)<br/><small>Optional</small>                                                                                             |
| conditions               | The observed state of the resource                                                                                                                                                                                                                                                                                        | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| createMode               | The mode to create a new PostgreSQL server.                                                                                                                                                                                                                                                                               | [ServerProperties_CreateMode_STATUS](#ServerProperties_CreateMode_STATUS)<br/><small>Optional</small>                                                   |
| fullyQualifiedDomainName | The fully qualified domain name of a server.                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                      |
| highAvailability         | High availability properties of a server.                                                                                                                                                                                                                                                                                 | [HighAvailability_STATUS](#HighAvailability_STATUS)<br/><small>Optional</small>                                                                         |
| id                       | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName} | string<br/><small>Optional</small>                                                                                                                      |
| location                 | The geo-location where the resource lives                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| maintenanceWindow        | Maintenance window properties of a server.                                                                                                                                                                                                                                                                                | [MaintenanceWindow_STATUS](#MaintenanceWindow_STATUS)<br/><small>Optional</small>                                                                       |
| minorVersion             | The minor version of the server.                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                      |
| name                     | The name of the resource                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| network                  | Network properties of a server.                                                                                                                                                                                                                                                                                           | [Network_STATUS](#Network_STATUS)<br/><small>Optional</small>                                                                                           |
| pointInTimeUTC           | Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when `createMode` is `PointInTimeRestore`.                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| sku                      | The SKU (pricing tier) of the server.                                                                                                                                                                                                                                                                                     | [Sku_STATUS](#Sku_STATUS)<br/><small>Optional</small>                                                                                                   |
| sourceServerResourceId   | The source server resource ID to restore from. It's required when `createMode` is `PointInTimeRestore`.                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| state                    | A state of a server that is visible to user.                                                                                                                                                                                                                                                                              | [ServerProperties_State_STATUS](#ServerProperties_State_STATUS)<br/><small>Optional</small>                                                             |
| storage                  | Storage properties of a server.                                                                                                                                                                                                                                                                                           | [Storage_STATUS](#Storage_STATUS)<br/><small>Optional</small>                                                                                           |
| systemData               | The system metadata relating to this resource.                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags                     | Resource tags.                                                                                                                                                                                                                                                                                                            | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                     | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| version                  | PostgreSQL Server version.                                                                                                                                                                                                                                                                                                | [ServerVersion_STATUS](#ServerVersion_STATUS)<br/><small>Optional</small>                                                                               |

FlexibleServerList{#FlexibleServerList}
---------------------------------------

Generator information: - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.DBforPostgreSQL/&ZeroWidthSpace;flexibleServers/&ZeroWidthSpace;{serverName}

| Property                                                                            | Description | Type                                                            |
|-------------------------------------------------------------------------------------|-------------|-----------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                 |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                 |
| items                                                                               |             | [FlexibleServer[]](#FlexibleServer)<br/><small>Optional</small> |

FlexibleServersConfiguration{#FlexibleServersConfiguration}
-----------------------------------------------------------

Generator information: - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.DBforPostgreSQL/&ZeroWidthSpace;flexibleServers/&ZeroWidthSpace;{serverName}/&ZeroWidthSpace;configurations/&ZeroWidthSpace;{configurationName}

Used by: [FlexibleServersConfigurationList](#FlexibleServersConfigurationList).

| Property                                                                                | Description | Type                                                                                                    |
|-----------------------------------------------------------------------------------------|-------------|---------------------------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                                                         |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                                                         |
| spec                                                                                    |             | [FlexibleServersConfiguration_Spec](#FlexibleServersConfiguration_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [FlexibleServersConfiguration_STATUS](#FlexibleServersConfiguration_STATUS)<br/><small>Optional</small> |

### FlexibleServersConfiguration_Spec {#FlexibleServersConfiguration_Spec}

| Property     | Description                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                                 |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                     | [FlexibleServersConfigurationOperatorSpec](#FlexibleServersConfigurationOperatorSpec)<br/><small>Optional</small>                                                    |
| owner        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a dbforpostgresql.azure.com/FlexibleServer resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| source       | Source of the configuration.                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                                                                                                   |
| value        | Value of the configuration.                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                                   |

### FlexibleServersConfiguration_STATUS{#FlexibleServersConfiguration_STATUS}

| Property               | Description                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                    |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| allowedValues          | Allowed values of the configuration.                                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                      |
| conditions             | The observed state of the resource                                                                                                                                                                                                                                                                                        | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| dataType               | Data type of the configuration.                                                                                                                                                                                                                                                                                           | [ConfigurationProperties_DataType_STATUS](#ConfigurationProperties_DataType_STATUS)<br/><small>Optional</small>                                         |
| defaultValue           | Default value of the configuration.                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| description            | Description of the configuration.                                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                      |
| documentationLink      | Configuration documentation link.                                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                      |
| id                     | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName} | string<br/><small>Optional</small>                                                                                                                      |
| isConfigPendingRestart | Configuration is pending restart or not.                                                                                                                                                                                                                                                                                  | bool<br/><small>Optional</small>                                                                                                                        |
| isDynamicConfig        | Configuration dynamic or static.                                                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>                                                                                                                        |
| isReadOnly             | Configuration read-only or not.                                                                                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                        |
| name                   | The name of the resource                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| source                 | Source of the configuration.                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                      |
| systemData             | The system metadata relating to this resource.                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type                   | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| unit                   | Configuration unit.                                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| value                  | Value of the configuration.                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |

FlexibleServersConfigurationList{#FlexibleServersConfigurationList}
-------------------------------------------------------------------

Generator information: - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.DBforPostgreSQL/&ZeroWidthSpace;flexibleServers/&ZeroWidthSpace;{serverName}/&ZeroWidthSpace;configurations/&ZeroWidthSpace;{configurationName}

| Property                                                                            | Description | Type                                                                                        |
|-------------------------------------------------------------------------------------|-------------|---------------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                                             |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                                             |
| items                                                                               |             | [FlexibleServersConfiguration[]](#FlexibleServersConfiguration)<br/><small>Optional</small> |

FlexibleServersDatabase{#FlexibleServersDatabase}
-------------------------------------------------

Generator information: - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/Databases.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.DBforPostgreSQL/&ZeroWidthSpace;flexibleServers/&ZeroWidthSpace;{serverName}/&ZeroWidthSpace;databases/&ZeroWidthSpace;{databaseName}

Used by: [FlexibleServersDatabaseList](#FlexibleServersDatabaseList).

| Property                                                                                | Description | Type                                                                                          |
|-----------------------------------------------------------------------------------------|-------------|-----------------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                                               |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                                               |
| spec                                                                                    |             | [FlexibleServersDatabase_Spec](#FlexibleServersDatabase_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [FlexibleServersDatabase_STATUS](#FlexibleServersDatabase_STATUS)<br/><small>Optional</small> |

### FlexibleServersDatabase_Spec {#FlexibleServersDatabase_Spec}

| Property     | Description                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                                 |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| charset      | The charset of the database.                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                                                                                                   |
| collation    | The collation of the database.                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                     | [FlexibleServersDatabaseOperatorSpec](#FlexibleServersDatabaseOperatorSpec)<br/><small>Optional</small>                                                              |
| owner        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a dbforpostgresql.azure.com/FlexibleServer resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |

### FlexibleServersDatabase_STATUS{#FlexibleServersDatabase_STATUS}

| Property   | Description                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                    |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| charset    | The charset of the database.                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                      |
| collation  | The collation of the database.                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| conditions | The observed state of the resource                                                                                                                                                                                                                                                                                        | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| id         | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName} | string<br/><small>Optional</small>                                                                                                                      |
| name       | The name of the resource                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| systemData | The system metadata relating to this resource.                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type       | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |

FlexibleServersDatabaseList{#FlexibleServersDatabaseList}
---------------------------------------------------------

Generator information: - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/Databases.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.DBforPostgreSQL/&ZeroWidthSpace;flexibleServers/&ZeroWidthSpace;{serverName}/&ZeroWidthSpace;databases/&ZeroWidthSpace;{databaseName}

| Property                                                                            | Description | Type                                                                              |
|-------------------------------------------------------------------------------------|-------------|-----------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                                   |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                                   |
| items                                                                               |             | [FlexibleServersDatabase[]](#FlexibleServersDatabase)<br/><small>Optional</small> |

FlexibleServersFirewallRule{#FlexibleServersFirewallRule}
---------------------------------------------------------

Generator information: - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.DBforPostgreSQL/&ZeroWidthSpace;flexibleServers/&ZeroWidthSpace;{serverName}/&ZeroWidthSpace;firewallRules/&ZeroWidthSpace;{firewallRuleName}

Used by: [FlexibleServersFirewallRuleList](#FlexibleServersFirewallRuleList).

| Property                                                                                | Description | Type                                                                                                  |
|-----------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                                                       |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                                                       |
| spec                                                                                    |             | [FlexibleServersFirewallRule_Spec](#FlexibleServersFirewallRule_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [FlexibleServersFirewallRule_STATUS](#FlexibleServersFirewallRule_STATUS)<br/><small>Optional</small> |

### FlexibleServersFirewallRule_Spec {#FlexibleServersFirewallRule_Spec}

| Property       | Description                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                                 |
|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName      | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| endIpAddress   | The end IP address of the server firewall rule. Must be IPv4 format.                                                                                                                                                                                                                                | string<br/><small>Required</small>                                                                                                                                   |
| operatorSpec   | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                     | [FlexibleServersFirewallRuleOperatorSpec](#FlexibleServersFirewallRuleOperatorSpec)<br/><small>Optional</small>                                                      |
| owner          | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a dbforpostgresql.azure.com/FlexibleServer resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| startIpAddress | The start IP address of the server firewall rule. Must be IPv4 format.                                                                                                                                                                                                                              | string<br/><small>Required</small>                                                                                                                                   |

### FlexibleServersFirewallRule_STATUS{#FlexibleServersFirewallRule_STATUS}

| Property       | Description                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                    |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions     | The observed state of the resource                                                                                                                                                                                                                                                                                        | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| endIpAddress   | The end IP address of the server firewall rule. Must be IPv4 format.                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                      |
| id             | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName} | string<br/><small>Optional</small>                                                                                                                      |
| name           | The name of the resource                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| startIpAddress | The start IP address of the server firewall rule. Must be IPv4 format.                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| systemData     | The system metadata relating to this resource.                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type           | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |

FlexibleServersFirewallRuleList{#FlexibleServersFirewallRuleList}
-----------------------------------------------------------------

Generator information: - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.DBforPostgreSQL/&ZeroWidthSpace;flexibleServers/&ZeroWidthSpace;{serverName}/&ZeroWidthSpace;firewallRules/&ZeroWidthSpace;{firewallRuleName}

| Property                                                                            | Description | Type                                                                                      |
|-------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                                           |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                                           |
| items                                                                               |             | [FlexibleServersFirewallRule[]](#FlexibleServersFirewallRule)<br/><small>Optional</small> |

FlexibleServer_Spec{#FlexibleServer_Spec}
-----------------------------------------

Used by: [FlexibleServer](#FlexibleServer).

| Property                      | Description                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                 |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| administratorLogin            | The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                                   |
| administratorLoginPassword    | The administrator login password (required for server creation).                                                                                                                                                                                                                             | [genruntime.SecretReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference)<br/><small>Optional</small>               |
| availabilityZone              | availability zone information of the server.                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |
| azureName                     | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| backup                        | Backup properties of a server.                                                                                                                                                                                                                                                               | [Backup](#Backup)<br/><small>Optional</small>                                                                                                                        |
| createMode                    | The mode to create a new PostgreSQL server.                                                                                                                                                                                                                                                  | [ServerProperties_CreateMode](#ServerProperties_CreateMode)<br/><small>Optional</small>                                                                              |
| highAvailability              | High availability properties of a server.                                                                                                                                                                                                                                                    | [HighAvailability](#HighAvailability)<br/><small>Optional</small>                                                                                                    |
| location                      | The geo-location where the resource lives                                                                                                                                                                                                                                                    | string<br/><small>Required</small>                                                                                                                                   |
| maintenanceWindow             | Maintenance window properties of a server.                                                                                                                                                                                                                                                   | [MaintenanceWindow](#MaintenanceWindow)<br/><small>Optional</small>                                                                                                  |
| network                       | Network properties of a server.                                                                                                                                                                                                                                                              | [Network](#Network)<br/><small>Optional</small>                                                                                                                      |
| operatorSpec                  | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                              | [FlexibleServerOperatorSpec](#FlexibleServerOperatorSpec)<br/><small>Optional</small>                                                                                |
| owner                         | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| pointInTimeUTC                | Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when `createMode` is `PointInTimeRestore`.                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                                   |
| sku                           | The SKU (pricing tier) of the server.                                                                                                                                                                                                                                                        | [Sku](#Sku)<br/><small>Optional</small>                                                                                                                              |
| sourceServerResourceReference | The source server resource ID to restore from. It's required when `createMode` is `PointInTimeRestore`.                                                                                                                                                                                      | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>           |
| storage                       | Storage properties of a server.                                                                                                                                                                                                                                                              | [Storage](#Storage)<br/><small>Optional</small>                                                                                                                      |
| tags                          | Resource tags.                                                                                                                                                                                                                                                                               | map[string]string<br/><small>Optional</small>                                                                                                                        |
| version                       | PostgreSQL Server version.                                                                                                                                                                                                                                                                   | [ServerVersion](#ServerVersion)<br/><small>Optional</small>                                                                                                          |

FlexibleServer_STATUS{#FlexibleServer_STATUS}
---------------------------------------------

Used by: [FlexibleServer](#FlexibleServer).

| Property                 | Description                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                    |
|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| administratorLogin       | The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).                                                                                                                                                                                        | string<br/><small>Optional</small>                                                                                                                      |
| availabilityZone         | availability zone information of the server.                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                      |
| backup                   | Backup properties of a server.                                                                                                                                                                                                                                                                                            | [Backup_STATUS](#Backup_STATUS)<br/><small>Optional</small>                                                                                             |
| conditions               | The observed state of the resource                                                                                                                                                                                                                                                                                        | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| createMode               | The mode to create a new PostgreSQL server.                                                                                                                                                                                                                                                                               | [ServerProperties_CreateMode_STATUS](#ServerProperties_CreateMode_STATUS)<br/><small>Optional</small>                                                   |
| fullyQualifiedDomainName | The fully qualified domain name of a server.                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                      |
| highAvailability         | High availability properties of a server.                                                                                                                                                                                                                                                                                 | [HighAvailability_STATUS](#HighAvailability_STATUS)<br/><small>Optional</small>                                                                         |
| id                       | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName} | string<br/><small>Optional</small>                                                                                                                      |
| location                 | The geo-location where the resource lives                                                                                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| maintenanceWindow        | Maintenance window properties of a server.                                                                                                                                                                                                                                                                                | [MaintenanceWindow_STATUS](#MaintenanceWindow_STATUS)<br/><small>Optional</small>                                                                       |
| minorVersion             | The minor version of the server.                                                                                                                                                                                                                                                                                          | string<br/><small>Optional</small>                                                                                                                      |
| name                     | The name of the resource                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| network                  | Network properties of a server.                                                                                                                                                                                                                                                                                           | [Network_STATUS](#Network_STATUS)<br/><small>Optional</small>                                                                                           |
| pointInTimeUTC           | Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when `createMode` is `PointInTimeRestore`.                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |
| sku                      | The SKU (pricing tier) of the server.                                                                                                                                                                                                                                                                                     | [Sku_STATUS](#Sku_STATUS)<br/><small>Optional</small>                                                                                                   |
| sourceServerResourceId   | The source server resource ID to restore from. It's required when `createMode` is `PointInTimeRestore`.                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| state                    | A state of a server that is visible to user.                                                                                                                                                                                                                                                                              | [ServerProperties_State_STATUS](#ServerProperties_State_STATUS)<br/><small>Optional</small>                                                             |
| storage                  | Storage properties of a server.                                                                                                                                                                                                                                                                                           | [Storage_STATUS](#Storage_STATUS)<br/><small>Optional</small>                                                                                           |
| systemData               | The system metadata relating to this resource.                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags                     | Resource tags.                                                                                                                                                                                                                                                                                                            | map[string]string<br/><small>Optional</small>                                                                                                           |
| type                     | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| version                  | PostgreSQL Server version.                                                                                                                                                                                                                                                                                                | [ServerVersion_STATUS](#ServerVersion_STATUS)<br/><small>Optional</small>                                                                               |

FlexibleServersConfiguration_Spec{#FlexibleServersConfiguration_Spec}
---------------------------------------------------------------------

Used by: [FlexibleServersConfiguration](#FlexibleServersConfiguration).

| Property     | Description                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                                 |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                     | [FlexibleServersConfigurationOperatorSpec](#FlexibleServersConfigurationOperatorSpec)<br/><small>Optional</small>                                                    |
| owner        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a dbforpostgresql.azure.com/FlexibleServer resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| source       | Source of the configuration.                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                                                                                                   |
| value        | Value of the configuration.                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                                   |

FlexibleServersConfiguration_STATUS{#FlexibleServersConfiguration_STATUS}
-------------------------------------------------------------------------

Used by: [FlexibleServersConfiguration](#FlexibleServersConfiguration).

| Property               | Description                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                    |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| allowedValues          | Allowed values of the configuration.                                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                      |
| conditions             | The observed state of the resource                                                                                                                                                                                                                                                                                        | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| dataType               | Data type of the configuration.                                                                                                                                                                                                                                                                                           | [ConfigurationProperties_DataType_STATUS](#ConfigurationProperties_DataType_STATUS)<br/><small>Optional</small>                                         |
| defaultValue           | Default value of the configuration.                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| description            | Description of the configuration.                                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                      |
| documentationLink      | Configuration documentation link.                                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                      |
| id                     | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName} | string<br/><small>Optional</small>                                                                                                                      |
| isConfigPendingRestart | Configuration is pending restart or not.                                                                                                                                                                                                                                                                                  | bool<br/><small>Optional</small>                                                                                                                        |
| isDynamicConfig        | Configuration dynamic or static.                                                                                                                                                                                                                                                                                          | bool<br/><small>Optional</small>                                                                                                                        |
| isReadOnly             | Configuration read-only or not.                                                                                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                        |
| name                   | The name of the resource                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| source                 | Source of the configuration.                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                      |
| systemData             | The system metadata relating to this resource.                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type                   | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |
| unit                   | Configuration unit.                                                                                                                                                                                                                                                                                                       | string<br/><small>Optional</small>                                                                                                                      |
| value                  | Value of the configuration.                                                                                                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                      |

FlexibleServersDatabase_Spec{#FlexibleServersDatabase_Spec}
-----------------------------------------------------------

Used by: [FlexibleServersDatabase](#FlexibleServersDatabase).

| Property     | Description                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                                 |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName    | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| charset      | The charset of the database.                                                                                                                                                                                                                                                                        | string<br/><small>Optional</small>                                                                                                                                   |
| collation    | The collation of the database.                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                     | [FlexibleServersDatabaseOperatorSpec](#FlexibleServersDatabaseOperatorSpec)<br/><small>Optional</small>                                                              |
| owner        | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a dbforpostgresql.azure.com/FlexibleServer resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |

FlexibleServersDatabase_STATUS{#FlexibleServersDatabase_STATUS}
---------------------------------------------------------------

Used by: [FlexibleServersDatabase](#FlexibleServersDatabase).

| Property   | Description                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                    |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| charset    | The charset of the database.                                                                                                                                                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                      |
| collation  | The collation of the database.                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| conditions | The observed state of the resource                                                                                                                                                                                                                                                                                        | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| id         | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName} | string<br/><small>Optional</small>                                                                                                                      |
| name       | The name of the resource                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| systemData | The system metadata relating to this resource.                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type       | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |

FlexibleServersFirewallRule_Spec{#FlexibleServersFirewallRule_Spec}
-------------------------------------------------------------------

Used by: [FlexibleServersFirewallRule](#FlexibleServersFirewallRule).

| Property       | Description                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                                 |
|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName      | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| endIpAddress   | The end IP address of the server firewall rule. Must be IPv4 format.                                                                                                                                                                                                                                | string<br/><small>Required</small>                                                                                                                                   |
| operatorSpec   | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                     | [FlexibleServersFirewallRuleOperatorSpec](#FlexibleServersFirewallRuleOperatorSpec)<br/><small>Optional</small>                                                      |
| owner          | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a dbforpostgresql.azure.com/FlexibleServer resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| startIpAddress | The start IP address of the server firewall rule. Must be IPv4 format.                                                                                                                                                                                                                              | string<br/><small>Required</small>                                                                                                                                   |

FlexibleServersFirewallRule_STATUS{#FlexibleServersFirewallRule_STATUS}
-----------------------------------------------------------------------

Used by: [FlexibleServersFirewallRule](#FlexibleServersFirewallRule).

| Property       | Description                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                    |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions     | The observed state of the resource                                                                                                                                                                                                                                                                                        | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| endIpAddress   | The end IP address of the server firewall rule. Must be IPv4 format.                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                      |
| id             | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName} | string<br/><small>Optional</small>                                                                                                                      |
| name           | The name of the resource                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| startIpAddress | The start IP address of the server firewall rule. Must be IPv4 format.                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| systemData     | The system metadata relating to this resource.                                                                                                                                                                                                                                                                            | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| type           | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                      |

Backup{#Backup}
---------------

Backup properties of a server

Used by: [FlexibleServer_Spec](#FlexibleServer_Spec).

| Property            | Description                                                               | Type                                                                                |
|---------------------|---------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| backupRetentionDays | Backup retention days for the server.                                     | int<br/><small>Optional</small>                                                     |
| geoRedundantBackup  | A value indicating whether Geo-Redundant backup is enabled on the server. | [Backup_GeoRedundantBackup](#Backup_GeoRedundantBackup)<br/><small>Optional</small> |

Backup_STATUS{#Backup_STATUS}
-----------------------------

Backup properties of a server

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS).

| Property            | Description                                                               | Type                                                                                              |
|---------------------|---------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| backupRetentionDays | Backup retention days for the server.                                     | int<br/><small>Optional</small>                                                                   |
| earliestRestoreDate | The earliest restore point time (ISO8601 format) for server.              | string<br/><small>Optional</small>                                                                |
| geoRedundantBackup  | A value indicating whether Geo-Redundant backup is enabled on the server. | [Backup_GeoRedundantBackup_STATUS](#Backup_GeoRedundantBackup_STATUS)<br/><small>Optional</small> |

ConfigurationProperties_DataType_STATUS{#ConfigurationProperties_DataType_STATUS}
---------------------------------------------------------------------------------

Used by: [FlexibleServersConfiguration_STATUS](#FlexibleServersConfiguration_STATUS).

| Value         | Description |
|---------------|-------------|
| "Boolean"     |             |
| "Enumeration" |             |
| "Integer"     |             |
| "Numeric"     |             |

FlexibleServerOperatorSpec{#FlexibleServerOperatorSpec}
-------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [FlexibleServer_Spec](#FlexibleServer_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| configMaps           | configures where to place operator written ConfigMaps.                                        | [FlexibleServerOperatorConfigMaps](#FlexibleServerOperatorConfigMaps)<br/><small>Optional</small>                                                                   |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secrets              | configures where to place Azure generated secrets.                                            | [FlexibleServerOperatorSecrets](#FlexibleServerOperatorSecrets)<br/><small>Optional</small>                                                                         |

FlexibleServersConfigurationOperatorSpec{#FlexibleServersConfigurationOperatorSpec}
-----------------------------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [FlexibleServersConfiguration_Spec](#FlexibleServersConfiguration_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

FlexibleServersDatabaseOperatorSpec{#FlexibleServersDatabaseOperatorSpec}
-------------------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [FlexibleServersDatabase_Spec](#FlexibleServersDatabase_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

FlexibleServersFirewallRuleOperatorSpec{#FlexibleServersFirewallRuleOperatorSpec}
---------------------------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [FlexibleServersFirewallRule_Spec](#FlexibleServersFirewallRule_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

HighAvailability{#HighAvailability}
-----------------------------------

High availability properties of a server

Used by: [FlexibleServer_Spec](#FlexibleServer_Spec).

| Property                | Description                                   | Type                                                                        |
|-------------------------|-----------------------------------------------|-----------------------------------------------------------------------------|
| mode                    | The HA mode for the server.                   | [HighAvailability_Mode](#HighAvailability_Mode)<br/><small>Optional</small> |
| standbyAvailabilityZone | availability zone information of the standby. | string<br/><small>Optional</small>                                          |

HighAvailability_STATUS{#HighAvailability_STATUS}
-------------------------------------------------

High availability properties of a server

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS).

| Property                | Description                                     | Type                                                                                        |
|-------------------------|-------------------------------------------------|---------------------------------------------------------------------------------------------|
| mode                    | The HA mode for the server.                     | [HighAvailability_Mode_STATUS](#HighAvailability_Mode_STATUS)<br/><small>Optional</small>   |
| standbyAvailabilityZone | availability zone information of the standby.   | string<br/><small>Optional</small>                                                          |
| state                   | A state of a HA server that is visible to user. | [HighAvailability_State_STATUS](#HighAvailability_State_STATUS)<br/><small>Optional</small> |

MaintenanceWindow{#MaintenanceWindow}
-------------------------------------

Maintenance window properties of a server.

Used by: [FlexibleServer_Spec](#FlexibleServer_Spec).

| Property     | Description                                            | Type                               |
|--------------|--------------------------------------------------------|------------------------------------|
| customWindow | indicates whether custom window is enabled or disabled | string<br/><small>Optional</small> |
| dayOfWeek    | day of week for maintenance window                     | int<br/><small>Optional</small>    |
| startHour    | start hour for maintenance window                      | int<br/><small>Optional</small>    |
| startMinute  | start minute for maintenance window                    | int<br/><small>Optional</small>    |

MaintenanceWindow_STATUS{#MaintenanceWindow_STATUS}
---------------------------------------------------

Maintenance window properties of a server.

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS).

| Property     | Description                                            | Type                               |
|--------------|--------------------------------------------------------|------------------------------------|
| customWindow | indicates whether custom window is enabled or disabled | string<br/><small>Optional</small> |
| dayOfWeek    | day of week for maintenance window                     | int<br/><small>Optional</small>    |
| startHour    | start hour for maintenance window                      | int<br/><small>Optional</small>    |
| startMinute  | start minute for maintenance window                    | int<br/><small>Optional</small>    |

Network{#Network}
-----------------

Network properties of a server

Used by: [FlexibleServer_Spec](#FlexibleServer_Spec).

| Property                           | Description                       | Type                                                                                                                                                       |
|------------------------------------|-----------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| delegatedSubnetResourceReference   | delegated subnet arm resource id. | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| privateDnsZoneArmResourceReference | private dns zone arm resource id. | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

Network_STATUS{#Network_STATUS}
-------------------------------

Network properties of a server

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS).

| Property                    | Description                             | Type                                                                                                  |
|-----------------------------|-----------------------------------------|-------------------------------------------------------------------------------------------------------|
| delegatedSubnetResourceId   | delegated subnet arm resource id.       | string<br/><small>Optional</small>                                                                    |
| privateDnsZoneArmResourceId | private dns zone arm resource id.       | string<br/><small>Optional</small>                                                                    |
| publicNetworkAccess         | public network access is enabled or not | [Network_PublicNetworkAccess_STATUS](#Network_PublicNetworkAccess_STATUS)<br/><small>Optional</small> |

ServerProperties_CreateMode{#ServerProperties_CreateMode}
---------------------------------------------------------

Used by: [FlexibleServer_Spec](#FlexibleServer_Spec).

| Value                | Description |
|----------------------|-------------|
| "Create"             |             |
| "Default"            |             |
| "PointInTimeRestore" |             |
| "Update"             |             |

ServerProperties_CreateMode_STATUS{#ServerProperties_CreateMode_STATUS}
-----------------------------------------------------------------------

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS).

| Value                | Description |
|----------------------|-------------|
| "Create"             |             |
| "Default"            |             |
| "PointInTimeRestore" |             |
| "Update"             |             |

ServerProperties_State_STATUS{#ServerProperties_State_STATUS}
-------------------------------------------------------------

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Dropping" |             |
| "Ready"    |             |
| "Starting" |             |
| "Stopped"  |             |
| "Stopping" |             |
| "Updating" |             |

ServerVersion{#ServerVersion}
-----------------------------

The version of a server.

Used by: [FlexibleServer_Spec](#FlexibleServer_Spec).

| Value | Description |
|-------|-------------|
| "11"  |             |
| "12"  |             |
| "13"  |             |
| "14"  |             |

ServerVersion_STATUS{#ServerVersion_STATUS}
-------------------------------------------

The version of a server.

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS).

| Value | Description |
|-------|-------------|
| "11"  |             |
| "12"  |             |
| "13"  |             |
| "14"  |             |

Sku{#Sku}
---------

Sku information related properties of a server.

Used by: [FlexibleServer_Spec](#FlexibleServer_Spec).

| Property | Description                                                                  | Type                                              |
|----------|------------------------------------------------------------------------------|---------------------------------------------------|
| name     | The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3. | string<br/><small>Required</small>                |
| tier     | The tier of the particular SKU, e.g. Burstable.                              | [Sku_Tier](#Sku_Tier)<br/><small>Required</small> |

Sku_STATUS{#Sku_STATUS}
-----------------------

Sku information related properties of a server.

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS).

| Property | Description                                                                  | Type                                                            |
|----------|------------------------------------------------------------------------------|-----------------------------------------------------------------|
| name     | The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3. | string<br/><small>Optional</small>                              |
| tier     | The tier of the particular SKU, e.g. Burstable.                              | [Sku_Tier_STATUS](#Sku_Tier_STATUS)<br/><small>Optional</small> |

Storage{#Storage}
-----------------

Storage properties of a server

Used by: [FlexibleServer_Spec](#FlexibleServer_Spec).

| Property      | Description                       | Type                            |
|---------------|-----------------------------------|---------------------------------|
| storageSizeGB | Max storage allowed for a server. | int<br/><small>Optional</small> |

Storage_STATUS{#Storage_STATUS}
-------------------------------

Storage properties of a server

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS).

| Property      | Description                       | Type                            |
|---------------|-----------------------------------|---------------------------------|
| storageSizeGB | Max storage allowed for a server. | int<br/><small>Optional</small> |

SystemData_STATUS{#SystemData_STATUS}
-------------------------------------

Metadata pertaining to creation and last modification of the resource.

Used by: [FlexibleServer_STATUS](#FlexibleServer_STATUS), [FlexibleServersConfiguration_STATUS](#FlexibleServersConfiguration_STATUS), [FlexibleServersDatabase_STATUS](#FlexibleServersDatabase_STATUS), and [FlexibleServersFirewallRule_STATUS](#FlexibleServersFirewallRule_STATUS).

| Property           | Description                                           | Type                                                                                                      |
|--------------------|-------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| createdAt          | The timestamp of resource creation (UTC).             | string<br/><small>Optional</small>                                                                        |
| createdBy          | The identity that created the resource.               | string<br/><small>Optional</small>                                                                        |
| createdByType      | The type of identity that created the resource.       | [SystemData_CreatedByType_STATUS](#SystemData_CreatedByType_STATUS)<br/><small>Optional</small>           |
| lastModifiedAt     | The timestamp of resource last modification (UTC)     | string<br/><small>Optional</small>                                                                        |
| lastModifiedBy     | The identity that last modified the resource.         | string<br/><small>Optional</small>                                                                        |
| lastModifiedByType | The type of identity that last modified the resource. | [SystemData_LastModifiedByType_STATUS](#SystemData_LastModifiedByType_STATUS)<br/><small>Optional</small> |

Backup_GeoRedundantBackup{#Backup_GeoRedundantBackup}
-----------------------------------------------------

Used by: [Backup](#Backup).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Enabled"  |             |

Backup_GeoRedundantBackup_STATUS{#Backup_GeoRedundantBackup_STATUS}
-------------------------------------------------------------------

Used by: [Backup_STATUS](#Backup_STATUS).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Enabled"  |             |

FlexibleServerOperatorConfigMaps{#FlexibleServerOperatorConfigMaps}
-------------------------------------------------------------------

Used by: [FlexibleServerOperatorSpec](#FlexibleServerOperatorSpec).

| Property                 | Description                                                                                                          | Type                                                                                                                                                             |
|--------------------------|----------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| fullyQualifiedDomainName | indicates where the FullyQualifiedDomainName config map should be placed. If omitted, no config map will be created. | [genruntime.ConfigMapDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapDestination)<br/><small>Optional</small> |

FlexibleServerOperatorSecrets{#FlexibleServerOperatorSecrets}
-------------------------------------------------------------

Used by: [FlexibleServerOperatorSpec](#FlexibleServerOperatorSpec).

| Property                 | Description                                                                                                                    | Type                                                                                                                                                       |
|--------------------------|--------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| fullyQualifiedDomainName | indicates where the FullyQualifiedDomainName secret should be placed. If omitted, the secret will not be retrieved from Azure. | [genruntime.SecretDestination](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination)<br/><small>Optional</small> |

HighAvailability_Mode{#HighAvailability_Mode}
---------------------------------------------

Used by: [HighAvailability](#HighAvailability).

| Value           | Description |
|-----------------|-------------|
| "Disabled"      |             |
| "SameZone"      |             |
| "ZoneRedundant" |             |

HighAvailability_Mode_STATUS{#HighAvailability_Mode_STATUS}
-----------------------------------------------------------

Used by: [HighAvailability_STATUS](#HighAvailability_STATUS).

| Value           | Description |
|-----------------|-------------|
| "Disabled"      |             |
| "SameZone"      |             |
| "ZoneRedundant" |             |

HighAvailability_State_STATUS{#HighAvailability_State_STATUS}
-------------------------------------------------------------

Used by: [HighAvailability_STATUS](#HighAvailability_STATUS).

| Value             | Description |
|-------------------|-------------|
| "CreatingStandby" |             |
| "FailingOver"     |             |
| "Healthy"         |             |
| "NotEnabled"      |             |
| "RemovingStandby" |             |
| "ReplicatingData" |             |

Network_PublicNetworkAccess_STATUS{#Network_PublicNetworkAccess_STATUS}
-----------------------------------------------------------------------

Used by: [Network_STATUS](#Network_STATUS).

| Value      | Description |
|------------|-------------|
| "Disabled" |             |
| "Enabled"  |             |

Sku_Tier{#Sku_Tier}
-------------------

Used by: [Sku](#Sku).

| Value             | Description |
|-------------------|-------------|
| "Burstable"       |             |
| "GeneralPurpose"  |             |
| "MemoryOptimized" |             |

Sku_Tier_STATUS{#Sku_Tier_STATUS}
---------------------------------

Used by: [Sku_STATUS](#Sku_STATUS).

| Value             | Description |
|-------------------|-------------|
| "Burstable"       |             |
| "GeneralPurpose"  |             |
| "MemoryOptimized" |             |

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
