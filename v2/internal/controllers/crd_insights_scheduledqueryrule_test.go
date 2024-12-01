/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20220615"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Insights_ScheduledQueryRule_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	component := newAppInsightsComponent(tc, rg)
	tc.CreateResourceAndWait(component)

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
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
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
