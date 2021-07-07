/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	network "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.network/v1alpha1api20201101"
	resources "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	storage "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.storage/v1alpha1api20210401"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/reconcilers"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func waitForOwnerMissingError(tc testcommon.KubePerTestContext, obj controllerutil.Object) {
	objectKey, err := client.ObjectKeyFromObject(obj)
	tc.Expect(err).ToNot(HaveOccurred())

	tc.Eventually(func() string {
		tc.GetResource(objectKey, obj)
		return obj.GetAnnotations()[reconcilers.ResourceErrorAnnotation]
	}).Should(MatchRegexp("owner.*is not ready"))
}

func doNotWait(_ testcommon.KubePerTestContext, _ controllerutil.Object) {}

func storageAccountAndResourceGroupProvisionedOutOfOrderHelper(t *testing.T, waitHelper func(tc testcommon.KubePerTestContext, obj controllerutil.Object)) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create the resource group in-memory but don't submit it yet
	rg := tc.NewTestResourceGroup()

	// Custom namer because storage accounts have strict names
	namer := tc.Namer.WithSeparator("")

	// Create a storage account
	accessTier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("stor")),
		Spec: storage.StorageAccounts_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Kind:     storage.StorageAccountsSpecKindBlobStorage,
			Sku: storage.Sku{
				Name: storage.SkuNameStandardLRS,
			},
			AccessTier: &accessTier,
		},
	}

	// Create the storage account - initially this will not succeed, but it should keep trying
	tc.G.Expect(tc.KubeClient.Create(tc.Ctx, acct)).To(Succeed())

	waitHelper(tc, acct)

	// The resource group should be created successfully
	_, err := tc.CreateTestResourceGroup(rg, testcommon.WaitForCreation)
	tc.Expect(err).ToNot(HaveOccurred())

	// The storage account should also be created successfully
	tc.G.Eventually(acct, tc.RemainingTime()).Should(tc.Match.BeProvisioned())
}

func subnetAndVNETCreatedProvisionedOutOfOrder(t *testing.T, waitHelper func(tc testcommon.KubePerTestContext, obj controllerutil.Object)) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateNewTestResourceGroupAndWait()

	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetworks_Spec{
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Location: testcommon.DefaultTestRegion,
			AddressSpace: network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}

	subnet := &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworksSubnets_Spec{
			Owner:         testcommon.AsOwner(vnet.ObjectMeta),
			AddressPrefix: "10.0.0.0/24",
		},
	}

	// Create the subnet - initially this will not succeed, but it should keep trying
	tc.G.Expect(tc.KubeClient.Create(tc.Ctx, subnet)).To(Succeed())

	waitHelper(tc, subnet)

	// Now created the vnet
	tc.CreateResourceAndWait(vnet)
	// The subnet account should also be created successfully eventually
	tc.G.Eventually(subnet, tc.RemainingTime()).Should(tc.Match.BeProvisioned())
}

func Test_StorageAccount_CreatedBeforeResourceGroup(t *testing.T) {
	storageAccountAndResourceGroupProvisionedOutOfOrderHelper(t, waitForOwnerMissingError)
}

func Test_StorageAccount_CreatedInParallelWithResourceGroup(t *testing.T) {
	storageAccountAndResourceGroupProvisionedOutOfOrderHelper(t, doNotWait)
}

func Test_Subnet_CreatedBeforeVNET(t *testing.T) {
	subnetAndVNETCreatedProvisionedOutOfOrder(t, waitForOwnerMissingError)
}

func Test_Subnet_CreatedInParallelWithVNET(t *testing.T) {
	subnetAndVNETCreatedProvisionedOutOfOrder(t, doNotWait)
}

func Test_CreateResourceGroupThatAlreadyExists_ReconcilesSuccessfully(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateNewTestResourceGroupAndWait()

	// Create another resource group that points to the same Azure resource
	rgCopy := &resources.ResourceGroup{
		ObjectMeta: ctrl.ObjectMeta{
			Name:      rg.Name,
			Namespace: rg.Namespace,
		},
		Spec: *rg.Spec.DeepCopy(),
	}
	rgCopy.Spec.AzureName = rgCopy.Name
	rgCopy.Name = rgCopy.Name + "duplicate"

	tc.CreateResourceAndWait(rgCopy)
}

func Test_CreateStorageAccountThatAlreadyExists_ReconcilesSuccessfully(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateNewTestResourceGroupAndWait()

	// Create another resource group that points to the same Azure resource
	// Custom namer because storage accounts have strict names
	namer := tc.Namer.WithSeparator("")

	// Create a storage account
	accessTier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("stor")),
		Spec: storage.StorageAccounts_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Kind:     storage.StorageAccountsSpecKindBlobStorage,
			Sku: storage.Sku{
				Name: storage.SkuNameStandardLRS,
			},
			AccessTier: &accessTier,
		},
	}

	acctCopy := acct.DeepCopy()

	tc.CreateResourcesAndWait(acct)

	// Patch the account to remove the finalizer
	patcher := tc.NewResourcePatcher(acct)
	controllerutil.RemoveFinalizer(acct, "generated.infra.azure.com/finalizer")
	patcher.Patch(acct)

	// Delete the account
	tc.DeleteResourceAndWait(acct)

	// Create it again
	tc.CreateResourcesAndWait(acctCopy)
}
