// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"

// ResourceClient contains the helper functions for interacting with SQL servers / databases
type ResourceClient interface {
	CreateOrUpdateSQLServerImpl(properties sql.ServerProperties) (result *string, err error)
	CreateOrUpdateDBImpl(dbName string, properties sql.DatabaseProperties) (result *string, err error)
	DeleteDBImpl(dbName string) (result bool, err error)
	DeleteSQLServerImpl() (result bool, err error)
}
