---
title: Compute Supported Resources
linktitle: Compute
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `compute.azure.com/*` as one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Latest Released Versions

These resource(s) are the latest versions available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                         | ARM Version | CRD Version | Supported From | Sample                                                                                                                                 |
|----------------------------------|-------------|-------------|----------------|----------------------------------------------------------------------------------------------------------------------------------------|
| AvailabilitySet                  | 2024-11-01  | v20241101   | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v20241101/v20241101_availabilityset.yaml)          |
| CapacityReservation              | 2025-04-01  | v20250401   | v2.16.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v20250401/v20250401_capacityreservation.yaml)      |
| CapacityReservationGroup         | 2025-04-01  | v20250401   | v2.16.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v20250401/v20250401_capacityreservationgroup.yaml) |
| Disk                             | 2024-03-02  | v20240302   | v2.9.0         | -                                                                                                                                      |
| DiskAccess                       | 2024-03-02  | v20240302   | v2.9.0         | -                                                                                                                                      |
| DiskEncryptionSet                | 2024-03-02  | v20240302   | v2.9.0         | -                                                                                                                                      |
| Image                            | 2022-03-01  | v20220301   | v2.0.0         | -                                                                                                                                      |
| Snapshot                         | 2024-03-02  | v20240302   | v2.9.0         | -                                                                                                                                      |
| VirtualMachine                   | 2022-03-01  | v20220301   | v2.0.0         | -                                                                                                                                      |
| VirtualMachineScaleSet           | 2022-03-01  | v20220301   | v2.0.0         | -                                                                                                                                      |
| VirtualMachineScaleSetsExtension | 2022-03-01  | v20220301   | v2.6.0         | -                                                                                                                                      |
| VirtualMachinesExtension         | 2022-03-01  | v20220301   | v2.6.0         | -                                                                                                                                      |

### Other Supported Versions

These are older versions of resources still available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.
| Resource                         | ARM Version | CRD Version | Supported From | Sample |
|----------------------------------|-------------|-------------|----------------|--------|
| Disk                             | 2020-09-30  | v20200930   | v2.0.0         | -      |
| DiskEncryptionSet                | 2022-07-02  | v20220702   | v2.3.0         | -      |
| Image                            | 2021-07-01  | v20210701   | v2.0.0         | -      |
| Snapshot                         | 2020-09-30  | v20200930   | v2.0.0         | -      |
| VirtualMachine                   | 2020-12-01  | v20201201   | v2.0.0         | -      |
| VirtualMachineScaleSet           | 2020-12-01  | v20201201   | v2.0.0         | -      |
| VirtualMachineScaleSetsExtension | 2020-12-01  | v20201201   | v2.6.0         | -      |
| VirtualMachinesExtension         | 2020-12-01  | v20201201   | v2.6.0         | -      |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                               | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                         |
|------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| [AvailabilitySet]({{< relref "/reference/compute/v1api20241101#AvailabilitySet" >}})                                   | 2024-11-01  | v1api20241101 | v2.15.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api20241101/v1api20241101_availabilityset.yaml)          |
| [CapacityReservation]({{< relref "/reference/compute/v1api20250401#CapacityReservation" >}})                           | 2025-04-01  | v1api20250401 | v2.16.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api20250401/v1api20250401_capacityreservation.yaml)      |
| [CapacityReservationGroup]({{< relref "/reference/compute/v1api20250401#CapacityReservationGroup" >}})                 | 2025-04-01  | v1api20250401 | v2.16.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api20250401/v1api20250401_capacityreservationgroup.yaml) |
| [Disk]({{< relref "/reference/compute/v1api20240302#Disk" >}})                                                         | 2024-03-02  | v1api20240302 | v2.9.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20240302_disk.yaml)                             |
| [Disk]({{< relref "/reference/compute/v1api20200930#Disk" >}})                                                         | 2020-09-30  | v1api20200930 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20200930_disk.yaml)                             |
| [DiskAccess]({{< relref "/reference/compute/v1api20240302#DiskAccess" >}})                                             | 2024-03-02  | v1api20240302 | v2.9.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20240302_diskaccess.yaml)                       |
| [DiskEncryptionSet]({{< relref "/reference/compute/v1api20240302#DiskEncryptionSet" >}})                               | 2024-03-02  | v1api20240302 | v2.9.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20240302_diskencryptionset.yaml)                |
| [DiskEncryptionSet]({{< relref "/reference/compute/v1api20220702#DiskEncryptionSet" >}})                               | 2022-07-02  | v1api20220702 | v2.3.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220702_diskencryptionset.yaml)                |
| [Image]({{< relref "/reference/compute/v1api20220301#Image" >}})                                                       | 2022-03-01  | v1api20220301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_image.yaml)                            |
| [Image]({{< relref "/reference/compute/v1api20210701#Image" >}})                                                       | 2021-07-01  | v1api20210701 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20210701_image.yaml)                            |
| [Snapshot]({{< relref "/reference/compute/v1api20240302#Snapshot" >}})                                                 | 2024-03-02  | v1api20240302 | v2.9.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20240302_snapshot.yaml)                         |
| [Snapshot]({{< relref "/reference/compute/v1api20200930#Snapshot" >}})                                                 | 2020-09-30  | v1api20200930 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20200930_snapshot.yaml)                         |
| [VirtualMachine]({{< relref "/reference/compute/v1api20220301#VirtualMachine" >}})                                     | 2022-03-01  | v1api20220301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_virtualmachine.yaml)                   |
| [VirtualMachine]({{< relref "/reference/compute/v1api20201201#VirtualMachine" >}})                                     | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20201201_virtualmachine.yaml)                   |
| [VirtualMachineScaleSet]({{< relref "/reference/compute/v1api20220301#VirtualMachineScaleSet" >}})                     | 2022-03-01  | v1api20220301 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_virtualmachinescaleset.yaml)           |
| [VirtualMachineScaleSet]({{< relref "/reference/compute/v1api20201201#VirtualMachineScaleSet" >}})                     | 2020-12-01  | v1api20201201 | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20201201_virtualmachinescaleset.yaml)           |
| [VirtualMachineScaleSetsExtension]({{< relref "/reference/compute/v1api20220301#VirtualMachineScaleSetsExtension" >}}) | 2022-03-01  | v1api20220301 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_virtualmachinescalesetsextension.yaml) |
| [VirtualMachineScaleSetsExtension]({{< relref "/reference/compute/v1api20201201#VirtualMachineScaleSetsExtension" >}}) | 2020-12-01  | v1api20201201 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20201201_virtualmachinescalesetsextension.yaml) |
| [VirtualMachinesExtension]({{< relref "/reference/compute/v1api20220301#VirtualMachinesExtension" >}})                 | 2022-03-01  | v1api20220301 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20220301_virtualmachinesextension.yaml)         |
| [VirtualMachinesExtension]({{< relref "/reference/compute/v1api20201201#VirtualMachinesExtension" >}})                 | 2020-12-01  | v1api20201201 | v2.6.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/compute/v1api/v1api20201201_virtualmachinesextension.yaml)         |

