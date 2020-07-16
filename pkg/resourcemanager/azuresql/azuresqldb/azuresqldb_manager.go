// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqldb

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

// SqlDbManager is the client for the resource manager for SQL databases
type SqlDbManager interface {
	CreateOrUpdateDB(ctx context.Context,
		resourceGroupName string,
		location string,
		serverName string,
		tags map[string]*string,
		properties azuresqlshared.SQLDatabaseProperties) (pollingUrl string, db *sql.Database, err error)

	DeleteDB(ctx context.Context,
		resourceGroupName string,
		serverName string,
		databaseName string) (future *sql.DatabasesDeleteFuture, err error)

	GetDB(ctx context.Context,
		resourceGroupName string,
		serverName string, databaseName string) (sql.Database, error)

	GetServer(ctx context.Context,
		resourceGroupName string,
		serverName string) (result sql.Server, err error)

	AddLongTermRetention(ctx context.Context,
		resourceGroupName string,
		serverName string,
		databaseName string,
		weeklyRetention string,
		monthlyRetention string,
		yearlyRetention string,
		weekOfYear int32) (*http.Response, error)

	resourcemanager.ARMClient
}
