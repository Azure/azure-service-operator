---
title: CognitiveServices Supported Resources
linktitle: CognitiveServices
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `cognitiveservices.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                                                                             | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                     |
|--------------------------------------------------------------------------------------|-------------|---------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------------|
| [Account]({{< relref "/reference/cognitiveservices/v1api20250601#Account" >}})       | 2025-06-01  | v1api20250601 | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cognitiveservices/v1api20250601/v1api20250601_account.yaml)    |
| [Deployment]({{< relref "/reference/cognitiveservices/v1api20250601#Deployment" >}}) | 2025-06-01  | v1api20250601 | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/cognitiveservices/v1api20250601/v1api20250601_deployment.yaml) |

