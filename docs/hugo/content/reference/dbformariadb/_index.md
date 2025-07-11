---
title: DBforMariaDB Supported Resources
linktitle: DBforMariaDB
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `dbformariadb.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                              | ARM Version | CRD Version   | Supported From | Sample                                                                                                                           |
|---------------------------------------------------------------------------------------|-------------|---------------|----------------|----------------------------------------------------------------------------------------------------------------------------------|
| [Configuration]({{< relref "/reference/dbformariadb/v1api20180601#Configuration" >}}) | 2018-06-01  | v1api20180601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformariadb/v1api/v1api20180601_configuration.yaml) |
| [Database]({{< relref "/reference/dbformariadb/v1api20180601#Database" >}})           | 2018-06-01  | v1api20180601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformariadb/v1api/v1api20180601_database.yaml)      |
| [Server]({{< relref "/reference/dbformariadb/v1api20180601#Server" >}})               | 2018-06-01  | v1api20180601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dbformariadb/v1api/v1api20180601_server.yaml)        |

