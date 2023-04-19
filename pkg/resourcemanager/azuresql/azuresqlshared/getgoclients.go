// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlshared

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	sql3 "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// GetGoDbClient retrieves a DatabasesClient
func GetGoDbClient(creds config.Credentials) (sql.DatabasesClient, error) {
	dbClient := sql.NewDatabasesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return sql.DatabasesClient{}, err
	}
	dbClient.Authorizer = a
	dbClient.AddToUserAgent(config.UserAgent())
	return dbClient, nil
}

// GetGoServersClient retrieves a ServersClient
func GetGoServersClient(creds config.Credentials) (sql.ServersClient, error) {
	serversClient := sql.NewServersClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return sql.ServersClient{}, err
	}
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient, nil
}

// GetGoFailoverGroupsClient retrieves a FailoverGroupsClient
func GetGoFailoverGroupsClient(creds config.Credentials) (sql.FailoverGroupsClient, error) {
	failoverGroupsClient := sql.NewFailoverGroupsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return sql.FailoverGroupsClient{}, err
	}
	failoverGroupsClient.Authorizer = a
	failoverGroupsClient.AddToUserAgent(config.UserAgent())
	return failoverGroupsClient, nil
}

// GetGoFirewallClient retrieves a FirewallRulesClient
func GetGoFirewallClient(creds config.Credentials) (sql.FirewallRulesClient, error) {
	firewallClient := sql.NewFirewallRulesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return sql.FirewallRulesClient{}, err
	}
	firewallClient.Authorizer = a
	firewallClient.AddToUserAgent(config.UserAgent())
	return firewallClient, nil
}

// GetGoVNetRulesClient retrieves a VirtualNetworkRulesClient
func GetGoVNetRulesClient(creds config.Credentials) (sql.VirtualNetworkRulesClient, error) {
	vnetRulesClient := sql.NewVirtualNetworkRulesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return sql.VirtualNetworkRulesClient{}, err
	}
	vnetRulesClient.Authorizer = a
	vnetRulesClient.AddToUserAgent(config.UserAgent())
	return vnetRulesClient, nil
}

// GetNetworkSubnetClient retrieves a Subnetclient
func GetGoNetworkSubnetClient(creds config.Credentials) (network.SubnetsClient, error) {
	subnetsClient := network.NewSubnetsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return network.SubnetsClient{}, err
	}
	subnetsClient.Authorizer = a
	subnetsClient.AddToUserAgent(config.UserAgent())
	return subnetsClient, nil
}

// GetBackupLongTermRetentionPoliciesClient retrieves a BackupLongTermRetentionPoliciesClient
func GetBackupLongTermRetentionPoliciesClient(creds config.Credentials) (sql3.BackupLongTermRetentionPoliciesClient, error) {
	backupClient := sql3.NewBackupLongTermRetentionPoliciesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return sql3.BackupLongTermRetentionPoliciesClient{}, err
	}
	backupClient.Authorizer = a
	backupClient.AddToUserAgent(config.UserAgent())
	return backupClient, nil
}

// GetBackupShortTermRetentionPoliciesClient retrieves a BackupShortTermRetentionPoliciesClient
func GetBackupShortTermRetentionPoliciesClient(creds config.Credentials) (sql3.BackupShortTermRetentionPoliciesClient, error) {
	backupClient := sql3.NewBackupShortTermRetentionPoliciesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return sql3.BackupShortTermRetentionPoliciesClient{}, err
	}
	backupClient.Authorizer = a
	backupClient.AddToUserAgent(config.UserAgent())
	return backupClient, nil
}

func GetSubscriptionCredentials(creds config.Credentials, subscriptionID string) config.Credentials {
	if subscriptionID == "" {
		return creds
	}

	return creds.WithSubscriptionID(subscriptionID)
}
