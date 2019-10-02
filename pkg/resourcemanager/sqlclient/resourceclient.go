// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
)

// ResourceClient contains the helper functions for interacting with SQL servers / databases
type ResourceClient interface {
	CreateOrUpdateSQLServer(properties SQLServerProperties) (result sql.Server, err error)
	CreateOrUpdateSQLFirewallRule(ruleName string, startIP string, endIP string) (result bool, err error)
	CreateOrUpdateDB(properties SQLDatabaseProperties) (result sql.Database, err error)
	GetServer() (result sql.Server, err error)
	GetDB(databaseName string) (sql.Database, error)
	GetSQLFirewallRule(ruleName string) (result sql.FirewallRule, err error)
	DeleteDB(databaseName string) (result autorest.Response, err error)
	DeleteSQLServer() (result autorest.Response, err error)
	DeleteSQLFirewallRule(ruleName string) (err error)
	IsAsyncNotCompleted(err error) (result bool)
	CheckNameAvailability() (result AvailabilityResponse, err error)
}
