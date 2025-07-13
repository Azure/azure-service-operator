---
title: SignalRService Supported Resources
linktitle: SignalRService
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `signalrservice.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                        | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                         |
|-------------------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| [CustomCertificate]({{< relref "/reference/signalrservice/v1api20240301#CustomCertificate" >}}) | 2024-03-01  | v1api20240301 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/signalrservice/v1api20240301/v1api20240301_customcertificate.yaml) |
| [CustomDomain]({{< relref "/reference/signalrservice/v1api20240301#CustomDomain" >}})           | 2024-03-01  | v1api20240301 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/signalrservice/v1api20240301/v1api20240301_customdomain.yaml)      |
| [Replica]({{< relref "/reference/signalrservice/v1api20240301#Replica" >}})                     | 2024-03-01  | v1api20240301 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/signalrservice/v1api20240301/v1api20240301_replica.yaml)           |
| [SignalR]({{< relref "/reference/signalrservice/v1api20240301#SignalR" >}})                     | 2024-03-01  | v1api20240301 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/signalrservice/v1api20240301/v1api20240301_signalr.yaml)           |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                    | ARM Version | CRD Version   | Supported From | Sample                                                                                                                               |
|-----------------------------------------------------------------------------|-------------|---------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------|
| [SignalR]({{< relref "/reference/signalrservice/v1api20211001#SignalR" >}}) | 2021-10-01  | v1api20211001 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/signalrservice/v1api20211001/v1api20211001_signalr.yaml) |

