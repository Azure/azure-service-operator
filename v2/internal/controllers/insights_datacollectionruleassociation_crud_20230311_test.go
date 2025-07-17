/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20230311"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Insights_DataCollectionRuleAssociation_v20230311_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a storage account to use as the parent resource for the association
	sa := newStorageAccount(tc, rg)

	// Create a data collection endpoint to associate with the storage account
	endpoint := &insights.DataCollectionEndpoint{
		ObjectMeta: tc.MakeObjectMeta("dce"),
		Spec: insights.DataCollectionEndpoint_Spec{
			Description: to.Ptr("Data collection endpoint for association testing"),
			Kind:        to.Ptr(insights.DataCollectionEndpoint_Kind_Spec_Linux),
			Location:    tc.AzureRegion,
			Owner:       testcommon.AsOwner(rg),
		},
	}

	// Create the data collection rule association
	association := &insights.DataCollectionRuleAssociation{
		ObjectMeta: tc.MakeObjectMetaWithName("configuration-access-endpoint"),
		Spec: insights.DataCollectionRuleAssociation_Spec{
			AzureName:                       "configurationAccessEndpoint",
			DataCollectionEndpointReference: tc.MakeReferenceFromResource(endpoint),
			Description:                     to.Ptr("Data collection rule association for testing"),
			Owner:                           tc.AsExtensionOwner(sa),
		},
	}

	tc.CreateResourcesAndWait(sa, endpoint, association)

	// Verify status fields are populated correctly
	tc.Expect(association.Status.Id).ToNot(BeNil())
	tc.Expect(association.Status.Name).ToNot(BeNil())
	tc.Expect(association.Status.Type).ToNot(BeNil())

	armId := *association.Status.Id

	// Test update - associations don't typically support tags, but we can update the description
	old := association.DeepCopy()
	association.Spec.Description = to.Ptr("Updated data collection rule association description")

	tc.PatchResourceAndWait(old, association)
	tc.Expect(association.Status.Description).To(Equal(association.Spec.Description))

	tc.DeleteResourceAndWait(association)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
