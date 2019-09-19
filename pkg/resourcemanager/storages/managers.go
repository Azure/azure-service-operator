package storages

type StorageManagers struct {
	Storage       StorageManager
	BlobContainer BlobContainerManager
}

var AzureStorageManagers = StorageManagers{
	Storage:       &azureStorageManager{},
	BlobContainer: &azureBlobContainerManager{},
}
