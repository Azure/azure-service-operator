// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type MockSqlFirewallRuleManager struct {
	sqlFirewallRules []MockSqlFirewallRuleResource
}

type MockSqlFirewallRuleResource struct {
	resourceGroupName string
	sqlServerName     string
	sqlFirewallRule   sql.FirewallRule
}

func NewMockSqlFirewallRuleManager() *MockSqlFirewallRuleManager {
	return &MockSqlFirewallRuleManager{}
}

func findSqlFirewallRule(res []MockSqlFirewallRuleResource, predicate func(MockSqlFirewallRuleResource) bool) (int, MockSqlFirewallRuleResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, MockSqlFirewallRuleResource{}
}

// CreateOrUpdateSQLFirewallRule creates a sql firewall rule
func (manager *MockSqlFirewallRuleManager) CreateOrUpdateSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string, startIP string, endIP string) (result bool, err error) {
	index, _ := findSqlFirewallRule(manager.sqlFirewallRules, func(s MockSqlFirewallRuleResource) bool {
		return s.resourceGroupName == resourceGroupName && s.sqlServerName == serverName && *s.sqlFirewallRule.Name == ruleName
	})

	sqlFR := sql.FirewallRule{
		Name: to.StringPtr(ruleName),
	}

	q := MockSqlFirewallRuleResource{
		resourceGroupName: resourceGroupName,
		sqlServerName:     serverName,
		sqlFirewallRule:   sqlFR,
	}

	if index == -1 {
		manager.sqlFirewallRules = append(manager.sqlFirewallRules, q)
	}

	return true, nil
}

// GetSQLFirewallRule gets a sql firewall rule
func (manager *MockSqlFirewallRuleManager) GetSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (result sql.FirewallRule, err error) {
	index, _ := findSqlFirewallRule(manager.sqlFirewallRules, func(s MockSqlFirewallRuleResource) bool {
		return s.resourceGroupName == resourceGroupName && s.sqlServerName == serverName && *s.sqlFirewallRule.Name == ruleName
	})

	if index == -1 {
		return sql.FirewallRule{}, errors.New("Sql Firewall Rule Not Found")
	}

	return manager.sqlFirewallRules[index].sqlFirewallRule, nil
}

//GetServer gets a sql server
func (manager *MockSqlFirewallRuleManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	sqlManager := MockSqlServerManager{}
	return sqlManager.GetServer(ctx, resourceGroupName, serverName)
}

// DeleteSQLFirewallRule deletes a sql firewall rule
func (manager *MockSqlFirewallRuleManager) DeleteSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error) {

	sqlFirewallRules := manager.sqlFirewallRules

	index, _ := findSqlFirewallRule(manager.sqlFirewallRules, func(s MockSqlFirewallRuleResource) bool {
		return s.resourceGroupName == resourceGroupName && s.sqlServerName == serverName && *s.sqlFirewallRule.Name == ruleName
	})

	if index == -1 {
		return errors.New("Sql Firewall Rule Found")
	}

	manager.sqlFirewallRules = append(sqlFirewallRules[:index], sqlFirewallRules[index+1:]...)

	return nil
}

func (fw *MockSqlFirewallRuleManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	return true, nil
}

func (fw *MockSqlFirewallRuleManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	return true, nil
}

func (g *MockSqlFirewallRuleManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}

func (*MockSqlFirewallRuleManager) convert(obj runtime.Object) (*azurev1alpha1.AzureSqlFirewallRule, error) {
	local, ok := obj.(*azurev1alpha1.AzureSqlFirewallRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
