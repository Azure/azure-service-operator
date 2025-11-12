---
title: Quota Supported Resources
linktitle: Quota
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `quota.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                       | ARM Version | CRD Version   | Supported From | Sample                                                                                                                    |
|----------------------------------------------------------------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------|
| [Quota]({{< relref "/reference/quota/v1api20250901#Quota" >}}) | 2025-09-01  | v1api20250901 | v2.16.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/quota/v1api20250901/v1api20250901_quota.yaml) |

