// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azurevirtualmachine

package controllers

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"golang.org/x/crypto/ssh"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func generateRandomSshPublicKeyString() string {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicRsaKey, _ := ssh.NewPublicKey(&privateKey.PublicKey)
	sshPublicKeyData := string(ssh.MarshalAuthorizedKey(publicRsaKey))
	return sshPublicKeyData
}

func TestVirtualMachineControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgName := GenerateTestResourceNameWithRandom("rg", 10)
	vmName := GenerateTestResourceNameWithRandom("vm", 10)
	vmSize := "Standard_DS1_v2"
	osType := azurev1alpha1.OSType("Linux")
	adminUserName := GenerateTestResourceNameWithRandom("u", 10)
	sshPublicKeyData := GenerateTestResourceNameWithRandom("ssh", 10)
	nicName := GenerateTestResourceNameWithRandom("nic", 10)
	platformImageUrn := "Canonical:UbuntuServer:16.04-LTS:latest"

	// Create a VM
	vmInstance := &azurev1alpha1.AzureVirtualMachine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vmName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureVirtualMachineSpec{
			Location:             tc.resourceGroupLocation,
			ResourceGroup:        rgName,
			VMSize:               vmSize,
			OSType:               osType,
			AdminUserName:        adminUserName,
			SSHPublicKeyData:     sshPublicKeyData,
			NetworkInterfaceName: nicName,
			PlatformImageURN:     platformImageUrn,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, vmInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, vmInstance)
}

func TestVirtualMachineHappyPathWithNicPipVNetAndSubnet(t *testing.T) {
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

	// Create a VM
	vmName := GenerateTestResourceNameWithRandom("vm", 10)
	vmSize := "Standard_DS1_v2"
	osType := azurev1alpha1.OSType("Linux")
	vmImageUrn := "Canonical:UbuntuServer:16.04-LTS:latest"
	userName := "azureuser"

	sshPublicKeyData := generateRandomSshPublicKeyString()

	vmInstance := &azurev1alpha1.AzureVirtualMachine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vmName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureVirtualMachineSpec{
			Location:             tc.resourceGroupLocation,
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

	EnsureDelete(ctx, t, tc, vmInstance)

	EnsureDelete(ctx, t, tc, nicInstance)

	EnsureDelete(ctx, t, tc, pipInstance)

	EnsureDelete(ctx, t, tc, vnetInstance)
}
