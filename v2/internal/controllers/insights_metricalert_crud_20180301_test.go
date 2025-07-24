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
	tc.CreateResourceAndWait(sa)

	alert := &insights.MetricAlert{
		ObjectMeta: tc.MakeObjectMeta("alert"),
		Spec: insights.MetricAlert_Spec{
			AzureName: "",
			Criteria: &insights.MetricAlertCriteria{
				MicrosoftAzureMonitorSingleResourceMultipleMetric: &insights.MetricAlertSingleResourceMultipleMetricCriteria{
					AllOf: []insights.MetricCriteria{
						{
							CriterionType:        to.Ptr(insights.MetricCriteria_CriterionType_StaticThresholdCriterion),
							MetricName:           to.Ptr("TestMetric"),
							MetricNamespace:      to.Ptr("namespace"),
							Name:                 to.Ptr("Criteria1"),
							Operator:             to.Ptr(insights.MetricCriteria_Operator_GreaterThan),
							SkipMetricValidation: to.Ptr(true),
							Threshold:            to.Ptr(10.0),
							TimeAggregation:      to.Ptr(insights.MetricCriteria_TimeAggregation_Count),
						},
					},
					OdataType: to.Ptr(insights.MetricAlertSingleResourceMultipleMetricCriteria_OdataType_MicrosoftAzureMonitorSingleResourceMultipleMetricCriteria),
				},
			},
			Enabled:             to.Ptr(false),
			EvaluationFrequency: to.Ptr("PT5M"),
			Location:            to.Ptr("global"),
			Owner:               testcommon.AsOwner(rg),
			ScopesReferences: []genruntime.ResourceReference{
				*tc.MakeReferenceFromResource(sa),
			},
			Severity:   to.Ptr(0),
			WindowSize: to.Ptr("PT5M"),
		},
	}

	tc.CreateResourceAndWait(alert)

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
