// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresql

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	helpers "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest"
)

// MockGoSDKClient struct
type MockGoSDKClient struct {
	Server                             sql.Server
	Database                           sql.Database
	FirewallRule                       sql.FirewallRule
	FailoverGroup                      sql.FailoverGroup
	DatabasesCreateOrUpdateFuture      sql.DatabasesCreateOrUpdateFuture
	FailoverGroupsCreateOrUpdateFuture sql.FailoverGroupsCreateOrUpdateFuture
}

// CreateOrUpdateSQLServer creates a new sql server
func (sdk *MockGoSDKClient) CreateOrUpdateSQLServer(ctx context.Context, resourceGroupName string, location string, serverName string, properties azuresqlshared.SQLServerProperties, forceUpdate bool) (result sql.Server, err error) {
	var sqlServer = sql.Server{
		Response: helpers.GetRestResponse(http.StatusCreated),
	}

	sdk.Server = sqlServer

	return sqlServer, nil
}

//DeleteSQLServer return StatusOK
func (sdk *MockGoSDKClient) DeleteSQLServer(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error) {

	return helpers.GetRestResponse(http.StatusOK), nil
}

//GetServer get server
func (sdk *MockGoSDKClient) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {

	state := "Ready"
	serverProperties := sql.ServerProperties{State: &state}
	var sqlServer = sql.Server{
		Response:         helpers.GetRestResponse(http.StatusCreated),
		ServerProperties: &serverProperties,
	}

	sdk.Server = sqlServer

	return sqlServer, nil
}

//CreateOrUpdateSQLFirewallRule create or
func (sdk *MockGoSDKClient) CreateOrUpdateSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string, startIP string, endIP string) (result bool, err error) {

	return true, nil
}

//DeleteSQLFirewallRule delete sql firewall
func (sdk *MockGoSDKClient) DeleteSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error) {
	return nil
}

//DeleteDB delete database
func (sdk *MockGoSDKClient) DeleteDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result autorest.Response, err error) {

	return helpers.GetRestResponse(http.StatusOK), nil
}

//GetSQLFirewallRule get sql firewall rule
func (sdk *MockGoSDKClient) GetSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (result sql.FirewallRule, err error) {

	var sqlFirewallRule = sql.FirewallRule{
		Response: helpers.GetRestResponse(http.StatusCreated),
	}

	sdk.FirewallRule = sqlFirewallRule

	return sqlFirewallRule, nil
}

//GetDB get database
func (sdk *MockGoSDKClient) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error) {

	var sqlDatabase = sql.Database{
		Response: helpers.GetRestResponse(http.StatusCreated),
	}

	sdk.Database = sqlDatabase

	return sqlDatabase, nil
}

//CreateOrUpdateDB create or update DB
func (sdk *MockGoSDKClient) CreateOrUpdateDB(ctx context.Context, resourceGroupName string, location string, serverName string, properties azuresqlshared.SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error) {

	var sqlDatabasesCreateOrUpdateFuture = sql.DatabasesCreateOrUpdateFuture{}

	return sqlDatabasesCreateOrUpdateFuture, nil
}

//CreateOrUpdateFailoverGroup create or update failover group
func (sdk *MockGoSDKClient) CreateOrUpdateFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string, properties azuresqlshared.SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error) {

	var sqlFailoverGroupsCreateOrUpdateFuture = sql.FailoverGroupsCreateOrUpdateFuture{}
	sdk.FailoverGroupsCreateOrUpdateFuture = sqlFailoverGroupsCreateOrUpdateFuture

	return sqlFailoverGroupsCreateOrUpdateFuture, nil

}

//DeleteFailoverGroup delete fail over group
func (sdk *MockGoSDKClient) DeleteFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result autorest.Response, err error) {

	return helpers.GetRestResponse(http.StatusOK), nil
}

//DeleteFailoverGroup delete fail over group
func (sdk *MockGoSDKClient) GetFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string) (sql.FailoverGroup, error) {
	var sqlFailovergroup = sql.FailoverGroup{
		Response: helpers.GetRestResponse(http.StatusCreated),
	}

	sdk.FailoverGroup = sqlFailovergroup

	return sqlFailovergroup, nil
}
