---
title: ManagedIdentity Supported Resources
linktitle: ManagedIdentity
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `managedidentity.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                    | ARM Version | CRD Version   | Supported From | Sample |
|-----------------------------|-------------|---------------|----------------|--------|
| FederatedIdentityCredential | 2023-01-31  | v1api20230131 | v2.5.0         | -      |
| UserAssignedIdentity        | 2023-01-31  | v1api20230131 | v2.5.0         | -      |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                                                                 | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                           |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [FederatedIdentityCredential](https://azure.github.io/azure-service-operator/reference/managedidentity/v1api20220131preview/#managedidentity.azure.com/v1api20220131preview.FederatedIdentityCredential) | 2022-01-31-preview | v1api20220131preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api20230131/v1api20220131preview_federatedidentitycredential.yaml) |
| [UserAssignedIdentity](https://azure.github.io/azure-service-operator/reference/managedidentity/v1api20181130/#managedidentity.azure.com/v1api20181130.UserAssignedIdentity)                             | 2018-11-30         | v1api20181130        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api20230131/v1api20181130_userassignedidentity.yaml)               |

