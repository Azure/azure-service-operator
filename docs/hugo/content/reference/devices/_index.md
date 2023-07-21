---
title: Devices Supported Resources
linktitle: Devices
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `devices.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                         | ARM Version | CRD Version   | Supported From | Sample                                                                                                               |
|----------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|----------------------------------------------------------------------------------------------------------------------|
| [IotHub](https://azure.github.io/azure-service-operator/reference/devices/v1api20210702/#devices.azure.com/v1api20210702.IotHub) | 2021-07-02  | v1api20210702 | v2.1.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/devices/v1api/v1api20210702_iothub.yaml) |

