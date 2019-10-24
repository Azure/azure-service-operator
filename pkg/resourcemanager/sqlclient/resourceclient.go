// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
)

// ResourceClient contains the helper functions for interacting with SQL servers / databases
type ResourceClient interface {
	CreateOrUpdateSQLServer(ctx context.Context, properties SQLServerProperties) (result sql.Server, err error)
	CreateOrUpdateSQLFirewallRule(ctx context.Context, ruleName string, startIP string, endIP string) (result bool, err error)
	CreateOrUpdateDB(ctx context.Context, properties SQLDatabaseProperties) (result sql.Database, err error)
	CreateOrUpdateFailoverGroup(ctx context.Context, failovergroupname string, properties SQLFailoverGroupProperties) (result sql.FailoverGroup, err error)
	GetServer(ctx context.Context) (result sql.Server, err error)
	GetSQLFirewallRule(ctx context.Context, ruleName string) (result sql.FirewallRule, err error)
	GetDB(ctx context.Context, databaseName string) (sql.Database, error)
	GetFailoverGroup(ctx context.Context, failovergroupname string) (sql.FailoverGroup, error)
	DeleteDB(ctx context.Context, databaseName string) (result autorest.Response, err error)
	DeleteSQLFirewallRule(ctx context.Context, ruleName string) (err error)
	DeleteSQLServer(ctx context.Context) (result autorest.Response, err error)
	DeleteFailoverGroup(ctx context.Context, failoverGroupName string)
	CheckNameAvailability(ctx context.Context) (result AvailabilityResponse, err error)
}
