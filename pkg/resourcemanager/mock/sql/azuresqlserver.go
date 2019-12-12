// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"errors"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	sqlclient "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/go-autorest/autorest"
)

type MockSqlServerManager struct {
	resourceGroupName string
	sqlServers        []MockSqlServerResource
}

type MockSqlServerResource struct {
	resourceGroupName string
	sqlServer         sql.Server
}

func NewMockSqlServerManager() *MockSqlServerManager {
	return &MockSqlServerManager{}
}

func findSqlServer(res []MockSqlServerResource, predicate func(MockSqlServerResource) bool) (int, error) {
	for index, r := range res {
		if predicate(r) {
			return index, nil
		}
	}
	return -1, errors.New("not found")
}

// CreateOrUpdateSqlServer creates a new sql server
func (manager *MockSqlServerManager) CreateOrUpdateSQLServer(ctx context.Context, resourceGroupName string, location string, serverName string, properties sqlclient.SQLServerProperties) (result sql.Server, err error) {
	index, _ := findSqlServer(manager.sqlServers, func(s MockSqlServerResource) bool {
		return s.resourceGroupName == resourceGroupName && *s.sqlServer.Name == serverName
	})

	sqlS := sql.Server{
		Response: helpers.GetRestResponse(http.StatusCreated),
		Location: to.StringPtr(location),
		Name:     to.StringPtr(serverName),
	}

	q := MockSqlServerResource{
		resourceGroupName: resourceGroupName,
		sqlServer:         sqlS,
	}

	if index == -1 {
		manager.sqlServers = append(manager.sqlServers, q)
	}

	return q.sqlServer, nil
}

// DeleteSQLServer removes the sqlserver
func (manager *MockSqlServerManager) DeleteSQLServer(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error) {
	sqlServers := manager.sqlServers

	index, _ := findSqlServer(sqlServers, func(s MockSqlServerResource) bool {
		return s.resourceGroupName == resourceGroupName && *s.sqlServer.Name == serverName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("Sql Server Not Found")
	}

	manager.sqlServers = append(sqlServers[:index], sqlServers[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}

// GetServer gets a sql server
func (manager *MockSqlServerManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	index, _ := findSqlServer(manager.sqlServers, func(s MockSqlServerResource) bool {
		return s.resourceGroupName == resourceGroupName && *s.sqlServer.Name == serverName
	})

	if index == -1 {
		return sql.Server{}, errors.New("Sql Server Not Found")
	}

	state := "Ready"
	serverProperties := sql.ServerProperties{State: &state}
	server := manager.sqlServers[index].sqlServer
	server.ServerProperties = &serverProperties

	return server, nil
}
