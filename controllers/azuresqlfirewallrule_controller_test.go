// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azuresqlfirewall
// +build all azuresqlfirewall

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAzureSqlFirewallRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	//rgName := tc.resourceGroupName
	sqlServerName := GenerateTestResourceNameWithRandom("sqlfwrule-test-srv", 10)
	sqlFirewallRuleName := GenerateTestResourceNameWithRandom("fwrule-dev", 10)

	// Create the SqlFirewallRule object and expect the Reconcile to be created
	sqlFirewallRuleInstance := &v1beta1.AzureSqlFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlFirewallRuleName,
			Namespace: "default",
		},
		Spec: v1beta1.AzureSqlFirewallRuleSpec{
			ResourceGroup:  GenerateTestResourceNameWithRandom("rg-fake-srv", 10),
			Server:         sqlServerName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, sqlFirewallRuleInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, sqlFirewallRuleInstance)

}
