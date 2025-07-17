/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20221001"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Insights_Autoscalesetting_v20221001_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVNet20201101(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/8"})
	subnet := newSubnet20201101(tc, vnet, "10.0.0.0/24")
	publicIP := newPublicIP20201101(tc, testcommon.AsOwner(rg))
	lb := newLoadBalancerForVMSS(tc, rg, publicIP)

	tc.CreateResourcesAndWait(vnet, subnet, publicIP, lb)

	vmss := newVMSS20220301(tc, rg, lb, subnet)
	tc.CreateResourceAndWait(vmss)

	objectMeta := tc.MakeObjectMeta("setting")
	setting := &insights.AutoscaleSetting{
		ObjectMeta: objectMeta,
		Spec: insights.AutoscaleSetting_Spec{
			Location:                   tc.AzureRegion,
			Name:                       &objectMeta.Name,
			Owner:                      testcommon.AsOwner(rg),
			TargetResourceUriReference: tc.MakeReferenceFromResource(vmss),
			Profiles: []insights.AutoscaleProfile{
				{
					Capacity: &insights.ScaleCapacity{
						Default: to.Ptr("1"),
						Maximum: to.Ptr("3"),
						Minimum: to.Ptr("1"),
					},
					Name: to.Ptr("profile"),
					Rules: []insights.ScaleRule{
						{
							MetricTrigger: &insights.MetricTrigger{
								MetricName:                 to.Ptr("Percentage CPU"),
								MetricResourceUriReference: tc.MakeReferenceFromResource(vmss),
								Operator:                   to.Ptr(insights.MetricTrigger_Operator_GreaterThan),
								Statistic:                  to.Ptr(insights.MetricTrigger_Statistic_Average),
								Threshold:                  to.Ptr(75.0),
								TimeGrain:                  to.Ptr("PT1M"),
								TimeWindow:                 to.Ptr("PT5M"),
								TimeAggregation:            to.Ptr(insights.MetricTrigger_TimeAggregation_Average),
							},
							ScaleAction: &insights.ScaleAction{
								Cooldown:  to.Ptr("PT1M"),
								Direction: to.Ptr(insights.ScaleAction_Direction_Increase),
								Type:      to.Ptr(insights.ScaleAction_Type_ChangeCount),
								Value:     to.Ptr("1"),
							},
						},
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(setting)

	tc.Expect(setting.Status.Id).ToNot(BeNil())
	armId := *setting.Status.Id

	old := setting.DeepCopy()
	key := "foo"
	setting.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, setting)
	tc.Expect(setting.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(setting)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
