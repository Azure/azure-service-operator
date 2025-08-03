---
title: Cache Supported Resources
linktitle: Cache
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `cache.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                                                                                           | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                      |
|----------------------------------------------------------------------------------------------------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| [RedisEnterprise]({{< relref "/reference/cache/v1api20250401#RedisEnterprise" >}})                 | 2025-04-01  | v1api20250401 | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20250401/v1api20250401_redisenterprise.yaml)         |
| [RedisEnterpriseDatabase]({{< relref "/reference/cache/v1api20250401#RedisEnterpriseDatabase" >}}) | 2025-04-01  | v1api20250401 | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20250401/v1api20250401_redisenterprisedatabase.yaml) |

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                           | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                      |
|----------------------------------------------------------------------------------------------------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| [Redis]({{< relref "/reference/cache/v1api20230801#Redis" >}})                                     | 2023-08-01  | v1api20230801 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230801/v1api20230801_redis.yaml)                   |
| [RedisEnterprise]({{< relref "/reference/cache/v1api20230701#RedisEnterprise" >}})                 | 2023-07-01  | v1api20230701 | v2.3.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230701/v1api20230701_redisenterprise.yaml)         |
| [RedisEnterpriseDatabase]({{< relref "/reference/cache/v1api20230701#RedisEnterpriseDatabase" >}}) | 2023-07-01  | v1api20230701 | v2.3.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230701/v1api20230701_redisenterprisedatabase.yaml) |
| [RedisFirewallRule]({{< relref "/reference/cache/v1api20230801#RedisFirewallRule" >}})             | 2023-08-01  | v1api20230801 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230801/v1api20230801_redisfirewallrule.yaml)       |
| [RedisLinkedServer]({{< relref "/reference/cache/v1api20230801#RedisLinkedServer" >}})             | 2023-08-01  | v1api20230801 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230801/v1api20230801_redislinkedserver.yaml)       |
| [RedisPatchSchedule]({{< relref "/reference/cache/v1api20230801#RedisPatchSchedule" >}})           | 2023-08-01  | v1api20230801 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230801/v1api20230801_redispatchschedule.yaml)      |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                           | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                      |
|----------------------------------------------------------------------------------------------------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| [Redis]({{< relref "/reference/cache/v1api20230401#Redis" >}})                                     | 2023-04-01  | v1api20230401 | v2.3.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230401/v1api20230401_redis.yaml)                   |
| [Redis]({{< relref "/reference/cache/v1api20201201#Redis" >}})                                     | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20201201/v1api20201201_redis.yaml)                   |
| [RedisEnterprise]({{< relref "/reference/cache/v1api20210301#RedisEnterprise" >}})                 | 2021-03-01  | v1api20210301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20210301/v1api20210301_redisenterprise.yaml)         |
| [RedisEnterpriseDatabase]({{< relref "/reference/cache/v1api20210301#RedisEnterpriseDatabase" >}}) | 2021-03-01  | v1api20210301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20210301/v1api20210301_redisenterprisedatabase.yaml) |
| [RedisFirewallRule]({{< relref "/reference/cache/v1api20230401#RedisFirewallRule" >}})             | 2023-04-01  | v1api20230401 | v2.3.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230401/v1api20230401_redisfirewallrule.yaml)       |
| [RedisFirewallRule]({{< relref "/reference/cache/v1api20201201#RedisFirewallRule" >}})             | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20201201/v1api20201201_redisfirewallrule.yaml)       |
| [RedisLinkedServer]({{< relref "/reference/cache/v1api20230401#RedisLinkedServer" >}})             | 2023-04-01  | v1api20230401 | v2.3.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230401/v1api20230401_redislinkedserver.yaml)       |
| [RedisLinkedServer]({{< relref "/reference/cache/v1api20201201#RedisLinkedServer" >}})             | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20201201/v1api20201201_redislinkedserver.yaml)       |
| [RedisPatchSchedule]({{< relref "/reference/cache/v1api20230401#RedisPatchSchedule" >}})           | 2023-04-01  | v1api20230401 | v2.3.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20230401/v1api20230401_redispatchschedule.yaml)      |
| [RedisPatchSchedule]({{< relref "/reference/cache/v1api20201201#RedisPatchSchedule" >}})           | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api20201201/v1api20201201_redispatchschedule.yaml)      |

