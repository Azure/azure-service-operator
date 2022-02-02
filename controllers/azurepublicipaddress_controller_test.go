// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azurepublicipaddress
// +build all azurepublicipaddress

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPublicIPAddressIdleTimeoutIsOutOfRange(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	pipName := GenerateTestResourceNameWithRandom("pip1", 10)
	publicIPAllocationMethod := "Static"
	idleTimeoutInMinutes := -1
	publicIPAddressVersion := "IPv4"
	skuName := "Basic"

	// Create a Public IP Address
	pipInstance := &azurev1alpha1.AzurePublicIPAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pipName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzurePublicIPAddressSpec{
			Location:                 rgLocation,
			ResourceGroup:            rgName,
			PublicIPAllocationMethod: publicIPAllocationMethod,
			IdleTimeoutInMinutes:     idleTimeoutInMinutes,
			PublicIPAddressVersion:   publicIPAddressVersion,
			SkuName:                  skuName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, pipInstance, errhelp.PublicIPIdleTimeoutIsOutOfRange, false)

	EnsureDelete(ctx, t, tc, pipInstance)
}

func TestPublicIPAddressHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	pipName := GenerateTestResourceNameWithRandom("pip2", 10)
	publicIPAllocationMethod := "Static"
	idleTimeoutInMinutes := 10
	publicIPAddressVersion := "IPv4"
	skuName := "Basic"

	// Create a Public IP Address
	pipInstance := &azurev1alpha1.AzurePublicIPAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pipName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzurePublicIPAddressSpec{
			Location:                 rgLocation,
			ResourceGroup:            rgName,
			PublicIPAllocationMethod: publicIPAllocationMethod,
			IdleTimeoutInMinutes:     idleTimeoutInMinutes,
			PublicIPAddressVersion:   publicIPAddressVersion,
			SkuName:                  skuName,
		},
	}

	EnsureInstance(ctx, t, tc, pipInstance)

	EnsureDelete(ctx, t, tc, pipInstance)
}

func TestPublicIPAddressControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("pip-rand-rg", 10)
	rgLocation := tc.resourceGroupLocation
	pipName := GenerateTestResourceNameWithRandom("pip4", 10)
	publicIPAllocationMethod := "Static"
	idleTimeoutInMinutes := 10
	publicIPAddressVersion := "IPv4"
	skuName := "Basic"

	// Create a Public IP Address
	pipInstance := &azurev1alpha1.AzurePublicIPAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pipName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzurePublicIPAddressSpec{
			Location:                 rgLocation,
			ResourceGroup:            rgName,
			PublicIPAllocationMethod: publicIPAllocationMethod,
			IdleTimeoutInMinutes:     idleTimeoutInMinutes,
			PublicIPAddressVersion:   publicIPAddressVersion,
			SkuName:                  skuName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, pipInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, pipInstance)
}
