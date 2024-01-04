---
title: ApiManagement Supported Resources
linktitle: ApiManagement
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `apimanagement.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                                                                                                                                                   | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                    |
|------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| [ProductApi](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.ProductApi)       | 2022-08-01  | v1api20220801 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_productapi.yaml)    |
| [ProductPolicy](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.ProductPolicy) | 2022-08-01  | v1api20220801 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_productpolicy.yaml) |
| [User](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.User)                   | 2022-08-01  | v1api20220801 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_user.yaml)          |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                     | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                     |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------------|
| [Api](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.Api)                       | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_api.yaml)            |
| [ApiVersionSet](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.ApiVersionSet)   | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_apiversionset.yaml)  |
| [Backend](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.Backend)               | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_backend.yaml)        |
| [NamedValue](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.NamedValue)         | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_namedvalue.yaml)     |
| [Policy](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.Policy)                 | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_policy.yaml)         |
| [PolicyFragment](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.PolicyFragment) | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_policyfragment.yaml) |
| [Product](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.Product)               | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_product.yaml)        |
| [Service](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.Service)               | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_service.yaml)        |
| [Subscription](https://azure.github.io/azure-service-operator/reference/apimanagement/v1api20220801/#apimanagement.azure.com/v1api20220801.Subscription)     | 2022-08-01  | v1api20220801 | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/apimanagement/v1api20220801/v1api20220801_subscription.yaml)   |

