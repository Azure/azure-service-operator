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
	"fmt"
	"log"
	// "github.com/Azure/azure-sdk-for-go/services/datalake/store/2016-11-01/filesystem"
)

type azureFileSystemManager struct{}

// func (_ *azureFileSystemManager) CreateFileSystem(ctx context.Context, filesystem string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, datalakeName string) (*autorest.Response, error) {
// 	fsClient := getFileSystemClient(datalakeName)

// 	req, err := fsClient.CreatePreparer(ctx, filesystem, "", "", to.Int32Ptr(100), "")	// TODO: check to make sure filesystem name conforms correctly
// 	// result, err := fsClient.Create(ctx, filesystem, "", "", to.Int32Ptr(100), "")
// 	response, err := fsClient.CreateSender(req)
// 	result, err := fsClient.CreateResponder(response)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &result, err
// }
func (_ *azureFileSystemManager) CreateFileSystem(ctx context.Context, filesystemName string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, datalakeName string) (*autorest.Response, error) {
	client := getFileSystemClient(datalakeName)

	// bear minimum logic to check auth
	rep, err := client.Create(ctx, "test", "", "", nil, "")


	if err != nil {
		return nil, err
	}
	fmt.Println(rep)
	return nil, err
}

func (_ *azureFileSystemManager) GetFileSystem() {

}

func (_ *azureFileSystemManager) DeleteFileSystem() {

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
	fsClient.AddToUserAgent(config.UserAgent())

	return fsClient
}
