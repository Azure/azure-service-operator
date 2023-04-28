---
title: "Supported Resources"
linkTitle: "Supported Resources"
weight: 40
menu:
  main:
    weight: 40
layout: single
cascade:
- type: docs
- render: always
---
These are the resources with Azure Service Operator support committed to our **main** branch,
grouped by the originating ARM service.
(Newly supported resources will appear in this list prior to inclusion in any ASO release.)


## Appconfiguration

Supporting 1 resource: ConfigurationStore

| Resource           | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                      |
|--------------------|-------------|----------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| ConfigurationStore | 2022-05-01  | v1api20220501  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/appconfiguration/v1api/v1api20220501_configurationstore.yaml)   |
| ConfigurationStore | 2022-05-01  | v1beta20220501 | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/appconfiguration/v1beta/v1beta20220501_configurationstore.yaml) |

## Authorization

Supporting 1 resource: RoleAssignment

| Resource       | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                                      |
|----------------|--------------------|-----------------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| RoleAssignment | 2020-08-01-preview | v1api20200801preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/authorization/v1api/v1api20200801preview_roleassignment.yaml)   |
| RoleAssignment | 2020-08-01-preview | v1beta20200801preview | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/authorization/v1beta/v1beta20200801preview_roleassignment.yaml) |

## Batch

Supporting 1 resource: BatchAccount

| Resource                                                                                                                                   | ARM Version | CRD Version    | Supported From | Sample                                                                                                                     |
|--------------------------------------------------------------------------------------------------------------------------------------------|-------------|----------------|----------------|----------------------------------------------------------------------------------------------------------------------------|
| BatchAccount                                                                                                                               | 2021-01-01  | v1api20210101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/batch/v1api/v1api20210101_batchaccount.yaml)   |
| [BatchAccount](https://azure.github.io/azure-service-operator/reference/batch/v1beta20210101/#batch.azure.com/v1beta20210101.BatchAccount) | 2021-01-01  | v1beta20210101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/batch/v1beta/v1beta20210101_batchaccount.yaml) |

## Cache

Supporting 6 resources: Redis, RedisEnterprise, RedisEnterpriseDatabase, RedisFirewallRule, RedisLinkedServer, RedisPatchSchedule

| Resource                | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                |
|-------------------------|-------------|----------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------|
| Redis                   | 2020-12-01  | v1api20201201  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20201201_redis.yaml)                     |
| Redis                   | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20201201_redis.yaml)                   |
| RedisEnterprise         | 2021-03-01  | v1api20210301  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20210301_redisenterprise.yaml)           |
| RedisEnterprise         | 2021-03-01  | v1beta20210301 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20210301_redisenterprise.yaml)         |
| RedisEnterpriseDatabase | 2021-03-01  | v1api20210301  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20210301_redisenterprisedatabase.yaml)   |
| RedisEnterpriseDatabase | 2021-03-01  | v1beta20210301 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20210301_redisenterprisedatabase.yaml) |
| RedisFirewallRule       | 2020-12-01  | v1api20201201  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20201201_redisfirewallrule.yaml)         |
| RedisFirewallRule       | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20201201_redisfirewallrule.yaml)       |
| RedisLinkedServer       | 2020-12-01  | v1api20201201  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20201201_redislinkedserver.yaml)         |
| RedisLinkedServer       | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20201201_redislinkedserver.yaml)       |
| RedisPatchSchedule      | 2020-12-01  | v1api20201201  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20201201_redispatchschedule.yaml)        |
| RedisPatchSchedule      | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20201201_redispatchschedule.yaml)      |

## Cdn

Supporting 2 resources: Profile, ProfilesEndpoint

| Resource         | ARM Version | CRD Version    | Supported From | Sample                                                                                                                       |
|------------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------|
| Profile          | 2021-06-01  | v1api20210601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api/v1api20210601_profile.yaml)            |
| Profile          | 2021-06-01  | v1beta20210601 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1beta/v1beta20210601_profile.yaml)          |
| ProfilesEndpoint | 2021-06-01  | v1api20210601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api/v1api20210601_profilesendpoint.yaml)   |
| ProfilesEndpoint | 2021-06-01  | v1beta20210601 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1beta/v1beta20210601_profilesendpoint.yaml) |

## Compute

Supporting 5 resources: Disk, Image, Snapshot, VirtualMachine, VirtualMachineScaleSet

