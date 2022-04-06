# Supported Resources

These are the resources with Azure Service Operator support committed to our **main** branch, grouped by the originating ARM service. (Newly supported resources will appear in this list prior to inclusion in any ASO release.)

## Authorization

| Name           | ARM Version        | CRD Version                | Sample                                                                                                                                                            |
|----------------|--------------------|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| RoleAssignment | 2020-08-01-preview | v1alpha1api20200801preview | [RoleAssignment sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/authorization/v1alpha1api20200801preview_roleassignment.yaml) |

## Batch

| Name         | ARM Version | CRD Version         | Sample                                                                                                                                         |
|--------------|-------------|---------------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| BatchAccount | 2021-01-01  | v1alpha1api20210101 | [BatchAccount sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/batch/v1alpha1api20210101_batchaccount.yaml) |

## Cache

| Name                    | ARM Version | CRD Version         | Sample                                                                                                                                                               |
|-------------------------|-------------|---------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Redis                   | 2020-12-01  | v1alpha1api20201201 | [Redis sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redis.yaml)                                     |
| RedisEnterprise         | 2021-03-01  | v1alpha1api20210301 | [RedisEnterprise sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20210301_redisenterprise.yaml)                 |
| RedisEnterpriseDatabase | 2021-03-01  | v1alpha1api20210301 | [RedisEnterpriseDatabase sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20210301_redisenterprisedatabase.yaml) |
| RedisFirewallRule       | 2020-12-01  | v1alpha1api20201201 | [RedisFirewallRule sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redisfirewallrule.yaml)             |
| RedisLinkedServer       | 2020-12-01  | v1alpha1api20201201 | [RedisLinkedServer sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redislinkedserver.yaml)             |
| RedisPatchSchedule      | 2020-12-01  | v1alpha1api20201201 | [RedisPatchSchedule sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redispatchschedule.yaml)           |

## Compute

| Name                   | ARM Version | CRD Version         | Sample                                                                                                                                                               |
|------------------------|-------------|---------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Disk                   | 2020-09-30  | v1alpha1api20200930 | [Disk sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20200930_disk.yaml)                                     |
| Image                  | 2021-07-01  | v1alpha1api20210701 | [Image sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20210701_image.yaml)                                   |
| Snapshot               | 2020-09-30  | v1alpha1api20200930 | [Snapshot sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20200930_snapshot.yaml)                             |
| VirtualMachine         | 2020-12-01  | v1alpha1api20201201 | [VirtualMachine sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20201201_virtualmachine.yaml)                 |
| VirtualMachineScaleSet | 2020-12-01  | v1alpha1api20201201 | [VirtualMachineScaleSet sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20201201_virtualmachinescaleset.yaml) |

## Containerregistry

| Name     | ARM Version | CRD Version         | Sample                                                                                                                                             |
|----------|-------------|---------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| Registry | 2021-09-01  | v1alpha1api20210901 | [Registry sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerregistry/v1alpha1api20210901_registry.yaml) |

## Containerservice

| Name                     | ARM Version | CRD Version         | Sample                                                                                                                                                                            |
|--------------------------|-------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ManagedCluster           | 2021-05-01  | v1alpha1api20210501 | [ManagedCluster sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerservice/v1alpha1api20210501_managedcluster.yaml)                     |
| ManagedClustersAgentPool | 2021-05-01  | v1alpha1api20210501 | [ManagedClustersAgentPool sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerservice/v1alpha1api20210501_managedclustersagentpool.yaml) |

## Dbformysql

| Name                        | ARM Version | CRD Version         | Sample                                                                                                                                                                            |
|-----------------------------|-------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| FlexibleServer              | 2021-05-01  | v1alpha1api20210501 | [FlexibleServer sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbformysql/v1alpha1api20210501_flexibleserver.yaml)                           |
| FlexibleServersDatabase     | 2021-05-01  | v1alpha1api20210501 | [FlexibleServersDatabase sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbformysql/v1alpha1api20210501_flexibleserversdatabase.yaml)         |
| FlexibleServersFirewallRule | 2021-05-01  | v1alpha1api20210501 | [FlexibleServersFirewallRule sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbformysql/v1alpha1api20210501_flexibleserversfirewallrule.yaml) |

## Dbforpostgresql

