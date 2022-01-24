// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azurenetworkinterface
// +build all azurenetworkinterface

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNetworkInterfaceNonExistingSubResources(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	nicName := GenerateTestResourceNameWithRandom("nic", 10)
	vnetName := GenerateTestResourceNameWithRandom("vnt", 10)
	subnetName := GenerateTestResourceNameWithRandom("snt", 10)
	pipName := GenerateTestResourceNameWithRandom("pip", 10)

	// Create a NIC
	nicInstance := &azurev1alpha1.AzureNetworkInterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nicName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureNetworkInterfaceSpec{
			Location:            tc.resourceGroupLocation,
			ResourceGroup:       tc.resourceGroupName,
			VNetName:            vnetName,
			SubnetName:          subnetName,
			PublicIPAddressName: pipName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, nicInstance, errhelp.InvalidResourceReference, false)

	EnsureDelete(ctx, t, tc, nicInstance)

}

func TestNetworkInterfaceHappyPathWithVNetAndSubnet(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Create a vnet with a subnet
	vnetName := GenerateTestResourceNameWithRandom("vnt", 10)
	subnetName := GenerateTestResourceNameWithRandom("snt", 10)
	vnetSubNetInstance := azurev1alpha1.VNetSubnets{
		SubnetName:          subnetName,
		SubnetAddressPrefix: "110.1.0.0/16",
	}
	pipName := GenerateTestResourceNameWithRandom("pip3", 10)
	vnetInstance := &azurev1alpha1.VirtualNetwork{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vnetName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.VirtualNetworkSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			AddressSpace:  "110.0.0.0/8",
			Subnets:       []azurev1alpha1.VNetSubnets{vnetSubNetInstance},
		},
	}

	EnsureInstance(ctx, t, tc, vnetInstance)

	// Create a Public IP Address
	publicIPAllocationMethod := "Static"
	idleTimeoutInMinutes := 10
	publicIPAddressVersion := "IPv4"
	skuName := "Basic"
	pipInstance := &azurev1alpha1.AzurePublicIPAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pipName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzurePublicIPAddressSpec{
			Location:                 tc.resourceGroupLocation,
			ResourceGroup:            tc.resourceGroupName,
			PublicIPAllocationMethod: publicIPAllocationMethod,
			IdleTimeoutInMinutes:     idleTimeoutInMinutes,
			PublicIPAddressVersion:   publicIPAddressVersion,
			SkuName:                  skuName,
		},
	}

	EnsureInstance(ctx, t, tc, pipInstance)

	// Create a NIC
	nicName := GenerateTestResourceNameWithRandom("nic", 10)
	nicInstance := &azurev1alpha1.AzureNetworkInterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nicName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureNetworkInterfaceSpec{
			Location:            tc.resourceGroupLocation,
			ResourceGroup:       tc.resourceGroupName,
			VNetName:            vnetName,
			SubnetName:          subnetName,
			PublicIPAddressName: pipName,
		},
	}

	EnsureInstance(ctx, t, tc, nicInstance)

	EnsureDelete(ctx, t, tc, nicInstance)

	EnsureDelete(ctx, t, tc, pipInstance)

	EnsureDelete(ctx, t, tc, vnetInstance)
}

func TestNetworkInterfaceControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("nic-rand-rg", 10)

	// Create a vnet with a subnet
	vnetName := GenerateTestResourceNameWithRandom("vnt5", 10)
	subnetName := GenerateTestResourceNameWithRandom("snt5", 10)
	vnetSubNetInstance := azurev1alpha1.VNetSubnets{
		SubnetName:          subnetName,
		SubnetAddressPrefix: "110.1.0.0/16",
	}
	pipName := GenerateTestResourceNameWithRandom("pip5", 10)
	vnetInstance := &azurev1alpha1.VirtualNetwork{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vnetName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.VirtualNetworkSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			AddressSpace:  "110.0.0.0/8",
			Subnets:       []azurev1alpha1.VNetSubnets{vnetSubNetInstance},
		},
	}

	EnsureInstance(ctx, t, tc, vnetInstance)

	// Create a Public IP Address
	publicIPAllocationMethod := "Static"
	idleTimeoutInMinutes := 10
	publicIPAddressVersion := "IPv4"
	skuName := "Basic"
	pipInstance := &azurev1alpha1.AzurePublicIPAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pipName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzurePublicIPAddressSpec{
			Location:                 tc.resourceGroupLocation,
			ResourceGroup:            tc.resourceGroupName,
			PublicIPAllocationMethod: publicIPAllocationMethod,
			IdleTimeoutInMinutes:     idleTimeoutInMinutes,
			PublicIPAddressVersion:   publicIPAddressVersion,
			SkuName:                  skuName,
		},
	}

	EnsureInstance(ctx, t, tc, pipInstance)

	// Create a NIC
	nicName := GenerateTestResourceNameWithRandom("nic5", 10)
	nicInstance := &azurev1alpha1.AzureNetworkInterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nicName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureNetworkInterfaceSpec{
			Location:            tc.resourceGroupLocation,
			ResourceGroup:       rgName,
			VNetName:            vnetName,
			SubnetName:          subnetName,
			PublicIPAddressName: pipName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, nicInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, nicInstance)

	EnsureDelete(ctx, t, tc, pipInstance)

	EnsureDelete(ctx, t, tc, vnetInstance)
}
