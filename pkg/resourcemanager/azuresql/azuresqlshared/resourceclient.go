// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlshared

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
)

// ResourceClient contains the helper functions for interacting with SQL servers / databases
type ResourceClient interface {
	CreateOrUpdateSQLServer(ctx context.Context, resourceGroupName string, location string, serverName string, properties SQLServerProperties, forceUpdate bool) (result sql.Server, err error)
	DeleteSQLServer(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error)
	GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error)
	DeleteSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error)
	CreateOrUpdateFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string, properties SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error)
	DeleteFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result autorest.Response, err error)
	DeleteDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result autorest.Response, err error)
	GetSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (result sql.FirewallRule, err error)
	GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error)
	GetFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string) (sql.FailoverGroup, error)
	CreateOrUpdateDB(ctx context.Context, resourceGroupName string, location string, serverName string, properties SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error)
	CreateOrUpdateSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string, startIP string, endIP string) (result bool, err error)
}
