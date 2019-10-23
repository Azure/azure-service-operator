/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sqlclient

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	helpers "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	sqlclient "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	"github.com/Azure/go-autorest/autorest"
)

// MockGoSDKClient struct
type MockGoSDKClient struct {
	sqlServer                             sql.Server
	sqlDatabase                           sql.Database
	sqlFirewallRule                       sql.FirewallRule
	sqlDatabasesCreateOrUpdateFuture      sql.DatabasesCreateOrUpdateFuture
	sqlFailoverGroupsCreateOrUpdateFuture sql.FailoverGroupsCreateOrUpdateFuture
}

// CreateOrUpdateSQLServer creates a new sql server
func (sdk *MockGoSDKClient) CreateOrUpdateSQLServer(ctx context.Context, resourceGroupName string, location string, serverName string, properties sqlclient.SQLServerProperties) (result sql.Server, err error) {
	var sqlServer = sql.Server{
		Response: helpers.GetRestResponse(http.StatusCreated),
	}

	sdk.sqlServer = sqlServer

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

	sdk.sqlServer = sqlServer

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

	sdk.sqlFirewallRule = sqlFirewallRule

	return sqlFirewallRule, nil
}

//GetDB get database
func (sdk *MockGoSDKClient) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error) {

	var sqlDatabase = sql.Database{
		Response: helpers.GetRestResponse(http.StatusCreated),
	}

	sdk.sqlDatabase = sqlDatabase

	return sqlDatabase, nil
}

//CreateOrUpdateDB create or update DB
func (sdk *MockGoSDKClient) CreateOrUpdateDB(ctx context.Context, resourceGroupName string, location string, serverName string, properties sqlclient.SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error) {

	var sqlDatabasesCreateOrUpdateFuture = sql.DatabasesCreateOrUpdateFuture{}

	return sqlDatabasesCreateOrUpdateFuture, nil
}

//CreateOrUpdateFailoverGroup create or update failover group
func (sdk *MockGoSDKClient) CreateOrUpdateFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string, properties sqlclient.SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error) {

	var sqlFailoverGroupsCreateOrUpdateFuture = sql.FailoverGroupsCreateOrUpdateFuture{}
	sdk.sqlFailoverGroupsCreateOrUpdateFuture = sqlFailoverGroupsCreateOrUpdateFuture

	return sqlFailoverGroupsCreateOrUpdateFuture, nil

}

//DeleteFailoverGroup delete fail over group
func (sdk *MockGoSDKClient) DeleteFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result autorest.Response, err error) {

	return helpers.GetRestResponse(http.StatusOK), nil
}
