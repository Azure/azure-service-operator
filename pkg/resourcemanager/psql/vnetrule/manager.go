// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type PostgreSqlVNetRuleManager interface {
	CreateOrUpdatePostgreSQLVNetRule(
		ctx context.Context,
		resourceGroupName string,
		serverName string,
		ruleName string,
		VNetRG string,
		VNetName string,
		SubnetName string,
		IgnoreServiceEndpoint bool) (result psql.VirtualNetworkRule, err error)
	DeletePostgreSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error)
	GetPostgreSQLVNetRulesClient(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (result psql.VirtualNetworkRule, err error)
	resourcemanager.ARMClient
}
