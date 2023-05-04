// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqldb

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
)

// SqlDbManager is the client for the resource manager for SQL databases
type SqlDbManager interface {
	CreateOrUpdateDB(ctx context.Context,
		subscriptionID string,
		resourceGroupName string,
		location string,
		serverName string,
		tags map[string]*string,
		properties azuresqlshared.SQLDatabaseProperties) (pollingUrl string, db *sql.Database, err error)

	DeleteDB(ctx context.Context,
		subscriptionID string,
		resourceGroupName string,
		serverName string,
		databaseName string) (future *sql.DatabasesDeleteFuture, err error)

	GetDB(ctx context.Context,
		subscriptionID string,
		resourceGroupName string,
		serverName string, databaseName string) (sql.Database, error)

	GetServer(ctx context.Context,
		subscriptionID string,
		resourceGroupName string,
		serverName string) (result sql.Server, err error)

	AddLongTermRetention(ctx context.Context,
		subscriptionID string,
		resourceGroupName string,
		serverName string,
		databaseName string,
		policy azuresqlshared.SQLDatabaseBackupLongTermRetentionPolicy) (*sql.BackupLongTermRetentionPoliciesCreateOrUpdateFuture, error)

	resourcemanager.ARMClient
}
