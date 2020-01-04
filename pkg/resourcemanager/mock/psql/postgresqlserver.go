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
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
)

type MockPostgreSqlServerManager struct {
	resourceGroupName string
	PostgresqlServers []MockPostgreSqlServerResource
	SecretClient      secrets.SecretClient
	Scheme            *runtime.Scheme
	Log               logr.Logger
}

type MockPostgreSqlServerResource struct {
	resourceGroupName string
	PostgresqlServer  postgresql.Server
}

func NewMockPSQLServerClient(log logr.Logger, secretclient secrets.SecretClient, scheme *runtime.Scheme) *MockPostgreSqlServerManager {
	return &MockPostgreSqlServerManager{
		Log:          log,
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func findPostgreSqlServer(res []MockPostgreSqlServerResource, predicate func(MockPostgreSqlServerResource) bool) (int, error) {
	for index, r := range res {
		if predicate(r) {
			return index, nil
		}
	}
	return -1, errors.New("not found")
}

// CreateServerIfValid creates a new Postgresql server
func (manager *MockPostgreSqlServerManager) CreateServerIfValid(ctx context.Context, servername string, resourcegroup string, location string, tags map[string]*string, serverversion postgresql.ServerVersion, sslenforcement postgresql.SslEnforcementEnum, skuInfo postgresql.Sku, adminlogin string, adminpassword string) (postgresql.ServersCreateFuture, error) {
	index, _ := findPostgreSqlServer(manager.PostgresqlServers, func(s MockPostgreSqlServerResource) bool {
		return s.resourceGroupName == resourcegroup && *s.PostgresqlServer.Name == servername
	})

	postgresqlS := postgresql.Server{
		Response: helpers.GetRestResponse(http.StatusCreated),
		Location: to.StringPtr(location),
		Name:     to.StringPtr(servername),
	}

	q := MockPostgreSqlServerResource{
		resourceGroupName: resourcegroup,
		PostgresqlServer:  postgresqlS,
	}

	if index == -1 {
		manager.PostgresqlServers = append(manager.PostgresqlServers, q)
	}

	f := postgresql.ServersCreateFuture{}

	return f, nil
}

// DeleteServer removes the postgresqlserver
func (manager *MockPostgreSqlServerManager) DeleteServer(ctx context.Context, resourcegroup string, servername string) (string, error) {
	PostgresqlServers := manager.PostgresqlServers

	index, _ := findPostgreSqlServer(PostgresqlServers, func(s MockPostgreSqlServerResource) bool {
		return s.resourceGroupName == resourcegroup && *s.PostgresqlServer.Name == servername
	})

	if index == -1 {
		return "Not Found", nil
	}

	manager.PostgresqlServers = append(PostgresqlServers[:index], PostgresqlServers[index+1:]...)

	return "Deleted", nil
}

// GetServer gets a postgresql server
func (manager *MockPostgreSqlServerManager) GetServer(ctx context.Context, resourcegroup string, servername string) (postgresql.Server, error) {
	index, _ := findPostgreSqlServer(manager.PostgresqlServers, func(s MockPostgreSqlServerResource) bool {
		return s.resourceGroupName == resourcegroup && *s.PostgresqlServer.Name == servername
	})

	if index == -1 {
		return postgresql.Server{}, errors.New("Sql Server Not Found")
	}

	state := postgresql.ServerState("Ready")
	serverProperties := postgresql.ServerProperties{UserVisibleState: state}
	server := manager.PostgresqlServers[index].PostgresqlServer
	server.ServerProperties = &serverProperties

	return server, nil
}

func (manager *MockPostgreSqlServerManager) convert(obj runtime.Object) (*v1alpha1.PostgreSQLServer, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (manager *MockPostgreSqlServerManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := manager.convert(obj)
	if err != nil {
		return true, err
	}

	tags := map[string]*string{}
	_, _ = manager.CreateServerIfValid(
		ctx,
		instance.Name,
		instance.Spec.ResourceGroup,
		instance.Spec.Location,
		tags,
		postgresql.ServerVersion("10"),
		postgresql.SslEnforcementEnumEnabled,
		postgresql.Sku{},
		"",
		"",
	)

	instance.Status.Provisioned = true

	return true, nil
}
func (manager *MockPostgreSqlServerManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := manager.convert(obj)
	if err != nil {
		return true, err
	}

	_, _ = manager.DeleteServer(ctx, instance.Spec.ResourceGroup, instance.Name)

	return false, nil
}
func (manager *MockPostgreSqlServerManager) GetParents(runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}
