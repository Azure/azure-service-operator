---
title: MachineLearningServices Supported Resources
linktitle: MachineLearningServices
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `machinelearningservices.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                       | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                     |
|----------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Registry]({{< relref "/reference/machinelearningservices/v1api20240401#Registry" >}})                         | 2024-04-01  | v1api20240401 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api20240401/v1api20240401_registry.yaml)             |
| [Workspace]({{< relref "/reference/machinelearningservices/v1api20240401#Workspace" >}})                       | 2024-04-01  | v1api20240401 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api20240401/v1api20240401_workspace.yaml)            |
| [WorkspacesCompute]({{< relref "/reference/machinelearningservices/v1api20240401#WorkspacesCompute" >}})       | 2024-04-01  | v1api20240401 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api20240401/v1api20240401_workspacescompute.yaml)    |
| [WorkspacesConnection]({{< relref "/reference/machinelearningservices/v1api20240401#WorkspacesConnection" >}}) | 2024-04-01  | v1api20240401 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api20240401/v1api20240401_workspacesconnection.yaml) |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                       | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                     |
|----------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Workspace]({{< relref "/reference/machinelearningservices/v1api20210701#Workspace" >}})                       | 2021-07-01  | v1api20210701 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api20210701/v1api20210701_workspace.yaml)            |
| [WorkspacesCompute]({{< relref "/reference/machinelearningservices/v1api20210701#WorkspacesCompute" >}})       | 2021-07-01  | v1api20210701 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api20210701/v1api20210701_workspacescompute.yaml)    |
| [WorkspacesConnection]({{< relref "/reference/machinelearningservices/v1api20210701#WorkspacesConnection" >}}) | 2021-07-01  | v1api20210701 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api20210701/v1api20210701_workspacesconnection.yaml) |

