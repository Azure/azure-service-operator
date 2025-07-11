---
title: Entra Supported Resources
linktitle: Entra
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `entra.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                            | ARM Version | CRD Version | Supported From | Sample                                                                                                      |
|---------------------------------------------------------------------|-------------|-------------|----------------|-------------------------------------------------------------------------------------------------------------|
| [SecurityGroup]({{< relref "/reference/entra/v1#SecurityGroup" >}}) | v1          | v1          | v2.14.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/entra/v1/v1_securitygroup.yaml) |

