// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest/to"
)

func getServersClient() sql.ServersClient {
	serversClient := sql.NewServersClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient
}

func getDbClient() sql.DatabasesClient {
	dbClient := sql.NewDatabasesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	dbClient.Authorizer = a
	dbClient.AddToUserAgent(config.UserAgent())
	return dbClient
}

// CreateOrUpdateSQLServerImpl creates a SQL server in Azure
func (sdk GoSDKClient) CreateOrUpdateSQLServerImpl(properties sql.ServerProperties) (result *string, err error) {
	serversClient := getServersClient()

	future, err := serversClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		sql.Server{
			Location:         to.StringPtr(config.Location()),
			ServerProperties: &properties,
		})
	if err != nil {
		return nil, fmt.Errorf("cannot create sql server: %v", err)
	}

	err = future.WaitForCompletionRef(sdk.Ctx, serversClient.Client)
	if err != nil {
		return nil, fmt.Errorf("cannot get the sql server create or update future response: %v", err)
	}

	server, err := future.Result(serversClient)
	if err != nil {
		return nil, fmt.Errorf("cannot get the sql server instance: %v", err)
	}

	return server.ServerProperties.State, nil
}

// CreateOrUpdateDBImpl creates or updates a DB in Azure
func (sdk GoSDKClient) CreateOrUpdateDBImpl(dbName string, properties sql.DatabaseProperties) (result *string, err error) {
	dbClient := getDbClient()

	future, err := dbClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		dbName,
		sql.Database{
			Location:           to.StringPtr(sdk.Location),
			DatabaseProperties: &properties,
		})
	if err != nil {
		return nil, fmt.Errorf("cannot create sql database: %v", err)
	}

	err = future.WaitForCompletionRef(sdk.Ctx, dbClient.Client)
	if err != nil {
		return nil, fmt.Errorf("cannot get the sql database create or update future response: %v", err)
	}

	db, err := future.Result(dbClient)
	if err != nil {
		return nil, fmt.Errorf("cannot get the db instance: %v", err)
	}

	return db.DatabaseProperties.Status, nil
}

// DeleteDBImpl deletes a DB
func (sdk GoSDKClient) DeleteDBImpl(dbName string) (result bool, err error) {
	dbClient := getDbClient()

	_, err = dbClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		dbName,
	)
	if err != nil {
		return false, fmt.Errorf("cannot delete db: %v", err)
	}

	return true, nil
}

// DeleteSQLServerImpl deletes a DB
func (sdk GoSDKClient) DeleteSQLServerImpl() (result bool, err error) {
	serversClient := getServersClient()

	future, err := serversClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
	if err != nil {
		return false, fmt.Errorf("cannot delete sql server: %v", err)
	}

	err = future.WaitForCompletionRef(sdk.Ctx, serversClient.Client)
	if err != nil {
		return false, fmt.Errorf("cannot get the sql server delete future response: %v", err)
	}

	return true, nil
}
