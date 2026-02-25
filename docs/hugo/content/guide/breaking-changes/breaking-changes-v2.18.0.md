---
title: "v2.18.0 Breaking Changes"
linkTitle: "v2.18.0"
weight: -55  # This should be 5 lower than the previous breaking change document
---
## Breaking changes

- **`Storage.Tier` removed from `StorageAccount` spec (`storage/v20250601`)**: The `Tier` property on `StorageAccount.Spec.Sku` has been removed. This field was marked as `readonly` in the upstream TypeSpec migration ([azure-rest-api-specs openapi.json#L12317](https://github.com/Azure/azure-rest-api-specs/blob/main/specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/openapi.json#L12317)). The `Tier` value has always been effectively ignored by Azure - it is computed from the `Name` field (e.g., `Standard_LRS` → Standard tier). Users specifying this field in their CRDs will need to remove it.

- **`DataEncryption.GeoBackupEncryptionKeyStatus` and `DataEncryption.PrimaryEncryptionKeyStatus` removed from `dbforpostgresql/FlexibleServer` spec (`v20250801`)**: These two fields have been removed from the `DataEncryption` object on the PostgreSQL FlexibleServer spec. These fields were always read-only but never documented as such until now. The upstream spec now treats these as readonly/status-only properties. Users specifying these fields will need to remove them from their CRDs.

- **`insights/v1api20240101preview/ScheduledQueryRule` API version removed**: The preview API version `2024-01-01-preview` for `ScheduledQueryRule` has been completely removed as it was deprecated and removed upstream. Users using this API version should migrate to a supported version if they haven't already

## Upcoming breaking changes

* Will remove containerservice ManagedCluster and AgentPool api version v1api20240402preview in v2.19 of ASO.