| Resource               | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                 |
|------------------------|-------------|----------------|----------------|----------------------------------------------------------------------------------------------------------------------------------------|
| Disk                   | 2020-09-30  | v1api20200930  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20200930_disk.yaml)                     |
| Disk                   | 2020-09-30  | v1beta20200930 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1beta/v1beta20200930_disk.yaml)                   |
| Image                  | 2022-03-01  | v1api20220301  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_image.yaml)                    |
| Image                  | 2021-07-01  | v1api20210701  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20210701_image.yaml)                    |
| Image                  | 2022-03-01  | v1beta20220301 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1beta/v1beta20220301_image.yaml)                  |
| Image                  | 2021-07-01  | v1beta20210701 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1beta/v1beta20210701_image.yaml)                  |
| Snapshot               | 2020-09-30  | v1api20200930  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20200930_snapshot.yaml)                 |
| Snapshot               | 2020-09-30  | v1beta20200930 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1beta/v1beta20200930_snapshot.yaml)               |
| VirtualMachine         | 2022-03-01  | v1api20220301  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_virtualmachine.yaml)           |
| VirtualMachine         | 2020-12-01  | v1api20201201  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20201201_virtualmachine.yaml)           |
| VirtualMachine         | 2022-03-01  | v1beta20220301 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1beta/v1beta20220301_virtualmachine.yaml)         |
| VirtualMachine         | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1beta/v1beta20201201_virtualmachine.yaml)         |
| VirtualMachineScaleSet | 2022-03-01  | v1api20220301  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_virtualmachinescaleset.yaml)   |
| VirtualMachineScaleSet | 2020-12-01  | v1api20201201  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20201201_virtualmachinescaleset.yaml)   |
| VirtualMachineScaleSet | 2022-03-01  | v1beta20220301 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1beta/v1beta20220301_virtualmachinescaleset.yaml) |
| VirtualMachineScaleSet | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1beta/v1beta20201201_virtualmachinescaleset.yaml) |

## Containerinstance

Supporting 1 resource: ContainerGroup

| Resource       | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                   |
|----------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------|
| ContainerGroup | 2021-10-01  | v1api20211001  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerinstance/v1api/v1api20211001_containergroup.yaml)   |
| ContainerGroup | 2021-10-01  | v1beta20211001 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerinstance/v1beta/v1beta20211001_containergroup.yaml) |

## Containerregistry

Supporting 1 resource: Registry

| Resource | ARM Version | CRD Version    | Supported From | Sample                                                                                                                             |
|----------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------------|
| Registry | 2021-09-01  | v1api20210901  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerregistry/v1api/v1api20210901_registry.yaml)   |
| Registry | 2021-09-01  | v1beta20210901 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerregistry/v1beta/v1beta20210901_registry.yaml) |

## Containerservice

Supporting 2 resources: ManagedCluster, ManagedClustersAgentPool

| Resource                                                                                                                                                                                 | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                                |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ManagedCluster                                                                                                                                                                           | 2023-02-02-preview | v1api20230202preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20230202preview/v1api20230202preview_managedcluster.yaml)           |
| ManagedCluster                                                                                                                                                                           | 2023-02-01         | v1api20230201        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20230201/v1api20230201_managedcluster.yaml)                         |
| ManagedCluster                                                                                                                                                                           | 2021-05-01         | v1api20210501        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20210501/v1api20210501_managedcluster.yaml)                         |
| [ManagedCluster](https://azure.github.io/azure-service-operator/reference/containerservice/v1beta20210501/#containerservice.azure.com/v1beta20210501.ManagedCluster)                     | 2021-05-01         | v1beta20210501       | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1beta20210501/v1beta20210501_managedcluster.yaml)                       |
| ManagedClustersAgentPool                                                                                                                                                                 | 2023-02-02-preview | v1api20230202preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20230202preview/v1api20230202preview_managedclustersagentpool.yaml) |
| ManagedClustersAgentPool                                                                                                                                                                 | 2023-02-01         | v1api20230201        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20230201/v1api20230201_managedclustersagentpool.yaml)               |
| ManagedClustersAgentPool                                                                                                                                                                 | 2021-05-01         | v1api20210501        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20210501/v1api20210501_managedclustersagentpool.yaml)               |
| [ManagedClustersAgentPool](https://azure.github.io/azure-service-operator/reference/containerservice/v1beta20210501/#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPool) | 2021-05-01         | v1beta20210501       | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1beta20210501/v1beta20210501_managedclustersagentpool.yaml)             |

## Datafactory

Supporting 1 resource: Factory

| Resource | ARM Version | CRD Version   | Supported From | Sample                                                                                                                    |
|----------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------|
| Factory  | 2018-06-01  | v1api20180601 | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/datafactory/v1api/v1api20180601_factory.yaml) |

## Dbformariadb

Supporting 3 resources: Configuration, Database, Server

| Resource      | ARM Version | CRD Version    | Supported From | Sample                                                                                                                             |
|---------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------------|
| Configuration | 2018-06-01  | v1api20180601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformariadb/v1api/v1api20180601_configuration.yaml)   |
| Configuration | 2018-06-01  | v1beta20180601 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformariadb/v1beta/v1beta20180601_configuration.yaml) |
| Database      | 2018-06-01  | v1api20180601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformariadb/v1api/v1api20180601_database.yaml)        |
| Database      | 2018-06-01  | v1beta20180601 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformariadb/v1beta/v1beta20180601_database.yaml)      |
| Server        | 2018-06-01  | v1api20180601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformariadb/v1api/v1api20180601_server.yaml)          |
| Server        | 2018-06-01  | v1beta20180601 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformariadb/v1beta/v1beta20180601_server.yaml)        |

## Dbformysql

Azure Database for MySQL - Single Server is on the retirement path and is [scheduled for retirement by September 16, 2024](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server). We will not be supporting it in ASO v2.

