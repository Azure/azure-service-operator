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

func NewDBClient() sql.DatabasesClient {
	return getGoDbClient()
}

// getGoFirewallClient retrieves a FirewallRulesClient
func getGoFirewallClient() sql.FirewallRulesClient {
	firewallClient := sql.NewFirewallRulesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	firewallClient.Authorizer = a
	firewallClient.AddToUserAgent(config.UserAgent())
	return firewallClient
}

// CreateOrUpdateSQLServer creates a SQL server in Azure
func (sdk GoSDKClient) CreateOrUpdateSQLServer(properties SQLServerProperties) (result sql.Server, err error) {
	serversClient := getGoServersClient()
	serverProp := SQLServerPropertiesToServer(properties)

	// issue the creation
	future, err := serversClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		sql.Server{
			Location:         to.StringPtr(sdk.Location),
			ServerProperties: &serverProp,
		})
	if err != nil {
		return result, err
	}

	return future.Result(serversClient)
}

// CreateOrUpdateSQLFirewallRule creates or updates a firewall rule
// based on code from: https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/sql/sql.go#L111
// to allow allow Azure services to connect example: https://docs.microsoft.com/en-us/azure/sql-database/sql-database-firewall-configure#manage-firewall-rules-using-azure-cli
func (sdk GoSDKClient) CreateOrUpdateSQLFirewallRule(ruleName string, startIP string, endIP string) (result bool, err error) {
	serversClient := getGoServersClient()
	firewallClient := getGoFirewallClient()

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := serversClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
	if err != nil || *server.State != "Ready" {
		return false, err
	}

	_, err = firewallClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		ruleName,
		sql.FirewallRule{
			FirewallRuleProperties: &sql.FirewallRuleProperties{
				StartIPAddress: to.StringPtr(startIP),
				EndIPAddress:   to.StringPtr(endIP),
			},
		},
	)
	result = false
	if err == nil {
		result = true
	}

	return result, err
}

// CreateOrUpdateDB creates or updates a DB in Azure
func (sdk GoSDKClient) CreateOrUpdateDB(properties SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error) {
	dbClient := getGoDbClient()
	dbProp := SQLDatabasePropertiesToDatabase(properties)

	return dbClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		properties.DatabaseName,
		sql.Database{
			Location:           to.StringPtr(sdk.Location),
			DatabaseProperties: &dbProp,
		})

}

// CreateOrUpdateDB creates or updates a DB in Azure
func (sdk GoSDKClient) GetDB(name string) (sql.Database, error) {
	dbClient := getGoDbClient()

	return dbClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		name,
		"serviceTierAdvisors, transparentDataEncryption",
	)
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
				StatusCode: 200,
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

// DeleteSQLFirewallRule deletes a firewall rule
func (sdk GoSDKClient) DeleteSQLFirewallRule(ruleName string) (err error) {
	serversClient := getGoServersClient()
	firewallClient := getGoFirewallClient()

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := serversClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
	if err != nil || *server.State != "Ready" {
		return err
	}

	_, err = firewallClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		ruleName,
	)

	return err
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
				StatusCode: 200,
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

// GetServer returns a server
func (sdk GoSDKClient) GetServer() (sql.Server, error) {
	serversClient := getGoServersClient()

	return serversClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
}

func (sdk GoSDKClient) CheckNameAvailablity() (result AvailabilityResponse, err error) {
	serversClient := getGoServersClient()
	typeOfService := "Microsoft.Sql/servers"

	response, err := serversClient.CheckNameAvailability(
		sdk.Ctx,
		sql.CheckNameAvailabilityRequest {
			Name: to.StringPtr(sdk.ServerName),
			Type: to.StringPtr(typeOfService),
		},
	)
	if err != nil {
		return result, err
	}

	return ToAvailabilityResponse(response), err
}