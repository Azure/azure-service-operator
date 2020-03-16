// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlshared

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// GetGoDbClient retrieves a DatabasesClient
func GetGoDbClient() sql.DatabasesClient {
	dbClient := sql.NewDatabasesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	dbClient.Authorizer = a
	dbClient.AddToUserAgent(config.UserAgent())
	return dbClient
}

// GetGoServersClient retrieves a ServersClient
func GetGoServersClient() sql.ServersClient {
	serversClient := sql.NewServersClientWithBaseURI(config.BaseURI(), config.SubscriptionID())

	a, _ := iam.GetResourceManagementAuthorizer()
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient
}

// GetGoFailoverGroupsClient retrieves a FailoverGroupsClient
func GetGoFailoverGroupsClient() sql.FailoverGroupsClient {
	failoverGroupsClient := sql.NewFailoverGroupsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	failoverGroupsClient.Authorizer = a
	failoverGroupsClient.AddToUserAgent(config.UserAgent())
	return failoverGroupsClient
}

// GetGoFirewallClient retrieves a FirewallRulesClient
func GetGoFirewallClient() sql.FirewallRulesClient {
	firewallClient := sql.NewFirewallRulesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	firewallClient.Authorizer = a
	firewallClient.AddToUserAgent(config.UserAgent())
	return firewallClient
}
