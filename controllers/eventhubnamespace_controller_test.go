// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all eventhubnamespace

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	config "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/stretchr/testify/assert"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestEventHubNamespaceControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgLocation string
	rgLocation = tc.resourceGroupLocation

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	// setting this rg name tells the mocks to set a proper error
	resourceGroupName := "gone"
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

func TestEventHubNamespaceControllerBasicTierNetworkRule(t *testing.T) {
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

	// Create a VNET
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

	// Create EventhubNamespace network rule for this namespace and expect success
	subnetID := "/subscriptions/" + config.SubscriptionID() + "/resourceGroups/" + rgName + "/providers/Microsoft.Network/virtualNetworks/" + VNetName + "/subnets/" + subnetName
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

	// Create the Eventhub namespace object as prereq
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
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

	EnsureInstanceWithResult(ctx, t, tc, eventhubNamespaceInstance, errhelp.BadRequest, false)

	// Delete the namespace
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)

}

func TestEventHubNamespaceControllerHappy(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)
	var err error

	var rgName string = tc.resourceGroupName
	var rgLocation string = tc.resourceGroupLocation
	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Create the Eventhub namespace object and expect the Reconcile to be created
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
		},
	}

	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance)

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

	err = tc.k8sClient.Delete(ctx, eventhubNamespaceInstance)
	assert.Equal(nil, err, "delete eventhubns in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubNamespaceInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for eventHubnamespaceInstance to be gone from k8s")
}

func TestEventHubNamespaceControllerHappyWithNetworkRule(t *testing.T) {
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

	// Create a VNET
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

	// Create EventhubNamespace network rule for this namespace and expect success
	subnetID := "/subscriptions/" + config.SubscriptionID() + "/resourceGroups/" + rgName + "/providers/Microsoft.Network/virtualNetworks/" + VNetName + "/subnets/" + subnetName
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

	// Create the Eventhub namespace object as prereq
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
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

	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance)

	// Delete the namespace
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)

}
