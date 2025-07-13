---
title: DBforMySQL Supported Resources
linktitle: DBforMySQL
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `dbformysql.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

Azure Database for MySQL - Single Server is on the retirement path and is [scheduled for retirement by September 16, 2024](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server). We will not be supporting it in ASO v2.

Existing instances of *Single Server* can be migrated to *Azure Database for MySQL - Flexible Server* using the [Azure Database migration Service](https://azure.microsoft.com/en-us/products/database-migration).

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                          | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                |
|-------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| [FlexibleServer]({{< relref "/reference/dbformysql/v1api20231230#FlexibleServer" >}})                             | 2023-12-30  | v1api20231230 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20231230/v1api20231230_flexibleserver.yaml)               |
| [FlexibleServersAdministrator]({{< relref "/reference/dbformysql/v1api20231230#FlexibleServersAdministrator" >}}) | 2023-12-30  | v1api20231230 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20231230/v1api20231230_flexibleserversadministrator.yaml) |
| [FlexibleServersConfiguration]({{< relref "/reference/dbformysql/v1api20231230#FlexibleServersConfiguration" >}}) | 2023-12-30  | v1api20231230 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20231230/v1api20231230_flexibleserversconfiguration.yaml) |
| [FlexibleServersDatabase]({{< relref "/reference/dbformysql/v1api20231230#FlexibleServersDatabase" >}})           | 2023-12-30  | v1api20231230 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20231230/v1api20231230_flexibleserversdatabase.yaml)      |
| [FlexibleServersFirewallRule]({{< relref "/reference/dbformysql/v1api20231230#FlexibleServersFirewallRule" >}})   | 2023-12-30  | v1api20231230 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20231230/v1api20231230_flexibleserversfirewallrule.yaml)  |
| [User]({{< relref "/reference/dbformysql/v1#User" >}})                                                            | v1          | v1            | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api/v1_user.yaml)                                            |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                          | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                |
|-------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| [FlexibleServer]({{< relref "/reference/dbformysql/v1api20230630#FlexibleServer" >}})                             | 2023-06-30  | v1api20230630 | v2.7.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20230630/v1api20230630_flexibleserver.yaml)               |
| [FlexibleServer]({{< relref "/reference/dbformysql/v1api20210501#FlexibleServer" >}})                             | 2021-05-01  | v1api20210501 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20220101/v1api20210501_flexibleserver.yaml)               |
| [FlexibleServersAdministrator]({{< relref "/reference/dbformysql/v1api20230630#FlexibleServersAdministrator" >}}) | 2023-06-30  | v1api20230630 | v2.7.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20230630/v1api20230630_flexibleserversadministrator.yaml) |
| [FlexibleServersAdministrator]({{< relref "/reference/dbformysql/v1api20220101#FlexibleServersAdministrator" >}}) | 2022-01-01  | v1api20220101 | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20220101/v1api20220101_flexibleserversadministrator.yaml) |
| [FlexibleServersConfiguration]({{< relref "/reference/dbformysql/v1api20230630#FlexibleServersConfiguration" >}}) | 2023-06-30  | v1api20230630 | v2.7.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20230630/v1api20230630_flexibleserversconfiguration.yaml) |
| [FlexibleServersConfiguration]({{< relref "/reference/dbformysql/v1api20220101#FlexibleServersConfiguration" >}}) | 2022-01-01  | v1api20220101 | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20220101/v1api20220101_flexibleserversconfiguration.yaml) |
| [FlexibleServersDatabase]({{< relref "/reference/dbformysql/v1api20230630#FlexibleServersDatabase" >}})           | 2023-06-30  | v1api20230630 | v2.7.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20230630/v1api20230630_flexibleserversdatabase.yaml)      |
| [FlexibleServersDatabase]({{< relref "/reference/dbformysql/v1api20210501#FlexibleServersDatabase" >}})           | 2021-05-01  | v1api20210501 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20210501/v1api20210501_flexibleserversdatabase.yaml)      |
| [FlexibleServersFirewallRule]({{< relref "/reference/dbformysql/v1api20230630#FlexibleServersFirewallRule" >}})   | 2023-06-30  | v1api20230630 | v2.7.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20230630/v1api20230630_flexibleserversfirewallrule.yaml)  |
| [FlexibleServersFirewallRule]({{< relref "/reference/dbformysql/v1api20210501#FlexibleServersFirewallRule" >}})   | 2021-05-01  | v1api20210501 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformysql/v1api20210501/v1api20210501_flexibleserversfirewallrule.yaml)  |

