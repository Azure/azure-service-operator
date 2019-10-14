package adlsgen2s

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"log"
)

type azureFileSystemManager struct{}

func (_ *azureFileSystemManager) CreateFileSystem(ctx context.Context, filesystem string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, accountName string) (*autorest.Response, error) {
	fsClient := getFsClient(accountName)
	// TODO: check to make sure filesystem name conforms correctly
	result, err := fsClient.Create(ctx, filesystem, xMsProperties, xMsClientRequestID, timeout, xMsDate)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (_ *azureFileSystemManager) GetFileSystem() {

}

func (_ *azureFileSystemManager) DeleteFileSystem() {

}

func getFileSystemClient(accountName string) storagedatalake.FilesystemClient {
	xmsversion := "2019-10-31"
	fsClient := storagedatalake.NewFilesystemClient(xmsversion, accountName)

	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	fsClient.Authorizer = a
	fsClient.AddToUserAgent(config.UserAgent())
	return fsClient
}
