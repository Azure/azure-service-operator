/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/api/authorization/customizations"
	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v20201001"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Authorization_RoleEligibilityScheduleRequest_20201001_CRUD(t *testing.T) {
	t.Parallel()

	customizations.DisableBuiltInRoleDefinitionsCaching()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "identity-settings"
	principalIdKey := "principalId"
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

	request := &authorization.RoleEligibilityScheduleRequest{
		ObjectMeta: tc.MakeObjectMeta("eligible"),
		Spec: authorization.RoleEligibilityScheduleRequest_Spec{
			Owner: tc.AsExtensionOwner(rg),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.WellKnownResourceReference{
				WellKnownName: "Reader",
			},
			RequestType: to.Ptr(authorization.RoleEligibilityScheduleRequestProperties_RequestType_AdminAssign),
			ScheduleInfo: &authorization.RoleEligibilityScheduleRequestProperties_ScheduleInfo{
				Expiration: &authorization.RoleEligibilityScheduleRequestProperties_ScheduleInfo_Expiration{
					Duration: to.Ptr("P30D"),
					Type: to.Ptr(
						authorization.RoleEligibilityScheduleRequestProperties_ScheduleInfo_Expiration_Type_AfterDuration,
					),
				},
			},
		},
	}
	tc.AddAnnotation(&request.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourcesAndWait(mi, request)

	tc.Expect(request.Status.Id).ToNot(BeNil())
	tc.Expect(request.Status.TargetRoleEligibilityScheduleId).ToNot(BeNil())

	removeRequest := &authorization.RoleEligibilityScheduleRequest{
		ObjectMeta: tc.MakeObjectMeta("remove"),
		Spec: authorization.RoleEligibilityScheduleRequest_Spec{
			Owner: tc.AsExtensionOwner(rg),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.WellKnownResourceReference{
				WellKnownName: "Reader",
			},
			RequestType:                     to.Ptr(authorization.RoleEligibilityScheduleRequestProperties_RequestType_AdminRemove),
			TargetRoleEligibilityScheduleId: request.Status.TargetRoleEligibilityScheduleId,
		},
	}
	tc.AddAnnotation(&removeRequest.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourceAndWait(removeRequest)
	tc.Expect(removeRequest.Status.Id).ToNot(BeNil())

	tc.DeleteResourcesAndWait(removeRequest, request)
}
