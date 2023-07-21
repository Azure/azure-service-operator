---
title: Cache Supported Resources
linktitle: Cache
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `cache.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                       | ARM Version | CRD Version   | Supported From | Sample                                                                                                                              |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------|
| [Redis](https://azure.github.io/azure-service-operator/reference/cache/v1api20201201/#cache.azure.com/v1api20201201.Redis)                                     | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20201201_redis.yaml)                   |
| [RedisEnterprise](https://azure.github.io/azure-service-operator/reference/cache/v1api20210301/#cache.azure.com/v1api20210301.RedisEnterprise)                 | 2021-03-01  | v1api20210301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20210301_redisenterprise.yaml)         |
| [RedisEnterpriseDatabase](https://azure.github.io/azure-service-operator/reference/cache/v1api20210301/#cache.azure.com/v1api20210301.RedisEnterpriseDatabase) | 2021-03-01  | v1api20210301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20210301_redisenterprisedatabase.yaml) |
| [RedisFirewallRule](https://azure.github.io/azure-service-operator/reference/cache/v1api20201201/#cache.azure.com/v1api20201201.RedisFirewallRule)             | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20201201_redisfirewallrule.yaml)       |
| [RedisLinkedServer](https://azure.github.io/azure-service-operator/reference/cache/v1api20201201/#cache.azure.com/v1api20201201.RedisLinkedServer)             | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20201201_redislinkedserver.yaml)       |
| [RedisPatchSchedule](https://azure.github.io/azure-service-operator/reference/cache/v1api20201201/#cache.azure.com/v1api20201201.RedisPatchSchedule)           | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1api/v1api20201201_redispatchschedule.yaml)      |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                                                                         | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|----------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------|
| [Redis](https://azure.github.io/azure-service-operator/reference/cache/v1beta20201201/#cache.azure.com/v1beta20201201.Redis)                                     | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20201201_redis.yaml)                   |
| [RedisEnterprise](https://azure.github.io/azure-service-operator/reference/cache/v1beta20210301/#cache.azure.com/v1beta20210301.RedisEnterprise)                 | 2021-03-01  | v1beta20210301 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20210301_redisenterprise.yaml)         |
| [RedisEnterpriseDatabase](https://azure.github.io/azure-service-operator/reference/cache/v1beta20210301/#cache.azure.com/v1beta20210301.RedisEnterpriseDatabase) | 2021-03-01  | v1beta20210301 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20210301_redisenterprisedatabase.yaml) |
| [RedisFirewallRule](https://azure.github.io/azure-service-operator/reference/cache/v1beta20201201/#cache.azure.com/v1beta20201201.RedisFirewallRule)             | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20201201_redisfirewallrule.yaml)       |
| [RedisLinkedServer](https://azure.github.io/azure-service-operator/reference/cache/v1beta20201201/#cache.azure.com/v1beta20201201.RedisLinkedServer)             | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20201201_redislinkedserver.yaml)       |
| [RedisPatchSchedule](https://azure.github.io/azure-service-operator/reference/cache/v1beta20201201/#cache.azure.com/v1beta20201201.RedisPatchSchedule)           | 2020-12-01  | v1beta20201201 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cache/v1beta/v1beta20201201_redispatchschedule.yaml)      |

