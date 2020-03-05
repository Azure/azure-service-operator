// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
)

type FileSystemManager interface {
	CreateFileSystem(ctx context.Context, goupName string, filesystem string, timeout *int32, xMsDate string, accountName string) (*autorest.Response, error)

	GetFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (autorest.Response, error)

	DeleteFileSystem(ctx context.Context, groupName string, filesystemName string, timeout *int32, xMsDate string, datalakeName string) (autorest.Response, error)
}