| Name                         | ARM Version | CRD Version         | Sample                                                                                                                                                                                   |
|------------------------------|-------------|---------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| FlexibleServer               | 2021-06-01  | v1alpha1api20210601 | [FlexibleServer sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbforpostgresql/v1alpha1api20210601_flexibleserver.yaml)                             |
| FlexibleServersConfiguration | 2021-06-01  | v1alpha1api20210601 | [FlexibleServersConfiguration sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbforpostgresql/v1alpha1api20210601_flexibleserversconfiguration.yaml) |
| FlexibleServersDatabase      | 2021-06-01  | v1alpha1api20210601 | [FlexibleServersDatabase sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbforpostgresql/v1alpha1api20210601_flexibleserversdatabase.yaml)           |
| FlexibleServersFirewallRule  | 2021-06-01  | v1alpha1api20210601 | [FlexibleServersFirewallRule sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbforpostgresql/v1alpha1api20210601_flexibleserversfirewallrule.yaml)   |

## Documentdb

| Name                                       | ARM Version | CRD Version         | Sample                                                                                                                                                                                                          |
|--------------------------------------------|-------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| DatabaseAccount                            | 2021-05-15  | v1alpha1api20210515 | [DatabaseAccount sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_databaseaccount.yaml)                                                       |
| MongodbDatabase                            | 2021-05-15  | v1alpha1api20210515 | [MongodbDatabase sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_mongodbdatabase.yaml)                                                       |
| MongodbDatabaseCollection                  | 2021-05-15  | v1alpha1api20210515 | [MongodbDatabaseCollection sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_mongodbdatabasecollection.yaml)                                   |
| MongodbDatabaseCollectionThroughputSetting | 2021-05-15  | v1alpha1api20210515 | [MongodbDatabaseCollectionThroughputSetting sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_mongodbdatabasecollectionthroughputsetting.yaml) |
| MongodbDatabaseThroughputSetting           | 2021-05-15  | v1alpha1api20210515 | [MongodbDatabaseThroughputSetting sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_mongodbdatabasethroughputsetting.yaml)                     |
| SqlDatabase                                | 2021-05-15  | v1alpha1api20210515 | [SqlDatabase sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabase.yaml)                                                               |
| SqlDatabaseContainer                       | 2021-05-15  | v1alpha1api20210515 | [SqlDatabaseContainer sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontainer.yaml)                                             |
| SqlDatabaseContainerStoredProcedure        | 2021-05-15  | v1alpha1api20210515 | [SqlDatabaseContainerStoredProcedure sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontainerstoredprocedure.yaml)               |
| SqlDatabaseContainerThroughputSetting      | 2021-05-15  | v1alpha1api20210515 | [SqlDatabaseContainerThroughputSetting sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontainerthroughputsetting.yaml)           |
| SqlDatabaseContainerTrigger                | 2021-05-15  | v1alpha1api20210515 | [SqlDatabaseContainerTrigger sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontainertrigger.yaml)                               |
| SqlDatabaseContainerUserDefinedFunction    | 2021-05-15  | v1alpha1api20210515 | [SqlDatabaseContainerUserDefinedFunction sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontaineruserdefinedfunction.yaml)       |
| SqlDatabaseThroughputSetting               | 2021-05-15  | v1alpha1api20210515 | [SqlDatabaseThroughputSetting sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasethroughputsetting.yaml)                             |

## Eventgrid

| Name              | ARM Version | CRD Version         | Sample                                                                                                                                                       |
|-------------------|-------------|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Domain            | 2020-06-01  | v1alpha1api20200601 | [Domain sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_domain.yaml)                       |
| DomainsTopic      | 2020-06-01  | v1alpha1api20200601 | [DomainsTopic sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_domainstopic.yaml)           |
| EventSubscription | 2020-06-01  | v1alpha1api20200601 | [EventSubscription sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_eventsubscription.yaml) |
| Topic             | 2020-06-01  | v1alpha1api20200601 | [Topic sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_topic.yaml)                         |

## Eventhub

| Name                                 | ARM Version | CRD Version         | Sample                                                                                                                                                                                            |
|--------------------------------------|-------------|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Namespace                            | 2021-11-01  | v1alpha1api20211101 | [Namespace sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespace.yaml)                                                       |
| NamespacesAuthorizationRule          | 2021-11-01  | v1alpha1api20211101 | [NamespacesAuthorizationRule sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespacesauthorizationrule.yaml)                   |
| NamespacesEventhub                   | 2021-11-01  | v1alpha1api20211101 | [NamespacesEventhub sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespaceseventhub.yaml)                                     |
| NamespacesEventhubsAuthorizationRule | 2021-11-01  | v1alpha1api20211101 | [NamespacesEventhubsAuthorizationRule sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespaceseventhubsauthorizationrule.yaml) |
| NamespacesEventhubsConsumerGroup     | 2021-11-01  | v1alpha1api20211101 | [NamespacesEventhubsConsumerGroup sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespaceseventhubsconsumergroup.yaml)         |

## Insights

