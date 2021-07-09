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
	"sigs.k8s.io/controller-runtime/pkg/client"

	compute "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.compute/v1alpha1api20201201"
	network "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.network/v1alpha1api20201101"
	resources "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func newVNETForVMSS(tc testcommon.KubePerTestContext, owner genruntime.KnownResourceReference) *network.VirtualNetwork {
	return &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetworks_Spec{
			Owner:    owner,
			Location: testcommon.DefaultTestRegion,
			AddressSpace: network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}
}

func newSubnetForVMSS(tc testcommon.KubePerTestContext, owner genruntime.KnownResourceReference) *network.VirtualNetworksSubnet {
	return &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworksSubnets_Spec{
			Owner:         owner,
			AddressPrefix: "10.0.0.0/24",
		},
	}
}

func newPublicIPAddressForVMSS(tc testcommon.KubePerTestContext, owner genruntime.KnownResourceReference) *network.PublicIPAddress {
	publicIPAddressSku := network.PublicIPAddressSkuNameStandard
	return &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddresses_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &network.PublicIPAddressSku{
				Name: &publicIPAddressSku,
			},
			PublicIPAllocationMethod: network.PublicIPAddressPropertiesFormatPublicIPAllocationMethodStatic,
		},
	}
}

func newLoadBalancerForVMSS(tc testcommon.KubePerTestContext, rg *resources.ResourceGroup, publicIPAddress *network.PublicIPAddress) *network.LoadBalancer {
	loadBalancerSku := network.LoadBalancerSkuNameStandard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network.InboundNatPoolPropertiesFormatProtocolTcp
	return &network.LoadBalancer{
		ObjectMeta: tc.MakeObjectMetaWithName(lbName),
		Spec: network.LoadBalancers_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &network.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			FrontendIPConfigurations: []network.LoadBalancers_Spec_Properties_FrontendIPConfigurations{
				{
					Name: lbFrontendName,
					PublicIPAddress: &network.SubResource{
						Reference: tc.MakeReferenceFromResource(publicIPAddress),
					},
				},
			},
			InboundNatPools: []network.LoadBalancers_Spec_Properties_InboundNatPools{
				{
					Name: "MyFancyNatPool",
					FrontendIPConfiguration: &network.SubResource{
						Reference: genruntime.ResourceReference{
							// TODO: Getting this is SUPER awkward
							ARMID: tc.MakeARMId(rg.Name, "Microsoft.Network", "loadBalancers", lbName, "frontendIPConfigurations", lbFrontendName),
						},
					},
					Protocol:               &protocol,
					FrontendPortRangeStart: to.IntPtr(50_000),
					FrontendPortRangeEnd:   to.IntPtr(51_000),
					BackendPort:            to.IntPtr(22),
				},
			},
		},
	}
}

func newVMSS(
	tc testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	loadBalancer *network.LoadBalancer,
	subnet *network.VirtualNetworksSubnet) *compute.VirtualMachineScaleSet {

	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	upgradePolicyMode := compute.UpgradePolicyModeAutomatic
	adminUsername := "adminUser"

	inboundNATPoolRef := genruntime.ResourceReference{
		// TODO: It is the most awkward thing in the world that this is not a fully fledged resource
		ARMID: *loadBalancer.Status.InboundNatPools[0].Id,
	}

	return &compute.VirtualMachineScaleSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vmss")),
		Spec: compute.VirtualMachineScaleSets_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &compute.Sku{
				Name:     to.StringPtr("STANDARD_D1_v2"),
				Capacity: to.IntPtr(1),
			},
			PlatformFaultDomainCount: to.IntPtr(3),
			SinglePlacementGroup:     to.BoolPtr(false),
			UpgradePolicy: &compute.UpgradePolicy{
				Mode: &upgradePolicyMode,
			},
			VirtualMachineProfile: &compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile{
				StorageProfile: &compute.VirtualMachineScaleSetStorageProfile{
					ImageReference: &compute.ImageReference{
						Publisher: to.StringPtr("Canonical"),
						Offer:     to.StringPtr("UbuntuServer"),
						Sku:       to.StringPtr("18.04-lts"),
						Version:   to.StringPtr("latest"),
					},
				},
				OsProfile: &compute.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.StringPtr("computer"),
					AdminUsername:      &adminUsername,
					LinuxConfiguration: &compute.LinuxConfiguration{
						DisablePasswordAuthentication: to.BoolPtr(true),
						Ssh: &compute.SshConfiguration{
							PublicKeys: []compute.SshPublicKey{
								{
									KeyData: sshPublicKey,
									Path:    to.StringPtr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUsername)),
								},
							},
						},
					},
				},
				NetworkProfile: &compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile{
					NetworkInterfaceConfigurations: []compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations{
						{
							Name:    "mynicconfig",
							Primary: to.BoolPtr(true),
							IpConfigurations: []compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations{
								{
									Name: "myipconfiguration",
									Subnet: &compute.ApiEntityReference{
										Reference: tc.MakeReferencePtrFromResource(subnet),
									},
									LoadBalancerInboundNatPools: []compute.SubResource{
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

func Test_VMSS_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateNewTestResourceGroupAndWait()

	vnet := newVNETForVMSS(tc, testcommon.AsOwner(rg.ObjectMeta))
	subnet := newSubnetForVMSS(tc, testcommon.AsOwner(vnet.ObjectMeta))
	publicIPAddress := newPublicIPAddressForVMSS(tc, testcommon.AsOwner(rg.ObjectMeta))
	loadBalancer := newLoadBalancerForVMSS(tc, rg, publicIPAddress)
	tc.CreateResourcesAndWait(vnet, subnet, loadBalancer, publicIPAddress)
	vmss := newVMSS(tc, rg, loadBalancer, subnet)

	tc.CreateResourceAndWait(vmss)
	tc.Expect(vmss.Status.Id).ToNot(BeNil())
	armId := *vmss.Status.Id

	// Perform a simple patch to add a basic custom script extension
	patcher := tc.NewResourcePatcher(vmss)
	extensionName := "mycustomextension"
	vmss.Spec.VirtualMachineProfile.ExtensionProfile = &compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile{
		Extensions: []compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_Extensions{
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
	patcher.Patch(vmss)

	objectKey, err := client.ObjectKeyFromObject(vmss)
	tc.Expect(err).ToNot(HaveOccurred())

	// Ensure state eventually gets updated in k8s from change in Azure.
	tc.Eventually(func() string {
		var updatedVMSS compute.VirtualMachineScaleSet
		tc.GetResource(objectKey, &updatedVMSS)

		vmProfile := updatedVMSS.Status.VirtualMachineProfile
		if vmProfile == nil {
			return ""
		}

		if vmProfile.ExtensionProfile == nil {
			return ""
		}

		if len(vmProfile.ExtensionProfile.Extensions) == 0 {
			return ""
		}

		return *updatedVMSS.Status.VirtualMachineProfile.ExtensionProfile.Extensions[0].Name
	}).Should(BeEquivalentTo(extensionName))

	// Delete VMSS
	tc.DeleteResourceAndWait(vmss)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, string(compute.VirtualMachineScaleSetsSpecAPIVersion20201201))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
