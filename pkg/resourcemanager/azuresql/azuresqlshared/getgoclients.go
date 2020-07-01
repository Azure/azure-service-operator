// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlshared

import (
	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	sql3 "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
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

func GetGoDbClientWithCreds(creds map[string]string) (sql.DatabasesClient, error) {
	dbClient := sql.NewDatabasesClientWithBaseURI(config.BaseURI(), creds["AZURE_SUBSCRIPTION_ID"])
	a, err := iam.GetResourceManagementAuthorizerWithCreds(creds)
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

// GetGoServersClient retrieves a ServersClient
func GetGoServersClientWithCreds(creds map[string]string) (sql.ServersClient, error) {
	serversClient := sql.NewServersClientWithBaseURI(config.BaseURI(), creds["AZURE_SUBSCRIPTION_ID"])
	a, err := iam.GetResourceManagementAuthorizerWithCreds(creds)
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

// GetGoFailoverGroupsClientWithCreds retrieves a FailoverGroupsClient
func GetGoFailoverGroupsClientWithCreds(creds map[string]string) (sql.FailoverGroupsClient, error) {
	failoverGroupsClient := sql.NewFailoverGroupsClientWithBaseURI(config.BaseURI(), creds["AZURE_SUBSCRIPTION_ID"])
	a, err := iam.GetResourceManagementAuthorizerWithCreds(creds)
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

// GetGoFirewallClientWithCreds retrieves a FirewallRulesClient
func GetGoFirewallClientWithCreds(creds map[string]string) (sql.FirewallRulesClient, error) {
	firewallClient := sql.NewFirewallRulesClientWithBaseURI(config.BaseURI(), creds["AZURE_SUBSCRIPTION_ID"])
	a, err := iam.GetResourceManagementAuthorizerWithCreds(creds)
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
func GetGoVNetRulesClientWithCreds(creds map[string]string) (sql.VirtualNetworkRulesClient, error) {
	VNetRulesClient := sql.NewVirtualNetworkRulesClientWithBaseURI(config.BaseURI(), creds["AZURE_SUBSCRIPTION_ID"])
	a, err := iam.GetResourceManagementAuthorizerWithCreds(creds)
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

// GetGoNetworkSubnetClientWithCreds retrieves a Subnetclient
func GetGoNetworkSubnetClientWithCreds(creds map[string]string) (network.SubnetsClient, error) {
	SubnetsClient := network.NewSubnetsClientWithBaseURI(config.BaseURI(), creds["AZURE_SUBSCRIPTION_ID"])
	a, err := iam.GetResourceManagementAuthorizerWithCreds(creds)
	if err != nil {
		return network.SubnetsClient{}, err
	}
	SubnetsClient.Authorizer = a
	SubnetsClient.AddToUserAgent(config.UserAgent())
	return SubnetsClient, nil
}

// GetBackupLongTermRetentionPoliciesClient retrieves a Subnetclient
func GetBackupLongTermRetentionPoliciesClientWithCreds(creds map[string]string) (sql3.BackupLongTermRetentionPoliciesClient, error) {
	BackupClient := sql3.NewBackupLongTermRetentionPoliciesClientWithBaseURI(config.BaseURI(), creds["AZURE_SUBSCRIPTION_ID"])
	a, err := iam.GetResourceManagementAuthorizerWithCreds(creds)
	if err != nil {
		return sql3.BackupLongTermRetentionPoliciesClient{}, err
	}
	BackupClient.Authorizer = a
	BackupClient.AddToUserAgent(config.UserAgent())
	return BackupClient, nil
}
