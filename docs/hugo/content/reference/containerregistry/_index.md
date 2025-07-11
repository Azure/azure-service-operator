---
title: ContainerRegistry Supported Resources
linktitle: ContainerRegistry
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `containerregistry.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                               | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                              |
|--------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| [Registry]({{< relref "/reference/containerregistry/v1api20230701#Registry" >}})                       | 2023-07-01  | v1api20230701 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerregistry/v1api20230701/v1api20230701_registry.yaml)            |
| [RegistryReplication]({{< relref "/reference/containerregistry/v1api20230701#RegistryReplication" >}}) | 2023-07-01  | v1api20230701 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerregistry/v1api20230701/v1api20230701_registryreplication.yaml) |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                         | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                   |
|----------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------|
| [Registry]({{< relref "/reference/containerregistry/v1api20210901#Registry" >}}) | 2021-09-01  | v1api20210901 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerregistry/v1api20210901/v1api20210901_registry.yaml) |

