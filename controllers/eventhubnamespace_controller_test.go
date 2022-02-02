// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || eventhubnamespace
// +build all eventhubnamespace

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	config "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestEventHubNamespaceControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgLocation string
	rgLocation = tc.resourceGroupLocation
	resourceGroupName := GenerateTestResourceNameWithRandom("rg", 10)
	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Create the EventHubNamespace object and expect the Reconcile to be created
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: resourceGroupName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, eventhubNamespaceInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)
}

func TestEventHubNamespaceControllerNetworkRules(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgName string = tc.resourceGroupName
	var rgLocation string = tc.resourceGroupLocation

	// Create a VNET as prereq for the test
	VNetName := GenerateTestResourceNameWithRandom("vnet", 10)
	subnetName := "subnet-test"
	VNetSubNetInstance := azurev1alpha1.VNetSubnets{
		SubnetName:          subnetName,
		SubnetAddressPrefix: "110.1.0.0/16",
	}

	// Create a VNET as prereq
	VNetInstance := &azurev1alpha1.VirtualNetwork{
		ObjectMeta: metav1.ObjectMeta{
			Name:      VNetName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.VirtualNetworkSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			AddressSpace:  "110.0.0.0/8",
			Subnets:       []azurev1alpha1.VNetSubnets{VNetSubNetInstance},
		},
	}

	EnsureInstance(ctx, t, tc, VNetInstance)

	// Create EventhubNamespace network rule using the above VNET
	subnetID := "/subscriptions/" + config.GlobalCredentials().SubscriptionID() + "/resourceGroups/" + rgName + "/providers/Microsoft.Network/virtualNetworks/" + VNetName + "/subnets/" + subnetName
	vnetRules := []azurev1alpha1.VirtualNetworkRules{
		{
			SubnetID:                     subnetID,
			IgnoreMissingServiceEndpoint: true,
		},
	}
	ipmask := "1.1.1.1"
	ipRules := []azurev1alpha1.IPRules{
		{
			IPMask: &ipmask,
		},
	}

	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Add this to the Eventhubnamespace that is Basic Tier and expect to fail
	eventhubNamespaceInstance1 := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.EventhubNamespaceSku{
				Name:     "Basic",
				Tier:     "Basic",
				Capacity: 1,
			},
			NetworkRule: &azurev1alpha1.EventhubNamespaceNetworkRule{
				DefaultAction:       "deny",
				VirtualNetworkRules: &vnetRules,
				IPRules:             &ipRules,
			},
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, eventhubNamespaceInstance1, errhelp.BadRequest, false)

	eventhubNamespaceName = GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Add the network rule to a namespace that is Standard SKU and expect to pass
	eventhubNamespaceInstance2 := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.EventhubNamespaceSku{
				Name:     "Standard",
				Tier:     "Standard",
				Capacity: 1,
			},
			NetworkRule: &azurev1alpha1.EventhubNamespaceNetworkRule{
				DefaultAction:       "deny",
				VirtualNetworkRules: &vnetRules,
				IPRules:             &ipRules,
			},
		},
	}

	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance2)

	// Delete the namespace
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance2)

	// Delete the namespace
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance1)

	EnsureDelete(ctx, t, tc, VNetInstance)

}
