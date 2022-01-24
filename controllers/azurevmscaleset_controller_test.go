// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azurevmscaleset
// +build all azurevmscaleset

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestVMScaleSetControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgName := GenerateTestResourceNameWithRandom("rg", 10)
	vmName := GenerateTestResourceNameWithRandom("vm", 10)
	vmSize := "Standard_DS1_v2"
	capacity := 2
	osType := azurev1alpha1.OSType("Linux")
	adminUserName := GenerateTestResourceNameWithRandom("u", 10)
	sshPublicKeyData := GenerateTestResourceNameWithRandom("ssh", 10)
	platformImageUrn := "Canonical:UbuntuServer:16.04-LTS:latest"
	vnetName := GenerateTestResourceNameWithRandom("vn", 10)
	subnetName := "test"
	lbName := GenerateTestResourceNameWithRandom("lb", 10)
	beName := GenerateTestResourceNameWithRandom("be", 10)
	natName := GenerateTestResourceNameWithRandom("nat", 10)

	// Create a VMSS
	vmssInstance := &azurev1alpha1.AzureVMScaleSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vmName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureVMScaleSetSpec{
			Location:               tc.resourceGroupLocation,
			ResourceGroup:          rgName,
			VMSize:                 vmSize,
			Capacity:               capacity,
			OSType:                 osType,
			AdminUserName:          adminUserName,
			SSHPublicKeyData:       sshPublicKeyData,
			PlatformImageURN:       platformImageUrn,
			VirtualNetworkName:     vnetName,
			SubnetName:             subnetName,
			LoadBalancerName:       lbName,
			BackendAddressPoolName: beName,
			InboundNatPoolName:     natName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, vmssInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, vmssInstance)
}

func TestVMScaleSetHappyPathWithLbPipVNetAndSubnet(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// location := tc.resourceGroupLocation
	location := "australiaeast"

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
			Location:      location,
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
			Location:                 location,
			ResourceGroup:            tc.resourceGroupName,
			PublicIPAllocationMethod: publicIPAllocationMethod,
			IdleTimeoutInMinutes:     idleTimeoutInMinutes,
			PublicIPAddressVersion:   publicIPAddressVersion,
			SkuName:                  skuName,
		},
	}

	EnsureInstance(ctx, t, tc, pipInstance)

	// Create a Load Balancer
	lbName := GenerateTestResourceNameWithRandom("lb", 10)
	bpName := "test"
	natName := "test"
	portRangeStart := 100
	portRangeEnd := 200
	backendPort := 300
	lbInstance := &azurev1alpha1.AzureLoadBalancer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      lbName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureLoadBalancerSpec{
			Location:               location,
			ResourceGroup:          tc.resourceGroupName,
			PublicIPAddressName:    pipName,
			BackendAddressPoolName: bpName,
			InboundNatPoolName:     natName,
			FrontendPortRangeStart: portRangeStart,
			FrontendPortRangeEnd:   portRangeEnd,
			BackendPort:            backendPort,
		},
	}

	EnsureInstance(ctx, t, tc, lbInstance)

	// Create a VMSS
	vmName := GenerateTestResourceNameWithRandom("vm", 10)
	vmSize := "Standard_DS1_v2"
	capacity := 3
	osType := azurev1alpha1.OSType("Linux")
	vmImageUrn := "Canonical:UbuntuServer:16.04-LTS:latest"
	userName := "azureuser"
	sshPublicKeyData := GenerateRandomSshPublicKeyString()

	vmssInstance := &azurev1alpha1.AzureVMScaleSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vmName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureVMScaleSetSpec{
			Location:               location,
			ResourceGroup:          tc.resourceGroupName,
			VMSize:                 vmSize,
			Capacity:               capacity,
			OSType:                 osType,
			AdminUserName:          userName,
			SSHPublicKeyData:       sshPublicKeyData,
			PlatformImageURN:       vmImageUrn,
			VirtualNetworkName:     vnetName,
			SubnetName:             subnetName,
			LoadBalancerName:       lbName,
			BackendAddressPoolName: bpName,
			InboundNatPoolName:     natName,
		},
	}

	EnsureInstance(ctx, t, tc, vmssInstance)

	EnsureDelete(ctx, t, tc, vmssInstance)

	EnsureDelete(ctx, t, tc, lbInstance)

	EnsureDelete(ctx, t, tc, pipInstance)

	EnsureDelete(ctx, t, tc, vnetInstance)
}