Existing instances of *Single Server* can be migrated to *Azure Database for MySQL - Flexible Server* using the [Azure Database migration Service](https://azure.microsoft.com/en-us/products/database-migration).

Supporting 3 resources: FlexibleServer, FlexibleServersDatabase, FlexibleServersFirewallRule

| Resource                    | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                         |
|-----------------------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| FlexibleServer              | 2021-05-01  | v1api20210501  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api/v1api20210501_flexibleserver.yaml)                |
| FlexibleServer              | 2021-05-01  | v1beta20210501 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1beta/v1beta20210501_flexibleserver.yaml)              |
| FlexibleServersDatabase     | 2021-05-01  | v1api20210501  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api/v1api20210501_flexibleserversdatabase.yaml)       |
| FlexibleServersDatabase     | 2021-05-01  | v1beta20210501 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1beta/v1beta20210501_flexibleserversdatabase.yaml)     |
| FlexibleServersFirewallRule | 2021-05-01  | v1api20210501  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api/v1api20210501_flexibleserversfirewallrule.yaml)   |
| FlexibleServersFirewallRule | 2021-05-01  | v1beta20210501 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1beta/v1beta20210501_flexibleserversfirewallrule.yaml) |

## Dbforpostgresql

Supporting 4 resources: FlexibleServer, FlexibleServersConfiguration, FlexibleServersDatabase, FlexibleServersFirewallRule

| Resource                     | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                                                                     |
|------------------------------|--------------------|-----------------------|----------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| FlexibleServer               | 2022-01-20-preview | v1api20220120preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1api20220120preview/v1api20220120preview_flexibleserver.yaml)                 |
| FlexibleServer               | 2021-06-01         | v1api20210601         | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1api20210601/v1api20210601_flexibleserver.yaml)                               |
| FlexibleServer               | 2022-01-20-preview | v1beta20220120preview | v2.0.0-beta.4  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1beta20220120preview/v1beta20220120preview_flexibleserver.yaml)               |
| FlexibleServer               | 2021-06-01         | v1beta20210601        | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1beta20210601/v1beta20210601_flexibleserver.yaml)                             |
| FlexibleServersConfiguration | 2022-01-20-preview | v1api20220120preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1api20220120preview/v1api20220120preview_flexibleserversconfiguration.yaml)   |
| FlexibleServersConfiguration | 2021-06-01         | v1api20210601         | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1api20210601/v1api20210601_flexibleserversconfiguration.yaml)                 |
| FlexibleServersConfiguration | 2022-01-20-preview | v1beta20220120preview | v2.0.0-beta.4  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1beta20220120preview/v1beta20220120preview_flexibleserversconfiguration.yaml) |
| FlexibleServersConfiguration | 2021-06-01         | v1beta20210601        | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1beta20210601/v1beta20210601_flexibleserversconfiguration.yaml)               |
| FlexibleServersDatabase      | 2022-01-20-preview | v1api20220120preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1api20220120preview/v1api20220120preview_flexibleserversdatabase.yaml)        |
| FlexibleServersDatabase      | 2021-06-01         | v1api20210601         | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1api20210601/v1api20210601_flexibleserversdatabase.yaml)                      |
| FlexibleServersDatabase      | 2022-01-20-preview | v1beta20220120preview | v2.0.0-beta.4  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1beta20220120preview/v1beta20220120preview_flexibleserversdatabase.yaml)      |
| FlexibleServersDatabase      | 2021-06-01         | v1beta20210601        | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1beta20210601/v1beta20210601_flexibleserversdatabase.yaml)                    |
| FlexibleServersFirewallRule  | 2022-01-20-preview | v1api20220120preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1api20220120preview/v1api20220120preview_flexibleserversfirewallrule.yaml)    |
| FlexibleServersFirewallRule  | 2021-06-01         | v1api20210601         | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1api20210601/v1api20210601_flexibleserversfirewallrule.yaml)                  |
| FlexibleServersFirewallRule  | 2022-01-20-preview | v1beta20220120preview | v2.0.0-beta.4  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1beta20220120preview/v1beta20220120preview_flexibleserversfirewallrule.yaml)  |
| FlexibleServersFirewallRule  | 2021-06-01         | v1beta20210601        | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbforpostgresql/v1beta20210601/v1beta20210601_flexibleserversfirewallrule.yaml)                |

## Documentdb

Supporting 13 resources: DatabaseAccount, MongodbDatabase, MongodbDatabaseCollection, MongodbDatabaseCollectionThroughputSetting, MongodbDatabaseThroughputSetting, SqlDatabase, SqlDatabaseContainer, SqlDatabaseContainerStoredProcedure, SqlDatabaseContainerThroughputSetting, SqlDatabaseContainerTrigger, SqlDatabaseContainerUserDefinedFunction, SqlDatabaseThroughputSetting, SqlRoleAssignment

