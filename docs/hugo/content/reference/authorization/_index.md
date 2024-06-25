---
title: Authorization Supported Resources
linktitle: Authorization
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `authorization.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                                   | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                   |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|
| [RoleAssignment](https://azure.github.io/azure-service-operator/reference/authorization/v1api20220401/#authorization.azure.com/v1api20220401.RoleAssignment)               | 2022-04-01         | v1api20220401        | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/authorization/v1api20220401/v1api20220401_roleassignment.yaml)               |
| [RoleAssignment](https://azure.github.io/azure-service-operator/reference/authorization/v1api20200801preview/#authorization.azure.com/v1api20200801preview.RoleAssignment) | 2020-08-01-preview | v1api20200801preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/authorization/v1api20200801preview/v1api20200801preview_roleassignment.yaml) |
| [RoleDefinition](https://azure.github.io/azure-service-operator/reference/authorization/v1api20220401/#authorization.azure.com/v1api20220401.RoleDefinition)               | 2022-04-01         | v1api20220401        | v2.8.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/authorization/v1api20220401/v1api20220401_roledefinition.yaml)               |

