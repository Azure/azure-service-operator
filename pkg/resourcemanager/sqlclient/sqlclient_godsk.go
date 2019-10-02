// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

// getGoServersClient retrieves a ServersClient
func getGoServersClient() sql.ServersClient {
	serversClient := sql.NewServersClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient
}

// getGoDbClient retrieves a DatabasesClient
func getGoDbClient() sql.DatabasesClient {
	dbClient := sql.NewDatabasesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	dbClient.Authorizer = a
	dbClient.AddToUserAgent(config.UserAgent())
	return dbClient
}

// CreateOrUpdateSQLServer creates a SQL server in Azure
func (sdk GoSDKClient) CreateOrUpdateSQLServer(properties SQLServerProperties) (result sql.Server, err error) {
	serversClient := getGoServersClient()
	serverProp := SQLServerPropertiesToServer(properties)

	// check to see if the server exists, if it does then short-circuit
	server, err := serversClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
	if err == nil && *server.State == "Ready" {
		return server, nil
	}

	// issue the creation
	future, err := serversClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		sql.Server{
			Location:         to.StringPtr(config.Location()),
			ServerProperties: &serverProp,
		})
	if err != nil {
		return result, err
	}

	return future.Result(serversClient)
}

// CreateOrUpdateDB creates or updates a DB in Azure
func (sdk GoSDKClient) CreateOrUpdateDB(properties SQLDatabaseProperties) (result sql.Database, err error) {
	dbClient := getGoDbClient()
	dbProp := SQLDatabasePropertiesToDatabase(properties)

	// check to see if the db exists, if it does then short-circuit
	db, err := dbClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		properties.DatabaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
	if err == nil && *db.Status == "Online" {
		return db, nil
	}

	future, err := dbClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		properties.DatabaseName,
		sql.Database{
			Location:           to.StringPtr(sdk.Location),
			DatabaseProperties: &dbProp,
		})
	if err != nil {
		return result, err
	}

	// TODO: Will needs to remove the sync
	err = future.WaitForCompletionRef(
		sdk.Ctx,
		dbClient.Client,
	)
	if err != nil {
		return result, err
	}

	return future.Result(dbClient)
}

// DeleteDB deletes a DB
func (sdk GoSDKClient) DeleteDB(databaseName string) (result autorest.Response, err error) {
	dbClient := getGoDbClient()

	// check to see if the db exists, if it does then short-circuit
	_, err = dbClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
	if err != nil {
		result = autorest.Response{
			Response: &http.Response{
				StatusCode: http.StatusOK,
			},
		}
		return result, nil
	}

	result, err = dbClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		databaseName,
	)

	return result, err
}

// DeleteSQLServer deletes a SQL server
func (sdk GoSDKClient) DeleteSQLServer() (result autorest.Response, err error) {
	serversClient := getGoServersClient()

	// check to see if the server exists, if it doesn't then short-circuit
	_, err = serversClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
	if err != nil {
		result = autorest.Response{
			Response: &http.Response{
				StatusCode: http.StatusOK,
			},
		}
		return result, nil
	}

	future, err := serversClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
	if err != nil {
		return result, err
	}

	return future.Result(serversClient)
}

// IsAsyncNotCompleted returns true if the error is due to async not completed
func (sdk GoSDKClient) IsAsyncNotCompleted(err error) (result bool) {
	result = false
	if err != nil && strings.Contains(err.Error(), "asynchronous operation has not completed") {
		result = true
	} else if strings.Contains(err.Error(), "is busy with another operation") {
		result = true
	}
	return result
}
