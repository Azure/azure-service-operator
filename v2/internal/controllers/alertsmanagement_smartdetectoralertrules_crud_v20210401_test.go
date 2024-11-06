/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	alertsmanagement "github.com/Azure/azure-service-operator/v2/api/alertsmanagement/v1api20210401"
	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20230101"
	monitor "github.com/Azure/azure-service-operator/v2/api/monitor/v1api20230403"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AlertsManagement_SmartDetectorAlertRules_CRUD(t *testing.T) {
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
	ag := &insights.ActionGroup{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("actiongroup")),
		Spec: insights.ActionGroup_Spec{
			Enabled:        to.Ptr(false),
			GroupShortName: to.Ptr("ag"),
			Location:       to.Ptr("global"),
			Owner:          testcommon.AsOwner(rg),
		},
	}

	// Create a Smart Detector alert rule
	// This was adapted from https://learn.microsoft.com/en-us/rest/api/monitor/smart-detector-alert-rules/create-or-update?view=rest-monitor-2019-06-01&tabs=HTTP#create-or-update-a-smart-detector-alert-rule
	state := alertsmanagement.AlertRuleProperties_State("Enabled")
	severity := alertsmanagement.AlertRuleProperties_Severity("Sev3")
	detectorId := &alertsmanagement.Detector{
		Reference: &genruntime.ResourceReference{
			ARMID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/detectors/detectorName",
		},
	}
	actionGroup := &alertsmanagement.ActionGroupsInformation{
		GroupReferences: []genruntime.ResourceReference{
			*tc.MakeReferenceFromResource(ag),
		},
	}
	alertRule := &alertsmanagement.SmartDetectorAlertRule{
		ObjectMeta: tc.MakeObjectMeta("smartalertrules"),
		Spec: alertsmanagement.SmartDetectorAlertRule_Spec{
			Location:     tc.AzureRegion,
			Owner:        testcommon.AsOwner(rg),
			State:        &state,
			Severity:     &severity,
			Frequency:    to.Ptr("PT1M"),
			Detector:     detectorId,
			ActionGroups: actionGroup,
			ScopeReferences: []genruntime.ResourceReference{
				*tc.MakeReferenceFromResource(acct),
			},
		},
	}
	tc.CreateResourcesAndWait(alertRule, acct)

	// Ensure that the status is what we expect
	tc.Expect(alertRule.Status.Id).ToNot(BeNil())
	armId := *alertRule.Status.Id

	tc.DeleteResourceAndWait(alertRule)

	// Ensure that the resource was really deleted in Azure
	ctx := context.Background()
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(ctx, armId, string(alertsmanagement.APIVersion_Value))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
