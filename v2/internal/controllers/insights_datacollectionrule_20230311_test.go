/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20230311"
	operationalinsights "github.com/Azure/azure-service-operator/v2/api/operationalinsights/v1api20210601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Insights_DataCollectionRule_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a Log Analytics workspace first, as DataCollectionRule needs it
	sku := operationalinsights.WorkspaceSku_Name_Standalone
	workspace := &operationalinsights.Workspace{
		ObjectMeta: tc.MakeObjectMeta("workspace"),
		Spec: operationalinsights.Workspace_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &operationalinsights.WorkspaceSku{
				Name: &sku,
			},
		},
	}

	tc.CreateResourceAndWait(workspace)

	// Create a DataCollectionRule
	dataCollectionRule := &insights.DataCollectionRule{
		ObjectMeta: tc.MakeObjectMeta("datacollectionrule"),
		Spec: insights.DataCollectionRule_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Description: to.Ptr("Test data collection rule for Azure Service Operator"),
			DataSources: &insights.DataSourcesSpec{
				PerformanceCounters: []insights.PerfCounterDataSource{
					{
						Name: to.Ptr("perfCounterDataSource"),
						CounterSpecifiers: []string{
							"\\Processor Information(_Total)\\% Processor Time",
							"\\Processor Information(_Total)\\% Privileged Time",
							"\\Processor Information(_Total)\\% User Time",
						},
						SamplingFrequencyInSeconds: to.Ptr(10),
						Streams: []insights.PerfCounterDataSource_Streams{
							insights.PerfCounterDataSource_Streams_MicrosoftPerf,
						},
					},
				},
			},
			Destinations: &insights.DestinationsSpec{
				LogAnalytics: []insights.LogAnalyticsDestination{
					{
						Name:                      to.Ptr("logAnalyticsDestination"),
						WorkspaceResourceReference: tc.MakeReferenceFromResource(workspace),
					},
				},
			},
			DataFlows: []insights.DataFlow{
				{
					Streams: []insights.DataFlow_Streams{
						insights.DataFlow_Streams_MicrosoftPerf,
					},
					Destinations: []string{
						"logAnalyticsDestination",
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(dataCollectionRule)

	tc.Expect(dataCollectionRule.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(dataCollectionRule.Status.Id).ToNot(BeNil())
	armId := *dataCollectionRule.Status.Id

	// Perform a simple patch - update the description
	old := dataCollectionRule.DeepCopy()
	dataCollectionRule.Spec.Description = to.Ptr("Updated description for test data collection rule")
	tc.PatchResourceAndWait(old, dataCollectionRule)
	tc.Expect(dataCollectionRule.Status.Description).To(Equal(to.Ptr("Updated description for test data collection rule")))

	tc.DeleteResourceAndWait(dataCollectionRule)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())

	// Clean up workspace
	tc.DeleteResourceAndWait(workspace)
}