// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type MySqlVNetRuleManager interface {
	CreateOrUpdateMySQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string, VNetRG string, VNetName string, SubnetName string, IgnoreServiceEndpoint bool) (result mysql.VirtualNetworkRule, err error)
	DeleteMySQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error)
	GetMySQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (result mysql.VirtualNetworkRule, err error)
	resourcemanager.ARMClient
}
