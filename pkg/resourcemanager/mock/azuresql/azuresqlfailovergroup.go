// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresql

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type MockSqlFailoverGroupManager struct {
	sqlFailoverGroups []MockSqlFailoverResource
}

type MockSqlFailoverResource struct {
	resourceGroupName string
	sqlServerName     string
	sqlFailover       sql.FailoverGroup
}

func NewMockSqlFailoverGroupManager() *MockSqlFailoverGroupManager {
	return &MockSqlFailoverGroupManager{}
}

func findSqlFailoverGroup(res []MockSqlFailoverResource, predicate func(MockSqlFailoverResource) bool) (int, MockSqlFailoverResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, MockSqlFailoverResource{}
}

func (manager *MockSqlFailoverGroupManager) CreateOrUpdateFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string, properties azuresqlshared.SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error) {
	index, _ := findSqlFailoverGroup(manager.sqlFailoverGroups, func(s MockSqlFailoverResource) bool {
		return s.resourceGroupName == resourceGroupName && s.sqlServerName == serverName && *s.sqlFailover.Name == failovergroupname
	})

	sqlFG := sql.FailoverGroup{
		Name: to.StringPtr(failovergroupname),
	}

	q := MockSqlFailoverResource{
		resourceGroupName: resourceGroupName,
		sqlServerName:     serverName,
		sqlFailover:       sqlFG,
	}

	if index == -1 {
		manager.sqlFailoverGroups = append(manager.sqlFailoverGroups, q)
	}

	return sql.FailoverGroupsCreateOrUpdateFuture{}, nil
}

//DeleteFailoverGroup deletes a failover group
func (manager *MockSqlFailoverGroupManager) DeleteFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result autorest.Response, err error) {
	sqlFailoverGroups := manager.sqlFailoverGroups

	index, _ := findSqlFailoverGroup(manager.sqlFailoverGroups, func(s MockSqlFailoverResource) bool {
		return s.resourceGroupName == resourceGroupName && s.sqlServerName == serverName && *s.sqlFailover.Name == failoverGroupName
	})
	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("Sql Firewall Rule Found")
	}

	manager.sqlFailoverGroups = append(sqlFailoverGroups[:index], sqlFailoverGroups[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}

//GetFailoverGroup gets a failover group
func (manager *MockSqlFailoverGroupManager) GetFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string) (result sql.FailoverGroup, err error) {

	index, _ := findSqlFailoverGroup(manager.sqlFailoverGroups, func(s MockSqlFailoverResource) bool {
		return s.resourceGroupName == resourceGroupName && s.sqlServerName == serverName && *s.sqlFailover.Name == failovergroupname
	})
	if index == -1 {
		return sql.FailoverGroup{}, errors.New("Sql FailoverGroup Not Found")
	}

	return manager.sqlFailoverGroups[index].sqlFailover, nil
}

//GetServer gets a sql server
func (manager *MockSqlFailoverGroupManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	sqlManager := MockSqlServerManager{}
	return sqlManager.GetServer(ctx, resourceGroupName, serverName)
}

//GetDb gets a sql db server
func (manager *MockSqlFailoverGroupManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.Database, err error) {
	sqlManager := MockSqlDbManager{}
	return sqlManager.GetDB(ctx, resourceGroupName, serverName, databaseName)
}

func (fg *MockSqlFailoverGroupManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	return true, nil
}

func (fg *MockSqlFailoverGroupManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	return true, nil
}

func (g *MockSqlFailoverGroupManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}

func (*MockSqlFailoverGroupManager) convert(obj runtime.Object) (*azurev1alpha1.AzureSqlFailoverGroup, error) {
	local, ok := obj.(*azurev1alpha1.AzureSqlFailoverGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
