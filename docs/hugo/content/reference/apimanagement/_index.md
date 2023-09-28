---
title: ApiManagement Supported Resources
linktitle: ApiManagement
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `apimanagement.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource               | ARM Version | CRD Version   | Supported From | Sample |
|------------------------|-------------|---------------|----------------|--------|
| Service_Api            | 2022-08-01  | v1api20220801 | v2.4.0         | -      |
| Service_ApiVersionSet  | 2022-08-01  | v1api20220801 | v2.4.0         | -      |
| Service_Backend        | 2022-08-01  | v1api20220801 | v2.4.0         | -      |
| Service_NamedValue     | 2022-08-01  | v1api20220801 | v2.4.0         | -      |
| Service_Policy         | 2022-08-01  | v1api20220801 | v2.4.0         | -      |
| Service_PolicyFragment | 2022-08-01  | v1api20220801 | v2.4.0         | -      |
| Service_Product        | 2022-08-01  | v1api20220801 | v2.4.0         | -      |
| Service_Subscription   | 2022-08-01  | v1api20220801 | v2.4.0         | -      |

