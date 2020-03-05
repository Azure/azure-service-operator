// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import . "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"

var MockStorageManagers = StorageManagers{
	Storage:       &mockStorageManager{},
	BlobContainer: &mockBlobContainerManager{},
	FileSystem:    &mockFileSystemManager{},
}
