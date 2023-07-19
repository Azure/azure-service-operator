---
title: Search Supported Resources
linktitle: Search
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `search.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

The `authOptions` and `operatorSpec` properties on [`SearchService`](https://azure.github.io/azure-service-operator/reference/search/v1api20220901/#search.azure.com/v1api20220901.SearchService) are not supported in ASO v2.1.0; they'll be included in the next release.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                     | ARM Version | CRD Version   | Supported From | Sample                                                                                                                     |
|----------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|----------------------------------------------------------------------------------------------------------------------------|
| [SearchService](https://azure.github.io/azure-service-operator/reference/search/v1api20220901/#search.azure.com/v1api20220901.SearchService) | 2022-09-01  | v1api20220901 | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/search/v1api/v1api20220901_searchservice.yaml) |

