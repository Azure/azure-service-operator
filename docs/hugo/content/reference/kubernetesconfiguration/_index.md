---
title: KubernetesConfiguration Supported Resources
linktitle: KubernetesConfiguration
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `kubernetesconfiguration.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                 | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                  |
|----------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Extension]({{< relref "/reference/kubernetesconfiguration/v1api20241101#Extension" >}})                 | 2024-11-01  | v1api20241101 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kubernetesconfiguration/v1api20241101/v1api20241101_extension.yaml)         |
| [FluxConfiguration]({{< relref "/reference/kubernetesconfiguration/v1api20241101#FluxConfiguration" >}}) | 2024-11-01  | v1api20241101 | v2.13.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kubernetesconfiguration/v1api20241101/v1api20241101_fluxconfiguration.yaml) |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                 | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                  |
|----------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Extension]({{< relref "/reference/kubernetesconfiguration/v1api20230501#Extension" >}})                 | 2023-05-01  | v1api20230501 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kubernetesconfiguration/v1api20230501/v1api20230501_extension.yaml)         |
| [FluxConfiguration]({{< relref "/reference/kubernetesconfiguration/v1api20230501#FluxConfiguration" >}}) | 2023-05-01  | v1api20230501 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kubernetesconfiguration/v1api20230501/v1api20230501_fluxconfiguration.yaml) |

