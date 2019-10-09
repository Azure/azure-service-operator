package adlsgen2s

import (
"context"
"github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"


// apiv1 "github.com/Azure/azure-service-operator/api/v1"
)

type AdlsGen2Manager interface {
	CreateAdlsGen2(ctx context.Context, filesystem string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, accountName string) (*storagedatalake.FileSystem, error)

	GetAdlsGen2()

	DeleteAdlsGen2()
}
