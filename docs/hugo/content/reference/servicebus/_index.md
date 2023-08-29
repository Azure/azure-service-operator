---
title: ServiceBus Supported Resources
linktitle: ServiceBus
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `servicebus.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                          | ARM Version        | CRD Version          | Supported From | Sample |
|-----------------------------------|--------------------|----------------------|----------------|--------|
| Namespace                         | 2022-10-01-preview | v1api20221001preview | v2.3.0         | -      |
| Namespace                         | 2021-11-01         | v1api20211101        | v2.3.0         | -      |
| NamespacesAuthorizationRule       | 2022-10-01-preview | v1api20221001preview | v2.3.0         | -      |
| NamespacesAuthorizationRule       | 2021-11-01         | v1api20211101        | v2.3.0         | -      |
| NamespacesQueue                   | 2022-10-01-preview | v1api20221001preview | v2.3.0         | -      |
| NamespacesQueue                   | 2021-11-01         | v1api20211101        | v2.3.0         | -      |
| NamespacesTopic                   | 2022-10-01-preview | v1api20221001preview | v2.3.0         | -      |
| NamespacesTopic                   | 2021-11-01         | v1api20211101        | v2.3.0         | -      |
| NamespacesTopicsSubscription      | 2022-10-01-preview | v1api20221001preview | v2.3.0         | -      |
| NamespacesTopicsSubscription      | 2021-11-01         | v1api20211101        | v2.3.0         | -      |
| NamespacesTopicsSubscriptionsRule | 2022-10-01-preview | v1api20221001preview | v2.3.0         | -      |
| NamespacesTopicsSubscriptionsRule | 2021-11-01         | v1api20211101        | v2.3.0         | -      |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                                                                   | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                    |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Namespace](https://azure.github.io/azure-service-operator/reference/servicebus/v1api20210101preview/#servicebus.azure.com/v1api20210101preview.Namespace)                                                 | 2021-01-01-preview | v1api20210101preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespace.yaml)                         |
| [NamespacesAuthorizationRule](https://azure.github.io/azure-service-operator/reference/servicebus/v1api20210101preview/#servicebus.azure.com/v1api20210101preview.NamespacesAuthorizationRule)             | 2021-01-01-preview | v1api20210101preview | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespacesauthorizationrule.yaml)       |
| [NamespacesQueue](https://azure.github.io/azure-service-operator/reference/servicebus/v1api20210101preview/#servicebus.azure.com/v1api20210101preview.NamespacesQueue)                                     | 2021-01-01-preview | v1api20210101preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespacesqueue.yaml)                   |
| [NamespacesTopic](https://azure.github.io/azure-service-operator/reference/servicebus/v1api20210101preview/#servicebus.azure.com/v1api20210101preview.NamespacesTopic)                                     | 2021-01-01-preview | v1api20210101preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespacestopic.yaml)                   |
| [NamespacesTopicsSubscription](https://azure.github.io/azure-service-operator/reference/servicebus/v1api20210101preview/#servicebus.azure.com/v1api20210101preview.NamespacesTopicsSubscription)           | 2021-01-01-preview | v1api20210101preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespacestopicssubscription.yaml)      |
| [NamespacesTopicsSubscriptionsRule](https://azure.github.io/azure-service-operator/reference/servicebus/v1api20210101preview/#servicebus.azure.com/v1api20210101preview.NamespacesTopicsSubscriptionsRule) | 2021-01-01-preview | v1api20210101preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1api/v1api20210101preview_namespacestopicssubscriptionsrule.yaml) |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                                                                                                                     | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                                                      |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|-----------------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Namespace](https://azure.github.io/azure-service-operator/reference/servicebus/v1beta20210101preview/#servicebus.azure.com/v1beta20210101preview.Namespace)                                                 | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespace.yaml)                         |
| [NamespacesQueue](https://azure.github.io/azure-service-operator/reference/servicebus/v1beta20210101preview/#servicebus.azure.com/v1beta20210101preview.NamespacesQueue)                                     | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespacesqueue.yaml)                   |
| [NamespacesTopic](https://azure.github.io/azure-service-operator/reference/servicebus/v1beta20210101preview/#servicebus.azure.com/v1beta20210101preview.NamespacesTopic)                                     | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespacestopic.yaml)                   |
| [NamespacesTopicsSubscription](https://azure.github.io/azure-service-operator/reference/servicebus/v1beta20210101preview/#servicebus.azure.com/v1beta20210101preview.NamespacesTopicsSubscription)           | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespacestopicssubscription.yaml)      |
| [NamespacesTopicsSubscriptionsRule](https://azure.github.io/azure-service-operator/reference/servicebus/v1beta20210101preview/#servicebus.azure.com/v1beta20210101preview.NamespacesTopicsSubscriptionsRule) | 2021-01-01-preview | v1beta20210101preview | v2.0.0-beta.3  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/servicebus/v1beta/v1beta20210101preview_namespacestopicssubscriptionsrule.yaml) |

