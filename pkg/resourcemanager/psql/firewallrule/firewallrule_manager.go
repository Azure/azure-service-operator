// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"net/http"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type PostgreSQLFirewallRuleManager interface {
	CreateFirewallRule(ctx context.Context,
		resourcegroup string,
		servername string,
		firewallrulename string,
		startip string,
		endip string) (*http.Response, error)

	GetFirewallRule(ctx context.Context,
		resourcegroup string,
		servername string,
		firewallrulename string) (psql.FirewallRule, error)

	DeleteFirewallRule(ctx context.Context,
		resourcegroup string,
		servername string,
		firewallrulename string) (string, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
