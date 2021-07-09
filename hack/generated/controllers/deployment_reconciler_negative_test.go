/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/go-autorest/autorest/to"

	compute "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.compute/v1alpha1api20201201"
	resources "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	storage "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.storage/v1alpha1api20210401"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/reconcilers"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func newStorageAccountWithInvalidKeyExpiration(tc testcommon.KubePerTestContext, rg *resources.ResourceGroup) *storage.StorageAccount {
	// Custom namer because storage accounts have strict names
	namer := tc.Namer.WithSeparator("")

	// Create a storage account with an invalid key expiration period
	accessTier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	return &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("stor")),
		Spec: storage.StorageAccounts_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Kind:     storage.StorageAccountsSpecKindBlobStorage,
			Sku: storage.Sku{
				Name: storage.SkuNameStandardLRS,
			},
			AccessTier: &accessTier,
			KeyPolicy: &storage.KeyPolicy{
				KeyExpirationPeriodInDays: -260,
			},
		},
	}
}

func newVMSSWithInvalidPublisher(tc testcommon.KubePerTestContext, rg *resources.ResourceGroup) *compute.VirtualMachineScaleSet {
	upgradePolicyMode := compute.UpgradePolicyModeAutomatic
	adminUsername := "adminUser"
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
						Publisher: to.StringPtr("this publisher"),
						Offer:     to.StringPtr("does not"),
						Sku:       to.StringPtr("exist"),
						Version:   to.StringPtr("latest"),
					},
				},
				OsProfile: &compute.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.StringPtr("computer"),
					AdminUsername:      &adminUsername,
				},
				NetworkProfile: &compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile{
					NetworkInterfaceConfigurations: []compute.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations{
						{
							Name: "mynicconfig",
						},
					},
				},
			},
		},
	}
}

// There are two ways that a deployment can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long running operation polling. This ensures that the second case is handled correctly.
func Test_DeploymentAccepted_LongRunningOperationFails(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateNewTestResourceGroupAndWait()
	acct := newStorageAccountWithInvalidKeyExpiration(tc, rg)
	tc.CreateResourceAndWaitForFailure(acct)

	errAnnotation := acct.Annotations[reconcilers.ResourceErrorAnnotation]
	tc.Expect(errAnnotation).To(ContainSubstring("InvalidValuesForRequestParameters"))
	tc.Expect(errAnnotation).To(ContainSubstring("Values for request parameters are invalid: keyPolicy.keyExpirationPeriodInDays."))
}

// There are two ways that a deployment can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long running operation polling. This ensures that a resource in the second case
// can be updated to resolve the cause of the failure and successfully deployed.
func Test_DeploymentAccepted_LongRunningOperationFails_SucceedsAfterUpdate(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateNewTestResourceGroupAndWait()
	acct := newStorageAccountWithInvalidKeyExpiration(tc, rg)
	tc.CreateResourceAndWaitForFailure(acct)

	// Remove the bad property and ensure we can successfully provision
	patcher := tc.NewResourcePatcher(acct)
	acct.Spec.KeyPolicy = nil
	tc.PatchResourceAndWaitAfter(acct, patcher, armclient.FailedProvisioningState)

	// Ensure that the old failure information was cleared away
	objectKey, err := client.ObjectKeyFromObject(acct)
	tc.Expect(err).ToNot(HaveOccurred())
	updated := &storage.StorageAccount{}
	tc.GetResource(objectKey, updated)

	tc.Expect(updated.Annotations).ToNot(HaveKey(reconcilers.ResourceErrorAnnotation))
}

// There are two ways that a deployment can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long running operation polling. This ensures that the first case is handled correctly.
func Test_DeploymentRejected(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateNewTestResourceGroupAndWait()
	vmss := newVMSSWithInvalidPublisher(tc, rg)
	tc.CreateResourceAndWaitForFailure(vmss)

	errAnnotation := vmss.Annotations[reconcilers.ResourceErrorAnnotation]
	tc.Expect(errAnnotation).To(ContainSubstring("InvalidTemplateDeployment"))
	tc.Expect(errAnnotation).To(ContainSubstring("InvalidParameter"))
	tc.Expect(errAnnotation).To(ContainSubstring("The value of parameter imageReference.publisher is invalid"))
}

// There are two ways that a deployment can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long running operation polling. This ensures that a resource in the first case
// can be updated to resolve the cause of the failure and successfully deployed.
func Test_DeploymentRejected_SucceedsAfterUpdate(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateNewTestResourceGroupAndWait()

	vnet := newVNETForVMSS(tc, testcommon.AsOwner(rg.ObjectMeta))
	subnet := newSubnetForVMSS(tc, testcommon.AsOwner(vnet.ObjectMeta))
	publicIPAddress := newPublicIPAddressForVMSS(tc, testcommon.AsOwner(rg.ObjectMeta))
	loadBalancer := newLoadBalancerForVMSS(tc, rg, publicIPAddress)
	tc.CreateResourcesAndWait(vnet, subnet, loadBalancer, publicIPAddress)
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
	patcher := tc.NewResourcePatcher(vmss)
	vmss.Spec.VirtualMachineProfile.StorageProfile.ImageReference = originalImgRef
	tc.PatchResourceAndWaitAfter(vmss, patcher, armclient.FailedProvisioningState)

	// Ensure that the old failure information was cleared away
	objectKey, err := client.ObjectKeyFromObject(vmss)
	tc.Expect(err).ToNot(HaveOccurred())
	updated := &compute.VirtualMachineScaleSet{}
	tc.GetResource(objectKey, updated)

	tc.Expect(updated.Annotations).ToNot(HaveKey(reconcilers.ResourceErrorAnnotation))
}
