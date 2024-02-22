---
title: DataProtection Supported Resources
linktitle: DataProtection
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `dataprotection.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                                               | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                  |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| BackupVault                                                                                                                                                                            | 2023-12-01  | v1api20231201 | v2.2.0         | -                                                                                                                                                       |
| [BackupVault](https://azure.github.io/azure-service-operator/reference/dataprotection/v1api20230101/#dataprotection.azure.com/v1api20230101.BackupVault)                               | 2023-01-01  | v1api20230101 | v2.2.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dataprotection/v1api20230101/v1api20230101_backupvault.yaml)                |
| BackupVaultsBackupInstance                                                                                                                                                             | 2023-12-01  | v1api20231201 | v2.5.0         | -                                                                                                                                                       |
| [BackupVaultsBackupInstance](https://azure.github.io/azure-service-operator/reference/dataprotection/v1api20230101/#dataprotection.azure.com/v1api20230101.BackupVaultsBackupInstance) | 2023-01-01  | v1api20230101 | v2.5.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dataprotection/v1api20230101/v1api20230101_backupvaultsbackupinstance.yaml) |
| BackupVaultsBackupPolicy                                                                                                                                                               | 2023-12-01  | v1api20231201 | v2.2.0         | -                                                                                                                                                       |
| [BackupVaultsBackupPolicy](https://azure.github.io/azure-service-operator/reference/dataprotection/v1api20230101/#dataprotection.azure.com/v1api20230101.BackupVaultsBackupPolicy)     | 2023-01-01  | v1api20230101 | v2.2.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/dataprotection/v1api20230101/v1api20230101_backupvaultsbackuppolicy.yaml)   |

