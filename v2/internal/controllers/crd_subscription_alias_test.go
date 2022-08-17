/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801preview"
	subscription "github.com/Azure/azure-service-operator/v2/api/subscription/v1beta20211001"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// NOTE: Recording this test is a bit of a pain if you're in the MSFT tenant.
// See https://dev.azure.com/msazure/CloudNativeCompute/_wiki/wikis/CloudNativeCompute.wiki/357302/Recording-Subscription-Creation-test
// for how this should be done.
// If not in the MSFT tenant, just set TEST_BILLING_ID to a valid Billing Invoice ID and record the test as normal

func Test_SubscriptionAndAlias_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	if tc.AzureBillingInvoiceID == "" {
		t.Fatalf("%q enviornment variable must be set", testcommon.TestBillingIDVar)
	}

	if *isLive == true {
		t.Skip("Can't run in live mode as our tenant doesn't support subscription creation")
	}

	// TODO: Once ManagedIdentity is registered in our sub, we can use this instead of the hardcoded identity below
	// TODO: There's an issue registering it right now though so for now using a hardcoded identity
	//rg := tc.CreateTestResourceGroupAndWait()
	//// Create a dummy managed identity which we will assign to a role
	//mi := &managedidentity.UserAssignedIdentity{
	//	ObjectMeta: tc.MakeObjectMeta("mi"),
	//	Spec: managedidentity.UserAssignedIdentities_Spec{
	//		Location: tc.AzureRegion,
	//		Owner:    testcommon.AsOwner(rg),
	//	},
	//}
	//
	//tc.CreateResourceAndWait(mi)
	//tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	//tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	workload := subscription.PutAliasRequestPropertiesWorkload_Production
	sub := &subscription.Alias{
		ObjectMeta: tc.MakeObjectMeta("sub"),
		Spec: subscription.Aliases_Spec{
			Properties: &subscription.PutAliasRequestProperties{
				DisplayName:  to.StringPtr("Subscription for ASO testing"),
				Workload:     &workload,
				BillingScope: &tc.AzureBillingInvoiceID,
			},
		},
	}

	tc.CreateResourceAndWait(sub)
	tc.Expect(sub.Status.Id).ToNot(BeNil())
	armId := *sub.Status.Id

	// Assign an identity to a new role
	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())
	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(roleAssignmentGUID.String()),
		Spec: authorization.RoleAssignments_Spec{
			Location: tc.AzureRegion,
			Owner:    tc.AsExtensionOwner(sub),
			//PrincipalId: mi.Status.PrincipalId,
			PrincipalId: to.StringPtr("1605884e-16c3-4fc0-bf09-4220deecef02"), /// 1605884e-16c3-4fc0-bf09-4220deecef02 == my user in PPE
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: "/providers/Microsoft.Authorization/roleDefinitions/8e3af657-a8ff-443c-a75c-2fe8c4bcb635", // This is owner
			},
		},
	}

	tc.CreateResourceAndWait(roleAssignment)

	tc.DeleteResourceAndWait(sub)
	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(subscription.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())

}
