// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"net/http"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest/to"
)

type PSQLFirewallRuleClient struct {
}

func NewPSQLFirewallRuleClient() *PSQLFirewallRuleClient {
	return &PSQLFirewallRuleClient{}
}

func getPSQLFirewallRulesClient() (psql.FirewallRulesClient, error) {
	firewallRulesClient := psql.NewFirewallRulesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return psql.FirewallRulesClient{}, err
	}
	firewallRulesClient.Authorizer = a
	firewallRulesClient.AddToUserAgent(config.UserAgent())
	return firewallRulesClient, err
}

func (p *PSQLFirewallRuleClient) CreateFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string, startip string, endip string) (*http.Response, error) {

	client, err := getPSQLFirewallRulesClient()
	if err != nil {
		return &http.Response{
			StatusCode: 500,
		}, err
	}

	firewallRuleProperties := psql.FirewallRuleProperties{
		StartIPAddress: to.StringPtr(startip),
		EndIPAddress:   to.StringPtr(endip),
	}

	future, err := client.CreateOrUpdate(
		ctx,
		resourcegroup,
		servername,
		firewallrulename,
		psql.FirewallRule{
			FirewallRuleProperties: &firewallRuleProperties,
		},
	)
	if err != nil {
		return &http.Response{
			StatusCode: 500,
		}, err
	}

	return future.GetResult(client)
}

func (p *PSQLFirewallRuleClient) DeleteFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (status string, err error) {

	client, err := getPSQLFirewallRulesClient()
	if err != nil {
		return "", err
	}

	_, err = client.Get(ctx, resourcegroup, servername, firewallrulename)
	if err == nil { // FW rule present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername, firewallrulename)
		return future.Status(), err
	}

	// FW rule not present so return success anyway
	return "Firewall Rule not present", nil
}

func (p *PSQLFirewallRuleClient) GetFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (firewall psql.FirewallRule, err error) {

	client, err := getPSQLFirewallRulesClient()
	if err != nil {
		return psql.FirewallRule{}, err
	}

	return client.Get(ctx, resourcegroup, servername, firewallrulename)
}
