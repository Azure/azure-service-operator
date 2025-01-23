---
title: NotificationHubs Supported Resources
linktitle: NotificationHubs
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `notificationhubs.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                          | ARM Version | CRD Version   | Supported From | Sample |
|-----------------------------------|-------------|---------------|----------------|--------|
| Namespace                         | 2023-09-01  | v1api20230901 | v2.12.0        | -      |
| NamespacesAuthorizationRule       | 2023-09-01  | v1api20230901 | v2.12.0        | -      |
| NotificationHub                   | 2023-09-01  | v1api20230901 | v2.12.0        | -      |
| NotificationHubsAuthorizationRule | 2023-09-01  | v1api20230901 | v2.12.0        | -      |

