---
title: KeyVault Supported Resources
linktitle: KeyVault
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `keyvault.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                       | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                      |
|------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|-----------------------------------------------------------------------------------------------------------------------------|
| [Vault](https://azure.github.io/azure-service-operator/reference/keyvault/v1api20210401preview/#keyvault.azure.com/v1api20210401preview.Vault) | 2021-04-01-preview | v1api20210401preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/keyvault/v1api/v1api20210401preview_vault.yaml) |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                                                         | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                        |
|--------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|-----------------------|----------------|-------------------------------------------------------------------------------------------------------------------------------|
| [Vault](https://azure.github.io/azure-service-operator/reference/keyvault/v1beta20210401preview/#keyvault.azure.com/v1beta20210401preview.Vault) | 2021-04-01-preview | v1beta20210401preview | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/keyvault/v1beta/v1beta20210401preview_vault.yaml) |

