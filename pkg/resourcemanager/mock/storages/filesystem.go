/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storages

import (
	"context"
	"errors"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest"
	"net/http"
)

type fileSystemResource struct {
	resourceGroupName  string
	storageAccountName string
	filesystemName     string
}

type mockFileSystemManager struct {
	fileSystemResource []fileSystemResource
}

func findFileSystem(res []fileSystemResource, predicate func(fileSystemResource) bool) (int, fileSystemResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, fileSystemResource{}
}

// Creates a filesystem in a storage account
func (manager *mockFileSystemManager) CreateFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (*autorest.Response, error) {
	fs := fileSystemResource{
		resourceGroupName:  groupName,
		storageAccountName: datalakeName,
		filesystemName:     filesystemName,
	}

	manager.fileSystemResource = append(manager.fileSystemResource, fs)
	mockresponse := helpers.GetRestResponse(http.StatusOK)

	return &mockresponse, nil
}

// Gets a filesystem
func (manager *mockFileSystemManager) GetFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (autorest.Response, error) {
	groups := manager.fileSystemResource

	index, _ := findFileSystem(groups, func(g fileSystemResource) bool {
		return g.resourceGroupName == groupName &&
			g.storageAccountName == datalakeName &&
			g.filesystemName == filesystemName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errhelp.NewAzureError(errors.New("filesystem not found"))
	}
	return helpers.GetRestResponse(http.StatusOK), nil
}

// Deletes the filesystem
func (manager *mockFileSystemManager) DeleteFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (autorest.Response, error) {
	groups := manager.fileSystemResource

	index, _ := findFileSystem(groups, func(g fileSystemResource) bool {
		return g.resourceGroupName == groupName &&
			g.storageAccountName == datalakeName &&
			g.filesystemName == filesystemName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errhelp.NewAzureError(errors.New("filesystem not found"))
	}

	manager.fileSystemResource = append(groups[:index], groups[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}
