/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20180301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Insights_MetricAlert_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	sa := newStorageAccount(tc, rg)

	alert := &insights.MetricAlert{
		ObjectMeta: tc.MakeObjectMeta("alert"),
		Spec: insights.MetricAlert_Spec{
			AzureName: "",
			Criteria: &insights.MetricAlertCriteria{
				MicrosoftAzureMonitorSingleResourceMultipleMetric: &insights.MetricAlertSingleResourceMultipleMetricCriteria{
					AllOf: []insights.MetricCriteria{
						{
							CriterionType:        to.Ptr(insights.MetricCriteria_CriterionType_StaticThresholdCriterion),
							MetricName:           new("TestMetric"),
							MetricNamespace:      new("namespace"),
							Name:                 new("Criteria1"),
							Operator:             to.Ptr(insights.MetricCriteria_Operator_GreaterThan),
							SkipMetricValidation: new(true),
							Threshold:            new(10.0),
							TimeAggregation:      to.Ptr(insights.MetricCriteria_TimeAggregation_Count),
						},
					},
					OdataType: to.Ptr(insights.MetricAlertSingleResourceMultipleMetricCriteria_OdataType_MicrosoftAzureMonitorSingleResourceMultipleMetricCriteria),
				},
			},
			Enabled:             new(false),
			EvaluationFrequency: new("PT5M"),
			Location:            new("global"),
			Owner:               testcommon.AsOwner(rg),
			ScopesReferences: []genruntime.ResourceReference{
				*tc.MakeReferenceFromResource(sa),
			},
			Severity:   new(0),
			WindowSize: new("PT5M"),
		},
	}

	tc.CreateResourcesAndWait(sa, alert)

	tc.Expect(alert.Status.Id).ToNot(BeNil())
	armId := *alert.Status.Id

	old := alert.DeepCopy()
	key := "foo"
	alert.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, alert)
	tc.Expect(alert.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(alert)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
