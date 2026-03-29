---
title: Communication Supported Resources
linktitle: Communication
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `communication.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                                                                                         | ARM Version | CRD Version | Supported From | Sample                                                                                                                                   |
|--------------------------------------------------------------------------------------------------|-------------|-------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------|
| [CommunicationService]({{< relref "/reference/communication/v20230401#CommunicationService" >}}) | 2023-04-01  | v20230401   | v2.19.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/communication/v20230401/v20230401_communicationservice.yaml) |
| [Domain]({{< relref "/reference/communication/v20230401#Domain" >}})                             | 2023-04-01  | v20230401   | v2.19.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/communication/v20230401/v20230401_domain.yaml)               |
| [EmailService]({{< relref "/reference/communication/v20230401#EmailService" >}})                 | 2023-04-01  | v20230401   | v2.19.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/communication/v20230401/v20230401_emailservice.yaml)         |
| [SenderUsername]({{< relref "/reference/communication/v20230401#SenderUsername" >}})             | 2023-04-01  | v20230401   | v2.19.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/communication/v20230401/v20230401_senderusername.yaml)       |

