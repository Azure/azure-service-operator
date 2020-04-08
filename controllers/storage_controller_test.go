// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all storage

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	config "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestStorageControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	StorageAccountName := GenerateAlphaNumTestResourceName("sadev")

	// Create the ResourceGroup object and expect the Reconcile to be created
	saInstance := &azurev1alpha1.Storage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      StorageAccountName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.StorageSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			Sku: azurev1alpha1.StorageSku{
				Name: "Standard_RAGRS",
			},
			Kind:                   "StorageV2",
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
		},
	}

	// create rg
	EnsureInstance(ctx, t, tc, saInstance)

	// delete rg
	EnsureDelete(ctx, t, tc, saInstance)
}

func TestStorageControllerHappyPathWithNetworkRule(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	StorageAccountName := GenerateAlphaNumTestResourceName("cndev")

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	VNetName := GenerateTestResourceNameWithRandom("vnet", 10)
	subnetName := "subnet-test"

	subnetID := "/subscriptions/" + config.SubscriptionID() + "/resourceGroups/" + rgName + "/providers/Microsoft.Network/virtualNetworks/" + VNetName + "/subnets/" + subnetName
	vnetRules := []azurev1alpha1.VirtualNetworkRule{
		{
			SubnetId: &subnetID,
		},
	}
	ipAddress := "1.1.1.1"
	ipRange := "2.2.2.2/24"
	ipRules := []azurev1alpha1.IPRule{
		{
			IPAddressOrRange: &ipAddress,
		},
		{
			IPAddressOrRange: &ipRange,
		},
	}

	// Create the ResourceGroup object and expect the Reconcile to be created
	cnInstance := &azurev1alpha1.Storage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      StorageAccountName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.StorageSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.StorageSku{
				Name: "Standard_RAGRS",
			},
			Kind:                   "StorageV2",
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
			NetworkRule: &azurev1alpha1.StorageNetworkRuleSet{
				Bypass:              "AzureServices",
				VirtualNetworkRules: &vnetRules,
				IPRules:             &ipRules,
				DefaultAction:       "Deny",
			},
		},
	}

	// create rg
	EnsureInstance(ctx, t, tc, cnInstance)

	// delete rg
	EnsureDelete(ctx, t, tc, cnInstance)
}
