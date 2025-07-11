---
title: Web Supported Resources
linktitle: Web
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `web.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                               | ARM Version | CRD Version   | Supported From | Sample                                                                                                                       |
|----------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------|
| [ServerFarm]({{< relref "/reference/web/v1api20220301#ServerFarm" >}})                 | 2022-03-01  | v1api20220301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/web/v1api/v1api20220301_serverfarm.yaml)         |
| [Site]({{< relref "/reference/web/v1api20220301#Site" >}})                             | 2022-03-01  | v1api20220301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/web/v1api/v1api20220301_site.yaml)               |
| [SitesSourcecontrol]({{< relref "/reference/web/v1api20220301#SitesSourcecontrol" >}}) | 2022-03-01  | v1api20220301 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/web/v1api/v1api20220301_sitessourcecontrol.yaml) |

