# Storage Account Operator

This operator deploys an Azure Storage Account into a specified resource group at the specified location.

Learn more about Azure Storage Accounts [here](https://docs.microsoft.com/en-us/azure/storage/common/storage-account-overview).

Here is a [sample YAML](/config/samples/azure_v1alpha1_storageaccount.yaml) to provision an Azure Storage Account.

The spec is comprised of the following fields:

* Location
* ResourceGroup
* Sku
    * Name
* Kind
* AccessTier
* SupportsHttpsTrafficOnly
* DataLakeEnabled

### Required Fields

A Storage Account needs the following fields to deploy, along with a location and resource group.

* `Sku.Name` SKU of the storage account, possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS', 'PremiumZRS', 'StandardGZRS', 'StandardRAGZRS'
* `Kind` specify the kind of storage account, defaults to 'StorageV2'. Possible values include: 'BlobStorage', 'BlockBlobStorage', 'FileStorage', 'Storage', 'StorageV2'
* `AccessTier` specify the access tier, defaults to 'Hot'. Possible values include: 'Hot', 'Cold'
* `SupportsHttpTrafficOnly` allows only HTTPS traffic to storage account if set to true. Default value is true.

### Optional Fields

* `DataLakeEnabled` Enable a hierarchical namespace in your storage account, which requires the Storage Account `Kind` to be set to 'StorageV2'. Read more about Data Lake Storage hierarchical namespaces [here](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-namespace).

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.

## Secrets
After creating a storage account, the operator stores a JSON formatted secret with the following fields. For more details on where the secrets are stored, look [here](/docs/secrets.md).
* `key1` (depending on the number of keys, there could be up to keyn)
* `connectionString1` (depending on the number of keys, there could be up to connectionStringn)