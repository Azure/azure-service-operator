// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest/to"
)

type MySQLFirewallRuleClient struct {
}

func NewMySQLFirewallRuleClient() *MySQLFirewallRuleClient {
	return &MySQLFirewallRuleClient{}
}

func getMySQLFirewallRulesClient() mysql.FirewallRulesClient {
	firewallRulesClient := mysql.NewFirewallRulesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	firewallRulesClient.Authorizer = a
	firewallRulesClient.AddToUserAgent(config.UserAgent())
	return firewallRulesClient
}

func (m *MySQLFirewallRuleClient) CreateFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string, startip string, endip string) (future mysql.FirewallRulesCreateOrUpdateFuture, err error) {

	client := getMySQLFirewallRulesClient()

	firewallRuleProperties := mysql.FirewallRuleProperties{
		StartIPAddress: to.StringPtr(startip),
		EndIPAddress:   to.StringPtr(endip),
	}

	future, err = client.CreateOrUpdate(
		ctx,
		resourcegroup,
		servername,
		firewallrulename,
		mysql.FirewallRule{
			FirewallRuleProperties: &firewallRuleProperties,
		},
	)
	return future, err
}

func (m *MySQLFirewallRuleClient) DeleteFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (status string, err error) {

	client := getMySQLFirewallRulesClient()

	_, err = client.Get(ctx, resourcegroup, servername, firewallrulename)
	if err == nil { // FW rule present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername, firewallrulename)
		return future.Status(), err
	}
	// FW rule not present so return success anyway
	return "Firewall Rule not present", nil

}

func (m *MySQLFirewallRuleClient) GetFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (firewall mysql.FirewallRule, err error) {

	client := getMySQLFirewallRulesClient()

	return client.Get(ctx, resourcegroup, servername, firewallrulename)
}
