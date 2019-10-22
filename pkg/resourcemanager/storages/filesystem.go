package storages

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"log"
	"net/http"
)

type azureFileSystemManager struct{}

func (_ *azureFileSystemManager) CreateFileSystem(ctx context.Context, groupName string, filesystemName string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, datalakeName string) (*autorest.Response, error) {
	client := getFileSystemClient(ctx, groupName, datalakeName)

	// bear minimum logic to check auth
	result, err := client.Create(ctx, filesystemName, xMsProperties, xMsClientRequestID, timeout, xMsDate)

	if err != nil {
		return nil, err
	}

	return &result, err
}

func (_ *azureFileSystemManager) GetFileSystem(ctx context.Context, groupName string, filesystemName string, xMsClientRequestID string, xMsDate string, datalakeName string) (autorest.Response, error) {
	response := autorest.Response{Response: &http.Response{StatusCode: http.StatusNotFound}}
	client := getFileSystemClient(ctx, groupName, datalakeName)

	list, err := client.List(ctx, filesystemName, "", nil, xMsClientRequestID, nil, xMsDate)

	if len(*list.Filesystems) == 0 {
		return response, err
	}

	response = list.Response

	return response, err
}

func (_ *azureFileSystemManager) DeleteFileSystem(ctx context.Context, groupName string, filesystemName string, xMsClientRequestID string, xMsDate string, datalakeName string) (autorest.Response, error) {
	client := getFileSystemClient(ctx, groupName, datalakeName)

	return client.Delete(ctx, filesystemName, "", "", xMsClientRequestID, to.Int32Ptr(20), xMsDate)
}

func getFileSystemClient(ctx context.Context, groupName string, accountName string) storagedatalake.FilesystemClient {
	xmsversion := "2019-02-02"
	fsClient := storagedatalake.NewFilesystemClient(xmsversion, accountName)
	adlsClient := getStoragesClient()

	a, err := iam.GetSharedKeyAuthorizer(ctx, groupName, accountName, adlsClient)

	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	fsClient.Authorizer = a
	fsClient.AddToUserAgent(config.UserAgent())

	return fsClient
}
