// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

// ResourceClient contains the helper functions for interacting with SQL servers / databases
type ResourceClient interface {
	CreateOrUpdateSQLServer(properties SQLServerProperties) (result bool, err error)
	SQLServerReady() (result bool, err error)
	CreateOrUpdateDB(properties SQLDatabaseProperties) (result bool, err error)
	DeleteDB(databaseName string) (result bool, err error)
	DeleteSQLServer() (result bool, err error)
}
