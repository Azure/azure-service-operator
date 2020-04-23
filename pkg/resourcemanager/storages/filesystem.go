// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"context"
	"errors"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	"github.com/Azure/go-autorest/autorest"
)

type azureFileSystemManager struct{}

func (_ *azureFileSystemManager) CreateFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (*autorest.Response, error) {
	err := checkRGAndStorageAccount(ctx, groupName, datalakeName)
	if err != nil {
		return nil, err
	}

	client, err := getFileSystemClient(ctx, groupName, datalakeName)
	if err != nil {
		return nil, err
	}

	// empty parameters are optional request headers (properties, requestid)
	result, err := client.Create(ctx, filesystemName, "", "", timeout, xMsDate)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (_ *azureFileSystemManager) GetFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (autorest.Response, error) {
	response := autorest.Response{Response: &http.Response{StatusCode: http.StatusNotFound}}

	err := checkRGAndStorageAccount(ctx, groupName, datalakeName)
	if err != nil {
		return response, errors.New("unable to create filesystem")
	}
	response = autorest.Response{Response: &http.Response{StatusCode: http.StatusNotFound}}

	client, err := getFileSystemClient(ctx, groupName, datalakeName)
	if err != nil {
		return response, err
	}

	// empty parameters are optional request headers (continuation, maxresults, requestid)
	list, err := client.List(ctx, filesystemName, "", nil, "", timeout, xMsDate)

	if len(*list.Filesystems) == 0 {
		return response, err
	}

	response = list.Response

	return response, err
}

func (_ *azureFileSystemManager) DeleteFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (autorest.Response, error) {
	response := autorest.Response{Response: &http.Response{StatusCode: http.StatusAccepted}}
	err := checkRGAndStorageAccount(ctx, groupName, datalakeName)
	if err != nil {
		return response, nil
	}

	client, err := getFileSystemClient(ctx, groupName, datalakeName)
	if err != nil {
		return response, err
	}

	// empty parameters are optional request headers (ifmodifiedsince, ifunmodifiedsince, requestid)
	return client.Delete(ctx, filesystemName, "", "", "", timeout, xMsDate)
}

func getFileSystemClient(ctx context.Context, groupName string, accountName string) (storagedatalake.FilesystemClient, error) {
	xmsversion := "2019-02-02"
	fsClient := storagedatalake.NewFilesystemClient(xmsversion, accountName)

	accountKey, err := getAccountKey(ctx, groupName, accountName)
	if err != nil {
		return storagedatalake.FilesystemClient{}, err
	}

	a, err := iam.GetSharedKeyAuthorizer(accountName, accountKey)
	if err != nil {
		return storagedatalake.FilesystemClient{}, err
	}

	fsClient.Authorizer = a
	fsClient.AddToUserAgent(config.UserAgent())

	return fsClient, nil
}

func getResourcesClient() (resources.GroupsClient, error) {
	resourcesClient := resources.NewGroupsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return resources.GroupsClient{}, err
	}
	resourcesClient.Authorizer = a
	resourcesClient.AddToUserAgent(config.UserAgent())
	return resourcesClient, nil
}

func checkRGAndStorageAccount(ctx context.Context, groupName string, datalakeName string) error {
	rgClient, err := getResourcesClient()
	if err != nil {
		return err
	}

	storagesClient := storageaccount.New()

	response, err := rgClient.CheckExistence(ctx, groupName)
	if response.IsHTTPStatus(404) {
		return errhelp.NewAzureError(errors.New("ResourceGroupNotFound"))
	}

	_, err = storagesClient.ListKeys(ctx, groupName, datalakeName)
	if err != nil {
		return errhelp.NewAzureError(errors.New("ParentResourceNotFound"))
	}

	return err
}

func getAccountKey(ctx context.Context, groupName string, accountName string) (accountKey string, err error) {
	adlsClient := storageaccount.New()
	keys, err := adlsClient.ListKeys(ctx, groupName, accountName)
	if err != nil {
		return "", err
	}

	for _, key := range *keys.Keys {
		if *key.KeyName == "key1" {
			accountKey = *key.Value
		}
	}
	return accountKey, err
}
