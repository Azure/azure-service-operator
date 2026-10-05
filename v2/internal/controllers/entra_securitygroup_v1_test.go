/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	entra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Entra_SecurityGroup_v1_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: "mi-secret",
						Key:  "principal-id",
					},
				},
			},
		},
	}

	securityGroup := &entra.SecurityGroup{
		ObjectMeta: tc.MakeObjectMeta("securitygroup"),
		Spec: entra.SecurityGroupSpec{
			DisplayName:    to.Ptr("ASO Test Security Group"),
			MailNickname:   to.Ptr("asotest-security-group"),
			MembershipType: to.Ptr(entra.SecurityGroupMembershipTypeAssigned),
			OperatorSpec: &entra.SecurityGroupOperatorSpec{
				ConfigMaps: &entra.SecurityGroupOperatorConfigMaps{
					EntraID: &genruntime.ConfigMapDestination{
						Name: "entra-secret",
						Key:  "entra-id",
					},
				},
			},
			Owners: []entra.SecurityGroupMemberReference{
				{
					ObjectIDFromConfig: &genruntime.ConfigMapReference{
						Name: "mi-secret",
						Key:  "principal-id",
					},
				},
			},
		},
	}

	// Create the resources and wait for them to be ready
	tc.CreateResourcesAndWait(mi, securityGroup)

	// Make sure the configmaps were correctly created
	configMapList := &v1.ConfigMapList{}
	tc.ListResources(configMapList, client.InNamespace(tc.Namespace))
	tc.Expect(configMapList.Items).To(HaveLen(2))
	tc.Expect(configMapList.Items[0].Data).To(HaveKeyWithValue("entra-id", *securityGroup.Status.EntraID))

	// Save an update
	old := securityGroup.DeepCopy()
	securityGroup.Spec.DisplayName = to.Ptr("ASO Test Security Group Updated")
	tc.PatchResourceAndWait(old, securityGroup)

	// Make sure the group was updated
	tc.Expect(*securityGroup.Status.DisplayName).To(Equal("ASO Test Security Group Updated"))

	// Delete the resource and wait for it to be gone
	tc.DeleteResourceAndWait(securityGroup)
}
