package sqlclient

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

const typeOfService = "Microsoft.Sql/servers"

type azureSqlManager struct{}

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

// getGoFirewallClient retrieves a FirewallRulesClient
func getGoFirewallClient() sql.FirewallRulesClient {
	firewallClient := sql.NewFirewallRulesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	firewallClient.Authorizer = a
	firewallClient.AddToUserAgent(config.UserAgent())
	return firewallClient
}

// CreateOrUpdateSQLServer creates a SQL server in Azure
func (*azureSqlManager) CreateOrUpdateSQLServer(sdkClient GoSDKClient, properties SQLServerProperties) (result sql.Server, err error) {
	serversClient := getGoServersClient()
	serverProp := SQLServerPropertiesToServer(properties)

	// issue the creation
	future, err := serversClient.CreateOrUpdate(
		sdkClient.Ctx,
		sdkClient.ResourceGroupName,
		sdkClient.ServerName,
		sql.Server{
			Location:         to.StringPtr(sdkClient.Location),
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
func (manager *azureSqlManager) CreateOrUpdateSQLFirewallRule(sdkClient GoSDKClient, ruleName string, startIP string, endIP string) (result bool, err error) {

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := manager.GetServer(sdkClient)
	if err != nil || *server.State != "Ready" {
		return false, err
	}

	firewallClient := getGoFirewallClient()
	_, err = firewallClient.CreateOrUpdate(
		sdkClient.Ctx,
		sdkClient.ResourceGroupName,
		sdkClient.ServerName,
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
func (manager *azureSqlManager) CreateOrUpdateDB(sdkClient GoSDKClient, properties SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error) {
	dbClient := getGoDbClient()
	dbProp := SQLDatabasePropertiesToDatabase(properties)

	return dbClient.CreateOrUpdate(
		sdkClient.Ctx,
		sdkClient.ResourceGroupName,
		sdkClient.ServerName,
		properties.DatabaseName,
		sql.Database{
			Location:           to.StringPtr(sdkClient.Location),
			DatabaseProperties: &dbProp,
		})
}

// GetSQLFirewallRule returns a firewall rule
func (manager *azureSqlManager) GetSQLFirewallRule(sdkClient GoSDKClient, ruleName string) (result sql.FirewallRule, err error) {
	firewallClient := getGoFirewallClient()

	return firewallClient.Get(
		sdkClient.Ctx,
		sdkClient.ResourceGroupName,
		sdkClient.ServerName,
		ruleName,
	)
}

// GetDB retrieves a database
func (manager *azureSqlManager) GetDB(sdkClient GoSDKClient, databaseName string) (sql.Database, error) {
	dbClient := getGoDbClient()

	return dbClient.Get(
		sdkClient.Ctx,
		sdkClient.ResourceGroupName,
		sdkClient.ServerName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
}

// DeleteDB deletes a DB
func (manager *azureSqlManager) DeleteDB(sdkClient GoSDKClient, databaseName string) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := manager.GetServer(sdkClient)
	if err != nil || *server.State != "Ready" {
		return result, nil
	}

	// check to see if the db exists, if it doesn't then short-circuit
	_, err = manager.GetDB(sdkClient, databaseName)
	if err != nil {
		return result, nil
	}

	dbClient := getGoDbClient()
	result, err = dbClient.Delete(
		sdkClient.Ctx,
		sdkClient.ResourceGroupName,
		sdkClient.ServerName,
		databaseName,
	)

	return result, err
}

// DeleteSQLFirewallRule deletes a firewall rule
func (manager *azureSqlManager) DeleteSQLFirewallRule(sdkClient GoSDKClient, ruleName string) (err error) {

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := manager.GetServer(sdkClient)
	if err != nil || *server.State != "Ready" {
		return nil
	}

	// check to see if the rule exists, if it doesn't then short-circuit
	_, err = manager.GetSQLFirewallRule(sdkClient, ruleName)
	if err != nil {
		return nil
	}

	firewallClient := getGoFirewallClient()
	_, err = firewallClient.Delete(
		sdkClient.Ctx,
		sdkClient.ResourceGroupName,
		sdkClient.ServerName,
		ruleName,
	)

	return err
}

// DeleteSQLServer deletes a SQL server
func (manager *azureSqlManager) DeleteSQLServer(sdkClient GoSDKClient) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	_, err = manager.GetServer(sdkClient)
	if err != nil {
		return result, nil
	}

	serversClient := getGoServersClient()
	future, err := serversClient.Delete(
		sdkClient.Ctx,
		sdkClient.ResourceGroupName,
		sdkClient.ServerName,
	)
	if err != nil {
		return result, err
	}

	return future.Result(serversClient)
}

// CheckNameAvailability determines whether a SQL resource can be created with the specified name
func (manager *azureSqlManager) CheckNameAvailability(sdkClient GoSDKClient) (result AvailabilityResponse, err error) {
	serversClient := getGoServersClient()

	response, err := serversClient.CheckNameAvailability(
		sdkClient.Ctx,
		sql.CheckNameAvailabilityRequest{
			Name: to.StringPtr(sdkClient.ServerName),
			Type: to.StringPtr(typeOfService),
		},
	)
	if err != nil {
		return result, err
	}

	return ToAvailabilityResponse(response), err
}

// GetServer returns a SQL server
func (manager *azureSqlManager) GetServer(sdkClient GoSDKClient) (result sql.Server, err error) {
	serversClient := getGoServersClient()

	return serversClient.Get(
		sdkClient.Ctx,
		sdkClient.ResourceGroupName,
		sdkClient.ServerName,
	)
}
