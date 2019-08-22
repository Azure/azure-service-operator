// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import "github.com/Azure/go-autorest/autorest/azure"

// ResourceClient contains the helper functions for interacting with SQL servers / databases
type ResourceClient interface {
	CreateOrUpdateSQLServer(properties SQLServerProperties) (result azure.Future, err error)
	SQLServerReady() (result bool, err error)
	CreateOrUpdateDB(properties SQLDatabaseProperties) (result azure.Future, err error)
	DeleteDB(databaseName string) (result bool, err error)
	DeleteSQLServer() (result azure.Future, err error)
}
