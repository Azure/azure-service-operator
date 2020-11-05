/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	storage "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.storage/v20190401"
	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
	"github.com/Azure/k8s-infra/hack/generated/pkg/testcommon"
)

func Test_ResourceGroup_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	// Create a resource group
	rg := testContext.NewTestResourceGroup()
	err = testContext.KubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	// It should be created in Kubernetes
	g.Eventually(rg).Should(testContext.Match.BeProvisioned(ctx))

	g.Expect(rg.Status.Location).To(Equal(testContext.AzureRegion))
	g.Expect(rg.Status.Properties.ProvisioningState).To(Equal(string(armclient.SucceededProvisioningState)))
	g.Expect(rg.Status.ID).ToNot(BeNil())
	armId := rg.Status.ID

	// Delete the resource group
	err = testContext.KubeClient.Delete(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())
	g.Eventually(rg).Should(testContext.Match.BeDeleted(ctx))

	// Ensure that the resource group was really deleted in Azure
	// TODO: Do we want to just use an SDK here? This process is quite icky as is...
	exists, err := testContext.AzureClient.HeadResource(
		ctx,
		armId,
		"2020-06-01")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(exists).To(BeFalse())
}

func Test_StorageAccount_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	rg, err := testContext.CreateNewTestResourceGroup(testcommon.WaitForCreation)
	g.Expect(err).ToNot(HaveOccurred())

	// Custom namer because storage accounts have strict names
	namer := testContext.Namer.WithSeparator("")

	// Create a storage account
	accessTier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	acct := &storage.StorageAccount{
		ObjectMeta: testContext.MakeObjectMetaWithName(namer.GenerateName("stor")),
		Spec: storage.StorageAccounts_Spec{
			Location:   testContext.AzureRegion,
			ApiVersion: "2019-04-01", // TODO: This should be removed from the storage type eventually
			Owner:      testcommon.AsOwner(rg.ObjectMeta),
			Kind:       storage.StorageAccountsSpecKindBlobStorage,
			Sku: storage.Sku{
				Name: storage.SkuNameStandardLRS,
			},
			// TODO: They mark this property as optional but actually it is required
			Properties: &storage.StorageAccountPropertiesCreateParameters{
				AccessTier: &accessTier,
			},
		},
	}
	err = testContext.KubeClient.Create(ctx, acct)
	g.Expect(err).ToNot(HaveOccurred())

	// It should be created in Kubernetes
	g.Eventually(acct).Should(testContext.Match.BeProvisioned(ctx))

	g.Expect(acct.Status.Location).To(Equal(testContext.AzureRegion))
	expectedKind := storage.StorageAccountStatusKindBlobStorage
	g.Expect(acct.Status.Kind).To(Equal(&expectedKind))
	g.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	// Delete the storage account
	err = testContext.KubeClient.Delete(ctx, acct)
	g.Expect(err).ToNot(HaveOccurred())
	g.Eventually(acct).Should(testContext.Match.BeDeleted(ctx))

	// Ensure that the resource group was really deleted in Azure
	// TODO: Do we want to just use an SDK here? This process is quite icky as is...
	exists, err := testContext.AzureClient.HeadResource(
		ctx,
		armId,
		"2019-04-01")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(exists).To(BeFalse())
}
