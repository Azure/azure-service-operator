/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	storage "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.storage/v1alpha1api20190401"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_StorageAccount_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateNewTestResourceGroupAndWait()

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
			// TODO: They mark this property as optional but actually it is required
			Properties: &storage.StorageAccountPropertiesCreateParameters{
				AccessTier: &accessTier,
			},
		},
	}

	tc.CreateAndWait(acct)

	tc.Expect(acct.Status.Location).To(Equal(tc.AzureRegion))
	expectedKind := storage.StorageAccountStatusKindBlobStorage
	tc.Expect(acct.Status.Kind).To(Equal(&expectedKind))
	tc.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	// Run sub-tests on storage account
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Blob Services CRUD",
			Test: func(testContext testcommon.KubePerTestContext) {
				StorageAccount_BlobServices_CRUD(testContext, acct.ObjectMeta)
			},
		},
	)

	tc.DeleteAndWait(acct)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadResource(
		tc.Ctx,
		armId,
		"2019-04-01")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func StorageAccount_BlobServices_CRUD(testContext testcommon.KubePerTestContext, storageAccount metav1.ObjectMeta) {
	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: testContext.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccountsBlobServices_Spec{
			Owner: testcommon.AsOwner(storageAccount),
		},
	}

	testContext.CreateAndWait(blobService)
	// no DELETE, this is not a real resource

	testContext.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Container CRUD",
			Test: func(testContext testcommon.KubePerTestContext) {
				StorageAccount_BlobServices_Container_CRUD(testContext, blobService.ObjectMeta)
			},
		})

	// TODO: Delete doesn't seem to work?
	// — is this because it is not a real resource but properties on the storage account?
	// We probably need to ask the Storage team.
	/*
		err = testContext.KubeClient.Delete(ctx, blobService)
		g.Expect(err).ToNot(HaveOccurred())
		g.Eventually(blobService).Should(testContext.Match.BeDeleted())
	*/
}

func StorageAccount_BlobServices_Container_CRUD(testContext testcommon.KubePerTestContext, blobService metav1.ObjectMeta) {

	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: testContext.MakeObjectMeta("container"),
		Spec: storage.StorageAccountsBlobServicesContainers_Spec{
			Owner: testcommon.AsOwner(blobService),
		},
	}

	testContext.CreateAndWait(blobContainer)
	defer testContext.DeleteAndWait(blobContainer)
}
