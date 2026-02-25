---
title: AppConfiguration Supported Resources
linktitle: AppConfiguration
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `appconfiguration.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                        | ARM Version | CRD Version | Supported From | Sample                                                                                                                                    |
|-------------------------------------------------------------------------------------------------|-------------|-------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| [ConfigurationStore]({{< relref "/reference/appconfiguration/v20240601#ConfigurationStore" >}}) | 2024-06-01  | v20240601   | v2.18.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/appconfiguration/v20240601/v20240601_configurationstore.yaml) |
| [KeyValue]({{< relref "/reference/appconfiguration/v20240601#KeyValue" >}})                     | 2024-06-01  | v20240601   | v2.18.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/appconfiguration/v20240601/v20240601_keyvalue.yaml)           |
| [Replica]({{< relref "/reference/appconfiguration/v20240601#Replica" >}})                       | 2024-06-01  | v20240601   | v2.18.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/appconfiguration/v20240601/v20240601_replica.yaml)            |
| [Snapshot]({{< relref "/reference/appconfiguration/v20240601#Snapshot" >}})                     | 2024-06-01  | v20240601   | v2.18.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/appconfiguration/v20240601/v20240601_snapshot.yaml)           |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                            | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                    |
|-----------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| [ConfigurationStore]({{< relref "/reference/appconfiguration/v1api20220501#ConfigurationStore" >}}) | 2022-05-01  | v1api20220501 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/appconfiguration/v1api/v1api20220501_configurationstore.yaml) |

