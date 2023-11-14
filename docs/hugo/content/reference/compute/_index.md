---
title: Compute Supported Resources
linktitle: Compute
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `compute.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                         | ARM Version | CRD Version   | Supported From | Sample                                                                                                                               |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------|
| [Disk](https://azure.github.io/azure-service-operator/reference/compute/v1api20200930/#compute.azure.com/v1api20200930.Disk)                                     | 2020-09-30  | v1api20200930 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20200930_disk.yaml)                   |
| [DiskEncryptionSet](https://azure.github.io/azure-service-operator/reference/compute/v1api20220702/#compute.azure.com/v1api20220702.DiskEncryptionSet)           | 2022-07-02  | v1api20220702 | v2.3.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220702_diskencryptionset.yaml)      |
| [Image](https://azure.github.io/azure-service-operator/reference/compute/v1api20220301/#compute.azure.com/v1api20220301.Image)                                   | 2022-03-01  | v1api20220301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_image.yaml)                  |
| [Image](https://azure.github.io/azure-service-operator/reference/compute/v1api20210701/#compute.azure.com/v1api20210701.Image)                                   | 2021-07-01  | v1api20210701 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20210701_image.yaml)                  |
| [Snapshot](https://azure.github.io/azure-service-operator/reference/compute/v1api20200930/#compute.azure.com/v1api20200930.Snapshot)                             | 2020-09-30  | v1api20200930 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20200930_snapshot.yaml)               |
| [VirtualMachine](https://azure.github.io/azure-service-operator/reference/compute/v1api20220301/#compute.azure.com/v1api20220301.VirtualMachine)                 | 2022-03-01  | v1api20220301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_virtualmachine.yaml)         |
| [VirtualMachine](https://azure.github.io/azure-service-operator/reference/compute/v1api20201201/#compute.azure.com/v1api20201201.VirtualMachine)                 | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20201201_virtualmachine.yaml)         |
| [VirtualMachineScaleSet](https://azure.github.io/azure-service-operator/reference/compute/v1api20220301/#compute.azure.com/v1api20220301.VirtualMachineScaleSet) | 2022-03-01  | v1api20220301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_virtualmachinescaleset.yaml) |
| [VirtualMachineScaleSet](https://azure.github.io/azure-service-operator/reference/compute/v1api20201201/#compute.azure.com/v1api20201201.VirtualMachineScaleSet) | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20201201_virtualmachinescaleset.yaml) |

