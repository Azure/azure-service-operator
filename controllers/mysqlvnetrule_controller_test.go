// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || mysql
// +build all mysql

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

func TestMySqlVNetRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	mySqlServerName := GenerateTestResourceNameWithRandom("mysqlvnetrule-test-srv", 10)
	mySqlVNetRuleName := GenerateTestResourceNameWithRandom("vnetrule-dev", 10)

	// Create the SqlVnetRule object and expect the Reconcile to be created
	mySqlVNetRuleInstance := &azurev1alpha1.MySQLVNetRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mySqlVNetRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLVNetRuleSpec{
			ResourceGroup:                GenerateTestResourceNameWithRandom("rg-fake-srv", 10),
			Server:                       mySqlServerName,
			VNetResourceGroup:            "vnet-rg",
			VNetName:                     "test-vnet",
			SubnetName:                   "subnet1",
			IgnoreMissingServiceEndpoint: true,
		},
	}

	err := tc.k8sClient.Create(ctx, mySqlVNetRuleInstance)
	assert.Equal(nil, err, "create mysqlvnetrule in k8s")

	mySqlVNETRuleNamespacedName := types.NamespacedName{Name: mySqlVNetRuleName, Namespace: "default"}

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, mySqlVNETRuleNamespacedName, mySqlVNetRuleInstance)
		if err == nil {
			return HasFinalizer(mySqlVNetRuleInstance, finalizerName)
		} else {
			return false
		}
	}, tc.timeout, tc.retry, "wait for mysqlvnetrule to have finalizer")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, mySqlVNETRuleNamespacedName, mySqlVNetRuleInstance)
		if err == nil {
			return strings.Contains(mySqlVNetRuleInstance.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
		} else {
			return false
		}
	}, tc.timeout, tc.retry, "wait for mysqlvnetrule to have rg not found error")

	err = tc.k8sClient.Delete(ctx, mySqlVNetRuleInstance)
	assert.Equal(nil, err, "delete mysqlvnetrule in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, mySqlVNETRuleNamespacedName, mySqlVNetRuleInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for mysqlvnetrule to be gone from k8s")
}

func RunMySqlVNetRuleHappyPath(t *testing.T, mySqlServerName string, rgLocation string) {
	defer PanicRecover(t)
	ctx := context.Background()

	mySqlVNetRuleName := GenerateTestResourceNameWithRandom("vnet-rule", 10)
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

	mySqlVNetRuleInstance := &azurev1alpha1.MySQLVNetRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mySqlVNetRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLVNetRuleSpec{
			ResourceGroup:                tc.resourceGroupName,
			Server:                       mySqlServerName,
			VNetResourceGroup:            tc.resourceGroupName,
			VNetName:                     VNetName,
			SubnetName:                   subnetName,
			IgnoreMissingServiceEndpoint: true,
		},
	}

	// Create VNet Rule and ensure it was created
	EnsureInstance(ctx, t, tc, mySqlVNetRuleInstance)

	// Delete a VNet Rule and ensure it was deleted
	EnsureDelete(ctx, t, tc, mySqlVNetRuleInstance)

}
