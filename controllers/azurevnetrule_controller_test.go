// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver azuresqlservercombined testvnetrule

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlVNetRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

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

	err := tc.k8sClient.Create(ctx, sqlVNetRuleInstance)
	assert.Equal(nil, err, "create sqlvnetrule in k8s")

	sqlVNETRuleNamespacedName := types.NamespacedName{Name: sqlVNetRuleName, Namespace: "default"}

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlVNETRuleNamespacedName, sqlVNetRuleInstance)
		if err == nil {
			return HasFinalizer(sqlVNetRuleInstance, finalizerName)
		} else {
			return false
		}
	}, tc.timeout, tc.retry, "wait for sqlvnetrule to have finalizer")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlVNETRuleNamespacedName, sqlVNetRuleInstance)
		if err == nil {
			return strings.Contains(sqlVNetRuleInstance.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
		} else {
			return false
		}
	}, tc.timeout, tc.retry, "wait for sqlvnetrule to have rg not found error")

	err = tc.k8sClient.Delete(ctx, sqlVNetRuleInstance)
	assert.Equal(nil, err, "delete sqlvnetrule in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlVNETRuleNamespacedName, sqlVNetRuleInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for sqlvnetrule to be gone from k8s")
}

func RunAzureSqlVNetRuleHappyPath(t *testing.T, sqlServerName string) {
	defer PanicRecover()
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
			Location:      tc.resourceGroupLocation,
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
