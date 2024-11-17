/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	compute2022 "github.com/Azure/azure-service-operator/v2/api/compute/v1api20220301"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
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
				Name:     to.Ptr("STANDARD_D1_v2"),
				Capacity: to.Ptr(1),
			},
			PlatformFaultDomainCount: to.Ptr(3),
			SinglePlacementGroup:     to.Ptr(false),
			UpgradePolicy: &compute2022.UpgradePolicy{
				Mode: &upgradePolicyMode,
			},
			VirtualMachineProfile: &compute2022.VirtualMachineScaleSetVMProfile{
				StorageProfile: &compute2022.VirtualMachineScaleSetStorageProfile{
					ImageReference: &compute2022.ImageReference{
						Publisher: to.Ptr("Canonical"),
						Offer:     to.Ptr("UbuntuServer"),
						Sku:       to.Ptr("18.04-lts"),
						Version:   to.Ptr("latest"),
					},
				},
				OsProfile: &compute2022.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.Ptr("computer"),
					AdminUsername:      &adminUsername,
					LinuxConfiguration: &compute2022.LinuxConfiguration{
						DisablePasswordAuthentication: to.Ptr(true),
						Ssh: &compute2022.SshConfiguration{
							PublicKeys: []compute2022.SshPublicKeySpec{
								{
									KeyData: sshPublicKey,
									Path:    to.Ptr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUsername)),
								},
							},
						},
					},
				},
				NetworkProfile: &compute2022.VirtualMachineScaleSetNetworkProfile{
					NetworkInterfaceConfigurations: []compute2022.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name:    to.Ptr("mynicconfig"),
							Primary: to.Ptr(true),
							IpConfigurations: []compute2022.VirtualMachineScaleSetIPConfiguration{
								{
									Name: to.Ptr("myipconfiguration"),
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
	tc.AzureRegion = to.Ptr("westeurope")

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
	extensionName2 := "mycustomextension2"
	vmss.Spec.VirtualMachineProfile.ExtensionProfile = &compute2022.VirtualMachineScaleSetExtensionProfile{
		Extensions: []compute2022.VirtualMachineScaleSetExtension{
			{
				Name:               &extensionName,
				Publisher:          to.Ptr("Microsoft.Azure.Extensions"),
				Type:               to.Ptr("CustomScript"),
				TypeHandlerVersion: to.Ptr("2.0"),
				Settings: map[string]v1.JSON{
					"commandToExecute": {
						Raw: []byte(`"/bin/bash -c \"echo hello\""`),
					},
				},
			},
		},
	}
	tc.PatchResourceAndWait(old, vmss)
	found := checkExtensionExists2022(tc, vmss, extensionName)
	tc.Expect(found).To(BeTrue())

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "VMSS_Extension_20220301_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				VMSS_Extension_20220301_CRUD(tc, vmss, extensionName2)
			},
		},
	)

	objectKey := client.ObjectKeyFromObject(vmss)

	tc.Eventually(func() bool {
		var updated compute2022.VirtualMachineScaleSet
		tc.GetResource(objectKey, &updated)
		return checkExtensionExists2022(tc, &updated, extensionName)
	}).Should(BeTrue())

	// Delete VMSS
	tc.DeleteResourceAndWait(vmss)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(compute2022.APIVersion_Value))
}

func checkExtensionExists2022(tc *testcommon.KubePerTestContext, vmss *compute2022.VirtualMachineScaleSet, extensionName string) bool {
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

	return found
}

func VMSS_Extension_20220301_CRUD(tc *testcommon.KubePerTestContext, vmss *compute2022.VirtualMachineScaleSet, extensionName string) {
	extension := &compute2022.VirtualMachineScaleSetsExtension{
		ObjectMeta: tc.MakeObjectMetaWithName(extensionName),
		Spec: compute2022.VirtualMachineScaleSetsExtension_Spec{
			Owner:              testcommon.AsOwner(vmss),
			Publisher:          to.Ptr("Microsoft.ManagedServices"),
			Type:               to.Ptr("ApplicationHealthLinux"),
			TypeHandlerVersion: to.Ptr("1.0"),
		},
	}

	tc.CreateResourceAndWait(extension)
	tc.Expect(extension.Status.Id).ToNot(BeNil())
	armId := *extension.Status.Id

	tc.DeleteResourceAndWait(extension)

	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(compute2022.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
