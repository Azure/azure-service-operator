// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlshared

import (
	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// GetGoDbClient retrieves a DatabasesClient
func GetGoDbClient() (sql.DatabasesClient, error) {
	dbClient := sql.NewDatabasesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return sql.DatabasesClient{}, err
	}
	dbClient.Authorizer = a
	dbClient.AddToUserAgent(config.UserAgent())
	return dbClient, nil
}

// GetGoServersClient retrieves a ServersClient
func GetGoServersClient() (sql.ServersClient, error) {
	serversClient := sql.NewServersClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return sql.ServersClient{}, err
	}
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient, nil
}

// GetGoFailoverGroupsClient retrieves a FailoverGroupsClient
func GetGoFailoverGroupsClient() (sql.FailoverGroupsClient, error) {
	failoverGroupsClient := sql.NewFailoverGroupsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return sql.FailoverGroupsClient{}, err
	}
	failoverGroupsClient.Authorizer = a
	failoverGroupsClient.AddToUserAgent(config.UserAgent())
	return failoverGroupsClient, nil
}

// GetGoFirewallClient retrieves a FirewallRulesClient
func GetGoFirewallClient() (sql.FirewallRulesClient, error) {
	firewallClient := sql.NewFirewallRulesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return sql.FirewallRulesClient{}, err
	}
	firewallClient.Authorizer = a
	firewallClient.AddToUserAgent(config.UserAgent())
	return firewallClient, nil
}

// GetGoVNetRulesClient retrieves a VirtualNetworkRulesClient
func GetGoVNetRulesClient() (sql.VirtualNetworkRulesClient, error) {
	VNetRulesClient := sql.NewVirtualNetworkRulesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return sql.VirtualNetworkRulesClient{}, err
	}
	VNetRulesClient.Authorizer = a
	VNetRulesClient.AddToUserAgent(config.UserAgent())
	return VNetRulesClient, nil
}

// GetNetworkSubnetClient retrieves a Subnetclient
func GetGoNetworkSubnetClient() (network.SubnetsClient, error) {
	SubnetsClient := network.NewSubnetsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return network.SubnetsClient{}, err
	}
	SubnetsClient.Authorizer = a
	SubnetsClient.AddToUserAgent(config.UserAgent())
	return SubnetsClient, nil
}
