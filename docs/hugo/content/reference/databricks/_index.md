---
title: Databricks Supported Resources
linktitle: Databricks
no_list: true
---

To install the CRDs for these resources, your ASO configuration must include `databricks.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                            | ARM Version | CRD Version   | Supported From | Sample |
|-------------------------------------|-------------|---------------|----------------|--------|
| AccessConnector                     | 2024-05-01  | v1api20240501 | v2.17.0        | -      |
| Workspace                           | 2024-05-01  | v1api20240501 | v2.17.0        | -      |
| WorkspacesPrivateEndpointConnection | 2024-05-01  | v1api20240501 | v2.17.0        | -      |
| WorkspacesVirtualNetworkPeering     | 2024-05-01  | v1api20240501 | v2.17.0        | -      |

