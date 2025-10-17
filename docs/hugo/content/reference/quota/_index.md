---
title: Quota Supported Resources
linktitle: Quota
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `quota.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource | ARM Version | CRD Version   | Supported From | Sample                                                                                                                    |
|----------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------|
| Quota    | 2025-09-01  | v1api20250901 | v2.16.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/quota/v1api20250901/v1api20250901_quota.yaml) |

