---
title: EventHub Supported Resources
linktitle: EventHub
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `eventhub.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                             | ARM Version | CRD Version   | Supported From | Sample |
|--------------------------------------|-------------|---------------|----------------|--------|
| Namespace                            | 2024-01-01  | v1api20240101 | v2.12.0        | -      |
| NamespacesAuthorizationRule          | 2024-01-01  | v1api20240101 | v2.12.0        | -      |
| NamespacesEventhub                   | 2024-01-01  | v1api20240101 | v2.12.0        | -      |
| NamespacesEventhubsAuthorizationRule | 2024-01-01  | v1api20240101 | v2.12.0        | -      |
| NamespacesEventhubsConsumerGroup     | 2024-01-01  | v1api20240101 | v2.12.0        | -      |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                                                       | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                      |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Namespace](https://azure.github.io/azure-service-operator/reference/eventhub/v1api20211101/#eventhub.azure.com/v1api20211101.Namespace)                                                       | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespace.yaml)                            |
| [NamespacesAuthorizationRule](https://azure.github.io/azure-service-operator/reference/eventhub/v1api20211101/#eventhub.azure.com/v1api20211101.NamespacesAuthorizationRule)                   | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespacesauthorizationrule.yaml)          |
| [NamespacesEventhub](https://azure.github.io/azure-service-operator/reference/eventhub/v1api20211101/#eventhub.azure.com/v1api20211101.NamespacesEventhub)                                     | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespaceseventhub.yaml)                   |
| [NamespacesEventhubsAuthorizationRule](https://azure.github.io/azure-service-operator/reference/eventhub/v1api20211101/#eventhub.azure.com/v1api20211101.NamespacesEventhubsAuthorizationRule) | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespaceseventhubsauthorizationrule.yaml) |
| [NamespacesEventhubsConsumerGroup](https://azure.github.io/azure-service-operator/reference/eventhub/v1api20211101/#eventhub.azure.com/v1api20211101.NamespacesEventhubsConsumerGroup)         | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespaceseventhubsconsumergroup.yaml)     |

