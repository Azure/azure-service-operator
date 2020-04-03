// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type MySQLFirewallRuleManager interface {
	//convert(obj runtime.Object) (*v1alpha1.PostgreSQLFirewallRule, error)

	CreateFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string, startip string, endip string) (mysql.FirewallRulesCreateOrUpdateFuture, error)
	DeleteFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (string, error)
	GetFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (mysql.FirewallRule, error)
	// also embed async client methods
	resourcemanager.ARMClient
}
