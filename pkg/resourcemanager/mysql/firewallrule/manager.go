// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type PostgreSQLFirewallRuleManager interface {
	//convert(obj runtime.Object) (*v1alpha1.PostgreSQLFirewallRule, error)

	CreateFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string, startip string, endip string) (psql.FirewallRulesCreateOrUpdateFuture, error)
	DeleteFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (string, error)
	GetFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (psql.FirewallRule, error)
	// also embed async client methods
	resourcemanager.ARMClient
}
