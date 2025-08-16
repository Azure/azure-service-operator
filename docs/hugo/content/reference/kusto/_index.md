---
title: Kusto Supported Resources
linktitle: Kusto
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `kusto.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource            | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                  |
|---------------------|-------------|---------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| Cluster             | 2024-04-13  | v1api20240413 | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kusto/v1api20240413/v1api20240413_cluster.yaml)             |
| DataConnection      | 2024-04-13  | v1api20240413 | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kusto/v1api20240413/v1api20240413_dataconnection.yaml)      |
| Database            | 2024-04-13  | v1api20240413 | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kusto/v1api20240413/v1api20240413_database.yaml)            |
| PrincipalAssignment | 2024-04-13  | v1api20240413 | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kusto/v1api20240413/v1api20240413_principalassignment.yaml) |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                         | ARM Version | CRD Version   | Supported From | Sample                                                                                                                             |
|----------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------|
| [Cluster]({{< relref "/reference/kusto/v1api20230815#Cluster" >}})               | 2023-08-15  | v1api20230815 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kusto/v1api20230815/v1api20230815_cluster.yaml)        |
| [DataConnection]({{< relref "/reference/kusto/v1api20230815#DataConnection" >}}) | 2023-08-15  | v1api20230815 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kusto/v1api20230815/v1api20230815_dataconnection.yaml) |
| [Database]({{< relref "/reference/kusto/v1api20230815#Database" >}})             | 2023-08-15  | v1api20230815 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kusto/v1api20230815/v1api20230815_database.yaml)       |

