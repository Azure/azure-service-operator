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

	compute2022 "github.com/Azure/azure-service-operator/v2/api/compute/v1api20220301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// This test cannot be run in record/replay mode because the state it looks for
// is not "stable" (the reconciler keeps retrying). Since it keeps retrying there isn't a deterministic number of
// retry attempts it makes which means a recording test may run out of recorded retries.
func Test_VMSSExceedsQuota_ContinuesToRetry(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	adminUsername := "adminUser"
	vmss := &compute2022.VirtualMachineScaleSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vmss")),
		Spec: compute2022.VirtualMachineScaleSet_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &compute2022.Sku{
				Name:     to.Ptr("STANDARD_D1_v2"),
				Capacity: to.Ptr(10000), // We can't possibly have this quota
			},
			PlatformFaultDomainCount: to.Ptr(3),
			SinglePlacementGroup:     to.Ptr(false),
			UpgradePolicy: &compute2022.UpgradePolicy{
				Mode: to.Ptr(compute2022.UpgradePolicy_Mode_Automatic),
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
					AdminUsername:      to.Ptr(adminUsername),
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
								},
							},
						},
					},
				},
			},
		},
	}

	tc.CreateResourceAndWaitForState(vmss, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(vmss.Status.Conditions[0].Reason).To(Equal("OperationNotAllowed"))
}
