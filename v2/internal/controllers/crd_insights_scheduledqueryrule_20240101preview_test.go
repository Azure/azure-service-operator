/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"
	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20240101preview"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Insights_ScheduledQueryRule_20240101preview_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	tc.AzureRegion = to.Ptr("eastus")

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
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
	tc.CreateResourcesAndWait(mi)

	component := newAppInsightsComponent(tc, rg)
	tc.CreateResourceAndWait(component)

	armID := *component.Status.Id

	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())
	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(roleAssignmentGUID.String()),
		Spec: authorization.RoleAssignment_Spec{
			Owner: &genruntime.ArbitraryOwnerReference{
				ARMID: armID,
			},
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: "/providers/Microsoft.Authorization/roleDefinitions/8e3af657-a8ff-443c-a75c-2fe8c4bcb635", // This is owner
			},
		},
	}
	tc.CreateResourceAndWait(roleAssignment)

	tc.Expect(roleAssignment.Status.Id).ToNot(BeNil())

	rule := &insights.ScheduledQueryRule{
		ObjectMeta: tc.MakeObjectMeta("rule"),
		Spec: insights.ScheduledQueryRule_Spec{
			Criteria: &insights.ScheduledQueryRuleCriteria{
				AllOf: []insights.Condition{
					{
						FailingPeriods: &insights.Condition_FailingPeriods{
							MinFailingPeriodsToAlert:  to.Ptr(1),
							NumberOfEvaluationPeriods: to.Ptr(1),
						},
						Operator:                  to.Ptr(insights.Condition_Operator_LessThan),
						Query:                     to.Ptr("requests | summarize CountByCountry=count() by client_CountryOrRegion"),
						ResourceIdColumnReference: nil,
						Threshold:                 to.Ptr(10.0),
						TimeAggregation:           to.Ptr(insights.Condition_TimeAggregation_Count),
					},
				},
			},
			EvaluationFrequency: to.Ptr("PT10M"),
			Identity: &insights.Identity{
				Type: to.Ptr(insights.Identity_Type_UserAssigned),
				UserAssignedIdentities: []insights.UserAssignedIdentityDetails{
					{
						Reference: *tc.MakeReferenceFromResource(mi),
					},
				},
			},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			ScopesReferences: []genruntime.ResourceReference{
				*tc.MakeReferenceFromResource(component),
			},
			Severity:   to.Ptr(insights.ScheduledQueryRuleProperties_Severity_0),
			WindowSize: to.Ptr("PT10M"),
		},
	}

	tc.CreateResourceAndWait(rule)

	tc.Expect(rule.Status.Id).ToNot(BeNil())
	armId := *rule.Status.Id

	old := rule.DeepCopy()
	key := "foo"
	rule.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, rule)
	tc.Expect(rule.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(rule)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
