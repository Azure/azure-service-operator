// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlfirewallrule

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	"github.com/Azure/go-autorest/autorest/to"
)

type AzureSqlFirewallRuleManager struct {
	creds config.Credentials
}

func NewAzureSqlFirewallRuleManager(creds config.Credentials) *AzureSqlFirewallRuleManager {
	return &AzureSqlFirewallRuleManager{creds: creds}
}

// GetServer returns a SQL server
func (m *AzureSqlFirewallRuleManager) GetServer(ctx context.Context, subscriptionID string, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient, err := azuresqlshared.GetGoServersClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
	if err != nil {
		return sql.Server{}, err
	}

	return serversClient.Get(
		ctx,
		resourceGroupName,
		serverName,
	)
}

// GetSQLFirewallRule returns a firewall rule
func (m *AzureSqlFirewallRuleManager) GetSQLFirewallRule(ctx context.Context, subscriptionID string, resourceGroupName string, serverName string, ruleName string) (result sql.FirewallRule, err error) {
	firewallClient, err := azuresqlshared.GetGoFirewallClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
	if err != nil {
		return sql.FirewallRule{}, err
	}

	return firewallClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		ruleName,
	)
}

// DeleteSQLFirewallRule deletes a firewall rule
func (m *AzureSqlFirewallRuleManager) DeleteSQLFirewallRule(ctx context.Context, subscriptionID string, resourceGroupName string, serverName string, ruleName string) (err error) {

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := m.GetServer(ctx, subscriptionID, resourceGroupName, serverName)
	if err != nil || *server.State != "Ready" {
		return nil
	}

	// check to see if the rule exists, if it doesn't then short-circuit
	_, err = m.GetSQLFirewallRule(ctx, subscriptionID, resourceGroupName, serverName, ruleName)
	if err != nil {
		return nil
	}

	firewallClient, err := azuresqlshared.GetGoFirewallClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
	if err != nil {
		return err
	}

	_, err = firewallClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
		ruleName,
	)

	return err
}

// CreateOrUpdateSQLFirewallRule creates or updates a firewall rule
// based on code from: https://github.com/Azure-Samples/azure-m-for-go-samples/blob/master/sql/sql.go#L111
// to allow allow Azure services to connect example: https://docs.microsoft.com/en-us/azure/sql-database/sql-database-firewall-configure#manage-firewall-rules-using-azure-cli
func (m *AzureSqlFirewallRuleManager) CreateOrUpdateSQLFirewallRule(
	ctx context.Context,
	subscriptionID string,
	resourceGroupName string,
	serverName string,
	ruleName string,
	startIP string,
	endIP string,
) (result bool, err error) {

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := m.GetServer(ctx, subscriptionID, resourceGroupName, serverName)
	if err != nil || *server.State != "Ready" {
		return false, err
	}

	firewallClient, err := azuresqlshared.GetGoFirewallClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
	if err != nil {
		return false, err
	}

	_, err = firewallClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
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
