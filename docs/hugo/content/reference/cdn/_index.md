---
title: Cdn Supported Resources
linktitle: Cdn
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `cdn.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                     | ARM Version | CRD Version   | Supported From | Sample                                                                                                                     |
|----------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|----------------------------------------------------------------------------------------------------------------------------|
| [Profile](https://azure.github.io/azure-service-operator/reference/cdn/v1api20210601/#cdn.azure.com/v1api20210601.Profile)                   | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api/v1api20210601_profile.yaml)          |
| [ProfilesEndpoint](https://azure.github.io/azure-service-operator/reference/cdn/v1api20210601/#cdn.azure.com/v1api20210601.ProfilesEndpoint) | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api/v1api20210601_profilesendpoint.yaml) |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                                                       | ARM Version | CRD Version    | Supported From | Sample                                                                                                                       |
|------------------------------------------------------------------------------------------------------------------------------------------------|-------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------------------|
| [Profile](https://azure.github.io/azure-service-operator/reference/cdn/v1beta20210601/#cdn.azure.com/v1beta20210601.Profile)                   | 2021-06-01  | v1beta20210601 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1beta/v1beta20210601_profile.yaml)          |
| [ProfilesEndpoint](https://azure.github.io/azure-service-operator/reference/cdn/v1beta20210601/#cdn.azure.com/v1beta20210601.ProfilesEndpoint) | 2021-06-01  | v1beta20210601 | v2.0.0-beta.1  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1beta/v1beta20210601_profilesendpoint.yaml) |

