---
title: Cdn Supported Resources
linktitle: Cdn
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `cdn.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                                                                                                                                   | ARM Version | CRD Version   | Supported From | Sample                                                                                                                            |
|--------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------|
| [AfdCustomDomain](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.AfdCustomDomain) | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_afdcustomdomain.yaml) |
| [AfdEndpoint](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.AfdEndpoint)         | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_afdendpoint.yaml)     |
| [AfdOrigin](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.AfdOrigin)             | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_afdorigin.yaml)       |
| [AfdOriginGroup](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.AfdOriginGroup)   | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_afdorigingroup.yaml)  |
| [Profile](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.Profile)                 | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_profile.yaml)         |
| [Route](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.Route)                     | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_route.yaml)           |
| [Rule](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.Rule)                       | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_rule.yaml)            |
| [RuleSet](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.RuleSet)                 | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_ruleset.yaml)         |
| [Secret](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.Secret)                   | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_secret.yaml)          |
| [SecurityPolicy](https://azure.github.io/azure-service-operator/reference/cdn/v1api20230501/#cdn.azure.com/v1api20230501.SecurityPolicy)   | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_securitypolicy.yaml)  |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                     | ARM Version | CRD Version   | Supported From | Sample                                                                                                                             |
|----------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------|
| [Profile](https://azure.github.io/azure-service-operator/reference/cdn/v1api20210601/#cdn.azure.com/v1api20210601.Profile)                   | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20210601/v1api20210601_profile.yaml)          |
| [ProfilesEndpoint](https://azure.github.io/azure-service-operator/reference/cdn/v1api20210601/#cdn.azure.com/v1api20210601.ProfilesEndpoint) | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20210601/v1api20210601_profilesendpoint.yaml) |

