// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type MySQLVNetRuleClient struct {
	creds config.Credentials
}

func NewMySQLVNetRuleClient(creds config.Credentials) *MySQLVNetRuleClient {
	return &MySQLVNetRuleClient{creds: creds}
}

// NewARMClient returns a new manager (but as an ARMClient).
func NewARMClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) resourcemanager.ARMClient {
	return NewMySQLVNetRuleClient(creds)
}

func getMySQLVNetRulesClient(creds config.Credentials) mysql.VirtualNetworkRulesClient {
	VNetRulesClient := mysql.NewVirtualNetworkRulesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	VNetRulesClient.Authorizer = a
	VNetRulesClient.AddToUserAgent(config.UserAgent())
	return VNetRulesClient
}

// GetNetworkSubnetClient retrieves a Subnetclient
func GetGoNetworkSubnetClient(creds config.Credentials, subscription string) network.SubnetsClient {
	SubnetsClient := network.NewSubnetsClientWithBaseURI(config.BaseURI(), subscription)
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	SubnetsClient.Authorizer = a
	SubnetsClient.AddToUserAgent(config.UserAgent())
	return SubnetsClient
}

// GetSQLVNetRule returns a VNet rule
func (c *MySQLVNetRuleClient) GetSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (result mysql.VirtualNetworkRule, err error) {
	VNetRulesClient := getMySQLVNetRulesClient(c.creds)

	return VNetRulesClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		ruleName,
	)
}

// DeleteSQLVNetRule deletes a VNet rule
func (c *MySQLVNetRuleClient) DeleteSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error) {

	// check to see if the rule exists, if it doesn't then short-circuit
	_, err = c.GetSQLVNetRule(ctx, resourceGroupName, serverName, ruleName)
	if err != nil {
		return nil
	}

	VNetRulesClient := getMySQLVNetRulesClient(c.creds)
	_, err = VNetRulesClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
		ruleName,
	)

	return err
}

// CreateOrUpdateSQLVNetRule creates or updates a VNet rule
// based on code from: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql#VirtualNetworkRulesClient.CreateOrUpdate
func (c *MySQLVNetRuleClient) CreateOrUpdateSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string, VNetRG string, VNetName string, SubnetName string, subscription string, IgnoreServiceEndpoint bool) (vnr mysql.VirtualNetworkRule, err error) {

	VNetRulesClient := getMySQLVNetRulesClient(c.creds)
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
	parameters := mysql.VirtualNetworkRule{
		VirtualNetworkRuleProperties: &mysql.VirtualNetworkRuleProperties{
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
