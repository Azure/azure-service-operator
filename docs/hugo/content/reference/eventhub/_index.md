---
title: EventHub Supported Resources
linktitle: EventHub
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `eventhub.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                                        | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                      |
|---------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Namespace]({{< relref "/reference/eventhub/v1api20240101#Namespace" >}})                                                       | 2024-01-01  | v1api20240101 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20240101/v1api20240101_namespace.yaml)                            |
| [NamespacesAuthorizationRule]({{< relref "/reference/eventhub/v1api20240101#NamespacesAuthorizationRule" >}})                   | 2024-01-01  | v1api20240101 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20240101/v1api20240101_namespacesauthorizationrule.yaml)          |
| [NamespacesEventhub]({{< relref "/reference/eventhub/v1api20240101#NamespacesEventhub" >}})                                     | 2024-01-01  | v1api20240101 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20240101/v1api20240101_namespaceseventhub.yaml)                   |
| [NamespacesEventhubsAuthorizationRule]({{< relref "/reference/eventhub/v1api20240101#NamespacesEventhubsAuthorizationRule" >}}) | 2024-01-01  | v1api20240101 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20240101/v1api20240101_namespaceseventhubsauthorizationrule.yaml) |
| [NamespacesEventhubsConsumerGroup]({{< relref "/reference/eventhub/v1api20240101#NamespacesEventhubsConsumerGroup" >}})         | 2024-01-01  | v1api20240101 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20240101/v1api20240101_namespaceseventhubsconsumergroup.yaml)     |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                                                        | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                      |
|---------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Namespace]({{< relref "/reference/eventhub/v1api20211101#Namespace" >}})                                                       | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespace.yaml)                            |
| [NamespacesAuthorizationRule]({{< relref "/reference/eventhub/v1api20211101#NamespacesAuthorizationRule" >}})                   | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespacesauthorizationrule.yaml)          |
| [NamespacesEventhub]({{< relref "/reference/eventhub/v1api20211101#NamespacesEventhub" >}})                                     | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespaceseventhub.yaml)                   |
| [NamespacesEventhubsAuthorizationRule]({{< relref "/reference/eventhub/v1api20211101#NamespacesEventhubsAuthorizationRule" >}}) | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespaceseventhubsauthorizationrule.yaml) |
| [NamespacesEventhubsConsumerGroup]({{< relref "/reference/eventhub/v1api20211101#NamespacesEventhubsConsumerGroup" >}})         | 2021-11-01  | v1api20211101 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventhub/v1api20211101/v1api20211101_namespaceseventhubsconsumergroup.yaml)     |

