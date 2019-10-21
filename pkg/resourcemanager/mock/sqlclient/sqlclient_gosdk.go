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
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	helpers "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	sqlclient "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	"github.com/Azure/go-autorest/autorest"
)

// MockSqlManager struct
type MockSqlManager struct {
	sqlServer                             sql.Server
	sqlDatabase                           sql.Database
	sqlFirewallRule                       sql.FirewallRule
	sqlDatabasesCreateOrUpdateFuture      sql.DatabasesCreateOrUpdateFuture
	sqlFailoverGroupsCreateOrUpdateFuture sql.FailoverGroupsCreateOrUpdateFuture
}

// CreateOrUpdateSQLServer creates a new sql server
func (manager *MockSqlManager) CreateOrUpdateSQLServer(sdkClient sqlclient.GoSDKClient, properties sqlclient.SQLServerProperties) (result sql.Server, err error) {
	var sqlServer = sql.Server{
		Response: helpers.GetRestResponse(http.StatusCreated),
	}

	manager.sqlServer = sqlServer

	return sqlServer, nil
}

//DeleteSQLServer return StatusOK
func (manager *MockSqlManager) DeleteSQLServer(sdkClient sqlclient.GoSDKClient) (result autorest.Response, err error) {

	return helpers.GetRestResponse(http.StatusOK), nil
}

//GetServer get server
func (manager *MockSqlManager) GetServer(sdkClient sqlclient.GoSDKClient) (result sql.Server, err error) {

	state := "Ready"
	serverProperties := sql.ServerProperties{State: &state}
	var sqlServer = sql.Server{
		Response:         helpers.GetRestResponse(http.StatusCreated),
		ServerProperties: &serverProperties,
	}

	manager.sqlServer = sqlServer

	return sqlServer, nil
}

//CreateOrUpdateSQLFirewallRule create or
func (manager *MockSqlManager) CreateOrUpdateSQLFirewallRule(sdkClient sqlclient.GoSDKClient, ruleName string, startIP string, endIP string) (result bool, err error) {

	return true, nil
}

//DeleteSQLFirewallRule delete sql firewall
func (manager *MockSqlManager) DeleteSQLFirewallRule(sdkClient sqlclient.GoSDKClient, ruleName string) (err error) {
	return nil
}

//DeleteDB delete database
func (manager *MockSqlManager) DeleteDB(sdkClient sqlclient.GoSDKClient, databaseName string) (result autorest.Response, err error) {

	return helpers.GetRestResponse(http.StatusOK), nil
}

//GetSQLFirewallRule get sql firewall rule
func (manager *MockSqlManager) GetSQLFirewallRule(sdkClient sqlclient.GoSDKClient, ruleName string) (result sql.FirewallRule, err error) {

	var sqlFirewallRule = sql.FirewallRule{
		Response: helpers.GetRestResponse(http.StatusCreated),
	}

	manager.sqlFirewallRule = sqlFirewallRule

	return sqlFirewallRule, nil
}

//GetDB get database
func (manager *MockSqlManager) GetDB(sdkClient sqlclient.GoSDKClient, databaseName string) (sql.Database, error) {

	var sqlDatabase = sql.Database{
		Response: helpers.GetRestResponse(http.StatusCreated),
	}

	manager.sqlDatabase = sqlDatabase

	return sqlDatabase, nil
}

//CreateOrUpdateDB create or update DB
func (manager *MockSqlManager) CreateOrUpdateDB(sdkClient sqlclient.GoSDKClient, properties sqlclient.SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error) {

	var sqlDatabasesCreateOrUpdateFuture = sql.DatabasesCreateOrUpdateFuture{}
	manager.sqlDatabasesCreateOrUpdateFuture = sqlDatabasesCreateOrUpdateFuture

	return sqlDatabasesCreateOrUpdateFuture, nil
}

//CreateOrUpdateFailoverGroup create or update failover group
func (manager *MockSqlManager) CreateOrUpdateFailoverGroup(sdkClient sqlclient.GoSDKClient, failovergroupname string, properties sqlclient.SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error) {

	var sqlFailoverGroupsCreateOrUpdateFuture = sql.FailoverGroupsCreateOrUpdateFuture{}
	manager.sqlFailoverGroupsCreateOrUpdateFuture = sqlFailoverGroupsCreateOrUpdateFuture

	return sqlFailoverGroupsCreateOrUpdateFuture, nil

}

//DeleteFailoverGroup delete fail over group
func (manager *MockSqlManager) DeleteFailoverGroup(sdkClient sqlclient.GoSDKClient, failoverGroupName string) (result autorest.Response, err error) {

	return helpers.GetRestResponse(http.StatusOK), nil
}
