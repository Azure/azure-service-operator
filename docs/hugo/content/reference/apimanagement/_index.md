---
title: ApiManagement Supported Resources
linktitle: ApiManagement
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `apimanagement.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                                                                                                                                       | ARM Version | CRD Version   | Supported From | Sample                                                                                                                              |
|------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------|
| [Service](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.Service) | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_service.yaml) |

