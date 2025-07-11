---
title: ManagedIdentity Supported Resources
linktitle: ManagedIdentity
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `managedidentity.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                             | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                    |
|----------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| [FederatedIdentityCredential]({{< relref "/reference/managedidentity/v1api20230131#FederatedIdentityCredential" >}}) | 2023-01-31  | v1api20230131 | v2.5.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api20230131/v1api20230131_federatedidentitycredential.yaml) |
| [UserAssignedIdentity]({{< relref "/reference/managedidentity/v1api20230131#UserAssignedIdentity" >}})               | 2023-01-31  | v1api20230131 | v2.5.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api20230131/v1api20230131_userassignedidentity.yaml)        |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                                    | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                                  |
|-----------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [FederatedIdentityCredential]({{< relref "/reference/managedidentity/v1api20220131preview#FederatedIdentityCredential" >}}) | 2022-01-31-preview | v1api20220131preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api20220131preview/v1api20220131preview_federatedidentitycredential.yaml) |
| [UserAssignedIdentity]({{< relref "/reference/managedidentity/v1api20181130#UserAssignedIdentity" >}})                      | 2018-11-30         | v1api20181130        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/managedidentity/v1api20181130/v1api20181130_userassignedidentity.yaml)                      |

