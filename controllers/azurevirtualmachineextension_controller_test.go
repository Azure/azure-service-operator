// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azurevirtualmachineextension
// +build all azurevirtualmachineextension

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestVirtualMachineExtensionControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgName := GenerateTestResourceNameWithRandom("rg", 10)
	extName := GenerateTestResourceNameWithRandom("ext", 10)
	vmName := GenerateTestResourceNameWithRandom("vm", 10)
	// Try to create a VM extension
	extInstance := &azurev1alpha1.AzureVirtualMachineExtension{
		ObjectMeta: metav1.ObjectMeta{
			Name:      extName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureVirtualMachineExtensionSpec{
			Location:                tc.resourceGroupLocation,
			ResourceGroup:           rgName,
			VMName:                  vmName,
			AutoUpgradeMinorVersion: true,
			ForceUpdateTag:          "test",
			Publisher:               "test",
			TypeName:                "test",
			TypeHandlerVersion:      "1.0",
			Settings:                "{}",
			ProtectedSettings:       "{}",
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, extInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, extInstance)
}

func TestVirtualMachineExtensionHappyPath(t *testing.T) {
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

	// Create a NIC
	nicName := GenerateTestResourceNameWithRandom("nic", 10)
	nicInstance := &azurev1alpha1.AzureNetworkInterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nicName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureNetworkInterfaceSpec{
			Location:            location,
			ResourceGroup:       tc.resourceGroupName,
			VNetName:            vnetName,
			SubnetName:          subnetName,
			PublicIPAddressName: pipName,
		},
	}

	EnsureInstance(ctx, t, tc, nicInstance)

	// Create a VM
	vmName := GenerateTestResourceNameWithRandom("vm", 10)
	vmSize := "Standard_DS1_v2"
	osType := azurev1alpha1.OSType("Linux")
	vmImageUrn := "Canonical:UbuntuServer:16.04-LTS:latest"
	userName := "azureuser"

	sshPublicKeyData := GenerateRandomSshPublicKeyString()

	vmInstance := &azurev1alpha1.AzureVirtualMachine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vmName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureVirtualMachineSpec{
			Location:             location,
			ResourceGroup:        tc.resourceGroupName,
			VMSize:               vmSize,
			OSType:               osType,
			AdminUserName:        userName,
			SSHPublicKeyData:     sshPublicKeyData,
			NetworkInterfaceName: nicName,
			PlatformImageURN:     vmImageUrn,
		},
	}

	EnsureInstance(ctx, t, tc, vmInstance)

	// To create a VM extension
	extName := GenerateTestResourceNameWithRandom("ext", 10)
	extInstance := &azurev1alpha1.AzureVirtualMachineExtension{
		ObjectMeta: metav1.ObjectMeta{
			Name:      extName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureVirtualMachineExtensionSpec{
			Location:                location,
			ResourceGroup:           tc.resourceGroupName,
			VMName:                  vmName,
			AutoUpgradeMinorVersion: true,
			ForceUpdateTag:          "test",
			Publisher:               "Microsoft.Azure.Extensions",
			TypeName:                "CustomScript",
			TypeHandlerVersion:      "2.1",
			Settings:                "{\"commandToExecute\":\"echo 'hello, world'\"}",
			ProtectedSettings:       "{}",
		},
	}

	EnsureInstance(ctx, t, tc, extInstance)

	EnsureDelete(ctx, t, tc, extInstance)

	EnsureDelete(ctx, t, tc, vmInstance)

	EnsureDelete(ctx, t, tc, nicInstance)

	EnsureDelete(ctx, t, tc, pipInstance)

	EnsureDelete(ctx, t, tc, vnetInstance)
}
