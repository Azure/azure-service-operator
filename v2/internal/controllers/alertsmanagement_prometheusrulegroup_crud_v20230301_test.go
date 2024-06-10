/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	alertsmanagement "github.com/Azure/azure-service-operator/v2/api/alertsmanagement/v1api20230301"
	monitor "github.com/Azure/azure-service-operator/v2/api/monitor/v1api20230403"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AlertsManagement_PrometheusRuleGroup_CRUD(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	acct := &monitor.Account{
		ObjectMeta: tc.MakeObjectMeta("acct"),
		Spec: monitor.Account_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	// Create a Prometheus Rule Group
	// This was adapted from https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/prometheus-rule-groups#template-example-for-a-prometheus-rule-group
	ruleGroup := &alertsmanagement.PrometheusRuleGroup{
		ObjectMeta: tc.MakeObjectMeta("promrule"),
		Spec: alertsmanagement.PrometheusRuleGroup_Spec{
			Location:    tc.AzureRegion,
			Owner:       testcommon.AsOwner(rg),
			Enabled:     to.Ptr(true),
			ClusterName: to.Ptr("mycluster"),
			Interval:    to.Ptr("PT1M"),
			ScopesReferences: []genruntime.ResourceReference{
				*tc.MakeReferenceFromResource(acct),
			},
			Rules: []alertsmanagement.PrometheusRule{
				{
					Enabled:    to.Ptr(true),
					Expression: to.Ptr("1 - avg without (cpu) (sum without (mode)(rate(node_cpu_seconds_total{job=\"node\", mode=~\"idle|iowait|steal\"}[5m])))"),
					Labels: map[string]string{
						"workload_type": "job",
					},
					Record: to.Ptr("instance:node_cpu_utilisation:rate5m"),
				},
			},
		},
	}
	tc.CreateResourcesAndWait(ruleGroup, acct)

	// Ensure that the status is what we expect
	tc.Expect(ruleGroup.Status.Id).ToNot(BeNil())
	armId := *ruleGroup.Status.Id

	tc.DeleteResourceAndWait(ruleGroup)

	// Ensure that the resource was really deleted in Azure
	ctx := context.Background()
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(ctx, armId, string(alertsmanagement.APIVersion_Value))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
