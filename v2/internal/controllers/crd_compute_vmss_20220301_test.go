/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	compute2022 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20220301"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func newVMSS20220301(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	loadBalancer *network.LoadBalancer,
	subnet *network.VirtualNetworksSubnet,
) *compute2022.VirtualMachineScaleSet {
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	upgradePolicyMode := compute2022.UpgradePolicy_Mode_Automatic
	adminUsername := "adminUser"

	inboundNATPoolRef := genruntime.ResourceReference{
		// TODO: It is the most awkward thing in the world that this is not a fully fledged resource
		ARMID: *loadBalancer.Status.InboundNatPools[0].Id,
	}

	return &compute2022.VirtualMachineScaleSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vmss")),
		Spec: compute2022.VirtualMachineScaleSet_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &compute2022.Sku{
				Name:     to.StringPtr("STANDARD_D1_v2"),
				Capacity: to.IntPtr(1),
			},
			PlatformFaultDomainCount: to.IntPtr(3),
			SinglePlacementGroup:     to.BoolPtr(false),
			UpgradePolicy: &compute2022.UpgradePolicy{
				Mode: &upgradePolicyMode,
			},
			VirtualMachineProfile: &compute2022.VirtualMachineScaleSet_Properties_VirtualMachineProfile_Spec{
				StorageProfile: &compute2022.VirtualMachineScaleSetStorageProfile{
					ImageReference: &compute2022.ImageReference{
						Publisher: to.StringPtr("Canonical"),
						Offer:     to.StringPtr("UbuntuServer"),
						Sku:       to.StringPtr("18.04-lts"),
						Version:   to.StringPtr("latest"),
					},
				},
				OsProfile: &compute2022.VirtualMachineScaleSet_Properties_VirtualMachineProfile_OsProfile_Spec{
					ComputerNamePrefix: to.StringPtr("computer"),
					AdminUsername:      &adminUsername,
					LinuxConfiguration: &compute2022.LinuxConfiguration{
						DisablePasswordAuthentication: to.BoolPtr(true),
						Ssh: &compute2022.SshConfiguration{
							PublicKeys: []compute2022.SshPublicKey{
								{
									KeyData: sshPublicKey,
									Path:    to.StringPtr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUsername)),
								},
							},
						},
					},
				},
				NetworkProfile: &compute2022.VirtualMachineScaleSet_Properties_VirtualMachineProfile_NetworkProfile_Spec{
					NetworkInterfaceConfigurations: []compute2022.VirtualMachineScaleSet_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Spec{
						{
							Name:    to.StringPtr("mynicconfig"),
							Primary: to.BoolPtr(true),
							IpConfigurations: []compute2022.VirtualMachineScaleSet_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Spec{
								{
									Name: to.StringPtr("myipconfiguration"),
									Subnet: &compute2022.ApiEntityReference{
										Reference: tc.MakeReferenceFromResource(subnet),
									},
									LoadBalancerInboundNatPools: []compute2022.SubResource{
										{
											Reference: &inboundNATPoolRef,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func Test_Compute_VMSS_20220301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	// Move to a different reason where we have quota
	tc.AzureRegion = to.StringPtr("westeurope")

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	publicIPAddress := newPublicIPAddressForVMSS(tc, testcommon.AsOwner(rg))
	loadBalancer := newLoadBalancerForVMSS(tc, rg, publicIPAddress)
	// Have to create the vnet first there's a race between it and subnet creation that
	// can change the body of the VNET PUT (because VNET PUT contains subnets)
	tc.CreateResourceAndWait(vnet)
	tc.CreateResourcesAndWait(subnet, loadBalancer, publicIPAddress)
	vmss := newVMSS20220301(tc, rg, loadBalancer, subnet)

	tc.CreateResourceAndWait(vmss)
	tc.Expect(vmss.Status.Id).ToNot(BeNil())
	armId := *vmss.Status.Id

	// Perform a simple patch to add a basic custom script extension
	old := vmss.DeepCopy()
	extensionName := "mycustomextension"
	vmss.Spec.VirtualMachineProfile.ExtensionProfile = &compute2022.VirtualMachineScaleSet_Properties_VirtualMachineProfile_ExtensionProfile_Spec{
		Extensions: []compute2022.VirtualMachineScaleSet_Properties_VirtualMachineProfile_ExtensionProfile_Extensions_Spec{
			{
				Name:               &extensionName,
				Publisher:          to.StringPtr("Microsoft.Azure.Extensions"),
				Type:               to.StringPtr("CustomScript"),
				TypeHandlerVersion: to.StringPtr("2.0"),
				Settings: map[string]v1.JSON{
					"commandToExecute": {
						Raw: []byte(`"/bin/bash -c \"echo hello\""`),
					},
				},
			},
		},
	}
	tc.PatchResourceAndWait(old, vmss)
	tc.Expect(vmss.Status.VirtualMachineProfile).ToNot(BeNil())
	tc.Expect(vmss.Status.VirtualMachineProfile.ExtensionProfile).ToNot(BeNil())
	tc.Expect(len(vmss.Status.VirtualMachineProfile.ExtensionProfile.Extensions)).To(BeNumerically(">", 0))

	found := false
	for _, extension := range vmss.Status.VirtualMachineProfile.ExtensionProfile.Extensions {
		tc.Expect(extension.Name).ToNot(BeNil())
		if *extension.Name == extensionName {
			found = true
		}
	}
	tc.Expect(found).To(BeTrue())

	// Delete VMSS
	tc.DeleteResourceAndWait(vmss)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(compute2022.APIVersion_Value))
}
