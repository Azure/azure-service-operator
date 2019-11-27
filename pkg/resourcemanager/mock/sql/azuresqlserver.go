// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
)

type MockSqlServerManager struct {
	sqlServers []sql.Server
}

func findSqlServer(res []sql.Server, predicate func(sql.Server) bool) (int, sql.Server) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, sql.Server{}
}

func (manager *MockSqlServerManager) CreateOrUpdateSQLServer(ctx context.Context, resourceGroupName string, location string, serverName string, properties SQLServerProperties) (result sql.Server, err error) {
		index, _ := findSqlServer(manager.sqlServers, func(s sql.Servers) bool {
		return *s.Name == serverName
	})

	q := sql.Server{
		Response: helpers.GetRestResponse(201),
		Location: to.StringPtr(location),
		Name:     to.StringPtr(serverName),
		// todo: add resourcegroup 
	}

	if index == -1 {
		manager.sqlServers = append(manager.sqlServers, q)
	}

	return r, nil
}