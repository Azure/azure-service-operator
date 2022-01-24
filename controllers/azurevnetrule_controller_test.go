// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azuresqlserver || azuresqlservercombined || testvnetrule
// +build all azuresqlserver azuresqlservercombined testvnetrule

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAzureSqlVNetRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	sqlServerName := GenerateTestResourceNameWithRandom("sqlvnetrule-test-srv", 10)
	sqlVNetRuleName := GenerateTestResourceNameWithRandom("vnetrule-dev", 10)

	// Create the SqlVnetRule object and expect the Reconcile to be created
	sqlVNetRuleInstance := &azurev1alpha1.AzureSQLVNetRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlVNetRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSQLVNetRuleSpec{
			ResourceGroup:                GenerateTestResourceNameWithRandom("rg-fake-srv", 10),
			Server:                       sqlServerName,
			VNetResourceGroup:            "vnet-rg",
			VNetName:                     "test-vnet",
			SubnetName:                   "subnet1",
			IgnoreMissingServiceEndpoint: true,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, sqlVNetRuleInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, sqlVNetRuleInstance)
}

func RunAzureSqlVNetRuleHappyPath(t *testing.T, sqlServerName string, rgLocation string) {
	defer PanicRecover(t)
	ctx := context.Background()

	sqlVNetRuleName := GenerateTestResourceNameWithRandom("vnet-rule", 10)
	VNetName := GenerateTestResourceNameWithRandom("vnet", 10)
	subnetName := "subnet-test"
	VNetSubNetInstance := azurev1alpha1.VNetSubnets{
		SubnetName:          subnetName,
		SubnetAddressPrefix: "110.1.0.0/16",
	}

	// Create a VNET
	VNetInstance := &azurev1alpha1.VirtualNetwork{
		ObjectMeta: metav1.ObjectMeta{
			Name:      VNetName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.VirtualNetworkSpec{
			Location:      rgLocation,
			ResourceGroup: tc.resourceGroupName,
			AddressSpace:  "110.0.0.0/8",
			Subnets:       []azurev1alpha1.VNetSubnets{VNetSubNetInstance},
		},
	}

	EnsureInstance(ctx, t, tc, VNetInstance)

	// Create a VNet Rule

	sqlVNetRuleInstance := &azurev1alpha1.AzureSQLVNetRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlVNetRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSQLVNetRuleSpec{
			ResourceGroup:                tc.resourceGroupName,
			Server:                       sqlServerName,
			VNetResourceGroup:            tc.resourceGroupName,
			VNetName:                     VNetName,
			SubnetName:                   subnetName,
			IgnoreMissingServiceEndpoint: true,
		},
	}

	// Create VNet Rule and ensure it was created
	EnsureInstance(ctx, t, tc, sqlVNetRuleInstance)

	// Delete a VNet Rule and ensure it was deleted
	EnsureDelete(ctx, t, tc, sqlVNetRuleInstance)

}
