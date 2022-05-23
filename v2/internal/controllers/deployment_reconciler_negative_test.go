/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/go-autorest/autorest/to"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func newStorageAccountWithInvalidKeyExpiration(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *storage.StorageAccount {
	// Custom namer because storage accounts have strict names

	// Create a storage account with an invalid key expiration period
	accessTier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	kind := storage.StorageAccountsSpecKindBlobStorage
	sku := storage.SkuNameStandardLRS
	return &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccounts_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     &kind,
			Sku: &storage.Sku{
				Name: &sku,
			},
			AccessTier: &accessTier,
			KeyPolicy: &storage.KeyPolicy{
				KeyExpirationPeriodInDays: to.IntPtr(-260),
			},
		},
	}
}

func newVMSSWithInvalidPublisher(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *compute.VirtualMachineScaleSet {
	upgradePolicyMode := compute.UpgradePolicyModeAutomatic
	adminUsername := "adminUser"
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
						Publisher: to.StringPtr("this publisher"),
						Offer:     to.StringPtr("does not"),
						Sku:       to.StringPtr("exist"),
						Version:   to.StringPtr("latest"),
					},
				},
				OsProfile: &compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_OsProfile{
					ComputerNamePrefix: to.StringPtr("computer"),
					AdminUsername:      &adminUsername,
				},
				NetworkProfile: &compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile{
					NetworkInterfaceConfigurations: []compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations{
						{
							Name: to.StringPtr("mynicconfig"),
							IpConfigurations: []compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations{
								{
									Name: to.StringPtr("test"),
								},
							},
						},
					},
				},
			},
		},
	}
}

// There are two ways that a long-running operation can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long running operation polling. This ensures that the second case is handled correctly.
func Test_OperationAccepted_LongRunningOperationFails(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	acct := newStorageAccountWithInvalidKeyExpiration(tc, rg)
	tc.CreateResourceAndWaitForFailure(acct)

	ready, ok := conditions.GetCondition(acct, conditions.ConditionTypeReady)
	tc.Expect(ok).To(BeTrue())

	tc.Expect(ready.Status).To(Equal(metav1.ConditionFalse))
	tc.Expect(ready.Severity).To(Equal(conditions.ConditionSeverityError))
	tc.Expect(ready.Reason).To(Equal("InvalidValuesForRequestParameters"))
	tc.Expect(ready.Message).To(ContainSubstring("Values for request parameters are invalid: keyPolicy.keyExpirationPeriodInDays."))
}

// There are two ways that a long-running operation can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long-running operation polling. This ensures that a resource in the second case
// can be updated to resolve the cause of the failure and successfully deployed.
func Test_OperationAccepted_LongRunningOperationFails_SucceedsAfterUpdate(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	acct := newStorageAccountWithInvalidKeyExpiration(tc, rg)
	tc.CreateResourceAndWaitForFailure(acct)

	// Remove the bad property and ensure we can successfully provision
	old := acct.DeepCopy()
	acct.Spec.KeyPolicy = nil
	tc.PatchResourceAndWait(old, acct)

	// Ensure that the old failure information was cleared away
	objectKey := client.ObjectKeyFromObject(acct)
	updated := &storage.StorageAccount{}
	tc.GetResource(objectKey, updated)

	ready, ok := conditions.GetCondition(acct, conditions.ConditionTypeReady)
	tc.Expect(ok).To(BeTrue())

	tc.Expect(ready.Status).To(Equal(metav1.ConditionTrue))
	tc.Expect(ready.Severity).To(Equal(conditions.ConditionSeverityNone))
	tc.Expect(ready.Reason).To(Equal(conditions.ReasonSucceeded))
	tc.Expect(ready.Message).To(Equal(""))
}

// There are two ways that a long-running operation can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long-running operation polling. This ensures that the first case is handled correctly.
func Test_OperationRejected(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	vmss := newVMSSWithInvalidPublisher(tc, rg)
	tc.CreateResourceAndWaitForFailure(vmss)

	ready, ok := conditions.GetCondition(vmss, conditions.ConditionTypeReady)
	tc.Expect(ok).To(BeTrue())

	tc.Expect(ready.Status).To(Equal(metav1.ConditionFalse))
	tc.Expect(ready.Severity).To(Equal(conditions.ConditionSeverityError))
	tc.Expect(ready.Reason).To(ContainSubstring("InvalidParameter"))
	tc.Expect(ready.Message).To(ContainSubstring("The value of parameter imageReference.publisher is invalid"))
}

// There are two ways that a long-running operation can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long-running operation polling. This ensures that a resource in the first case
// can be updated to resolve the cause of the failure and successfully deployed.
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
	vmss := newVMSS(tc, rg, loadBalancer, subnet)
	imgRef := vmss.Spec.VirtualMachineProfile.StorageProfile.ImageReference
	originalImgRef := imgRef.DeepCopy()

	// Set the VMSS to have an invalid image
	imgRef.Publisher = to.StringPtr("this publisher")
	imgRef.Offer = to.StringPtr("does not")
	imgRef.Sku = to.StringPtr("exist")
	imgRef.Version = to.StringPtr("latest")

	tc.CreateResourceAndWaitForFailure(vmss)

	// Now fix the image reference and the VMSS should successfully deploy
	old := vmss.DeepCopy()
	vmss.Spec.VirtualMachineProfile.StorageProfile.ImageReference = originalImgRef
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
