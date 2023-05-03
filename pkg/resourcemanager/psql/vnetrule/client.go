// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

type PostgreSQLVNetRuleClient struct {
	creds config.Credentials
}

func NewPostgreSQLVNetRuleClient(creds config.Credentials) *PostgreSQLVNetRuleClient {
	return &PostgreSQLVNetRuleClient{creds: creds}
}

func GetPostgreSQLVNetRulesClient(creds config.Credentials) psql.VirtualNetworkRulesClient {
	VNetRulesClient := psql.NewVirtualNetworkRulesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	VNetRulesClient.Authorizer = a
	VNetRulesClient.AddToUserAgent(config.UserAgent())
	return VNetRulesClient
}

// retrieves the Subnetclient
func GetGoNetworkSubnetClient(creds config.Credentials, subscription string) network.SubnetsClient {
	SubnetsClient := network.NewSubnetsClientWithBaseURI(config.BaseURI(), subscription)
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	SubnetsClient.Authorizer = a
	SubnetsClient.AddToUserAgent(config.UserAgent())
	return SubnetsClient
}

// GetPostgreSQLVNetRule returns a VNet rule
func (c *PostgreSQLVNetRuleClient) GetPostgreSQLVNetRule(
	ctx context.Context,
	resourceGroupName string,
	serverName string,
	ruleName string) (result psql.VirtualNetworkRule, err error) {

	VNetRulesClient := GetPostgreSQLVNetRulesClient(c.creds)

	return VNetRulesClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		ruleName,
	)
}

// deletes a VNet rule
func (c *PostgreSQLVNetRuleClient) DeletePostgreSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error) {

	// check to see if the rule exists, if it doesn't then short-circuit
	_, err = c.GetPostgreSQLVNetRule(ctx, resourceGroupName, serverName, ruleName)
	if err != nil {
		return nil
	}

	VNetRulesClient := GetPostgreSQLVNetRulesClient(c.creds)
	_, err = VNetRulesClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
		ruleName,
	)

	return err
}

// creates or updates a VNet rule
func (c *PostgreSQLVNetRuleClient) CreateOrUpdatePostgreSQLVNetRule(
	ctx context.Context,
	resourceGroupName string,
	serverName string,
	ruleName string,
	VNetRG string,
	VNetName string,
	SubnetName string,
	subscription string,
	IgnoreServiceEndpoint bool) (vnr psql.VirtualNetworkRule, err error) {

	VNetRulesClient := GetPostgreSQLVNetRulesClient(c.creds)
	// Subnet may be in another subscription
	if subscription == "" {
		subscription = c.creds.SubscriptionID()
	}
	SubnetClient := GetGoNetworkSubnetClient(c.creds, subscription)

	// Get ARM Resource ID of Subnet based on the VNET name, Subnet name and Subnet Address Prefix
	subnet, err := SubnetClient.Get(ctx, VNetRG, VNetName, SubnetName, "")
	if err != nil {
		return vnr, err
	}
	subnetResourceID := *subnet.ID

	// Populate parameters with the right ID
	parameters := psql.VirtualNetworkRule{
		VirtualNetworkRuleProperties: &psql.VirtualNetworkRuleProperties{
			VirtualNetworkSubnetID:           &subnetResourceID,
			IgnoreMissingVnetServiceEndpoint: &IgnoreServiceEndpoint,
		},
	}

	// Call CreateOrUpdate
	result, err := VNetRulesClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		ruleName,
		parameters,
	)
	return result.Result(VNetRulesClient)
}
