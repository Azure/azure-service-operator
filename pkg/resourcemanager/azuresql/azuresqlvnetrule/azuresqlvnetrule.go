// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlvnetrule

import (
	"context"

	sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
)

type AzureSqlVNetRuleManager struct {
	creds config.Credentials
}

func NewAzureSqlVNetRuleManager(creds config.Credentials) *AzureSqlVNetRuleManager {
	return &AzureSqlVNetRuleManager{creds: creds}
}

// GetSQLVNetRule returns a VNet rule
func (m *AzureSqlVNetRuleManager) GetSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (result sql.VirtualNetworkRule, err error) {
	VNetRulesClient, err := azuresqlshared.GetGoVNetRulesClient(m.creds)
	if err != nil {
		return sql.VirtualNetworkRule{}, err
	}

	return VNetRulesClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		ruleName,
	)
}

// DeleteSQLVNetRule deletes a VNet rule
func (m *AzureSqlVNetRuleManager) DeleteSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error) {

	// check to see if the rule exists, if it doesn't then short-circuit
	_, err = m.GetSQLVNetRule(ctx, resourceGroupName, serverName, ruleName)
	if err != nil {
		return nil
	}

	VNetRulesClient, err := azuresqlshared.GetGoVNetRulesClient(m.creds)
	if err != nil {
		return err
	}

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
func (m *AzureSqlVNetRuleManager) CreateOrUpdateSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string, VNetRG string, VNetName string, SubnetName string, subscription string, IgnoreServiceEndpoint bool) (vnr sql.VirtualNetworkRule, err error) {

	VNetRulesClient, err := azuresqlshared.GetGoVNetRulesClient(m.creds)
	if err != nil {
		return sql.VirtualNetworkRule{}, err
	}

	if subscription == "" {
		subscription = m.creds.SubscriptionID()
	}
	SubnetClient, err := azuresqlshared.GetGoNetworkSubnetClient(m.creds, subscription)
	if err != nil {
		return sql.VirtualNetworkRule{}, err
	}

	// Get ARM Resource ID of Subnet based on the VNET name, Subnet name and Subnet Address Prefix
	subnet, err := SubnetClient.Get(ctx, VNetRG, VNetName, SubnetName, "")
	if err != nil {
		return vnr, err
	}
	subnetResourceID := *subnet.ID

	// Populate parameters with the right ID
	parameters := sql.VirtualNetworkRule{
		VirtualNetworkRuleProperties: &sql.VirtualNetworkRuleProperties{
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
