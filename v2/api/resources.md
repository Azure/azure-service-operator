# Supported Resources

These are the resources with Azure Service Operator support committed to our **main** branch, grouped by the originating ARM service. (Newly supported resources will appear in this list prior to inclusion in any ASO release.)

## authorization


### ARM version 2020-08-01-preview

- RoleAssignment ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/authorization/v1alpha1api20200801preview_roleassignment.yaml))

Use CRD version `v1alpha1api20200801preview`

## batch


### ARM version 2021-01-01

- BatchAccount ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/batch/v1alpha1api20210101_batchaccount.yaml))

Use CRD version `v1alpha1api20210101`

## cache


### ARM version 2020-12-01

- Redis ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redis.yaml))
- RedisFirewallRule ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redisfirewallrule.yaml))
- RedisLinkedServer ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redislinkedserver.yaml))
- RedisPatchSchedule ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redispatchschedule.yaml))

Use CRD version `v1alpha1api20201201`


### ARM version 2021-03-01

- RedisEnterprise ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20210301_redisenterprise.yaml))
- RedisEnterpriseDatabase ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20210301_redisenterprisedatabase.yaml))

Use CRD version `v1alpha1api20210301`

## compute


### ARM version 2020-09-30

- Disk ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20200930_disk.yaml))
- Snapshot ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20200930_snapshot.yaml))

Use CRD version `v1alpha1api20200930`


### ARM version 2020-12-01

- VirtualMachine ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20201201_virtualmachine.yaml))
- VirtualMachineScaleSet ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20201201_virtualmachinescaleset.yaml))

Use CRD version `v1alpha1api20201201`


### ARM version 2021-07-01

- Image ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20210701_image.yaml))

Use CRD version `v1alpha1api20210701`

## containerregistry


### ARM version 2021-09-01

- Registry ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerregistry/v1alpha1api20210901_registry.yaml))

Use CRD version `v1alpha1api20210901`

## containerservice


### ARM version 2021-05-01

- ManagedCluster ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerservice/v1alpha1api20210501_managedcluster.yaml))
- ManagedClustersAgentPool ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerservice/v1alpha1api20210501_managedclustersagentpool.yaml))

Use CRD version `v1alpha1api20210501`

## dbformysql


### ARM version 2021-05-01

- FlexibleServer ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbformysql/v1alpha1api20210501_flexibleserver.yaml))
- FlexibleServersDatabase ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbformysql/v1alpha1api20210501_flexibleserversdatabase.yaml))
- FlexibleServersFirewallRule ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbformysql/v1alpha1api20210501_flexibleserversfirewallrule.yaml))

Use CRD version `v1alpha1api20210501`

## dbforpostgresql


### ARM version 2021-06-01

- FlexibleServer ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbforpostgresql/v1alpha1api20210601_flexibleserver.yaml))
- FlexibleServersConfiguration ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbforpostgresql/v1alpha1api20210601_flexibleserversconfiguration.yaml))
- FlexibleServersDatabase ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbforpostgresql/v1alpha1api20210601_flexibleserversdatabase.yaml))
- FlexibleServersFirewallRule ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/dbforpostgresql/v1alpha1api20210601_flexibleserversfirewallrule.yaml))

Use CRD version `v1alpha1api20210601`

## documentdb


### ARM version 2021-05-15

- DatabaseAccount ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_databaseaccount.yaml))
- MongodbDatabase ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_mongodbdatabase.yaml))
- MongodbDatabaseCollection ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_mongodbdatabasecollection.yaml))
- MongodbDatabaseCollectionThroughputSetting ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_mongodbdatabasecollectionthroughputsetting.yaml))
- MongodbDatabaseThroughputSetting ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_mongodbdatabasethroughputsetting.yaml))
- SqlDatabase ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabase.yaml))
- SqlDatabaseContainer ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontainer.yaml))
- SqlDatabaseContainerStoredProcedure ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontainerstoredprocedure.yaml))
- SqlDatabaseContainerThroughputSetting ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontainerthroughputsetting.yaml))
- SqlDatabaseContainerTrigger ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontainertrigger.yaml))
- SqlDatabaseContainerUserDefinedFunction ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasecontaineruserdefinedfunction.yaml))
- SqlDatabaseThroughputSetting ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/documentdb/v1alpha1api20210515_sqldatabasethroughputsetting.yaml))

