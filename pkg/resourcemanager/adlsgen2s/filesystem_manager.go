package adlsgen2s

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"
)

type FileSystemManager interface {
	CreateFileSystem(ctx context.Context, filesystem string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, accountName string) (*storagedatalake.Filesystem, error)

	GetFileSystem()

	DeleteFileSystem()
}
