// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresql

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
)

type MockSqlDbManager struct {
	resourceGroupName string
	sqlDbs            []MockSqlDbResource
}

type MockSqlDbResource struct {
	resourceGroupName string
	sqlServerName     string
	sqlDb             sql.Database
}

func NewMockSqlDbManager() *MockSqlDbManager {
	return &MockSqlDbManager{}
}

func findSqlDb(res []MockSqlDbResource, predicate func(MockSqlDbResource) bool) (int, MockSqlDbResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, MockSqlDbResource{}
}

//CreateorUpdateDB creates a sql Db
func (manager *MockSqlDbManager) CreateOrUpdateDB(ctx context.Context, resourceGroupName string, location string, serverName string, properties azuresqlshared.SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error) {
	index, _ := findSqlDb(manager.sqlDbs, func(s MockSqlDbResource) bool {
		return s.resourceGroupName == resourceGroupName && s.sqlServerName == serverName && *s.sqlDb.Name == properties.DatabaseName
	})

	sqlD := sql.Database{
		Location: to.StringPtr(location),
		Name:     to.StringPtr(properties.DatabaseName),
	}

	q := MockSqlDbResource{
		resourceGroupName: resourceGroupName,
		sqlServerName:     serverName,
		sqlDb:             sqlD,
	}

	if index == -1 {
		manager.sqlDbs = append(manager.sqlDbs, q)
	}

	return sql.DatabasesCreateOrUpdateFuture{}, nil

}

func (manager *MockSqlDbManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.Database, err error) {
	index, _ := findSqlDb(manager.sqlDbs, func(s MockSqlDbResource) bool {
		return s.resourceGroupName == resourceGroupName && s.sqlServerName == serverName && *s.sqlDb.Name == databaseName
	})

	if index == -1 {
		return sql.Database{}, errors.New("Sql Db Not Found")
	}

	return manager.sqlDbs[index].sqlDb, nil
}

//GetServer gets a sql server
func (manager *MockSqlDbManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	sqlManager := MockSqlServerManager{}
	return sqlManager.GetServer(ctx, resourceGroupName, serverName)
}

// DeleteDb removes the sql db
func (manager *MockSqlDbManager) DeleteDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result autorest.Response, err error) {
	sqlDbs := manager.sqlDbs

	index, _ := findSqlDb(manager.sqlDbs, func(s MockSqlDbResource) bool {
		return s.resourceGroupName == resourceGroupName && s.sqlServerName == serverName && *s.sqlDb.Name == databaseName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("Sql Db Not Found")
	}

	manager.sqlDbs = append(sqlDbs[:index], sqlDbs[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}

func (db *MockSqlDbManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := db.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	dbName := instance.ObjectMeta.Name
	dbEdition := instance.Spec.Edition

	azureSqlDatabaseProperties := azuresqlshared.SQLDatabaseProperties{
		DatabaseName: dbName,
		Edition:      dbEdition,
	}

	_, err = db.CreateOrUpdateDB(ctx, groupName, location, server, azureSqlDatabaseProperties)
	if err != nil {
		return false, err
	}

	instance.Status.Provisioning = true
	instance.Status.Provisioned = true

	return true, nil
}

func (db *MockSqlDbManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := db.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	dbName := instance.ObjectMeta.Name

	_, err = db.DeleteDB(ctx, groupName, server, dbName)
	if err != nil {
		return false, err
	}

	return false, nil
}

func (g *MockSqlDbManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return nil, nil
}

func (*MockSqlDbManager) convert(obj runtime.Object) (*azurev1alpha1.AzureSqlDatabase, error) {
	local, ok := obj.(*azurev1alpha1.AzureSqlDatabase)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
