// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqldb

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest"
)

// SqlDbManager is the client for the resource manager for SQL databases
type SqlDbManager interface {
	CreateOrUpdateDB(ctx context.Context,
		resourceGroupName string,
		location string,
		serverName string,
		properties azuresqlshared.SQLDatabaseProperties) (*http.Response, error)

	DeleteDB(ctx context.Context,
		resourceGroupName string,
		serverName string,
		databaseName string) (result autorest.Response, err error)

	GetDB(ctx context.Context,
		resourceGroupName string,
		serverName string, databaseName string) (sql.Database, error)

	GetServer(ctx context.Context,
		resourceGroupName string,
		serverName string) (result sql.Server, err error)

	resourcemanager.ARMClient
}
