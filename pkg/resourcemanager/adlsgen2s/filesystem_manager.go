package adlsgen2s

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
)

type FileSystemManager interface {
	CreateFileSystem(ctx context.Context, goupName string, filesystem string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, accountName string) (*autorest.Response, error)

	GetFileSystem(ctx context.Context, groupName string, filesystemName string, xMsClientRequestID string, xMsDate string, datalakeName string) (autorest.Response, error)

	DeleteFileSystem(ctx context.Context, groupName string, filesystemName string, xMsClientRequestID string, xMsDate string, datalakeName string) (autorest.Response, error)
}
