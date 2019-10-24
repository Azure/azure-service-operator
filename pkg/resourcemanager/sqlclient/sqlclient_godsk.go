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
func (sdk *GoSDKClient) CreateOrUpdateSQLServer(ctx context.Context, properties SQLServerProperties) (result sql.Server, err error) {
	serversClient := getGoServersClient()
	serverProp := SQLServerPropertiesToServer(properties)

	// issue the creation
	future, err := serversClient.CreateOrUpdate(
		ctx,
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
func (sdk *GoSDKClient) CreateOrUpdateSQLFirewallRule(ctx context.Context, ruleName string, startIP string, endIP string) (result bool, err error) {

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := sdk.GetServer(ctx)
	if err != nil || *server.State != "Ready" {
		return false, err
	}

	firewallClient := getGoFirewallClient()
	_, err = firewallClient.CreateOrUpdate(
		ctx,
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
func (sdk *GoSDKClient) CreateOrUpdateDB(ctx context.Context, properties SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error) {
	dbClient := getGoDbClient()
	dbProp := SQLDatabasePropertiesToDatabase(properties)

	return dbClient.CreateOrUpdate(
		ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		properties.DatabaseName,
		sql.Database{
			Location:           to.StringPtr(sdk.Location),
			DatabaseProperties: &dbProp,
		})
}

// CreateOrUpdateFailoverGroup creates a failover group
func (sdk *GoSDKClient) CreateOrUpdateFailoverGroup(ctx context.Context, failovergroupname string, properties SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error) {
	failoverGroupsClient := getGoFailoverGroupsClient()

	// Construct a PartnerInfo object from the server name
	// Get resource ID from the servername to use
	secServerSDKClient := GoSDKClient{
		ResourceGroupName: properties.SecondaryServerResourceGroup,
		ServerName:        properties.SecondaryServerName,
		Location:          "", // We dont get the location from the user for the secondary server as it is not required
	}
	server, err := secServerSDKClient.GetServer(ctx)
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
		database, err := sdk.GetDB(ctx, each)
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
		ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		failovergroupname,
		failoverGroup)

}

// GetServer returns a SQL server
func (sdk *GoSDKClient) GetServer(ctx context.Context) (result sql.Server, err error) {
	serversClient := getGoServersClient()

	return serversClient.Get(
		ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
}

// GetSQLFirewallRule returns a firewall rule
func (sdk *GoSDKClient) GetSQLFirewallRule(ctx context.Context, ruleName string) (result sql.FirewallRule, err error) {
	firewallClient := getGoFirewallClient()

	return firewallClient.Get(
		ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		ruleName,
	)
}

// GetDB retrieves a database
func (sdk *GoSDKClient) GetDB(ctx context.Context, databaseName string) (sql.Database, error) {
	dbClient := getGoDbClient()

	return dbClient.Get(
		ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
}

// GetFailoverGroup retrieves a failover group
func (sdk *GoSDKClient) GetFailoverGroup(ctx context.Context, failovergroupname string) (sql.FailoverGroup, error) {
	failoverGroupsClient := getGoFailoverGroupsClient()

	return failoverGroupsClient.Get(
		ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		failovergroupname,
	)
}

// DeleteDB deletes a DB
func (sdk *GoSDKClient) DeleteDB(ctx context.Context, databaseName string) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := sdk.GetServer(ctx)
	if err != nil || *server.State != "Ready" {
		return result, nil
	}

	// check to see if the db exists, if it doesn't then short-circuit
	_, err = sdk.GetDB(ctx, databaseName)
	if err != nil {
		return result, nil
	}

	dbClient := getGoDbClient()
	result, err = dbClient.Delete(
		ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		databaseName,
	)

	return result, err
}

// DeleteSQLFirewallRule deletes a firewall rule
func (sdk *GoSDKClient) DeleteSQLFirewallRule(ctx context.Context, ruleName string) (err error) {

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := sdk.GetServer(ctx)
	if err != nil || *server.State != "Ready" {
		return nil
	}

	// check to see if the rule exists, if it doesn't then short-circuit
	_, err = sdk.GetSQLFirewallRule(ctx, ruleName)
	if err != nil {
		return nil
	}

	firewallClient := getGoFirewallClient()
	_, err = firewallClient.Delete(
		ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
		ruleName,
	)

	return err
}

// DeleteSQLServer deletes a SQL server
func (sdk *GoSDKClient) DeleteSQLServer(ctx context.Context) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	_, err = sdk.GetServer(ctx)
	if err != nil {
		return result, nil
	}

	serversClient := getGoServersClient()
	future, err := serversClient.Delete(
		ctx,
		sdk.ResourceGroupName,
		sdk.ServerName,
	)
	if err != nil {
		return result, err
	}

	return future.Result(serversClient)
}

// DeleteFailoverGroup deletes a failover group
func (sdk *GoSDKClient) DeleteFailoverGroup(ctx context.Context, failoverGroupName string) (result autorest.Response, err error) {

	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	_, err = sdk.GetServer(ctx)
	if err != nil {
		return result, nil
	}

	// check to see if the failover group exists, if it doesn't then short-circuit
	_, err = sdk.GetFailoverGroup(ctx, failoverGroupName)
	if err != nil {
		return result, nil
	}

	failoverGroupsClient := getGoFailoverGroupsClient()
	future, err := failoverGroupsClient.Delete(
		ctx,
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
func (sdk *GoSDKClient) CheckNameAvailability(ctx context.Context) (result AvailabilityResponse, err error) {
	serversClient := getGoServersClient()

	response, err := serversClient.CheckNameAvailability(
		ctx,
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
