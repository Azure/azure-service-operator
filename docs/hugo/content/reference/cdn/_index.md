---
title: Cdn Supported Resources
linktitle: Cdn
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `cdn.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                           | ARM Version | CRD Version   | Supported From | Sample                                                                                                                             |
|------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------|
| [AfdCustomDomain]({{< relref "/reference/cdn/v1api20230501#AfdCustomDomain" >}})   | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_afdcustomdomain.yaml)  |
| [AfdEndpoint]({{< relref "/reference/cdn/v1api20230501#AfdEndpoint" >}})           | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_afdendpoint.yaml)      |
| [AfdOrigin]({{< relref "/reference/cdn/v1api20230501#AfdOrigin" >}})               | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_afdorigin.yaml)        |
| [AfdOriginGroup]({{< relref "/reference/cdn/v1api20230501#AfdOriginGroup" >}})     | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_afdorigingroup.yaml)   |
| [Profile]({{< relref "/reference/cdn/v1api20230501#Profile" >}})                   | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_profile.yaml)          |
| [ProfilesEndpoint]({{< relref "/reference/cdn/v1api20210601#ProfilesEndpoint" >}}) | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20210601/v1api20210601_profilesendpoint.yaml) |
| [Route]({{< relref "/reference/cdn/v1api20230501#Route" >}})                       | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_route.yaml)            |
| [Rule]({{< relref "/reference/cdn/v1api20230501#Rule" >}})                         | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_rule.yaml)             |
| [RuleSet]({{< relref "/reference/cdn/v1api20230501#RuleSet" >}})                   | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_ruleset.yaml)          |
| [Secret]({{< relref "/reference/cdn/v1api20230501#Secret" >}})                     | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_secret.yaml)           |
| [SecurityPolicy]({{< relref "/reference/cdn/v1api20230501#SecurityPolicy" >}})     | 2023-05-01  | v1api20230501 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20230501/v1api20230501_securitypolicy.yaml)   |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                         | ARM Version | CRD Version   | Supported From | Sample                                                                                                                    |
|------------------------------------------------------------------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------|
| [Profile]({{< relref "/reference/cdn/v1api20210601#Profile" >}}) | 2021-06-01  | v1api20210601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cdn/v1api20210601/v1api20210601_profile.yaml) |

