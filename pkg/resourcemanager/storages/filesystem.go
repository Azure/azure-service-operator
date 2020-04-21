// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest"
)

type azureFileSystemManager struct {
	SecretClient secrets.SecretClient
}

func (f *azureFileSystemManager) CreateFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (*autorest.Response, error) {
	err := checkRGAndStorageAccount(ctx, groupName, datalakeName, f.SecretClient)
	if err != nil {
		return nil, err
	}
	client := getFileSystemClient(ctx, groupName, datalakeName, f.SecretClient)
	// empty parameters are optional request headers (properties, requestid)
	result, err := client.Create(ctx, filesystemName, "", "", timeout, xMsDate)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (f *azureFileSystemManager) GetFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (autorest.Response, error) {
	response := autorest.Response{Response: &http.Response{StatusCode: http.StatusNotFound}}

	err := checkRGAndStorageAccount(ctx, groupName, datalakeName, f.SecretClient)
	if err != nil {
		return response, errors.New("unable to create filesystem")
	}
	response = autorest.Response{Response: &http.Response{StatusCode: http.StatusNotFound}}
	client := getFileSystemClient(ctx, groupName, datalakeName, f.SecretClient)

	// empty parameters are optional request headers (continuation, maxresults, requestid)
	list, err := client.List(ctx, filesystemName, "", nil, "", timeout, xMsDate)

	if len(*list.Filesystems) == 0 {
		return response, err
	}

	response = list.Response

	return response, err
}

func (f *azureFileSystemManager) DeleteFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (autorest.Response, error) {
	response := autorest.Response{Response: &http.Response{StatusCode: http.StatusAccepted}}
	err := checkRGAndStorageAccount(ctx, groupName, datalakeName, f.SecretClient)
	if err != nil {
		return response, nil
	}
	client := getFileSystemClient(ctx, groupName, datalakeName, f.SecretClient)
	// empty parameters are optional request headers (ifmodifiedsince, ifunmodifiedsince, requestid)
	return client.Delete(ctx, filesystemName, "", "", "", timeout, xMsDate)
}

func getFileSystemClient(ctx context.Context, groupName string, accountName string, secretClient secrets.SecretClient) storagedatalake.FilesystemClient {
	xmsversion := "2019-02-02"
	fsClient := storagedatalake.NewFilesystemClient(xmsversion, accountName)

	accountKey, err := getAccountKey(ctx, groupName, accountName, secretClient)
	if err != nil {
		log.Printf("failed to get the account key for the authorizer: %v\n", err)
	}

	a, err := iam.GetSharedKeyAuthorizer(accountName, accountKey)

	if err != nil {
		log.Printf("failed to initialize authorizer: %v\n", err)
	}
	fsClient.Authorizer = a
	fsClient.AddToUserAgent(config.UserAgent())

	return fsClient
}

func getResourcesClient() resources.GroupsClient {
	resourcesClient := resources.NewGroupsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	resourcesClient.Authorizer = a
	resourcesClient.AddToUserAgent(config.UserAgent())
	return resourcesClient
}

func checkRGAndStorageAccount(ctx context.Context, groupName string, datalakeName string, secretClient secrets.SecretClient) error {
	rgClient := getResourcesClient()
	storagesClient := storageaccount.New(secretClient)

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

func getAccountKey(ctx context.Context, groupName string, accountName string, secretClient secrets.SecretClient) (accountKey string, err error) {
	adlsClient := storageaccount.New(secretClient)
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
