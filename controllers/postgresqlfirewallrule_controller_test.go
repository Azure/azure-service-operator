// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || psqlfirewallrule
// +build all psqlfirewallrule

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPSQLFirewallRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("psqlsrv-rg", 10)

	postgreSQLServerName := GenerateTestResourceNameWithRandom("psql-srv", 10)
	postgreSQLFirewallRuleName := GenerateTestResourceNameWithRandom("psql-fwrule", 10)

	// Create the PostgreSQLFirewallRule object and expect the Reconcile to be created
	postgreSQLFirewallRuleInstance := &azurev1alpha1.PostgreSQLFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLFirewallRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLFirewallRuleSpec{
			ResourceGroup:  rgName,
			Server:         postgreSQLServerName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, postgreSQLFirewallRuleInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, postgreSQLFirewallRuleInstance)

}

func TestPSQLFirewallRuleControllerNoServer(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName

	postgreSQLServerName := GenerateTestResourceNameWithRandom("psql-srv", 10)
	postgreSQLFirewallRuleName := GenerateTestResourceNameWithRandom("psql-fwrule", 10)

	// Create the PostgreSQLFirewallRule object and expect the Reconcile to be created
	postgreSQLFirewallRuleInstance := &azurev1alpha1.PostgreSQLFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLFirewallRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLFirewallRuleSpec{
			ResourceGroup:  rgName,
			Server:         postgreSQLServerName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, postgreSQLFirewallRuleInstance, errhelp.ResourceNotFound, false)
	EnsureDelete(ctx, t, tc, postgreSQLFirewallRuleInstance)

}