| Name      | ARM Version        | CRD Version                | Sample                                                                                                                                         |
|-----------|--------------------|----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| Component | 2020-02-02         | v1alpha1api20200202        | [Component sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/insights/v1alpha1api20200202_component.yaml)    |
| Webtest   | 2018-05-01-preview | v1alpha1api20180501preview | [Webtest sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/insights/v1alpha1api20180501preview_webtest.yaml) |

## Managedidentity

| Name                 | ARM Version | CRD Version         | Sample                                                                                                                                                                   |
|----------------------|-------------|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| UserAssignedIdentity | 2018-11-30  | v1alpha1api20181130 | [UserAssignedIdentity sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/managedidentity/v1alpha1api20181130_userassignedidentity.yaml) |

## Network

| Name                                 | ARM Version | CRD Version         | Sample                                                                                                                                                                                           |
|--------------------------------------|-------------|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| LoadBalancer                         | 2020-11-01  | v1alpha1api20201101 | [LoadBalancer sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_loadbalancer.yaml)                                                 |
| NetworkInterface                     | 2020-11-01  | v1alpha1api20201101 | [NetworkInterface sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_networkinterface.yaml)                                         |
| NetworkSecurityGroup                 | 2020-11-01  | v1alpha1api20201101 | [NetworkSecurityGroup sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_networksecuritygroup.yaml)                                 |
| NetworkSecurityGroupsSecurityRule    | 2020-11-01  | v1alpha1api20201101 | [NetworkSecurityGroupsSecurityRule sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_networksecuritygroupssecurityrule.yaml)       |
| PublicIPAddress                      | 2020-11-01  | v1alpha1api20201101 | [PublicIPAddress sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_publicipaddress.yaml)                                           |
| VirtualNetwork                       | 2020-11-01  | v1alpha1api20201101 | [VirtualNetwork sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetwork.yaml)                                             |
| VirtualNetworkGateway                | 2020-11-01  | v1alpha1api20201101 | [VirtualNetworkGateway sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetworkgateway.yaml)                               |
| VirtualNetworksSubnet                | 2020-11-01  | v1alpha1api20201101 | [VirtualNetworksSubnet sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetworkssubnet.yaml)                               |
| VirtualNetworksVirtualNetworkPeering | 2020-11-01  | v1alpha1api20201101 | [VirtualNetworksVirtualNetworkPeering sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetworksvirtualnetworkpeering.yaml) |

## Operationalinsights

| Name      | ARM Version | CRD Version         | Sample                                                                                                                                                 |
|-----------|-------------|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| Workspace | 2021-06-01  | v1alpha1api20210601 | [Workspace sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/operationalinsights/v1alpha1api20210601_workspace.yaml) |

## Servicebus

| Name            | ARM Version        | CRD Version                | Sample                                                                                                                                                           |
|-----------------|--------------------|----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Namespace       | 2021-01-01-preview | v1alpha1api20210101preview | [Namespace sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/servicebus/v1alpha1api20210101preview_namespace.yaml)             |
| NamespacesQueue | 2021-01-01-preview | v1alpha1api20210101preview | [NamespacesQueue sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/servicebus/v1alpha1api20210101preview_namespacesqueue.yaml) |
| NamespacesTopic | 2021-01-01-preview | v1alpha1api20210101preview | [NamespacesTopic sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/servicebus/v1alpha1api20210101preview_namespacestopic.yaml) |

## Signalrservice

| Name    | ARM Version | CRD Version         | Sample                                                                                                                                        |
|---------|-------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| SignalR | 2021-10-01  | v1alpha1api20211001 | [SignalR sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/signalrservice/v1alpha1api20211001_signalr.yaml) |

## Storage

| Name                                 | ARM Version | CRD Version         | Sample                                                                                                                                                                                           |
|--------------------------------------|-------------|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| StorageAccount                       | 2021-04-01  | v1alpha1api20210401 | [StorageAccount sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccount.yaml)                                             |
| StorageAccountsBlobService           | 2021-04-01  | v1alpha1api20210401 | [StorageAccountsBlobService sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsblobservice.yaml)                     |
| StorageAccountsBlobServicesContainer | 2021-04-01  | v1alpha1api20210401 | [StorageAccountsBlobServicesContainer sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsblobservicescontainer.yaml) |
| StorageAccountsManagementPolicy      | 2021-04-01  | v1alpha1api20210401 | [StorageAccountsManagementPolicy sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsmanagementpolicy.yaml)           |
| StorageAccountsQueueService          | 2021-04-01  | v1alpha1api20210401 | [StorageAccountsQueueService sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsqueueservice.yaml)                   |
| StorageAccountsQueueServicesQueue    | 2021-04-01  | v1alpha1api20210401 | [StorageAccountsQueueServicesQueue sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsqueueservicesqueue.yaml)       |

