// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package psql

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	postgresql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	resourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
)

type MockPostgreSqlDbManager struct {
	resourceGroupName string
	PostgresqlDbs     []MockPostgreSqlDbResource
	Log               logr.Logger
}

type MockPostgreSqlDbResource struct {
	resourceGroupName    string
	PostgresqlServerName string
	PostgresqlDb         postgresql.Database
}

func NewMockPostgreSqlDbManager(log logr.Logger) *MockPostgreSqlDbManager {
	return &MockPostgreSqlDbManager{
		Log: log,
	}
}

func findPostgreSqlDb(res []MockPostgreSqlDbResource, predicate func(MockPostgreSqlDbResource) bool) (int, error) {
	for index, r := range res {
		if predicate(r) {
			return index, nil
		}
	}
	return -1, errors.New("not found")
}

//CreateorUpdateDB creates a Postgresql Db
func (manager *MockPostgreSqlDbManager) CreateDatabaseIfValid(ctx context.Context, databasename string, servername string, resourcegroup string) (postgresql.DatabasesCreateOrUpdateFuture, error) {
	index, _ := findPostgreSqlDb(manager.PostgresqlDbs, func(s MockPostgreSqlDbResource) bool {
		return s.resourceGroupName == resourcegroup && s.PostgresqlServerName == servername && *s.PostgresqlDb.Name == databasename
	})

	PostgresqlD := postgresql.Database{
		Response: helpers.GetRestResponse(http.StatusCreated),
		Name:     to.StringPtr(databasename),
	}

	q := MockPostgreSqlDbResource{
		resourceGroupName:    resourcegroup,
		PostgresqlServerName: servername,
		PostgresqlDb:         PostgresqlD,
	}

	if index == -1 {
		manager.PostgresqlDbs = append(manager.PostgresqlDbs, q)
	}

	return postgresql.DatabasesCreateOrUpdateFuture{}, nil

}

func (manager *MockPostgreSqlDbManager) GetDatabase(ctx context.Context, resourcegroup string, servername string, database string) (postgresql.Database, error) {
	index, _ := findPostgreSqlDb(manager.PostgresqlDbs, func(s MockPostgreSqlDbResource) bool {
		return s.resourceGroupName == resourcegroup && s.PostgresqlServerName == servername && *s.PostgresqlDb.Name == database
	})

	if index == -1 {
		return postgresql.Database{}, errors.New("Sql Db Not Found")
	}

	return manager.PostgresqlDbs[index].PostgresqlDb, nil
}

// DeleteDb removes the Postgresql db
func (manager *MockPostgreSqlDbManager) DeleteDatabase(ctx context.Context, databasename string, servername string, resourcegroup string) (string, error) {
	PostgresqlDbs := manager.PostgresqlDbs

	index, _ := findPostgreSqlDb(manager.PostgresqlDbs, func(s MockPostgreSqlDbResource) bool {
		return s.resourceGroupName == resourcegroup && s.PostgresqlServerName == servername && *s.PostgresqlDb.Name == databasename
	})

	if index == -1 {
		return "Not found", nil
	}

	manager.PostgresqlDbs = append(PostgresqlDbs[:index], PostgresqlDbs[index+1:]...)

	return "Deleted", nil
}

func (manager *MockPostgreSqlDbManager) convert(obj runtime.Object) (*v1alpha1.PostgreSQLDatabase, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLDatabase)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (manager *MockPostgreSqlDbManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.EnsureOption) (bool, error) {
	instance, err := manager.convert(obj)
	if err != nil {
		return true, err
	}

	_, _ = manager.CreateDatabaseIfValid(ctx, instance.Name, instance.Spec.Server, instance.Spec.ResourceGroup)

	instance.Status.Provisioned = true

	return true, nil
}

func (manager *MockPostgreSqlDbManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := manager.convert(obj)
	if err != nil {
		return true, err
	}

	_, _ = manager.DeleteDatabase(ctx, instance.Name, instance.Spec.Server, instance.Spec.ResourceGroup)

	return false, nil
}

func (manager *MockPostgreSqlDbManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}
