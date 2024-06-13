/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// See https://learn.microsoft.com/en-us/azure/role-based-access-control/custom-roles-template and
// https://learn.microsoft.com/en-us/azure/role-based-access-control/custom-roles-rest#create-a-custom-role
func Test_Authorization_RoleDefinitionSubscriptionScope_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	subARMID := fmt.Sprintf("/subscriptions/%s", tc.AzureSubscription)

	// We have to manually set AzureName to a stable UUID because generating it based on the
	// owner.armId means that the seed to the random UUID changes between replay and recording.
	// In recording mode, it will be /subscriptions/<real-sub-id> and in replay it will be
	// /subscriptions/<all-0s-uuid>
	// Unlike other UUID replacements we can't just fix this up with the recorder in the replay
	// because the ID is used as a seed to generate a UUID rather than just going into the HTTP
	// recording directly.
	azureName, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	om := tc.MakeObjectMeta("roledef")
	roleDef := &authorization.RoleDefinition{
		ObjectMeta: om,
		Spec: authorization.RoleDefinition_Spec{
			Owner: &genruntime.ArbitraryOwnerReference{
				ARMID: subARMID,
			},
			Type:      to.Ptr("customRole"),
			AzureName: azureName.String(),
			AssignableScopesReferences: []genruntime.ResourceReference{
				{
					ARMID: subARMID,
				},
			},
			RoleName: to.Ptr(om.Name),
			Permissions: []authorization.Permission{
				{
					Actions: []string{
						"Microsoft.Resources/subscriptions/resourceGroups/read",
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(roleDef)

	tc.Expect(roleDef.Status.Id).ToNot(BeNil())
	armId := *roleDef.Status.Id

	tc.DeleteResourceAndWait(roleDef)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(authorization.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Test_Authorization_RoleDefinitionResourceGroupScope_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.NewTestResourceGroup()

	om := tc.MakeObjectMeta("roledef")
	roleDef := &authorization.RoleDefinition{
		ObjectMeta: om,
		Spec: authorization.RoleDefinition_Spec{
			Owner: tc.AsExtensionOwner(rg),
			Type:  to.Ptr("customRole"),
			AssignableScopesReferences: []genruntime.ResourceReference{
				*tc.MakeReferenceFromResource(rg),
			},
			RoleName: to.Ptr(om.Name),
			Permissions: []authorization.Permission{
				{
					Actions: []string{
						"Microsoft.Resources/subscriptions/resourceGroups/read",
					},
				},
			},
		},
	}

	tc.CreateResourcesAndWait(rg, roleDef)

	tc.Expect(roleDef.Status.Id).ToNot(BeNil())
	armId := *roleDef.Status.Id

	tc.DeleteResourceAndWait(roleDef)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(authorization.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
