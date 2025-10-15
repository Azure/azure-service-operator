---
title: RedHatOpenShift Supported Resources
linktitle: RedHatOpenShift
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `redhatopenshift.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource                         | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                                       |
|----------------------------------|--------------------|----------------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| HcpOpenShiftCluster              | 2024-06-10-preview | v1api20240610preview | v2.16.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/redhatopenshift/v1api20240610preview/v1api20240610preview_hcpopenshiftcluster.yaml)              |
| HcpOpenShiftClustersExternalAuth | 2024-06-10-preview | v1api20240610preview | v2.16.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/redhatopenshift/v1api20240610preview/v1api20240610preview_hcpopenshiftclustersexternalauth.yaml) |
| HcpOpenShiftClustersNodePool     | 2024-06-10-preview | v1api20240610preview | v2.16.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/redhatopenshift/v1api20240610preview/v1api20240610preview_hcpopenshiftclustersnodepool.yaml)     |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                                                                                       | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                 |
|------------------------------------------------------------------------------------------------|-------------|---------------|----------------|----------------------------------------------------------------------------------------------------------------------------------------|
| [OpenShiftCluster]({{< relref "/reference/redhatopenshift/v1api20231122#OpenShiftCluster" >}}) | 2023-11-22  | v1api20231122 | v2.9.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/redhatopenshift/v1api/v1api20231122_openshiftcluster.yaml) |