| Resource                                   | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                                                 |
|--------------------------------------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| DatabaseAccount                            | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1api/v1api20210515_databaseaccount.yaml)                               |
| DatabaseAccount                            | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1beta/v1beta20210515_databaseaccount.yaml)                             |
| MongodbDatabase                            | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1api/v1api20210515_mongodbdatabase.yaml)                               |
| MongodbDatabase                            | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1beta/v1beta20210515_mongodbdatabase.yaml)                             |
| MongodbDatabaseCollection                  | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1api/v1api20210515_mongodbdatabasecollection.yaml)                     |
| MongodbDatabaseCollection                  | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1beta/v1beta20210515_mongodbdatabasecollection.yaml)                   |
| MongodbDatabaseCollectionThroughputSetting | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1api/v1api20210515_mongodbdatabasecollectionthroughputsetting.yaml)    |
| MongodbDatabaseCollectionThroughputSetting | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1beta/v1beta20210515_mongodbdatabasecollectionthroughputsetting.yaml)  |
| MongodbDatabaseThroughputSetting           | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1api/v1api20210515_mongodbdatabasethroughputsetting.yaml)              |
| MongodbDatabaseThroughputSetting           | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/mongodb/v1beta/v1beta20210515_mongodbdatabasethroughputsetting.yaml)            |
| SqlDatabase                                | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1api/v1api20210515_sqldatabase.yaml)                               |
| SqlDatabase                                | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1beta/v1beta20210515_sqldatabase.yaml)                             |
| SqlDatabaseContainer                       | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1api/v1api20210515_sqldatabasecontainer.yaml)                      |
| SqlDatabaseContainer                       | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1beta/v1beta20210515_sqldatabasecontainer.yaml)                    |
| SqlDatabaseContainerStoredProcedure        | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1api/v1api20210515_sqldatabasecontainerstoredprocedure.yaml)       |
| SqlDatabaseContainerStoredProcedure        | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1beta/v1beta20210515_sqldatabasecontainerstoredprocedure.yaml)     |
| SqlDatabaseContainerThroughputSetting      | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1api/v1api20210515_sqldatabasecontainerthroughputsetting.yaml)     |
| SqlDatabaseContainerThroughputSetting      | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1beta/v1beta20210515_sqldatabasecontainerthroughputsetting.yaml)   |
| SqlDatabaseContainerTrigger                | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1api/v1api20210515_sqldatabasecontainertrigger.yaml)               |
| SqlDatabaseContainerTrigger                | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1beta/v1beta20210515_sqldatabasecontainertrigger.yaml)             |
| SqlDatabaseContainerUserDefinedFunction    | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1api/v1api20210515_sqldatabasecontaineruserdefinedfunction.yaml)   |
| SqlDatabaseContainerUserDefinedFunction    | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1beta/v1beta20210515_sqldatabasecontaineruserdefinedfunction.yaml) |
| SqlDatabaseThroughputSetting               | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1api/v1api20210515_sqldatabasethroughputsetting.yaml)              |
| SqlDatabaseThroughputSetting               | 2021-05-15  | v1beta20210515 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1beta/v1beta20210515_sqldatabasethroughputsetting.yaml)            |
| SqlRoleAssignment                          | 2021-05-15  | v1api20210515  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1api/v1api20210515_sqlroleassignment.yaml)                         |
| SqlRoleAssignment                          | 2021-05-15  | v1beta20210515 | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/sqldatabase/v1beta/v1beta20210515_sqlroleassignment.yaml)                       |

## Eventgrid

Supporting 4 resources: Domain, DomainsTopic, EventSubscription, Topic

| Resource          | ARM Version | CRD Version    | Supported From | Sample                                                                                                                              |
|-------------------|-------------|----------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------|
| Domain            | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1api/v1api20200601_domain.yaml)              |
| Domain            | 2020-06-01  | v1beta20200601 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1beta/v1beta20200601_domain.yaml)            |
| DomainsTopic      | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1api/v1api20200601_domainstopic.yaml)        |
| DomainsTopic      | 2020-06-01  | v1beta20200601 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1beta/v1beta20200601_domainstopic.yaml)      |
| EventSubscription | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1api/v1api20200601_eventsubscription.yaml)   |
| EventSubscription | 2020-06-01  | v1beta20200601 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1beta/v1beta20200601_eventsubscription.yaml) |
| Topic             | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1api/v1api20200601_topic.yaml)               |
| Topic             | 2020-06-01  | v1beta20200601 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1beta/v1beta20200601_topic.yaml)             |

## Eventhub

Supporting 5 resources: Namespace, NamespacesAuthorizationRule, NamespacesEventhub, NamespacesEventhubsAuthorizationRule, NamespacesEventhubsConsumerGroup

