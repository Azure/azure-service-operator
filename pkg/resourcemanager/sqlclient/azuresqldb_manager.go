// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"
)

func NewAzureSqlDbManager(log logr.Logger) *AzureSqlDbManager {
	return &AzureSqlDbManager{Log: log}
}

type SqlDBManager interface {
	CreateOrUpdateDB(ctx context.Context, resourceGroupName string, location string, serverName string, properties SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error)
	DeleteDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result autorest.Response, err error)
	GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error)
}
