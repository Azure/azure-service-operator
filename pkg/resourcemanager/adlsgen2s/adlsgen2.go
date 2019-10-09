package adlsgen2s

import (
	"context"
	"log"
	"github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
)

type azureAdlsGen2Manager struct{}

func getAdlsGen2Client() {}

func (_ *azureAdlsGen2Manager) CreateAdlsGen2(ctx context.Context, filesystem string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, accountName string) (*storagedatalake.FileSystem, error) {
	fsClient := getFsClient(accountName)
	// TODO: check to make sure filesystem name conforms correctly
	future, err := fsClient.Create(ctx, filesystem, xMsProperties, xMsClientRequestID, timeout, xMsDate)
	if err != nil {
		return nil, err
	}

	err = future.WaitForCompletionRef(ctx, fsClient.FileSystemClient)
	if err != nil {
		return nil, err
	}
	result, err := future.Result(fsClient)
	return &result, err
}

func (_ *azureAdlsGen2Manager) GetAdlsGen2() {

}

func (_ *azureAdlsGen2Manager) DeleteAdlsGen2() {

}

func getFsClient(accountName string) storagedatalake.FileSystemClient {
	xmsversion := "2019-10-31"
	fsClient := storagedatalake.NewFileSystemClient(xmsversion, accountName)
	
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	fsClient.Authorizer = a
	fsClient.AddToUserAgent(config.UserAgent())
	return fsClient
}
