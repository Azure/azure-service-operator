---
title: RedHatOpenShift Supported Resources
linktitle: RedHatOpenShift
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `redhatopenshift.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource         | ARM Version | CRD Version   | Supported From | Sample |
|------------------|-------------|---------------|----------------|--------|
| OpenShiftCluster | 2023-11-22  | v1api20231122 | v2.8.0         | -      |

