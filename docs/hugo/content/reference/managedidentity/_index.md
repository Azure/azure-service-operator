---
title: ManagedIdentity Supported Resources
linktitle: ManagedIdentity
no_list: true
---
### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.


| Resource                                                                                                                                                                                                 | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                   |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|
| [FederatedIdentityCredential](https://azure.github.io/azure-service-operator/reference/managedidentity/v1api20220131preview/#managedidentity.azure.com/v1api20220131preview.FederatedIdentityCredential) | 2022-01-31-preview | v1api20220131preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api/v1api20220131preview_federatedidentitycredential.yaml) |
| [UserAssignedIdentity](https://azure.github.io/azure-service-operator/reference/managedidentity/v1api20181130/#managedidentity.azure.com/v1api20181130.UserAssignedIdentity)                             | 2018-11-30         | v1api20181130        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api/v1api20181130_userassignedidentity.yaml)               |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                                                                                                                   | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                                                     |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|-----------------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [FederatedIdentityCredential](https://azure.github.io/azure-service-operator/reference/managedidentity/v1beta20220131preview/#managedidentity.azure.com/v1beta20220131preview.FederatedIdentityCredential) | 2022-01-31-preview | v1beta20220131preview | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1beta/v1beta20220131preview_federatedidentitycredential.yaml) |
| [UserAssignedIdentity](https://azure.github.io/azure-service-operator/reference/managedidentity/v1beta20181130/#managedidentity.azure.com/v1beta20181130.UserAssignedIdentity)                             | 2018-11-30         | v1beta20181130        | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1beta/v1beta20181130_userassignedidentity.yaml)               |

