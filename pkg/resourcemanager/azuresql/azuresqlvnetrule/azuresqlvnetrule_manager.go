// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlvnetrule

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type SqlVNetRuleManager interface {
	CreateOrUpdateSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string, VNetRG string, VNetName string, SubnetName string, IgnoreServiceEndpoint bool) (result sql.VirtualNetworkRule, err error)
	DeleteSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error)
	GetSQLVNetRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (result sql.VirtualNetworkRule, err error)
	resourcemanager.ARMClient
}
