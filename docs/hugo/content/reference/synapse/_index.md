---
title: Synapse Supported Resources
linktitle: Synapse
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `synapse.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                       | ARM Version | CRD Version   | Supported From | Sample                                                                                                                              |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------|
| [Workspace](https://azure.github.io/azure-service-operator/reference/synapse/v1api20210601/#synapse.azure.com/v1api20210601.Workspace)                         | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/synapse/v1api/v1api20210601_workspace.yaml)             |
| [WorkspacesBigDataPool](https://azure.github.io/azure-service-operator/reference/synapse/v1api20210601/#synapse.azure.com/v1api20210601.WorkspacesBigDataPool) | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/synapse/v1api/v1api20210601_workspacesbigdatapool.yaml) |

