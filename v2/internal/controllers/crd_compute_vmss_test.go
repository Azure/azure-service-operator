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

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func newVMVirtualNetwork(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.VirtualNetwork {
	return &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetwork_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}
}

func newVMSubnet(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.VirtualNetworksSubnet {
	return &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworksSubnets_Spec{
			Owner:         owner,
			AddressPrefix: to.StringPtr("10.0.0.0/24"),
		},
	}
}

func newPublicIPAddressForVMSS(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.PublicIPAddress {
	publicIPAddressSku := network.PublicIPAddressSkuNameStandard
	allocationMethod := network.PublicIPAddressPropertiesFormatPublicIPAllocationMethodStatic
	return &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddresses_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &network.PublicIPAddressSku{
				Name: &publicIPAddressSku,
			},
			PublicIPAllocationMethod: &allocationMethod,
		},
	}
}

func newLoadBalancerForVMSS(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, publicIPAddress *network.PublicIPAddress) *network.LoadBalancer {
	loadBalancerSku := network.LoadBalancerSkuNameStandard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network.InboundNatPoolPropertiesFormatProtocolTcp

	// TODO: Getting this is SUPER awkward
	frontIPConfigurationARMID, err := genericarmclient.MakeResourceGroupScopeARMID(
		tc.AzureSubscription,
		rg.Name,
		"Microsoft.Network",
		"loadBalancers",
		lbName,
		"frontendIPConfigurations",
		lbFrontendName)
	if err != nil {
		panic(err)
	}

	return &network.LoadBalancer{
		ObjectMeta: tc.MakeObjectMetaWithName(lbName),
		Spec: network.LoadBalancers_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			FrontendIPConfigurations: []network.LoadBalancers_Spec_Properties_FrontendIPConfigurations{
				{
					Name: &lbFrontendName,
					PublicIPAddress: &network.SubResource{
						Reference: tc.MakeReferenceFromResource(publicIPAddress),
					},
				},
			},
			InboundNatPools: []network.LoadBalancers_Spec_Properties_InboundNatPools{
				{
					Name: to.StringPtr("MyFancyNatPool"),
					FrontendIPConfiguration: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: frontIPConfigurationARMID,
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
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	loadBalancer *network.LoadBalancer,
	subnet *network.VirtualNetworksSubnet,
) *compute.VirtualMachineScaleSet {
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
			Owner:    testcommon.AsOwner(rg),
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
				OsProfile: &compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_OsProfile{
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
							Name:    to.StringPtr("mynicconfig"),
							Primary: to.BoolPtr(true),
							IpConfigurations: []compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations{
								{
									Name: to.StringPtr("myipconfiguration"),
									Subnet: &compute.ApiEntityReference{
										Reference: tc.MakeReferenceFromResource(subnet),
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

func Test_Compute_VMSS_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	publicIPAddress := newPublicIPAddressForVMSS(tc, testcommon.AsOwner(rg))
	loadBalancer := newLoadBalancerForVMSS(tc, rg, publicIPAddress)
	// Have to create the vnet first there's a race between it and subnet creation that
	// can change the body of the VNET PUT (because VNET PUT contains subnets)
	tc.CreateResourceAndWait(vnet)
	tc.CreateResourcesAndWait(subnet, loadBalancer, publicIPAddress)
	vmss := newVMSS(tc, rg, loadBalancer, subnet)

	tc.CreateResourceAndWait(vmss)
	tc.Expect(vmss.Status.Id).ToNot(BeNil())
	armId := *vmss.Status.Id

	// Perform a simple patch to add a basic custom script extension
	old := vmss.DeepCopy()
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
	tc.PatchResourceAndWait(old, vmss)
	tc.Expect(vmss.Status.VirtualMachineProfile).ToNot(BeNil())
	tc.Expect(vmss.Status.VirtualMachineProfile.ExtensionProfile).ToNot(BeNil())
	tc.Expect(vmss.Status.VirtualMachineProfile.ExtensionProfile.Extensions).To(HaveLen(1))
	tc.Expect(vmss.Status.VirtualMachineProfile.ExtensionProfile.Extensions[0].Name).ToNot(BeNil())
	tc.Expect(*vmss.Status.VirtualMachineProfile.ExtensionProfile.Extensions[0].Name).To(Equal(extensionName))

	// Delete VMSS
	tc.DeleteResourceAndWait(vmss)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(compute.APIVersionValue))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
