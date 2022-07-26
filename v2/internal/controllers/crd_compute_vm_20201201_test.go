/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	compute2020 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func createVMPasswordSecretAndRef(tc *testcommon.KubePerTestContext) genruntime.SecretReference {
	password := tc.Namer.GeneratePasswordOfLength(40)

	passwordKey := "password"
	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta("vmsecret"),
		StringData: map[string]string{
			passwordKey: password,
		},
	}

	tc.CreateResource(secret)

	secretRef := genruntime.SecretReference{
		Name: secret.Name,
		Key:  passwordKey,
	}
	return secretRef
}

func newVirtualMachine20201201(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	networkInterface *network.NetworkInterface,
	secretRef genruntime.SecretReference,
) *compute2020.VirtualMachine {
	adminUsername := "bloom"
	size := compute2020.HardwareProfileVmSizeStandardD1V2

	return &compute2020.VirtualMachine{
		ObjectMeta: tc.MakeObjectMeta("vm"),
		Spec: compute2020.VirtualMachines_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			HardwareProfile: &compute2020.HardwareProfile{
				VmSize: &size,
			},
			OsProfile: &compute2020.VirtualMachines_Spec_Properties_OsProfile{
				AdminUsername: &adminUsername,
				// Specifying AdminPassword here rather than SSH Key to ensure that handling and injection
				// of secrets works.
				AdminPassword: &secretRef,
				ComputerName:  to.StringPtr("poppy"),
			},
			StorageProfile: &compute2020.StorageProfile{
				ImageReference: &compute2020.ImageReference{
					Offer:     to.StringPtr("UbuntuServer"),
					Publisher: to.StringPtr("Canonical"),
					Sku:       to.StringPtr("18.04-LTS"),
					Version:   to.StringPtr("latest"),
				},
			},
			NetworkProfile: &compute2020.VirtualMachines_Spec_Properties_NetworkProfile{
				NetworkInterfaces: []compute2020.VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfaces{{
					Reference: tc.MakeReferenceFromResource(networkInterface),
				}},
			},
		},
	}
}

func newVMNetworkInterface(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, subnet *network.VirtualNetworksSubnet) *network.NetworkInterface {
	dynamic := network.NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethodDynamic
	return &network.NetworkInterface{
		ObjectMeta: tc.MakeObjectMeta("nic"),
		Spec: network.NetworkInterfaces_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			IpConfigurations: []network.NetworkInterfaces_Spec_Properties_IpConfigurations{{
				Name:                      to.StringPtr("ipconfig1"),
				PrivateIPAllocationMethod: &dynamic,
				Subnet: &network.SubResource{
					Reference: tc.MakeReferenceFromResource(subnet),
				},
			}},
		},
	}
}

func Test_Compute_VM_20201201_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	tc.AzureRegion = to.StringPtr("westeurope")

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	networkInterface := newVMNetworkInterface(tc, testcommon.AsOwner(rg), subnet)
	// Inefficient but avoids triggering the vnet/subnets problem.
	// https://github.com/Azure/azure-service-operator/issues/1944
	tc.CreateResourceAndWait(vnet)
	tc.CreateResourcesAndWait(subnet, networkInterface)
	secret := createVMPasswordSecretAndRef(tc)
	vm := newVirtualMachine20201201(tc, rg, networkInterface, secret)

	tc.CreateResourceAndWait(vm)
	tc.Expect(vm.Status.Id).ToNot(BeNil())
	armId := *vm.Status.Id

	// Perform a simple patch to turn on boot diagnostics
	old := vm.DeepCopy()
	vm.Spec.DiagnosticsProfile = &compute2020.DiagnosticsProfile{
		BootDiagnostics: &compute2020.BootDiagnostics{
			Enabled: to.BoolPtr(true),
		},
	}

	tc.PatchResourceAndWait(old, vm)
	tc.Expect(vm.Status.DiagnosticsProfile).ToNot(BeNil())
	tc.Expect(vm.Status.DiagnosticsProfile.BootDiagnostics).ToNot(BeNil())
	tc.Expect(vm.Status.DiagnosticsProfile.BootDiagnostics.Enabled).ToNot(BeNil())
	tc.Expect(*vm.Status.DiagnosticsProfile.BootDiagnostics.Enabled).To(BeTrue())

	// Delete VM and resources.
	tc.DeleteResourcesAndWait(vm, networkInterface, subnet, vnet, rg)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(compute2020.APIVersionValue))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
