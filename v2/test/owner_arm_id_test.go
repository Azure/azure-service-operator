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

	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// This test cannot be run in record/replay mode because the state it looks for at the beginning (Ready = false with warning)
// is not "stable" (the reconciler keeps retrying). Since it keeps retrying there isn't a deterministic number of
// retry attempts it makes which means a recording test may run out of recorded retries.
func Test_OwnerIsARMIDThatDoesntExist_ResourceFailsWithWarning(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	dummyResourceGroupID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", tc.AzureSubscription, "nonexistrg")

	// Now create a storage account
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsARMIDOwner(dummyResourceGroupID),
			Kind:     to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
			Sku: &storage.Sku{
				Name: to.Ptr(storage.SkuName_Standard_LRS),
			},
			AccessTier: to.Ptr(storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot),
		},
	}

	// This isn't a fatal state because fatal errors won't reconcile periodically and if the resource is created in Azure
	// we want to proceed
	tc.CreateResourceAndWaitForState(acct, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	tc.Expect(acct.Status.Conditions[0].Message).To(ContainSubstring("Resource group 'nonexistrg' could not be found."))

	// Ensure that delete of the acct goes through properly
	tc.DeleteResourceAndWait(acct)
}

func Test_OwnerIsARMID_OwnerDeleted_ResourceFailsWithWarning(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// Get the rg's ARM ID
	tc.Expect(rg.Status.Id).ToNot(BeNil())
	armID := *rg.Status.Id

	// Now create a storage account
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsARMIDOwner(armID),
			Kind:     to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
			Sku: &storage.Sku{
				Name: to.Ptr(storage.SkuName_Standard_LRS),
			},
			AccessTier: to.Ptr(storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot),
		},
	}

	// Create the storage account from ARM ID
	tc.CreateResourceAndWait(acct)

	// Delete the resource group
	tc.DeleteResourceAndWait(rg)

	// Now try to update the storage acct
	old := acct.DeepCopy()
	acct.Spec.Tags = map[string]string{"tag1": "value1"}
	tc.PatchResourceAndWaitForState(old, acct, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	tc.Expect(acct.Status.Conditions[0].Message).To(ContainSubstring(fmt.Sprintf("Resource group '%s' could not be found.", rg.Name)))

	// Ensure that delete of the acct goes through properly
	tc.DeleteResourceAndWait(acct)
}
