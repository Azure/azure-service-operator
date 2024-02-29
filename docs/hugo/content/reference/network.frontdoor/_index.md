---
title: Network Supported Resources
linktitle: Network
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `network.frontdoor.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                                                                                                                                                                                         | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                               |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| [WebApplicationFirewallPolicy](https://azure.github.io/azure-service-operator/reference/network.frontdoor/v1api20220501/#network.frontdoor.azure.com/v1api20220501.WebApplicationFirewallPolicy) | 2022-05-01  | v1api20220501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/network.frontdoor/v1api/v1api20220501_webapplicationfirewallpolicy.yaml) |