| Resource                             | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                                |
|--------------------------------------|-------------|----------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| Namespace                            | 2021-11-01  | v1api20211101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api/v1api20211101_namespace.yaml)                              |
| Namespace                            | 2021-11-01  | v1beta20211101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1beta/v1beta20211101_namespace.yaml)                            |
| NamespacesAuthorizationRule          | 2021-11-01  | v1api20211101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api/v1api20211101_namespacesauthorizationrule.yaml)            |
| NamespacesAuthorizationRule          | 2021-11-01  | v1beta20211101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1beta/v1beta20211101_namespacesauthorizationrule.yaml)          |
| NamespacesEventhub                   | 2021-11-01  | v1api20211101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api/v1api20211101_namespaceseventhub.yaml)                     |
| NamespacesEventhub                   | 2021-11-01  | v1beta20211101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1beta/v1beta20211101_namespaceseventhub.yaml)                   |
| NamespacesEventhubsAuthorizationRule | 2021-11-01  | v1api20211101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api/v1api20211101_namespaceseventhubsauthorizationrule.yaml)   |
| NamespacesEventhubsAuthorizationRule | 2021-11-01  | v1beta20211101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1beta/v1beta20211101_namespaceseventhubsauthorizationrule.yaml) |
| NamespacesEventhubsConsumerGroup     | 2021-11-01  | v1api20211101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api/v1api20211101_namespaceseventhubsconsumergroup.yaml)       |
| NamespacesEventhubsConsumerGroup     | 2021-11-01  | v1beta20211101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1beta/v1beta20211101_namespaceseventhubsconsumergroup.yaml)     |

## Insights

Supporting 2 resources: Component, Webtest

| Resource  | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                          |
|-----------|--------------------|-----------------------|----------------|---------------------------------------------------------------------------------------------------------------------------------|
| Component | 2020-02-02         | v1api20200202         | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20200202_component.yaml)        |
| Component | 2020-02-02         | v1beta20200202        | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1beta/v1beta20200202_component.yaml)      |
| Webtest   | 2018-05-01-preview | v1api20180501preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20180501preview_webtest.yaml)   |
| Webtest   | 2018-05-01-preview | v1beta20180501preview | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1beta/v1beta20180501preview_webtest.yaml) |

## Keyvault

Supporting 1 resource: Vault

| Resource | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                        |
|----------|--------------------|-----------------------|----------------|-------------------------------------------------------------------------------------------------------------------------------|
| Vault    | 2021-04-01-preview | v1api20210401preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/keyvault/v1api/v1api20210401preview_vault.yaml)   |
| Vault    | 2021-04-01-preview | v1beta20210401preview | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/keyvault/v1beta/v1beta20210401preview_vault.yaml) |

## Machinelearningservices

Supporting 3 resources: Workspace, WorkspacesCompute, WorkspacesConnection

| Resource             | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                               |
|----------------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| Workspace            | 2021-07-01  | v1api20210701  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api/v1api20210701_workspace.yaml)              |
| Workspace            | 2021-07-01  | v1beta20210701 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1beta/v1beta20210701_workspace.yaml)            |
| WorkspacesCompute    | 2021-07-01  | v1api20210701  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api/v1api20210701_workspacescompute.yaml)      |
| WorkspacesCompute    | 2021-07-01  | v1beta20210701 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1beta/v1beta20210701_workspacescompute.yaml)    |
| WorkspacesConnection | 2021-07-01  | v1api20210701  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api/v1api20210701_workspacesconnection.yaml)   |
| WorkspacesConnection | 2021-07-01  | v1beta20210701 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1beta/v1beta20210701_workspacesconnection.yaml) |

## Managedidentity

Supporting 2 resources: FederatedIdentityCredential, UserAssignedIdentity

| Resource                    | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                                                     |
|-----------------------------|--------------------|-----------------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| FederatedIdentityCredential | 2022-01-31-preview | v1api20220131preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api/v1api20220131preview_federatedidentitycredential.yaml)   |
| FederatedIdentityCredential | 2022-01-31-preview | v1beta20220131preview | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1beta/v1beta20220131preview_federatedidentitycredential.yaml) |
| UserAssignedIdentity        | 2018-11-30         | v1api20181130         | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api/v1api20181130_userassignedidentity.yaml)                 |
| UserAssignedIdentity        | 2018-11-30         | v1beta20181130        | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1beta/v1beta20181130_userassignedidentity.yaml)               |

## Network

Supporting 26 resources: BastionHost, LoadBalancer, NatGateway, NetworkInterface, NetworkSecurityGroup, NetworkSecurityGroupsSecurityRule, PrivateDnsZone, PrivateDnsZonesAAAARecord, PrivateDnsZonesARecord, PrivateDnsZonesCNAMERecord, PrivateDnsZonesMXRecord, PrivateDnsZonesPTRRecord, PrivateDnsZonesSRVRecord, PrivateDnsZonesTXTRecord, PrivateDnsZonesVirtualNetworkLink, PrivateEndpoint, PrivateEndpointsPrivateDnsZoneGroup, PrivateLinkService, PublicIPAddress, PublicIPPrefix, RouteTable, RouteTablesRoute, VirtualNetwork, VirtualNetworkGateway, VirtualNetworksSubnet, VirtualNetworksVirtualNetworkPeering

