// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresqldb

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"

	"github.com/Azure/go-autorest/autorest"
)

type SqlDbManager interface {
	CreateOrUpdateDB(ctx context.Context, resourceGroupName string, location string, serverName string, properties azuresqlshared.SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error)
	DeleteDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result autorest.Response, err error)
	GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error)
	GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error)
}
