---
title: AppConfiguration Supported Resources
linktitle: AppConfiguration
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `appconfiguration.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                                   | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                    |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| [ConfigurationStore](https://azure.github.io/azure-service-operator/reference/appconfiguration/v1api20220501/#appconfiguration.azure.com/v1api20220501.ConfigurationStore) | 2022-05-01  | v1api20220501 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/appconfiguration/v1api/v1api20220501_configurationstore.yaml) |

