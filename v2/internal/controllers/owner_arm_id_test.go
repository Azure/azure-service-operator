/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20200801preview"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_OwnerIsARMIDOfResourceGroup_ResourceSuccessfullyReconciled(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// Get the rg's ARM ID
	tc.Expect(rg.Status.Id).ToNot(BeNil())
	armID := *rg.Status.Id

	// Now create a storage account
	acct := newStorageAccount20230101(tc, rg)
	acct.Spec.Owner = testcommon.AsARMIDOwner(armID)

	// Create the storage account from ARM ID
	tc.CreateResourceAndWait(acct)
	tc.Expect(acct.Status.Id).ToNot(BeNil())
	acctARMID := *acct.Status.Id

	// Update the storage account
	old := acct.DeepCopy()
	acct.Spec.Tags = map[string]string{"tag1": "value1"}
	tc.PatchResourceAndWait(old, acct)
	tc.Expect(acct.Status.Tags).To(HaveKey("tag1"))

	tc.DeleteResourceAndWait(acct)

	// Ensure that the account was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		acctARMID,
		string(storage.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Test_OwnerIsARMIDOfParent_ChildResourceSuccessfullyReconciled(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// Now create a storage account
	acct := newStorageAccount20230101(tc, rg)
	tc.CreateResourceAndWait(acct)

	// and a blob service
	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccountsBlobService_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}
	tc.CreateResourceAndWait(blobService)

	tc.Expect(blobService.Status.Id).ToNot(BeNil())
	armID := *blobService.Status.Id

	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMeta("container"),
		Spec: storage.StorageAccountsBlobServicesContainer_Spec{
			Owner: testcommon.AsARMIDOwner(armID),
		},
	}
	tc.CreateResourceAndWait(blobContainer)
	tc.Expect(blobContainer.Status.Id).ToNot(BeNil())
	containerARMID := *blobContainer.Status.Id

	// Update the container
	old := blobContainer.DeepCopy()
	blobContainer.Spec.Metadata = map[string]string{"tag1": "value1"}
	tc.PatchResourceAndWait(old, blobContainer)
	tc.Expect(blobContainer.Status.Metadata).To(HaveKey("tag1"))

	tc.DeleteResourceAndWait(blobContainer)

	// Ensure that the container was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		containerARMID,
		string(storage.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Test_OwnerIsARMID_ExtensionResourceSuccessfullyReconciled(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
	principalIdKey := "principalId"

	// Create a dummy managed identity which we will assign to a role
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  principalIdKey,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(mi)
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())
	tc.Expect(mi.Status.Id).ToNot(BeNil())
	armID := *mi.Status.Id

	// Now assign a new role to that  managed identity
	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMeta("roleassignment"),
		Spec: authorization.RoleAssignment_Spec{
			Owner: &genruntime.ArbitraryOwnerReference{
				ARMID: armID,
			},
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is contributor
			},
		},
	}
	tc.CreateResourceAndWait(roleAssignment)

	tc.Expect(roleAssignment.Status.Id).ToNot(BeNil())
	armId := *roleAssignment.Status.Id

	tc.DeleteResourceAndWait(roleAssignment)

	// Ensure that the role assignment was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(authorization.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
