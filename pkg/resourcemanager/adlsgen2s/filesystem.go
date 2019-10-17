package adlsgen2s

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	// "github.com/google/uuid"
	// "github.com/Azure/go-autorest/autorest/azure/auth"
	// "github.com/Azure/go-autorest/autorest/to"
	"log"
	// "github.com/Azure/azure-sdk-for-go/services/datalake/store/2016-11-01/filesystem"
)

type azureFileSystemManager struct{}


func (_ *azureFileSystemManager) CreateFileSystem(ctx context.Context, filesystemName string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, datalakeName string) (*autorest.Response, error) {
	client := getFileSystemClient(datalakeName)
	// req, err := client.CreatePreparer(ctx, filesystemName, "", "", nil, "")
	// if err != nil {
	// 	fmt.Println("failed to create preparer")
	// 	return nil, err
	// }
	// fmt.Println(req)
	// resp, err := client.CreateSender(req)
	resp, err := client.GetProperties(ctx, filesystemName, xMsClientRequestID, timeout, xMsDate)
	xMsProperties = resp.Header.Get("xMsProperties")
	xMsClientRequestID = resp.Header.Get("xMsClientRequestID")
	xMsDate = resp.Header.Get("xMsDate")

	// result, err := client.CreateResponder(resp)

	// bear minimum logic to check auth
	result, err := client.Create(ctx, filesystemName, xMsProperties, xMsClientRequestID, timeout, xMsDate)

	if err != nil {
		return nil, err
	}

	return &result, err
}

func (_ *azureFileSystemManager) GetFileSystem(ctx context.Context, filesystemName string, xMsClientRequestID string, xMsDate string, datalakeName string) (autorest.Response, error) {
	client := getFileSystemClient(datalakeName)
	
	list, err := client.List(ctx, filesystemName, "", nil, xMsClientRequestID, nil, xMsDate)
	response := list.Response
	return response, err
}

func (_ *azureFileSystemManager) DeleteFileSystem(ctx context.Context, filesystemName string, xMsClientRequestID string, xMsDate string, datalakeName string) (autorest.Response, error) {
	client := getFileSystemClient(datalakeName)

	return client.Delete(ctx, filesystemName, "", "", xMsClientRequestID, nil, xMsDate)
}

func getFileSystemClient(accountName string) storagedatalake.FilesystemClient {
	// I think this is where the issue is. 
	// In all other examples, when you create a new "Client" you pass the subscription ID as a param. 
	// NewFileSystemClient seems to be an anomaly 
	xmsversion := "2019-02-02"
	fsClient := storagedatalake.NewFilesystemClient(xmsversion, accountName)
	a, err := iam.GetResourceManagementAuthorizer()

	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	fsClient.Authorizer = a
	// fsClient.AddToUserAgent(storagedatalake.UserAgent())
	fsClient.AddToUserAgent(config.UserAgent())

	return fsClient
}