Use CRD version `v1alpha1api20210515`

## eventgrid


### ARM version 2020-06-01

- Domain ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_domain.yaml))
- DomainsTopic ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_domainstopic.yaml))
- EventSubscription ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_eventsubscription.yaml))
- Topic ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_topic.yaml))

Use CRD version `v1alpha1api20200601`

## eventhub


### ARM version 2021-11-01

- Namespace ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespace.yaml))
- NamespacesAuthorizationRule ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespacesauthorizationrule.yaml))
- NamespacesEventhub ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespaceseventhub.yaml))
- NamespacesEventhubsAuthorizationRule ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespaceseventhubsauthorizationrule.yaml))
- NamespacesEventhubsConsumerGroup ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventhub/v1alpha1api20211101_namespaceseventhubsconsumergroup.yaml))

Use CRD version `v1alpha1api20211101`

## insights


### ARM version 2018-05-01-preview

- Webtest ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/insights/v1alpha1api20180501preview_webtest.yaml))

Use CRD version `v1alpha1api20180501preview`


### ARM version 2020-02-02

- Component ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/insights/v1alpha1api20200202_component.yaml))

Use CRD version `v1alpha1api20200202`

## managedidentity


### ARM version 2018-11-30

- UserAssignedIdentity ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/managedidentity/v1alpha1api20181130_userassignedidentity.yaml))

Use CRD version `v1alpha1api20181130`

## network


### ARM version 2020-11-01

- LoadBalancer ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_loadbalancer.yaml))
- NetworkInterface ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_networkinterface.yaml))
- NetworkSecurityGroup ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_networksecuritygroup.yaml))
- NetworkSecurityGroupsSecurityRule ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_networksecuritygroupssecurityrule.yaml))
- PublicIPAddress ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_publicipaddress.yaml))
- VirtualNetwork ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetwork.yaml))
- VirtualNetworkGateway ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetworkgateway.yaml))
- VirtualNetworksSubnet ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetworkssubnet.yaml))
- VirtualNetworksVirtualNetworkPeering ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetworksvirtualnetworkpeering.yaml))

Use CRD version `v1alpha1api20201101`

## operationalinsights


### ARM version 2021-06-01

- Workspace ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/operationalinsights/v1alpha1api20210601_workspace.yaml))

Use CRD version `v1alpha1api20210601`

## servicebus


### ARM version 2021-01-01-preview

- Namespace ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/servicebus/v1alpha1api20210101preview_namespace.yaml))
- NamespacesQueue ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/servicebus/v1alpha1api20210101preview_namespacesqueue.yaml))
- NamespacesTopic ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/servicebus/v1alpha1api20210101preview_namespacestopic.yaml))

Use CRD version `v1alpha1api20210101preview`

## signalrservice


### ARM version 2021-10-01

- SignalR ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/signalrservice/v1alpha1api20211001_signalr.yaml))

Use CRD version `v1alpha1api20211001`

## storage


### ARM version 2021-04-01

- StorageAccount ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccount.yaml))
- StorageAccountsBlobService ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsblobservice.yaml))
- StorageAccountsBlobServicesContainer ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsblobservicescontainer.yaml))
- StorageAccountsManagementPolicy ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsmanagementpolicy.yaml))
- StorageAccountsQueueService ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsqueueservice.yaml))
- StorageAccountsQueueServicesQueue ([sample](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsqueueservicesqueue.yaml))

Use CRD version `v1alpha1api20210401`

