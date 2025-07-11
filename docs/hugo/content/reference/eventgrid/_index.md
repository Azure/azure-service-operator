---
title: EventGrid Supported Resources
linktitle: EventGrid
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `eventgrid.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                   | ARM Version | CRD Version   | Supported From | Sample                                                                                                                            |
|--------------------------------------------------------------------------------------------|-------------|---------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------|
| [Domain]({{< relref "/reference/eventgrid/v1api20200601#Domain" >}})                       | 2020-06-01  | v1api20200601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1api/v1api20200601_domain.yaml)            |
| [DomainsTopic]({{< relref "/reference/eventgrid/v1api20200601#DomainsTopic" >}})           | 2020-06-01  | v1api20200601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1api/v1api20200601_domainstopic.yaml)      |
| [EventSubscription]({{< relref "/reference/eventgrid/v1api20200601#EventSubscription" >}}) | 2020-06-01  | v1api20200601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1api/v1api20200601_eventsubscription.yaml) |
| [Topic]({{< relref "/reference/eventgrid/v1api20200601#Topic" >}})                         | 2020-06-01  | v1api20200601 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/eventgrid/v1api/v1api20200601_topic.yaml)             |

