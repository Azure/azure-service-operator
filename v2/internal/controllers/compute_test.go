/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"

	. "github.com/onsi/gomega"

	compute2020 "github.com/Azure/azure-service-operator/v2/api/compute/v1api20201201"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	v1 "k8s.io/api/core/v1"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func newVirtualMachine20201201(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	networkInterface *network.NetworkInterface,
	secretRef genruntime.SecretReference,
) *compute2020.VirtualMachine {
	adminUsername := "bloom"
	size := compute2020.HardwareProfile_VmSize_Standard_D1_V2

	return &compute2020.VirtualMachine{
		ObjectMeta: tc.MakeObjectMeta("vm"),
		Spec: compute2020.VirtualMachine_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			HardwareProfile: &compute2020.HardwareProfile{
				VmSize: &size,
			},
			OsProfile: &compute2020.OSProfile{
				AdminUsername: &adminUsername,
				// Specifying AdminPassword here rather than SSH Key to ensure that handling and injection
				// of secrets works.
				AdminPassword: &secretRef,
				ComputerName:  to.Ptr("poppy"),
			},
			StorageProfile: &compute2020.StorageProfile{
				ImageReference: &compute2020.ImageReference{
					Offer:     to.Ptr("UbuntuServer"),
					Publisher: to.Ptr("Canonical"),
					Sku:       to.Ptr("18.04-LTS"),
					Version:   to.Ptr("latest"),
				},
			},
			NetworkProfile: &compute2020.NetworkProfile{
				NetworkInterfaces: []compute2020.NetworkInterfaceReference{
					{
						Reference: tc.MakeReferenceFromResource(networkInterface),
					},
				},
			},
		},
	}
}

func newVMNetworkInterface(
	tc *testcommon.KubePerTestContext,
	owner *genruntime.KnownResourceReference,
	subnet *network.VirtualNetworksSubnet,
) *network.NetworkInterface {
	dynamic := network.IPAllocationMethod_Dynamic
	return &network.NetworkInterface{
		ObjectMeta: tc.MakeObjectMeta("nic"),
		Spec: network.NetworkInterface_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			IpConfigurations: []network.NetworkInterfaceIPConfiguration_NetworkInterface_SubResourceEmbedded{{
				Name:                      to.Ptr("ipconfig1"),
				PrivateIPAllocationMethod: &dynamic,
				Subnet: &network.Subnet_NetworkInterface_SubResourceEmbedded{
					Reference: tc.MakeReferenceFromResource(subnet),
				},
			}},
		},
	}
}

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

// newVMSS20201201 creates a new VirtualMachineScaleSet for testing
func newVMSS20201201(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	loadBalancer *network.LoadBalancer,
	subnet *network.VirtualNetworksSubnet,
) *compute2020.VirtualMachineScaleSet {
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	upgradePolicyMode := compute2020.UpgradePolicy_Mode_Automatic
	adminUsername := "adminUser"

	inboundNATPoolRef := genruntime.ResourceReference{
		// TODO: It is the most awkward thing in the world that this is not a fully fledged resource
		ARMID: *loadBalancer.Status.InboundNatPools[0].Id,
	}

	return &compute2020.VirtualMachineScaleSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vmss")),
		Spec: compute2020.VirtualMachineScaleSet_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &compute2020.Sku{
				Name:     to.Ptr("STANDARD_D1_v2"),
				Capacity: to.Ptr(1),
			},
			PlatformFaultDomainCount: to.Ptr(3),
			SinglePlacementGroup:     to.Ptr(false),
			UpgradePolicy: &compute2020.UpgradePolicy{
				Mode: &upgradePolicyMode,
			},
			VirtualMachineProfile: &compute2020.VirtualMachineScaleSetVMProfile{
				StorageProfile: &compute2020.VirtualMachineScaleSetStorageProfile{
					ImageReference: &compute2020.ImageReference{
						Publisher: to.Ptr("Canonical"),
						Offer:     to.Ptr("UbuntuServer"),
						Sku:       to.Ptr("18.04-lts"),
						Version:   to.Ptr("latest"),
					},
				},
				OsProfile: &compute2020.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.Ptr("computer"),
					AdminUsername:      &adminUsername,
					LinuxConfiguration: &compute2020.LinuxConfiguration{
						DisablePasswordAuthentication: to.Ptr(true),
						Ssh: &compute2020.SshConfiguration{
							PublicKeys: []compute2020.SshPublicKeySpec{
								{
									KeyData: sshPublicKey,
									Path:    to.Ptr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUsername)),
								},
							},
						},
					},
				},
				NetworkProfile: &compute2020.VirtualMachineScaleSetNetworkProfile{
					NetworkInterfaceConfigurations: []compute2020.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name:    to.Ptr("mynicconfig"),
							Primary: to.Ptr(true),
							IpConfigurations: []compute2020.VirtualMachineScaleSetIPConfiguration{
								{
									Name: to.Ptr("myipconfiguration"),
									Subnet: &compute2020.ApiEntityReference{
										Reference: tc.MakeReferenceFromResource(subnet),
									},
									LoadBalancerInboundNatPools: []compute2020.SubResource{
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
