// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

// CreateOrUpdateSQLServer creates an instance of a SQL server
func CreateOrUpdateSQLServer(provider ResourceClient, properties SQLServerProperties) (result *string, err error) {
	return provider.CreateOrUpdateSQLServerImpl(SQLServerPropertiesToServer(properties))
}

// CreateOrUpdateDB creates an API endpoint on an API Management Service. Returns "true" if successful.
func CreateOrUpdateDB(provider ResourceClient, dbName string, properties SQLDatabaseProperties) (result *string, err error) {
	return provider.CreateOrUpdateDBImpl(dbName, SQLDatabasePropertiesToDatabase(properties))
}

// DeleteDB deletes a database
func DeleteDB(provider ResourceClient, dbName string) (result bool, err error) {
	return provider.DeleteDBImpl(dbName)
}

// DeleteSQLServer deletes a SQL server
func DeleteSQLServer(provider ResourceClient) (result bool, err error) {
	return provider.DeleteSQLServerImpl()
}
