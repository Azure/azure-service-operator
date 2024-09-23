---
title: Kusto Supported Resources
linktitle: Kusto
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `kusto.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource | ARM Version | CRD Version   | Supported From | Sample                                                                                                                      |
|----------|-------------|---------------|----------------|-----------------------------------------------------------------------------------------------------------------------------|
| Cluster  | 2023-08-15  | v1api20230815 | v2.10.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/kusto/v1api20230815/v1api20230815_cluster.yaml) |
| Database | 2023-08-15  | v1api20230815 | v2.10.0        | -                                                                                                                           |

