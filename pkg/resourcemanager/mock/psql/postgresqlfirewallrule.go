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

type MockPostgreSqlFirewallRuleManager struct {
	resourceGroupName       string
	PostgresqlFirewallRules []MockPostgreSqlFirewallRuleResource
	Log                     logr.Logger
}

type MockPostgreSqlFirewallRuleResource struct {
	resourceGroupName      string
	PostgresqlServerName   string
	PostgresqlFirewallRule postgresql.FirewallRule
}

func NewMockPostgreSqlFirewallRuleManager(log logr.Logger) *MockPostgreSqlFirewallRuleManager {
	return &MockPostgreSqlFirewallRuleManager{
		Log: log,
	}
}

func findPostgreSqlFirewallRule(res []MockPostgreSqlFirewallRuleResource, predicate func(MockPostgreSqlFirewallRuleResource) bool) (int, error) {
	for index, r := range res {
		if predicate(r) {
			return index, nil
		}
	}
	return -1, errors.New("not found")
}

//CreateorUpdateFirewallRule creates a Postgresql FirewallRule
func (manager *MockPostgreSqlFirewallRuleManager) CreateFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string, startip string, endip string) (postgresql.FirewallRulesCreateOrUpdateFuture, error) {
	index, _ := findPostgreSqlFirewallRule(manager.PostgresqlFirewallRules, func(s MockPostgreSqlFirewallRuleResource) bool {
		return s.resourceGroupName == resourcegroup && s.PostgresqlServerName == servername && *s.PostgresqlFirewallRule.Name == firewallrulename
	})

	PostgresqlF := postgresql.FirewallRule{
		Response: helpers.GetRestResponse(http.StatusCreated),
		Name:     to.StringPtr(firewallrulename),
		FirewallRuleProperties: &postgresql.FirewallRuleProperties{
			StartIPAddress: to.StringPtr(startip),
			EndIPAddress:   to.StringPtr(endip),
		},
	}

	q := MockPostgreSqlFirewallRuleResource{
		resourceGroupName:      resourcegroup,
		PostgresqlServerName:   servername,
		PostgresqlFirewallRule: PostgresqlF,
	}

	if index == -1 {
		manager.PostgresqlFirewallRules = append(manager.PostgresqlFirewallRules, q)
	}

	return postgresql.FirewallRulesCreateOrUpdateFuture{}, nil

}

func (manager *MockPostgreSqlFirewallRuleManager) GetFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (postgresql.FirewallRule, error) {
	index, _ := findPostgreSqlFirewallRule(manager.PostgresqlFirewallRules, func(s MockPostgreSqlFirewallRuleResource) bool {
		return s.resourceGroupName == resourcegroup && s.PostgresqlServerName == servername && *s.PostgresqlFirewallRule.Name == firewallrulename
	})

	if index == -1 {
		return postgresql.FirewallRule{}, errors.New("Sql Fw rule Not Found")
	}

	return manager.PostgresqlFirewallRules[index].PostgresqlFirewallRule, nil
}

// DeleteFirewallRule removes the Postgresql FirewallRule
func (manager *MockPostgreSqlFirewallRuleManager) DeleteFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (string, error) {
	PostgresqlFirewallRules := manager.PostgresqlFirewallRules

	index, _ := findPostgreSqlFirewallRule(manager.PostgresqlFirewallRules, func(s MockPostgreSqlFirewallRuleResource) bool {
		return s.resourceGroupName == resourcegroup && s.PostgresqlServerName == servername && *s.PostgresqlFirewallRule.Name == firewallrulename
	})

	if index == -1 {
		return "Not found", nil
	}

	manager.PostgresqlFirewallRules = append(PostgresqlFirewallRules[:index], PostgresqlFirewallRules[index+1:]...)

	return "Deleted", nil
}

func (manager *MockPostgreSqlFirewallRuleManager) convert(obj runtime.Object) (*v1alpha1.PostgreSQLFirewallRule, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLFirewallRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (manager *MockPostgreSqlFirewallRuleManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.EnsureOption) (bool, error) {
	instance, err := manager.convert(obj)
	if err != nil {
		return true, err
	}

	_, _ = manager.CreateFirewallRule(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Name, "0.0.0.0", "0.0.0.0")

	instance.Status.Provisioned = true

	return true, nil
}

func (manager *MockPostgreSqlFirewallRuleManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := manager.convert(obj)
	if err != nil {
		return true, err
	}

	_, _ = manager.DeleteFirewallRule(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Name)

	return false, nil
}

func (manager *MockPostgreSqlFirewallRuleManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}
