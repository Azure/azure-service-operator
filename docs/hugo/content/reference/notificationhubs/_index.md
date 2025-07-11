---
title: NotificationHubs Supported Resources
linktitle: NotificationHubs
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `notificationhubs.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                                          | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                           |
|-----------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Namespace]({{< relref "/reference/notificationhubs/v1api20230901#Namespace" >}})                                                 | 2023-09-01  | v1api20230901 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/notificationhubs/v1api20230901/v1api20230901_namespace.yaml)                         |
| [NamespacesAuthorizationRule]({{< relref "/reference/notificationhubs/v1api20230901#NamespacesAuthorizationRule" >}})             | 2023-09-01  | v1api20230901 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/notificationhubs/v1api20230901/v1api20230901_namespacesauthorizationrule.yaml)       |
| [NotificationHub]({{< relref "/reference/notificationhubs/v1api20230901#NotificationHub" >}})                                     | 2023-09-01  | v1api20230901 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/notificationhubs/v1api20230901/v1api20230901_notificationhub.yaml)                   |
| [NotificationHubsAuthorizationRule]({{< relref "/reference/notificationhubs/v1api20230901#NotificationHubsAuthorizationRule" >}}) | 2023-09-01  | v1api20230901 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/notificationhubs/v1api20230901/v1api20230901_notificationhubsauthorizationrule.yaml) |

