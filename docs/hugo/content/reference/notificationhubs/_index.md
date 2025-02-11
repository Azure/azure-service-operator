---
title: NotificationHubs Supported Resources
linktitle: NotificationHubs
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `notificationhubs.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                                                                                                                                                                                                 | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                           |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Namespace](https://azure.github.io/azure-service-operator/reference/notificationhubs/v1api20230901/#notificationhubs.azure.com/v1api20230901.Namespace)                                                 | 2023-09-01  | v1api20230901 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/notificationhubs/v1api20230901/v1api20230901_namespace.yaml)                         |
| [NamespacesAuthorizationRule](https://azure.github.io/azure-service-operator/reference/notificationhubs/v1api20230901/#notificationhubs.azure.com/v1api20230901.NamespacesAuthorizationRule)             | 2023-09-01  | v1api20230901 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/notificationhubs/v1api20230901/v1api20230901_namespacesauthorizationrule.yaml)       |
| [NotificationHub](https://azure.github.io/azure-service-operator/reference/notificationhubs/v1api20230901/#notificationhubs.azure.com/v1api20230901.NotificationHub)                                     | 2023-09-01  | v1api20230901 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/notificationhubs/v1api20230901/v1api20230901_notificationhub.yaml)                   |
| [NotificationHubsAuthorizationRule](https://azure.github.io/azure-service-operator/reference/notificationhubs/v1api20230901/#notificationhubs.azure.com/v1api20230901.NotificationHubsAuthorizationRule) | 2023-09-01  | v1api20230901 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/notificationhubs/v1api20230901/v1api20230901_notificationhubsauthorizationrule.yaml) |

