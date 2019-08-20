// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
)

// SQLServerProperties contains values needed for adding / updating SQL servers,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#Server
// also wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#ServerProperties
type SQLServerProperties struct {
}

// SQLDatabaseProperties contains values needed for adding / updating SQL servers,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#Database
// also wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#DatabaseProperties
type SQLDatabaseProperties struct {
}

// SQLServerPropertiesToServer translates APIProperties to APICreateOrUpdateParameter and underlying APICreateOrUpdateProperties
func SQLServerPropertiesToServer(parameter SQLServerProperties) (result sql.Server) {

	return result
}

// SQLDatabasePropertiesToDatabase translates APIProperties to APICreateOrUpdateParameter and underlying APICreateOrUpdateProperties
func SQLDatabasePropertiesToDatabase(parameter SQLDatabaseProperties) (result sql.Database) {

	return result
}
