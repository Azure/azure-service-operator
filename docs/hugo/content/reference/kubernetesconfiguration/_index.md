---
title: KubernetesConfiguration Supported Resources
linktitle: KubernetesConfiguration
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `kubernetesconfiguration.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource  | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                  |
|-----------|-------------|---------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| Extension | 2023-05-01  | v1api20230501 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kubernetesconfiguration/v1api/v1api20230501_extension.yaml) |

