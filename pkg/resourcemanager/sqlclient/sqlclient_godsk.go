// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"fmt"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
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

// getGoCBClient retrieves a DatabasesClient
func getGoCBClient() sql.DatabasesClient {
	dbClient := sql.NewDatabasesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	dbClient.Authorizer = a
	dbClient.AddToUserAgent(config.UserAgent())
	return dbClient
}

// CreateOrUpdateSQLServer creates a SQL server in Azure
func (sdk GoSDKClient) CreateOrUpdateSQLServer(properties SQLServerProperties) (s sql.Server, err error) {
	serversClient := getGoServersClient()
	serverProp := SQLServerPropertiesToServer(properties)

	future, err := serversClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		sql.Server{
			Location:         to.StringPtr(sdk.Location),
			ServerProperties: &serverProp,
		})

	log.Println(future)

	return future.Result(serversClient)
}

// SQLServerReady returns true if the SQL server is active
func (sdk GoSDKClient) SQLServerReady() (result bool, err error) {
	serversClient := getGoServersClient()

	server, err := serversClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
	if err != nil {
		if strings.Contains(err.Error(), "ResourceNotFound") {
			return false, nil
		}
		return false, fmt.Errorf("cannot get sql server: %v", err)
	}

	return *server.State == "Ready", err
}

// CreateOrUpdateDB creates or updates a DB in Azure
func (sdk GoSDKClient) CreateOrUpdateDB(properties SQLDatabaseProperties) (result bool, err error) {
	dbClient := getGoCBClient()
	dbProp := SQLDatabasePropertiesToDatabase(properties)

	_, err = dbClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		properties.DatabaseName,
		sql.Database{
			Location:           to.StringPtr(sdk.Location),
			DatabaseProperties: &dbProp,
		})
	if err != nil {
		return false, fmt.Errorf("cannot create sql database: %v", err)
	}

	return true, nil
}

// DeleteDB deletes a DB
func (sdk GoSDKClient) DeleteDB(databaseName string) (result bool, err error) {
	dbClient := getGoCBClient()

	_, err = dbClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		databaseName,
	)
	if err != nil {
		return false, fmt.Errorf("cannot delete db: %v", err)
	}

	return true, nil
}

// DeleteSQLServer deletes a DB
func (sdk GoSDKClient) DeleteSQLServer() (result bool, err error) {
	serversClient := getGoServersClient()

	_, err = serversClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
	if err != nil {
		return false, fmt.Errorf("cannot delete sql server: %v", err)
	}

	return true, nil
}
