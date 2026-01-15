/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1api20201201"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// This test cannot be run in record/replay mode because the state it looks for (Ready = false with warning)
// is not "stable" (the reconciler keeps retrying). Since it keeps retrying there isn't a deterministic number of
// retry attempts it makes which means a recording test may run out of recorded retries.
// There are two ways that a long-running operation can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long-running operation polling. This ensures that the first case is handled correctly.
// TODO: Should move to simulated error test once we have https://github.com/Azure/azure-service-operator/issues/2164
func Test_OperationRejected(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	vmss := newVMSSWithInvalidPublisher(tc, nil, nil, rg)
	tc.CreateResourceAndWaitForState(vmss, metav1.ConditionFalse, conditions.ConditionSeverityWarning)

	ready, ok := conditions.GetCondition(vmss, conditions.ConditionTypeReady)
	tc.Expect(ok).To(BeTrue())

	tc.Expect(ready.Status).To(Equal(metav1.ConditionFalse))
	tc.Expect(ready.Severity).To(Equal(conditions.ConditionSeverityWarning))
	tc.Expect(ready.Reason).To(ContainSubstring("InvalidParameter"))
	tc.Expect(ready.Message).To(ContainSubstring("The value of parameter imageReference.publisher is invalid"))
}

// This test cannot be run in record/replay mode because the state it looks for (Ready = false with warning)
// is not "stable" (the reconciler keeps retrying). Since it keeps retrying there isn't a deterministic number of
// retry attempts it makes which means a recording test may run out of recorded retries.
// There are two ways that a long-running operation can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long-running operation polling. This ensures that a resource in the first case
// can be updated to resolve the cause of the failure and successfully deployed.
// TODO: Should move to simulated error test once we have https://github.com/Azure/azure-service-operator/issues/2164
func Test_OperationRejected_SucceedsAfterUpdate(t *testing.T) {
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
	vmss := newVMSSWithInvalidPublisher(tc, loadBalancer, subnet, rg)

	tc.CreateResourceAndWaitForState(vmss, metav1.ConditionFalse, conditions.ConditionSeverityWarning)

	// Now fix the image reference and the VMSS should successfully deploy
	old := vmss.DeepCopy()
	vmss.Spec.VirtualMachineProfile.StorageProfile.ImageReference = newImageReference()
	tc.PatchResourceAndWait(old, vmss)

	// Ensure that the old failure information was cleared away
	objectKey := client.ObjectKeyFromObject(vmss)
	updated := &compute.VirtualMachineScaleSet{}
	tc.GetResource(objectKey, updated)

	ready, ok := conditions.GetCondition(updated, conditions.ConditionTypeReady)
	tc.Expect(ok).To(BeTrue())

	tc.Expect(ready.Status).To(Equal(metav1.ConditionTrue))
	tc.Expect(ready.Severity).To(Equal(conditions.ConditionSeverityNone))
	tc.Expect(ready.Reason).To(Equal(conditions.ReasonSucceeded))
	tc.Expect(ready.Message).To(Equal(""))
}

func newVMSSWithInvalidPublisher(tc *testcommon.KubePerTestContext, loadBalancer *network.LoadBalancer, subnet *network.VirtualNetworksSubnet, rg *resources.ResourceGroup) *compute.VirtualMachineScaleSet {
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	upgradePolicyMode := compute.UpgradePolicy_Mode_Automatic
	adminUsername := "adminUser"

	var inboundNATPoolRef *genruntime.ResourceReference
	if loadBalancer != nil {
		inboundNATPoolRef = &genruntime.ResourceReference{
			// TODO: It is the most awkward thing in the world that this is not a fully fledged resource
			ARMID: *loadBalancer.Status.InboundNatPools[0].Id,
		}
	}
	var subnetRef *compute.ApiEntityReference
	if subnet != nil {
		subnetRef = &compute.ApiEntityReference{
			Reference: tc.MakeReferenceFromResource(subnet),
		}
	}

	return &compute.VirtualMachineScaleSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vmss")),
		Spec: compute.VirtualMachineScaleSet_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &compute.Sku{
				Name:     to.Ptr("Standard_D2s_v3"),
				Capacity: to.Ptr(1),
			},
			PlatformFaultDomainCount: to.Ptr(3),
			SinglePlacementGroup:     to.Ptr(false),
			UpgradePolicy: &compute.UpgradePolicy{
				Mode: &upgradePolicyMode,
			},
			VirtualMachineProfile: &compute.VirtualMachineScaleSetVMProfile{
				StorageProfile: &compute.VirtualMachineScaleSetStorageProfile{
					ImageReference: &compute.ImageReference{
						Publisher: to.Ptr("this publisher"),
						Offer:     to.Ptr("does not"),
						Sku:       to.Ptr("exist"),
						Version:   to.Ptr("latest"),
					},
				},
				OsProfile: &compute.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.Ptr("computer"),
					AdminUsername:      &adminUsername,
					LinuxConfiguration: &compute.LinuxConfiguration{
						DisablePasswordAuthentication: to.Ptr(true),
						Ssh: &compute.SshConfiguration{
							PublicKeys: []compute.SshPublicKeySpec{
								{
									KeyData: sshPublicKey,
									Path:    to.Ptr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUsername)),
								},
							},
						},
					},
				},
				NetworkProfile: &compute.VirtualMachineScaleSetNetworkProfile{
					NetworkInterfaceConfigurations: []compute.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name:    to.Ptr("mynicconfig"),
							Primary: to.Ptr(true),
							IpConfigurations: []compute.VirtualMachineScaleSetIPConfiguration{
								{
									Name:   to.Ptr("myipconfiguration"),
									Subnet: subnetRef,
									LoadBalancerInboundNatPools: []compute.SubResource{
										{
											Reference: inboundNATPoolRef,
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

func newImageReference() *compute.ImageReference {
	return &compute.ImageReference{
		Publisher: to.Ptr("Canonical"),
		Offer:     to.Ptr("ubuntu-24_04-lts"),
		Sku:       to.Ptr("server"),
		Version:   to.Ptr("latest"),
	}
}

func newPublicIPAddressForVMSS(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.PublicIPAddress {
	publicIPAddressSku := network.PublicIPAddressSku_Name_Standard
	allocationMethod := network.IPAllocationMethod_Static
	return &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddress_Spec{
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
	loadBalancerSku := network.LoadBalancerSku_Name_Standard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network.TransportProtocol_Tcp

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
		Spec: network.LoadBalancer_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			FrontendIPConfigurations: []network.FrontendIPConfiguration_LoadBalancer_SubResourceEmbedded{
				{
					Name: &lbFrontendName,
					PublicIPAddress: &network.PublicIPAddressSpec_LoadBalancer_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(publicIPAddress),
					},
				},
			},
			InboundNatPools: []network.InboundNatPool{
				{
					Name: to.Ptr("MyFancyNatPool"),
					FrontendIPConfiguration: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: frontIPConfigurationARMID,
						},
					},
					Protocol:               &protocol,
					FrontendPortRangeStart: to.Ptr(50_000),
					FrontendPortRangeEnd:   to.Ptr(51_000),
					BackendPort:            to.Ptr(22),
				},
			},
		},
	}
}
