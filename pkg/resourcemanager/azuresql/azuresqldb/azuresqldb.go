// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqldb

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
)

type AzureSqlDbManager struct {
	Log logr.Logger
}

func NewAzureSqlDbManager(log logr.Logger) *AzureSqlDbManager {
	return &AzureSqlDbManager{Log: log}
}

// GetServer returns a SQL server
func (_ *AzureSqlDbManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient := azuresqlshared.GetGoServersClient()

	return serversClient.Get(
		ctx,
		resourceGroupName,
		serverName,
	)
}

// GetDB retrieves a database
func (_ *AzureSqlDbManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error) {
	dbClient := azuresqlshared.GetGoDbClient()

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
}

// DeleteDB deletes a DB
func (sdk *AzureSqlDbManager) DeleteDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := sdk.GetServer(ctx, resourceGroupName, serverName)
	if err != nil || *server.State != "Ready" {
		return result, nil
	}

	// check to see if the db exists, if it doesn't then short-circuit
	_, err = sdk.GetDB(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		return result, nil
	}

	dbClient := azuresqlshared.GetGoDbClient()
	result, err = dbClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)

	return result, err
}

// CreateOrUpdateDB creates or updates a DB in Azure
func (_ *AzureSqlDbManager) CreateOrUpdateDB(ctx context.Context, resourceGroupName string, location string, serverName string, properties azuresqlshared.SQLDatabaseProperties) (*http.Response, error) {
	dbClient := azuresqlshared.GetGoDbClient()
	dbProp := azuresqlshared.SQLDatabasePropertiesToDatabase(properties)

	future, err := dbClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		properties.DatabaseName,
		sql.Database{
			Location:           to.StringPtr(location),
			DatabaseProperties: &dbProp,
		})

	return future.Response(), err
}
