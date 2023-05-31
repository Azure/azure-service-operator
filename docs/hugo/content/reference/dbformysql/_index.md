---
title: DBforMySQL Supported Resources
linktitle: DBforMySQL
no_list: true
---
Azure Database for MySQL - Single Server is on the retirement path and is [scheduled for retirement by September 16, 2024](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server). We will not be supporting it in ASO v2.

Existing instances of *Single Server* can be migrated to *Azure Database for MySQL - Flexible Server* using the [Azure Database migration Service](https://azure.microsoft.com/en-us/products/database-migration).

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.


| Resource                                                                                                                                                                           | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                        |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| [FlexibleServersAdministrator](https://azure.github.io/azure-service-operator/reference/dbformysql/v1api20220101/#dbformysql.azure.com/v1api20220101.FlexibleServersAdministrator) | 2022-01-01  | v1api20220101 | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api/v1api20220101_flexibleserversadministrator.yaml) |
| [FlexibleServersConfiguration](https://azure.github.io/azure-service-operator/reference/dbformysql/v1api20220101/#dbformysql.azure.com/v1api20220101.FlexibleServersConfiguration) | 2022-01-01  | v1api20220101 | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api/v1api20220101_flexibleserversconfiguration.yaml) |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.


| Resource                                                                                                                                                                         | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                       |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|----------------------------------------------------------------------------------------------------------------------------------------------|
| [FlexibleServer](https://azure.github.io/azure-service-operator/reference/dbformysql/v1api20210501/#dbformysql.azure.com/v1api20210501.FlexibleServer)                           | 2021-05-01  | v1api20210501 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api/v1api20210501_flexibleserver.yaml)              |
| [FlexibleServersDatabase](https://azure.github.io/azure-service-operator/reference/dbformysql/v1api20210501/#dbformysql.azure.com/v1api20210501.FlexibleServersDatabase)         | 2021-05-01  | v1api20210501 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api/v1api20210501_flexibleserversdatabase.yaml)     |
| [FlexibleServersFirewallRule](https://azure.github.io/azure-service-operator/reference/dbformysql/v1api20210501/#dbformysql.azure.com/v1api20210501.FlexibleServersFirewallRule) | 2021-05-01  | v1api20210501 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api/v1api20210501_flexibleserversfirewallrule.yaml) |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                                                                                           | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                         |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| [FlexibleServer](https://azure.github.io/azure-service-operator/reference/dbformysql/v1beta20210501/#dbformysql.azure.com/v1beta20210501.FlexibleServer)                           | 2021-05-01  | v1beta20210501 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1beta/v1beta20210501_flexibleserver.yaml)              |
| [FlexibleServersDatabase](https://azure.github.io/azure-service-operator/reference/dbformysql/v1beta20210501/#dbformysql.azure.com/v1beta20210501.FlexibleServersDatabase)         | 2021-05-01  | v1beta20210501 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1beta/v1beta20210501_flexibleserversdatabase.yaml)     |
| [FlexibleServersFirewallRule](https://azure.github.io/azure-service-operator/reference/dbformysql/v1beta20210501/#dbformysql.azure.com/v1beta20210501.FlexibleServersFirewallRule) | 2021-05-01  | v1beta20210501 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1beta/v1beta20210501_flexibleserversfirewallrule.yaml) |