| Resource                             | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                                       |
|--------------------------------------|-------------|----------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| BastionHost                          | 2022-07-01  | v1api20220701  | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20220701/v1api20220701_bastionhost.yaml)                            |
| LoadBalancer                         | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_loadbalancer.yaml)                           |
| LoadBalancer                         | 2020-11-01  | v1beta20201101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_loadbalancer.yaml)                         |
| NatGateway                           | 2022-07-01  | v1api20220701  | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20220701/v1api20220701_natgateway.yaml)                             |
| NetworkInterface                     | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_networkinterface.yaml)                       |
| NetworkInterface                     | 2020-11-01  | v1beta20201101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_networkinterface.yaml)                     |
| NetworkSecurityGroup                 | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_networksecuritygroup.yaml)                   |
| NetworkSecurityGroup                 | 2020-11-01  | v1beta20201101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_networksecuritygroup.yaml)                 |
| NetworkSecurityGroupsSecurityRule    | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_networksecuritygroupssecurityrule.yaml)      |
| NetworkSecurityGroupsSecurityRule    | 2020-11-01  | v1beta20201101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_networksecuritygroupssecurityrule.yaml)    |
| PrivateDnsZone                       | 2018-09-01  | v1api20180901  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20180901/v1api20180901_privatednszone.yaml)                         |
| PrivateDnsZone                       | 2018-09-01  | v1beta20180901 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20180901/v1beta20180901_privatednszone.yaml)                       |
| PrivateDnsZonesAAAARecord            | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20200601/v1api20200601_privatednszonesaaaarecord.yaml)              |
| PrivateDnsZonesARecord               | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20200601/v1api20200601_privatednszonesarecord.yaml)                 |
| PrivateDnsZonesCNAMERecord           | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20200601/v1api20200601_privatednszonescnamerecord.yaml)             |
| PrivateDnsZonesMXRecord              | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20200601/v1api20200601_privatednszonesmxrecord.yaml)                |
| PrivateDnsZonesPTRRecord             | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20200601/v1api20200601_privatednszonesptrrecord.yaml)               |
| PrivateDnsZonesSRVRecord             | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20200601/v1api20200601_privatednszonessrvrecord.yaml)               |
| PrivateDnsZonesTXTRecord             | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20200601/v1api20200601_privatednszonestxtrecord.yaml)               |
| PrivateDnsZonesVirtualNetworkLink    | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20200601/v1api20200601_privatednszonesvirtualnetworklink.yaml)      |
| PrivateEndpoint                      | 2022-07-01  | v1api20220701  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20220701/v1api20220701_privateendpoint.yaml)                        |
| PrivateEndpointsPrivateDnsZoneGroup  | 2022-07-01  | v1api20220701  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20220701/v1api20220701_privateendpointsprivatednszonegroup.yaml)    |
| PrivateLinkService                   | 2022-07-01  | v1api20220701  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20220701/v1api20220701_privatelinkservice.yaml)                     |
| PublicIPAddress                      | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_publicipaddress.yaml)                        |
| PublicIPAddress                      | 2020-11-01  | v1beta20201101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_publicipaddress.yaml)                      |
| PublicIPPrefix                       | 2022-07-01  | v1api20220701  | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20220701/v1api20220701_publicipprefix.yaml)                         |
| RouteTable                           | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_routetable.yaml)                             |
| RouteTable                           | 2020-11-01  | v1beta20201101 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_routetable.yaml)                           |
| RouteTablesRoute                     | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_routetablesroute.yaml)                       |
| RouteTablesRoute                     | 2020-11-01  | v1beta20201101 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_routetablesroute.yaml)                     |
| VirtualNetwork                       | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_virtualnetwork.yaml)                         |
| VirtualNetwork                       | 2020-11-01  | v1beta20201101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_virtualnetwork.yaml)                       |
| VirtualNetworkGateway                | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_virtualnetworkgateway.yaml)                  |
| VirtualNetworkGateway                | 2020-11-01  | v1beta20201101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_virtualnetworkgateway.yaml)                |
| VirtualNetworksSubnet                | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_virtualnetworkssubnet.yaml)                  |
| VirtualNetworksSubnet                | 2020-11-01  | v1beta20201101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_virtualnetworkssubnet.yaml)                |
| VirtualNetworksVirtualNetworkPeering | 2020-11-01  | v1api20201101  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1api20201101/v1api20201101_virtualnetworksvirtualnetworkpeering.yaml)   |
| VirtualNetworksVirtualNetworkPeering | 2020-11-01  | v1beta20201101 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network/v1beta20201101/v1beta20201101_virtualnetworksvirtualnetworkpeering.yaml) |

## Operationalinsights

Supporting 1 resource: Workspace

| Resource  | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                |
|-----------|-------------|----------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------|
| Workspace | 2021-06-01  | v1api20210601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/synapse/v1api/v1api20210601_workspace.yaml)               |
| Workspace | 2021-06-01  | v1beta20210601 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/operationalinsights/v1beta/v1beta20210601_workspace.yaml) |

## Resources

Supporting 1 resource: ResourceGroup

| Resource      | ARM Version | CRD Version    | Supported From | Sample                                                                                                                          |
|---------------|-------------|----------------|----------------|---------------------------------------------------------------------------------------------------------------------------------|
| ResourceGroup | 2020-06-01  | v1api20200601  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/resources/v1api/v1api20200601_resourcegroup.yaml)   |
| ResourceGroup | 2020-06-01  | v1beta20200601 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/resources/v1beta/v1beta20200601_resourcegroup.yaml) |

## Servicebus

Supporting 5 resources: Namespace, NamespacesQueue, NamespacesTopic, NamespacesTopicsSubscription, NamespacesTopicsSubscriptionsRule

