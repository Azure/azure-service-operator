/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	// The testing package is imported for testing-related functionality.
	"testing"

	// The gomega package is used for assertions and expectations in tests.
	. "github.com/onsi/gomega"

	// The dataprotection package contains types and functions related to dataprotection resources.
	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	// The testcommon package includes common testing utilities.
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	// The to package includes utilities for converting values to pointers.
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Dataprotection_Backupinstace_CRUD(t *testing.T) {
	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// Create a test resource group and wait until the operation is completed, where the globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// create storage account and blob container
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
			Sku: &storage.Sku{
				Name: to.Ptr(storage.SkuName_Standard_LRS),
			},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: to.Ptr(storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot),
		},
	}
	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccounts_BlobService_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}
	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMeta("velero"),
		Spec: storage.StorageAccounts_BlobServices_Container_Spec{
			Owner: testcommon.AsOwner(blobService),
		},
	}

	tc.CreateResourceAndWait(acct, blobContainer)

	// create cluster

	// create extension

	// give permission to extension msi over SA

	// create vault

	// Create a new backupvault resource
	backupVault := newBackupVault(tc, rg, "asotestbackupvault")

	// give read permission to vault msi over SRG
	// give read permission to vault msi over cluster
	// create TA role binding

	// create policy

	//create backup instance

	// Note:
	// It is mandatory to create a backupvault before creating a backuppolicy

	// Create a BackupPolicy
	backupPolicy := newBackupPolicy(tc, backupVault, "asotestbackuppolicy")

	// Sequence of creating BackupVault and BackupPolicy is handled by ASO internally
	tc.CreateResourcesAndWait(backupVault, backupPolicy)

	// Assertions and Expectations
	armId := *backupPolicy.Status.Id

	// Note:
	// Patch Operations are currently not allowed on BackupInstance

	// Delete the backupinstance

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(dataprotection.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
