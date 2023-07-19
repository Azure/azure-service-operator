---
title: DataFactory Supported Resources
linktitle: DataFactory
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `datafactory.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                   | ARM Version | CRD Version   | Supported From | Sample                                                                                                                    |
|--------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------|
| [Factory](https://azure.github.io/azure-service-operator/reference/datafactory/v1api20180601/#datafactory.azure.com/v1api20180601.Factory) | 2018-06-01  | v1api20180601 | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/datafactory/v1api/v1api20180601_factory.yaml) |