| Resource                          | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                                                      |
|-----------------------------------|--------------------|-----------------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Namespace                         | 2021-01-01-preview | v1api20210101preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespace.yaml)                           |
| Namespace                         | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespace.yaml)                         |
| NamespacesQueue                   | 2021-01-01-preview | v1api20210101preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespacesqueue.yaml)                     |
| NamespacesQueue                   | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespacesqueue.yaml)                   |
| NamespacesTopic                   | 2021-01-01-preview | v1api20210101preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespacestopic.yaml)                     |
| NamespacesTopic                   | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespacestopic.yaml)                   |
| NamespacesTopicsSubscription      | 2021-01-01-preview | v1api20210101preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespacestopicssubscription.yaml)        |
| NamespacesTopicsSubscription      | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespacestopicssubscription.yaml)      |
| NamespacesTopicsSubscriptionsRule | 2021-01-01-preview | v1api20210101preview  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespacestopicssubscriptionsrule.yaml)   |
| NamespacesTopicsSubscriptionsRule | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespacestopicssubscriptionsrule.yaml) |

## Signalrservice

Supporting 1 resource: SignalR

| Resource                                                                                                                                           | ARM Version | CRD Version    | Supported From | Sample                                                                                                                         |
|----------------------------------------------------------------------------------------------------------------------------------------------------|-------------|----------------|----------------|--------------------------------------------------------------------------------------------------------------------------------|
| SignalR                                                                                                                                            | 2021-10-01  | v1api20211001  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/signalrservice/v1api/v1api20211001_signalr.yaml)   |
| [SignalR](https://azure.github.io/azure-service-operator/reference/signalrservice/v1beta20211001/#signalrservice.azure.com/v1beta20211001.SignalR) | 2021-10-01  | v1beta20211001 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/signalrservice/v1beta/v1beta20211001_signalr.yaml) |

## Sql

Supporting 22 resources: Server, ServersAdministrator, ServersAdvancedThreatProtectionSetting, ServersAuditingSetting, ServersAzureADOnlyAuthentication, ServersConnectionPolicy, ServersDatabase, ServersDatabasesAdvancedThreatProtectionSetting, ServersDatabasesAuditingSetting, ServersDatabasesBackupLongTermRetentionPolicy, ServersDatabasesBackupShortTermRetentionPolicy, ServersDatabasesSecurityAlertPolicy, ServersDatabasesTransparentDataEncryption, ServersDatabasesVulnerabilityAssessment, ServersElasticPool, ServersFailoverGroup, ServersFirewallRule, ServersIPV6FirewallRule, ServersOutboundFirewallRule, ServersSecurityAlertPolicy, ServersVirtualNetworkRule, ServersVulnerabilityAssessment

| Resource                                        | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                    |
|-------------------------------------------------|-------------|---------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| Server                                          | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_server.yaml)                                          |
| ServersAdministrator                            | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversadministrator.yaml)                            |
| ServersAdvancedThreatProtectionSetting          | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversadvancedthreatprotectionsetting.yaml)          |
| ServersAuditingSetting                          | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversauditingsetting.yaml)                          |
| ServersAzureADOnlyAuthentication                | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversazureadonlyauthentication.yaml)                |
| ServersConnectionPolicy                         | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversconnectionpolicy.yaml)                         |
| ServersDatabase                                 | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversdatabase.yaml)                                 |
| ServersDatabasesAdvancedThreatProtectionSetting | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversdatabasesadvancedthreatprotectionsetting.yaml) |
| ServersDatabasesAuditingSetting                 | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversdatabasesauditingsetting.yaml)                 |
| ServersDatabasesBackupLongTermRetentionPolicy   | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversdatabasesbackuplongtermretentionpolicy.yaml)   |
| ServersDatabasesBackupShortTermRetentionPolicy  | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversdatabasesbackupshorttermretentionpolicy.yaml)  |
| ServersDatabasesSecurityAlertPolicy             | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversdatabasessecurityalertpolicy.yaml)             |
| ServersDatabasesTransparentDataEncryption       | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversdatabasestransparentdataencryption.yaml)       |
| ServersDatabasesVulnerabilityAssessment         | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversdatabasesvulnerabilityassessment.yaml)         |
| ServersElasticPool                              | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serverselasticpool.yaml)                              |
| ServersFailoverGroup                            | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversfailovergroup.yaml)                            |
| ServersFirewallRule                             | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversfirewallrule.yaml)                             |
| ServersIPV6FirewallRule                         | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversipv6firewallrule.yaml)                         |
| ServersOutboundFirewallRule                     | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversoutboundfirewallrule.yaml)                     |
| ServersSecurityAlertPolicy                      | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serverssecurityalertpolicy.yaml)                      |
| ServersVirtualNetworkRule                       | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversvirtualnetworkrule.yaml)                       |
| ServersVulnerabilityAssessment                  | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/sql/v1api/v1api20211101_serversvulnerabilityassessment.yaml)                  |

## Storage

Supporting 6 resources: StorageAccount, StorageAccountsBlobService, StorageAccountsBlobServicesContainer, StorageAccountsManagementPolicy, StorageAccountsQueueService, StorageAccountsQueueServicesQueue

