---
title: MachineLearningServices Supported Resources
linktitle: MachineLearningServices
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `machinelearningservices.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                                                     | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                             |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| [Workspace](https://azure.github.io/azure-service-operator/reference/machinelearningservices/v1api20210701/#machinelearningservices.azure.com/v1api20210701.Workspace)                       | 2021-07-01  | v1api20210701 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api/v1api20210701_workspace.yaml)            |
| [WorkspacesCompute](https://azure.github.io/azure-service-operator/reference/machinelearningservices/v1api20210701/#machinelearningservices.azure.com/v1api20210701.WorkspacesCompute)       | 2021-07-01  | v1api20210701 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api/v1api20210701_workspacescompute.yaml)    |
| [WorkspacesConnection](https://azure.github.io/azure-service-operator/reference/machinelearningservices/v1api20210701/#machinelearningservices.azure.com/v1api20210701.WorkspacesConnection) | 2021-07-01  | v1api20210701 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1api/v1api20210701_workspacesconnection.yaml) |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                                                                                                       | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                               |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Workspace](https://azure.github.io/azure-service-operator/reference/machinelearningservices/v1beta20210701/#machinelearningservices.azure.com/v1beta20210701.Workspace)                       | 2021-07-01  | v1beta20210701 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1beta/v1beta20210701_workspace.yaml)            |
| [WorkspacesCompute](https://azure.github.io/azure-service-operator/reference/machinelearningservices/v1beta20210701/#machinelearningservices.azure.com/v1beta20210701.WorkspacesCompute)       | 2021-07-01  | v1beta20210701 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1beta/v1beta20210701_workspacescompute.yaml)    |
| [WorkspacesConnection](https://azure.github.io/azure-service-operator/reference/machinelearningservices/v1beta20210701/#machinelearningservices.azure.com/v1beta20210701.WorkspacesConnection) | 2021-07-01  | v1beta20210701 | v2.0.0-beta.2  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/machinelearningservices/v1beta/v1beta20210701_workspacesconnection.yaml) |

