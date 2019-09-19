package storages

import ."github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"

var MockStorageManagers = StorageManagers{
	Storage:       &mockStorageManager{},
	BlobContainer: &mockBlobContainerManager{},
}