| Resource                                                                                                                                                                                       | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                                           |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| StorageAccount                                                                                                                                                                                 | 2021-04-01  | v1api20210401  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/queueservice/v1api/v1api20210401_storageaccount.yaml)                        |
| [StorageAccount](https://azure.github.io/azure-service-operator/reference/storage/v1beta20210401/#storage.azure.com/v1beta20210401.StorageAccount)                                             | 2021-04-01  | v1beta20210401 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/queueservice/v1beta/v1beta20210401_storageaccount.yaml)                      |
| StorageAccountsBlobService                                                                                                                                                                     | 2021-04-01  | v1api20210401  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/blobservice/v1api/v1api20210401_storageaccountsblobservice.yaml)             |
| [StorageAccountsBlobService](https://azure.github.io/azure-service-operator/reference/storage/v1beta20210401/#storage.azure.com/v1beta20210401.StorageAccountsBlobService)                     | 2021-04-01  | v1beta20210401 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/blobservice/v1beta/v1beta20210401_storageaccountsblobservice.yaml)           |
| StorageAccountsBlobServicesContainer                                                                                                                                                           | 2021-04-01  | v1api20210401  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/blobservice/v1api/v1api20210401_storageaccountsblobservicescontainer.yaml)   |
| [StorageAccountsBlobServicesContainer](https://azure.github.io/azure-service-operator/reference/storage/v1beta20210401/#storage.azure.com/v1beta20210401.StorageAccountsBlobServicesContainer) | 2021-04-01  | v1beta20210401 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/blobservice/v1beta/v1beta20210401_storageaccountsblobservicescontainer.yaml) |
| StorageAccountsManagementPolicy                                                                                                                                                                | 2021-04-01  | v1api20210401  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/blobservice/v1api/v1api20210401_storageaccountsmanagementpolicy.yaml)        |
| [StorageAccountsManagementPolicy](https://azure.github.io/azure-service-operator/reference/storage/v1beta20210401/#storage.azure.com/v1beta20210401.StorageAccountsManagementPolicy)           | 2021-04-01  | v1beta20210401 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/blobservice/v1beta/v1beta20210401_storageaccountsmanagementpolicy.yaml)      |
| StorageAccountsQueueService                                                                                                                                                                    | 2021-04-01  | v1api20210401  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/queueservice/v1api/v1api20210401_storageaccountsqueueservice.yaml)           |
| [StorageAccountsQueueService](https://azure.github.io/azure-service-operator/reference/storage/v1beta20210401/#storage.azure.com/v1beta20210401.StorageAccountsQueueService)                   | 2021-04-01  | v1beta20210401 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/queueservice/v1beta/v1beta20210401_storageaccountsqueueservice.yaml)         |
| StorageAccountsQueueServicesQueue                                                                                                                                                              | 2021-04-01  | v1api20210401  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/queueservice/v1api/v1api20210401_storageaccountsqueueservicesqueue.yaml)     |
| [StorageAccountsQueueServicesQueue](https://azure.github.io/azure-service-operator/reference/storage/v1beta20210401/#storage.azure.com/v1beta20210401.StorageAccountsQueueServicesQueue)       | 2021-04-01  | v1beta20210401 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/storage/queueservice/v1beta/v1beta20210401_storageaccountsqueueservicesqueue.yaml)   |

## Subscription

Supporting 1 resource: Alias

| Resource                                                                                                                                   | ARM Version | CRD Version    | Supported From | Sample                                                                                                                     |
|--------------------------------------------------------------------------------------------------------------------------------------------|-------------|----------------|----------------|----------------------------------------------------------------------------------------------------------------------------|
| Alias                                                                                                                                      | 2021-10-01  | v1api20211001  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/subscription/v1api/v1api20211001_alias.yaml)   |
| [Alias](https://azure.github.io/azure-service-operator/reference/subscription/v1beta20211001/#subscription.azure.com/v1beta20211001.Alias) | 2021-10-01  | v1beta20211001 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/subscription/v1beta/v1beta20211001_alias.yaml) |

## Synapse

Supporting 2 resources: Workspace, WorkspacesBigDataPool

| Resource              | ARM Version | CRD Version   | Supported From | Sample                                                                                                                              |
|-----------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------|
| Workspace             | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/synapse/v1api/v1api20210601_workspace.yaml)             |
| WorkspacesBigDataPool | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/synapse/v1api/v1api20210601_workspacesbigdatapool.yaml) |

## Web

Supporting 2 resources: ServerFarm, Site

| Resource   | ARM Version | CRD Version    | Supported From | Sample                                                                                                                 |
|------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------|
| ServerFarm | 2022-03-01  | v1api20220301  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/web/v1api/v1api20220301_serverfarm.yaml)   |
| ServerFarm | 2022-03-01  | v1beta20220301 | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/web/v1beta/v1beta20220301_serverfarm.yaml) |
| Site       | 2022-03-01  | v1api20220301  | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/web/v1api/v1api20220301_site.yaml)         |
| Site       | 2022-03-01  | v1beta20220301 | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/web/v1beta/v1beta20220301_site.yaml)       |

