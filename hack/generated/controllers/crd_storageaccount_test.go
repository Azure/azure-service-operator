/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	storage "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.storage/v1alpha1api20210401"
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
			AccessTier: &accessTier,
		},
	}

	tc.CreateResourceAndWait(acct)

	tc.Expect(acct.Status.Location).To(Equal(tc.AzureRegion))
	expectedKind := storage.StorageAccountStatusKindBlobStorage
	tc.Expect(acct.Status.Kind).To(Equal(&expectedKind))
	tc.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	// Run sub-tests on storage account
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Blob Services CRUD",
			Test: func(tc testcommon.KubePerTestContext) {
				StorageAccount_BlobServices_CRUD(tc, acct.ObjectMeta)
			},
		},
	)

	tc.DeleteResourceAndWait(acct)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadResource(
		tc.Ctx,
		armId,
		string(storage.StorageAccountsSpecAPIVersion20210401))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func StorageAccount_BlobServices_CRUD(tc testcommon.KubePerTestContext, storageAccount metav1.ObjectMeta) {
	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccountsBlobServices_Spec{
			Owner: testcommon.AsOwner(storageAccount),
		},
	}

	tc.CreateResourceAndWait(blobService)
	// no DELETE, this is not a real resource

	tc.RunParallelSubtests(
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

func StorageAccount_BlobServices_Container_CRUD(tc testcommon.KubePerTestContext, blobService metav1.ObjectMeta) {
	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMeta("container"),
		Spec: storage.StorageAccountsBlobServicesContainers_Spec{
			Owner: testcommon.AsOwner(blobService),
		},
	}

	tc.CreateResourceAndWait(blobContainer)
	defer tc.DeleteResourceAndWait(blobContainer)
}
