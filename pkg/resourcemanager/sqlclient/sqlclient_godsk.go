package sqlclient

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

const typeOfService = "Microsoft.Sql/servers"

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

// getGoFailoverGroupsClient retrieves a FailoverGroupsClient
func getGoFailoverGroupsClient() sql.FailoverGroupsClient {
	failoverGroupsClient := sql.NewFailoverGroupsClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	failoverGroupsClient.Authorizer = a
	failoverGroupsClient.AddToUserAgent(config.UserAgent())
	return failoverGroupsClient
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

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := sdk.GetServer()
	if err != nil || *server.State != "Ready" {
		return false, err
	}

	firewallClient := getGoFirewallClient()
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

// CreateOrUpdateFailoverGroup creates a failover group
func (sdk GoSDKClient) CreateOrUpdateFailoverGroup(failovergroupname string, properties SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error) {
	failoverGroupsClient := getGoFailoverGroupsClient()
	serversClient := getGoServersClient()

	// Construct a PartnerInfo object from the server name
	// Get resource ID from the servername to use
	server, err := serversClient.Get(
		context.Background(),
		properties.SecondaryServerResourceGroup,
		properties.SecondaryServerName,
	)
	if err != nil {
		return result, nil
	}

	secServerResourceID := server.ID
	partnerServerInfo := sql.PartnerInfo{
		ID:              secServerResourceID,
		ReplicationRole: sql.Secondary,
	}

	partnerServerInfoArray := []sql.PartnerInfo{partnerServerInfo}

	var databaseIDArray []string

	// Parse the Databases in the Databaselist and form array of Resource IDs
	for _, each := range properties.DatabaseList {
		database, err := sdk.GetDB(each)
		if err != nil {
			return result, err
		}
		databaseIDArray = append(databaseIDArray, *database.ID)
	}

	// Construct FailoverGroupProperties struct
	failoverGroupProperties := sql.FailoverGroupProperties{
		ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpoint{
			FailoverPolicy:                         properties.FailoverPolicy,
			FailoverWithDataLossGracePeriodMinutes: &properties.FailoverGracePeriod,
		},
		PartnerServers: &partnerServerInfoArray,
		Databases:      &databaseIDArray,
	}

	failoverGroup := sql.FailoverGroup{
		FailoverGroupProperties: &failoverGroupProperties,
	}

	return failoverGroupsClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		failovergroupname,
		failoverGroup)

}

// GetSQLFirewallRule returns a firewall rule
func (sdk GoSDKClient) GetSQLFirewallRule(ruleName string) (result sql.FirewallRule, err error) {
	firewallClient := getGoFirewallClient()

	return firewallClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		ruleName,
	)
}

// GetDB retrieves a database
func (sdk GoSDKClient) GetDB(databaseName string) (sql.Database, error) {
	dbClient := getGoDbClient()

	return dbClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
}

// DeleteDB deletes a DB
func (sdk GoSDKClient) DeleteDB(databaseName string) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := sdk.GetServer()
	if err != nil || *server.State != "Ready" {
		return result, nil
	}

	// check to see if the db exists, if it doesn't then short-circuit
	_, err = sdk.GetDB(databaseName)
	if err != nil {
		return result, nil
	}

	dbClient := getGoDbClient()
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

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := sdk.GetServer()
	if err != nil || *server.State != "Ready" {
		return nil
	}

	// check to see if the rule exists, if it doesn't then short-circuit
	_, err = sdk.GetSQLFirewallRule(ruleName)
	if err != nil {
		return nil
	}

	firewallClient := getGoFirewallClient()
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
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	_, err = sdk.GetServer()
	if err != nil {
		return result, nil
	}

	serversClient := getGoServersClient()
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

// GetFailoverGroup retrieves a failover group
func (sdk GoSDKClient) GetFailoverGroup(failovergroupname string) (sql.FailoverGroup, error) {
	failoverGroupsClient := getGoFailoverGroupsClient()

	return failoverGroupsClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		failovergroupname,
	)
}

// DeleteFailoverGroup deletes a failover group
func (sdk GoSDKClient) DeleteFailoverGroup(failoverGroupName string) (result autorest.Response, err error) {

	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the failover group exists, if it doesn't then short-circuit
	_, err = sdk.GetFailoverGroup(failoverGroupName)
	if err != nil {
		return result, nil
	}

	failoverGroupsClient := getGoFailoverGroupsClient()
	future, err := failoverGroupsClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		failoverGroupName,
	)
	if err != nil {
		return result, err
	}

	return future.Result(failoverGroupsClient)
}

// CheckNameAvailability determines whether a SQL resource can be created with the specified name
func (sdk GoSDKClient) CheckNameAvailability() (result AvailabilityResponse, err error) {
	serversClient := getGoServersClient()

	response, err := serversClient.CheckNameAvailability(
		sdk.Ctx,
		sql.CheckNameAvailabilityRequest{
			Name: to.StringPtr(sdk.ServerName),
			Type: to.StringPtr(typeOfService),
		},
	)
	if err != nil {
		return result, err
	}

	return ToAvailabilityResponse(response), err
}

// GetServer returns a SQL server
func (sdk GoSDKClient) GetServer() (result sql.Server, err error) {
	serversClient := getGoServersClient()

	return serversClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
}
