// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || mysql || mysqlfirewallrule
// +build all mysql mysqlfirewallrule

package controllers

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
)

func TestMySQLFirewallRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	server := GenerateTestResourceNameWithRandom("mysqlsrv", 10)
	ruleName := GenerateTestResourceNameWithRandom("mysql-fw", 10)

	// Create the mySQLDB object and expect the Reconcile to be created
	ruleInstance := &azurev1alpha1.MySQLFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ruleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLFirewallRuleSpec{
			Server:         server,
			ResourceGroup:  rgName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, ruleInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, ruleInstance)
}

func TestMySQLFirewallRuleControllerNoServer(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	server := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	rgName := tc.resourceGroupName
	ruleName := GenerateTestResourceNameWithRandom("mysql-fw", 10)

	// Create the mySQLDB object and expect the Reconcile to be created
	ruleInstance := &azurev1alpha1.MySQLFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ruleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLFirewallRuleSpec{
			Server:         server,
			ResourceGroup:  rgName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}
	EnsureInstanceWithResult(ctx, t, tc, ruleInstance, errhelp.ResourceNotFound, false)
	EnsureDelete(ctx, t, tc, ruleInstance)
}
