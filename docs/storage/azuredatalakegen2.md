# Azure Data Lake Gen2 Operator 

## Resources Supported

The Azure Data Lake Gen2 operator can be used to provision the following resources

1. Azure Data Lake Gen2 enabled storage account - Deploys a data lake gen2 enabled storage account using the storage operator
2. Azure Data Lake Gen2 Filesystem - Deploys a filesystem inside of a data lake gen2 enabled storage account

### Azure Data Lake Gen 2 Enabled Storage Account

A sample YAML for the Data Lake Gen2 FileSystem is [here](/config/samples/azure_v1alpha1_azuredatalakegen2storage.yaml)

Important Values:

- Kind: the Custom Resource Definition (CRD) name
- Metadata.Name: the name of the data lake enabled storage account that will be created
- Spec.Location: Azure region where you want to create the data lake enabled storage account
- Spec.ResourceGroup: Name of the resource group under which you want to create the data lake enabled storage account
- Spec.Kind: The kind of storage account that you want to create. This value must be `StorageV2` in order for it to be a data lake
- Spec.AccessTier - The access tier for the storage account. Choose "Hot" access tier for frequently accessed data and "Cool" for infrequently accessed data
- Spec.SupportsHttpsTrafficOnly - When "enabled", requires connection requests to the storage account to be only over HTTPS
- Spec.DataLakeEnabled: If set to `true`, the storage account will have hierarchical namespace enabled and will, therefore, be a data lake enabled storage account

### Azure Data Lake Gen 2 FileSystem

A sample YAML for the Data Lake Gen2 FileSystem is [here](/config/samples/azure_v1alpha1_azuredatalakegen2filesystem.yaml)

Important Values:

- Kind: the Custom Resource Definition (CRD) name
- Metadata.Name: the name of the data lake enabled storage account that will be created
- Spec.StorageAccountName: Name of the data lake enabled storage account in which you would like to create the filesystem
- Spec.ResourceGroup: Name of the resource group under which the storage account lives

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
